# HeliosSearchIndexedObjectsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterIdentifiers** | Pointer to **[]string** | List of Clusters Identifiers to filter from. The format is clusterId:clusterIncarnationId. | [optional] 
**RegionIds** | Pointer to **[]string** | List of Regions to filter from. | [optional] 
**Count** | Pointer to **NullableInt32** | Specifies the number of indexed objects to be fetched. | [optional] 
**ObjectType** | **NullableString** | Specifies the object type to be searched for. | 
**SourceUUIDs** | Pointer to **[]string** | Specifies a list of source UUIDs. Only matches found in these sources will be returned. | [optional] 
**EmailParams** | Pointer to [**EmailHeliosSearchParams**](EmailHeliosSearchParams.md) |  | [optional] 
**FileParams** | Pointer to [**SearchFileRequestParamsBase**](SearchFileRequestParamsBase.md) |  | [optional] 
**CassandraParams** | Pointer to [**CassandraSearchParams**](CassandraSearchParams.md) |  | [optional] 
**CouchbaseParams** | Pointer to [**CouchbaseSearchParams**](CouchbaseSearchParams.md) |  | [optional] 
**HbaseParams** | Pointer to [**HbaseSearchParams**](HbaseSearchParams.md) |  | [optional] 
**HiveParams** | Pointer to [**HiveSearchParams**](HiveSearchParams.md) |  | [optional] 
**MongodbParams** | Pointer to [**MongodbSearchParams**](MongodbSearchParams.md) |  | [optional] 
**HdfsParams** | Pointer to [**HdfsSearchParams**](HdfsSearchParams.md) |  | [optional] 
**ExchangeParams** | Pointer to [**SearchExchangeObjectsRequestParams**](SearchExchangeObjectsRequestParams.md) |  | [optional] 
**PublicFolderParams** | Pointer to [**SearchPublicFolderRequestParams**](SearchPublicFolderRequestParams.md) |  | [optional] 
**MsTeamsParams** | Pointer to [**SearchMsTeamsRequestParams**](SearchMsTeamsRequestParams.md) |  | [optional] 

## Methods

### NewHeliosSearchIndexedObjectsRequest

`func NewHeliosSearchIndexedObjectsRequest(objectType NullableString, ) *HeliosSearchIndexedObjectsRequest`

NewHeliosSearchIndexedObjectsRequest instantiates a new HeliosSearchIndexedObjectsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosSearchIndexedObjectsRequestWithDefaults

`func NewHeliosSearchIndexedObjectsRequestWithDefaults() *HeliosSearchIndexedObjectsRequest`

NewHeliosSearchIndexedObjectsRequestWithDefaults instantiates a new HeliosSearchIndexedObjectsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterIdentifiers

`func (o *HeliosSearchIndexedObjectsRequest) GetClusterIdentifiers() []string`

GetClusterIdentifiers returns the ClusterIdentifiers field if non-nil, zero value otherwise.

### GetClusterIdentifiersOk

`func (o *HeliosSearchIndexedObjectsRequest) GetClusterIdentifiersOk() (*[]string, bool)`

GetClusterIdentifiersOk returns a tuple with the ClusterIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIdentifiers

`func (o *HeliosSearchIndexedObjectsRequest) SetClusterIdentifiers(v []string)`

SetClusterIdentifiers sets ClusterIdentifiers field to given value.

### HasClusterIdentifiers

`func (o *HeliosSearchIndexedObjectsRequest) HasClusterIdentifiers() bool`

HasClusterIdentifiers returns a boolean if a field has been set.

### SetClusterIdentifiersNil

`func (o *HeliosSearchIndexedObjectsRequest) SetClusterIdentifiersNil(b bool)`

 SetClusterIdentifiersNil sets the value for ClusterIdentifiers to be an explicit nil

### UnsetClusterIdentifiers
`func (o *HeliosSearchIndexedObjectsRequest) UnsetClusterIdentifiers()`

UnsetClusterIdentifiers ensures that no value is present for ClusterIdentifiers, not even an explicit nil
### GetRegionIds

`func (o *HeliosSearchIndexedObjectsRequest) GetRegionIds() []string`

GetRegionIds returns the RegionIds field if non-nil, zero value otherwise.

### GetRegionIdsOk

`func (o *HeliosSearchIndexedObjectsRequest) GetRegionIdsOk() (*[]string, bool)`

GetRegionIdsOk returns a tuple with the RegionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionIds

`func (o *HeliosSearchIndexedObjectsRequest) SetRegionIds(v []string)`

SetRegionIds sets RegionIds field to given value.

### HasRegionIds

`func (o *HeliosSearchIndexedObjectsRequest) HasRegionIds() bool`

HasRegionIds returns a boolean if a field has been set.

### SetRegionIdsNil

`func (o *HeliosSearchIndexedObjectsRequest) SetRegionIdsNil(b bool)`

 SetRegionIdsNil sets the value for RegionIds to be an explicit nil

### UnsetRegionIds
`func (o *HeliosSearchIndexedObjectsRequest) UnsetRegionIds()`

UnsetRegionIds ensures that no value is present for RegionIds, not even an explicit nil
### GetCount

`func (o *HeliosSearchIndexedObjectsRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *HeliosSearchIndexedObjectsRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *HeliosSearchIndexedObjectsRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *HeliosSearchIndexedObjectsRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *HeliosSearchIndexedObjectsRequest) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *HeliosSearchIndexedObjectsRequest) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetObjectType

`func (o *HeliosSearchIndexedObjectsRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HeliosSearchIndexedObjectsRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HeliosSearchIndexedObjectsRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### SetObjectTypeNil

`func (o *HeliosSearchIndexedObjectsRequest) SetObjectTypeNil(b bool)`

 SetObjectTypeNil sets the value for ObjectType to be an explicit nil

### UnsetObjectType
`func (o *HeliosSearchIndexedObjectsRequest) UnsetObjectType()`

UnsetObjectType ensures that no value is present for ObjectType, not even an explicit nil
### GetSourceUUIDs

`func (o *HeliosSearchIndexedObjectsRequest) GetSourceUUIDs() []string`

GetSourceUUIDs returns the SourceUUIDs field if non-nil, zero value otherwise.

### GetSourceUUIDsOk

`func (o *HeliosSearchIndexedObjectsRequest) GetSourceUUIDsOk() (*[]string, bool)`

GetSourceUUIDsOk returns a tuple with the SourceUUIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUUIDs

`func (o *HeliosSearchIndexedObjectsRequest) SetSourceUUIDs(v []string)`

SetSourceUUIDs sets SourceUUIDs field to given value.

### HasSourceUUIDs

`func (o *HeliosSearchIndexedObjectsRequest) HasSourceUUIDs() bool`

HasSourceUUIDs returns a boolean if a field has been set.

### GetEmailParams

`func (o *HeliosSearchIndexedObjectsRequest) GetEmailParams() EmailHeliosSearchParams`

GetEmailParams returns the EmailParams field if non-nil, zero value otherwise.

### GetEmailParamsOk

`func (o *HeliosSearchIndexedObjectsRequest) GetEmailParamsOk() (*EmailHeliosSearchParams, bool)`

GetEmailParamsOk returns a tuple with the EmailParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailParams

`func (o *HeliosSearchIndexedObjectsRequest) SetEmailParams(v EmailHeliosSearchParams)`

SetEmailParams sets EmailParams field to given value.

### HasEmailParams

`func (o *HeliosSearchIndexedObjectsRequest) HasEmailParams() bool`

HasEmailParams returns a boolean if a field has been set.

### GetFileParams

`func (o *HeliosSearchIndexedObjectsRequest) GetFileParams() SearchFileRequestParamsBase`

GetFileParams returns the FileParams field if non-nil, zero value otherwise.

### GetFileParamsOk

`func (o *HeliosSearchIndexedObjectsRequest) GetFileParamsOk() (*SearchFileRequestParamsBase, bool)`

GetFileParamsOk returns a tuple with the FileParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileParams

`func (o *HeliosSearchIndexedObjectsRequest) SetFileParams(v SearchFileRequestParamsBase)`

SetFileParams sets FileParams field to given value.

### HasFileParams

`func (o *HeliosSearchIndexedObjectsRequest) HasFileParams() bool`

HasFileParams returns a boolean if a field has been set.

### GetCassandraParams

`func (o *HeliosSearchIndexedObjectsRequest) GetCassandraParams() CassandraSearchParams`

GetCassandraParams returns the CassandraParams field if non-nil, zero value otherwise.

### GetCassandraParamsOk

`func (o *HeliosSearchIndexedObjectsRequest) GetCassandraParamsOk() (*CassandraSearchParams, bool)`

GetCassandraParamsOk returns a tuple with the CassandraParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraParams

`func (o *HeliosSearchIndexedObjectsRequest) SetCassandraParams(v CassandraSearchParams)`

SetCassandraParams sets CassandraParams field to given value.

### HasCassandraParams

`func (o *HeliosSearchIndexedObjectsRequest) HasCassandraParams() bool`

HasCassandraParams returns a boolean if a field has been set.

### GetCouchbaseParams

`func (o *HeliosSearchIndexedObjectsRequest) GetCouchbaseParams() CouchbaseSearchParams`

GetCouchbaseParams returns the CouchbaseParams field if non-nil, zero value otherwise.

### GetCouchbaseParamsOk

`func (o *HeliosSearchIndexedObjectsRequest) GetCouchbaseParamsOk() (*CouchbaseSearchParams, bool)`

GetCouchbaseParamsOk returns a tuple with the CouchbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouchbaseParams

`func (o *HeliosSearchIndexedObjectsRequest) SetCouchbaseParams(v CouchbaseSearchParams)`

SetCouchbaseParams sets CouchbaseParams field to given value.

### HasCouchbaseParams

`func (o *HeliosSearchIndexedObjectsRequest) HasCouchbaseParams() bool`

HasCouchbaseParams returns a boolean if a field has been set.

### GetHbaseParams

`func (o *HeliosSearchIndexedObjectsRequest) GetHbaseParams() HbaseSearchParams`

GetHbaseParams returns the HbaseParams field if non-nil, zero value otherwise.

### GetHbaseParamsOk

`func (o *HeliosSearchIndexedObjectsRequest) GetHbaseParamsOk() (*HbaseSearchParams, bool)`

GetHbaseParamsOk returns a tuple with the HbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseParams

`func (o *HeliosSearchIndexedObjectsRequest) SetHbaseParams(v HbaseSearchParams)`

SetHbaseParams sets HbaseParams field to given value.

### HasHbaseParams

`func (o *HeliosSearchIndexedObjectsRequest) HasHbaseParams() bool`

HasHbaseParams returns a boolean if a field has been set.

### GetHiveParams

`func (o *HeliosSearchIndexedObjectsRequest) GetHiveParams() HiveSearchParams`

GetHiveParams returns the HiveParams field if non-nil, zero value otherwise.

### GetHiveParamsOk

`func (o *HeliosSearchIndexedObjectsRequest) GetHiveParamsOk() (*HiveSearchParams, bool)`

GetHiveParamsOk returns a tuple with the HiveParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveParams

`func (o *HeliosSearchIndexedObjectsRequest) SetHiveParams(v HiveSearchParams)`

SetHiveParams sets HiveParams field to given value.

### HasHiveParams

`func (o *HeliosSearchIndexedObjectsRequest) HasHiveParams() bool`

HasHiveParams returns a boolean if a field has been set.

### GetMongodbParams

`func (o *HeliosSearchIndexedObjectsRequest) GetMongodbParams() MongodbSearchParams`

GetMongodbParams returns the MongodbParams field if non-nil, zero value otherwise.

### GetMongodbParamsOk

`func (o *HeliosSearchIndexedObjectsRequest) GetMongodbParamsOk() (*MongodbSearchParams, bool)`

GetMongodbParamsOk returns a tuple with the MongodbParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbParams

`func (o *HeliosSearchIndexedObjectsRequest) SetMongodbParams(v MongodbSearchParams)`

SetMongodbParams sets MongodbParams field to given value.

### HasMongodbParams

`func (o *HeliosSearchIndexedObjectsRequest) HasMongodbParams() bool`

HasMongodbParams returns a boolean if a field has been set.

### GetHdfsParams

`func (o *HeliosSearchIndexedObjectsRequest) GetHdfsParams() HdfsSearchParams`

GetHdfsParams returns the HdfsParams field if non-nil, zero value otherwise.

### GetHdfsParamsOk

`func (o *HeliosSearchIndexedObjectsRequest) GetHdfsParamsOk() (*HdfsSearchParams, bool)`

GetHdfsParamsOk returns a tuple with the HdfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsParams

`func (o *HeliosSearchIndexedObjectsRequest) SetHdfsParams(v HdfsSearchParams)`

SetHdfsParams sets HdfsParams field to given value.

### HasHdfsParams

`func (o *HeliosSearchIndexedObjectsRequest) HasHdfsParams() bool`

HasHdfsParams returns a boolean if a field has been set.

### GetExchangeParams

`func (o *HeliosSearchIndexedObjectsRequest) GetExchangeParams() SearchExchangeObjectsRequestParams`

GetExchangeParams returns the ExchangeParams field if non-nil, zero value otherwise.

### GetExchangeParamsOk

`func (o *HeliosSearchIndexedObjectsRequest) GetExchangeParamsOk() (*SearchExchangeObjectsRequestParams, bool)`

GetExchangeParamsOk returns a tuple with the ExchangeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeParams

`func (o *HeliosSearchIndexedObjectsRequest) SetExchangeParams(v SearchExchangeObjectsRequestParams)`

SetExchangeParams sets ExchangeParams field to given value.

### HasExchangeParams

`func (o *HeliosSearchIndexedObjectsRequest) HasExchangeParams() bool`

HasExchangeParams returns a boolean if a field has been set.

### GetPublicFolderParams

`func (o *HeliosSearchIndexedObjectsRequest) GetPublicFolderParams() SearchPublicFolderRequestParams`

GetPublicFolderParams returns the PublicFolderParams field if non-nil, zero value otherwise.

### GetPublicFolderParamsOk

`func (o *HeliosSearchIndexedObjectsRequest) GetPublicFolderParamsOk() (*SearchPublicFolderRequestParams, bool)`

GetPublicFolderParamsOk returns a tuple with the PublicFolderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicFolderParams

`func (o *HeliosSearchIndexedObjectsRequest) SetPublicFolderParams(v SearchPublicFolderRequestParams)`

SetPublicFolderParams sets PublicFolderParams field to given value.

### HasPublicFolderParams

`func (o *HeliosSearchIndexedObjectsRequest) HasPublicFolderParams() bool`

HasPublicFolderParams returns a boolean if a field has been set.

### GetMsTeamsParams

`func (o *HeliosSearchIndexedObjectsRequest) GetMsTeamsParams() SearchMsTeamsRequestParams`

GetMsTeamsParams returns the MsTeamsParams field if non-nil, zero value otherwise.

### GetMsTeamsParamsOk

`func (o *HeliosSearchIndexedObjectsRequest) GetMsTeamsParamsOk() (*SearchMsTeamsRequestParams, bool)`

GetMsTeamsParamsOk returns a tuple with the MsTeamsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsTeamsParams

`func (o *HeliosSearchIndexedObjectsRequest) SetMsTeamsParams(v SearchMsTeamsRequestParams)`

SetMsTeamsParams sets MsTeamsParams field to given value.

### HasMsTeamsParams

`func (o *HeliosSearchIndexedObjectsRequest) HasMsTeamsParams() bool`

HasMsTeamsParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


