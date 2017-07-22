package model

import (
	"github.com/jinzhu/gorm"
	"time"
	"github.com/graphql-go/graphql"
	"fmt"
)

type User struct {
	gorm.Model
	Name        string `gorm:"size:255"`
	PasswordMD5 string `gorm:"size:255"`
	Email       string `gorm:"not null"`
	Token       string
	Note        string    `gorm:"size:2000"`
	Groups      []Group `gorm:"many2many:user_group"`
	Reports     []Report
	Comments    []Comment
}

type Comment struct {
	gorm.Model
	UserID   uint `gorm:"index"`
	ReportID uint `gorm:"index"`
	Date     time.Time `json:"date" json:"date"`
	Content  string `json:"content" `
}

type Group struct {
	gorm.Model
	Name  string `gorm:"size:255" json:"name"`
	Users []User    `gorm:"many2many:user_group" json:"users"`
}

type Report struct {
	gorm.Model
	Date          time.Time
	Summerization string `gorm:"size:1000"`
	UserID        uint `gorm:"index"`
	Todoes        []Todo
	Comments      []Comment
}

type Todo struct {
	gorm.Model
	Content  string
	Status   uint
	ReportID uint `gorm:"index"`
}

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "User",
	Description: "...",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.Int,
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user := p.Source.(User)
				return user.ID, nil
			},
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "...",
		},
		"email": &graphql.Field{
			Type:        graphql.String,
			Description: "...",
		},
		//"groups": &graphql.Field{
		//	Type:        graphql.NewList(groupType),
		//	Description: "Which posts they have written.",
		//	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		//
		//	},
		//},
	},
})
var groupType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Group",
	Description: "Group of users",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.Int,
			Description: "The group's id",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				group := p.Source.(Group)
				return group.ID, nil
			},
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "The group's name",
		},
		"users": &graphql.Field{
			Type:        graphql.NewList(userType),
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var users []User
				group := p.Source.(Group)
				db.Model(&group).Association("Users").Find(&users)

				return users, nil
			},
		},
	},
})
var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"getGroup": &graphql.Field{
			Type: groupType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "...",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				fmt.Println("hehehe");
				var group Group
				idQuery, isOK := p.Args["id"].(int)
				if isOK {
					db.First(&group, idQuery)
					return group, nil
				}

				return Group{}, nil
			},
		},
	},
})

var SchemaQL, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
})
