package helper

import "github.com/kataras/iris/v12"

// ErrResponse represent error response
type ErrResponse struct {
	reponse Response
}

// CreateErrorResponse create error response
func CreateErrorResponse(ctx iris.Context, err error) ErrResponse {
	Logger.Error(err)
	return ErrResponse{
		CreateResponse(ctx).WithMessage(err.Error()),
	}
}

// BadRequest 400
func (err ErrResponse) BadRequest() Response {
	return err.reponse.WithStatus(400)
}

// InternalServer 500
func (err ErrResponse) InternalServer() Response {
	return err.reponse.WithStatus(500)
}

// Unauthorized 401
func (err ErrResponse) Unauthorized() Response {
	return err.reponse.WithStatus(401)
}

// NotFound 404
func (err ErrResponse) NotFound() Response {
	return err.reponse.WithStatus(404)
}
