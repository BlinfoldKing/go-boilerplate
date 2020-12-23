package helper

import (
	"github.com/kataras/iris/v12"
)

// Response represent http reponse
type Response struct {
	ok      bool
	content struct {
		Status  int         `json:"status"`
		Data    interface{} `json:"data"`
		Message *string     `json:"message"`
	}
	context iris.Context
}

// CreateResponse create new response
func CreateResponse(ctx iris.Context) Response {
	response := Response{}

	response.context = ctx

	return response
}

// Ok http 200
func (response Response) Ok() Response {
	response.content.Status = 200
	return response
}

// WithData set data
func (response Response) WithData(data interface{}) Response {
	response.content.Data = data
	return response
}

// WithStatus set status
func (response Response) WithStatus(status int) Response {
	response.content.Status = status
	return response
}

// WithMessage set message
func (response Response) WithMessage(message string) Response {
	response.content.Message = &message
	return response
}

// JSON send response as JSON
func (response Response) JSON() {
	Logger.
		WithField("content", response.content).
		Debug()
	response.context.JSON(response.content)
}
