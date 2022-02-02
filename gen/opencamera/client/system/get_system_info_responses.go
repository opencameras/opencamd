// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/opencameras/opencamd/gen/opencamera/models"
)

// GetSystemInfoReader is a Reader for the GetSystemInfo structure.
type GetSystemInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSystemInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSystemInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSystemInfoUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSystemInfoOK creates a GetSystemInfoOK with default headers values
func NewGetSystemInfoOK() *GetSystemInfoOK {
	return &GetSystemInfoOK{}
}

/* GetSystemInfoOK describes a response with status code 200, with default header values.

Success
*/
type GetSystemInfoOK struct {
	Payload *models.SystemInfo
}

func (o *GetSystemInfoOK) Error() string {
	return fmt.Sprintf("[GET /system][%d] getSystemInfoOK  %+v", 200, o.Payload)
}
func (o *GetSystemInfoOK) GetPayload() *models.SystemInfo {
	return o.Payload
}

func (o *GetSystemInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSystemInfoUnauthorized creates a GetSystemInfoUnauthorized with default headers values
func NewGetSystemInfoUnauthorized() *GetSystemInfoUnauthorized {
	return &GetSystemInfoUnauthorized{}
}

/* GetSystemInfoUnauthorized describes a response with status code 401, with default header values.

Access token is missing or invalid
*/
type GetSystemInfoUnauthorized struct {
}

func (o *GetSystemInfoUnauthorized) Error() string {
	return fmt.Sprintf("[GET /system][%d] getSystemInfoUnauthorized ", 401)
}

func (o *GetSystemInfoUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
