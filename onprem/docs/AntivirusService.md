# \AntivirusService

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAntivirusServiceGroups**](AntivirusService.md#GetAntivirusServiceGroups) | **Get** /antivirus-service/groups | Get Antivirus Service groups.
[**GetIcapUriConnectionStatus**](AntivirusService.md#GetIcapUriConnectionStatus) | **Get** /antivirus-service/icap-uri-connection-status | Get ICAP Uri connection status.
[**GetInfectedFiles**](AntivirusService.md#GetInfectedFiles) | **Get** /antivirus-service/infected-files | Get infected files.



## GetAntivirusServiceGroups

> AntivirusServiceGroups GetAntivirusServiceGroups(ctx).Execute()

Get Antivirus Service groups.



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

    request := onprem.ApiGetAntivirusServiceGroupsRequest{
    }

    resp, r, err := api_client.AntivirusService.GetAntivirusServiceGroups(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AntivirusService.GetAntivirusServiceGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAntivirusServiceGroups`: AntivirusServiceGroups
    fmt.Fprintf(os.Stdout, "Response from `AntivirusService.GetAntivirusServiceGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAntivirusServiceGroupsRequest struct via the builder pattern


### Return type

[**AntivirusServiceGroups**](AntivirusServiceGroups.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIcapUriConnectionStatus

> IcapUriConnectionStatusList GetIcapUriConnectionStatus(ctx).Uris(uris).Execute()

Get ICAP Uri connection status.



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
    uris := []string{"Inner_example"} // []string | Specifies a list of URIs to check connection status. (optional)

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

    request := onprem.ApiGetIcapUriConnectionStatusRequest{
        Uris: &uris
    }

    resp, r, err := api_client.AntivirusService.GetIcapUriConnectionStatus(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AntivirusService.GetIcapUriConnectionStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIcapUriConnectionStatus`: IcapUriConnectionStatusList
    fmt.Fprintf(os.Stdout, "Response from `AntivirusService.GetIcapUriConnectionStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIcapUriConnectionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uris** | **[]string** | Specifies a list of URIs to check connection status. | 

### Return type

[**IcapUriConnectionStatusList**](IcapUriConnectionStatusList.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfectedFiles

> InfectedFiles GetInfectedFiles(ctx).ViewIds(viewIds).Path(path).States(states).MaxCount(maxCount).Cookie(cookie).Execute()

Get infected files.



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
    viewIds := []int64{int64(123)} // []int64 | Specifies a list of view ids. Only infected files from these views will be returned. (optional)
    path := "path_example" // string | Specifies the file path. (optional)
    states := []string{"States_example"} // []string | Specifies the file states. (optional)
    maxCount := int64(789) // int64 | Specifies the max number of files to be returned. (optional)
    cookie := "cookie_example" // string | Specifies the pagination cookie. (optional)

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

    request := onprem.ApiGetInfectedFilesRequest{
        ViewIds: &viewIds
        Path: &path
        States: &states
        MaxCount: &maxCount
        Cookie: &cookie
    }

    resp, r, err := api_client.AntivirusService.GetInfectedFiles(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AntivirusService.GetInfectedFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInfectedFiles`: InfectedFiles
    fmt.Fprintf(os.Stdout, "Response from `AntivirusService.GetInfectedFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInfectedFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **viewIds** | **[]int64** | Specifies a list of view ids. Only infected files from these views will be returned. | 
 **path** | **string** | Specifies the file path. | 
 **states** | **[]string** | Specifies the file states. | 
 **maxCount** | **int64** | Specifies the max number of files to be returned. | 
 **cookie** | **string** | Specifies the pagination cookie. | 

### Return type

[**InfectedFiles**](InfectedFiles.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

