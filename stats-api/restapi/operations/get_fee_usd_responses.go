// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/blockpane/fio-web-utils/stats-api/models"
)

// GetFeeUsdOKCode is the HTTP code returned for type GetFeeUsdOK
const GetFeeUsdOKCode int = 200

/*GetFeeUsdOK An array of prices for each action in USD

swagger:response getFeeUsdOK
*/
type GetFeeUsdOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Price `json:"body,omitempty"`
}

// NewGetFeeUsdOK creates GetFeeUsdOK with default headers values
func NewGetFeeUsdOK() *GetFeeUsdOK {

	return &GetFeeUsdOK{}
}

// WithPayload adds the payload to the get fee usd o k response
func (o *GetFeeUsdOK) WithPayload(payload []*models.Price) *GetFeeUsdOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get fee usd o k response
func (o *GetFeeUsdOK) SetPayload(payload []*models.Price) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFeeUsdOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Price, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetFeeUsdServiceUnavailableCode is the HTTP code returned for type GetFeeUsdServiceUnavailable
const GetFeeUsdServiceUnavailableCode int = 503

/*GetFeeUsdServiceUnavailable Data is stale, has not been updated for several minutes

swagger:response getFeeUsdServiceUnavailable
*/
type GetFeeUsdServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetFeeUsdServiceUnavailable creates GetFeeUsdServiceUnavailable with default headers values
func NewGetFeeUsdServiceUnavailable() *GetFeeUsdServiceUnavailable {

	return &GetFeeUsdServiceUnavailable{}
}

// WithPayload adds the payload to the get fee usd service unavailable response
func (o *GetFeeUsdServiceUnavailable) WithPayload(payload *models.Error) *GetFeeUsdServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get fee usd service unavailable response
func (o *GetFeeUsdServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFeeUsdServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}