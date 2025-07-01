package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Controller) AuthMiddleware(gc *gin.Context) {

	cookieToken, _ := gc.Cookie("authToken")

	authHeaderToken := gc.GetHeader("Authorization")

	var finalToken string

	switch {
	case cookieToken != "":
		finalToken = cookieToken
	case authHeaderToken != "":
		finalToken = authHeaderToken
	default:
		gc.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization token missing"})
		return
	}

	userID, err := c.auth.CheckToken(finalToken)
	if err != nil {
		gc.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	gc.Set("user_id", userID)

	gc.Next()
	return
}
