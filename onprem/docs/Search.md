# \Search

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchIndexedObjects**](Search.md#SearchIndexedObjects) | **Post** /data-protect/search/indexed-objects | List indexed objects.
[**SearchObjects**](Search.md#SearchObjects) | **Get** /data-protect/search/objects | List Objects.
[**SearchProtectedObjects**](Search.md#SearchProtectedObjects) | **Get** /data-protect/search/protected-objects | List Protected Objects.



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
    onprem "onprem"
)

func main() {
    body := *openapiclient.NewSearchIndexedObjectsRequest("ObjectType_example") // SearchIndexedObjectsRequest | Specifies the parameters to search for indexed objects.

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

    request := onprem.ApiSearchIndexedObjectsRequest{
        Body: &body
    }

    resp, r, err := api_client.Search.SearchIndexedObjects(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Search.SearchIndexedObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchIndexedObjects`: SearchIndexedObjectsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `Search.SearchIndexedObjects`: %v\n", resp)
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

> ObjectsSearchResponseBody SearchObjects(ctx).SearchString(searchString).Environments(environments).ProtectionTypes(protectionTypes).TenantIds(tenantIds).IncludeTenants(includeTenants).ProtectionGroupIds(protectionGroupIds).ObjectIds(objectIds).OsTypes(osTypes).SourceIds(sourceIds).SourceUuids(sourceUuids).IsProtected(isProtected).IsDeleted(isDeleted).LastRunStatusList(lastRunStatusList).RegionIds(regionIds).ClusterIdentifiers(clusterIdentifiers).StorageDomainIds(storageDomainIds).IncludeDeletedObjects(includeDeletedObjects).PaginationCookie(paginationCookie).Count(count).MustHaveTagIds(mustHaveTagIds).MightHaveTagIds(mightHaveTagIds).MustHaveSnapshotTagIds(mustHaveSnapshotTagIds).MightHaveSnapshotTagIds(mightHaveSnapshotTagIds).Execute()

List Objects.



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
    searchString := "searchString_example" // string | Specifies the search string to filter the objects. This search string will be applicable for objectnames. User can specify a wildcard character '*' as a suffix to a string where all object names are matched with the prefix string. For example, if vm1 and vm2 are the names of objects, user can specify vm* to list the objects. If not specified, then all the objects will be returned which will match other filtering criteria. (optional)
    environments := []string{"Environments_example"} // []string | Specifies the environment type to filter objects. (optional)
    protectionTypes := []string{"ProtectionTypes_example"} // []string | Specifies the protection type to filter objects. (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
    includeTenants := true // bool | If true, the response will include Objects which belongs to all tenants which the current user has permission to see. (optional)
    protectionGroupIds := []string{"Inner_example"} // []string | Specifies a list of Protection Group ids to filter the objects. If specified, the objects protected by specified Protection Group ids will be returned. (optional)
    objectIds := []int64{int64(123)} // []int64 | Specifies a list of Object ids to filter. (optional)
    osTypes := []string{"OsTypes_example"} // []string | Specifies the operating system types to filter objects on. (optional)
    sourceIds := []int64{int64(123)} // []int64 | Specifies a list of Protection Source object ids to filter the objects. If specified, the object which are present in those Sources will be returned. (optional)
    sourceUuids := []string{"Inner_example"} // []string | Specifies a list of Protection Source object uuids to filter the objects. If specified, the object which are present in those Sources will be returned. (optional)
    isProtected := true // bool | Specifies the protection status of objects. If set to true, only protected objects will be returned. If set to false, only unprotected objects will be returned. If not specified, all objects will be returned. (optional)
    isDeleted := true // bool | Specifies the deletion status of objects. If set to true, only deleted objects will be returned. If set to false, only not deleted objects will be returned. If not specified, all objects will be returned. (optional)
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

    request := onprem.ApiSearchObjectsRequest{
        SearchString: &searchString
        Environments: &environments
        ProtectionTypes: &protectionTypes
        TenantIds: &tenantIds
        IncludeTenants: &includeTenants
        ProtectionGroupIds: &protectionGroupIds
        ObjectIds: &objectIds
        OsTypes: &osTypes
        SourceIds: &sourceIds
        SourceUuids: &sourceUuids
        IsProtected: &isProtected
        IsDeleted: &isDeleted
        LastRunStatusList: &lastRunStatusList
        RegionIds: &regionIds
        ClusterIdentifiers: &clusterIdentifiers
        StorageDomainIds: &storageDomainIds
        IncludeDeletedObjects: &includeDeletedObjects
        PaginationCookie: &paginationCookie
        Count: &count
        MustHaveTagIds: &mustHaveTagIds
        MightHaveTagIds: &mightHaveTagIds
        MustHaveSnapshotTagIds: &mustHaveSnapshotTagIds
        MightHaveSnapshotTagIds: &mightHaveSnapshotTagIds
    }

    resp, r, err := api_client.Search.SearchObjects(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Search.SearchObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchObjects`: ObjectsSearchResponseBody
    fmt.Fprintf(os.Stdout, "Response from `Search.SearchObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchString** | **string** | Specifies the search string to filter the objects. This search string will be applicable for objectnames. User can specify a wildcard character &#39;*&#39; as a suffix to a string where all object names are matched with the prefix string. For example, if vm1 and vm2 are the names of objects, user can specify vm* to list the objects. If not specified, then all the objects will be returned which will match other filtering criteria. | 
 **environments** | **[]string** | Specifies the environment type to filter objects. | 
 **protectionTypes** | **[]string** | Specifies the protection type to filter objects. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **includeTenants** | **bool** | If true, the response will include Objects which belongs to all tenants which the current user has permission to see. | 
 **protectionGroupIds** | **[]string** | Specifies a list of Protection Group ids to filter the objects. If specified, the objects protected by specified Protection Group ids will be returned. | 
 **objectIds** | **[]int64** | Specifies a list of Object ids to filter. | 
 **osTypes** | **[]string** | Specifies the operating system types to filter objects on. | 
 **sourceIds** | **[]int64** | Specifies a list of Protection Source object ids to filter the objects. If specified, the object which are present in those Sources will be returned. | 
 **sourceUuids** | **[]string** | Specifies a list of Protection Source object uuids to filter the objects. If specified, the object which are present in those Sources will be returned. | 
 **isProtected** | **bool** | Specifies the protection status of objects. If set to true, only protected objects will be returned. If set to false, only unprotected objects will be returned. If not specified, all objects will be returned. | 
 **isDeleted** | **bool** | Specifies the deletion status of objects. If set to true, only deleted objects will be returned. If set to false, only not deleted objects will be returned. If not specified, all objects will be returned. | 
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

> ProtectedObjectsSearchResponseBody SearchProtectedObjects(ctx).SearchString(searchString).Environments(environments).SnapshotActions(snapshotActions).TenantIds(tenantIds).IncludeTenants(includeTenants).ProtectionGroupIds(protectionGroupIds).ObjectIds(objectIds).StorageDomainIds(storageDomainIds).SubResultSize(subResultSize).FilterSnapshotFromUsecs(filterSnapshotFromUsecs).FilterSnapshotToUsecs(filterSnapshotToUsecs).OsTypes(osTypes).SourceIds(sourceIds).RunInstanceIds(runInstanceIds).CdpProtectedOnly(cdpProtectedOnly).RegionIds(regionIds).Execute()

List Protected Objects.



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
    searchString := "searchString_example" // string | Specifies the search string to filter the objects. This search string will be applicable for objectnames and Protection Group names. User can specify a wildcard character '*' as a suffix to a string where all object and their Protection Group names are matched with the prefix string. For example, if vm1 and vm2 are the names of objects, user can specify vm* to list the objects. If not specified, then all the objects with Protection Groups will be returned which will match other filtering criteria. (optional)
    environments := []string{"Environments_example"} // []string | Specifies the environment type to filter objects. (optional)
    snapshotActions := []string{"SnapshotActions_example"} // []string | Specifies a list of recovery actions. Only snapshots that applies to these actions will be returned. (optional)
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

    request := onprem.ApiSearchProtectedObjectsRequest{
        SearchString: &searchString
        Environments: &environments
        SnapshotActions: &snapshotActions
        TenantIds: &tenantIds
        IncludeTenants: &includeTenants
        ProtectionGroupIds: &protectionGroupIds
        ObjectIds: &objectIds
        StorageDomainIds: &storageDomainIds
        SubResultSize: &subResultSize
        FilterSnapshotFromUsecs: &filterSnapshotFromUsecs
        FilterSnapshotToUsecs: &filterSnapshotToUsecs
        OsTypes: &osTypes
        SourceIds: &sourceIds
        RunInstanceIds: &runInstanceIds
        CdpProtectedOnly: &cdpProtectedOnly
        RegionIds: &regionIds
    }

    resp, r, err := api_client.Search.SearchProtectedObjects(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Search.SearchProtectedObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchProtectedObjects`: ProtectedObjectsSearchResponseBody
    fmt.Fprintf(os.Stdout, "Response from `Search.SearchProtectedObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchProtectedObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchString** | **string** | Specifies the search string to filter the objects. This search string will be applicable for objectnames and Protection Group names. User can specify a wildcard character &#39;*&#39; as a suffix to a string where all object and their Protection Group names are matched with the prefix string. For example, if vm1 and vm2 are the names of objects, user can specify vm* to list the objects. If not specified, then all the objects with Protection Groups will be returned which will match other filtering criteria. | 
 **environments** | **[]string** | Specifies the environment type to filter objects. | 
 **snapshotActions** | **[]string** | Specifies a list of recovery actions. Only snapshots that applies to these actions will be returned. | 
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

