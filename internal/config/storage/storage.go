package storage

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/NikCool98/url-short/internal/config"
	"github.com/lib/pq"
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
		log.Fatalf("Ошибка подключения к БД %v", err)
	}
	storageDB := &Storage{
		DB: db,
	}
	return storageDB, nil
}

func (s *Storage) SaveURL(urlToSave string, alias string) (int64, error) {
	const op = "storage.postgres.SaveURL"
	var lastInsertId int64

	query := `INSERT INTO url (url, alias) VALUES ($1, $2) RETURNING id;`

	err := s.DB.QueryRow(query, urlToSave, alias).Scan(&lastInsertId)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return 0, fmt.Errorf("Эти данные уже были внесены в таблицу ранее: %w", err)
		}
	}
	return lastInsertId, nil
}
