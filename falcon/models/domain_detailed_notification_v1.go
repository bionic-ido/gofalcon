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

// DomainDetailedNotificationV1 domain detailed notification v1
//
// swagger:model domain.DetailedNotificationV1
type DomainDetailedNotificationV1 struct {

	// details
	// Required: true
	Details *DomainNotificationDetailsV1 `json:"details"`

	// notification
	// Required: true
	Notification *DomainNotificationV1 `json:"notification"`
}

// Validate validates this domain detailed notification v1
func (m *DomainDetailedNotificationV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotification(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainDetailedNotificationV1) validateDetails(formats strfmt.Registry) error {

	if err := validate.Required("details", "body", m.Details); err != nil {
		return err
	}

	if m.Details != nil {
		if err := m.Details.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("details")
			}
			return err
		}
	}

	return nil
}

func (m *DomainDetailedNotificationV1) validateNotification(formats strfmt.Registry) error {

	if err := validate.Required("notification", "body", m.Notification); err != nil {
		return err
	}

	if m.Notification != nil {
		if err := m.Notification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notification")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("notification")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this domain detailed notification v1 based on the context it is used
func (m *DomainDetailedNotificationV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNotification(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainDetailedNotificationV1) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.Details != nil {
		if err := m.Details.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("details")
			}
			return err
		}
	}

	return nil
}

func (m *DomainDetailedNotificationV1) contextValidateNotification(ctx context.Context, formats strfmt.Registry) error {

	if m.Notification != nil {
		if err := m.Notification.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notification")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("notification")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainDetailedNotificationV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainDetailedNotificationV1) UnmarshalBinary(b []byte) error {
	var res DomainDetailedNotificationV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
