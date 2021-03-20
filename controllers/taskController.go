package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/radish-miyazaki/todo-app/models"
)

type TaskHandler struct {
	Db *gorm.DB
}

// 一覧取得
func (h *TaskHandler) GetAll(c *gin.Context) {
	var tds []models.Todo
	h.Db.Find(&tds)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"todos": tds,
	})
}

// 新規作成
func (h *TaskHandler) Create(c *gin.Context) {
	t, _ := c.GetPostForm("title")
	h.Db.Create(&models.Todo{Title: t})
	c.Redirect(http.StatusMovedPermanently, "/")
}

// 編集画面
func (h *TaskHandler) Edit(c *gin.Context) {
	t := models.Todo{}
	id := c.Param("id")
	h.Db.First(&t, id)
	c.HTML(http.StatusOK, "edit.html", gin.H{"todo": t})
}

// 更新
func (h *TaskHandler) Update(c *gin.Context) {
	t := models.Todo{}
	id := c.Param("id")
	tt, _ := c.GetPostForm("title")
	h.Db.First(&t, id)
	t.Title = tt
	h.Db.Save(&t)
	c.Redirect(http.StatusMovedPermanently, "/")
}

// 削除
func (h *TaskHandler) Delete(c *gin.Context) {
	t := models.Todo{}
	id := c.Param("id")
	h.Db.First(&t, id)
	h.Db.Delete(&t)
	c.Redirect(http.StatusMovedPermanently, "/")
}
