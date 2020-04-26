package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

//get town list
func Get(db *sqlx.DB) ([]TownListDS, error) {
	fmt.Println("repository Get running!")

	var tl []TownListDS
	err := db.Select(&tl, "SELECT * FROM town_list")
	if err != nil {
		return nil, err
	}

	return tl, nil
}