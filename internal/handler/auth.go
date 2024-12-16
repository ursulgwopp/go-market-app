package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/market-api/internal/models"
)

// @Summary SignUp
// @Tags auth
// @Description Create New Account
// @ID sign-up
// @Accept  json
// @Produce  json
// @Param Input body models.SignUpRequest true "Sign Up Info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input models.SignUpRequest

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.SignUp(input)
	if err != nil {
		if err.Error() == "password must contain at least one uppercase letter" ||
			err.Error() == "password must contain at least one lowercase letter" ||
			err.Error() == "password must contain at least one digit" ||
			err.Error() == "password must contain at least one special character" {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		if err.Error() == "pq: duplicate key value violates unique constraint \"users_username_key\"" {
			newErrorResponse(c, http.StatusConflict, "username is not unique")
			return
		}

		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary SignIn
// @Tags auth
// @Description Sign Into Account
// @ID sign-in
// @Accept  json
// @Produce  json
// @Param Input body models.SignInRequest true "Sign In Info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input models.SignInRequest

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input)
	if err != nil {
		if err.Error() == "username does not exists" {
			newErrorResponse(c, http.StatusBadRequest, "username does not exists")
			return
		}

		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
