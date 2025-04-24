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
}

// http.Handler interface implementation
func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s - %s\n", r.Method, r.URL.Path)
	switch r.URL.Path {
	case "/":
		fmt.Fprintln(w, "Hello, World!")
	case "/products":
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

	case "/customers":
		fmt.Fprintln(w, "The customers list will be served")
	default:
		http.Error(w, "resource not found", http.StatusNotFound)
	}

}

func main() {
	appServer := &AppServer{}
	if err := http.ListenAndServe(":8080", appServer); err != nil {
		log.Println("Error :", err)
	}
}
