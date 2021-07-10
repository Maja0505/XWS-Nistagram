package handler

import (
	"XWS-Nistagram/AuthenticationService/service"
	"fmt"
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/persist"
	fileadapter "github.com/casbin/casbin/persist/file-adapter"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

type AuthorizationHandler struct{
	AuthenticationService *service.AuthenticationService
}
// Authorize determines if current subject has been authorized to take an action on an object.
func (handler *AuthorizationHandler) Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := handler.AuthenticationService.TokenValid(c.Request)
		fmt.Println(c.Request.Header.Get("path"))
		fmt.Println(c.Request.Header.Get("Authorization"))
		fmt.Println(c.Request.Header.Get("method"))

		if err != nil {
			c.JSON(http.StatusUnauthorized, "user hasn't logged in yet")
			c.Abort()
			return
		}
		metadata, err := handler.AuthenticationService.ExtractTokenMetadata(c.Request)
		fmt.Println(metadata.UserId)
		fmt.Println(metadata.AccessUuid)
		fmt.Println(metadata.Role)

		if err != nil {
			c.JSON(http.StatusUnauthorized, "unauthorized")
			return
		}
		// casbin enforces policy
		//pwd, _ := os.Getwd()
		adapter := fileadapter.NewAdapter("C:\\Users\\NEMANJA\\Projects\\Go\\src\\XWS-Nistagram\\AuthenticationService\\model\\authorization\\policy.csv")
		ok, err := enforce(metadata.Role, c.Request.Header.Get("path"),c.Request.Header.Get("method"), adapter)
		fmt.Println(metadata.Role)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(500, "error occurred when authorizing user")
			return
		}
		if !ok {
			c.AbortWithStatusJSON(403, "forbidden")
			return
		}
		c.Next()
	}
}



func enforce(sub string, obj string, act string, adapter persist.Adapter) (bool, error) {
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	enforcer := casbin.NewEnforcer("C:\\Users\\NEMANJA\\Projects\\Go\\src\\XWS-Nistagram\\AuthenticationService\\model\\authorization\\auth_model.conf", adapter)
	err := enforcer.LoadPolicy()
	if err != nil {
		return false, fmt.Errorf("failed to load policy from DB: %w", err)
	}
	ok := enforcer.Enforce(sub, obj, act)
	return ok, nil
}