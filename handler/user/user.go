package user

import (
	"apiserver/model"
	"apiserver/pkg/errno"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *CreateRequest) checkParam() error {
	if r.Username == "" {
		return errno.New(errno.ErrValidation, nil).Add("username is empty.")
	}
	if r.Password == "'" {
		return errno.New(errno.ErrValidation, nil).Add("password is empty.")
	}

	return nil
}

type CreateResponse struct {
	Username string `json:"username"`
}

type ListRequest struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	UserList   []*model.UserInfo `json:"userList"`
}
