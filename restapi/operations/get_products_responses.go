// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"integrator-docs/models"
)

// GetProductsOKCode is the HTTP code returned for type GetProductsOK
const GetProductsOKCode int = 200

/*
GetProductsOK lists products

swagger:response getProductsOK
*/
type GetProductsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Product `json:"body,omitempty"`
}

// NewGetProductsOK creates GetProductsOK with default headers values
func NewGetProductsOK() *GetProductsOK {

	return &GetProductsOK{}
}

// WithPayload adds the payload to the get products o k response
func (o *GetProductsOK) WithPayload(payload []*models.Product) *GetProductsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get products o k response
func (o *GetProductsOK) SetPayload(payload []*models.Product) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProductsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Product, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
