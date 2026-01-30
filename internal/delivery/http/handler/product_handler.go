package handler

import (
	"net/http"
	"strconv"

	"ecommerce/internal/delivery/http/dto"
	"ecommerce/internal/domain/entity"
	"ecommerce/internal/domain/service"
	"ecommerce/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req dto.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	categoryID, err := uuid.Parse(req.CategoryID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid Category ID")
		return
	}

	product := &entity.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		CategoryID:  categoryID,
		ImageURLs:   req.ImageURLs,
	}

	if err := h.service.CreateProduct(c.Request.Context(), product); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create product")
		return
	}

	utils.JSONResponse(c, http.StatusCreated, "Product created successfully", product)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	product, err := h.service.GetProduct(c.Request.Context(), id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Product not found")
		return
	}

	utils.JSONResponse(c, http.StatusOK, "Product retrieved", product)
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	products, total, err := h.service.ListProducts(c.Request.Context(), page, limit)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to list products")
		return
	}

	utils.JSONResponse(c, http.StatusOK, "Products retrieved", gin.H{
		"products": products,
		"total":    total,
		"page":     page,
		"limit":    limit,
	})
}

func (h *ProductHandler) CreateCategory(c *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	category := &entity.Category{
		Name:        req.Name,
		Description: req.Description,
		Slug:        req.Slug,
	}

	if err := h.service.CreateCategory(c.Request.Context(), category); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create category")
		return
	}

	utils.JSONResponse(c, http.StatusCreated, "Category created successfully", category)
}

func (h *ProductHandler) ListCategories(c *gin.Context) {
	categories, err := h.service.ListCategories(c.Request.Context())
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to list categories")
		return
	}
	utils.JSONResponse(c, http.StatusOK, "Categories retrieved", categories)
}

func (h *ProductHandler) SearchProducts(c *gin.Context) {
	query := c.Query("q")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	products, total, err := h.service.SearchProducts(c.Request.Context(), query, page, limit)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Search failed")
		return
	}

	utils.JSONResponse(c, http.StatusOK, "Search results", gin.H{
		"products": products,
		"total":    total,
		"page":     page,
		"limit":    limit,
	})
}
