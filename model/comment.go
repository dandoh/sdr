package model

import (
	"github.com/jinzhu/gorm"
	"github.com/graphql-go/graphql"
)

type Comment struct {
	gorm.Model
	UserID   uint `gorm:"index"`
	ReportID uint `gorm:"index"`
	Content  string `json:"content" `
}

var commentType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Comment",
	Description: "...",
	Fields: graphql.Fields{
		"commentId": &graphql.Field{
			Type:        graphql.Int,
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				comment := p.Source.(Comment)
				return comment.ID, nil
			},
		},
		//
		"content": &graphql.Field{
			Type:        graphql.String,
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				comment := p.Source.(Comment)
				return comment.Content, nil
			},
		},


	},


})

func createComment(content string, userID uint, reportID uint) int {
	comment := Comment{Content: content, UserID: userID, ReportID: reportID}
	insertComment(&comment)
	saveSubscribe(int(userID), int(reportID), comment.UpdatedAt)
	return int(comment.ID)
}
