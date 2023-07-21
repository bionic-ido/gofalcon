// Code generated by go-swagger; DO NOT EDIT.

package sample_uploads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new sample uploads API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sample uploads API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ArchiveDeleteV1(params *ArchiveDeleteV1Params, opts ...ClientOption) (*ArchiveDeleteV1Accepted, error)

	ArchiveGetV1(params *ArchiveGetV1Params, opts ...ClientOption) (*ArchiveGetV1OK, error)

	ArchiveListV1(params *ArchiveListV1Params, opts ...ClientOption) (*ArchiveListV1OK, error)

	ArchiveUploadV1(params *ArchiveUploadV1Params, opts ...ClientOption) (*ArchiveUploadV1OK, *ArchiveUploadV1Accepted, error)

	ArchiveUploadV2(params *ArchiveUploadV2Params, opts ...ClientOption) (*ArchiveUploadV2OK, *ArchiveUploadV2Accepted, error)

	DeleteSampleV3(params *DeleteSampleV3Params, opts ...ClientOption) (*DeleteSampleV3OK, error)

	ExtractionCreateV1(params *ExtractionCreateV1Params, opts ...ClientOption) (*ExtractionCreateV1OK, *ExtractionCreateV1Accepted, error)

	ExtractionGetV1(params *ExtractionGetV1Params, opts ...ClientOption) (*ExtractionGetV1OK, error)

	ExtractionListV1(params *ExtractionListV1Params, opts ...ClientOption) (*ExtractionListV1OK, error)

	GetSampleV3(params *GetSampleV3Params, opts ...ClientOption) (*GetSampleV3OK, error)

	UploadSampleV3(params *UploadSampleV3Params, opts ...ClientOption) (*UploadSampleV3OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ArchiveDeleteV1 deletes an archive that was uploaded previously
*/
func (a *Client) ArchiveDeleteV1(params *ArchiveDeleteV1Params, opts ...ClientOption) (*ArchiveDeleteV1Accepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewArchiveDeleteV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "ArchiveDeleteV1",
		Method:             "DELETE",
		PathPattern:        "/archives/entities/archives/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ArchiveDeleteV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ArchiveDeleteV1Accepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ArchiveDeleteV1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ArchiveGetV1 retrieves the archives upload operation statuses status done means that archive was processed successfully status error means that archive was not processed successfully
*/
func (a *Client) ArchiveGetV1(params *ArchiveGetV1Params, opts ...ClientOption) (*ArchiveGetV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewArchiveGetV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "ArchiveGetV1",
		Method:             "GET",
		PathPattern:        "/archives/entities/archives/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ArchiveGetV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ArchiveGetV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ArchiveGetV1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ArchiveListV1 retrieves the archives files in chunks
*/
func (a *Client) ArchiveListV1(params *ArchiveListV1Params, opts ...ClientOption) (*ArchiveListV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewArchiveListV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "ArchiveListV1",
		Method:             "GET",
		PathPattern:        "/archives/entities/archive-files/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ArchiveListV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ArchiveListV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ArchiveListV1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ArchiveUploadV1 uploads an archive and extracts files list from it operation is asynchronous use archives entities archives v1 to check the status after uploading use archives entities extractions v1 to copy the file to internal storage making it available for content analysis this method is deprecated in favor of archives entities archives v2
*/
func (a *Client) ArchiveUploadV1(params *ArchiveUploadV1Params, opts ...ClientOption) (*ArchiveUploadV1OK, *ArchiveUploadV1Accepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewArchiveUploadV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "ArchiveUploadV1",
		Method:             "POST",
		PathPattern:        "/archives/entities/archives/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/octet-stream", "application/x-7z-compressed", "application/x-zip-compressed", "application/zip"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ArchiveUploadV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ArchiveUploadV1OK:
		return value, nil, nil
	case *ArchiveUploadV1Accepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for sample_uploads: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ArchiveUploadV2 uploads an archive and extracts files list from it operation is asynchronous use archives entities archives v1 to check the status after uploading use archives entities extractions v1 to copy the file to internal storage making it available for content analysis
*/
func (a *Client) ArchiveUploadV2(params *ArchiveUploadV2Params, opts ...ClientOption) (*ArchiveUploadV2OK, *ArchiveUploadV2Accepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewArchiveUploadV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "ArchiveUploadV2",
		Method:             "POST",
		PathPattern:        "/archives/entities/archives/v2",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ArchiveUploadV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ArchiveUploadV2OK:
		return value, nil, nil
	case *ArchiveUploadV2Accepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ArchiveUploadV2Default)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteSampleV3 removes a sample including file meta and submissions from the collection
*/
func (a *Client) DeleteSampleV3(params *DeleteSampleV3Params, opts ...ClientOption) (*DeleteSampleV3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSampleV3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteSampleV3",
		Method:             "DELETE",
		PathPattern:        "/samples/entities/samples/v3",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSampleV3Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteSampleV3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteSampleV3Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ExtractionCreateV1 extracts files from an uploaded archive and copies them to internal storage making it available for content analysis
*/
func (a *Client) ExtractionCreateV1(params *ExtractionCreateV1Params, opts ...ClientOption) (*ExtractionCreateV1OK, *ExtractionCreateV1Accepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtractionCreateV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExtractionCreateV1",
		Method:             "POST",
		PathPattern:        "/archives/entities/extractions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtractionCreateV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ExtractionCreateV1OK:
		return value, nil, nil
	case *ExtractionCreateV1Accepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ExtractionCreateV1Default)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ExtractionGetV1 retrieves the files extraction operation statuses status done means that all files were processed successfully status error means that at least one of the file could not be processed
*/
func (a *Client) ExtractionGetV1(params *ExtractionGetV1Params, opts ...ClientOption) (*ExtractionGetV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtractionGetV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExtractionGetV1",
		Method:             "GET",
		PathPattern:        "/archives/entities/extractions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtractionGetV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ExtractionGetV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ExtractionGetV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ExtractionListV1 retrieves the files extractions in chunks status done means that all files were processed successfully status error means that at least one of the file could not be processed
*/
func (a *Client) ExtractionListV1(params *ExtractionListV1Params, opts ...ClientOption) (*ExtractionListV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtractionListV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExtractionListV1",
		Method:             "GET",
		PathPattern:        "/archives/entities/extraction-files/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtractionListV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ExtractionListV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ExtractionListV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetSampleV3 retrieves the file associated with the given ID s h a256
*/
func (a *Client) GetSampleV3(params *GetSampleV3Params, opts ...ClientOption) (*GetSampleV3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSampleV3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSampleV3",
		Method:             "GET",
		PathPattern:        "/samples/entities/samples/v3",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSampleV3Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSampleV3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSampleV3Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UploadSampleV3 uploads a file for further cloud analysis after uploading call the specific analysis API endpoint
*/
func (a *Client) UploadSampleV3(params *UploadSampleV3Params, opts ...ClientOption) (*UploadSampleV3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadSampleV3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "UploadSampleV3",
		Method:             "POST",
		PathPattern:        "/samples/entities/samples/v3",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/octet-stream", "multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UploadSampleV3Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UploadSampleV3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UploadSampleV3Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
