package model

import (
	_"github.com/dandoh/sdr/util"
	"github.com/jinzhu/gorm"
	"github.com/graphql-go/graphql"
	_"fmt"

	"time"

)

type Subscribe struct {
	gorm.Model
	UserId       uint `gorm:"index"`
	ReportId 	 uint `gorm:"index"`
	LastUpdatedAt time.Time
	NumberCommentsNotSeen int
}


var subscribeType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Subscribe",
	Description: "...",
	Fields: graphql.Fields{
		/*
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
		*/

		"userCommentLast":&graphql.Field{
			Type: userType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				subscribe := p.Source.(Subscribe)
				user, isExist := getUserCommentLastInReport(int(subscribe.ReportId))
				if (isExist){
					return user, nil
				}
				return User{}, nil
			},
		},

		"report" : &graphql.Field{
			Type: reportType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				subscribe := p.Source.(Subscribe)
				return findReportByID(subscribe.ReportId), nil
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

		"lastUpdatedAt" : &graphql.Field{
			Type:        graphql.String,
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				subscribe := p.Source.(Subscribe)
				return subscribe.LastUpdatedAt.String(), nil
			},
		},



		//Password and Token haven't been declared.
	},
})

func getUpdatedSubscribe(userId int, reportId int) Subscribe{
	var num int
	subscribe := findSubscribeByIds(userId, reportId)
	db.Model(&Comment{}).Where("report_id = ? AND updated_at > ?",
		reportId, subscribe.LastUpdatedAt).Count(&num)

	subscribe.NumberCommentsNotSeen = num
	db.Save(&subscribe)
	return subscribe
}

