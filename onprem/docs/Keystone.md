# \Keystone

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKeystone**](Keystone.md#CreateKeystone) | **Post** /keystones | Create a Keystone configuration.
[**DeleteKeystone**](Keystone.md#DeleteKeystone) | **Delete** /keystones/{id} | Delete a Keystone configuration.
[**GetKeystones**](Keystone.md#GetKeystones) | **Get** /keystones | Get Keystones.
[**GetKeystonesById**](Keystone.md#GetKeystonesById) | **Get** /keystones/{id} | Get a Keystone by its id.
[**UpdateKeystone**](Keystone.md#UpdateKeystone) | **Put** /keystones/{id} | Update a Keystone configuration.



## CreateKeystone

> Keystone CreateKeystone(ctx).Body(body).Execute()

Create a Keystone configuration.



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
    body := *openapiclient.NewCreateKeystoneRequest(*openapiclient.NewKeystoneAdminParams("Domain_example", "Username_example"), *openapiclient.NewKeystoneScopeParams("Type_example"), "Name_example", "AuthUrl_example") // CreateKeystoneRequest | Specifies the paremters to create a Keystone configuration.

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

    request := onprem.ApiCreateKeystoneRequest{
        Body: &body
    }

    resp, r, err := api_client.Keystone.CreateKeystone(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Keystone.CreateKeystone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKeystone`: Keystone
    fmt.Fprintf(os.Stdout, "Response from `Keystone.CreateKeystone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKeystoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateKeystoneRequest**](CreateKeystoneRequest.md) | Specifies the paremters to create a Keystone configuration. | 

### Return type

[**Keystone**](Keystone.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKeystone

> DeleteKeystone(ctx, id).AdminPassword(adminPassword).Execute()

Delete a Keystone configuration.



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
    id := int64(789) // int64 | Specifies the Keystone id.
    adminPassword := "adminPassword_example" // string | Specifies the password of Keystone administrator.

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

    request := onprem.ApiDeleteKeystoneRequest{
        Id: &id
        AdminPassword: &adminPassword
    }

    resp, r, err := api_client.Keystone.DeleteKeystone(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Keystone.DeleteKeystone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the Keystone id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeystoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adminPassword** | **string** | Specifies the password of Keystone administrator. | 

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


## GetKeystones

> Keystones GetKeystones(ctx).Names(names).TenantIds(tenantIds).IncludeTenants(includeTenants).Execute()

Get Keystones.



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
    names := []string{"Inner_example"} // []string | Specifies a list of Keystone names. (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
    includeTenants := true // bool | If true, the response will include Keystones which were created by all tenants which the current user has permission to see. If false, then only Keystones created by the current user will be returned. (optional)

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

    request := onprem.ApiGetKeystonesRequest{
        Names: &names
        TenantIds: &tenantIds
        IncludeTenants: &includeTenants
    }

    resp, r, err := api_client.Keystone.GetKeystones(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Keystone.GetKeystones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKeystones`: Keystones
    fmt.Fprintf(os.Stdout, "Response from `Keystone.GetKeystones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKeystonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **names** | **[]string** | Specifies a list of Keystone names. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **includeTenants** | **bool** | If true, the response will include Keystones which were created by all tenants which the current user has permission to see. If false, then only Keystones created by the current user will be returned. | 

### Return type

[**Keystones**](Keystones.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKeystonesById

> Keystone GetKeystonesById(ctx, id).Execute()

Get a Keystone by its id.



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
    id := int64(789) // int64 | Specifies the Keystone id.

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

    request := onprem.ApiGetKeystonesByIdRequest{
        Id: &id
    }

    resp, r, err := api_client.Keystone.GetKeystonesById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Keystone.GetKeystonesById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKeystonesById`: Keystone
    fmt.Fprintf(os.Stdout, "Response from `Keystone.GetKeystonesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the Keystone id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeystonesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Keystone**](Keystone.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKeystone

> Keystone UpdateKeystone(ctx, id).Body(body).Execute()

Update a Keystone configuration.



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
    id := int64(789) // int64 | Specifies the Keystone id.
    body := *openapiclient.NewUpdateKeystoneRequest(*openapiclient.NewKeystoneAdminParams("Domain_example", "Username_example"), *openapiclient.NewKeystoneScopeParams("Type_example"), "Name_example", "AuthUrl_example") // UpdateKeystoneRequest | Specifies the paremters to update a Keystone configuration.

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

    request := onprem.ApiUpdateKeystoneRequest{
        Id: &id
        Body: &body
    }

    resp, r, err := api_client.Keystone.UpdateKeystone(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Keystone.UpdateKeystone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKeystone`: Keystone
    fmt.Fprintf(os.Stdout, "Response from `Keystone.UpdateKeystone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the Keystone id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKeystoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateKeystoneRequest**](UpdateKeystoneRequest.md) | Specifies the paremters to update a Keystone configuration. | 

### Return type

[**Keystone**](Keystone.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

