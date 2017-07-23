package model

func getGroupsContainUser(user User) []Group{
	var groups []Group
	db.Model(&user).Association("Groups").Find(&groups)
	return groups
}

func getReportsOfUser(user User) []Report{
	var reports []Report
	db.Model(&user).Association("Reports").Find(&reports)
	return reports
}

func getCommentsOfUser(user User) []Comment{
	var comments []Comment
	db.Model(&user).Association("Comments").Find(&comments)
	return comments
}

func getUserById(id uint) User{
	var user User
	db.First(&user, id)
	return user
}