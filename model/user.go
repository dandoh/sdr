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

func getUserById(id int) User{
	var user User
	db.First(&user, id)
	return user
}

func getUserByEmail(email string) User{
	var user User
	db.Where("email = ?",email).First(&user)
	return user
}

func getGroupsByUserId(id int) []Group{
	var groups []Group
	var user User = getUserById(id)
	db.Model(&user).Association("Groups").Find(&groups)
	return groups
}

func isEmailExisted(email string) bool{
	var user User
	var count int
	db.Where("email = ?",email).Find(&user).Count(&count)
	if (count > 0){
		return true
	}
	return false
}


func createAccount(name string, password string, email string) bool{
	var user User = User{Name:name, PasswordMD5:password, Email:email}
	db.Create(&user)
	return true
}

func updateNoteForUser(note string, userId int) bool{
	user := getUserById(userId)
	db.Model(&user).Update("note", note)
	return true
}