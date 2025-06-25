package controller

import (
	"github.com/alhaos/quick-menu-api/internal/repository"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	repo *repository.Repository
}

// New is constructor
func New(repo *repository.Repository) *Controller {
	return &Controller{repo: repo}
}

func SetupRouter(router *gin.Engine, c *Controller) {

	api := router.Group("/api/clients")
	{
		client := api.Group("/:client_id")
		{
			categories := client.Group("/categories")
			{
				categories.POST("", c.CreateCategoryController)
				categories.GET("/:id", c.GetCategoryByIdController)
				categories.PATCH("", c.UpdateCategoryController)
				categories.DELETE("/:id", c.DeleteCategoryByIdController)
				categories.GET("", c.ListAllCategoriesController)
			}
			items := client.Group("/items")
			{
				items.POST("", c.CreateItemController)
			}
		}
	}
}
