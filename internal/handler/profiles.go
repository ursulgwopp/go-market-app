package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/go-market-app/internal/models"
)

// @Summary GetProfile
// @Security ApiKeyAuth
// @Tags profile
// @Description Get Your Profile
// @ID get-profile
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /api/profile [get]
func (h *Handler) getProfile(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.Profile.GetProfile(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary DeleteProfile
// @Security ApiKeyAuth
// @Tags profile
// @Description Delete Profile
// @ID delete-profile
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /api/profile/delete [delete]
func (h *Handler) deleteProfile(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.User.DeleteUser(userId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "ok"})
}

// @Summary Deposit
// @Security ApiKeyAuth
// @Tags profile
// @Description Deposit
// @ID deposit
// @Accept  json
// @Produce  json
// @Param amount query int true "Amount"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /api/profile/deposit [post]
func (h *Handler) deposit(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	amount_ := c.Query("amount")
	amount, err := strconv.Atoi(amount_)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Profile.Deposit(userId, amount); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "ok"})
}
