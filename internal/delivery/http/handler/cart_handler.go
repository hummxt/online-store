package handler

import (
	"net/http"

	"ecommerce/internal/delivery/http/dto"
	"ecommerce/internal/domain/service"
	"ecommerce/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CartHandler struct {
	service service.CartService
}

func NewCartHandler(service service.CartService) *CartHandler {
	return &CartHandler{service: service}
}

func (h *CartHandler) GetCart(c *gin.Context) {
	userID, _ := c.Get("user_id")
	cart, err := h.service.GetCart(c.Request.Context(), userID.(uuid.UUID))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve cart")
		return
	}
	utils.JSONResponse(c, http.StatusOK, "Cart retrieved", cart)
}

func (h *CartHandler) AddToCart(c *gin.Context) {
	var req dto.AddToCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, _ := c.Get("user_id")
	if err := h.service.AddToCart(c.Request.Context(), userID.(uuid.UUID), req.ProductID, req.Quantity); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(c, http.StatusOK, "Item added to cart", nil)
}

func (h *CartHandler) UpdateCartItem(c *gin.Context) {
	productID, err := uuid.Parse(c.Param("product_id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	var req dto.UpdateCartItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, _ := c.Get("user_id")
	if err := h.service.UpdateCartItem(c.Request.Context(), userID.(uuid.UUID), productID, req.Quantity); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(c, http.StatusOK, "Cart updated", nil)
}
