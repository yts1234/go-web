package main

import "fmt"

// Main func
// func main() {
// 	mux := http.NewServeMux()

// 	mux.HandleFunc("/", httphandler.HomeHandler)
// 	mux.HandleFunc("/helo", httphandler.HelloHandler)
// 	mux.HandleFunc("/product", httphandler.ProductHandler)
// 	mux.HandleFunc("/welcome", httphandler.WelcomeHandler)
// 	mux.HandleFunc("/signin", jwthandler.Signin)
// 	mux.HandleFunc("/pasetosignin", pasetohandler.PasetoSignIn)
// 	mux.HandleFunc("/welcomepaseto", httphandler.WelcomePaseto)

// 	fileServer := http.FileServer(http.Dir("assets"))
// 	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

// 	log.Println("Starting web on port 8080")

// 	err := http.ListenAndServe(":8080", mux)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// Test function for exercise
const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
