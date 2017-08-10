package model

import (
	"github.com/dandoh/sdr/util"
	"github.com/jinzhu/gorm"
	"github.com/graphql-go/graphql"
	"fmt"
)

type User struct {
	gorm.Model
	Name        string `gorm:"size:255; unique"`
	PasswordMD5 string `gorm:"size:255"`
	Email       string `gorm:"not null; unique"`
	Token       string
	//Note        string    `gorm:"size:2000"`
	Groups      []Group `gorm:"many2many:user_group"`
	Reports     []Report
	Comments    []Comment
}

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "User",
	Description: "...",
	Fields: graphql.Fields{
		"userId": &graphql.Field{
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
/*
		"note": &graphql.Field{
			Type:        graphql.String,
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user := p.Source.(User)
				return user.Note, nil
			},
		},
*/

		"groups": &graphql.Field{
			Type:        graphql.NewList(groupType),
			Description: "Which posts they have written.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user := p.Source.(User)
				return findGroupsContainUser(&user), nil
			},
		},

		"reports": &graphql.Field{
			Type:        graphql.NewList(reportType),
			Description: "Which posts they have written.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user := p.Source.(User)
				return findReportsOfUser(&user), nil
			},
		},

		"comments": &graphql.Field{
			Type:        graphql.NewList(commentType),
			Description: "Which posts they have written.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user := p.Source.(User)
				return findCommentsOfUser(&user), nil
			},
		},
		//Password and Token haven't been declared.
	},
})

func GetUserID(email string, password string) (uint, bool) {
	user := findUserByEmail(email);
	fmt.Printf("Received %s, expected %s", util.GetMD5Hash(password), user.PasswordMD5);
	if user.PasswordMD5 == util.GetMD5Hash(password) {
		return user.ID, true
	}
	return 0, false
}

func findUsersByGroupID(id int) (users []User) {
	group := findGroupByID(id)
	users = findUsersByGroup(&group)
	return
}

func insertUserToGroupByEmail(email string, groupID int) bool {
	user := findUserByEmail(email)
	group := findGroupByID(groupID)
	insertUserToGroup(&user, &group);
	return true
}

func insertUserToGroupByID(userID int, groupID int) bool {
	user := findUserByID(userID)
	group := findGroupByID(groupID)
	insertUserToGroup(&user, &group);
	return true
}

func deleteUserInGroupByEmail(emailUser string, groupId int) bool {
	user := findUserByEmail(emailUser)
	group := findGroupByID(groupId)
	deleteUserInGroup(&user, &group);
	return true
}

func IsUserExisted(name string, email string) bool {
	var user User
	var count int
	db.Where("email = ?", email).Find(&user).Count(&count)
	if count > 0 {
		return true
	}
	db.Where("name = ?", name).Find(&user).Count(&count)
	if count > 0{
		return true
	}
	return false
}

func CreateUser(user *User){
	db.Create(user)
	return
}