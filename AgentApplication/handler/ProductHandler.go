package handler

import (
	"XWS-Nistagram/AgentApplication/dto"
	"XWS-Nistagram/AgentApplication/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductHandler struct{
	Service *service.ProductService
}

func (handler *ProductHandler) CreateProduct(c *gin.Context){
	var product dto.ProductDTO
	if err := c.ShouldBindJSON(&product); err != nil{
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	err := handler.Service.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK,"Product successfully created")

}

func (handler *ProductHandler) UpdateProduct(c *gin.Context){
	var product dto.ProductDTO
	if err := c.ShouldBindJSON(&product); err != nil{
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	err := handler.Service.UpdateProduct(product)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK,"Product successfully updated")

}

func (handler *ProductHandler) DeleteProduct(c *gin.Context){
	var id string
	id = c.Query("id")
	if id==""{
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	err := handler.Service.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK,"Product successfully id")

}

func (handler *ProductHandler) FindAll(c *gin.Context) {
	products,err := handler.Service.FindAll()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK,*products)

}

func (handler *ProductHandler) FindById(c *gin.Context) {
	product,err := handler.Service.FindById(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK,product)

}

func (handler *ProductHandler) SaveProduct(c *gin.Context){
	var product dto.ProductDTO
	if err := c.ShouldBindJSON(&product); err != nil{
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	err := handler.Service.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK,"Product successfully created")

}