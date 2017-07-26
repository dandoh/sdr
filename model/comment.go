package model

func CreateComment(content string, userID uint, reportID uint) bool {
	comment := Comment{Content: content, UserID: userID, ReportID: reportID}
	db.Create(&comment)
	return true
}
