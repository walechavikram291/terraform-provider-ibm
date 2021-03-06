// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// NewPcloudPvminstancesVolumesPutParams creates a new PcloudPvminstancesVolumesPutParams object
// with the default values initialized.
func NewPcloudPvminstancesVolumesPutParams() *PcloudPvminstancesVolumesPutParams {
	var ()
	return &PcloudPvminstancesVolumesPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudPvminstancesVolumesPutParamsWithTimeout creates a new PcloudPvminstancesVolumesPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPcloudPvminstancesVolumesPutParamsWithTimeout(timeout time.Duration) *PcloudPvminstancesVolumesPutParams {
	var ()
	return &PcloudPvminstancesVolumesPutParams{

		timeout: timeout,
	}
}

// NewPcloudPvminstancesVolumesPutParamsWithContext creates a new PcloudPvminstancesVolumesPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewPcloudPvminstancesVolumesPutParamsWithContext(ctx context.Context) *PcloudPvminstancesVolumesPutParams {
	var ()
	return &PcloudPvminstancesVolumesPutParams{

		Context: ctx,
	}
}

// NewPcloudPvminstancesVolumesPutParamsWithHTTPClient creates a new PcloudPvminstancesVolumesPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPcloudPvminstancesVolumesPutParamsWithHTTPClient(client *http.Client) *PcloudPvminstancesVolumesPutParams {
	var ()
	return &PcloudPvminstancesVolumesPutParams{
		HTTPClient: client,
	}
}

/*PcloudPvminstancesVolumesPutParams contains all the parameters to send to the API endpoint
for the pcloud pvminstances volumes put operation typically these are written to a http.Request
*/
type PcloudPvminstancesVolumesPutParams struct {

	/*Body
	  Parameters to update a volume attached to a PVMInstance

	*/
	Body *models.PVMInstanceVolumeUpdate
	/*CloudInstanceID
	  Cloud Instance ID of a PCloud Instance

	*/
	CloudInstanceID string
	/*PvmInstanceID
	  PCloud PVM Instance ID

	*/
	PvmInstanceID string
	/*VolumeID
	  Volume ID

	*/
	VolumeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pcloud pvminstances volumes put params
func (o *PcloudPvminstancesVolumesPutParams) WithTimeout(timeout time.Duration) *PcloudPvminstancesVolumesPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud pvminstances volumes put params
func (o *PcloudPvminstancesVolumesPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud pvminstances volumes put params
func (o *PcloudPvminstancesVolumesPutParams) WithContext(ctx context.Context) *PcloudPvminstancesVolumesPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud pvminstances volumes put params
func (o *PcloudPvminstancesVolumesPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud pvminstances volumes put params
func (o *PcloudPvminstancesVolumesPutParams) WithHTTPClient(client *http.Client) *PcloudPvminstancesVolumesPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud pvminstances volumes put params
func (o *PcloudPvminstancesVolumesPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud pvminstances volumes put params
func (o *PcloudPvminstancesVolumesPutParams) WithBody(body *models.PVMInstanceVolumeUpdate) *PcloudPvminstancesVolumesPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud pvminstances volumes put params
func (o *PcloudPvminstancesVolumesPutParams) SetBody(body *models.PVMInstanceVolumeUpdate) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud pvminstances volumes put params
func (o *PcloudPvminstancesVolumesPutParams) WithCloudInstanceID(cloudInstanceID string) *PcloudPvminstancesVolumesPutParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud pvminstances volumes put params
func (o *PcloudPvminstancesVolumesPutParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithPvmInstanceID adds the pvmInstanceID to the pcloud pvminstances volumes put params
func (o *PcloudPvminstancesVolumesPutParams) WithPvmInstanceID(pvmInstanceID string) *PcloudPvminstancesVolumesPutParams {
	o.SetPvmInstanceID(pvmInstanceID)
	return o
}

// SetPvmInstanceID adds the pvmInstanceId to the pcloud pvminstances volumes put params
func (o *PcloudPvminstancesVolumesPutParams) SetPvmInstanceID(pvmInstanceID string) {
	o.PvmInstanceID = pvmInstanceID
}

// WithVolumeID adds the volumeID to the pcloud pvminstances volumes put params
func (o *PcloudPvminstancesVolumesPutParams) WithVolumeID(volumeID string) *PcloudPvminstancesVolumesPutParams {
	o.SetVolumeID(volumeID)
	return o
}

// SetVolumeID adds the volumeId to the pcloud pvminstances volumes put params
func (o *PcloudPvminstancesVolumesPutParams) SetVolumeID(volumeID string) {
	o.VolumeID = volumeID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudPvminstancesVolumesPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param pvm_instance_id
	if err := r.SetPathParam("pvm_instance_id", o.PvmInstanceID); err != nil {
		return err
	}

	// path param volume_id
	if err := r.SetPathParam("volume_id", o.VolumeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
