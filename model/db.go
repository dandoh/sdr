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
		db.DropTable(&Subscribe{})
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Group{})
	db.AutoMigrate(&Report{})
	db.AutoMigrate(&Comment{})
	db.AutoMigrate(&Todo{})
	db.AutoMigrate(&Subscribe{})

	if dev {

		user1 := User{Name: "Nhan", Email: "Dandoh@gmail.com", PasswordMD5: util.GetMD5Hash("haha"), Avatar: faker.Avatar().String()}
		user2 := User{Name: "De", Email: "De@gmail.com", PasswordMD5: util.GetMD5Hash("haha"), Avatar: faker.Avatar().String()}
		user3 := User{Name: "Shiki", Email: "Shiki@gmail.com", PasswordMD5: util.GetMD5Hash("haha"), Avatar: faker.Avatar().String()}
		db.Create(&user1)
		db.Create(&user2)
		db.Create(&user3)

		group1 := Group{Name: faker.Company().Name(), Purpose: "Team Building"}
		group2 := Group{Name: faker.Company().Name(), Purpose: "Learning React Native"}
		group3 := Group{Name: faker.Company().Name(), Purpose: "Learning Linux"}

		db.Create(&group1)
		db.Create(&group2)
		db.Create(&group3)

		db.Model(&user1).Association("Groups").Append(group1)
		db.Model(&user1).Association("Groups").Append(group3)
		db.Model(&user1).Association("Groups").Append(group2)
		//db.Model(&user2).Association("Groups").Append(group3)
		db.Model(&user2).Association("Groups").Append(group1)
		db.Model(&user3).Association("Groups").Append(group1)

		todayreport1 := Report{
			Note:   faker.Lorem().Sentence(5),
			UserID: user1.ID,
		}
		todayreport2 := Report{
			Note:  faker.Lorem().Sentence(5),
			UserID: user2.ID,
		}
		todayreport3 := Report{
			Note: faker.Lorem().Sentence(10),
			UserID: user3.ID,
		}


		db.Create(&todayreport1)
		db.Create(&todayreport2)
		db.Create(&todayreport3)


		//////////////////////////////// Today Report 1////////////////////////
		Todo1_1 := Todo{
			State:        1, // haven't done yet
			Content:      faker.Company().Bs(),
			EstimateTime: 150,
			SpentTime:    120,
			ReportID:     todayreport1.ID,
		}

		Todo1_2 := Todo{
			State:        0, // haven't done yet
			Content:      faker.Company().Bs(),
			EstimateTime: 100,
			SpentTime:    100,
			ReportID:     todayreport1.ID,
		}

		Todo1_3 := Todo{
			State:        0, // haven't done yet
			Content:      faker.Company().Bs(),
			EstimateTime: 130,
			SpentTime:    130,
			ReportID:     todayreport1.ID,
		}



		Comment1_1 := Comment{
			UserID:   user1.ID,
			ReportID: todayreport1.ID,
			Content:  faker.Company().Bs(),
		}

		Comment1_2 := Comment{
			UserID:   user2.ID,
			ReportID: todayreport1.ID,
			Content:  faker.Company().Bs(),
		}


		Comment1_3 := Comment{
			UserID:   user3.ID,
			ReportID: todayreport1.ID,
			Content:  faker.Company().Bs(),
		}


		db.Create(&Todo1_1)
		db.Create(&Todo1_2)
		db.Create(&Todo1_3)
		db.Create(&Comment1_1)
		db.Create(&Comment1_2)
		db.Create(&Comment1_3)
		///////////////////////////////////Today report 2//////////////

		Todo2_1 := Todo{
			State:        0, // haven't done yet
			Content:      faker.Company().Bs(),
			EstimateTime: 150,
			SpentTime:    120,
			ReportID:     todayreport2.ID,
		}

		Todo2_2 := Todo{
			State:        0, // haven't done yet
			Content:      faker.Company().Bs(),
			EstimateTime: 100,
			SpentTime:    100,
			ReportID:     todayreport2.ID,
		}

		Todo2_3 := Todo{
			State:        0, // haven't done yet
			Content:      faker.Company().Bs(),
			EstimateTime: 130,
			SpentTime:    130,
			ReportID:     todayreport2.ID,
		}



		Comment2_1 := Comment{
			UserID:   user1.ID,
			ReportID: todayreport2.ID,
			Content:  faker.Company().Bs(),
		}

		Comment2_2 := Comment{
			UserID:   user2.ID,
			ReportID: todayreport2.ID,
			Content:  faker.Company().Bs(),
		}


		Comment2_3 := Comment{
			UserID:   user3.ID,
			ReportID: todayreport2.ID,
			Content:  faker.Company().Bs(),
		}


		db.Create(&Todo2_1)
		db.Create(&Todo2_2)
		db.Create(&Todo2_3)
		db.Create(&Comment2_1)
		db.Create(&Comment2_2)
		db.Create(&Comment2_3)
		/////////////////////////////////Today report 3////////////



		Todo3_1 := Todo{
			State:        0, // haven't done yet
			Content:      faker.Company().Bs(),
			EstimateTime: 150,
			SpentTime:    120,
			ReportID:     todayreport3.ID,
		}

		Todo3_2 := Todo{
			State:        0, // haven't done yet
			Content:      faker.Company().Bs(),
			EstimateTime: 100,
			SpentTime:    100,
			ReportID:     todayreport3.ID,
		}

		Todo3_3 := Todo{
			State:        0, // haven't done yet
			Content:      faker.Company().Bs(),
			EstimateTime: 130,
			SpentTime:    130,
			ReportID:     todayreport3.ID,
		}



		Comment3_1 := Comment{
			UserID:   user1.ID,
			ReportID: todayreport3.ID,
			Content:  faker.Company().Bs(),
		}

		Comment3_2 := Comment{
			UserID:   user2.ID,
			ReportID: todayreport3.ID,
			Content:  faker.Company().Bs(),
		}


		Comment3_3 := Comment{
			UserID:   user3.ID,
			ReportID: todayreport3.ID,
			Content:  faker.Company().Bs(),
		}


		db.Create(&Todo3_1)
		db.Create(&Todo3_2)
		db.Create(&Todo3_3)
		db.Create(&Comment3_1)
		db.Create(&Comment3_2)
		db.Create(&Comment3_3)


		Subscribe1_1 := Subscribe{
			UserId: user1.ID,
			ReportId: todayreport1.ID,
			NumberCommentsNotSeen: 0,
		}

		Subscribe1_2 := Subscribe{
			UserId: user1.ID,
			ReportId: todayreport2.ID,
			NumberCommentsNotSeen: 0,
		}
		Subscribe1_3 := Subscribe{
			UserId: user1.ID,
			ReportId: todayreport3.ID,
			NumberCommentsNotSeen: 0,
		}
		Subscribe2_1 := Subscribe{
			UserId: user2.ID,
			ReportId: todayreport1.ID,
			NumberCommentsNotSeen: 0,
		}
		Subscribe2_2 := Subscribe{
			UserId: user2.ID,
			ReportId: todayreport2.ID,
			NumberCommentsNotSeen: 0,
		}
		Subscribe2_3 := Subscribe{
			UserId: user2.ID,
			ReportId: todayreport3.ID,
			NumberCommentsNotSeen: 0,
		}
		Subscribe3_1 := Subscribe{
			UserId: user3.ID,
			ReportId: todayreport1.ID,
			NumberCommentsNotSeen: 0,
		}
		Subscribe3_2 := Subscribe{
			UserId: user3.ID,
			ReportId: todayreport2.ID,
			NumberCommentsNotSeen: 0,
		}
		Subscribe3_3 := Subscribe{
			UserId: user3.ID,
			ReportId: todayreport3.ID,
			NumberCommentsNotSeen: 0,
		}

		db.Create(&Subscribe1_1)
		db.Create(&Subscribe1_2)
		db.Create(&Subscribe1_3)
		db.Create(&Subscribe2_1)
		db.Create(&Subscribe2_2)
		db.Create(&Subscribe2_3)
		db.Create(&Subscribe3_1)
		db.Create(&Subscribe3_2)
		db.Create(&Subscribe3_3)

		Subscribe1_1.LastUpdatedAt = Subscribe1_1.UpdatedAt
		db.Save(&Subscribe1_1)
		Subscribe1_2.LastUpdatedAt = Subscribe1_2.UpdatedAt
		db.Save(&Subscribe1_2)
		Subscribe1_3.LastUpdatedAt = Subscribe1_3.UpdatedAt
		db.Save(&Subscribe1_3)
		Subscribe2_1.LastUpdatedAt = Subscribe2_1.UpdatedAt
		db.Save(&Subscribe2_1)
		Subscribe2_2.LastUpdatedAt = Subscribe2_2.UpdatedAt
		db.Save(&Subscribe2_2)
		Subscribe2_3.LastUpdatedAt = Subscribe2_3.UpdatedAt
		db.Save(&Subscribe2_3)
		Subscribe3_1.LastUpdatedAt = Subscribe3_1.UpdatedAt
		db.Save(&Subscribe3_1)
		Subscribe3_2.LastUpdatedAt = Subscribe3_2.UpdatedAt
		db.Save(&Subscribe3_2)
		Subscribe3_3.LastUpdatedAt = Subscribe3_3.UpdatedAt
		db.Save(&Subscribe3_3)
	}



}

func Get() *gorm.DB {
	return db
}

func Close() {
	db.Close()
}
