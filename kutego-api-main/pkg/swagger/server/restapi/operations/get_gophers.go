// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetCatsHandlerFunc turns a function with the right signature into a get cats handler
type GetCatsHandlerFunc func(GetCatsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCatsHandlerFunc) Handle(params GetCatsParams) middleware.Responder {
	return fn(params)
}

// GetCatsHandler interface for that can handle valid get cats params
type GetCatsHandler interface {
	Handle(GetCatsParams) middleware.Responder
}

// NewGetCats creates a new http.Handler for the get cats operation
func NewGetCats(ctx *middleware.Context, handler GetCatsHandler) *GetCats {
	return &GetCats{Context: ctx, Handler: handler}
}

/* GetCats swagger:route GET /cats getCats

List all the cat

*/
type GetCats struct {
	Context *middleware.Context
	Handler GetCatsHandler
}

func (o *GetCats) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetCatsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
