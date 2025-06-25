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

	err = c.repo.CreateItem(&item)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusCreated, item)
}
