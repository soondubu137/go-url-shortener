package svc

import (
	"fmt"
	"go-url-shortener/internal/config"
	"go-url-shortener/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	URLMapModel model.UrlMapModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		c.URLMapDB.User,
		c.URLMapDB.Password,
		c.URLMapDB.Host,
		c.URLMapDB.Port,
		c.URLMapDB.DB,
	))
	return &ServiceContext{
		Config:      c,
		URLMapModel: model.NewUrlMapModel(conn),
	}
}
