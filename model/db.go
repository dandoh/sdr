package model

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"sdr/util"
)

const (
	DB_USER     = "Dandoh"
	DB_PASSWORD = "dandoh"
	DB_NAME     = "sdr_gorm"
)

var db *gorm.DB

const dev = true

func Init() {
	var err error
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err = gorm.Open("postgres", dbinfo)

	if err != nil {
		panic("failed to connect database")
	}

	if dev {
		db.DropTable(&User{})
		db.DropTable(&Group{})
		db.DropTable(&Report{})
		db.DropTable(&Comment{})
		db.DropTable(&Todo{})
	}

	db.LogMode(true)
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Group{})
	db.AutoMigrate(&Report{})
	db.AutoMigrate(&Comment{})
	db.AutoMigrate(&Todo{})

	if dev {

		user1 := User{Name: "Nhan", Email: "Dandoh@gmail.com", PasswordMD5: util.GetMD5Hash("haha")}
		user2 := User{Name: "De", Email: "De@gmail.com", PasswordMD5: util.GetMD5Hash("haha")}
		user3 := User{Name: "Shiki", Email: "Shiki@gmail.com", PasswordMD5: util.GetMD5Hash("haha")}
		db.Create(&user1)
		db.Create(&user2)
		db.Create(&user3)

		group1 := Group{Name: "intern"}
		group2 := Group{Name: "intern2"}
		group3 := Group{Name: "intern3"}
		db.Create(&group1)
		db.Create(&group2)
		db.Create(&group3)

		report1 := Report{
			Summerization: "This is summerization of report 1",
			UserID:user1.ID,
			GroupID:group1.ID,
		}
		report2 := Report{
			Summerization: "This is summerization of report 2",
			UserID:user2.ID,
			GroupID:group1.ID,
		}


		db.Create(&report1)
		db.Create(&report2)

		db.Model(&user1).Association("Groups").Append(group1)
		db.Model(&user1).Association("Groups").Append(group3)
		db.Model(&user1).Association("Groups").Append(group2)
		db.Model(&user2).Association("Groups").Append(group3)
		db.Model(&user2).Association("Groups").Append(group1)
		db.Model(&user3).Association("Groups").Append(group1)
		db.Model(&user3).Association("Groups").Append(group2)

		Todo1 := Todo{
			State: true,
			Content:" Content todo1",
			ReportID: report1.ID,
		}

		db.Create(&Todo1)
	}

}

func Get() *gorm.DB {
	return db
}

func Close() {
	db.Close()
}
