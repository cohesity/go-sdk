# HeliosCommonSearchIndexedObjectsRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterIdentifiers** | Pointer to **[]string** | List of Clusters Identifiers to filter from. The format is clusterId:clusterIncarnationId. | [optional] 
**RegionIds** | Pointer to **[]string** | List of Regions to filter from. | [optional] 
**Count** | Pointer to **NullableInt32** | Specifies the number of indexed objects to be fetched. | [optional] 
**ObjectType** | **NullableString** | Specifies the object type to be searched for. | 
**SourceUUIDs** | Pointer to **[]string** | Specifies a list of source UUIDs. Only matches found in these sources will be returned. | [optional] 

## Methods

### NewHeliosCommonSearchIndexedObjectsRequestParams

`func NewHeliosCommonSearchIndexedObjectsRequestParams(objectType NullableString, ) *HeliosCommonSearchIndexedObjectsRequestParams`

NewHeliosCommonSearchIndexedObjectsRequestParams instantiates a new HeliosCommonSearchIndexedObjectsRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosCommonSearchIndexedObjectsRequestParamsWithDefaults

`func NewHeliosCommonSearchIndexedObjectsRequestParamsWithDefaults() *HeliosCommonSearchIndexedObjectsRequestParams`

NewHeliosCommonSearchIndexedObjectsRequestParamsWithDefaults instantiates a new HeliosCommonSearchIndexedObjectsRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterIdentifiers

`func (o *HeliosCommonSearchIndexedObjectsRequestParams) GetClusterIdentifiers() []string`

GetClusterIdentifiers returns the ClusterIdentifiers field if non-nil, zero value otherwise.

### GetClusterIdentifiersOk

`func (o *HeliosCommonSearchIndexedObjectsRequestParams) GetClusterIdentifiersOk() (*[]string, bool)`

GetClusterIdentifiersOk returns a tuple with the ClusterIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIdentifiers

`func (o *HeliosCommonSearchIndexedObjectsRequestParams) SetClusterIdentifiers(v []string)`

SetClusterIdentifiers sets ClusterIdentifiers field to given value.

### HasClusterIdentifiers

`func (o *HeliosCommonSearchIndexedObjectsRequestParams) HasClusterIdentifiers() bool`

HasClusterIdentifiers returns a boolean if a field has been set.

### SetClusterIdentifiersNil

`func (o *HeliosCommonSearchIndexedObjectsRequestParams) SetClusterIdentifiersNil(b bool)`

 SetClusterIdentifiersNil sets the value for ClusterIdentifiers to be an explicit nil

### UnsetClusterIdentifiers
`func (o *HeliosCommonSearchIndexedObjectsRequestParams) UnsetClusterIdentifiers()`

UnsetClusterIdentifiers ensures that no value is present for ClusterIdentifiers, not even an explicit nil
### GetRegionIds

`func (o *HeliosCommonSearchIndexedObjectsRequestParams) GetRegionIds() []string`

GetRegionIds returns the RegionIds field if non-nil, zero value otherwise.

### GetRegionIdsOk

`func (o *HeliosCommonSearchIndexedObjectsRequestParams) GetRegionIdsOk() (*[]string, bool)`

GetRegionIdsOk returns a tuple with the RegionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionIds

`func (o *HeliosCommonSearchIndexedObjectsRequestParams) SetRegionIds(v []string)`

SetRegionIds sets RegionIds field to given value.

### HasRegionIds

`func (o *HeliosCommonSearchIndexedObjectsRequestParams) HasRegionIds() bool`

HasRegionIds returns a boolean if a field has been set.

### SetRegionIdsNil

`func (o *HeliosCommonSearchIndexedObjectsRequestParams) SetRegionIdsNil(b bool)`

 SetRegionIdsNil sets the value for RegionIds to be an explicit nil

### UnsetRegionIds
`func (o *HeliosCommonSearchIndexedObjectsRequestParams) UnsetRegionIds()`

UnsetRegionIds ensures that no value is present for RegionIds, not even an explicit nil
### GetCount

`func (o *HeliosCommonSearchIndexedObjectsRequestParams) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *HeliosCommonSearchIndexedObjectsRequestParams) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *HeliosCommonSearchIndexedObjectsRequestParams) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *HeliosCommonSearchIndexedObjectsRequestParams) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *HeliosCommonSearchIndexedObjectsRequestParams) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *HeliosCommonSearchIndexedObjectsRequestParams) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetObjectType

`func (o *HeliosCommonSearchIndexedObjectsRequestParams) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HeliosCommonSearchIndexedObjectsRequestParams) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HeliosCommonSearchIndexedObjectsRequestParams) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### SetObjectTypeNil

`func (o *HeliosCommonSearchIndexedObjectsRequestParams) SetObjectTypeNil(b bool)`

 SetObjectTypeNil sets the value for ObjectType to be an explicit nil

### UnsetObjectType
`func (o *HeliosCommonSearchIndexedObjectsRequestParams) UnsetObjectType()`

UnsetObjectType ensures that no value is present for ObjectType, not even an explicit nil
### GetSourceUUIDs

`func (o *HeliosCommonSearchIndexedObjectsRequestParams) GetSourceUUIDs() []string`

GetSourceUUIDs returns the SourceUUIDs field if non-nil, zero value otherwise.

### GetSourceUUIDsOk

`func (o *HeliosCommonSearchIndexedObjectsRequestParams) GetSourceUUIDsOk() (*[]string, bool)`

GetSourceUUIDsOk returns a tuple with the SourceUUIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUUIDs

`func (o *HeliosCommonSearchIndexedObjectsRequestParams) SetSourceUUIDs(v []string)`

SetSourceUUIDs sets SourceUUIDs field to given value.

### HasSourceUUIDs

`func (o *HeliosCommonSearchIndexedObjectsRequestParams) HasSourceUUIDs() bool`

HasSourceUUIDs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


