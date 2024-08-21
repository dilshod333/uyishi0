package pkg

import (
	"database/sql"
	"fmt"
	"conn/config"

	_ "github.com/lib/pq"
)

func ConnectToDBForSuit(cfg config.Config) (*sql.DB, func()) {
	psqlString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)

	connDb, err := sql.Open("postgres", psqlString)
	if err != nil {
		return nil, func() {}
	}

	cleanUpfunc := func() {
		connDb.Close()
	}

	return connDb, cleanUpfunc
}

func ConnectToDB(cfg config.Config) (*sql.DB, error) {
	psqlString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)

	connDb, err := sql.Open("postgres", psqlString)
	if err != nil {
		return nil, err
	}
	err = connDb.Ping()
	fmt.Println("88888888888", err)

	return connDb, nil
}
