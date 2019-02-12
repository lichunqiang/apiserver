package user

import (
	"github.com/gin-gonic/gin"
	"apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"
	"strconv"
)

// DELETE /v1/users/1212
func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := model.DeleteUser(uint64(id)); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
	}

	handler.SendResponse(c, nil, nil)
}
