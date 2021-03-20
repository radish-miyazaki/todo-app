package router

import (
	"github.com/gin-gonic/gin"
	"github.com/radish-miyazaki/todo-app/controllers"
	"github.com/radish-miyazaki/todo-app/db"
)

func Router() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*.html")

	h := controllers.TaskHandler{
		Db: db.Get(),
	}

	router.GET("/", h.GetAll)
	router.POST("/", h.Create)
	router.GET("/:id", h.Edit)
	router.POST("/update/:id", h.Update)
	router.POST("/delete/:id", h.Delete)

	router.Run(":3000")
}
