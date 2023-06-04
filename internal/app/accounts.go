package app

import (
	"net/http"

	"github.com/agnjunio/gobank/internal/common/database"
	"github.com/gin-gonic/gin"
)

func GetAccounts(c *gin.Context) {
	accounts, err := database.GetAccounts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"accounts": accounts,
	})
}

// func CreateAccount(c *gin.Context) {
// 	c.Request.Body
// }

func NewAccounts(router *gin.Engine) *gin.RouterGroup {
	accounts := router.Group("accounts")

	accounts.GET("/", GetAccounts)
	// accounts.POST("/", CreateAccount)

	return accounts
}