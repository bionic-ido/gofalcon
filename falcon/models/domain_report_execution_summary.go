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

// DomainReportExecutionSummary domain report execution summary
//
// swagger:model domain.ReportExecutionSummary
type DomainReportExecutionSummary struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// last updated on
	// Required: true
	// Format: date-time
	LastUpdatedOn *strfmt.DateTime `json:"last_updated_on"`

	// report file reference
	// Required: true
	ReportFileReference *string `json:"report_file_reference"`

	// result metadata
	// Required: true
	ResultMetadata *DomainResultMetadata `json:"result_metadata"`

	// status
	// Required: true
	Status *string `json:"status"`

	// status msg
	// Required: true
	StatusMsg *string `json:"status_msg"`
}

// Validate validates this domain report execution summary
func (m *DomainReportExecutionSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportFileReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResultMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusMsg(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainReportExecutionSummary) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DomainReportExecutionSummary) validateLastUpdatedOn(formats strfmt.Registry) error {

	if err := validate.Required("last_updated_on", "body", m.LastUpdatedOn); err != nil {
		return err
	}

	if err := validate.FormatOf("last_updated_on", "body", "date-time", m.LastUpdatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainReportExecutionSummary) validateReportFileReference(formats strfmt.Registry) error {

	if err := validate.Required("report_file_reference", "body", m.ReportFileReference); err != nil {
		return err
	}

	return nil
}

func (m *DomainReportExecutionSummary) validateResultMetadata(formats strfmt.Registry) error {

	if err := validate.Required("result_metadata", "body", m.ResultMetadata); err != nil {
		return err
	}

	if m.ResultMetadata != nil {
		if err := m.ResultMetadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result_metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("result_metadata")
			}
			return err
		}
	}

	return nil
}

func (m *DomainReportExecutionSummary) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *DomainReportExecutionSummary) validateStatusMsg(formats strfmt.Registry) error {

	if err := validate.Required("status_msg", "body", m.StatusMsg); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this domain report execution summary based on the context it is used
func (m *DomainReportExecutionSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResultMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainReportExecutionSummary) contextValidateResultMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.ResultMetadata != nil {

		if err := m.ResultMetadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result_metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("result_metadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainReportExecutionSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainReportExecutionSummary) UnmarshalBinary(b []byte) error {
	var res DomainReportExecutionSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
