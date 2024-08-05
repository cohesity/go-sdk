# \SupportAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSupportUserConfig**](SupportAPI.md#GetSupportUserConfig) | **Get** /support-user/config | Get support user configuration.
[**UpdateSupportUserConfig**](SupportAPI.md#UpdateSupportUserConfig) | **Put** /support-user/config | Update support user configuration.
[**ValidateSupportUserCreds**](SupportAPI.md#ValidateSupportUserCreds) | **Post** /support-user/config/validate | Validates the support user credentials.



## GetSupportUserConfig

> SupportUserConfig GetSupportUserConfig(ctx).Execute()

Get support user configuration.



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
	resp, r, err := apiClient.SupportAPI.GetSupportUserConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportAPI.GetSupportUserConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportUserConfig`: SupportUserConfig
	fmt.Fprintf(os.Stdout, "Response from `SupportAPI.GetSupportUserConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportUserConfigRequest struct via the builder pattern


### Return type

[**SupportUserConfig**](SupportUserConfig.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSupportUserConfig

> SuccessResp UpdateSupportUserConfig(ctx).Body(body).Execute()

Update support user configuration.



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
	body := *openapiclient.NewUpdateSupportUserParams() // UpdateSupportUserParams | Specifies the support user configuration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportAPI.UpdateSupportUserConfig(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportAPI.UpdateSupportUserConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSupportUserConfig`: SuccessResp
	fmt.Fprintf(os.Stdout, "Response from `SupportAPI.UpdateSupportUserConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSupportUserConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateSupportUserParams**](UpdateSupportUserParams.md) | Specifies the support user configuration. | 

### Return type

[**SuccessResp**](SuccessResp.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateSupportUserCreds

> SuccessResp ValidateSupportUserCreds(ctx).Body(body).Execute()

Validates the support user credentials.



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
	body := *openapiclient.NewValidateSupportUserCredParams("Password_example") // ValidateSupportUserCredParams | Specifies the support user credentials.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportAPI.ValidateSupportUserCreds(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportAPI.ValidateSupportUserCreds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateSupportUserCreds`: SuccessResp
	fmt.Fprintf(os.Stdout, "Response from `SupportAPI.ValidateSupportUserCreds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateSupportUserCredsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ValidateSupportUserCredParams**](ValidateSupportUserCredParams.md) | Specifies the support user credentials. | 

### Return type

[**SuccessResp**](SuccessResp.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

