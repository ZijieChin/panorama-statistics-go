package main

import (
	"Panorama-Statistics/database"
	"Panorama-Statistics/router"
	"Panorama-Statistics/util"
	"fmt"
)

// @title Panorama Statistics
// @version 1.0
// @description Panorama Statistics api

// @contact.name ZijieQin
// @contact.email qinzijie1@swirecocacola.com
// @BasePath /v1
func main() {
	fmt.Println("Reading configuration...")
	config := util.GetConfig()
	fmt.Println("Success.")
	database.SyncDatabase()
	fmt.Println("Database connected.")
	engine := router.InitRouter()
	fmt.Println("Service successfully initialized on port", config.Port, ".")
	engine.Run(":" + config.Port)
}
