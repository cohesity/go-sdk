# \NetworkInformationService

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNisNetgroup**](NetworkInformationService.md#CreateNisNetgroup) | **Post** /nis-netgroups | Create an NIS netgroup.
[**CreateNisProvider**](NetworkInformationService.md#CreateNisProvider) | **Post** /nis-providers | Create an NIS Provider.
[**DeleteNisNetgroupByName**](NetworkInformationService.md#DeleteNisNetgroupByName) | **Delete** /nis-netgroups/{name} | Delete an NIS netgroup by name.
[**DeleteNisProviderByDomainName**](NetworkInformationService.md#DeleteNisProviderByDomainName) | **Delete** /nis-providers/{domain} | Delete an NIS Provider by domain name.
[**GetNisNetgroupByName**](NetworkInformationService.md#GetNisNetgroupByName) | **Get** /nis-netgroups/{name} | Get an NIS netgroup by name.
[**GetNisNetgroups**](NetworkInformationService.md#GetNisNetgroups) | **Get** /nis-netgroups | Get a list of NIS netgroups.
[**GetNisProviderByDomainName**](NetworkInformationService.md#GetNisProviderByDomainName) | **Get** /nis-providers/{domain} | Get an NIS Provider by domain name.
[**GetNisProviders**](NetworkInformationService.md#GetNisProviders) | **Get** /nis-providers | Get a list of NIS Providers.
[**UpdateNisNetgroupByName**](NetworkInformationService.md#UpdateNisNetgroupByName) | **Put** /nis-netgroups/{name} | Update an NIS netgroup by name.
[**UpdateNisProviderByDomainName**](NetworkInformationService.md#UpdateNisProviderByDomainName) | **Put** /nis-providers/{domain} | Update an NIS Provider by domain name.



## CreateNisNetgroup

> NisNetgroup CreateNisNetgroup(ctx).Body(body).Execute()

Create an NIS netgroup.



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
    body := *openapiclient.NewNisNetgroup("Name_example", "Domain_example") // NisNetgroup | Specifies the parameters to create an NIS netgroup.

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

    request := onprem.ApiCreateNisNetgroupRequest{
        Body: &body
    }

    resp, r, err := api_client.NetworkInformationService.CreateNisNetgroup(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkInformationService.CreateNisNetgroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNisNetgroup`: NisNetgroup
    fmt.Fprintf(os.Stdout, "Response from `NetworkInformationService.CreateNisNetgroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNisNetgroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NisNetgroup**](NisNetgroup.md) | Specifies the parameters to create an NIS netgroup. | 

### Return type

[**NisNetgroup**](NisNetgroup.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNisProvider

> NisProvider CreateNisProvider(ctx).Body(body).Execute()

Create an NIS Provider.



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
    body := *openapiclient.NewNisProvider("Domain_example", "MasterServerHostname_example") // NisProvider | Specifies the parameters to create an NIS provider entry.

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

    request := onprem.ApiCreateNisProviderRequest{
        Body: &body
    }

    resp, r, err := api_client.NetworkInformationService.CreateNisProvider(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkInformationService.CreateNisProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNisProvider`: NisProvider
    fmt.Fprintf(os.Stdout, "Response from `NetworkInformationService.CreateNisProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNisProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NisProvider**](NisProvider.md) | Specifies the parameters to create an NIS provider entry. | 

### Return type

[**NisProvider**](NisProvider.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNisNetgroupByName

> DeleteNisNetgroupByName(ctx, name).Body(body).Execute()

Delete an NIS netgroup by name.



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
    name := "name_example" // string | Specifies name of the NIS netgroup.
    body := *openapiclient.NewNisNetgroup("Name_example", "Domain_example") // NisNetgroup | Request to delete the NIS netgroup.

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

    request := onprem.ApiDeleteNisNetgroupByNameRequest{
        Name: &name
        Body: &body
    }

    resp, r, err := api_client.NetworkInformationService.DeleteNisNetgroupByName(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkInformationService.DeleteNisNetgroupByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Specifies name of the NIS netgroup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNisNetgroupByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NisNetgroup**](NisNetgroup.md) | Request to delete the NIS netgroup. | 

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


## DeleteNisProviderByDomainName

> DeleteNisProviderByDomainName(ctx, domain).Execute()

Delete an NIS Provider by domain name.



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
    domain := "domain_example" // string | Specifies domain name of an NIS Provider.

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

    request := onprem.ApiDeleteNisProviderByDomainNameRequest{
        Domain: &domain
    }

    resp, r, err := api_client.NetworkInformationService.DeleteNisProviderByDomainName(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkInformationService.DeleteNisProviderByDomainName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Specifies domain name of an NIS Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNisProviderByDomainNameRequest struct via the builder pattern


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


## GetNisNetgroupByName

> NisNetgroup GetNisNetgroupByName(ctx, name).Execute()

Get an NIS netgroup by name.



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
    name := "name_example" // string | Specifies name of the NIS netgroup.

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

    request := onprem.ApiGetNisNetgroupByNameRequest{
        Name: &name
    }

    resp, r, err := api_client.NetworkInformationService.GetNisNetgroupByName(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkInformationService.GetNisNetgroupByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNisNetgroupByName`: NisNetgroup
    fmt.Fprintf(os.Stdout, "Response from `NetworkInformationService.GetNisNetgroupByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Specifies name of the NIS netgroup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNisNetgroupByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NisNetgroup**](NisNetgroup.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNisNetgroups

> NisNetgroups GetNisNetgroups(ctx).NetgroupNames(netgroupNames).Execute()

Get a list of NIS netgroups.



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
    netgroupNames := []string{"Inner_example"} // []string | Filter by a list of NIS netgroup names. (optional)

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

    request := onprem.ApiGetNisNetgroupsRequest{
        NetgroupNames: &netgroupNames
    }

    resp, r, err := api_client.NetworkInformationService.GetNisNetgroups(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkInformationService.GetNisNetgroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNisNetgroups`: NisNetgroups
    fmt.Fprintf(os.Stdout, "Response from `NetworkInformationService.GetNisNetgroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNisNetgroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **netgroupNames** | **[]string** | Filter by a list of NIS netgroup names. | 

### Return type

[**NisNetgroups**](NisNetgroups.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNisProviderByDomainName

> NisProvider GetNisProviderByDomainName(ctx, domain).Execute()

Get an NIS Provider by domain name.



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
    domain := "domain_example" // string | Specifies domain of an NIS Provider.

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

    request := onprem.ApiGetNisProviderByDomainNameRequest{
        Domain: &domain
    }

    resp, r, err := api_client.NetworkInformationService.GetNisProviderByDomainName(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkInformationService.GetNisProviderByDomainName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNisProviderByDomainName`: NisProvider
    fmt.Fprintf(os.Stdout, "Response from `NetworkInformationService.GetNisProviderByDomainName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Specifies domain of an NIS Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNisProviderByDomainNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NisProvider**](NisProvider.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNisProviders

> NisProviders GetNisProviders(ctx).DomainNames(domainNames).Execute()

Get a list of NIS Providers.



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
    domainNames := []string{"Inner_example"} // []string | Filter by a list of NIS domain names. (optional)

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

    request := onprem.ApiGetNisProvidersRequest{
        DomainNames: &domainNames
    }

    resp, r, err := api_client.NetworkInformationService.GetNisProviders(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkInformationService.GetNisProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNisProviders`: NisProviders
    fmt.Fprintf(os.Stdout, "Response from `NetworkInformationService.GetNisProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNisProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainNames** | **[]string** | Filter by a list of NIS domain names. | 

### Return type

[**NisProviders**](NisProviders.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNisNetgroupByName

> NisNetgroup UpdateNisNetgroupByName(ctx, name).Body(body).Execute()

Update an NIS netgroup by name.



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
    name := "name_example" // string | Specifies name of the NIS netgroup.
    body := *openapiclient.NewNisNetgroup("Name_example", "Domain_example") // NisNetgroup | Request to update the NIS netgroup.

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

    request := onprem.ApiUpdateNisNetgroupByNameRequest{
        Name: &name
        Body: &body
    }

    resp, r, err := api_client.NetworkInformationService.UpdateNisNetgroupByName(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkInformationService.UpdateNisNetgroupByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNisNetgroupByName`: NisNetgroup
    fmt.Fprintf(os.Stdout, "Response from `NetworkInformationService.UpdateNisNetgroupByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Specifies name of the NIS netgroup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNisNetgroupByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NisNetgroup**](NisNetgroup.md) | Request to update the NIS netgroup. | 

### Return type

[**NisNetgroup**](NisNetgroup.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNisProviderByDomainName

> NisProvider UpdateNisProviderByDomainName(ctx, domain).Body(body).Execute()

Update an NIS Provider by domain name.



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
    domain := "domain_example" // string | Specifies domain name of an NIS Provider.
    body := *openapiclient.NewNisProvider("Domain_example", "MasterServerHostname_example") // NisProvider | Request to update an NIS Provider.

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

    request := onprem.ApiUpdateNisProviderByDomainNameRequest{
        Domain: &domain
        Body: &body
    }

    resp, r, err := api_client.NetworkInformationService.UpdateNisProviderByDomainName(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkInformationService.UpdateNisProviderByDomainName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNisProviderByDomainName`: NisProvider
    fmt.Fprintf(os.Stdout, "Response from `NetworkInformationService.UpdateNisProviderByDomainName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Specifies domain name of an NIS Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNisProviderByDomainNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NisProvider**](NisProvider.md) | Request to update an NIS Provider. | 

### Return type

[**NisProvider**](NisProvider.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

