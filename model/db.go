package model

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

const (
	DB_USER     = "Dandoh"
	DB_PASSWORD = "dandoh"
	DB_NAME     = "sdr_gorm"
)

var db *gorm.DB

func Init() {
	var err error
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err = gorm.Open("postgres", dbinfo)

	if err != nil {
		panic("failed to connect database")
	}

	if !db.HasTable(&Group{}) {
		db.CreateTable(&Group{})
	}

	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})
	}

	if !db.HasTable(&Report{}) {
		db.CreateTable(&Report{})
	}

	if !db.HasTable(&Todo{}) {
		db.CreateTable(&Todo{})
	}

	if !db.HasTable(&Comment{}) {
		db.CreateTable(&Comment{})
	}
	db.LogMode(true)

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Group{})
	db.AutoMigrate(&Report{})
	db.AutoMigrate(&Comment{})
	db.AutoMigrate(&Todo{})

	user1 := User{Name:"Nhan", Email:"Dandoh@mgail.com", PasswordMD5:"haha"}
	user2 := User{Name:"De", Email:"De@mgail.com", PasswordMD5:"haha"}
	user3 := User{Name:"Shiki", Email:"Shiki@mgail.com", PasswordMD5:"haha"}
	db.Create(&user1)
	db.Create(&user2)
	db.Create(&user3)

	group1 := Group{Name:"intern"}
	group2 := Group{Name:"intern2"}
	group3 := Group{Name:"intern3"}
	db.Create(&group1)
	db.Create(&group2)
	db.Create(&group3)

	db.Model(&user1).Association("Groups").Append(group1)
	db.Model(&user1).Association("Groups").Append(group3)
	db.Model(&user1).Association("Groups").Append(group2)
	db.Model(&user2).Association("Groups").Append(group3)
	db.Model(&user3).Association("Groups").Append(group1)
	db.Model(&user3).Association("Groups").Append(group2)

}

func Get() *gorm.DB {
	return db
}

func Close() {
	db.DropTable(&User{})
	db.DropTable(&Group{})
	db.DropTable(&Report{})
	db.DropTable(&Comment{})
	db.DropTable(&Todo{})
}
