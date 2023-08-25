package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zenorachi/dynamic-user-segmentation/internal/entity"
	"net/http"
)

func (h *Handler) initOperationsRoutes(api *gin.RouterGroup) {
	operations := api.Group("/operations", h.userIdentity)
	{
		operations.POST("/add_segments", h.addSegmentsById)
		operations.DELETE("/delete_segments", h.deleteSegmentsById)
	}
}

type operationSegmentsByIdInput struct {
	UserID     int   `json:"user_id" binding:"required"`
	SegmentIDs []int `json:"segment_ids" binding:"required"`
}

type operationSegmentsByIdError struct {
	ID    int    `json:"id"`
	Error string `json:"error"`
}

type operationSegmentsResponse struct {
	OperationIDs []int                        `json:"operation_ids,omitempty"`
	Errors       []operationSegmentsByIdError `json:"errors,omitempty"`
}

func (h *Handler) addSegmentsById(c *gin.Context) {
	var input operationSegmentsByIdInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidInput.Error())
		return
	}

	var response operationSegmentsResponse
	fmt.Println(input.SegmentIDs)
	for _, segmentId := range input.SegmentIDs {
		relation := entity.Relation{
			UserID:    input.UserID,
			SegmentID: segmentId,
		}

		operationId, err := h.services.Operations.CreateBySegmentID(c, relation)
		if err != nil {
			if errors.Is(err, entity.ErrUserDoesNotExist) {
				newErrorResponse(c, http.StatusBadRequest, err.Error())
				return
			} else if errors.Is(err, entity.ErrSegmentDoesNotExist) ||
				errors.Is(err, entity.ErrRelationAlreadyExists) {
				response.Errors = append(response.Errors, operationSegmentsByIdError{
					ID:    segmentId,
					Error: err.Error(),
				})
				continue
			} else {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
		}

		response.OperationIDs = append(response.OperationIDs, operationId)
	}

	newResponse(c, http.StatusCreated, "message", response)
}

func (h *Handler) deleteSegmentsById(c *gin.Context) {
	var input operationSegmentsByIdInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidInput.Error())
		return
	}

	var response operationSegmentsResponse
	for _, segmentId := range input.SegmentIDs {
		relation := entity.Relation{
			UserID:    input.UserID,
			SegmentID: segmentId,
		}

		operationId, err := h.services.Operations.DeleteBySegmentID(c, relation)
		if err != nil {
			if errors.Is(err, entity.ErrRelationDoesNotExist) {
				response.Errors = append(response.Errors, operationSegmentsByIdError{
					ID:    segmentId,
					Error: err.Error(),
				})
				continue
			} else {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
		}

		response.OperationIDs = append(response.OperationIDs, operationId)
	}

	newResponse(c, http.StatusCreated, "message", response)
}
