// Code generated by go-swagger; DO NOT EDIT.

package top

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// TopListHandlerFunc turns a function with the right signature into a top list handler
type TopListHandlerFunc func(TopListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn TopListHandlerFunc) Handle(params TopListParams) middleware.Responder {
	return fn(params)
}

// TopListHandler interface for that can handle valid top list params
type TopListHandler interface {
	Handle(TopListParams) middleware.Responder
}

// NewTopList creates a new http.Handler for the top list operation
func NewTopList(ctx *middleware.Context, handler TopListHandler) *TopList {
	return &TopList{Context: ctx, Handler: handler}
}

/* TopList swagger:route GET /top/{number} top topList

Get top x players

It returns a list of the first x players from the top

*/
type TopList struct {
	Context *middleware.Context
	Handler TopListHandler
}

func (o *TopList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewTopListParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
