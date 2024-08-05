# \KeystoneAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKeystone**](KeystoneAPI.md#CreateKeystone) | **Post** /keystones | Create a Keystone configuration.
[**DeleteKeystone**](KeystoneAPI.md#DeleteKeystone) | **Delete** /keystones/{id} | Delete a Keystone configuration.
[**GetKeystones**](KeystoneAPI.md#GetKeystones) | **Get** /keystones | Get Keystones.
[**GetKeystonesById**](KeystoneAPI.md#GetKeystonesById) | **Get** /keystones/{id} | Get a Keystone by its id.
[**UpdateKeystone**](KeystoneAPI.md#UpdateKeystone) | **Put** /keystones/{id} | Update a Keystone configuration.



## CreateKeystone

> Keystone CreateKeystone(ctx).Body(body).Execute()

Create a Keystone configuration.



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
	body := *openapiclient.NewCreateKeystoneRequest(*openapiclient.NewKeystoneCredentialsAdminCreds("Domain_example", "Username_example"), *openapiclient.NewKeystoneCredentialsScope("Type_example"), "AuthUrl_example", "Name_example") // CreateKeystoneRequest | Specifies the paremters to create a Keystone configuration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeystoneAPI.CreateKeystone(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeystoneAPI.CreateKeystone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateKeystone`: Keystone
	fmt.Fprintf(os.Stdout, "Response from `KeystoneAPI.CreateKeystone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKeystoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateKeystoneRequest**](CreateKeystoneRequest.md) | Specifies the paremters to create a Keystone configuration. | 

### Return type

[**Keystone**](Keystone.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKeystone

> DeleteKeystone(ctx, id).AdminPassword(adminPassword).Execute()

Delete a Keystone configuration.



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
	id := int64(789) // int64 | Specifies the Keystone id.
	adminPassword := "adminPassword_example" // string | Specifies the password of Keystone administrator.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KeystoneAPI.DeleteKeystone(context.Background(), id).AdminPassword(adminPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeystoneAPI.DeleteKeystone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the Keystone id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeystoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adminPassword** | **string** | Specifies the password of Keystone administrator. | 

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


## GetKeystones

> Keystones GetKeystones(ctx).Names(names).TenantIds(tenantIds).IncludeTenants(includeTenants).Execute()

Get Keystones.



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
	names := []string{"Inner_example"} // []string | Specifies a list of Keystone names. (optional)
	tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
	includeTenants := true // bool | If true, the response will include Keystones which were created by all tenants which the current user has permission to see. If false, then only Keystones created by the current user will be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeystoneAPI.GetKeystones(context.Background()).Names(names).TenantIds(tenantIds).IncludeTenants(includeTenants).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeystoneAPI.GetKeystones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKeystones`: Keystones
	fmt.Fprintf(os.Stdout, "Response from `KeystoneAPI.GetKeystones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKeystonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **names** | **[]string** | Specifies a list of Keystone names. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **includeTenants** | **bool** | If true, the response will include Keystones which were created by all tenants which the current user has permission to see. If false, then only Keystones created by the current user will be returned. | 

### Return type

[**Keystones**](Keystones.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKeystonesById

> Keystone GetKeystonesById(ctx, id).Execute()

Get a Keystone by its id.



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
	id := int64(789) // int64 | Specifies the Keystone id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeystoneAPI.GetKeystonesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeystoneAPI.GetKeystonesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKeystonesById`: Keystone
	fmt.Fprintf(os.Stdout, "Response from `KeystoneAPI.GetKeystonesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the Keystone id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeystonesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Keystone**](Keystone.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKeystone

> Keystone UpdateKeystone(ctx, id).Body(body).Execute()

Update a Keystone configuration.



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
	id := int64(789) // int64 | Specifies the Keystone id.
	body := *openapiclient.NewUpdateKeystoneRequest(*openapiclient.NewKeystoneCredentialsAdminCreds("Domain_example", "Username_example"), *openapiclient.NewKeystoneCredentialsScope("Type_example"), "AuthUrl_example", "Name_example") // UpdateKeystoneRequest | Specifies the paremters to update a Keystone configuration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeystoneAPI.UpdateKeystone(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeystoneAPI.UpdateKeystone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateKeystone`: Keystone
	fmt.Fprintf(os.Stdout, "Response from `KeystoneAPI.UpdateKeystone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the Keystone id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKeystoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateKeystoneRequest**](UpdateKeystoneRequest.md) | Specifies the paremters to update a Keystone configuration. | 

### Return type

[**Keystone**](Keystone.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

