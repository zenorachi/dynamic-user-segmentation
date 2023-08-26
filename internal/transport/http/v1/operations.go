package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zenorachi/dynamic-user-segmentation/internal/entity"
	"github.com/zenorachi/dynamic-user-segmentation/tools"
)

func (h *Handler) initOperationsRoutes(api *gin.RouterGroup) {
	operations := api.Group("/operations", h.userIdentity)
	{
		operations.POST("/add_segments", h.addSegmentsById)
		operations.DELETE("/delete_segments", h.deleteSegmentsById)
		operations.POST("/add_segments_by_name", h.addSegmentsByName)
		operations.DELETE("/delete_segments_by_name", h.deleteSegmentsByName)
	}
}

type operationSegmentsByIdInput struct {
	UserID      int    `json:"user_id" binding:"required"`
	SegmentsIDs []int  `json:"segment_ids" binding:"required"`
	TTL         string `json:"ttl,omitempty"`
}

func (h *Handler) addSegmentsById(c *gin.Context) {
	var input operationSegmentsByIdInput
	if err := c.BindJSON(&input); err != nil || len(input.SegmentsIDs) == 0 {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidInput.Error())
		return
	}

	ttl, err := tools.ParseTTL(input.TTL)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidTTL.Error())
		return
	}

	operations, err := h.services.Operations.CreateBySegmentIDs(c, input.UserID, input.SegmentsIDs)
	if err != nil {
		if errors.Is(err, entity.ErrUserDoesNotExist) || errors.Is(err, entity.ErrSegmentDoesNotExist) ||
			errors.Is(err, entity.ErrRelationAlreadyExists) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if input.TTL != "" {
		go func() { h.services.Operations.DeleteAfterTTLBySegmentIDs(c, input.UserID, input.SegmentsIDs, ttl) }()
	}

	newResponse(c, http.StatusCreated, "operation_ids", operations)
}

func (h *Handler) deleteSegmentsById(c *gin.Context) {
	var input operationSegmentsByIdInput

	if err := c.BindJSON(&input); err != nil || len(input.SegmentsIDs) == 0 {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidInput.Error())
		return
	}

	operations, err := h.services.Operations.DeleteBySegmentIDs(c, input.UserID, input.SegmentsIDs)
	if err != nil {
		if errors.Is(err, entity.ErrUserDoesNotExist) || errors.Is(err, entity.ErrSegmentDoesNotExist) ||
			errors.Is(err, entity.ErrRelationDoesNotExist) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	newResponse(c, http.StatusOK, "operation_ids", operations)
}

type operationSegmentsByNameInput struct {
	UserID        int      `json:"user_id" binding:"required"`
	SegmentsNames []string `json:"segment_names" binding:"required"`
	TTL           string   `json:"ttl,omitempty"`
}

func (h *Handler) addSegmentsByName(c *gin.Context) {
	var input operationSegmentsByNameInput
	if err := c.BindJSON(&input); err != nil || len(input.SegmentsNames) == 0 {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidInput.Error())
		return
	}

	ttl, err := tools.ParseTTL(input.TTL)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidTTL.Error())
		return
	}

	operations, err := h.services.Operations.CreateBySegmentNames(c, input.UserID, input.SegmentsNames)
	if err != nil {
		if errors.Is(err, entity.ErrUserDoesNotExist) || errors.Is(err, entity.ErrSegmentDoesNotExist) ||
			errors.Is(err, entity.ErrRelationAlreadyExists) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if input.TTL != "" {
		go func() { h.services.Operations.DeleteAfterTTLBySegmentNames(c, input.UserID, input.SegmentsNames, ttl) }()
	}

	newResponse(c, http.StatusCreated, "operation_ids", operations)
}

func (h *Handler) deleteSegmentsByName(c *gin.Context) {
	var input operationSegmentsByNameInput
	if err := c.BindJSON(&input); err != nil || len(input.SegmentsNames) == 0 {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidInput.Error())
		return
	}

	operations, err := h.services.Operations.DeleteBySegmentNames(c, input.UserID, input.SegmentsNames)
	if err != nil {
		if errors.Is(err, entity.ErrUserDoesNotExist) || errors.Is(err, entity.ErrSegmentDoesNotExist) ||
			errors.Is(err, entity.ErrRelationDoesNotExist) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	newResponse(c, http.StatusOK, "operation_ids", operations)
}
