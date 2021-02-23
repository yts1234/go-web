package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/yts1234/go-web/entity"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("views/index.html", "views/layout.html")
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, but rest easy this maybe in our side", http.StatusInternalServerError)
		return
	}

	data := map[string]string{
		"title":   "Web GO",
		"content": "First dynamic Golang Web",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, but rest easy this maybe in our side", http.StatusInternalServerError)
		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello dari 8080"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("views/product.html", "views/layout.html")

	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, but rest easy this maybe in our side", http.StatusInternalServerError)
		return
	}
	// data := map[string]interface{}{
	// 	"title":   "Product",
	// 	"content": "",
	// }

	// data2 := entity.Product{ID: 1, Name: "Macbook Pro M1", Price: 22000000, Stock: 2}
	data2 := []entity.Product{
		{ID: 1, Name: "Macbook Pro M1", Price: 22000000, Stock: 8},
		{ID: 2, Name: "Macbook Air M1", Price: 17000000, Stock: 2},
		{ID: 3, Name: "Macbook 2012", Price: 7000000, Stock: 11},
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		// w.Write([]byte("Product Page"))
		// data["content"] = "Product Page"
		err = tmpl.Execute(w, data2)
		return
	}
	// idNumb, err := strconv.Atoi(id)
	// if err != nil || idNumb < 1 {
	// 	http.NotFound(w, r)
	// 	return
	// }
	// // w.Write([]byte("Product: ", idNumb))
	// data["content"] = data2
	// err = tmpl.Execute(w, data)
	// if err != nil {
	// 	log.Fatalf("Error: %s", err)
	// }

}


