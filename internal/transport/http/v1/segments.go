package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zenorachi/dynamic-user-segmentation/internal/entity"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) initSegmentsRoutes(api *gin.RouterGroup) {
	segments := api.Group("/segments", h.userIdentity)
	{
		segments.POST("/create", h.createSegment)
		segments.GET("/", h.getAllSegments)
		segments.GET("/:segment_id", h.getSegmentById)
		segments.GET("/active/:user_id", h.getActiveUserSegments)
		segments.DELETE("/delete", h.deleteSegmentByName)
		segments.DELETE("/delete_by_id", h.deleteSegmentById)
	}
}

type createSegmentInput struct {
	Name string `json:"name" binding:"required,min=2,max=64"`
}

func (h *Handler) createSegment(c *gin.Context) {
	var input createSegmentInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidInput.Error())
		return
	}

	id, err := h.services.Segments.Create(c, entity.Segment{Name: input.Name})
	if err != nil {
		if errors.Is(err, entity.ErrSegmentAlreadyExists) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	newResponse(c, http.StatusCreated, "id", id)
}

func (h *Handler) getAllSegments(c *gin.Context) {
	segments, err := h.services.Segments.GetAll(c)
	if err != nil {
		if errors.Is(err, entity.ErrSegmentDoesNotExist) {
			newResponse(c, http.StatusOK, "message", "no available segments")
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	newResponse(c, http.StatusOK, "segments", segments)
}

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

	newResponse(c, http.StatusOK, "segment", segment)
}

func (h *Handler) getActiveUserSegments(c *gin.Context) {
	paramId := strings.Trim(c.Param("user_id"), "/")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid parameter (id)")
		return
	}

	segments, err := h.services.Segments.GetActiveByUserID(c, id)
	if err != nil {
		if errors.Is(err, entity.ErrUserDoesNotExist) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	newResponse(c, http.StatusOK, "active_segments", segments)
}

type deleteByNameInput struct {
	Name string `json:"name" binding:"required"`
}

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

	newResponse(c, http.StatusNoContent, "", "")
}

type deleteByIdInput struct {
	ID int `json:"id" binding:"required"`
}

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

	newResponse(c, http.StatusNoContent, "", "")
}
