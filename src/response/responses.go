package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int         `json:"status,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
	Message string      `json:"message,omitempty"`
}

func New(options ...Option) Response {

	//
	res := &Response{Status: 200}

	//
	for _, option := range options {
		option(res)
	}

	//
	return *res
}

func New404(message string) Response {
	return Response{
		Status:  404,
		Message: message,
	}
}

func New400(message string) Response {
	return Response{
		Status:  400,
		Message: message,
	}
}

func (res Response) WriteLoud(w http.ResponseWriter) error {

	// this comes before, so go won't try to sniff the content-type itself and return plain/text
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(res.Status)

	//
	data, err := json.Marshal(res)
	if err != nil {
		return err
	}

	//
	_, err = w.Write(data)
	return err
}

func (res Response) Write(w http.ResponseWriter) {
	err := res.WriteLoud(w)
	if err != nil {
		panic(err)
	}
}

type Option func(*Response)

func WithStatus(status int) Option {
	return func(response *Response) {
		response.Status = status
	}
}

func WithStatusValidationError() Option {
	return func(response *Response) {
		response.Status = http.StatusUnprocessableEntity
	}
}

func WithPayload(payload interface{}) Option {
	return func(response *Response) {
		response.Payload = payload
	}
}

func WithMessage(message string) Option {
	return func(response *Response) {
		response.Message = message
	}
}
