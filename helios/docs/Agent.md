# \Agent

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUpgradeTask**](Agent.md#CreateUpgradeTask) | **Post** /data-protect/agents/upgrade-tasks | Create an upgrade task
[**GetUpgradeTasks**](Agent.md#GetUpgradeTasks) | **Get** /data-protect/agents/upgrade-tasks | Get upgrade tasks
[**McmGetAgentImageDetails**](Agent.md#McmGetAgentImageDetails) | **Get** /mcm/data-protect/agents/images | Get agent images details.
[**PerformActionOnAgentUpgradeTask**](Agent.md#PerformActionOnAgentUpgradeTask) | **Post** /data-protect/agents/upgrade-tasks/actions | Perform action on an upgrade task.



## CreateUpgradeTask

> AgentUpgradeTaskState CreateUpgradeTask(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Create an upgrade task



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
    body := *openapiclient.NewCreateUpgradeTaskRequest() // CreateUpgradeTaskRequest | Specifies parameters to create a schedule based agent upgrade task.
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

    request := helios.ApiCreateUpgradeTaskRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Agent.CreateUpgradeTask(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Agent.CreateUpgradeTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUpgradeTask`: AgentUpgradeTaskState
    fmt.Fprintf(os.Stdout, "Response from `Agent.CreateUpgradeTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUpgradeTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateUpgradeTaskRequest**](CreateUpgradeTaskRequest.md) | Specifies parameters to create a schedule based agent upgrade task. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**AgentUpgradeTaskState**](AgentUpgradeTaskState.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUpgradeTasks

> AgentUpgradeTaskStates GetUpgradeTasks(ctx).AccessClusterId(accessClusterId).RegionId(regionId).Ids(ids).TenantIds(tenantIds).IncludeTenants(includeTenants).Execute()

Get upgrade tasks



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
    ids := []int64{int64(123)} // []int64 | Specifies IDs of tasks to be fetched. (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
    includeTenants := true // bool | If true, the response will include upgrade tasks which were created by all tenants which the current user has permission to see. If false, then only upgrade tasks created by the current user will be returned. (optional)

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

    request := helios.ApiGetUpgradeTasksRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        Ids: &ids
        TenantIds: &tenantIds
        IncludeTenants: &includeTenants
    }

    resp, r, err := api_client.Agent.GetUpgradeTasks(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Agent.GetUpgradeTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUpgradeTasks`: AgentUpgradeTaskStates
    fmt.Fprintf(os.Stdout, "Response from `Agent.GetUpgradeTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUpgradeTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **ids** | **[]int64** | Specifies IDs of tasks to be fetched. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **includeTenants** | **bool** | If true, the response will include upgrade tasks which were created by all tenants which the current user has permission to see. If false, then only upgrade tasks created by the current user will be returned. | 

### Return type

[**AgentUpgradeTaskStates**](AgentUpgradeTaskStates.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McmGetAgentImageDetails

> McmAgentImagesResponse McmGetAgentImageDetails(ctx).RegionId(regionId).Platform(platform).PackageType(packageType).Execute()

Get agent images details.



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
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    platform := "platform_example" // string | Specifies a platform for which agent information need to be fetched. (optional)
    packageType := "packageType_example" // string | Specifies a package type for which agent information need to be fetched. (optional)

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

    request := helios.ApiMcmGetAgentImageDetailsRequest{
        RegionId: &regionId
        Platform: &platform
        PackageType: &packageType
    }

    resp, r, err := api_client.Agent.McmGetAgentImageDetails(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Agent.McmGetAgentImageDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `McmGetAgentImageDetails`: McmAgentImagesResponse
    fmt.Fprintf(os.Stdout, "Response from `Agent.McmGetAgentImageDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMcmGetAgentImageDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **platform** | **string** | Specifies a platform for which agent information need to be fetched. | 
 **packageType** | **string** | Specifies a package type for which agent information need to be fetched. | 

### Return type

[**McmAgentImagesResponse**](McmAgentImagesResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PerformActionOnAgentUpgradeTask

> AgentUpgradeTaskActionObject PerformActionOnAgentUpgradeTask(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Perform action on an upgrade task.



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
    body := *openapiclient.NewAgentUpgradeTaskActionRequest() // AgentUpgradeTaskActionRequest | Specifies the parameters to perform an action on an agent upgrade task.
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

    request := helios.ApiPerformActionOnAgentUpgradeTaskRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Agent.PerformActionOnAgentUpgradeTask(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Agent.PerformActionOnAgentUpgradeTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PerformActionOnAgentUpgradeTask`: AgentUpgradeTaskActionObject
    fmt.Fprintf(os.Stdout, "Response from `Agent.PerformActionOnAgentUpgradeTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPerformActionOnAgentUpgradeTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AgentUpgradeTaskActionRequest**](AgentUpgradeTaskActionRequest.md) | Specifies the parameters to perform an action on an agent upgrade task. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**AgentUpgradeTaskActionObject**](AgentUpgradeTaskActionObject.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

