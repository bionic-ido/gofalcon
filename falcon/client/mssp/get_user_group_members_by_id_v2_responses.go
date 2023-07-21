// Code generated by go-swagger; DO NOT EDIT.

package mssp

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

// GetUserGroupMembersByIDV2Reader is a Reader for the GetUserGroupMembersByIDV2 structure.
type GetUserGroupMembersByIDV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserGroupMembersByIDV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserGroupMembersByIDV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewGetUserGroupMembersByIDV2MultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserGroupMembersByIDV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserGroupMembersByIDV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUserGroupMembersByIDV2TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /mssp/entities/user-group-members/v2] getUserGroupMembersByIDV2", response, response.Code())
	}
}

// NewGetUserGroupMembersByIDV2OK creates a GetUserGroupMembersByIDV2OK with default headers values
func NewGetUserGroupMembersByIDV2OK() *GetUserGroupMembersByIDV2OK {
	return &GetUserGroupMembersByIDV2OK{}
}

/*
GetUserGroupMembersByIDV2OK describes a response with status code 200, with default header values.

OK
*/
type GetUserGroupMembersByIDV2OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainUserGroupMembersResponseV1
}

// IsSuccess returns true when this get user group members by Id v2 o k response has a 2xx status code
func (o *GetUserGroupMembersByIDV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user group members by Id v2 o k response has a 3xx status code
func (o *GetUserGroupMembersByIDV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user group members by Id v2 o k response has a 4xx status code
func (o *GetUserGroupMembersByIDV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user group members by Id v2 o k response has a 5xx status code
func (o *GetUserGroupMembersByIDV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user group members by Id v2 o k response a status code equal to that given
func (o *GetUserGroupMembersByIDV2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get user group members by Id v2 o k response
func (o *GetUserGroupMembersByIDV2OK) Code() int {
	return 200
}

func (o *GetUserGroupMembersByIDV2OK) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/user-group-members/v2][%d] getUserGroupMembersByIdV2OK  %+v", 200, o.Payload)
}

func (o *GetUserGroupMembersByIDV2OK) String() string {
	return fmt.Sprintf("[GET /mssp/entities/user-group-members/v2][%d] getUserGroupMembersByIdV2OK  %+v", 200, o.Payload)
}

func (o *GetUserGroupMembersByIDV2OK) GetPayload() *models.DomainUserGroupMembersResponseV1 {
	return o.Payload
}

func (o *GetUserGroupMembersByIDV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainUserGroupMembersResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGroupMembersByIDV2MultiStatus creates a GetUserGroupMembersByIDV2MultiStatus with default headers values
func NewGetUserGroupMembersByIDV2MultiStatus() *GetUserGroupMembersByIDV2MultiStatus {
	return &GetUserGroupMembersByIDV2MultiStatus{}
}

/*
GetUserGroupMembersByIDV2MultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type GetUserGroupMembersByIDV2MultiStatus struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainUserGroupMembersResponseV1
}

// IsSuccess returns true when this get user group members by Id v2 multi status response has a 2xx status code
func (o *GetUserGroupMembersByIDV2MultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user group members by Id v2 multi status response has a 3xx status code
func (o *GetUserGroupMembersByIDV2MultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user group members by Id v2 multi status response has a 4xx status code
func (o *GetUserGroupMembersByIDV2MultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user group members by Id v2 multi status response has a 5xx status code
func (o *GetUserGroupMembersByIDV2MultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this get user group members by Id v2 multi status response a status code equal to that given
func (o *GetUserGroupMembersByIDV2MultiStatus) IsCode(code int) bool {
	return code == 207
}

// Code gets the status code for the get user group members by Id v2 multi status response
func (o *GetUserGroupMembersByIDV2MultiStatus) Code() int {
	return 207
}

func (o *GetUserGroupMembersByIDV2MultiStatus) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/user-group-members/v2][%d] getUserGroupMembersByIdV2MultiStatus  %+v", 207, o.Payload)
}

func (o *GetUserGroupMembersByIDV2MultiStatus) String() string {
	return fmt.Sprintf("[GET /mssp/entities/user-group-members/v2][%d] getUserGroupMembersByIdV2MultiStatus  %+v", 207, o.Payload)
}

func (o *GetUserGroupMembersByIDV2MultiStatus) GetPayload() *models.DomainUserGroupMembersResponseV1 {
	return o.Payload
}

func (o *GetUserGroupMembersByIDV2MultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainUserGroupMembersResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGroupMembersByIDV2BadRequest creates a GetUserGroupMembersByIDV2BadRequest with default headers values
func NewGetUserGroupMembersByIDV2BadRequest() *GetUserGroupMembersByIDV2BadRequest {
	return &GetUserGroupMembersByIDV2BadRequest{}
}

/*
GetUserGroupMembersByIDV2BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetUserGroupMembersByIDV2BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this get user group members by Id v2 bad request response has a 2xx status code
func (o *GetUserGroupMembersByIDV2BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user group members by Id v2 bad request response has a 3xx status code
func (o *GetUserGroupMembersByIDV2BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user group members by Id v2 bad request response has a 4xx status code
func (o *GetUserGroupMembersByIDV2BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user group members by Id v2 bad request response has a 5xx status code
func (o *GetUserGroupMembersByIDV2BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get user group members by Id v2 bad request response a status code equal to that given
func (o *GetUserGroupMembersByIDV2BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get user group members by Id v2 bad request response
func (o *GetUserGroupMembersByIDV2BadRequest) Code() int {
	return 400
}

func (o *GetUserGroupMembersByIDV2BadRequest) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/user-group-members/v2][%d] getUserGroupMembersByIdV2BadRequest  %+v", 400, o.Payload)
}

func (o *GetUserGroupMembersByIDV2BadRequest) String() string {
	return fmt.Sprintf("[GET /mssp/entities/user-group-members/v2][%d] getUserGroupMembersByIdV2BadRequest  %+v", 400, o.Payload)
}

func (o *GetUserGroupMembersByIDV2BadRequest) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetUserGroupMembersByIDV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGroupMembersByIDV2Forbidden creates a GetUserGroupMembersByIDV2Forbidden with default headers values
func NewGetUserGroupMembersByIDV2Forbidden() *GetUserGroupMembersByIDV2Forbidden {
	return &GetUserGroupMembersByIDV2Forbidden{}
}

/*
GetUserGroupMembersByIDV2Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetUserGroupMembersByIDV2Forbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this get user group members by Id v2 forbidden response has a 2xx status code
func (o *GetUserGroupMembersByIDV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user group members by Id v2 forbidden response has a 3xx status code
func (o *GetUserGroupMembersByIDV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user group members by Id v2 forbidden response has a 4xx status code
func (o *GetUserGroupMembersByIDV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user group members by Id v2 forbidden response has a 5xx status code
func (o *GetUserGroupMembersByIDV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get user group members by Id v2 forbidden response a status code equal to that given
func (o *GetUserGroupMembersByIDV2Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get user group members by Id v2 forbidden response
func (o *GetUserGroupMembersByIDV2Forbidden) Code() int {
	return 403
}

func (o *GetUserGroupMembersByIDV2Forbidden) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/user-group-members/v2][%d] getUserGroupMembersByIdV2Forbidden  %+v", 403, o.Payload)
}

func (o *GetUserGroupMembersByIDV2Forbidden) String() string {
	return fmt.Sprintf("[GET /mssp/entities/user-group-members/v2][%d] getUserGroupMembersByIdV2Forbidden  %+v", 403, o.Payload)
}

func (o *GetUserGroupMembersByIDV2Forbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetUserGroupMembersByIDV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGroupMembersByIDV2TooManyRequests creates a GetUserGroupMembersByIDV2TooManyRequests with default headers values
func NewGetUserGroupMembersByIDV2TooManyRequests() *GetUserGroupMembersByIDV2TooManyRequests {
	return &GetUserGroupMembersByIDV2TooManyRequests{}
}

/*
GetUserGroupMembersByIDV2TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetUserGroupMembersByIDV2TooManyRequests struct {

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

// IsSuccess returns true when this get user group members by Id v2 too many requests response has a 2xx status code
func (o *GetUserGroupMembersByIDV2TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user group members by Id v2 too many requests response has a 3xx status code
func (o *GetUserGroupMembersByIDV2TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user group members by Id v2 too many requests response has a 4xx status code
func (o *GetUserGroupMembersByIDV2TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user group members by Id v2 too many requests response has a 5xx status code
func (o *GetUserGroupMembersByIDV2TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get user group members by Id v2 too many requests response a status code equal to that given
func (o *GetUserGroupMembersByIDV2TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get user group members by Id v2 too many requests response
func (o *GetUserGroupMembersByIDV2TooManyRequests) Code() int {
	return 429
}

func (o *GetUserGroupMembersByIDV2TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/user-group-members/v2][%d] getUserGroupMembersByIdV2TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserGroupMembersByIDV2TooManyRequests) String() string {
	return fmt.Sprintf("[GET /mssp/entities/user-group-members/v2][%d] getUserGroupMembersByIdV2TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserGroupMembersByIDV2TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetUserGroupMembersByIDV2TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
