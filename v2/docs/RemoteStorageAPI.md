# \RemoteStorageAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRemoteStorageRegistration**](RemoteStorageAPI.md#DeleteRemoteStorageRegistration) | **Delete** /remote-storage/{id} | Delete Remote Storage Registration
[**GetRegisteredRemoteStorageList**](RemoteStorageAPI.md#GetRegisteredRemoteStorageList) | **Get** /remote-storage | Get Registered Remote Storage Servers List
[**GetRemoteStorageDetails**](RemoteStorageAPI.md#GetRemoteStorageDetails) | **Get** /remote-storage/{id} | Get remote storage details
[**RegisterNewRemoteStorage**](RemoteStorageAPI.md#RegisterNewRemoteStorage) | **Post** /remote-storage | Register Remote Storage
[**UpdateRemoteStorageRegistration**](RemoteStorageAPI.md#UpdateRemoteStorageRegistration) | **Patch** /remote-storage/{id} | Update Remote Storage Config



## DeleteRemoteStorageRegistration

> DeleteRemoteStorageRegistration(ctx, id).Execute()

Delete Remote Storage Registration



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
	id := int64(789) // int64 | Specifies the registration id of the registered remote storage.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RemoteStorageAPI.DeleteRemoteStorageRegistration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteStorageAPI.DeleteRemoteStorageRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the registration id of the registered remote storage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRemoteStorageRegistrationRequest struct via the builder pattern


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


## GetRegisteredRemoteStorageList

> RegisteredRemoteStorageList GetRegisteredRemoteStorageList(ctx).Execute()

Get Registered Remote Storage Servers List



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
	resp, r, err := apiClient.RemoteStorageAPI.GetRegisteredRemoteStorageList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteStorageAPI.GetRegisteredRemoteStorageList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRegisteredRemoteStorageList`: RegisteredRemoteStorageList
	fmt.Fprintf(os.Stdout, "Response from `RemoteStorageAPI.GetRegisteredRemoteStorageList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegisteredRemoteStorageListRequest struct via the builder pattern


### Return type

[**RegisteredRemoteStorageList**](RegisteredRemoteStorageList.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteStorageDetails

> RemoteStorageInfo GetRemoteStorageDetails(ctx, id).IncludeAvailableSpace(includeAvailableSpace).IncludeAvailableDataVips(includeAvailableDataVips).IncludeArrayInfo(includeArrayInfo).Execute()

Get remote storage details



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
	id := int64(789) // int64 | Specifies the id of the registered remote storage.
	includeAvailableSpace := true // bool | Specifies whether to include available capacity on remote storage. (optional) (default to false)
	includeAvailableDataVips := true // bool | Specifies whether to include available data vips on remote storage. (optional) (default to false)
	includeArrayInfo := true // bool | Includes flashblade specific info like name, software os and version of pure flashblade. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteStorageAPI.GetRemoteStorageDetails(context.Background(), id).IncludeAvailableSpace(includeAvailableSpace).IncludeAvailableDataVips(includeAvailableDataVips).IncludeArrayInfo(includeArrayInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteStorageAPI.GetRemoteStorageDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemoteStorageDetails`: RemoteStorageInfo
	fmt.Fprintf(os.Stdout, "Response from `RemoteStorageAPI.GetRemoteStorageDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the id of the registered remote storage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteStorageDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeAvailableSpace** | **bool** | Specifies whether to include available capacity on remote storage. | [default to false]
 **includeAvailableDataVips** | **bool** | Specifies whether to include available data vips on remote storage. | [default to false]
 **includeArrayInfo** | **bool** | Includes flashblade specific info like name, software os and version of pure flashblade. | [default to false]

### Return type

[**RemoteStorageInfo**](RemoteStorageInfo.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterNewRemoteStorage

> RemoteStorageInfo RegisterNewRemoteStorage(ctx).Body(body).Execute()

Register Remote Storage



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
	body := *openapiclient.NewRemoteStorageInfo("Product_example") // RemoteStorageInfo | Specifies the parameters to register a remote storage management server.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteStorageAPI.RegisterNewRemoteStorage(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteStorageAPI.RegisterNewRemoteStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterNewRemoteStorage`: RemoteStorageInfo
	fmt.Fprintf(os.Stdout, "Response from `RemoteStorageAPI.RegisterNewRemoteStorage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterNewRemoteStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RemoteStorageInfo**](RemoteStorageInfo.md) | Specifies the parameters to register a remote storage management server. | 

### Return type

[**RemoteStorageInfo**](RemoteStorageInfo.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRemoteStorageRegistration

> RemoteStorageInfo UpdateRemoteStorageRegistration(ctx, id).Body(body).Execute()

Update Remote Storage Config



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
	id := int64(789) // int64 | Specifies the registration id of the registered remote storage.
	body := *openapiclient.NewRemoteStorageInfo("Product_example") // RemoteStorageInfo | Specifies the parameters to update the registration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteStorageAPI.UpdateRemoteStorageRegistration(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteStorageAPI.UpdateRemoteStorageRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRemoteStorageRegistration`: RemoteStorageInfo
	fmt.Fprintf(os.Stdout, "Response from `RemoteStorageAPI.UpdateRemoteStorageRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the registration id of the registered remote storage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRemoteStorageRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RemoteStorageInfo**](RemoteStorageInfo.md) | Specifies the parameters to update the registration. | 

### Return type

[**RemoteStorageInfo**](RemoteStorageInfo.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

