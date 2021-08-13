# \Tag

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTag**](Tag.md#CreateTag) | **Post** /tags | Create a Tag
[**DeleteTag**](Tag.md#DeleteTag) | **Delete** /tags/{id} | Delete a Tag
[**GetTagById**](Tag.md#GetTagById) | **Get** /tags/{id} | Get Tag by id.
[**GetTags**](Tag.md#GetTags) | **Get** /tags | Get tags based on filters.
[**UpdateTag**](Tag.md#UpdateTag) | **Put** /tags/{id} | Update a Tag



## CreateTag

> Tag CreateTag(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Create a Tag



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
    body := *openapiclient.NewTag("Name_example", "Namespace_example") // Tag | Request to create a Tag.
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

    request := helios.ApiCreateTagRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Tag.CreateTag(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tag.CreateTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTag`: Tag
    fmt.Fprintf(os.Stdout, "Response from `Tag.CreateTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Tag**](Tag.md) | Request to create a Tag. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**Tag**](Tag.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTag

> DeleteTag(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Delete a Tag



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
    id := "id_example" // string | Specifies the Id of the tag.
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

    request := helios.ApiDeleteTagRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Tag.DeleteTag(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tag.DeleteTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the Id of the tag. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagRequest struct via the builder pattern


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


## GetTagById

> Tag GetTagById(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Get Tag by id.



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
    id := "id_example" // string | Specifies the Id of the tag.
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

    request := helios.ApiGetTagByIdRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Tag.GetTagById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tag.GetTagById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTagById`: Tag
    fmt.Fprintf(os.Stdout, "Response from `Tag.GetTagById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the Id of the tag. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**Tag**](Tag.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTags

> []Tag GetTags(ctx).AccessClusterId(accessClusterId).RegionId(regionId).Ids(ids).Names(names).Namespaces(namespaces).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeMarkedForDeletion(includeMarkedForDeletion).Execute()

Get tags based on filters.



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
    ids := []string{"Inner_example"} // []string | Filter by a list of Tag Ids. If Ids are mentioned all other fields will be ignored. (optional)
    names := []string{"Inner_example"} // []string | Filter by a list of Tag names. (optional)
    namespaces := []string{"Inner_example"} // []string | Filter by a list of Namespaces. (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which tags are to be returned. (optional)
    includeTenants := true // bool | IncludeTenants specifies if tags of all the tenants under the hierarchy of the logged in user's organization should be returned. False, by default. (optional)
    includeMarkedForDeletion := true // bool | Specifies if tags marked for deletion should be shown. These are tags which are undergoing deletion. False, by default. (optional)

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

    request := helios.ApiGetTagsRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        Ids: &ids
        Names: &names
        Namespaces: &namespaces
        TenantIds: &tenantIds
        IncludeTenants: &includeTenants
        IncludeMarkedForDeletion: &includeMarkedForDeletion
    }

    resp, r, err := api_client.Tag.GetTags(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tag.GetTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTags`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `Tag.GetTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **ids** | **[]string** | Filter by a list of Tag Ids. If Ids are mentioned all other fields will be ignored. | 
 **names** | **[]string** | Filter by a list of Tag names. | 
 **namespaces** | **[]string** | Filter by a list of Namespaces. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which tags are to be returned. | 
 **includeTenants** | **bool** | IncludeTenants specifies if tags of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. False, by default. | 
 **includeMarkedForDeletion** | **bool** | Specifies if tags marked for deletion should be shown. These are tags which are undergoing deletion. False, by default. | 

### Return type

[**[]Tag**](Tag.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTag

> Tag UpdateTag(ctx, id).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Update a Tag



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
    id := "id_example" // string | Specifies the Id of the tag.
    body := *openapiclient.NewTag("Name_example", "Namespace_example") // Tag | Request to update a tag.
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

    request := helios.ApiUpdateTagRequest{
        Id: &id
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Tag.UpdateTag(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Tag.UpdateTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTag`: Tag
    fmt.Fprintf(os.Stdout, "Response from `Tag.UpdateTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the Id of the tag. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Tag**](Tag.md) | Request to update a tag. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**Tag**](Tag.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

