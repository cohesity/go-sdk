# \Tenant

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignPropertiesToTenant**](Tenant.md#AssignPropertiesToTenant) | **Put** /tenants/{id}/assignments | Update assginment of properties for a tenant.
[**CreateTenant**](Tenant.md#CreateTenant) | **Post** /tenants | Create a new Tenant.
[**DeleteTenant**](Tenant.md#DeleteTenant) | **Delete** /tenants/{id} | Delete Tenant with given ID.
[**GetAssignedPropertiesForTenant**](Tenant.md#GetAssignedPropertiesForTenant) | **Get** /tenants/{id}/assignments | Get tenant assignments.
[**GetOnPremTenantConfig**](Tenant.md#GetOnPremTenantConfig) | **Get** /clusters/tenant-config | Get Tenants Config.
[**GetTenantByID**](Tenant.md#GetTenantByID) | **Get** /tenants/{id} | Get Tenant by ID.
[**GetTenantSwift**](Tenant.md#GetTenantSwift) | **Get** /tenants/swift | Get a Swift configuration.
[**ListTenants**](Tenant.md#ListTenants) | **Get** /tenants | Get a list of Tenants.
[**PerformTenantAction**](Tenant.md#PerformTenantAction) | **Post** /tenants/{id}/actions | Perform actions on a Tenant.
[**RegisterSwift**](Tenant.md#RegisterSwift) | **Post** /tenants/swift/register | Register Swift service on a Keystone server.
[**UnregisterSwift**](Tenant.md#UnregisterSwift) | **Post** /tenants/swift/unregister | Unregister Swift service from a Keystone server.
[**UpdateOnPremTenantConfig**](Tenant.md#UpdateOnPremTenantConfig) | **Post** /clusters/tenant-config | Update Tenants Config.
[**UpdateTenant**](Tenant.md#UpdateTenant) | **Put** /tenants/{id} | Update Tenant.
[**UpdateTenantSwift**](Tenant.md#UpdateTenantSwift) | **Put** /tenants/swift | Update a Swift configuration.



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
    onprem "onprem"
)

func main() {
    id := "id_example" // string | The Tenant id.
    body := *openapiclient.NewTenantAssignmentsParams() // TenantAssignmentsParams | 

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiAssignPropertiesToTenantRequest{
        Id: &id
        Body: &body
    }

    resp, r, err := api_client.Tenant.AssignPropertiesToTenant(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tenant.AssignPropertiesToTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignPropertiesToTenant`: TenantAssignments
    fmt.Fprintf(os.Stdout, "Response from `Tenant.AssignPropertiesToTenant`: %v\n", resp)
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
    onprem "onprem"
)

func main() {
    body := *openapiclient.NewCreateTenantParams("Name_example", "TenantIdSuffix_example") // CreateTenantParams | 

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiCreateTenantRequest{
        Body: &body
    }

    resp, r, err := api_client.Tenant.CreateTenant(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tenant.CreateTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTenant`: TenantInfo
    fmt.Fprintf(os.Stdout, "Response from `Tenant.CreateTenant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateTenantParams**](CreateTenantParams.md) |  | 

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
    onprem "onprem"
)

func main() {
    id := "id_example" // string | The Tenant id.

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiDeleteTenantRequest{
        Id: &id
    }

    resp, r, err := api_client.Tenant.DeleteTenant(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tenant.DeleteTenant``: %v\n", err)
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
    onprem "onprem"
)

func main() {
    id := "id_example" // string | The Tenant id.

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiGetAssignedPropertiesForTenantRequest{
        Id: &id
    }

    resp, r, err := api_client.Tenant.GetAssignedPropertiesForTenant(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tenant.GetAssignedPropertiesForTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssignedPropertiesForTenant`: TenantAssignmentProperties
    fmt.Fprintf(os.Stdout, "Response from `Tenant.GetAssignedPropertiesForTenant`: %v\n", resp)
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
    onprem "onprem"
)

func main() {

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiGetOnPremTenantConfigRequest{
    }

    resp, r, err := api_client.Tenant.GetOnPremTenantConfig(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tenant.GetOnPremTenantConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOnPremTenantConfig`: OnPremTenantConfig
    fmt.Fprintf(os.Stdout, "Response from `Tenant.GetOnPremTenantConfig`: %v\n", resp)
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


## GetTenantByID

> TenantInfo GetTenantByID(ctx, id).Execute()

Get Tenant by ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    id := "id_example" // string | The Tenant id.

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiGetTenantByIDRequest{
        Id: &id
    }

    resp, r, err := api_client.Tenant.GetTenantByID(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tenant.GetTenantByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenantByID`: TenantInfo
    fmt.Fprintf(os.Stdout, "Response from `Tenant.GetTenantByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Tenant id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TenantInfo**](TenantInfo.md)

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
    onprem "onprem"
)

func main() {
    tenantId := "tenantId_example" // string | Specifies the tenant Id. (optional)

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiGetTenantSwiftRequest{
        TenantId: &tenantId
    }

    resp, r, err := api_client.Tenant.GetTenantSwift(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tenant.GetTenantSwift``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenantSwift`: SwiftParams
    fmt.Fprintf(os.Stdout, "Response from `Tenant.GetTenantSwift`: %v\n", resp)
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

> TenantsInfo ListTenants(ctx).Ids(ids).Statuses(statuses).Execute()

Get a list of Tenants.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    ids := []string{"Inner_example"} // []string | List of tenantIds to filter. (optional)
    statuses := []string{"Statuses_example"} // []string | Filter by current status of tenant. If left blank, only active and inactive tenants are returned. (optional)

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiListTenantsRequest{
        Ids: &ids
        Statuses: &statuses
    }

    resp, r, err := api_client.Tenant.ListTenants(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tenant.ListTenants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTenants`: TenantsInfo
    fmt.Fprintf(os.Stdout, "Response from `Tenant.ListTenants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTenantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** | List of tenantIds to filter. | 
 **statuses** | **[]string** | Filter by current status of tenant. If left blank, only active and inactive tenants are returned. | 

### Return type

[**TenantsInfo**](TenantsInfo.md)

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
    onprem "onprem"
)

func main() {
    id := "id_example" // string | The Tenant id.
    body := *openapiclient.NewTenantActionBody("Action_example") // TenantActionBody | Specifies the parameters to perform an action on a Tenant.

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiPerformTenantActionRequest{
        Id: &id
        Body: &body
    }

    resp, r, err := api_client.Tenant.PerformTenantAction(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tenant.PerformTenantAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PerformTenantAction`: TenantInfo
    fmt.Fprintf(os.Stdout, "Response from `Tenant.PerformTenantAction`: %v\n", resp)
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
    onprem "onprem"
)

func main() {
    body := *openapiclient.NewRegisterSwiftParams("TenantId_example") // RegisterSwiftParams | Specifies the parameters to register a Swift service on Keystone server.

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiRegisterSwiftRequest{
        Body: &body
    }

    resp, r, err := api_client.Tenant.RegisterSwift(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tenant.RegisterSwift``: %v\n", err)
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
    onprem "onprem"
)

func main() {
    body := *openapiclient.NewUnregisterSwiftParams("TenantId_example") // UnregisterSwiftParams | Specifies the parameters to unregister a Swift service from Keystone server.

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiUnregisterSwiftRequest{
        Body: &body
    }

    resp, r, err := api_client.Tenant.UnregisterSwift(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tenant.UnregisterSwift``: %v\n", err)
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
    onprem "onprem"
)

func main() {
    body := *openapiclient.NewOnPremTenantConfig(false, false) // OnPremTenantConfig | 

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiUpdateOnPremTenantConfigRequest{
        Body: &body
    }

    resp, r, err := api_client.Tenant.UpdateOnPremTenantConfig(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tenant.UpdateOnPremTenantConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOnPremTenantConfig`: OnPremTenantConfig
    fmt.Fprintf(os.Stdout, "Response from `Tenant.UpdateOnPremTenantConfig`: %v\n", resp)
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
    onprem "onprem"
)

func main() {
    id := "id_example" // string | 
    body := *openapiclient.NewUpdateTenantBody() // UpdateTenantBody | 

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiUpdateTenantRequest{
        Id: &id
        Body: &body
    }

    resp, r, err := api_client.Tenant.UpdateTenant(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tenant.UpdateTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTenant`: TenantInfo
    fmt.Fprintf(os.Stdout, "Response from `Tenant.UpdateTenant`: %v\n", resp)
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
    onprem "onprem"
)

func main() {
    body := *openapiclient.NewSwiftParams("TenantId_example") // SwiftParams | Specifies the parameters to update a Swift configuration.

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiUpdateTenantSwiftRequest{
        Body: &body
    }

    resp, r, err := api_client.Tenant.UpdateTenantSwift(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tenant.UpdateTenantSwift``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTenantSwift`: SwiftParams
    fmt.Fprintf(os.Stdout, "Response from `Tenant.UpdateTenantSwift`: %v\n", resp)
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

