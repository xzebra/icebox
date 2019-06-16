package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	foodDB "github.com/openfoodfacts/openfoodfacts-go"
)

func handleProduct(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(filepath.Join("public", "templates", "product.html"))
		if err != nil {
			http.Error(w, err.Error(), 500)
		} else {
			api := foodDB.NewHttpApiOperator("es", "", "")
			fmt.Printf("Requested info of product %s\n", r.URL.Query().Get("barcode"))
			product, err := api.GetProduct(r.URL.Query().Get("barcode"))
			if err != nil {
				http.ServeFile(w, r, filepath.Join("public/static/product_not_found.html"))
			} else {
				fmt.Println("INFO:", product.GenericName)
				err := t.Execute(w, product)
				if err != nil {
					http.ServeFile(w, r, filepath.Join("public/static/404.html"))
				}
			}
		}
	}
