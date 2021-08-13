# \KerberosProvider

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetKerberosProviderById**](KerberosProvider.md#GetKerberosProviderById) | **Get** /kerberos-providers/{id} | Get the Registered Kerberos Provider by id.
[**GetKerberosProviders**](KerberosProvider.md#GetKerberosProviders) | **Get** /kerberos-providers | Get the list of Kerberos Providers.
[**RegisterKerberosProvider**](KerberosProvider.md#RegisterKerberosProvider) | **Post** /kerberos-providers/register | Register a Kerberos Authentication Provider.
[**UnregisterKerberosProvider**](KerberosProvider.md#UnregisterKerberosProvider) | **Post** /kerberos-providers/{id}/unregister | Unregister a Kerberos Provider.
[**UpdateKerberosProvider**](KerberosProvider.md#UpdateKerberosProvider) | **Put** /kerberos-providers/{id} | Update the Kerberos Provider Registration.



## GetKerberosProviderById

> KerberosProvider GetKerberosProviderById(ctx, id).Execute()

Get the Registered Kerberos Provider by id.



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
    id := "id_example" // string | Specifies the id which will be of the pattern cluster_id:clusterincarnation_id:resource_id.

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

    request := onprem.ApiGetKerberosProviderByIdRequest{
        Id: &id
    }

    resp, r, err := api_client.KerberosProvider.GetKerberosProviderById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KerberosProvider.GetKerberosProviderById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKerberosProviderById`: KerberosProvider
    fmt.Fprintf(os.Stdout, "Response from `KerberosProvider.GetKerberosProviderById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id which will be of the pattern cluster_id:clusterincarnation_id:resource_id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKerberosProviderByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KerberosProvider**](KerberosProvider.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKerberosProviders

> KerberosProviders GetKerberosProviders(ctx).RealmNames(realmNames).Ids(ids).KdcServers(kdcServers).Execute()

Get the list of Kerberos Providers.



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
    realmNames := []string{"Inner_example"} // []string | Filter by a list of realm names. (optional)
    ids := []int64{int64(123)} // []int64 | Filter by a list of Kerberos Provider Ids. (optional)
    kdcServers := []string{"Inner_example"} // []string | Filter by a list of KDC servers. (optional)

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

    request := onprem.ApiGetKerberosProvidersRequest{
        RealmNames: &realmNames
        Ids: &ids
        KdcServers: &kdcServers
    }

    resp, r, err := api_client.KerberosProvider.GetKerberosProviders(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KerberosProvider.GetKerberosProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKerberosProviders`: KerberosProviders
    fmt.Fprintf(os.Stdout, "Response from `KerberosProvider.GetKerberosProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKerberosProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **realmNames** | **[]string** | Filter by a list of realm names. | 
 **ids** | **[]int64** | Filter by a list of Kerberos Provider Ids. | 
 **kdcServers** | **[]string** | Filter by a list of KDC servers. | 

### Return type

[**KerberosProviders**](KerberosProviders.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterKerberosProvider

> KerberosProvider RegisterKerberosProvider(ctx).Body(body).Execute()

Register a Kerberos Authentication Provider.



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
    body := *openapiclient.NewRegisterOrUpdateKerberosProviderRequest("RealmName_example", []string{"KdcServers_example"}, "AdminServer_example", []string{"HostAlias_example"}, "AdminPassword_example") // RegisterOrUpdateKerberosProviderRequest | Specifies the parameters to Register a Kerberos Provider.

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

    request := onprem.ApiRegisterKerberosProviderRequest{
        Body: &body
    }

    resp, r, err := api_client.KerberosProvider.RegisterKerberosProvider(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KerberosProvider.RegisterKerberosProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterKerberosProvider`: KerberosProvider
    fmt.Fprintf(os.Stdout, "Response from `KerberosProvider.RegisterKerberosProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterKerberosProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RegisterOrUpdateKerberosProviderRequest**](RegisterOrUpdateKerberosProviderRequest.md) | Specifies the parameters to Register a Kerberos Provider. | 

### Return type

[**KerberosProvider**](KerberosProvider.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnregisterKerberosProvider

> UnregisterKerberosProvider UnregisterKerberosProvider(ctx, id).Body(body).Execute()

Unregister a Kerberos Provider.



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
    id := "id_example" // string | Specifies the id.
    body := *openapiclient.NewUnregisterKerberosRequest("AdminPassword_example") // UnregisterKerberosRequest | Request to unregister a Kerberos Provider.

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

    request := onprem.ApiUnregisterKerberosProviderRequest{
        Id: &id
        Body: &body
    }

    resp, r, err := api_client.KerberosProvider.UnregisterKerberosProvider(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KerberosProvider.UnregisterKerberosProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnregisterKerberosProvider`: UnregisterKerberosProvider
    fmt.Fprintf(os.Stdout, "Response from `KerberosProvider.UnregisterKerberosProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnregisterKerberosProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UnregisterKerberosRequest**](UnregisterKerberosRequest.md) | Request to unregister a Kerberos Provider. | 

### Return type

[**UnregisterKerberosProvider**](UnregisterKerberosProvider.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKerberosProvider

> KerberosProvider UpdateKerberosProvider(ctx, id).Body(body).Execute()

Update the Kerberos Provider Registration.



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
    id := "id_example" // string | Specifies the id which will be of the pattern cluster_id:clusterincarnation_id:resource_id.
    body := *openapiclient.NewRegisterOrUpdateKerberosProviderRequest("RealmName_example", []string{"KdcServers_example"}, "AdminServer_example", []string{"HostAlias_example"}, "AdminPassword_example") // RegisterOrUpdateKerberosProviderRequest | Request to update a Kerberos Provider.

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

    request := onprem.ApiUpdateKerberosProviderRequest{
        Id: &id
        Body: &body
    }

    resp, r, err := api_client.KerberosProvider.UpdateKerberosProvider(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KerberosProvider.UpdateKerberosProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKerberosProvider`: KerberosProvider
    fmt.Fprintf(os.Stdout, "Response from `KerberosProvider.UpdateKerberosProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id which will be of the pattern cluster_id:clusterincarnation_id:resource_id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKerberosProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RegisterOrUpdateKerberosProviderRequest**](RegisterOrUpdateKerberosProviderRequest.md) | Request to update a Kerberos Provider. | 

### Return type

[**KerberosProvider**](KerberosProvider.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

