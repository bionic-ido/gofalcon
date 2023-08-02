// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IoaEventAggregate ioa event aggregate
//
// swagger:model ioa.EventAggregate
type IoaEventAggregate struct {

	// cid severity
	CidSeverity int64 `json:"cid_severity,omitempty"`

	// confidence
	Confidence int32 `json:"confidence,omitempty"`

	// count
	Count int32 `json:"count,omitempty"`

	// events
	Events []string `json:"events"`

	// first timestamp
	FirstTimestamp string `json:"first_timestamp,omitempty"`

	// join keys
	JoinKeys []string `json:"join_keys"`

	// last timestamp
	LastTimestamp string `json:"last_timestamp,omitempty"`

	// resource
	// Required: true
	Resource *Resource `json:"resource"`

	// score
	Score int32 `json:"score,omitempty"`

	// threatintel
	Threatintel *DetectionAggregateThreatIntel `json:"threatintel,omitempty"`

	// timestamps
	Timestamps []string `json:"timestamps"`
}

// Validate validates this ioa event aggregate
func (m *IoaEventAggregate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThreatintel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoaEventAggregate) validateResource(formats strfmt.Registry) error {

	if err := validate.Required("resource", "body", m.Resource); err != nil {
		return err
	}

	if m.Resource != nil {
		if err := m.Resource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

func (m *IoaEventAggregate) validateThreatintel(formats strfmt.Registry) error {
	if swag.IsZero(m.Threatintel) { // not required
		return nil
	}

	if m.Threatintel != nil {
		if err := m.Threatintel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("threatintel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("threatintel")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ioa event aggregate based on the context it is used
func (m *IoaEventAggregate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThreatintel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoaEventAggregate) contextValidateResource(ctx context.Context, formats strfmt.Registry) error {

	if m.Resource != nil {

		if err := m.Resource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

func (m *IoaEventAggregate) contextValidateThreatintel(ctx context.Context, formats strfmt.Registry) error {

	if m.Threatintel != nil {

		if swag.IsZero(m.Threatintel) { // not required
			return nil
		}

		if err := m.Threatintel.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("threatintel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("threatintel")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoaEventAggregate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoaEventAggregate) UnmarshalBinary(b []byte) error {
	var res IoaEventAggregate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
