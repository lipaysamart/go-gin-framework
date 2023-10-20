package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Creating(c *gin.Context) {
	msg := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	c.JSON(http.StatusOK, gin.H{
		"status":  "posted",
		"message": msg,
		"nick":    nick,
	})
}
