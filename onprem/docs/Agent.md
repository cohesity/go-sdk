# \Agent

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUpgradeTask**](Agent.md#CreateUpgradeTask) | **Post** /data-protect/agents/upgrade-tasks | Create an upgrade task
[**GetUpgradeTasks**](Agent.md#GetUpgradeTasks) | **Get** /data-protect/agents/upgrade-tasks | Get upgrade tasks
[**PerformActionOnAgentUpgradeTask**](Agent.md#PerformActionOnAgentUpgradeTask) | **Post** /data-protect/agents/upgrade-tasks/actions | Perform action on an upgrade task.



## CreateUpgradeTask

> AgentUpgradeTaskState CreateUpgradeTask(ctx).Body(body).Execute()

Create an upgrade task



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
    body := *openapiclient.NewCreateUpgradeTaskRequest() // CreateUpgradeTaskRequest | Specifies parameters to create a schedule based agent upgrade task.

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

    request := onprem.ApiCreateUpgradeTaskRequest{
        Body: &body
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

> AgentUpgradeTaskStates GetUpgradeTasks(ctx).Ids(ids).TenantIds(tenantIds).IncludeTenants(includeTenants).Execute()

Get upgrade tasks



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

    request := onprem.ApiGetUpgradeTasksRequest{
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


## PerformActionOnAgentUpgradeTask

> AgentUpgradeTaskActionObject PerformActionOnAgentUpgradeTask(ctx).Body(body).Execute()

Perform action on an upgrade task.



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
    body := *openapiclient.NewAgentUpgradeTaskActionRequest() // AgentUpgradeTaskActionRequest | Specifies the parameters to perform an action on an agent upgrade task.

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

    request := onprem.ApiPerformActionOnAgentUpgradeTaskRequest{
        Body: &body
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

