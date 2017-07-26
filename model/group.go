package model

func getGroupByID(id int) (group Group) {
	db.First(&group, id)
	return
}

func getUsersInGroup(group Group) (users []User) {
	db.Model(&group).Association("Users").Find(&users)
	return
}

func getUsersByGroupID(id int) (users []User) {
	group := getGroupByID(id)
	db.Model(&group).Association("Users").Find(&users)
	return
}

func getReportsByGroupID(id int) (reports []Report) {
	group := getGroupByID(id)
	db.Model(&group).Association("Reports").Find(&reports)
	return
}

func getReportsByGroupName(name string) (reports []Report) {
	group := getGroupByName(name)
	db.Model(&group).Association("Reports").Find(&reports)
	return
}

func isNameGroupExisted(name string) bool {
	var group Group
	var count int
	db.Where("name = ?", name).Find(&group).Count(&count)
	if count > 0 {
		return true
	}
	return false
}

func createGroup(name string) bool {
	db.Create(&Group{Name: name})
	return true
}

func getGroupByName(name string) (group Group) {
	db.Where("name = ?", name).First(&group)
	return
}

func addUserToGroup(email string, groupID int) bool {
	user := getUserByEmail(email)
	group := getGroupByID(groupID)
	db.Model(&user).Association("Groups").Append(group)
	return true
}

func addUserToGroupByID(userID int, groupName string) bool {
	user := getUserByID(userID)
	group := getGroupByName(groupName)
	db.Model(&user).Association("Groups").Append(group)
	return true
}

func deleteUserInGroup(emailUser string, nameGroup string) bool {
	user := getUserByEmail(emailUser)
	group := getGroupByName(nameGroup)
	db.Model(&user).Association("Groups").Delete(group)
	return true
}
