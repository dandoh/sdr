package model

import (
	"github.com/jinzhu/gorm"
	"github.com/graphql-go/graphql"
)

type Report struct {
	gorm.Model
	Note     string `gorm:"size:1000"`
	UserID   uint `gorm:"index"`
	Todoes   []Todo
	Comments []Comment
}

var reportType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Report",
	Description: "...",
	Fields: graphql.Fields{
		"reportId": &graphql.Field{
			Type:        graphql.Int,
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				report := p.Source.(Report)
				return report.ID, nil
			},
		},



		"note": &graphql.Field{
			Type:        graphql.String,
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				report := p.Source.(Report)
				return report.Note, nil
			},
		},


		"todoes": &graphql.Field{
			Type:        graphql.NewList(todoType),
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				report := p.Source.(Report)
				return findTodoesOfReport(&report), nil
			},
		},

		"comments": &graphql.Field{
			Type:        graphql.NewList(commentType),
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				report := p.Source.(Report)
				return findCommentsOfReport(&report), nil
			},
		},

	},


})
/*
func createReport(todoes []Todo, userId int) {
	var report Report = Report{UserID: uint(userId)} // TODO - fix this later
	insertReport(&report)

	for _, todo := range todoes {
		todo.ReportID = report.ID
		insertTodo(&todo)
	}
	return
}
*/

func createReport(userId int) int{
	var report Report = Report{UserID: uint(userId)}
	insertReport(&report)
	return  int(report.ID)
}
/*
func updateReport(reportId int, todoes []Todo, note string) {
	report := findReportByID(uint(reportId))
	updateNoteOfReport(note, &report)

	//Delete old to-do list
	deleteTodoesOfReport(&report)

	//Create new to-do list
	for _, todo := range todoes {
		todo.ReportID = uint(reportId)
		insertTodo(&todo)
	}
	return
}
*/
func updateNoteOfReport(note string, reportId int) string{
	report := findReportByID(uint(reportId))
	report.Note = note
	saveReport(&report)
	return report.Note
}



