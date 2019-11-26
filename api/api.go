package api

import (
	"test/dto"
	"test/logic"

	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		var customerDTO = dto.Customer{}
		var resp = logic.GetAll(customerDTO)
		c.JSON(200, resp)
	})
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
