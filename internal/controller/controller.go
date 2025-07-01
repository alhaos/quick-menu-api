package controller

import (
	"github.com/alhaos/quick-menu-api/internal/authService"
	"github.com/alhaos/quick-menu-api/internal/repository"
	"github.com/gin-gonic/gin"
	"time"
)

// Controller provides HTTP handlers for authentication and user management.
// It encapsulates dependencies required for handling user-related operations.
//
// Fields:
//
//	repo   - Repository layer for database operations
//	auth   - Authentication service for token generation and validation
//	secret - Secret key used for cryptographic operations (e.g., JWT signing)
//
// Usage Example:
//
//	c := &Controller{
//	    repo:   repository.New(),
//	    auth:   authService.New(secret),
//	    secret: []byte("your-256-bit-secret"),
//	}
//
// Security Note:
//   - The secret should be properly managed (not hardcoded)
//   - Consider using environment variables for secret configuration
type Controller struct {
	repo   *repository.Repository
	auth   *authService.AuthService
	secret []byte
}

// New creates and returns a new Controller instance with the specified dependencies.
//
// This is the preferred way to initialize a Controller as it ensures all required
// dependencies are properly injected. The returned Controller is ready to handle HTTP requests.
//
// Parameters:
//
//	repo - Repository instance for data access operations (must not be nil)
//	auth - AuthService instance for authentication operations (must not be nil)
//
// Returns:
//
//	*Controller - initialized controller instance
//
// Example:
//
//	ctrl := New(
//	    repository.New(),
//	    authService.New("secret-key"),
//	)
//
// Panics:
//   - Will panic if either repo or auth parameters are nil to fail fast during initialization
func New(repo *repository.Repository, auth *authService.AuthService) *Controller {
	return &Controller{repo: repo, auth: auth}
}

// SetupRouter configures all API routes for the application.
//
// Organizes routes into logical groups with proper middleware application.
// Routes are divided into three main sections:
//   - Private routes (require authentication)
//   - Public routes (no authentication)
//   - Auth routes (authentication handlers)
//
// Parameters:
//
//	router - The Gin engine instance to configure (must not be nil)
//	c      - Controller instance containing route handlers (must not be nil)
//
// Route Structure:
//
//	/api
//	  /private
//	    /categories       (CRUD operations)
//	    /items            (CRUD operations)
//	  /public
//	    /info             (Public information)
//	  /auth
//	    /login            (Authentication)
//	    /check            (Session check)
//	    /logoff           (Session termination)
//
// Middleware:
//   - All private routes are protected with AuthMiddleware
//   - Public routes have no authentication requirements
//
// Example:
//
//	router := gin.Default()
//	ctrl := NewController(repo, authService)
//	SetupRouter(router, ctrl)
//
// Notes:
//   - Panics if either parameter is nil
//   - Route groups follow RESTful conventions
//   - All private routes require valid JWT in Authorization header
func SetupRouter(router *gin.Engine, c *Controller) {
	api := router.Group("/api")
	{
		private := api.Group("/private")
		{
			categories := private.Group("/categories")
			{
				categories.POST("", c.AuthMiddleware, c.CreateCategoryController)
				categories.GET("/:id", c.AuthMiddleware, c.GetCategoryByIdController)
				categories.PATCH("", c.AuthMiddleware, c.UpdateCategoryController)
				categories.DELETE("/:id", c.AuthMiddleware, c.DeleteCategoryByIdController)
				categories.GET("", c.AuthMiddleware, c.ListAllCategoriesController)
			}

			items := private.Group("/items")
			{
				items.POST("", c.AuthMiddleware, c.CreateItemController)
				items.GET("/:id", c.AuthMiddleware, c.GetItemByIdController)
				items.DELETE("/:id", c.AuthMiddleware, c.DeleteItemByIdController)
				items.PATCH("", c.AuthMiddleware, c.UpdateItemByIdController)
				items.GET("", c.AuthMiddleware, c.ListAllItemsController)
			}
		}
		public := api.Group("/public")
		{
			info := public.Group("/info")
			{
				info.GET("", c.InfoHandler())
			}
		}
		auth := api.Group("/auth")
		{
			auth.POST("/login", c.LoginController)
			auth.GET("/check", c.AuthCheckController)
			auth.GET("/logoff", c.LogoffController)
		}
	}
}

// InfoHandler
// todo: remove this in prodaction this is for test only
func (c *Controller) InfoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"current time is": time.Now().Format(time.DateTime)})
	}
}
