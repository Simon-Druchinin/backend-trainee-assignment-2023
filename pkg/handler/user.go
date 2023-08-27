package handler

import "github.com/gin-gonic/gin"

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

}

func (h *Handler) addUserToSegment(c *gin.Context) {

}

func (h *Handler) deleteUserFromSegment(c *gin.Context) {

}
