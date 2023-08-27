package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary		Show user active segments
// @Tags		users
// @Description	Show user active segments
// @ID			show-user-active-segments
// @Accept		json
// @Produce		json
// @Success		200
// @Param		id path int true "User ID"
// @Failure		400 {object} errorResponse
// @Failure		404 {object} errorResponse
// @Failure		500 {object} errorResponse
// @Router		/api/users/{id}/show_active_segments [get]
func (h *Handler) showUserActiveSegments(c *gin.Context) {
	user_id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return
	}

	exists, err := h.services.Authorization.UserExists(user_id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf(`User with id '%d' does not exist`, user_id))
		return
	}

	activeSegments, err := h.services.User.GetActiveSegment(user_id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, activeSegments)
}

func (h *Handler) addUserToSegment(c *gin.Context) {

}

func (h *Handler) deleteUserFromSegment(c *gin.Context) {

}
