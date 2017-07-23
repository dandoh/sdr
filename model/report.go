package model


func getTodoesOfReport(report Report) []Todo{
	var todoes []Todo
	db.Model(&report).Association("Todo").Find(&todoes)
	return todoes
}


func getCommentsOfReport(report Report) []Comment{
	var comments []Comment
	db.Model(&report).Association("Comment").Find(&comments)
	return comments
}

func getReportById(id uint) Report{
	var report Report
	db.First(&report, id)
	return report
}