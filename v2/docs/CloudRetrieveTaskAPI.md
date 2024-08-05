# \CloudRetrieveTaskAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCloudRetrieveTask**](CloudRetrieveTaskAPI.md#CreateCloudRetrieveTask) | **Post** /data-protect/retrieve | Create a cloud retrieve task.
[**GetCloudRetrieveTaskByJobId**](CloudRetrieveTaskAPI.md#GetCloudRetrieveTaskByJobId) | **Get** /data-protect/retrieve/{jobId} | List details about the cloud retrieve task with the specific job id.
[**GetCloudRetrieveTasks**](CloudRetrieveTaskAPI.md#GetCloudRetrieveTasks) | **Get** /data-protect/retrieve | Get the list of cloud retrieve tasks.



## CreateCloudRetrieveTask

> CreateCloudRetrieveTaskRespBody CreateCloudRetrieveTask(ctx).Body(body).Execute()

Create a cloud retrieve task.



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
	body := *openapiclient.NewCreateCloudRetrieveTaskRequest() // CreateCloudRetrieveTaskRequest | Specifies the parameters to create a cloud retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudRetrieveTaskAPI.CreateCloudRetrieveTask(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudRetrieveTaskAPI.CreateCloudRetrieveTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCloudRetrieveTask`: CreateCloudRetrieveTaskRespBody
	fmt.Fprintf(os.Stdout, "Response from `CloudRetrieveTaskAPI.CreateCloudRetrieveTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCloudRetrieveTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateCloudRetrieveTaskRequest**](CreateCloudRetrieveTaskRequest.md) | Specifies the parameters to create a cloud retrieve. | 

### Return type

[**CreateCloudRetrieveTaskRespBody**](CreateCloudRetrieveTaskRespBody.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudRetrieveTaskByJobId

> CloudRetrieveTask GetCloudRetrieveTaskByJobId(ctx, jobId).Execute()

List details about the cloud retrieve task with the specific job id.



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
	jobId := int64(789) // int64 | Specifies a job id of the cloud retrieve task.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudRetrieveTaskAPI.GetCloudRetrieveTaskByJobId(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudRetrieveTaskAPI.GetCloudRetrieveTaskByJobId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudRetrieveTaskByJobId`: CloudRetrieveTask
	fmt.Fprintf(os.Stdout, "Response from `CloudRetrieveTaskAPI.GetCloudRetrieveTaskByJobId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **int64** | Specifies a job id of the cloud retrieve task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudRetrieveTaskByJobIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudRetrieveTask**](CloudRetrieveTask.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudRetrieveTasks

> CloudRetrieveTasks GetCloudRetrieveTasks(ctx).Execute()

Get the list of cloud retrieve tasks.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudRetrieveTaskAPI.GetCloudRetrieveTasks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudRetrieveTaskAPI.GetCloudRetrieveTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudRetrieveTasks`: CloudRetrieveTasks
	fmt.Fprintf(os.Stdout, "Response from `CloudRetrieveTaskAPI.GetCloudRetrieveTasks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudRetrieveTasksRequest struct via the builder pattern


### Return type

[**CloudRetrieveTasks**](CloudRetrieveTasks.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

