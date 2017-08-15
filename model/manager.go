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

func isUserInGroupAlready(email string, groupId int) bool{
	var users []User
	group := findGroupByID(groupId)
	users = findUsersByGroup(&group)
	for _, user := range users {

		if user.Email == email {
			return true
		}
	}

	return false

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

func saveReport(report *Report){
	db.Save(report)
}


func deleteTodoesOfReport(report *Report){
	todoes := findTodoesOfReport(report)
	for _, todo := range todoes {
		db.Model(report).Association("Todoes").Delete(todo)
	}
	return
}

func deleteTodo(todoId int) bool{

	todo := findTodoById(todoId)
	report := findReportByID(uint(todo.ReportID))
	db.Model(report).Association("Todoes").Delete(todo)
	return true
}

func updateTodo(todoId int, content string, state int, estimateTime int, spentTime int) bool{
	todo := findTodoById(todoId)
	todo.Content = content
	todo.State = state
	todo.EstimateTime = estimateTime
	todo.SpentTime = spentTime
	db.Save(todo)
	return true
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

func insertGroup(name string, purpose string) uint {
	var group Group = Group{Name: name, Purpose: purpose}
	db.Create(&group)
	return group.ID
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

//Todo

func findTodoById(todoId int) (todo Todo){
	db.Where("id = ?", todoId).First(&todo)
	return
}