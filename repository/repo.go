package repository

import (
	"GO_MINIPROJECT/model"
	"fmt"

	"gorm.io/gorm"
)

type IBukuRepo interface {
	GetAllBook() ([]model.Book, error)
	InsertBook(buku model.Book) error
	GetBookByName(Nama_buku string) (model.Book, error)
	DeleteBookByName(Nama_buku string) error
	InsertOperator(petugas model.Operator) error
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

func (r BukuRepo) GetBookByName(Nama_buku string) (model.Book, error) {
	var NamaBuku model.Book
	err := r.db.Where("Nama_buku = ?", Nama_buku).First(&NamaBuku).Error
	if err != nil {
		fmt.Println("Error while Get the book : ", err)
	}

	return NamaBuku, err
}

func (r BukuRepo) InsertBook(buku model.Book) error {
	err := r.db.Create(&buku).Error
	if err != nil {
		fmt.Println("Error while InsertBook", err)
	}
	return err
}

func (r BukuRepo) DeleteBookByName(Nama_buku string) error {
	err := r.db.Delete(&model.Book{}, "Nama_buku = ?", Nama_buku).Error
	if err != nil {
		fmt.Println("Error While Delete Book", err)
	}
	return err
}

func (r BukuRepo) InsertOperator(petugas model.Operator) error {
	err := r.db.Create(&petugas).Error
	if err != nil {
		fmt.Println("Error while InserOperator", err)
	}
	return err
}
