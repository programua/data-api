package interactor

import (
	"data_api.com/data_api/repository"
	"github.com/jmoiron/sqlx"
)

//
//town list interact
//  - call repository list.go
//  - get town name list or error
//  - return town name list or error
//

type TmpList struct {}

//call repository & return town list
func (t TmpList) List(db *sqlx.DB) ([]repository.TownListDS, error) {
	townList, err := repository.Get(db)

	if err != nil {
		return nil, err
	}

	return townList, err
}