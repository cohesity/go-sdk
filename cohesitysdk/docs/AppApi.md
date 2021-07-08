# \AppApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApps**](AppApi.md#GetApps) | **Get** /public/apps | Lists the apps.
[**InstallApp**](AppApi.md#InstallApp) | **Post** /public/apps/{appUid}/versions/{version} | Starts the application installation on the cluster.
[**UninstallApp**](AppApi.md#UninstallApp) | **Delete** /public/apps/{appUid}/versions/{version} | Starts the application uninstall from the cluster.
[**UploadApp**](AppApi.md#UploadApp) | **Post** /public/apps | Upload and install an app from image.



## GetApps

> []App GetApps(ctx).Execute()

Lists the apps.



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

    request := cohesitysdk.ApiGetAppsRequest{
    }

    resp, r, err := api_client.AppApi.GetApps(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppApi.GetApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApps`: []App
    fmt.Fprintf(os.Stdout, "Response from `AppApi.GetApps`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppsRequest struct via the builder pattern


### Return type

[**[]App**](App.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallApp

> App InstallApp(ctx, appUid, version).Execute()

Starts the application installation on the cluster.



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
    appUid := int64(789) // int64 | Specifies the app Id.
    version := int64(789) // int64 | Specifies the app version.

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

    request := cohesitysdk.ApiInstallAppRequest{
        AppUid: &appUid
        Version: &version
    }

    resp, r, err := api_client.AppApi.InstallApp(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppApi.InstallApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstallApp`: App
    fmt.Fprintf(os.Stdout, "Response from `AppApi.InstallApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appUid** | **int64** | Specifies the app Id. | 
**version** | **int64** | Specifies the app version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstallAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**App**](App.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UninstallApp

> UninstallApp(ctx, appUid, version).Execute()

Starts the application uninstall from the cluster.



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
    appUid := int64(789) // int64 | Specifies the app Id.
    version := int64(789) // int64 | Specifies the app version.

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

    request := cohesitysdk.ApiUninstallAppRequest{
        AppUid: &appUid
        Version: &version
    }

    resp, r, err := api_client.AppApi.UninstallApp(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppApi.UninstallApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appUid** | **int64** | Specifies the app Id. | 
**version** | **int64** | Specifies the app version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUninstallAppRequest struct via the builder pattern


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


## UploadApp

> App UploadApp(ctx).Execute()

Upload and install an app from image.



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

    request := cohesitysdk.ApiUploadAppRequest{
    }

    resp, r, err := api_client.AppApi.UploadApp(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppApi.UploadApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadApp`: App
    fmt.Fprintf(os.Stdout, "Response from `AppApi.UploadApp`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUploadAppRequest struct via the builder pattern


### Return type

[**App**](App.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

