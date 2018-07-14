package user

import (
	"github.com/gin-gonic/gin"
	. "github.com/lichunqiang/apiserver/handler"
	"github.com/lichunqiang/apiserver/model"
	"github.com/lichunqiang/apiserver/pkg/errno"
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
