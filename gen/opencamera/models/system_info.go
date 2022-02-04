// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SystemInfo system info
//
// swagger:model SystemInfo
type SystemInfo struct {

	// firmware version
	// Example: 1.0.1
	FirmwareVersion string `json:"firmwareVersion,omitempty"`

	// live session config
	LiveSessionConfig *LiveSessionConfig `json:"live_session_config,omitempty"`

	// free sd card disk in GB
	// Example: 10
	SdcardFree int64 `json:"sdcard_free,omitempty"`

	// total sd card disk in GB
	// Example: 10
	SdcardTotal int64 `json:"sdcard_total,omitempty"`
}

// Validate validates this system info
func (m *SystemInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLiveSessionConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SystemInfo) validateLiveSessionConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.LiveSessionConfig) { // not required
		return nil
	}

	if m.LiveSessionConfig != nil {
		if err := m.LiveSessionConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("live_session_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this system info based on the context it is used
func (m *SystemInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLiveSessionConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SystemInfo) contextValidateLiveSessionConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.LiveSessionConfig != nil {
		if err := m.LiveSessionConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("live_session_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SystemInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemInfo) UnmarshalBinary(b []byte) error {
	var res SystemInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
