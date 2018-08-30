package controllers

import "net/http"

// Response represents the response format of the API
type Response struct {
	// The API response data
	Data interface{} `json:"data,omitempty"`
	// The API generated message
	Message     string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
	TraceID     string `json:"trace_id,omitempty"`
	// Status Code of the request
	Status int `json:"status"`
	// Success indicates if the request is success
	Success bool `json:"success"`
}

func GetInternalServerError() *Response {
	resp := new(Response)
	resp.Status = http.StatusInternalServerError
	resp.Success = false
	resp.Message = "Something went wrong, Lets try again!"
	return resp

}

func GetBadResponse(message string) *Response {
	resp := new(Response)
	resp.Status = http.StatusBadRequest
	resp.Success = false
	resp.Message = message
	return resp

}

func GetUnauthorizedResponse(message string) *Response {
	resp := new(Response)
	resp.Status = http.StatusUnauthorized
	resp.Success = false
	resp.Message = message
	return resp

}

func GetForbiddenResponse(message string) *Response {
	resp := new(Response)
	resp.Status = http.StatusForbidden
	resp.Success = false
	resp.Message = message
	return resp

}

func GetSuccessResponse(data interface{}) *Response {
	resp := new(Response)
	resp.Status = http.StatusOK
	resp.Success = true
	resp.Message = "Success"
	resp.Data = data
	return resp
}

func GetSuccessResponseWithMessage(data interface{}, message string) *Response {
	resp := new(Response)
	resp.Status = http.StatusOK
	resp.Success = true
	resp.Message = message
	resp.Data = data
	return resp
}

type Controller struct {
}
