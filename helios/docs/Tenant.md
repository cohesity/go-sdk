# \Tenant

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignClusterToTenant**](Tenant.md#AssignClusterToTenant) | **Post** /mcm/tenants/{id}/clusters | Assign a Cluster to a tenant.
[**AssignPropertiesToTenant**](Tenant.md#AssignPropertiesToTenant) | **Put** /tenants/{id}/assignments | Update assginment of properties for a tenant.
[**ConfirmTenant**](Tenant.md#ConfirmTenant) | **Post** /mcm/tenants/{id}/manage | Enable Helios Management for Tenant.
[**CreateHeliosTenant**](Tenant.md#CreateHeliosTenant) | **Post** /mcm/tenants | Create a new Tenant on Helios.
[**CreateTenant**](Tenant.md#CreateTenant) | **Post** /tenants | Create a new Tenant.
[**DeleteHeliosTenant**](Tenant.md#DeleteHeliosTenant) | **Delete** /mcm/tenants/{id} | Delete a Tenant on Helios.
[**DeleteTenant**](Tenant.md#DeleteTenant) | **Delete** /tenants/{id} | Delete Tenant with given ID.
[**GetAccountTenantConfig**](Tenant.md#GetAccountTenantConfig) | **Get** /mcm/accounts/tenant-config | Get Tenants Config.
[**GetAssignedPropertiesForTenant**](Tenant.md#GetAssignedPropertiesForTenant) | **Get** /tenants/{id}/assignments | Get tenant assignments.
[**GetClustersTenantConfig**](Tenant.md#GetClustersTenantConfig) | **Get** /mcm/clusters/tenant-config | Get Tenant&#39;s config for all clusters.
[**GetHeliosTenants**](Tenant.md#GetHeliosTenants) | **Get** /mcm/tenants | Get a list of tenants.
[**GetOnPremTenantConfig**](Tenant.md#GetOnPremTenantConfig) | **Get** /clusters/tenant-config | Get Tenants Config.
[**GetTenantByID**](Tenant.md#GetTenantByID) | **Get** /tenants/{id} | Get Tenant by ID.
[**GetTenantStats**](Tenant.md#GetTenantStats) | **Get** /mcm/tenants/stats | Get Tenant Statistics.
[**GetTenantSwift**](Tenant.md#GetTenantSwift) | **Get** /tenants/swift | Get a Swift configuration.
[**HeliosAssignPropertiesToTenant**](Tenant.md#HeliosAssignPropertiesToTenant) | **Put** /mcm/tenants/{id}/assignments | Assign properties to a tenant.
[**ListTenants**](Tenant.md#ListTenants) | **Get** /tenants | Get a list of Tenants.
[**PerformHeliosTenantAction**](Tenant.md#PerformHeliosTenantAction) | **Post** /mcm/tenants/{id}/actions | Perform actions on a Helios Tenant.
[**PerformTenantAction**](Tenant.md#PerformTenantAction) | **Post** /tenants/{id}/actions | Perform actions on a Tenant.
[**RegisterSwift**](Tenant.md#RegisterSwift) | **Post** /tenants/swift/register | Register Swift service on a Keystone server.
[**UnregisterSwift**](Tenant.md#UnregisterSwift) | **Post** /tenants/swift/unregister | Unregister Swift service from a Keystone server.
[**UpdateAccountTenantConfig**](Tenant.md#UpdateAccountTenantConfig) | **Patch** /mcm/accounts/tenant-config | Update Tenant&#39;s Config.
[**UpdateClustersTenantConfig**](Tenant.md#UpdateClustersTenantConfig) | **Post** /mcm/clusters/tenant-config | Update Tenant&#39;s config for clusters.
[**UpdateHeliosTenant**](Tenant.md#UpdateHeliosTenant) | **Patch** /mcm/tenants/{id} | Update Tenant properties on Helios.
[**UpdateOnPremTenantConfig**](Tenant.md#UpdateOnPremTenantConfig) | **Post** /clusters/tenant-config | Update Tenants Config.
[**UpdateTenant**](Tenant.md#UpdateTenant) | **Put** /tenants/{id} | Update Tenant.
[**UpdateTenantSwift**](Tenant.md#UpdateTenantSwift) | **Put** /tenants/swift | Update a Swift configuration.



## AssignClusterToTenant

> HeliosTenant AssignClusterToTenant(ctx, id).Body(body).RegionId(regionId).Execute()

Assign a Cluster to a tenant.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | 
    body := *openapiclient.NewAssignClusterToTenantParamsBody("ClusterIdentifier_example", *openapiclient.NewTenantNetwork(false)) // AssignClusterToTenantParamsBody | 
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiAssignClusterToTenantRequest{
        Id: &id
        Body: &body
        RegionId: &regionId
    }

    resp, r, err := api_client.Tenant.AssignClusterToTenant(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tenant.AssignClusterToTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignClusterToTenant`: HeliosTenant
    fmt.Fprintf(os.Stdout, "Response from `Tenant.AssignClusterToTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignClusterToTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AssignClusterToTenantParamsBody**](AssignClusterToTenantParamsBody.md) |  | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**HeliosTenant**](HeliosTenant.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssignPropertiesToTenant

> TenantAssignments AssignPropertiesToTenant(ctx, id).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Update assginment of properties for a tenant.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | The Tenant id.
    body := *openapiclient.NewTenantAssignmentsParams() // TenantAssignmentsParams | 
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiAssignPropertiesToTenantRequest{
        Id: &id
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
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
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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


## ConfirmTenant

> HeliosTenant ConfirmTenant(ctx, id).Body(body).RegionId(regionId).Execute()

Enable Helios Management for Tenant.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | 
    body := *openapiclient.NewConfirmTenantParamsBody("Name_example") // ConfirmTenantParamsBody | 
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiConfirmTenantRequest{
        Id: &id
        Body: &body
        RegionId: &regionId
    }

    resp, r, err := api_client.Tenant.ConfirmTenant(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tenant.ConfirmTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfirmTenant`: HeliosTenant
    fmt.Fprintf(os.Stdout, "Response from `Tenant.ConfirmTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ConfirmTenantParamsBody**](ConfirmTenantParamsBody.md) |  | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**HeliosTenant**](HeliosTenant.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateHeliosTenant

> HeliosTenant CreateHeliosTenant(ctx).Body(body).RegionId(regionId).Execute()

Create a new Tenant on Helios.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    body := *openapiclient.NewCreateTenantParams("Name_example", "TenantIdSuffix_example") // CreateTenantParams | 
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiCreateHeliosTenantRequest{
        Body: &body
        RegionId: &regionId
    }

    resp, r, err := api_client.Tenant.CreateHeliosTenant(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tenant.CreateHeliosTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateHeliosTenant`: HeliosTenant
    fmt.Fprintf(os.Stdout, "Response from `Tenant.CreateHeliosTenant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHeliosTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateTenantParams**](CreateTenantParams.md) |  | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**HeliosTenant**](HeliosTenant.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTenant

> TenantInfo CreateTenant(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Create a new Tenant.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    body := *openapiclient.NewCreateTenantParamsOthers("Name_example", "TenantIdSuffix_example") // CreateTenantParamsOthers | 
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiCreateTenantRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
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
 **body** | [**CreateTenantParamsOthers**](CreateTenantParamsOthers.md) |  | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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


## DeleteHeliosTenant

> DeleteHeliosTenant(ctx, id).RegionId(regionId).DeleteClusterTenants(deleteClusterTenants).Execute()

Delete a Tenant on Helios.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | 
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    deleteClusterTenants := true // bool | Wether or not to delete the tenants on the cluster, default is true. (optional) (default to true)

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

    request := helios.ApiDeleteHeliosTenantRequest{
        Id: &id
        RegionId: &regionId
        DeleteClusterTenants: &deleteClusterTenants
    }

    resp, r, err := api_client.Tenant.DeleteHeliosTenant(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tenant.DeleteHeliosTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHeliosTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **deleteClusterTenants** | **bool** | Wether or not to delete the tenants on the cluster, default is true. | [default to true]

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


## DeleteTenant

> DeleteTenant(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Delete Tenant with given ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | The Tenant id.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiDeleteTenantRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
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

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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


## GetAccountTenantConfig

> TenantConfig GetAccountTenantConfig(ctx).RegionId(regionId).Execute()

Get Tenants Config.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiGetAccountTenantConfigRequest{
        RegionId: &regionId
    }

    resp, r, err := api_client.Tenant.GetAccountTenantConfig(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tenant.GetAccountTenantConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountTenantConfig`: TenantConfig
    fmt.Fprintf(os.Stdout, "Response from `Tenant.GetAccountTenantConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountTenantConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**TenantConfig**](TenantConfig.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssignedPropertiesForTenant

> TenantAssignments GetAssignedPropertiesForTenant(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Get tenant assignments.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | The Tenant id.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiGetAssignedPropertiesForTenantRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Tenant.GetAssignedPropertiesForTenant(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tenant.GetAssignedPropertiesForTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssignedPropertiesForTenant`: TenantAssignments
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

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**TenantAssignments**](TenantAssignments.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClustersTenantConfig

> ClustersTenantConfigs GetClustersTenantConfig(ctx).RegionId(regionId).Execute()

Get Tenant's config for all clusters.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiGetClustersTenantConfigRequest{
        RegionId: &regionId
    }

    resp, r, err := api_client.Tenant.GetClustersTenantConfig(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tenant.GetClustersTenantConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClustersTenantConfig`: ClustersTenantConfigs
    fmt.Fprintf(os.Stdout, "Response from `Tenant.GetClustersTenantConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersTenantConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**ClustersTenantConfigs**](ClustersTenantConfigs.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHeliosTenants

> ListTenantData GetHeliosTenants(ctx).RegionId(regionId).ClusterIdentifiers(clusterIdentifiers).TenantIds(tenantIds).Statuses(statuses).ManagedOnHelios(managedOnHelios).Execute()

Get a list of tenants.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    clusterIdentifiers := []string{"Inner_example"} // []string | Specifies the list of cluster identifiers. The format is clusterId:clusterIncarnationId. (optional)
    tenantIds := []string{"Inner_example"} // []string | List of Tenant Ids to filter from. (optional)
    statuses := []string{"Statuses_example"} // []string | Status of the tenant on Helios. Only Active and Inactive are returned if this field is not specified. If specified, only helios managed tenants are returned. (optional)
    managedOnHelios := true // bool | Wether managed on helios or not. (optional)

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

    request := helios.ApiGetHeliosTenantsRequest{
        RegionId: &regionId
        ClusterIdentifiers: &clusterIdentifiers
        TenantIds: &tenantIds
        Statuses: &statuses
        ManagedOnHelios: &managedOnHelios
    }

    resp, r, err := api_client.Tenant.GetHeliosTenants(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tenant.GetHeliosTenants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHeliosTenants`: ListTenantData
    fmt.Fprintf(os.Stdout, "Response from `Tenant.GetHeliosTenants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHeliosTenantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **clusterIdentifiers** | **[]string** | Specifies the list of cluster identifiers. The format is clusterId:clusterIncarnationId. | 
 **tenantIds** | **[]string** | List of Tenant Ids to filter from. | 
 **statuses** | **[]string** | Status of the tenant on Helios. Only Active and Inactive are returned if this field is not specified. If specified, only helios managed tenants are returned. | 
 **managedOnHelios** | **bool** | Wether managed on helios or not. | 

### Return type

[**ListTenantData**](ListTenantData.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnPremTenantConfig

> OnPremTenantConfig GetOnPremTenantConfig(ctx).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Get Tenants Config.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiGetOnPremTenantConfigRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
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



### Other Parameters

Other parameters are passed through a pointer to a apiGetOnPremTenantConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

> TenantInfo GetTenantByID(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Get Tenant by ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | The Tenant id.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiGetTenantByIDRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
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

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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


## GetTenantStats

> []HeliosTenantResources GetTenantStats(ctx).RegionId(regionId).TenantIds(tenantIds).Execute()

Get Tenant Statistics.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    tenantIds := []string{"Inner_example"} // []string | List of Tenant Ids to filter from. (optional)

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

    request := helios.ApiGetTenantStatsRequest{
        RegionId: &regionId
        TenantIds: &tenantIds
    }

    resp, r, err := api_client.Tenant.GetTenantStats(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tenant.GetTenantStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenantStats`: []HeliosTenantResources
    fmt.Fprintf(os.Stdout, "Response from `Tenant.GetTenantStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **tenantIds** | **[]string** | List of Tenant Ids to filter from. | 

### Return type

[**[]HeliosTenantResources**](HeliosTenantResources.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantSwift

> SwiftParams GetTenantSwift(ctx).AccessClusterId(accessClusterId).RegionId(regionId).TenantId(tenantId).Execute()

Get a Swift configuration.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
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

    request := helios.ApiGetTenantSwiftRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
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
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
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


## HeliosAssignPropertiesToTenant

> HeliosTenantAssignmentsResult HeliosAssignPropertiesToTenant(ctx, id).Body(body).RegionId(regionId).Execute()

Assign properties to a tenant.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | The Tenant id.
    body := *openapiclient.NewHeliosTenantAssignmentsParams() // HeliosTenantAssignmentsParams | 
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiHeliosAssignPropertiesToTenantRequest{
        Id: &id
        Body: &body
        RegionId: &regionId
    }

    resp, r, err := api_client.Tenant.HeliosAssignPropertiesToTenant(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tenant.HeliosAssignPropertiesToTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HeliosAssignPropertiesToTenant`: HeliosTenantAssignmentsResult
    fmt.Fprintf(os.Stdout, "Response from `Tenant.HeliosAssignPropertiesToTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Tenant id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeliosAssignPropertiesToTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HeliosTenantAssignmentsParams**](HeliosTenantAssignmentsParams.md) |  | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**HeliosTenantAssignmentsResult**](HeliosTenantAssignmentsResult.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTenants

> TenantsInfo ListTenants(ctx).AccessClusterId(accessClusterId).RegionId(regionId).Ids(ids).Statuses(statuses).Execute()

Get a list of Tenants.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
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

    request := helios.ApiListTenantsRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
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
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
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


## PerformHeliosTenantAction

> HeliosTenant PerformHeliosTenantAction(ctx, id).Body(body).RegionId(regionId).Execute()

Perform actions on a Helios Tenant.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | The Tenant id.
    body := *openapiclient.NewTenantActionBody("Action_example") // TenantActionBody | Specifies the parameters to perform an action on a Tenant.
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiPerformHeliosTenantActionRequest{
        Id: &id
        Body: &body
        RegionId: &regionId
    }

    resp, r, err := api_client.Tenant.PerformHeliosTenantAction(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tenant.PerformHeliosTenantAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PerformHeliosTenantAction`: HeliosTenant
    fmt.Fprintf(os.Stdout, "Response from `Tenant.PerformHeliosTenantAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Tenant id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPerformHeliosTenantActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**TenantActionBody**](TenantActionBody.md) | Specifies the parameters to perform an action on a Tenant. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**HeliosTenant**](HeliosTenant.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PerformTenantAction

> TenantInfo PerformTenantAction(ctx, id).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Perform actions on a Tenant.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | The Tenant id.
    body := *openapiclient.NewTenantActionBody("Action_example") // TenantActionBody | Specifies the parameters to perform an action on a Tenant.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiPerformTenantActionRequest{
        Id: &id
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
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
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

> RegisterSwift(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Register Swift service on a Keystone server.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    body := *openapiclient.NewRegisterSwiftParams("TenantId_example") // RegisterSwiftParams | Specifies the parameters to register a Swift service on Keystone server.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiRegisterSwiftRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
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
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

> UnregisterSwift(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Unregister Swift service from a Keystone server.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    body := *openapiclient.NewUnregisterSwiftParams("TenantId_example") // UnregisterSwiftParams | Specifies the parameters to unregister a Swift service from Keystone server.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiUnregisterSwiftRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
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
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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


## UpdateAccountTenantConfig

> TenantConfig UpdateAccountTenantConfig(ctx).Body(body).RegionId(regionId).Execute()

Update Tenant's Config.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    body := *openapiclient.NewTenantConfig(false) // TenantConfig | 
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiUpdateAccountTenantConfigRequest{
        Body: &body
        RegionId: &regionId
    }

    resp, r, err := api_client.Tenant.UpdateAccountTenantConfig(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tenant.UpdateAccountTenantConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccountTenantConfig`: TenantConfig
    fmt.Fprintf(os.Stdout, "Response from `Tenant.UpdateAccountTenantConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountTenantConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TenantConfig**](TenantConfig.md) |  | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**TenantConfig**](TenantConfig.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClustersTenantConfig

> UpdateClustersTenantConfigsResponse UpdateClustersTenantConfig(ctx).Body(body).RegionId(regionId).Execute()

Update Tenant's config for clusters.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    body := *openapiclient.NewClustersTenantConfigs() // ClustersTenantConfigs | 
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiUpdateClustersTenantConfigRequest{
        Body: &body
        RegionId: &regionId
    }

    resp, r, err := api_client.Tenant.UpdateClustersTenantConfig(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tenant.UpdateClustersTenantConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateClustersTenantConfig`: UpdateClustersTenantConfigsResponse
    fmt.Fprintf(os.Stdout, "Response from `Tenant.UpdateClustersTenantConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClustersTenantConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ClustersTenantConfigs**](ClustersTenantConfigs.md) |  | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**UpdateClustersTenantConfigsResponse**](UpdateClustersTenantConfigsResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHeliosTenant

> HeliosTenant UpdateHeliosTenant(ctx, id).Body(body).RegionId(regionId).Execute()

Update Tenant properties on Helios.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | 
    body := *openapiclient.NewUpdateTenantParams() // UpdateTenantParams | 
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiUpdateHeliosTenantRequest{
        Id: &id
        Body: &body
        RegionId: &regionId
    }

    resp, r, err := api_client.Tenant.UpdateHeliosTenant(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tenant.UpdateHeliosTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateHeliosTenant`: HeliosTenant
    fmt.Fprintf(os.Stdout, "Response from `Tenant.UpdateHeliosTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHeliosTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateTenantParams**](UpdateTenantParams.md) |  | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**HeliosTenant**](HeliosTenant.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOnPremTenantConfig

> OnPremTenantConfig UpdateOnPremTenantConfig(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Update Tenants Config.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    body := *openapiclient.NewOnPremTenantConfig(false, false) // OnPremTenantConfig | 
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiUpdateOnPremTenantConfigRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
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
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

> TenantInfo UpdateTenant(ctx, id).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Update Tenant.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | 
    body := *openapiclient.NewUpdateTenantBody() // UpdateTenantBody | 
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiUpdateTenantRequest{
        Id: &id
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
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
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

> SwiftParams UpdateTenantSwift(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Update a Swift configuration.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    body := *openapiclient.NewSwiftParams("TenantId_example") // SwiftParams | Specifies the parameters to update a Swift configuration.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiUpdateTenantSwiftRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
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
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

