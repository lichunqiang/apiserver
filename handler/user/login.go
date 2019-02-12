package user

import (
	"github.com/gin-gonic/gin"
	. "apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"
	"apiserver/pkg/token"
)

func Login(c *gin.Context) {
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, err, nil)
		return
	}

	d, err := model.GetUser(u.Username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	if err := d.ValidatePassword(u.Password); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	t, err := token.Sign(c, token.Context{
		ID:       d.Id,
		Username: d.Username,
	}, "")

	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}

	SendResponse(c, nil, model.Token{Token: t})
}
