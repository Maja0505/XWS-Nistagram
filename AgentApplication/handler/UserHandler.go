package handler

import (
	"XWS-Nistagram/AgentApplication/dto"
	"XWS-Nistagram/AgentApplication/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct{
	Service *service.UserService
}

func (handler *UserHandler) RegisterUser(c *gin.Context){
	var u dto.UserDTO
	if err := c.ShouldBindJSON(&u); err != nil || (u.FirstName=="" || u.LastName=="" || u.Email=="" || u.Password =="" || u.ConfirmedPassword=="" ) {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	err := handler.Service.RegisterUser(u)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK,"User successfully registered")

}

func (handler *UserHandler) Login(c *gin.Context){
	var loginForm dto.LoginFormDTO
	if err := c.ShouldBindJSON(&loginForm); err != nil || (loginForm.Email=="" || loginForm.Password=="" ) {
		c.JSON(http.StatusUnprocessableEntity, "Invalid form")
		return
	}

	found,credentialsValid,user :=handler.Service.CheckCredentials(loginForm.Email,loginForm.Password)
	if !found{
		c.JSON(http.StatusUnauthorized, "User not found")
		return
	}
	if !credentialsValid{
		c.JSON(http.StatusUnauthorized, "Please provide valid credentials")
		return
	}
	c.JSON(http.StatusOK, &user)

}

