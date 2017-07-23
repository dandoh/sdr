package model

func getGroupById(id int) Group {
	var group Group
	db.First(&group, id)
	return group
}

func getUsersInGroup(group Group) []User{
	var users []User
	db.Model(&group).Association("Users").Find(&users)
	return users
}

func getReportsByGroupId(id int) []Report{
	var reports []Report
	var group Group = getGroupById(id)
	db.Model(&group).Association("Reports").Find(&reports)
	return reports
}