package main

import (
	"ahmadarif/gin-gorm/config"
	"ahmadarif/gin-gorm/routes"
)

func main() {
	defer config.DB.Close()
	routes.InitRoutes()
}
