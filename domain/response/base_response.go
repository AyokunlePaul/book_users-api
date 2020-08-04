package response

import "net/http"

type BaseResponse struct {
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Status  int         `json:"-"`
	Data    interface{} `json:"data,omitempty"`
}

func NewCreateResponse(message string, data interface{}) *BaseResponse {
	return &BaseResponse{
		Message: message,
		Success: true,
		Data:    data,
		Status:  http.StatusCreated,
	}
}

func NewOkResponse(message string, data interface{}) *BaseResponse {
	return &BaseResponse{
		Message: message,
		Success: true,
		Status:  http.StatusOK,
		Data:    data,
	}
}

func NewBadRequestError(message string) *BaseResponse {
	return &BaseResponse{
		Message: message,
		Success: false,
		Status:  http.StatusBadRequest,
	}
}

func NewNotFoundError(message string) *BaseResponse {
	return &BaseResponse{
		Message: message,
		Success: false,
		Status:  http.StatusNotFound,
	}
}

func NewInternalServerError(message string) *BaseResponse {
	return &BaseResponse{
		Message: message,
		Success: false,
		Status:  http.StatusInternalServerError,
	}
}
