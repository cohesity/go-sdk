# \NodeGroup

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNodeGroup**](NodeGroup.md#CreateNodeGroup) | **Post** /node-groups | Create a Node Group.
[**DeleteNodeGroup**](NodeGroup.md#DeleteNodeGroup) | **Delete** /node-groups/{groupName} | Delete a Node Group.
[**GetNodeGroupByName**](NodeGroup.md#GetNodeGroupByName) | **Get** /node-groups/{groupName} | List Node Groups for a given Group Name.
[**GetNodeGroups**](NodeGroup.md#GetNodeGroups) | **Get** /node-groups | List Node Groups based on provided filtering parameters.
[**UpdateNodeGroup**](NodeGroup.md#UpdateNodeGroup) | **Put** /node-groups/{groupName} | Update a Node Group.



## CreateNodeGroup

> NodeGroupResponse CreateNodeGroup(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Create a Node Group.



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
    body := *openapiclient.NewNodeGroupRequest("Name_example") // NodeGroupRequest | Request to create a Node Group.
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

    request := helios.ApiCreateNodeGroupRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.NodeGroup.CreateNodeGroup(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeGroup.CreateNodeGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNodeGroup`: NodeGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `NodeGroup.CreateNodeGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNodeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NodeGroupRequest**](NodeGroupRequest.md) | Request to create a Node Group. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

> DeleteNodeGroup(ctx, groupName).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Delete a Node Group.



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
    groupName := "groupName_example" // string | Specifies a unique name of the Node Group to delete.
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

    request := helios.ApiDeleteNodeGroupRequest{
        GroupName: &groupName
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.NodeGroup.DeleteNodeGroup(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeGroup.DeleteNodeGroup``: %v\n", err)
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


## GetNodeGroupByName

> NodeGroupResponse GetNodeGroupByName(ctx, groupName).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

List Node Groups for a given Group Name.



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
    groupName := "groupName_example" // string | Specifies a unique id of Node Group to return.
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

    request := helios.ApiGetNodeGroupByNameRequest{
        GroupName: &groupName
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.NodeGroup.GetNodeGroupByName(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeGroup.GetNodeGroupByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodeGroupByName`: NodeGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `NodeGroup.GetNodeGroupByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** | Specifies a unique id of Node Group to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeGroupByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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


## GetNodeGroups

> NodeGroupResponse GetNodeGroups(ctx).AccessClusterId(accessClusterId).RegionId(regionId).GroupNames(groupNames).GroupType(groupType).Execute()

List Node Groups based on provided filtering parameters.



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
    groupNames := []string{"Inner_example"} // []string | Filter node groups by a list of node group names. (optional)
    groupType := int32(56) // int32 | Filter node groups by a node group type. (optional)

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

    request := helios.ApiGetNodeGroupsRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        GroupNames: &groupNames
        GroupType: &groupType
    }

    resp, r, err := api_client.NodeGroup.GetNodeGroups(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeGroup.GetNodeGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodeGroups`: NodeGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `NodeGroup.GetNodeGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
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

> NodeGroupResponse UpdateNodeGroup(ctx, groupName).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Update a Node Group.



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
    groupName := "groupName_example" // string | Specifies a unique name of the Node Group to update.
    body := *openapiclient.NewNodeGroupRequest("Name_example") // NodeGroupRequest | Request to update a Node Group.
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

    request := helios.ApiUpdateNodeGroupRequest{
        GroupName: &groupName
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.NodeGroup.UpdateNodeGroup(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeGroup.UpdateNodeGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNodeGroup`: NodeGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `NodeGroup.UpdateNodeGroup`: %v\n", resp)
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
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

