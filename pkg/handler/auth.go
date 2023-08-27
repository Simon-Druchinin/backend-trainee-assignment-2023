package handler

import (
	"net/http"
	"user_segmentation"

	"github.com/gin-gonic/gin"
)

// @Summary		Register
// @Tags		auth
// @Description	Create user
// @ID			create-user
// @Produce		application/json
// @Success		201	{object} successBaseResponse "id"
// @Failure		400 {object} errorResponse
// @Failure		404 {object} errorResponse
// @Failure		500 {object} errorResponse
// @Router		/auth/register [post]
func (h *Handler) register(c *gin.Context) {
	var input user_segmentation.User

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, successBaseResponse{
		Id: id,
	})
}
