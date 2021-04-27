package handler

import (
	"fmt"
	"net/http"
	"sso/model"

	"framework/net"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	login := model.NewLogin()
	// err := c.BindJSON(login)
	err := c.ShouldBind(login)
	if err != nil {
		c.JSON(http.StatusOK, net.BaseRepsonse{
			Code:    net.ERROR_HTTP_INNER_ERROR,
			Data:    struct{}{},
			Message: fmt.Sprintf(net.ErrorCodeToString(net.ERROR_HTTP_INNER_ERROR), err),
		})
		return
	}
	c.JSON(200, net.BaseRepsonse{
		Code:    net.ERROR_CODE_OK,
		Data:    nil,
		Message: net.ErrorCodeToString(net.ERROR_CODE_OK),
	})
}

func SignOut(c *gin.Context) {
	c.String(200, "this is logout post")
}

func SignUp(c *gin.Context) {
	c.String(200, "this is register")
}
