package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/go-market-app/internal/models"
)

// @Summary AddProduct
// @Security ApiKeyAuth
// @Tags products
// @Description Add Product
// @ID add-product
// @Accept  json
// @Produce  json
// @Param Input body models.ProductRequest true "Product Info"
// @Success 200 {object} Response
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /api/products/ [post]
func (h *Handler) addProduct(c *gin.Context) {
	var input models.ProductRequest

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Product.AddProduct(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

// @Summary ListProducts
// @Security ApiKeyAuth
// @Tags products
// @Description List Products
// @ID list-products
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Product
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /api/products/ [get]
func (h *Handler) listProducts(c *gin.Context) {
	products, err := h.services.Product.GetAllProducts()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, products)
}

// @Summary GetProduct
// @Security ApiKeyAuth
// @Tags products
// @Description Get Product
// @ID get-product
// @Accept  json
// @Produce  json
// @Param id path int true "ProductID"
// @Success 200 {object} models.Product
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /api/products/{id} [get]
func (h *Handler) getProduct(c *gin.Context) {
	id, err := parseId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	product, err := h.services.GetProductByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}

// @Summary UpdateProduct
// @Security ApiKeyAuth
// @Tags products
// @Description Update Product
// @ID update-product
// @Accept  json
// @Produce  json
// @Param Input body models.ProductRequest true "Product Info"
// @Param id path int true "ProductID"
// @Success 200 {object} Response
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /api/products/{id} [put]
func (h *Handler) updateProduct(c *gin.Context) {
	id, err := parseId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input models.ProductRequest

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Product.UpdateProduct(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

// @Summary DeleteProduct
// @Security ApiKeyAuth
// @Tags products
// @Description Delete Product
// @ID delete-product
// @Accept  json
// @Produce  json
// @Param id path int true "ProductID"
// @Success 200 {object} Response
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /api/products/{id} [delete]
func (h *Handler) deleteProduct(c *gin.Context) {
	id, err := parseId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Product.DeleteProduct(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
