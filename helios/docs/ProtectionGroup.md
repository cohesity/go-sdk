# \ProtectionGroup

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProtectionGroup**](ProtectionGroup.md#CreateProtectionGroup) | **Post** /data-protect/protection-groups | Create a Protection Group.
[**CreateProtectionGroupRun**](ProtectionGroup.md#CreateProtectionGroupRun) | **Post** /data-protect/protection-groups/{id}/runs | Create a new protection run.
[**DeleteProtectionGroup**](ProtectionGroup.md#DeleteProtectionGroup) | **Delete** /data-protect/protection-groups/{id} | Delete a Protection Group.
[**GetProtectionGroupById**](ProtectionGroup.md#GetProtectionGroupById) | **Get** /data-protect/protection-groups/{id} | List details about single Protection Group.
[**GetProtectionGroupRun**](ProtectionGroup.md#GetProtectionGroupRun) | **Get** /data-protect/protection-groups/{id}/runs/{runId} | Get a run for a Protection Group.
[**GetProtectionGroupRuns**](ProtectionGroup.md#GetProtectionGroupRuns) | **Get** /data-protect/protection-groups/{id}/runs | Get the list of runs for a Protection Group.
[**GetProtectionGroups**](ProtectionGroup.md#GetProtectionGroups) | **Get** /data-protect/protection-groups | Get the list of Protection Groups.
[**GetProtectionRunProgress**](ProtectionGroup.md#GetProtectionRunProgress) | **Get** /data-protect/runs/{runId}/progress | Get the progress of a run.
[**GetProtectionRuns**](ProtectionGroup.md#GetProtectionRuns) | **Get** /data-protect/runs/summary | Get the list of runs.
[**GetRunDebugLogs**](ProtectionGroup.md#GetRunDebugLogs) | **Get** /data-protect/protection-groups/{id}/runs/{runId}/debug-logs | Get the debug logs for a run from a Protection Group.
[**GetRunDebugLogsForObject**](ProtectionGroup.md#GetRunDebugLogsForObject) | **Get** /data-protect/protection-groups/{id}/runs/{runId}/objects/{objectId}/debug-logs | Get the debug logs for a particular object in a run from a Protection Group.
[**GetRunErrorsReport**](ProtectionGroup.md#GetRunErrorsReport) | **Get** /data-protect/protection-groups/{id}/runs/{runId}/objects/{objectId}/download-messages | Get the CSV of errors/warnings for a given run and an object.
[**PerformActionOnProtectionGroupRun**](ProtectionGroup.md#PerformActionOnProtectionGroupRun) | **Post** /data-protect/protection-groups/{id}/runs/actions | Actions on protection group run.
[**UpdateProtectionGroup**](ProtectionGroup.md#UpdateProtectionGroup) | **Put** /data-protect/protection-groups/{id} | Update a Protection Group.
[**UpdateProtectionGroupRun**](ProtectionGroup.md#UpdateProtectionGroupRun) | **Put** /data-protect/protection-groups/{id}/runs | Update runs for a particular Protection Group.
[**UpdateProtectionGroupsState**](ProtectionGroup.md#UpdateProtectionGroupsState) | **Post** /data-protect/protection-groups/states | Perform an action like pause, resume, active, deactivate on all specified Protection Groups.



## CreateProtectionGroup

> ProtectionGroup CreateProtectionGroup(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Create a Protection Group.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    body := *openapiclient.NewCreateOrUpdateProtectionGroupRequest("Name_example", "PolicyId_example", "Environment_example") // CreateOrUpdateProtectionGroupRequest | Specifies the parameters to create a Protection Group.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := helios.ApiCreateProtectionGroupRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.ProtectionGroup.CreateProtectionGroup(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroup.CreateProtectionGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProtectionGroup`: ProtectionGroup
    fmt.Fprintf(os.Stdout, "Response from `ProtectionGroup.CreateProtectionGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProtectionGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateOrUpdateProtectionGroupRequest**](CreateOrUpdateProtectionGroupRequest.md) | Specifies the parameters to create a Protection Group. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

> CreateProtectionGroupRunResponseBody CreateProtectionGroupRun(ctx, id).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Create a new protection run.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | Specifies a unique id of the Protection Group.
    body := *openapiclient.NewCreateProtectionGroupRunRequest("RunType_example") // CreateProtectionGroupRunRequest | Specifies the parameters to start a protection run.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := helios.ApiCreateProtectionGroupRunRequest{
        Id: &id
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.ProtectionGroup.CreateProtectionGroupRun(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroup.CreateProtectionGroupRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProtectionGroupRun`: CreateProtectionGroupRunResponseBody
    fmt.Fprintf(os.Stdout, "Response from `ProtectionGroup.CreateProtectionGroupRun`: %v\n", resp)
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
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

> DeleteProtectionGroup(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).DeleteSnapshots(deleteSnapshots).Execute()

Delete a Protection Group.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | Specifies a unique id of the Protection Group.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    deleteSnapshots := true // bool | Specifies if Snapshots generated by the Protection Group should also be deleted when the Protection Group is deleted. (optional)

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := helios.ApiDeleteProtectionGroupRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        DeleteSnapshots: &deleteSnapshots
    }

    resp, r, err := api_client.ProtectionGroup.DeleteProtectionGroup(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroup.DeleteProtectionGroup``: %v\n", err)
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

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
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

> ProtectionGroup GetProtectionGroupById(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).IncludeLastRunInfo(includeLastRunInfo).PruneExcludedSourceIds(pruneExcludedSourceIds).Execute()

List details about single Protection Group.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | Specifies a unique id of the Protection Group.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    includeLastRunInfo := true // bool | If true, the response will include last run info. If it is false or not specified, the last run info won't be returned. (optional)
    pruneExcludedSourceIds := true // bool | If true, the response will not include the list of excluded source IDs in groups that contain this field. This can be set to true in order to improve performance if excluded source IDs are not needed by the user. (optional)

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := helios.ApiGetProtectionGroupByIdRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        IncludeLastRunInfo: &includeLastRunInfo
        PruneExcludedSourceIds: &pruneExcludedSourceIds
    }

    resp, r, err := api_client.ProtectionGroup.GetProtectionGroupById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroup.GetProtectionGroupById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtectionGroupById`: ProtectionGroup
    fmt.Fprintf(os.Stdout, "Response from `ProtectionGroup.GetProtectionGroupById`: %v\n", resp)
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

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **includeLastRunInfo** | **bool** | If true, the response will include last run info. If it is false or not specified, the last run info won&#39;t be returned. | 
 **pruneExcludedSourceIds** | **bool** | If true, the response will not include the list of excluded source IDs in groups that contain this field. This can be set to true in order to improve performance if excluded source IDs are not needed by the user. | 

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

> ProtectionGroupRun GetProtectionGroupRun(ctx, id, runId).AccessClusterId(accessClusterId).RegionId(regionId).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeObjectDetails(includeObjectDetails).Execute()

Get a run for a Protection Group.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | Specifies a unique id of the Protection Group.
    runId := "runId_example" // string | Specifies a unique run id of the Protection Group run.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which the run is to be returned. (optional)
    includeTenants := true // bool | If true, the response will include Protection Group Runs which were created by all tenants which the current user has permission to see. If false, then only Protection Groups created by the current user will be returned. If it's not specified, it is true by default. (optional)
    includeObjectDetails := true // bool | Specifies if the result includes the object details for a protection run. If set to true, details of the protected object will be returned. If set to false or not specified, details will not be returned. (optional)

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := helios.ApiGetProtectionGroupRunRequest{
        Id: &id
        RunId: &runId
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        TenantIds: &tenantIds
        IncludeTenants: &includeTenants
        IncludeObjectDetails: &includeObjectDetails
    }

    resp, r, err := api_client.ProtectionGroup.GetProtectionGroupRun(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroup.GetProtectionGroupRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtectionGroupRun`: ProtectionGroupRun
    fmt.Fprintf(os.Stdout, "Response from `ProtectionGroup.GetProtectionGroupRun`: %v\n", resp)
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


 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which the run is to be returned. | 
 **includeTenants** | **bool** | If true, the response will include Protection Group Runs which were created by all tenants which the current user has permission to see. If false, then only Protection Groups created by the current user will be returned. If it&#39;s not specified, it is true by default. | 
 **includeObjectDetails** | **bool** | Specifies if the result includes the object details for a protection run. If set to true, details of the protected object will be returned. If set to false or not specified, details will not be returned. | 

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

> ProtectionGroupRuns GetProtectionGroupRuns(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).RunId(runId).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).TenantIds(tenantIds).IncludeTenants(includeTenants).RunTypes(runTypes).IncludeObjectDetails(includeObjectDetails).LocalBackupRunStatus(localBackupRunStatus).ReplicationRunStatus(replicationRunStatus).ArchivalRunStatus(archivalRunStatus).CloudSpinRunStatus(cloudSpinRunStatus).NumRuns(numRuns).ExcludeNonRestorableRuns(excludeNonRestorableRuns).Execute()

Get the list of runs for a Protection Group.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | Specifies a unique id of the Protection Group.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    runId := "runId_example" // string | Specifies the protection run id. (optional)
    startTimeUsecs := int64(789) // int64 | Filter by a start time. Specify the start time as a Unix epoch Timestamp (in microseconds). (optional)
    endTimeUsecs := int64(789) // int64 | Filter by a end time. Specify the start time as a Unix epoch Timestamp (in microseconds). (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
    includeTenants := true // bool | If true, the response will include Protection Group Runs which were created by all tenants which the current user has permission to see. If false, then only Protection Group Runs created by the current user will be returned. (optional)
    runTypes := []string{"RunTypes_example"} // []string | Filter by run type. Only protection run matching the specified types will be returned. (optional)
    includeObjectDetails := true // bool | Specifies if the result includes the object details for each protection run. If set to true, details of the protected object will be returned. If set to false or not specified, details will not be returned. (optional)
    localBackupRunStatus := []string{"LocalBackupRunStatus_example"} // []string | Specifies a list of local backup status, runs matching the status will be returned.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. (optional)
    replicationRunStatus := []string{"ReplicationRunStatus_example"} // []string | Specifies a list of replication status, runs matching the status will be returned.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. (optional)
    archivalRunStatus := []string{"ArchivalRunStatus_example"} // []string | Specifies a list of archival status, runs matching the status will be returned.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. (optional)
    cloudSpinRunStatus := []string{"CloudSpinRunStatus_example"} // []string | Specifies a list of cloud spin status, runs matching the status will be returned.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. (optional)
    numRuns := int64(789) // int64 | Specifies the max number of runs. If not specified, at most 100 runs will be returned. (optional)
    excludeNonRestorableRuns := true // bool | Specifies whether to exclude non restorable runs. Run is treated restorable only if there is atleast one object snapshot (which may be either a local or an archival snapshot) which is not deleted or expired. Default value is false. (optional) (default to false)

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := helios.ApiGetProtectionGroupRunsRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        RunId: &runId
        StartTimeUsecs: &startTimeUsecs
        EndTimeUsecs: &endTimeUsecs
        TenantIds: &tenantIds
        IncludeTenants: &includeTenants
        RunTypes: &runTypes
        IncludeObjectDetails: &includeObjectDetails
        LocalBackupRunStatus: &localBackupRunStatus
        ReplicationRunStatus: &replicationRunStatus
        ArchivalRunStatus: &archivalRunStatus
        CloudSpinRunStatus: &cloudSpinRunStatus
        NumRuns: &numRuns
        ExcludeNonRestorableRuns: &excludeNonRestorableRuns
    }

    resp, r, err := api_client.ProtectionGroup.GetProtectionGroupRuns(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroup.GetProtectionGroupRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtectionGroupRuns`: ProtectionGroupRuns
    fmt.Fprintf(os.Stdout, "Response from `ProtectionGroup.GetProtectionGroupRuns`: %v\n", resp)
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

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **runId** | **string** | Specifies the protection run id. | 
 **startTimeUsecs** | **int64** | Filter by a start time. Specify the start time as a Unix epoch Timestamp (in microseconds). | 
 **endTimeUsecs** | **int64** | Filter by a end time. Specify the start time as a Unix epoch Timestamp (in microseconds). | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **includeTenants** | **bool** | If true, the response will include Protection Group Runs which were created by all tenants which the current user has permission to see. If false, then only Protection Group Runs created by the current user will be returned. | 
 **runTypes** | **[]string** | Filter by run type. Only protection run matching the specified types will be returned. | 
 **includeObjectDetails** | **bool** | Specifies if the result includes the object details for each protection run. If set to true, details of the protected object will be returned. If set to false or not specified, details will not be returned. | 
 **localBackupRunStatus** | **[]string** | Specifies a list of local backup status, runs matching the status will be returned.&lt;br&gt; &#39;Running&#39; indicates that the run is still running.&lt;br&gt; &#39;Canceled&#39; indicates that the run has been canceled.&lt;br&gt; &#39;Canceling&#39; indicates that the run is in the process of being canceled.&lt;br&gt; &#39;Failed&#39; indicates that the run has failed.&lt;br&gt; &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening.&lt;br&gt; &#39;Succeeded&#39; indicates that the run has finished successfully.&lt;br&gt; &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages. | 
 **replicationRunStatus** | **[]string** | Specifies a list of replication status, runs matching the status will be returned.&lt;br&gt; &#39;Running&#39; indicates that the run is still running.&lt;br&gt; &#39;Canceled&#39; indicates that the run has been canceled.&lt;br&gt; &#39;Canceling&#39; indicates that the run is in the process of being canceled.&lt;br&gt; &#39;Failed&#39; indicates that the run has failed.&lt;br&gt; &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening.&lt;br&gt; &#39;Succeeded&#39; indicates that the run has finished successfully.&lt;br&gt; &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages. | 
 **archivalRunStatus** | **[]string** | Specifies a list of archival status, runs matching the status will be returned.&lt;br&gt; &#39;Running&#39; indicates that the run is still running.&lt;br&gt; &#39;Canceled&#39; indicates that the run has been canceled.&lt;br&gt; &#39;Canceling&#39; indicates that the run is in the process of being canceled.&lt;br&gt; &#39;Failed&#39; indicates that the run has failed.&lt;br&gt; &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening.&lt;br&gt; &#39;Succeeded&#39; indicates that the run has finished successfully.&lt;br&gt; &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages. | 
 **cloudSpinRunStatus** | **[]string** | Specifies a list of cloud spin status, runs matching the status will be returned.&lt;br&gt; &#39;Running&#39; indicates that the run is still running.&lt;br&gt; &#39;Canceled&#39; indicates that the run has been canceled.&lt;br&gt; &#39;Canceling&#39; indicates that the run is in the process of being canceled.&lt;br&gt; &#39;Failed&#39; indicates that the run has failed.&lt;br&gt; &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening.&lt;br&gt; &#39;Succeeded&#39; indicates that the run has finished successfully.&lt;br&gt; &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages. | 
 **numRuns** | **int64** | Specifies the max number of runs. If not specified, at most 100 runs will be returned. | 
 **excludeNonRestorableRuns** | **bool** | Specifies whether to exclude non restorable runs. Run is treated restorable only if there is atleast one object snapshot (which may be either a local or an archival snapshot) which is not deleted or expired. Default value is false. | [default to false]

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

> ProtectionGroups GetProtectionGroups(ctx).AccessClusterId(accessClusterId).RegionId(regionId).Ids(ids).Names(names).PolicyIds(policyIds).StorageDomainId(storageDomainId).IncludeGroupsWithDatalockOnly(includeGroupsWithDatalockOnly).Environments(environments).Office365Workloads(office365Workloads).IsActive(isActive).IsDeleted(isDeleted).IsPaused(isPaused).LastRunLocalBackupStatus(lastRunLocalBackupStatus).LastRunReplicationStatus(lastRunReplicationStatus).LastRunArchivalStatus(lastRunArchivalStatus).LastRunCloudSpinStatus(lastRunCloudSpinStatus).LastRunAnyStatus(lastRunAnyStatus).IsLastRunSlaViolated(isLastRunSlaViolated).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeLastRunInfo(includeLastRunInfo).PruneExcludedSourceIds(pruneExcludedSourceIds).Execute()

Get the list of Protection Groups.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
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
    lastRunLocalBackupStatus := []string{"LastRunLocalBackupStatus_example"} // []string | Filter by last local backup run status.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. (optional)
    lastRunReplicationStatus := []string{"LastRunReplicationStatus_example"} // []string | Filter by last remote replication run status.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. (optional)
    lastRunArchivalStatus := []string{"LastRunArchivalStatus_example"} // []string | Filter by last cloud archival run status.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. (optional)
    lastRunCloudSpinStatus := []string{"LastRunCloudSpinStatus_example"} // []string | Filter by last cloud spin run status.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. (optional)
    lastRunAnyStatus := []string{"LastRunAnyStatus_example"} // []string | Filter by last any run status.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. (optional)
    isLastRunSlaViolated := true // bool | If true, return Protection Groups for which last run SLA was violated. (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
    includeTenants := true // bool | If true, the response will include Protection Groups which were created by all tenants which the current user has permission to see. If false, then only Protection Groups created by the current user will be returned. (optional)
    includeLastRunInfo := true // bool | If true, the response will include last run info. If it is false or not specified, the last run info won't be returned. (optional)
    pruneExcludedSourceIds := true // bool | If true, the response will not include the list of excluded source IDs in groups that contain this field. This can be set to true in order to improve performance if excluded source IDs are not needed by the user. (optional)

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := helios.ApiGetProtectionGroupsRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        Ids: &ids
        Names: &names
        PolicyIds: &policyIds
        StorageDomainId: &storageDomainId
        IncludeGroupsWithDatalockOnly: &includeGroupsWithDatalockOnly
        Environments: &environments
        Office365Workloads: &office365Workloads
        IsActive: &isActive
        IsDeleted: &isDeleted
        IsPaused: &isPaused
        LastRunLocalBackupStatus: &lastRunLocalBackupStatus
        LastRunReplicationStatus: &lastRunReplicationStatus
        LastRunArchivalStatus: &lastRunArchivalStatus
        LastRunCloudSpinStatus: &lastRunCloudSpinStatus
        LastRunAnyStatus: &lastRunAnyStatus
        IsLastRunSlaViolated: &isLastRunSlaViolated
        TenantIds: &tenantIds
        IncludeTenants: &includeTenants
        IncludeLastRunInfo: &includeLastRunInfo
        PruneExcludedSourceIds: &pruneExcludedSourceIds
    }

    resp, r, err := api_client.ProtectionGroup.GetProtectionGroups(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroup.GetProtectionGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtectionGroups`: ProtectionGroups
    fmt.Fprintf(os.Stdout, "Response from `ProtectionGroup.GetProtectionGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
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
 **lastRunLocalBackupStatus** | **[]string** | Filter by last local backup run status.&lt;br&gt; &#39;Running&#39; indicates that the run is still running.&lt;br&gt; &#39;Canceled&#39; indicates that the run has been canceled.&lt;br&gt; &#39;Canceling&#39; indicates that the run is in the process of being canceled.&lt;br&gt; &#39;Failed&#39; indicates that the run has failed.&lt;br&gt; &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening.&lt;br&gt; &#39;Succeeded&#39; indicates that the run has finished successfully.&lt;br&gt; &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages. | 
 **lastRunReplicationStatus** | **[]string** | Filter by last remote replication run status.&lt;br&gt; &#39;Running&#39; indicates that the run is still running.&lt;br&gt; &#39;Canceled&#39; indicates that the run has been canceled.&lt;br&gt; &#39;Canceling&#39; indicates that the run is in the process of being canceled.&lt;br&gt; &#39;Failed&#39; indicates that the run has failed.&lt;br&gt; &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening.&lt;br&gt; &#39;Succeeded&#39; indicates that the run has finished successfully.&lt;br&gt; &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages. | 
 **lastRunArchivalStatus** | **[]string** | Filter by last cloud archival run status.&lt;br&gt; &#39;Running&#39; indicates that the run is still running.&lt;br&gt; &#39;Canceled&#39; indicates that the run has been canceled.&lt;br&gt; &#39;Canceling&#39; indicates that the run is in the process of being canceled.&lt;br&gt; &#39;Failed&#39; indicates that the run has failed.&lt;br&gt; &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening.&lt;br&gt; &#39;Succeeded&#39; indicates that the run has finished successfully.&lt;br&gt; &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages. | 
 **lastRunCloudSpinStatus** | **[]string** | Filter by last cloud spin run status.&lt;br&gt; &#39;Running&#39; indicates that the run is still running.&lt;br&gt; &#39;Canceled&#39; indicates that the run has been canceled.&lt;br&gt; &#39;Canceling&#39; indicates that the run is in the process of being canceled.&lt;br&gt; &#39;Failed&#39; indicates that the run has failed.&lt;br&gt; &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening.&lt;br&gt; &#39;Succeeded&#39; indicates that the run has finished successfully.&lt;br&gt; &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages. | 
 **lastRunAnyStatus** | **[]string** | Filter by last any run status.&lt;br&gt; &#39;Running&#39; indicates that the run is still running.&lt;br&gt; &#39;Canceled&#39; indicates that the run has been canceled.&lt;br&gt; &#39;Canceling&#39; indicates that the run is in the process of being canceled.&lt;br&gt; &#39;Failed&#39; indicates that the run has failed.&lt;br&gt; &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening.&lt;br&gt; &#39;Succeeded&#39; indicates that the run has finished successfully.&lt;br&gt; &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages. | 
 **isLastRunSlaViolated** | **bool** | If true, return Protection Groups for which last run SLA was violated. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **includeTenants** | **bool** | If true, the response will include Protection Groups which were created by all tenants which the current user has permission to see. If false, then only Protection Groups created by the current user will be returned. | 
 **includeLastRunInfo** | **bool** | If true, the response will include last run info. If it is false or not specified, the last run info won&#39;t be returned. | 
 **pruneExcludedSourceIds** | **bool** | If true, the response will not include the list of excluded source IDs in groups that contain this field. This can be set to true in order to improve performance if excluded source IDs are not needed by the user. | 

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

> GetProtectionRunProgressBody GetProtectionRunProgress(ctx, runId).AccessClusterId(accessClusterId).RegionId(regionId).Objects(objects).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeFinishedTasks(includeFinishedTasks).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).MaxTasksNum(maxTasksNum).ExcludeObjectDetails(excludeObjectDetails).IncludeEventLogs(includeEventLogs).MaxLogLevel(maxLogLevel).Execute()

Get the progress of a run.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    runId := "runId_example" // string | Specifies a unique run id of the Protection Run.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
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

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := helios.ApiGetProtectionRunProgressRequest{
        RunId: &runId
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        Objects: &objects
        TenantIds: &tenantIds
        IncludeTenants: &includeTenants
        IncludeFinishedTasks: &includeFinishedTasks
        StartTimeUsecs: &startTimeUsecs
        EndTimeUsecs: &endTimeUsecs
        MaxTasksNum: &maxTasksNum
        ExcludeObjectDetails: &excludeObjectDetails
        IncludeEventLogs: &includeEventLogs
        MaxLogLevel: &maxLogLevel
    }

    resp, r, err := api_client.ProtectionGroup.GetProtectionRunProgress(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroup.GetProtectionRunProgress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtectionRunProgress`: GetProtectionRunProgressBody
    fmt.Fprintf(os.Stdout, "Response from `ProtectionGroup.GetProtectionRunProgress`: %v\n", resp)
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

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
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


## GetProtectionRuns

> ProtectionRunsSummary GetProtectionRuns(ctx).AccessClusterId(accessClusterId).RegionId(regionId).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).RunStatus(runStatus).Execute()

Get the list of runs.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    startTimeUsecs := int64(789) // int64 | Filter by a start time. Specify the start time as a Unix epoch Timestamp (in microseconds), only runs executing after this time will be returned. By default it is endTimeUsecs minus an hour. (optional)
    endTimeUsecs := int64(789) // int64 | Filter by a end time. Specify the start time as a Unix epoch Timestamp (in microseconds), only runs executing before this time will be returned. By default it is current time. (optional)
    runStatus := []string{"RunStatus_example"} // []string | Specifies a list of status, runs matching the status will be returned.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. (optional)

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := helios.ApiGetProtectionRunsRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        StartTimeUsecs: &startTimeUsecs
        EndTimeUsecs: &endTimeUsecs
        RunStatus: &runStatus
    }

    resp, r, err := api_client.ProtectionGroup.GetProtectionRuns(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroup.GetProtectionRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtectionRuns`: ProtectionRunsSummary
    fmt.Fprintf(os.Stdout, "Response from `ProtectionGroup.GetProtectionRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **startTimeUsecs** | **int64** | Filter by a start time. Specify the start time as a Unix epoch Timestamp (in microseconds), only runs executing after this time will be returned. By default it is endTimeUsecs minus an hour. | 
 **endTimeUsecs** | **int64** | Filter by a end time. Specify the start time as a Unix epoch Timestamp (in microseconds), only runs executing before this time will be returned. By default it is current time. | 
 **runStatus** | **[]string** | Specifies a list of status, runs matching the status will be returned.&lt;br&gt; &#39;Running&#39; indicates that the run is still running.&lt;br&gt; &#39;Canceled&#39; indicates that the run has been canceled.&lt;br&gt; &#39;Canceling&#39; indicates that the run is in the process of being canceled.&lt;br&gt; &#39;Failed&#39; indicates that the run has failed.&lt;br&gt; &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening.&lt;br&gt; &#39;Succeeded&#39; indicates that the run has finished successfully.&lt;br&gt; &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages. | 

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

> GetRunDebugLogs(ctx, id, runId).AccessClusterId(accessClusterId).RegionId(regionId).ObjectId(objectId).Execute()

Get the debug logs for a run from a Protection Group.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | Specifies a unique id of the Protection Group.
    runId := "runId_example" // string | Specifies a unique run id of the Protection Group run.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    objectId := "objectId_example" // string | Specifies the id of the object for which debug logs are to be returned.  (optional)

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := helios.ApiGetRunDebugLogsRequest{
        Id: &id
        RunId: &runId
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        ObjectId: &objectId
    }

    resp, r, err := api_client.ProtectionGroup.GetRunDebugLogs(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroup.GetRunDebugLogs``: %v\n", err)
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


 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
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

> GetRunDebugLogsForObject(ctx, id, runId, objectId).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Get the debug logs for a particular object in a run from a Protection Group.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | Specifies a unique id of the Protection Group.
    runId := "runId_example" // string | Specifies a unique run id of the Protection Group run.
    objectId := "objectId_example" // string | Specifies the id of the object for which debug logs are to be returned. 
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := helios.ApiGetRunDebugLogsForObjectRequest{
        Id: &id
        RunId: &runId
        ObjectId: &objectId
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.ProtectionGroup.GetRunDebugLogsForObject(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroup.GetRunDebugLogsForObject``: %v\n", err)
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



 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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


## GetRunErrorsReport

> GetRunErrorsReport(ctx, id, runId, objectId).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Get the CSV of errors/warnings for a given run and an object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | Specifies a unique id of the Protection Group.
    runId := "runId_example" // string | Specifies a unique run id of the Protection Group run.
    objectId := "objectId_example" // string | Specifies the id of the object for which errors/warnings are to be returned. 
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := helios.ApiGetRunErrorsReportRequest{
        Id: &id
        RunId: &runId
        ObjectId: &objectId
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.ProtectionGroup.GetRunErrorsReport(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroup.GetRunErrorsReport``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetRunErrorsReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

> PerformRunActionResponse PerformActionOnProtectionGroupRun(ctx, id).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Actions on protection group run.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | Specifies a unique id of the Protection Group.
    body := *openapiclient.NewPerformActionOnProtectionGroupRunRequest("Action_example") // PerformActionOnProtectionGroupRunRequest | Specifies the parameters to perform an action on a protection run.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := helios.ApiPerformActionOnProtectionGroupRunRequest{
        Id: &id
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.ProtectionGroup.PerformActionOnProtectionGroupRun(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroup.PerformActionOnProtectionGroupRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PerformActionOnProtectionGroupRun`: PerformRunActionResponse
    fmt.Fprintf(os.Stdout, "Response from `ProtectionGroup.PerformActionOnProtectionGroupRun`: %v\n", resp)
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
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

> ProtectionGroup UpdateProtectionGroup(ctx, id).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Update a Protection Group.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | Specifies the id of the Protection Group.
    body := *openapiclient.NewCreateOrUpdateProtectionGroupRequest("Name_example", "PolicyId_example", "Environment_example") // CreateOrUpdateProtectionGroupRequest | Specifies the parameters to update a Protection Group.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := helios.ApiUpdateProtectionGroupRequest{
        Id: &id
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.ProtectionGroup.UpdateProtectionGroup(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroup.UpdateProtectionGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProtectionGroup`: ProtectionGroup
    fmt.Fprintf(os.Stdout, "Response from `ProtectionGroup.UpdateProtectionGroup`: %v\n", resp)
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
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

> UpdateProtectionGroupRunResponseBody UpdateProtectionGroupRun(ctx, id).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Update runs for a particular Protection Group.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | Specifies a unique id of the Protection Group.
    body := *openapiclient.NewUpdateProtectionGroupRunRequestBody([]openapiclient.UpdateProtectionGroupRunParams{*openapiclient.NewUpdateProtectionGroupRunParams("RunId_example")}) // UpdateProtectionGroupRunRequestBody | Specifies the parameters to update a Protection Group Run.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := helios.ApiUpdateProtectionGroupRunRequest{
        Id: &id
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.ProtectionGroup.UpdateProtectionGroupRun(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroup.UpdateProtectionGroupRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProtectionGroupRun`: UpdateProtectionGroupRunResponseBody
    fmt.Fprintf(os.Stdout, "Response from `ProtectionGroup.UpdateProtectionGroupRun`: %v\n", resp)
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
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

> UpdateProtectionGroupsState UpdateProtectionGroupsState(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Perform an action like pause, resume, active, deactivate on all specified Protection Groups.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    body := *openapiclient.NewUpdateProtectionGroupsStateRequest("Action_example", []string{"Ids_example"}) // UpdateProtectionGroupsStateRequest | Specifies the parameters to perform an action of list of Protection Groups.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := helios.ApiUpdateProtectionGroupsStateRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.ProtectionGroup.UpdateProtectionGroupsState(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionGroup.UpdateProtectionGroupsState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProtectionGroupsState`: UpdateProtectionGroupsState
    fmt.Fprintf(os.Stdout, "Response from `ProtectionGroup.UpdateProtectionGroupsState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProtectionGroupsStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateProtectionGroupsStateRequest**](UpdateProtectionGroupsStateRequest.md) | Specifies the parameters to perform an action of list of Protection Groups. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

