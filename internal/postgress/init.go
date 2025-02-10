package postgress

import (
	"context"
	"fmt"

	"avito-shop-service/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgre struct {
	db *sqlx.DB
}

func NewConnection(ctx context.Context, cfg config.DB) (*sqlx.DB, error) {

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Name, cfg.Password)

	db, err := sqlx.ConnectContext(ctx, "postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func New(conn *sqlx.DB) *Postgre {
	return &Postgre{db: conn}
}
