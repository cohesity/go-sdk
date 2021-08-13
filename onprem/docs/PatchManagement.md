# \PatchManagement

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyPatches**](PatchManagement.md#ApplyPatches) | **Post** /patch-management/available-patches | Apply patches
[**GetAppliedPatches**](PatchManagement.md#GetAppliedPatches) | **Get** /patch-management/applied-patches | Get applied patches
[**GetAvailablePatches**](PatchManagement.md#GetAvailablePatches) | **Get** /patch-management/available-patches | Get available patches
[**GetPatchOperationStatus**](PatchManagement.md#GetPatchOperationStatus) | **Get** /patch-management/operation-status | Get patch operation status
[**GetPatchesHistory**](PatchManagement.md#GetPatchesHistory) | **Get** /patch-management/patches-history | Get patches history
[**ImportPatches**](PatchManagement.md#ImportPatches) | **Put** /patch-management/available-patches | Import patches
[**RevertPatches**](PatchManagement.md#RevertPatches) | **Post** /patch-management/applied-patches | Revert patches



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
    onprem "onprem"
)

func main() {
    body := *openapiclient.NewApplyPatchesRequest() // ApplyPatchesRequest | Request to apply patches.

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

    request := onprem.ApiApplyPatchesRequest{
        Body: &body
    }

    resp, r, err := api_client.PatchManagement.ApplyPatches(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchManagement.ApplyPatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplyPatches`: []ServicePatchLevel
    fmt.Fprintf(os.Stdout, "Response from `PatchManagement.ApplyPatches`: %v\n", resp)
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
    onprem "onprem"
)

func main() {
    service := "service_example" // string | Specifies optional service name whose current patch is returned. If it is not specified, all the applied patches are returned. (optional)
    includeDetails := true // bool | Specifies whether to return the details of all the fixes in the patch. By default, returns only the most recent fix made for the service in the patch. (optional)

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

    request := onprem.ApiGetAppliedPatchesRequest{
        Service: &service
        IncludeDetails: &includeDetails
    }

    resp, r, err := api_client.PatchManagement.GetAppliedPatches(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchManagement.GetAppliedPatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppliedPatches`: []AppliedPatch
    fmt.Fprintf(os.Stdout, "Response from `PatchManagement.GetAppliedPatches`: %v\n", resp)
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
    onprem "onprem"
)

func main() {
    service := "service_example" // string | Specifies optional service name whose available patch is returned. If it is not specified, available patches for all the serivces are returned. (optional)
    includeDetails := true // bool | Specifies whether to return the description of all the fixes in the patch. By default, returns only the most recent fix made for the service in the patch. (optional)

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

    request := onprem.ApiGetAvailablePatchesRequest{
        Service: &service
        IncludeDetails: &includeDetails
    }

    resp, r, err := api_client.PatchManagement.GetAvailablePatches(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchManagement.GetAvailablePatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvailablePatches`: []AvailablePatch
    fmt.Fprintf(os.Stdout, "Response from `PatchManagement.GetAvailablePatches`: %v\n", resp)
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
    onprem "onprem"
)

func main() {
    includeDetails := true // bool | Specifies whether to return details of all service patch opertions on all nodes. By default, returns whether there is a patch operation in progress or not. (optional)

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

    request := onprem.ApiGetPatchOperationStatusRequest{
        IncludeDetails: &includeDetails
    }

    resp, r, err := api_client.PatchManagement.GetPatchOperationStatus(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchManagement.GetPatchOperationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPatchOperationStatus`: PatchOperationStatus
    fmt.Fprintf(os.Stdout, "Response from `PatchManagement.GetPatchOperationStatus`: %v\n", resp)
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
    onprem "onprem"
)

func main() {
    service := "service_example" // string | Specifies optional service name whose patch operation history is returned. If it is not specified, patch operations of all the serivces are returned. (optional)

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

    request := onprem.ApiGetPatchesHistoryRequest{
        Service: &service
    }

    resp, r, err := api_client.PatchManagement.GetPatchesHistory(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchManagement.GetPatchesHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPatchesHistory`: []PatchOperation
    fmt.Fprintf(os.Stdout, "Response from `PatchManagement.GetPatchesHistory`: %v\n", resp)
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

> []PatchDetail ImportPatches(ctx).FileName(fileName).Checksum(checksum).Patch(patch).Execute()

Import patches



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    fileName := "fileName_example" // string | 
    checksum := "checksum_example" // string | 
    patch := os.NewFile(1234, "some_file") // *os.File | 

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

    request := onprem.ApiImportPatchesRequest{
        FileName: &fileName
        Checksum: &checksum
        Patch: &patch
    }

    resp, r, err := api_client.PatchManagement.ImportPatches(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchManagement.ImportPatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportPatches`: []PatchDetail
    fmt.Fprintf(os.Stdout, "Response from `PatchManagement.ImportPatches`: %v\n", resp)
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
    onprem "onprem"
)

func main() {
    body := *openapiclient.NewRevertPatchesRequest("Service_example") // RevertPatchesRequest | Request to revert patches.

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

    request := onprem.ApiRevertPatchesRequest{
        Body: &body
    }

    resp, r, err := api_client.PatchManagement.RevertPatches(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchManagement.RevertPatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevertPatches`: []ServicePatchLevel
    fmt.Fprintf(os.Stdout, "Response from `PatchManagement.RevertPatches`: %v\n", resp)
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

