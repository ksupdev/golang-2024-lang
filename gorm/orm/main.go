package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:password@tcp(localhost:3306)/ormdb?charset=utf8mb4&parseTime=true"
	dialector := mysql.Open(dsn)
	db, err := gorm.Open(dialector)
	if err != nil {
		panic(err)
	}

	db.Migrator().CreateTable(Gender{})

}

type Gender struct {
	ID   uint
	Name string
}
