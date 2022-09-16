package model

type MessageResponse struct {
	Data     interface{} `json:"data"`
	Messages Responses   `json:"messages"`
	Errors   Responses   `json:"errors"`
}

type Response struct {
	Code    StatusCode `json:"code,omitempty"`
	Message string     `json:"message,omitempty"`
}

type Responses []Response
