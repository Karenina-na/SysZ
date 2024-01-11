package controller

import (
	"SysZ/src/config"
	"SysZ/src/controller/util"
	"SysZ/src/entity"
	"SysZ/src/exception"
	"SysZ/src/util/encryption"
	"SysZ/src/util/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RootController
//
//	@Description: RootController
//	@param c	*gin.Context
func RootController(c *gin.Context) {
	root := entity.NewRootModel()
	err := c.BindJSON(&root)
	if err != nil {
		exception.HandleException(exception.NewUserError("GetServersController", "参数绑定错误-"+err.Error()))
		c.JSON(http.StatusOK, entity.NewFalseResult("false", "参数绑定错误-"+err.Error()))
		return
	}
	account := encryption.Base64Decode(root.Account)
	password := encryption.Base64Decode(root.Password)
	if (account == config.Root.RootAccount) && (password == config.Root.RootPassword) {
		t, err := token.GenerateToken(config.Root.RootAccount, config.Root.RootPassword)
		if err != nil {
			util.Handle(err, c)
			return
		}
		c.JSON(http.StatusOK, entity.NewSuccessResult(struct {
			Token string `json:"token"`
		}{Token: t}))
		return
	}
	c.JSON(http.StatusOK, entity.NewFalseResult("false", "密码错误"))
}
