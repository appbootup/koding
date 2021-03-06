package j_stack_template

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

	"koding/remoteapi/models"
)

// NewPostRemoteAPIJStackTemplateBuildIDParams creates a new PostRemoteAPIJStackTemplateBuildIDParams object
// with the default values initialized.
func NewPostRemoteAPIJStackTemplateBuildIDParams() *PostRemoteAPIJStackTemplateBuildIDParams {
	var ()
	return &PostRemoteAPIJStackTemplateBuildIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJStackTemplateBuildIDParamsWithTimeout creates a new PostRemoteAPIJStackTemplateBuildIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJStackTemplateBuildIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJStackTemplateBuildIDParams {
	var ()
	return &PostRemoteAPIJStackTemplateBuildIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJStackTemplateBuildIDParamsWithContext creates a new PostRemoteAPIJStackTemplateBuildIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJStackTemplateBuildIDParamsWithContext(ctx context.Context) *PostRemoteAPIJStackTemplateBuildIDParams {
	var ()
	return &PostRemoteAPIJStackTemplateBuildIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJStackTemplateBuildIDParams contains all the parameters to send to the API endpoint
for the post remote API j stack template build ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJStackTemplateBuildIDParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector
	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j stack template build ID params
func (o *PostRemoteAPIJStackTemplateBuildIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJStackTemplateBuildIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j stack template build ID params
func (o *PostRemoteAPIJStackTemplateBuildIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j stack template build ID params
func (o *PostRemoteAPIJStackTemplateBuildIDParams) WithContext(ctx context.Context) *PostRemoteAPIJStackTemplateBuildIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j stack template build ID params
func (o *PostRemoteAPIJStackTemplateBuildIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j stack template build ID params
func (o *PostRemoteAPIJStackTemplateBuildIDParams) WithBody(body models.DefaultSelector) *PostRemoteAPIJStackTemplateBuildIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j stack template build ID params
func (o *PostRemoteAPIJStackTemplateBuildIDParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the post remote API j stack template build ID params
func (o *PostRemoteAPIJStackTemplateBuildIDParams) WithID(id string) *PostRemoteAPIJStackTemplateBuildIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j stack template build ID params
func (o *PostRemoteAPIJStackTemplateBuildIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJStackTemplateBuildIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
