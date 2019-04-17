package main

import (
	"github.com/bb3104/gorm_sample/db_model"
)

func main() {
	db := db_model.GormConnect()
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&db_model.User{})

	defer db.Close()

	db.LogMode(true)
}
