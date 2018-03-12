// Code generated by go-swagger; DO NOT EDIT.

package pilot_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSubtaskStatusParams creates a new GetSubtaskStatusParams object
// with the default values initialized.
func NewGetSubtaskStatusParams() *GetSubtaskStatusParams {
	var ()
	return &GetSubtaskStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSubtaskStatusParamsWithTimeout creates a new GetSubtaskStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSubtaskStatusParamsWithTimeout(timeout time.Duration) *GetSubtaskStatusParams {
	var ()
	return &GetSubtaskStatusParams{

		timeout: timeout,
	}
}

// NewGetSubtaskStatusParamsWithContext creates a new GetSubtaskStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSubtaskStatusParamsWithContext(ctx context.Context) *GetSubtaskStatusParams {
	var ()
	return &GetSubtaskStatusParams{

		Context: ctx,
	}
}

// NewGetSubtaskStatusParamsWithHTTPClient creates a new GetSubtaskStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSubtaskStatusParamsWithHTTPClient(client *http.Client) *GetSubtaskStatusParams {
	var ()
	return &GetSubtaskStatusParams{
		HTTPClient: client,
	}
}

/*GetSubtaskStatusParams contains all the parameters to send to the API endpoint
for the get subtask status operation typically these are written to a http.Request
*/
type GetSubtaskStatusParams struct {

	/*SubtaskID*/
	SubtaskID []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get subtask status params
func (o *GetSubtaskStatusParams) WithTimeout(timeout time.Duration) *GetSubtaskStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get subtask status params
func (o *GetSubtaskStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get subtask status params
func (o *GetSubtaskStatusParams) WithContext(ctx context.Context) *GetSubtaskStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get subtask status params
func (o *GetSubtaskStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get subtask status params
func (o *GetSubtaskStatusParams) WithHTTPClient(client *http.Client) *GetSubtaskStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get subtask status params
func (o *GetSubtaskStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSubtaskID adds the subtaskID to the get subtask status params
func (o *GetSubtaskStatusParams) WithSubtaskID(subtaskID []string) *GetSubtaskStatusParams {
	o.SetSubtaskID(subtaskID)
	return o
}

// SetSubtaskID adds the subtaskId to the get subtask status params
func (o *GetSubtaskStatusParams) SetSubtaskID(subtaskID []string) {
	o.SubtaskID = subtaskID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSubtaskStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesSubtaskID := o.SubtaskID

	joinedSubtaskID := swag.JoinByFormat(valuesSubtaskID, "")
	// query array param subtask_id
	if err := r.SetQueryParam("subtask_id", joinedSubtaskID...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
