package pgsql

import (
	"github.com/inspectorvitya/x-technology-test/internal/config"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
)

type StoragePgSql struct {
	db *sqlx.DB
}

func New(cfg config.Config) (*StoragePgSql, error) {
	db, err := sqlx.Open("pgx", cfg.DbURL)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	s := &StoragePgSql{
		db: db,
	}
	return s, nil
}

func (s *StoragePgSql) Close() {
	if err := s.db.Close(); err != nil {
		log.Fatalln(err)
	}
}
