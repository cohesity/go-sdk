# \TestDataManagement

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTdmTask**](TestDataManagement.md#CreateTdmTask) | **Post** /tdm/tasks | Create a TDM task
[**DeleteTdmSnapshotById**](TestDataManagement.md#DeleteTdmSnapshotById) | **Delete** /tdm/snapshots/{id} | Delete a snapshot by ID
[**GetTdmObjectById**](TestDataManagement.md#GetTdmObjectById) | **Get** /tdm/objects/{id} | Get TDM object by ID
[**GetTdmObjects**](TestDataManagement.md#GetTdmObjects) | **Get** /tdm/objects | Get all TDM objects
[**GetTdmTaskById**](TestDataManagement.md#GetTdmTaskById) | **Get** /tdm/tasks/{id} | Get a TDM task by ID
[**GetTdmTasks**](TestDataManagement.md#GetTdmTasks) | **Get** /tdm/tasks | Get all TDM tasks
[**GetTdmTimelineEventsByObjectId**](TestDataManagement.md#GetTdmTimelineEventsByObjectId) | **Get** /tdm/objects/{id}/timeline-events | Get timeline events of object
[**UpdateTdmSnapshotById**](TestDataManagement.md#UpdateTdmSnapshotById) | **Put** /tdm/snapshots/{id} | Update a snapshot by ID



## CreateTdmTask

> TdmTask CreateTdmTask(ctx).Body(body).Execute()

Create a TDM task



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
    body := *openapiclient.NewCreateTdmTaskRequest("Action_example") // CreateTdmTaskRequest | Specifies the parameters to create a TDM task.

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

    request := onprem.ApiCreateTdmTaskRequest{
        Body: &body
    }

    resp, r, err := api_client.TestDataManagement.CreateTdmTask(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestDataManagement.CreateTdmTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTdmTask`: TdmTask
    fmt.Fprintf(os.Stdout, "Response from `TestDataManagement.CreateTdmTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTdmTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateTdmTaskRequest**](CreateTdmTaskRequest.md) | Specifies the parameters to create a TDM task. | 

### Return type

[**TdmTask**](TdmTask.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTdmSnapshotById

> DeleteTdmSnapshotById(ctx, id).Execute()

Delete a snapshot by ID



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
    id := "id_example" // string | Specifies the ID of the snapshot.

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

    request := onprem.ApiDeleteTdmSnapshotByIdRequest{
        Id: &id
    }

    resp, r, err := api_client.TestDataManagement.DeleteTdmSnapshotById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestDataManagement.DeleteTdmSnapshotById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the ID of the snapshot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTdmSnapshotByIdRequest struct via the builder pattern


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


## GetTdmObjectById

> TdmObject GetTdmObjectById(ctx, id).Execute()

Get TDM object by ID



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
    id := "id_example" // string | Specifies the ID of the TDM object.

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

    request := onprem.ApiGetTdmObjectByIdRequest{
        Id: &id
    }

    resp, r, err := api_client.TestDataManagement.GetTdmObjectById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestDataManagement.GetTdmObjectById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTdmObjectById`: TdmObject
    fmt.Fprintf(os.Stdout, "Response from `TestDataManagement.GetTdmObjectById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the ID of the TDM object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTdmObjectByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TdmObject**](TdmObject.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTdmObjects

> TdmObjects GetTdmObjects(ctx).Ids(ids).Environments(environments).Name(name).TaskIds(taskIds).Statuses(statuses).PaginationCookie(paginationCookie).Execute()

Get all TDM objects



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
    ids := []string{"Inner_example"} // []string | Get the objects matching specifies IDs. (optional)
    environments := []string{"Environments_example"} // []string | Get the objects matching specified environments. (optional)
    name := "name_example" // string | Get the objects matching specified name. (optional)
    taskIds := []string{"Inner_example"} // []string | Get the objects belonging to the specified TDM task IDs. (optional)
    statuses := []string{"Statuses_example"} // []string | Get the objects matching specified statuses. (optional)
    paginationCookie := "paginationCookie_example" // string | Get the next set of objects by specifying the cookie string, as received from the server in the last call. (optional)

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

    request := onprem.ApiGetTdmObjectsRequest{
        Ids: &ids
        Environments: &environments
        Name: &name
        TaskIds: &taskIds
        Statuses: &statuses
        PaginationCookie: &paginationCookie
    }

    resp, r, err := api_client.TestDataManagement.GetTdmObjects(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestDataManagement.GetTdmObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTdmObjects`: TdmObjects
    fmt.Fprintf(os.Stdout, "Response from `TestDataManagement.GetTdmObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTdmObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** | Get the objects matching specifies IDs. | 
 **environments** | **[]string** | Get the objects matching specified environments. | 
 **name** | **string** | Get the objects matching specified name. | 
 **taskIds** | **[]string** | Get the objects belonging to the specified TDM task IDs. | 
 **statuses** | **[]string** | Get the objects matching specified statuses. | 
 **paginationCookie** | **string** | Get the next set of objects by specifying the cookie string, as received from the server in the last call. | 

### Return type

[**TdmObjects**](TdmObjects.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTdmTaskById

> TdmTask GetTdmTaskById(ctx, id).Execute()

Get a TDM task by ID



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
    id := "id_example" // string | Specifies the ID of the TDM task.

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

    request := onprem.ApiGetTdmTaskByIdRequest{
        Id: &id
    }

    resp, r, err := api_client.TestDataManagement.GetTdmTaskById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestDataManagement.GetTdmTaskById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTdmTaskById`: TdmTask
    fmt.Fprintf(os.Stdout, "Response from `TestDataManagement.GetTdmTaskById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the ID of the TDM task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTdmTaskByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TdmTask**](TdmTask.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTdmTasks

> TdmTasks GetTdmTasks(ctx).Ids(ids).Actions(actions).Environments(environments).CreatedAfterUsecs(createdAfterUsecs).CreatedBeforeUsecs(createdBeforeUsecs).Statuses(statuses).ObjectIds(objectIds).PaginationCookie(paginationCookie).Execute()

Get all TDM tasks



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
    ids := []string{"Inner_example"} // []string | Get the tasks matching specified IDs. (optional)
    actions := []string{"Actions_example"} // []string | Get the tasks matching specified actions. (optional)
    environments := []string{"Environments_example"} // []string | Get the tasks matching specified environments. (optional)
    createdAfterUsecs := int64(789) // int64 | Get the tasks created after the specified time (in usecs from epoch). (optional)
    createdBeforeUsecs := int64(789) // int64 | Get the tasks created before the specified time (in usecs from epoch). (optional)
    statuses := []string{"Statuses_example"} // []string | Get the tasks matching specified statuses. (optional)
    objectIds := []string{"Inner_example"} // []string | Get the tasks for the specified TDM object IDs. (optional)
    paginationCookie := "paginationCookie_example" // string | Get the next set of tasks by specifying the cookie string, as received from the server in the last call. (optional)

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

    request := onprem.ApiGetTdmTasksRequest{
        Ids: &ids
        Actions: &actions
        Environments: &environments
        CreatedAfterUsecs: &createdAfterUsecs
        CreatedBeforeUsecs: &createdBeforeUsecs
        Statuses: &statuses
        ObjectIds: &objectIds
        PaginationCookie: &paginationCookie
    }

    resp, r, err := api_client.TestDataManagement.GetTdmTasks(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestDataManagement.GetTdmTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTdmTasks`: TdmTasks
    fmt.Fprintf(os.Stdout, "Response from `TestDataManagement.GetTdmTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTdmTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** | Get the tasks matching specified IDs. | 
 **actions** | **[]string** | Get the tasks matching specified actions. | 
 **environments** | **[]string** | Get the tasks matching specified environments. | 
 **createdAfterUsecs** | **int64** | Get the tasks created after the specified time (in usecs from epoch). | 
 **createdBeforeUsecs** | **int64** | Get the tasks created before the specified time (in usecs from epoch). | 
 **statuses** | **[]string** | Get the tasks matching specified statuses. | 
 **objectIds** | **[]string** | Get the tasks for the specified TDM object IDs. | 
 **paginationCookie** | **string** | Get the next set of tasks by specifying the cookie string, as received from the server in the last call. | 

### Return type

[**TdmTasks**](TdmTasks.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTdmTimelineEventsByObjectId

> TdmObjectTimelineEvents GetTdmTimelineEventsByObjectId(ctx, id).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Execute()

Get timeline events of object



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
    id := "id_example" // string | Specifies the ID of the TDM object.
    createdAfter := int64(789) // int64 | Get the events created after the specified time (in usecs from epoch). (optional)
    createdBefore := int64(789) // int64 | Get the events created before the specified time (in usecs from epoch). (optional)

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

    request := onprem.ApiGetTdmTimelineEventsByObjectIdRequest{
        Id: &id
        CreatedAfter: &createdAfter
        CreatedBefore: &createdBefore
    }

    resp, r, err := api_client.TestDataManagement.GetTdmTimelineEventsByObjectId(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestDataManagement.GetTdmTimelineEventsByObjectId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTdmTimelineEventsByObjectId`: TdmObjectTimelineEvents
    fmt.Fprintf(os.Stdout, "Response from `TestDataManagement.GetTdmTimelineEventsByObjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the ID of the TDM object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTdmTimelineEventsByObjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createdAfter** | **int64** | Get the events created after the specified time (in usecs from epoch). | 
 **createdBefore** | **int64** | Get the events created before the specified time (in usecs from epoch). | 

### Return type

[**TdmObjectTimelineEvents**](TdmObjectTimelineEvents.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTdmSnapshotById

> TdmSnapshot UpdateTdmSnapshotById(ctx, id).Body(body).Execute()

Update a snapshot by ID



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
    id := "id_example" // string | Specifies the ID of the snapshot.
    body := *openapiclient.NewUpdateTdmSnapshotRequest() // UpdateTdmSnapshotRequest | Specifies the parameters to update the snapshot.

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

    request := onprem.ApiUpdateTdmSnapshotByIdRequest{
        Id: &id
        Body: &body
    }

    resp, r, err := api_client.TestDataManagement.UpdateTdmSnapshotById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestDataManagement.UpdateTdmSnapshotById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTdmSnapshotById`: TdmSnapshot
    fmt.Fprintf(os.Stdout, "Response from `TestDataManagement.UpdateTdmSnapshotById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the ID of the snapshot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTdmSnapshotByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateTdmSnapshotRequest**](UpdateTdmSnapshotRequest.md) | Specifies the parameters to update the snapshot. | 

### Return type

[**TdmSnapshot**](TdmSnapshot.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

