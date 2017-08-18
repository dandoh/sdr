package model

import (
	_"github.com/dandoh/sdr/util"
	"github.com/jinzhu/gorm"
	"github.com/graphql-go/graphql"
	_"fmt"


)

type Subscribe struct {
	gorm.Model
	UserId       uint
	ReportId uint
	NumberCommentsNotSeen int
}


var subscribeType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Subscribe",
	Description: "...",
	Fields: graphql.Fields{
		"userId": &graphql.Field{
			Type:        graphql.Int,
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				subscribe := p.Source.(Subscribe)
				return subscribe.UserId, nil
			},
		},

		"reportId": &graphql.Field{
			Type:        graphql.Int,
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				subscribe := p.Source.(Subscribe)
				return subscribe.ReportId, nil
			},
		},

		"numberCommentsNotSeen": &graphql.Field{
			Type:        graphql.Int,
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				subscribe := p.Source.(Subscribe)
				return subscribe.NumberCommentsNotSeen, nil
			},
		},



		//Password and Token haven't been declared.
	},
})

func getNumCommentsNotSeenInSubscribe(userId int, reportId int) int{
	var num int
	subscribe := findSubscribeByIds(userId, reportId)
	db.Model(&Comment{}).Where("report_id = ? AND updated_at > ?",
		reportId, subscribe.UpdatedAt).Count(&num)

	subscribe.NumberCommentsNotSeen = num
	db.Save(&subscribe)
	return subscribe.NumberCommentsNotSeen
}

