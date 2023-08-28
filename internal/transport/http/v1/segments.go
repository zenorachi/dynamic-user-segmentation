package v1

import (
	"errors"
	//_ "google.golang.org/genproto/googleapis/apps/drive/labels/v2beta"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zenorachi/dynamic-user-segmentation/internal/entity"
)

func (h *Handler) initSegmentsRoutes(api *gin.RouterGroup) {
	segments := api.Group("/segments", h.userIdentity)
	{
		segments.POST("/create", h.createSegment)
		segments.GET("/", h.getAllSegments)
		segments.GET("/:segment_id", h.getSegmentById)
		segments.GET("/active_users/:segment_id", h.getActiveUsers)
		segments.DELETE("/delete", h.deleteSegmentByName)
		segments.DELETE("/delete_by_id", h.deleteSegmentById)
	}
}

type createSegmentInput struct {
	Name          string  `json:"name" binding:"required,min=2,max=64"`
	AssignPercent float64 `json:"assign_percent,omitempty"`
}

type createSegmentResponse struct {
	ID int `json:"id"`
}

// @Summary Create segment
// @Security Bearer
// @Description create new segment
// @Tags segments
// @Accept json
// @Produce json
// @Param input body createSegmentInput true "input"
// @Success 201 {object} createSegmentResponse
// @Failure 400 {object} errorResponse
// @Failure 409 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v1/segments/create [post]
func (h *Handler) createSegment(c *gin.Context) {
	var input createSegmentInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidInput.Error())
		return
	}

	id, err := h.services.Segments.Create(c, entity.Segment{Name: input.Name, AssignPercent: input.AssignPercent})
	if err != nil {
		if errors.Is(err, entity.ErrSegmentAlreadyExists) {
			newErrorResponse(c, http.StatusConflict, err.Error())
		} else if errors.Is(err, entity.ErrInvalidAssignPercent) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	newResponse(c, http.StatusCreated, createSegmentResponse{ID: id})
}

type getAllSegmentsResponse struct {
	Segments []entity.Segment `json:"segments"`
}

// @Summary Get all segments
// @Security Bearer
// @Description getting all segments
// @Tags segments
// @Produce json
// @Success 200 {object} getAllSegmentsResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v1/segments/ [get]
func (h *Handler) getAllSegments(c *gin.Context) {
	segments, err := h.services.Segments.GetAll(c)
	if err != nil {
		if errors.Is(err, entity.ErrSegmentDoesNotExist) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	newResponse(c, http.StatusOK, getAllSegmentsResponse{Segments: segments})
}

type getSegmentByIdResponse struct {
	Segment entity.Segment `json:"segment"`
}

// @Summary Get Segment By ID
// @Security Bearer
// @Description getting segment by id
// @Tags segments
// @Produce json
// @Param segment_id path int true "Segment ID"
// @Success 200 {object} getSegmentByIdResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v1/segments/{id} [get]
func (h *Handler) getSegmentById(c *gin.Context) {
	paramId := strings.Trim(c.Param("segment_id"), "/")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid parameter (id)")
		return
	}

	segment, err := h.services.Segments.GetByID(c, id)
	if err != nil {
		if errors.Is(err, entity.ErrSegmentDoesNotExist) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	newResponse(c, http.StatusOK, getSegmentByIdResponse{Segment: segment})
}

type getActiveUsersResponse struct {
	Users []entity.User `json:"users"`
}

// @Summary Get Active Users By ID
// @Security Bearer
// @Description getting active users by id
// @Tags segment-users
// @Produce json
// @Param segment_id path int true "Segment ID"
// @Success 200 {object} getActiveUsersResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v1/segments/active_users/:segment_id [get]
func (h *Handler) getActiveUsers(c *gin.Context) {
	paramId := strings.Trim(c.Param("segment_id"), "/")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid parameter (id)")
		return
	}

	users, err := h.services.Segments.GetActiveUsersBySegmentID(c, id)
	if err != nil {
		if errors.Is(err, entity.ErrSegmentDoesNotExist) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	newResponse(c, http.StatusOK, getActiveUsersResponse{Users: users})
}

type deleteByNameInput struct {
	Name string `json:"name" binding:"required"`
}

// @Summary Delete Segment By Name
// @Security Bearer
// @Description deletion segment by name
// @Tags segments
// @Accept json
// @Param input body deleteByNameInput true "input"
// @Success 204 "No Content"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v1/segments/delete/ [delete]
func (h *Handler) deleteSegmentByName(c *gin.Context) {
	var input deleteByNameInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidInput.Error())
		return
	}

	err := h.services.Segments.DeleteByName(c, input.Name)
	if err != nil {
		if errors.Is(err, entity.ErrSegmentDoesNotExist) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	newResponse(c, http.StatusNoContent, nil)
}

type deleteByIdInput struct {
	ID int `json:"id" binding:"required"`
}

// @Summary Delete Segment By ID
// @Security Bearer
// @Description deletion segment by id
// @Tags segments
// @Accept json
// @Param input body deleteByIdInput true "input"
// @Success 204 "No Content"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v1/segments/delete_by_id/ [delete]
func (h *Handler) deleteSegmentById(c *gin.Context) {
	var input deleteByIdInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidInput.Error())
		return
	}

	err := h.services.Segments.DeleteByID(c, input.ID)
	if err != nil {
		if errors.Is(err, entity.ErrSegmentDoesNotExist) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	newResponse(c, http.StatusNoContent, nil)
}
