package user

import (
	"go-api-demo/model"
	"go-api-demo/pkg/auth"
	"go-api-demo/pkg/errno"
	"go-api-demo/pkg/token"

	. "go-api-demo/handler"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context)  {
	var u model.UserModel

	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	d, err := model.GetUser(u.Username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	// 比较数据库密码和传输来的密码
	if err := auth.Compare(d.Password, u.Password); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	t, err := token.Sign(c, token.Context{
		ID: d.Id,
		Username: d.Username,
	}, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}

	SendResponse(c, nil, model.Token{Token: t})	
}