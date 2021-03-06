// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/blockpane/fio-web-utils/stats-api/models"
)

// GetSuggestSetfeemultOKCode is the HTTP code returned for type GetSuggestSetfeemultOK
const GetSuggestSetfeemultOKCode int = 200

/*GetSuggestSetfeemultOK The suggested fee multiplier

swagger:response getSuggestSetfeemultOK
*/
type GetSuggestSetfeemultOK struct {

	/*
	  In: Body
	*/
	Payload *GetSuggestSetfeemultOKBody `json:"body,omitempty"`
}

// NewGetSuggestSetfeemultOK creates GetSuggestSetfeemultOK with default headers values
func NewGetSuggestSetfeemultOK() *GetSuggestSetfeemultOK {

	return &GetSuggestSetfeemultOK{}
}

// WithPayload adds the payload to the get suggest setfeemult o k response
func (o *GetSuggestSetfeemultOK) WithPayload(payload *GetSuggestSetfeemultOKBody) *GetSuggestSetfeemultOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get suggest setfeemult o k response
func (o *GetSuggestSetfeemultOK) SetPayload(payload *GetSuggestSetfeemultOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSuggestSetfeemultOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSuggestSetfeemultServiceUnavailableCode is the HTTP code returned for type GetSuggestSetfeemultServiceUnavailable
const GetSuggestSetfeemultServiceUnavailableCode int = 503

/*GetSuggestSetfeemultServiceUnavailable Data is stale, has not been updated for several minutes

swagger:response getSuggestSetfeemultServiceUnavailable
*/
type GetSuggestSetfeemultServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSuggestSetfeemultServiceUnavailable creates GetSuggestSetfeemultServiceUnavailable with default headers values
func NewGetSuggestSetfeemultServiceUnavailable() *GetSuggestSetfeemultServiceUnavailable {

	return &GetSuggestSetfeemultServiceUnavailable{}
}

// WithPayload adds the payload to the get suggest setfeemult service unavailable response
func (o *GetSuggestSetfeemultServiceUnavailable) WithPayload(payload *models.Error) *GetSuggestSetfeemultServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get suggest setfeemult service unavailable response
func (o *GetSuggestSetfeemultServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSuggestSetfeemultServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
