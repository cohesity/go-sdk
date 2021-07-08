# \RemoteRestoreApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRemoteVaultRestoreTask**](RemoteRestoreApi.md#CreateRemoteVaultRestoreTask) | **Post** /public/remoteVaults/restoreTasks | Create a remote Vault restore task. (CloudRetrieve)
[**CreateRemoteVaultSearchJob**](RemoteRestoreApi.md#CreateRemoteVaultSearchJob) | **Post** /public/remoteVaults/searchJobs | Create a search of a remote Vault. (CloudRetrieve)
[**GetRemoteVaultSearchJobResults**](RemoteRestoreApi.md#GetRemoteVaultSearchJobResults) | **Get** /public/remoteVaults/searchJobResults | List details about the Job Runs of Protection Jobs found by a single search of a remote Vault. (CloudRetrieve)
[**ListRemoteVaultRestoreTasks**](RemoteRestoreApi.md#ListRemoteVaultRestoreTasks) | **Get** /public/remoteVaults/restoreTasks | List the remote Vault restore tasks that have completed or are running on this Cohesity Cluster. (CloudRetrieve)
[**ListRemoteVaultSearchJobById**](RemoteRestoreApi.md#ListRemoteVaultSearchJobById) | **Get** /public/remoteVaults/searchJobs/{id} | List details about a single search Job of a remote Vault. (CloudRetrieve)
[**ListRemoteVaultSearchJobs**](RemoteRestoreApi.md#ListRemoteVaultSearchJobs) | **Get** /public/remoteVaults/searchJobs | List all the searches of remote Vaults. (CloudRetrieve)
[**StopRemoteVaultSearchJob**](RemoteRestoreApi.md#StopRemoteVaultSearchJob) | **Delete** /public/remoteVaults/searchJobs | Stop a search of a remote Vault (External Target). (CloudRetrieve)
[**UploadVaultEncryptionKeys**](RemoteRestoreApi.md#UploadVaultEncryptionKeys) | **Put** /public/remoteVaults/encryptionKeys/{id} | Upload the encryption keys required to restore data from a remote Vault. (CloudRetrieve)



## CreateRemoteVaultRestoreTask

> UniversalId CreateRemoteVaultRestoreTask(ctx).Body(body).Execute()

Create a remote Vault restore task. (CloudRetrieve)



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
    body := *openapiclient.NewCreateRemoteVaultRestoreTaskParameters("TODO", "TaskName_example", NullableInt64(123)) // CreateRemoteVaultRestoreTaskParameters | Request to create a remote Vault restore task.

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

    request := cohesitysdk.ApiCreateRemoteVaultRestoreTaskRequest{
        Body: &body
    }

    resp, r, err := api_client.RemoteRestoreApi.CreateRemoteVaultRestoreTask(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteRestoreApi.CreateRemoteVaultRestoreTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRemoteVaultRestoreTask`: UniversalId
    fmt.Fprintf(os.Stdout, "Response from `RemoteRestoreApi.CreateRemoteVaultRestoreTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRemoteVaultRestoreTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateRemoteVaultRestoreTaskParameters**](CreateRemoteVaultRestoreTaskParameters.md) | Request to create a remote Vault restore task. | 

### Return type

[**UniversalId**](UniversalId.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRemoteVaultSearchJob

> CreatedRemoteVaultSearchJobUid CreateRemoteVaultSearchJob(ctx).Body(body).Execute()

Create a search of a remote Vault. (CloudRetrieve)



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
    body := *openapiclient.NewCreateRemoteVaultSearchJobParameters("SearchJobName_example", NullableInt64(123)) // CreateRemoteVaultSearchJobParameters | Request to create a search of a remote Vault.

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

    request := cohesitysdk.ApiCreateRemoteVaultSearchJobRequest{
        Body: &body
    }

    resp, r, err := api_client.RemoteRestoreApi.CreateRemoteVaultSearchJob(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteRestoreApi.CreateRemoteVaultSearchJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRemoteVaultSearchJob`: CreatedRemoteVaultSearchJobUid
    fmt.Fprintf(os.Stdout, "Response from `RemoteRestoreApi.CreateRemoteVaultSearchJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRemoteVaultSearchJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateRemoteVaultSearchJobParameters**](CreateRemoteVaultSearchJobParameters.md) | Request to create a search of a remote Vault. | 

### Return type

[**CreatedRemoteVaultSearchJobUid**](CreatedRemoteVaultSearchJobUid.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteVaultSearchJobResults

> RemoteVaultSearchJobResults GetRemoteVaultSearchJobResults(ctx).SearchJobId(searchJobId).ClusterId(clusterId).ClusterIncarnationId(clusterIncarnationId).PageCount(pageCount).ClusterName(clusterName).Cookie(cookie).Execute()

List details about the Job Runs of Protection Jobs found by a single search of a remote Vault. (CloudRetrieve)



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
    searchJobId := int64(789) // int64 | Specifies the id of the remote Vault search Job assigned by the Cohesity Cluster. Used in combination with the clusterId and clusterIncarnationId to uniquely identify the search Job.
    clusterId := int64(789) // int64 | Specifies the Cohesity Cluster id where the search Job was created. Used in combination with the searchJobId and clusterIncarnationId to uniquely identify the search Job.
    clusterIncarnationId := int64(789) // int64 | Specifies the incarnation id of the Cohesity Cluster where the search Job was created. Used in combination with the searchJobId and clusterId to uniquely identify the search Job.
    pageCount := int32(56) // int32 | Specifies the number of Protection Jobs to return in the response to support pagination. (optional)
    clusterName := "clusterName_example" // string | Optionally filter the result by the remote Cohesity Cluster name. (optional)
    cookie := "cookie_example" // string | Specifies the opaque string cookie returned by the previous response, to get next set of results. Used in combination with pageCount to support pagination. (optional)

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

    request := cohesitysdk.ApiGetRemoteVaultSearchJobResultsRequest{
        SearchJobId: &searchJobId
        ClusterId: &clusterId
        ClusterIncarnationId: &clusterIncarnationId
        PageCount: &pageCount
        ClusterName: &clusterName
        Cookie: &cookie
    }

    resp, r, err := api_client.RemoteRestoreApi.GetRemoteVaultSearchJobResults(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteRestoreApi.GetRemoteVaultSearchJobResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRemoteVaultSearchJobResults`: RemoteVaultSearchJobResults
    fmt.Fprintf(os.Stdout, "Response from `RemoteRestoreApi.GetRemoteVaultSearchJobResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteVaultSearchJobResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchJobId** | **int64** | Specifies the id of the remote Vault search Job assigned by the Cohesity Cluster. Used in combination with the clusterId and clusterIncarnationId to uniquely identify the search Job. | 
 **clusterId** | **int64** | Specifies the Cohesity Cluster id where the search Job was created. Used in combination with the searchJobId and clusterIncarnationId to uniquely identify the search Job. | 
 **clusterIncarnationId** | **int64** | Specifies the incarnation id of the Cohesity Cluster where the search Job was created. Used in combination with the searchJobId and clusterId to uniquely identify the search Job. | 
 **pageCount** | **int32** | Specifies the number of Protection Jobs to return in the response to support pagination. | 
 **clusterName** | **string** | Optionally filter the result by the remote Cohesity Cluster name. | 
 **cookie** | **string** | Specifies the opaque string cookie returned by the previous response, to get next set of results. Used in combination with pageCount to support pagination. | 

### Return type

[**RemoteVaultSearchJobResults**](RemoteVaultSearchJobResults.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRemoteVaultRestoreTasks

> []RemoteVaultRestoreTaskStatus ListRemoteVaultRestoreTasks(ctx).Execute()

List the remote Vault restore tasks that have completed or are running on this Cohesity Cluster. (CloudRetrieve)



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

    request := cohesitysdk.ApiListRemoteVaultRestoreTasksRequest{
    }

    resp, r, err := api_client.RemoteRestoreApi.ListRemoteVaultRestoreTasks(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteRestoreApi.ListRemoteVaultRestoreTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRemoteVaultRestoreTasks`: []RemoteVaultRestoreTaskStatus
    fmt.Fprintf(os.Stdout, "Response from `RemoteRestoreApi.ListRemoteVaultRestoreTasks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRemoteVaultRestoreTasksRequest struct via the builder pattern


### Return type

[**[]RemoteVaultRestoreTaskStatus**](RemoteVaultRestoreTaskStatus.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRemoteVaultSearchJobById

> RemoteVaultSearchJobInformation ListRemoteVaultSearchJobById(ctx, id).Execute()

List details about a single search Job of a remote Vault. (CloudRetrieve)



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
    id := int64(789) // int64 | Specifies the id of the remote Vault search Job to return.

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

    request := cohesitysdk.ApiListRemoteVaultSearchJobByIdRequest{
        Id: &id
    }

    resp, r, err := api_client.RemoteRestoreApi.ListRemoteVaultSearchJobById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteRestoreApi.ListRemoteVaultSearchJobById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRemoteVaultSearchJobById`: RemoteVaultSearchJobInformation
    fmt.Fprintf(os.Stdout, "Response from `RemoteRestoreApi.ListRemoteVaultSearchJobById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the id of the remote Vault search Job to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRemoteVaultSearchJobByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RemoteVaultSearchJobInformation**](RemoteVaultSearchJobInformation.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRemoteVaultSearchJobs

> []RemoteVaultSearchJobInformation ListRemoteVaultSearchJobs(ctx).Execute()

List all the searches of remote Vaults. (CloudRetrieve)



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

    request := cohesitysdk.ApiListRemoteVaultSearchJobsRequest{
    }

    resp, r, err := api_client.RemoteRestoreApi.ListRemoteVaultSearchJobs(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteRestoreApi.ListRemoteVaultSearchJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRemoteVaultSearchJobs`: []RemoteVaultSearchJobInformation
    fmt.Fprintf(os.Stdout, "Response from `RemoteRestoreApi.ListRemoteVaultSearchJobs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRemoteVaultSearchJobsRequest struct via the builder pattern


### Return type

[**[]RemoteVaultSearchJobInformation**](RemoteVaultSearchJobInformation.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopRemoteVaultSearchJob

> StopRemoteVaultSearchJob(ctx).Body(body).Execute()

Stop a search of a remote Vault (External Target). (CloudRetrieve)



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
    body := *openapiclient.NewStopRemoteVaultSearchJobParameters() // StopRemoteVaultSearchJobParameters | Request to stop a Remote Vault Search Job.

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

    request := cohesitysdk.ApiStopRemoteVaultSearchJobRequest{
        Body: &body
    }

    resp, r, err := api_client.RemoteRestoreApi.StopRemoteVaultSearchJob(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteRestoreApi.StopRemoteVaultSearchJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStopRemoteVaultSearchJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**StopRemoteVaultSearchJobParameters**](StopRemoteVaultSearchJobParameters.md) | Request to stop a Remote Vault Search Job. | 

### Return type

 (empty response body)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadVaultEncryptionKeys

> UploadVaultEncryptionKeys(ctx, id).Body(body).Execute()

Upload the encryption keys required to restore data from a remote Vault. (CloudRetrieve)



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
    id := int64(789) // int64 | Specifies a unique id of the Vault.
    body := []openapiclient.VaultEncryptionKey{*openapiclient.NewVaultEncryptionKey()} // []VaultEncryptionKey | Request to upload encryption keys of a remote Vault. (optional)

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

    request := cohesitysdk.ApiUploadVaultEncryptionKeysRequest{
        Id: &id
        Body: &body
    }

    resp, r, err := api_client.RemoteRestoreApi.UploadVaultEncryptionKeys(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteRestoreApi.UploadVaultEncryptionKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the Vault. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadVaultEncryptionKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**[]VaultEncryptionKey**](VaultEncryptionKey.md) | Request to upload encryption keys of a remote Vault. | 

### Return type

 (empty response body)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

