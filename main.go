package main

import (
	"fmt"
	"net/http"
	"quoteapp/controller"
	"quoteapp/db"
	"quoteapp/repository"
	"quoteapp/routes"
)

func main() {
	db := db.NewDB(
		"192.168.0.104",
		"postgres",
		"password",
		"quoteapp",
		5432,
	)

	db.Migrate()
	conn := db.DB()

	repository := repository.NewQuoteRepository(conn)
	controller := controller.NewQuoteController(repository)

	router := routes.NewRoute(controller)

	fmt.Println("starting api server at http://localhost:8080")
	http.ListenAndServe(":8080", router.Run())
}
