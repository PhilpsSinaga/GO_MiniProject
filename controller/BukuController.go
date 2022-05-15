package controller

import (
	"GO_MINIPROJECT/model"
	"GO_MINIPROJECT/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IBukuController interface {
	GetAllBook(c echo.Context) error
	InsertBook(c echo.Context) error
}

type BukuController struct {
	iBukuRepo repository.IBukuRepo
}

func NewBukuController(iBukuRepo repository.IBukuRepo) IBukuController {
	return &BukuController{iBukuRepo: iBukuRepo}
}

var appJSON = "application/json"

func (ctrl BukuController) GetAllBook(c echo.Context) error {
	book, err := ctrl.iBukuRepo.GetAllBook()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message":     "Berhasil Melihat Semua Data Buku",
		"Daftar Buku": book,
	})
}

func (ctrl BukuController) InsertBook(c echo.Context) error {
	NamaBuku := model.Book{}
	c.Bind(&NamaBuku)
	err := ctrl.iBukuRepo.InsertBook(NamaBuku)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Succes insert data Buku",
	})

}
