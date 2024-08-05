# \RecoveryAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelRecoveryById**](RecoveryAPI.md#CancelRecoveryById) | **Post** /data-protect/recoveries/{id}/cancel | Cancel Recovery for a given id.
[**CreateDownloadFilesAndFoldersRecovery**](RecoveryAPI.md#CreateDownloadFilesAndFoldersRecovery) | **Post** /data-protect/recoveries/download-files-folders | Create a download files and folders recovery.
[**CreateRecovery**](RecoveryAPI.md#CreateRecovery) | **Post** /data-protect/recoveries | Performs a Recovery.
[**DeleteRecoveryCloneTaskById**](RecoveryAPI.md#DeleteRecoveryCloneTaskById) | **Delete** /data-protect/recoveries/clone/{id} | Delete a restore clone task
[**DownloadFilesFromRecovery**](RecoveryAPI.md#DownloadFilesFromRecovery) | **Get** /data-protect/recoveries/{id}/download-files | Download files from the given download file recovery.
[**FetchUptierData**](RecoveryAPI.md#FetchUptierData) | **Get** /data-protect/recoveries/fetch-uptier-data | Fetches the uptier data.
[**GetRecoveries**](RecoveryAPI.md#GetRecoveries) | **Get** /data-protect/recoveries | Lists the Recoveries.
[**GetRecoveryById**](RecoveryAPI.md#GetRecoveryById) | **Get** /data-protect/recoveries/{id} | Get Recovery for a given id.
[**GetRecoveryDebugLogs**](RecoveryAPI.md#GetRecoveryDebugLogs) | **Get** /data-protect/recoveries/{id}/debug-logs | Get the debug logs for a particular recovery operation.
[**GetRecoveryErrorsReport**](RecoveryAPI.md#GetRecoveryErrorsReport) | **Get** /data-protect/recoveries/{id}/download-messages | Get the CSV of errors/warnings for a given recovery operation.
[**TearDownRecoveryById**](RecoveryAPI.md#TearDownRecoveryById) | **Post** /data-protect/recoveries/{id}/tear-down | Tear down Recovery for a given id.
[**VirtualDiskInformation**](RecoveryAPI.md#VirtualDiskInformation) | **Get** /data-protect/recoveries/virtual-disks | Fetches information of virtual disks



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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := "id_example" // string | Specifies the id of a Recovery.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RecoveryAPI.CancelRecoveryById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecoveryAPI.CancelRecoveryById``: %v\n", err)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewDownloadFilesAndFoldersRequestParams("Name_example", *openapiclient.NewCommonRecoverObjectSnapshotParams("SnapshotId_example")) // DownloadFilesAndFoldersRequestParams | Specifies the parameters to create a download files and folder recovery.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecoveryAPI.CreateDownloadFilesAndFoldersRecovery(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecoveryAPI.CreateDownloadFilesAndFoldersRecovery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDownloadFilesAndFoldersRecovery`: Recovery
	fmt.Fprintf(os.Stdout, "Response from `RecoveryAPI.CreateDownloadFilesAndFoldersRecovery`: %v\n", resp)
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

> Recovery CreateRecovery(ctx).Body(body).RequestInitiatorType(requestInitiatorType).Execute()

Performs a Recovery.



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
	body := *openapiclient.NewCreateRecoveryRequest("Name_example", "SnapshotEnvironment_example") // CreateRecoveryRequest | Specifies the parameters to create a Recovery.
	requestInitiatorType := "requestInitiatorType_example" // string | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecoveryAPI.CreateRecovery(context.Background()).Body(body).RequestInitiatorType(requestInitiatorType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecoveryAPI.CreateRecovery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRecovery`: Recovery
	fmt.Fprintf(os.Stdout, "Response from `RecoveryAPI.CreateRecovery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRecoveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateRecoveryRequest**](CreateRecoveryRequest.md) | Specifies the parameters to create a Recovery. | 
 **requestInitiatorType** | **string** | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | 

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


## DeleteRecoveryCloneTaskById

> DeleteRecoveryCloneTaskById(ctx, id).Execute()

Delete a restore clone task



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
	id := int64(789) // int64 | Specifies a unique id of the Clone Task to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RecoveryAPI.DeleteRecoveryCloneTaskById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecoveryAPI.DeleteRecoveryCloneTaskById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the Clone Task to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecoveryCloneTaskByIdRequest struct via the builder pattern


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


## DownloadFilesFromRecovery

> DownloadFilesFromRecovery(ctx, id).StartOffset(startOffset).Length(length).FileType(fileType).SourceName(sourceName).StartTime(startTime).IncludeTenants(includeTenants).Execute()

Download files from the given download file recovery.



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
	id := "id_example" // string | Specifies the id of a Recovery.
	startOffset := int64(789) // int64 | Specifies the start offset of file chunk to be downloaded. (optional)
	length := int64(789) // int64 | Specifies the length of bytes to download. This can not be greater than 8MB (8388608 byets) (optional)
	fileType := "fileType_example" // string | Specifies the downloaded type, i.e: error, success_files_list (optional)
	sourceName := "sourceName_example" // string | Specifies the name of the source on which restore is done (optional)
	startTime := "startTime_example" // string | Specifies the start time of restore task (optional)
	includeTenants := true // bool | Specifies if objects of all the organizations under the hierarchy of the logged in user's organization should be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RecoveryAPI.DownloadFilesFromRecovery(context.Background(), id).StartOffset(startOffset).Length(length).FileType(fileType).SourceName(sourceName).StartTime(startTime).IncludeTenants(includeTenants).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecoveryAPI.DownloadFilesFromRecovery``: %v\n", err)
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
 **fileType** | **string** | Specifies the downloaded type, i.e: error, success_files_list | 
 **sourceName** | **string** | Specifies the name of the source on which restore is done | 
 **startTime** | **string** | Specifies the start time of restore task | 
 **includeTenants** | **bool** | Specifies if objects of all the organizations under the hierarchy of the logged in user&#39;s organization should be returned. | 

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


## FetchUptierData

> FetchUptierDataResponse FetchUptierData(ctx).ArchiveUId(archiveUId).Execute()

Fetches the uptier data.



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
	archiveUId := "archiveUId_example" // string | Archive UID of the current restore.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecoveryAPI.FetchUptierData(context.Background()).ArchiveUId(archiveUId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecoveryAPI.FetchUptierData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchUptierData`: FetchUptierDataResponse
	fmt.Fprintf(os.Stdout, "Response from `RecoveryAPI.FetchUptierData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchUptierDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **archiveUId** | **string** | Archive UID of the current restore. | 

### Return type

[**FetchUptierDataResponse**](FetchUptierDataResponse.md)

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
	openapiclient "github.com/cohesity/go-sdk"
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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecoveryAPI.GetRecoveries(context.Background()).Ids(ids).ReturnOnlyChildRecoveries(returnOnlyChildRecoveries).TenantIds(tenantIds).IncludeTenants(includeTenants).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).StorageDomainId(storageDomainId).SnapshotTargetType(snapshotTargetType).ArchivalTargetType(archivalTargetType).SnapshotEnvironments(snapshotEnvironments).Status(status).RecoveryActions(recoveryActions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecoveryAPI.GetRecoveries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecoveries`: Recoveries
	fmt.Fprintf(os.Stdout, "Response from `RecoveryAPI.GetRecoveries`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := "id_example" // string | Specifies the id of a Recovery.
	includeTenants := true // bool | Specifies if objects of all the organizations under the hierarchy of the logged in user's organization should be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecoveryAPI.GetRecoveryById(context.Background(), id).IncludeTenants(includeTenants).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecoveryAPI.GetRecoveryById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecoveryById`: Recovery
	fmt.Fprintf(os.Stdout, "Response from `RecoveryAPI.GetRecoveryById`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := "id_example" // string | Specifies the id of a Recovery job.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RecoveryAPI.GetRecoveryDebugLogs(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecoveryAPI.GetRecoveryDebugLogs``: %v\n", err)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := "id_example" // string | Specifies a unique ID of a Recovery.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RecoveryAPI.GetRecoveryErrorsReport(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecoveryAPI.GetRecoveryErrorsReport``: %v\n", err)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := "id_example" // string | Specifies the id of a Recovery.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RecoveryAPI.TearDownRecoveryById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecoveryAPI.TearDownRecoveryById``: %v\n", err)
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


## VirtualDiskInformation

> VirtualDiskInformationResponseParams VirtualDiskInformation(ctx).ClusterId(clusterId).ClusterIncarnationId(clusterIncarnationId).JobId(jobId).ObjectId(objectId).SnapshotId(snapshotId).PointInTimeUsecs(pointInTimeUsecs).VaultId(vaultId).VaultName(vaultName).VaultType(vaultType).Execute()

Fetches information of virtual disks



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
	clusterId := int64(789) // int64 | Specifies the Cohesity Cluster id where the Job was created.
	clusterIncarnationId := int64(789) // int64 | Specifies the incarnation id of the Cohesity Cluster where the Job was created.
	jobId := int64(789) // int64 | Specifies the id of the Job that captured the snapshot.
	objectId := int64(789) // int64 | Specifies the Id of the Protection Source object.
	snapshotId := "snapshotId_example" // string | Specifies the snapshot id. (optional)
	pointInTimeUsecs := int64(789) // int64 | Specifies the Id of the vault where snapshot was taken (optional)
	vaultId := int64(789) // int64 | Specifies the Id of the vault where snapshot was taken (optional)
	vaultName := "vaultName_example" // string | Specifies the name of the vault where snapshot was taken (optional)
	vaultType := "vaultType_example" // string | Specifies the External Target type. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecoveryAPI.VirtualDiskInformation(context.Background()).ClusterId(clusterId).ClusterIncarnationId(clusterIncarnationId).JobId(jobId).ObjectId(objectId).SnapshotId(snapshotId).PointInTimeUsecs(pointInTimeUsecs).VaultId(vaultId).VaultName(vaultName).VaultType(vaultType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecoveryAPI.VirtualDiskInformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VirtualDiskInformation`: VirtualDiskInformationResponseParams
	fmt.Fprintf(os.Stdout, "Response from `RecoveryAPI.VirtualDiskInformation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVirtualDiskInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **int64** | Specifies the Cohesity Cluster id where the Job was created. | 
 **clusterIncarnationId** | **int64** | Specifies the incarnation id of the Cohesity Cluster where the Job was created. | 
 **jobId** | **int64** | Specifies the id of the Job that captured the snapshot. | 
 **objectId** | **int64** | Specifies the Id of the Protection Source object. | 
 **snapshotId** | **string** | Specifies the snapshot id. | 
 **pointInTimeUsecs** | **int64** | Specifies the Id of the vault where snapshot was taken | 
 **vaultId** | **int64** | Specifies the Id of the vault where snapshot was taken | 
 **vaultName** | **string** | Specifies the name of the vault where snapshot was taken | 
 **vaultType** | **string** | Specifies the External Target type. | 

### Return type

[**VirtualDiskInformationResponseParams**](VirtualDiskInformationResponseParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

