# \HeliosOnPremAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHeliosOnPremConfig**](HeliosOnPremAPI.md#GetHeliosOnPremConfig) | **Get** /helios-onprem/config | Retreive Helios OnPrem Configuration
[**UpdateHeliosOnPremConfig**](HeliosOnPremAPI.md#UpdateHeliosOnPremConfig) | **Put** /helios-onprem/config | Update Helios OnPrem Configuration



## GetHeliosOnPremConfig

> HeliosOnPremConfig GetHeliosOnPremConfig(ctx).Execute()

Retreive Helios OnPrem Configuration



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
	resp, r, err := apiClient.HeliosOnPremAPI.GetHeliosOnPremConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HeliosOnPremAPI.GetHeliosOnPremConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHeliosOnPremConfig`: HeliosOnPremConfig
	fmt.Fprintf(os.Stdout, "Response from `HeliosOnPremAPI.GetHeliosOnPremConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHeliosOnPremConfigRequest struct via the builder pattern


### Return type

[**HeliosOnPremConfig**](HeliosOnPremConfig.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHeliosOnPremConfig

> HeliosOnPremConfig UpdateHeliosOnPremConfig(ctx).Body(body).Execute()

Update Helios OnPrem Configuration



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
	body := *openapiclient.NewHeliosOnPremConfig("KubernetesSubnetCidr_example", "Name_example") // HeliosOnPremConfig | Specifies the parameters for config update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HeliosOnPremAPI.UpdateHeliosOnPremConfig(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HeliosOnPremAPI.UpdateHeliosOnPremConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHeliosOnPremConfig`: HeliosOnPremConfig
	fmt.Fprintf(os.Stdout, "Response from `HeliosOnPremAPI.UpdateHeliosOnPremConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHeliosOnPremConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HeliosOnPremConfig**](HeliosOnPremConfig.md) | Specifies the parameters for config update. | 

### Return type

[**HeliosOnPremConfig**](HeliosOnPremConfig.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

