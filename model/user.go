package model

import (
	"github.com/lichunqiang/apiserver/pkg/auth"
	"gopkg.in/go-playground/validator.v9"
)

type UserModel struct {
	BaseModel

	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

func (u *UserModel) TableName() string {
	return "tb_users"
}

//create
func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error
}

//update
func (u *UserModel) Update() error {
	return DB.Self.Update(&u).Error
}

//validate the password
func (u *UserModel) ValidatePassword(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

//encrypt password
func (u *UserModel) EncryptPassword() (err error) {
	u.Password, err = auth.Encrypt(u.Password)

	return err
}

//validate the model
func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

//get user by username
func GetUser(username string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.First(&u, "username = ?", username)

	return u, d.Error
}

func GetUserById(id uint64) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("id = ?", id).First(&u)

	return u, d.Error
}

//delete user by id
func DeleteUser(id uint64) error {
	user := &UserModel{}
	user.BaseModel.Id = id

	return DB.Self.Delete(&user).Error
}

//func ListUser(username string, offset, limit int) ([]*UserModel, uint64, error) {
//	if limit == 0 || limit > constvar.MaxPageLimit {
//		limit = constvar.DefaultLimit
//	}
//
//users := make([]*UserModel)
//
//}
