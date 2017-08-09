package model

import (
	"errors"
	_"time"
	"github.com/graphql-go/graphql"
	_"fmt"
	_"container/list"
	_"github.com/labstack/gommon/email"
	"fmt"

	_"github.com/jinzhu/gorm"


)


//var TodoInputType = graphql.NewInputObject(
//	graphql.InputObjectConfig{
//		Name: "TodoInputType",
//		Fields: graphql.InputObjectConfigFieldMap{
//			"content": &graphql.InputObjectFieldConfig{
//				Type: graphql.String,
//			},
//			"state": &graphql.InputObjectFieldConfig{
//				Type: graphql.Int,
//			},
//			"estimateTime": &graphql.InputObjectFieldConfig{
//				Type: graphql.Int,
//			},
//			"spentTime": &graphql.InputObjectFieldConfig{
//				Type: graphql.Int,
//			},
//		},
//	},
//)


var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"group": &graphql.Field{
			Type: groupType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int     ,
					Description: "...",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				idQuery, isOK := p.Args["id"].(int)
				if isOK {
					return findGroupByID(idQuery), nil
				}
				return Group{}, nil
			},
		},

		"report": &graphql.Field{
			Type: reportType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "...",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				idQuery, isOK := p.Args["id"].(int)
				if isOK {
					return findReportByID(uint(idQuery)), nil
				}

				return Report{}, nil
			},

		},


		"groups": &graphql.Field{
			Type: graphql.NewList(groupType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				authorContext := p.Context.Value("authorContext").(AuthorContext)
				fmt.Printf("%+v", authorContext)
				return findGroupsByUserID(int(authorContext.AuthorID)), nil

			},

		},

		"usersOfGroup": &graphql.Field{
			Type: graphql.NewList(userType),
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "...",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				idQuery, isOK := p.Args["id"].(int)
				if isOK {
					return findUsersByGroupID(idQuery), nil
				}

				return User{}, nil
			},

		},

		"note": &graphql.Field{
			Type: graphql.String,

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				authorContext := p.Context.Value("authorContext").(AuthorContext)
				return findUserByID(int(authorContext.AuthorID)).Note, nil
			},
		},

	},
})

var mutateType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutex",
	Fields: graphql.Fields{
		"addGroup": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				name := p.Args["name"].(string)
				authorContext := p.Context.Value("authorContext").(AuthorContext)
				if !isNameGroupExisted(name) {
					insertGroup(name)
					insertUserToGroupByID(int(authorContext.AuthorID), name);
					return true, nil
				} else {
					return false, errors.New("Group name existed")
				}
			},
		},

		"addUserToGroup": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"email": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},

				"groupId": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				email := p.Args["email"].(string)
				groupId := p.Args["groupId"].(int)
				insertUserToGroupByEmail(email, groupId)
				//fmt.Print(emails)
				return true, nil

			},

		},


		"deleteUserInGroup": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"userEmail": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},

				"groupId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				email := p.Args["userEmail"].(string)
				groupId := p.Args["groupId"].(int)

				//deleteUserInGroup haven't been implemented yet!!!
				deleteUserInGroupByEmail(email, groupId)
				return true, nil
			},


		},


		"addComment": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"content": &graphql.ArgumentConfig{
					Type: graphql.String,
				},

				"reportId": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				content := p.Args["content"].(string)
				reportId := p.Args["reportId"].(int)
				authorContext := p.Context.Value("authorContext").(AuthorContext)

				return createComment(content, uint(authorContext.AuthorID), uint(reportId)), nil
			},
		},


	},
})

var QLSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    queryType,
	Mutation: mutateType,
})

// this is for cyclic dependencies
func InitType() {
	groupType.AddFieldConfig("users",
		&graphql.Field{Type: graphql.NewList(userType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				group := p.Source.(Group)
				return findUsersInGroup(group), nil
			}, })

	reportType.AddFieldConfig("user",
		&graphql.Field{Type: userType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				report := p.Source.(Report)
				return findUserByID(int(report.UserID)), nil
			},


		}, )

	commentType.AddFieldConfig("user", &graphql.Field{
		Type:        userType,
		Description: "...",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			comment := p.Source.(Comment)
			return findUserByID(int(comment.UserID)), nil
		},
	});
	commentType.AddFieldConfig("report", &graphql.Field{
		Type:        reportType,
		Description: "...",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			comment := p.Source.(Comment)
			return findReportByID(comment.ReportID), nil
		},
	})
}
