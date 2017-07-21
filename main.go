package main


import(
	"fmt"

    "net/http"
  	"github.com/jinzhu/gorm"
  	_"github.com/jinzhu/gorm/dialects/postgres"
)

import "project/model"
import "project/handler"

const(
  DB_USER = "postgres"
  DB_PASSWORD = "123456789"
  DB_NAME = "Scoville_Info"
)

func main(){
  dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
  db, err := gorm.Open("postgres", dbinfo)

  if (err != nil){
     panic("There is some thing wrong with making connection to database")
  }
  
  if (!db.HasTable(&model.Group{})){
    db.CreateTable(&model.Group{})
  }
  
  if (!db.HasTable(&model.User{})){
    db.CreateTable(&model.User{})
  }

  if (!db.HasTable(&model.Report{})){
    db.CreateTable(&model.Report{})
  }

  if (!db.HasTable(&model.Todo{})){
    db.CreateTable(&model.Todo{})
  }

  if (!db.HasTable(&model.Comment{})){
    db.CreateTable(&model.Comment{})
  }
 

  http.HandleFunc("/signup", handler.SignupPage)
  http.ListenAndServe(":7777", nil)

}