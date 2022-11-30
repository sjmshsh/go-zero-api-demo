package svc

import (
	config2 "Go-Zero/user-api/internal/config"
	"Go-Zero/user-api/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config2.Config
	UserModel model.UserModel
}

func NewServiceContext(c config2.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
