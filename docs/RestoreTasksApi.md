# \RestoreTasksApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdObjectsRestoreStatus**](RestoreTasksApi.md#AdObjectsRestoreStatus) | **Get** /public/restore/adObjects/status | 
[**CancelRestoreTask**](RestoreTasksApi.md#CancelRestoreTask) | **Put** /public/restore/tasks/cancel/{id} | Cancel a recover or clone task with specified id.
[**CompareAdObjects**](RestoreTasksApi.md#CompareAdObjects) | **Post** /public/restore/adObjectAttributes | Compares the AD Object from both Snapshot and Production AD and returns the attributes with status whether they differ or not.
[**CreateApplicationsCloneTask**](RestoreTasksApi.md#CreateApplicationsCloneTask) | **Post** /public/restore/applicationsClone | Create a Restore Task for cloning Applications like SQL Databases.
[**CreateApplicationsRecoverTask**](RestoreTasksApi.md#CreateApplicationsRecoverTask) | **Post** /public/restore/applicationsRecover | Create a Restore Task for recovering Applications like SQL Databases.
[**CreateCloneTask**](RestoreTasksApi.md#CreateCloneTask) | **Post** /public/restore/clone | Create a Restore Task for cloning VMs or a View.
[**CreateDeployTask**](RestoreTasksApi.md#CreateDeployTask) | **Post** /public/restore/deploy | Create a Restore Task for deploying VMs or a View on cloud.
[**CreateDownloadFilesAndFolders**](RestoreTasksApi.md#CreateDownloadFilesAndFolders) | **Post** /public/restore/downloadFilesAndFolders | Create a Download Task for downloading files and folders.
[**CreateRecoverTask**](RestoreTasksApi.md#CreateRecoverTask) | **Post** /public/restore/recover | Create a Restore Task for recovering VMs or instantly mounting volumes.
[**CreateRestoreFilesTask**](RestoreTasksApi.md#CreateRestoreFilesTask) | **Post** /public/restore/files | Create a Restore Task for recovering files and folders.
[**GetAdDomainRootTopology**](RestoreTasksApi.md#GetAdDomainRootTopology) | **Get** /public/restore/adDomainRootTopology | Gets Root Topology for an AD Domain.
[**GetAdObjects**](RestoreTasksApi.md#GetAdObjects) | **Get** /public/restore/adObjects/searchResults | 
[**GetFileFstatInformation**](RestoreTasksApi.md#GetFileFstatInformation) | **Get** /public/restore/files/fstats | Get the fstat information about file provided using query parameters.
[**GetFileSnapshotsInformation**](RestoreTasksApi.md#GetFileSnapshotsInformation) | **Get** /public/restore/files/snapshotsInformation | Get the information about snapshots that contain the specified file or folder. In addition, information about the file or folder is provided.
[**GetOneDriveDocuments**](RestoreTasksApi.md#GetOneDriveDocuments) | **Get** /public/restore/office365/onedrive/documents | Returns the OneDrive files and folders.
[**GetOutlookEmails**](RestoreTasksApi.md#GetOutlookEmails) | **Get** /public/restore/office365/outlook/emails | Returns the Outlook emails and folders containing emails.
[**GetRestorePointsForTimeRange**](RestoreTasksApi.md#GetRestorePointsForTimeRange) | **Post** /public/restore/pointsForTimeRange | List Restore Points in a give time range.
[**GetRestoreTaskById**](RestoreTasksApi.md#GetRestoreTaskById) | **Get** /public/restore/tasks/{id} | List details about a single Restore Task.
[**GetRestoreTasks**](RestoreTasksApi.md#GetRestoreTasks) | **Get** /public/restore/tasks | List the Restore Tasks filtered by the specified parameters.
[**GetSharepointDocuments**](RestoreTasksApi.md#GetSharepointDocuments) | **Get** /public/restore/office365/sharepoint/documents | Returns the Sharepoint Site&#39;s files and folders.
[**GetVirtualDiskInformation**](RestoreTasksApi.md#GetVirtualDiskInformation) | **Get** /public/restore/virtualDiskInformation | Fetches information of virtual disk.
[**GetVmDirectoryList**](RestoreTasksApi.md#GetVmDirectoryList) | **Get** /public/restore/vms/directoryList | Get the directory list based on the given directory name and other query parameters.
[**GetVmVolumesInformation**](RestoreTasksApi.md#GetVmVolumesInformation) | **Get** /public/restore/vms/volumesInformation | Get information about the logical volumes found in a VM.
[**ListStorageProfiles**](RestoreTasksApi.md#ListStorageProfiles) | **Get** /public/virtualDatacenters/{id}/storageProfiles | Fetches information of virtual disk.
[**PublicDestroyCloneTask**](RestoreTasksApi.md#PublicDestroyCloneTask) | **Delete** /public/restore/clone/{id} | Destroy a clone task with specified id.
[**SearchAdObjects**](RestoreTasksApi.md#SearchAdObjects) | **Get** /public/restore/adObjects | Searches for AD Objects in both Production and Snapshot AD from given search base dn and offset.
[**SearchObjects**](RestoreTasksApi.md#SearchObjects) | **Get** /public/restore/objects | Find backup objects that match the specified search and filter criteria on the Cohesity Cluster.
[**SearchProductionAdObjects**](RestoreTasksApi.md#SearchProductionAdObjects) | **Post** /public/restore/adObjects | Searches for AD Objects that match the list of object guids, sam account names and distinguished names provided in the request.
[**SearchRestoredFiles**](RestoreTasksApi.md#SearchRestoredFiles) | **Get** /public/restore/files | Search for files and folders to recover that match the specified search and filter criteria on the Cohesity Cluster.
[**UpdateRestoreTask**](RestoreTasksApi.md#UpdateRestoreTask) | **Put** /public/restore/recover | 



## AdObjectsRestoreStatus

> AdObjectsRestoreStatus AdObjectsRestoreStatus(ctx).RestoreTaskId(restoreTaskId).Execute()





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
    restoreTaskId := int64(789) // int64 | Specifies the restoreTaskId corresponding to which we need to get information about the restore of the AD objects. (optional)

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

    request := cohesitysdk.ApiAdObjectsRestoreStatusRequest{
        restoreTaskId: &RestoreTaskId
    }

    resp, r, err := api_client.RestoreTasksApi.AdObjectsRestoreStatus(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreTasksApi.AdObjectsRestoreStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdObjectsRestoreStatus`: AdObjectsRestoreStatus
    fmt.Fprintf(os.Stdout, "Response from `RestoreTasksApi.AdObjectsRestoreStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdObjectsRestoreStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restoreTaskId** | **int64** | Specifies the restoreTaskId corresponding to which we need to get information about the restore of the AD objects. | 

### Return type

[**AdObjectsRestoreStatus**](AdObjectsRestoreStatus.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelRestoreTask

> CancelRestoreTask(ctx, id).Execute()

Cancel a recover or clone task with specified id.

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
    id := int64(789) // int64 | Specifies a unique id for the Restore Task.

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

    request := cohesitysdk.ApiCancelRestoreTaskRequest{
        id: &Id
    }

    resp, r, err := api_client.RestoreTasksApi.CancelRestoreTask(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreTasksApi.CancelRestoreTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id for the Restore Task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelRestoreTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompareAdObjects

> []ComparedADObject CompareAdObjects(ctx).Body(body).Execute()

Compares the AD Object from both Snapshot and Production AD and returns the attributes with status whether they differ or not.



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
    body := *openapiclient.NewCompareAdObjectsRequest(NullableInt64(123), []openapiclient.GuidPair{*openapiclient.NewGuidPair()}) // CompareAdObjectsRequest | Specifies the Request to compare the AD Objects from both Snapshot and Production AD.

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

    request := cohesitysdk.ApiCompareAdObjectsRequest{
        body: &Body
    }

    resp, r, err := api_client.RestoreTasksApi.CompareAdObjects(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreTasksApi.CompareAdObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CompareAdObjects`: []ComparedADObject
    fmt.Fprintf(os.Stdout, "Response from `RestoreTasksApi.CompareAdObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompareAdObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CompareAdObjectsRequest**](CompareAdObjectsRequest.md) | Specifies the Request to compare the AD Objects from both Snapshot and Production AD. | 

### Return type

[**[]ComparedADObject**](ComparedADObject.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplicationsCloneTask

> RestoreTask CreateApplicationsCloneTask(ctx).Body(body).Execute()

Create a Restore Task for cloning Applications like SQL Databases.



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
    body := *openapiclient.NewApplicationsRestoreTaskRequest("ApplicationEnvironment_example", *openapiclient.NewRestoreObjectDetails(), "Name_example") // ApplicationsRestoreTaskRequest | Request to create a Restore Task for cloning Applications like SQL DB.

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

    request := cohesitysdk.ApiCreateApplicationsCloneTaskRequest{
        body: &Body
    }

    resp, r, err := api_client.RestoreTasksApi.CreateApplicationsCloneTask(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreTasksApi.CreateApplicationsCloneTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplicationsCloneTask`: RestoreTask
    fmt.Fprintf(os.Stdout, "Response from `RestoreTasksApi.CreateApplicationsCloneTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationsCloneTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApplicationsRestoreTaskRequest**](ApplicationsRestoreTaskRequest.md) | Request to create a Restore Task for cloning Applications like SQL DB. | 

### Return type

[**RestoreTask**](RestoreTask.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplicationsRecoverTask

> RestoreTask CreateApplicationsRecoverTask(ctx).Body(body).Execute()

Create a Restore Task for recovering Applications like SQL Databases.



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
    body := *openapiclient.NewApplicationsRestoreTaskRequest("ApplicationEnvironment_example", *openapiclient.NewRestoreObjectDetails(), "Name_example") // ApplicationsRestoreTaskRequest | Request to create a Restore Task for recovering Applications like SQL DB. volumes to mount points.

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

    request := cohesitysdk.ApiCreateApplicationsRecoverTaskRequest{
        body: &Body
    }

    resp, r, err := api_client.RestoreTasksApi.CreateApplicationsRecoverTask(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreTasksApi.CreateApplicationsRecoverTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplicationsRecoverTask`: RestoreTask
    fmt.Fprintf(os.Stdout, "Response from `RestoreTasksApi.CreateApplicationsRecoverTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationsRecoverTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApplicationsRestoreTaskRequest**](ApplicationsRestoreTaskRequest.md) | Request to create a Restore Task for recovering Applications like SQL DB. volumes to mount points. | 

### Return type

[**RestoreTask**](RestoreTask.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCloneTask

> RestoreTask CreateCloneTask(ctx).Body(body).Execute()

Create a Restore Task for cloning VMs or a View.



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
    body := *openapiclient.NewCloneTaskRequest("Name_example", "Type_example") // CloneTaskRequest | Request to create a Restore Task for cloning VMs or a View.

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

    request := cohesitysdk.ApiCreateCloneTaskRequest{
        body: &Body
    }

    resp, r, err := api_client.RestoreTasksApi.CreateCloneTask(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreTasksApi.CreateCloneTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCloneTask`: RestoreTask
    fmt.Fprintf(os.Stdout, "Response from `RestoreTasksApi.CreateCloneTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCloneTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CloneTaskRequest**](CloneTaskRequest.md) | Request to create a Restore Task for cloning VMs or a View. | 

### Return type

[**RestoreTask**](RestoreTask.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeployTask

> RestoreTask CreateDeployTask(ctx).Body(body).Execute()

Create a Restore Task for deploying VMs or a View on cloud.



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
    body := *openapiclient.NewDeployTaskRequest("Name_example") // DeployTaskRequest | Request to create a Restore Task for deploying VMs or a View on cloud.

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

    request := cohesitysdk.ApiCreateDeployTaskRequest{
        body: &Body
    }

    resp, r, err := api_client.RestoreTasksApi.CreateDeployTask(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreTasksApi.CreateDeployTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeployTask`: RestoreTask
    fmt.Fprintf(os.Stdout, "Response from `RestoreTasksApi.CreateDeployTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeployTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DeployTaskRequest**](DeployTaskRequest.md) | Request to create a Restore Task for deploying VMs or a View on cloud. | 

### Return type

[**RestoreTask**](RestoreTask.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDownloadFilesAndFolders

> RestoreTask CreateDownloadFilesAndFolders(ctx).Body(body).Execute()

Create a Download Task for downloading files and folders.



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
    body := *openapiclient.NewDownloadFilesAndFoldersParams("Name_example") // DownloadFilesAndFoldersParams | Request to create a task for downloading list of files or folders.

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

    request := cohesitysdk.ApiCreateDownloadFilesAndFoldersRequest{
        body: &Body
    }

    resp, r, err := api_client.RestoreTasksApi.CreateDownloadFilesAndFolders(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreTasksApi.CreateDownloadFilesAndFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDownloadFilesAndFolders`: RestoreTask
    fmt.Fprintf(os.Stdout, "Response from `RestoreTasksApi.CreateDownloadFilesAndFolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDownloadFilesAndFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DownloadFilesAndFoldersParams**](DownloadFilesAndFoldersParams.md) | Request to create a task for downloading list of files or folders. | 

### Return type

[**RestoreTask**](RestoreTask.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRecoverTask

> RestoreTask CreateRecoverTask(ctx).Body(body).Execute()

Create a Restore Task for recovering VMs or instantly mounting volumes.



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
    body := *openapiclient.NewRecoverTaskRequest("Name_example", "Type_example") // RecoverTaskRequest | Request to create a Restore Task for recovering VMs or mounting volumes to mount points.

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

    request := cohesitysdk.ApiCreateRecoverTaskRequest{
        body: &Body
    }

    resp, r, err := api_client.RestoreTasksApi.CreateRecoverTask(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreTasksApi.CreateRecoverTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRecoverTask`: RestoreTask
    fmt.Fprintf(os.Stdout, "Response from `RestoreTasksApi.CreateRecoverTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRecoverTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RecoverTaskRequest**](RecoverTaskRequest.md) | Request to create a Restore Task for recovering VMs or mounting volumes to mount points. | 

### Return type

[**RestoreTask**](RestoreTask.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRestoreFilesTask

> RestoreTask CreateRestoreFilesTask(ctx).Body(body).Execute()

Create a Restore Task for recovering files and folders.



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
    body := *openapiclient.NewRestoreFilesTaskRequest() // RestoreFilesTaskRequest | Request to create a Restore Task for recovering files or folders.

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

    request := cohesitysdk.ApiCreateRestoreFilesTaskRequest{
        body: &Body
    }

    resp, r, err := api_client.RestoreTasksApi.CreateRestoreFilesTask(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreTasksApi.CreateRestoreFilesTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRestoreFilesTask`: RestoreTask
    fmt.Fprintf(os.Stdout, "Response from `RestoreTasksApi.CreateRestoreFilesTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRestoreFilesTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RestoreFilesTaskRequest**](RestoreFilesTaskRequest.md) | Request to create a Restore Task for recovering files or folders. | 

### Return type

[**RestoreTask**](RestoreTask.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdDomainRootTopology

> []AdRootTopologyObject GetAdDomainRootTopology(ctx).RestoreTaskId(restoreTaskId).Execute()

Gets Root Topology for an AD Domain.



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
    restoreTaskId := int64(789) // int64 | Specifies the restoreTaskId corresponding to which we need to get the ad topology.

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

    request := cohesitysdk.ApiGetAdDomainRootTopologyRequest{
        restoreTaskId: &RestoreTaskId
    }

    resp, r, err := api_client.RestoreTasksApi.GetAdDomainRootTopology(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreTasksApi.GetAdDomainRootTopology``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdDomainRootTopology`: []AdRootTopologyObject
    fmt.Fprintf(os.Stdout, "Response from `RestoreTasksApi.GetAdDomainRootTopology`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdDomainRootTopologyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restoreTaskId** | **int64** | Specifies the restoreTaskId corresponding to which we need to get the ad topology. | 

### Return type

[**[]AdRootTopologyObject**](AdRootTopologyObject.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdObjects

> FileSearchResults GetAdObjects(ctx).Name(name).SamAccountName(samAccountName).ObjectType(objectType).Email(email).RegisteredSourceIds(registeredSourceIds).JobIds(jobIds).ViewBoxIds(viewBoxIds).Domain(domain).TenantId(tenantId).AllUnderHierarchy(allUnderHierarchy).Execute()





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
    name := "name_example" // string | Specifies the name of the AD object. (optional)
    samAccountName := "samAccountName_example" // string | Specifies the sam account name of the AD object. (optional)
    objectType := "objectType_example" // string | Specifies the type of the AD Object. The type may be user, computer, group or ou(organizational unit). (optional)
    email := "email_example" // string | Specifies the email of the AD object of type user or group. (optional)
    registeredSourceIds := []int64{int64(123)} // []int64 | Specifies the Active Directory Application Server Ids which contains the AD objects. (optional)
    jobIds := []int64{int64(123)} // []int64 | Specifies the protection job Ids which have backed up Active Directory Application Server. (optional)
    viewBoxIds := []int64{int64(123)} // []int64 | Filter by a list of Domains (View Boxes) ids. Only items stored in the listed Domains (View Boxes) are returned. (optional)
    domain := "domain_example" // string | domain of the AD object. (optional)
    tenantId := "tenantId_example" // string | TenantId specifies the tenant whose action resulted in the audit log. (optional)
    allUnderHierarchy := true // bool | AllUnderHierarchy specifies if logs of all the tenants under the hierarchy of tenant with id TenantId should be returned. (optional)

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

    request := cohesitysdk.ApiGetAdObjectsRequest{
        name: &Name
        samAccountName: &SamAccountName
        objectType: &ObjectType
        email: &Email
        registeredSourceIds: &RegisteredSourceIds
        jobIds: &JobIds
        viewBoxIds: &ViewBoxIds
        domain: &Domain
        tenantId: &TenantId
        allUnderHierarchy: &AllUnderHierarchy
    }

    resp, r, err := api_client.RestoreTasksApi.GetAdObjects(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreTasksApi.GetAdObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdObjects`: FileSearchResults
    fmt.Fprintf(os.Stdout, "Response from `RestoreTasksApi.GetAdObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Specifies the name of the AD object. | 
 **samAccountName** | **string** | Specifies the sam account name of the AD object. | 
 **objectType** | **string** | Specifies the type of the AD Object. The type may be user, computer, group or ou(organizational unit). | 
 **email** | **string** | Specifies the email of the AD object of type user or group. | 
 **registeredSourceIds** | **[]int64** | Specifies the Active Directory Application Server Ids which contains the AD objects. | 
 **jobIds** | **[]int64** | Specifies the protection job Ids which have backed up Active Directory Application Server. | 
 **viewBoxIds** | **[]int64** | Filter by a list of Domains (View Boxes) ids. Only items stored in the listed Domains (View Boxes) are returned. | 
 **domain** | **string** | domain of the AD object. | 
 **tenantId** | **string** | TenantId specifies the tenant whose action resulted in the audit log. | 
 **allUnderHierarchy** | **bool** | AllUnderHierarchy specifies if logs of all the tenants under the hierarchy of tenant with id TenantId should be returned. | 

### Return type

[**FileSearchResults**](FileSearchResults.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileFstatInformation

> FileFstatResult GetFileFstatInformation(ctx).JobId(jobId).JobUidObjectId(jobUidObjectId).EntityId(entityId).JobInstanceId(jobInstanceId).JobStartTimeUsecs(jobStartTimeUsecs).FilePath(filePath).AttemptNum(attemptNum).VolumeName(volumeName).ViewBoxId(viewBoxId).ViewName(viewName).VolumeInfoCookie(volumeInfoCookie).UseLibrarian(useLibrarian).Execute()

Get the fstat information about file provided using query parameters.

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
    jobId := int64(789) // int64 | JobId is the id of the local job that took the snapshot, which may or may not match the JobUidObjectId below depending on whether the object originally belonged to this local job or to a different remote job.
    jobUidObjectId := int64(789) // int64 | JobUidObjectId is the globally unique id of the job that the object originally belonged to. If this object originally belonged to a job from a remote cluster, this field will contain the JobId of the remote job, else it will contain the JobId of the local job.
    entityId := int64(789) // int64 | EntityId is the Id of the VM.
    jobInstanceId := int64(789) // int64 | JobInstanceId is the Id of the job run that backed up the entity.
    jobStartTimeUsecs := int64(789) // int64 | JobStartTimeUsecs is the start time in usecs of the job run that backed up the entity.
    filePath := "filePath_example" // string | FilePath is the full path of the file or directory whose stat needed.
    attemptNum := int64(789) // int64 | AttemptNum is the attempt number of the run that successfully created the snapshot. (optional)
    volumeName := "volumeName_example" // string | VolumeName is the name of the volume that needs to be browsed. This should match the name returned in VolumeInfo. (optional)
    viewBoxId := int64(789) // int64 | Id of the View Box if a View is being browsed. (optional)
    viewName := "viewName_example" // string | Name of the View if a View is being browsed. (optional)
    volumeInfoCookie := int32(56) // int32 | VolumeInfoCookie is the cookie to be passed in calls to reading a VM dir for this volume. (optional)
    useLibrarian := true // bool | Specifies whether to use Librarian for file stat. This will be true if the browse is enabled via librarian. (optional)

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

    request := cohesitysdk.ApiGetFileFstatInformationRequest{
        jobId: &JobId
        jobUidObjectId: &JobUidObjectId
        entityId: &EntityId
        jobInstanceId: &JobInstanceId
        jobStartTimeUsecs: &JobStartTimeUsecs
        filePath: &FilePath
        attemptNum: &AttemptNum
        volumeName: &VolumeName
        viewBoxId: &ViewBoxId
        viewName: &ViewName
        volumeInfoCookie: &VolumeInfoCookie
        useLibrarian: &UseLibrarian
    }

    resp, r, err := api_client.RestoreTasksApi.GetFileFstatInformation(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreTasksApi.GetFileFstatInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFileFstatInformation`: FileFstatResult
    fmt.Fprintf(os.Stdout, "Response from `RestoreTasksApi.GetFileFstatInformation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFileFstatInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobId** | **int64** | JobId is the id of the local job that took the snapshot, which may or may not match the JobUidObjectId below depending on whether the object originally belonged to this local job or to a different remote job. | 
 **jobUidObjectId** | **int64** | JobUidObjectId is the globally unique id of the job that the object originally belonged to. If this object originally belonged to a job from a remote cluster, this field will contain the JobId of the remote job, else it will contain the JobId of the local job. | 
 **entityId** | **int64** | EntityId is the Id of the VM. | 
 **jobInstanceId** | **int64** | JobInstanceId is the Id of the job run that backed up the entity. | 
 **jobStartTimeUsecs** | **int64** | JobStartTimeUsecs is the start time in usecs of the job run that backed up the entity. | 
 **filePath** | **string** | FilePath is the full path of the file or directory whose stat needed. | 
 **attemptNum** | **int64** | AttemptNum is the attempt number of the run that successfully created the snapshot. | 
 **volumeName** | **string** | VolumeName is the name of the volume that needs to be browsed. This should match the name returned in VolumeInfo. | 
 **viewBoxId** | **int64** | Id of the View Box if a View is being browsed. | 
 **viewName** | **string** | Name of the View if a View is being browsed. | 
 **volumeInfoCookie** | **int32** | VolumeInfoCookie is the cookie to be passed in calls to reading a VM dir for this volume. | 
 **useLibrarian** | **bool** | Specifies whether to use Librarian for file stat. This will be true if the browse is enabled via librarian. | 

### Return type

[**FileFstatResult**](FileFstatResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileSnapshotsInformation

> []FileSnapshotInformation GetFileSnapshotsInformation(ctx).JobId(jobId).ClusterId(clusterId).ClusterIncarnationId(clusterIncarnationId).SourceId(sourceId).Filename(filename).Execute()

Get the information about snapshots that contain the specified file or folder. In addition, information about the file or folder is provided.

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
    jobId := int64(789) // int64 | Specifies the id of the Job that captured the snapshots. These snapshots are searched for the specified files or folders. This field is required.
    clusterId := int64(789) // int64 | Specifies the Cohesity Cluster id where the Job was created. This field is required.
    clusterIncarnationId := int64(789) // int64 | Specifies the incarnation id of the Cohesity Cluster where the Job was created. An incarnation id is generated when a Cohesity Cluster is initially created. This field is required.
    sourceId := int64(789) // int64 | Specifies the id of the Protection Source object (such as a VM) to search. When a Job Run executes, snapshots of the specified Protection Source object are captured. This operation searches the snapshots of the object for the file or folder. This field is required.
    filename := "filename_example" // string | Specifies the name of the file or folder to find in the snapshots. This field is required.

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

    request := cohesitysdk.ApiGetFileSnapshotsInformationRequest{
        jobId: &JobId
        clusterId: &ClusterId
        clusterIncarnationId: &ClusterIncarnationId
        sourceId: &SourceId
        filename: &Filename
    }

    resp, r, err := api_client.RestoreTasksApi.GetFileSnapshotsInformation(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreTasksApi.GetFileSnapshotsInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFileSnapshotsInformation`: []FileSnapshotInformation
    fmt.Fprintf(os.Stdout, "Response from `RestoreTasksApi.GetFileSnapshotsInformation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFileSnapshotsInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobId** | **int64** | Specifies the id of the Job that captured the snapshots. These snapshots are searched for the specified files or folders. This field is required. | 
 **clusterId** | **int64** | Specifies the Cohesity Cluster id where the Job was created. This field is required. | 
 **clusterIncarnationId** | **int64** | Specifies the incarnation id of the Cohesity Cluster where the Job was created. An incarnation id is generated when a Cohesity Cluster is initially created. This field is required. | 
 **sourceId** | **int64** | Specifies the id of the Protection Source object (such as a VM) to search. When a Job Run executes, snapshots of the specified Protection Source object are captured. This operation searches the snapshots of the object for the file or folder. This field is required. | 
 **filename** | **string** | Specifies the name of the file or folder to find in the snapshots. This field is required. | 

### Return type

[**[]FileSnapshotInformation**](FileSnapshotInformation.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOneDriveDocuments

> FileSearchResults GetOneDriveDocuments(ctx).TenantId(tenantId).AllUnderHierarchy(allUnderHierarchy).DocumentName(documentName).DomainIds(domainIds).MailboxIds(mailboxIds).ProtectionJobIds(protectionJobIds).Execute()

Returns the OneDrive files and folders.



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
    tenantId := "tenantId_example" // string | TenantId specifies the tenant whose action resulted in the audit log. (optional)
    allUnderHierarchy := true // bool | AllUnderHierarchy specifies if logs of all the tenants under the hierarchy of tenant with id TenantId should be returned. (optional)
    documentName := "documentName_example" // string | Specifies the document(file/folder) name. (optional)
    domainIds := []int64{int64(123)} // []int64 | Specifies the domain Ids in which Users' OneDrives are registered. (optional)
    mailboxIds := []int64{int64(123)} // []int64 | Specifies the Office365 User Ids which is teh owner of the OneDrive. (optional)
    protectionJobIds := []int64{int64(123)} // []int64 | Specifies the protection job Ids which have backed up mailbox(es) continaing emails/folders. (optional)

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

    request := cohesitysdk.ApiGetOneDriveDocumentsRequest{
        tenantId: &TenantId
        allUnderHierarchy: &AllUnderHierarchy
        documentName: &DocumentName
        domainIds: &DomainIds
        mailboxIds: &MailboxIds
        protectionJobIds: &ProtectionJobIds
    }

    resp, r, err := api_client.RestoreTasksApi.GetOneDriveDocuments(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreTasksApi.GetOneDriveDocuments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOneDriveDocuments`: FileSearchResults
    fmt.Fprintf(os.Stdout, "Response from `RestoreTasksApi.GetOneDriveDocuments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOneDriveDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | TenantId specifies the tenant whose action resulted in the audit log. | 
 **allUnderHierarchy** | **bool** | AllUnderHierarchy specifies if logs of all the tenants under the hierarchy of tenant with id TenantId should be returned. | 
 **documentName** | **string** | Specifies the document(file/folder) name. | 
 **domainIds** | **[]int64** | Specifies the domain Ids in which Users&#39; OneDrives are registered. | 
 **mailboxIds** | **[]int64** | Specifies the Office365 User Ids which is teh owner of the OneDrive. | 
 **protectionJobIds** | **[]int64** | Specifies the protection job Ids which have backed up mailbox(es) continaing emails/folders. | 

### Return type

[**FileSearchResults**](FileSearchResults.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutlookEmails

> FileSearchResults GetOutlookEmails(ctx).HasAttachments(hasAttachments).SenderAddress(senderAddress).RecipientAddresses(recipientAddresses).CcRecipientAddresses(ccRecipientAddresses).BccRecipientAddresses(bccRecipientAddresses).SentTimeSeconds(sentTimeSeconds).ReceivedTimeSeconds(receivedTimeSeconds).ReceivedStartTime(receivedStartTime).ReceivedEndTime(receivedEndTime).EmailSubject(emailSubject).FolderName(folderName).ShowOnlyEmailFolders(showOnlyEmailFolders).DomainIds(domainIds).MailboxIds(mailboxIds).ProtectionJobIds(protectionJobIds).TenantId(tenantId).AllUnderHierarchy(allUnderHierarchy).Execute()

Returns the Outlook emails and folders containing emails.



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
    hasAttachments := true // bool | Specifies whether the emails have any attachments. (optional)
    senderAddress := "senderAddress_example" // string | Specifies the email address of the sender. (optional)
    recipientAddresses := []string{"Inner_example"} // []string | Specifies the email addresses of the recipients. (optional)
    ccRecipientAddresses := []string{"Inner_example"} // []string | Specifies the email addresses of the cc recipients. (optional)
    bccRecipientAddresses := []string{"Inner_example"} // []string | Specifies the email addresses of the bcc recipients. (optional)
    sentTimeSeconds := int64(789) // int64 | Specifies the unix time when the email was sent. (optional)
    receivedTimeSeconds := int64(789) // int64 | Specifies the unix time when the email was received. (optional)
    receivedStartTime := int64(789) // int64 | Specifies the unix start time for querying on email's received time. (optional)
    receivedEndTime := int64(789) // int64 | Specifies the unix end time for querying on email's received time. (optional)
    emailSubject := "emailSubject_example" // string | Specifies the subject of the email. (optional)
    folderName := "folderName_example" // string | Specifies the parent folder name of the email. (optional)
    showOnlyEmailFolders := true // bool | Specifies whether the query result should include only Email folders. (optional)
    domainIds := []int64{int64(123)} // []int64 | Specifies the domain Ids in which mailboxes are registered. (optional)
    mailboxIds := []int64{int64(123)} // []int64 | Specifies the mailbox Ids which contains the emails/folders. (optional)
    protectionJobIds := []int64{int64(123)} // []int64 | Specifies the protection job Ids which have backed up mailbox(es) continaing emails/folders. (optional)
    tenantId := "tenantId_example" // string | TenantId specifies the tenant whose action resulted in the audit log. (optional)
    allUnderHierarchy := true // bool | AllUnderHierarchy specifies if logs of all the tenants under the hierarchy of tenant with id TenantId should be returned. (optional)

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

    request := cohesitysdk.ApiGetOutlookEmailsRequest{
        hasAttachments: &HasAttachments
        senderAddress: &SenderAddress
        recipientAddresses: &RecipientAddresses
        ccRecipientAddresses: &CcRecipientAddresses
        bccRecipientAddresses: &BccRecipientAddresses
        sentTimeSeconds: &SentTimeSeconds
        receivedTimeSeconds: &ReceivedTimeSeconds
        receivedStartTime: &ReceivedStartTime
        receivedEndTime: &ReceivedEndTime
        emailSubject: &EmailSubject
        folderName: &FolderName
        showOnlyEmailFolders: &ShowOnlyEmailFolders
        domainIds: &DomainIds
        mailboxIds: &MailboxIds
        protectionJobIds: &ProtectionJobIds
        tenantId: &TenantId
        allUnderHierarchy: &AllUnderHierarchy
    }

    resp, r, err := api_client.RestoreTasksApi.GetOutlookEmails(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreTasksApi.GetOutlookEmails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOutlookEmails`: FileSearchResults
    fmt.Fprintf(os.Stdout, "Response from `RestoreTasksApi.GetOutlookEmails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOutlookEmailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hasAttachments** | **bool** | Specifies whether the emails have any attachments. | 
 **senderAddress** | **string** | Specifies the email address of the sender. | 
 **recipientAddresses** | **[]string** | Specifies the email addresses of the recipients. | 
 **ccRecipientAddresses** | **[]string** | Specifies the email addresses of the cc recipients. | 
 **bccRecipientAddresses** | **[]string** | Specifies the email addresses of the bcc recipients. | 
 **sentTimeSeconds** | **int64** | Specifies the unix time when the email was sent. | 
 **receivedTimeSeconds** | **int64** | Specifies the unix time when the email was received. | 
 **receivedStartTime** | **int64** | Specifies the unix start time for querying on email&#39;s received time. | 
 **receivedEndTime** | **int64** | Specifies the unix end time for querying on email&#39;s received time. | 
 **emailSubject** | **string** | Specifies the subject of the email. | 
 **folderName** | **string** | Specifies the parent folder name of the email. | 
 **showOnlyEmailFolders** | **bool** | Specifies whether the query result should include only Email folders. | 
 **domainIds** | **[]int64** | Specifies the domain Ids in which mailboxes are registered. | 
 **mailboxIds** | **[]int64** | Specifies the mailbox Ids which contains the emails/folders. | 
 **protectionJobIds** | **[]int64** | Specifies the protection job Ids which have backed up mailbox(es) continaing emails/folders. | 
 **tenantId** | **string** | TenantId specifies the tenant whose action resulted in the audit log. | 
 **allUnderHierarchy** | **bool** | AllUnderHierarchy specifies if logs of all the tenants under the hierarchy of tenant with id TenantId should be returned. | 

### Return type

[**FileSearchResults**](FileSearchResults.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRestorePointsForTimeRange

> RestorePointsForTimeRange GetRestorePointsForTimeRange(ctx).Body(body).Execute()

List Restore Points in a give time range.



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
    body := *openapiclient.NewRestorePointsForTimeRangeParam([]openapiclient.UniversalId{*openapiclient.NewUniversalId()}) // RestorePointsForTimeRangeParam | 

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

    request := cohesitysdk.ApiGetRestorePointsForTimeRangeRequest{
        body: &Body
    }

    resp, r, err := api_client.RestoreTasksApi.GetRestorePointsForTimeRange(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreTasksApi.GetRestorePointsForTimeRange``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRestorePointsForTimeRange`: RestorePointsForTimeRange
    fmt.Fprintf(os.Stdout, "Response from `RestoreTasksApi.GetRestorePointsForTimeRange`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRestorePointsForTimeRangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RestorePointsForTimeRangeParam**](RestorePointsForTimeRangeParam.md) |  | 

### Return type

[**RestorePointsForTimeRange**](RestorePointsForTimeRange.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRestoreTaskById

> RestoreTask GetRestoreTaskById(ctx, id).Execute()

List details about a single Restore Task.



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
    id := int64(789) // int64 | Specifies a unique id for the Restore Task.

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

    request := cohesitysdk.ApiGetRestoreTaskByIdRequest{
        id: &Id
    }

    resp, r, err := api_client.RestoreTasksApi.GetRestoreTaskById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreTasksApi.GetRestoreTaskById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRestoreTaskById`: RestoreTask
    fmt.Fprintf(os.Stdout, "Response from `RestoreTasksApi.GetRestoreTaskById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id for the Restore Task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRestoreTaskByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestoreTask**](RestoreTask.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRestoreTasks

> []RestoreTask GetRestoreTasks(ctx).TaskIds(taskIds).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).PageCount(pageCount).TaskTypes(taskTypes).Environment(environment).Execute()

List the Restore Tasks filtered by the specified parameters.



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
    taskIds := []int64{int64(123)} // []int64 | Filter by a list of Restore Task ids. (optional)
    startTimeUsecs := int64(789) // int64 | Filter by a start time specified as a Unix epoch Timestamp (in microseconds). All Restore Tasks (both completed and running) on the Cohesity Cluster that started after the specified start time but before the specified end time are returned. If not set, the start time is creation time of the Cohesity Cluster. (optional)
    endTimeUsecs := int64(789) // int64 | Filter by an end time specified as a Unix epoch Timestamp (in microseconds). All Restore Tasks (both completed and running) on the Cohesity Cluster that started after the specified start time but before the specified end time are returned. If not set, the end time is the current time. (optional)
    pageCount := int64(789) // int64 | Specifies the number of completed Restore Tasks to return in the response for pagination purposes. Running Restore Tasks are always returned. The newest completed Restore Tasks are returned. (optional)
    taskTypes := []string{"Inner_example"} // []string | Filter by the types of Restore Tasks such as 'kRecoverVMs', 'kCloneVMs', 'kCloneView' or 'kMountVolumes'. (optional)
    environment := "environment_example" // string | Specifies the environment like VMware, SQL, where the Protection Source exists. Supported environment types such as 'kView', 'kSQL', 'kVMware', etc. NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter. 'kVMware' indicates the VMware Protection Source environment. 'kHyperV' indicates the HyperV Protection Source environment. 'kSQL' indicates the SQL Protection Source environment. 'kView' indicates the View Protection Source environment. 'kPuppeteer' indicates the Cohesity's Remote Adapter. 'kPhysical' indicates the physical Protection Source environment. 'kPure' indicates the Pure Storage Protection Source environment. 'Nimble' indicates the Nimble Storage Protection Source environment. 'kAzure' indicates the Microsoft's Azure Protection Source environment. 'kNetapp' indicates the Netapp Protection Source environment. 'kAgent' indicates the Agent Protection Source environment. 'kGenericNas' indicates the Generic Network Attached Storage Protection Source environment. 'kAcropolis' indicates the Acropolis Protection Source environment. 'kPhsicalFiles' indicates the Physical Files Protection Source environment. 'kIsilon' indicates the Dell EMC's Isilon Protection Source environment. 'kGPFS' indicates IBM's GPFS Protection Source environment. 'kKVM' indicates the KVM Protection Source environment. 'kAWS' indicates the AWS Protection Source environment. 'kExchange' indicates the Exchange Protection Source environment. 'kHyperVVSS' indicates the HyperV VSS Protection Source environment. 'kOracle' indicates the Oracle Protection Source environment. 'kGCP' indicates the Google Cloud Platform Protection Source environment. 'kFlashBlade' indicates the Flash Blade Protection Source environment. 'kAWSNative' indicates the AWS Native Protection Source environment. 'kO365' indicates the Office 365 Protection Source environment. 'kO365Outlook' indicates Office 365 outlook Protection Source environment. 'kHyperFlex' indicates the Hyper Flex Protection Source environment. 'kGCPNative' indicates the GCP Native Protection Source environment. 'kAzureNative' indicates the Azure Native Protection Source environment. 'kKubernetes' indicates a Kubernetes Protection Source environment. 'kElastifile' indicates Elastifile Protection Source environment. 'kAD' indicates Active Directory Protection Source environment. 'kRDSSnapshotManager' indicates AWS RDS Protection Source environment. 'kCassandra' indicates Cassandra Protection Source environment. 'kMongoDB' indicates MongoDB Protection Source environment. 'kCouchbase' indicates Couchbase Protection Source environment. 'kHdfs' indicates Hdfs Protection Source environment. 'kHive' indicates Hive Protection Source environment. 'kHBase' indicates HBase Protection Source environment. (optional)

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

    request := cohesitysdk.ApiGetRestoreTasksRequest{
        taskIds: &TaskIds
        startTimeUsecs: &StartTimeUsecs
        endTimeUsecs: &EndTimeUsecs
        pageCount: &PageCount
        taskTypes: &TaskTypes
        environment: &Environment
    }

    resp, r, err := api_client.RestoreTasksApi.GetRestoreTasks(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreTasksApi.GetRestoreTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRestoreTasks`: []RestoreTask
    fmt.Fprintf(os.Stdout, "Response from `RestoreTasksApi.GetRestoreTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRestoreTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskIds** | **[]int64** | Filter by a list of Restore Task ids. | 
 **startTimeUsecs** | **int64** | Filter by a start time specified as a Unix epoch Timestamp (in microseconds). All Restore Tasks (both completed and running) on the Cohesity Cluster that started after the specified start time but before the specified end time are returned. If not set, the start time is creation time of the Cohesity Cluster. | 
 **endTimeUsecs** | **int64** | Filter by an end time specified as a Unix epoch Timestamp (in microseconds). All Restore Tasks (both completed and running) on the Cohesity Cluster that started after the specified start time but before the specified end time are returned. If not set, the end time is the current time. | 
 **pageCount** | **int64** | Specifies the number of completed Restore Tasks to return in the response for pagination purposes. Running Restore Tasks are always returned. The newest completed Restore Tasks are returned. | 
 **taskTypes** | **[]string** | Filter by the types of Restore Tasks such as &#39;kRecoverVMs&#39;, &#39;kCloneVMs&#39;, &#39;kCloneView&#39; or &#39;kMountVolumes&#39;. | 
 **environment** | **string** | Specifies the environment like VMware, SQL, where the Protection Source exists. Supported environment types such as &#39;kView&#39;, &#39;kSQL&#39;, &#39;kVMware&#39;, etc. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. &#39;kVMware&#39; indicates the VMware Protection Source environment. &#39;kHyperV&#39; indicates the HyperV Protection Source environment. &#39;kSQL&#39; indicates the SQL Protection Source environment. &#39;kView&#39; indicates the View Protection Source environment. &#39;kPuppeteer&#39; indicates the Cohesity&#39;s Remote Adapter. &#39;kPhysical&#39; indicates the physical Protection Source environment. &#39;kPure&#39; indicates the Pure Storage Protection Source environment. &#39;Nimble&#39; indicates the Nimble Storage Protection Source environment. &#39;kAzure&#39; indicates the Microsoft&#39;s Azure Protection Source environment. &#39;kNetapp&#39; indicates the Netapp Protection Source environment. &#39;kAgent&#39; indicates the Agent Protection Source environment. &#39;kGenericNas&#39; indicates the Generic Network Attached Storage Protection Source environment. &#39;kAcropolis&#39; indicates the Acropolis Protection Source environment. &#39;kPhsicalFiles&#39; indicates the Physical Files Protection Source environment. &#39;kIsilon&#39; indicates the Dell EMC&#39;s Isilon Protection Source environment. &#39;kGPFS&#39; indicates IBM&#39;s GPFS Protection Source environment. &#39;kKVM&#39; indicates the KVM Protection Source environment. &#39;kAWS&#39; indicates the AWS Protection Source environment. &#39;kExchange&#39; indicates the Exchange Protection Source environment. &#39;kHyperVVSS&#39; indicates the HyperV VSS Protection Source environment. &#39;kOracle&#39; indicates the Oracle Protection Source environment. &#39;kGCP&#39; indicates the Google Cloud Platform Protection Source environment. &#39;kFlashBlade&#39; indicates the Flash Blade Protection Source environment. &#39;kAWSNative&#39; indicates the AWS Native Protection Source environment. &#39;kO365&#39; indicates the Office 365 Protection Source environment. &#39;kO365Outlook&#39; indicates Office 365 outlook Protection Source environment. &#39;kHyperFlex&#39; indicates the Hyper Flex Protection Source environment. &#39;kGCPNative&#39; indicates the GCP Native Protection Source environment. &#39;kAzureNative&#39; indicates the Azure Native Protection Source environment. &#39;kKubernetes&#39; indicates a Kubernetes Protection Source environment. &#39;kElastifile&#39; indicates Elastifile Protection Source environment. &#39;kAD&#39; indicates Active Directory Protection Source environment. &#39;kRDSSnapshotManager&#39; indicates AWS RDS Protection Source environment. &#39;kCassandra&#39; indicates Cassandra Protection Source environment. &#39;kMongoDB&#39; indicates MongoDB Protection Source environment. &#39;kCouchbase&#39; indicates Couchbase Protection Source environment. &#39;kHdfs&#39; indicates Hdfs Protection Source environment. &#39;kHive&#39; indicates Hive Protection Source environment. &#39;kHBase&#39; indicates HBase Protection Source environment. | 

### Return type

[**[]RestoreTask**](RestoreTask.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSharepointDocuments

> FileSearchResults GetSharepointDocuments(ctx).TenantId(tenantId).AllUnderHierarchy(allUnderHierarchy).DocumentName(documentName).DomainIds(domainIds).SiteIds(siteIds).ProtectionJobIds(protectionJobIds).Execute()

Returns the Sharepoint Site's files and folders.



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
    tenantId := "tenantId_example" // string | TenantId specifies the tenant whose action resulted in the audit log. (optional)
    allUnderHierarchy := true // bool | AllUnderHierarchy specifies if logs of all the tenants under the hierarchy of tenant with id TenantId should be returned. (optional)
    documentName := "documentName_example" // string | Specifies the document(file/folder) name. (optional)
    domainIds := []int64{int64(123)} // []int64 | Specifies the domain Ids in which Sharepoint Site's are registered. (optional)
    siteIds := []int64{int64(123)} // []int64 | Specifies the Office365 Sharepoint Site Id. (optional)
    protectionJobIds := []int64{int64(123)} // []int64 | Specifies the protection job Ids which have backed up sites containing the documents. (optional)

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

    request := cohesitysdk.ApiGetSharepointDocumentsRequest{
        tenantId: &TenantId
        allUnderHierarchy: &AllUnderHierarchy
        documentName: &DocumentName
        domainIds: &DomainIds
        siteIds: &SiteIds
        protectionJobIds: &ProtectionJobIds
    }

    resp, r, err := api_client.RestoreTasksApi.GetSharepointDocuments(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreTasksApi.GetSharepointDocuments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSharepointDocuments`: FileSearchResults
    fmt.Fprintf(os.Stdout, "Response from `RestoreTasksApi.GetSharepointDocuments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSharepointDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | TenantId specifies the tenant whose action resulted in the audit log. | 
 **allUnderHierarchy** | **bool** | AllUnderHierarchy specifies if logs of all the tenants under the hierarchy of tenant with id TenantId should be returned. | 
 **documentName** | **string** | Specifies the document(file/folder) name. | 
 **domainIds** | **[]int64** | Specifies the domain Ids in which Sharepoint Site&#39;s are registered. | 
 **siteIds** | **[]int64** | Specifies the Office365 Sharepoint Site Id. | 
 **protectionJobIds** | **[]int64** | Specifies the protection job Ids which have backed up sites containing the documents. | 

### Return type

[**FileSearchResults**](FileSearchResults.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualDiskInformation

> []VirtualDiskInformation GetVirtualDiskInformation(ctx).ClusterId(clusterId).ClusterIncarnationId(clusterIncarnationId).JobId(jobId).JobRunId(jobRunId).StartTimeUsecs(startTimeUsecs).SourceId(sourceId).VaultId(vaultId).VaultName(vaultName).VaultType(vaultType).Execute()

Fetches information of virtual disk.

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
    clusterId := int64(789) // int64 | Specifies the Cohesity Cluster id where the Job was created.
    clusterIncarnationId := int64(789) // int64 | Specifies the incarnation id of the Cohesity Cluster where the Job was created.
    jobId := int64(789) // int64 | Specifies the id of the Job that captured the snapshot.
    jobRunId := int64(789) // int64 | Specifies the id of the Job Run that captured the snapshot.
    startTimeUsecs := int64(789) // int64 | Specifies the start time of the job run as a Unix epoch Timestamp in microseconds.
    sourceId := int64(789) // int64 | Specifies the Id of the Protection Source object.
    vaultId := int64(789) // int64 | Specifies the Id of the vault where snapshot was taken. (optional)
    vaultName := "vaultName_example" // string | Specifies the name of the vault where snapshot was taken. (optional)
    vaultType := "vaultType_example" // string | Specifes the type of the vault where snapshot was taken. (optional)

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

    request := cohesitysdk.ApiGetVirtualDiskInformationRequest{
        clusterId: &ClusterId
        clusterIncarnationId: &ClusterIncarnationId
        jobId: &JobId
        jobRunId: &JobRunId
        startTimeUsecs: &StartTimeUsecs
        sourceId: &SourceId
        vaultId: &VaultId
        vaultName: &VaultName
        vaultType: &VaultType
    }

    resp, r, err := api_client.RestoreTasksApi.GetVirtualDiskInformation(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreTasksApi.GetVirtualDiskInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualDiskInformation`: []VirtualDiskInformation
    fmt.Fprintf(os.Stdout, "Response from `RestoreTasksApi.GetVirtualDiskInformation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualDiskInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **int64** | Specifies the Cohesity Cluster id where the Job was created. | 
 **clusterIncarnationId** | **int64** | Specifies the incarnation id of the Cohesity Cluster where the Job was created. | 
 **jobId** | **int64** | Specifies the id of the Job that captured the snapshot. | 
 **jobRunId** | **int64** | Specifies the id of the Job Run that captured the snapshot. | 
 **startTimeUsecs** | **int64** | Specifies the start time of the job run as a Unix epoch Timestamp in microseconds. | 
 **sourceId** | **int64** | Specifies the Id of the Protection Source object. | 
 **vaultId** | **int64** | Specifies the Id of the vault where snapshot was taken. | 
 **vaultName** | **string** | Specifies the name of the vault where snapshot was taken. | 
 **vaultType** | **string** | Specifes the type of the vault where snapshot was taken. | 

### Return type

[**[]VirtualDiskInformation**](VirtualDiskInformation.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVmDirectoryList

> VmDirectoryListResult GetVmDirectoryList(ctx).JobId(jobId).JobUidObjectId(jobUidObjectId).EntityId(entityId).JobInstanceId(jobInstanceId).JobStartTimeUsecs(jobStartTimeUsecs).DirPath(dirPath).AttemptNum(attemptNum).VolumeName(volumeName).ViewBoxId(viewBoxId).ViewName(viewName).MaxEntries(maxEntries).VolumeInfoCookie(volumeInfoCookie).Cookie(cookie).StatFileEntries(statFileEntries).BrowseIndexedData(browseIndexedData).Execute()

Get the directory list based on the given directory name and other query parameters.

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
    jobId := int64(789) // int64 | JobId is the id of the local job that took the snapshot, which may or may not match the JobUidObjectId below depending on whether the object originally belonged to this local job or to a different remote job.
    jobUidObjectId := int64(789) // int64 | JobUidObjectId is the globally unique id of the job that the object originally belonged to. If this object originally belonged to a job from a remote cluster, this field will contain the JobId of the remote job, else it will contain the JobId of the local job.
    entityId := int64(789) // int64 | EntityId is the Id of the VM.
    jobInstanceId := int64(789) // int64 | JobInstanceId is the Id of the job run that backed up the entity.
    jobStartTimeUsecs := int64(789) // int64 | JobStartTimeUsecs is the start time in usecs of the job run that backed up the entity.
    dirPath := "dirPath_example" // string | DirPath is the full path of the directory whose contents need to be listed.
    attemptNum := int64(789) // int64 | AttemptNum is the attempt number of the run that successfully created the snapshot. (optional)
    volumeName := "volumeName_example" // string | VolumeName is the name of the volume that needs to be browsed. This should match the name returned in VolumeInfo. (optional)
    viewBoxId := int64(789) // int64 | Id of the View Box if a View is being browsed. (optional)
    viewName := "viewName_example" // string | Name of the View if a View is being browsed. (optional)
    maxEntries := int32(56) // int32 | MaxEntries is the maximum number of entries to return in this call. If there are more entries, server will return a cookie in the response that can be used to continue enumeration from the last call. (optional)
    volumeInfoCookie := int32(56) // int32 | VolumeInfoCookie is the cookie to be passed in calls to reading a VM dir for this volume. (optional)
    cookie := "cookie_example" // string | Cookie is used for paginating results. If ReadDirResult returned partial results, it will also return a cookie that can be used to resume the listing. The value returned in ReadDirResult should be passed in the next call. The first call should not have this value set. Note that this value is only a suggestion and server is free to do a short read (return fewer entries along with a cookie). (optional)
    statFileEntries := true // bool | StatFileEntries specifies whether file stat data is returned. (optional)
    browseIndexedData := true // bool | Specifies whether to use indexed data in Librarian for browse. (optional)

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

    request := cohesitysdk.ApiGetVmDirectoryListRequest{
        jobId: &JobId
        jobUidObjectId: &JobUidObjectId
        entityId: &EntityId
        jobInstanceId: &JobInstanceId
        jobStartTimeUsecs: &JobStartTimeUsecs
        dirPath: &DirPath
        attemptNum: &AttemptNum
        volumeName: &VolumeName
        viewBoxId: &ViewBoxId
        viewName: &ViewName
        maxEntries: &MaxEntries
        volumeInfoCookie: &VolumeInfoCookie
        cookie: &Cookie
        statFileEntries: &StatFileEntries
        browseIndexedData: &BrowseIndexedData
    }

    resp, r, err := api_client.RestoreTasksApi.GetVmDirectoryList(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreTasksApi.GetVmDirectoryList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVmDirectoryList`: VmDirectoryListResult
    fmt.Fprintf(os.Stdout, "Response from `RestoreTasksApi.GetVmDirectoryList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVmDirectoryListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobId** | **int64** | JobId is the id of the local job that took the snapshot, which may or may not match the JobUidObjectId below depending on whether the object originally belonged to this local job or to a different remote job. | 
 **jobUidObjectId** | **int64** | JobUidObjectId is the globally unique id of the job that the object originally belonged to. If this object originally belonged to a job from a remote cluster, this field will contain the JobId of the remote job, else it will contain the JobId of the local job. | 
 **entityId** | **int64** | EntityId is the Id of the VM. | 
 **jobInstanceId** | **int64** | JobInstanceId is the Id of the job run that backed up the entity. | 
 **jobStartTimeUsecs** | **int64** | JobStartTimeUsecs is the start time in usecs of the job run that backed up the entity. | 
 **dirPath** | **string** | DirPath is the full path of the directory whose contents need to be listed. | 
 **attemptNum** | **int64** | AttemptNum is the attempt number of the run that successfully created the snapshot. | 
 **volumeName** | **string** | VolumeName is the name of the volume that needs to be browsed. This should match the name returned in VolumeInfo. | 
 **viewBoxId** | **int64** | Id of the View Box if a View is being browsed. | 
 **viewName** | **string** | Name of the View if a View is being browsed. | 
 **maxEntries** | **int32** | MaxEntries is the maximum number of entries to return in this call. If there are more entries, server will return a cookie in the response that can be used to continue enumeration from the last call. | 
 **volumeInfoCookie** | **int32** | VolumeInfoCookie is the cookie to be passed in calls to reading a VM dir for this volume. | 
 **cookie** | **string** | Cookie is used for paginating results. If ReadDirResult returned partial results, it will also return a cookie that can be used to resume the listing. The value returned in ReadDirResult should be passed in the next call. The first call should not have this value set. Note that this value is only a suggestion and server is free to do a short read (return fewer entries along with a cookie). | 
 **statFileEntries** | **bool** | StatFileEntries specifies whether file stat data is returned. | 
 **browseIndexedData** | **bool** | Specifies whether to use indexed data in Librarian for browse. | 

### Return type

[**VmDirectoryListResult**](VmDirectoryListResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVmVolumesInformation

> VmVolumesInformation GetVmVolumesInformation(ctx).JobId(jobId).ClusterId(clusterId).ClusterIncarnationId(clusterIncarnationId).JobRunId(jobRunId).StartedTimeUsecs(startedTimeUsecs).SourceId(sourceId).OriginalJobId(originalJobId).AttemptNumber(attemptNumber).SupportedVolumesOnly(supportedVolumesOnly).ComputeVolumeInfo(computeVolumeInfo).Execute()

Get information about the logical volumes found in a VM.



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
    jobId := int64(789) // int64 | Specifies the Job id for the Protection Job that is currently associated with the object. If the object was backed up on current Cohesity Cluster, this field contains the id for the Job that captured this backup object. If the object was backed up on a Primary Cluster and replicated to this Cohesity Cluster, a new Inactive Job is created, the object is now associated with new Inactive Job, and this field contains the id of the new Inactive Job.
    clusterId := int64(789) // int64 | Specifies the Cohesity Cluster id where the Job was created.
    clusterIncarnationId := int64(789) // int64 | Specifies the incarnation id of the Cohesity Cluster where the Job was created. An incarnation id is generated when a Cohesity Cluster is initially created.
    jobRunId := int64(789) // int64 | Specifies the id of the Job Run that captured the snapshot.
    startedTimeUsecs := int64(789) // int64 | Specifies the time when the Job Run starts capturing a snapshot. Specified as a Unix epoch Timestamp (in microseconds).
    sourceId := int64(789) // int64 | Specifies the id of the VM object to search for volumes.
    originalJobId := int64(789) // int64 | Specifies the id for the Protection Job that originally captured the snapshots of the original object. If the object was backed up on a Primary Cluster replicated to this Cohesity Cluster, and a new Inactive Job is created, this field still contains the id of the original Job and NOT the id of the new Inactive Job. This field is used in combination with the clusterId and clusterIncarnationId to uniquely identify a Job.
    attemptNumber := int64(789) // int64 | Specifies the number of the attempts made by the Job Run to capture a snapshot of the object. For example, if an snapshot is successfully captured after three attempts, this field equals 3. (optional)
    supportedVolumesOnly := true // bool | Specifies to return only supported volumes information. Unsupported volumes are not returned if this flag is set to true. Default is false. (optional)
    computeVolumeInfo := true // bool | Specifies whether to compute volume information if it is not found. If `ComputeVolumeInfo` is false and volume information is not found it skips computation of volume information and returns KNotFound. (optional)

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

    request := cohesitysdk.ApiGetVmVolumesInformationRequest{
        jobId: &JobId
        clusterId: &ClusterId
        clusterIncarnationId: &ClusterIncarnationId
        jobRunId: &JobRunId
        startedTimeUsecs: &StartedTimeUsecs
        sourceId: &SourceId
        originalJobId: &OriginalJobId
        attemptNumber: &AttemptNumber
        supportedVolumesOnly: &SupportedVolumesOnly
        computeVolumeInfo: &ComputeVolumeInfo
    }

    resp, r, err := api_client.RestoreTasksApi.GetVmVolumesInformation(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreTasksApi.GetVmVolumesInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVmVolumesInformation`: VmVolumesInformation
    fmt.Fprintf(os.Stdout, "Response from `RestoreTasksApi.GetVmVolumesInformation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVmVolumesInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobId** | **int64** | Specifies the Job id for the Protection Job that is currently associated with the object. If the object was backed up on current Cohesity Cluster, this field contains the id for the Job that captured this backup object. If the object was backed up on a Primary Cluster and replicated to this Cohesity Cluster, a new Inactive Job is created, the object is now associated with new Inactive Job, and this field contains the id of the new Inactive Job. | 
 **clusterId** | **int64** | Specifies the Cohesity Cluster id where the Job was created. | 
 **clusterIncarnationId** | **int64** | Specifies the incarnation id of the Cohesity Cluster where the Job was created. An incarnation id is generated when a Cohesity Cluster is initially created. | 
 **jobRunId** | **int64** | Specifies the id of the Job Run that captured the snapshot. | 
 **startedTimeUsecs** | **int64** | Specifies the time when the Job Run starts capturing a snapshot. Specified as a Unix epoch Timestamp (in microseconds). | 
 **sourceId** | **int64** | Specifies the id of the VM object to search for volumes. | 
 **originalJobId** | **int64** | Specifies the id for the Protection Job that originally captured the snapshots of the original object. If the object was backed up on a Primary Cluster replicated to this Cohesity Cluster, and a new Inactive Job is created, this field still contains the id of the original Job and NOT the id of the new Inactive Job. This field is used in combination with the clusterId and clusterIncarnationId to uniquely identify a Job. | 
 **attemptNumber** | **int64** | Specifies the number of the attempts made by the Job Run to capture a snapshot of the object. For example, if an snapshot is successfully captured after three attempts, this field equals 3. | 
 **supportedVolumesOnly** | **bool** | Specifies to return only supported volumes information. Unsupported volumes are not returned if this flag is set to true. Default is false. | 
 **computeVolumeInfo** | **bool** | Specifies whether to compute volume information if it is not found. If &#x60;ComputeVolumeInfo&#x60; is false and volume information is not found it skips computation of volume information and returns KNotFound. | 

### Return type

[**VmVolumesInformation**](VmVolumesInformation.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStorageProfiles

> ListStorageProfilesResponseBody ListStorageProfiles(ctx, id).Execute()

Fetches information of virtual disk.

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
    id := int64(789) // int64 | Specifies a unique id for the VDC.

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

    request := cohesitysdk.ApiListStorageProfilesRequest{
        id: &Id
    }

    resp, r, err := api_client.RestoreTasksApi.ListStorageProfiles(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreTasksApi.ListStorageProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStorageProfiles`: ListStorageProfilesResponseBody
    fmt.Fprintf(os.Stdout, "Response from `RestoreTasksApi.ListStorageProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id for the VDC. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStorageProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListStorageProfilesResponseBody**](ListStorageProfilesResponseBody.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublicDestroyCloneTask

> PublicDestroyCloneTask(ctx, id).Execute()

Destroy a clone task with specified id.

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
    id := int64(789) // int64 | Specifies a unique id of the Clone Task to destroy.

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

    request := cohesitysdk.ApiPublicDestroyCloneTaskRequest{
        id: &Id
    }

    resp, r, err := api_client.RestoreTasksApi.PublicDestroyCloneTask(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreTasksApi.PublicDestroyCloneTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the Clone Task to destroy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicDestroyCloneTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAdObjects

> []ADObject SearchAdObjects(ctx).RestoreTaskId(restoreTaskId).NumObjects(numObjects).SearchBaseDn(searchBaseDn).SubtreeSearchScope(subtreeSearchScope).CompareObjects(compareObjects).ExcludeSystemProperties(excludeSystemProperties).Filter(filter).RecordOffset(recordOffset).Execute()

Searches for AD Objects in both Production and Snapshot AD from given search base dn and offset.



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
    restoreTaskId := int64(789) // int64 | Specifies the restoreTaskId corresponding to which we need to search AD objects.
    numObjects := int32(56) // int32 | Specifies the number of AD Objects to be fetched. (optional)
    searchBaseDn := "searchBaseDn_example" // string | Specifies the search base distinguished name from where the search should begin in the hierarchy of the AD in both Production and Snapshot AD. (optional)
    subtreeSearchScope := true // bool | Specifies the search scope for the request. If subtree search scope is true all the children of Search Base DN are returned from given offset. If subtree search scope is false only all objects which are one level from the Search Base DN are returned. (optional)
    compareObjects := true // bool | Specifies the option to compare the properties from Snapshot AD and Production AD if specifed and sets kNotEqual flag in the result when there is mismatch. (optional)
    excludeSystemProperties := true // bool | Specifies the option to exclude the system attributes while comparing the the objects from the Production AD and Snapshot AD. (optional)
    filter := "filter_example" // string | Specifies the filter which can be used for searching the AD Objects from given Search Base DN. There are two types of filters supported. They are: 1) If the string does not contain LDAP delimiters '(' and ')', then it is assumed to be ANR search \"(anr=<ldap_filter>)\" Eg: \"a\" will result in query to return all ANR fields with \"a\" characters (case insensitive) in them 2) Search with OR and AND combination: \"(|(&(objectClass=user)(distinguishedName=CN=Jone Doe,OU=Users,DC=corp,DC=cohesity,DC=com))(&(objectClass=user) (sAMAccountName=jdoe)))\" (optional)
    recordOffset := int32(56) // int32 | Specifies the offset from which AD objects should be searched in both the Snapshot and Production AD. (optional)

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

    request := cohesitysdk.ApiSearchAdObjectsRequest{
        restoreTaskId: &RestoreTaskId
        numObjects: &NumObjects
        searchBaseDn: &SearchBaseDn
        subtreeSearchScope: &SubtreeSearchScope
        compareObjects: &CompareObjects
        excludeSystemProperties: &ExcludeSystemProperties
        filter: &Filter
        recordOffset: &RecordOffset
    }

    resp, r, err := api_client.RestoreTasksApi.SearchAdObjects(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreTasksApi.SearchAdObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchAdObjects`: []ADObject
    fmt.Fprintf(os.Stdout, "Response from `RestoreTasksApi.SearchAdObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAdObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restoreTaskId** | **int64** | Specifies the restoreTaskId corresponding to which we need to search AD objects. | 
 **numObjects** | **int32** | Specifies the number of AD Objects to be fetched. | 
 **searchBaseDn** | **string** | Specifies the search base distinguished name from where the search should begin in the hierarchy of the AD in both Production and Snapshot AD. | 
 **subtreeSearchScope** | **bool** | Specifies the search scope for the request. If subtree search scope is true all the children of Search Base DN are returned from given offset. If subtree search scope is false only all objects which are one level from the Search Base DN are returned. | 
 **compareObjects** | **bool** | Specifies the option to compare the properties from Snapshot AD and Production AD if specifed and sets kNotEqual flag in the result when there is mismatch. | 
 **excludeSystemProperties** | **bool** | Specifies the option to exclude the system attributes while comparing the the objects from the Production AD and Snapshot AD. | 
 **filter** | **string** | Specifies the filter which can be used for searching the AD Objects from given Search Base DN. There are two types of filters supported. They are: 1) If the string does not contain LDAP delimiters &#39;(&#39; and &#39;)&#39;, then it is assumed to be ANR search \&quot;(anr&#x3D;&lt;ldap_filter&gt;)\&quot; Eg: \&quot;a\&quot; will result in query to return all ANR fields with \&quot;a\&quot; characters (case insensitive) in them 2) Search with OR and AND combination: \&quot;(|(&amp;(objectClass&#x3D;user)(distinguishedName&#x3D;CN&#x3D;Jone Doe,OU&#x3D;Users,DC&#x3D;corp,DC&#x3D;cohesity,DC&#x3D;com))(&amp;(objectClass&#x3D;user) (sAMAccountName&#x3D;jdoe)))\&quot; | 
 **recordOffset** | **int32** | Specifies the offset from which AD objects should be searched in both the Snapshot and Production AD. | 

### Return type

[**[]ADObject**](ADObject.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchObjects

> ObjectSearchResults SearchObjects(ctx).Search(search).JobIds(jobIds).RegisteredSourceIds(registeredSourceIds).ViewBoxIds(viewBoxIds).Environments(environments).Office365SourceTypes(office365SourceTypes).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).StartIndex(startIndex).PageCount(pageCount).OperatingSystems(operatingSystems).Application(application).OwnerEntityId(ownerEntityId).TenantId(tenantId).AllUnderHierarchy(allUnderHierarchy).Execute()

Find backup objects that match the specified search and filter criteria on the Cohesity Cluster.



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
    search := "search_example" // string | Filter by searching for sub-strings in the item name. The specified string can match any part of the item name. For example: \"vm\" or \"123\" both match the item name of \"vm-123\". (optional)
    jobIds := []int64{int64(123)} // []int64 | Filter by a list of Protection Job ids. Only items backed up by the specified Jobs are listed. (optional)
    registeredSourceIds := []int64{int64(123)} // []int64 | Filter by a list of Registered Sources ids. Only items from the listed Registered Sources are returned. (optional)
    viewBoxIds := []int64{int64(123)} // []int64 | Filter by a list of Domains (View Boxes) ids. Only items stored in the listed Domains (View Boxes) are returned. (optional)
    environments := []string{"Environments_example"} // []string | Filter by environment types such as 'kVMware', 'kView', etc. Only items from the specified environment types are returned. NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter. (optional)
    office365SourceTypes := []string{"Office365SourceTypes_example"} // []string | Filter by Office365 types such as 'kUser', 'kSite', etc. Only items from the specified source types are returned. (optional)
    startTimeUsecs := int64(789) // int64 | Filter by backup completion time by specifying a backup completion start and end times. Specified as a Unix epoch Timestamp (in microseconds). Only items created by backups that completed between the specified start and end times are returned. (optional)
    endTimeUsecs := int64(789) // int64 | Filter by backup completion time by specify a backup completion start and end times. Specified as a Unix epoch Timestamp (in microseconds). Only items created by backups that completed between the specified start and end times are returned. (optional)
    startIndex := int64(789) // int64 | Specifies an index number that can be used to return subsets of items in multiple requests. Break up the items to return into multiple requests by setting pageCount and using startIndex to return a subsets of items. For example, set startIndex to 0 to get the first set of items for the first request. Increment startIndex by pageCount to get the next set of items for a next request. Continue until all items are returned and therefore the total number of returned items is equal to totalCount. (optional)
    pageCount := int64(789) // int64 | Limit the number of items to return in the response for pagination purposes. (optional)
    operatingSystems := []string{"Inner_example"} // []string | Filter by the Operating Systems running on VMs and Physical Servers. This filter is applicable only to VMs and physical servers. (optional)
    application := "application_example" // string | Filter by application when the environment type is kSQL. For example, if SQL is specified the SQL databases are returned. (optional)
    ownerEntityId := int64(789) // int64 | Filter objects by the Entity id of the owner VM. For example, if a ownerEntityId is provided while searching for SQL databases, only SQL databases belonging to the VM with the specified id are returned. ownerEntityId is only significant if application is set to SQL. (optional)
    tenantId := "tenantId_example" // string | TenantId specifies the tenant whose action resulted in the audit log. (optional)
    allUnderHierarchy := true // bool | AllUnderHierarchy specifies if logs of all the tenants under the hierarchy of tenant with id TenantId should be returned. (optional)

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

    request := cohesitysdk.ApiSearchObjectsRequest{
        search: &Search
        jobIds: &JobIds
        registeredSourceIds: &RegisteredSourceIds
        viewBoxIds: &ViewBoxIds
        environments: &Environments
        office365SourceTypes: &Office365SourceTypes
        startTimeUsecs: &StartTimeUsecs
        endTimeUsecs: &EndTimeUsecs
        startIndex: &StartIndex
        pageCount: &PageCount
        operatingSystems: &OperatingSystems
        application: &Application
        ownerEntityId: &OwnerEntityId
        tenantId: &TenantId
        allUnderHierarchy: &AllUnderHierarchy
    }

    resp, r, err := api_client.RestoreTasksApi.SearchObjects(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreTasksApi.SearchObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchObjects`: ObjectSearchResults
    fmt.Fprintf(os.Stdout, "Response from `RestoreTasksApi.SearchObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Filter by searching for sub-strings in the item name. The specified string can match any part of the item name. For example: \&quot;vm\&quot; or \&quot;123\&quot; both match the item name of \&quot;vm-123\&quot;. | 
 **jobIds** | **[]int64** | Filter by a list of Protection Job ids. Only items backed up by the specified Jobs are listed. | 
 **registeredSourceIds** | **[]int64** | Filter by a list of Registered Sources ids. Only items from the listed Registered Sources are returned. | 
 **viewBoxIds** | **[]int64** | Filter by a list of Domains (View Boxes) ids. Only items stored in the listed Domains (View Boxes) are returned. | 
 **environments** | **[]string** | Filter by environment types such as &#39;kVMware&#39;, &#39;kView&#39;, etc. Only items from the specified environment types are returned. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. | 
 **office365SourceTypes** | **[]string** | Filter by Office365 types such as &#39;kUser&#39;, &#39;kSite&#39;, etc. Only items from the specified source types are returned. | 
 **startTimeUsecs** | **int64** | Filter by backup completion time by specifying a backup completion start and end times. Specified as a Unix epoch Timestamp (in microseconds). Only items created by backups that completed between the specified start and end times are returned. | 
 **endTimeUsecs** | **int64** | Filter by backup completion time by specify a backup completion start and end times. Specified as a Unix epoch Timestamp (in microseconds). Only items created by backups that completed between the specified start and end times are returned. | 
 **startIndex** | **int64** | Specifies an index number that can be used to return subsets of items in multiple requests. Break up the items to return into multiple requests by setting pageCount and using startIndex to return a subsets of items. For example, set startIndex to 0 to get the first set of items for the first request. Increment startIndex by pageCount to get the next set of items for a next request. Continue until all items are returned and therefore the total number of returned items is equal to totalCount. | 
 **pageCount** | **int64** | Limit the number of items to return in the response for pagination purposes. | 
 **operatingSystems** | **[]string** | Filter by the Operating Systems running on VMs and Physical Servers. This filter is applicable only to VMs and physical servers. | 
 **application** | **string** | Filter by application when the environment type is kSQL. For example, if SQL is specified the SQL databases are returned. | 
 **ownerEntityId** | **int64** | Filter objects by the Entity id of the owner VM. For example, if a ownerEntityId is provided while searching for SQL databases, only SQL databases belonging to the VM with the specified id are returned. ownerEntityId is only significant if application is set to SQL. | 
 **tenantId** | **string** | TenantId specifies the tenant whose action resulted in the audit log. | 
 **allUnderHierarchy** | **bool** | AllUnderHierarchy specifies if logs of all the tenants under the hierarchy of tenant with id TenantId should be returned. | 

### Return type

[**ObjectSearchResults**](ObjectSearchResults.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchProductionAdObjects

> []ADObject SearchProductionAdObjects(ctx).Body(body).Execute()

Searches for AD Objects that match the list of object guids, sam account names and distinguished names provided in the request.



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
    body := *openapiclient.NewSearchProductionAdObjectsRequest() // SearchProductionAdObjectsRequest | Specifies the Request to search the AD Objects from Production AD.

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

    request := cohesitysdk.ApiSearchProductionAdObjectsRequest{
        body: &Body
    }

    resp, r, err := api_client.RestoreTasksApi.SearchProductionAdObjects(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreTasksApi.SearchProductionAdObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchProductionAdObjects`: []ADObject
    fmt.Fprintf(os.Stdout, "Response from `RestoreTasksApi.SearchProductionAdObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchProductionAdObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SearchProductionAdObjectsRequest**](SearchProductionAdObjectsRequest.md) | Specifies the Request to search the AD Objects from Production AD. | 

### Return type

[**[]ADObject**](ADObject.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchRestoredFiles

> FileSearchResults SearchRestoredFiles(ctx).MustHaveTags(mustHaveTags).MightHaveTags(mightHaveTags).MustHaveSnapshotTags(mustHaveSnapshotTags).MightHaveSnapshotTags(mightHaveSnapshotTags).Paginate(paginate).PageSize(pageSize).PaginationCookie(paginationCookie).Search(search).JobIds(jobIds).RegisteredSourceIds(registeredSourceIds).ViewBoxIds(viewBoxIds).Environments(environments).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).StartIndex(startIndex).PageCount(pageCount).SourceIds(sourceIds).FolderOnly(folderOnly).TenantId(tenantId).AllUnderHierarchy(allUnderHierarchy).Execute()

Search for files and folders to recover that match the specified search and filter criteria on the Cohesity Cluster.



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
    mustHaveTags := []string{"Inner_example"} // []string | Specifies tags which must be all present in the document. (optional)
    mightHaveTags := []string{"Inner_example"} // []string | Specifies list of tags, one of which might be present in the document. These are OR'ed together and the resulting criteria AND'ed with the rest of the query. (optional)
    mustHaveSnapshotTags := []string{"Inner_example"} // []string | Specifies snapshot tags which must be all present in the document. (optional)
    mightHaveSnapshotTags := []string{"Inner_example"} // []string | Specifies list of snapshot tags, one of which might be present in the document. These are OR'ed together and the resulting criteria AND'ed with the rest of the query. (optional)
    paginate := true // bool | Specifies bool to control pagination of search results. Only valid for librarian queries. If this is set to true and a pagination cookie is provided, search will be resumed. (optional)
    pageSize := int32(56) // int32 | Specifies pagesize for pagination. Only valid for librarian queries. Effective only when Paginate is set to true. (optional)
    paginationCookie := "paginationCookie_example" // string | Specifies cookie for resuming search if pagination is being used. Only valid for librarian queries. Effective only when Paginate is set to true. (optional)
    search := "search_example" // string | Filter by searching for sub-strings in the item name. The specified string can match any part of the item name. For example: \"vm\" or \"123\" both match the item name of \"vm-123\". (optional)
    jobIds := []int64{int64(123)} // []int64 | Filter by a list of Protection Job ids. Only items backed up by the specified Jobs are listed. (optional)
    registeredSourceIds := []int64{int64(123)} // []int64 | Filter by a list of Registered Sources ids. Only items from the listed Registered Sources are returned. (optional)
    viewBoxIds := []int64{int64(123)} // []int64 | Filter by a list of Domains (View Boxes) ids. Only items stored in the listed Domains (View Boxes) are returned. (optional)
    environments := []string{"Environments_example"} // []string | Filter by environment types such as 'kVMware', 'kView', etc. Only items from the specified environment types are returned. NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter. (optional)
    startTimeUsecs := int64(789) // int64 | Filter by backup completion time by specifying a backup completion start and end times. Specified as a Unix epoch Timestamp (in microseconds). Only items created by backups that completed between the specified start and end times are returned. (optional)
    endTimeUsecs := int64(789) // int64 | Filter by backup completion time by specify a backup completion start and end times. Specified as a Unix epoch Timestamp (in microseconds). Only items created by backups that completed between the specified start and end times are returned. (optional)
    startIndex := int64(789) // int64 | Specifies an index number that can be used to return subsets of items in multiple requests. Break up the items to return into multiple requests by setting pageCount and using startIndex to return a subsets of items. For example, set startIndex to 0 to get the first set of items for the first request. Increment startIndex by pageCount to get the next set of items for a next request. Continue until all items are returned and therefore the total number of returned items is equal to totalCount. (optional)
    pageCount := int64(789) // int64 | Limit the number of items to return in the response for pagination purposes. (optional)
    sourceIds := []int64{int64(123)} // []int64 | Filter by source ids. Only files and folders found in the listed sources (such as VMs) are returned. (optional)
    folderOnly := true // bool | Filter by folders or files. If true, only folders are returned. If false, only files are returned. If not specified, both files and folders are returned. (optional)
    tenantId := "tenantId_example" // string | TenantId specifies the tenant whose action resulted in the audit log. (optional)
    allUnderHierarchy := true // bool | AllUnderHierarchy specifies if logs of all the tenants under the hierarchy of tenant with id TenantId should be returned. (optional)

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

    request := cohesitysdk.ApiSearchRestoredFilesRequest{
        mustHaveTags: &MustHaveTags
        mightHaveTags: &MightHaveTags
        mustHaveSnapshotTags: &MustHaveSnapshotTags
        mightHaveSnapshotTags: &MightHaveSnapshotTags
        paginate: &Paginate
        pageSize: &PageSize
        paginationCookie: &PaginationCookie
        search: &Search
        jobIds: &JobIds
        registeredSourceIds: &RegisteredSourceIds
        viewBoxIds: &ViewBoxIds
        environments: &Environments
        startTimeUsecs: &StartTimeUsecs
        endTimeUsecs: &EndTimeUsecs
        startIndex: &StartIndex
        pageCount: &PageCount
        sourceIds: &SourceIds
        folderOnly: &FolderOnly
        tenantId: &TenantId
        allUnderHierarchy: &AllUnderHierarchy
    }

    resp, r, err := api_client.RestoreTasksApi.SearchRestoredFiles(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreTasksApi.SearchRestoredFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchRestoredFiles`: FileSearchResults
    fmt.Fprintf(os.Stdout, "Response from `RestoreTasksApi.SearchRestoredFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRestoredFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mustHaveTags** | **[]string** | Specifies tags which must be all present in the document. | 
 **mightHaveTags** | **[]string** | Specifies list of tags, one of which might be present in the document. These are OR&#39;ed together and the resulting criteria AND&#39;ed with the rest of the query. | 
 **mustHaveSnapshotTags** | **[]string** | Specifies snapshot tags which must be all present in the document. | 
 **mightHaveSnapshotTags** | **[]string** | Specifies list of snapshot tags, one of which might be present in the document. These are OR&#39;ed together and the resulting criteria AND&#39;ed with the rest of the query. | 
 **paginate** | **bool** | Specifies bool to control pagination of search results. Only valid for librarian queries. If this is set to true and a pagination cookie is provided, search will be resumed. | 
 **pageSize** | **int32** | Specifies pagesize for pagination. Only valid for librarian queries. Effective only when Paginate is set to true. | 
 **paginationCookie** | **string** | Specifies cookie for resuming search if pagination is being used. Only valid for librarian queries. Effective only when Paginate is set to true. | 
 **search** | **string** | Filter by searching for sub-strings in the item name. The specified string can match any part of the item name. For example: \&quot;vm\&quot; or \&quot;123\&quot; both match the item name of \&quot;vm-123\&quot;. | 
 **jobIds** | **[]int64** | Filter by a list of Protection Job ids. Only items backed up by the specified Jobs are listed. | 
 **registeredSourceIds** | **[]int64** | Filter by a list of Registered Sources ids. Only items from the listed Registered Sources are returned. | 
 **viewBoxIds** | **[]int64** | Filter by a list of Domains (View Boxes) ids. Only items stored in the listed Domains (View Boxes) are returned. | 
 **environments** | **[]string** | Filter by environment types such as &#39;kVMware&#39;, &#39;kView&#39;, etc. Only items from the specified environment types are returned. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. | 
 **startTimeUsecs** | **int64** | Filter by backup completion time by specifying a backup completion start and end times. Specified as a Unix epoch Timestamp (in microseconds). Only items created by backups that completed between the specified start and end times are returned. | 
 **endTimeUsecs** | **int64** | Filter by backup completion time by specify a backup completion start and end times. Specified as a Unix epoch Timestamp (in microseconds). Only items created by backups that completed between the specified start and end times are returned. | 
 **startIndex** | **int64** | Specifies an index number that can be used to return subsets of items in multiple requests. Break up the items to return into multiple requests by setting pageCount and using startIndex to return a subsets of items. For example, set startIndex to 0 to get the first set of items for the first request. Increment startIndex by pageCount to get the next set of items for a next request. Continue until all items are returned and therefore the total number of returned items is equal to totalCount. | 
 **pageCount** | **int64** | Limit the number of items to return in the response for pagination purposes. | 
 **sourceIds** | **[]int64** | Filter by source ids. Only files and folders found in the listed sources (such as VMs) are returned. | 
 **folderOnly** | **bool** | Filter by folders or files. If true, only folders are returned. If false, only files are returned. If not specified, both files and folders are returned. | 
 **tenantId** | **string** | TenantId specifies the tenant whose action resulted in the audit log. | 
 **allUnderHierarchy** | **bool** | AllUnderHierarchy specifies if logs of all the tenants under the hierarchy of tenant with id TenantId should be returned. | 

### Return type

[**FileSearchResults**](FileSearchResults.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRestoreTask

> RestoreTask UpdateRestoreTask(ctx).Body(body).Execute()





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
    body := *openapiclient.NewUpdateRestoreTaskParams() // UpdateRestoreTaskParams | 

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

    request := cohesitysdk.ApiUpdateRestoreTaskRequest{
        body: &Body
    }

    resp, r, err := api_client.RestoreTasksApi.UpdateRestoreTask(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreTasksApi.UpdateRestoreTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRestoreTask`: RestoreTask
    fmt.Fprintf(os.Stdout, "Response from `RestoreTasksApi.UpdateRestoreTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRestoreTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateRestoreTaskParams**](UpdateRestoreTaskParams.md) |  | 

### Return type

[**RestoreTask**](RestoreTask.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

