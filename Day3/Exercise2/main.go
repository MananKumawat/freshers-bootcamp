package main

import (
	"awesomeProject/Day3/Exercise2/Config"
	"awesomeProject/Day3/Exercise2/Models"
	"awesomeProject/Day3/Exercise2/Routes"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	//"github.com/jinzhu/gorm"
)
var err error
func main() {

	Config.DB, err = gorm.Open(mysql.Open(Config.DbURL(Config.BuildDBConfig())), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("Status:", err)
	}

	Config.DB.AutoMigrate(&Models.Subject{})
	Config.DB.AutoMigrate(&Models.Student{})

	r := Routes.SetupRouter()
	//running
	r.Run()

}