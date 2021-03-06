package httphandler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/o1egl/paseto"
	"github.com/yts1234/go-web/entity"
	"github.com/yts1234/go-web/handler/jwthandler"
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

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			log.Print("Error: User cookie is not set")
			http.Error(w, "401 Unathorized", http.StatusUnauthorized)
			// w.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Print("Error: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tknStr := c.Value

	claims := &jwthandler.Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret_key"), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Print("Error: Signature is invalid ", err)
			http.Error(w, "Signature is invalid ", http.StatusUnauthorized)
			return
		}
		log.Print("Error: ", err)
		http.Error(w, "Bad Request ", http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		log.Print("Error: Token invalid ", err)
		http.Error(w, "Token is invalid ", http.StatusUnauthorized)
		return
	}
	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
}

func WelcomePaseto(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("pasetoToken")
	if err != nil {
		if err == http.ErrNoCookie {
			log.Print("Error: User cookie is not set, ", err)
			http.Error(w, "401 Unathorized", http.StatusUnauthorized)
			// w.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Print("Error: ", err)
		http.Error(w, "Error is happening in our side, rest easy", http.StatusInternalServerError)
		return
	}
	tknStr := c.Value

	var newJsonToken paseto.JSONToken
	var newFooter string
	err = paseto.NewV2().Decrypt(tknStr, []byte("YELLOW SUBMARINE, BLACK WIZARDRY"), &newJsonToken, &newFooter)
	if err != nil {
		if err == paseto.ErrInvalidSignature {
			log.Print("Error invalid: ", err)
			http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
			return
		} else if err == paseto.ErrTokenValidationError {
			log.Print("Error invalid: ", err)
			http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
			return
		} else {
			log.Print("Error: ", err)
			http.Error(w, "400 Bad Request", http.StatusBadRequest)
			return
		}
	}
	w.Write([]byte(fmt.Sprintf("Welcome %s! from: %s", newJsonToken.Audience, newFooter)))
}
