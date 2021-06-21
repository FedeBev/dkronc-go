// Code generated by go-swagger; DO NOT EDIT.

package members

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

// NewGetMemberParams creates a new GetMemberParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMemberParams() *GetMemberParams {
	return &GetMemberParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMemberParamsWithTimeout creates a new GetMemberParams object
// with the ability to set a timeout on a request.
func NewGetMemberParamsWithTimeout(timeout time.Duration) *GetMemberParams {
	return &GetMemberParams{
		timeout: timeout,
	}
}

// NewGetMemberParamsWithContext creates a new GetMemberParams object
// with the ability to set a context for a request.
func NewGetMemberParamsWithContext(ctx context.Context) *GetMemberParams {
	return &GetMemberParams{
		Context: ctx,
	}
}

// NewGetMemberParamsWithHTTPClient creates a new GetMemberParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMemberParamsWithHTTPClient(client *http.Client) *GetMemberParams {
	return &GetMemberParams{
		HTTPClient: client,
	}
}

/* GetMemberParams contains all the parameters to send to the API endpoint
   for the get member operation.

   Typically these are written to a http.Request.
*/
type GetMemberParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMemberParams) WithDefaults() *GetMemberParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMemberParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get member params
func (o *GetMemberParams) WithTimeout(timeout time.Duration) *GetMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get member params
func (o *GetMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get member params
func (o *GetMemberParams) WithContext(ctx context.Context) *GetMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get member params
func (o *GetMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get member params
func (o *GetMemberParams) WithHTTPClient(client *http.Client) *GetMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get member params
func (o *GetMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
