# \View

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateShare**](View.md#CreateShare) | **Post** /file-services/shares | Create a Share.
[**CreateView**](View.md#CreateView) | **Post** /file-services/views | Create a View
[**CreateViewTemplate**](View.md#CreateViewTemplate) | **Post** /file-services/view-template | Create a View Template
[**DeleteShare**](View.md#DeleteShare) | **Delete** /file-services/shares/{name} | Delete a Share.
[**DeleteView**](View.md#DeleteView) | **Delete** /file-services/views/{id} | Delete a View
[**DeleteViewDirectoryQuota**](View.md#DeleteViewDirectoryQuota) | **Delete** /file-services/views/{id}/directory-quotas | Delete directory quota for the View.
[**DeleteViewTemplate**](View.md#DeleteViewTemplate) | **Delete** /file-services/view-template/{id} | Delete a View Template
[**GetShareByName**](View.md#GetShareByName) | **Get** /file-services/shares/{name} | Get a Share by name.
[**GetShares**](View.md#GetShares) | **Get** /file-services/shares | Get Shares.
[**GetViewById**](View.md#GetViewById) | **Get** /file-services/views/{id} | Get a View by Id
[**GetViewDirectoryQuotas**](View.md#GetViewDirectoryQuotas) | **Get** /file-services/views/{id}/directory-quotas | Get directory quotas for the View.
[**GetViewUserQuotas**](View.md#GetViewUserQuotas) | **Get** /file-services/views/{id}/user-quotas | Get user quotas for the View.
[**GetViewUserQuotasConfig**](View.md#GetViewUserQuotasConfig) | **Get** /file-services/views/{id}/user-quotas-config | Get View user quotas config.
[**GetViews**](View.md#GetViews) | **Get** /file-services/views | List Views
[**GetViewsSummary**](View.md#GetViewsSummary) | **Get** /file-services/views-summary | Get Views summary.
[**ReadViewTemplateById**](View.md#ReadViewTemplateById) | **Get** /file-services/view-template/{id} | Read a View Template by Id
[**ReadViewTemplates**](View.md#ReadViewTemplates) | **Get** /file-services/view-template | List View Templates
[**UpdateShare**](View.md#UpdateShare) | **Put** /file-services/shares/{name} | Update a Share.
[**UpdateView**](View.md#UpdateView) | **Put** /file-services/views/{id} | Update a View
[**UpdateViewDirectoryQuota**](View.md#UpdateViewDirectoryQuota) | **Put** /file-services/views/{id}/directory-quotas | Upadte directory quota for the View.
[**UpdateViewTemplate**](View.md#UpdateViewTemplate) | **Put** /file-services/view-template/{id} | Update a View Template



## CreateShare

> Share CreateShare(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Create a Share.



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
    body := *openapiclient.NewShare("Name_example", "ViewName_example", "ViewPath_example") // Share | Specifies the request to create a Share.
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

    request := helios.ApiCreateShareRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.View.CreateShare(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `View.CreateShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateShare`: Share
    fmt.Fprintf(os.Stdout, "Response from `View.CreateShare`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Share**](Share.md) | Specifies the request to create a Share. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**Share**](Share.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateView

> View CreateView(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Create a View



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
    body := *openapiclient.NewCreateViewRequest() // CreateViewRequest | Request to create a View.
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

    request := helios.ApiCreateViewRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.View.CreateView(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `View.CreateView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateView`: View
    fmt.Fprintf(os.Stdout, "Response from `View.CreateView`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateViewRequest**](CreateViewRequest.md) | Request to create a View. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**View**](View.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateViewTemplate

> Template CreateViewTemplate(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Create a View Template



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
    body := *openapiclient.NewTemplate() // Template | Request to create a view template.
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

    request := helios.ApiCreateViewTemplateRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.View.CreateViewTemplate(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `View.CreateViewTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateViewTemplate`: Template
    fmt.Fprintf(os.Stdout, "Response from `View.CreateViewTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateViewTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Template**](Template.md) | Request to create a view template. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**Template**](Template.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteShare

> DeleteShare(ctx, name).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Delete a Share.



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
    name := "name_example" // string | Specifies the Share name to delete.
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

    request := helios.ApiDeleteShareRequest{
        Name: &name
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.View.DeleteShare(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `View.DeleteShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Specifies the Share name to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteShareRequest struct via the builder pattern


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


## DeleteView

> DeleteView(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Delete a View



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
    id := int64(789) // int64 | Specifies a unique id of the View to delete.
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

    request := helios.ApiDeleteViewRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.View.DeleteView(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `View.DeleteView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the View to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteViewRequest struct via the builder pattern


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


## DeleteViewDirectoryQuota

> DeleteViewDirectoryQuota(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).DirectoryPath(directoryPath).DeleteAllDirectoryQuotas(deleteAllDirectoryQuotas).Execute()

Delete directory quota for the View.



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
    id := int64(789) // int64 | Specifies the View id.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    directoryPath := "directoryPath_example" // string | Specifies the directory path to delete. Exactly one of 'directoryPath' and 'deleteAllDirectoryQuotas' should be provided. (optional)
    deleteAllDirectoryQuotas := true // bool | Specifies whether to delete all directory quotas for this view. Exactly one of 'directoryPath' and 'deleteAllDirectoryQuotas' should be provided. (optional)

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

    request := helios.ApiDeleteViewDirectoryQuotaRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        DirectoryPath: &directoryPath
        DeleteAllDirectoryQuotas: &deleteAllDirectoryQuotas
    }

    resp, r, err := api_client.View.DeleteViewDirectoryQuota(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `View.DeleteViewDirectoryQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the View id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteViewDirectoryQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **directoryPath** | **string** | Specifies the directory path to delete. Exactly one of &#39;directoryPath&#39; and &#39;deleteAllDirectoryQuotas&#39; should be provided. | 
 **deleteAllDirectoryQuotas** | **bool** | Specifies whether to delete all directory quotas for this view. Exactly one of &#39;directoryPath&#39; and &#39;deleteAllDirectoryQuotas&#39; should be provided. | 

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


## DeleteViewTemplate

> DeleteViewTemplate(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Delete a View Template



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
    id := int64(789) // int64 | Specifies a unique id of the view template to delete.
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

    request := helios.ApiDeleteViewTemplateRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.View.DeleteViewTemplate(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `View.DeleteViewTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the view template to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteViewTemplateRequest struct via the builder pattern


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


## GetShareByName

> Share GetShareByName(ctx, name).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Get a Share by name.



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
    name := "name_example" // string | Specifies the Share name.
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

    request := helios.ApiGetShareByNameRequest{
        Name: &name
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.View.GetShareByName(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `View.GetShareByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShareByName`: Share
    fmt.Fprintf(os.Stdout, "Response from `View.GetShareByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Specifies the Share name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShareByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**Share**](Share.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShares

> Shares GetShares(ctx).AccessClusterId(accessClusterId).RegionId(regionId).Name(name).MatchPartialName(matchPartialName).MaxCount(maxCount).Cookie(cookie).TenantIds(tenantIds).IncludeTenants(includeTenants).Execute()

Get Shares.



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
    name := "name_example" // string | Specifies the Share name. (optional)
    matchPartialName := true // bool | If true, the share nane is matched by any partial rather than exactly matched. (optional)
    maxCount := int32(56) // int32 | Specifies a limit on the number of Views returned. (optional)
    cookie := "cookie_example" // string | Specifies the cookie. (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
    includeTenants := true // bool | IncludeTenants specifies if objects of all the tenants under the hierarchy of the logged in user's organization should be returned. (optional)

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

    request := helios.ApiGetSharesRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        Name: &name
        MatchPartialName: &matchPartialName
        MaxCount: &maxCount
        Cookie: &cookie
        TenantIds: &tenantIds
        IncludeTenants: &includeTenants
    }

    resp, r, err := api_client.View.GetShares(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `View.GetShares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShares`: Shares
    fmt.Fprintf(os.Stdout, "Response from `View.GetShares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **name** | **string** | Specifies the Share name. | 
 **matchPartialName** | **bool** | If true, the share nane is matched by any partial rather than exactly matched. | 
 **maxCount** | **int32** | Specifies a limit on the number of Views returned. | 
 **cookie** | **string** | Specifies the cookie. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **includeTenants** | **bool** | IncludeTenants specifies if objects of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. | 

### Return type

[**Shares**](Shares.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViewById

> View GetViewById(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Get a View by Id



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
    id := int64(789) // int64 | Specifies a unique id of the View to fetch.
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

    request := helios.ApiGetViewByIdRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.View.GetViewById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `View.GetViewById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetViewById`: View
    fmt.Fprintf(os.Stdout, "Response from `View.GetViewById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the View to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetViewByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**View**](View.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViewDirectoryQuotas

> ViewDirectoryQuotas GetViewDirectoryQuotas(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).MaxCount(maxCount).Cookie(cookie).Execute()

Get directory quotas for the View.



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
    id := int64(789) // int64 | Specifies the View id.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    maxCount := int64(789) // int64 | Specifies a limit on the number of quotas returned. (optional)
    cookie := int64(789) // int64 | Specifies the cookie. (optional)

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

    request := helios.ApiGetViewDirectoryQuotasRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        MaxCount: &maxCount
        Cookie: &cookie
    }

    resp, r, err := api_client.View.GetViewDirectoryQuotas(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `View.GetViewDirectoryQuotas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetViewDirectoryQuotas`: ViewDirectoryQuotas
    fmt.Fprintf(os.Stdout, "Response from `View.GetViewDirectoryQuotas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the View id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetViewDirectoryQuotasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **maxCount** | **int64** | Specifies a limit on the number of quotas returned. | 
 **cookie** | **int64** | Specifies the cookie. | 

### Return type

[**ViewDirectoryQuotas**](ViewDirectoryQuotas.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViewUserQuotas

> ViewUserQuotas GetViewUserQuotas(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).MaxCount(maxCount).Cookie(cookie).Execute()

Get user quotas for the View.



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
    id := int64(789) // int64 | Specifies the View id.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    maxCount := int64(789) // int64 | Specifies a limit on the number of quotas returned. (optional)
    cookie := "cookie_example" // string | Specifies the cookie. (optional)

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

    request := helios.ApiGetViewUserQuotasRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        MaxCount: &maxCount
        Cookie: &cookie
    }

    resp, r, err := api_client.View.GetViewUserQuotas(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `View.GetViewUserQuotas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetViewUserQuotas`: ViewUserQuotas
    fmt.Fprintf(os.Stdout, "Response from `View.GetViewUserQuotas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the View id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetViewUserQuotasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **maxCount** | **int64** | Specifies a limit on the number of quotas returned. | 
 **cookie** | **string** | Specifies the cookie. | 

### Return type

[**ViewUserQuotas**](ViewUserQuotas.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViewUserQuotasConfig

> ViewUserQuotasConfig GetViewUserQuotasConfig(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Get View user quotas config.



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
    id := int64(789) // int64 | Specifies the View id.
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

    request := helios.ApiGetViewUserQuotasConfigRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.View.GetViewUserQuotasConfig(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `View.GetViewUserQuotasConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetViewUserQuotasConfig`: ViewUserQuotasConfig
    fmt.Fprintf(os.Stdout, "Response from `View.GetViewUserQuotasConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the View id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetViewUserQuotasConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**ViewUserQuotasConfig**](ViewUserQuotasConfig.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViews

> GetViewsResult GetViews(ctx).AccessClusterId(accessClusterId).RegionId(regionId).ViewNames(viewNames).ViewIds(viewIds).StorageDomainIds(storageDomainIds).StorageDomainNames(storageDomainNames).ProtocolAccesses(protocolAccesses).MatchPartialNames(matchPartialNames).MaxCount(maxCount).IncludeInternalViews(includeInternalViews).IncludeProtectionGroups(includeProtectionGroups).MaxViewId(maxViewId).IncludeInactive(includeInactive).ProtectionGroupIds(protectionGroupIds).ViewProtectionGroupIds(viewProtectionGroupIds).ViewCountOnly(viewCountOnly).SummaryOnly(summaryOnly).SortByLogicalUsage(sortByLogicalUsage).InternalAccessSids(internalAccessSids).MatchAliasNames(matchAliasNames).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeStats(includeStats).IncludeFileCountBySize(includeFileCountBySize).IncludeViewsWithAntivirusEnabledOnly(includeViewsWithAntivirusEnabledOnly).IncludeViewsWithDataLockEnabledOnly(includeViewsWithDataLockEnabledOnly).FilerAuditLogEnabled(filerAuditLogEnabled).Categories(categories).ViewProtectionTypes(viewProtectionTypes).LastRunLocalBackupStatuses(lastRunLocalBackupStatuses).LastRunReplicationStatuses(lastRunReplicationStatuses).LastRunArchivalStatuses(lastRunArchivalStatuses).Execute()

List Views



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
    viewNames := []string{"Inner_example"} // []string | Filter by a list of View names. (optional)
    viewIds := []int64{int64(123)} // []int64 | Filter by a list of View ids. (optional)
    storageDomainIds := []int64{int64(123)} // []int64 | Filter by a list of Storage Domains (View Boxes) specified by id. (optional)
    storageDomainNames := []string{"Inner_example"} // []string | Filter by a list of View Box names. (optional)
    protocolAccesses := []string{"ProtocolAccesses_example"} // []string | Filter by a list of protocol accesses. Only views with protocol accesses in these specified accesses list will be returned. (optional)
    matchPartialNames := true // bool | If true, the names in viewNames are matched by any partial rather than exactly matched. (optional)
    maxCount := int32(56) // int32 | Specifies a limit on the number of Views returned. (optional)
    includeInternalViews := true // bool | Specifies if internal Views created by the Cohesity Cluster are also returned. In addition, regular Views are returned. (optional)
    includeProtectionGroups := true // bool | Specifies if Protection Groups information needs to be returned along with view metadata. By default, if not set or set to true, Group information is returned. (optional)
    maxViewId := int64(789) // int64 | If the number of Views to return exceeds the maxCount specified in the original request, specify the id of the last View from the viewList in the previous response to get the next set of Views. (optional)
    includeInactive := true // bool | Specifies if inactive Views on this Remote Cluster (which have Snapshots copied by replication) should also be returned. Inactive Views are not counted towards the maxCount. By default, this field is set to false. (optional)
    protectionGroupIds := []int64{int64(123)} // []int64 | This field will be deprecated. Filter by Protection Group ids. Return Views that are being protected by listed Groups, which are specified by ids. If both protectionGroupIds and viewProtectionGroupIds are specified, only viewProtectionGroupIds will be used. (optional)
    viewProtectionGroupIds := []string{"Inner_example"} // []string | Filter by Protection Group ids. Return Views that are being protected by listed Groups, which are specified by ids. (optional)
    viewCountOnly := true // bool | Whether to get just the total number of views with the given input filters. If the flag is true, we ignore the parameter 'maxViews' for the count. Also, if flag is true, list of views will not be returned. (optional)
    summaryOnly := true // bool | Whether to get only view summary including 'name', 'viewId', 'storageDomainName', 'storageDomainId' and 'tenantId'. (optional)
    sortByLogicalUsage := true // bool | If set to true, the list is sorted descending by logical usage. (optional)
    internalAccessSids := []string{"Inner_example"} // []string | Sids of restricted principals who can access the view. This is an internal field and therefore does not have json tag. (optional)
    matchAliasNames := true // bool | If true, view aliases are also matched with the names in viewNames. (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
    includeTenants := true // bool | IncludeTenants specifies if objects of all the tenants under the hierarchy of the logged in user's organization should be returned. (optional)
    includeStats := true // bool | If set to true, stats of views will be returned. By default this parameter is set to false. (optional)
    includeFileCountBySize := true // bool | Whether to include View file count by size. (optional)
    includeViewsWithAntivirusEnabledOnly := true // bool | If set to true, the list will contain only the views for which antivirus scan is enabled. (optional)
    includeViewsWithDataLockEnabledOnly := true // bool | If set to true, the list will contain only the views for which either file level data lock is enabled or view level data lock is enabled. (optional)
    filerAuditLogEnabled := true // bool | If set to true, only views with filer audit log enabled will be returned. If set to false, only views with filer audit log disabled will be returned. (optional)
    categories := []string{"Categories_example"} // []string | Filter by a list of View categories. (optional)
    viewProtectionTypes := []string{"ViewProtectionTypes_example"} // []string | Filter by a list of View protection types. (optional)
    lastRunLocalBackupStatuses := []string{"LastRunLocalBackupStatuses_example"} // []string | Filter by last local backup run status of the view.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. (optional)
    lastRunReplicationStatuses := []string{"LastRunReplicationStatuses_example"} // []string | Filter by last remote replication run status of the view.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. (optional)
    lastRunArchivalStatuses := []string{"LastRunArchivalStatuses_example"} // []string | Filter by last cloud archival run status of the view.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. (optional)

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

    request := helios.ApiGetViewsRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        ViewNames: &viewNames
        ViewIds: &viewIds
        StorageDomainIds: &storageDomainIds
        StorageDomainNames: &storageDomainNames
        ProtocolAccesses: &protocolAccesses
        MatchPartialNames: &matchPartialNames
        MaxCount: &maxCount
        IncludeInternalViews: &includeInternalViews
        IncludeProtectionGroups: &includeProtectionGroups
        MaxViewId: &maxViewId
        IncludeInactive: &includeInactive
        ProtectionGroupIds: &protectionGroupIds
        ViewProtectionGroupIds: &viewProtectionGroupIds
        ViewCountOnly: &viewCountOnly
        SummaryOnly: &summaryOnly
        SortByLogicalUsage: &sortByLogicalUsage
        InternalAccessSids: &internalAccessSids
        MatchAliasNames: &matchAliasNames
        TenantIds: &tenantIds
        IncludeTenants: &includeTenants
        IncludeStats: &includeStats
        IncludeFileCountBySize: &includeFileCountBySize
        IncludeViewsWithAntivirusEnabledOnly: &includeViewsWithAntivirusEnabledOnly
        IncludeViewsWithDataLockEnabledOnly: &includeViewsWithDataLockEnabledOnly
        FilerAuditLogEnabled: &filerAuditLogEnabled
        Categories: &categories
        ViewProtectionTypes: &viewProtectionTypes
        LastRunLocalBackupStatuses: &lastRunLocalBackupStatuses
        LastRunReplicationStatuses: &lastRunReplicationStatuses
        LastRunArchivalStatuses: &lastRunArchivalStatuses
    }

    resp, r, err := api_client.View.GetViews(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `View.GetViews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetViews`: GetViewsResult
    fmt.Fprintf(os.Stdout, "Response from `View.GetViews`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **viewNames** | **[]string** | Filter by a list of View names. | 
 **viewIds** | **[]int64** | Filter by a list of View ids. | 
 **storageDomainIds** | **[]int64** | Filter by a list of Storage Domains (View Boxes) specified by id. | 
 **storageDomainNames** | **[]string** | Filter by a list of View Box names. | 
 **protocolAccesses** | **[]string** | Filter by a list of protocol accesses. Only views with protocol accesses in these specified accesses list will be returned. | 
 **matchPartialNames** | **bool** | If true, the names in viewNames are matched by any partial rather than exactly matched. | 
 **maxCount** | **int32** | Specifies a limit on the number of Views returned. | 
 **includeInternalViews** | **bool** | Specifies if internal Views created by the Cohesity Cluster are also returned. In addition, regular Views are returned. | 
 **includeProtectionGroups** | **bool** | Specifies if Protection Groups information needs to be returned along with view metadata. By default, if not set or set to true, Group information is returned. | 
 **maxViewId** | **int64** | If the number of Views to return exceeds the maxCount specified in the original request, specify the id of the last View from the viewList in the previous response to get the next set of Views. | 
 **includeInactive** | **bool** | Specifies if inactive Views on this Remote Cluster (which have Snapshots copied by replication) should also be returned. Inactive Views are not counted towards the maxCount. By default, this field is set to false. | 
 **protectionGroupIds** | **[]int64** | This field will be deprecated. Filter by Protection Group ids. Return Views that are being protected by listed Groups, which are specified by ids. If both protectionGroupIds and viewProtectionGroupIds are specified, only viewProtectionGroupIds will be used. | 
 **viewProtectionGroupIds** | **[]string** | Filter by Protection Group ids. Return Views that are being protected by listed Groups, which are specified by ids. | 
 **viewCountOnly** | **bool** | Whether to get just the total number of views with the given input filters. If the flag is true, we ignore the parameter &#39;maxViews&#39; for the count. Also, if flag is true, list of views will not be returned. | 
 **summaryOnly** | **bool** | Whether to get only view summary including &#39;name&#39;, &#39;viewId&#39;, &#39;storageDomainName&#39;, &#39;storageDomainId&#39; and &#39;tenantId&#39;. | 
 **sortByLogicalUsage** | **bool** | If set to true, the list is sorted descending by logical usage. | 
 **internalAccessSids** | **[]string** | Sids of restricted principals who can access the view. This is an internal field and therefore does not have json tag. | 
 **matchAliasNames** | **bool** | If true, view aliases are also matched with the names in viewNames. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **includeTenants** | **bool** | IncludeTenants specifies if objects of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. | 
 **includeStats** | **bool** | If set to true, stats of views will be returned. By default this parameter is set to false. | 
 **includeFileCountBySize** | **bool** | Whether to include View file count by size. | 
 **includeViewsWithAntivirusEnabledOnly** | **bool** | If set to true, the list will contain only the views for which antivirus scan is enabled. | 
 **includeViewsWithDataLockEnabledOnly** | **bool** | If set to true, the list will contain only the views for which either file level data lock is enabled or view level data lock is enabled. | 
 **filerAuditLogEnabled** | **bool** | If set to true, only views with filer audit log enabled will be returned. If set to false, only views with filer audit log disabled will be returned. | 
 **categories** | **[]string** | Filter by a list of View categories. | 
 **viewProtectionTypes** | **[]string** | Filter by a list of View protection types. | 
 **lastRunLocalBackupStatuses** | **[]string** | Filter by last local backup run status of the view.&lt;br&gt; &#39;Running&#39; indicates that the run is still running.&lt;br&gt; &#39;Canceled&#39; indicates that the run has been canceled.&lt;br&gt; &#39;Canceling&#39; indicates that the run is in the process of being canceled.&lt;br&gt; &#39;Failed&#39; indicates that the run has failed.&lt;br&gt; &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening.&lt;br&gt; &#39;Succeeded&#39; indicates that the run has finished successfully.&lt;br&gt; &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages. | 
 **lastRunReplicationStatuses** | **[]string** | Filter by last remote replication run status of the view.&lt;br&gt; &#39;Running&#39; indicates that the run is still running.&lt;br&gt; &#39;Canceled&#39; indicates that the run has been canceled.&lt;br&gt; &#39;Canceling&#39; indicates that the run is in the process of being canceled.&lt;br&gt; &#39;Failed&#39; indicates that the run has failed.&lt;br&gt; &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening.&lt;br&gt; &#39;Succeeded&#39; indicates that the run has finished successfully.&lt;br&gt; &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages. | 
 **lastRunArchivalStatuses** | **[]string** | Filter by last cloud archival run status of the view.&lt;br&gt; &#39;Running&#39; indicates that the run is still running.&lt;br&gt; &#39;Canceled&#39; indicates that the run has been canceled.&lt;br&gt; &#39;Canceling&#39; indicates that the run is in the process of being canceled.&lt;br&gt; &#39;Failed&#39; indicates that the run has failed.&lt;br&gt; &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening.&lt;br&gt; &#39;Succeeded&#39; indicates that the run has finished successfully.&lt;br&gt; &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages. | 

### Return type

[**GetViewsResult**](GetViewsResult.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViewsSummary

> ViewsSummary GetViewsSummary(ctx).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Get Views summary.



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

    request := helios.ApiGetViewsSummaryRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.View.GetViewsSummary(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `View.GetViewsSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetViewsSummary`: ViewsSummary
    fmt.Fprintf(os.Stdout, "Response from `View.GetViewsSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetViewsSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**ViewsSummary**](ViewsSummary.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadViewTemplateById

> Template ReadViewTemplateById(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Read a View Template by Id



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
    id := int64(789) // int64 | Specifies a unique id of the view template.
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

    request := helios.ApiReadViewTemplateByIdRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.View.ReadViewTemplateById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `View.ReadViewTemplateById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadViewTemplateById`: Template
    fmt.Fprintf(os.Stdout, "Response from `View.ReadViewTemplateById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the view template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadViewTemplateByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**Template**](Template.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadViewTemplates

> GetViewTemplatesResult ReadViewTemplates(ctx).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

List View Templates



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

    request := helios.ApiReadViewTemplatesRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.View.ReadViewTemplates(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `View.ReadViewTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadViewTemplates`: GetViewTemplatesResult
    fmt.Fprintf(os.Stdout, "Response from `View.ReadViewTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadViewTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**GetViewTemplatesResult**](GetViewTemplatesResult.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateShare

> Share UpdateShare(ctx, name).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Update a Share.



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
    name := "name_example" // string | Specifies the Share name to update.
    body := *openapiclient.NewUpdateShareParam() // UpdateShareParam | Specifies the request to update a Share.
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

    request := helios.ApiUpdateShareRequest{
        Name: &name
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.View.UpdateShare(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `View.UpdateShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateShare`: Share
    fmt.Fprintf(os.Stdout, "Response from `View.UpdateShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Specifies the Share name to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateShareParam**](UpdateShareParam.md) | Specifies the request to update a Share. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**Share**](Share.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateView

> View UpdateView(ctx, id).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Update a View



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
    id := int64(789) // int64 | Specifies a unique id of the View to update.
    body := *openapiclient.NewView() // View | Request to update a view.
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

    request := helios.ApiUpdateViewRequest{
        Id: &id
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.View.UpdateView(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `View.UpdateView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateView`: View
    fmt.Fprintf(os.Stdout, "Response from `View.UpdateView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the View to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**View**](View.md) | Request to update a view. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**View**](View.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateViewDirectoryQuota

> ViewDirectoryQuota UpdateViewDirectoryQuota(ctx, id).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Upadte directory quota for the View.



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
    id := int64(789) // int64 | Specifies the View id.
    body := *openapiclient.NewViewDirectoryQuota("DirectoryPath_example", *openapiclient.NewViewDirectoryQuotaPolicy()) // ViewDirectoryQuota | Specifies the request to update directory quota.
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

    request := helios.ApiUpdateViewDirectoryQuotaRequest{
        Id: &id
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.View.UpdateViewDirectoryQuota(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `View.UpdateViewDirectoryQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateViewDirectoryQuota`: ViewDirectoryQuota
    fmt.Fprintf(os.Stdout, "Response from `View.UpdateViewDirectoryQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the View id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateViewDirectoryQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ViewDirectoryQuota**](ViewDirectoryQuota.md) | Specifies the request to update directory quota. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**ViewDirectoryQuota**](ViewDirectoryQuota.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateViewTemplate

> Template UpdateViewTemplate(ctx, id).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Update a View Template



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
    id := int64(789) // int64 | Specifies a unique id of the view template.
    body := *openapiclient.NewTemplate() // Template | Request to update a view template.
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

    request := helios.ApiUpdateViewTemplateRequest{
        Id: &id
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.View.UpdateViewTemplate(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `View.UpdateViewTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateViewTemplate`: Template
    fmt.Fprintf(os.Stdout, "Response from `View.UpdateViewTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the view template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateViewTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Template**](Template.md) | Request to update a view template. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**Template**](Template.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

