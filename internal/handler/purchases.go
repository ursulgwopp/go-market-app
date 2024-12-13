package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/go-market-app/internal/models"
)

// @Summary MakePurchase
// @Security ApiKeyAuth
// @Tags purchases
// @Description Make Purchase
// @ID make-purchase
// @Accept  json
// @Produce  json
// @Param quantity query int false "Product Quantity"
// @Param id path int true "ProductID"
// @Success 200 {object} Response
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /api/purchases/{id} [post]
func (h *Handler) makePurchase(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	productId, err := parseId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	quantity_ := c.Query("quantity")
	quantity, err := strconv.Atoi(quantity_)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	purchaseId, err := h.services.Purchase.MakePurchase(models.Purchase{UserId: id, ProductId: productId}, quantity)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": purchaseId,
	})
}

// @Summary GetUserPurchases
// @Security ApiKeyAuth
// @Tags purchases
// @Description Get User Purchases
// @ID get-user-purchases
// @Accept  json
// @Produce  json
// @Param id path int true "UserID"
// @Success 200 {array} models.Purchase
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /api/purchases/user/{id} [get]
func (h *Handler) getUserPurchases(c *gin.Context) {
	userId, err := parseId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	purchases, err := h.services.GetUserPurchases(userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, purchases)
}

// @Summary GetProductPurchases
// @Security ApiKeyAuth
// @Tags purchases
// @Description Get Product Purchases
// @ID get-product-purchases
// @Accept  json
// @Produce  json
// @Param id path int true "ProductID"
// @Success 200 {array} models.Purchase
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /api/purchases/product/{id} [get]
func (h *Handler) getProductPurhases(c *gin.Context) {
	productId, err := parseId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	purchases, err := h.services.GetProductPurchases(productId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, purchases)
}
