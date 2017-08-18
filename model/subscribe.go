package model

import (
	_"github.com/dandoh/sdr/util"
	"github.com/jinzhu/gorm"
	_"github.com/graphql-go/graphql"
	_"fmt"
)

type Subscribe struct {
	gorm.Model
	UserId       uint
	ReportId uint
	CommentsNotSeen int
}