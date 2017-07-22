package main

import (
	_"github.com/jinzhu/gorm/dialects/postgres"
	"sdr/model"
	"github.com/graphql-go/handler"
	"net/http"
)

func main() {
	// initialize database
	model.Init();
	h := handler.New(&handler.Config{
		Schema: &model.SchemaQL,
		Pretty: true,
	})

	http.Handle("/graphql", h)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.ListenAndServe(":8080", nil)

	defer model.Close();
}
