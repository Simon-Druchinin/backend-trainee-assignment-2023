package handler

import (
	"fmt"
	"net/http"
	"user_segmentation"

	"github.com/gin-gonic/gin"
)

// @Summary		Create segment
// @Tags		segments
// @Description	Create a segment
// @ID			create-segment
// @Accept		json
// @Produce		json
// @Success		201	{object} successBaseResponse "id"
// @Param		input body user_segmentation.Segment true "Segment data"
// @Failure		400 {object} errorResponse
// @Failure		404 {object} errorResponse
// @Failure		500 {object} errorResponse
// @Router		/api/segments [post]
func (h *Handler) createSegment(c *gin.Context) {
	var input user_segmentation.Segment

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	exists, err := h.services.Segment.Exists(input.Slug)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if exists {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf(`Segment with name '%s' already exists`, input.Slug))
		return
	}

	id, err := h.services.Segment.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, successBaseResponse{
		Id: id,
	})
}

// @Summary		Delete segment
// @Tags		segments
// @Description	Delete a segment
// @ID			delete-segment
// @Accept		json
// @Produce		json
// @Success		200 {object} statusResponse
// @Param		slug path string true "Segment slug"
// @Failure		400 {object} errorResponse
// @Failure		404 {object} errorResponse
// @Failure		500 {object} errorResponse
// @Router		/api/segments/{slug} [delete]
func (h *Handler) deleteSegment(c *gin.Context) {
	slug := c.Param("slug")

	exists, err := h.services.Segment.Exists(slug)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf(`Segment with name '%s' does not exist`, slug))
		return
	}

	h.services.Segment.Delete(slug)

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
