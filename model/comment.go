package model

func CreateComment(content string, userId uint, reportId uint) bool {
	var comment Comment = Comment{Content:content, UserID: userId, ReportID: reportId}
	db.Create(&comment)
	return true
}