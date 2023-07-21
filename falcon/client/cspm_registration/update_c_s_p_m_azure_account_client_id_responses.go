// Code generated by go-swagger; DO NOT EDIT.

package cspm_registration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// UpdateCSPMAzureAccountClientIDReader is a Reader for the UpdateCSPMAzureAccountClientID structure.
type UpdateCSPMAzureAccountClientIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCSPMAzureAccountClientIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateCSPMAzureAccountClientIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateCSPMAzureAccountClientIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateCSPMAzureAccountClientIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateCSPMAzureAccountClientIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateCSPMAzureAccountClientIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /cloud-connect-cspm-azure/entities/client-id/v1] UpdateCSPMAzureAccountClientID", response, response.Code())
	}
}

// NewUpdateCSPMAzureAccountClientIDCreated creates a UpdateCSPMAzureAccountClientIDCreated with default headers values
func NewUpdateCSPMAzureAccountClientIDCreated() *UpdateCSPMAzureAccountClientIDCreated {
	return &UpdateCSPMAzureAccountClientIDCreated{}
}

/*
UpdateCSPMAzureAccountClientIDCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateCSPMAzureAccountClientIDCreated struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAzureServicePrincipalResponseV1
}

// IsSuccess returns true when this update c s p m azure account client Id created response has a 2xx status code
func (o *UpdateCSPMAzureAccountClientIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update c s p m azure account client Id created response has a 3xx status code
func (o *UpdateCSPMAzureAccountClientIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update c s p m azure account client Id created response has a 4xx status code
func (o *UpdateCSPMAzureAccountClientIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update c s p m azure account client Id created response has a 5xx status code
func (o *UpdateCSPMAzureAccountClientIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update c s p m azure account client Id created response a status code equal to that given
func (o *UpdateCSPMAzureAccountClientIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update c s p m azure account client Id created response
func (o *UpdateCSPMAzureAccountClientIDCreated) Code() int {
	return 201
}

func (o *UpdateCSPMAzureAccountClientIDCreated) Error() string {
	return fmt.Sprintf("[PATCH /cloud-connect-cspm-azure/entities/client-id/v1][%d] updateCSPMAzureAccountClientIdCreated  %+v", 201, o.Payload)
}

func (o *UpdateCSPMAzureAccountClientIDCreated) String() string {
	return fmt.Sprintf("[PATCH /cloud-connect-cspm-azure/entities/client-id/v1][%d] updateCSPMAzureAccountClientIdCreated  %+v", 201, o.Payload)
}

func (o *UpdateCSPMAzureAccountClientIDCreated) GetPayload() *models.RegistrationAzureServicePrincipalResponseV1 {
	return o.Payload
}

func (o *UpdateCSPMAzureAccountClientIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.RegistrationAzureServicePrincipalResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCSPMAzureAccountClientIDBadRequest creates a UpdateCSPMAzureAccountClientIDBadRequest with default headers values
func NewUpdateCSPMAzureAccountClientIDBadRequest() *UpdateCSPMAzureAccountClientIDBadRequest {
	return &UpdateCSPMAzureAccountClientIDBadRequest{}
}

/*
UpdateCSPMAzureAccountClientIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateCSPMAzureAccountClientIDBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAzureServicePrincipalResponseV1
}

// IsSuccess returns true when this update c s p m azure account client Id bad request response has a 2xx status code
func (o *UpdateCSPMAzureAccountClientIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update c s p m azure account client Id bad request response has a 3xx status code
func (o *UpdateCSPMAzureAccountClientIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update c s p m azure account client Id bad request response has a 4xx status code
func (o *UpdateCSPMAzureAccountClientIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update c s p m azure account client Id bad request response has a 5xx status code
func (o *UpdateCSPMAzureAccountClientIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update c s p m azure account client Id bad request response a status code equal to that given
func (o *UpdateCSPMAzureAccountClientIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update c s p m azure account client Id bad request response
func (o *UpdateCSPMAzureAccountClientIDBadRequest) Code() int {
	return 400
}

func (o *UpdateCSPMAzureAccountClientIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /cloud-connect-cspm-azure/entities/client-id/v1][%d] updateCSPMAzureAccountClientIdBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateCSPMAzureAccountClientIDBadRequest) String() string {
	return fmt.Sprintf("[PATCH /cloud-connect-cspm-azure/entities/client-id/v1][%d] updateCSPMAzureAccountClientIdBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateCSPMAzureAccountClientIDBadRequest) GetPayload() *models.RegistrationAzureServicePrincipalResponseV1 {
	return o.Payload
}

func (o *UpdateCSPMAzureAccountClientIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.RegistrationAzureServicePrincipalResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCSPMAzureAccountClientIDForbidden creates a UpdateCSPMAzureAccountClientIDForbidden with default headers values
func NewUpdateCSPMAzureAccountClientIDForbidden() *UpdateCSPMAzureAccountClientIDForbidden {
	return &UpdateCSPMAzureAccountClientIDForbidden{}
}

/*
UpdateCSPMAzureAccountClientIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateCSPMAzureAccountClientIDForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this update c s p m azure account client Id forbidden response has a 2xx status code
func (o *UpdateCSPMAzureAccountClientIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update c s p m azure account client Id forbidden response has a 3xx status code
func (o *UpdateCSPMAzureAccountClientIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update c s p m azure account client Id forbidden response has a 4xx status code
func (o *UpdateCSPMAzureAccountClientIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update c s p m azure account client Id forbidden response has a 5xx status code
func (o *UpdateCSPMAzureAccountClientIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update c s p m azure account client Id forbidden response a status code equal to that given
func (o *UpdateCSPMAzureAccountClientIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update c s p m azure account client Id forbidden response
func (o *UpdateCSPMAzureAccountClientIDForbidden) Code() int {
	return 403
}

func (o *UpdateCSPMAzureAccountClientIDForbidden) Error() string {
	return fmt.Sprintf("[PATCH /cloud-connect-cspm-azure/entities/client-id/v1][%d] updateCSPMAzureAccountClientIdForbidden  %+v", 403, o.Payload)
}

func (o *UpdateCSPMAzureAccountClientIDForbidden) String() string {
	return fmt.Sprintf("[PATCH /cloud-connect-cspm-azure/entities/client-id/v1][%d] updateCSPMAzureAccountClientIdForbidden  %+v", 403, o.Payload)
}

func (o *UpdateCSPMAzureAccountClientIDForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateCSPMAzureAccountClientIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCSPMAzureAccountClientIDTooManyRequests creates a UpdateCSPMAzureAccountClientIDTooManyRequests with default headers values
func NewUpdateCSPMAzureAccountClientIDTooManyRequests() *UpdateCSPMAzureAccountClientIDTooManyRequests {
	return &UpdateCSPMAzureAccountClientIDTooManyRequests{}
}

/*
UpdateCSPMAzureAccountClientIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateCSPMAzureAccountClientIDTooManyRequests struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	/* Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this update c s p m azure account client Id too many requests response has a 2xx status code
func (o *UpdateCSPMAzureAccountClientIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update c s p m azure account client Id too many requests response has a 3xx status code
func (o *UpdateCSPMAzureAccountClientIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update c s p m azure account client Id too many requests response has a 4xx status code
func (o *UpdateCSPMAzureAccountClientIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update c s p m azure account client Id too many requests response has a 5xx status code
func (o *UpdateCSPMAzureAccountClientIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update c s p m azure account client Id too many requests response a status code equal to that given
func (o *UpdateCSPMAzureAccountClientIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the update c s p m azure account client Id too many requests response
func (o *UpdateCSPMAzureAccountClientIDTooManyRequests) Code() int {
	return 429
}

func (o *UpdateCSPMAzureAccountClientIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /cloud-connect-cspm-azure/entities/client-id/v1][%d] updateCSPMAzureAccountClientIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateCSPMAzureAccountClientIDTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /cloud-connect-cspm-azure/entities/client-id/v1][%d] updateCSPMAzureAccountClientIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateCSPMAzureAccountClientIDTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateCSPMAzureAccountClientIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header X-RateLimit-RetryAfter
	hdrXRateLimitRetryAfter := response.GetHeader("X-RateLimit-RetryAfter")

	if hdrXRateLimitRetryAfter != "" {
		valxRateLimitRetryAfter, err := swag.ConvertInt64(hdrXRateLimitRetryAfter)
		if err != nil {
			return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", hdrXRateLimitRetryAfter)
		}
		o.XRateLimitRetryAfter = valxRateLimitRetryAfter
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCSPMAzureAccountClientIDInternalServerError creates a UpdateCSPMAzureAccountClientIDInternalServerError with default headers values
func NewUpdateCSPMAzureAccountClientIDInternalServerError() *UpdateCSPMAzureAccountClientIDInternalServerError {
	return &UpdateCSPMAzureAccountClientIDInternalServerError{}
}

/*
UpdateCSPMAzureAccountClientIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateCSPMAzureAccountClientIDInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAzureServicePrincipalResponseV1
}

// IsSuccess returns true when this update c s p m azure account client Id internal server error response has a 2xx status code
func (o *UpdateCSPMAzureAccountClientIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update c s p m azure account client Id internal server error response has a 3xx status code
func (o *UpdateCSPMAzureAccountClientIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update c s p m azure account client Id internal server error response has a 4xx status code
func (o *UpdateCSPMAzureAccountClientIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update c s p m azure account client Id internal server error response has a 5xx status code
func (o *UpdateCSPMAzureAccountClientIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update c s p m azure account client Id internal server error response a status code equal to that given
func (o *UpdateCSPMAzureAccountClientIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update c s p m azure account client Id internal server error response
func (o *UpdateCSPMAzureAccountClientIDInternalServerError) Code() int {
	return 500
}

func (o *UpdateCSPMAzureAccountClientIDInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /cloud-connect-cspm-azure/entities/client-id/v1][%d] updateCSPMAzureAccountClientIdInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateCSPMAzureAccountClientIDInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /cloud-connect-cspm-azure/entities/client-id/v1][%d] updateCSPMAzureAccountClientIdInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateCSPMAzureAccountClientIDInternalServerError) GetPayload() *models.RegistrationAzureServicePrincipalResponseV1 {
	return o.Payload
}

func (o *UpdateCSPMAzureAccountClientIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.RegistrationAzureServicePrincipalResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
