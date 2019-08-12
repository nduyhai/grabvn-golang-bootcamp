package bootcamp

import "github.com/gin-gonic/gin"

type WebContext struct {
	query QueryService
}

func NewWebContext() *WebContext {
	return &WebContext{NewQueryServiceImpl()}
}

func (w *WebContext) QueryUser(ctx *gin.Context) {
	user := ctx.Query("user")
	res := w.query.User(user)

	ctx.JSON(200, res)
}
