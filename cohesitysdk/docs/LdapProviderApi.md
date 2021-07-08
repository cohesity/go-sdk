# \LdapProviderApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLdapProvider**](LdapProviderApi.md#CreateLdapProvider) | **Post** /public/ldapProvider | Create a LDAP provider.
[**DeleteLdapProvider**](LdapProviderApi.md#DeleteLdapProvider) | **Delete** /public/ldapProvider/{id} | Delete an LDAP provider.
[**GetLdapProvider**](LdapProviderApi.md#GetLdapProvider) | **Get** /public/ldapProvider | Lists the LDAP providers.
[**GetLdapProviderStatus**](LdapProviderApi.md#GetLdapProviderStatus) | **Get** /public/ldapProvider/{id}/status | Get the connection status of an LDAP provider.
[**UpdateLdapProvider**](LdapProviderApi.md#UpdateLdapProvider) | **Put** /public/ldapProvider | Update an LDAP provider.



## CreateLdapProvider

> LdapProviderResponse CreateLdapProvider(ctx).Body(body).Execute()

Create a LDAP provider.



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
    body := *openapiclient.NewLdapProvider() // LdapProvider | Request to configure a LDAP provider.

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

    request := cohesitysdk.ApiCreateLdapProviderRequest{
        Body: &body
    }

    resp, r, err := api_client.LdapProviderApi.CreateLdapProvider(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LdapProviderApi.CreateLdapProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLdapProvider`: LdapProviderResponse
    fmt.Fprintf(os.Stdout, "Response from `LdapProviderApi.CreateLdapProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLdapProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LdapProvider**](LdapProvider.md) | Request to configure a LDAP provider. | 

### Return type

[**LdapProviderResponse**](LdapProviderResponse.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLdapProvider

> DeleteLdapProvider(ctx, id).Execute()

Delete an LDAP provider.

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
    id := int64(789) // int64 | Specifies the LDAP Id.

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

    request := cohesitysdk.ApiDeleteLdapProviderRequest{
        Id: &id
    }

    resp, r, err := api_client.LdapProviderApi.DeleteLdapProvider(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LdapProviderApi.DeleteLdapProvider``: %v\n", err)
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

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLdapProvider

> []LdapProviderResponse GetLdapProvider(ctx).Ids(ids).TenantIds(tenantIds).AllUnderHierarchy(allUnderHierarchy).Execute()

Lists the LDAP providers.

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
    ids := []int64{int64(123)} // []int64 | Specifies the ids of the LDAP providers to fetch. (optional)
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

    request := cohesitysdk.ApiGetLdapProviderRequest{
        Ids: &ids
        TenantIds: &tenantIds
        AllUnderHierarchy: &allUnderHierarchy
    }

    resp, r, err := api_client.LdapProviderApi.GetLdapProvider(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LdapProviderApi.GetLdapProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLdapProvider`: []LdapProviderResponse
    fmt.Fprintf(os.Stdout, "Response from `LdapProviderApi.GetLdapProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLdapProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]int64** | Specifies the ids of the LDAP providers to fetch. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **allUnderHierarchy** | **bool** | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. | 

### Return type

[**[]LdapProviderResponse**](LdapProviderResponse.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLdapProviderStatus

> GetLdapProviderStatus(ctx, id).Execute()

Get the connection status of an LDAP provider.

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
    id := int64(789) // int64 | Specifies the LDAP Id.

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

    request := cohesitysdk.ApiGetLdapProviderStatusRequest{
        Id: &id
    }

    resp, r, err := api_client.LdapProviderApi.GetLdapProviderStatus(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LdapProviderApi.GetLdapProviderStatus``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetLdapProviderStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLdapProvider

> LdapProviderResponse UpdateLdapProvider(ctx).Body(body).Execute()

Update an LDAP provider.



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
    body := *openapiclient.NewUpdateLdapProviderParam() // UpdateLdapProviderParam | Request to update a LDAP provider.

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

    request := cohesitysdk.ApiUpdateLdapProviderRequest{
        Body: &body
    }

    resp, r, err := api_client.LdapProviderApi.UpdateLdapProvider(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LdapProviderApi.UpdateLdapProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLdapProvider`: LdapProviderResponse
    fmt.Fprintf(os.Stdout, "Response from `LdapProviderApi.UpdateLdapProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLdapProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateLdapProviderParam**](UpdateLdapProviderParam.md) | Request to update a LDAP provider. | 

### Return type

[**LdapProviderResponse**](LdapProviderResponse.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

