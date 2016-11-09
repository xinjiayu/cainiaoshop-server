package server

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	v1 := router.Group("v1")
	{
		v1.GET()
	}
	router.Run("8080")
}
