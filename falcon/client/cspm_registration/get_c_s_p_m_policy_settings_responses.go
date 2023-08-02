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

// GetCSPMPolicySettingsReader is a Reader for the GetCSPMPolicySettings structure.
type GetCSPMPolicySettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCSPMPolicySettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCSPMPolicySettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewGetCSPMPolicySettingsMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCSPMPolicySettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCSPMPolicySettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCSPMPolicySettingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCSPMPolicySettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /settings/entities/policy/v1] GetCSPMPolicySettings", response, response.Code())
	}
}

// NewGetCSPMPolicySettingsOK creates a GetCSPMPolicySettingsOK with default headers values
func NewGetCSPMPolicySettingsOK() *GetCSPMPolicySettingsOK {
	return &GetCSPMPolicySettingsOK{}
}

/*
GetCSPMPolicySettingsOK describes a response with status code 200, with default header values.

OK
*/
type GetCSPMPolicySettingsOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationPolicySettingsResponseV1
}

// IsSuccess returns true when this get c s p m policy settings o k response has a 2xx status code
func (o *GetCSPMPolicySettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get c s p m policy settings o k response has a 3xx status code
func (o *GetCSPMPolicySettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s p m policy settings o k response has a 4xx status code
func (o *GetCSPMPolicySettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get c s p m policy settings o k response has a 5xx status code
func (o *GetCSPMPolicySettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get c s p m policy settings o k response a status code equal to that given
func (o *GetCSPMPolicySettingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get c s p m policy settings o k response
func (o *GetCSPMPolicySettingsOK) Code() int {
	return 200
}

func (o *GetCSPMPolicySettingsOK) Error() string {
	return fmt.Sprintf("[GET /settings/entities/policy/v1][%d] getCSPMPolicySettingsOK  %+v", 200, o.Payload)
}

func (o *GetCSPMPolicySettingsOK) String() string {
	return fmt.Sprintf("[GET /settings/entities/policy/v1][%d] getCSPMPolicySettingsOK  %+v", 200, o.Payload)
}

func (o *GetCSPMPolicySettingsOK) GetPayload() *models.RegistrationPolicySettingsResponseV1 {
	return o.Payload
}

func (o *GetCSPMPolicySettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

	o.Payload = new(models.RegistrationPolicySettingsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMPolicySettingsMultiStatus creates a GetCSPMPolicySettingsMultiStatus with default headers values
func NewGetCSPMPolicySettingsMultiStatus() *GetCSPMPolicySettingsMultiStatus {
	return &GetCSPMPolicySettingsMultiStatus{}
}

/*
GetCSPMPolicySettingsMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type GetCSPMPolicySettingsMultiStatus struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationPolicySettingsResponseV1
}

// IsSuccess returns true when this get c s p m policy settings multi status response has a 2xx status code
func (o *GetCSPMPolicySettingsMultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get c s p m policy settings multi status response has a 3xx status code
func (o *GetCSPMPolicySettingsMultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s p m policy settings multi status response has a 4xx status code
func (o *GetCSPMPolicySettingsMultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this get c s p m policy settings multi status response has a 5xx status code
func (o *GetCSPMPolicySettingsMultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this get c s p m policy settings multi status response a status code equal to that given
func (o *GetCSPMPolicySettingsMultiStatus) IsCode(code int) bool {
	return code == 207
}

// Code gets the status code for the get c s p m policy settings multi status response
func (o *GetCSPMPolicySettingsMultiStatus) Code() int {
	return 207
}

func (o *GetCSPMPolicySettingsMultiStatus) Error() string {
	return fmt.Sprintf("[GET /settings/entities/policy/v1][%d] getCSPMPolicySettingsMultiStatus  %+v", 207, o.Payload)
}

func (o *GetCSPMPolicySettingsMultiStatus) String() string {
	return fmt.Sprintf("[GET /settings/entities/policy/v1][%d] getCSPMPolicySettingsMultiStatus  %+v", 207, o.Payload)
}

func (o *GetCSPMPolicySettingsMultiStatus) GetPayload() *models.RegistrationPolicySettingsResponseV1 {
	return o.Payload
}

func (o *GetCSPMPolicySettingsMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

	o.Payload = new(models.RegistrationPolicySettingsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMPolicySettingsBadRequest creates a GetCSPMPolicySettingsBadRequest with default headers values
func NewGetCSPMPolicySettingsBadRequest() *GetCSPMPolicySettingsBadRequest {
	return &GetCSPMPolicySettingsBadRequest{}
}

/*
GetCSPMPolicySettingsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetCSPMPolicySettingsBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationPolicySettingsResponseV1
}

// IsSuccess returns true when this get c s p m policy settings bad request response has a 2xx status code
func (o *GetCSPMPolicySettingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c s p m policy settings bad request response has a 3xx status code
func (o *GetCSPMPolicySettingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s p m policy settings bad request response has a 4xx status code
func (o *GetCSPMPolicySettingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get c s p m policy settings bad request response has a 5xx status code
func (o *GetCSPMPolicySettingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get c s p m policy settings bad request response a status code equal to that given
func (o *GetCSPMPolicySettingsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get c s p m policy settings bad request response
func (o *GetCSPMPolicySettingsBadRequest) Code() int {
	return 400
}

func (o *GetCSPMPolicySettingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /settings/entities/policy/v1][%d] getCSPMPolicySettingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetCSPMPolicySettingsBadRequest) String() string {
	return fmt.Sprintf("[GET /settings/entities/policy/v1][%d] getCSPMPolicySettingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetCSPMPolicySettingsBadRequest) GetPayload() *models.RegistrationPolicySettingsResponseV1 {
	return o.Payload
}

func (o *GetCSPMPolicySettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

	o.Payload = new(models.RegistrationPolicySettingsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMPolicySettingsForbidden creates a GetCSPMPolicySettingsForbidden with default headers values
func NewGetCSPMPolicySettingsForbidden() *GetCSPMPolicySettingsForbidden {
	return &GetCSPMPolicySettingsForbidden{}
}

/*
GetCSPMPolicySettingsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCSPMPolicySettingsForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this get c s p m policy settings forbidden response has a 2xx status code
func (o *GetCSPMPolicySettingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c s p m policy settings forbidden response has a 3xx status code
func (o *GetCSPMPolicySettingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s p m policy settings forbidden response has a 4xx status code
func (o *GetCSPMPolicySettingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get c s p m policy settings forbidden response has a 5xx status code
func (o *GetCSPMPolicySettingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get c s p m policy settings forbidden response a status code equal to that given
func (o *GetCSPMPolicySettingsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get c s p m policy settings forbidden response
func (o *GetCSPMPolicySettingsForbidden) Code() int {
	return 403
}

func (o *GetCSPMPolicySettingsForbidden) Error() string {
	return fmt.Sprintf("[GET /settings/entities/policy/v1][%d] getCSPMPolicySettingsForbidden  %+v", 403, o.Payload)
}

func (o *GetCSPMPolicySettingsForbidden) String() string {
	return fmt.Sprintf("[GET /settings/entities/policy/v1][%d] getCSPMPolicySettingsForbidden  %+v", 403, o.Payload)
}

func (o *GetCSPMPolicySettingsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCSPMPolicySettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

// NewGetCSPMPolicySettingsTooManyRequests creates a GetCSPMPolicySettingsTooManyRequests with default headers values
func NewGetCSPMPolicySettingsTooManyRequests() *GetCSPMPolicySettingsTooManyRequests {
	return &GetCSPMPolicySettingsTooManyRequests{}
}

/*
GetCSPMPolicySettingsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetCSPMPolicySettingsTooManyRequests struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

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

// IsSuccess returns true when this get c s p m policy settings too many requests response has a 2xx status code
func (o *GetCSPMPolicySettingsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c s p m policy settings too many requests response has a 3xx status code
func (o *GetCSPMPolicySettingsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s p m policy settings too many requests response has a 4xx status code
func (o *GetCSPMPolicySettingsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get c s p m policy settings too many requests response has a 5xx status code
func (o *GetCSPMPolicySettingsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get c s p m policy settings too many requests response a status code equal to that given
func (o *GetCSPMPolicySettingsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get c s p m policy settings too many requests response
func (o *GetCSPMPolicySettingsTooManyRequests) Code() int {
	return 429
}

func (o *GetCSPMPolicySettingsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /settings/entities/policy/v1][%d] getCSPMPolicySettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCSPMPolicySettingsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /settings/entities/policy/v1][%d] getCSPMPolicySettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCSPMPolicySettingsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCSPMPolicySettingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

// NewGetCSPMPolicySettingsInternalServerError creates a GetCSPMPolicySettingsInternalServerError with default headers values
func NewGetCSPMPolicySettingsInternalServerError() *GetCSPMPolicySettingsInternalServerError {
	return &GetCSPMPolicySettingsInternalServerError{}
}

/*
GetCSPMPolicySettingsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetCSPMPolicySettingsInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationPolicySettingsResponseV1
}

// IsSuccess returns true when this get c s p m policy settings internal server error response has a 2xx status code
func (o *GetCSPMPolicySettingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c s p m policy settings internal server error response has a 3xx status code
func (o *GetCSPMPolicySettingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s p m policy settings internal server error response has a 4xx status code
func (o *GetCSPMPolicySettingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get c s p m policy settings internal server error response has a 5xx status code
func (o *GetCSPMPolicySettingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get c s p m policy settings internal server error response a status code equal to that given
func (o *GetCSPMPolicySettingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get c s p m policy settings internal server error response
func (o *GetCSPMPolicySettingsInternalServerError) Code() int {
	return 500
}

func (o *GetCSPMPolicySettingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /settings/entities/policy/v1][%d] getCSPMPolicySettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCSPMPolicySettingsInternalServerError) String() string {
	return fmt.Sprintf("[GET /settings/entities/policy/v1][%d] getCSPMPolicySettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCSPMPolicySettingsInternalServerError) GetPayload() *models.RegistrationPolicySettingsResponseV1 {
	return o.Payload
}

func (o *GetCSPMPolicySettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

	o.Payload = new(models.RegistrationPolicySettingsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
