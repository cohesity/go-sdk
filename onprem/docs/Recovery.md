# \Recovery

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelRecoveryById**](Recovery.md#CancelRecoveryById) | **Post** /data-protect/recoveries/{id}/cancel | Cancel Recovery for a given id.
[**CreateDownloadFilesAndFoldersRecovery**](Recovery.md#CreateDownloadFilesAndFoldersRecovery) | **Post** /data-protect/recoveries/download-files-folders | Create a download files and folders recovery.
[**CreateRecovery**](Recovery.md#CreateRecovery) | **Post** /data-protect/recoveries | Performs a Recovery.
[**DownloadFilesFromRecovery**](Recovery.md#DownloadFilesFromRecovery) | **Get** /data-protect/recoveries/{id}/download-files | Download files from the given download file recovery.
[**DownloadIndexedFile**](Recovery.md#DownloadIndexedFile) | **Get** /data-protect/snapshots/{snapshotsId}/download-file | Download an indexed file.
[**GetRecoveries**](Recovery.md#GetRecoveries) | **Get** /data-protect/recoveries | Lists the Recoveries.
[**GetRecoveryById**](Recovery.md#GetRecoveryById) | **Get** /data-protect/recoveries/{id} | Get Recovery for a given id.
[**GetRecoveryDebugLogs**](Recovery.md#GetRecoveryDebugLogs) | **Get** /data-protect/recoveries/{id}/debug-logs | Get the debug logs for a particular recovery operation.
[**GetRecoveryErrorsReport**](Recovery.md#GetRecoveryErrorsReport) | **Get** /data-protect/recoveries/{id}/download-messages | Get the CSV of errors/warnings for a given recovery operation.
[**TearDownRecoveryById**](Recovery.md#TearDownRecoveryById) | **Post** /data-protect/recoveries/{id}/tear-down | Tear down Recovery for a given id.



## CancelRecoveryById

> CancelRecoveryById(ctx, id).Execute()

Cancel Recovery for a given id.



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
    id := "id_example" // string | Specifies the id of a Recovery.

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

    request := onprem.ApiCancelRecoveryByIdRequest{
        Id: &id
    }

    resp, r, err := api_client.Recovery.CancelRecoveryById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Recovery.CancelRecoveryById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id of a Recovery. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelRecoveryByIdRequest struct via the builder pattern


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


## CreateDownloadFilesAndFoldersRecovery

> Recovery CreateDownloadFilesAndFoldersRecovery(ctx).Body(body).Execute()

Create a download files and folders recovery.



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
    body := *openapiclient.NewDownloadFilesAndFoldersRequestParams("Name_example", *openapiclient.NewCommonRecoverObjectSnapshotParams("SnapshotId_example"), []openapiclient.FilesAndFoldersObject{*openapiclient.NewFilesAndFoldersObject("AbsolutePath_example")}) // DownloadFilesAndFoldersRequestParams | Specifies the parameters to create a download files and folder recovery.

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

    request := onprem.ApiCreateDownloadFilesAndFoldersRecoveryRequest{
        Body: &body
    }

    resp, r, err := api_client.Recovery.CreateDownloadFilesAndFoldersRecovery(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Recovery.CreateDownloadFilesAndFoldersRecovery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDownloadFilesAndFoldersRecovery`: Recovery
    fmt.Fprintf(os.Stdout, "Response from `Recovery.CreateDownloadFilesAndFoldersRecovery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDownloadFilesAndFoldersRecoveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DownloadFilesAndFoldersRequestParams**](DownloadFilesAndFoldersRequestParams.md) | Specifies the parameters to create a download files and folder recovery. | 

### Return type

[**Recovery**](Recovery.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRecovery

> Recovery CreateRecovery(ctx).Body(body).Execute()

Performs a Recovery.



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
    body := *openapiclient.NewCreateRecoveryRequest("Name_example", "SnapshotEnvironment_example") // CreateRecoveryRequest | Specifies the parameters to create a Recovery.

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

    request := onprem.ApiCreateRecoveryRequest{
        Body: &body
    }

    resp, r, err := api_client.Recovery.CreateRecovery(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Recovery.CreateRecovery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRecovery`: Recovery
    fmt.Fprintf(os.Stdout, "Response from `Recovery.CreateRecovery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRecoveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateRecoveryRequest**](CreateRecoveryRequest.md) | Specifies the parameters to create a Recovery. | 

### Return type

[**Recovery**](Recovery.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadFilesFromRecovery

> DownloadFilesFromRecovery(ctx, id).StartOffset(startOffset).Length(length).Execute()

Download files from the given download file recovery.



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
    id := "id_example" // string | Specifies the id of a Recovery.
    startOffset := int64(789) // int64 | Specifies the start offset of file chunk to be downloaded. (optional)
    length := int64(789) // int64 | Specifies the length of bytes to download. This can not be greater than 8MB (8388608 byets) (optional)

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

    request := onprem.ApiDownloadFilesFromRecoveryRequest{
        Id: &id
        StartOffset: &startOffset
        Length: &length
    }

    resp, r, err := api_client.Recovery.DownloadFilesFromRecovery(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Recovery.DownloadFilesFromRecovery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id of a Recovery. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadFilesFromRecoveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startOffset** | **int64** | Specifies the start offset of file chunk to be downloaded. | 
 **length** | **int64** | Specifies the length of bytes to download. This can not be greater than 8MB (8388608 byets) | 

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


## DownloadIndexedFile

> DownloadIndexedFile(ctx, snapshotsId).FilePath(filePath).RetryAttempt(retryAttempt).StartOffset(startOffset).Length(length).Execute()

Download an indexed file.



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
    snapshotsId := "snapshotsId_example" // string | Specifies the snapshot id to download from.
    filePath := "filePath_example" // string | Specifies the path to the file to download. If no path is specified and snapshot environment is kVMWare, VMX file for VMware will be downloaded. For other snapshot environments, this field must be specified. (optional)
    retryAttempt := int64(789) // int64 | Specifies the number of attempts the protection run took to create this file. (optional)
    startOffset := int64(789) // int64 | Specifies the start offset of file chunk to be downloaded. (optional)
    length := int64(789) // int64 | Specifies the length of bytes to download. This can not be greater than 8MB (8388608 byets) (optional)

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

    request := onprem.ApiDownloadIndexedFileRequest{
        SnapshotsId: &snapshotsId
        FilePath: &filePath
        RetryAttempt: &retryAttempt
        StartOffset: &startOffset
        Length: &length
    }

    resp, r, err := api_client.Recovery.DownloadIndexedFile(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Recovery.DownloadIndexedFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotsId** | **string** | Specifies the snapshot id to download from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadIndexedFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filePath** | **string** | Specifies the path to the file to download. If no path is specified and snapshot environment is kVMWare, VMX file for VMware will be downloaded. For other snapshot environments, this field must be specified. | 
 **retryAttempt** | **int64** | Specifies the number of attempts the protection run took to create this file. | 
 **startOffset** | **int64** | Specifies the start offset of file chunk to be downloaded. | 
 **length** | **int64** | Specifies the length of bytes to download. This can not be greater than 8MB (8388608 byets) | 

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


## GetRecoveries

> Recoveries GetRecoveries(ctx).Ids(ids).ReturnOnlyChildRecoveries(returnOnlyChildRecoveries).TenantIds(tenantIds).IncludeTenants(includeTenants).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).StorageDomainId(storageDomainId).SnapshotTargetType(snapshotTargetType).ArchivalTargetType(archivalTargetType).SnapshotEnvironments(snapshotEnvironments).Status(status).RecoveryActions(recoveryActions).Execute()

Lists the Recoveries.



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
    ids := []string{"Inner_example"} // []string | Filter Recoveries for given ids. (optional)
    returnOnlyChildRecoveries := true // bool | Returns only child recoveries if passed as true. This filter should always be used along with 'ids' filter.  (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the organizations for which recoveries are to be returned. (optional)
    includeTenants := true // bool | Specifies if objects of all the organizations under the hierarchy of the logged in user's organization should be returned. (optional)
    startTimeUsecs := int64(789) // int64 | Returns the recoveries which are started after the specific time. This value should be in Unix timestamp epoch in microseconds. (optional)
    endTimeUsecs := int64(789) // int64 | Returns the recoveries which are started before the specific time. This value should be in Unix timestamp epoch in microseconds. (optional)
    storageDomainId := int64(789) // int64 | Filter by Storage Domain id. Only recoveries writing data to this Storage Domain will be returned. (optional)
    snapshotTargetType := []string{"SnapshotTargetType_example"} // []string | Specifies the snapshot's target type from which recovery has been performed. (optional)
    archivalTargetType := []string{"ArchivalTargetType_example"} // []string | Specifies the snapshot's archival target type from which recovery has been performed. This parameter applies only if 'snapshotTargetType' is 'Archival'. (optional)
    snapshotEnvironments := []string{"SnapshotEnvironments_example"} // []string | Specifies the list of snapshot environment types to filter Recoveries. If empty, Recoveries related to all environments will be returned. (optional)
    status := []string{"Status_example"} // []string | Specifies the list of run status to filter Recoveries. If empty, Recoveries with all run status will be returned. (optional)
    recoveryActions := []string{"RecoveryActions_example"} // []string | Specifies the list of recovery actions to filter Recoveries. If empty, Recoveries related to all actions will be returned. (optional)

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

    request := onprem.ApiGetRecoveriesRequest{
        Ids: &ids
        ReturnOnlyChildRecoveries: &returnOnlyChildRecoveries
        TenantIds: &tenantIds
        IncludeTenants: &includeTenants
        StartTimeUsecs: &startTimeUsecs
        EndTimeUsecs: &endTimeUsecs
        StorageDomainId: &storageDomainId
        SnapshotTargetType: &snapshotTargetType
        ArchivalTargetType: &archivalTargetType
        SnapshotEnvironments: &snapshotEnvironments
        Status: &status
        RecoveryActions: &recoveryActions
    }

    resp, r, err := api_client.Recovery.GetRecoveries(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Recovery.GetRecoveries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecoveries`: Recoveries
    fmt.Fprintf(os.Stdout, "Response from `Recovery.GetRecoveries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecoveriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** | Filter Recoveries for given ids. | 
 **returnOnlyChildRecoveries** | **bool** | Returns only child recoveries if passed as true. This filter should always be used along with &#39;ids&#39; filter.  | 
 **tenantIds** | **[]string** | TenantIds contains ids of the organizations for which recoveries are to be returned. | 
 **includeTenants** | **bool** | Specifies if objects of all the organizations under the hierarchy of the logged in user&#39;s organization should be returned. | 
 **startTimeUsecs** | **int64** | Returns the recoveries which are started after the specific time. This value should be in Unix timestamp epoch in microseconds. | 
 **endTimeUsecs** | **int64** | Returns the recoveries which are started before the specific time. This value should be in Unix timestamp epoch in microseconds. | 
 **storageDomainId** | **int64** | Filter by Storage Domain id. Only recoveries writing data to this Storage Domain will be returned. | 
 **snapshotTargetType** | **[]string** | Specifies the snapshot&#39;s target type from which recovery has been performed. | 
 **archivalTargetType** | **[]string** | Specifies the snapshot&#39;s archival target type from which recovery has been performed. This parameter applies only if &#39;snapshotTargetType&#39; is &#39;Archival&#39;. | 
 **snapshotEnvironments** | **[]string** | Specifies the list of snapshot environment types to filter Recoveries. If empty, Recoveries related to all environments will be returned. | 
 **status** | **[]string** | Specifies the list of run status to filter Recoveries. If empty, Recoveries with all run status will be returned. | 
 **recoveryActions** | **[]string** | Specifies the list of recovery actions to filter Recoveries. If empty, Recoveries related to all actions will be returned. | 

### Return type

[**Recoveries**](Recoveries.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecoveryById

> Recovery GetRecoveryById(ctx, id).IncludeTenants(includeTenants).Execute()

Get Recovery for a given id.



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
    id := "id_example" // string | Specifies the id of a Recovery.
    includeTenants := true // bool | Specifies if objects of all the organizations under the hierarchy of the logged in user's organization should be returned. (optional)

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

    request := onprem.ApiGetRecoveryByIdRequest{
        Id: &id
        IncludeTenants: &includeTenants
    }

    resp, r, err := api_client.Recovery.GetRecoveryById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Recovery.GetRecoveryById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecoveryById`: Recovery
    fmt.Fprintf(os.Stdout, "Response from `Recovery.GetRecoveryById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id of a Recovery. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecoveryByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeTenants** | **bool** | Specifies if objects of all the organizations under the hierarchy of the logged in user&#39;s organization should be returned. | 

### Return type

[**Recovery**](Recovery.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecoveryDebugLogs

> GetRecoveryDebugLogs(ctx, id).Execute()

Get the debug logs for a particular recovery operation.



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
    id := "id_example" // string | Specifies the id of a Recovery job.

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

    request := onprem.ApiGetRecoveryDebugLogsRequest{
        Id: &id
    }

    resp, r, err := api_client.Recovery.GetRecoveryDebugLogs(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Recovery.GetRecoveryDebugLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id of a Recovery job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecoveryDebugLogsRequest struct via the builder pattern


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


## GetRecoveryErrorsReport

> GetRecoveryErrorsReport(ctx, id).Execute()

Get the CSV of errors/warnings for a given recovery operation.



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
    id := "id_example" // string | Specifies a unique ID of a Recovery.

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

    request := onprem.ApiGetRecoveryErrorsReportRequest{
        Id: &id
    }

    resp, r, err := api_client.Recovery.GetRecoveryErrorsReport(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Recovery.GetRecoveryErrorsReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique ID of a Recovery. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecoveryErrorsReportRequest struct via the builder pattern


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


## TearDownRecoveryById

> TearDownRecoveryById(ctx, id).Execute()

Tear down Recovery for a given id.



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
    id := "id_example" // string | Specifies the id of a Recovery.

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

    request := onprem.ApiTearDownRecoveryByIdRequest{
        Id: &id
    }

    resp, r, err := api_client.Recovery.TearDownRecoveryById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Recovery.TearDownRecoveryById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id of a Recovery. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTearDownRecoveryByIdRequest struct via the builder pattern


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

