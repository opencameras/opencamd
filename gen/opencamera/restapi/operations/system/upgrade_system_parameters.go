// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// UpgradeSystemMaxParseMemory sets the maximum size in bytes for
// the multipart form parser for this operation.
//
// The default value is 32 MB.
// The multipart parser stores up to this + 10MB.
var UpgradeSystemMaxParseMemory int64 = 32 << 20

// NewUpgradeSystemParams creates a new UpgradeSystemParams object
//
// There are no default values defined in the spec.
func NewUpgradeSystemParams() UpgradeSystemParams {

	return UpgradeSystemParams{}
}

// UpgradeSystemParams contains all the bound params for the upgrade system operation
// typically these are obtained from a http.Request
//
// swagger:parameters upgradeSystem
type UpgradeSystemParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The firmware image to upload.
	  In: formData
	*/
	ImgFile io.ReadCloser
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpgradeSystemParams() beforehand.
func (o *UpgradeSystemParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(UpgradeSystemMaxParseMemory); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}

	imgFile, imgFileHeader, err := r.FormFile("img_file")
	if err != nil && err != http.ErrMissingFile {
		res = append(res, errors.New(400, "reading file %q failed: %v", "imgFile", err))
	} else if err == http.ErrMissingFile {
		// no-op for missing but optional file parameter
	} else if err := o.bindImgFile(imgFile, imgFileHeader); err != nil {
		res = append(res, err)
	} else {
		o.ImgFile = &runtime.File{Data: imgFile, Header: imgFileHeader}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindImgFile binds file parameter ImgFile.
//
// The only supported validations on files are MinLength and MaxLength
func (o *UpgradeSystemParams) bindImgFile(file multipart.File, header *multipart.FileHeader) error {
	return nil
}
