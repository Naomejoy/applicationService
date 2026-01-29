package service

type PaginationMeta struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"pageSize"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"totalPages"`
}

type ListResponse struct {
	Data interface{}    `json:"data"`
	Meta PaginationMeta `json:"meta"`
}

type CreateApplicationRequest struct {
	Name   string `json:"name" binding:"required"`
	UserID uint64 `json:"userId" binding:"required"`
}

type UpdateApplicationRequest struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type AddStatusRequest struct {
	Status string `json:"status" binding:"required"`
	UserID uint64 `json:"userId" binding:"required"`
}

type AddFileTypeRequest struct {
	FileTypeName string `json:"fileTypeName" binding:"required"`
}
