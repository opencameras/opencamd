// Code generated by go-swagger; DO NOT EDIT.

package media

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/opencameras/opencamd/gen/opencamera/models"
)

// DownloadVideosHandlerFunc turns a function with the right signature into a download videos handler
type DownloadVideosHandlerFunc func(DownloadVideosParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn DownloadVideosHandlerFunc) Handle(params DownloadVideosParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// DownloadVideosHandler interface for that can handle valid download videos params
type DownloadVideosHandler interface {
	Handle(DownloadVideosParams, *models.Principal) middleware.Responder
}

// NewDownloadVideos creates a new http.Handler for the download videos operation
func NewDownloadVideos(ctx *middleware.Context, handler DownloadVideosHandler) *DownloadVideos {
	return &DownloadVideos{Context: ctx, Handler: handler}
}

/* DownloadVideos swagger:route GET /media/vod/{start}/{end} media downloadVideos

Download recorded videos

Download record videos

*/
type DownloadVideos struct {
	Context *middleware.Context
	Handler DownloadVideosHandler
}

func (o *DownloadVideos) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDownloadVideosParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
