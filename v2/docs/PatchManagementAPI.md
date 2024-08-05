# \PatchManagementAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyPatches**](PatchManagementAPI.md#ApplyPatches) | **Post** /patch-management/available-patches | Apply patches
[**GetAppliedPatches**](PatchManagementAPI.md#GetAppliedPatches) | **Get** /patch-management/applied-patches | Get applied patches
[**GetAvailablePatches**](PatchManagementAPI.md#GetAvailablePatches) | **Get** /patch-management/available-patches | Get available patches
[**GetPatchOperationStatus**](PatchManagementAPI.md#GetPatchOperationStatus) | **Get** /patch-management/operation-status | Get patch operation status
[**GetPatchesHistory**](PatchManagementAPI.md#GetPatchesHistory) | **Get** /patch-management/patches-history | Get patches history
[**ImportPatches**](PatchManagementAPI.md#ImportPatches) | **Put** /patch-management/available-patches | Import patches
[**RemovePatch**](PatchManagementAPI.md#RemovePatch) | **Post** /patch-management/patch-remove | Cleans up the given patch file in data directory.
[**RevertPatches**](PatchManagementAPI.md#RevertPatches) | **Post** /patch-management/applied-patches | Revert patches



## ApplyPatches

> []ServicePatchLevel ApplyPatches(ctx).Body(body).Execute()

Apply patches



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
	body := *openapiclient.NewApplyPatchesRequest() // ApplyPatchesRequest | Request to apply patches.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchManagementAPI.ApplyPatches(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchManagementAPI.ApplyPatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyPatches`: []ServicePatchLevel
	fmt.Fprintf(os.Stdout, "Response from `PatchManagementAPI.ApplyPatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplyPatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApplyPatchesRequest**](ApplyPatchesRequest.md) | Request to apply patches. | 

### Return type

[**[]ServicePatchLevel**](ServicePatchLevel.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppliedPatches

> []AppliedPatch GetAppliedPatches(ctx).Service(service).IncludeDetails(includeDetails).Execute()

Get applied patches



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
	service := "service_example" // string | Specifies optional service name whose current patch is returned. If it is not specified, all the applied patches are returned. (optional)
	includeDetails := true // bool | Specifies whether to return the details of all the fixes in the patch. By default, returns only the most recent fix made for the service in the patch. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchManagementAPI.GetAppliedPatches(context.Background()).Service(service).IncludeDetails(includeDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchManagementAPI.GetAppliedPatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppliedPatches`: []AppliedPatch
	fmt.Fprintf(os.Stdout, "Response from `PatchManagementAPI.GetAppliedPatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAppliedPatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service** | **string** | Specifies optional service name whose current patch is returned. If it is not specified, all the applied patches are returned. | 
 **includeDetails** | **bool** | Specifies whether to return the details of all the fixes in the patch. By default, returns only the most recent fix made for the service in the patch. | 

### Return type

[**[]AppliedPatch**](AppliedPatch.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailablePatches

> []AvailablePatch GetAvailablePatches(ctx).Service(service).IncludeDetails(includeDetails).Execute()

Get available patches



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
	service := "service_example" // string | Specifies optional service name whose available patch is returned. If it is not specified, available patches for all the serivces are returned. (optional)
	includeDetails := true // bool | Specifies whether to return the description of all the fixes in the patch. By default, returns only the most recent fix made for the service in the patch. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchManagementAPI.GetAvailablePatches(context.Background()).Service(service).IncludeDetails(includeDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchManagementAPI.GetAvailablePatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvailablePatches`: []AvailablePatch
	fmt.Fprintf(os.Stdout, "Response from `PatchManagementAPI.GetAvailablePatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailablePatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service** | **string** | Specifies optional service name whose available patch is returned. If it is not specified, available patches for all the serivces are returned. | 
 **includeDetails** | **bool** | Specifies whether to return the description of all the fixes in the patch. By default, returns only the most recent fix made for the service in the patch. | 

### Return type

[**[]AvailablePatch**](AvailablePatch.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatchOperationStatus

> PatchOperationStatus GetPatchOperationStatus(ctx).IncludeDetails(includeDetails).Execute()

Get patch operation status



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
	includeDetails := true // bool | Specifies whether to return details of all service patch opertions on all nodes. By default, returns whether there is a patch operation in progress or not. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchManagementAPI.GetPatchOperationStatus(context.Background()).IncludeDetails(includeDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchManagementAPI.GetPatchOperationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatchOperationStatus`: PatchOperationStatus
	fmt.Fprintf(os.Stdout, "Response from `PatchManagementAPI.GetPatchOperationStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPatchOperationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeDetails** | **bool** | Specifies whether to return details of all service patch opertions on all nodes. By default, returns whether there is a patch operation in progress or not. | 

### Return type

[**PatchOperationStatus**](PatchOperationStatus.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatchesHistory

> []PatchOperation GetPatchesHistory(ctx).Service(service).Execute()

Get patches history



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
	service := "service_example" // string | Specifies optional service name whose patch operation history is returned. If it is not specified, patch operations of all the serivces are returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchManagementAPI.GetPatchesHistory(context.Background()).Service(service).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchManagementAPI.GetPatchesHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatchesHistory`: []PatchOperation
	fmt.Fprintf(os.Stdout, "Response from `PatchManagementAPI.GetPatchesHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPatchesHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service** | **string** | Specifies optional service name whose patch operation history is returned. If it is not specified, patch operations of all the serivces are returned. | 

### Return type

[**[]PatchOperation**](PatchOperation.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportPatches

> []PatchDetail ImportPatches(ctx).FileName(fileName).Checksum(checksum).Patch(patch).IsUpgradeAndPatch(isUpgradeAndPatch).Execute()

Import patches



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
	fileName := "fileName_example" // string | 
	checksum := "checksum_example" // string | 
	patch := os.NewFile(1234, "some_file") // *os.File | 
	isUpgradeAndPatch := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchManagementAPI.ImportPatches(context.Background()).FileName(fileName).Checksum(checksum).Patch(patch).IsUpgradeAndPatch(isUpgradeAndPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchManagementAPI.ImportPatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportPatches`: []PatchDetail
	fmt.Fprintf(os.Stdout, "Response from `PatchManagementAPI.ImportPatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportPatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileName** | **string** |  | 
 **checksum** | **string** |  | 
 **patch** | ***os.File** |  | 
 **isUpgradeAndPatch** | **bool** |  | [default to false]

### Return type

[**[]PatchDetail**](PatchDetail.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemovePatch

> RemovePatch(ctx).Body(body).Execute()

Cleans up the given patch file in data directory.



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
	body := *openapiclient.NewRemovePatchRequest("PatchName_example") // RemovePatchRequest | Request to remove the patch file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PatchManagementAPI.RemovePatch(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchManagementAPI.RemovePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemovePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RemovePatchRequest**](RemovePatchRequest.md) | Request to remove the patch file. | 

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


## RevertPatches

> []ServicePatchLevel RevertPatches(ctx).Body(body).Execute()

Revert patches



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
	body := *openapiclient.NewRevertPatchesRequest("Service_example") // RevertPatchesRequest | Request to revert patches.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchManagementAPI.RevertPatches(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchManagementAPI.RevertPatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RevertPatches`: []ServicePatchLevel
	fmt.Fprintf(os.Stdout, "Response from `PatchManagementAPI.RevertPatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevertPatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RevertPatchesRequest**](RevertPatchesRequest.md) | Request to revert patches. | 

### Return type

[**[]ServicePatchLevel**](ServicePatchLevel.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

