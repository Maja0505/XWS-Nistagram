package repository

import (
	"XWS-Nistagram/AgentApplication/dto"
	"XWS-Nistagram/AgentApplication/model"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository struct{
	Database *gorm.DB
}

func (repository *ProductRepository) CreateProduct(product *model.Product) error {
	result := repository.Database.Create(product)
	fmt.Println(result.RowsAffected)
	if result.RowsAffected == 0 {
		return fmt.Errorf("ticket not created")
	}
	fmt.Println("Product Created")
	return nil
}

func (repository *ProductRepository) UpdateProduct(updatedProduct *dto.ProductDTO) error {
	id, err := uuid.Parse(updatedProduct.ID)
	if err != nil {
		print(err)
		return err
	}
	result := repository.Database.Model(&model.Product{}).Where("id = ?", id).Updates(map[string]interface{}{"name": updatedProduct.Name,
		"description": updatedProduct.Description,"available_quantity":updatedProduct.AvailableQuantity ,"image":updatedProduct.Image,"price":updatedProduct.Price})
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	fmt.Println("updating")
	return nil
}

func (repository *ProductRepository) DeleteProduct(productId string) error {

	//repository.Database.Delete(&model.Product{},productId)
	//	repository.Database.Where("id LIKE ?", id).Delete(&model.Product{})
	product,err:=repository.FindById(productId)
	if err!=nil{
		fmt.Println(err)
		return err
	}
	repository.Database.Delete(&product)
	return nil
}

func (repository *ProductRepository) FindAll() (*[]model.Product,error){
	var products []model.Product
	if err:= repository.Database.Find(&products).Error;err!=nil{
		fmt.Println(err)
		return nil,err
	}
	return &products,nil
}

func (repository *ProductRepository) FindById(productId string) (*model.Product,error){
	var product model.Product
	if err:=repository.Database.First(&product, "id = ?", productId).Error; err != nil {
		return nil,err
	}
	return &product,nil
}