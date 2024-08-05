# \BaseosPatchManagementAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyBaseosPatch**](BaseosPatchManagementAPI.md#ApplyBaseosPatch) | **Post** /patch-management/baseos-patch-apply | Applies the given baseos patch.
[**DownloadBaseosPatch**](BaseosPatchManagementAPI.md#DownloadBaseosPatch) | **Post** /patch-management/baseos-patch-download | Downloads the given baseos patch.
[**GetBaseosPatchList**](BaseosPatchManagementAPI.md#GetBaseosPatchList) | **Get** /patch-management/baseos-patch-list | Get available baseos patches
[**GetBaseosPatchLog**](BaseosPatchManagementAPI.md#GetBaseosPatchLog) | **Get** /patch-management/baseos-patch-log | Get Baseos patch application log
[**RemoveBaseosPatch**](BaseosPatchManagementAPI.md#RemoveBaseosPatch) | **Post** /patch-management/baseos-patch-remove | Cleans up the given baseos patch files.



## ApplyBaseosPatch

> ApplyBaseosPatch(ctx).Body(body).Execute()

Applies the given baseos patch.



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
	body := *openapiclient.NewApplyBaseosPatchRequest("PatchName_example") // ApplyBaseosPatchRequest | Request to apply a baseos patch.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BaseosPatchManagementAPI.ApplyBaseosPatch(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaseosPatchManagementAPI.ApplyBaseosPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplyBaseosPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApplyBaseosPatchRequest**](ApplyBaseosPatchRequest.md) | Request to apply a baseos patch. | 

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


## DownloadBaseosPatch

> DownloadBaseosPatch(ctx).Body(body).Execute()

Downloads the given baseos patch.



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
	body := *openapiclient.NewDownloadBaseosPatchRequest("PatchUrl_example") // DownloadBaseosPatchRequest | Request to download a new baseos patch.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BaseosPatchManagementAPI.DownloadBaseosPatch(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaseosPatchManagementAPI.DownloadBaseosPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadBaseosPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DownloadBaseosPatchRequest**](DownloadBaseosPatchRequest.md) | Request to download a new baseos patch. | 

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


## GetBaseosPatchList

> []BaseosPatchListItem GetBaseosPatchList(ctx).Execute()

Get available baseos patches



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
	resp, r, err := apiClient.BaseosPatchManagementAPI.GetBaseosPatchList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaseosPatchManagementAPI.GetBaseosPatchList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBaseosPatchList`: []BaseosPatchListItem
	fmt.Fprintf(os.Stdout, "Response from `BaseosPatchManagementAPI.GetBaseosPatchList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBaseosPatchListRequest struct via the builder pattern


### Return type

[**[]BaseosPatchListItem**](BaseosPatchListItem.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBaseosPatchLog

> BaseosPatchLog GetBaseosPatchLog(ctx).PatchName(patchName).Execute()

Get Baseos patch application log



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
	patchName := "patchName_example" // string | Name of the hotfix with security patch

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaseosPatchManagementAPI.GetBaseosPatchLog(context.Background()).PatchName(patchName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaseosPatchManagementAPI.GetBaseosPatchLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBaseosPatchLog`: BaseosPatchLog
	fmt.Fprintf(os.Stdout, "Response from `BaseosPatchManagementAPI.GetBaseosPatchLog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBaseosPatchLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchName** | **string** | Name of the hotfix with security patch | 

### Return type

[**BaseosPatchLog**](BaseosPatchLog.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveBaseosPatch

> RemoveBaseosPatch(ctx).Body(body).Execute()

Cleans up the given baseos patch files.



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
	body := *openapiclient.NewRemoveBaseosPatchRequest("PatchName_example") // RemoveBaseosPatchRequest | Request to remove baseos patch files.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BaseosPatchManagementAPI.RemoveBaseosPatch(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaseosPatchManagementAPI.RemoveBaseosPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveBaseosPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RemoveBaseosPatchRequest**](RemoveBaseosPatchRequest.md) | Request to remove baseos patch files. | 

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

