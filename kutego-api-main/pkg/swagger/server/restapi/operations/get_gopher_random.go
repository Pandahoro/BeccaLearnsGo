// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetCatRandomHandlerFunc turns a function with the right signature into a get cat random handler
type GetCatRandomHandlerFunc func(GetCatRandomParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCatRandomHandlerFunc) Handle(params GetCatRandomParams) middleware.Responder {
	return fn(params)
}

// GetCatRandomHandler interface for that can handle valid get cat random params
type GetCatRandomHandler interface {
	Handle(GetCatRandomParams) middleware.Responder
}

// NewGetCatRandom creates a new http.Handler for the get cat random operation
func NewGetCatRandom(ctx *middleware.Context, handler GetCatRandomHandler) *GetCatRandom {
	return &GetCatRandom{Context: ctx, Handler: handler}
}

/* GetCatRandom swagger:route GET /cat/random getCatRandom

Return a random Cat Image

*/
type GetCatRandom struct {
	Context *middleware.Context
	Handler GetCatRandomHandler
}

func (o *GetCatRandom) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetCatRandomParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}