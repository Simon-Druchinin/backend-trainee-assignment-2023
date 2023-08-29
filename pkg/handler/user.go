package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userSegments struct {
	User_id int
	Slugs   []string
}

func (h *Handler) validateUserIdParam(paramName string, c *gin.Context) (int, bool) {
	user_id, err := strconv.Atoi(c.Param(paramName))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return 0, false
	}

	exists, err := h.services.Authorization.UserExists(user_id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return 0, false
	}

	if !exists {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf(`User with id '%d' does not exist`, user_id))
		return 0, false
	}

	return user_id, true
}

// @Summary		Show user active segments
// @Tags		users
// @Description	Show user active segments
// @ID			show-user-active-segments
// @Accept		json
// @Produce		json
// @Success		200 {object} userSegments
// @Param		id path int true "User ID"
// @Failure		400 {object} errorResponse
// @Failure		404 {object} errorResponse
// @Failure		500 {object} errorResponse
// @Router		/api/users/{id}/show_active_segments [get]
func (h *Handler) showUserActiveSegments(c *gin.Context) {
	user_id, valid := h.validateUserIdParam("id", c)

	if !valid {
		return
	}

	activeSegments, err := h.services.User.GetActiveSegment(user_id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var output userSegments

	output.User_id = user_id
	output.Slugs = []string{}

	for _, m := range activeSegments {
		output.Slugs = append(output.Slugs, m.Slug)
	}

	c.JSON(http.StatusOK, output)
}

// @Summary		Add user to segment
// @Tags		users
// @Description	Add user to segment
// @ID			add-user-to-segment
// @Accept		json
// @Produce		json
// @Success		201 {object} statusResponse
// @Param		id path int true "User ID"
// @Param		input body []string true "Segment data"
// @Failure		400 {object} errorResponse
// @Failure		404 {object} errorResponse
// @Failure		500 {object} errorResponse
// @Router		/api/users/{id}/add_to_segment [post]
func (h *Handler) addUserToSegment(c *gin.Context) {
	user_id, valid := h.validateUserIdParam("id", c)

	if !valid {
		return
	}

	var input []string

	if err := c.BindJSON(&input); err != nil || len(input) == 0 {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	for _, slug := range input {
		exists, err := h.services.Segment.Exists(slug)

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		if !exists {
			newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf(`Segment with name '%s' does not exist`, slug))
			return
		}

		exists, err = h.services.User.SegmentRelationExists(user_id, slug)

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		if exists {
			newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf(`Relation ('%d':'%s') already exists`, user_id, slug))
			return
		}
	}

	for _, slug := range input {
		_, err := h.services.User.AddToSegment(user_id, slug)

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusCreated, statusResponse{"ok"})
}

// @Summary		Delete user from segment
// @Tags		users
// @Description	Delete user from segment
// @ID			delete-user-from-segment
// @Accept		json
// @Produce		json
// @Success		200 {object} statusResponse
// @Param		id path int true "User ID"
// @Param		slug path string true "Segment slug"
// @Failure		400 {object} errorResponse
// @Failure		404 {object} errorResponse
// @Failure		500 {object} errorResponse
// @Router		/api/users/{id}/delete_from_segment/{slug} [delete]
func (h *Handler) deleteUserFromSegment(c *gin.Context) {
	user_id, valid := h.validateUserIdParam("id", c)

	if !valid {
		return
	}

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

	exists, err = h.services.User.SegmentRelationExists(user_id, slug)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf(`Relation ('%d':'%s') does not exist`, user_id, slug))
		return
	}

	err = h.services.User.DeleteSegmentRelation(user_id, slug)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, statusResponse{"ok"})
}
