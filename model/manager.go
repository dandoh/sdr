package model

// User
func findGroupsContainUser(user *User) (groups []Group) {
	db.Model(user).Association("Groups").Find(&groups)
	return
}

func findReportsOfUser(user *User) (reports []Report) {
	db.Model(user).Association("Reports").Find(&reports)
	return
}

func findCommentsOfUser(user *User) (comments []Comment) {
	db.Model(user).Association("Comments").Find(&comments)
	return
}

func findUsersByGroup(group *Group) (users []User) {
	db.Model(&group).Association("Users").Find(&users)
	return
}

func findUsersInGroup(group Group) (users []User) {
	db.Model(&group).Association("Users").Find(&users)
	return
}

func findUserByID(id int) (user User) {
	db.First(&user, id)
	return
}

func findUserByEmail(email string) (user User) {
	db.Where("email = ?", email).First(&user)
	return
}

func updateUserNote(note string, userID int) {
	db.Table("users").Where("id = ?", userID).Update("note", note)
	return
}

func findUserByName(name string) (user User) {
	db.Where("name = ?", name).First(&user)
	return
}

// Report
func insertReport(report *Report){
	db.Create(report)
	return
}

func findTodoesOfReport(report *Report) (todoes []Todo) {
	db.Model(report).Association("Todoes").Find(&todoes)
	return
}

func findCommentsOfReport(report *Report) (comments []Comment) {
	db.Model(report).Association("Comments").Find(&comments)
	return
}

func findReportsByGroup(group *Group) (reports []Report) {
	db.Model(group).Association("Reports").Find(&reports)
	return
}

func findReportByID(id uint) (report Report) {
	db.First(&report, id)
	return
}

func updateNoteOfReport(note string, report *Report){
	report.Note = note
	db.Save(report)
}


func deleteTodoesOfReport(report *Report){
	todoes := findTodoesOfReport(report)
	for _, todo := range todoes {
		db.Model(report).Association("Todoes").Delete(todo)
	}
	return
}


// Groups
func findGroupByID(id int) (group Group) {
	db.First(&group, id)
	return
}

func findGroupsByUser(user *User) (groups []Group) {
	db.Model(user).Association("Groups").Find(&groups);
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

func insertGroup(name string, purpose string) {
	db.Create(&Group{Name: name, Purpose: purpose})
}

func findGroupByName(name string) (group Group) {
	db.Where("name = ?", name).First(&group)
	return
}

func insertUserToGroup(user *User, group *Group) {
	db.Model(user).Association("Groups").Append(group)
}

func deleteUserInGroup(user *User, group *Group) {
	db.Model(user).Association("Groups").Delete(group)
}

// comment
func insertComment(comment *Comment) {
	db.Create(&comment)
}



func insertTodo(todo *Todo){
	db.Create(todo)
	return
}