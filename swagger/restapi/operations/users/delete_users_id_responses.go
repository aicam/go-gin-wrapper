package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/hiromaily/go-gin-wrapper/swagger/models"
)

/*DeleteUsersIDOK An array of user.

swagger:response deleteUsersIdOK
*/
type DeleteUsersIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.UserID `json:"body,omitempty"`
}

// NewDeleteUsersIDOK creates DeleteUsersIDOK with default headers values
func NewDeleteUsersIDOK() *DeleteUsersIDOK {
	return &DeleteUsersIDOK{}
}

// WithPayload adds the payload to the delete users Id o k response
func (o *DeleteUsersIDOK) WithPayload(payload *models.UserID) *DeleteUsersIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete users Id o k response
func (o *DeleteUsersIDOK) SetPayload(payload *models.UserID) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUsersIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteUsersIDDefault Unexpected error.

swagger:response deleteUsersIdDefault
*/
type DeleteUsersIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteUsersIDDefault creates DeleteUsersIDDefault with default headers values
func NewDeleteUsersIDDefault(code int) *DeleteUsersIDDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteUsersIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete users ID default response
func (o *DeleteUsersIDDefault) WithStatusCode(code int) *DeleteUsersIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete users ID default response
func (o *DeleteUsersIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete users ID default response
func (o *DeleteUsersIDDefault) WithPayload(payload *models.Error) *DeleteUsersIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete users ID default response
func (o *DeleteUsersIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUsersIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}