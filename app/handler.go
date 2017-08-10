package app

import (
	"github.com/graphql-go/handler"
	"github.com/graphql-go/graphql"
	"fmt"
	"net/http"
	"github.com/dandoh/sdr/model"
	"encoding/json"
	"github.com/dandoh/sdr/auth"
)

func GraphqlHandlerFunc() http.HandlerFunc {
	model.Init()
	model.InitType()
	return func(w http.ResponseWriter, r *http.Request) {
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
}


func AppHandler() http.Handler {
	return auth.RequireAuth(http.HandlerFunc(GraphqlHandlerFunc()))
}

