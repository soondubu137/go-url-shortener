package svc

import (
	"fmt"
	"go-url-shortener/internal/config"
	"go-url-shortener/internal/utils/idgenerator"
	"go-url-shortener/model"

	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	URLMapModel model.UrlMapModel
	IdGenerator idgenerator.IdGenerator
	Filter      *bloom.Filter
}

func NewServiceContext(c config.Config) *ServiceContext {
	// initialize the bloom filter
	store := redis.New(c.CacheRedis[0].Host, func(r *redis.Redis) {
		r.Type = redis.NodeType
	})
	filter := bloom.New(store, "url_filter", 20*1024)

	// create a connection to the URL map database
	conn := sqlx.NewMysql(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		c.URLMapDB.User,
		c.URLMapDB.Password,
		c.URLMapDB.Host,
		c.URLMapDB.Port,
		c.URLMapDB.DB,
	))

	return &ServiceContext{
		Config:      c,
		URLMapModel: model.NewUrlMapModel(conn, c.CacheRedis),
		// can be replaced with other IdGenerator implementations
		IdGenerator: idgenerator.NewDefaultIdGenerator(c),
		Filter:      filter,
	}
}
