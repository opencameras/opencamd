// Code generated by go-swagger; DO NOT EDIT.

package system

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

// NewUpgradeSystemParams creates a new UpgradeSystemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpgradeSystemParams() *UpgradeSystemParams {
	return &UpgradeSystemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpgradeSystemParamsWithTimeout creates a new UpgradeSystemParams object
// with the ability to set a timeout on a request.
func NewUpgradeSystemParamsWithTimeout(timeout time.Duration) *UpgradeSystemParams {
	return &UpgradeSystemParams{
		timeout: timeout,
	}
}

// NewUpgradeSystemParamsWithContext creates a new UpgradeSystemParams object
// with the ability to set a context for a request.
func NewUpgradeSystemParamsWithContext(ctx context.Context) *UpgradeSystemParams {
	return &UpgradeSystemParams{
		Context: ctx,
	}
}

// NewUpgradeSystemParamsWithHTTPClient creates a new UpgradeSystemParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpgradeSystemParamsWithHTTPClient(client *http.Client) *UpgradeSystemParams {
	return &UpgradeSystemParams{
		HTTPClient: client,
	}
}

/* UpgradeSystemParams contains all the parameters to send to the API endpoint
   for the upgrade system operation.

   Typically these are written to a http.Request.
*/
type UpgradeSystemParams struct {

	/* ImgFile.

	   The firmware image to upload.
	*/
	ImgFile runtime.NamedReadCloser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upgrade system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpgradeSystemParams) WithDefaults() *UpgradeSystemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upgrade system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpgradeSystemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upgrade system params
func (o *UpgradeSystemParams) WithTimeout(timeout time.Duration) *UpgradeSystemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upgrade system params
func (o *UpgradeSystemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upgrade system params
func (o *UpgradeSystemParams) WithContext(ctx context.Context) *UpgradeSystemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upgrade system params
func (o *UpgradeSystemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upgrade system params
func (o *UpgradeSystemParams) WithHTTPClient(client *http.Client) *UpgradeSystemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upgrade system params
func (o *UpgradeSystemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImgFile adds the imgFile to the upgrade system params
func (o *UpgradeSystemParams) WithImgFile(imgFile runtime.NamedReadCloser) *UpgradeSystemParams {
	o.SetImgFile(imgFile)
	return o
}

// SetImgFile adds the imgFile to the upgrade system params
func (o *UpgradeSystemParams) SetImgFile(imgFile runtime.NamedReadCloser) {
	o.ImgFile = imgFile
}

// WriteToRequest writes these params to a swagger request
func (o *UpgradeSystemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ImgFile != nil {

		if o.ImgFile != nil {
			// form file param img_file
			if err := r.SetFileParam("img_file", o.ImgFile); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
