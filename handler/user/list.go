package user

import (
	"go-api-demo/pkg/errno"
	"go-api-demo/service"

	. "go-api-demo/handler"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context)  {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		SendResponse(c, err, nil)
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: int(count),
		UserList: infos,
	})
}