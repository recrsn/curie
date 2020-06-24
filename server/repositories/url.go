package repositories

import (
	"github.com/go-pg/pg/v9"

	"github.com/curie/server/models"
)

type UrlRepository struct {
	db *pg.DB
}

func (u *UrlRepository) GetAll(limit, offset int) ([]models.Url, error) {
	var urls []models.Url

	err := u.db.Model(&urls).
		Order("id ASC").Limit(limit).Offset(offset).Select()

	if err != nil {
		return nil, err
	}

	return urls, nil
}
