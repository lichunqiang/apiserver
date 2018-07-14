package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	. "github.com/lichunqiang/apiserver/handler"
	"github.com/lichunqiang/apiserver/model"
	"github.com/lichunqiang/apiserver/pkg/errno"
	"github.com/lichunqiang/apiserver/util"
)

func Create(c *gin.Context) {
	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqId(c)})

	var r = CreateRequest{}

	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	if err := r.checkParam(); err != nil {
		SendResponse(c, err, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}
	//validate
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	//encrypt password
	if err := u.EncryptPassword(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	//insert
	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Username: u.Username,
	}

	SendResponse(c, nil, rsp)
}
