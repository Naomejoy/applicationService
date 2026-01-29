package api

import (
	"net/http"
	"strconv"

	"github.com/Naomejoy/app-service/internal/service"

	"github.com/gin-gonic/gin"
)

type StatusHandler struct {
	statusService *service.ApplicationStatusService
}

func NewStatusHandler(statusService *service.ApplicationStatusService) *StatusHandler {
	return &StatusHandler{statusService: statusService}
}

// @Summary Add status to an application
// @Description Add a new status for an application
// @Tags ApplicationStatus
// @Accept json
// @Produce json
// @Param id path int true "Application ID"
// @Param input body service.AddStatusRequest true "Status payload"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /applications/{id}/status [post]
func (h *StatusHandler) AddStatus(c *gin.Context) {
	appID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req service.AddStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.statusService.Add(appID, req.UserID, req.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "status added"})
}

// @Summary List statuses of an application
// @Description List all statuses for an application with pagination
// @Tags ApplicationStatus
// @Produce json
// @Param id path int true "Application ID"
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Page size" default(20)
// @Success 200 {object} service.ListResponse
// @Failure 500 {object} map[string]string
// @Router /applications/{id}/statuses [get]
func (h *StatusHandler) ListStatuses(c *gin.Context) {
	appID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))

	resp, err := h.statusService.List(appID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
