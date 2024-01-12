package controller

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

// HelloController
//
//	@Description: HelloController
//	@param c	*gin.Context
func HelloController(c *gin.Context) {
	tmpl, err := template.ParseFiles("./web/hello.html")
	if err != nil {
		c.JSON(200, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}
	err = tmpl.Execute(c.Writer, "如果你看到这条消息，说明你的网页服务器已经正常运行了")
	if err != nil {
		c.JSON(200, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}
}
