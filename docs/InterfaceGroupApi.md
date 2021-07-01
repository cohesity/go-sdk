# \InterfaceGroupApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInterfaceGroup**](InterfaceGroupApi.md#CreateInterfaceGroup) | **Post** /public/interfaceGroups | Create an interface group on the Cohesity Cluster.
[**DeleteInterfaceGroup**](InterfaceGroupApi.md#DeleteInterfaceGroup) | **Delete** /public/interfaceGroups/{name} | Delete the specified interface group from the Cohesity Cluster.
[**GetInterfaceGroups**](InterfaceGroupApi.md#GetInterfaceGroups) | **Get** /public/interfaceGroups | List the interface groups for the Cohesity Cluster.
[**UpdateInterfaceGroup**](InterfaceGroupApi.md#UpdateInterfaceGroup) | **Put** /public/interfaceGroups | Update an interface group on the Cohesity Cluster.



## CreateInterfaceGroup

> InterfaceGroup CreateInterfaceGroup(ctx).Body(body).Execute()

Create an interface group on the Cohesity Cluster.



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
    body := *openapiclient.NewInterfaceGroup() // InterfaceGroup |  (optional)

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

    request := cohesitysdk.ApiCreateInterfaceGroupRequest{
        body: &Body
    }

    resp, r, err := api_client.InterfaceGroupApi.CreateInterfaceGroup(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterfaceGroupApi.CreateInterfaceGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInterfaceGroup`: InterfaceGroup
    fmt.Fprintf(os.Stdout, "Response from `InterfaceGroupApi.CreateInterfaceGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInterfaceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InterfaceGroup**](InterfaceGroup.md) |  | 

### Return type

[**InterfaceGroup**](InterfaceGroup.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInterfaceGroup

> DeleteInterfaceGroup(ctx, name).Execute()

Delete the specified interface group from the Cohesity Cluster.



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
    name := "name_example" // string | Request to delete one interface group.

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

    request := cohesitysdk.ApiDeleteInterfaceGroupRequest{
        name: &Name
    }

    resp, r, err := api_client.InterfaceGroupApi.DeleteInterfaceGroup(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterfaceGroupApi.DeleteInterfaceGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Request to delete one interface group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInterfaceGroupRequest struct via the builder pattern


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


## GetInterfaceGroups

> []InterfaceGroup GetInterfaceGroups(ctx).Execute()

List the interface groups for the Cohesity Cluster.



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

    request := cohesitysdk.ApiGetInterfaceGroupsRequest{
    }

    resp, r, err := api_client.InterfaceGroupApi.GetInterfaceGroups(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterfaceGroupApi.GetInterfaceGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInterfaceGroups`: []InterfaceGroup
    fmt.Fprintf(os.Stdout, "Response from `InterfaceGroupApi.GetInterfaceGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInterfaceGroupsRequest struct via the builder pattern


### Return type

[**[]InterfaceGroup**](InterfaceGroup.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInterfaceGroup

> InterfaceGroup UpdateInterfaceGroup(ctx).Body(body).Execute()

Update an interface group on the Cohesity Cluster.



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
    body := *openapiclient.NewInterfaceGroup() // InterfaceGroup |  (optional)

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

    request := cohesitysdk.ApiUpdateInterfaceGroupRequest{
        body: &Body
    }

    resp, r, err := api_client.InterfaceGroupApi.UpdateInterfaceGroup(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterfaceGroupApi.UpdateInterfaceGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInterfaceGroup`: InterfaceGroup
    fmt.Fprintf(os.Stdout, "Response from `InterfaceGroupApi.UpdateInterfaceGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInterfaceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InterfaceGroup**](InterfaceGroup.md) |  | 

### Return type

[**InterfaceGroup**](InterfaceGroup.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

