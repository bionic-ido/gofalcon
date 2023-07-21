// Code generated by go-swagger; DO NOT EDIT.

package spotlight_evaluation_logic

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

// CombinedQueryEvaluationLogicReader is a Reader for the CombinedQueryEvaluationLogic structure.
type CombinedQueryEvaluationLogicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CombinedQueryEvaluationLogicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCombinedQueryEvaluationLogicOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCombinedQueryEvaluationLogicBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCombinedQueryEvaluationLogicForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCombinedQueryEvaluationLogicTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCombinedQueryEvaluationLogicInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /spotlight/combined/evaluation-logic/v1] combinedQueryEvaluationLogic", response, response.Code())
	}
}

// NewCombinedQueryEvaluationLogicOK creates a CombinedQueryEvaluationLogicOK with default headers values
func NewCombinedQueryEvaluationLogicOK() *CombinedQueryEvaluationLogicOK {
	return &CombinedQueryEvaluationLogicOK{}
}

/*
CombinedQueryEvaluationLogicOK describes a response with status code 200, with default header values.

OK
*/
type CombinedQueryEvaluationLogicOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainSPAPIEvaluationLogicCombinedResponseV1
}

// IsSuccess returns true when this combined query evaluation logic o k response has a 2xx status code
func (o *CombinedQueryEvaluationLogicOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this combined query evaluation logic o k response has a 3xx status code
func (o *CombinedQueryEvaluationLogicOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this combined query evaluation logic o k response has a 4xx status code
func (o *CombinedQueryEvaluationLogicOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this combined query evaluation logic o k response has a 5xx status code
func (o *CombinedQueryEvaluationLogicOK) IsServerError() bool {
	return false
}

// IsCode returns true when this combined query evaluation logic o k response a status code equal to that given
func (o *CombinedQueryEvaluationLogicOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the combined query evaluation logic o k response
func (o *CombinedQueryEvaluationLogicOK) Code() int {
	return 200
}

func (o *CombinedQueryEvaluationLogicOK) Error() string {
	return fmt.Sprintf("[GET /spotlight/combined/evaluation-logic/v1][%d] combinedQueryEvaluationLogicOK  %+v", 200, o.Payload)
}

func (o *CombinedQueryEvaluationLogicOK) String() string {
	return fmt.Sprintf("[GET /spotlight/combined/evaluation-logic/v1][%d] combinedQueryEvaluationLogicOK  %+v", 200, o.Payload)
}

func (o *CombinedQueryEvaluationLogicOK) GetPayload() *models.DomainSPAPIEvaluationLogicCombinedResponseV1 {
	return o.Payload
}

func (o *CombinedQueryEvaluationLogicOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainSPAPIEvaluationLogicCombinedResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCombinedQueryEvaluationLogicBadRequest creates a CombinedQueryEvaluationLogicBadRequest with default headers values
func NewCombinedQueryEvaluationLogicBadRequest() *CombinedQueryEvaluationLogicBadRequest {
	return &CombinedQueryEvaluationLogicBadRequest{}
}

/*
CombinedQueryEvaluationLogicBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CombinedQueryEvaluationLogicBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64
}

// IsSuccess returns true when this combined query evaluation logic bad request response has a 2xx status code
func (o *CombinedQueryEvaluationLogicBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this combined query evaluation logic bad request response has a 3xx status code
func (o *CombinedQueryEvaluationLogicBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this combined query evaluation logic bad request response has a 4xx status code
func (o *CombinedQueryEvaluationLogicBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this combined query evaluation logic bad request response has a 5xx status code
func (o *CombinedQueryEvaluationLogicBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this combined query evaluation logic bad request response a status code equal to that given
func (o *CombinedQueryEvaluationLogicBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the combined query evaluation logic bad request response
func (o *CombinedQueryEvaluationLogicBadRequest) Code() int {
	return 400
}

func (o *CombinedQueryEvaluationLogicBadRequest) Error() string {
	return fmt.Sprintf("[GET /spotlight/combined/evaluation-logic/v1][%d] combinedQueryEvaluationLogicBadRequest ", 400)
}

func (o *CombinedQueryEvaluationLogicBadRequest) String() string {
	return fmt.Sprintf("[GET /spotlight/combined/evaluation-logic/v1][%d] combinedQueryEvaluationLogicBadRequest ", 400)
}

func (o *CombinedQueryEvaluationLogicBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	return nil
}

// NewCombinedQueryEvaluationLogicForbidden creates a CombinedQueryEvaluationLogicForbidden with default headers values
func NewCombinedQueryEvaluationLogicForbidden() *CombinedQueryEvaluationLogicForbidden {
	return &CombinedQueryEvaluationLogicForbidden{}
}

/*
CombinedQueryEvaluationLogicForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CombinedQueryEvaluationLogicForbidden struct {

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

// IsSuccess returns true when this combined query evaluation logic forbidden response has a 2xx status code
func (o *CombinedQueryEvaluationLogicForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this combined query evaluation logic forbidden response has a 3xx status code
func (o *CombinedQueryEvaluationLogicForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this combined query evaluation logic forbidden response has a 4xx status code
func (o *CombinedQueryEvaluationLogicForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this combined query evaluation logic forbidden response has a 5xx status code
func (o *CombinedQueryEvaluationLogicForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this combined query evaluation logic forbidden response a status code equal to that given
func (o *CombinedQueryEvaluationLogicForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the combined query evaluation logic forbidden response
func (o *CombinedQueryEvaluationLogicForbidden) Code() int {
	return 403
}

func (o *CombinedQueryEvaluationLogicForbidden) Error() string {
	return fmt.Sprintf("[GET /spotlight/combined/evaluation-logic/v1][%d] combinedQueryEvaluationLogicForbidden  %+v", 403, o.Payload)
}

func (o *CombinedQueryEvaluationLogicForbidden) String() string {
	return fmt.Sprintf("[GET /spotlight/combined/evaluation-logic/v1][%d] combinedQueryEvaluationLogicForbidden  %+v", 403, o.Payload)
}

func (o *CombinedQueryEvaluationLogicForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CombinedQueryEvaluationLogicForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCombinedQueryEvaluationLogicTooManyRequests creates a CombinedQueryEvaluationLogicTooManyRequests with default headers values
func NewCombinedQueryEvaluationLogicTooManyRequests() *CombinedQueryEvaluationLogicTooManyRequests {
	return &CombinedQueryEvaluationLogicTooManyRequests{}
}

/*
CombinedQueryEvaluationLogicTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CombinedQueryEvaluationLogicTooManyRequests struct {

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

// IsSuccess returns true when this combined query evaluation logic too many requests response has a 2xx status code
func (o *CombinedQueryEvaluationLogicTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this combined query evaluation logic too many requests response has a 3xx status code
func (o *CombinedQueryEvaluationLogicTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this combined query evaluation logic too many requests response has a 4xx status code
func (o *CombinedQueryEvaluationLogicTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this combined query evaluation logic too many requests response has a 5xx status code
func (o *CombinedQueryEvaluationLogicTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this combined query evaluation logic too many requests response a status code equal to that given
func (o *CombinedQueryEvaluationLogicTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the combined query evaluation logic too many requests response
func (o *CombinedQueryEvaluationLogicTooManyRequests) Code() int {
	return 429
}

func (o *CombinedQueryEvaluationLogicTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /spotlight/combined/evaluation-logic/v1][%d] combinedQueryEvaluationLogicTooManyRequests  %+v", 429, o.Payload)
}

func (o *CombinedQueryEvaluationLogicTooManyRequests) String() string {
	return fmt.Sprintf("[GET /spotlight/combined/evaluation-logic/v1][%d] combinedQueryEvaluationLogicTooManyRequests  %+v", 429, o.Payload)
}

func (o *CombinedQueryEvaluationLogicTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CombinedQueryEvaluationLogicTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCombinedQueryEvaluationLogicInternalServerError creates a CombinedQueryEvaluationLogicInternalServerError with default headers values
func NewCombinedQueryEvaluationLogicInternalServerError() *CombinedQueryEvaluationLogicInternalServerError {
	return &CombinedQueryEvaluationLogicInternalServerError{}
}

/*
CombinedQueryEvaluationLogicInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CombinedQueryEvaluationLogicInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64
}

// IsSuccess returns true when this combined query evaluation logic internal server error response has a 2xx status code
func (o *CombinedQueryEvaluationLogicInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this combined query evaluation logic internal server error response has a 3xx status code
func (o *CombinedQueryEvaluationLogicInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this combined query evaluation logic internal server error response has a 4xx status code
func (o *CombinedQueryEvaluationLogicInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this combined query evaluation logic internal server error response has a 5xx status code
func (o *CombinedQueryEvaluationLogicInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this combined query evaluation logic internal server error response a status code equal to that given
func (o *CombinedQueryEvaluationLogicInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the combined query evaluation logic internal server error response
func (o *CombinedQueryEvaluationLogicInternalServerError) Code() int {
	return 500
}

func (o *CombinedQueryEvaluationLogicInternalServerError) Error() string {
	return fmt.Sprintf("[GET /spotlight/combined/evaluation-logic/v1][%d] combinedQueryEvaluationLogicInternalServerError ", 500)
}

func (o *CombinedQueryEvaluationLogicInternalServerError) String() string {
	return fmt.Sprintf("[GET /spotlight/combined/evaluation-logic/v1][%d] combinedQueryEvaluationLogicInternalServerError ", 500)
}

func (o *CombinedQueryEvaluationLogicInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	return nil
}
