# \TenantAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignPropertiesToTenant**](TenantAPI.md#AssignPropertiesToTenant) | **Put** /tenants/{id}/assignments | Update assginment of properties for a tenant.
[**CreateTenant**](TenantAPI.md#CreateTenant) | **Post** /tenants | Create a new Tenant.
[**DeleteTenant**](TenantAPI.md#DeleteTenant) | **Delete** /tenants/{id} | Delete Tenant with given ID.
[**GetAssignedPropertiesForTenant**](TenantAPI.md#GetAssignedPropertiesForTenant) | **Get** /tenants/{id}/assignments | Get tenant assignments.
[**GetOnPremTenantConfig**](TenantAPI.md#GetOnPremTenantConfig) | **Get** /clusters/tenant-config | Get Tenants Config.
[**GetTenantSwift**](TenantAPI.md#GetTenantSwift) | **Get** /tenants/swift | Get a Swift configuration.
[**ListTenants**](TenantAPI.md#ListTenants) | **Get** /tenants | Get a list of Tenants.
[**PerformTenantAction**](TenantAPI.md#PerformTenantAction) | **Post** /tenants/{id}/actions | Perform actions on a Tenant.
[**RegisterSwift**](TenantAPI.md#RegisterSwift) | **Post** /tenants/swift/register | Register Swift service on a Keystone server.
[**UnregisterSwift**](TenantAPI.md#UnregisterSwift) | **Post** /tenants/swift/unregister | Unregister Swift service from a Keystone server.
[**UpdateOnPremTenantConfig**](TenantAPI.md#UpdateOnPremTenantConfig) | **Post** /clusters/tenant-config | Update Tenants Config.
[**UpdateTenant**](TenantAPI.md#UpdateTenant) | **Put** /tenants/{id} | Update Tenant.
[**UpdateTenantSwift**](TenantAPI.md#UpdateTenantSwift) | **Put** /tenants/swift | Update a Swift configuration.



## AssignPropertiesToTenant

> TenantAssignments AssignPropertiesToTenant(ctx, id).Body(body).Execute()

Update assginment of properties for a tenant.



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
	id := "id_example" // string | The Tenant id.
	body := *openapiclient.NewTenantAssignmentsParams() // TenantAssignmentsParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.AssignPropertiesToTenant(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.AssignPropertiesToTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignPropertiesToTenant`: TenantAssignments
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.AssignPropertiesToTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Tenant id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignPropertiesToTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**TenantAssignmentsParams**](TenantAssignmentsParams.md) |  | 

### Return type

[**TenantAssignments**](TenantAssignments.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTenant

> TenantInfo CreateTenant(ctx).Body(body).Execute()

Create a new Tenant.

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
	body := *openapiclient.NewCreateTenantRequest("Name_example", "TenantIdSuffix_example") // CreateTenantRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.CreateTenant(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.CreateTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTenant`: TenantInfo
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.CreateTenant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateTenantRequest**](CreateTenantRequest.md) |  | 

### Return type

[**TenantInfo**](TenantInfo.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTenant

> DeleteTenant(ctx, id).Execute()

Delete Tenant with given ID.

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
	id := "id_example" // string | The Tenant id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TenantAPI.DeleteTenant(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.DeleteTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Tenant id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantRequest struct via the builder pattern


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


## GetAssignedPropertiesForTenant

> TenantAssignmentProperties GetAssignedPropertiesForTenant(ctx, id).Execute()

Get tenant assignments.



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
	id := "id_example" // string | The Tenant id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.GetAssignedPropertiesForTenant(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.GetAssignedPropertiesForTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssignedPropertiesForTenant`: TenantAssignmentProperties
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.GetAssignedPropertiesForTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Tenant id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssignedPropertiesForTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TenantAssignmentProperties**](TenantAssignmentProperties.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnPremTenantConfig

> OnPremTenantConfig GetOnPremTenantConfig(ctx).Execute()

Get Tenants Config.



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
	resp, r, err := apiClient.TenantAPI.GetOnPremTenantConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.GetOnPremTenantConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnPremTenantConfig`: OnPremTenantConfig
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.GetOnPremTenantConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOnPremTenantConfigRequest struct via the builder pattern


### Return type

[**OnPremTenantConfig**](OnPremTenantConfig.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantSwift

> SwiftParams GetTenantSwift(ctx).TenantId(tenantId).Execute()

Get a Swift configuration.



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
	tenantId := "tenantId_example" // string | Specifies the tenant Id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.GetTenantSwift(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.GetTenantSwift``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantSwift`: SwiftParams
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.GetTenantSwift`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantSwiftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Specifies the tenant Id. | 

### Return type

[**SwiftParams**](SwiftParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTenants

> []TenantInfo ListTenants(ctx).Ids(ids).Statuses(statuses).LivenessModes(livenessModes).OwnershipModes(ownershipModes).Execute()

Get a list of Tenants.

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
	ids := []*string{"Inner_example"} // []*string | List of tenantIds to filter. (optional)
	statuses := []*string{"Statuses_example"} // []*string | Filter by current status of tenant. If left blank, only active and inactive tenants are returned. (optional)
	livenessModes := []*string{"LivenessModes_example"} // []*string | Filter by liveness modes of the tenant. This filter only applies is tenant metadata is added for external vendor such as 'IBM'. In all other cases, the values provided for this filter will be ignored. (optional)
	ownershipModes := []*string{"OwnershipModes_example"} // []*string | Filter by ownership modes of the tenant. This filter only applies is tenant metadata is added for external vendor such as 'IBM'. In all other cases, the values provided for this filter will be ignored. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.ListTenants(context.Background()).Ids(ids).Statuses(statuses).LivenessModes(livenessModes).OwnershipModes(ownershipModes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.ListTenants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTenants`: []TenantInfo
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.ListTenants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTenantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** | List of tenantIds to filter. | 
 **statuses** | **[]string** | Filter by current status of tenant. If left blank, only active and inactive tenants are returned. | 
 **livenessModes** | **[]string** | Filter by liveness modes of the tenant. This filter only applies is tenant metadata is added for external vendor such as &#39;IBM&#39;. In all other cases, the values provided for this filter will be ignored. | 
 **ownershipModes** | **[]string** | Filter by ownership modes of the tenant. This filter only applies is tenant metadata is added for external vendor such as &#39;IBM&#39;. In all other cases, the values provided for this filter will be ignored. | 

### Return type

[**[]TenantInfo**](TenantInfo.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PerformTenantAction

> TenantInfo PerformTenantAction(ctx, id).Body(body).Execute()

Perform actions on a Tenant.



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
	id := "id_example" // string | The Tenant id.
	body := *openapiclient.NewTenantActionBody("Action_example") // TenantActionBody | Specifies the parameters to perform an action on a Tenant.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.PerformTenantAction(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.PerformTenantAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PerformTenantAction`: TenantInfo
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.PerformTenantAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Tenant id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPerformTenantActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**TenantActionBody**](TenantActionBody.md) | Specifies the parameters to perform an action on a Tenant. | 

### Return type

[**TenantInfo**](TenantInfo.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterSwift

> RegisterSwift(ctx).Body(body).Execute()

Register Swift service on a Keystone server.



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
	body := *openapiclient.NewRegisterSwiftParams("TenantId_example") // RegisterSwiftParams | Specifies the parameters to register a Swift service on Keystone server.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TenantAPI.RegisterSwift(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.RegisterSwift``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterSwiftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RegisterSwiftParams**](RegisterSwiftParams.md) | Specifies the parameters to register a Swift service on Keystone server. | 

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


## UnregisterSwift

> UnregisterSwift(ctx).Body(body).Execute()

Unregister Swift service from a Keystone server.



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
	body := *openapiclient.NewUnregisterSwiftParams("TenantId_example") // UnregisterSwiftParams | Specifies the parameters to unregister a Swift service from Keystone server.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TenantAPI.UnregisterSwift(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.UnregisterSwift``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnregisterSwiftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UnregisterSwiftParams**](UnregisterSwiftParams.md) | Specifies the parameters to unregister a Swift service from Keystone server. | 

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


## UpdateOnPremTenantConfig

> OnPremTenantConfig UpdateOnPremTenantConfig(ctx).Body(body).Execute()

Update Tenants Config.



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
	body := *openapiclient.NewOnPremTenantConfig(false, false) // OnPremTenantConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.UpdateOnPremTenantConfig(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.UpdateOnPremTenantConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOnPremTenantConfig`: OnPremTenantConfig
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.UpdateOnPremTenantConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOnPremTenantConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OnPremTenantConfig**](OnPremTenantConfig.md) |  | 

### Return type

[**OnPremTenantConfig**](OnPremTenantConfig.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTenant

> TenantInfo UpdateTenant(ctx, id).Body(body).Execute()

Update Tenant.



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
	id := "id_example" // string | 
	body := *openapiclient.NewUpdateTenantBody() // UpdateTenantBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.UpdateTenant(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.UpdateTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTenant`: TenantInfo
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.UpdateTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateTenantBody**](UpdateTenantBody.md) |  | 

### Return type

[**TenantInfo**](TenantInfo.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTenantSwift

> SwiftParams UpdateTenantSwift(ctx).Body(body).Execute()

Update a Swift configuration.



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
	body := *openapiclient.NewSwiftParams("TenantId_example") // SwiftParams | Specifies the parameters to update a Swift configuration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.UpdateTenantSwift(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.UpdateTenantSwift``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTenantSwift`: SwiftParams
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.UpdateTenantSwift`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantSwiftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SwiftParams**](SwiftParams.md) | Specifies the parameters to update a Swift configuration. | 

### Return type

[**SwiftParams**](SwiftParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

