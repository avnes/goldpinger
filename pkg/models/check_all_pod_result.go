// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CheckAllPodResult check all pod result
// swagger:model CheckAllPodResult
type CheckAllPodResult struct {

	// host IP
	// Format: ipv4
	HostIP strfmt.IPv4 `json:"HostIP,omitempty"`

	// o k
	OK *bool `json:"OK,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// response
	Response CheckResults `json:"response,omitempty"`

	// status code
	StatusCode int32 `json:"status-code,omitempty"`
}

// Validate validates this check all pod result
func (m *CheckAllPodResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CheckAllPodResult) validateHostIP(formats strfmt.Registry) error {

	if swag.IsZero(m.HostIP) { // not required
		return nil
	}

	if err := validate.FormatOf("HostIP", "body", "ipv4", m.HostIP.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CheckAllPodResult) validateResponse(formats strfmt.Registry) error {

	if swag.IsZero(m.Response) { // not required
		return nil
	}

	if err := m.Response.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("response")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CheckAllPodResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CheckAllPodResult) UnmarshalBinary(b []byte) error {
	var res CheckAllPodResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
