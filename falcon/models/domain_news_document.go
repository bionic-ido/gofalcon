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

// DomainNewsDocument domain news document
//
// swagger:model domain.NewsDocument
type DomainNewsDocument struct {

	// active
	Active bool `json:"active,omitempty"`

	// actors
	// Required: true
	Actors []*DomainSimpleActor `json:"actors"`

	// attachments
	Attachments []*DomainFile `json:"attachments"`

	// created date
	// Required: true
	CreatedDate *int64 `json:"created_date"`

	// description
	Description string `json:"description,omitempty"`

	// entitlements
	Entitlements []*DomainEntity `json:"entitlements"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// image
	Image *DomainImage `json:"image,omitempty"`

	// last modified date
	// Required: true
	LastModifiedDate *int64 `json:"last_modified_date"`

	// motivations
	// Required: true
	Motivations []*DomainEntity `json:"motivations"`

	// name
	// Required: true
	Name *string `json:"name"`

	// notify users
	NotifyUsers bool `json:"notify_users,omitempty"`

	// rich text description
	RichTextDescription string `json:"rich_text_description,omitempty"`

	// short description
	ShortDescription string `json:"short_description,omitempty"`

	// slug
	// Required: true
	Slug *string `json:"slug"`

	// sub type
	SubType *DomainEntity `json:"sub_type,omitempty"`

	// tags
	// Required: true
	Tags []*DomainEntity `json:"tags"`

	// target countries
	// Required: true
	TargetCountries []*DomainEntity `json:"target_countries"`

	// target industries
	// Required: true
	TargetIndustries []*DomainEntity `json:"target_industries"`

	// thumbnail
	// Required: true
	Thumbnail *DomainImage `json:"thumbnail"`

	// topic
	Topic *DomainEntity `json:"topic,omitempty"`

	// type
	Type *DomainEntity `json:"type,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this domain news document
func (m *DomainNewsDocument) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttachments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntitlements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMotivations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlug(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetCountries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetIndustries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThumbnail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainNewsDocument) validateActors(formats strfmt.Registry) error {

	if err := validate.Required("actors", "body", m.Actors); err != nil {
		return err
	}

	for i := 0; i < len(m.Actors); i++ {
		if swag.IsZero(m.Actors[i]) { // not required
			continue
		}

		if m.Actors[i] != nil {
			if err := m.Actors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("actors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("actors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainNewsDocument) validateAttachments(formats strfmt.Registry) error {
	if swag.IsZero(m.Attachments) { // not required
		return nil
	}

	for i := 0; i < len(m.Attachments); i++ {
		if swag.IsZero(m.Attachments[i]) { // not required
			continue
		}

		if m.Attachments[i] != nil {
			if err := m.Attachments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attachments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attachments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainNewsDocument) validateCreatedDate(formats strfmt.Registry) error {

	if err := validate.Required("created_date", "body", m.CreatedDate); err != nil {
		return err
	}

	return nil
}

func (m *DomainNewsDocument) validateEntitlements(formats strfmt.Registry) error {
	if swag.IsZero(m.Entitlements) { // not required
		return nil
	}

	for i := 0; i < len(m.Entitlements); i++ {
		if swag.IsZero(m.Entitlements[i]) { // not required
			continue
		}

		if m.Entitlements[i] != nil {
			if err := m.Entitlements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entitlements" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("entitlements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainNewsDocument) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DomainNewsDocument) validateImage(formats strfmt.Registry) error {
	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if m.Image != nil {
		if err := m.Image.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *DomainNewsDocument) validateLastModifiedDate(formats strfmt.Registry) error {

	if err := validate.Required("last_modified_date", "body", m.LastModifiedDate); err != nil {
		return err
	}

	return nil
}

func (m *DomainNewsDocument) validateMotivations(formats strfmt.Registry) error {

	if err := validate.Required("motivations", "body", m.Motivations); err != nil {
		return err
	}

	for i := 0; i < len(m.Motivations); i++ {
		if swag.IsZero(m.Motivations[i]) { // not required
			continue
		}

		if m.Motivations[i] != nil {
			if err := m.Motivations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("motivations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("motivations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainNewsDocument) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DomainNewsDocument) validateSlug(formats strfmt.Registry) error {

	if err := validate.Required("slug", "body", m.Slug); err != nil {
		return err
	}

	return nil
}

func (m *DomainNewsDocument) validateSubType(formats strfmt.Registry) error {
	if swag.IsZero(m.SubType) { // not required
		return nil
	}

	if m.SubType != nil {
		if err := m.SubType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sub_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sub_type")
			}
			return err
		}
	}

	return nil
}

func (m *DomainNewsDocument) validateTags(formats strfmt.Registry) error {

	if err := validate.Required("tags", "body", m.Tags); err != nil {
		return err
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainNewsDocument) validateTargetCountries(formats strfmt.Registry) error {

	if err := validate.Required("target_countries", "body", m.TargetCountries); err != nil {
		return err
	}

	for i := 0; i < len(m.TargetCountries); i++ {
		if swag.IsZero(m.TargetCountries[i]) { // not required
			continue
		}

		if m.TargetCountries[i] != nil {
			if err := m.TargetCountries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("target_countries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("target_countries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainNewsDocument) validateTargetIndustries(formats strfmt.Registry) error {

	if err := validate.Required("target_industries", "body", m.TargetIndustries); err != nil {
		return err
	}

	for i := 0; i < len(m.TargetIndustries); i++ {
		if swag.IsZero(m.TargetIndustries[i]) { // not required
			continue
		}

		if m.TargetIndustries[i] != nil {
			if err := m.TargetIndustries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("target_industries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("target_industries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainNewsDocument) validateThumbnail(formats strfmt.Registry) error {

	if err := validate.Required("thumbnail", "body", m.Thumbnail); err != nil {
		return err
	}

	if m.Thumbnail != nil {
		if err := m.Thumbnail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("thumbnail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("thumbnail")
			}
			return err
		}
	}

	return nil
}

func (m *DomainNewsDocument) validateTopic(formats strfmt.Registry) error {
	if swag.IsZero(m.Topic) { // not required
		return nil
	}

	if m.Topic != nil {
		if err := m.Topic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("topic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("topic")
			}
			return err
		}
	}

	return nil
}

func (m *DomainNewsDocument) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this domain news document based on the context it is used
func (m *DomainNewsDocument) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAttachments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntitlements(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMotivations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTargetCountries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTargetIndustries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThumbnail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTopic(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainNewsDocument) contextValidateActors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Actors); i++ {

		if m.Actors[i] != nil {
			if err := m.Actors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("actors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("actors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainNewsDocument) contextValidateAttachments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Attachments); i++ {

		if m.Attachments[i] != nil {
			if err := m.Attachments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attachments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attachments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainNewsDocument) contextValidateEntitlements(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Entitlements); i++ {

		if m.Entitlements[i] != nil {
			if err := m.Entitlements[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entitlements" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("entitlements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainNewsDocument) contextValidateImage(ctx context.Context, formats strfmt.Registry) error {

	if m.Image != nil {
		if err := m.Image.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *DomainNewsDocument) contextValidateMotivations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Motivations); i++ {

		if m.Motivations[i] != nil {
			if err := m.Motivations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("motivations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("motivations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainNewsDocument) contextValidateSubType(ctx context.Context, formats strfmt.Registry) error {

	if m.SubType != nil {
		if err := m.SubType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sub_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sub_type")
			}
			return err
		}
	}

	return nil
}

func (m *DomainNewsDocument) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainNewsDocument) contextValidateTargetCountries(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TargetCountries); i++ {

		if m.TargetCountries[i] != nil {
			if err := m.TargetCountries[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("target_countries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("target_countries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainNewsDocument) contextValidateTargetIndustries(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TargetIndustries); i++ {

		if m.TargetIndustries[i] != nil {
			if err := m.TargetIndustries[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("target_industries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("target_industries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainNewsDocument) contextValidateThumbnail(ctx context.Context, formats strfmt.Registry) error {

	if m.Thumbnail != nil {
		if err := m.Thumbnail.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("thumbnail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("thumbnail")
			}
			return err
		}
	}

	return nil
}

func (m *DomainNewsDocument) contextValidateTopic(ctx context.Context, formats strfmt.Registry) error {

	if m.Topic != nil {
		if err := m.Topic.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("topic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("topic")
			}
			return err
		}
	}

	return nil
}

func (m *DomainNewsDocument) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainNewsDocument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainNewsDocument) UnmarshalBinary(b []byte) error {
	var res DomainNewsDocument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
