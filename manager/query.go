package manager

import(
	"fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    _"database/sql"
    _"github.com/go-sql-driver/mysql"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

import "project/model"

const(
  DB_USER = "postgres"
  DB_PASSWORD = "123456789"
  DB_NAME = "Scoville_Info"
)


func GetUserInfo(email string) model.User{
  dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
  db, err := gorm.Open("postgres", dbinfo)
  if (err != nil){
  	panic("wrong")
  	return model.User{}
  }
  var user model.User  
  db.Where("email= ?", email).First(& user) 	
  return user
}