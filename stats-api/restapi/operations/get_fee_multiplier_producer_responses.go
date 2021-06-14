// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/blockpane/fio-web-utils/stats-api/models"
)

// GetFeeMultiplierProducerOKCode is the HTTP code returned for type GetFeeMultiplierProducerOK
const GetFeeMultiplierProducerOKCode int = 200

/*GetFeeMultiplierProducerOK The currently used fee multiplier which is applied to the fee vote to determine price

swagger:response getFeeMultiplierProducerOK
*/
type GetFeeMultiplierProducerOK struct {

	/*
	  In: Body
	*/
	Payload *GetFeeMultiplierProducerOKBody `json:"body,omitempty"`
}

// NewGetFeeMultiplierProducerOK creates GetFeeMultiplierProducerOK with default headers values
func NewGetFeeMultiplierProducerOK() *GetFeeMultiplierProducerOK {

	return &GetFeeMultiplierProducerOK{}
}

// WithPayload adds the payload to the get fee multiplier producer o k response
func (o *GetFeeMultiplierProducerOK) WithPayload(payload *GetFeeMultiplierProducerOKBody) *GetFeeMultiplierProducerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get fee multiplier producer o k response
func (o *GetFeeMultiplierProducerOK) SetPayload(payload *GetFeeMultiplierProducerOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFeeMultiplierProducerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFeeMultiplierProducerBadRequestCode is the HTTP code returned for type GetFeeMultiplierProducerBadRequest
const GetFeeMultiplierProducerBadRequestCode int = 400

/*GetFeeMultiplierProducerBadRequest Invalid account format, should be a 12 character string

swagger:response getFeeMultiplierProducerBadRequest
*/
type GetFeeMultiplierProducerBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetFeeMultiplierProducerBadRequest creates GetFeeMultiplierProducerBadRequest with default headers values
func NewGetFeeMultiplierProducerBadRequest() *GetFeeMultiplierProducerBadRequest {

	return &GetFeeMultiplierProducerBadRequest{}
}

// WithPayload adds the payload to the get fee multiplier producer bad request response
func (o *GetFeeMultiplierProducerBadRequest) WithPayload(payload *models.Error) *GetFeeMultiplierProducerBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get fee multiplier producer bad request response
func (o *GetFeeMultiplierProducerBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFeeMultiplierProducerBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFeeMultiplierProducerNotFoundCode is the HTTP code returned for type GetFeeMultiplierProducerNotFound
const GetFeeMultiplierProducerNotFoundCode int = 404

/*GetFeeMultiplierProducerNotFound Did not find a matching producer

swagger:response getFeeMultiplierProducerNotFound
*/
type GetFeeMultiplierProducerNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetFeeMultiplierProducerNotFound creates GetFeeMultiplierProducerNotFound with default headers values
func NewGetFeeMultiplierProducerNotFound() *GetFeeMultiplierProducerNotFound {

	return &GetFeeMultiplierProducerNotFound{}
}

// WithPayload adds the payload to the get fee multiplier producer not found response
func (o *GetFeeMultiplierProducerNotFound) WithPayload(payload *models.Error) *GetFeeMultiplierProducerNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get fee multiplier producer not found response
func (o *GetFeeMultiplierProducerNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFeeMultiplierProducerNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFeeMultiplierProducerServiceUnavailableCode is the HTTP code returned for type GetFeeMultiplierProducerServiceUnavailable
const GetFeeMultiplierProducerServiceUnavailableCode int = 503

/*GetFeeMultiplierProducerServiceUnavailable Data is stale, has not been updated for several minutes

swagger:response getFeeMultiplierProducerServiceUnavailable
*/
type GetFeeMultiplierProducerServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetFeeMultiplierProducerServiceUnavailable creates GetFeeMultiplierProducerServiceUnavailable with default headers values
func NewGetFeeMultiplierProducerServiceUnavailable() *GetFeeMultiplierProducerServiceUnavailable {

	return &GetFeeMultiplierProducerServiceUnavailable{}
}

// WithPayload adds the payload to the get fee multiplier producer service unavailable response
func (o *GetFeeMultiplierProducerServiceUnavailable) WithPayload(payload *models.Error) *GetFeeMultiplierProducerServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get fee multiplier producer service unavailable response
func (o *GetFeeMultiplierProducerServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFeeMultiplierProducerServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
