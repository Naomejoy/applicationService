package api

import (
	"net/http"
	"strconv"

	"github.com/Naomejoy/app-service/internal/service"

	"github.com/gin-gonic/gin"
)

type FileTypeHandler struct {
	fileService *service.ApplicationFileTypeService
}

func NewFileTypeHandler(fileService *service.ApplicationFileTypeService) *FileTypeHandler {
	return &FileTypeHandler{fileService: fileService}
}

// @Summary Add file type to an application
// @Description Add a new file type to an application
// @Tags ApplicationFileTypes
// @Accept json
// @Produce json
// @Param id path int true "Application ID"
// @Param input body service.AddFileTypeRequest true "File type payload"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /applications/{id}/file-types [post]
func (h *FileTypeHandler) AddFileType(c *gin.Context) {
	appID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req service.AddFileTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.fileService.Add(appID, req.FileTypeName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "file type added"})
}

// @Summary Delete file type
// @Description Delete a file type from an application
// @Tags ApplicationFileTypes
// @Produce json
// @Param fileTypeId path int true "FileType ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /applications/{id}/file-types/{fileTypeId} [delete]
func (h *FileTypeHandler) DeleteFileType(c *gin.Context) {
	fileTypeID, _ := strconv.ParseUint(c.Param("fileTypeId"), 10, 64)
	if err := h.fileService.Delete(fileTypeID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

// @Summary List file types
// @Description List all file types for an application
// @Tags ApplicationFileTypes
// @Produce json
// @Param id path int true "Application ID"
// @Success 200 {array} domain.ApplicationUploadedFileType
// @Failure 500 {object} map[string]string
// @Router /applications/{id}/file-types [get]
func (h *FileTypeHandler) ListFileTypes(c *gin.Context) {
	appID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	fileTypes, err := h.fileService.List(appID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, fileTypes)
}
