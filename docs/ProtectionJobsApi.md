# \ProtectionJobsApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeProtectionJobState**](ProtectionJobsApi.md#ChangeProtectionJobState) | **Post** /public/protectionJobState/{id} | Pause future Runs or resume future Runs of the specified Protection Job.
[**CreateProtectionJob**](ProtectionJobsApi.md#CreateProtectionJob) | **Post** /public/protectionJobs | Create a Protection Job.
[**DeleteProtectionJob**](ProtectionJobsApi.md#DeleteProtectionJob) | **Delete** /public/protectionJobs/{id} | Delete a Protection Job.
[**GetProtectionJobAudit**](ProtectionJobsApi.md#GetProtectionJobAudit) | **Get** /public/protectionJobs/{id}/auditTrail | List a protection job audit.
[**GetProtectionJobById**](ProtectionJobsApi.md#GetProtectionJobById) | **Get** /public/protectionJobs/{id} | List details about single Protection Job.
[**GetProtectionJobs**](ProtectionJobsApi.md#GetProtectionJobs) | **Get** /public/protectionJobs | List Protections Jobs filtered by the specified parameters.
[**RunProtectionJob**](ProtectionJobsApi.md#RunProtectionJob) | **Post** /public/protectionJobs/run/{id} | Immediately execute a single Protection Job Run.
[**UpdateProtectionJob**](ProtectionJobsApi.md#UpdateProtectionJob) | **Put** /public/protectionJobs/{id} | Update a Protection Job.
[**UpdateProtectionJobsState**](ProtectionJobsApi.md#UpdateProtectionJobsState) | **Post** /public/protectionJobs/states | Perform an action like pause, resume, activate, deactivate on all the specified Protection Jobs.



## ChangeProtectionJobState

> ChangeProtectionJobState(ctx, id).Body(body).Execute()

Pause future Runs or resume future Runs of the specified Protection Job.



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
    body := *openapiclient.NewChangeProtectionJobStateParam() // ChangeProtectionJobStateParam |  (optional)

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

    request := cohesitysdk.ApiChangeProtectionJobStateRequest{
        id: &Id
        body: &Body
    }

    resp, r, err := api_client.ProtectionJobsApi.ChangeProtectionJobState(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionJobsApi.ChangeProtectionJobState``: %v\n", err)
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

Other parameters are passed through a pointer to a apiChangeProtectionJobStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ChangeProtectionJobStateParam**](ChangeProtectionJobStateParam.md) |  | 

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


## CreateProtectionJob

> ProtectionJob CreateProtectionJob(ctx).Body(body).Execute()

Create a Protection Job.



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
    body := *openapiclient.NewProtectionJobRequestBody("Name_example", "PolicyId_example", []openapiclient.SourceSpecialParameter{*openapiclient.NewSourceSpecialParameter()}, NullableInt64(123)) // ProtectionJobRequestBody | Request to create a Protection Job.

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

    request := cohesitysdk.ApiCreateProtectionJobRequest{
        body: &Body
    }

    resp, r, err := api_client.ProtectionJobsApi.CreateProtectionJob(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionJobsApi.CreateProtectionJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProtectionJob`: ProtectionJob
    fmt.Fprintf(os.Stdout, "Response from `ProtectionJobsApi.CreateProtectionJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProtectionJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProtectionJobRequestBody**](ProtectionJobRequestBody.md) | Request to create a Protection Job. | 

### Return type

[**ProtectionJob**](ProtectionJob.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProtectionJob

> DeleteProtectionJob(ctx, id).Body(body).Execute()

Delete a Protection Job.



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
    body := *openapiclient.NewDeleteProtectionJobParam() // DeleteProtectionJobParam | Request to delete a protection job. (optional)

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

    request := cohesitysdk.ApiDeleteProtectionJobRequest{
        id: &Id
        body: &Body
    }

    resp, r, err := api_client.ProtectionJobsApi.DeleteProtectionJob(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionJobsApi.DeleteProtectionJob``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteProtectionJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DeleteProtectionJobParam**](DeleteProtectionJobParam.md) | Request to delete a protection job. | 

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


## GetProtectionJobAudit

> []ProtectionJobAuditTrail GetProtectionJobAudit(ctx, id).Execute()

List a protection job audit.



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

    request := cohesitysdk.ApiGetProtectionJobAuditRequest{
        id: &Id
    }

    resp, r, err := api_client.ProtectionJobsApi.GetProtectionJobAudit(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionJobsApi.GetProtectionJobAudit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtectionJobAudit`: []ProtectionJobAuditTrail
    fmt.Fprintf(os.Stdout, "Response from `ProtectionJobsApi.GetProtectionJobAudit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the Protection Job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionJobAuditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ProtectionJobAuditTrail**](ProtectionJobAuditTrail.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtectionJobById

> ProtectionJob GetProtectionJobById(ctx, id).Execute()

List details about single Protection Job.



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

    request := cohesitysdk.ApiGetProtectionJobByIdRequest{
        id: &Id
    }

    resp, r, err := api_client.ProtectionJobsApi.GetProtectionJobById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionJobsApi.GetProtectionJobById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtectionJobById`: ProtectionJob
    fmt.Fprintf(os.Stdout, "Response from `ProtectionJobsApi.GetProtectionJobById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the Protection Job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionJobByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProtectionJob**](ProtectionJob.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtectionJobs

> []ProtectionJob GetProtectionJobs(ctx).Ids(ids).Names(names).PolicyIds(policyIds).Environments(environments).IsActive(isActive).IsDeleted(isDeleted).OnlyReturnBasicSummary(onlyReturnBasicSummary).IncludeLastRunAndStats(includeLastRunAndStats).IncludeRpoSnapshots(includeRpoSnapshots).IsLastRunSlaViolated(isLastRunSlaViolated).OnlyReturnDataMigrationJobs(onlyReturnDataMigrationJobs).PruneExcludedSourceIds(pruneExcludedSourceIds).TenantIds(tenantIds).AllUnderHierarchy(allUnderHierarchy).Execute()

List Protections Jobs filtered by the specified parameters.



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
    ids := []int64{int64(123)} // []int64 | Filter by a list of Protection Job ids. (optional)
    names := []string{"Inner_example"} // []string | Filter by a list of Protection Job names. (optional)
    policyIds := []string{"Inner_example"} // []string | Filter by Policy ids that are associated with Protection Jobs. Only Jobs associated with the specified Policy ids, are returned. (optional)
    environments := []string{"Environments_example"} // []string | Filter by environment types such as 'kVMware', 'kView', etc. Only Jobs protecting the specified environment types are returned. NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter. (optional)
    isActive := true // bool | Filter by Inactive or Active Jobs. If not set, all Inactive and Active Jobs are returned. If true, only Active Jobs are returned. If false, only Inactive Jobs are returned. When you create a Protection Job on a Primary Cluster with a replication schedule, the Cluster creates an Inactive copy of the Job on the Remote Cluster. In addition, when an Active and running Job is deactivated, the Job becomes Inactive. (optional)
    isDeleted := true // bool | If true, return only Protection Jobs that have been deleted but still have Snapshots associated with them. If false, return all Protection Jobs except those Jobs that have been deleted and still have Snapshots associated with them. A Job that is deleted with all its Snapshots is not returned for either of these cases. (optional)
    onlyReturnBasicSummary := true // bool | if true then only job descriptions and the most recent run of the job will be returned. (optional)
    includeLastRunAndStats := true // bool | If true, return the last Protection Run of the Job and the summary stats. (optional)
    includeRpoSnapshots := true // bool | If true, then the Protected Objects protected by RPO policies will also be returned. (optional)
    isLastRunSlaViolated := true // bool | IsLastRunSlaViolated is the parameter to filter the Protection Jobs based on the SLA violation status of the last Protection Run. (optional)
    onlyReturnDataMigrationJobs := true // bool | OnlyReturnDataMigrationJobs specifies if only data migration jobs should be returned. If not set, no data migration job will be returned. (optional)
    pruneExcludedSourceIds := true // bool | If true, the list of exclusion sources will be omitted from the response. This can be used to improve performance when the exclusion sources are not needed. (optional)
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

    request := cohesitysdk.ApiGetProtectionJobsRequest{
        ids: &Ids
        names: &Names
        policyIds: &PolicyIds
        environments: &Environments
        isActive: &IsActive
        isDeleted: &IsDeleted
        onlyReturnBasicSummary: &OnlyReturnBasicSummary
        includeLastRunAndStats: &IncludeLastRunAndStats
        includeRpoSnapshots: &IncludeRpoSnapshots
        isLastRunSlaViolated: &IsLastRunSlaViolated
        onlyReturnDataMigrationJobs: &OnlyReturnDataMigrationJobs
        pruneExcludedSourceIds: &PruneExcludedSourceIds
        tenantIds: &TenantIds
        allUnderHierarchy: &AllUnderHierarchy
    }

    resp, r, err := api_client.ProtectionJobsApi.GetProtectionJobs(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionJobsApi.GetProtectionJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtectionJobs`: []ProtectionJob
    fmt.Fprintf(os.Stdout, "Response from `ProtectionJobsApi.GetProtectionJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]int64** | Filter by a list of Protection Job ids. | 
 **names** | **[]string** | Filter by a list of Protection Job names. | 
 **policyIds** | **[]string** | Filter by Policy ids that are associated with Protection Jobs. Only Jobs associated with the specified Policy ids, are returned. | 
 **environments** | **[]string** | Filter by environment types such as &#39;kVMware&#39;, &#39;kView&#39;, etc. Only Jobs protecting the specified environment types are returned. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. | 
 **isActive** | **bool** | Filter by Inactive or Active Jobs. If not set, all Inactive and Active Jobs are returned. If true, only Active Jobs are returned. If false, only Inactive Jobs are returned. When you create a Protection Job on a Primary Cluster with a replication schedule, the Cluster creates an Inactive copy of the Job on the Remote Cluster. In addition, when an Active and running Job is deactivated, the Job becomes Inactive. | 
 **isDeleted** | **bool** | If true, return only Protection Jobs that have been deleted but still have Snapshots associated with them. If false, return all Protection Jobs except those Jobs that have been deleted and still have Snapshots associated with them. A Job that is deleted with all its Snapshots is not returned for either of these cases. | 
 **onlyReturnBasicSummary** | **bool** | if true then only job descriptions and the most recent run of the job will be returned. | 
 **includeLastRunAndStats** | **bool** | If true, return the last Protection Run of the Job and the summary stats. | 
 **includeRpoSnapshots** | **bool** | If true, then the Protected Objects protected by RPO policies will also be returned. | 
 **isLastRunSlaViolated** | **bool** | IsLastRunSlaViolated is the parameter to filter the Protection Jobs based on the SLA violation status of the last Protection Run. | 
 **onlyReturnDataMigrationJobs** | **bool** | OnlyReturnDataMigrationJobs specifies if only data migration jobs should be returned. If not set, no data migration job will be returned. | 
 **pruneExcludedSourceIds** | **bool** | If true, the list of exclusion sources will be omitted from the response. This can be used to improve performance when the exclusion sources are not needed. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **allUnderHierarchy** | **bool** | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. | 

### Return type

[**[]ProtectionJob**](ProtectionJob.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunProtectionJob

> RunProtectionJob(ctx, id).Body(body).Execute()

Immediately execute a single Protection Job Run.



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
    body := *openapiclient.NewRunProtectionJobParam() // RunProtectionJobParam | Specifies the type of backup. If not specified, the 'kRegular' backup is run.

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

    request := cohesitysdk.ApiRunProtectionJobRequest{
        id: &Id
        body: &Body
    }

    resp, r, err := api_client.ProtectionJobsApi.RunProtectionJob(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionJobsApi.RunProtectionJob``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRunProtectionJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RunProtectionJobParam**](RunProtectionJobParam.md) | Specifies the type of backup. If not specified, the &#39;kRegular&#39; backup is run. | 

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


## UpdateProtectionJob

> ProtectionJob UpdateProtectionJob(ctx, id).Body(body).Execute()

Update a Protection Job.



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
    body := *openapiclient.NewProtectionJobRequestBody("Name_example", "PolicyId_example", []openapiclient.SourceSpecialParameter{*openapiclient.NewSourceSpecialParameter()}, NullableInt64(123)) // ProtectionJobRequestBody | Request to update a protection job.

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

    request := cohesitysdk.ApiUpdateProtectionJobRequest{
        id: &Id
        body: &Body
    }

    resp, r, err := api_client.ProtectionJobsApi.UpdateProtectionJob(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionJobsApi.UpdateProtectionJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProtectionJob`: ProtectionJob
    fmt.Fprintf(os.Stdout, "Response from `ProtectionJobsApi.UpdateProtectionJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the Protection Job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProtectionJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ProtectionJobRequestBody**](ProtectionJobRequestBody.md) | Request to update a protection job. | 

### Return type

[**ProtectionJob**](ProtectionJob.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProtectionJobsState

> UpdateProtectionJobsState UpdateProtectionJobsState(ctx).Body(body).Execute()

Perform an action like pause, resume, activate, deactivate on all the specified Protection Jobs.



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
    body := *openapiclient.NewUpdateProtectionJobsStateRequestBody() // UpdateProtectionJobsStateRequestBody |  (optional)

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

    request := cohesitysdk.ApiUpdateProtectionJobsStateRequest{
        body: &Body
    }

    resp, r, err := api_client.ProtectionJobsApi.UpdateProtectionJobsState(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionJobsApi.UpdateProtectionJobsState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProtectionJobsState`: UpdateProtectionJobsState
    fmt.Fprintf(os.Stdout, "Response from `ProtectionJobsApi.UpdateProtectionJobsState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProtectionJobsStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateProtectionJobsStateRequestBody**](UpdateProtectionJobsStateRequestBody.md) |  | 

### Return type

[**UpdateProtectionJobsState**](UpdateProtectionJobsState.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

