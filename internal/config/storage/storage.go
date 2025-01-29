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
		log.Fatalf("DB connection error %v", err)
	}
	storageDB := &Storage{
		DB: db,
	}
	return storageDB, nil
}

func (s *Storage) SaveURL(urlToSave string, alias string) (int64, error) {
	var lastInsertId int64

	query := `INSERT INTO url (url, alias) VALUES ($1, $2) RETURNING id;`

	err := s.DB.QueryRow(query, urlToSave, alias).Scan(&lastInsertId)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return 0, fmt.Errorf("This data has already been entered into the table earlier: %w", err)
		}
	}
	return lastInsertId, nil
}

func (s *Storage) GetUrl(alias string) (string, error) {
	var resUrl string
	query := "SELECT url FROM url WHERE alias = $1"
	err := s.DB.QueryRow(query, alias).Scan(&resUrl)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("Url not found: %w", err)
		}
		return "", err
	}
	return resUrl, nil
}

func (s *Storage) DeleteUrl(alias string) error {
	query := "DELETE FROM url WHERE alias = $1"
	res, err := s.DB.Exec(query, alias)
	if err != nil {
		return fmt.Errorf("Fault to delete url: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("Url with this alias not found: %w", err)
	}
	return nil
}
