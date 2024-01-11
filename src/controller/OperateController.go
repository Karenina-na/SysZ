package controller

import (
	"SysZ/src/controller/util"
	"SysZ/src/entity"
	"SysZ/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetStatusController
//
//	@Description: 获取服务器状态
//	@param c	*gin.Context
func GetStatusController(c *gin.Context) {
	computer, err := service.GetServerStatus()
	if err != nil {
		util.Handle(err, c)
		return
	}
	c.JSON(http.StatusOK, entity.NewSuccessResult(computer))
}
