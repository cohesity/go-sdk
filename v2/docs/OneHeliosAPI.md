# \OneHeliosAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServicesStatus**](OneHeliosAPI.md#GetServicesStatus) | **Get** /helios/services/ondemand/{serviceName}/status | Get status of an ondemand service
[**InstallLogs**](OneHeliosAPI.md#InstallLogs) | **Get** /helios/services/install/logs | Get Helios install logs.
[**PerformServiceAction**](OneHeliosAPI.md#PerformServiceAction) | **Put** /helios/services/ondemand/{serviceName}/action | Perform action to enable/disable an on demand service.
[**ServicesHealth**](OneHeliosAPI.md#ServicesHealth) | **Get** /helios/health/services | Get Health Status for Services
[**UpgradeLogs**](OneHeliosAPI.md#UpgradeLogs) | **Get** /helios/services/upgrade/logs | Get Helios Upgrade Logs



## GetServicesStatus

> ServicesStatusResponse GetServicesStatus(ctx, serviceName).Execute()

Get status of an ondemand service



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	serviceName := "serviceName_example" // string | The name of the service to get status.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OneHeliosAPI.GetServicesStatus(context.Background(), serviceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OneHeliosAPI.GetServicesStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServicesStatus`: ServicesStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `OneHeliosAPI.GetServicesStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceName** | **string** | The name of the service to get status. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServicesStatusResponse**](ServicesStatusResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallLogs

> InstallLogsResponse InstallLogs(ctx).Execute()

Get Helios install logs.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OneHeliosAPI.InstallLogs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OneHeliosAPI.InstallLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InstallLogs`: InstallLogsResponse
	fmt.Fprintf(os.Stdout, "Response from `OneHeliosAPI.InstallLogs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInstallLogsRequest struct via the builder pattern


### Return type

[**InstallLogsResponse**](InstallLogsResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PerformServiceAction

> ServiceActionResponse PerformServiceAction(ctx, serviceName).Execute()

Perform action to enable/disable an on demand service.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	serviceName := "serviceName_example" // string | The name of the service.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OneHeliosAPI.PerformServiceAction(context.Background(), serviceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OneHeliosAPI.PerformServiceAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PerformServiceAction`: ServiceActionResponse
	fmt.Fprintf(os.Stdout, "Response from `OneHeliosAPI.PerformServiceAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceName** | **string** | The name of the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPerformServiceActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceActionResponse**](ServiceActionResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesHealth

> ServicesHealthGetResponse ServicesHealth(ctx).Execute()

Get Health Status for Services



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OneHeliosAPI.ServicesHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OneHeliosAPI.ServicesHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicesHealth`: ServicesHealthGetResponse
	fmt.Fprintf(os.Stdout, "Response from `OneHeliosAPI.ServicesHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicesHealthRequest struct via the builder pattern


### Return type

[**ServicesHealthGetResponse**](ServicesHealthGetResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeLogs

> UpgradeLogsResponse UpgradeLogs(ctx).Execute()

Get Helios Upgrade Logs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OneHeliosAPI.UpgradeLogs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OneHeliosAPI.UpgradeLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradeLogs`: UpgradeLogsResponse
	fmt.Fprintf(os.Stdout, "Response from `OneHeliosAPI.UpgradeLogs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeLogsRequest struct via the builder pattern


### Return type

[**UpgradeLogsResponse**](UpgradeLogsResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

