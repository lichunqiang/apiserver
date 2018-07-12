package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lichunqiang/apiserver/pkg/errno"
	"github.com/lexkong/log"
	"fmt"
	. "github.com/lichunqiang/apiserver/handler"
)

func Create(c *gin.Context) {
	var r struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if !c.Bind(&r) {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	log.Debugf("username is: [%s], password is: [%s]", r.Username, r.Password)

	var err error

	if r.Username == "" {
		err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")).Add(" add message")
		log.Errorf(err, "Get an error")
	}

	if errno.IsErrUserNotFound(err) {
		log.Debug("err type is ErrUserNotFound")
	}

	if r.Password == "" {
		err = fmt.Errorf("password is empty")
	}

	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	b := map[string]string {
		"name": "lighjt",
	}

	SendResponse(c, nil, b)
}

func Get(c *gin.Context)  {
	
}

func Delete(c *gin.Context)  {
	
}

func Update(c *gin.Context)  {
	
}

func List(c *gin.Context) {

}