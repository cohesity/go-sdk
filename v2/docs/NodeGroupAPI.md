# \NodeGroupAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNodeGroup**](NodeGroupAPI.md#CreateNodeGroup) | **Post** /node-groups | Create a Node Group.
[**DeleteNodeGroup**](NodeGroupAPI.md#DeleteNodeGroup) | **Delete** /node-groups/{groupName} | Delete a Node Group.
[**GetNodeGroups**](NodeGroupAPI.md#GetNodeGroups) | **Get** /node-groups | List Node Groups based on provided filtering parameters.
[**UpdateNodeGroup**](NodeGroupAPI.md#UpdateNodeGroup) | **Put** /node-groups/{groupName} | Update a Node Group.



## CreateNodeGroup

> NodeGroupResponse CreateNodeGroup(ctx).Body(body).Execute()

Create a Node Group.



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
	body := *openapiclient.NewNodeGroupRequest("Name_example") // NodeGroupRequest | Request to create a Node Group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeGroupAPI.CreateNodeGroup(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeGroupAPI.CreateNodeGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNodeGroup`: NodeGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `NodeGroupAPI.CreateNodeGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNodeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NodeGroupRequest**](NodeGroupRequest.md) | Request to create a Node Group. | 

### Return type

[**NodeGroupResponse**](NodeGroupResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNodeGroup

> DeleteNodeGroup(ctx, groupName).Execute()

Delete a Node Group.



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
	groupName := "groupName_example" // string | Specifies a unique name of the Node Group to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NodeGroupAPI.DeleteNodeGroup(context.Background(), groupName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeGroupAPI.DeleteNodeGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** | Specifies a unique name of the Node Group to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNodeGroupRequest struct via the builder pattern


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


## GetNodeGroups

> NodeGroupResponse GetNodeGroups(ctx).GroupNames(groupNames).GroupType(groupType).Execute()

List Node Groups based on provided filtering parameters.



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
	groupNames := []string{"Inner_example"} // []string | Filter node groups by a list of node group names. (optional)
	groupType := int32(56) // int32 | Filter node groups by a node group type. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeGroupAPI.GetNodeGroups(context.Background()).GroupNames(groupNames).GroupType(groupType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeGroupAPI.GetNodeGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodeGroups`: NodeGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `NodeGroupAPI.GetNodeGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupNames** | **[]string** | Filter node groups by a list of node group names. | 
 **groupType** | **int32** | Filter node groups by a node group type. | 

### Return type

[**NodeGroupResponse**](NodeGroupResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNodeGroup

> NodeGroupResponse UpdateNodeGroup(ctx, groupName).Body(body).Execute()

Update a Node Group.



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
	groupName := "groupName_example" // string | Specifies a unique name of the Node Group to update.
	body := *openapiclient.NewNodeGroupRequest("Name_example") // NodeGroupRequest | Request to update a Node Group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeGroupAPI.UpdateNodeGroup(context.Background(), groupName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeGroupAPI.UpdateNodeGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNodeGroup`: NodeGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `NodeGroupAPI.UpdateNodeGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** | Specifies a unique name of the Node Group to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNodeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NodeGroupRequest**](NodeGroupRequest.md) | Request to update a Node Group. | 

### Return type

[**NodeGroupResponse**](NodeGroupResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

