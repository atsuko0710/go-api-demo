package user

import (
	"go-api-demo/model"
	"go-api-demo/pkg/errno"

	. "go-api-demo/handler"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context)  {
	username := c.Param("username")
	
	user, err := model.GetUser(username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	SendResponse(c, nil, user)
}