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
	GetBookByName(c echo.Context) error
	DeleteBookByName(c echo.Context) error
	InsertOperator(c echo.Context) error
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
		"message": "Succes insert Book",
	})

}

func (ctrl BukuController) GetBookByName(c echo.Context) error {
	stringName := c.Param("Nama_buku")
	NamaBuku, err := ctrl.iBukuRepo.GetBookByName(stringName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, NamaBuku)
}

func (ctrl BukuController) DeleteBookByName(c echo.Context) error {
	stringName := c.Param("Nama_buku")
	err := ctrl.iBukuRepo.DeleteBookByName(stringName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes delete Book with Name '" + stringName + "'",
	})
}

func (ctrl BukuController) InsertOperator(c echo.Context) error {
	OperatorName := model.Operator{}
	c.Bind(&OperatorName)
	err := ctrl.iBukuRepo.InsertOperator(OperatorName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Succes insert Operator",
	})
}
