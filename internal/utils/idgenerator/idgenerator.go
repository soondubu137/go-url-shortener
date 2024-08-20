package idgenerator

import (
	"context"
	"fmt"
	"go-url-shortener/internal/config"
	"go-url-shortener/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

// IdGenerator is an interface with a single method Generate.
// Generate generates a new id based on the stub provided.
type IdGenerator interface {
	Generate(context.Context, string) (int64, error)
}

type defaultIdGenerator struct {
	sequenceModel model.SequenceModel
}

func (g *defaultIdGenerator) Generate(ctx context.Context, stub string) (int64, error) {
	return g.sequenceModel.Replace(ctx, stub)
}

// NewDefaultIdGenerator returns a new IdGenerator implemented with MySQL.
// Corresponding config for MySQL connection should be present in the config file.
func NewDefaultIdGenerator(c config.Config) IdGenerator {
	conn := sqlx.NewMysql(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		c.SequenceDB.User,
		c.SequenceDB.Password,
		c.SequenceDB.Host,
		c.SequenceDB.Port,
		c.SequenceDB.DB,
	))
	return &defaultIdGenerator{
		sequenceModel: model.NewSequenceModel(conn),
	}
}
