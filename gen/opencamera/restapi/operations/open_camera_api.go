// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/opencameras/opencamd/gen/opencamera/models"
	"github.com/opencameras/opencamd/gen/opencamera/restapi/operations/device"
	"github.com/opencameras/opencamd/gen/opencamera/restapi/operations/media"
	"github.com/opencameras/opencamd/gen/opencamera/restapi/operations/system"
	"github.com/opencameras/opencamd/gen/opencamera/restapi/operations/user"
)

// NewOpenCameraAPI creates a new OpenCamera instance
func NewOpenCameraAPI(spec *loads.Document) *OpenCameraAPI {
	return &OpenCameraAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		MultipartformConsumer: runtime.DiscardConsumer,

		JSONProducer: runtime.JSONProducer(),

		UserCreateUserHandler: user.CreateUserHandlerFunc(func(params user.CreateUserParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.CreateUser has not yet been implemented")
		}),
		MediaDownloadVideosHandler: media.DownloadVideosHandlerFunc(func(params media.DownloadVideosParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation media.DownloadVideos has not yet been implemented")
		}),
		MediaGetCameraSDPHandler: media.GetCameraSDPHandlerFunc(func(params media.GetCameraSDPParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation media.GetCameraSDP has not yet been implemented")
		}),
		DeviceGetDeviceInfoHandler: device.GetDeviceInfoHandlerFunc(func(params device.GetDeviceInfoParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation device.GetDeviceInfo has not yet been implemented")
		}),
		SystemGetSystemInfoHandler: system.GetSystemInfoHandlerFunc(func(params system.GetSystemInfoParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation system.GetSystemInfo has not yet been implemented")
		}),
		UserLogoutUserHandler: user.LogoutUserHandlerFunc(func(params user.LogoutUserParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.LogoutUser has not yet been implemented")
		}),
		UserResetPasswordHandler: user.ResetPasswordHandlerFunc(func(params user.ResetPasswordParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.ResetPassword has not yet been implemented")
		}),
		MediaStartLiveSessionHandler: media.StartLiveSessionHandlerFunc(func(params media.StartLiveSessionParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation media.StartLiveSession has not yet been implemented")
		}),
		UserTokenObtainHandler: user.TokenObtainHandlerFunc(func(params user.TokenObtainParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.TokenObtain has not yet been implemented")
		}),
		UserTokenRefreshHandler: user.TokenRefreshHandlerFunc(func(params user.TokenRefreshParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.TokenRefresh has not yet been implemented")
		}),
		MediaUpdateStunServersHandler: media.UpdateStunServersHandlerFunc(func(params media.UpdateStunServersParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation media.UpdateStunServers has not yet been implemented")
		}),
		SystemUpgradeSystemHandler: system.UpgradeSystemHandlerFunc(func(params system.UpgradeSystemParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation system.UpgradeSystem has not yet been implemented")
		}),

		// Applies when the Authorization header is set with the Basic scheme
		BasicAuthAuth: func(user string, pass string) (*models.Principal, error) {
			return nil, errors.NotImplemented("basic auth  (basicAuth) has not yet been implemented")
		},
		// Applies when the "x-token" header is set
		KeyAuth: func(token string) (*models.Principal, error) {
			return nil, errors.NotImplemented("api key auth (key) x-token from header param [x-token] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*OpenCameraAPI Open Camera APIs to provide a secure / open security camera solution. */
type OpenCameraAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// MultipartformConsumer registers a consumer for the following mime types:
	//   - multipart/form-data
	MultipartformConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// BasicAuthAuth registers a function that takes username and password and returns a principal
	// it performs authentication with basic auth
	BasicAuthAuth func(string, string) (*models.Principal, error)

	// KeyAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key x-token provided in the header
	KeyAuth func(string) (*models.Principal, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// UserCreateUserHandler sets the operation handler for the create user operation
	UserCreateUserHandler user.CreateUserHandler
	// MediaDownloadVideosHandler sets the operation handler for the download videos operation
	MediaDownloadVideosHandler media.DownloadVideosHandler
	// MediaGetCameraSDPHandler sets the operation handler for the get camera s d p operation
	MediaGetCameraSDPHandler media.GetCameraSDPHandler
	// DeviceGetDeviceInfoHandler sets the operation handler for the get device info operation
	DeviceGetDeviceInfoHandler device.GetDeviceInfoHandler
	// SystemGetSystemInfoHandler sets the operation handler for the get system info operation
	SystemGetSystemInfoHandler system.GetSystemInfoHandler
	// UserLogoutUserHandler sets the operation handler for the logout user operation
	UserLogoutUserHandler user.LogoutUserHandler
	// UserResetPasswordHandler sets the operation handler for the reset password operation
	UserResetPasswordHandler user.ResetPasswordHandler
	// MediaStartLiveSessionHandler sets the operation handler for the start live session operation
	MediaStartLiveSessionHandler media.StartLiveSessionHandler
	// UserTokenObtainHandler sets the operation handler for the token obtain operation
	UserTokenObtainHandler user.TokenObtainHandler
	// UserTokenRefreshHandler sets the operation handler for the token refresh operation
	UserTokenRefreshHandler user.TokenRefreshHandler
	// MediaUpdateStunServersHandler sets the operation handler for the update stun servers operation
	MediaUpdateStunServersHandler media.UpdateStunServersHandler
	// SystemUpgradeSystemHandler sets the operation handler for the upgrade system operation
	SystemUpgradeSystemHandler system.UpgradeSystemHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *OpenCameraAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *OpenCameraAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *OpenCameraAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *OpenCameraAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *OpenCameraAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *OpenCameraAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *OpenCameraAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *OpenCameraAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *OpenCameraAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the OpenCameraAPI
func (o *OpenCameraAPI) Validate() error {
	var unregistered []string

	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.BasicAuthAuth == nil {
		unregistered = append(unregistered, "BasicAuthAuth")
	}
	if o.KeyAuth == nil {
		unregistered = append(unregistered, "XTokenAuth")
	}

	if o.UserCreateUserHandler == nil {
		unregistered = append(unregistered, "user.CreateUserHandler")
	}
	if o.MediaDownloadVideosHandler == nil {
		unregistered = append(unregistered, "media.DownloadVideosHandler")
	}
	if o.MediaGetCameraSDPHandler == nil {
		unregistered = append(unregistered, "media.GetCameraSDPHandler")
	}
	if o.DeviceGetDeviceInfoHandler == nil {
		unregistered = append(unregistered, "device.GetDeviceInfoHandler")
	}
	if o.SystemGetSystemInfoHandler == nil {
		unregistered = append(unregistered, "system.GetSystemInfoHandler")
	}
	if o.UserLogoutUserHandler == nil {
		unregistered = append(unregistered, "user.LogoutUserHandler")
	}
	if o.UserResetPasswordHandler == nil {
		unregistered = append(unregistered, "user.ResetPasswordHandler")
	}
	if o.MediaStartLiveSessionHandler == nil {
		unregistered = append(unregistered, "media.StartLiveSessionHandler")
	}
	if o.UserTokenObtainHandler == nil {
		unregistered = append(unregistered, "user.TokenObtainHandler")
	}
	if o.UserTokenRefreshHandler == nil {
		unregistered = append(unregistered, "user.TokenRefreshHandler")
	}
	if o.MediaUpdateStunServersHandler == nil {
		unregistered = append(unregistered, "media.UpdateStunServersHandler")
	}
	if o.SystemUpgradeSystemHandler == nil {
		unregistered = append(unregistered, "system.UpgradeSystemHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *OpenCameraAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *OpenCameraAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "basicAuth":
			result[name] = o.BasicAuthenticator(func(username, password string) (interface{}, error) {
				return o.BasicAuthAuth(username, password)
			})

		case "key":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.KeyAuth(token)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *OpenCameraAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *OpenCameraAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *OpenCameraAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *OpenCameraAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the open camera API
func (o *OpenCameraAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *OpenCameraAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/user"] = user.NewCreateUser(o.context, o.UserCreateUserHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/media/vod/{start}/{end}"] = media.NewDownloadVideos(o.context, o.MediaDownloadVideosHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/media/sdp"] = media.NewGetCameraSDP(o.context, o.MediaGetCameraSDPHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/device"] = device.NewGetDeviceInfo(o.context, o.DeviceGetDeviceInfoHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/system"] = system.NewGetSystemInfo(o.context, o.SystemGetSystemInfoHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/logout"] = user.NewLogoutUser(o.context, o.UserLogoutUserHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/user/reset"] = user.NewResetPassword(o.context, o.UserResetPasswordHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/media/live/session"] = media.NewStartLiveSession(o.context, o.MediaStartLiveSessionHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/user/token/obtain"] = user.NewTokenObtain(o.context, o.UserTokenObtainHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/user/token/refresh"] = user.NewTokenRefresh(o.context, o.UserTokenRefreshHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/media/stubservers"] = media.NewUpdateStunServers(o.context, o.MediaUpdateStunServersHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/system/upgrade"] = system.NewUpgradeSystem(o.context, o.SystemUpgradeSystemHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *OpenCameraAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *OpenCameraAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *OpenCameraAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *OpenCameraAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *OpenCameraAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
