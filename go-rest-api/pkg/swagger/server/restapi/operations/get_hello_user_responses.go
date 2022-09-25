// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetHelloUserOKCode is the HTTP code returned for type GetHelloUserOK
const GetHelloUserOKCode int = 200

/*
GetHelloUserOK Returns greeting

swagger:response getHelloUserOK
*/
type GetHelloUserOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetHelloUserOK creates GetHelloUserOK with default headers values
func NewGetHelloUserOK() *GetHelloUserOK {

	return &GetHelloUserOK{}
}

// WithPayload adds the payload to the get hello user o k response
func (o *GetHelloUserOK) WithPayload(payload string) *GetHelloUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get hello user o k response
func (o *GetHelloUserOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHelloUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetHelloUserBadRequestCode is the HTTP code returned for type GetHelloUserBadRequest
const GetHelloUserBadRequestCode int = 400

/*
GetHelloUserBadRequest invalid characters in "user" provided

swagger:response getHelloUserBadRequest
*/
type GetHelloUserBadRequest struct {
}

// NewGetHelloUserBadRequest creates GetHelloUserBadRequest with default headers values
func NewGetHelloUserBadRequest() *GetHelloUserBadRequest {

	return &GetHelloUserBadRequest{}
}

// WriteResponse to the client
func (o *GetHelloUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
