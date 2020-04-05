package models

import (
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

func Load(db *pg.DB) error {
	models := []interface{}{(*Url)(nil)}

	for _, model := range models {
		err := db.CreateTable(model, &orm.CreateTableOptions{})

		if err != nil {
			return err
		}
	}
	return nil
}
