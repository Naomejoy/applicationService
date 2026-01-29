package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Naomejoy/app-service/domain"
	"github.com/Naomejoy/app-service/internal/repository"
	"github.com/Naomejoy/app-service/internal/service"
	"github.com/gin-gonic/gin"
)

type ApplicationHandler struct {
	appService *service.ApplicationService
}

func NewApplicationHandler(appService *service.ApplicationService) *ApplicationHandler {
	return &ApplicationHandler{appService: appService}
}

// @Summary Create a new application
// @Description Create a new application for a user
// @Tags Applications
// @Accept json
// @Produce json
// @Param input body service.CreateApplicationRequest true "Application payload"
// @Success 201 {object} domain.Application
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /applications [post]
// @Security ApiKeyAuth
func (h *ApplicationHandler) CreateApplication(c *gin.Context) {
	var req service.CreateApplicationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	app := &domain.Application{
		Name:   req.Name,
		UserID: req.UserID,
	}

	if err := h.appService.Create(app); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, app)
}

// @Summary Get application by ID
// @Description Retrieve a single application
// @Tags Applications
// @Produce json
// @Param id path int true "Application ID"
// @Success 200 {object} domain.Application
// @Failure 404 {object} map[string]string
// @Router /applications/{id} [get]
func (h *ApplicationHandler) GetApplication(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	app, err := h.appService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "application not found"})
		return
	}
	c.JSON(http.StatusOK, app)
}

// @Summary Update application
// @Description Update application by ID
// @Tags Applications
// @Accept json
// @Produce json
// @Param id path int true "Application ID"
// @Param input body service.UpdateApplicationRequest true "Update payload"
// @Success 200 {object} domain.Application
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /applications/{id} [put]
func (h *ApplicationHandler) UpdateApplication(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req service.UpdateApplicationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	app, err := h.appService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "application not found"})
		return
	}

	if req.Name != "" {
		app.Name = req.Name
	}
	if req.Code != "" {
		app.Code = req.Code
	}
	if req.Description != "" {
		app.Description = req.Description
	}

	if err := h.appService.Update(app); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, app)
}

// @Summary Delete application
// @Description Soft delete application by ID
// @Tags Applications
// @Produce json
// @Param id path int true "Application ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /applications/{id} [delete]
func (h *ApplicationHandler) DeleteApplication(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.appService.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

// @Summary List applications
// @Description List applications with filters, pagination and sorting
// @Tags Applications
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Page size" default(20)
// @Param q query string false "Search query"
// @Param userId query int false "User ID"
// @Param sort query string false "Sort column" default(created_at)
// @Param order query string false "Sort order" default(desc)
// @Param from query string false "Start date YYYY-MM-DD"
// @Param to query string false "End date YYYY-MM-DD"
// @Success 200 {object} service.ListResponse
// @Failure 500 {object} map[string]string
// @Router /applications [get]
func (h *ApplicationHandler) ListApplications(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	q := c.Query("q")
	userID, _ := strconv.ParseUint(c.DefaultQuery("userId", "0"), 10, 64)
	sort := c.DefaultQuery("sort", "created_at")
	order := c.DefaultQuery("order", "desc")
	from := c.Query("from")
	to := c.Query("to")

	params := repository.ApplicationListParams{
		Page:     page,
		PageSize: pageSize,
		Q:        q,
		UserID:   userID,
		Sort:     sort,
		Order:    order,
		From:     parseDatePtr(from),
		To:       parseDatePtr(to),
	}

	resp, err := h.appService.List(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
func parseDatePtr(dateStr string) *time.Time {
	if dateStr == "" {
		return nil
	}
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil
	}
	return &t
}
