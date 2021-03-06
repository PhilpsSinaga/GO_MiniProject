package main

import (
	"GO_MINIPROJECT/configuration"
	"GO_MINIPROJECT/controller"
	"GO_MINIPROJECT/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	db, err := configuration.InitDB()
	if err != nil {
		panic(err)
	}
	err = configuration.InitialMigration(db)
	if err != nil {
		panic(err)
	}

	e := echo.New()

	iBukuRepo := repository.NewBukuRepo(db)
	iBukuController := controller.NewBukuController(iBukuRepo)
	// Routes
	e.GET("/buku", iBukuController.GetAllBook)
	e.GET("/buku/:Nama_buku", iBukuController.GetBookByName)
	e.POST("/buku", iBukuController.InsertBook)
	e.DELETE("/buku/:Nama_buku", iBukuController.DeleteBookByName)

	e.POST("/operator", iBukuController.InsertOperator)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
