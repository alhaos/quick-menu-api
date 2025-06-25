package controller

import (
	"github.com/alhaos/quick-menu-api/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Controller) CreateItemController(gc *gin.Context) {

	var item model.Item

	err := gc.ShouldBindJSON(&item)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client_id, exist := gc.Get("client_id")
	if !exist {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": "client id not set"})
	}

	err = c.repo.CreateItem(client_id.(string), &item)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusCreated, item)
}
