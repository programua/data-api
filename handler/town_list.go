package handler

import (
	"data_api.com/data_api/interactor"
	"data_api.com/data_api/usecases"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
)

func TownListHandler(c *gin.Context) {
	db, err := c.Get("DB") // database client
	if err == false {
		log.Fatal("can't get mysql driver")
	}

	var il = interactor.TmpList{}
	usecases.List(il, db.(*sqlx.DB))
	//c.JSON(http.StatusOK, gin.H{"message": "town_list gin ok!"})
}