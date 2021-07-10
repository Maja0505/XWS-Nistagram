package handler

import (
	"XWS-Nistagram/AgentApplication/dto"
	"XWS-Nistagram/AgentApplication/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ShoppingCartHandler struct{
	Service *service.ShoppingCartService
}

func (handler *ShoppingCartHandler) CreateShoppingCart(c *gin.Context){
	var shoppingCart dto.ShoppingCartDTO
	if err := c.ShouldBindJSON(&shoppingCart); err != nil{
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	cart,err := handler.Service.CreateShoppingCart(shoppingCart)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK,cart)

}

func (handler *ShoppingCartHandler) UpdateOrderQuantity(c *gin.Context){
	orderID := c.Query("orderId")
	amount := c.Query("amount")

	if orderID=="" || amount==""{
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	err := handler.Service.UpdateOrderQuantity(orderID,amount)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK,"Product successfully updated")

}

func (handler *ShoppingCartHandler) DeleteShoppingCart(c *gin.Context){
	var id string
	id = c.Query("id")
	if id==""{
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	err := handler.Service.DeleteShoppingCart(id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK,"Product successfully id")

}


func (handler *ShoppingCartHandler) FindById(c *gin.Context) {
	shoppingCart,err := handler.Service.FindById(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK,shoppingCart)

}

func (handler *ShoppingCartHandler) AddOrderToShoppingCart(c *gin.Context) {
	var order dto.OrderDTO
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	err, shoppingCart := handler.Service.AddOrderToShoppingCart(&order)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, shoppingCart)
}

func (handler *ShoppingCartHandler) DeleteOrderFromShoppingCart(c *gin.Context) {
	orderId:=c.Query("orderId")
	shoppingCartId:=c.Query("shoppingCartId")
	if orderId == "" || shoppingCartId == ""{
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	err, shoppingCart := handler.Service.DeleteOrderFromShoppingCart(orderId,shoppingCartId)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, shoppingCart)
}

func (handler *ShoppingCartHandler) FindByUser(c *gin.Context) {
	shoppingCart,err := handler.Service.FindByUser(c.Query("userId"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK,shoppingCart)

}

func (handler *ShoppingCartHandler) CreatePurchase(c *gin.Context){
	var purchase dto.PurchaseDTO
	if err := c.ShouldBindJSON(&purchase); err != nil{
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	err := handler.Service.CreatePurchase(purchase)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK,"Product successfully created")

}

func (handler *ShoppingCartHandler) EmptyShoppingCart(c *gin.Context){
	var shoppingCartID =c.Query("shoppingCartId")
	if shoppingCartID ==""{
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	err,shoppingCart := handler.Service.EmptyShoppingCart(shoppingCartID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK,shoppingCart)

}

