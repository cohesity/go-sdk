// Code generated by go-swagger; DO NOT EDIT.

package platform

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
)

// NewGetClusterVlansParams creates a new GetClusterVlansParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClusterVlansParams() *GetClusterVlansParams {
	return &GetClusterVlansParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterVlansParamsWithTimeout creates a new GetClusterVlansParams object
// with the ability to set a timeout on a request.
func NewGetClusterVlansParamsWithTimeout(timeout time.Duration) *GetClusterVlansParams {
	return &GetClusterVlansParams{
		timeout: timeout,
	}
}

// NewGetClusterVlansParamsWithContext creates a new GetClusterVlansParams object
// with the ability to set a context for a request.
func NewGetClusterVlansParamsWithContext(ctx context.Context) *GetClusterVlansParams {
	return &GetClusterVlansParams{
		Context: ctx,
	}
}

// NewGetClusterVlansParamsWithHTTPClient creates a new GetClusterVlansParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClusterVlansParamsWithHTTPClient(client *http.Client) *GetClusterVlansParams {
	return &GetClusterVlansParams{
		HTTPClient: client,
	}
}

/*
GetClusterVlansParams contains all the parameters to send to the API endpoint

	for the get cluster vlans operation.

	Typically these are written to a http.Request.
*/
type GetClusterVlansParams struct {

	/* CompressIpsToRanges.

	   Compress vlan IPs to list of contigous IP ranges with startIp and endIp.
	*/
	CompressIpsToRanges *bool

	/* IncludeTenants.

	   If true, the response includes vlans which belongs to all the tenants the current user has permissions to see.

	   Default: true
	*/
	IncludeTenants *bool

	/* InterfaceNames.

	   Vlan interface names, it should be in interface_group_name.vlan_id format.
	*/
	InterfaceNames []string

	/* SkipPrimaryAndBondIface.

	   If true, vlan primary and bond interfaces are not returned in the response.
	*/
	SkipPrimaryAndBondIface *bool

	/* TenantIds.

	   Ids of the tenants, used to get vlans assigned to tenants.
	*/
	TenantIds []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cluster vlans params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterVlansParams) WithDefaults() *GetClusterVlansParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster vlans params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterVlansParams) SetDefaults() {
	var (
		compressIpsToRangesDefault = bool(false)

		includeTenantsDefault = bool(true)

		skipPrimaryAndBondIfaceDefault = bool(false)
	)

	val := GetClusterVlansParams{
		CompressIpsToRanges:     &compressIpsToRangesDefault,
		IncludeTenants:          &includeTenantsDefault,
		SkipPrimaryAndBondIface: &skipPrimaryAndBondIfaceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get cluster vlans params
func (o *GetClusterVlansParams) WithTimeout(timeout time.Duration) *GetClusterVlansParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster vlans params
func (o *GetClusterVlansParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster vlans params
func (o *GetClusterVlansParams) WithContext(ctx context.Context) *GetClusterVlansParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster vlans params
func (o *GetClusterVlansParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster vlans params
func (o *GetClusterVlansParams) WithHTTPClient(client *http.Client) *GetClusterVlansParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster vlans params
func (o *GetClusterVlansParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCompressIpsToRanges adds the compressIpsToRanges to the get cluster vlans params
func (o *GetClusterVlansParams) WithCompressIpsToRanges(compressIpsToRanges *bool) *GetClusterVlansParams {
	o.SetCompressIpsToRanges(compressIpsToRanges)
	return o
}

// SetCompressIpsToRanges adds the compressIpsToRanges to the get cluster vlans params
func (o *GetClusterVlansParams) SetCompressIpsToRanges(compressIpsToRanges *bool) {
	o.CompressIpsToRanges = compressIpsToRanges
}

// WithIncludeTenants adds the includeTenants to the get cluster vlans params
func (o *GetClusterVlansParams) WithIncludeTenants(includeTenants *bool) *GetClusterVlansParams {
	o.SetIncludeTenants(includeTenants)
	return o
}

// SetIncludeTenants adds the includeTenants to the get cluster vlans params
func (o *GetClusterVlansParams) SetIncludeTenants(includeTenants *bool) {
	o.IncludeTenants = includeTenants
}

// WithInterfaceNames adds the interfaceNames to the get cluster vlans params
func (o *GetClusterVlansParams) WithInterfaceNames(interfaceNames []string) *GetClusterVlansParams {
	o.SetInterfaceNames(interfaceNames)
	return o
}

// SetInterfaceNames adds the interfaceNames to the get cluster vlans params
func (o *GetClusterVlansParams) SetInterfaceNames(interfaceNames []string) {
	o.InterfaceNames = interfaceNames
}

// WithSkipPrimaryAndBondIface adds the skipPrimaryAndBondIface to the get cluster vlans params
func (o *GetClusterVlansParams) WithSkipPrimaryAndBondIface(skipPrimaryAndBondIface *bool) *GetClusterVlansParams {
	o.SetSkipPrimaryAndBondIface(skipPrimaryAndBondIface)
	return o
}

// SetSkipPrimaryAndBondIface adds the skipPrimaryAndBondIface to the get cluster vlans params
func (o *GetClusterVlansParams) SetSkipPrimaryAndBondIface(skipPrimaryAndBondIface *bool) {
	o.SkipPrimaryAndBondIface = skipPrimaryAndBondIface
}

// WithTenantIds adds the tenantIds to the get cluster vlans params
func (o *GetClusterVlansParams) WithTenantIds(tenantIds []string) *GetClusterVlansParams {
	o.SetTenantIds(tenantIds)
	return o
}

// SetTenantIds adds the tenantIds to the get cluster vlans params
func (o *GetClusterVlansParams) SetTenantIds(tenantIds []string) {
	o.TenantIds = tenantIds
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterVlansParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CompressIpsToRanges != nil {

		// query param compressIpsToRanges
		var qrCompressIpsToRanges bool

		if o.CompressIpsToRanges != nil {
			qrCompressIpsToRanges = *o.CompressIpsToRanges
		}
		qCompressIpsToRanges := swag.FormatBool(qrCompressIpsToRanges)
		if qCompressIpsToRanges != "" {

			if err := r.SetQueryParam("compressIpsToRanges", qCompressIpsToRanges); err != nil {
				return err
			}
		}
	}

	if o.IncludeTenants != nil {

		// query param includeTenants
		var qrIncludeTenants bool

		if o.IncludeTenants != nil {
			qrIncludeTenants = *o.IncludeTenants
		}
		qIncludeTenants := swag.FormatBool(qrIncludeTenants)
		if qIncludeTenants != "" {

			if err := r.SetQueryParam("includeTenants", qIncludeTenants); err != nil {
				return err
			}
		}
	}

	if o.InterfaceNames != nil {

		// binding items for interfaceNames
		joinedInterfaceNames := o.bindParamInterfaceNames(reg)

		// query array param interfaceNames
		if err := r.SetQueryParam("interfaceNames", joinedInterfaceNames...); err != nil {
			return err
		}
	}

	if o.SkipPrimaryAndBondIface != nil {

		// query param skipPrimaryAndBondIface
		var qrSkipPrimaryAndBondIface bool

		if o.SkipPrimaryAndBondIface != nil {
			qrSkipPrimaryAndBondIface = *o.SkipPrimaryAndBondIface
		}
		qSkipPrimaryAndBondIface := swag.FormatBool(qrSkipPrimaryAndBondIface)
		if qSkipPrimaryAndBondIface != "" {

			if err := r.SetQueryParam("skipPrimaryAndBondIface", qSkipPrimaryAndBondIface); err != nil {
				return err
			}
		}
	}

	if o.TenantIds != nil {

		// binding items for tenantIds
		joinedTenantIds := o.bindParamTenantIds(reg)

		// query array param tenantIds
		if err := r.SetQueryParam("tenantIds", joinedTenantIds...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetClusterVlans binds the parameter interfaceNames
func (o *GetClusterVlansParams) bindParamInterfaceNames(formats strfmt.Registry) []string {
	interfaceNamesIR := o.InterfaceNames

	var interfaceNamesIC []string
	for _, interfaceNamesIIR := range interfaceNamesIR { // explode []string

		interfaceNamesIIV := interfaceNamesIIR // string as string
		interfaceNamesIC = append(interfaceNamesIC, interfaceNamesIIV)
	}

	// items.CollectionFormat: ""
	interfaceNamesIS := swag.JoinByFormat(interfaceNamesIC, "")

	return interfaceNamesIS
}

// bindParamGetClusterVlans binds the parameter tenantIds
func (o *GetClusterVlansParams) bindParamTenantIds(formats strfmt.Registry) []string {
	tenantIdsIR := o.TenantIds

	var tenantIdsIC []string
	for _, tenantIdsIIR := range tenantIdsIR { // explode []string

		tenantIdsIIV := tenantIdsIIR // string as string
		tenantIdsIC = append(tenantIdsIC, tenantIdsIIV)
	}

	// items.CollectionFormat: ""
	tenantIdsIS := swag.JoinByFormat(tenantIdsIC, "")

	return tenantIdsIS
}
