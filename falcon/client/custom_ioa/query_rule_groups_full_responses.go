// Code generated by go-swagger; DO NOT EDIT.

package custom_ioa

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

// QueryRuleGroupsFullReader is a Reader for the QueryRuleGroupsFull structure.
type QueryRuleGroupsFullReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryRuleGroupsFullReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryRuleGroupsFullOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewQueryRuleGroupsFullForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewQueryRuleGroupsFullNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryRuleGroupsFullTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /ioarules/queries/rule-groups-full/v1] query-rule-groups-full", response, response.Code())
	}
}

// NewQueryRuleGroupsFullOK creates a QueryRuleGroupsFullOK with default headers values
func NewQueryRuleGroupsFullOK() *QueryRuleGroupsFullOK {
	return &QueryRuleGroupsFullOK{}
}

/*
QueryRuleGroupsFullOK describes a response with status code 200, with default header values.

OK
*/
type QueryRuleGroupsFullOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIRuleGroupsResponse
}

// IsSuccess returns true when this query rule groups full o k response has a 2xx status code
func (o *QueryRuleGroupsFullOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query rule groups full o k response has a 3xx status code
func (o *QueryRuleGroupsFullOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query rule groups full o k response has a 4xx status code
func (o *QueryRuleGroupsFullOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query rule groups full o k response has a 5xx status code
func (o *QueryRuleGroupsFullOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query rule groups full o k response a status code equal to that given
func (o *QueryRuleGroupsFullOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query rule groups full o k response
func (o *QueryRuleGroupsFullOK) Code() int {
	return 200
}

func (o *QueryRuleGroupsFullOK) Error() string {
	return fmt.Sprintf("[GET /ioarules/queries/rule-groups-full/v1][%d] queryRuleGroupsFullOK  %+v", 200, o.Payload)
}

func (o *QueryRuleGroupsFullOK) String() string {
	return fmt.Sprintf("[GET /ioarules/queries/rule-groups-full/v1][%d] queryRuleGroupsFullOK  %+v", 200, o.Payload)
}

func (o *QueryRuleGroupsFullOK) GetPayload() *models.APIRuleGroupsResponse {
	return o.Payload
}

func (o *QueryRuleGroupsFullOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIRuleGroupsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryRuleGroupsFullForbidden creates a QueryRuleGroupsFullForbidden with default headers values
func NewQueryRuleGroupsFullForbidden() *QueryRuleGroupsFullForbidden {
	return &QueryRuleGroupsFullForbidden{}
}

/*
QueryRuleGroupsFullForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryRuleGroupsFullForbidden struct {

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

// IsSuccess returns true when this query rule groups full forbidden response has a 2xx status code
func (o *QueryRuleGroupsFullForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query rule groups full forbidden response has a 3xx status code
func (o *QueryRuleGroupsFullForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query rule groups full forbidden response has a 4xx status code
func (o *QueryRuleGroupsFullForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query rule groups full forbidden response has a 5xx status code
func (o *QueryRuleGroupsFullForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query rule groups full forbidden response a status code equal to that given
func (o *QueryRuleGroupsFullForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query rule groups full forbidden response
func (o *QueryRuleGroupsFullForbidden) Code() int {
	return 403
}

func (o *QueryRuleGroupsFullForbidden) Error() string {
	return fmt.Sprintf("[GET /ioarules/queries/rule-groups-full/v1][%d] queryRuleGroupsFullForbidden  %+v", 403, o.Payload)
}

func (o *QueryRuleGroupsFullForbidden) String() string {
	return fmt.Sprintf("[GET /ioarules/queries/rule-groups-full/v1][%d] queryRuleGroupsFullForbidden  %+v", 403, o.Payload)
}

func (o *QueryRuleGroupsFullForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryRuleGroupsFullForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryRuleGroupsFullNotFound creates a QueryRuleGroupsFullNotFound with default headers values
func NewQueryRuleGroupsFullNotFound() *QueryRuleGroupsFullNotFound {
	return &QueryRuleGroupsFullNotFound{}
}

/*
QueryRuleGroupsFullNotFound describes a response with status code 404, with default header values.

Not Found
*/
type QueryRuleGroupsFullNotFound struct {

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

// IsSuccess returns true when this query rule groups full not found response has a 2xx status code
func (o *QueryRuleGroupsFullNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query rule groups full not found response has a 3xx status code
func (o *QueryRuleGroupsFullNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query rule groups full not found response has a 4xx status code
func (o *QueryRuleGroupsFullNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this query rule groups full not found response has a 5xx status code
func (o *QueryRuleGroupsFullNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this query rule groups full not found response a status code equal to that given
func (o *QueryRuleGroupsFullNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the query rule groups full not found response
func (o *QueryRuleGroupsFullNotFound) Code() int {
	return 404
}

func (o *QueryRuleGroupsFullNotFound) Error() string {
	return fmt.Sprintf("[GET /ioarules/queries/rule-groups-full/v1][%d] queryRuleGroupsFullNotFound  %+v", 404, o.Payload)
}

func (o *QueryRuleGroupsFullNotFound) String() string {
	return fmt.Sprintf("[GET /ioarules/queries/rule-groups-full/v1][%d] queryRuleGroupsFullNotFound  %+v", 404, o.Payload)
}

func (o *QueryRuleGroupsFullNotFound) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryRuleGroupsFullNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryRuleGroupsFullTooManyRequests creates a QueryRuleGroupsFullTooManyRequests with default headers values
func NewQueryRuleGroupsFullTooManyRequests() *QueryRuleGroupsFullTooManyRequests {
	return &QueryRuleGroupsFullTooManyRequests{}
}

/*
QueryRuleGroupsFullTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryRuleGroupsFullTooManyRequests struct {

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

// IsSuccess returns true when this query rule groups full too many requests response has a 2xx status code
func (o *QueryRuleGroupsFullTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query rule groups full too many requests response has a 3xx status code
func (o *QueryRuleGroupsFullTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query rule groups full too many requests response has a 4xx status code
func (o *QueryRuleGroupsFullTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query rule groups full too many requests response has a 5xx status code
func (o *QueryRuleGroupsFullTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query rule groups full too many requests response a status code equal to that given
func (o *QueryRuleGroupsFullTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query rule groups full too many requests response
func (o *QueryRuleGroupsFullTooManyRequests) Code() int {
	return 429
}

func (o *QueryRuleGroupsFullTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /ioarules/queries/rule-groups-full/v1][%d] queryRuleGroupsFullTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryRuleGroupsFullTooManyRequests) String() string {
	return fmt.Sprintf("[GET /ioarules/queries/rule-groups-full/v1][%d] queryRuleGroupsFullTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryRuleGroupsFullTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryRuleGroupsFullTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
