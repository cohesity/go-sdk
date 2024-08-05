# \ProtectionGroupAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProtectionGroup**](ProtectionGroupAPI.md#CreateProtectionGroup) | **Post** /data-protect/protection-groups | Create a Protection Group.
[**CreateProtectionGroupRun**](ProtectionGroupAPI.md#CreateProtectionGroupRun) | **Post** /data-protect/protection-groups/{id}/runs | Create a new protection run.
[**DeleteProtectionGroup**](ProtectionGroupAPI.md#DeleteProtectionGroup) | **Delete** /data-protect/protection-groups/{id} | Delete a Protection Group.
[**GetProtectionGroupById**](ProtectionGroupAPI.md#GetProtectionGroupById) | **Get** /data-protect/protection-groups/{id} | List details about single Protection Group.
[**GetProtectionGroupRun**](ProtectionGroupAPI.md#GetProtectionGroupRun) | **Get** /data-protect/protection-groups/{id}/runs/{runId} | Get a run for a Protection Group.
[**GetProtectionGroupRuns**](ProtectionGroupAPI.md#GetProtectionGroupRuns) | **Get** /data-protect/protection-groups/{id}/runs | Get the list of runs for a Protection Group.
[**GetProtectionGroups**](ProtectionGroupAPI.md#GetProtectionGroups) | **Get** /data-protect/protection-groups | Get the list of Protection Groups.
[**GetProtectionRunProgress**](ProtectionGroupAPI.md#GetProtectionRunProgress) | **Get** /data-protect/runs/{runId}/progress | Get the progress of a run.
[**GetProtectionRunStats**](ProtectionGroupAPI.md#GetProtectionRunStats) | **Get** /data-protect/runs/{runId}/stats | Get the stats for a run.
[**GetProtectionRuns**](ProtectionGroupAPI.md#GetProtectionRuns) | **Get** /data-protect/runs/summary | Get the list of runs.
[**GetRunDebugLogs**](ProtectionGroupAPI.md#GetRunDebugLogs) | **Get** /data-protect/protection-groups/{id}/runs/{runId}/debug-logs | Get the debug logs for a run from a Protection Group.
[**GetRunDebugLogsForObject**](ProtectionGroupAPI.md#GetRunDebugLogsForObject) | **Get** /data-protect/protection-groups/{id}/runs/{runId}/objects/{objectId}/debug-logs | Get the debug logs for a particular object in a run from a Protection Group.
[**GetRunMessagesReport**](ProtectionGroupAPI.md#GetRunMessagesReport) | **Get** /data-protect/protection-groups/{id}/runs/{runId}/objects/{objectId}/download-messages | Get the CSV of various Proto Messages for a given run and an object.
[**GetRunsReport**](ProtectionGroupAPI.md#GetRunsReport) | **Get** /data-protect/protection-groups/{id}/runs/{runId}/objects/{objectId}/downloadFiles | Get the CSV of errors/warnings for a given run and an object.
[**PerformActionOnProtectionGroupRun**](ProtectionGroupAPI.md#PerformActionOnProtectionGroupRun) | **Post** /data-protect/protection-groups/{id}/runs/actions | Actions on protection group run.
[**UpdateProtectionGroup**](ProtectionGroupAPI.md#UpdateProtectionGroup) | **Put** /data-protect/protection-groups/{id} | Update a Protection Group.
[**UpdateProtectionGroupRun**](ProtectionGroupAPI.md#UpdateProtectionGroupRun) | **Put** /data-protect/protection-groups/{id}/runs | Update runs for a particular Protection Group.
[**UpdateProtectionGroupsState**](ProtectionGroupAPI.md#UpdateProtectionGroupsState) | **Post** /data-protect/protection-groups/states | Perform an action like pause, resume, active, deactivate on all specified Protection Groups.



## CreateProtectionGroup

> ProtectionGroup CreateProtectionGroup(ctx).Body(body).Execute()

Create a Protection Group.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewCreateOrUpdateProtectionGroupRequest("Environment_example", "Name_example", "PolicyId_example") // CreateOrUpdateProtectionGroupRequest | Specifies the parameters to create a Protection Group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProtectionGroupAPI.CreateProtectionGroup(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroupAPI.CreateProtectionGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProtectionGroup`: ProtectionGroup
	fmt.Fprintf(os.Stdout, "Response from `ProtectionGroupAPI.CreateProtectionGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProtectionGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateOrUpdateProtectionGroupRequest**](CreateOrUpdateProtectionGroupRequest.md) | Specifies the parameters to create a Protection Group. | 

### Return type

[**ProtectionGroup**](ProtectionGroup.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProtectionGroupRun

> CreateProtectionGroupRunResponseBody CreateProtectionGroupRun(ctx, id).Body(body).Execute()

Create a new protection run.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := "id_example" // string | Specifies a unique id of the Protection Group.
	body := *openapiclient.NewCreateProtectionGroupRunRequest("RunType_example") // CreateProtectionGroupRunRequest | Specifies the parameters to start a protection run.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProtectionGroupAPI.CreateProtectionGroupRun(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroupAPI.CreateProtectionGroupRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProtectionGroupRun`: CreateProtectionGroupRunResponseBody
	fmt.Fprintf(os.Stdout, "Response from `ProtectionGroupAPI.CreateProtectionGroupRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the Protection Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProtectionGroupRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateProtectionGroupRunRequest**](CreateProtectionGroupRunRequest.md) | Specifies the parameters to start a protection run. | 

### Return type

[**CreateProtectionGroupRunResponseBody**](CreateProtectionGroupRunResponseBody.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProtectionGroup

> DeleteProtectionGroup(ctx, id).DeleteSnapshots(deleteSnapshots).Execute()

Delete a Protection Group.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := "id_example" // string | Specifies a unique id of the Protection Group.
	deleteSnapshots := true // bool | Specifies if Snapshots generated by the Protection Group should also be deleted when the Protection Group is deleted. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProtectionGroupAPI.DeleteProtectionGroup(context.Background(), id).DeleteSnapshots(deleteSnapshots).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroupAPI.DeleteProtectionGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the Protection Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProtectionGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteSnapshots** | **bool** | Specifies if Snapshots generated by the Protection Group should also be deleted when the Protection Group is deleted. | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtectionGroupById

> ProtectionGroup GetProtectionGroupById(ctx, id).RequestInitiatorType(requestInitiatorType).IncludeLastRunInfo(includeLastRunInfo).PruneExcludedSourceIds(pruneExcludedSourceIds).PruneSourceIds(pruneSourceIds).Execute()

List details about single Protection Group.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := "id_example" // string | Specifies a unique id of the Protection Group.
	requestInitiatorType := "requestInitiatorType_example" // string | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. (optional)
	includeLastRunInfo := true // bool | If true, the response will include last run info. If it is false or not specified, the last run info won't be returned. (optional)
	pruneExcludedSourceIds := true // bool | If true, the response will not include the list of excluded source IDs in groups that contain this field. This can be set to true in order to improve performance if excluded source IDs are not needed by the user. (optional)
	pruneSourceIds := true // bool | If true, the response will exclude the list of source IDs within the group specified. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProtectionGroupAPI.GetProtectionGroupById(context.Background(), id).RequestInitiatorType(requestInitiatorType).IncludeLastRunInfo(includeLastRunInfo).PruneExcludedSourceIds(pruneExcludedSourceIds).PruneSourceIds(pruneSourceIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroupAPI.GetProtectionGroupById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProtectionGroupById`: ProtectionGroup
	fmt.Fprintf(os.Stdout, "Response from `ProtectionGroupAPI.GetProtectionGroupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the Protection Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionGroupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestInitiatorType** | **string** | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | 
 **includeLastRunInfo** | **bool** | If true, the response will include last run info. If it is false or not specified, the last run info won&#39;t be returned. | 
 **pruneExcludedSourceIds** | **bool** | If true, the response will not include the list of excluded source IDs in groups that contain this field. This can be set to true in order to improve performance if excluded source IDs are not needed by the user. | 
 **pruneSourceIds** | **bool** | If true, the response will exclude the list of source IDs within the group specified. | 

### Return type

[**ProtectionGroup**](ProtectionGroup.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtectionGroupRun

> ProtectionGroupRun GetProtectionGroupRun(ctx, id, runId).RequestInitiatorType(requestInitiatorType).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeObjectDetails(includeObjectDetails).UseCachedData(useCachedData).Execute()

Get a run for a Protection Group.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := "id_example" // string | Specifies a unique id of the Protection Group.
	runId := "runId_example" // string | Specifies a unique run id of the Protection Group run.
	requestInitiatorType := "requestInitiatorType_example" // string | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. (optional)
	tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which the run is to be returned. (optional)
	includeTenants := true // bool | If true, the response will include Protection Group Runs which were created by all tenants which the current user has permission to see. If false, then only Protection Groups created by the current user will be returned. If it's not specified, it is true by default. (optional)
	includeObjectDetails := true // bool | Specifies if the result includes the object details for a protection run. If set to true, details of the protected object will be returned. If set to false or not specified, details will not be returned. (optional)
	useCachedData := true // bool | Specifies whether we can serve the GET request from the read replica cache. There is a lag of 15 seconds between the read replica and primary data source. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProtectionGroupAPI.GetProtectionGroupRun(context.Background(), id, runId).RequestInitiatorType(requestInitiatorType).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeObjectDetails(includeObjectDetails).UseCachedData(useCachedData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroupAPI.GetProtectionGroupRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProtectionGroupRun`: ProtectionGroupRun
	fmt.Fprintf(os.Stdout, "Response from `ProtectionGroupAPI.GetProtectionGroupRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the Protection Group. | 
**runId** | **string** | Specifies a unique run id of the Protection Group run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionGroupRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestInitiatorType** | **string** | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which the run is to be returned. | 
 **includeTenants** | **bool** | If true, the response will include Protection Group Runs which were created by all tenants which the current user has permission to see. If false, then only Protection Groups created by the current user will be returned. If it&#39;s not specified, it is true by default. | 
 **includeObjectDetails** | **bool** | Specifies if the result includes the object details for a protection run. If set to true, details of the protected object will be returned. If set to false or not specified, details will not be returned. | 
 **useCachedData** | **bool** | Specifies whether we can serve the GET request from the read replica cache. There is a lag of 15 seconds between the read replica and primary data source. | 

### Return type

[**ProtectionGroupRun**](ProtectionGroupRun.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtectionGroupRuns

> ProtectionGroupRuns GetProtectionGroupRuns(ctx, id).RequestInitiatorType(requestInitiatorType).RunId(runId).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).TenantIds(tenantIds).IncludeTenants(includeTenants).RunTypes(runTypes).IncludeObjectDetails(includeObjectDetails).LocalBackupRunStatus(localBackupRunStatus).ReplicationRunStatus(replicationRunStatus).ArchivalRunStatus(archivalRunStatus).CloudSpinRunStatus(cloudSpinRunStatus).NumRuns(numRuns).ExcludeNonRestorableRuns(excludeNonRestorableRuns).RunTags(runTags).UseCachedData(useCachedData).FilterByEndTime(filterByEndTime).SnapshotTargetTypes(snapshotTargetTypes).OnlyReturnSuccessfulCopyRun(onlyReturnSuccessfulCopyRun).FilterByCopyTaskEndTime(filterByCopyTaskEndTime).TruncateResponse(truncateResponse).OnlyReturnShellInfo(onlyReturnShellInfo).ExcludeErrorRuns(excludeErrorRuns).JobRunStartTimeUsecs(jobRunStartTimeUsecs).OnlyReturnDataMigrationJobs(onlyReturnDataMigrationJobs).IncludeExtensionInfo(includeExtensionInfo).IncludeRpoSnapshots(includeRpoSnapshots).SourceId(sourceId).Execute()

Get the list of runs for a Protection Group.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := "id_example" // string | Specifies a unique id of the Protection Group.
	requestInitiatorType := "requestInitiatorType_example" // string | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. (optional)
	runId := "runId_example" // string | Specifies the protection run id. (optional)
	startTimeUsecs := int64(789) // int64 | Start time for time range filter. Specify the start time as a Unix epoch Timestamp (in microseconds), only runs executing after this time will be returned. By default it is endTimeUsecs minus an hour. (optional)
	endTimeUsecs := int64(789) // int64 | End time for time range filter. Specify the end time as a Unix epoch Timestamp (in microseconds), only runs executing before this time will be returned. By default it is current time. (optional)
	tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
	includeTenants := true // bool | If true, the response will include Protection Group Runs which were created by all tenants which the current user has permission to see. If false, then only Protection Group Runs created by the current user will be returned. (optional)
	runTypes := []string{"RunTypes_example"} // []string | Filter by run type. Only protection run matching the specified types will be returned. (optional)
	includeObjectDetails := true // bool | Specifies if the result includes the object details for each protection run. If set to true, details of the protected object will be returned. If set to false or not specified, details will not be returned. (optional)
	localBackupRunStatus := []string{"LocalBackupRunStatus_example"} // []string | Specifies a list of local backup status, runs matching the status will be returned.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped. (optional)
	replicationRunStatus := []string{"ReplicationRunStatus_example"} // []string | Specifies a list of replication status, runs matching the status will be returned.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped. (optional)
	archivalRunStatus := []string{"ArchivalRunStatus_example"} // []string | Specifies a list of archival status, runs matching the status will be returned.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped. (optional)
	cloudSpinRunStatus := []string{"CloudSpinRunStatus_example"} // []string | Specifies a list of cloud spin status, runs matching the status will be returned.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped. (optional)
	numRuns := int64(789) // int64 | Specifies the max number of runs. If not specified, at most 100 runs will be returned. (optional)
	excludeNonRestorableRuns := true // bool | Specifies whether to exclude non restorable runs. Run is treated restorable only if there is atleast one object snapshot (which may be either a local or an archival snapshot) which is not deleted or expired. Default value is false. (optional) (default to false)
	runTags := []string{"Inner_example"} // []string | Specifies a list of tags for protection runs. If this is specified, only the runs which match these tags will be returned. (optional)
	useCachedData := true // bool | Specifies whether we can serve the GET request from the read replica cache. There is a lag of 15 seconds between the read replica and primary data source. (optional)
	filterByEndTime := true // bool | If true, the runs with backup end time within the specified time range will be returned. Otherwise, the runs with start time in the time range are returned. (optional)
	snapshotTargetTypes := []string{"SnapshotTargetTypes_example"} // []string | Specifies the snapshot's target type which should be filtered. Note: this field is only considered when, filterByCopyTaskEndTime is set to true, or else it is ignored. (optional)
	onlyReturnSuccessfulCopyRun := true // bool | If set to false, all copy_tasks in any given valid state will be considered. If left empty or set to true, only successful copy_tasks would be considered. Note: this field is only considered when, filterByCopyTaskEndTime is set to true, or else it is ignored. (optional)
	filterByCopyTaskEndTime := true // bool | If true, then the details of the runs for which any copyTask completed in the given timerange will be returned. Only one of filterByEndTime and filterByCopyTaskEndTime can be set. (optional)
	truncateResponse := true // bool | If set, magneto will truncate the response if it exceeds max size limit governed by magneto_http_rpc_response_size_limit_bytes (optional)
	onlyReturnShellInfo := true // bool | If set, returns only shell info such as run's start time, type, error if any. (optional)
	excludeErrorRuns := true // bool | Specifies whether to exclude runs with error. If no value is specified, then runs with errors are included. (optional)
	jobRunStartTimeUsecs := int64(789) // int64 | Return a specific Job Run by specifying a time and a group id. Specify the time when the Job Run started as a Unix epoch Timestamp (in microseconds). If this field is specified, jobId must also be specified. (optional)
	onlyReturnDataMigrationJobs := true // bool | Specifies if only data stubbing jobs should be returned. If not set, no data migration job will be returned. (optional)
	includeExtensionInfo := true // bool | Specifies if needs to include proto extensions if they are extended. (optional)
	includeRpoSnapshots := true // bool | If true, then the snapshots for Protection Sources protected by Rpo policies will also be returned. (optional)
	sourceId := int64(789) // int64 | Filter by source id. Only Job Runs protecting the specified source (such as a VM or View) are returned. The source id is assigned by the Cohesity Cluster. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProtectionGroupAPI.GetProtectionGroupRuns(context.Background(), id).RequestInitiatorType(requestInitiatorType).RunId(runId).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).TenantIds(tenantIds).IncludeTenants(includeTenants).RunTypes(runTypes).IncludeObjectDetails(includeObjectDetails).LocalBackupRunStatus(localBackupRunStatus).ReplicationRunStatus(replicationRunStatus).ArchivalRunStatus(archivalRunStatus).CloudSpinRunStatus(cloudSpinRunStatus).NumRuns(numRuns).ExcludeNonRestorableRuns(excludeNonRestorableRuns).RunTags(runTags).UseCachedData(useCachedData).FilterByEndTime(filterByEndTime).SnapshotTargetTypes(snapshotTargetTypes).OnlyReturnSuccessfulCopyRun(onlyReturnSuccessfulCopyRun).FilterByCopyTaskEndTime(filterByCopyTaskEndTime).TruncateResponse(truncateResponse).OnlyReturnShellInfo(onlyReturnShellInfo).ExcludeErrorRuns(excludeErrorRuns).JobRunStartTimeUsecs(jobRunStartTimeUsecs).OnlyReturnDataMigrationJobs(onlyReturnDataMigrationJobs).IncludeExtensionInfo(includeExtensionInfo).IncludeRpoSnapshots(includeRpoSnapshots).SourceId(sourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroupAPI.GetProtectionGroupRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProtectionGroupRuns`: ProtectionGroupRuns
	fmt.Fprintf(os.Stdout, "Response from `ProtectionGroupAPI.GetProtectionGroupRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the Protection Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionGroupRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestInitiatorType** | **string** | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | 
 **runId** | **string** | Specifies the protection run id. | 
 **startTimeUsecs** | **int64** | Start time for time range filter. Specify the start time as a Unix epoch Timestamp (in microseconds), only runs executing after this time will be returned. By default it is endTimeUsecs minus an hour. | 
 **endTimeUsecs** | **int64** | End time for time range filter. Specify the end time as a Unix epoch Timestamp (in microseconds), only runs executing before this time will be returned. By default it is current time. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **includeTenants** | **bool** | If true, the response will include Protection Group Runs which were created by all tenants which the current user has permission to see. If false, then only Protection Group Runs created by the current user will be returned. | 
 **runTypes** | **[]string** | Filter by run type. Only protection run matching the specified types will be returned. | 
 **includeObjectDetails** | **bool** | Specifies if the result includes the object details for each protection run. If set to true, details of the protected object will be returned. If set to false or not specified, details will not be returned. | 
 **localBackupRunStatus** | **[]string** | Specifies a list of local backup status, runs matching the status will be returned.&lt;br&gt; &#39;Running&#39; indicates that the run is still running.&lt;br&gt; &#39;Canceled&#39; indicates that the run has been canceled.&lt;br&gt; &#39;Canceling&#39; indicates that the run is in the process of being canceled.&lt;br&gt; &#39;Failed&#39; indicates that the run has failed.&lt;br&gt; &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening.&lt;br&gt; &#39;Succeeded&#39; indicates that the run has finished successfully.&lt;br&gt; &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages.&lt;br&gt; &#39;Paused&#39; indicates that the ongoing run has been paused.&lt;br&gt; &#39;Skipped&#39; indicates that the run was skipped. | 
 **replicationRunStatus** | **[]string** | Specifies a list of replication status, runs matching the status will be returned.&lt;br&gt; &#39;Running&#39; indicates that the run is still running.&lt;br&gt; &#39;Canceled&#39; indicates that the run has been canceled.&lt;br&gt; &#39;Canceling&#39; indicates that the run is in the process of being canceled.&lt;br&gt; &#39;Failed&#39; indicates that the run has failed.&lt;br&gt; &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening.&lt;br&gt; &#39;Succeeded&#39; indicates that the run has finished successfully.&lt;br&gt; &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages.&lt;br&gt; &#39;Paused&#39; indicates that the ongoing run has been paused.&lt;br&gt; &#39;Skipped&#39; indicates that the run was skipped. | 
 **archivalRunStatus** | **[]string** | Specifies a list of archival status, runs matching the status will be returned.&lt;br&gt; &#39;Running&#39; indicates that the run is still running.&lt;br&gt; &#39;Canceled&#39; indicates that the run has been canceled.&lt;br&gt; &#39;Canceling&#39; indicates that the run is in the process of being canceled.&lt;br&gt; &#39;Failed&#39; indicates that the run has failed.&lt;br&gt; &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening.&lt;br&gt; &#39;Succeeded&#39; indicates that the run has finished successfully.&lt;br&gt; &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages.&lt;br&gt; &#39;Paused&#39; indicates that the ongoing run has been paused.&lt;br&gt; &#39;Skipped&#39; indicates that the run was skipped. | 
 **cloudSpinRunStatus** | **[]string** | Specifies a list of cloud spin status, runs matching the status will be returned.&lt;br&gt; &#39;Running&#39; indicates that the run is still running.&lt;br&gt; &#39;Canceled&#39; indicates that the run has been canceled.&lt;br&gt; &#39;Canceling&#39; indicates that the run is in the process of being canceled.&lt;br&gt; &#39;Failed&#39; indicates that the run has failed.&lt;br&gt; &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening.&lt;br&gt; &#39;Succeeded&#39; indicates that the run has finished successfully.&lt;br&gt; &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages.&lt;br&gt; &#39;Paused&#39; indicates that the ongoing run has been paused.&lt;br&gt; &#39;Skipped&#39; indicates that the run was skipped. | 
 **numRuns** | **int64** | Specifies the max number of runs. If not specified, at most 100 runs will be returned. | 
 **excludeNonRestorableRuns** | **bool** | Specifies whether to exclude non restorable runs. Run is treated restorable only if there is atleast one object snapshot (which may be either a local or an archival snapshot) which is not deleted or expired. Default value is false. | [default to false]
 **runTags** | **[]string** | Specifies a list of tags for protection runs. If this is specified, only the runs which match these tags will be returned. | 
 **useCachedData** | **bool** | Specifies whether we can serve the GET request from the read replica cache. There is a lag of 15 seconds between the read replica and primary data source. | 
 **filterByEndTime** | **bool** | If true, the runs with backup end time within the specified time range will be returned. Otherwise, the runs with start time in the time range are returned. | 
 **snapshotTargetTypes** | **[]string** | Specifies the snapshot&#39;s target type which should be filtered. Note: this field is only considered when, filterByCopyTaskEndTime is set to true, or else it is ignored. | 
 **onlyReturnSuccessfulCopyRun** | **bool** | If set to false, all copy_tasks in any given valid state will be considered. If left empty or set to true, only successful copy_tasks would be considered. Note: this field is only considered when, filterByCopyTaskEndTime is set to true, or else it is ignored. | 
 **filterByCopyTaskEndTime** | **bool** | If true, then the details of the runs for which any copyTask completed in the given timerange will be returned. Only one of filterByEndTime and filterByCopyTaskEndTime can be set. | 
 **truncateResponse** | **bool** | If set, magneto will truncate the response if it exceeds max size limit governed by magneto_http_rpc_response_size_limit_bytes | 
 **onlyReturnShellInfo** | **bool** | If set, returns only shell info such as run&#39;s start time, type, error if any. | 
 **excludeErrorRuns** | **bool** | Specifies whether to exclude runs with error. If no value is specified, then runs with errors are included. | 
 **jobRunStartTimeUsecs** | **int64** | Return a specific Job Run by specifying a time and a group id. Specify the time when the Job Run started as a Unix epoch Timestamp (in microseconds). If this field is specified, jobId must also be specified. | 
 **onlyReturnDataMigrationJobs** | **bool** | Specifies if only data stubbing jobs should be returned. If not set, no data migration job will be returned. | 
 **includeExtensionInfo** | **bool** | Specifies if needs to include proto extensions if they are extended. | 
 **includeRpoSnapshots** | **bool** | If true, then the snapshots for Protection Sources protected by Rpo policies will also be returned. | 
 **sourceId** | **int64** | Filter by source id. Only Job Runs protecting the specified source (such as a VM or View) are returned. The source id is assigned by the Cohesity Cluster. | 

### Return type

[**ProtectionGroupRuns**](ProtectionGroupRuns.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtectionGroups

> ProtectionGroups GetProtectionGroups(ctx).RequestInitiatorType(requestInitiatorType).Ids(ids).Names(names).PolicyIds(policyIds).StorageDomainId(storageDomainId).IncludeGroupsWithDatalockOnly(includeGroupsWithDatalockOnly).Environments(environments).Office365Workloads(office365Workloads).IsActive(isActive).IsDeleted(isDeleted).IsPaused(isPaused).LastRunLocalBackupStatus(lastRunLocalBackupStatus).LastRunReplicationStatus(lastRunReplicationStatus).LastRunArchivalStatus(lastRunArchivalStatus).LastRunCloudSpinStatus(lastRunCloudSpinStatus).LastRunAnyStatus(lastRunAnyStatus).IsLastRunSlaViolated(isLastRunSlaViolated).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeLastRunInfo(includeLastRunInfo).PruneExcludedSourceIds(pruneExcludedSourceIds).PruneSourceIds(pruneSourceIds).UseCachedData(useCachedData).SourceIds(sourceIds).Execute()

Get the list of Protection Groups.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	requestInitiatorType := "requestInitiatorType_example" // string | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. (optional)
	ids := []string{"Inner_example"} // []string | Filter by a list of Protection Group ids. (optional)
	names := []string{"Inner_example"} // []string | Filter by a list of Protection Group names. (optional)
	policyIds := []string{"Inner_example"} // []string | Filter by Policy ids that are associated with Protection Groups. Only Protection Groups associated with the specified Policy ids, are returned. (optional)
	storageDomainId := int64(789) // int64 | Filter by Storage Domain id. Only Protection Groups writing data to this Storage Domain will be returned. (optional)
	includeGroupsWithDatalockOnly := true // bool | Whether to only return Protection Groups with a datalock. (optional)
	environments := []string{"Environments_example"} // []string | Filter by environment types such as 'kVMware', 'kView', etc. Only Protection Groups protecting the specified environment types are returned. (optional)
	office365Workloads := []string{"Office365Workloads_example"} // []string |  (optional)
	isActive := true // bool | Filter by Inactive or Active Protection Groups. If not set, all Inactive and Active Protection Groups are returned. If true, only Active Protection Groups are returned. If false, only Inactive Protection Groups are returned. When you create a Protection Group on a Primary Cluster with a replication schedule, the Cluster creates an Inactive copy of the Protection Group on the Remote Cluster. In addition, when an Active and running Protection Group is deactivated, the Protection Group becomes Inactive. (optional)
	isDeleted := true // bool | If true, return only Protection Groups that have been deleted but still have Snapshots associated with them. If false, return all Protection Groups except those Protection Groups that have been deleted and still have Snapshots associated with them. A Protection Group that is deleted with all its Snapshots is not returned for either of these cases. (optional)
	isPaused := true // bool | Filter by paused or non paused Protection Groups, If not set, all paused and non paused Protection Groups are returned. If true, only paused Protection Groups are returned. If false, only non paused Protection Groups are returned. (optional)
	lastRunLocalBackupStatus := []string{"LastRunLocalBackupStatus_example"} // []string | Filter by last local backup run status.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped. (optional)
	lastRunReplicationStatus := []string{"LastRunReplicationStatus_example"} // []string | Filter by last remote replication run status.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped. (optional)
	lastRunArchivalStatus := []string{"LastRunArchivalStatus_example"} // []string | Filter by last cloud archival run status.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped. (optional)
	lastRunCloudSpinStatus := []string{"LastRunCloudSpinStatus_example"} // []string | Filter by last cloud spin run status.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped. (optional)
	lastRunAnyStatus := []string{"LastRunAnyStatus_example"} // []string | Filter by last any run status.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped. (optional)
	isLastRunSlaViolated := true // bool | If true, return Protection Groups for which last run SLA was violated. (optional)
	tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
	includeTenants := true // bool | If true, the response will include Protection Groups which were created by all tenants which the current user has permission to see. If false, then only Protection Groups created by the current user will be returned. (optional)
	includeLastRunInfo := true // bool | If true, the response will include last run info. If it is false or not specified, the last run info won't be returned. (optional)
	pruneExcludedSourceIds := true // bool | If true, the response will not include the list of excluded source IDs in groups that contain this field. This can be set to true in order to improve performance if excluded source IDs are not needed by the user. (optional)
	pruneSourceIds := true // bool | If true, the response will exclude the list of source IDs within the group specified. (optional)
	useCachedData := true // bool | Specifies whether we can serve the GET request from the read replica cache. There is a lag of 15 seconds between the read replica and primary data source. (optional)
	sourceIds := []int64{int64(123)} // []int64 | Filter by Source ids that are associated with Protection Groups. Only Protection Groups associated with the specified Source ids, are returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProtectionGroupAPI.GetProtectionGroups(context.Background()).RequestInitiatorType(requestInitiatorType).Ids(ids).Names(names).PolicyIds(policyIds).StorageDomainId(storageDomainId).IncludeGroupsWithDatalockOnly(includeGroupsWithDatalockOnly).Environments(environments).Office365Workloads(office365Workloads).IsActive(isActive).IsDeleted(isDeleted).IsPaused(isPaused).LastRunLocalBackupStatus(lastRunLocalBackupStatus).LastRunReplicationStatus(lastRunReplicationStatus).LastRunArchivalStatus(lastRunArchivalStatus).LastRunCloudSpinStatus(lastRunCloudSpinStatus).LastRunAnyStatus(lastRunAnyStatus).IsLastRunSlaViolated(isLastRunSlaViolated).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeLastRunInfo(includeLastRunInfo).PruneExcludedSourceIds(pruneExcludedSourceIds).PruneSourceIds(pruneSourceIds).UseCachedData(useCachedData).SourceIds(sourceIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroupAPI.GetProtectionGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProtectionGroups`: ProtectionGroups
	fmt.Fprintf(os.Stdout, "Response from `ProtectionGroupAPI.GetProtectionGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestInitiatorType** | **string** | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | 
 **ids** | **[]string** | Filter by a list of Protection Group ids. | 
 **names** | **[]string** | Filter by a list of Protection Group names. | 
 **policyIds** | **[]string** | Filter by Policy ids that are associated with Protection Groups. Only Protection Groups associated with the specified Policy ids, are returned. | 
 **storageDomainId** | **int64** | Filter by Storage Domain id. Only Protection Groups writing data to this Storage Domain will be returned. | 
 **includeGroupsWithDatalockOnly** | **bool** | Whether to only return Protection Groups with a datalock. | 
 **environments** | **[]string** | Filter by environment types such as &#39;kVMware&#39;, &#39;kView&#39;, etc. Only Protection Groups protecting the specified environment types are returned. | 
 **office365Workloads** | **[]string** |  | 
 **isActive** | **bool** | Filter by Inactive or Active Protection Groups. If not set, all Inactive and Active Protection Groups are returned. If true, only Active Protection Groups are returned. If false, only Inactive Protection Groups are returned. When you create a Protection Group on a Primary Cluster with a replication schedule, the Cluster creates an Inactive copy of the Protection Group on the Remote Cluster. In addition, when an Active and running Protection Group is deactivated, the Protection Group becomes Inactive. | 
 **isDeleted** | **bool** | If true, return only Protection Groups that have been deleted but still have Snapshots associated with them. If false, return all Protection Groups except those Protection Groups that have been deleted and still have Snapshots associated with them. A Protection Group that is deleted with all its Snapshots is not returned for either of these cases. | 
 **isPaused** | **bool** | Filter by paused or non paused Protection Groups, If not set, all paused and non paused Protection Groups are returned. If true, only paused Protection Groups are returned. If false, only non paused Protection Groups are returned. | 
 **lastRunLocalBackupStatus** | **[]string** | Filter by last local backup run status.&lt;br&gt; &#39;Running&#39; indicates that the run is still running.&lt;br&gt; &#39;Canceled&#39; indicates that the run has been canceled.&lt;br&gt; &#39;Canceling&#39; indicates that the run is in the process of being canceled.&lt;br&gt; &#39;Failed&#39; indicates that the run has failed.&lt;br&gt; &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening.&lt;br&gt; &#39;Succeeded&#39; indicates that the run has finished successfully.&lt;br&gt; &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages.&lt;br&gt; &#39;Paused&#39; indicates that the ongoing run has been paused.&lt;br&gt; &#39;Skipped&#39; indicates that the run was skipped. | 
 **lastRunReplicationStatus** | **[]string** | Filter by last remote replication run status.&lt;br&gt; &#39;Running&#39; indicates that the run is still running.&lt;br&gt; &#39;Canceled&#39; indicates that the run has been canceled.&lt;br&gt; &#39;Canceling&#39; indicates that the run is in the process of being canceled.&lt;br&gt; &#39;Failed&#39; indicates that the run has failed.&lt;br&gt; &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening.&lt;br&gt; &#39;Succeeded&#39; indicates that the run has finished successfully.&lt;br&gt; &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages.&lt;br&gt; &#39;Paused&#39; indicates that the ongoing run has been paused.&lt;br&gt; &#39;Skipped&#39; indicates that the run was skipped. | 
 **lastRunArchivalStatus** | **[]string** | Filter by last cloud archival run status.&lt;br&gt; &#39;Running&#39; indicates that the run is still running.&lt;br&gt; &#39;Canceled&#39; indicates that the run has been canceled.&lt;br&gt; &#39;Canceling&#39; indicates that the run is in the process of being canceled.&lt;br&gt; &#39;Failed&#39; indicates that the run has failed.&lt;br&gt; &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening.&lt;br&gt; &#39;Succeeded&#39; indicates that the run has finished successfully.&lt;br&gt; &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages.&lt;br&gt; &#39;Paused&#39; indicates that the ongoing run has been paused.&lt;br&gt; &#39;Skipped&#39; indicates that the run was skipped. | 
 **lastRunCloudSpinStatus** | **[]string** | Filter by last cloud spin run status.&lt;br&gt; &#39;Running&#39; indicates that the run is still running.&lt;br&gt; &#39;Canceled&#39; indicates that the run has been canceled.&lt;br&gt; &#39;Canceling&#39; indicates that the run is in the process of being canceled.&lt;br&gt; &#39;Failed&#39; indicates that the run has failed.&lt;br&gt; &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening.&lt;br&gt; &#39;Succeeded&#39; indicates that the run has finished successfully.&lt;br&gt; &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages.&lt;br&gt; &#39;Paused&#39; indicates that the ongoing run has been paused.&lt;br&gt; &#39;Skipped&#39; indicates that the run was skipped. | 
 **lastRunAnyStatus** | **[]string** | Filter by last any run status.&lt;br&gt; &#39;Running&#39; indicates that the run is still running.&lt;br&gt; &#39;Canceled&#39; indicates that the run has been canceled.&lt;br&gt; &#39;Canceling&#39; indicates that the run is in the process of being canceled.&lt;br&gt; &#39;Failed&#39; indicates that the run has failed.&lt;br&gt; &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening.&lt;br&gt; &#39;Succeeded&#39; indicates that the run has finished successfully.&lt;br&gt; &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages.&lt;br&gt; &#39;Paused&#39; indicates that the ongoing run has been paused.&lt;br&gt; &#39;Skipped&#39; indicates that the run was skipped. | 
 **isLastRunSlaViolated** | **bool** | If true, return Protection Groups for which last run SLA was violated. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **includeTenants** | **bool** | If true, the response will include Protection Groups which were created by all tenants which the current user has permission to see. If false, then only Protection Groups created by the current user will be returned. | 
 **includeLastRunInfo** | **bool** | If true, the response will include last run info. If it is false or not specified, the last run info won&#39;t be returned. | 
 **pruneExcludedSourceIds** | **bool** | If true, the response will not include the list of excluded source IDs in groups that contain this field. This can be set to true in order to improve performance if excluded source IDs are not needed by the user. | 
 **pruneSourceIds** | **bool** | If true, the response will exclude the list of source IDs within the group specified. | 
 **useCachedData** | **bool** | Specifies whether we can serve the GET request from the read replica cache. There is a lag of 15 seconds between the read replica and primary data source. | 
 **sourceIds** | **[]int64** | Filter by Source ids that are associated with Protection Groups. Only Protection Groups associated with the specified Source ids, are returned. | 

### Return type

[**ProtectionGroups**](ProtectionGroups.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtectionRunProgress

> GetProtectionRunProgressBody GetProtectionRunProgress(ctx, runId).Objects(objects).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeFinishedTasks(includeFinishedTasks).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).MaxTasksNum(maxTasksNum).ExcludeObjectDetails(excludeObjectDetails).IncludeEventLogs(includeEventLogs).MaxLogLevel(maxLogLevel).RunTaskPath(runTaskPath).ObjectTaskPaths(objectTaskPaths).Execute()

Get the progress of a run.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	runId := "runId_example" // string | Specifies a unique run id of the Protection Run.
	objects := []int64{int64(123)} // []int64 | Specifies the objects whose progress will be returned. This only applies to protection group runs and will be ignored for object runs. If the objects are specified, the run progress will not be returned and only the progress of the specified objects will be returned. (optional)
	tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which the run is to be returned. (optional)
	includeTenants := true // bool | If true, the response will include Protection Group Runs which were created by all tenants which the current user has permission to see. If false, then only Protection Groups created by the current user will be returned. If it's not specified, it is true by default. (optional)
	includeFinishedTasks := true // bool | Specifies whether to return finished tasks. By default only active tasks are returned. (optional)
	startTimeUsecs := int64(789) // int64 | Specifies the time after which the progress task starts in Unix epoch Timestamp(in microseconds). (optional)
	endTimeUsecs := int64(789) // int64 | Specifies the time before which the progress task ends in Unix epoch Timestamp(in microseconds). (optional)
	maxTasksNum := int32(56) // int32 | Specifies the maximum number of tasks to return. (optional)
	excludeObjectDetails := true // bool | Specifies whether to return objects. By default all the task tree are returned. (optional)
	includeEventLogs := true // bool | Specifies whether to include event logs (optional)
	maxLogLevel := int32(56) // int32 | Specifies the number of levels till which to fetch the event logs. This is applicable only when includeEventLogs is true. (optional)
	runTaskPath := "runTaskPath_example" // string | Specifies the task path of the run or object run. This is applicable only if progress of a protection group with one or more object is required.If provided this will be used to fetch progress details directly without looking actual task path of the object. Objects field is stil expected else it changes the response format. (optional)
	objectTaskPaths := []string{"Inner_example"} // []string | Specifies the object level task path. This relates to the objectID. If provided this will take precedence over the objects, and will be used to fetch progress details directly without looking actuall task path of the object. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProtectionGroupAPI.GetProtectionRunProgress(context.Background(), runId).Objects(objects).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeFinishedTasks(includeFinishedTasks).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).MaxTasksNum(maxTasksNum).ExcludeObjectDetails(excludeObjectDetails).IncludeEventLogs(includeEventLogs).MaxLogLevel(maxLogLevel).RunTaskPath(runTaskPath).ObjectTaskPaths(objectTaskPaths).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroupAPI.GetProtectionRunProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProtectionRunProgress`: GetProtectionRunProgressBody
	fmt.Fprintf(os.Stdout, "Response from `ProtectionGroupAPI.GetProtectionRunProgress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Specifies a unique run id of the Protection Run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionRunProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objects** | **[]int64** | Specifies the objects whose progress will be returned. This only applies to protection group runs and will be ignored for object runs. If the objects are specified, the run progress will not be returned and only the progress of the specified objects will be returned. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which the run is to be returned. | 
 **includeTenants** | **bool** | If true, the response will include Protection Group Runs which were created by all tenants which the current user has permission to see. If false, then only Protection Groups created by the current user will be returned. If it&#39;s not specified, it is true by default. | 
 **includeFinishedTasks** | **bool** | Specifies whether to return finished tasks. By default only active tasks are returned. | 
 **startTimeUsecs** | **int64** | Specifies the time after which the progress task starts in Unix epoch Timestamp(in microseconds). | 
 **endTimeUsecs** | **int64** | Specifies the time before which the progress task ends in Unix epoch Timestamp(in microseconds). | 
 **maxTasksNum** | **int32** | Specifies the maximum number of tasks to return. | 
 **excludeObjectDetails** | **bool** | Specifies whether to return objects. By default all the task tree are returned. | 
 **includeEventLogs** | **bool** | Specifies whether to include event logs | 
 **maxLogLevel** | **int32** | Specifies the number of levels till which to fetch the event logs. This is applicable only when includeEventLogs is true. | 
 **runTaskPath** | **string** | Specifies the task path of the run or object run. This is applicable only if progress of a protection group with one or more object is required.If provided this will be used to fetch progress details directly without looking actual task path of the object. Objects field is stil expected else it changes the response format. | 
 **objectTaskPaths** | **[]string** | Specifies the object level task path. This relates to the objectID. If provided this will take precedence over the objects, and will be used to fetch progress details directly without looking actuall task path of the object. | 

### Return type

[**GetProtectionRunProgressBody**](GetProtectionRunProgressBody.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtectionRunStats

> GetProtectionRunStatsBody GetProtectionRunStats(ctx, runId).Objects(objects).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeFinishedTasks(includeFinishedTasks).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).MaxTasksNum(maxTasksNum).ExcludeObjectDetails(excludeObjectDetails).RunTaskPath(runTaskPath).ObjectTaskPaths(objectTaskPaths).Execute()

Get the stats for a run.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	runId := "runId_example" // string | Specifies a unique run id of the Protection Run.
	objects := []int64{int64(123)} // []int64 | Specifies the objects whose stats will be returned. This only applies to protection group runs and will be ignored for object runs. If the objects are specified, the run stats will not be returned and only the stats of the specified objects will be returned. (optional)
	tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which the run is to be returned. (optional)
	includeTenants := true // bool | If true, the response will include Protection Group Runs which were created by all tenants which the current user has permission to see. If false, then only Protection Groups created by the current user will be returned. If it's not specified, it is true by default. (optional)
	includeFinishedTasks := true // bool | Specifies whether to return finished tasks. By default only active tasks are returned. (optional)
	startTimeUsecs := int64(789) // int64 | Specifies the time after which the stats task starts in Unix epoch Timestamp(in microseconds). (optional)
	endTimeUsecs := int64(789) // int64 | Specifies the time before which the stats task ends in Unix epoch Timestamp(in microseconds). (optional)
	maxTasksNum := int32(56) // int32 | Specifies the maximum number of tasks to return. (optional)
	excludeObjectDetails := true // bool | Specifies whether to return objects. By default all the task tree are returned. (optional)
	runTaskPath := "runTaskPath_example" // string | Specifies the task path of the run or object run. This is applicable only if stats of a protection group with one or more object is required. If provided this will be used to fetch stats details directly without looking actual task path of the object. Objects field is stil expected else it changes the response format. (optional)
	objectTaskPaths := []string{"Inner_example"} // []string | Specifies the object level task path. This relates to the objectID. If provided this will take precedence over the objects, and will be used to fetch stats details directly without looking actuall task path of the object. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProtectionGroupAPI.GetProtectionRunStats(context.Background(), runId).Objects(objects).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeFinishedTasks(includeFinishedTasks).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).MaxTasksNum(maxTasksNum).ExcludeObjectDetails(excludeObjectDetails).RunTaskPath(runTaskPath).ObjectTaskPaths(objectTaskPaths).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroupAPI.GetProtectionRunStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProtectionRunStats`: GetProtectionRunStatsBody
	fmt.Fprintf(os.Stdout, "Response from `ProtectionGroupAPI.GetProtectionRunStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Specifies a unique run id of the Protection Run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionRunStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objects** | **[]int64** | Specifies the objects whose stats will be returned. This only applies to protection group runs and will be ignored for object runs. If the objects are specified, the run stats will not be returned and only the stats of the specified objects will be returned. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which the run is to be returned. | 
 **includeTenants** | **bool** | If true, the response will include Protection Group Runs which were created by all tenants which the current user has permission to see. If false, then only Protection Groups created by the current user will be returned. If it&#39;s not specified, it is true by default. | 
 **includeFinishedTasks** | **bool** | Specifies whether to return finished tasks. By default only active tasks are returned. | 
 **startTimeUsecs** | **int64** | Specifies the time after which the stats task starts in Unix epoch Timestamp(in microseconds). | 
 **endTimeUsecs** | **int64** | Specifies the time before which the stats task ends in Unix epoch Timestamp(in microseconds). | 
 **maxTasksNum** | **int32** | Specifies the maximum number of tasks to return. | 
 **excludeObjectDetails** | **bool** | Specifies whether to return objects. By default all the task tree are returned. | 
 **runTaskPath** | **string** | Specifies the task path of the run or object run. This is applicable only if stats of a protection group with one or more object is required. If provided this will be used to fetch stats details directly without looking actual task path of the object. Objects field is stil expected else it changes the response format. | 
 **objectTaskPaths** | **[]string** | Specifies the object level task path. This relates to the objectID. If provided this will take precedence over the objects, and will be used to fetch stats details directly without looking actuall task path of the object. | 

### Return type

[**GetProtectionRunStatsBody**](GetProtectionRunStatsBody.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtectionRuns

> ProtectionRunsSummary GetProtectionRuns(ctx).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).RunStatus(runStatus).Execute()

Get the list of runs.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	startTimeUsecs := int64(789) // int64 | Start time for time range filter. Specify the start time as a Unix epoch Timestamp (in microseconds), only runs executing after this time will be returned. By default it is endTimeUsecs minus an hour. (optional)
	endTimeUsecs := int64(789) // int64 | End time for time range filter. Specify the end time as a Unix epoch Timestamp (in microseconds), only runs executing before this time will be returned. By default it is current time. (optional)
	runStatus := []string{"RunStatus_example"} // []string | Specifies a list of status, runs matching the status will be returned.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Skipped' indicates that the run was skipped. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProtectionGroupAPI.GetProtectionRuns(context.Background()).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).RunStatus(runStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroupAPI.GetProtectionRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProtectionRuns`: ProtectionRunsSummary
	fmt.Fprintf(os.Stdout, "Response from `ProtectionGroupAPI.GetProtectionRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTimeUsecs** | **int64** | Start time for time range filter. Specify the start time as a Unix epoch Timestamp (in microseconds), only runs executing after this time will be returned. By default it is endTimeUsecs minus an hour. | 
 **endTimeUsecs** | **int64** | End time for time range filter. Specify the end time as a Unix epoch Timestamp (in microseconds), only runs executing before this time will be returned. By default it is current time. | 
 **runStatus** | **[]string** | Specifies a list of status, runs matching the status will be returned.&lt;br&gt; &#39;Running&#39; indicates that the run is still running.&lt;br&gt; &#39;Canceled&#39; indicates that the run has been canceled.&lt;br&gt; &#39;Canceling&#39; indicates that the run is in the process of being canceled.&lt;br&gt; &#39;Failed&#39; indicates that the run has failed.&lt;br&gt; &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening.&lt;br&gt; &#39;Succeeded&#39; indicates that the run has finished successfully.&lt;br&gt; &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages.&lt;br&gt; &#39;Skipped&#39; indicates that the run was skipped. | 

### Return type

[**ProtectionRunsSummary**](ProtectionRunsSummary.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRunDebugLogs

> GetRunDebugLogs(ctx, id, runId).ObjectId(objectId).Execute()

Get the debug logs for a run from a Protection Group.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := "id_example" // string | Specifies a unique id of the Protection Group.
	runId := "runId_example" // string | Specifies a unique run id of the Protection Group run.
	objectId := "objectId_example" // string | Specifies the id of the object for which debug logs are to be returned.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProtectionGroupAPI.GetRunDebugLogs(context.Background(), id, runId).ObjectId(objectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroupAPI.GetRunDebugLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the Protection Group. | 
**runId** | **string** | Specifies a unique run id of the Protection Group run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRunDebugLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **objectId** | **string** | Specifies the id of the object for which debug logs are to be returned.  | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRunDebugLogsForObject

> GetRunDebugLogsForObject(ctx, id, runId, objectId).Execute()

Get the debug logs for a particular object in a run from a Protection Group.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := "id_example" // string | Specifies a unique id of the Protection Group.
	runId := "runId_example" // string | Specifies a unique run id of the Protection Group run.
	objectId := "objectId_example" // string | Specifies the id of the object for which debug logs are to be returned. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProtectionGroupAPI.GetRunDebugLogsForObject(context.Background(), id, runId, objectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroupAPI.GetRunDebugLogsForObject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the Protection Group. | 
**runId** | **string** | Specifies a unique run id of the Protection Group run. | 
**objectId** | **string** | Specifies the id of the object for which debug logs are to be returned.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRunDebugLogsForObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRunMessagesReport

> GetRunMessagesReport(ctx, id, runId, objectId).FileType(fileType).Name(name).Execute()

Get the CSV of various Proto Messages for a given run and an object.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := "id_example" // string | Specifies a unique id of the Protection Group.
	runId := "runId_example" // string | Specifies a unique run id of the Protection Group run.
	objectId := "objectId_example" // string | Specifies the id of the object for which errors/warnings are to be returned. 
	fileType := "fileType_example" // string | Specifies the downloaded type, i.e: inclusion_exclusion_reports, error_files_list. default: error_files_list (optional)
	name := "name_example" // string | Specifies the name of the source being backed up (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProtectionGroupAPI.GetRunMessagesReport(context.Background(), id, runId, objectId).FileType(fileType).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroupAPI.GetRunMessagesReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the Protection Group. | 
**runId** | **string** | Specifies a unique run id of the Protection Group run. | 
**objectId** | **string** | Specifies the id of the object for which errors/warnings are to be returned.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRunMessagesReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fileType** | **string** | Specifies the downloaded type, i.e: inclusion_exclusion_reports, error_files_list. default: error_files_list | 
 **name** | **string** | Specifies the name of the source being backed up | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRunsReport

> GetRunsReport(ctx, id, runId, objectId).FileType(fileType).Name(name).Execute()

Get the CSV of errors/warnings for a given run and an object.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := "id_example" // string | Specifies a unique id of the Protection Group.
	runId := "runId_example" // string | Specifies a unique run id of the Protection Group run.
	objectId := "objectId_example" // string | Specifies the id of the object for which errors/warnings are to be returned. 
	fileType := "fileType_example" // string | Specifies the downloaded type, i.e: success_files_list, default: success_files_list (optional)
	name := "name_example" // string | Specifies the name of the source being backed up (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProtectionGroupAPI.GetRunsReport(context.Background(), id, runId, objectId).FileType(fileType).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroupAPI.GetRunsReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the Protection Group. | 
**runId** | **string** | Specifies a unique run id of the Protection Group run. | 
**objectId** | **string** | Specifies the id of the object for which errors/warnings are to be returned.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRunsReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fileType** | **string** | Specifies the downloaded type, i.e: success_files_list, default: success_files_list | 
 **name** | **string** | Specifies the name of the source being backed up | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PerformActionOnProtectionGroupRun

> PerformRunActionResponse PerformActionOnProtectionGroupRun(ctx, id).Body(body).Execute()

Actions on protection group run.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := "id_example" // string | Specifies a unique id of the Protection Group.
	body := *openapiclient.NewPerformActionOnProtectionGroupRunRequest("Action_example") // PerformActionOnProtectionGroupRunRequest | Specifies the parameters to perform an action on a protection run.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProtectionGroupAPI.PerformActionOnProtectionGroupRun(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroupAPI.PerformActionOnProtectionGroupRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PerformActionOnProtectionGroupRun`: PerformRunActionResponse
	fmt.Fprintf(os.Stdout, "Response from `ProtectionGroupAPI.PerformActionOnProtectionGroupRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the Protection Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPerformActionOnProtectionGroupRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PerformActionOnProtectionGroupRunRequest**](PerformActionOnProtectionGroupRunRequest.md) | Specifies the parameters to perform an action on a protection run. | 

### Return type

[**PerformRunActionResponse**](PerformRunActionResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProtectionGroup

> ProtectionGroup UpdateProtectionGroup(ctx, id).Body(body).Execute()

Update a Protection Group.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := "id_example" // string | Specifies the id of the Protection Group.
	body := *openapiclient.NewCreateOrUpdateProtectionGroupRequest("Environment_example", "Name_example", "PolicyId_example") // CreateOrUpdateProtectionGroupRequest | Specifies the parameters to update a Protection Group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProtectionGroupAPI.UpdateProtectionGroup(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroupAPI.UpdateProtectionGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProtectionGroup`: ProtectionGroup
	fmt.Fprintf(os.Stdout, "Response from `ProtectionGroupAPI.UpdateProtectionGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id of the Protection Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProtectionGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateOrUpdateProtectionGroupRequest**](CreateOrUpdateProtectionGroupRequest.md) | Specifies the parameters to update a Protection Group. | 

### Return type

[**ProtectionGroup**](ProtectionGroup.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProtectionGroupRun

> UpdateProtectionGroupRunResponseBody UpdateProtectionGroupRun(ctx, id).Body(body).Execute()

Update runs for a particular Protection Group.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := "id_example" // string | Specifies a unique id of the Protection Group.
	body := *openapiclient.NewUpdateProtectionGroupRunRequestBody([]openapiclient.UpdateProtectionGroupRunParams{*openapiclient.NewUpdateProtectionGroupRunParams("RunId_example")}) // UpdateProtectionGroupRunRequestBody | Specifies the parameters to update a Protection Group Run.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProtectionGroupAPI.UpdateProtectionGroupRun(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroupAPI.UpdateProtectionGroupRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProtectionGroupRun`: UpdateProtectionGroupRunResponseBody
	fmt.Fprintf(os.Stdout, "Response from `ProtectionGroupAPI.UpdateProtectionGroupRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the Protection Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProtectionGroupRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateProtectionGroupRunRequestBody**](UpdateProtectionGroupRunRequestBody.md) | Specifies the parameters to update a Protection Group Run. | 

### Return type

[**UpdateProtectionGroupRunResponseBody**](UpdateProtectionGroupRunResponseBody.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProtectionGroupsState

> UpdateProtectionGroupsState UpdateProtectionGroupsState(ctx).Body(body).Execute()

Perform an action like pause, resume, active, deactivate on all specified Protection Groups.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewUpdateProtectionGroupsStateRequest("Action_example", []string{"Ids_example"}) // UpdateProtectionGroupsStateRequest | Specifies the parameters to perform an action of list of Protection Groups.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProtectionGroupAPI.UpdateProtectionGroupsState(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroupAPI.UpdateProtectionGroupsState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProtectionGroupsState`: UpdateProtectionGroupsState
	fmt.Fprintf(os.Stdout, "Response from `ProtectionGroupAPI.UpdateProtectionGroupsState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProtectionGroupsStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateProtectionGroupsStateRequest**](UpdateProtectionGroupsStateRequest.md) | Specifies the parameters to perform an action of list of Protection Groups. | 

### Return type

[**UpdateProtectionGroupsState**](UpdateProtectionGroupsState.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

