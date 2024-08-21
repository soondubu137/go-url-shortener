package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	// custom config

	CypherKey  string
	CacheRedis cache.CacheConf

	// URLMapDB is the config for the url_map table
	URLMapDB struct {
		Host     string
		Port     int
		User     string
		Password string
		DB       string
	}

	// SequenceDB is the config for the sequence table
	SequenceDB struct {
		Host     string
		Port     int
		User     string
		Password string
		DB       string
	}
}
