package main

import (
	"guizizhan/config"
	"guizizhan/pkg/mysql"
	"guizizhan/router"
)

func main() {
	config.InitConfig()
	db, err := mysql.InitMySQL()
	if err != nil {
		println("err:", err)
	}

	e := router.RouterInit(db)
	e.Run("10.131.122.150:8080")
}
