package main

import (
	"fmt"
	"quoteapp/db"
	"quoteapp/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	conn := db.NewDB().DB()

	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	fmt.Println("starting api server at http://localhost:8080")
	err := routes.NewRoute(conn, server).Run()
	if err != nil {
		panic(err)
	}
}
