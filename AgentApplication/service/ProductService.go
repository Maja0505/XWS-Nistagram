package service

import (
	"XWS-Nistagram/AgentApplication/dto"
	"XWS-Nistagram/AgentApplication/mapper"
	"XWS-Nistagram/AgentApplication/model"
	"XWS-Nistagram/AgentApplication/repository"
	"fmt"
)

type  ProductService struct{
	Repository *repository.ProductRepository
}

func (service *ProductService) CreateProduct(product dto.ProductDTO) error {
	fmt.Println("Creating product")
	err := service.Repository.CreateProduct(mapper.MapDTOToProduct(&product))
	if err != nil {
		return err
	}
	return nil

}

func (service *ProductService) UpdateProduct( product dto.ProductDTO) error {
	return service.Repository.UpdateProduct(&product)

}

func (service *ProductService) DeleteProduct(productId string) error {
	return service.Repository.DeleteProduct(productId)

}

func (service *ProductService) FindAll() (*[]model.Product,error) {
	return service.Repository.FindAll()
}

func (service *ProductService) FindById(productId string) (*model.Product,error) {
	return service.Repository.FindById(productId)
}
