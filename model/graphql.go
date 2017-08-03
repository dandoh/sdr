package model

import (
	"errors"
	_"time"
	"github.com/graphql-go/graphql"
	_"fmt"
	_"container/list"
	_"github.com/labstack/gommon/email"
	"fmt"
)

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"group": &graphql.Field{
			Type: groupType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.Int,
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

		"reportsOfGroup": &graphql.Field{
			Type: graphql.NewList(reportType),
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "...",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				idQuery, isOK := p.Args["id"].(int)
				if isOK {
					return findReportsByGroupID(idQuery), nil
				}

				return []Report{}, nil
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

		"addUserByEmail": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"email": &graphql.ArgumentConfig{
					Type: graphql.String,
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

		"updateNote": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"note": &graphql.ArgumentConfig{
					Type: graphql.String,
				},

			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				note := p.Args["note"].(string)
				authorContext := p.Context.Value("authorContext").(AuthorContext)
				updateUserNote(note, int(authorContext.AuthorID))
				return true, nil
			},
		},

		"deleteUserInGroup": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"userEmail": &graphql.ArgumentConfig{
					Type: graphql.String,
				},

				"groupName": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				email := p.Args["userEmail"].(string)
				groupName := p.Args["groupName"].(string)

				//deleteUserInGroup haven't been implemented yet!!!
				deleteUserInGroupByEmail(email, groupName)
				return true, nil
			},


		},
		"createReport": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"contentTodoes": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},

				"states": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.Int),
				},


				"groupId": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				contentTodoesArgs := p.Args["contentTodoes"].([]interface{})
				statesArgs := p.Args["states"].([]interface{})
				contentTodoes := make([]string, len(contentTodoesArgs))
				states := make([]int, len(statesArgs))
				for i := range contentTodoes {
					contentTodoes[i] = contentTodoesArgs[i].(string)
					fmt.Println(contentTodoes[i]);
				}

				for i := range states {
					states[i] = statesArgs[i].(int)
					fmt.Println(states[i]);
				}

				groupId := p.Args["groupId"].(int)
				authorContext := p.Context.Value("authorContext").(AuthorContext)
				createReport(contentTodoes, states, int(authorContext.AuthorID), groupId)
				return true, nil

			},
		},

		"updateReport": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"reportId": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"contentTodoes": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},

				"states": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.Int),
				},

				"summary": &graphql.ArgumentConfig{
					Type: graphql.String,
				},

				"status": &graphql.ArgumentConfig{
					Type: graphql.String,
				},

			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				contentTodoesArgs := p.Args["contentTodoes"].([]interface{})
				statesArgs := p.Args["states"].([]interface{})
				contentTodoes := make([]string, len(contentTodoesArgs))
				states := make([]int, len(statesArgs))
				for i := range contentTodoes {
					contentTodoes[i] = contentTodoesArgs[i].(string)
					fmt.Println(contentTodoes[i]);
				}

				for i := range states {
					states[i] = statesArgs[i].(int)
					fmt.Println(states[i]);
				}

				reportId := p.Args["reportId"].(int)
				summary := ""
				if p.Args["summary"] != nil {
					summary = p.Args["summary"].(string)
				}
				//summary := p.Args["summary"](string)
				status := p.Args["status"].(string)

				//authorContext := p.Context.Value("authorContext").(AuthorContext)
				updateReport(reportId, contentTodoes, states, summary, status)
				return true, nil

			},

		},


		"createComment": &graphql.Field{
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

	reportType.AddFieldConfig("group",
		&graphql.Field{Type: groupType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				report := p.Source.(Report)
				return findGroupByID(int(report.GroupID)), nil
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
