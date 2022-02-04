// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/opencameras/opencamd/gen/opencamera/models"
	"github.com/opencameras/opencamd/gen/opencamera/restapi/operations"
	"github.com/opencameras/opencamd/gen/opencamera/restapi/operations/device"
	"github.com/opencameras/opencamd/gen/opencamera/restapi/operations/media"
	"github.com/opencameras/opencamd/gen/opencamera/restapi/operations/system"
	"github.com/opencameras/opencamd/gen/opencamera/restapi/operations/user"
)

//go:generate swagger generate server --target ../../opencamera --name OpenCamera --spec ../../../spec/opencameras.yaml --principal models.User --exclude-main

func configureFlags(api *operations.OpenCameraAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.OpenCameraAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()
	api.MultipartformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the Authorization header is set with the Basic scheme
	if api.BasicAuthAuth == nil {
		api.BasicAuthAuth = func(user string, pass string) (*models.User, error) {
			return nil, errors.NotImplemented("basic auth  (basicAuth) has not yet been implemented")
		}
	}
	// Applies when the "x-token" header is set
	if api.KeyAuth == nil {
		api.KeyAuth = func(token string) (*models.User, error) {
			return nil, errors.NotImplemented("api key auth (key) x-token from header param [x-token] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// system.UpgradeSystemMaxParseMemory = 32 << 20

	if api.UserCreateUserHandler == nil {
		api.UserCreateUserHandler = user.CreateUserHandlerFunc(func(params user.CreateUserParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation user.CreateUser has not yet been implemented")
		})
	}
	if api.MediaDownloadVideosHandler == nil {
		api.MediaDownloadVideosHandler = media.DownloadVideosHandlerFunc(func(params media.DownloadVideosParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation media.DownloadVideos has not yet been implemented")
		})
	}
	if api.DeviceGetDeviceInfoHandler == nil {
		api.DeviceGetDeviceInfoHandler = device.GetDeviceInfoHandlerFunc(func(params device.GetDeviceInfoParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation device.GetDeviceInfo has not yet been implemented")
		})
	}
	if api.SystemGetSystemInfoHandler == nil {
		api.SystemGetSystemInfoHandler = system.GetSystemInfoHandlerFunc(func(params system.GetSystemInfoParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation system.GetSystemInfo has not yet been implemented")
		})
	}
	if api.UserLogoutUserHandler == nil {
		api.UserLogoutUserHandler = user.LogoutUserHandlerFunc(func(params user.LogoutUserParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation user.LogoutUser has not yet been implemented")
		})
	}
	if api.UserResetPasswordHandler == nil {
		api.UserResetPasswordHandler = user.ResetPasswordHandlerFunc(func(params user.ResetPasswordParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation user.ResetPassword has not yet been implemented")
		})
	}
	if api.MediaStartLiveSessionHandler == nil {
		api.MediaStartLiveSessionHandler = media.StartLiveSessionHandlerFunc(func(params media.StartLiveSessionParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation media.StartLiveSession has not yet been implemented")
		})
	}
	if api.UserTokenObtainHandler == nil {
		api.UserTokenObtainHandler = user.TokenObtainHandlerFunc(func(params user.TokenObtainParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation user.TokenObtain has not yet been implemented")
		})
	}
	if api.UserTokenRefreshHandler == nil {
		api.UserTokenRefreshHandler = user.TokenRefreshHandlerFunc(func(params user.TokenRefreshParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation user.TokenRefresh has not yet been implemented")
		})
	}
	if api.MediaUpdateLiveConfigHandler == nil {
		api.MediaUpdateLiveConfigHandler = media.UpdateLiveConfigHandlerFunc(func(params media.UpdateLiveConfigParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation media.UpdateLiveConfig has not yet been implemented")
		})
	}
	if api.SystemUpgradeSystemHandler == nil {
		api.SystemUpgradeSystemHandler = system.UpgradeSystemHandlerFunc(func(params system.UpgradeSystemParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation system.UpgradeSystem has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
