package main

import (
	"gin-practice/config"
	"gin-practice/initialize"
)

func main() {
	config.SELF_VIPER = initialize.GetConfig()
	if config.SELF_DB = initialize.GetGorm(); config.SELF_DB != nil {
		db, _ := config.SELF_DB.DB()
		initialize.CreateTables(config.SELF_DB)
		defer db.Close()
	}
}
