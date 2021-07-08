# \ProtectionObjectsApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProtectionObjectSummary**](ProtectionObjectsApi.md#GetProtectionObjectSummary) | **Get** /public/protectionObjects/summary | Protect an Object.
[**ProtectObject**](ProtectionObjectsApi.md#ProtectObject) | **Post** /public/protectionObjects | Protect an Object.
[**UnprotectObject**](ProtectionObjectsApi.md#UnprotectObject) | **Delete** /public/protectionObjects | Unprotect a Protected Object.
[**UpdateProtectionObject**](ProtectionObjectsApi.md#UpdateProtectionObject) | **Put** /public/protectionObjects | Update a Protected Object.



## GetProtectionObjectSummary

> ProtectionObjectSummary GetProtectionObjectSummary(ctx).ProtectionSourceId(protectionSourceId).Execute()

Protect an Object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    protectionSourceId := int64(789) // int64 | Specifies the id of the Protection Source.

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

    request := cohesitysdk.ApiGetProtectionObjectSummaryRequest{
        ProtectionSourceId: &protectionSourceId
    }

    resp, r, err := api_client.ProtectionObjectsApi.GetProtectionObjectSummary(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionObjectsApi.GetProtectionObjectSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtectionObjectSummary`: ProtectionObjectSummary
    fmt.Fprintf(os.Stdout, "Response from `ProtectionObjectsApi.GetProtectionObjectSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionObjectSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **protectionSourceId** | **int64** | Specifies the id of the Protection Source. | 

### Return type

[**ProtectionObjectSummary**](ProtectionObjectSummary.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProtectObject

> []ProtectedObject ProtectObject(ctx).Body(body).Execute()

Protect an Object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewProtectObjectParameters([]int64{int64(123)}, "RpoPolicyId_example") // ProtectObjectParameters |  (optional)

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

    request := cohesitysdk.ApiProtectObjectRequest{
        Body: &body
    }

    resp, r, err := api_client.ProtectionObjectsApi.ProtectObject(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionObjectsApi.ProtectObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProtectObject`: []ProtectedObject
    fmt.Fprintf(os.Stdout, "Response from `ProtectionObjectsApi.ProtectObject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProtectObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProtectObjectParameters**](ProtectObjectParameters.md) |  | 

### Return type

[**[]ProtectedObject**](ProtectedObject.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnprotectObject

> UnprotectObject(ctx).Body(body).Execute()

Unprotect a Protected Object.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewUnprotectObjectParams(NullableInt64(123), "RpoPolicyId_example") // UnprotectObjectParams |  (optional)

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

    request := cohesitysdk.ApiUnprotectObjectRequest{
        Body: &body
    }

    resp, r, err := api_client.ProtectionObjectsApi.UnprotectObject(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionObjectsApi.UnprotectObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnprotectObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UnprotectObjectParams**](UnprotectObjectParams.md) |  | 

### Return type

 (empty response body)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProtectionObject

> ProtectionJob UpdateProtectionObject(ctx).Body(body).Execute()

Update a Protected Object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewUpdateProtectionObjectParameters(*openapiclient.NewUniversalId()) // UpdateProtectionObjectParameters |  (optional)

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

    request := cohesitysdk.ApiUpdateProtectionObjectRequest{
        Body: &body
    }

    resp, r, err := api_client.ProtectionObjectsApi.UpdateProtectionObject(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionObjectsApi.UpdateProtectionObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProtectionObject`: ProtectionJob
    fmt.Fprintf(os.Stdout, "Response from `ProtectionObjectsApi.UpdateProtectionObject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProtectionObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateProtectionObjectParameters**](UpdateProtectionObjectParameters.md) |  | 

### Return type

[**ProtectionJob**](ProtectionJob.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

