package service

import (
	"cainiaoshop-server/db"

	"github.com/gin-gonic/gin"
)

//GetRollRecommend 获取滚动推荐数据
func GetRollRecommend(c *gin.Context) {
	recommends, err := db.GetRollRecommend()
	if err != nil {
		c.JSON(400, gin.H{
			"err": 1,
		})
		return
	}
	c.JSON(200, recommends)
}

//GetRecommend 获取首页推荐内容
func GetRecommend(c *gin.Context) {

}
