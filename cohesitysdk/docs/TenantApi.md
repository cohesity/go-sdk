# \TenantApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTenant**](TenantApi.md#CreateTenant) | **Post** /public/tenants | Create/add a new tenant to the Cohesity Cluster.
[**DeleteTenant**](TenantApi.md#DeleteTenant) | **Delete** /public/tenants | Delete an existing tenant on the Cohesity Cluster.
[**DownloadTenantsProxy**](TenantApi.md#DownloadTenantsProxy) | **Get** /public/tenants/proxy/image | Downloads the tenants proxy.
[**GetTenants**](TenantApi.md#GetTenants) | **Get** /public/tenants | List the tenants on the cohesity cluster filtered by tenant ID prefixed to list tenants for the respective tenant admin.
[**GetTenantsProxies**](TenantApi.md#GetTenantsProxies) | **Get** /public/tenants/proxies | List proxies for tenant which are running within tenant&#39;s environment.
[**GetTenantsProxyConfig**](TenantApi.md#GetTenantsProxyConfig) | **Get** /public/tenants/proxy/config | Get proxy config for tenant.
[**UpdateTenant**](TenantApi.md#UpdateTenant) | **Put** /public/tenants | Update an existing tenant on the Cohesity Cluster.
[**UpdateTenantActiveDirectory**](TenantApi.md#UpdateTenantActiveDirectory) | **Put** /public/tenants/activeDirectory | Update Active Directory for an existing tenant on the Cohesity Cluster.
[**UpdateTenantEntity**](TenantApi.md#UpdateTenantEntity) | **Put** /public/tenants/entity | Update entity permission for an existing tenant on the Cohesity Cluster.
[**UpdateTenantGroups**](TenantApi.md#UpdateTenantGroups) | **Put** /public/tenants/groups | Update existing groups to an existing tenant on the Cohesity Cluster.
[**UpdateTenantLdapProvider**](TenantApi.md#UpdateTenantLdapProvider) | **Put** /public/tenants/ldapProvider | Update Ldap Providers for an existing tenant on the Cohesity Cluster.
[**UpdateTenantProtectionJob**](TenantApi.md#UpdateTenantProtectionJob) | **Put** /public/tenants/protectionJob | Update protection job for an existing tenant on the Cohesity Cluster.
[**UpdateTenantProtectionPolicy**](TenantApi.md#UpdateTenantProtectionPolicy) | **Put** /public/tenants/policy | Update protection policy permission for an existing tenant on the Cohesity Cluster.
[**UpdateTenantUsers**](TenantApi.md#UpdateTenantUsers) | **Put** /public/tenants/users | Update existing users to an existing tenant on the Cohesity Cluster.
[**UpdateTenantView**](TenantApi.md#UpdateTenantView) | **Put** /public/tenants/view | Update views permission for an existing tenant on the Cohesity Cluster.
[**UpdateTenantViewBox**](TenantApi.md#UpdateTenantViewBox) | **Put** /public/tenants/viewBox | Update view box for an existing tenant on the Cohesity Cluster.
[**UpdateTenantVlan**](TenantApi.md#UpdateTenantVlan) | **Put** /public/tenants/vlan | Update vlan for an existing tenant on the Cohesity Cluster.



## CreateTenant

> Tenant CreateTenant(ctx).Body(body).Execute()

Create/add a new tenant to the Cohesity Cluster.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewTenantCreateParameters() // TenantCreateParameters | Request to add or create a new tenant. (optional)

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

    request := cohesitysdk.ApiCreateTenantRequest{
        Body: &body
    }

    resp, r, err := api_client.TenantApi.CreateTenant(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.CreateTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTenant`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenantApi.CreateTenant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TenantCreateParameters**](TenantCreateParameters.md) | Request to add or create a new tenant. | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTenant

> []Tenant DeleteTenant(ctx).Body(body).Execute()

Delete an existing tenant on the Cohesity Cluster.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewTenantIdData() // TenantIdData |  (optional)

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

    request := cohesitysdk.ApiDeleteTenantRequest{
        Body: &body
    }

    resp, r, err := api_client.TenantApi.DeleteTenant(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.DeleteTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTenant`: []Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenantApi.DeleteTenant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TenantIdData**](TenantIdData.md) |  | 

### Return type

[**[]Tenant**](Tenant.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadTenantsProxy

> []int32 DownloadTenantsProxy(ctx).Id(id).Execute()

Downloads the tenants proxy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    id := "id_example" // string | Specifies the id of the tenant. (optional)

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

    request := cohesitysdk.ApiDownloadTenantsProxyRequest{
        Id: &id
    }

    resp, r, err := api_client.TenantApi.DownloadTenantsProxy(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.DownloadTenantsProxy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadTenantsProxy`: []int32
    fmt.Fprintf(os.Stdout, "Response from `TenantApi.DownloadTenantsProxy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadTenantsProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Specifies the id of the tenant. | 

### Return type

**[]int32**

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenants

> []Tenant GetTenants(ctx).Ids(ids).Properties(properties).Hierarchy(hierarchy).IncludeSelf(includeSelf).SkipInvalidIds(skipInvalidIds).Status(status).Execute()

List the tenants on the cohesity cluster filtered by tenant ID prefixed to list tenants for the respective tenant admin.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    ids := []string{"Inner_example"} // []string |  (optional)
    properties := []string{"Properties_example"} // []string | 'ViewBox' indicates view box data for tenant. 'Vlan' indicates vlan data for tenant. 'ProtectionPolicy' indicates protection policy for tenant. 'ProtectionJob' indicates protection job for tenant. 'Entity' indicates entity data for tenant. 'Views' indicates view data for tenant. 'ActiveDirectory' indicates active directory for tenant. 'LdapProvider' indicates ldap provider for tenant. 'SwiftConfig' indicates Swift configuration for tenant. (optional)
    hierarchy := true // bool |  (optional)
    includeSelf := true // bool |  (optional)
    skipInvalidIds := true // bool |  (optional)
    status := []string{"Status_example"} // []string | Filter by tenant status. (optional)

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

    request := cohesitysdk.ApiGetTenantsRequest{
        Ids: &ids
        Properties: &properties
        Hierarchy: &hierarchy
        IncludeSelf: &includeSelf
        SkipInvalidIds: &skipInvalidIds
        Status: &status
    }

    resp, r, err := api_client.TenantApi.GetTenants(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.GetTenants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenants`: []Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenantApi.GetTenants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** |  | 
 **properties** | **[]string** | &#39;ViewBox&#39; indicates view box data for tenant. &#39;Vlan&#39; indicates vlan data for tenant. &#39;ProtectionPolicy&#39; indicates protection policy for tenant. &#39;ProtectionJob&#39; indicates protection job for tenant. &#39;Entity&#39; indicates entity data for tenant. &#39;Views&#39; indicates view data for tenant. &#39;ActiveDirectory&#39; indicates active directory for tenant. &#39;LdapProvider&#39; indicates ldap provider for tenant. &#39;SwiftConfig&#39; indicates Swift configuration for tenant. | 
 **hierarchy** | **bool** |  | 
 **includeSelf** | **bool** |  | 
 **skipInvalidIds** | **bool** |  | 
 **status** | **[]string** | Filter by tenant status. | 

### Return type

[**[]Tenant**](Tenant.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantsProxies

> []TenantProxy GetTenantsProxies(ctx).Ids(ids).Execute()

List proxies for tenant which are running within tenant's environment.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    ids := []string{"Inner_example"} // []string |  (optional)

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

    request := cohesitysdk.ApiGetTenantsProxiesRequest{
        Ids: &ids
    }

    resp, r, err := api_client.TenantApi.GetTenantsProxies(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.GetTenantsProxies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenantsProxies`: []TenantProxy
    fmt.Fprintf(os.Stdout, "Response from `TenantApi.GetTenantsProxies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantsProxiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** |  | 

### Return type

[**[]TenantProxy**](TenantProxy.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantsProxyConfig

> []int32 GetTenantsProxyConfig(ctx).Id(id).ValidateOnly(validateOnly).Execute()

Get proxy config for tenant.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    id := "id_example" // string | Specifies the id of the tenant. (optional)
    validateOnly := true // bool | Specifies whether to only validate the config request. (optional)

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

    request := cohesitysdk.ApiGetTenantsProxyConfigRequest{
        Id: &id
        ValidateOnly: &validateOnly
    }

    resp, r, err := api_client.TenantApi.GetTenantsProxyConfig(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.GetTenantsProxyConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenantsProxyConfig`: []int32
    fmt.Fprintf(os.Stdout, "Response from `TenantApi.GetTenantsProxyConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantsProxyConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Specifies the id of the tenant. | 
 **validateOnly** | **bool** | Specifies whether to only validate the config request. | 

### Return type

**[]int32**

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTenant

> Tenant UpdateTenant(ctx).Body(body).Execute()

Update an existing tenant on the Cohesity Cluster.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewTenantUpdate() // TenantUpdate | Request to update existing tenant. (optional)

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

    request := cohesitysdk.ApiUpdateTenantRequest{
        Body: &body
    }

    resp, r, err := api_client.TenantApi.UpdateTenant(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.UpdateTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTenant`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenantApi.UpdateTenant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TenantUpdate**](TenantUpdate.md) | Request to update existing tenant. | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTenantActiveDirectory

> TenantActiveDirectoryUpdate UpdateTenantActiveDirectory(ctx).Body(body).Execute()

Update Active Directory for an existing tenant on the Cohesity Cluster.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewTenantActiveDirectoryUpdateParameters() // TenantActiveDirectoryUpdateParameters | Request to update existing active directories. (optional)

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

    request := cohesitysdk.ApiUpdateTenantActiveDirectoryRequest{
        Body: &body
    }

    resp, r, err := api_client.TenantApi.UpdateTenantActiveDirectory(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.UpdateTenantActiveDirectory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTenantActiveDirectory`: TenantActiveDirectoryUpdate
    fmt.Fprintf(os.Stdout, "Response from `TenantApi.UpdateTenantActiveDirectory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantActiveDirectoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TenantActiveDirectoryUpdateParameters**](TenantActiveDirectoryUpdateParameters.md) | Request to update existing active directories. | 

### Return type

[**TenantActiveDirectoryUpdate**](TenantActiveDirectoryUpdate.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTenantEntity

> TenantEntityUpdate UpdateTenantEntity(ctx).Body(body).Execute()

Update entity permission for an existing tenant on the Cohesity Cluster.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewTenantEntityUpdateParameters() // TenantEntityUpdateParameters | Request to associate entity for tenant. (optional)

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

    request := cohesitysdk.ApiUpdateTenantEntityRequest{
        Body: &body
    }

    resp, r, err := api_client.TenantApi.UpdateTenantEntity(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.UpdateTenantEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTenantEntity`: TenantEntityUpdate
    fmt.Fprintf(os.Stdout, "Response from `TenantApi.UpdateTenantEntity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TenantEntityUpdateParameters**](TenantEntityUpdateParameters.md) | Request to associate entity for tenant. | 

### Return type

[**TenantEntityUpdate**](TenantEntityUpdate.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTenantGroups

> []Group UpdateTenantGroups(ctx).Body(body).Execute()

Update existing groups to an existing tenant on the Cohesity Cluster.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewTenantGroupUpdateParameters() // TenantGroupUpdateParameters | Request to update existing tenant groups. (optional)

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

    request := cohesitysdk.ApiUpdateTenantGroupsRequest{
        Body: &body
    }

    resp, r, err := api_client.TenantApi.UpdateTenantGroups(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.UpdateTenantGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTenantGroups`: []Group
    fmt.Fprintf(os.Stdout, "Response from `TenantApi.UpdateTenantGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TenantGroupUpdateParameters**](TenantGroupUpdateParameters.md) | Request to update existing tenant groups. | 

### Return type

[**[]Group**](Group.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTenantLdapProvider

> TenantLdapProviderUpdate UpdateTenantLdapProvider(ctx).Body(body).Execute()

Update Ldap Providers for an existing tenant on the Cohesity Cluster.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewTenantLdapProviderUpdateParameters() // TenantLdapProviderUpdateParameters | Request to update existing ldap providers. (optional)

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

    request := cohesitysdk.ApiUpdateTenantLdapProviderRequest{
        Body: &body
    }

    resp, r, err := api_client.TenantApi.UpdateTenantLdapProvider(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.UpdateTenantLdapProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTenantLdapProvider`: TenantLdapProviderUpdate
    fmt.Fprintf(os.Stdout, "Response from `TenantApi.UpdateTenantLdapProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantLdapProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TenantLdapProviderUpdateParameters**](TenantLdapProviderUpdateParameters.md) | Request to update existing ldap providers. | 

### Return type

[**TenantLdapProviderUpdate**](TenantLdapProviderUpdate.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTenantProtectionJob

> TenantProtectionJobUpdate UpdateTenantProtectionJob(ctx).Body(body).Execute()

Update protection job for an existing tenant on the Cohesity Cluster.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewTenantProtectionJobUpdateParameters() // TenantProtectionJobUpdateParameters | Request to update existing protection jobs. (optional)

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

    request := cohesitysdk.ApiUpdateTenantProtectionJobRequest{
        Body: &body
    }

    resp, r, err := api_client.TenantApi.UpdateTenantProtectionJob(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.UpdateTenantProtectionJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTenantProtectionJob`: TenantProtectionJobUpdate
    fmt.Fprintf(os.Stdout, "Response from `TenantApi.UpdateTenantProtectionJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantProtectionJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TenantProtectionJobUpdateParameters**](TenantProtectionJobUpdateParameters.md) | Request to update existing protection jobs. | 

### Return type

[**TenantProtectionJobUpdate**](TenantProtectionJobUpdate.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTenantProtectionPolicy

> TenantProtectionPolicyUpdate UpdateTenantProtectionPolicy(ctx).Body(body).Execute()

Update protection policy permission for an existing tenant on the Cohesity Cluster.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewTenantProtectionPolicyUpdateParameters() // TenantProtectionPolicyUpdateParameters | Request to update existing protection policies. (optional)

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

    request := cohesitysdk.ApiUpdateTenantProtectionPolicyRequest{
        Body: &body
    }

    resp, r, err := api_client.TenantApi.UpdateTenantProtectionPolicy(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.UpdateTenantProtectionPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTenantProtectionPolicy`: TenantProtectionPolicyUpdate
    fmt.Fprintf(os.Stdout, "Response from `TenantApi.UpdateTenantProtectionPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantProtectionPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TenantProtectionPolicyUpdateParameters**](TenantProtectionPolicyUpdateParameters.md) | Request to update existing protection policies. | 

### Return type

[**TenantProtectionPolicyUpdate**](TenantProtectionPolicyUpdate.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTenantUsers

> []User UpdateTenantUsers(ctx).Body(body).Execute()

Update existing users to an existing tenant on the Cohesity Cluster.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewTenantUserUpdateParameters() // TenantUserUpdateParameters | Request to update existing users. (optional)

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

    request := cohesitysdk.ApiUpdateTenantUsersRequest{
        Body: &body
    }

    resp, r, err := api_client.TenantApi.UpdateTenantUsers(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.UpdateTenantUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTenantUsers`: []User
    fmt.Fprintf(os.Stdout, "Response from `TenantApi.UpdateTenantUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TenantUserUpdateParameters**](TenantUserUpdateParameters.md) | Request to update existing users. | 

### Return type

[**[]User**](User.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTenantView

> TenantViewUpdate UpdateTenantView(ctx).Body(body).Execute()

Update views permission for an existing tenant on the Cohesity Cluster.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewTenantViewUpdateParameters() // TenantViewUpdateParameters | Request to update existing views. (optional)

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

    request := cohesitysdk.ApiUpdateTenantViewRequest{
        Body: &body
    }

    resp, r, err := api_client.TenantApi.UpdateTenantView(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.UpdateTenantView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTenantView`: TenantViewUpdate
    fmt.Fprintf(os.Stdout, "Response from `TenantApi.UpdateTenantView`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TenantViewUpdateParameters**](TenantViewUpdateParameters.md) | Request to update existing views. | 

### Return type

[**TenantViewUpdate**](TenantViewUpdate.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTenantViewBox

> TenantViewBoxUpdate UpdateTenantViewBox(ctx).Body(body).Execute()

Update view box for an existing tenant on the Cohesity Cluster.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewTenantViewBoxUpdateParameters() // TenantViewBoxUpdateParameters | Request to update existing tenant view box. (optional)

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

    request := cohesitysdk.ApiUpdateTenantViewBoxRequest{
        Body: &body
    }

    resp, r, err := api_client.TenantApi.UpdateTenantViewBox(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.UpdateTenantViewBox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTenantViewBox`: TenantViewBoxUpdate
    fmt.Fprintf(os.Stdout, "Response from `TenantApi.UpdateTenantViewBox`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantViewBoxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TenantViewBoxUpdateParameters**](TenantViewBoxUpdateParameters.md) | Request to update existing tenant view box. | 

### Return type

[**TenantViewBoxUpdate**](TenantViewBoxUpdate.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTenantVlan

> TenantVlanUpdate UpdateTenantVlan(ctx).Body(body).Execute()

Update vlan for an existing tenant on the Cohesity Cluster.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewTenantVlanUpdateParameters() // TenantVlanUpdateParameters | Request to update existing tenant vlan. (optional)

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

    request := cohesitysdk.ApiUpdateTenantVlanRequest{
        Body: &body
    }

    resp, r, err := api_client.TenantApi.UpdateTenantVlan(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.UpdateTenantVlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTenantVlan`: TenantVlanUpdate
    fmt.Fprintf(os.Stdout, "Response from `TenantApi.UpdateTenantVlan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TenantVlanUpdateParameters**](TenantVlanUpdateParameters.md) | Request to update existing tenant vlan. | 

### Return type

[**TenantVlanUpdate**](TenantVlanUpdate.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

