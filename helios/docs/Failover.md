# \Failover

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelFailover**](Failover.md#CancelFailover) | **Post** /data-protect/failover/{id}/cancel | Cancel failover workflow.
[**CancelViewFailover**](Failover.md#CancelViewFailover) | **Post** /data-protect/failover/views/{id}/cancel | Cancel View Failover Task.
[**CreatePlannedRun**](Failover.md#CreatePlannedRun) | **Post** /data-protect/failover/{id}/planned-run | Create a planned run for backup and replication.
[**CreateViewFailover**](Failover.md#CreateViewFailover) | **Post** /data-protect/failover/views/{id} | Create View Failover Task.
[**GetViewFailover**](Failover.md#GetViewFailover) | **Get** /data-protect/failover/views/{id} | Get View Failover.
[**InitFailover**](Failover.md#InitFailover) | **Post** /data-protect/failover/{id} | Initiate a failover request.
[**ObjectLinkage**](Failover.md#ObjectLinkage) | **Post** /data-protect/failover/{id}/object-linkage | Linking between replicated objects and failover objects
[**PollPlannedRuns**](Failover.md#PollPlannedRuns) | **Get** /data-protect/failover/planned-runs | Get the list of failover planned runs.
[**ReplicationBackupActivation**](Failover.md#ReplicationBackupActivation) | **Post** /data-protect/failover/{id}/backup-activation | Activate failover entity backup on replication clsuter.
[**SourceBackupDeactivation**](Failover.md#SourceBackupDeactivation) | **Post** /data-protect/failover/{id}/backup-deactivation | Deactivate failover entity backup on source clsuter.



## CancelFailover

> CancelFailover(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Cancel failover workflow.



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
    id := "id_example" // string | Specifies the id of the failover workflow.
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

    request := helios.ApiCancelFailoverRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Failover.CancelFailover(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Failover.CancelFailover``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id of the failover workflow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelFailoverRequest struct via the builder pattern


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


## CancelViewFailover

> CancelViewFailover(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Cancel View Failover Task.



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
    id := int64(789) // int64 | Specifies a view id to cancel it's failover.
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

    request := helios.ApiCancelViewFailoverRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Failover.CancelViewFailover(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Failover.CancelViewFailover``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a view id to cancel it&#39;s failover. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelViewFailoverRequest struct via the builder pattern


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


## CreatePlannedRun

> FailoverCreateRunResponse CreatePlannedRun(ctx, id).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Create a planned run for backup and replication.



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
    id := "id_example" // string | Specifies the id of the failover workflow.
    body := *openapiclient.NewFailoverRunConfiguration(NullableInt64(123), []openapiclient.FailoverObject{*openapiclient.NewFailoverObject(NullableInt64(123))}, "ProtectionGroupId_example") // FailoverRunConfiguration | Specifies the paramteres to create a planned run while failover workflow.
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

    request := helios.ApiCreatePlannedRunRequest{
        Id: &id
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Failover.CreatePlannedRun(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Failover.CreatePlannedRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePlannedRun`: FailoverCreateRunResponse
    fmt.Fprintf(os.Stdout, "Response from `Failover.CreatePlannedRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id of the failover workflow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlannedRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**FailoverRunConfiguration**](FailoverRunConfiguration.md) | Specifies the paramteres to create a planned run while failover workflow. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**FailoverCreateRunResponse**](FailoverCreateRunResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateViewFailover

> Failover CreateViewFailover(ctx, id).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Create View Failover Task.



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
    id := int64(789) // int64 | Specifies a view id to create an failover task.
    body := *openapiclient.NewCreateViewFailoverRequest("Type_example") // CreateViewFailoverRequest | Specifies the request body to create failover task.
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

    request := helios.ApiCreateViewFailoverRequest{
        Id: &id
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Failover.CreateViewFailover(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Failover.CreateViewFailover``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateViewFailover`: Failover
    fmt.Fprintf(os.Stdout, "Response from `Failover.CreateViewFailover`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a view id to create an failover task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateViewFailoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateViewFailoverRequest**](CreateViewFailoverRequest.md) | Specifies the request body to create failover task. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**Failover**](Failover.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViewFailover

> GetViewFailoverResponseBody GetViewFailover(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Get View Failover.



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
    id := int64(789) // int64 | Specifies a view id to create an failover task.
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

    request := helios.ApiGetViewFailoverRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Failover.GetViewFailover(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Failover.GetViewFailover``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetViewFailover`: GetViewFailoverResponseBody
    fmt.Fprintf(os.Stdout, "Response from `Failover.GetViewFailover`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a view id to create an failover task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetViewFailoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**GetViewFailoverResponseBody**](GetViewFailoverResponseBody.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitFailover

> InitFailoverResponse InitFailover(ctx, id).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Initiate a failover request.



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
    id := "id_example" // string | Specifies the id of the failover workflow.
    body := *openapiclient.NewInitFailoverRequest() // InitFailoverRequest | Specifies the parameters to initiate a failover. This failover request should be intiaited from replication cluster.
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

    request := helios.ApiInitFailoverRequest{
        Id: &id
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Failover.InitFailover(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Failover.InitFailover``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InitFailover`: InitFailoverResponse
    fmt.Fprintf(os.Stdout, "Response from `Failover.InitFailover`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id of the failover workflow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInitFailoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InitFailoverRequest**](InitFailoverRequest.md) | Specifies the parameters to initiate a failover. This failover request should be intiaited from replication cluster. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**InitFailoverResponse**](InitFailoverResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ObjectLinkage

> ObjectLinkage(ctx, id).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Linking between replicated objects and failover objects



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
    id := "id_example" // string | Specifies the id of the failover workflow.
    body := *openapiclient.NewObjectLinkingRequest() // ObjectLinkingRequest | Specifies the paramteres to create links between replicated objects and failover objects.
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

    request := helios.ApiObjectLinkageRequest{
        Id: &id
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Failover.ObjectLinkage(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Failover.ObjectLinkage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id of the failover workflow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiObjectLinkageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ObjectLinkingRequest**](ObjectLinkingRequest.md) | Specifies the paramteres to create links between replicated objects and failover objects. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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


## PollPlannedRuns

> FailoverRunsResponse PollPlannedRuns(ctx).FailoverIds(failoverIds).AccessClusterId(accessClusterId).RegionId(regionId).TenantIds(tenantIds).IncludeTenants(includeTenants).Execute()

Get the list of failover planned runs.



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
    failoverIds := []string{"Inner_example"} // []string | Get runs for specific failover workflows.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
    includeTenants := true // bool | If true, the response will include Protection Groups which were created by all tenants which the current user has permission to see. If false, then only Protection Groups created by the current user will be returned. (optional)

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

    request := helios.ApiPollPlannedRunsRequest{
        FailoverIds: &failoverIds
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        TenantIds: &tenantIds
        IncludeTenants: &includeTenants
    }

    resp, r, err := api_client.Failover.PollPlannedRuns(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Failover.PollPlannedRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PollPlannedRuns`: FailoverRunsResponse
    fmt.Fprintf(os.Stdout, "Response from `Failover.PollPlannedRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPollPlannedRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **failoverIds** | **[]string** | Get runs for specific failover workflows. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **includeTenants** | **bool** | If true, the response will include Protection Groups which were created by all tenants which the current user has permission to see. If false, then only Protection Groups created by the current user will be returned. | 

### Return type

[**FailoverRunsResponse**](FailoverRunsResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplicationBackupActivation

> ReplicationBackupActivationResult ReplicationBackupActivation(ctx, id).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Activate failover entity backup on replication clsuter.



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
    id := "id_example" // string | Specifies the id of the failover workflow.
    body := *openapiclient.NewReplicationBackupActivation() // ReplicationBackupActivation | Specifies the paramteres to activate the backup of failover entities.
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

    request := helios.ApiReplicationBackupActivationRequest{
        Id: &id
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Failover.ReplicationBackupActivation(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Failover.ReplicationBackupActivation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplicationBackupActivation`: ReplicationBackupActivationResult
    fmt.Fprintf(os.Stdout, "Response from `Failover.ReplicationBackupActivation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id of the failover workflow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplicationBackupActivationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ReplicationBackupActivation**](ReplicationBackupActivation.md) | Specifies the paramteres to activate the backup of failover entities. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**ReplicationBackupActivationResult**](ReplicationBackupActivationResult.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceBackupDeactivation

> SourceBackupDeactivation(ctx, id).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Deactivate failover entity backup on source clsuter.



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
    id := "id_example" // string | Specifies the id of the failover workflow.
    body := *openapiclient.NewSourceBackupDeactivation() // SourceBackupDeactivation | Specifies the paramteres to deactivate the backup of failover entities.
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

    request := helios.ApiSourceBackupDeactivationRequest{
        Id: &id
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Failover.SourceBackupDeactivation(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Failover.SourceBackupDeactivation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id of the failover workflow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSourceBackupDeactivationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SourceBackupDeactivation**](SourceBackupDeactivation.md) | Specifies the paramteres to deactivate the backup of failover entities. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

