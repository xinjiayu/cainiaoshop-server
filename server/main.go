package main

import (
	"net/http"

	"cainiaoshop-server/server/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.StaticFS("/v1/RollCommend", http.Dir("/home/hsulei/Documents/appresources/cainiaoshop/RollCommend"))
	router.StaticFS("/v1/Commends", http.Dir("/home/hsulei/Documents/appresources/cainiaoshop/Commend"))
	router.StaticFS("/v1/Hot", http.Dir("/home/hsulei/Documents/appresources/cainiaoshop/Hot"))

	v1 := router.Group("/v1")
	{
		v1.GET("/rollcommend", service.GetRollRecommend)
		v1.GET("/commend", service.GetRecommend)
		v1.GET("/hot", service.GetHotModel)
	}
	router.Run(":8080")
}
