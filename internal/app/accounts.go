package app

import (
	"net/http"

	"github.com/agnjunio/gobank/internal/common/database"
	"github.com/agnjunio/gobank/internal/common/models"
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

type AddAccountRequestBody struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateAccount(c *gin.Context) {
	body := AddAccountRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var account models.Account
	account.Name = body.Name
	account.Email = body.Email
	account.Balance = 0

	if err := database.AddAccount(&account); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, &account)
}

func NewAccounts(router *gin.Engine) *gin.RouterGroup {
	accounts := router.Group("accounts")

	accounts.GET("/", GetAccounts)
	accounts.POST("/", CreateAccount)

	return accounts
}
