// Copyright 2017-2022 the Kubeapps contributors.
// SPDX-License-Identifier: Apache-2.0

// Package response implements helpers for JSON responses.
package response

import (
	"encoding/json"
	"net/http"
)

/*
ErrorResponse describes a JSON error response with the following body:
	{
		"code": 404,
		"message": "not found"
	}
*/
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewErrorResponse returns a new ErrorResponse
func NewErrorResponse(code int, message string) ErrorResponse {
	return ErrorResponse{code, message}
}

func (e ErrorResponse) Write(w http.ResponseWriter) {
	responseBody, err := json.Marshal(e)
	if err != nil {
		return
	}
	w.WriteHeader(e.Code)
	_, err = w.Write(responseBody)
	if err != nil {
		return
	}
}

/*
DataResponse describes a JSON response containing resource data:
	{
		data: {...}
	}
If resource data is an array of objects, the data key will be an array:
	{
		data: [...]
	}
*/
type DataResponse struct {
	Code int         `json:"-"`
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta,omitempty"`
}

// NewDataResponse returns a new DataResponse
func NewDataResponse(resources interface{}) DataResponse {
	return DataResponse{http.StatusOK, resources, nil}
}

func (d DataResponse) Write(w http.ResponseWriter) {
	responseBody, err := json.Marshal(d)
	if err != nil {
		return
	}
	w.WriteHeader(d.Code)
	_, err = w.Write(responseBody)
	if err != nil {
		return
	}
}
