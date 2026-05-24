package handlers

import (
	"net/http"
	"store-server/internal/cart/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CartHandler struct {
	service *services.CartService
}

func NewCartHandler(service *services.CartService) *CartHandler {
	return &CartHandler{service: service}
}

func (h *CartHandler) GetCart(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	cart, err := h.service.GetCartByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "cart not found" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, cart)
}

func (h *CartHandler) CreateCart(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	cart, err := h.service.CreateCart(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create cart" + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"cart_id": cart})
}
