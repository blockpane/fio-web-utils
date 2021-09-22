// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetSuggestSetfeemultParams creates a new GetSuggestSetfeemultParams object
// with the default values initialized.
func NewGetSuggestSetfeemultParams() GetSuggestSetfeemultParams {

	var (
		// initialize parameters with default values

		targetDefault = float64(2)
	)

	return GetSuggestSetfeemultParams{
		Target: &targetDefault,
	}
}

// GetSuggestSetfeemultParams contains all the bound params for the get suggest setfeemult operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetSuggestSetfeemult
type GetSuggestSetfeemultParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	  Default: 2
	*/
	Target *float64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetSuggestSetfeemultParams() beforehand.
func (o *GetSuggestSetfeemultParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qTarget, qhkTarget, _ := qs.GetOK("target")
	if err := o.bindTarget(qTarget, qhkTarget, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindTarget binds and validates parameter Target from query.
func (o *GetSuggestSetfeemultParams) bindTarget(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetSuggestSetfeemultParams()
		return nil
	}

	value, err := swag.ConvertFloat64(raw)
	if err != nil {
		return errors.InvalidType("target", "query", "float64", raw)
	}
	o.Target = &value

	return nil
}