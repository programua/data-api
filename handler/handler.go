package handler

import (
	"data_api.com/data_api/repository/mysql"
	"fmt"
	"github.com/gin-gonic/gin"
)

// エラーハンドリングのどっかに defer db.Close()置きたい。そもそも必要なのか調べる
func MakeRouter() *gin.Engine {
	fmt.Println("MakeRouter is run")

	r := gin.Default()
	r.Use(mysql.NewMysqlClient())
	v1 := r.Group("/v1")
	{
		v1.POST("/top", TopHandler)
		v1.POST("/list", TownListHandler)
	}

	return r
}
