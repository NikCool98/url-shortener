package storage

import (
	"database/sql"
	"fmt"

	"github.com/NikCool98/url-short/internal/config"
	_ "github.com/lib/pq"
	_ "modernc.org/sqlite"
)

type Storage struct {
	DB *sql.DB
}

func ConnectDB(cfg *config.Config) (*Storage, error) {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name,
	)
	db, err := sql.Open(cfg.DB.Schema, connString)
	if err != nil {
		return nil, err
	}
	storageDB := &Storage{
		DB: db,
	}
	return storageDB, nil
}
