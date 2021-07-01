# \AppInstanceApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAppInstances**](AppInstanceApi.md#GetAppInstances) | **Get** /public/appInstances | Lists the app instances.
[**LaunchAppInstance**](AppInstanceApi.md#LaunchAppInstance) | **Post** /public/appInstances | Starts the application instance launch on the cluster.
[**UpdateAppInstanceSettings**](AppInstanceApi.md#UpdateAppInstanceSettings) | **Put** /public/appInstanceSettings/{appInstanceId} | Updates app instance settings.
[**UpdateAppInstanceState**](AppInstanceApi.md#UpdateAppInstanceState) | **Put** /public/appInstances/{appInstanceId}/states | Updates app instance state.



## GetAppInstances

> []AppInstance GetAppInstances(ctx).Execute()

Lists the app instances.



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

    request := cohesitysdk.ApiGetAppInstancesRequest{
    }

    resp, r, err := api_client.AppInstanceApi.GetAppInstances(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppInstanceApi.GetAppInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppInstances`: []AppInstance
    fmt.Fprintf(os.Stdout, "Response from `AppInstanceApi.GetAppInstances`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppInstancesRequest struct via the builder pattern


### Return type

[**[]AppInstance**](AppInstance.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LaunchAppInstance

> AppInstanceIdParameter LaunchAppInstance(ctx).Body(body).Execute()

Starts the application instance launch on the cluster.



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
    body := *openapiclient.NewLaunchAppInstance() // LaunchAppInstance | Request to launch app.

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

    request := cohesitysdk.ApiLaunchAppInstanceRequest{
        body: &Body
    }

    resp, r, err := api_client.AppInstanceApi.LaunchAppInstance(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppInstanceApi.LaunchAppInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LaunchAppInstance`: AppInstanceIdParameter
    fmt.Fprintf(os.Stdout, "Response from `AppInstanceApi.LaunchAppInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLaunchAppInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LaunchAppInstance**](LaunchAppInstance.md) | Request to launch app. | 

### Return type

[**AppInstanceIdParameter**](AppInstanceIdParameter.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppInstanceSettings

> UpdateAppInstanceSettings(ctx, appInstanceId).Body(body).Execute()

Updates app instance settings.



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
    appInstanceId := int64(789) // int64 | Specifies the app instance Id.
    body := *openapiclient.NewUpdateAppInstanceSettingsParameters() // UpdateAppInstanceSettingsParameters | Request to update app instance settings.

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

    request := cohesitysdk.ApiUpdateAppInstanceSettingsRequest{
        appInstanceId: &AppInstanceId
        body: &Body
    }

    resp, r, err := api_client.AppInstanceApi.UpdateAppInstanceSettings(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppInstanceApi.UpdateAppInstanceSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appInstanceId** | **int64** | Specifies the app instance Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppInstanceSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateAppInstanceSettingsParameters**](UpdateAppInstanceSettingsParameters.md) | Request to update app instance settings. | 

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


## UpdateAppInstanceState

> UpdateAppInstanceState(ctx, appInstanceId).Body(body).Execute()

Updates app instance state.



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
    appInstanceId := int64(789) // int64 | Specifies the app instance Id.
    body := *openapiclient.NewUpdateAppInstanceStateParameters() // UpdateAppInstanceStateParameters | Request to update app instance state.

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

    request := cohesitysdk.ApiUpdateAppInstanceStateRequest{
        appInstanceId: &AppInstanceId
        body: &Body
    }

    resp, r, err := api_client.AppInstanceApi.UpdateAppInstanceState(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppInstanceApi.UpdateAppInstanceState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appInstanceId** | **int64** | Specifies the app instance Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppInstanceStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateAppInstanceStateParameters**](UpdateAppInstanceStateParameters.md) | Request to update app instance state. | 

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

