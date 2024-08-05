# \SourceAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAzureApplications**](SourceAPI.md#CreateAzureApplications) | **Post** /data-protect/sources/microsoft365/azure-applications | Create Microsoft 365 Azure Applications for a given domain.
[**CreateOrUpdateAzureApplications**](SourceAPI.md#CreateOrUpdateAzureApplications) | **Put** /data-protect/sources/microsoft365/azure-applications | Create/Update Microsoft 365 Azure Applications for a given domain.
[**DeleteApplicationServersRegistration**](SourceAPI.md#DeleteApplicationServersRegistration) | **Delete** /data-protect/sources/application-servers/{id} | Delete an application server registration.
[**DeleteM365SelfServiceConfig**](SourceAPI.md#DeleteM365SelfServiceConfig) | **Delete** /data-protect/sources/microsoft365/self-service-config/{uuid} | Deletes the Self-Service configuration for a Microsoft365 source.
[**DeleteProtectionSourceRegistration**](SourceAPI.md#DeleteProtectionSourceRegistration) | **Delete** /data-protect/sources/registrations/{id} | Delete Protection Source Registration.
[**GenerateM365DeviceAccessToken**](SourceAPI.md#GenerateM365DeviceAccessToken) | **Post** /data-protect/sources/microsoft365/auth/token | Generate access token for Microsoft365 Device Authorization Grant flow.
[**GenerateM365DeviceCode**](SourceAPI.md#GenerateM365DeviceCode) | **Post** /data-protect/sources/microsoft365/auth/device-code | Generate device code for Microsoft365 Device Authorization Grant flow.
[**GetMicrosoft365SelfServiceConfig**](SourceAPI.md#GetMicrosoft365SelfServiceConfig) | **Get** /data-protect/sources/microsoft365/self-service-config | Get the list of Microsoft365 Self-Service configurations
[**GetProtectionSourceRegistration**](SourceAPI.md#GetProtectionSourceRegistration) | **Get** /data-protect/sources/registrations/{id} | Get a Protection Source registration.
[**GetProtectionSources**](SourceAPI.md#GetProtectionSources) | **Get** /data-protect/sources | Get a List of Protection Sources.
[**GetSourceAttributeFilters**](SourceAPI.md#GetSourceAttributeFilters) | **Get** /data-protect/sources/filters | List attribute filters for a source.
[**GetSourceRegistrations**](SourceAPI.md#GetSourceRegistrations) | **Get** /data-protect/sources/registrations | Get the list of Protection Source registrations.
[**GetVdcDetails**](SourceAPI.md#GetVdcDetails) | **Get** /data-protect/sources/virtual-datacenter/{id} | Get VDC Details.
[**PatchProtectionSourceRegistration**](SourceAPI.md#PatchProtectionSourceRegistration) | **Patch** /data-protect/sources/registrations/{id} | Perform Partial Update on Protection Source registration. Currently this API is supported only for Cassandra
[**ProtectionSourceById**](SourceAPI.md#ProtectionSourceById) | **Get** /data-protect/sources/{id} | Get a Protection Sources.
[**RefreshProtectionSourceById**](SourceAPI.md#RefreshProtectionSourceById) | **Post** /data-protect/sources/{id}/refresh | Refresh a Protection Source.
[**RegisterProtectionSource**](SourceAPI.md#RegisterProtectionSource) | **Post** /data-protect/sources/registrations | Register a Protection Source.
[**TestConnectionProtectionSource**](SourceAPI.md#TestConnectionProtectionSource) | **Post** /data-protect/sources/test-connection | Test connection to a source.
[**UpdateM365SelfServiceConfig**](SourceAPI.md#UpdateM365SelfServiceConfig) | **Put** /data-protect/sources/microsoft365/self-service-config/{uuid} | Create or Update the Self-Service configuration for a Microsoft365 source.
[**UpdateProtectionSourceRegistration**](SourceAPI.md#UpdateProtectionSourceRegistration) | **Put** /data-protect/sources/registrations/{id} | Update Protection Source registration.



## CreateAzureApplications

> CreateAzureApplicationResponseParams CreateAzureApplications(ctx).Body(body).Execute()

Create Microsoft 365 Azure Applications for a given domain.



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
	body := *openapiclient.NewCreateAzureApplicationRequestParams("AccessToken_example", int32(123)) // CreateAzureApplicationRequestParams | Specifies the parameters to create Azure applications within a given Microsoft365 source.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceAPI.CreateAzureApplications(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceAPI.CreateAzureApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAzureApplications`: CreateAzureApplicationResponseParams
	fmt.Fprintf(os.Stdout, "Response from `SourceAPI.CreateAzureApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAzureApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateAzureApplicationRequestParams**](CreateAzureApplicationRequestParams.md) | Specifies the parameters to create Azure applications within a given Microsoft365 source. | 

### Return type

[**CreateAzureApplicationResponseParams**](CreateAzureApplicationResponseParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateAzureApplications

> CreateAzureApplicationResponseParams CreateOrUpdateAzureApplications(ctx).Body(body).Execute()

Create/Update Microsoft 365 Azure Applications for a given domain.



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
	body := *openapiclient.NewCreateAzureApplicationRequestParams("AccessToken_example", int32(123)) // CreateAzureApplicationRequestParams | Specifies the parameters to create/update Azure applications within a given Microsoft365 source.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceAPI.CreateOrUpdateAzureApplications(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceAPI.CreateOrUpdateAzureApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrUpdateAzureApplications`: CreateAzureApplicationResponseParams
	fmt.Fprintf(os.Stdout, "Response from `SourceAPI.CreateOrUpdateAzureApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateAzureApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateAzureApplicationRequestParams**](CreateAzureApplicationRequestParams.md) | Specifies the parameters to create/update Azure applications within a given Microsoft365 source. | 

### Return type

[**CreateAzureApplicationResponseParams**](CreateAzureApplicationResponseParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationServersRegistration

> DeleteApplicationServersRegistration(ctx, id).Body(body).Execute()

Delete an application server registration.



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
	id := int64(789) // int64 | Specifies the id of the Application Server.
	body := *openapiclient.NewUnRegisterApplicationServersParams() // UnRegisterApplicationServersParams | Specifies the request to unregister a an application server.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SourceAPI.DeleteApplicationServersRegistration(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceAPI.DeleteApplicationServersRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the id of the Application Server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationServersRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UnRegisterApplicationServersParams**](UnRegisterApplicationServersParams.md) | Specifies the request to unregister a an application server. | 

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


## DeleteM365SelfServiceConfig

> DeleteM365SelfServiceConfig(ctx, uuid).Execute()

Deletes the Self-Service configuration for a Microsoft365 source.



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
	uuid := "uuid_example" // string | Specifies the UUID of the Microsoft365 Source.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SourceAPI.DeleteM365SelfServiceConfig(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceAPI.DeleteM365SelfServiceConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Specifies the UUID of the Microsoft365 Source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteM365SelfServiceConfigRequest struct via the builder pattern


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


## DeleteProtectionSourceRegistration

> DeleteProtectionSourceRegistration(ctx, id).Execute()

Delete Protection Source Registration.



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
	id := int64(789) // int64 | Specifies the ID of the Protection Source Registration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SourceAPI.DeleteProtectionSourceRegistration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceAPI.DeleteProtectionSourceRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the ID of the Protection Source Registration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProtectionSourceRegistrationRequest struct via the builder pattern


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


## GenerateM365DeviceAccessToken

> GenerateM365DeviceAccessTokenResponseParams GenerateM365DeviceAccessToken(ctx).Body(body).Execute()

Generate access token for Microsoft365 Device Authorization Grant flow.



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
	body := *openapiclient.NewGenerateM365DeviceAccessTokenRequestParams() // GenerateM365DeviceAccessTokenRequestParams | Specifies the parameters to validate and generate access token for authorizing the client within Microsoft365.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceAPI.GenerateM365DeviceAccessToken(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceAPI.GenerateM365DeviceAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateM365DeviceAccessToken`: GenerateM365DeviceAccessTokenResponseParams
	fmt.Fprintf(os.Stdout, "Response from `SourceAPI.GenerateM365DeviceAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateM365DeviceAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GenerateM365DeviceAccessTokenRequestParams**](GenerateM365DeviceAccessTokenRequestParams.md) | Specifies the parameters to validate and generate access token for authorizing the client within Microsoft365. | 

### Return type

[**GenerateM365DeviceAccessTokenResponseParams**](GenerateM365DeviceAccessTokenResponseParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateM365DeviceCode

> GenerateM365DeviceCodeResponseParams GenerateM365DeviceCode(ctx).Body(body).Execute()

Generate device code for Microsoft365 Device Authorization Grant flow.



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
	body := *openapiclient.NewGenerateM365DeviceCodeRequestParams() // GenerateM365DeviceCodeRequestParams | Specifies the parameters to generate the user and device code to initiate authentication with Microsoft365.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceAPI.GenerateM365DeviceCode(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceAPI.GenerateM365DeviceCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateM365DeviceCode`: GenerateM365DeviceCodeResponseParams
	fmt.Fprintf(os.Stdout, "Response from `SourceAPI.GenerateM365DeviceCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateM365DeviceCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GenerateM365DeviceCodeRequestParams**](GenerateM365DeviceCodeRequestParams.md) | Specifies the parameters to generate the user and device code to initiate authentication with Microsoft365. | 

### Return type

[**GenerateM365DeviceCodeResponseParams**](GenerateM365DeviceCodeResponseParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMicrosoft365SelfServiceConfig

> []CreateM365SelfServiceConfigRequestParams GetMicrosoft365SelfServiceConfig(ctx).TenantId(tenantId).WorkloadType(workloadType).Execute()

Get the list of Microsoft365 Self-Service configurations



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
	tenantId := "tenantId_example" // string | Specifies the Cohesity Tenant ID for the source owner.
	workloadType := "workloadType_example" // string | Specifies the workload type as filter for fetching Self-Service configuration types. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceAPI.GetMicrosoft365SelfServiceConfig(context.Background()).TenantId(tenantId).WorkloadType(workloadType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceAPI.GetMicrosoft365SelfServiceConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMicrosoft365SelfServiceConfig`: []CreateM365SelfServiceConfigRequestParams
	fmt.Fprintf(os.Stdout, "Response from `SourceAPI.GetMicrosoft365SelfServiceConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMicrosoft365SelfServiceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Specifies the Cohesity Tenant ID for the source owner. | 
 **workloadType** | **string** | Specifies the workload type as filter for fetching Self-Service configuration types. | 

### Return type

[**[]CreateM365SelfServiceConfigRequestParams**](CreateM365SelfServiceConfigRequestParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtectionSourceRegistration

> SourceRegistration GetProtectionSourceRegistration(ctx, id).RequestInitiatorType(requestInitiatorType).Execute()

Get a Protection Source registration.



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
	id := int64(789) // int64 | Specifies the id of the Protection Source registration.
	requestInitiatorType := "requestInitiatorType_example" // string | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceAPI.GetProtectionSourceRegistration(context.Background(), id).RequestInitiatorType(requestInitiatorType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceAPI.GetProtectionSourceRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProtectionSourceRegistration`: SourceRegistration
	fmt.Fprintf(os.Stdout, "Response from `SourceAPI.GetProtectionSourceRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the id of the Protection Source registration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionSourceRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestInitiatorType** | **string** | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | 

### Return type

[**SourceRegistration**](SourceRegistration.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtectionSources

> Sources GetProtectionSources(ctx).RequestInitiatorType(requestInitiatorType).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeSourceCredentials(includeSourceCredentials).EncryptionKey(encryptionKey).Execute()

Get a List of Protection Sources.



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
	requestInitiatorType := "requestInitiatorType_example" // string | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. (optional)
	tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which Sources are to be returned. (optional)
	includeTenants := true // bool | If true, the response will include Sources which belong belong to all tenants which the current user has permission to see. If false, then only Sources for the current user will be returned. (optional)
	includeSourceCredentials := true // bool | If true, the encrypted crednetial for the registered sources will be included. Credential is first encrypted with internal key and then reencrypted with user supplied encryption key. (optional)
	encryptionKey := "encryptionKey_example" // string | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceAPI.GetProtectionSources(context.Background()).RequestInitiatorType(requestInitiatorType).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeSourceCredentials(includeSourceCredentials).EncryptionKey(encryptionKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceAPI.GetProtectionSources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProtectionSources`: Sources
	fmt.Fprintf(os.Stdout, "Response from `SourceAPI.GetProtectionSources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestInitiatorType** | **string** | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which Sources are to be returned. | 
 **includeTenants** | **bool** | If true, the response will include Sources which belong belong to all tenants which the current user has permission to see. If false, then only Sources for the current user will be returned. | 
 **includeSourceCredentials** | **bool** | If true, the encrypted crednetial for the registered sources will be included. Credential is first encrypted with internal key and then reencrypted with user supplied encryption key. | 
 **encryptionKey** | **string** | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | 

### Return type

[**Sources**](Sources.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceAttributeFilters

> SourceAttributeFiltersResponseParams GetSourceAttributeFilters(ctx).SourceUuid(sourceUuid).Environment(environment).Execute()

List attribute filters for a source.



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
	sourceUuid := "sourceUuid_example" // string | Specifies the source UUID of the parent entity.
	environment := "environment_example" // string | Specifies the environment type of the Protection Source. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceAPI.GetSourceAttributeFilters(context.Background()).SourceUuid(sourceUuid).Environment(environment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceAPI.GetSourceAttributeFilters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSourceAttributeFilters`: SourceAttributeFiltersResponseParams
	fmt.Fprintf(os.Stdout, "Response from `SourceAPI.GetSourceAttributeFilters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceAttributeFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceUuid** | **string** | Specifies the source UUID of the parent entity. | 
 **environment** | **string** | Specifies the environment type of the Protection Source. | 

### Return type

[**SourceAttributeFiltersResponseParams**](SourceAttributeFiltersResponseParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceRegistrations

> SourceRegistrations GetSourceRegistrations(ctx).Ids(ids).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeSourceCredentials(includeSourceCredentials).EncryptionKey(encryptionKey).UseCachedData(useCachedData).IncludeExternalMetadata(includeExternalMetadata).IgnoreTenantMigrationInProgressCheck(ignoreTenantMigrationInProgressCheck).Execute()

Get the list of Protection Source registrations.



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
	ids := []int32{int32(123)} // []int32 | Ids specifies the list of source registration ids to return. If left empty, every source registration will be returned by default. (optional)
	tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
	includeTenants := true // bool | If true, the response will include Registrations which were created by all tenants which the current user has permission to see. If false, then only Registrations created by the current user will be returned. (optional)
	includeSourceCredentials := true // bool | If true, the encrypted crednetial for the registered sources will be included. Credential is first encrypted with internal key and then reencrypted with user supplied encryption key. (optional)
	encryptionKey := "encryptionKey_example" // string | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. (optional)
	useCachedData := true // bool | Specifies whether we can serve the GET request from the read replica cache. There is a lag of 15 seconds between the read replica and primary data source. (optional)
	includeExternalMetadata := true // bool | If true, the external entity metadata like maintenance mode config for the registered sources will be included. (optional)
	ignoreTenantMigrationInProgressCheck := true // bool | If true, tenant migration check will be ignored (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceAPI.GetSourceRegistrations(context.Background()).Ids(ids).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeSourceCredentials(includeSourceCredentials).EncryptionKey(encryptionKey).UseCachedData(useCachedData).IncludeExternalMetadata(includeExternalMetadata).IgnoreTenantMigrationInProgressCheck(ignoreTenantMigrationInProgressCheck).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceAPI.GetSourceRegistrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSourceRegistrations`: SourceRegistrations
	fmt.Fprintf(os.Stdout, "Response from `SourceAPI.GetSourceRegistrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceRegistrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]int32** | Ids specifies the list of source registration ids to return. If left empty, every source registration will be returned by default. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **includeTenants** | **bool** | If true, the response will include Registrations which were created by all tenants which the current user has permission to see. If false, then only Registrations created by the current user will be returned. | 
 **includeSourceCredentials** | **bool** | If true, the encrypted crednetial for the registered sources will be included. Credential is first encrypted with internal key and then reencrypted with user supplied encryption key. | 
 **encryptionKey** | **string** | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | 
 **useCachedData** | **bool** | Specifies whether we can serve the GET request from the read replica cache. There is a lag of 15 seconds between the read replica and primary data source. | 
 **includeExternalMetadata** | **bool** | If true, the external entity metadata like maintenance mode config for the registered sources will be included. | 
 **ignoreTenantMigrationInProgressCheck** | **bool** | If true, tenant migration check will be ignored | 

### Return type

[**SourceRegistrations**](SourceRegistrations.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVdcDetails

> VdcObject GetVdcDetails(ctx, id).Execute()

Get VDC Details.



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
	id := int64(789) // int64 | Specifies the ID of the VMware virtual datacenter.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceAPI.GetVdcDetails(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceAPI.GetVdcDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVdcDetails`: VdcObject
	fmt.Fprintf(os.Stdout, "Response from `SourceAPI.GetVdcDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the ID of the VMware virtual datacenter. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVdcDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VdcObject**](VdcObject.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchProtectionSourceRegistration

> SourceRegistration PatchProtectionSourceRegistration(ctx, id).Body(body).Execute()

Perform Partial Update on Protection Source registration. Currently this API is supported only for Cassandra



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
	id := int64(789) // int64 | Specifies the id of the Protection Source registration.
	body := *openapiclient.NewSourceRegistrationPatchRequestParams("Environment_example") // SourceRegistrationPatchRequestParams | Specifies the parameters to partially update the registration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceAPI.PatchProtectionSourceRegistration(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceAPI.PatchProtectionSourceRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProtectionSourceRegistration`: SourceRegistration
	fmt.Fprintf(os.Stdout, "Response from `SourceAPI.PatchProtectionSourceRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the id of the Protection Source registration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProtectionSourceRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SourceRegistrationPatchRequestParams**](SourceRegistrationPatchRequestParams.md) | Specifies the parameters to partially update the registration. | 

### Return type

[**SourceRegistration**](SourceRegistration.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProtectionSourceById

> Source ProtectionSourceById(ctx, id).Execute()

Get a Protection Sources.



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
	id := int64(789) // int64 | Specifies the id of the Protection Source.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceAPI.ProtectionSourceById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceAPI.ProtectionSourceById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProtectionSourceById`: Source
	fmt.Fprintf(os.Stdout, "Response from `SourceAPI.ProtectionSourceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the id of the Protection Source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProtectionSourceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Source**](Source.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshProtectionSourceById

> RefreshProtectionSourceById(ctx, id).Execute()

Refresh a Protection Source.



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
	id := int64(789) // int64 | Specifies the id of the Protection Source.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SourceAPI.RefreshProtectionSourceById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceAPI.RefreshProtectionSourceById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the id of the Protection Source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshProtectionSourceByIdRequest struct via the builder pattern


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


## RegisterProtectionSource

> SourceRegistration RegisterProtectionSource(ctx).Body(body).Execute()

Register a Protection Source.



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
	body := *openapiclient.NewSourceRegistrationRequestParams("Environment_example") // SourceRegistrationRequestParams | Specifies the parameters to register a Protection Source.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceAPI.RegisterProtectionSource(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceAPI.RegisterProtectionSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterProtectionSource`: SourceRegistration
	fmt.Fprintf(os.Stdout, "Response from `SourceAPI.RegisterProtectionSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterProtectionSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SourceRegistrationRequestParams**](SourceRegistrationRequestParams.md) | Specifies the parameters to register a Protection Source. | 

### Return type

[**SourceRegistration**](SourceRegistration.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestConnectionProtectionSource

> SourceConnectionResponseParams TestConnectionProtectionSource(ctx).Body(body).Execute()

Test connection to a source.



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
	body := *openapiclient.NewSourceConnectionRequestParams("Environment_example") // SourceConnectionRequestParams | Specifies the parameters to test connectivity with a source.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceAPI.TestConnectionProtectionSource(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceAPI.TestConnectionProtectionSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestConnectionProtectionSource`: SourceConnectionResponseParams
	fmt.Fprintf(os.Stdout, "Response from `SourceAPI.TestConnectionProtectionSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestConnectionProtectionSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SourceConnectionRequestParams**](SourceConnectionRequestParams.md) | Specifies the parameters to test connectivity with a source. | 

### Return type

[**SourceConnectionResponseParams**](SourceConnectionResponseParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateM365SelfServiceConfig

> CreateM365SelfServiceConfigRequestParams UpdateM365SelfServiceConfig(ctx, uuid).Body(body).Execute()

Create or Update the Self-Service configuration for a Microsoft365 source.



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
	uuid := "uuid_example" // string | Specifies the UUID of the Microsoft365 Source.
	body := *openapiclient.NewCreateM365SelfServiceConfigRequestParams("TenantId_example", "Uuid_example") // CreateM365SelfServiceConfigRequestParams | Specifies the parameters to enable Self-Service for a Microsoft365 source. This configuration will apply to all regions incase the same source is registered across regions.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceAPI.UpdateM365SelfServiceConfig(context.Background(), uuid).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceAPI.UpdateM365SelfServiceConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateM365SelfServiceConfig`: CreateM365SelfServiceConfigRequestParams
	fmt.Fprintf(os.Stdout, "Response from `SourceAPI.UpdateM365SelfServiceConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Specifies the UUID of the Microsoft365 Source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateM365SelfServiceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateM365SelfServiceConfigRequestParams**](CreateM365SelfServiceConfigRequestParams.md) | Specifies the parameters to enable Self-Service for a Microsoft365 source. This configuration will apply to all regions incase the same source is registered across regions. | 

### Return type

[**CreateM365SelfServiceConfigRequestParams**](CreateM365SelfServiceConfigRequestParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProtectionSourceRegistration

> SourceRegistration UpdateProtectionSourceRegistration(ctx, id).Body(body).Execute()

Update Protection Source registration.



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
	id := int64(789) // int64 | Specifies the id of the Protection Source registration.
	body := *openapiclient.NewSourceRegistrationUpdateRequestParams("Environment_example") // SourceRegistrationUpdateRequestParams | Specifies the parameters to update the registration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceAPI.UpdateProtectionSourceRegistration(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceAPI.UpdateProtectionSourceRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProtectionSourceRegistration`: SourceRegistration
	fmt.Fprintf(os.Stdout, "Response from `SourceAPI.UpdateProtectionSourceRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the id of the Protection Source registration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProtectionSourceRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SourceRegistrationUpdateRequestParams**](SourceRegistrationUpdateRequestParams.md) | Specifies the parameters to update the registration. | 

### Return type

[**SourceRegistration**](SourceRegistration.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

