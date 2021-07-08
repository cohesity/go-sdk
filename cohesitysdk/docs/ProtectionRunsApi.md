# \ProtectionRunsApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelProtectionJobRun**](ProtectionRunsApi.md#CancelProtectionJobRun) | **Post** /public/protectionRuns/cancel/{id} | Cancel a Protection Job run.
[**GetProtectionRunErrors**](ProtectionRunsApi.md#GetProtectionRunErrors) | **Get** /public/protectionRuns/errors | List Protection Job Run Errors filtered by the specified parameters.
[**GetProtectionRuns**](ProtectionRunsApi.md#GetProtectionRuns) | **Get** /public/protectionRuns | List Protection Job Runs filtered by the specified parameters.
[**UpdateProtectionRuns**](ProtectionRunsApi.md#UpdateProtectionRuns) | **Put** /public/protectionRuns | Update how long Protection Job Runs and their snapshots are retained on the Cohesity Cluster.



## CancelProtectionJobRun

> CancelProtectionJobRun(ctx, id).Body(body).Execute()

Cancel a Protection Job run.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    id := int64(789) // int64 | Specifies a unique id of the Protection Job.
    body := *openapiclient.NewCancelProtectionJobRunParam() // CancelProtectionJobRunParam |  (optional)

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

    request := cohesitysdk.ApiCancelProtectionJobRunRequest{
        Id: &id
        Body: &body
    }

    resp, r, err := api_client.ProtectionRunsApi.CancelProtectionJobRun(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionRunsApi.CancelProtectionJobRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the Protection Job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelProtectionJobRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CancelProtectionJobRunParam**](CancelProtectionJobRunParam.md) |  | 

### Return type

 (empty response body)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtectionRunErrors

> ProtectionRunErrors GetProtectionRunErrors(ctx).JobId(jobId).StartTimeUsecs(startTimeUsecs).TaskId(taskId).LimitNumberOfErrors(limitNumberOfErrors).PaginationCookie(paginationCookie).Execute()

List Protection Job Run Errors filtered by the specified parameters.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    jobId := int64(789) // int64 | Specifies the id of the Protection Job whose runs are to be returned. This field is required.
    startTimeUsecs := int64(789) // int64 | Specifies the time when the Job Run started as a Unix epoch Timestamp (in microseconds). This field is required
    taskId := int64(789) // int64 | Specifies the id of the Protection Run task for which errors are to be returned. This field is required to get the errors list.
    limitNumberOfErrors := int64(789) // int64 | Specifies the number of the results expected. (optional)
    paginationCookie := "paginationCookie_example" // string | Specifies the cookie for next set of results. (optional)

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

    request := cohesitysdk.ApiGetProtectionRunErrorsRequest{
        JobId: &jobId
        StartTimeUsecs: &startTimeUsecs
        TaskId: &taskId
        LimitNumberOfErrors: &limitNumberOfErrors
        PaginationCookie: &paginationCookie
    }

    resp, r, err := api_client.ProtectionRunsApi.GetProtectionRunErrors(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionRunsApi.GetProtectionRunErrors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtectionRunErrors`: ProtectionRunErrors
    fmt.Fprintf(os.Stdout, "Response from `ProtectionRunsApi.GetProtectionRunErrors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionRunErrorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobId** | **int64** | Specifies the id of the Protection Job whose runs are to be returned. This field is required. | 
 **startTimeUsecs** | **int64** | Specifies the time when the Job Run started as a Unix epoch Timestamp (in microseconds). This field is required | 
 **taskId** | **int64** | Specifies the id of the Protection Run task for which errors are to be returned. This field is required to get the errors list. | 
 **limitNumberOfErrors** | **int64** | Specifies the number of the results expected. | 
 **paginationCookie** | **string** | Specifies the cookie for next set of results. | 

### Return type

[**ProtectionRunErrors**](ProtectionRunErrors.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtectionRuns

> []ProtectionRunInstance GetProtectionRuns(ctx).JobId(jobId).IncludeRpoSnapshots(includeRpoSnapshots).StartedTimeUsecs(startedTimeUsecs).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).NumRuns(numRuns).ExcludeTasks(excludeTasks).SourceId(sourceId).RunTypes(runTypes).ExcludeErrorRuns(excludeErrorRuns).ExcludeNonRestoreableRuns(excludeNonRestoreableRuns).OnlyReturnShellInfo(onlyReturnShellInfo).Execute()

List Protection Job Runs filtered by the specified parameters.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    jobId := int64(789) // int64 | Filter by a Protection Job that is specified by id. If not specified, all Job Runs for all Protection Jobs are returned. (optional)
    includeRpoSnapshots := true // bool | If true, then the snapshots for Protection Sources protected by Rpo policies will also be returned. (optional)
    startedTimeUsecs := int64(789) // int64 | Return a specific Job Run by specifying a time and a jobId. Specify the time when the Job Run started as a Unix epoch Timestamp (in microseconds). If this field is specified, jobId must also be specified. (optional)
    startTimeUsecs := int64(789) // int64 | Filter by a start time. Only Job Runs that started after the specified time are returned. Specify the start time as a Unix epoch Timestamp (in microseconds). (optional)
    endTimeUsecs := int64(789) // int64 | Filter by a end time specified as a Unix epoch Timestamp (in microseconds). Only Job Runs that completed before the specified end time are returned. (optional)
    numRuns := int64(789) // int64 | Specify the number of Job Runs to return. The newest Job Runs are returned. (optional)
    excludeTasks := true // bool | If true, the individual backup status for all the objects protected by the Job Run are not populated in the response. For example in a VMware environment, the status of backing up each VM associated with a Job is not returned. (optional)
    sourceId := int64(789) // int64 | Filter by source id. Only Job Runs protecting the specified source (such as a VM or View) are returned. The source id is assigned by the Cohesity Cluster. (optional)
    runTypes := []string{"Inner_example"} // []string | Filter by run type such as 'kFull', 'kRegular' or 'kLog'. If not specified, Job Runs of all types are returned. (optional)
    excludeErrorRuns := true // bool | Filter out Jobs Runs with errors by setting this field to 'true'. If not set or set to 'false', Job Runs with errors are returned. (optional)
    excludeNonRestoreableRuns := true // bool | Filter out jobs runs that cannot be restored by setting this field to 'true'. If not set or set to 'false', Runs without any successful object will be returned. The default value is false. (optional)
    onlyReturnShellInfo := true // bool | If passed as true, then only returns the summary information about run including details such as runs start time, status, type etc. It does not include extra details such as attempt/task info etc. (optional)

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

    request := cohesitysdk.ApiGetProtectionRunsRequest{
        JobId: &jobId
        IncludeRpoSnapshots: &includeRpoSnapshots
        StartedTimeUsecs: &startedTimeUsecs
        StartTimeUsecs: &startTimeUsecs
        EndTimeUsecs: &endTimeUsecs
        NumRuns: &numRuns
        ExcludeTasks: &excludeTasks
        SourceId: &sourceId
        RunTypes: &runTypes
        ExcludeErrorRuns: &excludeErrorRuns
        ExcludeNonRestoreableRuns: &excludeNonRestoreableRuns
        OnlyReturnShellInfo: &onlyReturnShellInfo
    }

    resp, r, err := api_client.ProtectionRunsApi.GetProtectionRuns(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionRunsApi.GetProtectionRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtectionRuns`: []ProtectionRunInstance
    fmt.Fprintf(os.Stdout, "Response from `ProtectionRunsApi.GetProtectionRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobId** | **int64** | Filter by a Protection Job that is specified by id. If not specified, all Job Runs for all Protection Jobs are returned. | 
 **includeRpoSnapshots** | **bool** | If true, then the snapshots for Protection Sources protected by Rpo policies will also be returned. | 
 **startedTimeUsecs** | **int64** | Return a specific Job Run by specifying a time and a jobId. Specify the time when the Job Run started as a Unix epoch Timestamp (in microseconds). If this field is specified, jobId must also be specified. | 
 **startTimeUsecs** | **int64** | Filter by a start time. Only Job Runs that started after the specified time are returned. Specify the start time as a Unix epoch Timestamp (in microseconds). | 
 **endTimeUsecs** | **int64** | Filter by a end time specified as a Unix epoch Timestamp (in microseconds). Only Job Runs that completed before the specified end time are returned. | 
 **numRuns** | **int64** | Specify the number of Job Runs to return. The newest Job Runs are returned. | 
 **excludeTasks** | **bool** | If true, the individual backup status for all the objects protected by the Job Run are not populated in the response. For example in a VMware environment, the status of backing up each VM associated with a Job is not returned. | 
 **sourceId** | **int64** | Filter by source id. Only Job Runs protecting the specified source (such as a VM or View) are returned. The source id is assigned by the Cohesity Cluster. | 
 **runTypes** | **[]string** | Filter by run type such as &#39;kFull&#39;, &#39;kRegular&#39; or &#39;kLog&#39;. If not specified, Job Runs of all types are returned. | 
 **excludeErrorRuns** | **bool** | Filter out Jobs Runs with errors by setting this field to &#39;true&#39;. If not set or set to &#39;false&#39;, Job Runs with errors are returned. | 
 **excludeNonRestoreableRuns** | **bool** | Filter out jobs runs that cannot be restored by setting this field to &#39;true&#39;. If not set or set to &#39;false&#39;, Runs without any successful object will be returned. The default value is false. | 
 **onlyReturnShellInfo** | **bool** | If passed as true, then only returns the summary information about run including details such as runs start time, status, type etc. It does not include extra details such as attempt/task info etc. | 

### Return type

[**[]ProtectionRunInstance**](ProtectionRunInstance.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProtectionRuns

> UpdateProtectionRuns(ctx).Body(body).Execute()

Update how long Protection Job Runs and their snapshots are retained on the Cohesity Cluster.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewUpdateProtectionJobRunsParam() // UpdateProtectionJobRunsParam | Request to update the expiration time of Protection Job Runs.

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

    request := cohesitysdk.ApiUpdateProtectionRunsRequest{
        Body: &body
    }

    resp, r, err := api_client.ProtectionRunsApi.UpdateProtectionRuns(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionRunsApi.UpdateProtectionRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProtectionRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateProtectionJobRunsParam**](UpdateProtectionJobRunsParam.md) | Request to update the expiration time of Protection Job Runs. | 

### Return type

 (empty response body)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

