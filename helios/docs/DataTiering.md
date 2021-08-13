# \DataTiering

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelDataTieringAnalysisGroupRun**](DataTiering.md#CancelDataTieringAnalysisGroupRun) | **Post** /data-tiering/analysis-groups/{id}/runs/{runId}/cancel | Cancel data tiering analysis run.
[**CancelDataTieringTaskRun**](DataTiering.md#CancelDataTieringTaskRun) | **Post** /data-tiering/tasks/{id}/runs/{runId}/cancel | Cancel data tiering task.
[**CreateDataTieringAnalysisGroup**](DataTiering.md#CreateDataTieringAnalysisGroup) | **Post** /data-tiering/analysis-groups | Create a data tiering analysis group.
[**CreateDataTieringAnalysisGroupRun**](DataTiering.md#CreateDataTieringAnalysisGroupRun) | **Post** /data-tiering/analysis-groups/{id}/runs | Create a data tiering analysis group run.
[**CreateDataTieringTask**](DataTiering.md#CreateDataTieringTask) | **Post** /data-tiering/tasks | Create a data tiering task.
[**CreateDataTieringTaskRun**](DataTiering.md#CreateDataTieringTaskRun) | **Post** /data-tiering/tasks/{id}/runs | Create a data tiering tasks run.
[**DeleteDataTieringAnalysisGroup**](DataTiering.md#DeleteDataTieringAnalysisGroup) | **Delete** /data-tiering/analysis-groups/{id} | Delete data tiering analysis group.
[**DeleteDataTieringTask**](DataTiering.md#DeleteDataTieringTask) | **Delete** /data-tiering/tasks/{id} | delete the data tiering task.
[**GetDataTieringAnalysisGroupById**](DataTiering.md#GetDataTieringAnalysisGroupById) | **Get** /data-tiering/analysis-groups/{id} | Get data tiering analysis group by id.
[**GetDataTieringAnalysisGroups**](DataTiering.md#GetDataTieringAnalysisGroups) | **Get** /data-tiering/analysis-groups | Get the list of data tiering analysis groups.
[**GetDataTieringTaskById**](DataTiering.md#GetDataTieringTaskById) | **Get** /data-tiering/tasks/{id} | Get data tiering task by id.
[**GetDataTieringTasks**](DataTiering.md#GetDataTieringTasks) | **Get** /data-tiering/tasks | Get the list of data tiering tasks.
[**UpdateDataTieringAnalysisGroup**](DataTiering.md#UpdateDataTieringAnalysisGroup) | **Put** /data-tiering/analysis-groups/{id} | Update a data tiering analysis group. Currently, it supports updating sources only.
[**UpdateDataTieringAnalysisGroupTagsConfig**](DataTiering.md#UpdateDataTieringAnalysisGroupTagsConfig) | **Put** /data-tiering/analysis-groups/{id}/config | Update data tiering analysis group config.
[**UpdateDataTieringAnalysisGroupsState**](DataTiering.md#UpdateDataTieringAnalysisGroupsState) | **Post** /data-tiering/analysis-groups/states | Update data tiering analysis groups state.
[**UpdateDataTieringTask**](DataTiering.md#UpdateDataTieringTask) | **Put** /data-tiering/tasks/{id} | Update a data tiering task.
[**UpdateDataTieringTasksState**](DataTiering.md#UpdateDataTieringTasksState) | **Post** /data-tiering/tasks/states | Update data tiering source analysis tasks state.



## CancelDataTieringAnalysisGroupRun

> CancelDataTieringAnalysisGroupRun(ctx, id, runId).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Cancel data tiering analysis run.



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
    id := "id_example" // string | Specifies a unique id of data tiering group.
    runId := "runId_example" // string | Specifies a unique run id of data tiering group run.
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

    request := helios.ApiCancelDataTieringAnalysisGroupRunRequest{
        Id: &id
        RunId: &runId
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.DataTiering.CancelDataTieringAnalysisGroupRun(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTiering.CancelDataTieringAnalysisGroupRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of data tiering group. | 
**runId** | **string** | Specifies a unique run id of data tiering group run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelDataTieringAnalysisGroupRunRequest struct via the builder pattern


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


## CancelDataTieringTaskRun

> CancelDataTieringTaskRun(ctx, id, runId).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Cancel data tiering task.



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
    id := "id_example" // string | Specifies a unique id of data tiering task.
    runId := "runId_example" // string | Specifies a unique run id of data tiering task.
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

    request := helios.ApiCancelDataTieringTaskRunRequest{
        Id: &id
        RunId: &runId
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.DataTiering.CancelDataTieringTaskRun(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTiering.CancelDataTieringTaskRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of data tiering task. | 
**runId** | **string** | Specifies a unique run id of data tiering task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelDataTieringTaskRunRequest struct via the builder pattern


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


## CreateDataTieringAnalysisGroup

> DataTieringAnalysisGroup CreateDataTieringAnalysisGroup(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Create a data tiering analysis group.



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
    body := *openapiclient.NewCommonDataTieringAnalysisGroupParams("Name_example", *openapiclient.NewDataTieringSource()) // CommonDataTieringAnalysisGroupParams | Specifies the data tiering analysis group.
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

    request := helios.ApiCreateDataTieringAnalysisGroupRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.DataTiering.CreateDataTieringAnalysisGroup(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTiering.CreateDataTieringAnalysisGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDataTieringAnalysisGroup`: DataTieringAnalysisGroup
    fmt.Fprintf(os.Stdout, "Response from `DataTiering.CreateDataTieringAnalysisGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataTieringAnalysisGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CommonDataTieringAnalysisGroupParams**](CommonDataTieringAnalysisGroupParams.md) | Specifies the data tiering analysis group. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**DataTieringAnalysisGroup**](DataTieringAnalysisGroup.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDataTieringAnalysisGroupRun

> CreateDataTieringAnalysisGroupRun(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Body(body).Execute()

Create a data tiering analysis group run.



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
    id := "id_example" // string | Specifies the id of the data tiering analysis group.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    body := *openapiclient.NewDataTieringAnalysisRunRequest() // DataTieringAnalysisRunRequest | Specifies the request to run analysis group once. (optional)

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

    request := helios.ApiCreateDataTieringAnalysisGroupRunRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        Body: &body
    }

    resp, r, err := api_client.DataTiering.CreateDataTieringAnalysisGroupRun(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTiering.CreateDataTieringAnalysisGroupRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id of the data tiering analysis group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataTieringAnalysisGroupRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **body** | [**DataTieringAnalysisRunRequest**](DataTieringAnalysisRunRequest.md) | Specifies the request to run analysis group once. | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDataTieringTask

> DataTieringTask CreateDataTieringTask(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Create a data tiering task.



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
    body := *openapiclient.NewCreateOrUpdateDataTieringTaskRequest("Name_example", "Type_example") // CreateOrUpdateDataTieringTaskRequest | Specifies the parameters to create a data tiering task.
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

    request := helios.ApiCreateDataTieringTaskRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.DataTiering.CreateDataTieringTask(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTiering.CreateDataTieringTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDataTieringTask`: DataTieringTask
    fmt.Fprintf(os.Stdout, "Response from `DataTiering.CreateDataTieringTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataTieringTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateOrUpdateDataTieringTaskRequest**](CreateOrUpdateDataTieringTaskRequest.md) | Specifies the parameters to create a data tiering task. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**DataTieringTask**](DataTieringTask.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDataTieringTaskRun

> CreateDataTieringTaskRun(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Body(body).Execute()

Create a data tiering tasks run.



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
    id := "id_example" // string | Specifies the id of the data tiering tasks.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    body := *openapiclient.NewDataTieringTaskRunRequest() // DataTieringTaskRunRequest | Specifies the request to run tiering task once. (optional)

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

    request := helios.ApiCreateDataTieringTaskRunRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        Body: &body
    }

    resp, r, err := api_client.DataTiering.CreateDataTieringTaskRun(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTiering.CreateDataTieringTaskRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id of the data tiering tasks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataTieringTaskRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **body** | [**DataTieringTaskRunRequest**](DataTieringTaskRunRequest.md) | Specifies the request to run tiering task once. | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDataTieringAnalysisGroup

> DeleteDataTieringAnalysisGroup(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Delete data tiering analysis group.



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
    id := "id_example" // string | Specifies a unique id of the data tiering analysis group.
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

    request := helios.ApiDeleteDataTieringAnalysisGroupRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.DataTiering.DeleteDataTieringAnalysisGroup(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTiering.DeleteDataTieringAnalysisGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the data tiering analysis group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDataTieringAnalysisGroupRequest struct via the builder pattern


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


## DeleteDataTieringTask

> DeleteDataTieringTask(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

delete the data tiering task.



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
    id := "id_example" // string | Specifies the id of the data tiering task.
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

    request := helios.ApiDeleteDataTieringTaskRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.DataTiering.DeleteDataTieringTask(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTiering.DeleteDataTieringTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id of the data tiering task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDataTieringTaskRequest struct via the builder pattern


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


## GetDataTieringAnalysisGroupById

> DataTieringAnalysisGroup GetDataTieringAnalysisGroupById(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Get data tiering analysis group by id.



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
    id := "id_example" // string | Specifies a unique id of the data tiering analysis group.
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

    request := helios.ApiGetDataTieringAnalysisGroupByIdRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.DataTiering.GetDataTieringAnalysisGroupById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTiering.GetDataTieringAnalysisGroupById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataTieringAnalysisGroupById`: DataTieringAnalysisGroup
    fmt.Fprintf(os.Stdout, "Response from `DataTiering.GetDataTieringAnalysisGroupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the data tiering analysis group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataTieringAnalysisGroupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**DataTieringAnalysisGroup**](DataTieringAnalysisGroup.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataTieringAnalysisGroups

> []DataTieringAnalysisGroup GetDataTieringAnalysisGroups(ctx).AccessClusterId(accessClusterId).RegionId(regionId).Ids(ids).Execute()

Get the list of data tiering analysis groups.



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
    ids := []string{"Inner_example"} // []string | Filter by a list of Analysis Group IDs. (optional)

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

    request := helios.ApiGetDataTieringAnalysisGroupsRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        Ids: &ids
    }

    resp, r, err := api_client.DataTiering.GetDataTieringAnalysisGroups(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTiering.GetDataTieringAnalysisGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataTieringAnalysisGroups`: []DataTieringAnalysisGroup
    fmt.Fprintf(os.Stdout, "Response from `DataTiering.GetDataTieringAnalysisGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDataTieringAnalysisGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **ids** | **[]string** | Filter by a list of Analysis Group IDs. | 

### Return type

[**[]DataTieringAnalysisGroup**](DataTieringAnalysisGroup.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataTieringTaskById

> DataTieringTask GetDataTieringTaskById(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Get data tiering task by id.



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
    id := "id_example" // string | Specifies the id of the data tiering task.
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

    request := helios.ApiGetDataTieringTaskByIdRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.DataTiering.GetDataTieringTaskById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTiering.GetDataTieringTaskById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataTieringTaskById`: DataTieringTask
    fmt.Fprintf(os.Stdout, "Response from `DataTiering.GetDataTieringTaskById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id of the data tiering task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataTieringTaskByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**DataTieringTask**](DataTieringTask.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataTieringTasks

> []DataTieringTask GetDataTieringTasks(ctx).AccessClusterId(accessClusterId).RegionId(regionId).Ids(ids).IncludeDowntieredDataLocation(includeDowntieredDataLocation).Execute()

Get the list of data tiering tasks.



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
    ids := []string{"Inner_example"} // []string | Filter by a list of data tiering task ids. (optional)
    includeDowntieredDataLocation := true // bool | If true, it will also return a list of downtiered data locations for downtiered tasks. (optional)

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

    request := helios.ApiGetDataTieringTasksRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        Ids: &ids
        IncludeDowntieredDataLocation: &includeDowntieredDataLocation
    }

    resp, r, err := api_client.DataTiering.GetDataTieringTasks(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTiering.GetDataTieringTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataTieringTasks`: []DataTieringTask
    fmt.Fprintf(os.Stdout, "Response from `DataTiering.GetDataTieringTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDataTieringTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **ids** | **[]string** | Filter by a list of data tiering task ids. | 
 **includeDowntieredDataLocation** | **bool** | If true, it will also return a list of downtiered data locations for downtiered tasks. | 

### Return type

[**[]DataTieringTask**](DataTieringTask.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDataTieringAnalysisGroup

> DataTieringAnalysisGroup UpdateDataTieringAnalysisGroup(ctx, id).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Update a data tiering analysis group. Currently, it supports updating sources only.



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
    id := "id_example" // string | Specifies a unique id of the data tiering analysis group.
    body := *openapiclient.NewCommonDataTieringAnalysisGroupParams("Name_example", *openapiclient.NewDataTieringSource()) // CommonDataTieringAnalysisGroupParams | Specifies the data tiering analysis group.
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

    request := helios.ApiUpdateDataTieringAnalysisGroupRequest{
        Id: &id
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.DataTiering.UpdateDataTieringAnalysisGroup(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTiering.UpdateDataTieringAnalysisGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDataTieringAnalysisGroup`: DataTieringAnalysisGroup
    fmt.Fprintf(os.Stdout, "Response from `DataTiering.UpdateDataTieringAnalysisGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the data tiering analysis group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataTieringAnalysisGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CommonDataTieringAnalysisGroupParams**](CommonDataTieringAnalysisGroupParams.md) | Specifies the data tiering analysis group. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**DataTieringAnalysisGroup**](DataTieringAnalysisGroup.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDataTieringAnalysisGroupTagsConfig

> DataTieringTagConfig UpdateDataTieringAnalysisGroupTagsConfig(ctx, id).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Update data tiering analysis group config.



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
    id := "id_example" // string | Specifies a unique id of the data tiering analysis group.
    body := *openapiclient.NewDataTieringTagConfig([]openapiclient.DataTieringTagObject{*openapiclient.NewDataTieringTagObject()}) // DataTieringTagConfig | Specifies the data tiering analysis Tags Config.
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

    request := helios.ApiUpdateDataTieringAnalysisGroupTagsConfigRequest{
        Id: &id
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.DataTiering.UpdateDataTieringAnalysisGroupTagsConfig(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTiering.UpdateDataTieringAnalysisGroupTagsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDataTieringAnalysisGroupTagsConfig`: DataTieringTagConfig
    fmt.Fprintf(os.Stdout, "Response from `DataTiering.UpdateDataTieringAnalysisGroupTagsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the data tiering analysis group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataTieringAnalysisGroupTagsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DataTieringTagConfig**](DataTieringTagConfig.md) | Specifies the data tiering analysis Tags Config. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**DataTieringTagConfig**](DataTieringTagConfig.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDataTieringAnalysisGroupsState

> UpdateDataTieringState UpdateDataTieringAnalysisGroupsState(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Update data tiering analysis groups state.



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
    body := *openapiclient.NewUpdateDataTieringStateRequest("Action_example", []string{"Ids_example"}) // UpdateDataTieringStateRequest | Specifies the parameters to perform an action of list of data tiering analysis groups.
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

    request := helios.ApiUpdateDataTieringAnalysisGroupsStateRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.DataTiering.UpdateDataTieringAnalysisGroupsState(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTiering.UpdateDataTieringAnalysisGroupsState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDataTieringAnalysisGroupsState`: UpdateDataTieringState
    fmt.Fprintf(os.Stdout, "Response from `DataTiering.UpdateDataTieringAnalysisGroupsState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataTieringAnalysisGroupsStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateDataTieringStateRequest**](UpdateDataTieringStateRequest.md) | Specifies the parameters to perform an action of list of data tiering analysis groups. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**UpdateDataTieringState**](UpdateDataTieringState.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDataTieringTask

> DataTieringTask UpdateDataTieringTask(ctx, id).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Update a data tiering task.



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
    id := "id_example" // string | Specifies the id of the data tiering task.
    body := *openapiclient.NewCreateOrUpdateDataTieringTaskRequest("Name_example", "Type_example") // CreateOrUpdateDataTieringTaskRequest | Specifies the parameters to update a data tiering task.
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

    request := helios.ApiUpdateDataTieringTaskRequest{
        Id: &id
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.DataTiering.UpdateDataTieringTask(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTiering.UpdateDataTieringTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDataTieringTask`: DataTieringTask
    fmt.Fprintf(os.Stdout, "Response from `DataTiering.UpdateDataTieringTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id of the data tiering task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataTieringTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateOrUpdateDataTieringTaskRequest**](CreateOrUpdateDataTieringTaskRequest.md) | Specifies the parameters to update a data tiering task. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**DataTieringTask**](DataTieringTask.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDataTieringTasksState

> UpdateDataTieringState UpdateDataTieringTasksState(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Update data tiering source analysis tasks state.



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
    body := *openapiclient.NewUpdateDataTieringStateRequest("Action_example", []string{"Ids_example"}) // UpdateDataTieringStateRequest | Specifies the parameters to perform an action of list of data tiering tasks.
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

    request := helios.ApiUpdateDataTieringTasksStateRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.DataTiering.UpdateDataTieringTasksState(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTiering.UpdateDataTieringTasksState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDataTieringTasksState`: UpdateDataTieringState
    fmt.Fprintf(os.Stdout, "Response from `DataTiering.UpdateDataTieringTasksState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataTieringTasksStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateDataTieringStateRequest**](UpdateDataTieringStateRequest.md) | Specifies the parameters to perform an action of list of data tiering tasks. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**UpdateDataTieringState**](UpdateDataTieringState.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

