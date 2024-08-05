# \KerberosProviderAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetKerberosProviderById**](KerberosProviderAPI.md#GetKerberosProviderById) | **Get** /kerberos-providers/{id} | Get the Registered Kerberos Provider by id.
[**GetKerberosProviders**](KerberosProviderAPI.md#GetKerberosProviders) | **Get** /kerberos-providers | Get the list of Kerberos Providers.
[**RegisterKerberosProvider**](KerberosProviderAPI.md#RegisterKerberosProvider) | **Post** /kerberos-providers/register | Register a Kerberos Authentication Provider.
[**UnregisterKerberosProvider**](KerberosProviderAPI.md#UnregisterKerberosProvider) | **Post** /kerberos-providers/{id}/unregister | Unregister a Kerberos Provider.
[**UpdateKerberosProvider**](KerberosProviderAPI.md#UpdateKerberosProvider) | **Put** /kerberos-providers/{id} | Update the Kerberos Provider Registration.



## GetKerberosProviderById

> KerberosProvider GetKerberosProviderById(ctx, id).Execute()

Get the Registered Kerberos Provider by id.



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
	id := "id_example" // string | Specifies the id which will be of the pattern cluster_id:clusterincarnation_id:resource_id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KerberosProviderAPI.GetKerberosProviderById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KerberosProviderAPI.GetKerberosProviderById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKerberosProviderById`: KerberosProvider
	fmt.Fprintf(os.Stdout, "Response from `KerberosProviderAPI.GetKerberosProviderById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id which will be of the pattern cluster_id:clusterincarnation_id:resource_id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKerberosProviderByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KerberosProvider**](KerberosProvider.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKerberosProviders

> KerberosProviders GetKerberosProviders(ctx).RealmNames(realmNames).HasLDAP(hasLDAP).Ids(ids).KdcServers(kdcServers).Execute()

Get the list of Kerberos Providers.



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
	realmNames := []string{"Inner_example"} // []string | Filter by a list of realm names. (optional)
	hasLDAP := true // bool | Filter by whether LDAP is associated with the provider. (optional)
	ids := []int64{int64(123)} // []int64 | Filter by a list of Kerberos Provider Ids. (optional)
	kdcServers := []string{"Inner_example"} // []string | Filter by a list of KDC servers. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KerberosProviderAPI.GetKerberosProviders(context.Background()).RealmNames(realmNames).HasLDAP(hasLDAP).Ids(ids).KdcServers(kdcServers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KerberosProviderAPI.GetKerberosProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKerberosProviders`: KerberosProviders
	fmt.Fprintf(os.Stdout, "Response from `KerberosProviderAPI.GetKerberosProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKerberosProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **realmNames** | **[]string** | Filter by a list of realm names. | 
 **hasLDAP** | **bool** | Filter by whether LDAP is associated with the provider. | 
 **ids** | **[]int64** | Filter by a list of Kerberos Provider Ids. | 
 **kdcServers** | **[]string** | Filter by a list of KDC servers. | 

### Return type

[**KerberosProviders**](KerberosProviders.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterKerberosProvider

> KerberosProvider RegisterKerberosProvider(ctx).Body(body).Execute()

Register a Kerberos Authentication Provider.



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
	body := *openapiclient.NewRegisterOrUpdateKerberosProviderRequest("AdminServer_example", []string{"HostAlias_example"}, []string{"KdcServers_example"}, "RealmName_example", "AdminPassword_example") // RegisterOrUpdateKerberosProviderRequest | Specifies the parameters to Register a Kerberos Provider.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KerberosProviderAPI.RegisterKerberosProvider(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KerberosProviderAPI.RegisterKerberosProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterKerberosProvider`: KerberosProvider
	fmt.Fprintf(os.Stdout, "Response from `KerberosProviderAPI.RegisterKerberosProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterKerberosProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RegisterOrUpdateKerberosProviderRequest**](RegisterOrUpdateKerberosProviderRequest.md) | Specifies the parameters to Register a Kerberos Provider. | 

### Return type

[**KerberosProvider**](KerberosProvider.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnregisterKerberosProvider

> UnregisterKerberosProvider UnregisterKerberosProvider(ctx, id).Body(body).Execute()

Unregister a Kerberos Provider.



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
	id := "id_example" // string | Specifies the id.
	body := *openapiclient.NewUnregisterKerberosRequest("AdminPassword_example") // UnregisterKerberosRequest | Request to unregister a Kerberos Provider.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KerberosProviderAPI.UnregisterKerberosProvider(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KerberosProviderAPI.UnregisterKerberosProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnregisterKerberosProvider`: UnregisterKerberosProvider
	fmt.Fprintf(os.Stdout, "Response from `KerberosProviderAPI.UnregisterKerberosProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnregisterKerberosProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UnregisterKerberosRequest**](UnregisterKerberosRequest.md) | Request to unregister a Kerberos Provider. | 

### Return type

[**UnregisterKerberosProvider**](UnregisterKerberosProvider.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKerberosProvider

> KerberosProvider UpdateKerberosProvider(ctx, id).Body(body).Execute()

Update the Kerberos Provider Registration.



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
	id := "id_example" // string | Specifies the id which will be of the pattern cluster_id:clusterincarnation_id:resource_id.
	body := *openapiclient.NewRegisterOrUpdateKerberosProviderRequest("AdminServer_example", []string{"HostAlias_example"}, []string{"KdcServers_example"}, "RealmName_example", "AdminPassword_example") // RegisterOrUpdateKerberosProviderRequest | Request to update a Kerberos Provider.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KerberosProviderAPI.UpdateKerberosProvider(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KerberosProviderAPI.UpdateKerberosProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateKerberosProvider`: KerberosProvider
	fmt.Fprintf(os.Stdout, "Response from `KerberosProviderAPI.UpdateKerberosProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id which will be of the pattern cluster_id:clusterincarnation_id:resource_id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKerberosProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RegisterOrUpdateKerberosProviderRequest**](RegisterOrUpdateKerberosProviderRequest.md) | Request to update a Kerberos Provider. | 

### Return type

[**KerberosProvider**](KerberosProvider.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

