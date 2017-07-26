package model

import "github.com/dandoh/sdr/util"

func getGroupsContainUser(user User) (groups []Group) {
	db.Model(&user).Association("Groups").Find(&groups)
	return
}

func getReportsOfUser(user User) (reports []Report) {
	db.Model(&user).Association("Reports").Find(&reports)
	return
}

func getCommentsOfUser(user User) (comments []Comment) {
	db.Model(&user).Association("Comments").Find(&comments)
	return
}

func getUserByID(id int) (user User) {
	db.First(&user, id)
	return
}

func getUserByEmail(email string) (user User) {
	db.Where("email = ?", email).First(&user)
	return
}

func getGroupsByUserID(id int) (groups []Group) {
	user := getUserByID(id)
	db.Model(&user).Association("Groups").Find(&groups)
	return
}

func isEmailExisted(email string) bool {
	var user User
	var count int
	db.Where("email = ?", email).Find(&user).Count(&count)
	if count > 0 {
		return true
	}
	return false
}

func createAccount(name string, password string, email string) bool {
	user := User{
		Name:        name,
		PasswordMD5: password,
		Email:       email,
	}
	db.Create(&user)
	return true
}

func updateNoteForUser(note string, userID int) bool {
	user := getUserByID(userID)
	db.Model(&user).Update("note", note)
	return true
}

func GetUserID(username string, password string) (uint, bool) {
	var user User
	db.Where("name = ?", username).First(&user)
	if user.PasswordMD5 == util.GetMD5Hash(password) {
		return user.ID, true
	}
	return 0, false
}
