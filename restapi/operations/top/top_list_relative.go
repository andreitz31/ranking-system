// Code generated by go-swagger; DO NOT EDIT.

package top

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// TopListRelativeHandlerFunc turns a function with the right signature into a top list relative handler
type TopListRelativeHandlerFunc func(TopListRelativeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn TopListRelativeHandlerFunc) Handle(params TopListRelativeParams) middleware.Responder {
	return fn(params)
}

// TopListRelativeHandler interface for that can handle valid top list relative params
type TopListRelativeHandler interface {
	Handle(TopListRelativeParams) middleware.Responder
}

// NewTopListRelative creates a new http.Handler for the top list relative operation
func NewTopListRelative(ctx *middleware.Context, handler TopListRelativeHandler) *TopListRelative {
	return &TopListRelative{Context: ctx, Handler: handler}
}

/* TopListRelative swagger:route GET /top/position/{position}/players/{players} top topListRelative

Get relative top x players

It returns a list of the first x players from the top relative to a position

*/
type TopListRelative struct {
	Context *middleware.Context
	Handler TopListRelativeHandler
}

func (o *TopListRelative) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewTopListRelativeParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
