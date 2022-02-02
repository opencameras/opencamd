// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TokenObtainReader is a Reader for the TokenObtain structure.
type TokenObtainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TokenObtainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTokenObtainOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTokenObtainBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTokenObtainOK creates a TokenObtainOK with default headers values
func NewTokenObtainOK() *TokenObtainOK {
	return &TokenObtainOK{}
}

/* TokenObtainOK describes a response with status code 200, with default header values.

Encoded token
*/
type TokenObtainOK struct {
	Payload *TokenObtainOKBody
}

func (o *TokenObtainOK) Error() string {
	return fmt.Sprintf("[POST /user/token/obtain][%d] tokenObtainOK  %+v", 200, o.Payload)
}
func (o *TokenObtainOK) GetPayload() *TokenObtainOKBody {
	return o.Payload
}

func (o *TokenObtainOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(TokenObtainOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenObtainBadRequest creates a TokenObtainBadRequest with default headers values
func NewTokenObtainBadRequest() *TokenObtainBadRequest {
	return &TokenObtainBadRequest{}
}

/* TokenObtainBadRequest describes a response with status code 400, with default header values.

Invalid username/password supplied
*/
type TokenObtainBadRequest struct {
}

func (o *TokenObtainBadRequest) Error() string {
	return fmt.Sprintf("[POST /user/token/obtain][%d] tokenObtainBadRequest ", 400)
}

func (o *TokenObtainBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*TokenObtainOKBody token obtain o k body
swagger:model TokenObtainOKBody
*/
type TokenObtainOKBody struct {

	// access token
	AccessToken string `json:"access_token,omitempty"`

	// token is valid until this specified date in UTC.
	// Example: 2022-02-02T00:00:00Z
	// Format: date
	ValidUntil strfmt.Date `json:"valid_until,omitempty"`
}

// Validate validates this token obtain o k body
func (o *TokenObtainOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateValidUntil(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TokenObtainOKBody) validateValidUntil(formats strfmt.Registry) error {
	if swag.IsZero(o.ValidUntil) { // not required
		return nil
	}

	if err := validate.FormatOf("tokenObtainOK"+"."+"valid_until", "body", "date", o.ValidUntil.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this token obtain o k body based on context it is used
func (o *TokenObtainOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TokenObtainOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TokenObtainOKBody) UnmarshalBinary(b []byte) error {
	var res TokenObtainOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
