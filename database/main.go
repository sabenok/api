package database

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/noah-blockchain/noah-explorer-api/core"
)

func Connect(env *core.Environment) *pg.DB {
	options := &pg.Options{
		User:       env.DbUser,
		Password:   env.DbPassword,
		Database:   env.DbName,
		PoolSize:   env.DbPoolSize,
		Addr:       fmt.Sprintf("%s:%d", env.DbHost, env.DbPort),
		MaxRetries: 10,
	}

	db := pg.Connect(options)
	if db == nil {
		panic("Could not connect to database")
	}

	if env.IsDebug {
		db.AddQueryHook(dbLogger{})
	}

	return db
}

func Close(db *pg.DB) {
	err := db.Close()
	if err != nil {
		panic(fmt.Sprintf("Could not close connection to database: %s", err))
	}
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(q *pg.QueryEvent) {}

func (d dbLogger) AfterQuery(q *pg.QueryEvent) {
	fmt.Println(q.FormattedQuery())
}
