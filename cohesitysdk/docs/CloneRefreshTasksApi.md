# \CloneRefreshTasksApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCloneRefreshTask**](CloneRefreshTasksApi.md#CreateCloneRefreshTask) | **Post** /public/restore/applicationsClone/refresh | Create a Clone Refresh Task to refresh the clone with the new data.



## CreateCloneRefreshTask

> RestoreTaskWrapper CreateCloneRefreshTask(ctx).Body(body).Execute()

Create a Clone Refresh Task to refresh the clone with the new data.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewCloneRefreshRequest("Name_example") // CloneRefreshRequest | Request to create a Clone Refresh Task.

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

    request := cohesitysdk.ApiCreateCloneRefreshTaskRequest{
        Body: &body
    }

    resp, r, err := api_client.CloneRefreshTasksApi.CreateCloneRefreshTask(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloneRefreshTasksApi.CreateCloneRefreshTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCloneRefreshTask`: RestoreTaskWrapper
    fmt.Fprintf(os.Stdout, "Response from `CloneRefreshTasksApi.CreateCloneRefreshTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCloneRefreshTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CloneRefreshRequest**](CloneRefreshRequest.md) | Request to create a Clone Refresh Task. | 

### Return type

[**RestoreTaskWrapper**](RestoreTaskWrapper.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

