package main

import (
	"log"
	"net/http"

	"github.com/yts1234/go-web/handler/httphandler"
	"github.com/yts1234/go-web/handler/jwthandler"
	"github.com/yts1234/go-web/handler/pasetohandler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", httphandler.HomeHandler)
	mux.HandleFunc("/helo", httphandler.HelloHandler)
	mux.HandleFunc("/product", httphandler.ProductHandler)
	mux.HandleFunc("/welcome", httphandler.WelcomeHandler)
	mux.HandleFunc("/signin", jwthandler.Signin)
	mux.HandleFunc("/pasetosignin", pasetohandler.PasetoSignIn)
	mux.HandleFunc("/welcomepaseto", httphandler.WelcomePaseto)

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting web on port 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
