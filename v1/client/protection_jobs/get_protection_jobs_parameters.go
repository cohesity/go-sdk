// Code generated by go-swagger; DO NOT EDIT.

package protection_jobs

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

// NewGetProtectionJobsParams creates a new GetProtectionJobsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProtectionJobsParams() *GetProtectionJobsParams {
	return &GetProtectionJobsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProtectionJobsParamsWithTimeout creates a new GetProtectionJobsParams object
// with the ability to set a timeout on a request.
func NewGetProtectionJobsParamsWithTimeout(timeout time.Duration) *GetProtectionJobsParams {
	return &GetProtectionJobsParams{
		timeout: timeout,
	}
}

// NewGetProtectionJobsParamsWithContext creates a new GetProtectionJobsParams object
// with the ability to set a context for a request.
func NewGetProtectionJobsParamsWithContext(ctx context.Context) *GetProtectionJobsParams {
	return &GetProtectionJobsParams{
		Context: ctx,
	}
}

// NewGetProtectionJobsParamsWithHTTPClient creates a new GetProtectionJobsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProtectionJobsParamsWithHTTPClient(client *http.Client) *GetProtectionJobsParams {
	return &GetProtectionJobsParams{
		HTTPClient: client,
	}
}

/*
GetProtectionJobsParams contains all the parameters to send to the API endpoint

	for the get protection jobs operation.

	Typically these are written to a http.Request.
*/
type GetProtectionJobsParams struct {

	/* AllUnderHierarchy.

	     AllUnderHierarchy specifies if objects of all the tenants under the
	hierarchy of the logged in user's organization should be returned.
	*/
	AllUnderHierarchy *bool

	/* Environments.

	     Filter by environment types such as 'kVMware', 'kView', etc.
	Only Jobs protecting the specified environment types are returned.
	NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter.
	*/
	Environments []string

	/* Ids.

	   Filter by a list of Protection Job ids.
	*/
	Ids []int64

	/* IncludeLastRunAndStats.

	   If true, return the last Protection Run of the Job and the summary stats.
	*/
	IncludeLastRunAndStats *bool

	/* IncludeRpoSnapshots.

	     If true, then the Protected Objects protected by RPO policies will also
	be returned.
	*/
	IncludeRpoSnapshots *bool

	/* IncludeSourceNames.

	     If true, both names and ids of Sources inside the protection group will be
	sent in the response.
	If false, only Source ids will be sent in the response.
	*/
	IncludeSourceNames *bool

	/* IsActive.

	     Filter by Inactive or Active Jobs. If not set, all Inactive and
	Active Jobs are returned. If true, only Active Jobs are returned.
	If false, only Inactive Jobs are returned.
	When you create a Protection Job on a Primary Cluster
	with a replication schedule, the Cluster creates an
	Inactive copy of the Job on the Remote Cluster.
	In addition, when an Active and running Job is deactivated,
	the Job becomes Inactive.
	*/
	IsActive *bool

	/* IsDeleted.

	     If true, return only Protection Jobs that have been deleted but still
	have Snapshots associated with them.
	If false, return all Protection Jobs except those Jobs that have
	been deleted and still have Snapshots associated with them.
	A Job that is deleted with all its Snapshots is not returned for
	either of these cases.
	*/
	IsDeleted *bool

	/* IsLastRunSLAViolated.

	     IsLastRunSlaViolated is the parameter to filter the Protection Jobs based
	on the SLA violation status of the last Protection Run.
	*/
	IsLastRunSLAViolated *bool

	/* MaxResultCount.

	     Identifies the max number of items to be returned. This is specifically
	to be used with pagination.

	     Format: int64
	*/
	MaxResultCount *int64

	/* Names.

	   Filter by a list of Protection Job names.
	*/
	Names []string

	/* OnlyReturnBasicSummary.

	     if true then only job descriptions and the most recent run of the job
	will be returned.
	*/
	OnlyReturnBasicSummary *bool

	/* OnlyReturnDataMigrationJobs.

	     OnlyReturnDataMigrationJobs specifies if only data migration jobs should be
	returned. If not set, no data migration job will be returned.
	*/
	OnlyReturnDataMigrationJobs *bool

	/* PaginationCookie.

	   Pagination cookie to fetch the next set of results.
	*/
	PaginationCookie *string

	/* PolicyIds.

	     Filter by Policy ids that are associated with Protection Jobs.
	Only Jobs associated with the specified Policy ids, are returned.
	*/
	PolicyIds []string

	/* PruneExcludedSourceIds.

	     If true, the list of exclusion sources will be omitted from the response.
	This can be used to improve performance when the exclusion sources are
	not needed.
	*/
	PruneExcludedSourceIds *bool

	/* TenantIds.

	     TenantIds contains ids of the tenants for which objects are to be
	returned.
	*/
	TenantIds []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get protection jobs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProtectionJobsParams) WithDefaults() *GetProtectionJobsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get protection jobs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProtectionJobsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get protection jobs params
func (o *GetProtectionJobsParams) WithTimeout(timeout time.Duration) *GetProtectionJobsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get protection jobs params
func (o *GetProtectionJobsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get protection jobs params
func (o *GetProtectionJobsParams) WithContext(ctx context.Context) *GetProtectionJobsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get protection jobs params
func (o *GetProtectionJobsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get protection jobs params
func (o *GetProtectionJobsParams) WithHTTPClient(client *http.Client) *GetProtectionJobsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get protection jobs params
func (o *GetProtectionJobsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllUnderHierarchy adds the allUnderHierarchy to the get protection jobs params
func (o *GetProtectionJobsParams) WithAllUnderHierarchy(allUnderHierarchy *bool) *GetProtectionJobsParams {
	o.SetAllUnderHierarchy(allUnderHierarchy)
	return o
}

// SetAllUnderHierarchy adds the allUnderHierarchy to the get protection jobs params
func (o *GetProtectionJobsParams) SetAllUnderHierarchy(allUnderHierarchy *bool) {
	o.AllUnderHierarchy = allUnderHierarchy
}

// WithEnvironments adds the environments to the get protection jobs params
func (o *GetProtectionJobsParams) WithEnvironments(environments []string) *GetProtectionJobsParams {
	o.SetEnvironments(environments)
	return o
}

// SetEnvironments adds the environments to the get protection jobs params
func (o *GetProtectionJobsParams) SetEnvironments(environments []string) {
	o.Environments = environments
}

// WithIds adds the ids to the get protection jobs params
func (o *GetProtectionJobsParams) WithIds(ids []int64) *GetProtectionJobsParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get protection jobs params
func (o *GetProtectionJobsParams) SetIds(ids []int64) {
	o.Ids = ids
}

// WithIncludeLastRunAndStats adds the includeLastRunAndStats to the get protection jobs params
func (o *GetProtectionJobsParams) WithIncludeLastRunAndStats(includeLastRunAndStats *bool) *GetProtectionJobsParams {
	o.SetIncludeLastRunAndStats(includeLastRunAndStats)
	return o
}

// SetIncludeLastRunAndStats adds the includeLastRunAndStats to the get protection jobs params
func (o *GetProtectionJobsParams) SetIncludeLastRunAndStats(includeLastRunAndStats *bool) {
	o.IncludeLastRunAndStats = includeLastRunAndStats
}

// WithIncludeRpoSnapshots adds the includeRpoSnapshots to the get protection jobs params
func (o *GetProtectionJobsParams) WithIncludeRpoSnapshots(includeRpoSnapshots *bool) *GetProtectionJobsParams {
	o.SetIncludeRpoSnapshots(includeRpoSnapshots)
	return o
}

// SetIncludeRpoSnapshots adds the includeRpoSnapshots to the get protection jobs params
func (o *GetProtectionJobsParams) SetIncludeRpoSnapshots(includeRpoSnapshots *bool) {
	o.IncludeRpoSnapshots = includeRpoSnapshots
}

// WithIncludeSourceNames adds the includeSourceNames to the get protection jobs params
func (o *GetProtectionJobsParams) WithIncludeSourceNames(includeSourceNames *bool) *GetProtectionJobsParams {
	o.SetIncludeSourceNames(includeSourceNames)
	return o
}

// SetIncludeSourceNames adds the includeSourceNames to the get protection jobs params
func (o *GetProtectionJobsParams) SetIncludeSourceNames(includeSourceNames *bool) {
	o.IncludeSourceNames = includeSourceNames
}

// WithIsActive adds the isActive to the get protection jobs params
func (o *GetProtectionJobsParams) WithIsActive(isActive *bool) *GetProtectionJobsParams {
	o.SetIsActive(isActive)
	return o
}

// SetIsActive adds the isActive to the get protection jobs params
func (o *GetProtectionJobsParams) SetIsActive(isActive *bool) {
	o.IsActive = isActive
}

// WithIsDeleted adds the isDeleted to the get protection jobs params
func (o *GetProtectionJobsParams) WithIsDeleted(isDeleted *bool) *GetProtectionJobsParams {
	o.SetIsDeleted(isDeleted)
	return o
}

// SetIsDeleted adds the isDeleted to the get protection jobs params
func (o *GetProtectionJobsParams) SetIsDeleted(isDeleted *bool) {
	o.IsDeleted = isDeleted
}

// WithIsLastRunSLAViolated adds the isLastRunSLAViolated to the get protection jobs params
func (o *GetProtectionJobsParams) WithIsLastRunSLAViolated(isLastRunSLAViolated *bool) *GetProtectionJobsParams {
	o.SetIsLastRunSLAViolated(isLastRunSLAViolated)
	return o
}

// SetIsLastRunSLAViolated adds the isLastRunSlaViolated to the get protection jobs params
func (o *GetProtectionJobsParams) SetIsLastRunSLAViolated(isLastRunSLAViolated *bool) {
	o.IsLastRunSLAViolated = isLastRunSLAViolated
}

// WithMaxResultCount adds the maxResultCount to the get protection jobs params
func (o *GetProtectionJobsParams) WithMaxResultCount(maxResultCount *int64) *GetProtectionJobsParams {
	o.SetMaxResultCount(maxResultCount)
	return o
}

// SetMaxResultCount adds the maxResultCount to the get protection jobs params
func (o *GetProtectionJobsParams) SetMaxResultCount(maxResultCount *int64) {
	o.MaxResultCount = maxResultCount
}

// WithNames adds the names to the get protection jobs params
func (o *GetProtectionJobsParams) WithNames(names []string) *GetProtectionJobsParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the get protection jobs params
func (o *GetProtectionJobsParams) SetNames(names []string) {
	o.Names = names
}

// WithOnlyReturnBasicSummary adds the onlyReturnBasicSummary to the get protection jobs params
func (o *GetProtectionJobsParams) WithOnlyReturnBasicSummary(onlyReturnBasicSummary *bool) *GetProtectionJobsParams {
	o.SetOnlyReturnBasicSummary(onlyReturnBasicSummary)
	return o
}

// SetOnlyReturnBasicSummary adds the onlyReturnBasicSummary to the get protection jobs params
func (o *GetProtectionJobsParams) SetOnlyReturnBasicSummary(onlyReturnBasicSummary *bool) {
	o.OnlyReturnBasicSummary = onlyReturnBasicSummary
}

// WithOnlyReturnDataMigrationJobs adds the onlyReturnDataMigrationJobs to the get protection jobs params
func (o *GetProtectionJobsParams) WithOnlyReturnDataMigrationJobs(onlyReturnDataMigrationJobs *bool) *GetProtectionJobsParams {
	o.SetOnlyReturnDataMigrationJobs(onlyReturnDataMigrationJobs)
	return o
}

// SetOnlyReturnDataMigrationJobs adds the onlyReturnDataMigrationJobs to the get protection jobs params
func (o *GetProtectionJobsParams) SetOnlyReturnDataMigrationJobs(onlyReturnDataMigrationJobs *bool) {
	o.OnlyReturnDataMigrationJobs = onlyReturnDataMigrationJobs
}

// WithPaginationCookie adds the paginationCookie to the get protection jobs params
func (o *GetProtectionJobsParams) WithPaginationCookie(paginationCookie *string) *GetProtectionJobsParams {
	o.SetPaginationCookie(paginationCookie)
	return o
}

// SetPaginationCookie adds the paginationCookie to the get protection jobs params
func (o *GetProtectionJobsParams) SetPaginationCookie(paginationCookie *string) {
	o.PaginationCookie = paginationCookie
}

// WithPolicyIds adds the policyIds to the get protection jobs params
func (o *GetProtectionJobsParams) WithPolicyIds(policyIds []string) *GetProtectionJobsParams {
	o.SetPolicyIds(policyIds)
	return o
}

// SetPolicyIds adds the policyIds to the get protection jobs params
func (o *GetProtectionJobsParams) SetPolicyIds(policyIds []string) {
	o.PolicyIds = policyIds
}

// WithPruneExcludedSourceIds adds the pruneExcludedSourceIds to the get protection jobs params
func (o *GetProtectionJobsParams) WithPruneExcludedSourceIds(pruneExcludedSourceIds *bool) *GetProtectionJobsParams {
	o.SetPruneExcludedSourceIds(pruneExcludedSourceIds)
	return o
}

// SetPruneExcludedSourceIds adds the pruneExcludedSourceIds to the get protection jobs params
func (o *GetProtectionJobsParams) SetPruneExcludedSourceIds(pruneExcludedSourceIds *bool) {
	o.PruneExcludedSourceIds = pruneExcludedSourceIds
}

// WithTenantIds adds the tenantIds to the get protection jobs params
func (o *GetProtectionJobsParams) WithTenantIds(tenantIds []string) *GetProtectionJobsParams {
	o.SetTenantIds(tenantIds)
	return o
}

// SetTenantIds adds the tenantIds to the get protection jobs params
func (o *GetProtectionJobsParams) SetTenantIds(tenantIds []string) {
	o.TenantIds = tenantIds
}

// WriteToRequest writes these params to a swagger request
func (o *GetProtectionJobsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllUnderHierarchy != nil {

		// query param allUnderHierarchy
		var qrAllUnderHierarchy bool

		if o.AllUnderHierarchy != nil {
			qrAllUnderHierarchy = *o.AllUnderHierarchy
		}
		qAllUnderHierarchy := swag.FormatBool(qrAllUnderHierarchy)
		if qAllUnderHierarchy != "" {

			if err := r.SetQueryParam("allUnderHierarchy", qAllUnderHierarchy); err != nil {
				return err
			}
		}
	}

	if o.Environments != nil {

		// binding items for environments
		joinedEnvironments := o.bindParamEnvironments(reg)

		// query array param environments
		if err := r.SetQueryParam("environments", joinedEnvironments...); err != nil {
			return err
		}
	}

	if o.Ids != nil {

		// binding items for ids
		joinedIds := o.bindParamIds(reg)

		// query array param ids
		if err := r.SetQueryParam("ids", joinedIds...); err != nil {
			return err
		}
	}

	if o.IncludeLastRunAndStats != nil {

		// query param includeLastRunAndStats
		var qrIncludeLastRunAndStats bool

		if o.IncludeLastRunAndStats != nil {
			qrIncludeLastRunAndStats = *o.IncludeLastRunAndStats
		}
		qIncludeLastRunAndStats := swag.FormatBool(qrIncludeLastRunAndStats)
		if qIncludeLastRunAndStats != "" {

			if err := r.SetQueryParam("includeLastRunAndStats", qIncludeLastRunAndStats); err != nil {
				return err
			}
		}
	}

	if o.IncludeRpoSnapshots != nil {

		// query param includeRpoSnapshots
		var qrIncludeRpoSnapshots bool

		if o.IncludeRpoSnapshots != nil {
			qrIncludeRpoSnapshots = *o.IncludeRpoSnapshots
		}
		qIncludeRpoSnapshots := swag.FormatBool(qrIncludeRpoSnapshots)
		if qIncludeRpoSnapshots != "" {

			if err := r.SetQueryParam("includeRpoSnapshots", qIncludeRpoSnapshots); err != nil {
				return err
			}
		}
	}

	if o.IncludeSourceNames != nil {

		// query param includeSourceNames
		var qrIncludeSourceNames bool

		if o.IncludeSourceNames != nil {
			qrIncludeSourceNames = *o.IncludeSourceNames
		}
		qIncludeSourceNames := swag.FormatBool(qrIncludeSourceNames)
		if qIncludeSourceNames != "" {

			if err := r.SetQueryParam("includeSourceNames", qIncludeSourceNames); err != nil {
				return err
			}
		}
	}

	if o.IsActive != nil {

		// query param isActive
		var qrIsActive bool

		if o.IsActive != nil {
			qrIsActive = *o.IsActive
		}
		qIsActive := swag.FormatBool(qrIsActive)
		if qIsActive != "" {

			if err := r.SetQueryParam("isActive", qIsActive); err != nil {
				return err
			}
		}
	}

	if o.IsDeleted != nil {

		// query param isDeleted
		var qrIsDeleted bool

		if o.IsDeleted != nil {
			qrIsDeleted = *o.IsDeleted
		}
		qIsDeleted := swag.FormatBool(qrIsDeleted)
		if qIsDeleted != "" {

			if err := r.SetQueryParam("isDeleted", qIsDeleted); err != nil {
				return err
			}
		}
	}

	if o.IsLastRunSLAViolated != nil {

		// query param isLastRunSlaViolated
		var qrIsLastRunSLAViolated bool

		if o.IsLastRunSLAViolated != nil {
			qrIsLastRunSLAViolated = *o.IsLastRunSLAViolated
		}
		qIsLastRunSLAViolated := swag.FormatBool(qrIsLastRunSLAViolated)
		if qIsLastRunSLAViolated != "" {

			if err := r.SetQueryParam("isLastRunSlaViolated", qIsLastRunSLAViolated); err != nil {
				return err
			}
		}
	}

	if o.MaxResultCount != nil {

		// query param maxResultCount
		var qrMaxResultCount int64

		if o.MaxResultCount != nil {
			qrMaxResultCount = *o.MaxResultCount
		}
		qMaxResultCount := swag.FormatInt64(qrMaxResultCount)
		if qMaxResultCount != "" {

			if err := r.SetQueryParam("maxResultCount", qMaxResultCount); err != nil {
				return err
			}
		}
	}

	if o.Names != nil {

		// binding items for names
		joinedNames := o.bindParamNames(reg)

		// query array param names
		if err := r.SetQueryParam("names", joinedNames...); err != nil {
			return err
		}
	}

	if o.OnlyReturnBasicSummary != nil {

		// query param onlyReturnBasicSummary
		var qrOnlyReturnBasicSummary bool

		if o.OnlyReturnBasicSummary != nil {
			qrOnlyReturnBasicSummary = *o.OnlyReturnBasicSummary
		}
		qOnlyReturnBasicSummary := swag.FormatBool(qrOnlyReturnBasicSummary)
		if qOnlyReturnBasicSummary != "" {

			if err := r.SetQueryParam("onlyReturnBasicSummary", qOnlyReturnBasicSummary); err != nil {
				return err
			}
		}
	}

	if o.OnlyReturnDataMigrationJobs != nil {

		// query param onlyReturnDataMigrationJobs
		var qrOnlyReturnDataMigrationJobs bool

		if o.OnlyReturnDataMigrationJobs != nil {
			qrOnlyReturnDataMigrationJobs = *o.OnlyReturnDataMigrationJobs
		}
		qOnlyReturnDataMigrationJobs := swag.FormatBool(qrOnlyReturnDataMigrationJobs)
		if qOnlyReturnDataMigrationJobs != "" {

			if err := r.SetQueryParam("onlyReturnDataMigrationJobs", qOnlyReturnDataMigrationJobs); err != nil {
				return err
			}
		}
	}

	if o.PaginationCookie != nil {

		// query param paginationCookie
		var qrPaginationCookie string

		if o.PaginationCookie != nil {
			qrPaginationCookie = *o.PaginationCookie
		}
		qPaginationCookie := qrPaginationCookie
		if qPaginationCookie != "" {

			if err := r.SetQueryParam("paginationCookie", qPaginationCookie); err != nil {
				return err
			}
		}
	}

	if o.PolicyIds != nil {

		// binding items for policyIds
		joinedPolicyIds := o.bindParamPolicyIds(reg)

		// query array param policyIds
		if err := r.SetQueryParam("policyIds", joinedPolicyIds...); err != nil {
			return err
		}
	}

	if o.PruneExcludedSourceIds != nil {

		// query param pruneExcludedSourceIds
		var qrPruneExcludedSourceIds bool

		if o.PruneExcludedSourceIds != nil {
			qrPruneExcludedSourceIds = *o.PruneExcludedSourceIds
		}
		qPruneExcludedSourceIds := swag.FormatBool(qrPruneExcludedSourceIds)
		if qPruneExcludedSourceIds != "" {

			if err := r.SetQueryParam("pruneExcludedSourceIds", qPruneExcludedSourceIds); err != nil {
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

// bindParamGetProtectionJobs binds the parameter environments
func (o *GetProtectionJobsParams) bindParamEnvironments(formats strfmt.Registry) []string {
	environmentsIR := o.Environments

	var environmentsIC []string
	for _, environmentsIIR := range environmentsIR { // explode []string

		environmentsIIV := environmentsIIR // string as string
		environmentsIC = append(environmentsIC, environmentsIIV)
	}

	// items.CollectionFormat: ""
	environmentsIS := swag.JoinByFormat(environmentsIC, "")

	return environmentsIS
}

// bindParamGetProtectionJobs binds the parameter ids
func (o *GetProtectionJobsParams) bindParamIds(formats strfmt.Registry) []string {
	idsIR := o.Ids

	var idsIC []string
	for _, idsIIR := range idsIR { // explode []int64

		idsIIV := swag.FormatInt64(idsIIR) // int64 as string
		idsIC = append(idsIC, idsIIV)
	}

	// items.CollectionFormat: ""
	idsIS := swag.JoinByFormat(idsIC, "")

	return idsIS
}

// bindParamGetProtectionJobs binds the parameter names
func (o *GetProtectionJobsParams) bindParamNames(formats strfmt.Registry) []string {
	namesIR := o.Names

	var namesIC []string
	for _, namesIIR := range namesIR { // explode []string

		namesIIV := namesIIR // string as string
		namesIC = append(namesIC, namesIIV)
	}

	// items.CollectionFormat: ""
	namesIS := swag.JoinByFormat(namesIC, "")

	return namesIS
}

// bindParamGetProtectionJobs binds the parameter policyIds
func (o *GetProtectionJobsParams) bindParamPolicyIds(formats strfmt.Registry) []string {
	policyIdsIR := o.PolicyIds

	var policyIdsIC []string
	for _, policyIdsIIR := range policyIdsIR { // explode []string

		policyIdsIIV := policyIdsIIR // string as string
		policyIdsIC = append(policyIdsIC, policyIdsIIV)
	}

	// items.CollectionFormat: ""
	policyIdsIS := swag.JoinByFormat(policyIdsIC, "")

	return policyIdsIS
}

// bindParamGetProtectionJobs binds the parameter tenantIds
func (o *GetProtectionJobsParams) bindParamTenantIds(formats strfmt.Registry) []string {
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
