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
	"github.com/go-openapi/validate"
)

// NewGetFeeVotesProducerParams creates a new GetFeeVotesProducerParams object
// with the default values initialized.
func NewGetFeeVotesProducerParams() GetFeeVotesProducerParams {

	var (
		// initialize parameters with default values

		applyMultiplierDefault = bool(true)
		formatDefault          = string("suf")
	)

	return GetFeeVotesProducerParams{
		ApplyMultiplier: &applyMultiplierDefault,

		Format: &formatDefault,
	}
}

// GetFeeVotesProducerParams contains all the bound params for the get fee votes producer operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetFeeVotesProducer
type GetFeeVotesProducerParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Should the result apply the multiplier? Setting to 'true' is useful for calculating the actual cost that a producer has voted, most useful with 'format=usd'. Using 'false' will provide their base fio.fee::setfeevotes values, and only makes sense if 'format=suf'.
	  In: query
	  Default: true
	*/
	ApplyMultiplier *bool
	/*How to format the values. Options are 'suf' which is the smallest unit of FIO, 'float' which is whole FIO with up to 9 digits, and 'usd' which is the current price.
	  In: query
	  Default: "suf"
	*/
	Format *string
	/*The producer's account
	  Required: true
	  In: path
	*/
	Producer string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetFeeVotesProducerParams() beforehand.
func (o *GetFeeVotesProducerParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qApplyMultiplier, qhkApplyMultiplier, _ := qs.GetOK("apply_multiplier")
	if err := o.bindApplyMultiplier(qApplyMultiplier, qhkApplyMultiplier, route.Formats); err != nil {
		res = append(res, err)
	}

	qFormat, qhkFormat, _ := qs.GetOK("format")
	if err := o.bindFormat(qFormat, qhkFormat, route.Formats); err != nil {
		res = append(res, err)
	}

	rProducer, rhkProducer, _ := route.Params.GetOK("producer")
	if err := o.bindProducer(rProducer, rhkProducer, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindApplyMultiplier binds and validates parameter ApplyMultiplier from query.
func (o *GetFeeVotesProducerParams) bindApplyMultiplier(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetFeeVotesProducerParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("apply_multiplier", "query", "bool", raw)
	}
	o.ApplyMultiplier = &value

	return nil
}

// bindFormat binds and validates parameter Format from query.
func (o *GetFeeVotesProducerParams) bindFormat(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetFeeVotesProducerParams()
		return nil
	}
	o.Format = &raw

	if err := o.validateFormat(formats); err != nil {
		return err
	}

	return nil
}

// validateFormat carries on validations for parameter Format
func (o *GetFeeVotesProducerParams) validateFormat(formats strfmt.Registry) error {

	if err := validate.EnumCase("format", "query", *o.Format, []interface{}{"suf", "float", "usd"}, true); err != nil {
		return err
	}

	return nil
}

// bindProducer binds and validates parameter Producer from path.
func (o *GetFeeVotesProducerParams) bindProducer(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Producer = raw

	return nil
}
