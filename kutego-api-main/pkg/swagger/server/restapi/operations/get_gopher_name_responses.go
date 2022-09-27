// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetCatNameOKCode is the HTTP code returned for type GetCatNameOK
const GetCatNameOKCode int = 200

/*GetCatNameOK Returns the cat.

swagger:response getCatNameOK
*/
type GetCatNameOK struct {

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewGetCatNameOK creates GetCatNameOK with default headers values
func NewGetCatNameOK() *GetCatNameOK {

	return &GetCatNameOK{}
}

// WithPayload adds the payload to the get cat name o k response
func (o *GetCatNameOK) WithPayload(payload io.ReadCloser) *GetCatNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get cat name o k response
func (o *GetCatNameOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCatNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetCatNameBadRequestCode is the HTTP code returned for type GetCatNameBadRequest
const GetCatNameBadRequestCode int = 400

/*GetCatNameBadRequest Invalid characters in "name" were provided.

swagger:response getCatNameBadRequest
*/
type GetCatNameBadRequest struct {
}

// NewGetCatNameBadRequest creates GetCatNameBadRequest with default headers values
func NewGetCatNameBadRequest() *GetCatNameBadRequest {

	return &GetCatNameBadRequest{}
}

// WriteResponse to the client
func (o *GetCatNameBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}