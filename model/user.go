package model

import (
	"go-api-demo/pkg/auth"

	validator "github.com/go-playground/validator/v10"
)

type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" bindding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" bindding:"required" validate:"min=5,max=128"`
}

func (c *UserModel) TableName() string {
	return "tb_users"
}

// 插入数据
func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error
}

func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}