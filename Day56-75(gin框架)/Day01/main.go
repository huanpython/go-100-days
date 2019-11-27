/*
@Time : 2019/11/27 0027 下午 5:12
@Author : huanfuan
@File : main
@Software: GoLand
*/

package Day01

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	r.Run()
}
