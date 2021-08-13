# \Object

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BrowseObjectContents**](Object.md#BrowseObjectContents) | **Post** /data-protect/objects/{id}/browse | Fetch the contents (files &amp; folders) for the specified object.
[**CancelObjectRuns**](Object.md#CancelObjectRuns) | **Post** /data-protect/objects/runs/cancel | Cancel object runs.
[**ConstructMetaInfo**](Object.md#ConstructMetaInfo) | **Post** /data-protect/snapshots/{snapshotId}/meta-info | Construct meta info for any workflow from object snapshot and some other information.
[**FilterObjects**](Object.md#FilterObjects) | **Post** /data-protect/filter/objects | List all the filtered objects.
[**GetAllIndexedObjectSnapshots**](Object.md#GetAllIndexedObjectSnapshots) | **Get** /data-protect/objects/{objectId}/indexed-objects/snapshots | Get snapshots of indexed object.
[**GetIndexedObjectSnapshots**](Object.md#GetIndexedObjectSnapshots) | **Get** /data-protect/objects/{objectId}/protection-groups/{protectionGroupId}/indexed-objects/snapshots | Get snapshots of indexed object.
[**GetMcmObjectSnapshots**](Object.md#GetMcmObjectSnapshots) | **Get** /mcm/data-protect/objects/snapshots | List the snapshots for a given object.
[**GetMcmObjectStats**](Object.md#GetMcmObjectStats) | **Get** /mcm/data-protect/objects/stats | Get stats for a given object.
[**GetMcmObjectsActivity**](Object.md#GetMcmObjectsActivity) | **Post** /mcm/data-protect/objects/activity | Get Object activity on Helios.
[**GetObjectIdentifiers**](Object.md#GetObjectIdentifiers) | **Get** /data-protect/objects/object-identifiers | Get Object Identifiers
[**GetObjectRunByRunId**](Object.md#GetObjectRunByRunId) | **Get** /data-protect/objects/{id}/runs/{runId} | Get a run for an object.
[**GetObjectRuns**](Object.md#GetObjectRuns) | **Get** /data-protect/objects/{id}/runs | Get the list of runs for an object.
[**GetObjectSnapshotInfo**](Object.md#GetObjectSnapshotInfo) | **Get** /data-protect/snapshots/{snapshotId} | Get details of object snapshot.
[**GetObjectSnapshotVolumeInfo**](Object.md#GetObjectSnapshotVolumeInfo) | **Get** /data-protect/snapshots/{snapshotId}/volume | Get volume info of object snapshot.
[**GetObjectSnapshots**](Object.md#GetObjectSnapshots) | **Get** /data-protect/objects/{id}/snapshots | List the snapshots for a given object.
[**GetObjectStats**](Object.md#GetObjectStats) | **Get** /data-protect/objects/{id}/stats | Get stats for a given object.
[**GetObjectTree**](Object.md#GetObjectTree) | **Get** /data-protect/objects/{id}/tree | Get the objects tree hierarchy for for an Object.
[**GetObjectsLastRun**](Object.md#GetObjectsLastRun) | **Get** /data-protect/objects/last-run | Get last protection run of objects.
[**GetPITRangesForProtectedObject**](Object.md#GetPITRangesForProtectedObject) | **Get** /data-protect/objects/{id}/pit-ranges | Get PIT ranges for an object
[**GetProtectedObjectOfAnyTypeById**](Object.md#GetProtectedObjectOfAnyTypeById) | **Get** /data-protect/objects/{id} | Get an Object.
[**GetProtectedObjectsOfAnyType**](Object.md#GetProtectedObjectsOfAnyType) | **Get** /data-protect/objects | Get Objects.
[**GetSnapshotDiff**](Object.md#GetSnapshotDiff) | **Post** /data-protect/objects/{id}/snapshot-diff | Get diff between two snapshots of a given object.
[**GetSourceHierarchyObjects**](Object.md#GetSourceHierarchyObjects) | **Get** /data-protect/sources/{sourceId}/objects | List objects on a source which can be used for data protection.
[**McmGetTenantObjectIds**](Object.md#McmGetTenantObjectIds) | **Post** /mcm/tenants/object-ids | GetTenantObjectIds
[**ObjectsActions**](Object.md#ObjectsActions) | **Post** /data-protect/objects/actions | Actions on Objects
[**PerformActionOnObject**](Object.md#PerformActionOnObject) | **Post** /data-protect/objects/{id}/actions | Perform an action on an object.
[**UpdateObjectSnapshot**](Object.md#UpdateObjectSnapshot) | **Put** /data-protect/objects/{id}/snapshots/{snapshotId} | Update an object snapshot.



## BrowseObjectContents

> FileFolderInfo BrowseObjectContents(ctx, id).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Fetch the contents (files & folders) for the specified object.



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
    id := int64(789) // int64 | Specifies the id of the Object.
    body := *openapiclient.NewObjectBrowseRequest("Environment_example") // ObjectBrowseRequest | Specifies the parameters to fetch contents of an object.
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

    request := helios.ApiBrowseObjectContentsRequest{
        Id: &id
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Object.BrowseObjectContents(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Object.BrowseObjectContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BrowseObjectContents`: FileFolderInfo
    fmt.Fprintf(os.Stdout, "Response from `Object.BrowseObjectContents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the id of the Object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrowseObjectContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ObjectBrowseRequest**](ObjectBrowseRequest.md) | Specifies the parameters to fetch contents of an object. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**FileFolderInfo**](FileFolderInfo.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelObjectRuns

> CancelObjectRunsResults CancelObjectRuns(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Cancel object runs.



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
    body := *openapiclient.NewCancelObjectRunsRequest() // CancelObjectRunsRequest | Specifies the parameters to cancel object runs.
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

    request := helios.ApiCancelObjectRunsRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Object.CancelObjectRuns(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Object.CancelObjectRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelObjectRuns`: CancelObjectRunsResults
    fmt.Fprintf(os.Stdout, "Response from `Object.CancelObjectRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelObjectRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CancelObjectRunsRequest**](CancelObjectRunsRequest.md) | Specifies the parameters to cancel object runs. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**CancelObjectRunsResults**](CancelObjectRunsResults.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConstructMetaInfo

> ConstructMetaInfoResult ConstructMetaInfo(ctx, snapshotId).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Construct meta info for any workflow from object snapshot and some other information.



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
    snapshotId := "snapshotId_example" // string | Specifies the snapshot id.
    body := *openapiclient.NewConstructMetaInfoParams("Environment_example") // ConstructMetaInfoParams | Specifies the parameters to construct meta info for desired workflow.
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

    request := helios.ApiConstructMetaInfoRequest{
        SnapshotId: &snapshotId
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Object.ConstructMetaInfo(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Object.ConstructMetaInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConstructMetaInfo`: ConstructMetaInfoResult
    fmt.Fprintf(os.Stdout, "Response from `Object.ConstructMetaInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | Specifies the snapshot id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiConstructMetaInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ConstructMetaInfoParams**](ConstructMetaInfoParams.md) | Specifies the parameters to construct meta info for desired workflow. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**ConstructMetaInfoResult**](ConstructMetaInfoResult.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilterObjects

> FilteredObjectsResponseBody FilterObjects(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

List all the filtered objects.



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
    body := *openapiclient.NewFilterObjectsRequest("FilterType_example", []openapiclient.Filter{*openapiclient.NewFilter()}, []int64{int64(123)}) // FilterObjectsRequest | Specifies the parameters to filter objects.
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

    request := helios.ApiFilterObjectsRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Object.FilterObjects(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Object.FilterObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilterObjects`: FilteredObjectsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `Object.FilterObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilterObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FilterObjectsRequest**](FilterObjectsRequest.md) | Specifies the parameters to filter objects. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**FilteredObjectsResponseBody**](FilteredObjectsResponseBody.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllIndexedObjectSnapshots

> GetIndexedObjectSnapshotsResponseBody GetAllIndexedObjectSnapshots(ctx, objectId).IndexedObjectName(indexedObjectName).AccessClusterId(accessClusterId).RegionId(regionId).ProtectionGroupId(protectionGroupId).IncludeIndexedSnapshotsOnly(includeIndexedSnapshotsOnly).FromTimeUsecs(fromTimeUsecs).ToTimeUsecs(toTimeUsecs).RunTypes(runTypes).Execute()

Get snapshots of indexed object.



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
    objectId := int64(789) // int64 | Specifies the object id.
    indexedObjectName := "indexedObjectName_example" // string | Specifies the indexed object name.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    protectionGroupId := "protectionGroupId_example" // string | Specifies the protection group id. (optional)
    includeIndexedSnapshotsOnly := true // bool | Specifies whether to only return snapshots which are indexed. In an indexed snapshot files are guaranteed to exist, while in a non-indexed snapshot files may not exist. (optional) (default to false)
    fromTimeUsecs := int64(789) // int64 | Specifies the timestamp in Unix time epoch in microseconds to filter indexed object's snapshots which are taken after this value. (optional)
    toTimeUsecs := int64(789) // int64 | Specifies the timestamp in Unix time epoch in microseconds to filter indexed object's snapshots which are taken before this value. (optional)
    runTypes := []string{"RunTypes_example"} // []string | Filter by run type. Only protection run matching the specified types will be returned. By default, CDP hydration snapshots are not included, unless explicitly queried using this field. (optional)

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

    request := helios.ApiGetAllIndexedObjectSnapshotsRequest{
        ObjectId: &objectId
        IndexedObjectName: &indexedObjectName
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        ProtectionGroupId: &protectionGroupId
        IncludeIndexedSnapshotsOnly: &includeIndexedSnapshotsOnly
        FromTimeUsecs: &fromTimeUsecs
        ToTimeUsecs: &toTimeUsecs
        RunTypes: &runTypes
    }

    resp, r, err := api_client.Object.GetAllIndexedObjectSnapshots(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Object.GetAllIndexedObjectSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllIndexedObjectSnapshots`: GetIndexedObjectSnapshotsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `Object.GetAllIndexedObjectSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **int64** | Specifies the object id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllIndexedObjectSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **indexedObjectName** | **string** | Specifies the indexed object name. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **protectionGroupId** | **string** | Specifies the protection group id. | 
 **includeIndexedSnapshotsOnly** | **bool** | Specifies whether to only return snapshots which are indexed. In an indexed snapshot files are guaranteed to exist, while in a non-indexed snapshot files may not exist. | [default to false]
 **fromTimeUsecs** | **int64** | Specifies the timestamp in Unix time epoch in microseconds to filter indexed object&#39;s snapshots which are taken after this value. | 
 **toTimeUsecs** | **int64** | Specifies the timestamp in Unix time epoch in microseconds to filter indexed object&#39;s snapshots which are taken before this value. | 
 **runTypes** | **[]string** | Filter by run type. Only protection run matching the specified types will be returned. By default, CDP hydration snapshots are not included, unless explicitly queried using this field. | 

### Return type

[**GetIndexedObjectSnapshotsResponseBody**](GetIndexedObjectSnapshotsResponseBody.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndexedObjectSnapshots

> GetIndexedObjectSnapshotsResponseBody GetIndexedObjectSnapshots(ctx, protectionGroupId, objectId).IndexedObjectName(indexedObjectName).AccessClusterId(accessClusterId).RegionId(regionId).IncludeIndexedSnapshotsOnly(includeIndexedSnapshotsOnly).FromTimeUsecs(fromTimeUsecs).ToTimeUsecs(toTimeUsecs).RunTypes(runTypes).Execute()

Get snapshots of indexed object.



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
    protectionGroupId := "protectionGroupId_example" // string | Specifies the protection group id.
    objectId := int64(789) // int64 | Specifies the object id.
    indexedObjectName := "indexedObjectName_example" // string | Specifies the indexed object name.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    includeIndexedSnapshotsOnly := true // bool | Specifies whether to only return snapshots which are indexed. In an indexed snapshots file are guaranteened to exist, while in a non-indexed snapshots file may not exist. (optional) (default to false)
    fromTimeUsecs := int64(789) // int64 | Specifies the timestamp in Unix time epoch in microseconds to filter indexed object's snapshots which are taken after this value. (optional)
    toTimeUsecs := int64(789) // int64 | Specifies the timestamp in Unix time epoch in microseconds to filter indexed object's snapshots which are taken before this value. (optional)
    runTypes := []string{"RunTypes_example"} // []string | Filter by run type. Only protection run matching the specified types will be returned. By default, CDP hydration snapshots are not included, unless explicitly queried using this field. (optional)

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

    request := helios.ApiGetIndexedObjectSnapshotsRequest{
        ProtectionGroupId: &protectionGroupId
        ObjectId: &objectId
        IndexedObjectName: &indexedObjectName
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        IncludeIndexedSnapshotsOnly: &includeIndexedSnapshotsOnly
        FromTimeUsecs: &fromTimeUsecs
        ToTimeUsecs: &toTimeUsecs
        RunTypes: &runTypes
    }

    resp, r, err := api_client.Object.GetIndexedObjectSnapshots(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Object.GetIndexedObjectSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIndexedObjectSnapshots`: GetIndexedObjectSnapshotsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `Object.GetIndexedObjectSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**protectionGroupId** | **string** | Specifies the protection group id. | 
**objectId** | **int64** | Specifies the object id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexedObjectSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **indexedObjectName** | **string** | Specifies the indexed object name. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **includeIndexedSnapshotsOnly** | **bool** | Specifies whether to only return snapshots which are indexed. In an indexed snapshots file are guaranteened to exist, while in a non-indexed snapshots file may not exist. | [default to false]
 **fromTimeUsecs** | **int64** | Specifies the timestamp in Unix time epoch in microseconds to filter indexed object&#39;s snapshots which are taken after this value. | 
 **toTimeUsecs** | **int64** | Specifies the timestamp in Unix time epoch in microseconds to filter indexed object&#39;s snapshots which are taken before this value. | 
 **runTypes** | **[]string** | Filter by run type. Only protection run matching the specified types will be returned. By default, CDP hydration snapshots are not included, unless explicitly queried using this field. | 

### Return type

[**GetIndexedObjectSnapshotsResponseBody**](GetIndexedObjectSnapshotsResponseBody.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMcmObjectSnapshots

> GetObjectSnapshotsResponseBody GetMcmObjectSnapshots(ctx).GlobalId(globalId).RegionId(regionId).Uuid(uuid).FromTimeUsecs(fromTimeUsecs).ToTimeUsecs(toTimeUsecs).Execute()

List the snapshots for a given object.



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
    globalId := "globalId_example" // string | Specifies the global id of the Object.
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    uuid := "uuid_example" // string | Specifies the uuid of the Object. This field is deprecated. (optional)
    fromTimeUsecs := int64(789) // int64 | Specifies the timestamp in Unix time epoch in microseconds to filter Object's snapshots which are taken after this value. (optional)
    toTimeUsecs := int64(789) // int64 | Specifies the timestamp in Unix time epoch in microseconds to filter Object's snapshots which are taken before this value. (optional)

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

    request := helios.ApiGetMcmObjectSnapshotsRequest{
        GlobalId: &globalId
        RegionId: &regionId
        Uuid: &uuid
        FromTimeUsecs: &fromTimeUsecs
        ToTimeUsecs: &toTimeUsecs
    }

    resp, r, err := api_client.Object.GetMcmObjectSnapshots(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Object.GetMcmObjectSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMcmObjectSnapshots`: GetObjectSnapshotsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `Object.GetMcmObjectSnapshots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMcmObjectSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **globalId** | **string** | Specifies the global id of the Object. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **uuid** | **string** | Specifies the uuid of the Object. This field is deprecated. | 
 **fromTimeUsecs** | **int64** | Specifies the timestamp in Unix time epoch in microseconds to filter Object&#39;s snapshots which are taken after this value. | 
 **toTimeUsecs** | **int64** | Specifies the timestamp in Unix time epoch in microseconds to filter Object&#39;s snapshots which are taken before this value. | 

### Return type

[**GetObjectSnapshotsResponseBody**](GetObjectSnapshotsResponseBody.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMcmObjectStats

> ObjectStats GetMcmObjectStats(ctx).GlobalId(globalId).RegionId(regionId).EntityHash(entityHash).RegionIds(regionIds).ClusterIdentifiers(clusterIdentifiers).Execute()

Get stats for a given object.



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
    globalId := "globalId_example" // string | Specifies the global id of the Object. This field is required because we only fetches snapshots stats for one object.
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    entityHash := "entityHash_example" // string | Specifies the entity hash global id of the Object. This field is deprecated. (optional)
    regionIds := []string{"Inner_example"} // []string | Specifies a list of region ids. Only records from clusters having these region ids will be returned. (optional)
    clusterIdentifiers := []string{"Inner_example"} // []string | Specifies the list of cluster identifiers. Format is clusterId:clusterIncarnationId. Only records from clusters having these identifiers will be returned. (optional)

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

    request := helios.ApiGetMcmObjectStatsRequest{
        GlobalId: &globalId
        RegionId: &regionId
        EntityHash: &entityHash
        RegionIds: &regionIds
        ClusterIdentifiers: &clusterIdentifiers
    }

    resp, r, err := api_client.Object.GetMcmObjectStats(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Object.GetMcmObjectStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMcmObjectStats`: ObjectStats
    fmt.Fprintf(os.Stdout, "Response from `Object.GetMcmObjectStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMcmObjectStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **globalId** | **string** | Specifies the global id of the Object. This field is required because we only fetches snapshots stats for one object. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **entityHash** | **string** | Specifies the entity hash global id of the Object. This field is deprecated. | 
 **regionIds** | **[]string** | Specifies a list of region ids. Only records from clusters having these region ids will be returned. | 
 **clusterIdentifiers** | **[]string** | Specifies the list of cluster identifiers. Format is clusterId:clusterIncarnationId. Only records from clusters having these identifiers will be returned. | 

### Return type

[**ObjectStats**](ObjectStats.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMcmObjectsActivity

> McmObjectsActivity GetMcmObjectsActivity(ctx).RegionId(regionId).RegionIds(regionIds).Body(body).Execute()

Get Object activity on Helios.



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
    regionIds := []string{"Inner_example"} // []string | Filter by a list of region ids. (optional)
    body := *openapiclient.NewGetMcmObjectsActivityReqParams() // GetMcmObjectsActivityReqParams | Request parameters to filter object activity. (optional)

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

    request := helios.ApiGetMcmObjectsActivityRequest{
        RegionId: &regionId
        RegionIds: &regionIds
        Body: &body
    }

    resp, r, err := api_client.Object.GetMcmObjectsActivity(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Object.GetMcmObjectsActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMcmObjectsActivity`: McmObjectsActivity
    fmt.Fprintf(os.Stdout, "Response from `Object.GetMcmObjectsActivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMcmObjectsActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **regionIds** | **[]string** | Filter by a list of region ids. | 
 **body** | [**GetMcmObjectsActivityReqParams**](GetMcmObjectsActivityReqParams.md) | Request parameters to filter object activity. | 

### Return type

[**McmObjectsActivity**](McmObjectsActivity.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectIdentifiers

> LocalGlobalObjectIdList GetObjectIdentifiers(ctx).AccessClusterId(accessClusterId).RegionId(regionId).GlobalIds(globalIds).LocalIds(localIds).Execute()

Get Object Identifiers



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
    globalIds := []string{"Inner_example"} // []string | Get the object identifier matching specified global IDs. (optional)
    localIds := []int64{int64(123)} // []int64 | Get the object identifier matching specified local IDs. (optional)

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

    request := helios.ApiGetObjectIdentifiersRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        GlobalIds: &globalIds
        LocalIds: &localIds
    }

    resp, r, err := api_client.Object.GetObjectIdentifiers(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Object.GetObjectIdentifiers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectIdentifiers`: LocalGlobalObjectIdList
    fmt.Fprintf(os.Stdout, "Response from `Object.GetObjectIdentifiers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectIdentifiersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **globalIds** | **[]string** | Get the object identifier matching specified global IDs. | 
 **localIds** | **[]int64** | Get the object identifier matching specified local IDs. | 

### Return type

[**LocalGlobalObjectIdList**](LocalGlobalObjectIdList.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectRunByRunId

> ObjectProtectionRunSummary GetObjectRunByRunId(ctx, id, runId).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Get a run for an object.



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
    id := int64(789) // int64 | Specifies a unique id of the object.
    runId := "runId_example" // string | Specifies the id of the run.
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

    request := helios.ApiGetObjectRunByRunIdRequest{
        Id: &id
        RunId: &runId
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Object.GetObjectRunByRunId(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Object.GetObjectRunByRunId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectRunByRunId`: ObjectProtectionRunSummary
    fmt.Fprintf(os.Stdout, "Response from `Object.GetObjectRunByRunId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the object. | 
**runId** | **string** | Specifies the id of the run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectRunByRunIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**ObjectProtectionRunSummary**](ObjectProtectionRunSummary.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectRuns

> GetObjectRunsResponseBody GetObjectRuns(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).RunId(runId).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).TenantIds(tenantIds).IncludeTenants(includeTenants).RunTypes(runTypes).LocalBackupObjectStatus(localBackupObjectStatus).ReplicationObjectStatus(replicationObjectStatus).ArchivalObjectStatus(archivalObjectStatus).CloudSpinRunStatus(cloudSpinRunStatus).NumRuns(numRuns).PaginationCookie(paginationCookie).ExcludeNonRestorableRuns(excludeNonRestorableRuns).Execute()

Get the list of runs for an object.



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
    id := int64(789) // int64 | Specifies a unique id of the object.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    runId := "runId_example" // string | Specifies a unique id of the run. (optional)
    startTimeUsecs := int64(789) // int64 | Filter by a start time when the run starts. Specify the start time as a Unix epoch Timestamp (in microseconds). (optional)
    endTimeUsecs := int64(789) // int64 | Filter by a end time when the run starts. Specify the start time as a Unix epoch Timestamp (in microseconds). (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
    includeTenants := true // bool | If true, the response will include Protection Group Runs which were created by all tenants which the current user has permission to see. If false, then only Protection Group Runs created by the current user will be returned. (optional)
    runTypes := []string{"RunTypes_example"} // []string | Filter by run type. Only protection run matching the specified types will be returned. (optional)
    localBackupObjectStatus := []string{"LocalBackupObjectStatus_example"} // []string | Specifies a list of status for the object in the backup run. (optional)
    replicationObjectStatus := []string{"ReplicationObjectStatus_example"} // []string | Specifies a list of status for the object in the replication run. (optional)
    archivalObjectStatus := []string{"ArchivalObjectStatus_example"} // []string | Specifies a list of status for the object in the archival run. (optional)
    cloudSpinRunStatus := []string{"CloudSpinRunStatus_example"} // []string | Specifies a list of status for the object in the cloud spin run. (optional)
    numRuns := int64(789) // int64 | Specifies the max number of runs. If not specified, at most 100 runs will be returned. (optional)
    paginationCookie := "paginationCookie_example" // string | Specifies the pagination cookie with which subsequent parts of the response can be fetched. Users can use this to get next runs (optional)
    excludeNonRestorableRuns := true // bool | Specifies whether to exclude non restorable runs. Run is treated restorable only if there is atleast one object snapshot (which may be either a local or an archival snapshot) which is not deleted or expired. Default value is false. (optional) (default to false)

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

    request := helios.ApiGetObjectRunsRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        RunId: &runId
        StartTimeUsecs: &startTimeUsecs
        EndTimeUsecs: &endTimeUsecs
        TenantIds: &tenantIds
        IncludeTenants: &includeTenants
        RunTypes: &runTypes
        LocalBackupObjectStatus: &localBackupObjectStatus
        ReplicationObjectStatus: &replicationObjectStatus
        ArchivalObjectStatus: &archivalObjectStatus
        CloudSpinRunStatus: &cloudSpinRunStatus
        NumRuns: &numRuns
        PaginationCookie: &paginationCookie
        ExcludeNonRestorableRuns: &excludeNonRestorableRuns
    }

    resp, r, err := api_client.Object.GetObjectRuns(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Object.GetObjectRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectRuns`: GetObjectRunsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `Object.GetObjectRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **runId** | **string** | Specifies a unique id of the run. | 
 **startTimeUsecs** | **int64** | Filter by a start time when the run starts. Specify the start time as a Unix epoch Timestamp (in microseconds). | 
 **endTimeUsecs** | **int64** | Filter by a end time when the run starts. Specify the start time as a Unix epoch Timestamp (in microseconds). | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **includeTenants** | **bool** | If true, the response will include Protection Group Runs which were created by all tenants which the current user has permission to see. If false, then only Protection Group Runs created by the current user will be returned. | 
 **runTypes** | **[]string** | Filter by run type. Only protection run matching the specified types will be returned. | 
 **localBackupObjectStatus** | **[]string** | Specifies a list of status for the object in the backup run. | 
 **replicationObjectStatus** | **[]string** | Specifies a list of status for the object in the replication run. | 
 **archivalObjectStatus** | **[]string** | Specifies a list of status for the object in the archival run. | 
 **cloudSpinRunStatus** | **[]string** | Specifies a list of status for the object in the cloud spin run. | 
 **numRuns** | **int64** | Specifies the max number of runs. If not specified, at most 100 runs will be returned. | 
 **paginationCookie** | **string** | Specifies the pagination cookie with which subsequent parts of the response can be fetched. Users can use this to get next runs | 
 **excludeNonRestorableRuns** | **bool** | Specifies whether to exclude non restorable runs. Run is treated restorable only if there is atleast one object snapshot (which may be either a local or an archival snapshot) which is not deleted or expired. Default value is false. | [default to false]

### Return type

[**GetObjectRunsResponseBody**](GetObjectRunsResponseBody.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectSnapshotInfo

> ObjectSnapshot GetObjectSnapshotInfo(ctx, snapshotId).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Get details of object snapshot.



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
    snapshotId := "snapshotId_example" // string | Specifies the snapshot id.
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

    request := helios.ApiGetObjectSnapshotInfoRequest{
        SnapshotId: &snapshotId
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Object.GetObjectSnapshotInfo(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Object.GetObjectSnapshotInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectSnapshotInfo`: ObjectSnapshot
    fmt.Fprintf(os.Stdout, "Response from `Object.GetObjectSnapshotInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | Specifies the snapshot id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectSnapshotInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**ObjectSnapshot**](ObjectSnapshot.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectSnapshotVolumeInfo

> ObjectSnapshotVolumeInfo GetObjectSnapshotVolumeInfo(ctx, snapshotId).AccessClusterId(accessClusterId).RegionId(regionId).IncludeSupportedOnly(includeSupportedOnly).PointInTimeUsecs(pointInTimeUsecs).Execute()

Get volume info of object snapshot.



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
    snapshotId := "snapshotId_example" // string | Specifies the snapshot id.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    includeSupportedOnly := true // bool | Specifies whether to only return supported volumes. (optional)
    pointInTimeUsecs := float32(8.14) // float32 | Specifies the point-in-time timestamp (in usecs from epoch) between snapshots for which the volume info is to be returned. (optional)

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

    request := helios.ApiGetObjectSnapshotVolumeInfoRequest{
        SnapshotId: &snapshotId
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        IncludeSupportedOnly: &includeSupportedOnly
        PointInTimeUsecs: &pointInTimeUsecs
    }

    resp, r, err := api_client.Object.GetObjectSnapshotVolumeInfo(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Object.GetObjectSnapshotVolumeInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectSnapshotVolumeInfo`: ObjectSnapshotVolumeInfo
    fmt.Fprintf(os.Stdout, "Response from `Object.GetObjectSnapshotVolumeInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | Specifies the snapshot id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectSnapshotVolumeInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **includeSupportedOnly** | **bool** | Specifies whether to only return supported volumes. | 
 **pointInTimeUsecs** | **float32** | Specifies the point-in-time timestamp (in usecs from epoch) between snapshots for which the volume info is to be returned. | 

### Return type

[**ObjectSnapshotVolumeInfo**](ObjectSnapshotVolumeInfo.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectSnapshots

> GetObjectSnapshotsResponseBody GetObjectSnapshots(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).FromTimeUsecs(fromTimeUsecs).ToTimeUsecs(toTimeUsecs).RunStartFromTimeUsecs(runStartFromTimeUsecs).RunStartToTimeUsecs(runStartToTimeUsecs).RunTypes(runTypes).ProtectionGroupIds(protectionGroupIds).RunInstanceIds(runInstanceIds).RegionIds(regionIds).Execute()

List the snapshots for a given object.



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
    id := int64(789) // int64 | Specifies the id of the Object.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    fromTimeUsecs := int64(789) // int64 | Specifies the timestamp in Unix time epoch in microseconds to filter Object's snapshots which were taken after this value. (optional)
    toTimeUsecs := int64(789) // int64 | Specifies the timestamp in Unix time epoch in microseconds to filter Object's snapshots which were taken before this value. (optional)
    runStartFromTimeUsecs := int64(789) // int64 | Specifies the timestamp in Unix time epoch in microseconds to filter Object's snapshots which were run after this value. (optional)
    runStartToTimeUsecs := int64(789) // int64 | Specifies the timestamp in Unix time epoch in microseconds to filter Object's snapshots which were run before this value. (optional)
    runTypes := []string{"RunTypes_example"} // []string | Filter by run type. Only protection run matching the specified types will be returned. By default, CDP hydration snapshots are not included, unless explicitly queried using this field. (optional)
    protectionGroupIds := []string{"Inner_example"} // []string | If specified, this returns only the snapshots of the specified object ID, which belong to the provided protection group IDs. (optional)
    runInstanceIds := []int64{int64(123)} // []int64 | Filter by a list run instance ids. If specified, only snapshots created by these protection runs will be returned. (optional)
    regionIds := []string{"Inner_example"} // []string | Filter by a list of region ids. (optional)

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

    request := helios.ApiGetObjectSnapshotsRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        FromTimeUsecs: &fromTimeUsecs
        ToTimeUsecs: &toTimeUsecs
        RunStartFromTimeUsecs: &runStartFromTimeUsecs
        RunStartToTimeUsecs: &runStartToTimeUsecs
        RunTypes: &runTypes
        ProtectionGroupIds: &protectionGroupIds
        RunInstanceIds: &runInstanceIds
        RegionIds: &regionIds
    }

    resp, r, err := api_client.Object.GetObjectSnapshots(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Object.GetObjectSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectSnapshots`: GetObjectSnapshotsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `Object.GetObjectSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the id of the Object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **fromTimeUsecs** | **int64** | Specifies the timestamp in Unix time epoch in microseconds to filter Object&#39;s snapshots which were taken after this value. | 
 **toTimeUsecs** | **int64** | Specifies the timestamp in Unix time epoch in microseconds to filter Object&#39;s snapshots which were taken before this value. | 
 **runStartFromTimeUsecs** | **int64** | Specifies the timestamp in Unix time epoch in microseconds to filter Object&#39;s snapshots which were run after this value. | 
 **runStartToTimeUsecs** | **int64** | Specifies the timestamp in Unix time epoch in microseconds to filter Object&#39;s snapshots which were run before this value. | 
 **runTypes** | **[]string** | Filter by run type. Only protection run matching the specified types will be returned. By default, CDP hydration snapshots are not included, unless explicitly queried using this field. | 
 **protectionGroupIds** | **[]string** | If specified, this returns only the snapshots of the specified object ID, which belong to the provided protection group IDs. | 
 **runInstanceIds** | **[]int64** | Filter by a list run instance ids. If specified, only snapshots created by these protection runs will be returned. | 
 **regionIds** | **[]string** | Filter by a list of region ids. | 

### Return type

[**GetObjectSnapshotsResponseBody**](GetObjectSnapshotsResponseBody.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStats

> ObjectStats GetObjectStats(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).RegionIds(regionIds).Execute()

Get stats for a given object.



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
    id := int64(789) // int64 | Specifies the id of the Object.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    regionIds := []string{"Inner_example"} // []string | Filter by a list of region ids. (optional)

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

    request := helios.ApiGetObjectStatsRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        RegionIds: &regionIds
    }

    resp, r, err := api_client.Object.GetObjectStats(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Object.GetObjectStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectStats`: ObjectStats
    fmt.Fprintf(os.Stdout, "Response from `Object.GetObjectStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the id of the Object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **regionIds** | **[]string** | Filter by a list of region ids. | 

### Return type

[**ObjectStats**](ObjectStats.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectTree

> ObjectWithChildren GetObjectTree(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Get the objects tree hierarchy for for an Object.



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
    id := int64(789) // int64 | Specifies the id of the Object.
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

    request := helios.ApiGetObjectTreeRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Object.GetObjectTree(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Object.GetObjectTree``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectTree`: ObjectWithChildren
    fmt.Fprintf(os.Stdout, "Response from `Object.GetObjectTree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the id of the Object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectTreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**ObjectWithChildren**](ObjectWithChildren.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectsLastRun

> ObjectsLastRun GetObjectsLastRun(ctx).AccessClusterId(accessClusterId).RegionId(regionId).Ids(ids).TenantIds(tenantIds).IncludeTenants(includeTenants).PaginationCookie(paginationCookie).Count(count).Execute()

Get last protection run of objects.



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
    ids := []int64{int64(123)} // []int64 | Specifies a list of object ids, only last runs for these objects will be returned. (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
    includeTenants := true // bool | If true, the response will include Objects which belongs to all tenants which the current user has permission to see. (optional)
    paginationCookie := "paginationCookie_example" // string | Specifies the pagination cookie with which subsequent parts of the response can be fetched. (optional)
    count := int32(56) // int32 | Specifies the number of objects to be fetched for the specified pagination cookie. (optional)

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

    request := helios.ApiGetObjectsLastRunRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        Ids: &ids
        TenantIds: &tenantIds
        IncludeTenants: &includeTenants
        PaginationCookie: &paginationCookie
        Count: &count
    }

    resp, r, err := api_client.Object.GetObjectsLastRun(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Object.GetObjectsLastRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectsLastRun`: ObjectsLastRun
    fmt.Fprintf(os.Stdout, "Response from `Object.GetObjectsLastRun`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectsLastRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **ids** | **[]int64** | Specifies a list of object ids, only last runs for these objects will be returned. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **includeTenants** | **bool** | If true, the response will include Objects which belongs to all tenants which the current user has permission to see. | 
 **paginationCookie** | **string** | Specifies the pagination cookie with which subsequent parts of the response can be fetched. | 
 **count** | **int32** | Specifies the number of objects to be fetched for the specified pagination cookie. | 

### Return type

[**ObjectsLastRun**](ObjectsLastRun.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPITRangesForProtectedObject

> GetPITRangesProtectedObjectResponseBody GetPITRangesForProtectedObject(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).FromTimeUsecs(fromTimeUsecs).ToTimeUsecs(toTimeUsecs).ProtectionGroupIds(protectionGroupIds).Execute()

Get PIT ranges for an object



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
    id := int64(789) // int64 | Specifies the ID of the protected object.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    fromTimeUsecs := int64(789) // int64 | If specified, return the restore ranges that lie after this timestamp. This parameter is specified as the timestamp in Unix time epoch in microseconds. (optional)
    toTimeUsecs := int64(789) // int64 | If specified, return the restore ranges that lie before this timestamp. This parameter is specified as the timestamp in Unix time epoch in microseconds. (optional)
    protectionGroupIds := []string{"Inner_example"} // []string | If specified, return only the points in time corresponding to these protection group IDs. (optional)

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

    request := helios.ApiGetPITRangesForProtectedObjectRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        FromTimeUsecs: &fromTimeUsecs
        ToTimeUsecs: &toTimeUsecs
        ProtectionGroupIds: &protectionGroupIds
    }

    resp, r, err := api_client.Object.GetPITRangesForProtectedObject(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Object.GetPITRangesForProtectedObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPITRangesForProtectedObject`: GetPITRangesProtectedObjectResponseBody
    fmt.Fprintf(os.Stdout, "Response from `Object.GetPITRangesForProtectedObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the ID of the protected object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPITRangesForProtectedObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **fromTimeUsecs** | **int64** | If specified, return the restore ranges that lie after this timestamp. This parameter is specified as the timestamp in Unix time epoch in microseconds. | 
 **toTimeUsecs** | **int64** | If specified, return the restore ranges that lie before this timestamp. This parameter is specified as the timestamp in Unix time epoch in microseconds. | 
 **protectionGroupIds** | **[]string** | If specified, return only the points in time corresponding to these protection group IDs. | 

### Return type

[**GetPITRangesProtectedObjectResponseBody**](GetPITRangesProtectedObjectResponseBody.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtectedObjectOfAnyTypeById

> ProtectedObjectInfo GetProtectedObjectOfAnyTypeById(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).ObjectActionKey(objectActionKey).OnlyProtectedObjects(onlyProtectedObjects).StorageDomainId(storageDomainId).Environments(environments).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeLastRunInfo(includeLastRunInfo).OnlyAutoProtectedObjects(onlyAutoProtectedObjects).OnlyLeafObjects(onlyLeafObjects).Execute()

Get an Object.



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
    id := int64(789) // int64 | Specifies the id of the Object.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    objectActionKey := []string{"ObjectActionKey_example"} // []string | Filter by ObjectActionKey, uniquely represent protection of an object. An object can be protected in multiple ways but atmost once for a given combination of ObjectActionKey, when specified Only objects of given action_key are returned for corresponding object id and this vec's size needs to be same as 'id'. (optional)
    onlyProtectedObjects := true // bool | If true, the response will include only objects which have been protected. (optional)
    storageDomainId := int64(789) // int64 | Filter by Storage Domain id. Only Objects protected to this Storage Domain will be returned. (optional)
    environments := []string{"Environments_example"} // []string | Filter by environment types such as 'kVMware', 'kView', etc. Only Protected objects protecting the specified environment types are returned. (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
    includeTenants := true // bool | If true, the response will include Objects which were protected by all tenants which the current user has permission to see. If false, then only objects protected by the current user will be returned. (optional)
    includeLastRunInfo := true // bool | If true, the response will include information about the last protection run on this object. (optional)
    onlyAutoProtectedObjects := true // bool | If true, the response will include only the auto protected objects. (optional)
    onlyLeafObjects := true // bool | If true, the response will include only the leaf level objects. (optional)

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

    request := helios.ApiGetProtectedObjectOfAnyTypeByIdRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        ObjectActionKey: &objectActionKey
        OnlyProtectedObjects: &onlyProtectedObjects
        StorageDomainId: &storageDomainId
        Environments: &environments
        TenantIds: &tenantIds
        IncludeTenants: &includeTenants
        IncludeLastRunInfo: &includeLastRunInfo
        OnlyAutoProtectedObjects: &onlyAutoProtectedObjects
        OnlyLeafObjects: &onlyLeafObjects
    }

    resp, r, err := api_client.Object.GetProtectedObjectOfAnyTypeById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Object.GetProtectedObjectOfAnyTypeById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtectedObjectOfAnyTypeById`: ProtectedObjectInfo
    fmt.Fprintf(os.Stdout, "Response from `Object.GetProtectedObjectOfAnyTypeById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the id of the Object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectedObjectOfAnyTypeByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **objectActionKey** | **[]string** | Filter by ObjectActionKey, uniquely represent protection of an object. An object can be protected in multiple ways but atmost once for a given combination of ObjectActionKey, when specified Only objects of given action_key are returned for corresponding object id and this vec&#39;s size needs to be same as &#39;id&#39;. | 
 **onlyProtectedObjects** | **bool** | If true, the response will include only objects which have been protected. | 
 **storageDomainId** | **int64** | Filter by Storage Domain id. Only Objects protected to this Storage Domain will be returned. | 
 **environments** | **[]string** | Filter by environment types such as &#39;kVMware&#39;, &#39;kView&#39;, etc. Only Protected objects protecting the specified environment types are returned. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **includeTenants** | **bool** | If true, the response will include Objects which were protected by all tenants which the current user has permission to see. If false, then only objects protected by the current user will be returned. | 
 **includeLastRunInfo** | **bool** | If true, the response will include information about the last protection run on this object. | 
 **onlyAutoProtectedObjects** | **bool** | If true, the response will include only the auto protected objects. | 
 **onlyLeafObjects** | **bool** | If true, the response will include only the leaf level objects. | 

### Return type

[**ProtectedObjectInfo**](ProtectedObjectInfo.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtectedObjectsOfAnyType

> GetProtectedObjectsResponse GetProtectedObjectsOfAnyType(ctx).AccessClusterId(accessClusterId).RegionId(regionId).Ids(ids).ObjectActionKeys(objectActionKeys).PolicyIds(policyIds).ParentId(parentId).OnlyProtectedObjects(onlyProtectedObjects).StorageDomainId(storageDomainId).Environments(environments).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeLastRunInfo(includeLastRunInfo).OnlyAutoProtectedObjects(onlyAutoProtectedObjects).OnlyLeafObjects(onlyLeafObjects).RegionIds(regionIds).ProtectionTypes(protectionTypes).Execute()

Get Objects.



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
    ids := []int64{int64(123)} // []int64 | Filter by a list of Object ids. (optional)
    objectActionKeys := []string{"ObjectActionKeys_example"} // []string | Filter by ObjectActionKey, uniquely represent protection of an object. An object can be protected in multiple ways but atmost once for a given combination of ObjectActionKey, when specified Only objects of given action_key are returned for corresponding object id and this vec's size needs to be same as 'ids'. (optional)
    policyIds := []string{"Inner_example"} // []string | Filter by Policy ids that are associated with Protected Objects. (optional)
    parentId := int64(789) // int64 | Filter by Parent Id. Parent id is a unique object Id which may contain protected objects underneath in the source tree. (optional)
    onlyProtectedObjects := true // bool | If true, the response will include only objects which have been protected. (optional)
    storageDomainId := int64(789) // int64 | Filter by Storage Domain id. Only Objects protected to this Storage Domain will be returned. (optional)
    environments := []string{"Environments_example"} // []string | Filter by environment types such as 'kVMware', 'kView', etc. Only Protected objects protecting the specified environment types are returned. (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
    includeTenants := true // bool | If true, the response will include Objects which were protected by all tenants which the current user has permission to see. If false, then only objects protected by the current user will be returned. (optional)
    includeLastRunInfo := true // bool | If true, the response will include information about the last protection run on this object. (optional)
    onlyAutoProtectedObjects := true // bool | If true, the response will include only the auto protected objects. (optional)
    onlyLeafObjects := true // bool | If true, the response will include only the leaf level objects. (optional)
    regionIds := []string{"Inner_example"} // []string | Filter by a list of region ids. (optional)
    protectionTypes := []string{"ProtectionTypes_example"} // []string | Filter by a list of protection types. (optional)

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

    request := helios.ApiGetProtectedObjectsOfAnyTypeRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        Ids: &ids
        ObjectActionKeys: &objectActionKeys
        PolicyIds: &policyIds
        ParentId: &parentId
        OnlyProtectedObjects: &onlyProtectedObjects
        StorageDomainId: &storageDomainId
        Environments: &environments
        TenantIds: &tenantIds
        IncludeTenants: &includeTenants
        IncludeLastRunInfo: &includeLastRunInfo
        OnlyAutoProtectedObjects: &onlyAutoProtectedObjects
        OnlyLeafObjects: &onlyLeafObjects
        RegionIds: &regionIds
        ProtectionTypes: &protectionTypes
    }

    resp, r, err := api_client.Object.GetProtectedObjectsOfAnyType(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Object.GetProtectedObjectsOfAnyType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtectedObjectsOfAnyType`: GetProtectedObjectsResponse
    fmt.Fprintf(os.Stdout, "Response from `Object.GetProtectedObjectsOfAnyType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectedObjectsOfAnyTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **ids** | **[]int64** | Filter by a list of Object ids. | 
 **objectActionKeys** | **[]string** | Filter by ObjectActionKey, uniquely represent protection of an object. An object can be protected in multiple ways but atmost once for a given combination of ObjectActionKey, when specified Only objects of given action_key are returned for corresponding object id and this vec&#39;s size needs to be same as &#39;ids&#39;. | 
 **policyIds** | **[]string** | Filter by Policy ids that are associated with Protected Objects. | 
 **parentId** | **int64** | Filter by Parent Id. Parent id is a unique object Id which may contain protected objects underneath in the source tree. | 
 **onlyProtectedObjects** | **bool** | If true, the response will include only objects which have been protected. | 
 **storageDomainId** | **int64** | Filter by Storage Domain id. Only Objects protected to this Storage Domain will be returned. | 
 **environments** | **[]string** | Filter by environment types such as &#39;kVMware&#39;, &#39;kView&#39;, etc. Only Protected objects protecting the specified environment types are returned. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **includeTenants** | **bool** | If true, the response will include Objects which were protected by all tenants which the current user has permission to see. If false, then only objects protected by the current user will be returned. | 
 **includeLastRunInfo** | **bool** | If true, the response will include information about the last protection run on this object. | 
 **onlyAutoProtectedObjects** | **bool** | If true, the response will include only the auto protected objects. | 
 **onlyLeafObjects** | **bool** | If true, the response will include only the leaf level objects. | 
 **regionIds** | **[]string** | Filter by a list of region ids. | 
 **protectionTypes** | **[]string** | Filter by a list of protection types. | 

### Return type

[**GetProtectedObjectsResponse**](GetProtectedObjectsResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnapshotDiff

> SnapshotDiffResult GetSnapshotDiff(ctx, id).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Get diff between two snapshots of a given object.



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
    id := int64(789) // int64 | 
    body := *openapiclient.NewSnapshotDiffParams(int64(123), int64(123), int64(123), "EntityType_example", int64(123), int64(123), int64(123), int64(123), int64(123), int64(123)) // SnapshotDiffParams | 
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

    request := helios.ApiGetSnapshotDiffRequest{
        Id: &id
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Object.GetSnapshotDiff(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Object.GetSnapshotDiff``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSnapshotDiff`: SnapshotDiffResult
    fmt.Fprintf(os.Stdout, "Response from `Object.GetSnapshotDiff`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnapshotDiffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SnapshotDiffParams**](SnapshotDiffParams.md) |  | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**SnapshotDiffResult**](SnapshotDiffResult.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceHierarchyObjects

> SourceHierarchyObjectSummaries GetSourceHierarchyObjects(ctx, sourceId).AccessClusterId(accessClusterId).RegionId(regionId).ParentId(parentId).TenantIds(tenantIds).IncludeTenants(includeTenants).VmwareObjectTypes(vmwareObjectTypes).NetappObjectTypes(netappObjectTypes).O365ObjectTypes(o365ObjectTypes).CassandraObjectTypes(cassandraObjectTypes).MongodbObjectTypes(mongodbObjectTypes).CouchbaseObjectTypes(couchbaseObjectTypes).HdfsObjectTypes(hdfsObjectTypes).HbaseObjectTypes(hbaseObjectTypes).HiveObjectTypes(hiveObjectTypes).HypervObjectTypes(hypervObjectTypes).AzureObjectTypes(azureObjectTypes).KvmObjectTypes(kvmObjectTypes).AwsObjectTypes(awsObjectTypes).GcpObjectTypes(gcpObjectTypes).AcropolisObjectTypes(acropolisObjectTypes).GenericNasObjectTypes(genericNasObjectTypes).IsilonObjectTypes(isilonObjectTypes).FlashbladeObjectTypes(flashbladeObjectTypes).ElastifileObjectTypes(elastifileObjectTypes).GpfsObjectTypes(gpfsObjectTypes).PureObjectTypes(pureObjectTypes).NimbleObjectTypes(nimbleObjectTypes).PhysicalObjectTypes(physicalObjectTypes).KubernetesObjectTypes(kubernetesObjectTypes).ExchangeObjectTypes(exchangeObjectTypes).AdObjectTypes(adObjectTypes).MssqlObjectTypes(mssqlObjectTypes).OracleObjectTypes(oracleObjectTypes).Execute()

List objects on a source which can be used for data protection.



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
    sourceId := int64(789) // int64 | Specifies the source ID for which objects should be returned.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    parentId := int64(789) // int64 | Specifies the parent ID under which objects should be returned. (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
    includeTenants := true // bool | If true, the response will include Objects which belongs to all tenants which the current user has permission to see. (optional)
    vmwareObjectTypes := []string{"VmwareObjectTypes_example"} // []string | Specifies the VMware object types to filter objects. (optional)
    netappObjectTypes := []string{"NetappObjectTypes_example"} // []string | Specifies the Netapp object types to filter objects. (optional)
    o365ObjectTypes := []string{"O365ObjectTypes_example"} // []string | Specifies the Office 365 object types to filter objects. (optional)
    cassandraObjectTypes := []string{"CassandraObjectTypes_example"} // []string | Specifies the Cassandra object types to filter objects. (optional)
    mongodbObjectTypes := []string{"MongodbObjectTypes_example"} // []string | Specifies the Mongo DB object types to filter objects. (optional)
    couchbaseObjectTypes := []string{"CouchbaseObjectTypes_example"} // []string | Specifies the Couchbase object types to filter objects. (optional)
    hdfsObjectTypes := []string{"HdfsObjectTypes_example"} // []string | Specifies the HDFS object types to filter objects. (optional)
    hbaseObjectTypes := []string{"HbaseObjectTypes_example"} // []string | Specifies the Hbase object types to filter objects. (optional)
    hiveObjectTypes := []string{"HiveObjectTypes_example"} // []string | Specifies the Hive object types to filter objects. (optional)
    hypervObjectTypes := []string{"HypervObjectTypes_example"} // []string | Specifies the HyperV object types to filter objects. (optional)
    azureObjectTypes := []string{"AzureObjectTypes_example"} // []string | Specifies the Azure object types to filter objects. (optional)
    kvmObjectTypes := []string{"KvmObjectTypes_example"} // []string | Specifies the KVM object types to filter objects. (optional)
    awsObjectTypes := []string{"AwsObjectTypes_example"} // []string | Specifies the AWS object types to filter objects. (optional)
    gcpObjectTypes := []string{"GcpObjectTypes_example"} // []string | Specifies the GCP object types to filter objects. (optional)
    acropolisObjectTypes := []string{"AcropolisObjectTypes_example"} // []string | Specifies the Acropolis object types to filter objects. (optional)
    genericNasObjectTypes := []string{"GenericNasObjectTypes_example"} // []string | Specifies the generic NAS object types to filter objects. (optional)
    isilonObjectTypes := []string{"IsilonObjectTypes_example"} // []string | Specifies the Isilon object types to filter objects. (optional)
    flashbladeObjectTypes := []string{"FlashbladeObjectTypes_example"} // []string | Specifies the Flashblade object types to filter objects. (optional)
    elastifileObjectTypes := []string{"ElastifileObjectTypes_example"} // []string | Specifies the Elastifile object types to filter objects. (optional)
    gpfsObjectTypes := []string{"GpfsObjectTypes_example"} // []string | Specifies the GPFS object types to filter objects. (optional)
    pureObjectTypes := []string{"PureObjectTypes_example"} // []string | Specifies the Pure object types to filter objects. (optional)
    nimbleObjectTypes := []string{"NimbleObjectTypes_example"} // []string | Specifies the Nimble object types to filter objects. (optional)
    physicalObjectTypes := []string{"PhysicalObjectTypes_example"} // []string | Specifies the Physical object types to filter objects. (optional)
    kubernetesObjectTypes := []string{"KubernetesObjectTypes_example"} // []string | Specifies the Kubernetes object types to filter objects. (optional)
    exchangeObjectTypes := []string{"ExchangeObjectTypes_example"} // []string | Specifies the Exchange object types to filter objects. (optional)
    adObjectTypes := []string{"AdObjectTypes_example"} // []string | Specifies the AD object types to filter objects. (optional)
    mssqlObjectTypes := []string{"MssqlObjectTypes_example"} // []string | Specifies the MSSQL object types to filter objects. (optional)
    oracleObjectTypes := []string{"OracleObjectTypes_example"} // []string | Specifies the Oracle object types to filter objects. (optional)

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

    request := helios.ApiGetSourceHierarchyObjectsRequest{
        SourceId: &sourceId
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        ParentId: &parentId
        TenantIds: &tenantIds
        IncludeTenants: &includeTenants
        VmwareObjectTypes: &vmwareObjectTypes
        NetappObjectTypes: &netappObjectTypes
        O365ObjectTypes: &o365ObjectTypes
        CassandraObjectTypes: &cassandraObjectTypes
        MongodbObjectTypes: &mongodbObjectTypes
        CouchbaseObjectTypes: &couchbaseObjectTypes
        HdfsObjectTypes: &hdfsObjectTypes
        HbaseObjectTypes: &hbaseObjectTypes
        HiveObjectTypes: &hiveObjectTypes
        HypervObjectTypes: &hypervObjectTypes
        AzureObjectTypes: &azureObjectTypes
        KvmObjectTypes: &kvmObjectTypes
        AwsObjectTypes: &awsObjectTypes
        GcpObjectTypes: &gcpObjectTypes
        AcropolisObjectTypes: &acropolisObjectTypes
        GenericNasObjectTypes: &genericNasObjectTypes
        IsilonObjectTypes: &isilonObjectTypes
        FlashbladeObjectTypes: &flashbladeObjectTypes
        ElastifileObjectTypes: &elastifileObjectTypes
        GpfsObjectTypes: &gpfsObjectTypes
        PureObjectTypes: &pureObjectTypes
        NimbleObjectTypes: &nimbleObjectTypes
        PhysicalObjectTypes: &physicalObjectTypes
        KubernetesObjectTypes: &kubernetesObjectTypes
        ExchangeObjectTypes: &exchangeObjectTypes
        AdObjectTypes: &adObjectTypes
        MssqlObjectTypes: &mssqlObjectTypes
        OracleObjectTypes: &oracleObjectTypes
    }

    resp, r, err := api_client.Object.GetSourceHierarchyObjects(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Object.GetSourceHierarchyObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceHierarchyObjects`: SourceHierarchyObjectSummaries
    fmt.Fprintf(os.Stdout, "Response from `Object.GetSourceHierarchyObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **int64** | Specifies the source ID for which objects should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceHierarchyObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **parentId** | **int64** | Specifies the parent ID under which objects should be returned. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **includeTenants** | **bool** | If true, the response will include Objects which belongs to all tenants which the current user has permission to see. | 
 **vmwareObjectTypes** | **[]string** | Specifies the VMware object types to filter objects. | 
 **netappObjectTypes** | **[]string** | Specifies the Netapp object types to filter objects. | 
 **o365ObjectTypes** | **[]string** | Specifies the Office 365 object types to filter objects. | 
 **cassandraObjectTypes** | **[]string** | Specifies the Cassandra object types to filter objects. | 
 **mongodbObjectTypes** | **[]string** | Specifies the Mongo DB object types to filter objects. | 
 **couchbaseObjectTypes** | **[]string** | Specifies the Couchbase object types to filter objects. | 
 **hdfsObjectTypes** | **[]string** | Specifies the HDFS object types to filter objects. | 
 **hbaseObjectTypes** | **[]string** | Specifies the Hbase object types to filter objects. | 
 **hiveObjectTypes** | **[]string** | Specifies the Hive object types to filter objects. | 
 **hypervObjectTypes** | **[]string** | Specifies the HyperV object types to filter objects. | 
 **azureObjectTypes** | **[]string** | Specifies the Azure object types to filter objects. | 
 **kvmObjectTypes** | **[]string** | Specifies the KVM object types to filter objects. | 
 **awsObjectTypes** | **[]string** | Specifies the AWS object types to filter objects. | 
 **gcpObjectTypes** | **[]string** | Specifies the GCP object types to filter objects. | 
 **acropolisObjectTypes** | **[]string** | Specifies the Acropolis object types to filter objects. | 
 **genericNasObjectTypes** | **[]string** | Specifies the generic NAS object types to filter objects. | 
 **isilonObjectTypes** | **[]string** | Specifies the Isilon object types to filter objects. | 
 **flashbladeObjectTypes** | **[]string** | Specifies the Flashblade object types to filter objects. | 
 **elastifileObjectTypes** | **[]string** | Specifies the Elastifile object types to filter objects. | 
 **gpfsObjectTypes** | **[]string** | Specifies the GPFS object types to filter objects. | 
 **pureObjectTypes** | **[]string** | Specifies the Pure object types to filter objects. | 
 **nimbleObjectTypes** | **[]string** | Specifies the Nimble object types to filter objects. | 
 **physicalObjectTypes** | **[]string** | Specifies the Physical object types to filter objects. | 
 **kubernetesObjectTypes** | **[]string** | Specifies the Kubernetes object types to filter objects. | 
 **exchangeObjectTypes** | **[]string** | Specifies the Exchange object types to filter objects. | 
 **adObjectTypes** | **[]string** | Specifies the AD object types to filter objects. | 
 **mssqlObjectTypes** | **[]string** | Specifies the MSSQL object types to filter objects. | 
 **oracleObjectTypes** | **[]string** | Specifies the Oracle object types to filter objects. | 

### Return type

[**SourceHierarchyObjectSummaries**](SourceHierarchyObjectSummaries.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McmGetTenantObjectIds

> McmTenantObjectIds McmGetTenantObjectIds(ctx).Body(body).RegionId(regionId).Execute()

GetTenantObjectIds



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
    body := *openapiclient.NewMcmTenantObjectIdsParams("TenantId_example", []string{"ObjectHashes_example"}) // McmTenantObjectIdsParams | Specifies the parameters to fetch object IDs.
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

    request := helios.ApiMcmGetTenantObjectIdsRequest{
        Body: &body
        RegionId: &regionId
    }

    resp, r, err := api_client.Object.McmGetTenantObjectIds(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Object.McmGetTenantObjectIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `McmGetTenantObjectIds`: McmTenantObjectIds
    fmt.Fprintf(os.Stdout, "Response from `Object.McmGetTenantObjectIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMcmGetTenantObjectIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**McmTenantObjectIdsParams**](McmTenantObjectIdsParams.md) | Specifies the parameters to fetch object IDs. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**McmTenantObjectIds**](McmTenantObjectIds.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ObjectsActions

> ObjectsActions(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Actions on Objects



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
    body := *openapiclient.NewObjectsActionRequest() // ObjectsActionRequest | Specifies the paramteres to execute actions on given list of objects.
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

    request := helios.ApiObjectsActionsRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Object.ObjectsActions(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Object.ObjectsActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiObjectsActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ObjectsActionRequest**](ObjectsActionRequest.md) | Specifies the paramteres to execute actions on given list of objects. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PerformActionOnObject

> PerformActionOnObject(ctx, id).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Perform an action on an object.



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
    id := int64(789) // int64 | Specifies the id of the Object.
    body := *openapiclient.NewObjectActionRequest("Environment_example") // ObjectActionRequest | Specifies the parameters to perform an action on an object.
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

    request := helios.ApiPerformActionOnObjectRequest{
        Id: &id
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Object.PerformActionOnObject(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Object.PerformActionOnObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the id of the Object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPerformActionOnObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ObjectActionRequest**](ObjectActionRequest.md) | Specifies the parameters to perform an action on an object. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateObjectSnapshot

> ObjectSnapshot UpdateObjectSnapshot(ctx, id, snapshotId).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Update an object snapshot.



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
    id := int64(789) // int64 | Specifies the id of the Object.
    snapshotId := "snapshotId_example" // string | Specifies the id of the snapshot.<br> Note: 1. If the snapshotid of one of the apps is specified, it applies for all the databases in the Protection Run.<br> 2. In case of volume based jobs, please specify the snapshotid of the source not the database. if source snapshot is specified, applied to source snapshot. if database snapshotid is specified in case of volume based jobs, then it is applicable for host's snapshot.
    body := *openapiclient.NewUpdateObjectSnapshotRequest() // UpdateObjectSnapshotRequest | Specifies the parameters update an object snapshot.
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

    request := helios.ApiUpdateObjectSnapshotRequest{
        Id: &id
        SnapshotId: &snapshotId
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Object.UpdateObjectSnapshot(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Object.UpdateObjectSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateObjectSnapshot`: ObjectSnapshot
    fmt.Fprintf(os.Stdout, "Response from `Object.UpdateObjectSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the id of the Object. | 
**snapshotId** | **string** | Specifies the id of the snapshot.&lt;br&gt; Note: 1. If the snapshotid of one of the apps is specified, it applies for all the databases in the Protection Run.&lt;br&gt; 2. In case of volume based jobs, please specify the snapshotid of the source not the database. if source snapshot is specified, applied to source snapshot. if database snapshotid is specified in case of volume based jobs, then it is applicable for host&#39;s snapshot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateObjectSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateObjectSnapshotRequest**](UpdateObjectSnapshotRequest.md) | Specifies the parameters update an object snapshot. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**ObjectSnapshot**](ObjectSnapshot.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

