package main

import (
	_"github.com/jinzhu/gorm/dialects/postgres"
	"sdr/model"
	"github.com/graphql-go/handler"
	"net/http"
	_"fmt"
	"github.com/rs/cors"
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
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	c := cors.Default().Handler(h);
	http.ListenAndServe(":8080", c)
	defer model.Close();
}
