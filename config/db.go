package config

import (
	. "ahmadarif/gin-gorm/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func init() {
	//open a db connection
	var err error
	DB, err = gorm.Open("sqlite3", "tmp/gorm.db")
	if err != nil {
		panic(err)
	}

	//Migrate the schema
	DB.AutoMigrate(&Todo{})
}
