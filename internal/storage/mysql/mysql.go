package mysql

import (
	"database/sql"
	"fmt"

	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.mysql.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}
