package repository

//
//list data structure
//

//town list data structure
type TownListDS struct {
	Id       int    `db:"id"`
	TownName string `db:"town"`
	TownRuby string `db:"town_ruby"`
}
