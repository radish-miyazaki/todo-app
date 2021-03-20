package main

import (
	"github.com/radish-miyazaki/todo-app/db"
	"github.com/radish-miyazaki/todo-app/router"
)

func main() {
	db.Initialize()
	defer db.Close()

	router.Router()
}
