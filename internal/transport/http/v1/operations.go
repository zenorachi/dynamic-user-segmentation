package v1

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zenorachi/dynamic-user-segmentation/internal/entity"
)

const (
	defaultPageSize = 5
)

func (h *Handler) initOperationsRoutes(api *gin.RouterGroup) {
	operations := api.Group("/operations", h.userIdentity)
	{
		operations.POST("/add_segments", h.addSegmentsById)
		operations.POST("/add_segments_by_names", h.addSegmentsByName)
		operations.DELETE("/delete_segments", h.deleteSegmentsById)
		operations.DELETE("/delete_segments_by_names", h.deleteSegmentsByName)
		operations.GET("/history", h.getOperationsHistory)
	}
}

type operationsResponse struct {
	IDs []int `json:"operation_ids"`
}

type addSegmentsByIdInput struct {
	UserID      int    `json:"user_id" binding:"required"`
	SegmentsIDs []int  `json:"segment_ids" binding:"required"`
	TTL         string `json:"ttl,omitempty"`
}

// @Summary Add a user to segments by id
// @Security JWT
// @Description addition a user to segments by id
// @Tags operations
// @Accept json
// @Produce json
// @Param input body addSegmentsByIdInput true "input"
// @Success 201 {object} operationsResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v1/operations/add_segments/ [post]
func (h *Handler) addSegmentsById(c *gin.Context) {
	var input addSegmentsByIdInput
	if err := c.BindJSON(&input); err != nil || len(input.SegmentsIDs) == 0 {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidInput.Error())
		return
	}

	ttl, err := h.parseTTL(input.TTL)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidTTL.Error())
		return
	}

	operationIDs, err := h.services.Operations.CreateBySegmentIDs(c, input.UserID, input.SegmentsIDs)
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

	newResponse(c, http.StatusCreated, operationsResponse{IDs: operationIDs})
}

type addSegmentsByNameInput struct {
	UserID        int      `json:"user_id" binding:"required"`
	SegmentsNames []string `json:"segment_names" binding:"required"`
	TTL           string   `json:"ttl,omitempty"`
}

// @Summary Add a user to segments by name
// @Security JWT
// @Description addition a user to segments by name
// @Tags operations
// @Accept json
// @Produce json
// @Param input body addSegmentsByNameInput true "input"
// @Success 201 {object} operationsResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v1/operations/add_segments_by_names/ [post]
func (h *Handler) addSegmentsByName(c *gin.Context) {
	var input addSegmentsByNameInput
	if err := c.BindJSON(&input); err != nil || len(input.SegmentsNames) == 0 {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidInput.Error())
		return
	}

	ttl, err := h.parseTTL(input.TTL)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidTTL.Error())
		return
	}

	operationIDs, err := h.services.Operations.CreateBySegmentNames(c, input.UserID, input.SegmentsNames)
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

	newResponse(c, http.StatusCreated, operationsResponse{IDs: operationIDs})
}

type deleteSegmentsByIdInput struct {
	UserID      int   `json:"user_id" binding:"required"`
	SegmentsIDs []int `json:"segment_ids" binding:"required"`
}

// @Summary Delete User From Segments by ids
// @Security JWT
// @Description delete user-segments relation by ids
// @Tags operations
// @Accept json
// @Produce json
// @Param input body deleteSegmentsByIdInput true "input"
// @Success 200 {object} operationsResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v1/operations/delete_segments/ [delete]
func (h *Handler) deleteSegmentsById(c *gin.Context) {
	var input deleteSegmentsByIdInput

	if err := c.BindJSON(&input); err != nil || len(input.SegmentsIDs) == 0 {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidInput.Error())
		return
	}

	operationIDs, err := h.services.Operations.DeleteBySegmentIDs(c, input.UserID, input.SegmentsIDs)
	if err != nil {
		if errors.Is(err, entity.ErrUserDoesNotExist) || errors.Is(err, entity.ErrSegmentDoesNotExist) ||
			errors.Is(err, entity.ErrRelationDoesNotExist) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	newResponse(c, http.StatusOK, operationsResponse{IDs: operationIDs})
}

type deleteSegmentsByNameInput struct {
	UserID        int      `json:"user_id" binding:"required"`
	SegmentsNames []string `json:"segment_names" binding:"required"`
}

// @Summary Delete User From Segments By Names
// @Security JWT
// @Description delete user-segments relation by names
// @Tags operations
// @Accept json
// @Produce json
// @Param input body deleteSegmentsByNameInput true "input"
// @Success 200 {object} operationsResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v1/operations/delete_segments_by_names/ [delete]
func (h *Handler) deleteSegmentsByName(c *gin.Context) {
	var input deleteSegmentsByNameInput
	if err := c.BindJSON(&input); err != nil || len(input.SegmentsNames) == 0 {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidInput.Error())
		return
	}

	operationIDs, err := h.services.Operations.DeleteBySegmentNames(c, input.UserID, input.SegmentsNames)
	if err != nil {
		if errors.Is(err, entity.ErrUserDoesNotExist) || errors.Is(err, entity.ErrSegmentDoesNotExist) ||
			errors.Is(err, entity.ErrRelationDoesNotExist) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	newResponse(c, http.StatusOK, operationsResponse{IDs: operationIDs})
}

type getOperationsHistoryInput struct {
	UserIDs  []int `json:"user_ids"`
	Year     int   `json:"year" binding:"required"`
	Month    int   `json:"month" binding:"required"`
	PageSize int   `json:"page_size"`
}

type getOperationsHistoryResponse struct {
	Operations []entity.Operation `json:"operations"`
}

// @Summary Get Operations History
// @Security JWT
// @Description getting operations history
// @Tags operations
// @Accept json
// @Produce json
// @Param input body getOperationsHistoryInput true "input"
// @Success 200 {object} getOperationsHistoryResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v1/operations/history/ [get]
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

	newResponse(c, http.StatusOK, getOperationsHistoryResponse{Operations: pagedOperations})
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

func (h *Handler) parseTTL(ttl string) (time.Duration, error) {
	if ttl == "" {
		return 0, nil
	}

	ttlDuration, err := time.ParseDuration(ttl)
	if err != nil {
		return 0, err
	}

	return ttlDuration, nil
}
