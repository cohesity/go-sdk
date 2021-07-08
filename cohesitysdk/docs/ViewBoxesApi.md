# \ViewBoxesApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateViewBox**](ViewBoxesApi.md#CreateViewBox) | **Post** /public/viewBoxes | Create a Domain (View Box).
[**DeleteViewBox**](ViewBoxesApi.md#DeleteViewBox) | **Delete** /public/viewBoxes/{id} | Delete a Domain (View Box).
[**GetViewBoxById**](ViewBoxesApi.md#GetViewBoxById) | **Get** /public/viewBoxes/{id} | List details about a single Domain (View Box).
[**GetViewBoxes**](ViewBoxesApi.md#GetViewBoxes) | **Get** /public/viewBoxes | List Domains (View Boxes) filtered by the specified parameters.
[**UpdateViewBox**](ViewBoxesApi.md#UpdateViewBox) | **Put** /public/viewBoxes/{id} | Update a Domain (View Box).



## CreateViewBox

> ViewBox CreateViewBox(ctx).Body(body).Execute()

Create a Domain (View Box).



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
    body := *openapiclient.NewCreateViewBoxParams(NullableInt64(123), "Name_example") // CreateViewBoxParams | Request to create a Storage Domain (View Box) configuration.

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

    request := cohesitysdk.ApiCreateViewBoxRequest{
        Body: &body
    }

    resp, r, err := api_client.ViewBoxesApi.CreateViewBox(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewBoxesApi.CreateViewBox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateViewBox`: ViewBox
    fmt.Fprintf(os.Stdout, "Response from `ViewBoxesApi.CreateViewBox`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateViewBoxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateViewBoxParams**](CreateViewBoxParams.md) | Request to create a Storage Domain (View Box) configuration. | 

### Return type

[**ViewBox**](ViewBox.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteViewBox

> DeleteViewBox(ctx, id).Execute()

Delete a Domain (View Box).



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
    id := int64(789) // int64 | Id of the Storage Domain (View Box)

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

    request := cohesitysdk.ApiDeleteViewBoxRequest{
        Id: &id
    }

    resp, r, err := api_client.ViewBoxesApi.DeleteViewBox(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewBoxesApi.DeleteViewBox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Id of the Storage Domain (View Box) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteViewBoxRequest struct via the builder pattern


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


## GetViewBoxById

> ViewBox GetViewBoxById(ctx, id).FetchStats(fetchStats).Execute()

List details about a single Domain (View Box).



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
    id := int64(789) // int64 | Id of the Storage Domain (View Box)
    fetchStats := true // bool | Specifies whether to include usage and performance statistics. (optional)

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

    request := cohesitysdk.ApiGetViewBoxByIdRequest{
        Id: &id
        FetchStats: &fetchStats
    }

    resp, r, err := api_client.ViewBoxesApi.GetViewBoxById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewBoxesApi.GetViewBoxById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetViewBoxById`: ViewBox
    fmt.Fprintf(os.Stdout, "Response from `ViewBoxesApi.GetViewBoxById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Id of the Storage Domain (View Box) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetViewBoxByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fetchStats** | **bool** | Specifies whether to include usage and performance statistics. | 

### Return type

[**ViewBox**](ViewBox.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViewBoxes

> []ViewBox GetViewBoxes(ctx).TenantIds(tenantIds).AllUnderHierarchy(allUnderHierarchy).Ids(ids).Names(names).ClusterPartitionIds(clusterPartitionIds).FetchStats(fetchStats).FetchTimeSeriesSchema(fetchTimeSeriesSchema).TemplateId(templateId).Execute()

List Domains (View Boxes) filtered by the specified parameters.



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
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
    allUnderHierarchy := true // bool | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user's organization should be returned. (optional)
    ids := []int64{int64(123)} // []int64 | Filter by a list of Storage Domain (View Box) ids. If empty, View Boxes are not filtered by id. (optional)
    names := []string{"Inner_example"} // []string | Filter by a list of Storage Domain (View Box) Names. If empty, Storage Domains (View Boxes) are not filtered by Name. (optional)
    clusterPartitionIds := []int64{int64(123)} // []int64 | Filter by a list of Cluster Partition Ids. (optional)
    fetchStats := true // bool | Specifies whether to include usage and performance statistics. (optional)
    fetchTimeSeriesSchema := true // bool | Specifies whether to get time series schema info of the view box. (optional)
    templateId := int64(789) // int64 | Filter list of Storage Domain (View Box) by the properties of the template like dedup, compression. If empty, Storage Domains (View Boxes) are not filtered. (optional)

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

    request := cohesitysdk.ApiGetViewBoxesRequest{
        TenantIds: &tenantIds
        AllUnderHierarchy: &allUnderHierarchy
        Ids: &ids
        Names: &names
        ClusterPartitionIds: &clusterPartitionIds
        FetchStats: &fetchStats
        FetchTimeSeriesSchema: &fetchTimeSeriesSchema
        TemplateId: &templateId
    }

    resp, r, err := api_client.ViewBoxesApi.GetViewBoxes(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewBoxesApi.GetViewBoxes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetViewBoxes`: []ViewBox
    fmt.Fprintf(os.Stdout, "Response from `ViewBoxesApi.GetViewBoxes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetViewBoxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **allUnderHierarchy** | **bool** | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. | 
 **ids** | **[]int64** | Filter by a list of Storage Domain (View Box) ids. If empty, View Boxes are not filtered by id. | 
 **names** | **[]string** | Filter by a list of Storage Domain (View Box) Names. If empty, Storage Domains (View Boxes) are not filtered by Name. | 
 **clusterPartitionIds** | **[]int64** | Filter by a list of Cluster Partition Ids. | 
 **fetchStats** | **bool** | Specifies whether to include usage and performance statistics. | 
 **fetchTimeSeriesSchema** | **bool** | Specifies whether to get time series schema info of the view box. | 
 **templateId** | **int64** | Filter list of Storage Domain (View Box) by the properties of the template like dedup, compression. If empty, Storage Domains (View Boxes) are not filtered. | 

### Return type

[**[]ViewBox**](ViewBox.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateViewBox

> ViewBox UpdateViewBox(ctx, id).Body(body).Execute()

Update a Domain (View Box).



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
    id := int64(789) // int64 | Id of the Storage Domain (View Box)
    body := *openapiclient.NewCreateViewBoxParams(NullableInt64(123), "Name_example") // CreateViewBoxParams | Request to update a Storage Domain (View Box) configuration.

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

    request := cohesitysdk.ApiUpdateViewBoxRequest{
        Id: &id
        Body: &body
    }

    resp, r, err := api_client.ViewBoxesApi.UpdateViewBox(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewBoxesApi.UpdateViewBox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateViewBox`: ViewBox
    fmt.Fprintf(os.Stdout, "Response from `ViewBoxesApi.UpdateViewBox`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Id of the Storage Domain (View Box) | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateViewBoxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateViewBoxParams**](CreateViewBoxParams.md) | Request to update a Storage Domain (View Box) configuration. | 

### Return type

[**ViewBox**](ViewBox.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

