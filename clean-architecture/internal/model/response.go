package model

import (
	"clean-architecture/internal/i18n"
	"net/http"
	"time"
)

type Response struct {
	Status       int           `json:"status"`
	Message      string        `json:"message"`
	Data         interface{}   `json:"data"`
	ErrorDetails *ErrorDetails `json:"errorDetails"`
	Metadata     *Metadata     `json:"metadata"`
}

type ErrorDetails struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ErrorCode   string `json:"errorCode"`
}

type Metadata struct {
	Timestamp string `json:"timestamp"`
	TraceId   string `json:"traceId"`
}

func ResponseOk() *Response {
	return &Response{http.StatusOK, http.StatusText(http.StatusOK), nil, nil, nil}
}

func ResponseOkWithData(data interface{}) *Response {
	return &Response{http.StatusOK, http.StatusText(http.StatusOK), data, nil, nil}
}

func ResponseBadRequest(errorCode string, lang string, traceId string) *Response {
	title := i18n.Translate(lang, "title."+errorCode)
	description := i18n.Translate(lang, "description."+errorCode)

	errorDetails := &ErrorDetails{
		Title:       title,
		Description: description,
		ErrorCode:   errorCode,
	}

	metadata := &Metadata{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		TraceId:   traceId,
	}
	return &Response{http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil, errorDetails, metadata}
}
