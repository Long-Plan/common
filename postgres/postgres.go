package postgres

import (
	"context"
	"time"

	"github.com/go-pg/pg/v10"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewPostgresClient(cf *Config) (*pg.DB, error) {
	pg := pg.Connect(&pg.Options{
		Addr:     cf.Host + ":" + cf.Port,
		User:     cf.User,
		Password: cf.Password,
		Database: cf.DBName,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := pg.Ping(ctx)
	if err != nil {
		_ = pg.Close()
		return nil, err
	}
	return pg, nil
}
