# \RemoteStorage

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRemoteStorageRegistration**](RemoteStorage.md#DeleteRemoteStorageRegistration) | **Delete** /remote-storage/{id} | Delete Remote Storage Registration
[**GetRegisteredRemoteStorageList**](RemoteStorage.md#GetRegisteredRemoteStorageList) | **Get** /remote-storage | Get Registered Remote Storage Servers List
[**GetRemoteStorageDetails**](RemoteStorage.md#GetRemoteStorageDetails) | **Get** /remote-storage/{id} | Get remote storage details
[**RegisterNewRemoteStorage**](RemoteStorage.md#RegisterNewRemoteStorage) | **Post** /remote-storage | Register Remote Storage
[**UpdateRemoteStorageRegistration**](RemoteStorage.md#UpdateRemoteStorageRegistration) | **Patch** /remote-storage/{id} | Update Remote Storage Config



## DeleteRemoteStorageRegistration

> DeleteRemoteStorageRegistration(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Delete Remote Storage Registration



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
    id := int64(789) // int64 | Specifies the registration id of the registered remote storage.
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

    request := helios.ApiDeleteRemoteStorageRegistrationRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.RemoteStorage.DeleteRemoteStorageRegistration(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteStorage.DeleteRemoteStorageRegistration``: %v\n", err)
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

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

> RegisteredRemoteStorageList GetRegisteredRemoteStorageList(ctx).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Get Registered Remote Storage Servers List



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

    request := helios.ApiGetRegisteredRemoteStorageListRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.RemoteStorage.GetRegisteredRemoteStorageList(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteStorage.GetRegisteredRemoteStorageList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegisteredRemoteStorageList`: RegisteredRemoteStorageList
    fmt.Fprintf(os.Stdout, "Response from `RemoteStorage.GetRegisteredRemoteStorageList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRegisteredRemoteStorageListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

> RemoteStorageInfo GetRemoteStorageDetails(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).IncludeAvailableSpace(includeAvailableSpace).IncludeAvailableDataVips(includeAvailableDataVips).IncludeArrayInfo(includeArrayInfo).Execute()

Get remote storage details



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
    id := int64(789) // int64 | Specifies the id of the registered remote storage.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    includeAvailableSpace := true // bool | Specifies whether to include available capacity on remote storage. (optional) (default to false)
    includeAvailableDataVips := true // bool | Specifies whether to include available data vips on remote storage. (optional) (default to false)
    includeArrayInfo := true // bool | Includes flashblade specific info like name, software os and version of pure flashblade. (optional) (default to false)

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

    request := helios.ApiGetRemoteStorageDetailsRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        IncludeAvailableSpace: &includeAvailableSpace
        IncludeAvailableDataVips: &includeAvailableDataVips
        IncludeArrayInfo: &includeArrayInfo
    }

    resp, r, err := api_client.RemoteStorage.GetRemoteStorageDetails(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteStorage.GetRemoteStorageDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRemoteStorageDetails`: RemoteStorageInfo
    fmt.Fprintf(os.Stdout, "Response from `RemoteStorage.GetRemoteStorageDetails`: %v\n", resp)
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

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
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

> RemoteStorageInfo RegisterNewRemoteStorage(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Register Remote Storage



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
    body := *openapiclient.NewRemoteStorageInfo("Product_example") // RemoteStorageInfo | Specifies the parameters to register a remote storage management server.
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

    request := helios.ApiRegisterNewRemoteStorageRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.RemoteStorage.RegisterNewRemoteStorage(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteStorage.RegisterNewRemoteStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterNewRemoteStorage`: RemoteStorageInfo
    fmt.Fprintf(os.Stdout, "Response from `RemoteStorage.RegisterNewRemoteStorage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterNewRemoteStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RemoteStorageInfo**](RemoteStorageInfo.md) | Specifies the parameters to register a remote storage management server. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

> RemoteStorageInfo UpdateRemoteStorageRegistration(ctx, id).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Update Remote Storage Config



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
    id := int64(789) // int64 | Specifies the registration id of the registered remote storage.
    body := *openapiclient.NewRemoteStorageInfo("Product_example") // RemoteStorageInfo | Specifies the parameters to update the registration.
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

    request := helios.ApiUpdateRemoteStorageRegistrationRequest{
        Id: &id
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.RemoteStorage.UpdateRemoteStorageRegistration(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteStorage.UpdateRemoteStorageRegistration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRemoteStorageRegistration`: RemoteStorageInfo
    fmt.Fprintf(os.Stdout, "Response from `RemoteStorage.UpdateRemoteStorageRegistration`: %v\n", resp)
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
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

