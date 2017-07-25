package model


func getTodoesOfReport(report Report) []Todo{
	var todoes []Todo
	db.Model(&report).Association("Todoes").Find(&todoes)
	return todoes
}


func getCommentsOfReport(report Report) []Comment{
	var comments []Comment
	db.Model(&report).Association("Comments").Find(&comments)
	return comments
}

func getReportById(id uint) Report{
	var report Report
	db.First(&report, id)
	return report
}

func CreateReport(contentTodoes []string, states  []bool, summerization string, emailUser string, nameGroup string) bool {
	// A Report belongs to a user and a group.
	user := getUserByEmail(emailUser)
	group:= getGroupByName(nameGroup)
	var report Report = Report{Summerization:summerization, UserID:user.ID, GroupID:group.ID}
	db.Create(&report)
	todoes := make([]Todo, len(contentTodoes))
	for i,todo := range todoes{
		todo.Content = contentTodoes[i]
		todo.State = states[i]
		todo.ReportID = report.ID
		db.Create(&todo)
	}
	return true
}