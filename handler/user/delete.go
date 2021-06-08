package user

import (
	"go-api-demo/model"
	"go-api-demo/pkg/errno"
	"strconv"

	. "go-api-demo/handler"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context)  {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, nil)
}