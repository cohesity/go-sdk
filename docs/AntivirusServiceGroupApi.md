# \AntivirusServiceGroupApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AntivirusServiceGroupState**](AntivirusServiceGroupApi.md#AntivirusServiceGroupState) | **Put** /public/antivirusGroups/states | Change the state of an antivirus service group.
[**CreateAntivirusServiceGroup**](AntivirusServiceGroupApi.md#CreateAntivirusServiceGroup) | **Post** /public/antivirusGroups | Create an Antivirus service group.
[**DeleteAntivirusServiceGroup**](AntivirusServiceGroupApi.md#DeleteAntivirusServiceGroup) | **Delete** /public/antivirusGroups/{id} | Delete an antivirus service group corresponding to the specified antivirus service group Id.
[**DeleteInfectedFiles**](AntivirusServiceGroupApi.md#DeleteInfectedFiles) | **Delete** /public/infectedFiles | Delete the list of infected files.
[**GetAntivirusServiceGroup**](AntivirusServiceGroupApi.md#GetAntivirusServiceGroup) | **Get** /public/antivirusGroups | Lists the antivirus service groups.
[**GetIcapConnectionStatus**](AntivirusServiceGroupApi.md#GetIcapConnectionStatus) | **Get** /public/icapConnectionStatus | Check the Icap server connection status.
[**GetInfectedFiles**](AntivirusServiceGroupApi.md#GetInfectedFiles) | **Get** /public/infectedFiles | Lists the infected files.
[**UpdateAntivirusServiceGroup**](AntivirusServiceGroupApi.md#UpdateAntivirusServiceGroup) | **Put** /public/antivirusGroups | Update an antivirus service group.
[**UpdateInfectedFiles**](AntivirusServiceGroupApi.md#UpdateInfectedFiles) | **Put** /public/infectedFiles | Update the list of infected files.



## AntivirusServiceGroupState

> AntivirusServiceGroupStateParams AntivirusServiceGroupState(ctx).Body(body).Execute()

Change the state of an antivirus service group.



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
    body := *openapiclient.NewAntivirusServiceGroupStateParams(false, NullableInt64(123)) // AntivirusServiceGroupStateParams |  (optional)

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

    request := cohesitysdk.ApiAntivirusServiceGroupStateRequest{
        body: &Body
    }

    resp, r, err := api_client.AntivirusServiceGroupApi.AntivirusServiceGroupState(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AntivirusServiceGroupApi.AntivirusServiceGroupState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AntivirusServiceGroupState`: AntivirusServiceGroupStateParams
    fmt.Fprintf(os.Stdout, "Response from `AntivirusServiceGroupApi.AntivirusServiceGroupState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAntivirusServiceGroupStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AntivirusServiceGroupStateParams**](AntivirusServiceGroupStateParams.md) |  | 

### Return type

[**AntivirusServiceGroupStateParams**](AntivirusServiceGroupStateParams.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAntivirusServiceGroup

> AntivirusServiceGroup CreateAntivirusServiceGroup(ctx).Body(body).Execute()

Create an Antivirus service group.



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
    body := *openapiclient.NewAntivirusServiceGroupParams("Name_example") // AntivirusServiceGroupParams | Request to create an Antivirus Service Group.

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

    request := cohesitysdk.ApiCreateAntivirusServiceGroupRequest{
        body: &Body
    }

    resp, r, err := api_client.AntivirusServiceGroupApi.CreateAntivirusServiceGroup(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AntivirusServiceGroupApi.CreateAntivirusServiceGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAntivirusServiceGroup`: AntivirusServiceGroup
    fmt.Fprintf(os.Stdout, "Response from `AntivirusServiceGroupApi.CreateAntivirusServiceGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAntivirusServiceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AntivirusServiceGroupParams**](AntivirusServiceGroupParams.md) | Request to create an Antivirus Service Group. | 

### Return type

[**AntivirusServiceGroup**](AntivirusServiceGroup.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAntivirusServiceGroup

> DeleteAntivirusServiceGroup(ctx, id).Execute()

Delete an antivirus service group corresponding to the specified antivirus service group Id.



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
    id := int64(789) // int64 | Specifies the AntivirusServiceGroup Id.

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

    request := cohesitysdk.ApiDeleteAntivirusServiceGroupRequest{
        id: &Id
    }

    resp, r, err := api_client.AntivirusServiceGroupApi.DeleteAntivirusServiceGroup(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AntivirusServiceGroupApi.DeleteAntivirusServiceGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the AntivirusServiceGroup Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAntivirusServiceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInfectedFiles

> DeleteInfectedFileResponse DeleteInfectedFiles(ctx).Body(body).Execute()

Delete the list of infected files.



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
    body := *openapiclient.NewDeleteInfectedFileParams() // DeleteInfectedFileParams | Request to delete the list of infected files.

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

    request := cohesitysdk.ApiDeleteInfectedFilesRequest{
        body: &Body
    }

    resp, r, err := api_client.AntivirusServiceGroupApi.DeleteInfectedFiles(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AntivirusServiceGroupApi.DeleteInfectedFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteInfectedFiles`: DeleteInfectedFileResponse
    fmt.Fprintf(os.Stdout, "Response from `AntivirusServiceGroupApi.DeleteInfectedFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInfectedFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DeleteInfectedFileParams**](DeleteInfectedFileParams.md) | Request to delete the list of infected files. | 

### Return type

[**DeleteInfectedFileResponse**](DeleteInfectedFileResponse.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAntivirusServiceGroup

> []AntivirusServiceGroup GetAntivirusServiceGroup(ctx).Execute()

Lists the antivirus service groups.



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

    request := cohesitysdk.ApiGetAntivirusServiceGroupRequest{
    }

    resp, r, err := api_client.AntivirusServiceGroupApi.GetAntivirusServiceGroup(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AntivirusServiceGroupApi.GetAntivirusServiceGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAntivirusServiceGroup`: []AntivirusServiceGroup
    fmt.Fprintf(os.Stdout, "Response from `AntivirusServiceGroupApi.GetAntivirusServiceGroup`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAntivirusServiceGroupRequest struct via the builder pattern


### Return type

[**[]AntivirusServiceGroup**](AntivirusServiceGroup.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIcapConnectionStatus

> IcapConnectionStatusResponse GetIcapConnectionStatus(ctx).IcapUris(icapUris).Execute()

Check the Icap server connection status.



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
    icapUris := []string{"Inner_example"} // []string | Specifies the list of icap uri. (optional)

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

    request := cohesitysdk.ApiGetIcapConnectionStatusRequest{
        icapUris: &IcapUris
    }

    resp, r, err := api_client.AntivirusServiceGroupApi.GetIcapConnectionStatus(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AntivirusServiceGroupApi.GetIcapConnectionStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIcapConnectionStatus`: IcapConnectionStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `AntivirusServiceGroupApi.GetIcapConnectionStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIcapConnectionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **icapUris** | **[]string** | Specifies the list of icap uri. | 

### Return type

[**IcapConnectionStatusResponse**](IcapConnectionStatusResponse.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfectedFiles

> InfectedFiles GetInfectedFiles(ctx).ViewNames(viewNames).IncludeQuarantinedFiles(includeQuarantinedFiles).IncludeUnquarantinedFiles(includeUnquarantinedFiles).FilePath(filePath).PageCount(pageCount).PaginationCookie(paginationCookie).Execute()

Lists the infected files.



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
    viewNames := []string{"Inner_example"} // []string | Filter by a list of View names. (optional)
    includeQuarantinedFiles := true // bool | Specifies whether to include quarantined files in the result. (optional)
    includeUnquarantinedFiles := true // bool | Specifies whether to include unquarantined files in the result. (optional)
    filePath := "filePath_example" // string | Specifies the path of a file. If this is provided, infected file list would contain the scan and infection state of the file and pagination cookie will be ignored. (optional)
    pageCount := int64(789) // int64 | Specifies the number of items to return in the response for pagination purposes. Default value is 1000. (optional)
    paginationCookie := "paginationCookie_example" // string | Pagination cookie should be used from previous call to list infected files. It resumes (or gives the next set of values) from the result of the previous call. (optional)

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

    request := cohesitysdk.ApiGetInfectedFilesRequest{
        viewNames: &ViewNames
        includeQuarantinedFiles: &IncludeQuarantinedFiles
        includeUnquarantinedFiles: &IncludeUnquarantinedFiles
        filePath: &FilePath
        pageCount: &PageCount
        paginationCookie: &PaginationCookie
    }

    resp, r, err := api_client.AntivirusServiceGroupApi.GetInfectedFiles(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AntivirusServiceGroupApi.GetInfectedFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInfectedFiles`: InfectedFiles
    fmt.Fprintf(os.Stdout, "Response from `AntivirusServiceGroupApi.GetInfectedFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInfectedFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **viewNames** | **[]string** | Filter by a list of View names. | 
 **includeQuarantinedFiles** | **bool** | Specifies whether to include quarantined files in the result. | 
 **includeUnquarantinedFiles** | **bool** | Specifies whether to include unquarantined files in the result. | 
 **filePath** | **string** | Specifies the path of a file. If this is provided, infected file list would contain the scan and infection state of the file and pagination cookie will be ignored. | 
 **pageCount** | **int64** | Specifies the number of items to return in the response for pagination purposes. Default value is 1000. | 
 **paginationCookie** | **string** | Pagination cookie should be used from previous call to list infected files. It resumes (or gives the next set of values) from the result of the previous call. | 

### Return type

[**InfectedFiles**](InfectedFiles.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAntivirusServiceGroup

> AntivirusServiceGroup UpdateAntivirusServiceGroup(ctx).Body(body).Execute()

Update an antivirus service group.



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
    body := *openapiclient.NewUpdateAntivirusServiceGroupParams(NullableInt64(123), "Name_example") // UpdateAntivirusServiceGroupParams | Request to update an Antivirus Service Group.

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

    request := cohesitysdk.ApiUpdateAntivirusServiceGroupRequest{
        body: &Body
    }

    resp, r, err := api_client.AntivirusServiceGroupApi.UpdateAntivirusServiceGroup(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AntivirusServiceGroupApi.UpdateAntivirusServiceGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAntivirusServiceGroup`: AntivirusServiceGroup
    fmt.Fprintf(os.Stdout, "Response from `AntivirusServiceGroupApi.UpdateAntivirusServiceGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAntivirusServiceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateAntivirusServiceGroupParams**](UpdateAntivirusServiceGroupParams.md) | Request to update an Antivirus Service Group. | 

### Return type

[**AntivirusServiceGroup**](AntivirusServiceGroup.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInfectedFiles

> UpdateInfectedFileResponse UpdateInfectedFiles(ctx).Body(body).Execute()

Update the list of infected files.



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
    body := *openapiclient.NewUpdateInfectedFileParams() // UpdateInfectedFileParams | Request to update the list of infected files.

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

    request := cohesitysdk.ApiUpdateInfectedFilesRequest{
        body: &Body
    }

    resp, r, err := api_client.AntivirusServiceGroupApi.UpdateInfectedFiles(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AntivirusServiceGroupApi.UpdateInfectedFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInfectedFiles`: UpdateInfectedFileResponse
    fmt.Fprintf(os.Stdout, "Response from `AntivirusServiceGroupApi.UpdateInfectedFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInfectedFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateInfectedFileParams**](UpdateInfectedFileParams.md) | Request to update the list of infected files. | 

### Return type

[**UpdateInfectedFileResponse**](UpdateInfectedFileResponse.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

