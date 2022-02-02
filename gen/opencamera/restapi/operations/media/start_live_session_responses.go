// Code generated by go-swagger; DO NOT EDIT.

package media

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// StartLiveSessionOKCode is the HTTP code returned for type StartLiveSessionOK
const StartLiveSessionOKCode int = 200

/*StartLiveSessionOK successful operation

swagger:response startLiveSessionOK
*/
type StartLiveSessionOK struct {
}

// NewStartLiveSessionOK creates StartLiveSessionOK with default headers values
func NewStartLiveSessionOK() *StartLiveSessionOK {

	return &StartLiveSessionOK{}
}

// WriteResponse to the client
func (o *StartLiveSessionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}