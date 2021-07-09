package repository

import (
	"XWS-Nistagram/AgentApplication/model"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShoppingCartRepository struct{
	Database *gorm.DB
}

func (repository *ShoppingCartRepository) UpdateOrderQuantity(orderId string,quantity string) error {
	id, err := uuid.Parse(orderId)
	if err != nil {
		print(err)
		return err
	}
	result := repository.Database.Model(&model.Order{}).Where("id = ?", id).Updates(map[string]interface{}{"amount":quantity})
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	fmt.Println("updating")
	return nil
}

func (repository *ShoppingCartRepository) FindById(shoppingCartId string) (*model.ShoppingCart,error){
	var cart model.ShoppingCart
	if err:=repository.Database.Preload("Orders.Product").Preload("Orders", "shopping_cart_id = ?", shoppingCartId).Preload("User").Preload("Address").First(&cart, "id = ?", shoppingCartId).Error; err != nil {
		return nil,err
	}
	return &cart,nil
}

func (repository *ShoppingCartRepository) CreateShoppingCart(shoppingCart *model.ShoppingCart) (*model.ShoppingCart,error) {
	result := repository.Database.Create(&shoppingCart)
	fmt.Println(result.RowsAffected)
	if result.RowsAffected == 0 {
		return nil,fmt.Errorf("ticket not created")
	}
	fmt.Println("Shopping cart Created")
	return shoppingCart,nil
}

func (repository *ShoppingCartRepository) ShoppingCartExists(userId string) bool{
	var cart model.ShoppingCart
	if err:=repository.Database.First(&cart, "user_id = ?", userId).Error; err != gorm.ErrRecordNotFound {
		return true
	}
	return false
}

func (repository *ShoppingCartRepository) DeleteShoppingCart(shoppingCartId string) error {

	//repository.Database.Delete(&model.Product{},productId)
	//	repository.Database.Where("id LIKE ?", id).Delete(&model.Product{})
	shoppingCart,err:=repository.FindById(shoppingCartId)
	if err!=nil{
		fmt.Println(err)
		return err
	}
	repository.Database.Delete(&shoppingCart)
	return nil
}

func (repository *ShoppingCartRepository) AddOrderToShoppingCart(order *model.Order) (error,*model.ShoppingCart) {
	shoppingCart,err:=repository.FindById(order.ShoppingCartID.String())
	 if err!=nil{
		return err,nil
	 }
	repository.Database.Model(&shoppingCart).Association("Orders").Replace(append(shoppingCart.Orders, *order))
	//repository.Database.Model(&shoppingCart).Association("Orders").Append(&order)

	return nil,shoppingCart
}

func (repository *ShoppingCartRepository) DeleteOrderFromShoppingCart(orderId string,shoppingCartId string) (error,*model.ShoppingCart) {
	shoppingCart,err:=repository.FindById(shoppingCartId)
	var order model.Order
	if err:=repository.Database.Preload("Product").First(&order, "id = ?", orderId).Error; err != nil {
		return err,nil
	}
	if err!=nil{
		return err,nil
	}
	repository.Database.Model(&shoppingCart).Association("Orders").Delete(&order)

	return nil,shoppingCart
}

func (repository *ShoppingCartRepository) FindByUser(userId string) (*model.ShoppingCart,error){
	var cart model.ShoppingCart
	if err:=repository.Database.Preload("Orders.Product").Preload("Orders").Preload("User").Preload("Address").First(&cart, "user_id = ?", userId).Error; err != nil {
		return nil,err
	}
	return &cart,nil
}

func (repository *ShoppingCartRepository) CreatePurchase(purchase *model.Purchase) error {
	result := repository.Database.Model(&model.Purchase{}).Create(purchase)
	fmt.Println(result.RowsAffected)
	if result.RowsAffected == 0 {
		return fmt.Errorf("ticket not created")
	}
	fmt.Println("Shopping cart Created")
	return nil
}
