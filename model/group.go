package model

import (
	"github.com/jinzhu/gorm"
	"github.com/graphql-go/graphql"
)

type Group struct {
	gorm.Model
	Name    string `gorm:"size:255; unique_index" json:"name"`
	Purpose string `json:"purpose"`
	Users   []User    `gorm:"many2many:user_group" json:"users"`
}

var groupType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Group",
	Description: "Group of users",
	Fields: graphql.Fields{
		"groupId": &graphql.Field{
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
		"purpose": &graphql.Field{
			Type:        graphql.String,
			Description: "The group's purpose",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				group := p.Source.(Group)
				return group.Purpose, nil
			},
		},

	},
})

func findGroupsByUserID(id int) (groups []Group) {
	user := findUserByID(id)
	groups = findGroupsByUser(&user);
	return
}

func updateGroupInfo(groupId int, groupName string, purpose string, emails []string) (bool){
	users := findUsersByGroupID(groupId)
	for _, user := range users{
		deleteUserInGroupByEmail(user.Email, groupId)
	}
	group := findGroupByID(groupId)
	group.Name = groupName
	group.Purpose = purpose

	numNewUser := 0
	for _, email := range emails {
		if (!isUserInGroupAlready(email, groupId) && isEmailExisted(email)) {
			insertUserToGroupByEmail(email, groupId)
			numNewUser ++
		}
	}

	db.Save(&group)

	if (numNewUser == len(emails)) {
		return true
	}

	if (numNewUser != 0) {
		return false
	}

	return false
}
