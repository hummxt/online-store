package handler

import (
	"net/http"

	"ecommerce/internal/delivery/http/dto"
	"ecommerce/internal/domain/service"
	"ecommerce/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	val, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	userID, err := uuid.Parse(val.(string))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Invalid user ID")
		return
	}

	user, err := h.userService.GetProfile(c.Request.Context(), userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}

	resp := dto.UserResponse{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		Role:        user.Role,
	}

	utils.JSONResponse(c, http.StatusOK, "Profile retrieved", resp)
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	val, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	userID, err := uuid.Parse(val.(string))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Invalid user ID")
		return
	}

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.userService.UpdateProfile(c.Request.Context(), userID, &req); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update profile")
		return
	}

	utils.JSONResponse(c, http.StatusOK, "Profile updated successfully", nil)
}
