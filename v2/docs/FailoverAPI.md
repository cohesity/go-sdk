# \FailoverAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelFailover**](FailoverAPI.md#CancelFailover) | **Post** /data-protect/failover/{id}/cancel | Cancel failover workflow.
[**CancelViewFailover**](FailoverAPI.md#CancelViewFailover) | **Post** /data-protect/failover/views/{id}/cancel | Cancel View Failover Task.
[**CreatePlannedRun**](FailoverAPI.md#CreatePlannedRun) | **Post** /data-protect/failover/{id}/planned-run | Create a planned run for backup and replication.
[**CreateViewFailover**](FailoverAPI.md#CreateViewFailover) | **Post** /data-protect/failover/views/{id} | Create View Failover Task.
[**GetFailoverOps**](FailoverAPI.md#GetFailoverOps) | **Get** /data-protect/failover/views/{id}/operations | Gets all the failover operations which can be performed on this view.
[**GetTrackingViewId**](FailoverAPI.md#GetTrackingViewId) | **Get** /data-protect/failover/views/trackingViewId/{id} | Get tracking View Id
[**GetViewFailover**](FailoverAPI.md#GetViewFailover) | **Get** /data-protect/failover/views/{id} | Get View Failover.
[**InitFailover**](FailoverAPI.md#InitFailover) | **Post** /data-protect/failover/{id} | Initiate a failover request.
[**ObjectLinkage**](FailoverAPI.md#ObjectLinkage) | **Post** /data-protect/failover/{id}/object-linkage | Linking between replicated objects and failover objects
[**PollPlannedRuns**](FailoverAPI.md#PollPlannedRuns) | **Get** /data-protect/failover/planned-runs | Get the list of failover planned runs.
[**ReplicationBackupActivation**](FailoverAPI.md#ReplicationBackupActivation) | **Post** /data-protect/failover/{id}/backup-activation | Activate failover entity backup on replication clsuter.
[**SourceBackupDeactivation**](FailoverAPI.md#SourceBackupDeactivation) | **Post** /data-protect/failover/{id}/backup-deactivation | Deactivate failover entity backup on source clsuter.



## CancelFailover

> CancelFailover(ctx, id).Execute()

Cancel failover workflow.



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
	id := "id_example" // string | Specifies the id of the failover workflow.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FailoverAPI.CancelFailover(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FailoverAPI.CancelFailover``: %v\n", err)
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

> CancelViewFailover(ctx, id).Execute()

Cancel View Failover Task.



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
	id := int64(789) // int64 | Specifies a view id to cancel it's failover.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FailoverAPI.CancelViewFailover(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FailoverAPI.CancelViewFailover``: %v\n", err)
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

> FailoverCreateRunResponse CreatePlannedRun(ctx, id).Body(body).Execute()

Create a planned run for backup and replication.



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
	id := "id_example" // string | Specifies the id of the failover workflow.
	body := *openapiclient.NewFailoverRunConfiguration([]openapiclient.FailoverObject{*openapiclient.NewFailoverObject(NullableInt64(123))}, NullableInt64(123)) // FailoverRunConfiguration | Specifies the paramteres to create a planned run while failover workflow.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FailoverAPI.CreatePlannedRun(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FailoverAPI.CreatePlannedRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePlannedRun`: FailoverCreateRunResponse
	fmt.Fprintf(os.Stdout, "Response from `FailoverAPI.CreatePlannedRun`: %v\n", resp)
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

> Failover CreateViewFailover(ctx, id).Body(body).Execute()

Create View Failover Task.



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
	id := int64(789) // int64 | Specifies a view id to create an failover task.
	body := *openapiclient.NewCreateViewFailoverRequest("Type_example") // CreateViewFailoverRequest | Specifies the request body to create failover task.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FailoverAPI.CreateViewFailover(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FailoverAPI.CreateViewFailover``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateViewFailover`: Failover
	fmt.Fprintf(os.Stdout, "Response from `FailoverAPI.CreateViewFailover`: %v\n", resp)
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


## GetFailoverOps

> GetFailoverOpsResponse GetFailoverOps(ctx, id).Execute()

Gets all the failover operations which can be performed on this view.



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
	id := int64(789) // int64 | Specifies the view id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FailoverAPI.GetFailoverOps(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FailoverAPI.GetFailoverOps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFailoverOps`: GetFailoverOpsResponse
	fmt.Fprintf(os.Stdout, "Response from `FailoverAPI.GetFailoverOps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the view id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFailoverOpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetFailoverOpsResponse**](GetFailoverOpsResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrackingViewId

> GetTrackingViewIdResponse GetTrackingViewId(ctx, id).IsForwarded(isForwarded).Execute()

Get tracking View Id



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
	id := "id_example" // string | Specifies the view_uid of the source view.
	isForwarded := true // bool | Indicates whether the request is forwarded (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FailoverAPI.GetTrackingViewId(context.Background(), id).IsForwarded(isForwarded).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FailoverAPI.GetTrackingViewId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrackingViewId`: GetTrackingViewIdResponse
	fmt.Fprintf(os.Stdout, "Response from `FailoverAPI.GetTrackingViewId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the view_uid of the source view. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrackingViewIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isForwarded** | **bool** | Indicates whether the request is forwarded | 

### Return type

[**GetTrackingViewIdResponse**](GetTrackingViewIdResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViewFailover

> GetViewFailoverResponseBody GetViewFailover(ctx, id).Execute()

Get View Failover.



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
	id := int64(789) // int64 | Specifies a view id to create an failover task.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FailoverAPI.GetViewFailover(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FailoverAPI.GetViewFailover``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetViewFailover`: GetViewFailoverResponseBody
	fmt.Fprintf(os.Stdout, "Response from `FailoverAPI.GetViewFailover`: %v\n", resp)
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

> InitFailoverResponse InitFailover(ctx, id).Body(body).Execute()

Initiate a failover request.



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
	id := "id_example" // string | Specifies the id of the failover workflow.
	body := *openapiclient.NewInitFailoverRequest() // InitFailoverRequest | Specifies the parameters to initiate a failover. This failover request should be intiaited from replication cluster.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FailoverAPI.InitFailover(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FailoverAPI.InitFailover``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InitFailover`: InitFailoverResponse
	fmt.Fprintf(os.Stdout, "Response from `FailoverAPI.InitFailover`: %v\n", resp)
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

> ObjectLinkage(ctx, id).Body(body).Execute()

Linking between replicated objects and failover objects



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
	id := "id_example" // string | Specifies the id of the failover workflow.
	body := *openapiclient.NewObjectLinkingRequest() // ObjectLinkingRequest | Specifies the paramteres to create links between replicated objects and failover objects.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FailoverAPI.ObjectLinkage(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FailoverAPI.ObjectLinkage``: %v\n", err)
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

> FailoverRunsResponse PollPlannedRuns(ctx).FailoverIds(failoverIds).TenantIds(tenantIds).IncludeTenants(includeTenants).Execute()

Get the list of failover planned runs.



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
	failoverIds := []string{"Inner_example"} // []string | Get runs for specific failover workflows.
	tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
	includeTenants := true // bool | If true, the response will include Protection Groups which were created by all tenants which the current user has permission to see. If false, then only Protection Groups created by the current user will be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FailoverAPI.PollPlannedRuns(context.Background()).FailoverIds(failoverIds).TenantIds(tenantIds).IncludeTenants(includeTenants).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FailoverAPI.PollPlannedRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PollPlannedRuns`: FailoverRunsResponse
	fmt.Fprintf(os.Stdout, "Response from `FailoverAPI.PollPlannedRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPollPlannedRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **failoverIds** | **[]string** | Get runs for specific failover workflows. | 
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

> ReplicationBackupActivationResult ReplicationBackupActivation(ctx, id).Body(body).Execute()

Activate failover entity backup on replication clsuter.



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
	id := "id_example" // string | Specifies the id of the failover workflow.
	body := *openapiclient.NewReplicationBackupActivation() // ReplicationBackupActivation | Specifies the paramteres to activate the backup of failover entities.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FailoverAPI.ReplicationBackupActivation(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FailoverAPI.ReplicationBackupActivation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplicationBackupActivation`: ReplicationBackupActivationResult
	fmt.Fprintf(os.Stdout, "Response from `FailoverAPI.ReplicationBackupActivation`: %v\n", resp)
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

> SourceBackupDeactivation(ctx, id).Body(body).Execute()

Deactivate failover entity backup on source clsuter.



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
	id := "id_example" // string | Specifies the id of the failover workflow.
	body := *openapiclient.NewSourceBackupDeactivation() // SourceBackupDeactivation | Specifies the paramteres to deactivate the backup of failover entities.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FailoverAPI.SourceBackupDeactivation(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FailoverAPI.SourceBackupDeactivation``: %v\n", err)
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

