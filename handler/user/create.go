package user

import (
	"fmt"
	. "go-api-demo/handler"
	"go-api-demo/pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// Create creates a new user account.
func Create(c *gin.Context) {
	var r CreateRequest

	var err error
	if err := c.Bind(&r); err != nil {
		// c.JSON(http.StatusOK, gin.H{"error": errno.ErrBind})
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	admin2 := c.Param("username")
	log.Infof("URL username: %s", admin2)

	desc := c.Query("desc")
	log.Infof("URL key param desc: %s", desc)

	contentType := c.GetHeader("Content-Type")
	log.Infof("Header Content Type: %s", contentType)

	log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)
	if r.Username == "" {
		err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")).Add("This is add message.")
		SendResponse(c, err, nil)
		return
	}

	if errno.IsErrUserNotFound(err) {
		log.Debug("err type is ErrUserNotFound")
	}

	if r.Password == "" {
		err = fmt.Errorf("password is empty")
		SendResponse(c, err, nil)
	}

	rsp := CreateResponse{
		Username: r.Username,
	}
	SendResponse(c, nil, rsp)

	// code, message := errno.DecodeErr(err)
	// c.JSON(http.StatusOK, gin.H{"code": code, "message": message})
}
