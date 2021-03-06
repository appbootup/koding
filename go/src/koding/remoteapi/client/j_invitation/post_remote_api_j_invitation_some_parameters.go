package j_invitation

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

// NewPostRemoteAPIJInvitationSomeParams creates a new PostRemoteAPIJInvitationSomeParams object
// with the default values initialized.
func NewPostRemoteAPIJInvitationSomeParams() *PostRemoteAPIJInvitationSomeParams {
	var ()
	return &PostRemoteAPIJInvitationSomeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJInvitationSomeParamsWithTimeout creates a new PostRemoteAPIJInvitationSomeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJInvitationSomeParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJInvitationSomeParams {
	var ()
	return &PostRemoteAPIJInvitationSomeParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJInvitationSomeParamsWithContext creates a new PostRemoteAPIJInvitationSomeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJInvitationSomeParamsWithContext(ctx context.Context) *PostRemoteAPIJInvitationSomeParams {
	var ()
	return &PostRemoteAPIJInvitationSomeParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJInvitationSomeParams contains all the parameters to send to the API endpoint
for the post remote API j invitation some operation typically these are written to a http.Request
*/
type PostRemoteAPIJInvitationSomeParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j invitation some params
func (o *PostRemoteAPIJInvitationSomeParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJInvitationSomeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j invitation some params
func (o *PostRemoteAPIJInvitationSomeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j invitation some params
func (o *PostRemoteAPIJInvitationSomeParams) WithContext(ctx context.Context) *PostRemoteAPIJInvitationSomeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j invitation some params
func (o *PostRemoteAPIJInvitationSomeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j invitation some params
func (o *PostRemoteAPIJInvitationSomeParams) WithBody(body models.DefaultSelector) *PostRemoteAPIJInvitationSomeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j invitation some params
func (o *PostRemoteAPIJInvitationSomeParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJInvitationSomeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
