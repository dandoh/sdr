package model

import (
	"github.com/jinzhu/gorm"
	"time"
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
	Date          time.Time
	Summerization string `gorm:"size:1000"`
	UserID        uint `gorm:"index"`
	Todoes        []Todo
	Comments      []Comment
	GroupID       uint `gorm:"index"`
}

type Comment struct {
	gorm.Model
	UserID   uint `gorm:"index"`
	ReportID uint `gorm:"index"`
	Date     time.Time `json:"date" json:"date"`
	Content  string `json:"content" `
}

type Todo struct {
	gorm.Model
	Content  string
	Status   string
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

		"date": &graphql.Field{
			Type:        graphql.String,
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				report := p.Source.(Report)
				return report.Date.String(), nil
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

		"user": &graphql.Field{
			//			Type:      graphql.Type(userType),
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				comment := p.Source.(Comment)
				return getUserById(int(comment.UserID)), nil
			},
		},

		"report": &graphql.Field{
			//			Type:      graphql.Type(reportType),
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				comment := p.Source.(Comment)
				return getReportById(comment.ReportID), nil
			},
		},


		"date": &graphql.Field{
			Type:        graphql.String,
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				date := p.Source.(Comment)
				return date.Date.String(), nil
			},
		},


		"content": &graphql.Field{
			Type:        graphql.String,
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				date := p.Source.(Comment)
				return date.Content, nil
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

		"status": &graphql.Field{
			Type:        graphql.String,
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				todo := p.Source.(Todo)
				return todo.Status, nil
			},
		},

		"report": &graphql.Field{
			Type:        graphql.String,
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				todo := p.Source.(Todo)
				return getReportById(todo.ReportID), nil
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

	commentType.AddFieldConfig("user", &graphql.Field{Type: userType})
	commentType.AddFieldConfig("report", &graphql.Field{Type: reportType})
}
