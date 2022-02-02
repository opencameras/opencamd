// Code generated by go-swagger; DO NOT EDIT.

package media

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// UpdateStunServersOKCode is the HTTP code returned for type UpdateStunServersOK
const UpdateStunServersOKCode int = 200

/*UpdateStunServersOK successful operation

swagger:response updateStunServersOK
*/
type UpdateStunServersOK struct {
}

// NewUpdateStunServersOK creates UpdateStunServersOK with default headers values
func NewUpdateStunServersOK() *UpdateStunServersOK {

	return &UpdateStunServersOK{}
}

// WriteResponse to the client
func (o *UpdateStunServersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}