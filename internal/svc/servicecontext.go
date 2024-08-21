package svc

import (
	"fmt"
	"go-url-shortener/internal/config"
	"go-url-shortener/internal/utils/idgenerator"
	"go-url-shortener/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	URLMapModel model.UrlMapModel
	IdGenerator idgenerator.IdGenerator
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		c.URLMapDB.User,
		c.URLMapDB.Password,
		c.URLMapDB.Host,
		c.URLMapDB.Port,
		c.URLMapDB.DB,
	))

	return &ServiceContext{
		Config:      c,
		URLMapModel: model.NewUrlMapModel(conn),
		// can be replaced with other IdGenerator implementations
		IdGenerator: idgenerator.NewDefaultIdGenerator(c),
	}
}
