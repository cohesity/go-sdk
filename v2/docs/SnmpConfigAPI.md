# \SnmpConfigAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSnmpConfig**](SnmpConfigAPI.md#GetSnmpConfig) | **Get** /snmp/config | Get Snmp Config
[**UpdateSnmpConfig**](SnmpConfigAPI.md#UpdateSnmpConfig) | **Put** /snmp/config | Update Snmp Config



## GetSnmpConfig

> SnmpConfig GetSnmpConfig(ctx).Execute()

Get Snmp Config



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
	resp, r, err := apiClient.SnmpConfigAPI.GetSnmpConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnmpConfigAPI.GetSnmpConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnmpConfig`: SnmpConfig
	fmt.Fprintf(os.Stdout, "Response from `SnmpConfigAPI.GetSnmpConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnmpConfigRequest struct via the builder pattern


### Return type

[**SnmpConfig**](SnmpConfig.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSnmpConfig

> SnmpConfig UpdateSnmpConfig(ctx).Body(body).Execute()

Update Snmp Config



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
	body := *openapiclient.NewSnmpConfig() // SnmpConfig | Modify Snmp Config with given parameters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnmpConfigAPI.UpdateSnmpConfig(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnmpConfigAPI.UpdateSnmpConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSnmpConfig`: SnmpConfig
	fmt.Fprintf(os.Stdout, "Response from `SnmpConfigAPI.UpdateSnmpConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSnmpConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SnmpConfig**](SnmpConfig.md) | Modify Snmp Config with given parameters. | 

### Return type

[**SnmpConfig**](SnmpConfig.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

