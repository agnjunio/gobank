package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rsc.io/quote"
)

func Init() {
	router := NewApi()

	NewAccounts(router)

	router.Run()
}

func NewApi() *gin.Engine {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": quote.Hello(),
		})
	})

	return router
}
