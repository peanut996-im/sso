package handler

import (
	"net/http"

	"framework/api/model"
	"framework/db"
	"framework/logger"
	"framework/net"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	login := newLogin()
	err := c.ShouldBind(login)
	if err != nil {
		c.JSON(http.StatusOK, net.NewHttpInnerErrorResp(err))
		return
	}
	c.JSON(200, net.NewSuccessResponse(nil))
}

func SignOut(c *gin.Context) {
	c.String(200, "this is logout post")
}

func SignUp(c *gin.Context) {
	login := newLogin()
	mongo := db.GetLastMongoClient()
	// rds := db.GetLastRedisClient()

	err := c.ShouldBind(login)
	if err != nil {
		c.JSON(http.StatusOK, net.NewHttpInnerErrorResp(err))
		return
	}
	user := model.NewUser(login.Account, login.Password)
	res, err := mongo.InsertOne("user", user)
	if err != nil {
		logger.Error("mongo insert user err: %v", err)
		c.JSON(http.StatusOK, net.NewHttpInnerErrorResp(err))
		return
	}
	logger.Debug("mongo insert user success", res.InsertedID)
	c.JSON(200, net.NewSuccessResponse(nil))
}
