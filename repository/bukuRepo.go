package repository

import (
	"GO_MINIPROJECT/model"
	"fmt"

	"gorm.io/gorm"
)

type IBukuRepo interface {
	GetAllBook() ([]model.Book, error)
	InsertBook(buku model.Book) error
	// GetBookByName() (model.Book, error)
}

type BukuRepo struct {
	db *gorm.DB
}

func NewBukuRepo(db *gorm.DB) IBukuRepo {
	return &BukuRepo{db}
}

func (r BukuRepo) GetAllBook() ([]model.Book, error) {
	buku := []model.Book{}
	err := r.db.Find(&buku).Error
	if err != nil {
		fmt.Println("Error while GetAllBook", err)
	}
	return buku, err
}

// func (r BukuRepo) GetBookByName() (model.Book, error) {
// 	var NameBook model.Book
// }

func (r BukuRepo) InsertBook(buku model.Book) error {
	err := r.db.Save(&buku).Error
	if err != nil {
		fmt.Println("Error while InsertBook", err)
	}
	return err
}
