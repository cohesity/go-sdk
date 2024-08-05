# \ObjectAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateEntityMetadata**](ObjectAPI.md#AssociateEntityMetadata) | **Put** /data-protect/objects/metadata | Associate Metadata with Entity
[**BrowseObjectContents**](ObjectAPI.md#BrowseObjectContents) | **Post** /data-protect/objects/{id}/browse | Fetch the contents (files &amp; folders) for the specified object.
[**CancelObjectRuns**](ObjectAPI.md#CancelObjectRuns) | **Post** /data-protect/objects/runs/cancel | Cancel object runs.
[**ConstructMetaInfo**](ObjectAPI.md#ConstructMetaInfo) | **Post** /data-protect/snapshots/{snapshotId}/meta-info | Construct meta info for any workflow from object snapshot and some other information.
[**DeleteEntityMetadata**](ObjectAPI.md#DeleteEntityMetadata) | **Delete** /data-protect/objects/metadata/{id} | Delete Metadata with Entity
[**FilterObjects**](ObjectAPI.md#FilterObjects) | **Post** /data-protect/filter/objects | List all the filtered objects.
[**GetAllIndexedObjectSnapshots**](ObjectAPI.md#GetAllIndexedObjectSnapshots) | **Get** /data-protect/objects/{objectId}/indexed-objects/snapshots | Get snapshots of indexed object.
[**GetEntityMetadata**](ObjectAPI.md#GetEntityMetadata) | **Get** /data-protect/objects/{sourceId}/metadata | Get Metadata of Entities
[**GetIndexedObjectSnapshots**](ObjectAPI.md#GetIndexedObjectSnapshots) | **Get** /data-protect/objects/{objectId}/protection-groups/{protectionGroupId}/indexed-objects/snapshots | Get snapshots of indexed object.
[**GetObjectRunByRunId**](ObjectAPI.md#GetObjectRunByRunId) | **Get** /data-protect/objects/{id}/runs/{runId} | Get a run for an object.
[**GetObjectRuns**](ObjectAPI.md#GetObjectRuns) | **Get** /data-protect/objects/{id}/runs | Get the list of runs for an object.
[**GetObjectSnapshotInfo**](ObjectAPI.md#GetObjectSnapshotInfo) | **Get** /data-protect/snapshots/{snapshotId} | Get details of object snapshot.
[**GetObjectSnapshotVolumeInfo**](ObjectAPI.md#GetObjectSnapshotVolumeInfo) | **Get** /data-protect/snapshots/{snapshotId}/volume | Get volume info of object snapshot.
[**GetObjectSnapshots**](ObjectAPI.md#GetObjectSnapshots) | **Get** /data-protect/objects/{id}/snapshots | List the snapshots for a given object.
[**GetObjectStats**](ObjectAPI.md#GetObjectStats) | **Get** /data-protect/objects/{id}/stats | Get stats for a given object.
[**GetObjectTree**](ObjectAPI.md#GetObjectTree) | **Get** /data-protect/objects/{id}/tree | Get the objects tree hierarchy for for an Object.
[**GetObjectsLastRun**](ObjectAPI.md#GetObjectsLastRun) | **Get** /data-protect/objects/last-run | Get last protection run of objects.
[**GetPITRangesForProtectedObject**](ObjectAPI.md#GetPITRangesForProtectedObject) | **Get** /data-protect/objects/{id}/pit-ranges | Get PIT ranges for an object
[**GetProtectedObjectOfAnyTypeById**](ObjectAPI.md#GetProtectedObjectOfAnyTypeById) | **Get** /data-protect/objects/{id} | Get an Object.
[**GetProtectedObjectsOfAnyType**](ObjectAPI.md#GetProtectedObjectsOfAnyType) | **Get** /data-protect/objects | Get Objects.
[**GetSnapshotDiff**](ObjectAPI.md#GetSnapshotDiff) | **Post** /data-protect/objects/{id}/snapshot-diff | Get diff between two snapshots of a given object.
[**GetSourceHierarchyObjects**](ObjectAPI.md#GetSourceHierarchyObjects) | **Get** /data-protect/sources/{sourceId}/objects | List objects on a source which can be used for data protection.
[**ObjectsActions**](ObjectAPI.md#ObjectsActions) | **Post** /data-protect/objects/actions | Actions on Objects
[**PerformActionOnObject**](ObjectAPI.md#PerformActionOnObject) | **Post** /data-protect/objects/{id}/actions | Perform an action on an object.
[**UpdateObjectSnapshot**](ObjectAPI.md#UpdateObjectSnapshot) | **Put** /data-protect/objects/{id}/snapshots/{snapshotId} | Update an object snapshot.



## AssociateEntityMetadata

> AssociateEntityMetadataResult AssociateEntityMetadata(ctx).Body(body).Execute()

Associate Metadata with Entity



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
	body := *openapiclient.NewAssociateEntityMetadataRequest([]openapiclient.EntityMetadataParams{*openapiclient.NewEntityMetadataParams(int64(123))}, int64(123)) // AssociateEntityMetadataRequest | Specifies the parameters to associate metadata with entities in the entity hierarchy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAPI.AssociateEntityMetadata(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.AssociateEntityMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssociateEntityMetadata`: AssociateEntityMetadataResult
	fmt.Fprintf(os.Stdout, "Response from `ObjectAPI.AssociateEntityMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssociateEntityMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AssociateEntityMetadataRequest**](AssociateEntityMetadataRequest.md) | Specifies the parameters to associate metadata with entities in the entity hierarchy. | 

### Return type

[**AssociateEntityMetadataResult**](AssociateEntityMetadataResult.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrowseObjectContents

> FileFolderInfo BrowseObjectContents(ctx, id).Body(body).Execute()

Fetch the contents (files & folders) for the specified object.



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
	id := int64(789) // int64 | Specifies the id of the Object.
	body := *openapiclient.NewObjectBrowseRequest("Environment_example") // ObjectBrowseRequest | Specifies the parameters to fetch contents of an object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAPI.BrowseObjectContents(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.BrowseObjectContents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BrowseObjectContents`: FileFolderInfo
	fmt.Fprintf(os.Stdout, "Response from `ObjectAPI.BrowseObjectContents`: %v\n", resp)
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

> CancelObjectRunsResults CancelObjectRuns(ctx).Body(body).Execute()

Cancel object runs.



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
	body := *openapiclient.NewCancelObjectRunsRequest() // CancelObjectRunsRequest | Specifies the parameters to cancel object runs.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAPI.CancelObjectRuns(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.CancelObjectRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelObjectRuns`: CancelObjectRunsResults
	fmt.Fprintf(os.Stdout, "Response from `ObjectAPI.CancelObjectRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelObjectRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CancelObjectRunsRequest**](CancelObjectRunsRequest.md) | Specifies the parameters to cancel object runs. | 

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

> ConstructMetaInfoResult ConstructMetaInfo(ctx, snapshotId).Body(body).Execute()

Construct meta info for any workflow from object snapshot and some other information.



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
	snapshotId := "snapshotId_example" // string | Specifies the snapshot id.
	body := *openapiclient.NewConstructMetaInfoRequest("Environment_example") // ConstructMetaInfoRequest | Specifies the parameters to construct meta info for desired workflow.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAPI.ConstructMetaInfo(context.Background(), snapshotId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.ConstructMetaInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConstructMetaInfo`: ConstructMetaInfoResult
	fmt.Fprintf(os.Stdout, "Response from `ObjectAPI.ConstructMetaInfo`: %v\n", resp)
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

 **body** | [**ConstructMetaInfoRequest**](ConstructMetaInfoRequest.md) | Specifies the parameters to construct meta info for desired workflow. | 

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


## DeleteEntityMetadata

> DeleteEntityMetadata(ctx, id).MetadataType(metadataType).EnvironmentType(environmentType).Execute()

Delete Metadata with Entity



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
	id := int64(789) // int64 | Specifies a unique id of the Entity.
	metadataType := "metadataType_example" // string | Specifies the metadata type to be deleted. This is a required field currently and the API will error out if this field is not provided. (optional)
	environmentType := "environmentType_example" // string | Specifies the environment type for the Credentials metadata to be deleted. This will be only set when the metadata type is Credentials. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ObjectAPI.DeleteEntityMetadata(context.Background(), id).MetadataType(metadataType).EnvironmentType(environmentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.DeleteEntityMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the Entity. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEntityMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metadataType** | **string** | Specifies the metadata type to be deleted. This is a required field currently and the API will error out if this field is not provided. | 
 **environmentType** | **string** | Specifies the environment type for the Credentials metadata to be deleted. This will be only set when the metadata type is Credentials. | 

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


## FilterObjects

> FilteredObjectsResponseBody FilterObjects(ctx).Body(body).Execute()

List all the filtered objects.



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
	body := *openapiclient.NewFilterObjectsRequest("FilterType_example", []openapiclient.Filter{*openapiclient.NewFilter()}, []int64{int64(123)}) // FilterObjectsRequest | Specifies the parameters to filter objects.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAPI.FilterObjects(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.FilterObjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FilterObjects`: FilteredObjectsResponseBody
	fmt.Fprintf(os.Stdout, "Response from `ObjectAPI.FilterObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilterObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FilterObjectsRequest**](FilterObjectsRequest.md) | Specifies the parameters to filter objects. | 

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

> GetIndexedObjectSnapshotsResponseBody GetAllIndexedObjectSnapshots(ctx, objectId).IndexedObjectName(indexedObjectName).ProtectionGroupId(protectionGroupId).IncludeIndexedSnapshotsOnly(includeIndexedSnapshotsOnly).FromTimeUsecs(fromTimeUsecs).ToTimeUsecs(toTimeUsecs).RunTypes(runTypes).UseCachedData(useCachedData).ObjectActionKey(objectActionKey).Execute()

Get snapshots of indexed object.



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
	objectId := int64(789) // int64 | Specifies the object id.
	indexedObjectName := "indexedObjectName_example" // string | Specifies the indexed object name.
	protectionGroupId := "protectionGroupId_example" // string | Specifies the protection group id. (optional)
	includeIndexedSnapshotsOnly := true // bool | Specifies whether to only return snapshots which are indexed. In an indexed snapshot files are guaranteed to exist, while in a non-indexed snapshot files may not exist. (optional) (default to false)
	fromTimeUsecs := int64(789) // int64 | Specifies the timestamp in Unix time epoch in microseconds to filter indexed object's snapshots which are taken after this value. (optional)
	toTimeUsecs := int64(789) // int64 | Specifies the timestamp in Unix time epoch in microseconds to filter indexed object's snapshots which are taken before this value. (optional)
	runTypes := []string{"RunTypes_example"} // []string | Filter by run type. Only protection run matching the specified types will be returned. By default, CDP hydration snapshots are not included, unless explicitly queried using this field. (optional)
	useCachedData := true // bool | Specifies whether we can serve the GET request to the read replica cache. There is a lag of 15 seconds between the read replica and primary data source. (optional)
	objectActionKey := "objectActionKey_example" // string | Filter by ObjectActionKey, which uniquely represents backup type for a given version. An object can be protected in multiple ways but atmost once for a given combination of ObjectActionKey and ObjectId. When specified, only versions of given ObjectActionKey are returned for corresponding object id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAPI.GetAllIndexedObjectSnapshots(context.Background(), objectId).IndexedObjectName(indexedObjectName).ProtectionGroupId(protectionGroupId).IncludeIndexedSnapshotsOnly(includeIndexedSnapshotsOnly).FromTimeUsecs(fromTimeUsecs).ToTimeUsecs(toTimeUsecs).RunTypes(runTypes).UseCachedData(useCachedData).ObjectActionKey(objectActionKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.GetAllIndexedObjectSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllIndexedObjectSnapshots`: GetIndexedObjectSnapshotsResponseBody
	fmt.Fprintf(os.Stdout, "Response from `ObjectAPI.GetAllIndexedObjectSnapshots`: %v\n", resp)
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
 **protectionGroupId** | **string** | Specifies the protection group id. | 
 **includeIndexedSnapshotsOnly** | **bool** | Specifies whether to only return snapshots which are indexed. In an indexed snapshot files are guaranteed to exist, while in a non-indexed snapshot files may not exist. | [default to false]
 **fromTimeUsecs** | **int64** | Specifies the timestamp in Unix time epoch in microseconds to filter indexed object&#39;s snapshots which are taken after this value. | 
 **toTimeUsecs** | **int64** | Specifies the timestamp in Unix time epoch in microseconds to filter indexed object&#39;s snapshots which are taken before this value. | 
 **runTypes** | **[]string** | Filter by run type. Only protection run matching the specified types will be returned. By default, CDP hydration snapshots are not included, unless explicitly queried using this field. | 
 **useCachedData** | **bool** | Specifies whether we can serve the GET request to the read replica cache. There is a lag of 15 seconds between the read replica and primary data source. | 
 **objectActionKey** | **string** | Filter by ObjectActionKey, which uniquely represents backup type for a given version. An object can be protected in multiple ways but atmost once for a given combination of ObjectActionKey and ObjectId. When specified, only versions of given ObjectActionKey are returned for corresponding object id. | 

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


## GetEntityMetadata

> GetEntityMetadataResult GetEntityMetadata(ctx, sourceId).EntityIds(entityIds).Execute()

Get Metadata of Entities



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
	sourceId := int64(789) // int64 | Specifies the source ID for which objects should be returned.
	entityIds := []int64{int64(123)} // []int64 | EntityIds contains ids of the entities for which objects are to be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAPI.GetEntityMetadata(context.Background(), sourceId).EntityIds(entityIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.GetEntityMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEntityMetadata`: GetEntityMetadataResult
	fmt.Fprintf(os.Stdout, "Response from `ObjectAPI.GetEntityMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **int64** | Specifies the source ID for which objects should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntityMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entityIds** | **[]int64** | EntityIds contains ids of the entities for which objects are to be returned. | 

### Return type

[**GetEntityMetadataResult**](GetEntityMetadataResult.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndexedObjectSnapshots

> GetIndexedObjectSnapshotsResponseBody GetIndexedObjectSnapshots(ctx, protectionGroupId, objectId).IndexedObjectName(indexedObjectName).IncludeIndexedSnapshotsOnly(includeIndexedSnapshotsOnly).FromTimeUsecs(fromTimeUsecs).ToTimeUsecs(toTimeUsecs).RunTypes(runTypes).UseCachedData(useCachedData).ObjectActionKey(objectActionKey).Execute()

Get snapshots of indexed object.



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
	protectionGroupId := "protectionGroupId_example" // string | Specifies the protection group id.
	objectId := int64(789) // int64 | Specifies the object id.
	indexedObjectName := "indexedObjectName_example" // string | Specifies the indexed object name.
	includeIndexedSnapshotsOnly := true // bool | Specifies whether to only return snapshots which are indexed. In an indexed snapshots file are guaranteed to exist, while in a non-indexed snapshots file may not exist. (optional) (default to false)
	fromTimeUsecs := int64(789) // int64 | Specifies the timestamp in Unix time epoch in microseconds to filter indexed object's snapshots which are taken after this value. (optional)
	toTimeUsecs := int64(789) // int64 | Specifies the timestamp in Unix time epoch in microseconds to filter indexed object's snapshots which are taken before this value. (optional)
	runTypes := []string{"RunTypes_example"} // []string | Filter by run type. Only protection run matching the specified types will be returned. By default, CDP hydration snapshots are not included, unless explicitly queried using this field. (optional)
	useCachedData := true // bool | Specifies whether we can serve the GET request to the read replica cache. There is a lag of 15 seconds between the read replica and primary data source. (optional)
	objectActionKey := "objectActionKey_example" // string | Filter by ObjectActionKey, which uniquely represents backup type for a given version. An object can be protected in multiple ways but atmost once for a given combination of ObjectActionKey and ObjectId. When specified, only versions of given ObjectActionKey are returned for corresponding object id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAPI.GetIndexedObjectSnapshots(context.Background(), protectionGroupId, objectId).IndexedObjectName(indexedObjectName).IncludeIndexedSnapshotsOnly(includeIndexedSnapshotsOnly).FromTimeUsecs(fromTimeUsecs).ToTimeUsecs(toTimeUsecs).RunTypes(runTypes).UseCachedData(useCachedData).ObjectActionKey(objectActionKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.GetIndexedObjectSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIndexedObjectSnapshots`: GetIndexedObjectSnapshotsResponseBody
	fmt.Fprintf(os.Stdout, "Response from `ObjectAPI.GetIndexedObjectSnapshots`: %v\n", resp)
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
 **includeIndexedSnapshotsOnly** | **bool** | Specifies whether to only return snapshots which are indexed. In an indexed snapshots file are guaranteed to exist, while in a non-indexed snapshots file may not exist. | [default to false]
 **fromTimeUsecs** | **int64** | Specifies the timestamp in Unix time epoch in microseconds to filter indexed object&#39;s snapshots which are taken after this value. | 
 **toTimeUsecs** | **int64** | Specifies the timestamp in Unix time epoch in microseconds to filter indexed object&#39;s snapshots which are taken before this value. | 
 **runTypes** | **[]string** | Filter by run type. Only protection run matching the specified types will be returned. By default, CDP hydration snapshots are not included, unless explicitly queried using this field. | 
 **useCachedData** | **bool** | Specifies whether we can serve the GET request to the read replica cache. There is a lag of 15 seconds between the read replica and primary data source. | 
 **objectActionKey** | **string** | Filter by ObjectActionKey, which uniquely represents backup type for a given version. An object can be protected in multiple ways but atmost once for a given combination of ObjectActionKey and ObjectId. When specified, only versions of given ObjectActionKey are returned for corresponding object id. | 

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


## GetObjectRunByRunId

> ObjectProtectionRunSummary GetObjectRunByRunId(ctx, id, runId).Execute()

Get a run for an object.



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
	id := int64(789) // int64 | Specifies a unique id of the object.
	runId := "runId_example" // string | Specifies the id of the run.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAPI.GetObjectRunByRunId(context.Background(), id, runId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.GetObjectRunByRunId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectRunByRunId`: ObjectProtectionRunSummary
	fmt.Fprintf(os.Stdout, "Response from `ObjectAPI.GetObjectRunByRunId`: %v\n", resp)
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

> GetObjectRunsResponseBody GetObjectRuns(ctx, id).RunId(runId).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).TenantIds(tenantIds).IncludeTenants(includeTenants).RunTypes(runTypes).LocalBackupObjectStatus(localBackupObjectStatus).ReplicationObjectStatus(replicationObjectStatus).ArchivalObjectStatus(archivalObjectStatus).CloudSpinRunStatus(cloudSpinRunStatus).NumRuns(numRuns).PaginationCookie(paginationCookie).ExcludeNonRestorableRuns(excludeNonRestorableRuns).Execute()

Get the list of runs for an object.



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
	id := int64(789) // int64 | Specifies a unique id of the object.
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
	excludeNonRestorableRuns := true // bool | Specifies whether to exclude non restorable runs. Run is treated restorable only if there is at least one object snapshot (which may be either a local or an archival snapshot) which is not deleted or expired. Default value is false. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAPI.GetObjectRuns(context.Background(), id).RunId(runId).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).TenantIds(tenantIds).IncludeTenants(includeTenants).RunTypes(runTypes).LocalBackupObjectStatus(localBackupObjectStatus).ReplicationObjectStatus(replicationObjectStatus).ArchivalObjectStatus(archivalObjectStatus).CloudSpinRunStatus(cloudSpinRunStatus).NumRuns(numRuns).PaginationCookie(paginationCookie).ExcludeNonRestorableRuns(excludeNonRestorableRuns).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.GetObjectRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectRuns`: GetObjectRunsResponseBody
	fmt.Fprintf(os.Stdout, "Response from `ObjectAPI.GetObjectRuns`: %v\n", resp)
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
 **excludeNonRestorableRuns** | **bool** | Specifies whether to exclude non restorable runs. Run is treated restorable only if there is at least one object snapshot (which may be either a local or an archival snapshot) which is not deleted or expired. Default value is false. | [default to false]

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

> ObjectSnapshot GetObjectSnapshotInfo(ctx, snapshotId).Execute()

Get details of object snapshot.



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
	snapshotId := "snapshotId_example" // string | Specifies the snapshot id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAPI.GetObjectSnapshotInfo(context.Background(), snapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.GetObjectSnapshotInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectSnapshotInfo`: ObjectSnapshot
	fmt.Fprintf(os.Stdout, "Response from `ObjectAPI.GetObjectSnapshotInfo`: %v\n", resp)
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

> ObjectSnapshotVolumeInfo GetObjectSnapshotVolumeInfo(ctx, snapshotId).IncludeSupportedOnly(includeSupportedOnly).PointInTimeUsecs(pointInTimeUsecs).UseCachedData(useCachedData).Execute()

Get volume info of object snapshot.



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
	snapshotId := "snapshotId_example" // string | Specifies the snapshot id.
	includeSupportedOnly := true // bool | Specifies whether to only return supported volumes. (optional)
	pointInTimeUsecs := float32(8.14) // float32 | Specifies the point-in-time timestamp (in usecs from epoch) between snapshots for which the volume info is to be returned. (optional)
	useCachedData := true // bool | Specifies whether we can serve the GET request to the read replica cache. There is a lag of 15 seconds between the read replica and primary data source. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAPI.GetObjectSnapshotVolumeInfo(context.Background(), snapshotId).IncludeSupportedOnly(includeSupportedOnly).PointInTimeUsecs(pointInTimeUsecs).UseCachedData(useCachedData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.GetObjectSnapshotVolumeInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectSnapshotVolumeInfo`: ObjectSnapshotVolumeInfo
	fmt.Fprintf(os.Stdout, "Response from `ObjectAPI.GetObjectSnapshotVolumeInfo`: %v\n", resp)
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

 **includeSupportedOnly** | **bool** | Specifies whether to only return supported volumes. | 
 **pointInTimeUsecs** | **float32** | Specifies the point-in-time timestamp (in usecs from epoch) between snapshots for which the volume info is to be returned. | 
 **useCachedData** | **bool** | Specifies whether we can serve the GET request to the read replica cache. There is a lag of 15 seconds between the read replica and primary data source. | 

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

> GetObjectSnapshotsResponseBody GetObjectSnapshots(ctx, id).FromTimeUsecs(fromTimeUsecs).ToTimeUsecs(toTimeUsecs).RunStartFromTimeUsecs(runStartFromTimeUsecs).RunStartToTimeUsecs(runStartToTimeUsecs).SnapshotActions(snapshotActions).RunTypes(runTypes).ProtectionGroupIds(protectionGroupIds).RunInstanceIds(runInstanceIds).RegionIds(regionIds).ObjectActionKeys(objectActionKeys).Execute()

List the snapshots for a given object.



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
	id := int64(789) // int64 | Specifies the id of the Object.
	fromTimeUsecs := int64(789) // int64 | Specifies the timestamp in Unix time epoch in microseconds to filter Object's snapshots which were taken after this value. (optional)
	toTimeUsecs := int64(789) // int64 | Specifies the timestamp in Unix time epoch in microseconds to filter Object's snapshots which were taken before this value. (optional)
	runStartFromTimeUsecs := int64(789) // int64 | Specifies the timestamp in Unix time epoch in microseconds to filter Object's snapshots which were run after this value. (optional)
	runStartToTimeUsecs := int64(789) // int64 | Specifies the timestamp in Unix time epoch in microseconds to filter Object's snapshots which were run before this value. (optional)
	snapshotActions := []string{"SnapshotActions_example"} // []string | Specifies a list of recovery actions. Only snapshots that applies to these actions will be returned. (optional)
	runTypes := []string{"RunTypes_example"} // []string | Filter by run type. Only protection run matching the specified types will be returned. By default, CDP hydration snapshots are not included, unless explicitly queried using this field. (optional)
	protectionGroupIds := []string{"Inner_example"} // []string | If specified, this returns only the snapshots of the specified object ID, which belong to the provided protection group IDs. (optional)
	runInstanceIds := []int64{int64(123)} // []int64 | Filter by a list run instance ids. If specified, only snapshots created by these protection runs will be returned. (optional)
	regionIds := []string{"Inner_example"} // []string | Filter by a list of region ids. (optional)
	objectActionKeys := []string{"ObjectActionKeys_example"} // []string | Filter by ObjectActionKey, which uniquely represents protection of an object. An object can be protected in multiple ways but atmost once for a given combination of ObjectActionKey. When specified, only snapshots matching given action keys are returned for corresponding object. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAPI.GetObjectSnapshots(context.Background(), id).FromTimeUsecs(fromTimeUsecs).ToTimeUsecs(toTimeUsecs).RunStartFromTimeUsecs(runStartFromTimeUsecs).RunStartToTimeUsecs(runStartToTimeUsecs).SnapshotActions(snapshotActions).RunTypes(runTypes).ProtectionGroupIds(protectionGroupIds).RunInstanceIds(runInstanceIds).RegionIds(regionIds).ObjectActionKeys(objectActionKeys).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.GetObjectSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectSnapshots`: GetObjectSnapshotsResponseBody
	fmt.Fprintf(os.Stdout, "Response from `ObjectAPI.GetObjectSnapshots`: %v\n", resp)
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

 **fromTimeUsecs** | **int64** | Specifies the timestamp in Unix time epoch in microseconds to filter Object&#39;s snapshots which were taken after this value. | 
 **toTimeUsecs** | **int64** | Specifies the timestamp in Unix time epoch in microseconds to filter Object&#39;s snapshots which were taken before this value. | 
 **runStartFromTimeUsecs** | **int64** | Specifies the timestamp in Unix time epoch in microseconds to filter Object&#39;s snapshots which were run after this value. | 
 **runStartToTimeUsecs** | **int64** | Specifies the timestamp in Unix time epoch in microseconds to filter Object&#39;s snapshots which were run before this value. | 
 **snapshotActions** | **[]string** | Specifies a list of recovery actions. Only snapshots that applies to these actions will be returned. | 
 **runTypes** | **[]string** | Filter by run type. Only protection run matching the specified types will be returned. By default, CDP hydration snapshots are not included, unless explicitly queried using this field. | 
 **protectionGroupIds** | **[]string** | If specified, this returns only the snapshots of the specified object ID, which belong to the provided protection group IDs. | 
 **runInstanceIds** | **[]int64** | Filter by a list run instance ids. If specified, only snapshots created by these protection runs will be returned. | 
 **regionIds** | **[]string** | Filter by a list of region ids. | 
 **objectActionKeys** | **[]string** | Filter by ObjectActionKey, which uniquely represents protection of an object. An object can be protected in multiple ways but atmost once for a given combination of ObjectActionKey. When specified, only snapshots matching given action keys are returned for corresponding object. | 

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

> ObjectStats GetObjectStats(ctx, id).RegionIds(regionIds).Execute()

Get stats for a given object.



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
	id := int64(789) // int64 | Specifies the id of the Object.
	regionIds := []string{"Inner_example"} // []string | Filter by a list of region ids. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAPI.GetObjectStats(context.Background(), id).RegionIds(regionIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.GetObjectStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectStats`: ObjectStats
	fmt.Fprintf(os.Stdout, "Response from `ObjectAPI.GetObjectStats`: %v\n", resp)
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

> ObjectWithChildren GetObjectTree(ctx, id).Execute()

Get the objects tree hierarchy for for an Object.



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
	id := int64(789) // int64 | Specifies the id of the Object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAPI.GetObjectTree(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.GetObjectTree``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectTree`: ObjectWithChildren
	fmt.Fprintf(os.Stdout, "Response from `ObjectAPI.GetObjectTree`: %v\n", resp)
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

> ObjectsLastRun GetObjectsLastRun(ctx).Ids(ids).TenantIds(tenantIds).IncludeTenants(includeTenants).PaginationCookie(paginationCookie).Count(count).Execute()

Get last protection run of objects.



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
	ids := []int64{int64(123)} // []int64 | Specifies a list of object ids, only last runs for these objects will be returned. (optional)
	tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
	includeTenants := true // bool | If true, the response will include Objects which belongs to all tenants which the current user has permission to see. (optional)
	paginationCookie := "paginationCookie_example" // string | Specifies the pagination cookie with which subsequent parts of the response can be fetched. (optional)
	count := int32(56) // int32 | Specifies the number of objects to be fetched for the specified pagination cookie. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAPI.GetObjectsLastRun(context.Background()).Ids(ids).TenantIds(tenantIds).IncludeTenants(includeTenants).PaginationCookie(paginationCookie).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.GetObjectsLastRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectsLastRun`: ObjectsLastRun
	fmt.Fprintf(os.Stdout, "Response from `ObjectAPI.GetObjectsLastRun`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectsLastRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

> GetPITRangesProtectedObjectResponseBody GetPITRangesForProtectedObject(ctx, id).FromTimeUsecs(fromTimeUsecs).ToTimeUsecs(toTimeUsecs).ProtectionGroupIds(protectionGroupIds).Execute()

Get PIT ranges for an object



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
	id := int64(789) // int64 | Specifies the ID of the protected object.
	fromTimeUsecs := int64(789) // int64 | If specified, return the restore ranges that lie after this timestamp. This parameter is specified as the timestamp in Unix time epoch in microseconds. (optional)
	toTimeUsecs := int64(789) // int64 | If specified, return the restore ranges that lie before this timestamp. This parameter is specified as the timestamp in Unix time epoch in microseconds. (optional)
	protectionGroupIds := []string{"Inner_example"} // []string | If specified, return only the points in time corresponding to these protection group IDs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAPI.GetPITRangesForProtectedObject(context.Background(), id).FromTimeUsecs(fromTimeUsecs).ToTimeUsecs(toTimeUsecs).ProtectionGroupIds(protectionGroupIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.GetPITRangesForProtectedObject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPITRangesForProtectedObject`: GetPITRangesProtectedObjectResponseBody
	fmt.Fprintf(os.Stdout, "Response from `ObjectAPI.GetPITRangesForProtectedObject`: %v\n", resp)
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

> ProtectedObjectInfo GetProtectedObjectOfAnyTypeById(ctx, id).RequestInitiatorType(requestInitiatorType).ObjectActionKey(objectActionKey).OnlyProtectedObjects(onlyProtectedObjects).StorageDomainId(storageDomainId).Environments(environments).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeLastRunInfo(includeLastRunInfo).OnlyAutoProtectedObjects(onlyAutoProtectedObjects).OnlyLeafObjects(onlyLeafObjects).Execute()

Get an Object.



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
	id := int64(789) // int64 | Specifies the id of the Object.
	requestInitiatorType := "requestInitiatorType_example" // string | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. (optional)
	objectActionKey := []string{"ObjectActionKey_example"} // []string | Filter by ObjectActionKey, uniquely represent protection of an object. An object can be protected in multiple ways but atmost once for a given combination of ObjectActionKey, when specified Only objects of given action_key are returned for corresponding object id and this vec's size needs to be same as 'id'. (optional)
	onlyProtectedObjects := true // bool | If true, the response will include only objects which have been protected. (optional)
	storageDomainId := int64(789) // int64 | Filter by Storage Domain id. Only Objects protected to this Storage Domain will be returned. (optional)
	environments := []string{"Environments_example"} // []string | Filter by environment types such as 'kVMware', 'kView', etc. Only Protected objects protecting the specified environment types are returned. (optional)
	tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
	includeTenants := true // bool | If true, the response will include Objects which were protected by all tenants which the current user has permission to see. If false, then only objects protected by the current user will be returned. (optional)
	includeLastRunInfo := true // bool | If true, the response will include information about the last protection run on this object. (optional)
	onlyAutoProtectedObjects := true // bool | If true, the response will include only the auto protected objects. (optional)
	onlyLeafObjects := true // bool | If true, the response will include only the leaf level objects. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAPI.GetProtectedObjectOfAnyTypeById(context.Background(), id).RequestInitiatorType(requestInitiatorType).ObjectActionKey(objectActionKey).OnlyProtectedObjects(onlyProtectedObjects).StorageDomainId(storageDomainId).Environments(environments).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeLastRunInfo(includeLastRunInfo).OnlyAutoProtectedObjects(onlyAutoProtectedObjects).OnlyLeafObjects(onlyLeafObjects).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.GetProtectedObjectOfAnyTypeById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProtectedObjectOfAnyTypeById`: ProtectedObjectInfo
	fmt.Fprintf(os.Stdout, "Response from `ObjectAPI.GetProtectedObjectOfAnyTypeById`: %v\n", resp)
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

 **requestInitiatorType** | **string** | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | 
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

> GetProtectedObjectsResponse GetProtectedObjectsOfAnyType(ctx).RequestInitiatorType(requestInitiatorType).Ids(ids).ObjectActionKeys(objectActionKeys).PolicyIds(policyIds).ParentId(parentId).OnlyProtectedObjects(onlyProtectedObjects).StorageDomainId(storageDomainId).Environments(environments).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeLastRunInfo(includeLastRunInfo).OnlyAutoProtectedObjects(onlyAutoProtectedObjects).OnlyLeafObjects(onlyLeafObjects).RegionIds(regionIds).MaxCount(maxCount).Cookie(cookie).Execute()

Get Objects.



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
	requestInitiatorType := "requestInitiatorType_example" // string | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. (optional)
	ids := []int64{int64(123)} // []int64 | Filter by a list of Object ids. (optional)
	objectActionKeys := []string{"ObjectActionKeys_example"} // []string | Filter by ObjectActionKey, uniquely represent protection of an object. An object can be protected in multiple ways but atmost once for a given combination of ObjectActionKey, when specified Only objects of given action_key are returned for corresponding object id. The vec's size needs to be of either length one or same as the length of 'ids'. If the length of objectActionKey is one, it will be repeated as many number of times equal to the length of objectIds, as mandated by backend validation. If the length of objectActionKey and object ids are same then it will be passed as it is. (optional)
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
	maxCount := int32(56) // int32 | Specifies the max number of objects to return. (optional)
	cookie := "cookie_example" // string | Specifies the pagination cookie. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAPI.GetProtectedObjectsOfAnyType(context.Background()).RequestInitiatorType(requestInitiatorType).Ids(ids).ObjectActionKeys(objectActionKeys).PolicyIds(policyIds).ParentId(parentId).OnlyProtectedObjects(onlyProtectedObjects).StorageDomainId(storageDomainId).Environments(environments).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeLastRunInfo(includeLastRunInfo).OnlyAutoProtectedObjects(onlyAutoProtectedObjects).OnlyLeafObjects(onlyLeafObjects).RegionIds(regionIds).MaxCount(maxCount).Cookie(cookie).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.GetProtectedObjectsOfAnyType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProtectedObjectsOfAnyType`: GetProtectedObjectsResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectAPI.GetProtectedObjectsOfAnyType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectedObjectsOfAnyTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestInitiatorType** | **string** | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | 
 **ids** | **[]int64** | Filter by a list of Object ids. | 
 **objectActionKeys** | **[]string** | Filter by ObjectActionKey, uniquely represent protection of an object. An object can be protected in multiple ways but atmost once for a given combination of ObjectActionKey, when specified Only objects of given action_key are returned for corresponding object id. The vec&#39;s size needs to be of either length one or same as the length of &#39;ids&#39;. If the length of objectActionKey is one, it will be repeated as many number of times equal to the length of objectIds, as mandated by backend validation. If the length of objectActionKey and object ids are same then it will be passed as it is. | 
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
 **maxCount** | **int32** | Specifies the max number of objects to return. | 
 **cookie** | **string** | Specifies the pagination cookie. | 

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

> SnapshotDiffResult GetSnapshotDiff(ctx, id).Body(body).Execute()

Get diff between two snapshots of a given object.



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
	id := int64(789) // int64 | 
	body := *openapiclient.NewSnapshotDiffParams(int64(123), int64(123), int64(123), "EntityType_example", int64(123), int64(123), int64(123), int64(123), int64(123), int64(123)) // SnapshotDiffParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAPI.GetSnapshotDiff(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.GetSnapshotDiff``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnapshotDiff`: SnapshotDiffResult
	fmt.Fprintf(os.Stdout, "Response from `ObjectAPI.GetSnapshotDiff`: %v\n", resp)
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

> SourceHierarchyObjectSummaries GetSourceHierarchyObjects(ctx, sourceId).ParentId(parentId).TenantIds(tenantIds).IncludeTenants(includeTenants).VmwareObjectTypes(vmwareObjectTypes).NetappObjectTypes(netappObjectTypes).O365ObjectTypes(o365ObjectTypes).CassandraObjectTypes(cassandraObjectTypes).MongodbObjectTypes(mongodbObjectTypes).CouchbaseObjectTypes(couchbaseObjectTypes).HdfsObjectTypes(hdfsObjectTypes).HbaseObjectTypes(hbaseObjectTypes).HiveObjectTypes(hiveObjectTypes).HypervObjectTypes(hypervObjectTypes).AzureObjectTypes(azureObjectTypes).KvmObjectTypes(kvmObjectTypes).AwsObjectTypes(awsObjectTypes).GcpObjectTypes(gcpObjectTypes).AcropolisObjectTypes(acropolisObjectTypes).GenericNasObjectTypes(genericNasObjectTypes).IsilonObjectTypes(isilonObjectTypes).FlashbladeObjectTypes(flashbladeObjectTypes).ElastifileObjectTypes(elastifileObjectTypes).GpfsObjectTypes(gpfsObjectTypes).PureObjectTypes(pureObjectTypes).NimbleObjectTypes(nimbleObjectTypes).PhysicalObjectTypes(physicalObjectTypes).KubernetesObjectTypes(kubernetesObjectTypes).ExchangeObjectTypes(exchangeObjectTypes).AdObjectTypes(adObjectTypes).MssqlObjectTypes(mssqlObjectTypes).OracleObjectTypes(oracleObjectTypes).UseCachedData(useCachedData).Execute()

List objects on a source which can be used for data protection.



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
	sourceId := int64(789) // int64 | Specifies the source ID for which objects should be returned.
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
	useCachedData := true // bool | Specifies whether we can serve the GET request to the read replica cache. There is a lag of 15 seconds between the read replica and primary data source. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAPI.GetSourceHierarchyObjects(context.Background(), sourceId).ParentId(parentId).TenantIds(tenantIds).IncludeTenants(includeTenants).VmwareObjectTypes(vmwareObjectTypes).NetappObjectTypes(netappObjectTypes).O365ObjectTypes(o365ObjectTypes).CassandraObjectTypes(cassandraObjectTypes).MongodbObjectTypes(mongodbObjectTypes).CouchbaseObjectTypes(couchbaseObjectTypes).HdfsObjectTypes(hdfsObjectTypes).HbaseObjectTypes(hbaseObjectTypes).HiveObjectTypes(hiveObjectTypes).HypervObjectTypes(hypervObjectTypes).AzureObjectTypes(azureObjectTypes).KvmObjectTypes(kvmObjectTypes).AwsObjectTypes(awsObjectTypes).GcpObjectTypes(gcpObjectTypes).AcropolisObjectTypes(acropolisObjectTypes).GenericNasObjectTypes(genericNasObjectTypes).IsilonObjectTypes(isilonObjectTypes).FlashbladeObjectTypes(flashbladeObjectTypes).ElastifileObjectTypes(elastifileObjectTypes).GpfsObjectTypes(gpfsObjectTypes).PureObjectTypes(pureObjectTypes).NimbleObjectTypes(nimbleObjectTypes).PhysicalObjectTypes(physicalObjectTypes).KubernetesObjectTypes(kubernetesObjectTypes).ExchangeObjectTypes(exchangeObjectTypes).AdObjectTypes(adObjectTypes).MssqlObjectTypes(mssqlObjectTypes).OracleObjectTypes(oracleObjectTypes).UseCachedData(useCachedData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.GetSourceHierarchyObjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSourceHierarchyObjects`: SourceHierarchyObjectSummaries
	fmt.Fprintf(os.Stdout, "Response from `ObjectAPI.GetSourceHierarchyObjects`: %v\n", resp)
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
 **useCachedData** | **bool** | Specifies whether we can serve the GET request to the read replica cache. There is a lag of 15 seconds between the read replica and primary data source. | 

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


## ObjectsActions

> ObjectsActions(ctx).Body(body).Execute()

Actions on Objects



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
	body := *openapiclient.NewObjectsActionRequest() // ObjectsActionRequest | Specifies the parameters to execute actions on given list of objects.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ObjectAPI.ObjectsActions(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.ObjectsActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiObjectsActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ObjectsActionRequest**](ObjectsActionRequest.md) | Specifies the parameters to execute actions on given list of objects. | 

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

> PerformActionOnObject(ctx, id).Body(body).Execute()

Perform an action on an object.



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
	id := int64(789) // int64 | Specifies the id of the Object.
	body := *openapiclient.NewObjectActionRequest("Environment_example") // ObjectActionRequest | Specifies the parameters to perform an action on an object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ObjectAPI.PerformActionOnObject(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.PerformActionOnObject``: %v\n", err)
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

> ObjectSnapshot UpdateObjectSnapshot(ctx, id, snapshotId).Body(body).Execute()

Update an object snapshot.



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
	id := int64(789) // int64 | Specifies the id of the Object.
	snapshotId := "snapshotId_example" // string | Specifies the id of the snapshot.<br> Note: 1. If the snapshotid of one of the apps is specified, it applies for all the databases in the Protection Run.<br> 2. In case of volume based jobs, please specify the snapshotid of the source not the database. if source snapshot is specified, applied to source snapshot. if database snapshotid is specified in case of volume based jobs, then it is applicable for host's snapshot.
	body := *openapiclient.NewUpdateObjectSnapshotRequest() // UpdateObjectSnapshotRequest | Specifies the parameters update an object snapshot.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAPI.UpdateObjectSnapshot(context.Background(), id, snapshotId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.UpdateObjectSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateObjectSnapshot`: ObjectSnapshot
	fmt.Fprintf(os.Stdout, "Response from `ObjectAPI.UpdateObjectSnapshot`: %v\n", resp)
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

