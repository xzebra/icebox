package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	foodDB "github.com/openfoodfacts/openfoodfacts-go"
)

func handleProduct(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(filepath.Join("public", "templates", "product.html"))
		if err != nil {
			http.Error(w, err.Error(), 500)
		} else {
			api := foodDB.NewHttpApiOperator("es", "", "")
			barcode := r.URL.Query().Get("barcode")
			fmt.Printf("Requested info of product %s\n", barcode)
			product, err := api.GetProduct(barcode)

			if err != nil || barcode == "" || product.ProductName == "" {
				// handle non existing products before executing template
				http.ServeFile(w, r, filepath.Join("public", "static", "product_not_found.html"))
			} else {
				// trim product name
				extraTags := strings.Index(product.ProductName, " - ")
				if extraTags != -1 {
					product.ProductName = product.ProductName[:extraTags]
				}

				err := t.Execute(w, product)
				if err != nil {
					fmt.Printf("template error: %s\n", err.Error())
				}
			}
		}
}

func fileServerWithErrors(dir string) http.Handler {
	fileSystem := http.FileSystem(http.Dir(dir))
	fileServer := http.FileServer(fileSystem)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fileSystem.Open(path.Clean(r.URL.Path))
		if os.IsNotExist(err) {
			http.ServeFile(w, r, filepath.Join("public", "static", "404.html"))
			return
		}
		fileServer.ServeHTTP(w, r)
	})
}
