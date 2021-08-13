# \LDAP

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLdapConnectionStatus**](LDAP.md#GetLdapConnectionStatus) | **Get** /ldap/{id}/connection-status | Get LDAP connection status.
[**GetLdaps**](LDAP.md#GetLdaps) | **Get** /ldap | Get Groups.



## GetLdapConnectionStatus

> LdapStatus GetLdapConnectionStatus(ctx, id).Execute()

Get LDAP connection status.



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
    id := int64(789) // int64 | Specifies the LDAP id.

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

    request := onprem.ApiGetLdapConnectionStatusRequest{
        Id: &id
    }

    resp, r, err := api_client.LDAP.GetLdapConnectionStatus(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAP.GetLdapConnectionStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLdapConnectionStatus`: LdapStatus
    fmt.Fprintf(os.Stdout, "Response from `LDAP.GetLdapConnectionStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the LDAP id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLdapConnectionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LdapStatus**](LdapStatus.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLdaps

> Ldaps GetLdaps(ctx).Ids(ids).TenantIds(tenantIds).IncludeTenants(includeTenants).Execute()

Get Groups.



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
    ids := []int64{int64(123)} // []int64 | Specifies a list of ids to filter. (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which LDAPs are to be returned. (optional)
    includeTenants := true // bool | IncludeTenants specifies if LDAPs of all the tenants under the hierarchy of the logged in user's organization should be returned. (optional)

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

    request := onprem.ApiGetLdapsRequest{
        Ids: &ids
        TenantIds: &tenantIds
        IncludeTenants: &includeTenants
    }

    resp, r, err := api_client.LDAP.GetLdaps(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAP.GetLdaps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLdaps`: Ldaps
    fmt.Fprintf(os.Stdout, "Response from `LDAP.GetLdaps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLdapsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]int64** | Specifies a list of ids to filter. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which LDAPs are to be returned. | 
 **includeTenants** | **bool** | IncludeTenants specifies if LDAPs of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. | 

### Return type

[**Ldaps**](Ldaps.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

