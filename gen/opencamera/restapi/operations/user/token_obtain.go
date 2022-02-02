// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/opencameras/opencamd/gen/opencamera/models"
)

// TokenObtainHandlerFunc turns a function with the right signature into a token obtain handler
type TokenObtainHandlerFunc func(TokenObtainParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn TokenObtainHandlerFunc) Handle(params TokenObtainParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// TokenObtainHandler interface for that can handle valid token obtain params
type TokenObtainHandler interface {
	Handle(TokenObtainParams, *models.Principal) middleware.Responder
}

// NewTokenObtain creates a new http.Handler for the token obtain operation
func NewTokenObtain(ctx *middleware.Context, handler TokenObtainHandler) *TokenObtain {
	return &TokenObtain{Context: ctx, Handler: handler}
}

/* TokenObtain swagger:route POST /user/token/obtain user tokenObtain

Logs user into the system and obtain token

*/
type TokenObtain struct {
	Context *middleware.Context
	Handler TokenObtainHandler
}

func (o *TokenObtain) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewTokenObtainParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// TokenObtainOKBody token obtain o k body
//
// swagger:model TokenObtainOKBody
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