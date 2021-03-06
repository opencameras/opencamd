// Code generated by go-swagger; DO NOT EDIT.

package media

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new media API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for media API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DownloadVideos(params *DownloadVideosParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DownloadVideosOK, error)

	StartLiveSession(params *StartLiveSessionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StartLiveSessionOK, error)

	UpdateLiveConfig(params *UpdateLiveConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateLiveConfigOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DownloadVideos downloads recorded videos

  Download record videos
*/
func (a *Client) DownloadVideos(params *DownloadVideosParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DownloadVideosOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDownloadVideosParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "downloadVideos",
		Method:             "GET",
		PathPattern:        "/media/vod/{start}/{end}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DownloadVideosReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*DownloadVideosOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for downloadVideos: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StartLiveSession starts live media session

  Start media session to view camera live status
*/
func (a *Client) StartLiveSession(params *StartLiveSessionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StartLiveSessionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartLiveSessionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "startLiveSession",
		Method:             "POST",
		PathPattern:        "/media/live/session",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StartLiveSessionReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*StartLiveSessionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for startLiveSession: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateLiveConfig updates live session configuration

  Update live session configuration, such as ice servers
*/
func (a *Client) UpdateLiveConfig(params *UpdateLiveConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateLiveConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateLiveConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateLiveConfig",
		Method:             "PUT",
		PathPattern:        "/media/live/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateLiveConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*UpdateLiveConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateLiveConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
