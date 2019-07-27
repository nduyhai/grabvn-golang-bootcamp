package bootcamp

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type WebContext struct {
	DB *gorm.DB
}

func (w *WebContext) createTodo(ctx *gin.Context) {
	createTodo(ctx, w)
}

func (w *WebContext) getAllTodo(ctx *gin.Context) {
	getAllTodo(ctx, w)
}

func (w *WebContext) getTodoById(ctx *gin.Context) {
	getTodoById(ctx, w)
}
