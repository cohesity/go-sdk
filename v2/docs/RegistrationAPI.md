# \RegistrationAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHeliosRegConfig**](RegistrationAPI.md#GetHeliosRegConfig) | **Get** /helios-registration-config | Lists the Helios Registration Config.
[**HeliosClaim**](RegistrationAPI.md#HeliosClaim) | **Post** /helios-registration | Register to Helios.



## GetHeliosRegConfig

> HeliosRegConfig GetHeliosRegConfig(ctx).Execute()

Lists the Helios Registration Config.



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
	resp, r, err := apiClient.RegistrationAPI.GetHeliosRegConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistrationAPI.GetHeliosRegConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHeliosRegConfig`: HeliosRegConfig
	fmt.Fprintf(os.Stdout, "Response from `RegistrationAPI.GetHeliosRegConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHeliosRegConfigRequest struct via the builder pattern


### Return type

[**HeliosRegConfig**](HeliosRegConfig.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeliosClaim

> HeliosClaim(ctx).Body(body).Execute()

Register to Helios.



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
	body := *openapiclient.NewHeliosClaimRequest("RegistrationToken_example") // HeliosClaimRequest | Specifies the parameters to claim to Helios.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RegistrationAPI.HeliosClaim(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistrationAPI.HeliosClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHeliosClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HeliosClaimRequest**](HeliosClaimRequest.md) | Specifies the parameters to claim to Helios. | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

