package handler

import (
	"XWS-Nistagram/AuthenticationService/service"
	"fmt"
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/persist"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AuthorizationHandler struct{
	AuthenticationService *service.AuthenticationService
}
// Authorize determines if current subject has been authorized to take an action on an object.
func (handler *AuthorizationHandler) Authorize(obj string, act string, adapter persist.Adapter) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := handler.AuthenticationService.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "user hasn't logged in yet")
			c.Abort()
			return
		}
		metadata, err := handler.AuthenticationService.ExtractTokenMetadata(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "unauthorized")
			return
		}
		// casbin enforces policy
		ok, err := enforce(metadata.Role, obj, act, adapter)
		//ok, err := enforce(val.(string), obj, act, adapter)
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
	enforcer := casbin.NewEnforcer("C:\\Users\\danic\\GOprojects\\src\\XWS-Nistagram\\AuthenticationService\\model\\authorization\\auth_model.conf", adapter)
	err := enforcer.LoadPolicy()
	if err != nil {
		return false, fmt.Errorf("failed to load policy from DB: %w", err)
	}
	ok := enforcer.Enforce(sub, obj, act)
	return ok, nil
}