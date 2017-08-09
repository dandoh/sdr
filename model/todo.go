package model

import (
	"github.com/jinzhu/gorm"
	"github.com/graphql-go/graphql"
)

type Todo struct {
	gorm.Model
	Content  string
	State    int
	EstimateTime int
	ActualTime int
	ReportID uint `gorm:"index"`
}

var todoType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Todo",
	Description: "...",
	Fields: graphql.Fields{
		"todoId": &graphql.Field{
			Type:        graphql.Int,
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				todo := p.Source.(Todo)
				return todo.ID, nil
			},
		},

		"content": &graphql.Field{
			Type:        graphql.String,
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				todo := p.Source.(Todo)
				return todo.Content, nil
			},
		},


		"state": &graphql.Field{
			Type:        graphql.Int,
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				todo := p.Source.(Todo)
				return todo.State, nil
			},
		},

		"estimateTime": &graphql.Field{
			Type: graphql.Int,
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				todo := p.Source.(Todo)
				return todo.EstimateTime ,nil
			},

		},

		"actualTime": &graphql.Field{
			Type: graphql.Int,
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				todo := p.Source.(Todo)
				return todo.ActualTime, nil
			},

		},


	},

})



