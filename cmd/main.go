package main

import (
	"gin-mall/repository/db"
	"gin-mall/routes"
)

func main() {
	err := db.Init()
	if err != nil {
		panic(err)
	}
	r := routes.Routes()
	r.Run(":8080")
}
