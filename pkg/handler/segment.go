package handler

import (
	"net/http"
	"user_segmentation"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createSegment(c *gin.Context) {
	var input user_segmentation.Segment

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	id, err := h.services.Segment.CreateSegment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) deleteSegment(c *gin.Context) {

}
