// Code generated by go-swagger; DO NOT EDIT.

package device

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/opencameras/opencamd/gen/opencamera/models"
)

// GetDeviceInfoReader is a Reader for the GetDeviceInfo structure.
type GetDeviceInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDeviceInfoUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeviceInfoOK creates a GetDeviceInfoOK with default headers values
func NewGetDeviceInfoOK() *GetDeviceInfoOK {
	return &GetDeviceInfoOK{}
}

/* GetDeviceInfoOK describes a response with status code 200, with default header values.

Success
*/
type GetDeviceInfoOK struct {
	Payload *models.DeviceInfo
}

func (o *GetDeviceInfoOK) Error() string {
	return fmt.Sprintf("[GET /device][%d] getDeviceInfoOK  %+v", 200, o.Payload)
}
func (o *GetDeviceInfoOK) GetPayload() *models.DeviceInfo {
	return o.Payload
}

func (o *GetDeviceInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceInfoUnauthorized creates a GetDeviceInfoUnauthorized with default headers values
func NewGetDeviceInfoUnauthorized() *GetDeviceInfoUnauthorized {
	return &GetDeviceInfoUnauthorized{}
}

/* GetDeviceInfoUnauthorized describes a response with status code 401, with default header values.

Access token is missing or invalid
*/
type GetDeviceInfoUnauthorized struct {
}

func (o *GetDeviceInfoUnauthorized) Error() string {
	return fmt.Sprintf("[GET /device][%d] getDeviceInfoUnauthorized ", 401)
}

func (o *GetDeviceInfoUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
