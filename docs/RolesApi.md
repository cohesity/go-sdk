# \RolesApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRole**](RolesApi.md#CreateRole) | **Post** /public/roles | Create a new custom role.
[**DeleteRoles**](RolesApi.md#DeleteRoles) | **Delete** /public/roles | Delete one or more custom Roles.
[**GetRoles**](RolesApi.md#GetRoles) | **Get** /public/roles | List the roles defined on the Cohesity Cluster.
[**UpdateRole**](RolesApi.md#UpdateRole) | **Put** /public/roles/{name} | Update a user-defined custom role.



## CreateRole

> Role CreateRole(ctx).Body(body).Execute()

Create a new custom role.



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
    body := *openapiclient.NewRoleCreateParameters() // RoleCreateParameters | Request to create a new custom Role. (optional)

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

    request := cohesitysdk.ApiCreateRoleRequest{
        body: &Body
    }

    resp, r, err := api_client.RolesApi.CreateRole(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.CreateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRole`: Role
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.CreateRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RoleCreateParameters**](RoleCreateParameters.md) | Request to create a new custom Role. | 

### Return type

[**Role**](Role.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoles

> DeleteRoles(ctx).Body(body).Execute()

Delete one or more custom Roles.



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
    body := *openapiclient.NewRoleDeleteParameters([]string{"Names_example"}) // RoleDeleteParameters | Request to delete one or more Roles. (optional)

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

    request := cohesitysdk.ApiDeleteRolesRequest{
        body: &Body
    }

    resp, r, err := api_client.RolesApi.DeleteRoles(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.DeleteRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RoleDeleteParameters**](RoleDeleteParameters.md) | Request to delete one or more Roles. | 

### Return type

 (empty response body)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoles

> []Role GetRoles(ctx).Name(name).TenantIds(tenantIds).AllUnderHierarchy(allUnderHierarchy).Execute()

List the roles defined on the Cohesity Cluster.



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
    name := "name_example" // string | Specifies the name of the role. (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
    allUnderHierarchy := true // bool | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user's organization should be returned. (optional)

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

    request := cohesitysdk.ApiGetRolesRequest{
        name: &Name
        tenantIds: &TenantIds
        allUnderHierarchy: &AllUnderHierarchy
    }

    resp, r, err := api_client.RolesApi.GetRoles(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.GetRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoles`: []Role
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.GetRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Specifies the name of the role. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **allUnderHierarchy** | **bool** | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. | 

### Return type

[**[]Role**](Role.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> Role UpdateRole(ctx, name).Body(body).Execute()

Update a user-defined custom role.



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
    name := "name_example" // string | Specifies the name of the role to update.
    body := *openapiclient.NewRoleUpdateParameters() // RoleUpdateParameters | Request to update a custom role. (optional)

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

    request := cohesitysdk.ApiUpdateRoleRequest{
        name: &Name
        body: &Body
    }

    resp, r, err := api_client.RolesApi.UpdateRole(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.UpdateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRole`: Role
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.UpdateRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Specifies the name of the role to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RoleUpdateParameters**](RoleUpdateParameters.md) | Request to update a custom role. | 

### Return type

[**Role**](Role.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

