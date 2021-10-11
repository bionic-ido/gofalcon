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

// FalconxReportV1 falconx report v1
//
// swagger:model falconx.ReportV1
type FalconxReportV1 struct {

	// cid
	Cid string `json:"cid,omitempty"`

	// created timestamp
	CreatedTimestamp string `json:"created_timestamp,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// intel
	Intel []*FalconxIntelReportV1 `json:"intel"`

	// ioc report broad csv artifact id
	IocReportBroadCsvArtifactID string `json:"ioc_report_broad_csv_artifact_id,omitempty"`

	// ioc report broad json artifact id
	IocReportBroadJSONArtifactID string `json:"ioc_report_broad_json_artifact_id,omitempty"`

	// ioc report broad maec artifact id
	IocReportBroadMaecArtifactID string `json:"ioc_report_broad_maec_artifact_id,omitempty"`

	// ioc report broad stix artifact id
	IocReportBroadStixArtifactID string `json:"ioc_report_broad_stix_artifact_id,omitempty"`

	// ioc report strict csv artifact id
	IocReportStrictCsvArtifactID string `json:"ioc_report_strict_csv_artifact_id,omitempty"`

	// ioc report strict json artifact id
	IocReportStrictJSONArtifactID string `json:"ioc_report_strict_json_artifact_id,omitempty"`

	// ioc report strict maec artifact id
	IocReportStrictMaecArtifactID string `json:"ioc_report_strict_maec_artifact_id,omitempty"`

	// ioc report strict stix artifact id
	IocReportStrictStixArtifactID string `json:"ioc_report_strict_stix_artifact_id,omitempty"`

	// malquery
	Malquery []*FalconxMalqueryReportV1 `json:"malquery"`

	// origin
	Origin string `json:"origin,omitempty"`

	// sandbox
	Sandbox []*FalconxSandboxReportV1 `json:"sandbox"`

	// tags
	Tags []string `json:"tags"`

	// threat graph
	ThreatGraph *FalconxThreatGraphReportV1 `json:"threat_graph,omitempty"`

	// user id
	UserID string `json:"user_id,omitempty"`

	// user name
	UserName string `json:"user_name,omitempty"`

	// user tags
	UserTags []string `json:"user_tags"`

	// user uuid
	UserUUID string `json:"user_uuid,omitempty"`

	// verdict
	Verdict string `json:"verdict,omitempty"`
}

// Validate validates this falconx report v1
func (m *FalconxReportV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIntel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMalquery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSandbox(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThreatGraph(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FalconxReportV1) validateIntel(formats strfmt.Registry) error {
	if swag.IsZero(m.Intel) { // not required
		return nil
	}

	for i := 0; i < len(m.Intel); i++ {
		if swag.IsZero(m.Intel[i]) { // not required
			continue
		}

		if m.Intel[i] != nil {
			if err := m.Intel[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("intel" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("intel" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxReportV1) validateMalquery(formats strfmt.Registry) error {
	if swag.IsZero(m.Malquery) { // not required
		return nil
	}

	for i := 0; i < len(m.Malquery); i++ {
		if swag.IsZero(m.Malquery[i]) { // not required
			continue
		}

		if m.Malquery[i] != nil {
			if err := m.Malquery[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("malquery" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("malquery" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxReportV1) validateSandbox(formats strfmt.Registry) error {
	if swag.IsZero(m.Sandbox) { // not required
		return nil
	}

	for i := 0; i < len(m.Sandbox); i++ {
		if swag.IsZero(m.Sandbox[i]) { // not required
			continue
		}

		if m.Sandbox[i] != nil {
			if err := m.Sandbox[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sandbox" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sandbox" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxReportV1) validateThreatGraph(formats strfmt.Registry) error {
	if swag.IsZero(m.ThreatGraph) { // not required
		return nil
	}

	if m.ThreatGraph != nil {
		if err := m.ThreatGraph.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("threat_graph")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("threat_graph")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this falconx report v1 based on the context it is used
func (m *FalconxReportV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIntel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMalquery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSandbox(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThreatGraph(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FalconxReportV1) contextValidateIntel(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Intel); i++ {

		if m.Intel[i] != nil {
			if err := m.Intel[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("intel" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("intel" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxReportV1) contextValidateMalquery(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Malquery); i++ {

		if m.Malquery[i] != nil {
			if err := m.Malquery[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("malquery" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("malquery" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxReportV1) contextValidateSandbox(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Sandbox); i++ {

		if m.Sandbox[i] != nil {
			if err := m.Sandbox[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sandbox" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sandbox" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxReportV1) contextValidateThreatGraph(ctx context.Context, formats strfmt.Registry) error {

	if m.ThreatGraph != nil {
		if err := m.ThreatGraph.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("threat_graph")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("threat_graph")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FalconxReportV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FalconxReportV1) UnmarshalBinary(b []byte) error {
	var res FalconxReportV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
