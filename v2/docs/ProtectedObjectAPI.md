# \ProtectedObjectAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PerformActionOnProtectObjects**](ProtectedObjectAPI.md#PerformActionOnProtectObjects) | **Post** /data-protect/protected-objects/actions | Perform Actions on Protect Objects.
[**ProtectObjectsOfAnyType**](ProtectedObjectAPI.md#ProtectObjectsOfAnyType) | **Post** /data-protect/protected-objects | Create Object Backup.
[**UpdateProtectedObjectsOfAnyType**](ProtectedObjectAPI.md#UpdateProtectedObjectsOfAnyType) | **Put** /data-protect/protected-objects/{id} | Update Object Backup.



## PerformActionOnProtectObjects

> ProtectedObjectActionResponse PerformActionOnProtectObjects(ctx).Body(body).Execute()

Perform Actions on Protect Objects.



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
	body := *openapiclient.NewProtectdObjectsActionRequest("Action_example") // ProtectdObjectsActionRequest | Specifies the parameters to perform an action on an already protected object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProtectedObjectAPI.PerformActionOnProtectObjects(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtectedObjectAPI.PerformActionOnProtectObjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PerformActionOnProtectObjects`: ProtectedObjectActionResponse
	fmt.Fprintf(os.Stdout, "Response from `ProtectedObjectAPI.PerformActionOnProtectObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPerformActionOnProtectObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProtectdObjectsActionRequest**](ProtectdObjectsActionRequest.md) | Specifies the parameters to perform an action on an already protected object. | 

### Return type

[**ProtectedObjectActionResponse**](ProtectedObjectActionResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProtectObjectsOfAnyType

> CreateProtectedObjectsResponse ProtectObjectsOfAnyType(ctx).Body(body).RequestInitiatorType(requestInitiatorType).Execute()

Create Object Backup.



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
	body := *openapiclient.NewCreateProtectedObjectsRequest([]openapiclient.EnvSpecificObjectProtectionRequestParams{*openapiclient.NewEnvSpecificObjectProtectionRequestParams()}) // CreateProtectedObjectsRequest | Specifies the parameters to protect objects.
	requestInitiatorType := "requestInitiatorType_example" // string | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProtectedObjectAPI.ProtectObjectsOfAnyType(context.Background()).Body(body).RequestInitiatorType(requestInitiatorType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtectedObjectAPI.ProtectObjectsOfAnyType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProtectObjectsOfAnyType`: CreateProtectedObjectsResponse
	fmt.Fprintf(os.Stdout, "Response from `ProtectedObjectAPI.ProtectObjectsOfAnyType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProtectObjectsOfAnyTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateProtectedObjectsRequest**](CreateProtectedObjectsRequest.md) | Specifies the parameters to protect objects. | 
 **requestInitiatorType** | **string** | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | 

### Return type

[**CreateProtectedObjectsResponse**](CreateProtectedObjectsResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProtectedObjectsOfAnyType

> GetProtectedObjectResponse UpdateProtectedObjectsOfAnyType(ctx, id).Body(body).RequestInitiatorType(requestInitiatorType).Execute()

Update Object Backup.



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
	id := int64(789) // int64 | Specifies the id of the Protected Object.
	body := *openapiclient.NewUpdateProtectedObjectsRequest() // UpdateProtectedObjectsRequest | Specifies the parameters to perform an update on protected objects.
	requestInitiatorType := "requestInitiatorType_example" // string | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProtectedObjectAPI.UpdateProtectedObjectsOfAnyType(context.Background(), id).Body(body).RequestInitiatorType(requestInitiatorType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtectedObjectAPI.UpdateProtectedObjectsOfAnyType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProtectedObjectsOfAnyType`: GetProtectedObjectResponse
	fmt.Fprintf(os.Stdout, "Response from `ProtectedObjectAPI.UpdateProtectedObjectsOfAnyType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the id of the Protected Object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProtectedObjectsOfAnyTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateProtectedObjectsRequest**](UpdateProtectedObjectsRequest.md) | Specifies the parameters to perform an update on protected objects. | 
 **requestInitiatorType** | **string** | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | 

### Return type

[**GetProtectedObjectResponse**](GetProtectedObjectResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

