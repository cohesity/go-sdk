# \StorageDomainAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStorageDomain**](StorageDomainAPI.md#CreateStorageDomain) | **Post** /storage-domains | Create a Storage Domain.
[**DeleteStorageDomain**](StorageDomainAPI.md#DeleteStorageDomain) | **Delete** /storage-domains/{id} | Delete a Storage Domain.
[**GetStorageDomainById**](StorageDomainAPI.md#GetStorageDomainById) | **Get** /storage-domains/{id} | Get a Storage Domain by id.
[**GetStorageDomains**](StorageDomainAPI.md#GetStorageDomains) | **Get** /storage-domains | Get Storage Domains.
[**UpdateStorageDomain**](StorageDomainAPI.md#UpdateStorageDomain) | **Put** /storage-domains/{id} | Update a Storage Domain.



## CreateStorageDomain

> StorageDomain CreateStorageDomain(ctx).Body(body).Execute()

Create a Storage Domain.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewCreateStorageDomainParam(NullableInt64(123), "Name_example") // CreateStorageDomainParam | Specified the request to create a Storage Domain.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageDomainAPI.CreateStorageDomain(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageDomainAPI.CreateStorageDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStorageDomain`: StorageDomain
	fmt.Fprintf(os.Stdout, "Response from `StorageDomainAPI.CreateStorageDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorageDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateStorageDomainParam**](CreateStorageDomainParam.md) | Specified the request to create a Storage Domain. | 

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

> DeleteStorageDomain(ctx, id).Execute()

Delete a Storage Domain.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := int64(789) // int64 | Specified the Storage Domain id to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StorageDomainAPI.DeleteStorageDomain(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageDomainAPI.DeleteStorageDomain``: %v\n", err)
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

> StorageDomain GetStorageDomainById(ctx, id).IncludeStats(includeStats).IncludeTimeSeriesSchema(includeTimeSeriesSchema).IncludeFileCountBySize(includeFileCountBySize).IncludeTenants(includeTenants).Execute()

Get a Storage Domain by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := int64(789) // int64 | Specified the Storage Domain id to fetch.
	includeStats := true // bool | Whether to include Storage Domain stats in response. (optional)
	includeTimeSeriesSchema := true // bool | Whether to include Storage Domain time series schema in response. (optional)
	includeFileCountBySize := true // bool | Whether to include Storage Domain file count by size. (optional)
	includeTenants := true // bool | Whether to include Storage Domains that belong to Tenants. This param is only effective when the User has privilege to view Storage Domain details of a tenant. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageDomainAPI.GetStorageDomainById(context.Background(), id).IncludeStats(includeStats).IncludeTimeSeriesSchema(includeTimeSeriesSchema).IncludeFileCountBySize(includeFileCountBySize).IncludeTenants(includeTenants).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageDomainAPI.GetStorageDomainById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStorageDomainById`: StorageDomain
	fmt.Fprintf(os.Stdout, "Response from `StorageDomainAPI.GetStorageDomainById`: %v\n", resp)
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

 **includeStats** | **bool** | Whether to include Storage Domain stats in response. | 
 **includeTimeSeriesSchema** | **bool** | Whether to include Storage Domain time series schema in response. | 
 **includeFileCountBySize** | **bool** | Whether to include Storage Domain file count by size. | 
 **includeTenants** | **bool** | Whether to include Storage Domains that belong to Tenants. This param is only effective when the User has privilege to view Storage Domain details of a tenant. | 

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

> StorageDomains GetStorageDomains(ctx).Ids(ids).Names(names).ClusterPartitionIds(clusterPartitionIds).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeStats(includeStats).IncludeTimeSeriesSchema(includeTimeSeriesSchema).IncludeFileCountBySize(includeFileCountBySize).MatchPartialNames(matchPartialNames).ViewTemplateId(viewTemplateId).Execute()

Get Storage Domains.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageDomainAPI.GetStorageDomains(context.Background()).Ids(ids).Names(names).ClusterPartitionIds(clusterPartitionIds).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeStats(includeStats).IncludeTimeSeriesSchema(includeTimeSeriesSchema).IncludeFileCountBySize(includeFileCountBySize).MatchPartialNames(matchPartialNames).ViewTemplateId(viewTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageDomainAPI.GetStorageDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStorageDomains`: StorageDomains
	fmt.Fprintf(os.Stdout, "Response from `StorageDomainAPI.GetStorageDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

> StorageDomain UpdateStorageDomain(ctx, id).Body(body).Execute()

Update a Storage Domain.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := int64(789) // int64 | Specified the Storage Domain id to update.
	body := *openapiclient.NewUpdateStorageDomainParam(NullableInt64(123), "Name_example") // UpdateStorageDomainParam | Specified the request to update a Storage Domain.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageDomainAPI.UpdateStorageDomain(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageDomainAPI.UpdateStorageDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStorageDomain`: StorageDomain
	fmt.Fprintf(os.Stdout, "Response from `StorageDomainAPI.UpdateStorageDomain`: %v\n", resp)
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

