package model

import "net/http"

// CodeResponse is a generic success response
type CodeResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg,omitempty"`
}

// BooleanResponse is a generic success response containing code and a boolean value
type BooleanResponse struct {
	Code    int    `json:"code"`
	Data    bool   `json:"data"`
	Message string `json:"msg,omitempty"`
}

// CodeOK returns HTTP Status OK (200)
func CodeOK() CodeResponse {
	return CodeResponse{Code: http.StatusOK}
}

// NewBoolResponse returns new Boolean response
func NewBoolResponse(b bool) BooleanResponse {
	return BooleanResponse{Code: http.StatusOK, Data: b}
}
