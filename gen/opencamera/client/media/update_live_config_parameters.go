// Code generated by go-swagger; DO NOT EDIT.

package media

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

	"github.com/opencameras/opencamd/gen/opencamera/models"
)

// NewUpdateLiveConfigParams creates a new UpdateLiveConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateLiveConfigParams() *UpdateLiveConfigParams {
	return &UpdateLiveConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateLiveConfigParamsWithTimeout creates a new UpdateLiveConfigParams object
// with the ability to set a timeout on a request.
func NewUpdateLiveConfigParamsWithTimeout(timeout time.Duration) *UpdateLiveConfigParams {
	return &UpdateLiveConfigParams{
		timeout: timeout,
	}
}

// NewUpdateLiveConfigParamsWithContext creates a new UpdateLiveConfigParams object
// with the ability to set a context for a request.
func NewUpdateLiveConfigParamsWithContext(ctx context.Context) *UpdateLiveConfigParams {
	return &UpdateLiveConfigParams{
		Context: ctx,
	}
}

// NewUpdateLiveConfigParamsWithHTTPClient creates a new UpdateLiveConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateLiveConfigParamsWithHTTPClient(client *http.Client) *UpdateLiveConfigParams {
	return &UpdateLiveConfigParams{
		HTTPClient: client,
	}
}

/* UpdateLiveConfigParams contains all the parameters to send to the API endpoint
   for the update live config operation.

   Typically these are written to a http.Request.
*/
type UpdateLiveConfigParams struct {

	/* Body.

	   Live session configuration
	*/
	Body *models.LiveSessionConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update live config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateLiveConfigParams) WithDefaults() *UpdateLiveConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update live config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateLiveConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update live config params
func (o *UpdateLiveConfigParams) WithTimeout(timeout time.Duration) *UpdateLiveConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update live config params
func (o *UpdateLiveConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update live config params
func (o *UpdateLiveConfigParams) WithContext(ctx context.Context) *UpdateLiveConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update live config params
func (o *UpdateLiveConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update live config params
func (o *UpdateLiveConfigParams) WithHTTPClient(client *http.Client) *UpdateLiveConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update live config params
func (o *UpdateLiveConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update live config params
func (o *UpdateLiveConfigParams) WithBody(body *models.LiveSessionConfig) *UpdateLiveConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update live config params
func (o *UpdateLiveConfigParams) SetBody(body *models.LiveSessionConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateLiveConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}