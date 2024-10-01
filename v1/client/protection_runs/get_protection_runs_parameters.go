// Code generated by go-swagger; DO NOT EDIT.

package protection_runs

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

// NewGetProtectionRunsParams creates a new GetProtectionRunsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProtectionRunsParams() *GetProtectionRunsParams {
	return &GetProtectionRunsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProtectionRunsParamsWithTimeout creates a new GetProtectionRunsParams object
// with the ability to set a timeout on a request.
func NewGetProtectionRunsParamsWithTimeout(timeout time.Duration) *GetProtectionRunsParams {
	return &GetProtectionRunsParams{
		timeout: timeout,
	}
}

// NewGetProtectionRunsParamsWithContext creates a new GetProtectionRunsParams object
// with the ability to set a context for a request.
func NewGetProtectionRunsParamsWithContext(ctx context.Context) *GetProtectionRunsParams {
	return &GetProtectionRunsParams{
		Context: ctx,
	}
}

// NewGetProtectionRunsParamsWithHTTPClient creates a new GetProtectionRunsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProtectionRunsParamsWithHTTPClient(client *http.Client) *GetProtectionRunsParams {
	return &GetProtectionRunsParams{
		HTTPClient: client,
	}
}

/*
GetProtectionRunsParams contains all the parameters to send to the API endpoint

	for the get protection runs operation.

	Typically these are written to a http.Request.
*/
type GetProtectionRunsParams struct {

	/* CopyRunTypes.

	     Following field contains a list of valid CopyTarget i.e Replication,
	Archival representing the types of copy runs needed to be looked at
	the given time window, denoted via TimeRange.
	This input is only considered valid when
	filterByCopyTaskEndTime is set to true. Else it is ignored.
	*/
	CopyRunTypes []string

	/* EndTimeUsecs.

	     Filter by a end time specified as a Unix epoch Timestamp
	(in microseconds). Only Job Runs that completed before the specified
	end time are returned.

	     Format: int64
	*/
	EndTimeUsecs *int64

	/* ExcludeErrorRuns.

	     Filter out Jobs Runs with errors by setting this field to 'true'.
	If not set or set to 'false', Job Runs with errors are returned.
	*/
	ExcludeErrorRuns *bool

	/* ExcludeNonRestoreableRuns.

	     Filter out jobs runs that cannot be restored by setting this field to
	'true'. If not set or set to 'false', Runs without any successful
	object will be returned. The default value is false.
	*/
	ExcludeNonRestoreableRuns *bool

	/* ExcludeTasks.

	     If true, the individual backup status for all the objects protected by
	the Job Run are not populated in the response. For example in a VMware
	environment, the status of backing up each VM associated with a Job
	is not returned.
	*/
	ExcludeTasks *bool

	/* FilterByCopyTaskEndTime.

	     If true, then the details of the backup runs along with CopyRuns will be
	returned where those backup run has atleast one CopyTask that is completed
	in the given Time Range. If this field is true, then other filters such as
	filterByEndTime should not be applied.
	*/
	FilterByCopyTaskEndTime *bool

	/* FilterByEndTime.

	     If true, the runs with end time within the specified time range will be
	returned. Otherwise, the runs with start time in the time range are
	returned.
	*/
	FilterByEndTime *bool

	/* IncludeRpoSnapshots.

	     If true, then the snapshots for Protection Sources protected by Rpo
	policies will also be returned.
	*/
	IncludeRpoSnapshots *bool

	/* JobID.

	     Filter by a Protection Job that is specified by id.
	If not specified, all Job Runs for all Protection Jobs are returned.

	     Format: int64
	*/
	JobID *int64

	/* NumRuns.

	     Specify the number of Job Runs to return.
	The newest Job Runs are returned.

	     Format: int64
	*/
	NumRuns *int64

	/* OnlyIncludeSuccessfulCopyRuns.

	     If marked false, all CopyTasks in any finished state like cancelled,
	failed in the given time window will be considered. Otherwise if kept
	empty or marked as true, only the copy_tasks with kSuccess status will
	be considered. This input is only considered valid when
	filterByCopyTaskEndTime is set to true. Else it is ignored.
	*/
	OnlyIncludeSuccessfulCopyRuns *bool

	/* OnlyReturnShellInfo.

	     If passed as true, then only returns the summary information about run
	including details such as runs start time, status, type etc. It does not
	include extra details such as attempt/task info etc.
	*/
	OnlyReturnShellInfo *bool

	/* RunTypes.

	     Filter by run type such as 'kFull', 'kRegular' or 'kLog'.
	If not specified, Job Runs of all types are returned.
	*/
	RunTypes []string

	/* SourceID.

	     Filter by source id. Only Job Runs protecting the specified source
	(such as a VM or View) are returned. The source id
	is assigned by the Cohesity Cluster.

	     Format: int64
	*/
	SourceID *int64

	/* StartTimeUsecs.

	     Filter by a start time. Only Job Runs that started after the specified
	time are returned.
	Specify the start time as a Unix epoch Timestamp (in microseconds).

	     Format: int64
	*/
	StartTimeUsecs *int64

	/* StartedTimeUsecs.

	     Return a specific Job Run by specifying a time and a jobId.
	Specify the time when the Job Run started as a
	Unix epoch Timestamp (in microseconds).
	If this field is specified, jobId must also be
	specified.

	     Format: int64
	*/
	StartedTimeUsecs *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get protection runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProtectionRunsParams) WithDefaults() *GetProtectionRunsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get protection runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProtectionRunsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get protection runs params
func (o *GetProtectionRunsParams) WithTimeout(timeout time.Duration) *GetProtectionRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get protection runs params
func (o *GetProtectionRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get protection runs params
func (o *GetProtectionRunsParams) WithContext(ctx context.Context) *GetProtectionRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get protection runs params
func (o *GetProtectionRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get protection runs params
func (o *GetProtectionRunsParams) WithHTTPClient(client *http.Client) *GetProtectionRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get protection runs params
func (o *GetProtectionRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCopyRunTypes adds the copyRunTypes to the get protection runs params
func (o *GetProtectionRunsParams) WithCopyRunTypes(copyRunTypes []string) *GetProtectionRunsParams {
	o.SetCopyRunTypes(copyRunTypes)
	return o
}

// SetCopyRunTypes adds the copyRunTypes to the get protection runs params
func (o *GetProtectionRunsParams) SetCopyRunTypes(copyRunTypes []string) {
	o.CopyRunTypes = copyRunTypes
}

// WithEndTimeUsecs adds the endTimeUsecs to the get protection runs params
func (o *GetProtectionRunsParams) WithEndTimeUsecs(endTimeUsecs *int64) *GetProtectionRunsParams {
	o.SetEndTimeUsecs(endTimeUsecs)
	return o
}

// SetEndTimeUsecs adds the endTimeUsecs to the get protection runs params
func (o *GetProtectionRunsParams) SetEndTimeUsecs(endTimeUsecs *int64) {
	o.EndTimeUsecs = endTimeUsecs
}

// WithExcludeErrorRuns adds the excludeErrorRuns to the get protection runs params
func (o *GetProtectionRunsParams) WithExcludeErrorRuns(excludeErrorRuns *bool) *GetProtectionRunsParams {
	o.SetExcludeErrorRuns(excludeErrorRuns)
	return o
}

// SetExcludeErrorRuns adds the excludeErrorRuns to the get protection runs params
func (o *GetProtectionRunsParams) SetExcludeErrorRuns(excludeErrorRuns *bool) {
	o.ExcludeErrorRuns = excludeErrorRuns
}

// WithExcludeNonRestoreableRuns adds the excludeNonRestoreableRuns to the get protection runs params
func (o *GetProtectionRunsParams) WithExcludeNonRestoreableRuns(excludeNonRestoreableRuns *bool) *GetProtectionRunsParams {
	o.SetExcludeNonRestoreableRuns(excludeNonRestoreableRuns)
	return o
}

// SetExcludeNonRestoreableRuns adds the excludeNonRestoreableRuns to the get protection runs params
func (o *GetProtectionRunsParams) SetExcludeNonRestoreableRuns(excludeNonRestoreableRuns *bool) {
	o.ExcludeNonRestoreableRuns = excludeNonRestoreableRuns
}

// WithExcludeTasks adds the excludeTasks to the get protection runs params
func (o *GetProtectionRunsParams) WithExcludeTasks(excludeTasks *bool) *GetProtectionRunsParams {
	o.SetExcludeTasks(excludeTasks)
	return o
}

// SetExcludeTasks adds the excludeTasks to the get protection runs params
func (o *GetProtectionRunsParams) SetExcludeTasks(excludeTasks *bool) {
	o.ExcludeTasks = excludeTasks
}

// WithFilterByCopyTaskEndTime adds the filterByCopyTaskEndTime to the get protection runs params
func (o *GetProtectionRunsParams) WithFilterByCopyTaskEndTime(filterByCopyTaskEndTime *bool) *GetProtectionRunsParams {
	o.SetFilterByCopyTaskEndTime(filterByCopyTaskEndTime)
	return o
}

// SetFilterByCopyTaskEndTime adds the filterByCopyTaskEndTime to the get protection runs params
func (o *GetProtectionRunsParams) SetFilterByCopyTaskEndTime(filterByCopyTaskEndTime *bool) {
	o.FilterByCopyTaskEndTime = filterByCopyTaskEndTime
}

// WithFilterByEndTime adds the filterByEndTime to the get protection runs params
func (o *GetProtectionRunsParams) WithFilterByEndTime(filterByEndTime *bool) *GetProtectionRunsParams {
	o.SetFilterByEndTime(filterByEndTime)
	return o
}

// SetFilterByEndTime adds the filterByEndTime to the get protection runs params
func (o *GetProtectionRunsParams) SetFilterByEndTime(filterByEndTime *bool) {
	o.FilterByEndTime = filterByEndTime
}

// WithIncludeRpoSnapshots adds the includeRpoSnapshots to the get protection runs params
func (o *GetProtectionRunsParams) WithIncludeRpoSnapshots(includeRpoSnapshots *bool) *GetProtectionRunsParams {
	o.SetIncludeRpoSnapshots(includeRpoSnapshots)
	return o
}

// SetIncludeRpoSnapshots adds the includeRpoSnapshots to the get protection runs params
func (o *GetProtectionRunsParams) SetIncludeRpoSnapshots(includeRpoSnapshots *bool) {
	o.IncludeRpoSnapshots = includeRpoSnapshots
}

// WithJobID adds the jobID to the get protection runs params
func (o *GetProtectionRunsParams) WithJobID(jobID *int64) *GetProtectionRunsParams {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the get protection runs params
func (o *GetProtectionRunsParams) SetJobID(jobID *int64) {
	o.JobID = jobID
}

// WithNumRuns adds the numRuns to the get protection runs params
func (o *GetProtectionRunsParams) WithNumRuns(numRuns *int64) *GetProtectionRunsParams {
	o.SetNumRuns(numRuns)
	return o
}

// SetNumRuns adds the numRuns to the get protection runs params
func (o *GetProtectionRunsParams) SetNumRuns(numRuns *int64) {
	o.NumRuns = numRuns
}

// WithOnlyIncludeSuccessfulCopyRuns adds the onlyIncludeSuccessfulCopyRuns to the get protection runs params
func (o *GetProtectionRunsParams) WithOnlyIncludeSuccessfulCopyRuns(onlyIncludeSuccessfulCopyRuns *bool) *GetProtectionRunsParams {
	o.SetOnlyIncludeSuccessfulCopyRuns(onlyIncludeSuccessfulCopyRuns)
	return o
}

// SetOnlyIncludeSuccessfulCopyRuns adds the onlyIncludeSuccessfulCopyRuns to the get protection runs params
func (o *GetProtectionRunsParams) SetOnlyIncludeSuccessfulCopyRuns(onlyIncludeSuccessfulCopyRuns *bool) {
	o.OnlyIncludeSuccessfulCopyRuns = onlyIncludeSuccessfulCopyRuns
}

// WithOnlyReturnShellInfo adds the onlyReturnShellInfo to the get protection runs params
func (o *GetProtectionRunsParams) WithOnlyReturnShellInfo(onlyReturnShellInfo *bool) *GetProtectionRunsParams {
	o.SetOnlyReturnShellInfo(onlyReturnShellInfo)
	return o
}

// SetOnlyReturnShellInfo adds the onlyReturnShellInfo to the get protection runs params
func (o *GetProtectionRunsParams) SetOnlyReturnShellInfo(onlyReturnShellInfo *bool) {
	o.OnlyReturnShellInfo = onlyReturnShellInfo
}

// WithRunTypes adds the runTypes to the get protection runs params
func (o *GetProtectionRunsParams) WithRunTypes(runTypes []string) *GetProtectionRunsParams {
	o.SetRunTypes(runTypes)
	return o
}

// SetRunTypes adds the runTypes to the get protection runs params
func (o *GetProtectionRunsParams) SetRunTypes(runTypes []string) {
	o.RunTypes = runTypes
}

// WithSourceID adds the sourceID to the get protection runs params
func (o *GetProtectionRunsParams) WithSourceID(sourceID *int64) *GetProtectionRunsParams {
	o.SetSourceID(sourceID)
	return o
}

// SetSourceID adds the sourceId to the get protection runs params
func (o *GetProtectionRunsParams) SetSourceID(sourceID *int64) {
	o.SourceID = sourceID
}

// WithStartTimeUsecs adds the startTimeUsecs to the get protection runs params
func (o *GetProtectionRunsParams) WithStartTimeUsecs(startTimeUsecs *int64) *GetProtectionRunsParams {
	o.SetStartTimeUsecs(startTimeUsecs)
	return o
}

// SetStartTimeUsecs adds the startTimeUsecs to the get protection runs params
func (o *GetProtectionRunsParams) SetStartTimeUsecs(startTimeUsecs *int64) {
	o.StartTimeUsecs = startTimeUsecs
}

// WithStartedTimeUsecs adds the startedTimeUsecs to the get protection runs params
func (o *GetProtectionRunsParams) WithStartedTimeUsecs(startedTimeUsecs *int64) *GetProtectionRunsParams {
	o.SetStartedTimeUsecs(startedTimeUsecs)
	return o
}

// SetStartedTimeUsecs adds the startedTimeUsecs to the get protection runs params
func (o *GetProtectionRunsParams) SetStartedTimeUsecs(startedTimeUsecs *int64) {
	o.StartedTimeUsecs = startedTimeUsecs
}

// WriteToRequest writes these params to a swagger request
func (o *GetProtectionRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CopyRunTypes != nil {

		// binding items for copyRunTypes
		joinedCopyRunTypes := o.bindParamCopyRunTypes(reg)

		// query array param copyRunTypes
		if err := r.SetQueryParam("copyRunTypes", joinedCopyRunTypes...); err != nil {
			return err
		}
	}

	if o.EndTimeUsecs != nil {

		// query param endTimeUsecs
		var qrEndTimeUsecs int64

		if o.EndTimeUsecs != nil {
			qrEndTimeUsecs = *o.EndTimeUsecs
		}
		qEndTimeUsecs := swag.FormatInt64(qrEndTimeUsecs)
		if qEndTimeUsecs != "" {

			if err := r.SetQueryParam("endTimeUsecs", qEndTimeUsecs); err != nil {
				return err
			}
		}
	}

	if o.ExcludeErrorRuns != nil {

		// query param excludeErrorRuns
		var qrExcludeErrorRuns bool

		if o.ExcludeErrorRuns != nil {
			qrExcludeErrorRuns = *o.ExcludeErrorRuns
		}
		qExcludeErrorRuns := swag.FormatBool(qrExcludeErrorRuns)
		if qExcludeErrorRuns != "" {

			if err := r.SetQueryParam("excludeErrorRuns", qExcludeErrorRuns); err != nil {
				return err
			}
		}
	}

	if o.ExcludeNonRestoreableRuns != nil {

		// query param excludeNonRestoreableRuns
		var qrExcludeNonRestoreableRuns bool

		if o.ExcludeNonRestoreableRuns != nil {
			qrExcludeNonRestoreableRuns = *o.ExcludeNonRestoreableRuns
		}
		qExcludeNonRestoreableRuns := swag.FormatBool(qrExcludeNonRestoreableRuns)
		if qExcludeNonRestoreableRuns != "" {

			if err := r.SetQueryParam("excludeNonRestoreableRuns", qExcludeNonRestoreableRuns); err != nil {
				return err
			}
		}
	}

	if o.ExcludeTasks != nil {

		// query param excludeTasks
		var qrExcludeTasks bool

		if o.ExcludeTasks != nil {
			qrExcludeTasks = *o.ExcludeTasks
		}
		qExcludeTasks := swag.FormatBool(qrExcludeTasks)
		if qExcludeTasks != "" {

			if err := r.SetQueryParam("excludeTasks", qExcludeTasks); err != nil {
				return err
			}
		}
	}

	if o.FilterByCopyTaskEndTime != nil {

		// query param filterByCopyTaskEndTime
		var qrFilterByCopyTaskEndTime bool

		if o.FilterByCopyTaskEndTime != nil {
			qrFilterByCopyTaskEndTime = *o.FilterByCopyTaskEndTime
		}
		qFilterByCopyTaskEndTime := swag.FormatBool(qrFilterByCopyTaskEndTime)
		if qFilterByCopyTaskEndTime != "" {

			if err := r.SetQueryParam("filterByCopyTaskEndTime", qFilterByCopyTaskEndTime); err != nil {
				return err
			}
		}
	}

	if o.FilterByEndTime != nil {

		// query param filterByEndTime
		var qrFilterByEndTime bool

		if o.FilterByEndTime != nil {
			qrFilterByEndTime = *o.FilterByEndTime
		}
		qFilterByEndTime := swag.FormatBool(qrFilterByEndTime)
		if qFilterByEndTime != "" {

			if err := r.SetQueryParam("filterByEndTime", qFilterByEndTime); err != nil {
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

	if o.JobID != nil {

		// query param jobId
		var qrJobID int64

		if o.JobID != nil {
			qrJobID = *o.JobID
		}
		qJobID := swag.FormatInt64(qrJobID)
		if qJobID != "" {

			if err := r.SetQueryParam("jobId", qJobID); err != nil {
				return err
			}
		}
	}

	if o.NumRuns != nil {

		// query param numRuns
		var qrNumRuns int64

		if o.NumRuns != nil {
			qrNumRuns = *o.NumRuns
		}
		qNumRuns := swag.FormatInt64(qrNumRuns)
		if qNumRuns != "" {

			if err := r.SetQueryParam("numRuns", qNumRuns); err != nil {
				return err
			}
		}
	}

	if o.OnlyIncludeSuccessfulCopyRuns != nil {

		// query param onlyIncludeSuccessfulCopyRuns
		var qrOnlyIncludeSuccessfulCopyRuns bool

		if o.OnlyIncludeSuccessfulCopyRuns != nil {
			qrOnlyIncludeSuccessfulCopyRuns = *o.OnlyIncludeSuccessfulCopyRuns
		}
		qOnlyIncludeSuccessfulCopyRuns := swag.FormatBool(qrOnlyIncludeSuccessfulCopyRuns)
		if qOnlyIncludeSuccessfulCopyRuns != "" {

			if err := r.SetQueryParam("onlyIncludeSuccessfulCopyRuns", qOnlyIncludeSuccessfulCopyRuns); err != nil {
				return err
			}
		}
	}

	if o.OnlyReturnShellInfo != nil {

		// query param onlyReturnShellInfo
		var qrOnlyReturnShellInfo bool

		if o.OnlyReturnShellInfo != nil {
			qrOnlyReturnShellInfo = *o.OnlyReturnShellInfo
		}
		qOnlyReturnShellInfo := swag.FormatBool(qrOnlyReturnShellInfo)
		if qOnlyReturnShellInfo != "" {

			if err := r.SetQueryParam("onlyReturnShellInfo", qOnlyReturnShellInfo); err != nil {
				return err
			}
		}
	}

	if o.RunTypes != nil {

		// binding items for runTypes
		joinedRunTypes := o.bindParamRunTypes(reg)

		// query array param runTypes
		if err := r.SetQueryParam("runTypes", joinedRunTypes...); err != nil {
			return err
		}
	}

	if o.SourceID != nil {

		// query param sourceId
		var qrSourceID int64

		if o.SourceID != nil {
			qrSourceID = *o.SourceID
		}
		qSourceID := swag.FormatInt64(qrSourceID)
		if qSourceID != "" {

			if err := r.SetQueryParam("sourceId", qSourceID); err != nil {
				return err
			}
		}
	}

	if o.StartTimeUsecs != nil {

		// query param startTimeUsecs
		var qrStartTimeUsecs int64

		if o.StartTimeUsecs != nil {
			qrStartTimeUsecs = *o.StartTimeUsecs
		}
		qStartTimeUsecs := swag.FormatInt64(qrStartTimeUsecs)
		if qStartTimeUsecs != "" {

			if err := r.SetQueryParam("startTimeUsecs", qStartTimeUsecs); err != nil {
				return err
			}
		}
	}

	if o.StartedTimeUsecs != nil {

		// query param startedTimeUsecs
		var qrStartedTimeUsecs int64

		if o.StartedTimeUsecs != nil {
			qrStartedTimeUsecs = *o.StartedTimeUsecs
		}
		qStartedTimeUsecs := swag.FormatInt64(qrStartedTimeUsecs)
		if qStartedTimeUsecs != "" {

			if err := r.SetQueryParam("startedTimeUsecs", qStartedTimeUsecs); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetProtectionRuns binds the parameter copyRunTypes
func (o *GetProtectionRunsParams) bindParamCopyRunTypes(formats strfmt.Registry) []string {
	copyRunTypesIR := o.CopyRunTypes

	var copyRunTypesIC []string
	for _, copyRunTypesIIR := range copyRunTypesIR { // explode []string

		copyRunTypesIIV := copyRunTypesIIR // string as string
		copyRunTypesIC = append(copyRunTypesIC, copyRunTypesIIV)
	}

	// items.CollectionFormat: ""
	copyRunTypesIS := swag.JoinByFormat(copyRunTypesIC, "")

	return copyRunTypesIS
}

// bindParamGetProtectionRuns binds the parameter runTypes
func (o *GetProtectionRunsParams) bindParamRunTypes(formats strfmt.Registry) []string {
	runTypesIR := o.RunTypes

	var runTypesIC []string
	for _, runTypesIIR := range runTypesIR { // explode []string

		runTypesIIV := runTypesIIR // string as string
		runTypesIC = append(runTypesIC, runTypesIIV)
	}

	// items.CollectionFormat: ""
	runTypesIS := swag.JoinByFormat(runTypesIC, "")

	return runTypesIS
}
