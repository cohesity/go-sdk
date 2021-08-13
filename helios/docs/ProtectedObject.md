# \ProtectedObject

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PerformActionOnProtectObjects**](ProtectedObject.md#PerformActionOnProtectObjects) | **Post** /data-protect/protected-objects/actions | Perform Actions on Protect Objects.
[**ProtectObjectsOfAnyType**](ProtectedObject.md#ProtectObjectsOfAnyType) | **Post** /data-protect/protected-objects | Create Object Backup.
[**UpdateProtectedObjectsOfAnyType**](ProtectedObject.md#UpdateProtectedObjectsOfAnyType) | **Put** /data-protect/protected-objects/{id} | Update Object Backup.



## PerformActionOnProtectObjects

> ProtectedObjectActionResponse PerformActionOnProtectObjects(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Perform Actions on Protect Objects.



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
    body := *openapiclient.NewProtectdObjectsActionRequest("Action_example") // ProtectdObjectsActionRequest | Specifies the parameters to perform an action on an already protected object.
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

    request := helios.ApiPerformActionOnProtectObjectsRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.ProtectedObject.PerformActionOnProtectObjects(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectedObject.PerformActionOnProtectObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PerformActionOnProtectObjects`: ProtectedObjectActionResponse
    fmt.Fprintf(os.Stdout, "Response from `ProtectedObject.PerformActionOnProtectObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPerformActionOnProtectObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProtectdObjectsActionRequest**](ProtectdObjectsActionRequest.md) | Specifies the parameters to perform an action on an already protected object. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

> CreateProtectedObjectsResponse ProtectObjectsOfAnyType(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Create Object Backup.



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
    body := *openapiclient.NewCreateProtectedObjectsRequest([]openapiclient.EnvSpecificObjectProtectionRequestParams{*openapiclient.NewEnvSpecificObjectProtectionRequestParams()}) // CreateProtectedObjectsRequest | Specifies the parameters to protect objects.
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

    request := helios.ApiProtectObjectsOfAnyTypeRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.ProtectedObject.ProtectObjectsOfAnyType(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectedObject.ProtectObjectsOfAnyType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProtectObjectsOfAnyType`: CreateProtectedObjectsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProtectedObject.ProtectObjectsOfAnyType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProtectObjectsOfAnyTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateProtectedObjectsRequest**](CreateProtectedObjectsRequest.md) | Specifies the parameters to protect objects. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

> GetProtectedObjectResponse UpdateProtectedObjectsOfAnyType(ctx, id).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Update Object Backup.



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
    id := int64(789) // int64 | Specifies the id of the Protected Object.
    body := *openapiclient.NewUpdateProtectedObjectsRequest() // UpdateProtectedObjectsRequest | Specifies the parameters to perform an update on protected objects.
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

    request := helios.ApiUpdateProtectedObjectsOfAnyTypeRequest{
        Id: &id
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.ProtectedObject.UpdateProtectedObjectsOfAnyType(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectedObject.UpdateProtectedObjectsOfAnyType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProtectedObjectsOfAnyType`: GetProtectedObjectResponse
    fmt.Fprintf(os.Stdout, "Response from `ProtectedObject.UpdateProtectedObjectsOfAnyType`: %v\n", resp)
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
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

