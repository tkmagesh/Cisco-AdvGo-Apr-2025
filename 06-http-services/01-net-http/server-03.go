package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/agext/uuid"
)

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Cost  float64 `json:"cost"`
	Units int     `json:"units"`
}

type ProductsStore struct {
	list []Product
}

func NewProductsStore() *ProductsStore {
	return &ProductsStore{
		list: []Product{
			{Id: 100, Name: "Pen", Cost: 10, Units: 20},
			{Id: 101, Name: "Pencil", Cost: 5, Units: 50},
			{Id: 102, Name: "Marker", Cost: 50, Units: 25},
		},
	}
}

func (ps *ProductsStore) GetAll(ctx context.Context) ([]Product, error) {
	// TODO : implement logging
	log.Printf("[ProductsStore.GetAll()][%s] - returning products\n", ctx.Value("trace-id"))
	log.Printf("[ProductsStore.GetAll()][%s] - retrieving products from db\n", ctx.Value("trace-id"))
	time.Sleep(2 * time.Second) // simulating time consuming db communication
	if ctx.Err() == context.DeadlineExceeded {
		log.Printf("[ProductsStore.GetAll()][%s] - timeout from db\n", ctx.Value("trace-id"))
		return nil, errors.New("request timedout")
	}
	return ps.list[:], nil
}

func (ps *ProductsStore) Add(ctx context.Context, p Product) {
	// TODO : implement logging
	log.Printf("[ProductsStore.AddNew()][%s] - adding new product\n", ctx.Value("trace-id"))
	ps.list = append(ps.list, p)
}

type HandlerFunction func(http.ResponseWriter, *http.Request)
type MiddlewareFunction func(HandlerFunction) HandlerFunction

type AppServer struct {
	routes      map[string]HandlerFunction
	middlewares []MiddlewareFunction
}

/*
=> traceMiddleware(loggerMiddleware(timeoutMiddleware(IndexHandler)))
middlewares[ traceMiddleware, loggerMiddleware, timeoutMiddleware]
*/
func (appServer *AppServer) Add(pattern string, handlerFn HandlerFunction) {
	for i := len(appServer.middlewares) - 1; i >= 0; i-- {
		middleware := appServer.middlewares[i]
		handlerFn = middleware(handlerFn)
	}
	appServer.routes[pattern] = handlerFn
}

func (appServer *AppServer) Use(middleware MiddlewareFunction) {
	appServer.middlewares = append(appServer.middlewares, middleware)
}

func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handlerFn, exists := appServer.routes[r.URL.Path]; exists {
		handlerFn(w, r)
		return
	}
	http.Error(w, "resource not found", http.StatusNotFound)
}

func NewAppServer() *AppServer {
	return &AppServer{
		routes: make(map[string]HandlerFunction),
	}
}

// App specific  logic
// http.Handler interface implementation
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

var ps = NewProductsStore()

func ProductsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		log.Printf("[ProductsHandler][%s] - returning all the products\n", r.Context().Value("trace-id"))
		products, err := ps.GetAll(r.Context())
		if err != nil {
			http.Error(w, "request timeout", http.StatusRequestTimeout)
			return
		}
		if payload, err := json.Marshal(products); err == nil {
			fmt.Fprintln(w, string(payload))
			return
		}
		http.Error(w, "error serializing products", http.StatusInternalServerError)
	case http.MethodPost:
		log.Printf("[ProductsHandler][%s] - returning all the products\n", r.Context().Value("trace-id"))
		body, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		var newProduct Product
		if err := json.Unmarshal(body, &newProduct); err == nil {
			ps.Add(r.Context(), newProduct)
			w.WriteHeader(http.StatusCreated)
			return
		} else {
			fmt.Println("error :", err)
			http.Error(w, "error parsing payload", http.StatusBadRequest)
		}

	}
}

func CustomersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "The customers list will be served")
}

// Middlewares
func loggerMiddleware(handler HandlerFunction) HandlerFunction {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s - %s\n", r.Context().Value("trace-id"), r.Method, r.URL.Path)
		handler(w, r)
	}
}

func traceMiddleware(handler HandlerFunction) HandlerFunction {
	return func(w http.ResponseWriter, r *http.Request) {
		uuidObj := uuid.New()
		traceCtx := context.WithValue(r.Context(), "trace-id", uuidObj.String())
		handler(w, r.WithContext(traceCtx))
	}
}

func timeoutMiddleware(handler HandlerFunction) HandlerFunction {
	return func(w http.ResponseWriter, r *http.Request) {
		timeoutCtx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()
		handler(w, r.WithContext(timeoutCtx))
		if r.Context().Err() == context.DeadlineExceeded {
			log.Printf("%s - timedout\n", r.Context().Value("trace-id"))
		}
	}
}

func main() {

	appServer := NewAppServer()
	appServer.Use(traceMiddleware)
	appServer.Use(loggerMiddleware)
	appServer.Use(timeoutMiddleware)

	appServer.Add("/", IndexHandler)
	appServer.Add("/products", ProductsHandler)
	appServer.Add("/customers", CustomersHandler)
	if err := http.ListenAndServe(":8080", appServer); err != nil {
		log.Println("Error :", err)
	}
}
