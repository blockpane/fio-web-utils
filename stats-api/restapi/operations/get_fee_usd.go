// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetFeeUsdHandlerFunc turns a function with the right signature into a get fee usd handler
type GetFeeUsdHandlerFunc func(GetFeeUsdParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFeeUsdHandlerFunc) Handle(params GetFeeUsdParams) middleware.Responder {
	return fn(params)
}

// GetFeeUsdHandler interface for that can handle valid get fee usd params
type GetFeeUsdHandler interface {
	Handle(GetFeeUsdParams) middleware.Responder
}

// NewGetFeeUsd creates a new http.Handler for the get fee usd operation
func NewGetFeeUsd(ctx *middleware.Context, handler GetFeeUsdHandler) *GetFeeUsd {
	return &GetFeeUsd{Context: ctx, Handler: handler}
}

/* GetFeeUsd swagger:route GET /fee/usd getFeeUsd

List all fees in USD

*/
type GetFeeUsd struct {
	Context *middleware.Context
	Handler GetFeeUsdHandler
}

func (o *GetFeeUsd) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetFeeUsdParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}