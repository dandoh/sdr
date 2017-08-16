package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/dandoh/sdr/util"
	"os"
	"github.com/enodata/faker"
)

var db *gorm.DB

const dev = true

func Init() {
	var err error
	var (
		DBUSER     = os.Getenv("DBUSER")
		DBPASSWORD = os.Getenv("DBPASSWORD")
		DBNAME     = os.Getenv("DBNAME")
	)
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DBUSER, DBPASSWORD, DBNAME)
	db, err = gorm.Open("postgres", dbInfo)
	fmt.Println(dbInfo)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	db.LogMode(true)

	if dev {
		db.DropTable(&User{})
		db.DropTable(&Group{})
		db.DropTable("user_group")
		db.DropTable(&Report{})
		db.DropTable(&Comment{})
		db.DropTable(&Todo{})
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Group{})
	db.AutoMigrate(&Report{})
	db.AutoMigrate(&Comment{})
	db.AutoMigrate(&Todo{})

	if dev {

		user1 := User{Name: "Nhan", Email: "Dandoh@gmail.com", PasswordMD5: util.GetMD5Hash("haha"), Avatar: faker.Avatar().String()}
		user2 := User{Name: "De", Email: "De@gmail.com", PasswordMD5: util.GetMD5Hash("haha"), Avatar: faker.Avatar().String()}
		user3 := User{Name: "Shiki", Email: "Shiki@gmail.com", PasswordMD5: util.GetMD5Hash("haha"), Avatar: faker.Avatar().String()}
		db.Create(&user1)
		db.Create(&user2)
		db.Create(&user3)

		group1 := Group{Name: faker.Company().Name()}
		group2 := Group{Name: faker.Company().Name()}
		group3 := Group{Name: faker.Company().Name()}

		db.Create(&group1)
		db.Create(&group2)
		db.Create(&group3)

		report1 := Report{
			Note:   faker.Lorem().Sentence(5),
			UserID: user1.ID,
		}
		report2 := Report{
			Note:  faker.Lorem().Sentence(5),
			UserID: user2.ID,
		}

		db.Create(&report1)
		db.Create(&report2)

		db.Model(&user1).Association("Groups").Append(group1)
		db.Model(&user1).Association("Groups").Append(group3)
		db.Model(&user1).Association("Groups").Append(group2)
		//db.Model(&user2).Association("Groups").Append(group3)
		db.Model(&user2).Association("Groups").Append(group1)
		db.Model(&user3).Association("Groups").Append(group1)

		Todo1 := Todo{
			State:        0, // haven't done yet
			Content:      faker.Company().Bs(),
			EstimateTime: 150,
			SpentTime:    120,
			ReportID:     report1.ID,
		}

		Comment1 := Comment{
			UserID:   user1.ID,
			ReportID: report2.ID,
			Content:  faker.Company().Bs(),
		}

		Comment2 := Comment{
			UserID:   user2.ID,
			ReportID: report2.ID,
			Content:  faker.Company().Bs(),
		}

		db.Create(&Comment1)
		db.Create(&Comment2)
		db.Create(&Todo1)
	}

}

func Get() *gorm.DB {
	return db
}

func Close() {
	db.Close()
}
