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
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidInput.Error())
		return
	}

	var operations []int
	for _, segmentId := range input.SegmentsIDs {
		relation := entity.Relation{
			UserID:    input.UserID,
			SegmentID: segmentId,
		}

		operationId, err := h.services.Operations.CreateBySegmentID(c, relation)
		if err != nil {
			if errors.Is(err, entity.ErrUserDoesNotExist) ||
				errors.Is(err, entity.ErrSegmentDoesNotExist) ||
				errors.Is(err, entity.ErrRelationAlreadyExists) {
				newErrorResponse(c, http.StatusBadRequest, err.Error())
			} else {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
			}
			return
		}

		operations = append(operations, operationId)
	}

	newResponse(c, http.StatusCreated, "operations_ids", operations)
}

func (h *Handler) deleteSegmentsById(c *gin.Context) {
	var input operationSegmentsByIdInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidInput.Error())
		return
	}

	var operations []int
	for _, segmentId := range input.SegmentsIDs {
		relation := entity.Relation{
			UserID:    input.UserID,
			SegmentID: segmentId,
		}

		operationId, err := h.services.Operations.DeleteBySegmentID(c, relation)
		if err != nil {
			if errors.Is(err, entity.ErrRelationDoesNotExist) {
				newErrorResponse(c, http.StatusBadRequest, err.Error())
			} else {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
			}
			return
		}

		operations = append(operations, operationId)
	}

	newResponse(c, http.StatusCreated, "operations_ids", operations)
}

type operationSegmentsByNameInput struct {
	UserID        int      `json:"user_id" binding:"required"`
	SegmentsNames []string `json:"segments_names" binding:"required"`
}

func (h *Handler) addSegmentsByName(c *gin.Context) {
	var input operationSegmentsByNameInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidInput.Error())
		return
	}

	var operations []int
	for _, segmentName := range input.SegmentsNames {
		operationId, err := h.services.Operations.CreateBySegmentName(c, input.UserID, segmentName)
		if err != nil {
			if errors.Is(err, entity.ErrUserDoesNotExist) ||
				errors.Is(err, entity.ErrSegmentDoesNotExist) ||
				errors.Is(err, entity.ErrRelationAlreadyExists) {
				newErrorResponse(c, http.StatusBadRequest, err.Error())
			} else {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
			}
			return
		}

		operations = append(operations, operationId)
	}

	newResponse(c, http.StatusCreated, "operations_ids", operations)
}

func (h *Handler) deleteSegmentsByName(c *gin.Context) {
	var input operationSegmentsByNameInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidInput.Error())
		return
	}

	var operations []int
	for _, segmentName := range input.SegmentsNames {
		operationId, err := h.services.Operations.DeleteBySegmentName(c, input.UserID, segmentName)
		if err != nil {
			if errors.Is(err, entity.ErrRelationDoesNotExist) {
				newErrorResponse(c, http.StatusBadRequest, err.Error())
			} else {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
			}
			return
		}

		operations = append(operations, operationId)
	}

	newResponse(c, http.StatusCreated, "operations_ids", operations)
}
