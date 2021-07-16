package main

import (
	"Panorama-Statistics/database"
	"Panorama-Statistics/router"
	"Panorama-Statistics/util"
	"fmt"
	"net"
)

// @title Panorama Statistics
// @version 1.0
// @description Panorama Statistics api

// @contact.name Zijie Qin
// @contact.email qinzijie1@swirecocacola.com
// @BasePath /v1
func main() {
	fmt.Println("Reading configuration...")
	config := util.GetConfig()
	_, err := net.Dial("tcp", "0.0.0.0:"+config.Port)
	if err == nil {
		fmt.Println("Port", config.Port, "is occupied, please change field 'port' in config.yml and run again.")
		return
	} else {
		fmt.Println("Success. Service will be running on port", config.Port, "as configured.")
	}
	database.SyncDatabase()
	fmt.Println("Database connected.")
	engine := router.InitRouter()
	fmt.Println("Service successfully initialized on port", config.Port, ".")
	fmt.Println("Plaese do not close this window.")
	engine.Run(":" + config.Port)
}
