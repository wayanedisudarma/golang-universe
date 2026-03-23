package model

import "net/http"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseOk() *Response {
	return &Response{http.StatusOK, http.StatusText(http.StatusOK), nil}
}

func ResponseOkWithData(data interface{}) *Response {
	return &Response{http.StatusOK, http.StatusText(http.StatusOK), data}
}

func ResponseBadRequest(message string) *Response {
	return &Response{http.StatusBadRequest, message, nil}
}
