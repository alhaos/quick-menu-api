package controller

import (
	"github.com/alhaos/quick-menu-api/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginController handles user authentication request.
//
// Accepts JSON with user credentials (email and password), validates them,
// and returns JWT token on success. Sets token both in cookie and response body.
//
// Parameters:
//
//	gc - Gin context containing HTTP request/response
//
// Returns:
//
//	On success:
//	  - HTTP 200 with JSON {"token": "JWT_TOKEN"}
//	  - Sets "authToken" cookie with JWT
//	On error:
//	  - HTTP 400 for invalid JSON
//	  - HTTP 401 for invalid credentials
//
// Request body example:
//
//	{
//	  "email": "user@example.com",
//	  "password": "securepassword"
//	}
//
// Success response example:
//
//	{
//	  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
//	}
//
// Notes:
//   - Cookie is set with:
//   - MaxAge: 3600 seconds (1 hour)
//   - Path: "/"
//   - HttpOnly: true
//   - Secure: false (should be true in production)
//
// LoginController handles user authentication request.
//
// Accepts JSON with user credentials (email and password), validates them,
// and returns JWT token on success. Sets token both in cookie and response body.
//
// Parameters:
//
//	gc - Gin context containing HTTP request/response
//
// Returns:
//
//	On success:
//	  - HTTP 200 with JSON {"token": "JWT_TOKEN"}
//	  - Sets "authToken" cookie with JWT
//	On error:
//	  - HTTP 400 for invalid JSON
//	  - HTTP 401 for invalid credentials
//
// Request body example:
//
//	{
//	  "email": "user@example.com",
//	  "password": "securepassword"
//	}
//
// Success response example:
//
//	{
//	  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
//	}
//
// Notes:
//   - Cookie is set with:
//   - MaxAge: 3600 seconds (1 hour)
//   - Path: "/"
//   - HttpOnly: true
//   - Secure: false (should be true in production)
func (c *Controller) LoginController(gc *gin.Context) {

	var u model.User

	err := gc.ShouldBindJSON(&u)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := c.auth.Login(u)
	if err != nil {
		gc.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	gc.SetCookie("authToken",
		tokenString,
		3600,
		"/",
		"",
		false,
		true)

	gc.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// AuthCheckController checks for the presence of a token in the session or Cookie or in the Authorization header
func (c *Controller) AuthCheckController(gc *gin.Context) {

	tokenInCookie, err := gc.Cookie("authToken")
	if err != nil {
		gc.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"error": err.Error()},
		)
		return
	}

	tokenInHeader := gc.GetHeader("Authorization")

	var tokenSelected string

	switch {
	case tokenInHeader != "":
		tokenSelected = tokenInHeader
	case tokenInCookie != "":
		tokenSelected = tokenInCookie
	default:
		gc.JSON(http.StatusUnauthorized,
			gin.H{"error": "token not provided"},
		)
		return
	}

	userID, err := c.auth.CheckToken(tokenSelected)
	if err != nil {
		gc.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"error": err.Error()},
		)
		return
	}

	gc.JSON(http.StatusOK, gin.H{
		"authenticated": true,
		"user_id":       userID})
}

// LogoffController remove authToken from Cookie
func (c *Controller) LogoffController(gc *gin.Context) {

	gc.SetCookie(
		"authToken",
		"",
		-1,
		"/",
		"localhost",
		true,
		true,
	)

	gc.JSON(http.StatusOK, gin.H{"msg": "success logoff"})

}
