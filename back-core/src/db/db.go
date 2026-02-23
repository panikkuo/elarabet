package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	instance *pgxpool.Pool
)

func Init(dsn string) {
	var err error
	instance, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		panic("failed to connect to db: " + err.Error())
	}
}

func Get() *pgxpool.Pool {
	if instance == nil {
		panic("db not initialized. Call db.Init() first")
	}
	return instance
}
