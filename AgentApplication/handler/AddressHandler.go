package handler

import (
	"XWS-Nistagram/AgentApplication/dto"
	"XWS-Nistagram/AgentApplication/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddressHandler struct{
	Service *service.AddressService
}

func (handler *AddressHandler) CreateAddress(c *gin.Context){
	var address dto.AddressDTO
	if err := c.ShouldBindJSON(&address); err != nil{
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	err,result := handler.Service.CreateAddress(address)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK,result)

}
