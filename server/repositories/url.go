package repositories

import (
	"github.com/go-pg/pg/v9"
)

type UrlRepository struct {
	db *pg.DB
}