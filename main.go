package main

import (
	"fmt"
	"net/http"
	"quoteapp/controller"
	"quoteapp/db"
	"quoteapp/model"
	"quoteapp/routes"
)

func main() {
	db := db.NewDB(
		"192.168.0.101",
		"postgres",
		"password",
		"quoteapp",
		5432,
	)

	db.Migrate()
	conn := db.DB()

	quotemodel := model.NewQuoteModel(conn)
	quotecontroller := controller.NewQuoteController(quotemodel)

	router := routes.NewRoute(quotecontroller)

	fmt.Println("starting api server at http://localhost:8080")
	panic(http.ListenAndServe(":8080", router.Run()))
}
