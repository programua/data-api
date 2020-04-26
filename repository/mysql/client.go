package mysql

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

//connect mysql driver
func NewMysqlClient() gin.HandlerFunc {
	db, err := sqlx.Connect("mysql", "root:sherlock@tcp(localhost:3306)/data_api")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("database connect success!")

	return func(c *gin.Context){
		c.Set("DB", db)
		c.Next()
	}
}