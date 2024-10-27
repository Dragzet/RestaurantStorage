package handlers

import (
	"RestaurantStorage/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) AddProduct(c *gin.Context) {
	var product models.ProductModel
	if err := c.BindJSON(&product); err != nil {
		h.logger.Error("Failed to bind product", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind product"})
		return
	}
	if err := h.service.AddProduct(c, product); err != nil {
		h.logger.Error("Failed to add product", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handler) GetProductsList(c *gin.Context) {
	products, err := h.service.GetProductsList(c)
	if err != nil {
		h.logger.Error("Failed to get products list", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get products list"})
		return
	}
	if products == nil {
		c.JSON(http.StatusOK, gin.H{"status": "empty"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *Handler) ChangeProduct(c *gin.Context) {
	var product models.ProductModel
	if err := c.BindJSON(&product); err != nil {
		h.logger.Error("Failed to bind product", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind product"})
		return
	}
	if err := h.service.ChangeProduct(c, product.Name, product.Amount); err != nil {
		h.logger.Error("Failed to change product", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to change product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		h.logger.Error("Failed to get name", "error", "empty name")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get name"})
		return
	}
	if err := h.service.DeleteProduct(c, name); err != nil {
		h.logger.Error("Failed to delete product", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
