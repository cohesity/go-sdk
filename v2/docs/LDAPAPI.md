# \LDAPAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLdapProvider**](LDAPAPI.md#CreateLdapProvider) | **Post** /ldap | Create Ldap provider.
[**DeleteLdapProvider**](LDAPAPI.md#DeleteLdapProvider) | **Delete** /ldap/{id} | Delete LDAP provider.
[**GetLdapConnectionStatus**](LDAPAPI.md#GetLdapConnectionStatus) | **Get** /ldap/{id}/connection-status | Get LDAP connection status.
[**GetLdaps**](LDAPAPI.md#GetLdaps) | **Get** /ldap | Get Groups.
[**UpdateLdapProvider**](LDAPAPI.md#UpdateLdapProvider) | **Put** /ldap | Update Ldap provider.



## CreateLdapProvider

> Ldap CreateLdapProvider(ctx).Body(body).Execute()

Create Ldap provider.



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
	body := *openapiclient.NewLdap("AuthType_example", "BaseDistinguishedName_example", "Name_example") // Ldap | Specifies the parameters to create Ldap provider.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LDAPAPI.CreateLdapProvider(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LDAPAPI.CreateLdapProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLdapProvider`: Ldap
	fmt.Fprintf(os.Stdout, "Response from `LDAPAPI.CreateLdapProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLdapProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Ldap**](Ldap.md) | Specifies the parameters to create Ldap provider. | 

### Return type

[**Ldap**](Ldap.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLdapProvider

> DeleteLdapProvider(ctx, id).Execute()

Delete LDAP provider.



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
	id := int64(789) // int64 | Specifies the LDAP Id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LDAPAPI.DeleteLdapProvider(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LDAPAPI.DeleteLdapProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the LDAP Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLdapProviderRequest struct via the builder pattern


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


## GetLdapConnectionStatus

> LdapStatus GetLdapConnectionStatus(ctx, id).Execute()

Get LDAP connection status.



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
	id := int64(789) // int64 | Specifies the LDAP id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LDAPAPI.GetLdapConnectionStatus(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LDAPAPI.GetLdapConnectionStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLdapConnectionStatus`: LdapStatus
	fmt.Fprintf(os.Stdout, "Response from `LDAPAPI.GetLdapConnectionStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the LDAP id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLdapConnectionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LdapStatus**](LdapStatus.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLdaps

> Ldaps GetLdaps(ctx).Ids(ids).TenantIds(tenantIds).IncludeTenants(includeTenants).Execute()

Get Groups.



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
	ids := []int64{int64(123)} // []int64 | Specifies a list of ids to filter. (optional)
	tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which LDAPs are to be returned. (optional)
	includeTenants := true // bool | IncludeTenants specifies if LDAPs of all the tenants under the hierarchy of the logged in user's organization should be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LDAPAPI.GetLdaps(context.Background()).Ids(ids).TenantIds(tenantIds).IncludeTenants(includeTenants).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LDAPAPI.GetLdaps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLdaps`: Ldaps
	fmt.Fprintf(os.Stdout, "Response from `LDAPAPI.GetLdaps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLdapsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]int64** | Specifies a list of ids to filter. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which LDAPs are to be returned. | 
 **includeTenants** | **bool** | IncludeTenants specifies if LDAPs of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. | 

### Return type

[**Ldaps**](Ldaps.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLdapProvider

> Ldap UpdateLdapProvider(ctx).Body(body).Execute()

Update Ldap provider.



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
	body := *openapiclient.NewLdap("AuthType_example", "BaseDistinguishedName_example", "Name_example") // Ldap | Specifies the parameters to update Ldap provider.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LDAPAPI.UpdateLdapProvider(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LDAPAPI.UpdateLdapProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLdapProvider`: Ldap
	fmt.Fprintf(os.Stdout, "Response from `LDAPAPI.UpdateLdapProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLdapProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Ldap**](Ldap.md) | Specifies the parameters to update Ldap provider. | 

### Return type

[**Ldap**](Ldap.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

