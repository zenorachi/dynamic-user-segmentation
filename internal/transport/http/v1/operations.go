package v1

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zenorachi/dynamic-user-segmentation/internal/entity"
	"github.com/zenorachi/dynamic-user-segmentation/tools"
)

var (
	defaultPageSize = 5
)

func (h *Handler) initOperationsRoutes(api *gin.RouterGroup) {
	operations := api.Group("/operations", h.userIdentity)
	{
		operations.POST("/add_segments", h.addSegmentsById)
		operations.DELETE("/delete_segments", h.deleteSegmentsById)
		operations.POST("/add_segments_by_name", h.addSegmentsByName)
		operations.DELETE("/delete_segments_by_name", h.deleteSegmentsByName)
		operations.GET("/history", h.getOperationsHistory)
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

type getOperationsHistoryInput struct {
	UserIDs  []int `json:"user_ids"`
	Year     int   `json:"year" binding:"required"`
	Month    int   `json:"month" binding:"required"`
	PageSize int   `json:"page_size"`
}

func (h *Handler) getOperationsHistory(c *gin.Context) {
	var input getOperationsHistoryInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidInput.Error())
		return
	}

	operations, err := h.services.Operations.GetOperationsHistory(c, input.Year, input.Month, input.UserIDs...)
	if err != nil {
		if errors.Is(err, entity.ErrInvalidHistoryPeriod) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	pagedOperations, err := h.generateOperationsPagination(c, input.PageSize, operations)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newResponse(c, http.StatusOK, "operations", pagedOperations)
}

func (h *Handler) generateOperationsPagination(c *gin.Context, pageSize int, operations []entity.Operation) ([]entity.Operation, error) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		return nil, err
	}

	if pageSize <= 0 {
		pageSize = defaultPageSize
	}
	
	startIndex := (page - 1) * pageSize
	endIndex := startIndex + pageSize
	if endIndex > len(operations) {
		endIndex = len(operations)
	}

	return operations[startIndex:endIndex], nil
}
