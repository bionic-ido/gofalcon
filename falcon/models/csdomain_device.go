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

// CsdomainDevice csdomain device
//
// swagger:model csdomain.Device
type CsdomainDevice struct {

	// agent load flags
	AgentLoadFlags string `json:"agent_load_flags,omitempty"`

	// agent local time
	AgentLocalTime string `json:"agent_local_time,omitempty"`

	// agent version
	AgentVersion string `json:"agent_version,omitempty"`

	// bios manufacturer
	BiosManufacturer string `json:"bios_manufacturer,omitempty"`

	// bios version
	BiosVersion string `json:"bios_version,omitempty"`

	// build number
	BuildNumber string `json:"build_number,omitempty"`

	// cid
	// Required: true
	Cid *string `json:"cid"`

	// config id base
	ConfigIDBase string `json:"config_id_base,omitempty"`

	// config id build
	ConfigIDBuild string `json:"config_id_build,omitempty"`

	// config id platform
	ConfigIDPlatform string `json:"config_id_platform,omitempty"`

	// detection suppression status
	DetectionSuppressionStatus string `json:"detection_suppression_status,omitempty"`

	// device id
	// Required: true
	DeviceID *string `json:"device_id"`

	// device policies
	DevicePolicies *CsdomainMappedDevicePolicies `json:"device_policies,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// external ip
	ExternalIP string `json:"external_ip,omitempty"`

	// first login timestamp
	FirstLoginTimestamp string `json:"first_login_timestamp,omitempty"`

	// first login user
	FirstLoginUser string `json:"first_login_user,omitempty"`

	// first seen
	FirstSeen string `json:"first_seen,omitempty"`

	// group hash
	GroupHash string `json:"group_hash,omitempty"`

	// groups
	Groups []string `json:"groups"`

	// host hidden status
	HostHiddenStatus string `json:"host_hidden_status,omitempty"`

	// hostname
	Hostname string `json:"hostname,omitempty"`

	// instance id
	InstanceID string `json:"instance_id,omitempty"`

	// last login timestamp
	LastLoginTimestamp string `json:"last_login_timestamp,omitempty"`

	// last login user
	LastLoginUser string `json:"last_login_user,omitempty"`

	// last seen
	LastSeen string `json:"last_seen,omitempty"`

	// local ip
	LocalIP string `json:"local_ip,omitempty"`

	// mac address
	MacAddress string `json:"mac_address,omitempty"`

	// machine domain
	MachineDomain string `json:"machine_domain,omitempty"`

	// major version
	MajorVersion string `json:"major_version,omitempty"`

	// managed apps
	ManagedApps *CsdomainManagedApps `json:"managed_apps,omitempty"`

	// meta
	Meta *CsdomainDeviceMeta `json:"meta,omitempty"`

	// minor version
	MinorVersion string `json:"minor_version,omitempty"`

	// modified timestamp
	ModifiedTimestamp string `json:"modified_timestamp,omitempty"`

	// notes
	Notes []string `json:"notes"`

	// os version
	OsVersion string `json:"os_version,omitempty"`

	// ou
	Ou []string `json:"ou"`

	// platform id
	PlatformID string `json:"platform_id,omitempty"`

	// platform name
	PlatformName string `json:"platform_name,omitempty"`

	// pointer size
	PointerSize string `json:"pointer_size,omitempty"`

	// policies
	Policies []*CsdomainDevicePolicy `json:"policies"`

	// product type
	ProductType string `json:"product_type,omitempty"`

	// product type desc
	ProductTypeDesc string `json:"product_type_desc,omitempty"`

	// provision status
	ProvisionStatus string `json:"provision_status,omitempty"`

	// release group
	ReleaseGroup string `json:"release_group,omitempty"`

	// service pack major
	ServicePackMajor string `json:"service_pack_major,omitempty"`

	// service pack minor
	ServicePackMinor string `json:"service_pack_minor,omitempty"`

	// service provider
	ServiceProvider string `json:"service_provider,omitempty"`

	// service provider account id
	ServiceProviderAccountID string `json:"service_provider_account_id,omitempty"`

	// site name
	SiteName string `json:"site_name,omitempty"`

	// slow changing modified timestamp
	SlowChangingModifiedTimestamp string `json:"slow_changing_modified_timestamp,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// system manufacturer
	SystemManufacturer string `json:"system_manufacturer,omitempty"`

	// system product name
	SystemProductName string `json:"system_product_name,omitempty"`

	// tags
	Tags []string `json:"tags"`
}

// Validate validates this csdomain device
func (m *CsdomainDevice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevicePolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManagedApps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CsdomainDevice) validateCid(formats strfmt.Registry) error {

	if err := validate.Required("cid", "body", m.Cid); err != nil {
		return err
	}

	return nil
}

func (m *CsdomainDevice) validateDeviceID(formats strfmt.Registry) error {

	if err := validate.Required("device_id", "body", m.DeviceID); err != nil {
		return err
	}

	return nil
}

func (m *CsdomainDevice) validateDevicePolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.DevicePolicies) { // not required
		return nil
	}

	if m.DevicePolicies != nil {
		if err := m.DevicePolicies.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_policies")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device_policies")
			}
			return err
		}
	}

	return nil
}

func (m *CsdomainDevice) validateManagedApps(formats strfmt.Registry) error {
	if swag.IsZero(m.ManagedApps) { // not required
		return nil
	}

	if m.ManagedApps != nil {
		if err := m.ManagedApps.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("managed_apps")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("managed_apps")
			}
			return err
		}
	}

	return nil
}

func (m *CsdomainDevice) validateMeta(formats strfmt.Registry) error {
	if swag.IsZero(m.Meta) { // not required
		return nil
	}

	if m.Meta != nil {
		if err := m.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

func (m *CsdomainDevice) validatePolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.Policies) { // not required
		return nil
	}

	for i := 0; i < len(m.Policies); i++ {
		if swag.IsZero(m.Policies[i]) { // not required
			continue
		}

		if m.Policies[i] != nil {
			if err := m.Policies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("policies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this csdomain device based on the context it is used
func (m *CsdomainDevice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDevicePolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateManagedApps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMeta(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CsdomainDevice) contextValidateDevicePolicies(ctx context.Context, formats strfmt.Registry) error {

	if m.DevicePolicies != nil {

		if swag.IsZero(m.DevicePolicies) { // not required
			return nil
		}

		if err := m.DevicePolicies.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_policies")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device_policies")
			}
			return err
		}
	}

	return nil
}

func (m *CsdomainDevice) contextValidateManagedApps(ctx context.Context, formats strfmt.Registry) error {

	if m.ManagedApps != nil {

		if swag.IsZero(m.ManagedApps) { // not required
			return nil
		}

		if err := m.ManagedApps.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("managed_apps")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("managed_apps")
			}
			return err
		}
	}

	return nil
}

func (m *CsdomainDevice) contextValidateMeta(ctx context.Context, formats strfmt.Registry) error {

	if m.Meta != nil {

		if swag.IsZero(m.Meta) { // not required
			return nil
		}

		if err := m.Meta.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

func (m *CsdomainDevice) contextValidatePolicies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Policies); i++ {

		if m.Policies[i] != nil {

			if swag.IsZero(m.Policies[i]) { // not required
				return nil
			}

			if err := m.Policies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("policies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CsdomainDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CsdomainDevice) UnmarshalBinary(b []byte) error {
	var res CsdomainDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
