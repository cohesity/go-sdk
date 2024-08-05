# \IdentityProviderAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIdentity**](IdentityProviderAPI.md#CreateIdentity) | **Post** /identity-providers | Configure Identity Provider
[**CreateIdentityProvider**](IdentityProviderAPI.md#CreateIdentityProvider) | **Post** /idps | Configure identity provider
[**DeleteIdentity**](IdentityProviderAPI.md#DeleteIdentity) | **Delete** /identity-providers/{id} | Delete Identity Provider
[**DeleteIdentityProvider**](IdentityProviderAPI.md#DeleteIdentityProvider) | **Delete** /idps/{id} | Delete identity provider
[**GetIdentities**](IdentityProviderAPI.md#GetIdentities) | **Get** /identity-providers | Get Identities
[**GetIdentityProviders**](IdentityProviderAPI.md#GetIdentityProviders) | **Get** /idps | Get identity providers
[**IdpsLogin**](IdentityProviderAPI.md#IdpsLogin) | **Get** /idps/login | Login to cluster using idp
[**PerformIdentityAction**](IdentityProviderAPI.md#PerformIdentityAction) | **Post** /identity-providers/actions | Perform Identity Action
[**UpdateIdentity**](IdentityProviderAPI.md#UpdateIdentity) | **Put** /identity-providers/{id} | Update Identity Provider
[**UpdateIdentityProvider**](IdentityProviderAPI.md#UpdateIdentityProvider) | **Put** /idps/{id} | Update identity provider



## CreateIdentity

> IdentityConfig CreateIdentity(ctx).Body(body).Execute()

Configure Identity Provider



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
	body := *openapiclient.NewIdentityConfig("Domain_example", "IdentityProviderType_example") // IdentityConfig | Specifies parameters to configure Identity

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderAPI.CreateIdentity(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.CreateIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIdentity`: IdentityConfig
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.CreateIdentity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IdentityConfig**](IdentityConfig.md) | Specifies parameters to configure Identity | 

### Return type

[**IdentityConfig**](IdentityConfig.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIdentityProvider

> IdentityProviderConfiguration CreateIdentityProvider(ctx).Body(body).Execute()

Configure identity provider



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
	body := *openapiclient.NewCreateIdpRequestParams("Certificate_example", "IssuerId_example", "SsoUrl_example", "Domain_example", "Name_example") // CreateIdpRequestParams | Specifies parameters to configure identity provider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderAPI.CreateIdentityProvider(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.CreateIdentityProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIdentityProvider`: IdentityProviderConfiguration
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.CreateIdentityProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateIdpRequestParams**](CreateIdpRequestParams.md) | Specifies parameters to configure identity provider | 

### Return type

[**IdentityProviderConfiguration**](IdentityProviderConfiguration.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentity

> DeleteIdentity(ctx, id).Execute()

Delete Identity Provider



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
	id := int64(789) // int64 | Specifies id of identity provider configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentityProviderAPI.DeleteIdentity(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.DeleteIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies id of identity provider configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityRequest struct via the builder pattern


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


## DeleteIdentityProvider

> DeleteIdentityProvider(ctx, id).Execute()

Delete identity provider



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
	id := int64(789) // int64 | Specifies id of idp configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentityProviderAPI.DeleteIdentityProvider(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.DeleteIdentityProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies id of idp configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityProviderRequest struct via the builder pattern


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


## GetIdentities

> IdentityConfigs GetIdentities(ctx).Ids(ids).TenantIds(tenantIds).Domains(domains).IncludeAllTenants(includeAllTenants).Execute()

Get Identities



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
	ids := []int64{int64(123)} // []int64 | Specifies IDs of configured identity providers (optional)
	tenantIds := []string{"Inner_example"} // []string | Specifies the tenant id's to get IDPs configured on tenants (optional)
	domains := []string{"Inner_example"} // []string | Specifies domains of the IDP configurations (optional)
	includeAllTenants := true // bool | Specifies if IDP configurations on all the tenants under the hierarchy of the logged in user should be returned (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderAPI.GetIdentities(context.Background()).Ids(ids).TenantIds(tenantIds).Domains(domains).IncludeAllTenants(includeAllTenants).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.GetIdentities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentities`: IdentityConfigs
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.GetIdentities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]int64** | Specifies IDs of configured identity providers | 
 **tenantIds** | **[]string** | Specifies the tenant id&#39;s to get IDPs configured on tenants | 
 **domains** | **[]string** | Specifies domains of the IDP configurations | 
 **includeAllTenants** | **bool** | Specifies if IDP configurations on all the tenants under the hierarchy of the logged in user should be returned | 

### Return type

[**IdentityConfigs**](IdentityConfigs.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityProviders

> IdentityProviderConfigurations GetIdentityProviders(ctx).Ids(ids).TenantIds(tenantIds).Names(names).Domains(domains).IncludeAllTenants(includeAllTenants).Execute()

Get identity providers



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
	ids := []int64{int64(123)} // []int64 | Specifies ids of configured identity providers (optional)
	tenantIds := []string{"Inner_example"} // []string | Specifies the tenant id's to get idps configured on tenants (optional)
	names := []string{"Inner_example"} // []string | Specifies the names of the identity providers (optional)
	domains := []string{"Inner_example"} // []string | Specifies domains of the idp configurations (optional)
	includeAllTenants := true // bool | Specifies if idp configurations on all the tenants under the hierarchy of the logged in user should be returned (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderAPI.GetIdentityProviders(context.Background()).Ids(ids).TenantIds(tenantIds).Names(names).Domains(domains).IncludeAllTenants(includeAllTenants).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.GetIdentityProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentityProviders`: IdentityProviderConfigurations
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.GetIdentityProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]int64** | Specifies ids of configured identity providers | 
 **tenantIds** | **[]string** | Specifies the tenant id&#39;s to get idps configured on tenants | 
 **names** | **[]string** | Specifies the names of the identity providers | 
 **domains** | **[]string** | Specifies domains of the idp configurations | 
 **includeAllTenants** | **bool** | Specifies if idp configurations on all the tenants under the hierarchy of the logged in user should be returned | 

### Return type

[**IdentityProviderConfigurations**](IdentityProviderConfigurations.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdpsLogin

> Error IdpsLogin(ctx).TenantId(tenantId).Execute()

Login to cluster using idp



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
	tenantId := "tenantId_example" // string | Specifies an optional tenantId for which the SSO login should be done. If this is not specified, cluster SSO login is done. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderAPI.IdpsLogin(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.IdpsLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IdpsLogin`: Error
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.IdpsLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdpsLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Specifies an optional tenantId for which the SSO login should be done. If this is not specified, cluster SSO login is done. | 

### Return type

[**Error**](Error.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PerformIdentityAction

> IdentityAction PerformIdentityAction(ctx).Body(body).Execute()

Perform Identity Action



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
	body := *openapiclient.NewIdentityAction("IdentityProviderType_example") // IdentityAction | Specifies parameters perform an identity action.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderAPI.PerformIdentityAction(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.PerformIdentityAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PerformIdentityAction`: IdentityAction
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.PerformIdentityAction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPerformIdentityActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IdentityAction**](IdentityAction.md) | Specifies parameters perform an identity action. | 

### Return type

[**IdentityAction**](IdentityAction.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdentity

> IdentityConfig UpdateIdentity(ctx, id).Body(body).Execute()

Update Identity Provider



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
	id := int64(789) // int64 | Specifies id of identity provider configuration
	body := *openapiclient.NewIdentityConfig("Domain_example", "IdentityProviderType_example") // IdentityConfig | Specifies parameters to update identity provider configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderAPI.UpdateIdentity(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.UpdateIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIdentity`: IdentityConfig
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.UpdateIdentity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies id of identity provider configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IdentityConfig**](IdentityConfig.md) | Specifies parameters to update identity provider configuration | 

### Return type

[**IdentityConfig**](IdentityConfig.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdentityProvider

> IdentityProviderConfiguration UpdateIdentityProvider(ctx, id).Body(body).Execute()

Update identity provider



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
	id := int64(789) // int64 | Specifies id of idp configuration
	body := *openapiclient.NewUpdateIdpRequestParams("Certificate_example", "IssuerId_example", "SsoUrl_example") // UpdateIdpRequestParams | Specifies parameters to update identity provider configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderAPI.UpdateIdentityProvider(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.UpdateIdentityProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIdentityProvider`: IdentityProviderConfiguration
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.UpdateIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies id of idp configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateIdpRequestParams**](UpdateIdpRequestParams.md) | Specifies parameters to update identity provider configuration | 

### Return type

[**IdentityProviderConfiguration**](IdentityProviderConfiguration.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

