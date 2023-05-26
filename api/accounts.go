package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAccounts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"accounts": []int{},
	})
}

func NewAccounts(router *gin.Engine) *gin.RouterGroup {
	accounts := router.Group("accounts")

	accounts.GET("/", GetAccounts)

	return accounts
}
