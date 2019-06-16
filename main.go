package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	foodDB "github.com/openfoodfacts/openfoodfacts-go"
)

func main() {
	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(filepath.Join("public", "templates", "product.html"))
		if err != nil {
			http.Error(w, err.Error(), 500)
		} else {
			api := foodDB.NewHttpApiOperator("es", "", "")
			fmt.Printf("Requested info of product %s\n", r.URL.Query().Get("barcode"))
			product, err := api.GetProduct(r.URL.Query().Get("barcode"))
			if err != nil {
				http.Error(w, "could not connect to the server", 500)
			} else {
				fmt.Println("INFO:", product.Id, product.GenericName)
				err := t.Execute(w, product)
				if err != nil {
					http.Error(w, "could not get product info", 404)
				}
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("public", "static", "index.html"))
	})

	fmt.Println("Running server at port 8080")
	err := http.ListenAndServe("192.168.100.17:8080", nil)
	if err != nil {
		panic(err)
	}
}