# \SearchAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchIndexedObjects**](SearchAPI.md#SearchIndexedObjects) | **Post** /data-protect/search/indexed-objects | List indexed objects.
[**SearchObjects**](SearchAPI.md#SearchObjects) | **Get** /data-protect/search/objects | List Objects.
[**SearchProtectedObjects**](SearchAPI.md#SearchProtectedObjects) | **Get** /data-protect/search/protected-objects | List Protected Objects.



## SearchIndexedObjects

> SearchIndexedObjectsResponseBody SearchIndexedObjects(ctx).Body(body).Execute()

List indexed objects.



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
	body := *openapiclient.NewSearchIndexedObjectsRequest("ObjectType_example") // SearchIndexedObjectsRequest | Specifies the parameters to search for indexed objects.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchIndexedObjects(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchIndexedObjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchIndexedObjects`: SearchIndexedObjectsResponseBody
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchIndexedObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchIndexedObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SearchIndexedObjectsRequest**](SearchIndexedObjectsRequest.md) | Specifies the parameters to search for indexed objects. | 

### Return type

[**SearchIndexedObjectsResponseBody**](SearchIndexedObjectsResponseBody.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchObjects

> ObjectsSearchResponseBody SearchObjects(ctx).RequestInitiatorType(requestInitiatorType).SearchString(searchString).Environments(environments).ProtectionTypes(protectionTypes).TenantIds(tenantIds).IncludeTenants(includeTenants).ProtectionGroupIds(protectionGroupIds).ObjectIds(objectIds).OsTypes(osTypes).O365ObjectTypes(o365ObjectTypes).AzureObjectTypes(azureObjectTypes).AwsObjectTypes(awsObjectTypes).AzureUuids(azureUuids).SourceIds(sourceIds).SourceUuids(sourceUuids).IsProtected(isProtected).IsDeleted(isDeleted).LastRunStatusList(lastRunStatusList).RegionIds(regionIds).ClusterIdentifiers(clusterIdentifiers).StorageDomainIds(storageDomainIds).IncludeDeletedObjects(includeDeletedObjects).PaginationCookie(paginationCookie).Count(count).MustHaveTagIds(mustHaveTagIds).MightHaveTagIds(mightHaveTagIds).MustHaveSnapshotTagIds(mustHaveSnapshotTagIds).MightHaveSnapshotTagIds(mightHaveSnapshotTagIds).TagSearchName(tagSearchName).TagNames(tagNames).AnomalyTags(anomalyTags).DataClassificationTags(dataClassificationTags).ThreatTags(threatTags).TagNamesExcluded(tagNamesExcluded).TagTypes(tagTypes).TagCategories(tagCategories).TagSubCategories(tagSubCategories).IncludeHeliosTagInfoForObjects(includeHeliosTagInfoForObjects).ExternalFilters(externalFilters).Execute()

List Objects.



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
	searchString := "searchString_example" // string | Specifies the search string to filter the objects. This search string will be applicable for objectnames. User can specify a wildcard character '*' as a suffix to a string where all object names are matched with the prefix string. For example, if vm1 and vm2 are the names of objects, user can specify vm* to list the objects. If not specified, then all the objects will be returned which will match other filtering criteria. (optional)
	environments := []string{"Environments_example"} // []string | Specifies the environment type to filter objects. (optional)
	protectionTypes := []string{"ProtectionTypes_example"} // []string | Specifies the protection type to filter objects. (optional)
	tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
	includeTenants := true // bool | If true, the response will include Objects which belongs to all tenants which the current user has permission to see. (optional)
	protectionGroupIds := []string{"Inner_example"} // []string | Specifies a list of Protection Group ids to filter the objects. If specified, the objects protected by specified Protection Group ids will be returned. (optional)
	objectIds := []int64{int64(123)} // []int64 | Specifies a list of Object ids to filter. (optional)
	osTypes := []string{"OsTypes_example"} // []string | Specifies the operating system types to filter objects on. (optional)
	o365ObjectTypes := []string{"O365ObjectTypes_example"} // []string | Specifies the object types to filter objects on. Only applicable if the environment is o365. (optional)
	azureObjectTypes := []string{"AzureObjectTypes_example"} // []string | Specifies the object types to filter objects on. Only applicable if the environment is Azure. (optional)
	awsObjectTypes := []string{"AwsObjectTypes_example"} // []string | Specifies the object types to filter objects on. Only applicable if the environment is AWS. (optional)
	azureUuids := []string{"Inner_example"} // []string | Specifies the Azure UUID for the Microsoft365 objects. If specified, the objects with the matching Azure UUIDs will be returned. (optional)
	sourceIds := []int64{int64(123)} // []int64 | Specifies a list of Protection Source object ids to filter the objects. If specified, the object which are present in those Sources will be returned. (optional)
	sourceUuids := []string{"Inner_example"} // []string | Specifies a list of Protection Source object uuids to filter the objects. If specified, the object which are present in those Sources will be returned. (optional)
	isProtected := true // bool | Specifies the protection status of objects. If set to true, only protected objects will be returned. If set to false, only unprotected objects will be returned. If not specified, all objects will be returned. (optional)
	isDeleted := true // bool | If set to true, then objects which are deleted on atleast one cluster will be returned. If not set or set to false then objects which are registered on atleast one cluster are returned. (optional)
	lastRunStatusList := []string{"LastRunStatusList_example"} // []string | Specifies a list of status of the object's last protection run. Only objects with last run status of these will be returned. (optional)
	regionIds := []string{"Inner_example"} // []string | Specifies a list of region ids. Only records from clusters having these region ids will be returned. (optional)
	clusterIdentifiers := []string{"Inner_example"} // []string | Specifies the list of cluster identifiers. Format is clusterId:clusterIncarnationId. Only records from clusters having these identifiers will be returned. (optional)
	storageDomainIds := []string{"Inner_example"} // []string | Specifies the list of storage domain ids. Format is clusterId:clusterIncarnationId:storageDomainId. Only objects having protection in these storage domains will be returned. (optional)
	includeDeletedObjects := true // bool | Specifies whether to include deleted objects in response. These objects can't be protected but can be recovered. This field is deprecated. (optional)
	paginationCookie := "paginationCookie_example" // string | Specifies the pagination cookie with which subsequent parts of the response can be fetched. (optional)
	count := int32(56) // int32 | Specifies the number of objects to be fetched for the specified pagination cookie. (optional)
	mustHaveTagIds := []string{"Inner_example"} // []string | Specifies tags which must be all present in the document. (optional)
	mightHaveTagIds := []string{"Inner_example"} // []string | Specifies list of tags, one or more of which might be present in the document. These are OR'ed together and the resulting criteria AND'ed with the rest of the query. (optional)
	mustHaveSnapshotTagIds := []string{"Inner_example"} // []string | Specifies snapshot tags which must be all present in the document. (optional)
	mightHaveSnapshotTagIds := []string{"Inner_example"} // []string | Specifies list of snapshot tags, one or more of which might be present in the document. These are OR'ed together and the resulting criteria AND'ed with the rest of the query. (optional)
	tagSearchName := "tagSearchName_example" // string | Specifies the tag name to filter the tagged objects and snapshots. User can specify a wildcard character '*' as a suffix to a string where all object's tag names are matched with the prefix string. (optional)
	tagNames := []string{"Inner_example"} // []string | Specifies the tag names to filter the tagged objects and snapshots only for non system tags (optional)
	anomalyTags := []string{"AnomalyTags_example"} // []string | Specifies the Anomaly's tag names to filter the tagged snapshots (optional)
	dataClassificationTags := []string{"DataClassificationTags_example"} // []string | Specifies the Data classification's tag names to filter the tagged snapshots (optional)
	threatTags := []string{"ThreatTags_example"} // []string | Specifies the threat tag's names to filter the tagged snapshots (optional)
	tagNamesExcluded := []string{"Inner_example"} // []string | Specifies the tag names to not include in the tagged snapshots response (optional)
	tagTypes := []string{"TagTypes_example"} // []string | Specifies the tag type to filter the objects and snapshots. (optional)
	tagCategories := []string{"TagCategories_example"} // []string | Specifies the tag category to filter the objects and snapshots. (optional)
	tagSubCategories := []string{"TagSubCategories_example"} // []string | Specifies the tag subcategory to filter the objects and snapshots (optional)
	includeHeliosTagInfoForObjects := true // bool | Specifies whether to include helios tags information for objects in response. Default value is false (optional)
	externalFilters := []string{"Inner_example"} // []string | Specifies the key-value pairs to filtering the results for the search. Each filter is of the form 'key:value'. The filter 'externalFilters:k1:v1&externalFilters:k2:v2&externalFilters:k2:v3' returns the documents where each document will match the query (k1=v1) AND (k2=v2 OR k2 = v3). Allowed keys: - vmBiosUuid - graphUuid - arn - instanceId - bucketName - azureId (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchObjects(context.Background()).RequestInitiatorType(requestInitiatorType).SearchString(searchString).Environments(environments).ProtectionTypes(protectionTypes).TenantIds(tenantIds).IncludeTenants(includeTenants).ProtectionGroupIds(protectionGroupIds).ObjectIds(objectIds).OsTypes(osTypes).O365ObjectTypes(o365ObjectTypes).AzureObjectTypes(azureObjectTypes).AwsObjectTypes(awsObjectTypes).AzureUuids(azureUuids).SourceIds(sourceIds).SourceUuids(sourceUuids).IsProtected(isProtected).IsDeleted(isDeleted).LastRunStatusList(lastRunStatusList).RegionIds(regionIds).ClusterIdentifiers(clusterIdentifiers).StorageDomainIds(storageDomainIds).IncludeDeletedObjects(includeDeletedObjects).PaginationCookie(paginationCookie).Count(count).MustHaveTagIds(mustHaveTagIds).MightHaveTagIds(mightHaveTagIds).MustHaveSnapshotTagIds(mustHaveSnapshotTagIds).MightHaveSnapshotTagIds(mightHaveSnapshotTagIds).TagSearchName(tagSearchName).TagNames(tagNames).AnomalyTags(anomalyTags).DataClassificationTags(dataClassificationTags).ThreatTags(threatTags).TagNamesExcluded(tagNamesExcluded).TagTypes(tagTypes).TagCategories(tagCategories).TagSubCategories(tagSubCategories).IncludeHeliosTagInfoForObjects(includeHeliosTagInfoForObjects).ExternalFilters(externalFilters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchObjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchObjects`: ObjectsSearchResponseBody
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestInitiatorType** | **string** | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | 
 **searchString** | **string** | Specifies the search string to filter the objects. This search string will be applicable for objectnames. User can specify a wildcard character &#39;*&#39; as a suffix to a string where all object names are matched with the prefix string. For example, if vm1 and vm2 are the names of objects, user can specify vm* to list the objects. If not specified, then all the objects will be returned which will match other filtering criteria. | 
 **environments** | **[]string** | Specifies the environment type to filter objects. | 
 **protectionTypes** | **[]string** | Specifies the protection type to filter objects. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **includeTenants** | **bool** | If true, the response will include Objects which belongs to all tenants which the current user has permission to see. | 
 **protectionGroupIds** | **[]string** | Specifies a list of Protection Group ids to filter the objects. If specified, the objects protected by specified Protection Group ids will be returned. | 
 **objectIds** | **[]int64** | Specifies a list of Object ids to filter. | 
 **osTypes** | **[]string** | Specifies the operating system types to filter objects on. | 
 **o365ObjectTypes** | **[]string** | Specifies the object types to filter objects on. Only applicable if the environment is o365. | 
 **azureObjectTypes** | **[]string** | Specifies the object types to filter objects on. Only applicable if the environment is Azure. | 
 **awsObjectTypes** | **[]string** | Specifies the object types to filter objects on. Only applicable if the environment is AWS. | 
 **azureUuids** | **[]string** | Specifies the Azure UUID for the Microsoft365 objects. If specified, the objects with the matching Azure UUIDs will be returned. | 
 **sourceIds** | **[]int64** | Specifies a list of Protection Source object ids to filter the objects. If specified, the object which are present in those Sources will be returned. | 
 **sourceUuids** | **[]string** | Specifies a list of Protection Source object uuids to filter the objects. If specified, the object which are present in those Sources will be returned. | 
 **isProtected** | **bool** | Specifies the protection status of objects. If set to true, only protected objects will be returned. If set to false, only unprotected objects will be returned. If not specified, all objects will be returned. | 
 **isDeleted** | **bool** | If set to true, then objects which are deleted on atleast one cluster will be returned. If not set or set to false then objects which are registered on atleast one cluster are returned. | 
 **lastRunStatusList** | **[]string** | Specifies a list of status of the object&#39;s last protection run. Only objects with last run status of these will be returned. | 
 **regionIds** | **[]string** | Specifies a list of region ids. Only records from clusters having these region ids will be returned. | 
 **clusterIdentifiers** | **[]string** | Specifies the list of cluster identifiers. Format is clusterId:clusterIncarnationId. Only records from clusters having these identifiers will be returned. | 
 **storageDomainIds** | **[]string** | Specifies the list of storage domain ids. Format is clusterId:clusterIncarnationId:storageDomainId. Only objects having protection in these storage domains will be returned. | 
 **includeDeletedObjects** | **bool** | Specifies whether to include deleted objects in response. These objects can&#39;t be protected but can be recovered. This field is deprecated. | 
 **paginationCookie** | **string** | Specifies the pagination cookie with which subsequent parts of the response can be fetched. | 
 **count** | **int32** | Specifies the number of objects to be fetched for the specified pagination cookie. | 
 **mustHaveTagIds** | **[]string** | Specifies tags which must be all present in the document. | 
 **mightHaveTagIds** | **[]string** | Specifies list of tags, one or more of which might be present in the document. These are OR&#39;ed together and the resulting criteria AND&#39;ed with the rest of the query. | 
 **mustHaveSnapshotTagIds** | **[]string** | Specifies snapshot tags which must be all present in the document. | 
 **mightHaveSnapshotTagIds** | **[]string** | Specifies list of snapshot tags, one or more of which might be present in the document. These are OR&#39;ed together and the resulting criteria AND&#39;ed with the rest of the query. | 
 **tagSearchName** | **string** | Specifies the tag name to filter the tagged objects and snapshots. User can specify a wildcard character &#39;*&#39; as a suffix to a string where all object&#39;s tag names are matched with the prefix string. | 
 **tagNames** | **[]string** | Specifies the tag names to filter the tagged objects and snapshots only for non system tags | 
 **anomalyTags** | **[]string** | Specifies the Anomaly&#39;s tag names to filter the tagged snapshots | 
 **dataClassificationTags** | **[]string** | Specifies the Data classification&#39;s tag names to filter the tagged snapshots | 
 **threatTags** | **[]string** | Specifies the threat tag&#39;s names to filter the tagged snapshots | 
 **tagNamesExcluded** | **[]string** | Specifies the tag names to not include in the tagged snapshots response | 
 **tagTypes** | **[]string** | Specifies the tag type to filter the objects and snapshots. | 
 **tagCategories** | **[]string** | Specifies the tag category to filter the objects and snapshots. | 
 **tagSubCategories** | **[]string** | Specifies the tag subcategory to filter the objects and snapshots | 
 **includeHeliosTagInfoForObjects** | **bool** | Specifies whether to include helios tags information for objects in response. Default value is false | 
 **externalFilters** | **[]string** | Specifies the key-value pairs to filtering the results for the search. Each filter is of the form &#39;key:value&#39;. The filter &#39;externalFilters:k1:v1&amp;externalFilters:k2:v2&amp;externalFilters:k2:v3&#39; returns the documents where each document will match the query (k1&#x3D;v1) AND (k2&#x3D;v2 OR k2 &#x3D; v3). Allowed keys: - vmBiosUuid - graphUuid - arn - instanceId - bucketName - azureId | 

### Return type

[**ObjectsSearchResponseBody**](ObjectsSearchResponseBody.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchProtectedObjects

> ProtectedObjectsSearchResponseBody SearchProtectedObjects(ctx).RequestInitiatorType(requestInitiatorType).SearchString(searchString).Environments(environments).SnapshotActions(snapshotActions).ObjectActionKey(objectActionKey).TenantIds(tenantIds).IncludeTenants(includeTenants).ProtectionGroupIds(protectionGroupIds).ObjectIds(objectIds).StorageDomainIds(storageDomainIds).SubResultSize(subResultSize).FilterSnapshotFromUsecs(filterSnapshotFromUsecs).FilterSnapshotToUsecs(filterSnapshotToUsecs).OsTypes(osTypes).SourceIds(sourceIds).RunInstanceIds(runInstanceIds).CdpProtectedOnly(cdpProtectedOnly).RegionIds(regionIds).UseCachedData(useCachedData).Execute()

List Protected Objects.



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
	searchString := "searchString_example" // string | Specifies the search string to filter the objects. This search string will be applicable for objectnames and Protection Group names. User can specify a wildcard character '*' as a suffix to a string where all object and their Protection Group names are matched with the prefix string. For example, if vm1 and vm2 are the names of objects, user can specify vm* to list the objects. If not specified, then all the objects with Protection Groups will be returned which will match other filtering criteria. (optional)
	environments := []string{"Environments_example"} // []string | Specifies the environment type to filter objects. (optional)
	snapshotActions := []string{"SnapshotActions_example"} // []string | Specifies a list of recovery actions. Only snapshots that applies to these actions will be returned. (optional)
	objectActionKey := "objectActionKey_example" // string | Filter by ObjectActionKey, which uniquely represents protection of an object. An object can be protected in multiple ways but atmost once for a given combination of ObjectActionKey. When specified, latest snapshot info matching the objectActionKey is for corresponding object. (optional)
	tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
	includeTenants := true // bool | If true, the response will include Objects which belongs to all tenants which the current user has permission to see. (optional)
	protectionGroupIds := []string{"Inner_example"} // []string | Specifies a list of Protection Group ids to filter the objects. If specified, the objects protected by specified Protection Group ids will be returned. (optional)
	objectIds := []int64{int64(123)} // []int64 | Specifies a list of Object ids to filter. (optional)
	storageDomainIds := []int64{int64(123)} // []int64 | Specifies the Storage Domain ids to filter objects for which Protection Groups are writing data to Cohesity Views on the specified Storage Domains. (optional)
	subResultSize := int32(56) // int32 | Specifies the size of objects to be fetched for a single subresult. (optional)
	filterSnapshotFromUsecs := int64(789) // int64 | Specifies the timestamp in Unix time epoch in microseconds to filter the objects if the Object has a successful snapshot after this value. (optional)
	filterSnapshotToUsecs := int64(789) // int64 | Specifies the timestamp in Unix time epoch in microseconds to filter the objects if the Object has a successful snapshot before this value. (optional)
	osTypes := []string{"OsTypes_example"} // []string | Specifies the operating system types to filter objects on. (optional)
	sourceIds := []int64{int64(123)} // []int64 | Specifies a list of Protection Source object ids to filter the objects. If specified, the object which are present in those Sources will be returned. (optional)
	runInstanceIds := []int64{int64(123)} // []int64 | Specifies a list of run instance ids. If specified only objects belonging to the provided run id will be retunrned. (optional)
	cdpProtectedOnly := true // bool | Specifies whether to only return the CDP protected objects. (optional)
	regionIds := []string{"Inner_example"} // []string | Specifies a list of region ids. Only records from clusters having these region ids will be returned. (optional)
	useCachedData := true // bool | Specifies whether we can serve the GET request to the read replica cache cache. There is a lag of 15 seconds between the read replica and primary data source. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchProtectedObjects(context.Background()).RequestInitiatorType(requestInitiatorType).SearchString(searchString).Environments(environments).SnapshotActions(snapshotActions).ObjectActionKey(objectActionKey).TenantIds(tenantIds).IncludeTenants(includeTenants).ProtectionGroupIds(protectionGroupIds).ObjectIds(objectIds).StorageDomainIds(storageDomainIds).SubResultSize(subResultSize).FilterSnapshotFromUsecs(filterSnapshotFromUsecs).FilterSnapshotToUsecs(filterSnapshotToUsecs).OsTypes(osTypes).SourceIds(sourceIds).RunInstanceIds(runInstanceIds).CdpProtectedOnly(cdpProtectedOnly).RegionIds(regionIds).UseCachedData(useCachedData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchProtectedObjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchProtectedObjects`: ProtectedObjectsSearchResponseBody
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchProtectedObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchProtectedObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestInitiatorType** | **string** | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | 
 **searchString** | **string** | Specifies the search string to filter the objects. This search string will be applicable for objectnames and Protection Group names. User can specify a wildcard character &#39;*&#39; as a suffix to a string where all object and their Protection Group names are matched with the prefix string. For example, if vm1 and vm2 are the names of objects, user can specify vm* to list the objects. If not specified, then all the objects with Protection Groups will be returned which will match other filtering criteria. | 
 **environments** | **[]string** | Specifies the environment type to filter objects. | 
 **snapshotActions** | **[]string** | Specifies a list of recovery actions. Only snapshots that applies to these actions will be returned. | 
 **objectActionKey** | **string** | Filter by ObjectActionKey, which uniquely represents protection of an object. An object can be protected in multiple ways but atmost once for a given combination of ObjectActionKey. When specified, latest snapshot info matching the objectActionKey is for corresponding object. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **includeTenants** | **bool** | If true, the response will include Objects which belongs to all tenants which the current user has permission to see. | 
 **protectionGroupIds** | **[]string** | Specifies a list of Protection Group ids to filter the objects. If specified, the objects protected by specified Protection Group ids will be returned. | 
 **objectIds** | **[]int64** | Specifies a list of Object ids to filter. | 
 **storageDomainIds** | **[]int64** | Specifies the Storage Domain ids to filter objects for which Protection Groups are writing data to Cohesity Views on the specified Storage Domains. | 
 **subResultSize** | **int32** | Specifies the size of objects to be fetched for a single subresult. | 
 **filterSnapshotFromUsecs** | **int64** | Specifies the timestamp in Unix time epoch in microseconds to filter the objects if the Object has a successful snapshot after this value. | 
 **filterSnapshotToUsecs** | **int64** | Specifies the timestamp in Unix time epoch in microseconds to filter the objects if the Object has a successful snapshot before this value. | 
 **osTypes** | **[]string** | Specifies the operating system types to filter objects on. | 
 **sourceIds** | **[]int64** | Specifies a list of Protection Source object ids to filter the objects. If specified, the object which are present in those Sources will be returned. | 
 **runInstanceIds** | **[]int64** | Specifies a list of run instance ids. If specified only objects belonging to the provided run id will be retunrned. | 
 **cdpProtectedOnly** | **bool** | Specifies whether to only return the CDP protected objects. | 
 **regionIds** | **[]string** | Specifies a list of region ids. Only records from clusters having these region ids will be returned. | 
 **useCachedData** | **bool** | Specifies whether we can serve the GET request to the read replica cache cache. There is a lag of 15 seconds between the read replica and primary data source. | 

### Return type

[**ProtectedObjectsSearchResponseBody**](ProtectedObjectsSearchResponseBody.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

