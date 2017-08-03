package model

import (
	"github.com/jinzhu/gorm"
	"github.com/graphql-go/graphql"
)

type Report struct {
	gorm.Model
	Summary  string `gorm:"size:1000"`
	UserID   uint `gorm:"index"`
	Status   string
	Todoes   []Todo
	Comments []Comment
	GroupID  uint `gorm:"index"`
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


		"status": &graphql.Field{
			Type:        graphql.String,
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				report := p.Source.(Report)
				return report.Status, nil
			},
		},

		"summary": &graphql.Field{
			Type:        graphql.String,
			Description: "...",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				report := p.Source.(Report)
				return report.Summary, nil
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

func findReportsByGroupID(id int) (reports []Report) {
	group := findGroupByID(id)
	reports = findReportsByGroup(&group);
	return
}

func createReport(contentTodoes []string, states []int, userId int, groupId int) {
	var report Report = Report{UserID: uint(userId), GroupID: uint(groupId), Status: "Planned"} // TODO - fix this later
	insertReport(&report)
	todoes := make([]Todo, len(contentTodoes))
	for i, todo := range todoes {
		todo.Content = contentTodoes[i]
		todo.State = states[i]
		todo.ReportID = report.ID
		insertTodo(&todo)
	}
	return
}

func updateReport(reportId int, contentTodoes []string, states []int, summary string, status string) {
	report := findReportByID(uint(reportId))
	updateSummaryOfReport(summary, &report)
	updateStatusOfReport(status, &report)

	//Delete old to-do list
	deleteTodoesOfReport(&report)

	//Create new to-do list
	todoes := make([]Todo, len(contentTodoes))
	for i, todo := range todoes {
		todo.Content = contentTodoes[i]
		todo.State = states[i]
		todo.ReportID = report.ID
		insertTodo(&todo)
	}
	return
}
