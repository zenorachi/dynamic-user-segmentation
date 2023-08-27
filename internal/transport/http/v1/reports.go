package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zenorachi/dynamic-user-segmentation/internal/entity"
)

func (h *Handler) initReportsRoutes(api *gin.RouterGroup) {
	reports := api.Group("/reports", h.userIdentity)
	{
		reports.GET("/file", h.getReportFile)
		reports.GET("/link", h.getReportLink)
	}
}

type getReportInput struct {
	UserIDs []int `json:"user_ids"`
	Year    int   `json:"year" binding:"required"`
	Month   int   `json:"month" binding:"required"`
}

func (h *Handler) getReportFile(c *gin.Context) {
	var input getReportInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidInput.Error())
		return
	}

	file, err := h.services.Reports.CreateReportFile(c, input.Year, input.Month, input.UserIDs...)
	if err != nil {
		if errors.Is(err, entity.ErrInvalidHistoryPeriod) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.Data(http.StatusOK, "text/csv", file)
}

func (h *Handler) getReportLink(c *gin.Context) {
	var input getReportInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidInput.Error())
		return
	}

	link, err := h.services.Reports.CreateReportLink(c, input.Year, input.Month, input.UserIDs...)
	if err != nil {
		if errors.Is(err, entity.ErrInvalidHistoryPeriod) || errors.Is(err, entity.ErrGDriveIsNotAvailable) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	newResponse(c, http.StatusOK, "link", link)
}
