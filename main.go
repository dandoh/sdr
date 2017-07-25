package main

import (
	_"github.com/jinzhu/gorm/dialects/postgres"
	"sdr/model"
	"github.com/graphql-go/handler"
	"net/http"
	_"fmt"
	//"fmt"
)





func main() {
	// initialize database
	model.Init();
	model.InitType();

	h := handler.New(&handler.Config{
		Schema: &model.SchemaQL,
		Pretty: true,
	})

	http.Handle("/graphql", h)
/*
	http.HandleFunc("/graphql", 	func(w http.ResponseWriter, r *http.Request){
		var token string = r.URL.Query()["query"][0]
		fmt.Print("this is token:", token)
		if(token == "vuthede") {
			http.Handle("/graphql", h)
		}
	})
	*/
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.ListenAndServe(":8080", nil)

	defer model.Close();
}
