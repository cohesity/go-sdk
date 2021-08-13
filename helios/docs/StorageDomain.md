# \StorageDomain

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStorageDomain**](StorageDomain.md#CreateStorageDomain) | **Post** /storage-domains | Create a Storage Domain.
[**DeleteStorageDomain**](StorageDomain.md#DeleteStorageDomain) | **Delete** /storage-domains/{id} | Delete a Storage Domain.
[**GetStorageDomainById**](StorageDomain.md#GetStorageDomainById) | **Get** /storage-domains/{id} | Get a Storage Domain by id.
[**GetStorageDomains**](StorageDomain.md#GetStorageDomains) | **Get** /storage-domains | Get Storage Domains.
[**UpdateStorageDomain**](StorageDomain.md#UpdateStorageDomain) | **Put** /storage-domains/{id} | Update a Storage Domain.



## CreateStorageDomain

> StorageDomain CreateStorageDomain(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Create a Storage Domain.



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
    body := *openapiclient.NewCreateStorageDomainParam("Name_example", NullableInt64(123)) // CreateStorageDomainParam | Specified the request to create a Storage Domain.
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

    request := helios.ApiCreateStorageDomainRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.StorageDomain.CreateStorageDomain(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageDomain.CreateStorageDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStorageDomain`: StorageDomain
    fmt.Fprintf(os.Stdout, "Response from `StorageDomain.CreateStorageDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorageDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateStorageDomainParam**](CreateStorageDomainParam.md) | Specified the request to create a Storage Domain. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**StorageDomain**](StorageDomain.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStorageDomain

> DeleteStorageDomain(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Delete a Storage Domain.



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
    id := int64(789) // int64 | Specified the Storage Domain id to delete.
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

    request := helios.ApiDeleteStorageDomainRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.StorageDomain.DeleteStorageDomain(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageDomain.DeleteStorageDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specified the Storage Domain id to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStorageDomainRequest struct via the builder pattern


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


## GetStorageDomainById

> StorageDomain GetStorageDomainById(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).IncludeStats(includeStats).IncludeTimeSeriesSchema(includeTimeSeriesSchema).IncludeFileCountBySize(includeFileCountBySize).Execute()

Get a Storage Domain by id.



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
    id := int64(789) // int64 | Specified the Storage Domain id to fetch.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    includeStats := true // bool | Whether to include Storage Domain stats in response. (optional)
    includeTimeSeriesSchema := true // bool | Whether to include Storage Domain time series schema in response. (optional)
    includeFileCountBySize := true // bool | Whether to include Storage Domain file count by size. (optional)

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

    request := helios.ApiGetStorageDomainByIdRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        IncludeStats: &includeStats
        IncludeTimeSeriesSchema: &includeTimeSeriesSchema
        IncludeFileCountBySize: &includeFileCountBySize
    }

    resp, r, err := api_client.StorageDomain.GetStorageDomainById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageDomain.GetStorageDomainById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorageDomainById`: StorageDomain
    fmt.Fprintf(os.Stdout, "Response from `StorageDomain.GetStorageDomainById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specified the Storage Domain id to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageDomainByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **includeStats** | **bool** | Whether to include Storage Domain stats in response. | 
 **includeTimeSeriesSchema** | **bool** | Whether to include Storage Domain time series schema in response. | 
 **includeFileCountBySize** | **bool** | Whether to include Storage Domain file count by size. | 

### Return type

[**StorageDomain**](StorageDomain.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorageDomains

> StorageDomains GetStorageDomains(ctx).AccessClusterId(accessClusterId).RegionId(regionId).Ids(ids).Names(names).ClusterPartitionIds(clusterPartitionIds).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeStats(includeStats).IncludeTimeSeriesSchema(includeTimeSeriesSchema).IncludeFileCountBySize(includeFileCountBySize).MatchPartialNames(matchPartialNames).ViewTemplateId(viewTemplateId).Execute()

Get Storage Domains.



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
    ids := []int64{int64(123)} // []int64 | Filter by a list of Storage Domain ids. (optional)
    names := []string{"Inner_example"} // []string | Filter by a list of Storage Domain names. (optional)
    clusterPartitionIds := []int64{int64(123)} // []int64 | Filter by a list of cluster partition ids. (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which Storage Domains are to be returned. (optional)
    includeTenants := true // bool | IncludeTenants specifies if Storage Domains of all the tenants under the hierarchy of the logged in user's organization should be returned. (optional)
    includeStats := true // bool | Whether to include Storage Domain stats in response. (optional)
    includeTimeSeriesSchema := true // bool | Whether to include Storage Domain time series schema in response. (optional)
    includeFileCountBySize := true // bool | Whether to include Storage Domain file count by size. (optional)
    matchPartialNames := true // bool | If true, the names in viewNames are matched by any partial rather than exactly matched. (optional)
    viewTemplateId := int64(789) // int64 | Specifies a view template id for Storage Domain. Storage Domains with same deduplication and compression settings will be recommended. (optional)

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

    request := helios.ApiGetStorageDomainsRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        Ids: &ids
        Names: &names
        ClusterPartitionIds: &clusterPartitionIds
        TenantIds: &tenantIds
        IncludeTenants: &includeTenants
        IncludeStats: &includeStats
        IncludeTimeSeriesSchema: &includeTimeSeriesSchema
        IncludeFileCountBySize: &includeFileCountBySize
        MatchPartialNames: &matchPartialNames
        ViewTemplateId: &viewTemplateId
    }

    resp, r, err := api_client.StorageDomain.GetStorageDomains(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageDomain.GetStorageDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorageDomains`: StorageDomains
    fmt.Fprintf(os.Stdout, "Response from `StorageDomain.GetStorageDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **ids** | **[]int64** | Filter by a list of Storage Domain ids. | 
 **names** | **[]string** | Filter by a list of Storage Domain names. | 
 **clusterPartitionIds** | **[]int64** | Filter by a list of cluster partition ids. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which Storage Domains are to be returned. | 
 **includeTenants** | **bool** | IncludeTenants specifies if Storage Domains of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. | 
 **includeStats** | **bool** | Whether to include Storage Domain stats in response. | 
 **includeTimeSeriesSchema** | **bool** | Whether to include Storage Domain time series schema in response. | 
 **includeFileCountBySize** | **bool** | Whether to include Storage Domain file count by size. | 
 **matchPartialNames** | **bool** | If true, the names in viewNames are matched by any partial rather than exactly matched. | 
 **viewTemplateId** | **int64** | Specifies a view template id for Storage Domain. Storage Domains with same deduplication and compression settings will be recommended. | 

### Return type

[**StorageDomains**](StorageDomains.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStorageDomain

> StorageDomain UpdateStorageDomain(ctx, id).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Update a Storage Domain.



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
    id := int64(789) // int64 | Specified the Storage Domain id to update.
    body := *openapiclient.NewUpdateStorageDomainParam("Name_example", NullableInt64(123)) // UpdateStorageDomainParam | Specified the request to update a Storage Domain.
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

    request := helios.ApiUpdateStorageDomainRequest{
        Id: &id
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.StorageDomain.UpdateStorageDomain(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageDomain.UpdateStorageDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStorageDomain`: StorageDomain
    fmt.Fprintf(os.Stdout, "Response from `StorageDomain.UpdateStorageDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specified the Storage Domain id to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStorageDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateStorageDomainParam**](UpdateStorageDomainParam.md) | Specified the request to update a Storage Domain. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**StorageDomain**](StorageDomain.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

