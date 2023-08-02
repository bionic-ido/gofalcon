// Code generated by go-swagger; DO NOT EDIT.

package cspm_registration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetCSPMAzureUserScriptsAttachmentParams creates a new GetCSPMAzureUserScriptsAttachmentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCSPMAzureUserScriptsAttachmentParams() *GetCSPMAzureUserScriptsAttachmentParams {
	return &GetCSPMAzureUserScriptsAttachmentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCSPMAzureUserScriptsAttachmentParamsWithTimeout creates a new GetCSPMAzureUserScriptsAttachmentParams object
// with the ability to set a timeout on a request.
func NewGetCSPMAzureUserScriptsAttachmentParamsWithTimeout(timeout time.Duration) *GetCSPMAzureUserScriptsAttachmentParams {
	return &GetCSPMAzureUserScriptsAttachmentParams{
		timeout: timeout,
	}
}

// NewGetCSPMAzureUserScriptsAttachmentParamsWithContext creates a new GetCSPMAzureUserScriptsAttachmentParams object
// with the ability to set a context for a request.
func NewGetCSPMAzureUserScriptsAttachmentParamsWithContext(ctx context.Context) *GetCSPMAzureUserScriptsAttachmentParams {
	return &GetCSPMAzureUserScriptsAttachmentParams{
		Context: ctx,
	}
}

// NewGetCSPMAzureUserScriptsAttachmentParamsWithHTTPClient creates a new GetCSPMAzureUserScriptsAttachmentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCSPMAzureUserScriptsAttachmentParamsWithHTTPClient(client *http.Client) *GetCSPMAzureUserScriptsAttachmentParams {
	return &GetCSPMAzureUserScriptsAttachmentParams{
		HTTPClient: client,
	}
}

/*
GetCSPMAzureUserScriptsAttachmentParams contains all the parameters to send to the API endpoint

	for the get c s p m azure user scripts attachment operation.

	Typically these are written to a http.Request.
*/
type GetCSPMAzureUserScriptsAttachmentParams struct {

	// AccountType.
	AccountType *string

	/* SubscriptionIds.

	   Subscription IDs to generate script for. Defaults to all.
	*/
	SubscriptionIds []string

	/* Template.

	   Template to be rendered
	*/
	Template *string

	/* TenantID.

	   Tenant ID to generate script for. Defaults to most recently registered tenant.
	*/
	TenantID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get c s p m azure user scripts attachment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCSPMAzureUserScriptsAttachmentParams) WithDefaults() *GetCSPMAzureUserScriptsAttachmentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get c s p m azure user scripts attachment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCSPMAzureUserScriptsAttachmentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get c s p m azure user scripts attachment params
func (o *GetCSPMAzureUserScriptsAttachmentParams) WithTimeout(timeout time.Duration) *GetCSPMAzureUserScriptsAttachmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get c s p m azure user scripts attachment params
func (o *GetCSPMAzureUserScriptsAttachmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get c s p m azure user scripts attachment params
func (o *GetCSPMAzureUserScriptsAttachmentParams) WithContext(ctx context.Context) *GetCSPMAzureUserScriptsAttachmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get c s p m azure user scripts attachment params
func (o *GetCSPMAzureUserScriptsAttachmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get c s p m azure user scripts attachment params
func (o *GetCSPMAzureUserScriptsAttachmentParams) WithHTTPClient(client *http.Client) *GetCSPMAzureUserScriptsAttachmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get c s p m azure user scripts attachment params
func (o *GetCSPMAzureUserScriptsAttachmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountType adds the accountType to the get c s p m azure user scripts attachment params
func (o *GetCSPMAzureUserScriptsAttachmentParams) WithAccountType(accountType *string) *GetCSPMAzureUserScriptsAttachmentParams {
	o.SetAccountType(accountType)
	return o
}

// SetAccountType adds the accountType to the get c s p m azure user scripts attachment params
func (o *GetCSPMAzureUserScriptsAttachmentParams) SetAccountType(accountType *string) {
	o.AccountType = accountType
}

// WithSubscriptionIds adds the subscriptionIds to the get c s p m azure user scripts attachment params
func (o *GetCSPMAzureUserScriptsAttachmentParams) WithSubscriptionIds(subscriptionIds []string) *GetCSPMAzureUserScriptsAttachmentParams {
	o.SetSubscriptionIds(subscriptionIds)
	return o
}

// SetSubscriptionIds adds the subscriptionIds to the get c s p m azure user scripts attachment params
func (o *GetCSPMAzureUserScriptsAttachmentParams) SetSubscriptionIds(subscriptionIds []string) {
	o.SubscriptionIds = subscriptionIds
}

// WithTemplate adds the template to the get c s p m azure user scripts attachment params
func (o *GetCSPMAzureUserScriptsAttachmentParams) WithTemplate(template *string) *GetCSPMAzureUserScriptsAttachmentParams {
	o.SetTemplate(template)
	return o
}

// SetTemplate adds the template to the get c s p m azure user scripts attachment params
func (o *GetCSPMAzureUserScriptsAttachmentParams) SetTemplate(template *string) {
	o.Template = template
}

// WithTenantID adds the tenantID to the get c s p m azure user scripts attachment params
func (o *GetCSPMAzureUserScriptsAttachmentParams) WithTenantID(tenantID *string) *GetCSPMAzureUserScriptsAttachmentParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the get c s p m azure user scripts attachment params
func (o *GetCSPMAzureUserScriptsAttachmentParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCSPMAzureUserScriptsAttachmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountType != nil {

		// query param account_type
		var qrAccountType string

		if o.AccountType != nil {
			qrAccountType = *o.AccountType
		}
		qAccountType := qrAccountType
		if qAccountType != "" {

			if err := r.SetQueryParam("account_type", qAccountType); err != nil {
				return err
			}
		}
	}

	if o.SubscriptionIds != nil {

		// binding items for subscription_ids
		joinedSubscriptionIds := o.bindParamSubscriptionIds(reg)

		// query array param subscription_ids
		if err := r.SetQueryParam("subscription_ids", joinedSubscriptionIds...); err != nil {
			return err
		}
	}

	if o.Template != nil {

		// query param template
		var qrTemplate string

		if o.Template != nil {
			qrTemplate = *o.Template
		}
		qTemplate := qrTemplate
		if qTemplate != "" {

			if err := r.SetQueryParam("template", qTemplate); err != nil {
				return err
			}
		}
	}

	if o.TenantID != nil {

		// query param tenant-id
		var qrTenantID string

		if o.TenantID != nil {
			qrTenantID = *o.TenantID
		}
		qTenantID := qrTenantID
		if qTenantID != "" {

			if err := r.SetQueryParam("tenant-id", qTenantID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetCSPMAzureUserScriptsAttachment binds the parameter subscription_ids
func (o *GetCSPMAzureUserScriptsAttachmentParams) bindParamSubscriptionIds(formats strfmt.Registry) []string {
	subscriptionIdsIR := o.SubscriptionIds

	var subscriptionIdsIC []string
	for _, subscriptionIdsIIR := range subscriptionIdsIR { // explode []string

		subscriptionIdsIIV := subscriptionIdsIIR // string as string
		subscriptionIdsIC = append(subscriptionIdsIC, subscriptionIdsIIV)
	}

	// items.CollectionFormat: "multi"
	subscriptionIdsIS := swag.JoinByFormat(subscriptionIdsIC, "multi")

	return subscriptionIdsIS
}
