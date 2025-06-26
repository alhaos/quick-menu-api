package controller

import (
	"github.com/alhaos/quick-menu-api/internal/repository"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	repo   *repository.Repository
	secret []byte
}

// New is constructor
func New(repo *repository.Repository, secret []byte) *Controller {
	return &Controller{repo: repo, secret: secret}
}

func SetupRouter(router *gin.Engine, c *Controller) {

	api := router.Group("/api")
	{
		categories := api.Group("/categories")
		{
			categories.POST("", c.AuthMiddleware, c.CreateCategoryController)
			categories.GET("/:id", c.AuthMiddleware, c.GetCategoryByIdController)
			categories.PATCH("", c.AuthMiddleware, c.UpdateCategoryController)
			categories.DELETE("/:id", c.AuthMiddleware, c.DeleteCategoryByIdController)
			categories.GET("", c.AuthMiddleware, c.ListAllCategoriesController)
		}

		items := api.Group("/items")
		{
			items.POST("", c.AuthMiddleware, c.CreateItemController)
			items.GET("/:id", c.AuthMiddleware, c.GetItemByIdController)
			items.DELETE("/:id", c.AuthMiddleware, c.DeleteItemByIdController)
		}

		login := api.Group("/login")
		{
			login.POST("", c.LoginController)
		}
	}
}
