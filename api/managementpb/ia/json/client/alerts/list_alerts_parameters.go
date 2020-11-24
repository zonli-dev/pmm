// Code generated by go-swagger; DO NOT EDIT.

package alerts

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

// NewListAlertsParams creates a new ListAlertsParams object
// with the default values initialized.
func NewListAlertsParams() *ListAlertsParams {
	var ()
	return &ListAlertsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAlertsParamsWithTimeout creates a new ListAlertsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAlertsParamsWithTimeout(timeout time.Duration) *ListAlertsParams {
	var ()
	return &ListAlertsParams{

		timeout: timeout,
	}
}

// NewListAlertsParamsWithContext creates a new ListAlertsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAlertsParamsWithContext(ctx context.Context) *ListAlertsParams {
	var ()
	return &ListAlertsParams{

		Context: ctx,
	}
}

// NewListAlertsParamsWithHTTPClient creates a new ListAlertsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAlertsParamsWithHTTPClient(client *http.Client) *ListAlertsParams {
	var ()
	return &ListAlertsParams{
		HTTPClient: client,
	}
}

/*ListAlertsParams contains all the parameters to send to the API endpoint
for the list alerts operation typically these are written to a http.Request
*/
type ListAlertsParams struct {

	/*Body*/
	Body interface{}

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list alerts params
func (o *ListAlertsParams) WithTimeout(timeout time.Duration) *ListAlertsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list alerts params
func (o *ListAlertsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list alerts params
func (o *ListAlertsParams) WithContext(ctx context.Context) *ListAlertsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list alerts params
func (o *ListAlertsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list alerts params
func (o *ListAlertsParams) WithHTTPClient(client *http.Client) *ListAlertsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list alerts params
func (o *ListAlertsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list alerts params
func (o *ListAlertsParams) WithBody(body interface{}) *ListAlertsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list alerts params
func (o *ListAlertsParams) SetBody(body interface{}) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListAlertsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
