package helper

import "github.com/kataras/iris/v12"

// Response represent http reponse
type Response struct {
	content map[string]interface{}
	context iris.Context
}

// CreateResponse create new response
func CreateResponse(ctx iris.Context) Response {
	return Response{make(map[string]interface{}), ctx}
}

// Ok http 200
func (response Response) Ok() Response {
	response.content["status"] = 200
	return response
}

// WithData set data
func (response Response) WithData(data interface{}) Response {
	response.content["data"] = data
	return response
}

// WithStatus set status
func (response Response) WithStatus(status int) Response {
	response.content["status"] = status
	return response
}

// WithMessage set message
func (response Response) WithMessage(message string) Response {
	response.content["message"] = message
	return response
}

// JSON send response as JSON
func (response Response) JSON() {
	response.context.JSON(response.content)
	response.context.Next()
}
