
package user

import (
	. "go-api-demo/handler"
	"go-api-demo/model"
	"go-api-demo/pkg/errno"
	"go-api-demo/util"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// Create creates a new user account.
func Create(c *gin.Context) {
	log.Info("User Create functions called.cle", lager.Data{"X-Request-Id": util.GetReqID(c)})

	var r CreateRequest

	if err := c.Bind(&r); err != nil {
		// c.JSON(http.StatusOK, gin.H{"error": errno.ErrBind})
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// admin2 := c.Param("username")
	// log.Infof("URL username: %s", admin2)

	// desc := c.Query("desc")
	// log.Infof("URL key param desc: %s", desc)

	// contentType := c.GetHeader("Content-Type")
	// log.Infof("Header Content Type: %s", contentType)

	// log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}
	
	// 数据验证
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// 密码加密失败
	if err:= u.Encrypt(); err != nil{
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// 创建用户
	if err:= u.Create(); err != nil{
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	// if r.Username == "" {
	// 	err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")).Add("This is add message.")
	// 	SendResponse(c, err, nil)
	// 	return
	// }

	// if errno.IsErrUserNotFound(err) {
	// 	log.Debug("err type is ErrUserNotFound")
	// }

	// if r.Password == "" {
	// 	err = fmt.Errorf("password is empty")
	// 	SendResponse(c, err, nil)
	// }

	rsp := CreateResponse{
		Username: r.Username,
	}
	SendResponse(c, nil, rsp)

	// code, message := errno.DecodeErr(err)
	// c.JSON(http.StatusOK, gin.H{"code": code, "message": message})
}
