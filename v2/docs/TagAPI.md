# \TagAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTag**](TagAPI.md#CreateTag) | **Post** /tags | Create a Tag
[**DeleteTag**](TagAPI.md#DeleteTag) | **Delete** /tags/{id} | Delete a Tag
[**GetTagById**](TagAPI.md#GetTagById) | **Get** /tags/{id} | Get Tag by id.
[**GetTags**](TagAPI.md#GetTags) | **Get** /tags | Get tags based on filters.
[**UpdateTag**](TagAPI.md#UpdateTag) | **Put** /tags/{id} | Update a Tag



## CreateTag

> Tag CreateTag(ctx).Body(body).Execute()

Create a Tag



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
	body := *openapiclient.NewTag("Name_example", "Namespace_example") // Tag | Request to create a Tag.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagAPI.CreateTag(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagAPI.CreateTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTag`: Tag
	fmt.Fprintf(os.Stdout, "Response from `TagAPI.CreateTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Tag**](Tag.md) | Request to create a Tag. | 

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

> DeleteTag(ctx, id).Execute()

Delete a Tag



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
	id := "id_example" // string | Specifies the Id of the tag.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TagAPI.DeleteTag(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagAPI.DeleteTag``: %v\n", err)
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

> Tag GetTagById(ctx, id).Execute()

Get Tag by id.



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
	id := "id_example" // string | Specifies the Id of the tag.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagAPI.GetTagById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagAPI.GetTagById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTagById`: Tag
	fmt.Fprintf(os.Stdout, "Response from `TagAPI.GetTagById`: %v\n", resp)
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

> []Tag GetTags(ctx).Ids(ids).Names(names).Namespaces(namespaces).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeMarkedForDeletion(includeMarkedForDeletion).Execute()

Get tags based on filters.



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
	ids := []string{"Inner_example"} // []string | Filter by a list of Tag Ids. If Ids are mentioned all other fields will be ignored. (optional)
	names := []string{"Inner_example"} // []string | Filter by a list of Tag names. (optional)
	namespaces := []string{"Inner_example"} // []string | Filter by a list of Namespaces. (optional)
	tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which tags are to be returned. (optional)
	includeTenants := true // bool | IncludeTenants specifies if tags of all the tenants under the hierarchy of the logged in user's organization should be returned. False, by default. (optional)
	includeMarkedForDeletion := true // bool | Specifies if tags marked for deletion should be shown. These are tags which are undergoing deletion. False, by default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagAPI.GetTags(context.Background()).Ids(ids).Names(names).Namespaces(namespaces).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeMarkedForDeletion(includeMarkedForDeletion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagAPI.GetTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTags`: []Tag
	fmt.Fprintf(os.Stdout, "Response from `TagAPI.GetTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

> Tag UpdateTag(ctx, id).Body(body).Execute()

Update a Tag



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
	id := "id_example" // string | Specifies the Id of the tag.
	body := *openapiclient.NewTag("Name_example", "Namespace_example") // Tag | Request to update a tag.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagAPI.UpdateTag(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagAPI.UpdateTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTag`: Tag
	fmt.Fprintf(os.Stdout, "Response from `TagAPI.UpdateTag`: %v\n", resp)
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

