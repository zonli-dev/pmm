// Code generated by go-swagger; DO NOT EDIT.

package metrics_names

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new metrics names API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for metrics names API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetMetricsNames gets metrics names gets map of metrics names
*/
func (a *Client) GetMetricsNames(params *GetMetricsNamesParams) (*GetMetricsNamesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMetricsNamesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetMetricsNames",
		Method:             "POST",
		PathPattern:        "/v1/qan/GetMetricsNames",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetMetricsNamesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMetricsNamesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
