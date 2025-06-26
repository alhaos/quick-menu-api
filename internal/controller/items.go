package controller

import (
	"errors"
	"fmt"
	"github.com/alhaos/quick-menu-api/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateItemController create new item
func (c *Controller) CreateItemController(gc *gin.Context) {

	var item model.Item

	err := gc.ShouldBindJSON(&item)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clientId, err := extractClientId(gc)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	err = c.repo.CreateItem(clientId, &item)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusCreated, item)
}

func (c *Controller) GetItemByIdController(gc *gin.Context) {

	id := gc.Param("id")
	if id == "" {
		gc.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
	}

	clientId, err := extractClientId(gc)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	item, err := c.repo.GetItemById(clientId, id)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, item)
}

// extractClientId extract client_id variable from gin context
func extractClientId(gc *gin.Context) (string, error) {
	id, exist := gc.Get("user_id")
	if !exist {
		return "", errors.New("user_id not found in context")
	}
	idString, ok := id.(string)
	if !ok {
		return "", errors.New("user_id string convert failed")
	}
	return idString, nil
}

// DeleteItemByIdController ...
func (c *Controller) DeleteItemByIdController(gc *gin.Context) {

	id := gc.Param("id")
	if id == "" {
		gc.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
	}

	clientId, err := extractClientId(gc)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	err = c.repo.DeleteItemById(clientId, id)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("item with id %s removed", id)})
}
