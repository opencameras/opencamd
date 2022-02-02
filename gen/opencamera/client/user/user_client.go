// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateUser(params *CreateUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateUserOK, error)

	LogoutUser(params *LogoutUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LogoutUserOK, error)

	ResetPassword(params *ResetPasswordParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResetPasswordOK, error)

	TokenObtain(params *TokenObtainParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TokenObtainOK, error)

	TokenRefresh(params *TokenRefreshParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TokenRefreshOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateUser creates user

  Only one user is supported.
*/
func (a *Client) CreateUser(params *CreateUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createUser",
		Method:             "POST",
		PathPattern:        "/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateUserReader{formats: a.formats},
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
	success, ok := result.(*CreateUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LogoutUser logs out current logged in user session
*/
func (a *Client) LogoutUser(params *LogoutUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LogoutUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLogoutUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "logoutUser",
		Method:             "GET",
		PathPattern:        "/user/logout",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LogoutUserReader{formats: a.formats},
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
	success, ok := result.(*LogoutUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for logoutUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ResetPassword resets user password
*/
func (a *Client) ResetPassword(params *ResetPasswordParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResetPasswordOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResetPasswordParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "resetPassword",
		Method:             "PUT",
		PathPattern:        "/user/reset",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ResetPasswordReader{formats: a.formats},
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
	success, ok := result.(*ResetPasswordOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for resetPassword: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TokenObtain logs user into the system and obtain token
*/
func (a *Client) TokenObtain(params *TokenObtainParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TokenObtainOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTokenObtainParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "tokenObtain",
		Method:             "POST",
		PathPattern:        "/user/token/obtain",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TokenObtainReader{formats: a.formats},
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
	success, ok := result.(*TokenObtainOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tokenObtain: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TokenRefresh refreshes token
*/
func (a *Client) TokenRefresh(params *TokenRefreshParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TokenRefreshOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTokenRefreshParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "tokenRefresh",
		Method:             "POST",
		PathPattern:        "/user/token/refresh",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TokenRefreshReader{formats: a.formats},
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
	success, ok := result.(*TokenRefreshOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tokenRefresh: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
