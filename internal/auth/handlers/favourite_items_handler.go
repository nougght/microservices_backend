package handlers

import (
	"fmt"
	"net/http"
	"store-server/internal/auth/models"
	"store-server/internal/auth/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FavouriteItemsHandler struct {
	service *services.FavouriteItemsService
}

func NewFavouriteItemsHandler(service *services.FavouriteItemsService) *FavouriteItemsHandler {
	return &FavouriteItemsHandler{service: service}
}

func (h *FavouriteItemsHandler) AddToFavourites(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	var item models.FavouriteItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input" + err.Error()})
		return
	}
	item.UserID = userID
	productID, err := uuid.Parse(c.Param("product_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}
	item.ProductID = productID

	if err := h.service.AddToFavourites(c.Request.Context(), &item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add item to favourites" + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, item)
}

func (h *FavouriteItemsHandler) GetFavouritesByUserID(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	favourites, err := h.service.GetFavouritesByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve favourites" + err.Error()})
		return
	}
	fmt.Println(favourites)
	c.JSON(http.StatusOK, favourites)
}

func (h *FavouriteItemsHandler) DeleteFromFavourites(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
	productID, err := uuid.Parse(c.Param("product_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	if err := h.service.DeleteFromFavourites(c.Request.Context(), userID, productID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete item from favourites" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
