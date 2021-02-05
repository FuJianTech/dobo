package database

import (
	"dobo/config"
	"dobo/database/model"
	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)


type auth model.Auth


// DB global variable to access gorm
var DB *gorm.DB
var err error



// InitDB - function to initialize db
func InitDB() *gorm.DB {
	var db = DB

	configureDB := config.ConfigInit()

	driver := configureDB.Database.DbDriver
	username := configureDB.Database.DbUser
	password := configureDB.Database.DbPass
	database := configureDB.Database.DbName
	host := configureDB.Database.DbHost
	port := configureDB.Database.DbPort

	switch driver {
	case "mysql":
		db, err = gorm.Open(driver, username+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8mb4&parseTime=True&loc=Local")
		if err != nil {
			// fmt.Println("DB err: ", err)
			log.Error(err)
		}
		// Only for debugging
		if err == nil {
			log.Info("DB connection successful!")
		}
		break
	case "postgres":
		db, err = gorm.Open(driver, "host="+host+" port="+port+" user="+username+" dbname="+database+" password="+password)
		if err != nil {
			// fmt.Println("DB err: ", err)

			log.Error(err)
		}
		// Only for debugging
		if err == nil {
			log.Info("DB connection successful!")
		}
		break
	case "sqlite3":
		db, err = gorm.Open(driver, database)
		if err != nil {
			log.Error(err)
		}
		// Only for debugging
		if err == nil {
			log.Info("DB connection successful!")
		}
		break
	default:
		log.Error("The driver " + driver + " is not implemented yet")
	}
	//db.AutoMigrate(&userHobby{}, &hobby{}, &post{}, &user{}, &auth{})

	DB = db
	migrateTables()
	admin := auth{
		Email:    "admin@gibbon.com",
		//Password: "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92",
		Password: model.HashPass("123456"),
		UserName: "admin",
	}
	DB.FirstOrCreate(&admin, auth{Email: "admin@gibbon.com"})

	return DB
}


func migrateTables() {
	configureDB := config.ConfigInit()
	driver := configureDB.Database.DbDriver

	if driver == "mysql" {
		// db.Set() --> add table suffix during auto migration
		if err := DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&auth{}).Error; err != nil {
			//errorState = 1
			log.Error(err)
			//fmt.Println(err)
		} else {
			log.Info("New tables are  migrated successfully!")
		}
	} else {
		if err := DB.AutoMigrate(&auth{}).Error; err != nil {
			//errorState = 1
			log.Error(err)
			//fmt.Println(err)
		} else {
			log.Info("New tables are  migrated successfully!")
		}
	}
}


// GetDB - get a connection
func GetDB() *gorm.DB {
	return DB
}
