// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPingParams creates a new PingParams object
// with the default values initialized.
func NewPingParams() *PingParams {

	return &PingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPingParamsWithTimeout creates a new PingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPingParamsWithTimeout(timeout time.Duration) *PingParams {

	return &PingParams{

		timeout: timeout,
	}
}

// NewPingParamsWithContext creates a new PingParams object
// with the default values initialized, and the ability to set a context for a request
func NewPingParamsWithContext(ctx context.Context) *PingParams {

	return &PingParams{

		Context: ctx,
	}
}

// NewPingParamsWithHTTPClient creates a new PingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPingParamsWithHTTPClient(client *http.Client) *PingParams {

	return &PingParams{
		HTTPClient: client,
	}
}

/*PingParams contains all the parameters to send to the API endpoint
for the ping operation typically these are written to a http.Request
*/
type PingParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ping params
func (o *PingParams) WithTimeout(timeout time.Duration) *PingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ping params
func (o *PingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ping params
func (o *PingParams) WithContext(ctx context.Context) *PingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ping params
func (o *PingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ping params
func (o *PingParams) WithHTTPClient(client *http.Client) *PingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ping params
func (o *PingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
