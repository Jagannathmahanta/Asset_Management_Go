package main

import (
	"AssetManagementSystem/Config"
	"AssetManagementSystem/Routes"
	"github.com/joho/godotenv"

	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	godotenv.Load(".env")
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	//fmt.Println("db connected sucessfully")
	defer Config.DB.Close()
	//Config.DB.AutoMigrate(&Models.Users{})

	r := Routes.SetupRouter()
	//running
	r.Run()
}
