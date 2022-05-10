package repository

import (
	"GO_MINIPROJECT/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type IUserRepo interface {
// }

// type UserRepo struct {
// 	db *gorm.DB
// }

func Connect() {
	dsn := "root:password@tcp(127.0.0.1:3306)/buku?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var listbuku []model.Buku
	err = db.Raw("Select * From bukus;").Scan(&listbuku).Error
	fmt.Println(listbuku)
	fmt.Println(err)
}
