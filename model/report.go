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



func createTodayReportForUser(userId int) int{
	var report Report = Report{UserID: uint(userId)}
	yesterdayReport := findReportYestedayByUserId(userId)
	yesterdayTodoes := findTodoesOfReport(&yesterdayReport)
	insertReport(&report)
	for _,todo := range yesterdayTodoes{
		if todo.State == 0{
			insertTodo(&Todo{Content: todo.Content, State: todo.State,
				EstimateTime: todo.EstimateTime, SpentTime: todo.SpentTime, ReportID: report.ID})
		}
	}

	return  int(report.ID)
}

func createTodayReportForAllUsers() bool{
	users := findAllUsers()
	for _, user := range users{
		createTodayReportForUser(int(user.ID))
	}
	return true
}

func updateNoteOfReport(note string, reportId int) string{
	report := findReportByID(uint(reportId))
	report.Note = note
	saveReport(&report)
	return report.Note
}

func getAllReportsTodayByGroupId(groupId int) (reports []Report){
	users := findUsersByGroupID(groupId)
	for _, user := range users{
		 reportToday := findReportTodayByUserId(int(user.ID))
		if reportToday.ID != 0{
			reports = append(reports, reportToday)
		}
	}

	return
}

