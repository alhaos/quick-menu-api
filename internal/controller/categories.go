package controller

import (
	"github.com/alhaos/quick-menu-api/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Controller) CreateCategoryController(gc *gin.Context) {

	clientID := gc.Param("client_id")

	category := model.Category{}

	err := gc.ShouldBindJSON(&category)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = c.repo.CreateCategory(clientID, &category)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusCreated, gin.H{"category": category})
}

func (c *Controller) GetCategoryByIdController(gc *gin.Context) {

	id := gc.Param("id")
	clientId := gc.Param("clientId")

	category, err := c.repo.GetCategoryByID(id, clientId)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, gin.H{"category": category})
}

func (c *Controller) UpdateCategoryController(gc *gin.Context) {

	clientID := gc.Param("client_id")

	category := model.Category{}
	err := gc.ShouldBindJSON(&category)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = c.repo.UpdateCategory(clientID, &category)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, gin.H{"category": category})
}

func (c *Controller) DeleteCategoryByIdController(gc *gin.Context) {

	id := gc.Param("id")
	clientID := gc.Param("client_id")

	err := c.repo.DeleteCategoryByID(clientID, id)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusNoContent, nil)
}

func (c *Controller) ListAllCategoriesController(gc *gin.Context) {

	clientID := gc.Param("client_id")

	categories, err := c.repo.ListAllCategories(clientID)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, gin.H{"categories": categories})
}
