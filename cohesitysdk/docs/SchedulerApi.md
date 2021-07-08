# \SchedulerApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSchedulerJob**](SchedulerApi.md#CreateSchedulerJob) | **Post** /public/scheduler | Create an email report scheduler job.
[**DeleteSchedulerJobs**](SchedulerApi.md#DeleteSchedulerJobs) | **Delete** /public/scheduler | Delete one or more email report schedule jobs.
[**GetSchedulerJobs**](SchedulerApi.md#GetSchedulerJobs) | **Get** /public/scheduler | List the email report schedule jobs that are scheduled to run.
[**UpdateSchedulerJob**](SchedulerApi.md#UpdateSchedulerJob) | **Put** /public/scheduler | Update an existing email report schedule job.



## CreateSchedulerJob

> SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameter CreateSchedulerJob(ctx).Body(body).Execute()

Create an email report scheduler job.



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
    body := *openapiclient.NewSchedulerProtoSchedulerJob() // SchedulerProtoSchedulerJob | 

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

    request := cohesitysdk.ApiCreateSchedulerJobRequest{
        Body: &body
    }

    resp, r, err := api_client.SchedulerApi.CreateSchedulerJob(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulerApi.CreateSchedulerJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSchedulerJob`: SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameter
    fmt.Fprintf(os.Stdout, "Response from `SchedulerApi.CreateSchedulerJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSchedulerJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SchedulerProtoSchedulerJob**](SchedulerProtoSchedulerJob.md) |  | 

### Return type

[**SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameter**](SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameter.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSchedulerJobs

> DeleteSchedulerJobs(ctx).Ids(ids).Execute()

Delete one or more email report schedule jobs.



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
    ids := []int64{int64(123)} // []int64 | Array of ids (optional)

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

    request := cohesitysdk.ApiDeleteSchedulerJobsRequest{
        Ids: &ids
    }

    resp, r, err := api_client.SchedulerApi.DeleteSchedulerJobs(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulerApi.DeleteSchedulerJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSchedulerJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]int64** | Array of ids | 

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


## GetSchedulerJobs

> SchedulerProto GetSchedulerJobs(ctx).Ids(ids).TenantIds(tenantIds).AllUnderHierarchy(allUnderHierarchy).Execute()

List the email report schedule jobs that are scheduled to run.



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
    ids := []int64{int64(123)} // []int64 | Specifies ids of scheduler jobs to fetch. (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
    allUnderHierarchy := true // bool | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user's organization should be returned. (optional)

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

    request := cohesitysdk.ApiGetSchedulerJobsRequest{
        Ids: &ids
        TenantIds: &tenantIds
        AllUnderHierarchy: &allUnderHierarchy
    }

    resp, r, err := api_client.SchedulerApi.GetSchedulerJobs(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulerApi.GetSchedulerJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchedulerJobs`: SchedulerProto
    fmt.Fprintf(os.Stdout, "Response from `SchedulerApi.GetSchedulerJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSchedulerJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]int64** | Specifies ids of scheduler jobs to fetch. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **allUnderHierarchy** | **bool** | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. | 

### Return type

[**SchedulerProto**](SchedulerProto.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSchedulerJob

> SchedulerProtoSchedulerJob UpdateSchedulerJob(ctx).Body(body).Execute()

Update an existing email report schedule job.



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
    body := *openapiclient.NewSchedulerProtoSchedulerJob() // SchedulerProtoSchedulerJob | Update Job Parameter. (optional)

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

    request := cohesitysdk.ApiUpdateSchedulerJobRequest{
        Body: &body
    }

    resp, r, err := api_client.SchedulerApi.UpdateSchedulerJob(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulerApi.UpdateSchedulerJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSchedulerJob`: SchedulerProtoSchedulerJob
    fmt.Fprintf(os.Stdout, "Response from `SchedulerApi.UpdateSchedulerJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSchedulerJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SchedulerProtoSchedulerJob**](SchedulerProtoSchedulerJob.md) | Update Job Parameter. | 

### Return type

[**SchedulerProtoSchedulerJob**](SchedulerProtoSchedulerJob.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

