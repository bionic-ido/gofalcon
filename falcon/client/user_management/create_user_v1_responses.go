// Code generated by go-swagger; DO NOT EDIT.

package user_management

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

// CreateUserV1Reader is a Reader for the CreateUserV1 structure.
type CreateUserV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateUserV1Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateUserV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateUserV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateUserV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateUserV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /user-management/entities/users/v1] createUserV1", response, response.Code())
	}
}

// NewCreateUserV1Created creates a CreateUserV1Created with default headers values
func NewCreateUserV1Created() *CreateUserV1Created {
	return &CreateUserV1Created{}
}

/*
CreateUserV1Created describes a response with status code 201, with default header values.

OK
*/
type CreateUserV1Created struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainMsaEntitiesUsersResponse
}

// IsSuccess returns true when this create user v1 created response has a 2xx status code
func (o *CreateUserV1Created) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create user v1 created response has a 3xx status code
func (o *CreateUserV1Created) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user v1 created response has a 4xx status code
func (o *CreateUserV1Created) IsClientError() bool {
	return false
}

// IsServerError returns true when this create user v1 created response has a 5xx status code
func (o *CreateUserV1Created) IsServerError() bool {
	return false
}

// IsCode returns true when this create user v1 created response a status code equal to that given
func (o *CreateUserV1Created) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create user v1 created response
func (o *CreateUserV1Created) Code() int {
	return 201
}

func (o *CreateUserV1Created) Error() string {
	return fmt.Sprintf("[POST /user-management/entities/users/v1][%d] createUserV1Created  %+v", 201, o.Payload)
}

func (o *CreateUserV1Created) String() string {
	return fmt.Sprintf("[POST /user-management/entities/users/v1][%d] createUserV1Created  %+v", 201, o.Payload)
}

func (o *CreateUserV1Created) GetPayload() *models.DomainMsaEntitiesUsersResponse {
	return o.Payload
}

func (o *CreateUserV1Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainMsaEntitiesUsersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserV1BadRequest creates a CreateUserV1BadRequest with default headers values
func NewCreateUserV1BadRequest() *CreateUserV1BadRequest {
	return &CreateUserV1BadRequest{}
}

/*
CreateUserV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateUserV1BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainMsaEntitiesUsersResponse
}

// IsSuccess returns true when this create user v1 bad request response has a 2xx status code
func (o *CreateUserV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user v1 bad request response has a 3xx status code
func (o *CreateUserV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user v1 bad request response has a 4xx status code
func (o *CreateUserV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user v1 bad request response has a 5xx status code
func (o *CreateUserV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create user v1 bad request response a status code equal to that given
func (o *CreateUserV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create user v1 bad request response
func (o *CreateUserV1BadRequest) Code() int {
	return 400
}

func (o *CreateUserV1BadRequest) Error() string {
	return fmt.Sprintf("[POST /user-management/entities/users/v1][%d] createUserV1BadRequest  %+v", 400, o.Payload)
}

func (o *CreateUserV1BadRequest) String() string {
	return fmt.Sprintf("[POST /user-management/entities/users/v1][%d] createUserV1BadRequest  %+v", 400, o.Payload)
}

func (o *CreateUserV1BadRequest) GetPayload() *models.DomainMsaEntitiesUsersResponse {
	return o.Payload
}

func (o *CreateUserV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainMsaEntitiesUsersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserV1Forbidden creates a CreateUserV1Forbidden with default headers values
func NewCreateUserV1Forbidden() *CreateUserV1Forbidden {
	return &CreateUserV1Forbidden{}
}

/*
CreateUserV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateUserV1Forbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainMsaEntitiesUsersResponse
}

// IsSuccess returns true when this create user v1 forbidden response has a 2xx status code
func (o *CreateUserV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user v1 forbidden response has a 3xx status code
func (o *CreateUserV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user v1 forbidden response has a 4xx status code
func (o *CreateUserV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user v1 forbidden response has a 5xx status code
func (o *CreateUserV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create user v1 forbidden response a status code equal to that given
func (o *CreateUserV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create user v1 forbidden response
func (o *CreateUserV1Forbidden) Code() int {
	return 403
}

func (o *CreateUserV1Forbidden) Error() string {
	return fmt.Sprintf("[POST /user-management/entities/users/v1][%d] createUserV1Forbidden  %+v", 403, o.Payload)
}

func (o *CreateUserV1Forbidden) String() string {
	return fmt.Sprintf("[POST /user-management/entities/users/v1][%d] createUserV1Forbidden  %+v", 403, o.Payload)
}

func (o *CreateUserV1Forbidden) GetPayload() *models.DomainMsaEntitiesUsersResponse {
	return o.Payload
}

func (o *CreateUserV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainMsaEntitiesUsersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserV1TooManyRequests creates a CreateUserV1TooManyRequests with default headers values
func NewCreateUserV1TooManyRequests() *CreateUserV1TooManyRequests {
	return &CreateUserV1TooManyRequests{}
}

/*
CreateUserV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateUserV1TooManyRequests struct {

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

// IsSuccess returns true when this create user v1 too many requests response has a 2xx status code
func (o *CreateUserV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user v1 too many requests response has a 3xx status code
func (o *CreateUserV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user v1 too many requests response has a 4xx status code
func (o *CreateUserV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user v1 too many requests response has a 5xx status code
func (o *CreateUserV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this create user v1 too many requests response a status code equal to that given
func (o *CreateUserV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the create user v1 too many requests response
func (o *CreateUserV1TooManyRequests) Code() int {
	return 429
}

func (o *CreateUserV1TooManyRequests) Error() string {
	return fmt.Sprintf("[POST /user-management/entities/users/v1][%d] createUserV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateUserV1TooManyRequests) String() string {
	return fmt.Sprintf("[POST /user-management/entities/users/v1][%d] createUserV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateUserV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CreateUserV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateUserV1InternalServerError creates a CreateUserV1InternalServerError with default headers values
func NewCreateUserV1InternalServerError() *CreateUserV1InternalServerError {
	return &CreateUserV1InternalServerError{}
}

/*
CreateUserV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateUserV1InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainMsaEntitiesUsersResponse
}

// IsSuccess returns true when this create user v1 internal server error response has a 2xx status code
func (o *CreateUserV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user v1 internal server error response has a 3xx status code
func (o *CreateUserV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user v1 internal server error response has a 4xx status code
func (o *CreateUserV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create user v1 internal server error response has a 5xx status code
func (o *CreateUserV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create user v1 internal server error response a status code equal to that given
func (o *CreateUserV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create user v1 internal server error response
func (o *CreateUserV1InternalServerError) Code() int {
	return 500
}

func (o *CreateUserV1InternalServerError) Error() string {
	return fmt.Sprintf("[POST /user-management/entities/users/v1][%d] createUserV1InternalServerError  %+v", 500, o.Payload)
}

func (o *CreateUserV1InternalServerError) String() string {
	return fmt.Sprintf("[POST /user-management/entities/users/v1][%d] createUserV1InternalServerError  %+v", 500, o.Payload)
}

func (o *CreateUserV1InternalServerError) GetPayload() *models.DomainMsaEntitiesUsersResponse {
	return o.Payload
}

func (o *CreateUserV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainMsaEntitiesUsersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
