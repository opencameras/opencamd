// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// TokenRefreshOKCode is the HTTP code returned for type TokenRefreshOK
const TokenRefreshOKCode int = 200

/*TokenRefreshOK New extend date

swagger:response tokenRefreshOK
*/
type TokenRefreshOK struct {

	/*
	  In: Body
	*/
	Payload *TokenRefreshOKBody `json:"body,omitempty"`
}

// NewTokenRefreshOK creates TokenRefreshOK with default headers values
func NewTokenRefreshOK() *TokenRefreshOK {

	return &TokenRefreshOK{}
}

// WithPayload adds the payload to the token refresh o k response
func (o *TokenRefreshOK) WithPayload(payload *TokenRefreshOKBody) *TokenRefreshOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the token refresh o k response
func (o *TokenRefreshOK) SetPayload(payload *TokenRefreshOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TokenRefreshOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// TokenRefreshBadRequestCode is the HTTP code returned for type TokenRefreshBadRequest
const TokenRefreshBadRequestCode int = 400

/*TokenRefreshBadRequest Invalid username/password supplied

swagger:response tokenRefreshBadRequest
*/
type TokenRefreshBadRequest struct {
}

// NewTokenRefreshBadRequest creates TokenRefreshBadRequest with default headers values
func NewTokenRefreshBadRequest() *TokenRefreshBadRequest {

	return &TokenRefreshBadRequest{}
}

// WriteResponse to the client
func (o *TokenRefreshBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
