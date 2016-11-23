package service

import (
	"strconv"

	"cainiaoshop-server/db"

	"github.com/gin-gonic/gin"
)

//GetHotModel 得到热门商品
func GetHotModel(c *gin.Context) {
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return
	}
	hots, err := db.GetHotByLimit(page)
	if err != nil {
		return
	}
	c.JSON(200, hots)
}
