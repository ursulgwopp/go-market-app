package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary GetProfile
// @Security ApiKeyAuth
// @Tags users
// @Description Get Your Profile
// @ID get-profile
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /api/users/profile [get]
func (h *Handler) getProfile(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.services.User.GetUserByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary GetUserProfile
// @Security ApiKeyAuth
// @Tags users
// @Description Get User Profile
// @ID get-user-profile
// @Accept  json
// @Produce  json
// @Param id path int true "UserID"
// @Success 200 {object} models.User
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /api/users/{id} [get]
func (h *Handler) getUserProfile(c *gin.Context) {
	id, err := parseId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.User.GetUserByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary ListUsers
// @Security ApiKeyAuth
// @Tags users
// @Description List Users
// @ID list-users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /api/users/ [get]
func (h *Handler) listUsers(c *gin.Context) {
	user, err := h.services.User.GetAllUsers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
