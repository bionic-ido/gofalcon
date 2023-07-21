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
	"github.com/go-openapi/validate"
)

// APIScheduledReportV1 api scheduled report v1
//
// swagger:model api.ScheduledReportV1
type APIScheduledReportV1 struct {

	// can write
	// Required: true
	CanWrite *bool `json:"can_write"`

	// created on
	// Required: true
	// Format: date-time
	CreatedOn *strfmt.DateTime `json:"created_on"`

	// customer id
	// Required: true
	CustomerID *string `json:"customer_id"`

	// description
	// Required: true
	Description *string `json:"description"`

	// expiration on
	// Format: date-time
	ExpirationOn strfmt.DateTime `json:"expiration_on,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// last execution
	LastExecution *DomainReportExecutionSummary `json:"last_execution,omitempty"`

	// last updated on
	// Required: true
	// Format: date-time
	LastUpdatedOn *strfmt.DateTime `json:"last_updated_on"`

	// name
	// Required: true
	Name *string `json:"name"`

	// next execution on
	// Format: date-time
	NextExecutionOn strfmt.DateTime `json:"next_execution_on,omitempty"`

	// notifications
	// Required: true
	Notifications []*DomainNotifications `json:"notifications"`

	// report params
	// Required: true
	ReportParams *DomainReportParams `json:"report_params"`

	// schedule
	// Required: true
	Schedule *DomainSchedule `json:"schedule"`

	// shared with
	// Required: true
	SharedWith []string `json:"shared_with"`

	// start on
	// Format: date-time
	StartOn strfmt.DateTime `json:"start_on,omitempty"`

	// status
	// Required: true
	Status *string `json:"status"`

	// stop on
	// Format: date-time
	StopOn strfmt.DateTime `json:"stop_on,omitempty"`

	// tracking
	Tracking string `json:"tracking,omitempty"`

	// trigger reference
	// Required: true
	TriggerReference *string `json:"trigger_reference"`

	// type
	// Required: true
	Type *string `json:"type"`

	// user id
	// Required: true
	UserID *string `json:"user_id"`

	// user uuid
	// Required: true
	UserUUID *string `json:"user_uuid"`
}

// Validate validates this api scheduled report v1
func (m *APIScheduledReportV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCanWrite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpirationOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastExecution(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextExecutionOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotifications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSharedWith(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTriggerReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIScheduledReportV1) validateCanWrite(formats strfmt.Registry) error {

	if err := validate.Required("can_write", "body", m.CanWrite); err != nil {
		return err
	}

	return nil
}

func (m *APIScheduledReportV1) validateCreatedOn(formats strfmt.Registry) error {

	if err := validate.Required("created_on", "body", m.CreatedOn); err != nil {
		return err
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APIScheduledReportV1) validateCustomerID(formats strfmt.Registry) error {

	if err := validate.Required("customer_id", "body", m.CustomerID); err != nil {
		return err
	}

	return nil
}

func (m *APIScheduledReportV1) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *APIScheduledReportV1) validateExpirationOn(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpirationOn) { // not required
		return nil
	}

	if err := validate.FormatOf("expiration_on", "body", "date-time", m.ExpirationOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APIScheduledReportV1) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *APIScheduledReportV1) validateLastExecution(formats strfmt.Registry) error {
	if swag.IsZero(m.LastExecution) { // not required
		return nil
	}

	if m.LastExecution != nil {
		if err := m.LastExecution.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_execution")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last_execution")
			}
			return err
		}
	}

	return nil
}

func (m *APIScheduledReportV1) validateLastUpdatedOn(formats strfmt.Registry) error {

	if err := validate.Required("last_updated_on", "body", m.LastUpdatedOn); err != nil {
		return err
	}

	if err := validate.FormatOf("last_updated_on", "body", "date-time", m.LastUpdatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APIScheduledReportV1) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *APIScheduledReportV1) validateNextExecutionOn(formats strfmt.Registry) error {
	if swag.IsZero(m.NextExecutionOn) { // not required
		return nil
	}

	if err := validate.FormatOf("next_execution_on", "body", "date-time", m.NextExecutionOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APIScheduledReportV1) validateNotifications(formats strfmt.Registry) error {

	if err := validate.Required("notifications", "body", m.Notifications); err != nil {
		return err
	}

	for i := 0; i < len(m.Notifications); i++ {
		if swag.IsZero(m.Notifications[i]) { // not required
			continue
		}

		if m.Notifications[i] != nil {
			if err := m.Notifications[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("notifications" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("notifications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APIScheduledReportV1) validateReportParams(formats strfmt.Registry) error {

	if err := validate.Required("report_params", "body", m.ReportParams); err != nil {
		return err
	}

	if m.ReportParams != nil {
		if err := m.ReportParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("report_params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("report_params")
			}
			return err
		}
	}

	return nil
}

func (m *APIScheduledReportV1) validateSchedule(formats strfmt.Registry) error {

	if err := validate.Required("schedule", "body", m.Schedule); err != nil {
		return err
	}

	if m.Schedule != nil {
		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

func (m *APIScheduledReportV1) validateSharedWith(formats strfmt.Registry) error {

	if err := validate.Required("shared_with", "body", m.SharedWith); err != nil {
		return err
	}

	return nil
}

func (m *APIScheduledReportV1) validateStartOn(formats strfmt.Registry) error {
	if swag.IsZero(m.StartOn) { // not required
		return nil
	}

	if err := validate.FormatOf("start_on", "body", "date-time", m.StartOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APIScheduledReportV1) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *APIScheduledReportV1) validateStopOn(formats strfmt.Registry) error {
	if swag.IsZero(m.StopOn) { // not required
		return nil
	}

	if err := validate.FormatOf("stop_on", "body", "date-time", m.StopOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APIScheduledReportV1) validateTriggerReference(formats strfmt.Registry) error {

	if err := validate.Required("trigger_reference", "body", m.TriggerReference); err != nil {
		return err
	}

	return nil
}

func (m *APIScheduledReportV1) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *APIScheduledReportV1) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("user_id", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

func (m *APIScheduledReportV1) validateUserUUID(formats strfmt.Registry) error {

	if err := validate.Required("user_uuid", "body", m.UserUUID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this api scheduled report v1 based on the context it is used
func (m *APIScheduledReportV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLastExecution(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNotifications(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReportParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIScheduledReportV1) contextValidateLastExecution(ctx context.Context, formats strfmt.Registry) error {

	if m.LastExecution != nil {

		if swag.IsZero(m.LastExecution) { // not required
			return nil
		}

		if err := m.LastExecution.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_execution")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last_execution")
			}
			return err
		}
	}

	return nil
}

func (m *APIScheduledReportV1) contextValidateNotifications(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Notifications); i++ {

		if m.Notifications[i] != nil {

			if swag.IsZero(m.Notifications[i]) { // not required
				return nil
			}

			if err := m.Notifications[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("notifications" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("notifications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APIScheduledReportV1) contextValidateReportParams(ctx context.Context, formats strfmt.Registry) error {

	if m.ReportParams != nil {

		if err := m.ReportParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("report_params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("report_params")
			}
			return err
		}
	}

	return nil
}

func (m *APIScheduledReportV1) contextValidateSchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.Schedule != nil {

		if err := m.Schedule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIScheduledReportV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIScheduledReportV1) UnmarshalBinary(b []byte) error {
	var res APIScheduledReportV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
