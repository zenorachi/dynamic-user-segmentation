package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zenorachi/dynamic-user-segmentation/internal/entity"
	"net/http"
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
	UserID      int   `json:"user_id" binding:"required"`
	SegmentsIDs []int `json:"segments_ids" binding:"required"`
}

func (h *Handler) addSegmentsById(c *gin.Context) {
	var input operationSegmentsByIdInput
	if err := c.BindJSON(&input); err != nil || len(input.SegmentsIDs) == 0 {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidInput.Error())
		return
	}

	var relations []entity.Relation
	for _, segmentId := range input.SegmentsIDs {
		relations = append(relations, entity.Relation{UserID: input.UserID, SegmentID: segmentId})
	}

	operations, err := h.services.Operations.CreateBySegmentID(c, relations)
	if err != nil {
		if errors.Is(err, entity.ErrUserDoesNotExist) || errors.Is(err, entity.ErrSegmentDoesNotExist) ||
			errors.Is(err, entity.ErrRelationAlreadyExists) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	newResponse(c, http.StatusCreated, "operations_ids", operations)
}

func (h *Handler) deleteSegmentsById(c *gin.Context) {
	var input operationSegmentsByIdInput
	if err := c.BindJSON(&input); err != nil || len(input.SegmentsIDs) == 0 {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidInput.Error())
		return
	}

	var relations []entity.Relation
	for _, segmentId := range input.SegmentsIDs {
		relations = append(relations, entity.Relation{UserID: input.UserID, SegmentID: segmentId})
	}

	operations, err := h.services.Operations.DeleteBySegmentID(c, relations)
	if err != nil {
		if errors.Is(err, entity.ErrUserDoesNotExist) || errors.Is(err, entity.ErrSegmentDoesNotExist) ||
			errors.Is(err, entity.ErrRelationDoesNotExist) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	newResponse(c, http.StatusOK, "operations_ids", operations)
}

type operationSegmentsByNameInput struct {
	UserID        int      `json:"user_id" binding:"required"`
	SegmentsNames []string `json:"segments_names" binding:"required"`
}

func (h *Handler) addSegmentsByName(c *gin.Context) {
	var input operationSegmentsByNameInput
	if err := c.BindJSON(&input); err != nil || len(input.SegmentsNames) == 0 {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidInput.Error())
		return
	}

	operations, err := h.services.Operations.CreateBySegmentName(c, input.UserID, input.SegmentsNames)
	if err != nil {
		if errors.Is(err, entity.ErrUserDoesNotExist) || errors.Is(err, entity.ErrSegmentDoesNotExist) ||
			errors.Is(err, entity.ErrRelationAlreadyExists) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	newResponse(c, http.StatusCreated, "operations_ids", operations)
}

func (h *Handler) deleteSegmentsByName(c *gin.Context) {
	var input operationSegmentsByNameInput
	if err := c.BindJSON(&input); err != nil || len(input.SegmentsNames) == 0 {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidInput.Error())
		return
	}

	operations, err := h.services.Operations.DeleteBySegmentName(c, input.UserID, input.SegmentsNames)
	if err != nil {
		if errors.Is(err, entity.ErrUserDoesNotExist) || errors.Is(err, entity.ErrSegmentDoesNotExist) ||
			errors.Is(err, entity.ErrRelationDoesNotExist) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	newResponse(c, http.StatusOK, "operations_ids", operations)
}
