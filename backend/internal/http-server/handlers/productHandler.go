package handlers

import (
	"RestaurantStorage/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary      Add a new product
// @Description  Adds a new product to the storage.
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        product  body      models.ProductModel  true  "Product to add"
// @Success      200      {object}  string    "status: ok"
// @Failure      400      {object}  string   "error: Failed to bind product"
// @Failure      500      {object}  string    "error: Failed to add product"
// @Router       /product [post]
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

// @Summary      Get list of products
// @Description  Retrieves the list of all products in the storage.
// @Tags         Products
// @Produce      json
// @Success      200  {array}   models.ProductModel  "List of products"
// @Failure      500  {object}  string    "error: Failed to get products list"
// @Router       /product [get]
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

// @Summary      Update a product
// @Description  Updates the details of an existing product.
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        product  body      models.ProductModel  true  "Product to update"
// @Success      200      {object}  string    "status: ok"
// @Failure      400      {object} 	string    "error: Failed to bind product"
// @Failure      500      {object}  string    "error: Failed to change product"
// @Router       /product [put]
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

// @Summary      Delete a product
// @Description  Deletes a product from the storage by name.
// @Tags         Products
// @Produce      json
// @Param        name  query     string  true  "Product name"
// @Success      200   {object}  string  "status: ok"
// @Failure      400   {object}  string  "error: Failed to get name"
// @Failure      500   {object}  string  "error: Failed to delete product"
// @Router       /product [delete]
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
