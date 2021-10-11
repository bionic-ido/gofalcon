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

// DomainPublicIndicatorV3 domain public indicator v3
//
// swagger:model domain.PublicIndicatorV3
type DomainPublicIndicatorV3 struct {

	// marker
	// Required: true
	Marker *string `json:"_marker"`

	// actors
	// Required: true
	Actors []string `json:"actors"`

	// deleted
	// Required: true
	Deleted *bool `json:"deleted"`

	// domain types
	// Required: true
	DomainTypes []string `json:"domain_types"`

	// id
	// Required: true
	ID *string `json:"id"`

	// indicator
	// Required: true
	Indicator *string `json:"indicator"`

	// ip address types
	// Required: true
	IPAddressTypes []string `json:"ip_address_types"`

	// kill chains
	// Required: true
	KillChains []string `json:"kill_chains"`

	// labels
	// Required: true
	Labels []*DomainCSIXLabel `json:"labels"`

	// last updated
	// Required: true
	LastUpdated *int64 `json:"last_updated"`

	// malicious confidence
	// Required: true
	MaliciousConfidence *string `json:"malicious_confidence"`

	// malware families
	// Required: true
	MalwareFamilies []string `json:"malware_families"`

	// published date
	// Required: true
	PublishedDate *int64 `json:"published_date"`

	// relations
	// Required: true
	Relations []*DomainCSIXRelation `json:"relations"`

	// reports
	// Required: true
	Reports []string `json:"reports"`

	// targets
	// Required: true
	Targets []string `json:"targets"`

	// threat types
	// Required: true
	ThreatTypes []string `json:"threat_types"`

	// type
	// Required: true
	Type *string `json:"type"`

	// vulnerabilities
	// Required: true
	Vulnerabilities []string `json:"vulnerabilities"`
}

// Validate validates this domain public indicator v3
func (m *DomainPublicIndicatorV3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMarker(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomainTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndicator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPAddressTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKillChains(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaliciousConfidence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMalwareFamilies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublishedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReports(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThreatTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVulnerabilities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainPublicIndicatorV3) validateMarker(formats strfmt.Registry) error {

	if err := validate.Required("_marker", "body", m.Marker); err != nil {
		return err
	}

	return nil
}

func (m *DomainPublicIndicatorV3) validateActors(formats strfmt.Registry) error {

	if err := validate.Required("actors", "body", m.Actors); err != nil {
		return err
	}

	return nil
}

func (m *DomainPublicIndicatorV3) validateDeleted(formats strfmt.Registry) error {

	if err := validate.Required("deleted", "body", m.Deleted); err != nil {
		return err
	}

	return nil
}

func (m *DomainPublicIndicatorV3) validateDomainTypes(formats strfmt.Registry) error {

	if err := validate.Required("domain_types", "body", m.DomainTypes); err != nil {
		return err
	}

	return nil
}

func (m *DomainPublicIndicatorV3) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DomainPublicIndicatorV3) validateIndicator(formats strfmt.Registry) error {

	if err := validate.Required("indicator", "body", m.Indicator); err != nil {
		return err
	}

	return nil
}

func (m *DomainPublicIndicatorV3) validateIPAddressTypes(formats strfmt.Registry) error {

	if err := validate.Required("ip_address_types", "body", m.IPAddressTypes); err != nil {
		return err
	}

	return nil
}

func (m *DomainPublicIndicatorV3) validateKillChains(formats strfmt.Registry) error {

	if err := validate.Required("kill_chains", "body", m.KillChains); err != nil {
		return err
	}

	return nil
}

func (m *DomainPublicIndicatorV3) validateLabels(formats strfmt.Registry) error {

	if err := validate.Required("labels", "body", m.Labels); err != nil {
		return err
	}

	for i := 0; i < len(m.Labels); i++ {
		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if m.Labels[i] != nil {
			if err := m.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainPublicIndicatorV3) validateLastUpdated(formats strfmt.Registry) error {

	if err := validate.Required("last_updated", "body", m.LastUpdated); err != nil {
		return err
	}

	return nil
}

func (m *DomainPublicIndicatorV3) validateMaliciousConfidence(formats strfmt.Registry) error {

	if err := validate.Required("malicious_confidence", "body", m.MaliciousConfidence); err != nil {
		return err
	}

	return nil
}

func (m *DomainPublicIndicatorV3) validateMalwareFamilies(formats strfmt.Registry) error {

	if err := validate.Required("malware_families", "body", m.MalwareFamilies); err != nil {
		return err
	}

	return nil
}

func (m *DomainPublicIndicatorV3) validatePublishedDate(formats strfmt.Registry) error {

	if err := validate.Required("published_date", "body", m.PublishedDate); err != nil {
		return err
	}

	return nil
}

func (m *DomainPublicIndicatorV3) validateRelations(formats strfmt.Registry) error {

	if err := validate.Required("relations", "body", m.Relations); err != nil {
		return err
	}

	for i := 0; i < len(m.Relations); i++ {
		if swag.IsZero(m.Relations[i]) { // not required
			continue
		}

		if m.Relations[i] != nil {
			if err := m.Relations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("relations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainPublicIndicatorV3) validateReports(formats strfmt.Registry) error {

	if err := validate.Required("reports", "body", m.Reports); err != nil {
		return err
	}

	return nil
}

func (m *DomainPublicIndicatorV3) validateTargets(formats strfmt.Registry) error {

	if err := validate.Required("targets", "body", m.Targets); err != nil {
		return err
	}

	return nil
}

func (m *DomainPublicIndicatorV3) validateThreatTypes(formats strfmt.Registry) error {

	if err := validate.Required("threat_types", "body", m.ThreatTypes); err != nil {
		return err
	}

	return nil
}

func (m *DomainPublicIndicatorV3) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *DomainPublicIndicatorV3) validateVulnerabilities(formats strfmt.Registry) error {

	if err := validate.Required("vulnerabilities", "body", m.Vulnerabilities); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this domain public indicator v3 based on the context it is used
func (m *DomainPublicIndicatorV3) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRelations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainPublicIndicatorV3) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Labels); i++ {

		if m.Labels[i] != nil {
			if err := m.Labels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainPublicIndicatorV3) contextValidateRelations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Relations); i++ {

		if m.Relations[i] != nil {
			if err := m.Relations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("relations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainPublicIndicatorV3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainPublicIndicatorV3) UnmarshalBinary(b []byte) error {
	var res DomainPublicIndicatorV3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
