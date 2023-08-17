// Code generated by go-swagger; DO NOT EDIT.

package notes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetNotesOKCode is the HTTP code returned for type GetNotesOK
const GetNotesOKCode int = 200

/*
GetNotesOK OK

swagger:response getNotesOK
*/
type GetNotesOK struct {
}

// NewGetNotesOK creates GetNotesOK with default headers values
func NewGetNotesOK() *GetNotesOK {

	return &GetNotesOK{}
}

// WriteResponse to the client
func (o *GetNotesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// GetNotesBadRequestCode is the HTTP code returned for type GetNotesBadRequest
const GetNotesBadRequestCode int = 400

/*
GetNotesBadRequest Bad Request

swagger:response getNotesBadRequest
*/
type GetNotesBadRequest struct {
}

// NewGetNotesBadRequest creates GetNotesBadRequest with default headers values
func NewGetNotesBadRequest() *GetNotesBadRequest {

	return &GetNotesBadRequest{}
}

// WriteResponse to the client
func (o *GetNotesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetNotesInternalServerErrorCode is the HTTP code returned for type GetNotesInternalServerError
const GetNotesInternalServerErrorCode int = 500

/*
GetNotesInternalServerError Internal Server Error

swagger:response getNotesInternalServerError
*/
type GetNotesInternalServerError struct {
}

// NewGetNotesInternalServerError creates GetNotesInternalServerError with default headers values
func NewGetNotesInternalServerError() *GetNotesInternalServerError {

	return &GetNotesInternalServerError{}
}

// WriteResponse to the client
func (o *GetNotesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
