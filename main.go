package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/dandoh/sdr/model"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/cors"
	"log"
	"github.com/dandoh/sdr/auth"
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
	graphqlHandler := http.HandlerFunc(graphqlHandlerFunc)

	// login Handler
	mux.Handle("/", http.FileServer(http.Dir("./public")))
	mux.HandleFunc("/login", auth.LoginFunc)
	mux.HandleFunc("/signup", auth.SignupFunc)
	// add in addContext middlware
	mux.Handle("/graphql", auth.RequireAuth(graphqlHandler))

	return mux
}

func setupServer() {
	rootMux := setupMux();
	c := cors.AllowAll().Handler(rootMux);
	http.ListenAndServe(":8080", c)
}
func graphqlHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// get query
	opts := handler.NewRequestOptions(r)

	// execute graphql query
	params := graphql.Params{
		Schema:         model.QLSchema, // defined in another file
		RequestString:  opts.Query,
		VariableValues: opts.Variables,
		OperationName:  opts.OperationName,
		Context:        r.Context(), // pass http.Request.Context() to our graphql object
	}
	result := graphql.Do(params)
	fmt.Printf("%+v", result)

	// output JSON
	var buff []byte
	w.WriteHeader(http.StatusOK)
	/*
		if prettyPrintGraphQL {
			buff, _ = json.MarshalIndent(result, "", "\t")
		} else {
			buff, _ = json.Marshal(result)
		}
	*/
	buff, _ = json.Marshal(result)
	w.Write(buff)
}


