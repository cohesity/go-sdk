# \Patch

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyPatches**](Patch.md#ApplyPatches) | **Post** /patch-management/available-patches | Apply patches
[**GetAppliedPatches**](Patch.md#GetAppliedPatches) | **Get** /patch-management/applied-patches | Get applied patches
[**GetAvailablePatches**](Patch.md#GetAvailablePatches) | **Get** /patch-management/available-patches | Get available patches
[**GetPatchOperationStatus**](Patch.md#GetPatchOperationStatus) | **Get** /patch-management/operation-status | Get patch operation status
[**GetPatchesHistory**](Patch.md#GetPatchesHistory) | **Get** /patch-management/patches-history | Get patches history
[**ImportPatches**](Patch.md#ImportPatches) | **Put** /patch-management/available-patches | Import patches
[**RevertPatches**](Patch.md#RevertPatches) | **Post** /patch-management/applied-patches | Revert patches



## ApplyPatches

> []ServicePatchLevel ApplyPatches(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Apply patches



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
    body := *openapiclient.NewApplyPatchesRequest() // ApplyPatchesRequest | Request to apply patches.
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

    request := helios.ApiApplyPatchesRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Patch.ApplyPatches(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Patch.ApplyPatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplyPatches`: []ServicePatchLevel
    fmt.Fprintf(os.Stdout, "Response from `Patch.ApplyPatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplyPatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApplyPatchesRequest**](ApplyPatchesRequest.md) | Request to apply patches. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

> []AppliedPatch GetAppliedPatches(ctx).AccessClusterId(accessClusterId).RegionId(regionId).Service(service).IncludeDetails(includeDetails).Execute()

Get applied patches



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

    request := helios.ApiGetAppliedPatchesRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        Service: &service
        IncludeDetails: &includeDetails
    }

    resp, r, err := api_client.Patch.GetAppliedPatches(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Patch.GetAppliedPatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppliedPatches`: []AppliedPatch
    fmt.Fprintf(os.Stdout, "Response from `Patch.GetAppliedPatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAppliedPatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
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

> []AvailablePatch GetAvailablePatches(ctx).AccessClusterId(accessClusterId).RegionId(regionId).Service(service).IncludeDetails(includeDetails).Execute()

Get available patches



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

    request := helios.ApiGetAvailablePatchesRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        Service: &service
        IncludeDetails: &includeDetails
    }

    resp, r, err := api_client.Patch.GetAvailablePatches(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Patch.GetAvailablePatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvailablePatches`: []AvailablePatch
    fmt.Fprintf(os.Stdout, "Response from `Patch.GetAvailablePatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailablePatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
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

> PatchOperationStatus GetPatchOperationStatus(ctx).AccessClusterId(accessClusterId).RegionId(regionId).IncludeDetails(includeDetails).Execute()

Get patch operation status



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

    request := helios.ApiGetPatchOperationStatusRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        IncludeDetails: &includeDetails
    }

    resp, r, err := api_client.Patch.GetPatchOperationStatus(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Patch.GetPatchOperationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPatchOperationStatus`: PatchOperationStatus
    fmt.Fprintf(os.Stdout, "Response from `Patch.GetPatchOperationStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPatchOperationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
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

> []PatchOperation GetPatchesHistory(ctx).AccessClusterId(accessClusterId).RegionId(regionId).Service(service).Execute()

Get patches history



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

    request := helios.ApiGetPatchesHistoryRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        Service: &service
    }

    resp, r, err := api_client.Patch.GetPatchesHistory(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Patch.GetPatchesHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPatchesHistory`: []PatchOperation
    fmt.Fprintf(os.Stdout, "Response from `Patch.GetPatchesHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPatchesHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
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

> []PatchDetail ImportPatches(ctx).FileName(fileName).Checksum(checksum).Patch(patch).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Import patches



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
    fileName := "fileName_example" // string | 
    checksum := "checksum_example" // string | 
    patch := os.NewFile(1234, "some_file") // *os.File | 
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

    request := helios.ApiImportPatchesRequest{
        FileName: &fileName
        Checksum: &checksum
        Patch: &patch
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Patch.ImportPatches(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Patch.ImportPatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportPatches`: []PatchDetail
    fmt.Fprintf(os.Stdout, "Response from `Patch.ImportPatches`: %v\n", resp)
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
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

> []ServicePatchLevel RevertPatches(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Revert patches



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
    body := *openapiclient.NewRevertPatchesRequest("Service_example") // RevertPatchesRequest | Request to revert patches.
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

    request := helios.ApiRevertPatchesRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Patch.RevertPatches(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Patch.RevertPatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevertPatches`: []ServicePatchLevel
    fmt.Fprintf(os.Stdout, "Response from `Patch.RevertPatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevertPatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RevertPatchesRequest**](RevertPatchesRequest.md) | Request to revert patches. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

