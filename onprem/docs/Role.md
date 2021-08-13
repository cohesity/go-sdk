# \Role

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRole**](Role.md#CreateRole) | **Post** /roles | Create a Role.
[**DeleteRole**](Role.md#DeleteRole) | **Delete** /roles/{name} | Delete a Role.
[**GetRoles**](Role.md#GetRoles) | **Get** /roles | Get Roles.
[**UpdateRole**](Role.md#UpdateRole) | **Put** /roles/{name} | Update a Role.



## CreateRole

> Role CreateRole(ctx).Body(body).Execute()

Create a Role.



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
    body := *openapiclient.NewCreateRoleParams("Name_example") // CreateRoleParams | Specifies the request body to create a Role.

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

    request := onprem.ApiCreateRoleRequest{
        Body: &body
    }

    resp, r, err := api_client.Role.CreateRole(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Role.CreateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRole`: Role
    fmt.Fprintf(os.Stdout, "Response from `Role.CreateRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateRoleParams**](CreateRoleParams.md) | Specifies the request body to create a Role. | 

### Return type

[**Role**](Role.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRole

> DeleteRole(ctx, name).Execute()

Delete a Role.



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
    name := "name_example" // string | Specifies the name of Role to delete.

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

    request := onprem.ApiDeleteRoleRequest{
        Name: &name
    }

    resp, r, err := api_client.Role.DeleteRole(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Role.DeleteRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Specifies the name of Role to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleRequest struct via the builder pattern


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


## GetRoles

> Roles GetRoles(ctx).Names(names).TenantIds(tenantIds).IncludeTenants(includeTenants).Execute()

Get Roles.



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
    names := []string{"Inner_example"} // []string | Specifies a list of Role names. (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which Roles are to be returned. (optional)
    includeTenants := true // bool | If true, the response will include Roles which were created by all tenants which the current user has permission to see. If false, then only Roles created by the current user will be returned. (optional)

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

    request := onprem.ApiGetRolesRequest{
        Names: &names
        TenantIds: &tenantIds
        IncludeTenants: &includeTenants
    }

    resp, r, err := api_client.Role.GetRoles(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Role.GetRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoles`: Roles
    fmt.Fprintf(os.Stdout, "Response from `Role.GetRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **names** | **[]string** | Specifies a list of Role names. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which Roles are to be returned. | 
 **includeTenants** | **bool** | If true, the response will include Roles which were created by all tenants which the current user has permission to see. If false, then only Roles created by the current user will be returned. | 

### Return type

[**Roles**](Roles.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> Role UpdateRole(ctx, name).Body(body).Execute()

Update a Role.



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
    name := "name_example" // string | Specifies the name of Role to update.
    body := *openapiclient.NewUpdateRoleParams() // UpdateRoleParams | Specifies the request body to update a Role.

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

    request := onprem.ApiUpdateRoleRequest{
        Name: &name
        Body: &body
    }

    resp, r, err := api_client.Role.UpdateRole(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Role.UpdateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRole`: Role
    fmt.Fprintf(os.Stdout, "Response from `Role.UpdateRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Specifies the name of Role to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateRoleParams**](UpdateRoleParams.md) | Specifies the request body to update a Role. | 

### Return type

[**Role**](Role.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

