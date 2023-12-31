// Code generated by go-swagger; DO NOT EDIT.

package notes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"github.com/almalii/swagger-contracts/models"
)

// NewPatchNotesIDParams creates a new PatchNotesIDParams object
//
// There are no default values defined in the spec.
func NewPatchNotesIDParams() PatchNotesIDParams {

	return PatchNotesIDParams{}
}

// PatchNotesIDParams contains all the bound params for the patch notes ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters PatchNotesID
type PatchNotesIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Note ID
	  Required: true
	  In: path
	*/
	ID string
	/*Note info
	  Required: true
	  In: body
	*/
	Note *models.ControllerUpdateNoteRequest
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPatchNotesIDParams() beforehand.
func (o *PatchNotesIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ControllerUpdateNoteRequest
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("note", "body", ""))
			} else {
				res = append(res, errors.NewParseError("note", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Note = &body
			}
		}
	} else {
		res = append(res, errors.Required("note", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *PatchNotesIDParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ID = raw

	return nil
}
