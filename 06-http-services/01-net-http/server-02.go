package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Cost  float64 `json:"cost"`
	Units int     `json:"units"`
}

var products = []Product{
	{Id: 100, Name: "Pen", Cost: 10, Units: 20},
	{Id: 101, Name: "Pencil", Cost: 5, Units: 50},
	{Id: 102, Name: "Marker", Cost: 50, Units: 25},
}

type AppServer struct {
	routes map[string]func(http.ResponseWriter, *http.Request)
}

func (appServer *AppServer) Add(pattern string, handlerFn func(http.ResponseWriter, *http.Request)) {
	appServer.routes[pattern] = handlerFn
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
		routes: make(map[string]func(http.ResponseWriter, *http.Request)),
	}
}

// App specific  logic
// http.Handler interface implementation
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if payload, err := json.Marshal(products); err == nil {
			fmt.Fprintln(w, string(payload))
			return
		}
		http.Error(w, "error serializing products", http.StatusInternalServerError)
	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		var newProduct Product
		if err := json.Unmarshal(body, &newProduct); err == nil {
			products = append(products, newProduct)
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

func main() {
	appServer := NewAppServer()
	appServer.Add("/", IndexHandler)
	appServer.Add("/products", ProductsHandler)
	appServer.Add("/customers", CustomersHandler)
	if err := http.ListenAndServe(":8080", appServer); err != nil {
		log.Println("Error :", err)
	}
}
