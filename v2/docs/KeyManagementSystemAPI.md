# \KeyManagementSystemAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddKmsConfiguration**](KeyManagementSystemAPI.md#AddKmsConfiguration) | **Post** /kms | Add KMS
[**DeleteKmsConfig**](KeyManagementSystemAPI.md#DeleteKmsConfig) | **Delete** /kms/{id} | Delete KMS
[**GetKmsConfigurations**](KeyManagementSystemAPI.md#GetKmsConfigurations) | **Get** /kms | Get KMS
[**UpdateKmsConfiguration**](KeyManagementSystemAPI.md#UpdateKmsConfiguration) | **Put** /kms/{id} | Update KMS



## AddKmsConfiguration

> KmsConfiguration AddKmsConfiguration(ctx).Body(body).Execute()

Add KMS



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
	body := *openapiclient.NewKmsConfigurationCreateParams("Name_example", "Type_example") // KmsConfigurationCreateParams | Parameters to add KMS on the cluster.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyManagementSystemAPI.AddKmsConfiguration(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementSystemAPI.AddKmsConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddKmsConfiguration`: KmsConfiguration
	fmt.Fprintf(os.Stdout, "Response from `KeyManagementSystemAPI.AddKmsConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddKmsConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**KmsConfigurationCreateParams**](KmsConfigurationCreateParams.md) | Parameters to add KMS on the cluster. | 

### Return type

[**KmsConfiguration**](KmsConfiguration.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKmsConfig

> DeleteKmsConfig(ctx, id).Execute()

Delete KMS



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
	id := int64(789) // int64 | ID of the KMS configured on the cluster.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KeyManagementSystemAPI.DeleteKmsConfig(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementSystemAPI.DeleteKmsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of the KMS configured on the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKmsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKmsConfigurations

> KmsConfigurations GetKmsConfigurations(ctx).Ids(ids).IncludeRpaasKms(includeRpaasKms).Execute()

Get KMS



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
	ids := []int64{int64(123)} // []int64 | Ids of KMS configured on the cluster. (optional)
	includeRpaasKms := true // bool | If true, returns KMS that are configured by FortKnox. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyManagementSystemAPI.GetKmsConfigurations(context.Background()).Ids(ids).IncludeRpaasKms(includeRpaasKms).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementSystemAPI.GetKmsConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKmsConfigurations`: KmsConfigurations
	fmt.Fprintf(os.Stdout, "Response from `KeyManagementSystemAPI.GetKmsConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKmsConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]int64** | Ids of KMS configured on the cluster. | 
 **includeRpaasKms** | **bool** | If true, returns KMS that are configured by FortKnox. | 

### Return type

[**KmsConfigurations**](KmsConfigurations.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKmsConfiguration

> KmsConfiguration UpdateKmsConfiguration(ctx, id).Body(body).Execute()

Update KMS



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
	id := int64(789) // int64 | ID of the KMS configured on the cluster.
	body := *openapiclient.NewKmsConfigurationUpdateParams("Name_example") // KmsConfigurationUpdateParams | Parameters to update KMS on the cluster.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyManagementSystemAPI.UpdateKmsConfiguration(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementSystemAPI.UpdateKmsConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateKmsConfiguration`: KmsConfiguration
	fmt.Fprintf(os.Stdout, "Response from `KeyManagementSystemAPI.UpdateKmsConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of the KMS configured on the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKmsConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**KmsConfigurationUpdateParams**](KmsConfigurationUpdateParams.md) | Parameters to update KMS on the cluster. | 

### Return type

[**KmsConfiguration**](KmsConfiguration.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

