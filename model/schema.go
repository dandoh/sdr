package model

import (
	"github.com/jinzhu/gorm"
	_"time"
	"github.com/graphql-go/graphql"
	_"fmt"
	_"container/list"
	_"github.com/labstack/gommon/email"
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

type Group struct {
	gorm.Model
	Name    string `gorm:"size:255" json:"name"`
	Users   []User    `gorm:"many2many:user_group" json:"users"`
	Reports []Report
}

type Report struct {
	gorm.Model
	Summerization string `gorm:"size:1000"`
	UserID        uint `gorm:"index"`
	Status        string
	Todoes        []Todo
	Comments      []Comment
	GroupID       uint `gorm:"index"`
}

type Comment struct {
	gorm.Model
	UserID   uint `gorm:"index"`
	ReportID uint `gorm:"index"`
	Content  string `json:"content" `
}

type Todo struct {
	gorm.Model
	Content  string
	State    bool
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
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user := p.Source.(User)
				return user.Name, nil
			},
		},
		"email": &graphql.Field{
			Type:        graphql.String,
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user := p.Source.(User)
				return user.Email, nil
			},

		},

		"note": &graphql.Field{
			Type:        graphql.String,
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user := p.Source.(User)
				return user.Note, nil
			},
		},


		"groups": &graphql.Field{
			Type:        graphql.NewList(groupType),
			Description: "Which posts they have written.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user := p.Source.(User)
				return getGroupsContainUser(user), nil
			},
		},

		"reports": &graphql.Field{
			Type:        graphql.NewList(reportType),
			Description: "Which posts they have written.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user := p.Source.(User)
				return getReportsOfUser(user), nil
			},
		},

		"comments": &graphql.Field{
			Type:        graphql.NewList(commentType),
			Description: "Which posts they have written.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user := p.Source.(User)
				return getCommentsOfUser(user), nil
			},
		},
		//Password and Token haven't been declared.
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
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				group := p.Source.(Group)
				return group.Name, nil
			},
		},

		"reports": &graphql.Field{
			Type:        graphql.NewList(reportType),
			Description: "Which posts they have written.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				group := p.Source.(Group)
				return getReportsByGroupId(int(group.ID)), nil
			},
		},

	},
})

var reportType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Report",
	Description: "...",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.Int,
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				report := p.Source.(Report)
				return report.ID, nil
			},
		},


		"status": &graphql.Field{
			Type:        graphql.String,
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				report := p.Source.(Report)
				return report.Status, nil
			},
		},

		"summerization": &graphql.Field{
			Type:        graphql.String,
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				report := p.Source.(Report)
				return report.Summerization, nil
			},
		},


		"todoes": &graphql.Field{
			Type:        graphql.NewList(todoType),
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				report := p.Source.(Report)
				return getTodoesOfReport(report), nil
			},
		},

		"comments": &graphql.Field{
			Type:        graphql.NewList(commentType),
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				report := p.Source.(Report)
				return getCommentsOfReport(report), nil
			},
		},

	},


})

var commentType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Comment",
	Description: "...",
	Fields: graphql.Fields{
		"id": &graphql.Field{
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

var todoType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Todo",
	Description: "...",
	Fields: graphql.Fields{
		"id": &graphql.Field{
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
			Type:        graphql.Boolean,
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				todo := p.Source.(Todo)
				return todo.State, nil
			},
		},


	},

})

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"getGroupById": &graphql.Field{
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
					return getGroupById(idQuery), nil
				}
				return Group{}, nil
			},
		},

		"getReportsByGroupId": &graphql.Field{
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
					return getReportsByGroupId(idQuery), nil
				}

				return Group{}, nil
			},

		},

		"getReportsByGroupName": &graphql.Field{
			Type: graphql.NewList(reportType),
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "...",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				idQuery, isOK := p.Args["name"].(string)
				if isOK {
					return getReportsByGroupName(idQuery), nil
				}

				return Group{}, nil
			},

		},

		"getGroupsByUserId": &graphql.Field{
			Type: graphql.NewList(groupType),
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "...",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				idQuery, isOK := p.Args["id"].(int)
				if isOK {
					return getGroupsByUserId(idQuery), nil
				}

				return Group{}, nil
			},

		},

		"getUsersByGroupId": &graphql.Field{
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
					return getUsersByGroupId(idQuery), nil
				}

				return User{}, nil
			},

		},

		"getNoteByUserId": &graphql.Field{
			Type: graphql.String,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "...",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				idQuery, isOK := p.Args["id"].(int)
				if isOK {
					return getUserById(idQuery).Note, nil
				}

				return User{}, nil
			},
		},

	},
})

var mutateType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutex",
	Fields: graphql.Fields{
		"createAccount": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"email": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				name := p.Args["name"].(string)
				password := p.Args["password"].(string)
				email := p.Args["email"].(string)
				if !isEmailExisted(email) {
					return createAccount(name, password, email), nil
				}
				return false, nil

			},

		},
		"createGroup": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				name := p.Args["name"].(string)
				if !isNameGroupExisted(name) {
					return createGroup(name), nil
				}
				return false, nil

			},
		},
		"addUsersToGroup": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"emails": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},

				"groupName": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				emailsArgs := p.Args["emails"].([]interface{})
				emails := make([]string, len(emailsArgs))

				for i := range emails {
					emails[i] = emailsArgs[i].(string)
				}

				groupName := p.Args["groupName"].(string)
				for _, email := range emails {
					addUserToGroup(email, groupName)
				}

				//fmt.Print(emails)
				return true, nil

			},

		},
		"updateNoteForUser": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"note": &graphql.ArgumentConfig{
					Type: graphql.String,
				},

				"userId": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				note := p.Args["note"].(string)
				userId := p.Args["userId"].(int)
				return updateNoteForUser(note, userId), nil

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
				name := p.Args["groupName"].(string)

				//deleteUserInGroup haven't been implemented yet!!!
				return deleteUserInGroup(email, name), nil

			},


		},
		"createReport": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"contentTodoes": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},

				"states": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.Boolean),
				},

				"summerization": &graphql.ArgumentConfig{
					Type: graphql.String,
				},


				"userEmail": &graphql.ArgumentConfig{
					Type: graphql.String,
				},

				"groupName": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				contentTodoesArgs := p.Args["contentTodoes"].([]interface{})
				statesArgs := p.Args["states"].([]interface{})
				contentTodoes := make([]string, len(contentTodoesArgs))
				states := make([]bool, len(statesArgs))

				for i := range contentTodoes {
					contentTodoes[i] = contentTodoesArgs[i].(string)
				}

				for i := range states {
					states[i] = statesArgs[i].(bool)
				}

				userEmail := p.Args["userEmail"].(string)
				nameGroup := p.Args["groupName"].(string)
				summerization := p.Args["summerization"].(string)

				return CreateReport(contentTodoes, states, summerization, userEmail, nameGroup), nil
			},
		},
		"createComment": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"content": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"userId" : &graphql.ArgumentConfig{
					Type : graphql.Int,
				},
				"reportId" :&graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				content := p.Args["content"].(string)
				userId := p.Args["userId"].(int)
				reportId := p.Args["reportId"].(int)

				return CreateComment(content, uint(userId), uint(reportId)), nil
			},
		},


	},
})

var SchemaQL, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    queryType,
	Mutation: mutateType,
})

func InitType() {
	groupType.AddFieldConfig("users",
		&graphql.Field{Type: graphql.NewList(userType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				group := p.Source.(Group)
				return getUsersInGroup(group), nil
			}, })

	// TODO
	reportType.AddFieldConfig("user",
		&graphql.Field{Type: userType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				report := p.Source.(Report)
				return getUserById(int(report.UserID)), nil
			},


		}, )

	reportType.AddFieldConfig("group",
		&graphql.Field{Type: groupType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				report := p.Source.(Report)
				return getGroupById(int(report.GroupID)), nil
			},


		}, )

	commentType.AddFieldConfig("user", &graphql.Field{
		Type:       reportType,
		Description: "...",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			comment := p.Source.(Comment)
			return getReportById(comment.ReportID), nil
		},
	});

	commentType.AddFieldConfig("report", &graphql.Field{
		Type:        reportType,
		Description: "...",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			comment := p.Source.(Comment)
			return getReportById(comment.ReportID), nil
		},
	})
}
