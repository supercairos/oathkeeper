// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewMakeGenericDecisionParams creates a new MakeGenericDecisionParams object
// with the default values initialized.
func NewMakeGenericDecisionParams() *MakeGenericDecisionParams {

	return &MakeGenericDecisionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMakeGenericDecisionParamsWithTimeout creates a new MakeGenericDecisionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMakeGenericDecisionParamsWithTimeout(timeout time.Duration) *MakeGenericDecisionParams {

	return &MakeGenericDecisionParams{

		timeout: timeout,
	}
}

// NewMakeGenericDecisionParamsWithContext creates a new MakeGenericDecisionParams object
// with the default values initialized, and the ability to set a context for a request
func NewMakeGenericDecisionParamsWithContext(ctx context.Context) *MakeGenericDecisionParams {

	return &MakeGenericDecisionParams{

		Context: ctx,
	}
}

// NewMakeGenericDecisionParamsWithHTTPClient creates a new MakeGenericDecisionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMakeGenericDecisionParamsWithHTTPClient(client *http.Client) *MakeGenericDecisionParams {

	return &MakeGenericDecisionParams{
		HTTPClient: client,
	}
}

/*MakeGenericDecisionParams contains all the parameters to send to the API endpoint
for the make generic decision operation typically these are written to a http.Request
*/
type MakeGenericDecisionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the decisions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MakeGenericDecisionParams) WithDefaults() *MakeGenericDecisionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the decisions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MakeGenericDecisionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the make generic decision params
func (o *MakeGenericDecisionParams) WithTimeout(timeout time.Duration) *MakeGenericDecisionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the make generic decision params
func (o *MakeGenericDecisionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the make generic decision params
func (o *MakeGenericDecisionParams) WithContext(ctx context.Context) *MakeGenericDecisionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the make generic decision params
func (o *MakeGenericDecisionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the make generic decision params
func (o *MakeGenericDecisionParams) WithHTTPClient(client *http.Client) *MakeGenericDecisionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the make generic decision params
func (o *MakeGenericDecisionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *MakeGenericDecisionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
