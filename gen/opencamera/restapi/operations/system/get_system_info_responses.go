// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/opencameras/opencamd/gen/opencamera/models"
)

// GetSystemInfoOKCode is the HTTP code returned for type GetSystemInfoOK
const GetSystemInfoOKCode int = 200

/*GetSystemInfoOK Success

swagger:response getSystemInfoOK
*/
type GetSystemInfoOK struct {

	/*
	  In: Body
	*/
	Payload *models.SystemInfo `json:"body,omitempty"`
}

// NewGetSystemInfoOK creates GetSystemInfoOK with default headers values
func NewGetSystemInfoOK() *GetSystemInfoOK {

	return &GetSystemInfoOK{}
}

// WithPayload adds the payload to the get system info o k response
func (o *GetSystemInfoOK) WithPayload(payload *models.SystemInfo) *GetSystemInfoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get system info o k response
func (o *GetSystemInfoOK) SetPayload(payload *models.SystemInfo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSystemInfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSystemInfoUnauthorizedCode is the HTTP code returned for type GetSystemInfoUnauthorized
const GetSystemInfoUnauthorizedCode int = 401

/*GetSystemInfoUnauthorized Access token is missing or invalid

swagger:response getSystemInfoUnauthorized
*/
type GetSystemInfoUnauthorized struct {
}

// NewGetSystemInfoUnauthorized creates GetSystemInfoUnauthorized with default headers values
func NewGetSystemInfoUnauthorized() *GetSystemInfoUnauthorized {

	return &GetSystemInfoUnauthorized{}
}

// WriteResponse to the client
func (o *GetSystemInfoUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}