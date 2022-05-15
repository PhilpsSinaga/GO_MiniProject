package repository

import (
	"GO_MINIPROJECT/configuration"
	"GO_MINIPROJECT/model"
	"fmt"
)

func CreateNewPetugas(Petugas model.Operator) error {
	err := configuration.DB.Save(&Petugas).Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}
