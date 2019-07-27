package bootcamp

import "github.com/gin-gonic/gin"

func getAllAccount() gin.Accounts {
	return gin.Accounts{
		"jon": "123",
	}
}
