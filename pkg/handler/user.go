package handler

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"user_segmentation"

	"github.com/gin-gonic/gin"
)

type userSegments struct {
	User_id int      `json:"user_id"`
	Slugs   []string `json:"slugs"`
}

const usersSegmentsHistoryFile = "usersSegmentsHistory.csv"

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
// @Param		input body []string true "Segment data"
// @Failure		400 {object} errorResponse
// @Failure		404 {object} errorResponse
// @Failure		500 {object} errorResponse
// @Router		/api/users/{id}/delete_from_segment [delete]
func (h *Handler) deleteUserFromSegment(c *gin.Context) {
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

		if !exists {
			newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf(`Relation ('%d':'%s') does not exist`, user_id, slug))
			return
		}
	}

	for _, slug := range input {
		err := h.services.User.DeleteSegmentRelation(user_id, slug)

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusCreated, statusResponse{"ok"})
}

func (h *Handler) CSVExport(user_segments_history []user_segmentation.UserSegmentHistory) (*os.File, error) {

	file, err := os.Create(usersSegmentsHistoryFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var user_segments_history_csv [][]string
	for _, row := range user_segments_history {
		var temp []string
		temp = append(temp, fmt.Sprint(row.User_id), fmt.Sprint(row.Slug), fmt.Sprint(row.Operation_type), fmt.Sprint(row.Timestamp))
		user_segments_history_csv = append(user_segments_history_csv, temp)
	}

	for _, row := range user_segments_history_csv {
		err = writer.Write(row)
		if err != nil {
			return nil, err
		}
	}

	return file, nil
}

// @Summary		Show user segments history
// @Tags		users
// @Description	Show user segments history
// @ID			show-user-segments-history
// @Accept		json
// @Produce		text/csv
// @Success		200 {object} user_segmentation.UserSegmentHistory
// @Param		year path int true "Operation year"
// @Param		month path int true "Operation month"
// @Failure		400 {object} errorResponse
// @Failure		404 {object} errorResponse
// @Failure		500 {object} errorResponse
// @Router		/api/users/show_segments_history/{year}/{month} [get]
func (h *Handler) showUserSegmentsHistory(c *gin.Context) {
	month, err := strconv.Atoi(c.Param("month"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid month param")
		return
	}

	year, err := strconv.Atoi(c.Param("year"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid year param")
		return
	}

	user_segments_history, err := h.services.User.GetSegmentRelationHistory(month, year)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = h.CSVExport(user_segments_history)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.FileAttachment(fmt.Sprintf("./%s", usersSegmentsHistoryFile), usersSegmentsHistoryFile)
	c.Writer.Header().Set("attachment", fmt.Sprintf("filename=%s", usersSegmentsHistoryFile))
}
