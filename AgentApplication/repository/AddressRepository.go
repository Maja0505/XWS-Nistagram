package repository

import (
	"XWS-Nistagram/AgentApplication/model"
	"fmt"
	"gorm.io/gorm"
)

type AddressRepository struct{
	Database *gorm.DB
}

func (repository *AddressRepository) CreateAddress(address *model.Address) (error,model.Address) {
	result := repository.Database.Create(&address)
	fmt.Println(result.RowsAffected)
	if result.RowsAffected == 0 {
		return fmt.Errorf("address not created"),*address
	}
	fmt.Println("Product Created")
	return nil,*address
}
