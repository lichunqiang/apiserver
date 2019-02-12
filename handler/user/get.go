package user

import (
	"github.com/gin-gonic/gin"
	. "apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"
	"strconv"
)

func Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	u, err := model.GetUserById(uint64(id))

	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, u)
}
