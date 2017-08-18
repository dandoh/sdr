package main

import (
	"net/http"
	"github.com/joho/godotenv"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/cors"
	"log"
	"github.com/dandoh/sdr/auth"
	"github.com/dandoh/sdr/app"
	"github.com/dandoh/sdr/model"
	"fmt"

)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	model.Init()
	model.InitType()
	setupServer()

}

func setupMux() *http.ServeMux {
	mux := http.NewServeMux()

	// graphql Handler
	appHandler := app.AppHandler();

	// login Handler
	mux.Handle("/", http.FileServer(http.Dir("./public")))
	mux.HandleFunc("/signin", auth.LoginFunc)
	mux.HandleFunc("/signup", auth.SignupFunc)

	// add in addContext middlware
	mux.Handle("/graphql", appHandler)

	return mux
}

func setupServer() {
	rootMux := setupMux();
	c := cors.AllowAll().Handler(rootMux);
	fmt.Println(http.ListenAndServe(":8080", c))
}
