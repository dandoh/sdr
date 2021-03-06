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
	"time"
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
		"users": &graphql.Field{
			Type: graphql.NewList(userType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return findAllUsers(), nil
			},
		},



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

		"report": &graphql.Field{
			Type: reportType,
			Args: graphql.FieldConfigArgument{
				"reportId": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "...",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				authorContext := p.Context.Value("authorContext").(AuthorContext)
				idQuery, isOK := p.Args["reportId"].(int)
				if isOK {
					saveSubscribe(int(authorContext.AuthorID), idQuery, time.Now())
					return findReportByID(uint(idQuery)), nil
				}

				return Report{}, nil
			},

		},
		"user": &graphql.Field{
			Type: userType,
			Args: graphql.FieldConfigArgument{
				"userId": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				idQuery, isOK := p.Args["userId"].(int)
				if isOK {
					return findUserByID(idQuery), nil
				}

				return User{}, nil
			},
		},

		"reports": &graphql.Field{
			Type: graphql.NewList(reportType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				authorContext := p.Context.Value("authorContext").(AuthorContext)
				return getAllReportsOfUser(int(authorContext.AuthorID)), nil

			},

		},

		"oldReports": &graphql.Field{
			Type: graphql.NewList(reportType),
			Args: graphql.FieldConfigArgument{
				"fromDate": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "...",
				},

				"toDate": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "...",
				},

				"userId": &graphql.ArgumentConfig{
					Type: graphql.Int,
					Description: "...",
				},

			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				fromDate := p.Args["fromDate"].(string)
				toDate := p.Args["toDate"].(string)
				userId := p.Args["userId"].(int)
				reports, err := getOldReportsByUserId(userId, fromDate, toDate)

				if err != nil {
					return nil, errors.New("Date must be in form yyyy-mm-dd")
				}

				return reports, nil

			},

		},

		"reportToday": &graphql.Field{
			Type: reportType,

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				authorContext := p.Context.Value("authorContext").(AuthorContext)
				report := findReportTodayByUserId(int(authorContext.AuthorID))
				saveSubscribe(int(authorContext.AuthorID), int(report.ID), time.Now())
				return report, nil
				return Report{}, nil
			},

		},

		"reportsTodayOfGroup": &graphql.Field{
			Type: graphql.NewList(reportType),
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "...",
				},
			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				groupId := p.Args["groupId"].(int)
				return getAllReportsTodayByGroupId(int(groupId)), nil
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
				"groupId": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "...",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				idQuery, isOK := p.Args["groupId"].(int)
				if isOK {
					return findUsersByGroupID(idQuery), nil
				}

				return nil, nil
			},

		},

		"subscribes": &graphql.Field{
			Type: graphql.NewList(subscribeType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				authorContext := p.Context.Value("authorContext").(AuthorContext)
				return getAllSubscribesOfUser(int(authorContext.AuthorID)), nil
			},

		},


	},
})

var mutateType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutex",
	Fields: graphql.Fields{
		"addGroup": &graphql.Field{
			Type: graphql.Int,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"purpose": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},

			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				name := p.Args["name"].(string)
				purpose := p.Args["purpose"].(string)
				authorContext := p.Context.Value("authorContext").(AuthorContext)
				if !isNameGroupExisted(name) {
					groupID := insertGroup(name, purpose)
					insertUserToGroupByID(int(authorContext.AuthorID), int(groupID));
					fmt.Println(groupID)
					return groupID, nil
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
					Type: graphql.NewNonNull(graphql.Int),
				},
			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				email := p.Args["email"].(string)
				groupId := p.Args["groupId"].(int)
				if !isEmailExisted(email) {
					return false, errors.New("Email is not existed in database")
				}
				if !isUserInGroupAlready(email, groupId) {
					insertUserToGroupByEmail(email, groupId)
					return true, nil
				} else {
					return false, errors.New("User has joined in group already")
				}

			},

		},

		"addUsersToGroup": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"emails": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.NewNonNull(graphql.String)),
				},

				"groupId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				emailsArg := p.Args["emails"].([]interface{})

				groupId := p.Args["groupId"].(int)

				numNewUser := 0
				for _, email := range emailsArg {
					if (!isUserInGroupAlready(email.(string), groupId) && isEmailExisted(email.(string))) {
						insertUserToGroupByEmail(email.(string), groupId)
						numNewUser ++
					}
				}

				if (numNewUser == len(emailsArg)) {
					return true, nil
				}

				if (numNewUser != 0) {
					return false, errors.New("There is one or more unvalid emails")
				}

				return false, errors.New("All emails have existed")

			},

		},

		"changeGroupInfo": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},

				"groupName": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},

				"purpose": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},

				"emails": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.NewNonNull(graphql.String)),
				},
			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				groupId := p.Args["groupId"].(int)
				groupName := p.Args["groupName"].(string)
				purpose := p.Args["purpose"].(string)
				emailsArg := p.Args["emails"].([]interface{})

				emails := make([]string, len(emailsArg))
				for i,_ := range emails{
					emails[i] = emailsArg[i].(string)
				}

				success := updateGroupInfo(groupId, groupName, purpose, emails)
				if (success){
					return true, nil
				}
				return false, errors.New("At least a user is invalid. However, valid users were inserted into dbs")
			},


		},

		"deleteUserInGroup": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"email": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},

				"groupId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				email := p.Args["email"].(string)
				groupId := p.Args["groupId"].(int)
				if (isUserInGroupAlready(email, groupId)) {
					deleteUserInGroupByEmail(email, groupId)
					return true, nil
				}
				return false, errors.New("This group doesn't have email like that.")

			},


		},

		"addTodo": &graphql.Field{
			Type: graphql.Int,
			Args: graphql.FieldConfigArgument{
				"content": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"state": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"estimateTime": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"spentTime": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"reportId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				content := p.Args["content"].(string)
				state := p.Args["state"].(int)
				estimateTime := p.Args["estimateTime"].(int)
				spentTime := p.Args["spentTime"].(int)
				reportId := p.Args["reportId"].(int)
				//remember to check if reportId is existed or not... haven't implement yet
				return addTodo(content, state, estimateTime, spentTime, reportId), nil
			},
		},

		"deleteTodo": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"todoId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},


			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				todoId := p.Args["todoId"].(int)
				return deleteTodo(todoId), nil

			},
		},

		"updateTodo": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"todoId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"content": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"state": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"estimateTime": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"spentTime": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},

			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				todoId := p.Args["todoId"].(int)
				content := p.Args["content"].(string)
				state := p.Args["state"].(int)
				estimateTime := p.Args["estimateTime"].(int)
				spentTime := p.Args["spentTime"].(int)
				return updateTodo(todoId, content, state, estimateTime, spentTime), nil

			},
		},

		"addComment": &graphql.Field{
			Type: graphql.Int,
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
				//remember to check if reportId have existed or not.
				return createComment(content, uint(authorContext.AuthorID), uint(reportId)), nil
			},
		},

		"updateNote": &graphql.Field{
			Type: graphql.String,
			Args: graphql.FieldConfigArgument{
				"note": &graphql.ArgumentConfig{
					Type: graphql.String,
				},

				"reportId": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				note := p.Args["note"].(string)
				reportId := p.Args["reportId"].(int)
				return updateNoteOfReport(note, reportId), nil
			},
		},

		"saveSubscribe": &graphql.Field{
			Type: graphql.String,
			Args: graphql.FieldConfigArgument{
				"reportId": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				reportId := p.Args["reportId"].(int)
				authorContext := p.Context.Value("authorContext").(AuthorContext)
				return saveSubscribe(int(authorContext.AuthorID), reportId, time.Now()), nil
			},
		},


	},
})

var QLSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    queryType,
	Mutation: mutateType,
})

// for cyclic dependencies
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

	/*
	subscribeType.AddFieldConfig("user",
		&graphql.Field{Type: userType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				subscribe := p.Source.(Subscribe)
				return findUserByID(int(subscribe.UserId)), nil
			},
		}, )

	subscribeType.AddFieldConfig("report",
		&graphql.Field{Type: reportType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				subscribe := p.Source.(Subscribe)
				return findReportByID(subscribe.ReportId), nil
			},

		}, )

		*/

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
