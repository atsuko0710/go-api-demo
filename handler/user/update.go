package user

import (
	"go-api-demo/model"
	"go-api-demo/pkg/errno"
	"go-api-demo/util"
	"strconv"

	. "go-api-demo/handler"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

func Update(c *gin.Context) {
	log.Info("User Update functions called.cle", lager.Data{"X-Request-Id": util.GetReqID(c)})

	userId, _ := strconv.Atoi(c.Param("id"))

	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u.Id = uint64(userId)

	// 数据验证
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// 密码加密失败
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	if err := u.Update(); err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, nil)
}
