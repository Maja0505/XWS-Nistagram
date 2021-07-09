package service

import (
	"XWS-Nistagram/AgentApplication/dto"
	"XWS-Nistagram/AgentApplication/mapper"
	"XWS-Nistagram/AgentApplication/model"
	"XWS-Nistagram/AgentApplication/repository"
)

type  ShoppingCartService struct{
	Repository *repository.ShoppingCartRepository
	ProductRepository *repository.ProductRepository
	AddressRepository *repository.AddressRepository
}

func (service *ShoppingCartService) CreateShoppingCart(shoppingCart dto.ShoppingCartDTO) (*model.ShoppingCart,error) {
	cart,err := service.Repository.CreateShoppingCart(mapper.MapDTOToShoppingCart(&shoppingCart))
	if err != nil {
		return nil,err
	}
	return cart,nil

}

func (service *ShoppingCartService) UpdateOrderQuantity(orderId string,quantity string) error {

	return service.Repository.UpdateOrderQuantity(orderId,quantity)

}

func (service *ShoppingCartService) DeleteShoppingCart(shoppingCartId string) error {
	return service.Repository.DeleteShoppingCart(shoppingCartId)

}



func (service *ShoppingCartService) FindById(shoppingCartId string) (*model.ShoppingCart,error) {
	return service.Repository.FindById(shoppingCartId)
}

func (service *ShoppingCartService) AddOrderToShoppingCart(order *dto.OrderDTO) (error,*model.ShoppingCart) {
	dto:=mapper.MapDTOToOrder(order)
	err,shoppingCart := service.Repository.AddOrderToShoppingCart(dto)
	if err != nil {
		return err,nil
	}
	return nil,shoppingCart

}

func (service *ShoppingCartService) DeleteOrderFromShoppingCart(orderId string,shoppingCartId string) (error,*model.ShoppingCart) {
 	return service.Repository.DeleteOrderFromShoppingCart(orderId,shoppingCartId)
}

func (service *ShoppingCartService) FindByUser(userId string) (*model.ShoppingCart,error) {
	return service.Repository.FindByUser(userId)


}
func (service *ShoppingCartService) CreatePurchase(purchase dto.PurchaseDTO) error {
	err := service.Repository.CreatePurchase(mapper.MapDTOToPurchase(&purchase))
	if err != nil {
		return err
	}
	return nil

}





