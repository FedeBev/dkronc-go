// Code generated by go-swagger; DO NOT EDIT.

package jobs

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
	"github.com/go-openapi/swag"

	"github.com/FedeBev/dkronc-go/pkg/models"
)

// NewCreateOrUpdateJobParams creates a new CreateOrUpdateJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateOrUpdateJobParams() *CreateOrUpdateJobParams {
	return &CreateOrUpdateJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOrUpdateJobParamsWithTimeout creates a new CreateOrUpdateJobParams object
// with the ability to set a timeout on a request.
func NewCreateOrUpdateJobParamsWithTimeout(timeout time.Duration) *CreateOrUpdateJobParams {
	return &CreateOrUpdateJobParams{
		timeout: timeout,
	}
}

// NewCreateOrUpdateJobParamsWithContext creates a new CreateOrUpdateJobParams object
// with the ability to set a context for a request.
func NewCreateOrUpdateJobParamsWithContext(ctx context.Context) *CreateOrUpdateJobParams {
	return &CreateOrUpdateJobParams{
		Context: ctx,
	}
}

// NewCreateOrUpdateJobParamsWithHTTPClient creates a new CreateOrUpdateJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateOrUpdateJobParamsWithHTTPClient(client *http.Client) *CreateOrUpdateJobParams {
	return &CreateOrUpdateJobParams{
		HTTPClient: client,
	}
}

/* CreateOrUpdateJobParams contains all the parameters to send to the API endpoint
   for the create or update job operation.

   Typically these are written to a http.Request.
*/
type CreateOrUpdateJobParams struct {

	/* Body.

	   Updated job object
	*/
	Body *models.Job

	/* Runoncreate.

	   If present, regardless of any value, causes the job to be run immediately after being succesfully created or updated.
	*/
	Runoncreate bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create or update job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrUpdateJobParams) WithDefaults() *CreateOrUpdateJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create or update job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrUpdateJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create or update job params
func (o *CreateOrUpdateJobParams) WithTimeout(timeout time.Duration) *CreateOrUpdateJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create or update job params
func (o *CreateOrUpdateJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create or update job params
func (o *CreateOrUpdateJobParams) WithContext(ctx context.Context) *CreateOrUpdateJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create or update job params
func (o *CreateOrUpdateJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create or update job params
func (o *CreateOrUpdateJobParams) WithHTTPClient(client *http.Client) *CreateOrUpdateJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create or update job params
func (o *CreateOrUpdateJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create or update job params
func (o *CreateOrUpdateJobParams) WithBody(body *models.Job) *CreateOrUpdateJobParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create or update job params
func (o *CreateOrUpdateJobParams) SetBody(body *models.Job) {
	o.Body = body
}

// WithRunoncreate adds the runoncreate to the create or update job params
func (o *CreateOrUpdateJobParams) WithRunoncreate(runoncreate bool) *CreateOrUpdateJobParams {
	o.SetRunoncreate(runoncreate)
	return o
}

// SetRunoncreate adds the runoncreate to the create or update job params
func (o *CreateOrUpdateJobParams) SetRunoncreate(runoncreate bool) {
	o.Runoncreate = runoncreate
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOrUpdateJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// query param runoncreate
	qrRunoncreate := o.Runoncreate
	qRunoncreate := swag.FormatBool(qrRunoncreate)

	if err := r.SetQueryParam("runoncreate", qRunoncreate); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
