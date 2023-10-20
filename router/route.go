package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lipaysamart/go-gin-framework/controller"
)

func Route() {
	router := gin.Default()

	router.GET("/get", controller.Getting)
	router.POST("/post", controller.Creating)

	err := router.Run(":8090")
	if err != nil {
		return
	}
}
