package main

import (
	"github.com/bb3104/gorm_sample/db_model"
)

func main() {
	db := db_model.GormConnect()
	db_model.DbCreate()

	defer db.Close()

	db.LogMode(true)
}
