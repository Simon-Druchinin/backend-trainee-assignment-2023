package handler

import (
	"net/http"
	"user_segmentation"

	"github.com/gin-gonic/gin"
)

func (h *Handler) register(c *gin.Context) {
	var input user_segmentation.User

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
