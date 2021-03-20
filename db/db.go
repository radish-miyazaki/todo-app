package db

import (
	"github.com/jinzhu/gorm"
	"github.com/radish-miyazaki/todo-app/models"

	_ "github.com/mattn/go-sqlite3"
)

var db *gorm.DB

func Initialize() {
	db, _ = gorm.Open("sqlite3", "todo.sqlite3")

	db.AutoMigrate(&models.Todo{})
}

func Get() *gorm.DB {
	return db
}

func Close() {
	db.Close()
}
