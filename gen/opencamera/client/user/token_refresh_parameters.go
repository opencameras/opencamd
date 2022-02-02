// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewTokenRefreshParams creates a new TokenRefreshParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTokenRefreshParams() *TokenRefreshParams {
	return &TokenRefreshParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTokenRefreshParamsWithTimeout creates a new TokenRefreshParams object
// with the ability to set a timeout on a request.
func NewTokenRefreshParamsWithTimeout(timeout time.Duration) *TokenRefreshParams {
	return &TokenRefreshParams{
		timeout: timeout,
	}
}

// NewTokenRefreshParamsWithContext creates a new TokenRefreshParams object
// with the ability to set a context for a request.
func NewTokenRefreshParamsWithContext(ctx context.Context) *TokenRefreshParams {
	return &TokenRefreshParams{
		Context: ctx,
	}
}

// NewTokenRefreshParamsWithHTTPClient creates a new TokenRefreshParams object
// with the ability to set a custom HTTPClient for a request.
func NewTokenRefreshParamsWithHTTPClient(client *http.Client) *TokenRefreshParams {
	return &TokenRefreshParams{
		HTTPClient: client,
	}
}

/* TokenRefreshParams contains all the parameters to send to the API endpoint
   for the token refresh operation.

   Typically these are written to a http.Request.
*/
type TokenRefreshParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the token refresh params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TokenRefreshParams) WithDefaults() *TokenRefreshParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the token refresh params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TokenRefreshParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the token refresh params
func (o *TokenRefreshParams) WithTimeout(timeout time.Duration) *TokenRefreshParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the token refresh params
func (o *TokenRefreshParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the token refresh params
func (o *TokenRefreshParams) WithContext(ctx context.Context) *TokenRefreshParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the token refresh params
func (o *TokenRefreshParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the token refresh params
func (o *TokenRefreshParams) WithHTTPClient(client *http.Client) *TokenRefreshParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the token refresh params
func (o *TokenRefreshParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *TokenRefreshParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
