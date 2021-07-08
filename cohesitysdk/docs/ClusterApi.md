# \ClusterApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAppSettings**](ClusterApi.md#GetAppSettings) | **Get** /public/cluster/appSettings | Gets the app settings for the cluster.
[**GetBasicClusterInfo**](ClusterApi.md#GetBasicClusterInfo) | **Get** /public/basicClusterInfo | List details about the Cohesity Cluster such as the name, type, version, language, locale and domains. This operation does not require authentication.
[**GetCluster**](ClusterApi.md#GetCluster) | **Get** /public/cluster | List details about this Cohesity Cluster.
[**GetClusterStatus**](ClusterApi.md#GetClusterStatus) | **Get** /public/cluster/status | Get the status of a Cohesity Cluster.
[**UpdateAppSettings**](ClusterApi.md#UpdateAppSettings) | **Put** /public/cluster/appSettings | Update the app settings of the cluster.
[**UpdateCluster**](ClusterApi.md#UpdateCluster) | **Put** /public/cluster | Update the configuration of this Cohesity Cluster.



## GetAppSettings

> AppsConfig GetAppSettings(ctx).Execute()

Gets the app settings for the cluster.



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

    request := cohesitysdk.ApiGetAppSettingsRequest{
    }

    resp, r, err := api_client.ClusterApi.GetAppSettings(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterApi.GetAppSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppSettings`: AppsConfig
    fmt.Fprintf(os.Stdout, "Response from `ClusterApi.GetAppSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppSettingsRequest struct via the builder pattern


### Return type

[**AppsConfig**](AppsConfig.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBasicClusterInfo

> BasicClusterInfo GetBasicClusterInfo(ctx).Execute()

List details about the Cohesity Cluster such as the name, type, version, language, locale and domains. This operation does not require authentication.



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

    request := cohesitysdk.ApiGetBasicClusterInfoRequest{
    }

    resp, r, err := api_client.ClusterApi.GetBasicClusterInfo(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterApi.GetBasicClusterInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBasicClusterInfo`: BasicClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `ClusterApi.GetBasicClusterInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBasicClusterInfoRequest struct via the builder pattern


### Return type

[**BasicClusterInfo**](BasicClusterInfo.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCluster

> Cluster GetCluster(ctx).FetchStats(fetchStats).FetchTimeSeriesSchema(fetchTimeSeriesSchema).Execute()

List details about this Cohesity Cluster.



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
    fetchStats := true // bool | If 'true', also get statistics about the Cohesity Cluster. (optional)
    fetchTimeSeriesSchema := true // bool | Specifies whether to get time series schema info of the cluster. (optional)

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

    request := cohesitysdk.ApiGetClusterRequest{
        FetchStats: &fetchStats
        FetchTimeSeriesSchema: &fetchTimeSeriesSchema
    }

    resp, r, err := api_client.ClusterApi.GetCluster(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterApi.GetCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCluster`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClusterApi.GetCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fetchStats** | **bool** | If &#39;true&#39;, also get statistics about the Cohesity Cluster. | 
 **fetchTimeSeriesSchema** | **bool** | Specifies whether to get time series schema info of the cluster. | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterStatus

> ClusterStatusResult GetClusterStatus(ctx).Execute()

Get the status of a Cohesity Cluster.



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

    request := cohesitysdk.ApiGetClusterStatusRequest{
    }

    resp, r, err := api_client.ClusterApi.GetClusterStatus(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterApi.GetClusterStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterStatus`: ClusterStatusResult
    fmt.Fprintf(os.Stdout, "Response from `ClusterApi.GetClusterStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterStatusRequest struct via the builder pattern


### Return type

[**ClusterStatusResult**](ClusterStatusResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppSettings

> AppsConfig UpdateAppSettings(ctx).Body(body).Execute()

Update the app settings of the cluster.



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
    body := *openapiclient.NewAppsConfig() // AppsConfig | Update App Settings Parameter. (optional)

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

    request := cohesitysdk.ApiUpdateAppSettingsRequest{
        Body: &body
    }

    resp, r, err := api_client.ClusterApi.UpdateAppSettings(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterApi.UpdateAppSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAppSettings`: AppsConfig
    fmt.Fprintf(os.Stdout, "Response from `ClusterApi.UpdateAppSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AppsConfig**](AppsConfig.md) | Update App Settings Parameter. | 

### Return type

[**AppsConfig**](AppsConfig.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCluster

> Cluster UpdateCluster(ctx).Body(body).Execute()

Update the configuration of this Cohesity Cluster.



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
    body := *openapiclient.NewUpdateClusterParams() // UpdateClusterParams | Update Cluster Parameter. (optional)

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

    request := cohesitysdk.ApiUpdateClusterRequest{
        Body: &body
    }

    resp, r, err := api_client.ClusterApi.UpdateCluster(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterApi.UpdateCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCluster`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClusterApi.UpdateCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateClusterParams**](UpdateClusterParams.md) | Update Cluster Parameter. | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

