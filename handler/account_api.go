package handler

import(
 	"fmt"
    "net/http"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    _"database/sql"
    _"github.com/go-sql-driver/mysql"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

import "project/model"
import "project/manager"

const(
  DB_USER = "postgres"
  DB_PASSWORD = "123456789"
  DB_NAME = "Scoville_Info"
)


func SignupPage(res http.ResponseWriter, req *http.Request) {
  
  //connection
  dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER,DB_PASSWORD, DB_NAME)
  db, err := gorm.Open("postgres", dbinfo) 

  //Check if there is wrong with connection to database. 
  if(err != nil){
    panic("wrong")
  }

   if (req.Method != "POST") {
    http.ServeFile(res, req, "signup.html")
    
  }

  // Retreive data sent by client in signup page.
  username := req.FormValue("username")
  password := req.FormValue("password")
  email	   := req.FormValue("email")
  
  //Check if the username has existed. If there is not, insert infomation of user to the database.
  
  if (manager.GetUserInfo(email) == model.User{}) {
  	if(!db.HasTable(&model.User{})){
   		db.CreateTable(&model.User{})
	 }   

  	db.Create(&model.User{UserName: username, PassWord: password, Email: email, Token: "", Note : ""})   
  	return  
  }

  	return 
 

}