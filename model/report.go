package model

func getTodoesOfReport(report Report) (todoes []Todo) {
	db.Model(&report).Association("Todoes").Find(&todoes)
	return
}

func getCommentsOfReport(report Report) (comments []Comment) {
	db.Model(&report).Association("Comments").Find(&comments)
	return
}

func getReportByID(id uint) (report Report) {
	db.First(&report, id)
	return
}

func CreateReport(contentTodoes []string, states []bool, summerization string, emailUser string, nameGroup string) bool {
	// A Report belongs to a user and a group.
	user := getUserByEmail(emailUser)
	group := getGroupByName(nameGroup)
	report := Report{
		Summerization: summerization,
		UserID:        user.ID,
		GroupID:       group.ID,
	}
	db.Create(&report)
	todoes := make([]Todo, len(contentTodoes))
	for i, todo := range todoes {
		todo.Content = contentTodoes[i]
		todo.State = states[i]
		todo.ReportID = report.ID
		db.Create(&todo)
	}
	return true
}
