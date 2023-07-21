// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// QuarantineQuarantinedFile quarantine quarantined file
//
// swagger:model quarantine.QuarantinedFile
type QuarantineQuarantinedFile struct {

	// aid
	Aid string `json:"aid,omitempty"`

	// cid
	Cid string `json:"cid,omitempty"`

	// date created
	DateCreated string `json:"date_created,omitempty"`

	// date updated
	DateUpdated string `json:"date_updated,omitempty"`

	// detect ids
	DetectIds []string `json:"detect_ids"`

	// extracted
	Extracted bool `json:"extracted,omitempty"`

	// hostname
	Hostname string `json:"hostname,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// paths
	Paths []*QuarantineQuarantinedFilePath `json:"paths"`

	// primary module
	PrimaryModule bool `json:"primary_module,omitempty"`

	// sandbox report id
	SandboxReportID string `json:"sandbox_report_id,omitempty"`

	// sandbox report state
	SandboxReportState string `json:"sandbox_report_state,omitempty"`

	// sha256
	Sha256 string `json:"sha256,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this quarantine quarantined file
func (m *QuarantineQuarantinedFile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaths(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuarantineQuarantinedFile) validatePaths(formats strfmt.Registry) error {
	if swag.IsZero(m.Paths) { // not required
		return nil
	}

	for i := 0; i < len(m.Paths); i++ {
		if swag.IsZero(m.Paths[i]) { // not required
			continue
		}

		if m.Paths[i] != nil {
			if err := m.Paths[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("paths" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("paths" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this quarantine quarantined file based on the context it is used
func (m *QuarantineQuarantinedFile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePaths(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuarantineQuarantinedFile) contextValidatePaths(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Paths); i++ {

		if m.Paths[i] != nil {

			if swag.IsZero(m.Paths[i]) { // not required
				return nil
			}

			if err := m.Paths[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("paths" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("paths" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *QuarantineQuarantinedFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuarantineQuarantinedFile) UnmarshalBinary(b []byte) error {
	var res QuarantineQuarantinedFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
