package main

import (
	"net/http"

	"cainiaoshop-server/db"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.StaticFS("/v1/RollCommend", http.Dir("/home/hsulei/Documents/appresources/cainiaoshop/RollCommend"))
	router.StaticFS("/v1/Commends", http.Dir("/home/hsulei/Documents/appresources/cainiaoshop/Commend"))
	v1 := router.Group("/v1")
	{
		v1.GET("/rollcommend", getRollRecommend)
		v1.GET("/commend", getRecommend)
	}
	router.Run(":8080")
}

//getRollCommend 获取滚动推荐数据
func getRollRecommend(c *gin.Context) {
	recommends, err := db.GetRollRecommend()
	if err != nil {
		c.JSON(400, gin.H{
			"err": 1,
		})
		return
	}
	c.JSON(200, recommends)
}

//getRecommend 获取首页推荐内容
func getRecommend(c *gin.Context) {

}
