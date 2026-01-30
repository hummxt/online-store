package handler

import (
	"net/http"

	"ecommerce/internal/domain/service"
	"ecommerce/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrderHandler struct {
	service service.OrderService
}

func NewOrderHandler(service service.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) PlaceOrder(c *gin.Context) {
	userID, _ := c.Get("user_id")
	order, err := h.service.PlaceOrder(c.Request.Context(), userID.(uuid.UUID))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(c, http.StatusCreated, "Order placed successfully", order)
}

func (h *OrderHandler) GetUserOrders(c *gin.Context) {
	userID, _ := c.Get("user_id")
	orders, err := h.service.GetUserOrders(c.Request.Context(), userID.(uuid.UUID))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve orders")
		return
	}

	utils.JSONResponse(c, http.StatusOK, "Orders retrieved", orders)
}

func (h *OrderHandler) GetOrderDetails(c *gin.Context) {
	orderID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid Order ID")
		return
	}

	order, err := h.service.GetOrder(c.Request.Context(), orderID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Order not found")
		return
	}

	utils.JSONResponse(c, http.StatusOK, "Order details retrieved", order)
}
