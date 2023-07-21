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

// DomainScanProfile domain scan profile
//
// swagger:model domain.ScanProfile
type DomainScanProfile struct {

	// cid
	Cid string `json:"cid,omitempty"`

	// cloud ml level detection
	CloudMlLevelDetection int32 `json:"cloud_ml_level_detection,omitempty"`

	// cloud ml level prevention
	CloudMlLevelPrevention int32 `json:"cloud_ml_level_prevention,omitempty"`

	// cpu priority
	CPUPriority int32 `json:"cpu_priority,omitempty"`

	// created by
	CreatedBy string `json:"created_by,omitempty"`

	// created on
	// Format: date-time
	CreatedOn strfmt.DateTime `json:"created_on,omitempty"`

	// deleted
	// Required: true
	Deleted *bool `json:"deleted"`

	// description
	Description string `json:"description,omitempty"`

	// endpoint notification
	EndpointNotification bool `json:"endpoint_notification,omitempty"`

	// file paths
	FilePaths []string `json:"file_paths"`

	// host groups
	HostGroups []string `json:"host_groups"`

	// hosts
	Hosts []string `json:"hosts"`

	// id
	// Required: true
	ID *string `json:"id"`

	// initiated from
	InitiatedFrom string `json:"initiated_from,omitempty"`

	// last updated
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// max duration
	MaxDuration int32 `json:"max_duration,omitempty"`

	// max file size
	MaxFileSize int32 `json:"max_file_size,omitempty"`

	// metadata
	Metadata []*DomainScanProfileMetadata `json:"metadata"`

	// pause duration
	PauseDuration int32 `json:"pause_duration,omitempty"`

	// policy setting
	PolicySetting []int64 `json:"policy_setting"`

	// preemption priority
	PreemptionPriority int32 `json:"preemption_priority,omitempty"`

	// quarantine
	Quarantine bool `json:"quarantine,omitempty"`

	// scan exclusions
	ScanExclusions []string `json:"scan_exclusions"`

	// scan inclusions
	ScanInclusions []string `json:"scan_inclusions"`

	// schedule
	Schedule *DomainSchedule `json:"schedule,omitempty"`

	// sensor ml level detection
	SensorMlLevelDetection int32 `json:"sensor_ml_level_detection,omitempty"`

	// sensor ml level prevention
	SensorMlLevelPrevention int32 `json:"sensor_ml_level_prevention,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this domain scan profile
func (m *DomainScanProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainScanProfile) validateCreatedOn(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainScanProfile) validateDeleted(formats strfmt.Registry) error {

	if err := validate.Required("deleted", "body", m.Deleted); err != nil {
		return err
	}

	return nil
}

func (m *DomainScanProfile) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DomainScanProfile) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainScanProfile) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	for i := 0; i < len(m.Metadata); i++ {
		if swag.IsZero(m.Metadata[i]) { // not required
			continue
		}

		if m.Metadata[i] != nil {
			if err := m.Metadata[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metadata" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("metadata" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainScanProfile) validateSchedule(formats strfmt.Registry) error {
	if swag.IsZero(m.Schedule) { // not required
		return nil
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

// ContextValidate validate this domain scan profile based on the context it is used
func (m *DomainScanProfile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
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

func (m *DomainScanProfile) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Metadata); i++ {

		if m.Metadata[i] != nil {

			if swag.IsZero(m.Metadata[i]) { // not required
				return nil
			}

			if err := m.Metadata[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metadata" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("metadata" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainScanProfile) contextValidateSchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.Schedule != nil {

		if swag.IsZero(m.Schedule) { // not required
			return nil
		}

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
func (m *DomainScanProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainScanProfile) UnmarshalBinary(b []byte) error {
	var res DomainScanProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
