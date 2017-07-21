package model

import(
  "time"
  _"github.com/jinzhu/gorm"
  _"github.com/jinzhu/gorm/dialects/postgres"
)



type User struct{

  Id  int `gorm:"AUTO_INCREMENT"`
  UserName string `gorm:"size:255"`
  PassWord string `gorm:"size:255"`
  Email string     `gorm:"not null"`
  Token string
  Note string
}

type Group struct{

  ID  int `gorm:"AUTO_INCREMENT"`
  Name string `gorm:"size:255"`

  Users []User `gorm:"many2many:group_users"`
}

type Report struct{
 
  ReportId int `gorm:"AUTO_INCREMENT"`
  Date time.Time
  Summerization string
  User User
  UserId string
  Todoes []Todo
}

type Todo struct{
 
  Id int `gorm:"AUTO_INCREMENT"`
  Content string
  IsCompleted int

  Report Report
  ReportId string

}

type Comment struct{
  
  Id int `gorm:"AUTO_INCREMENT"`
  Content string
  Date time.Time

  User User
  UserId string

  Report Report
  ReportId string

}

