# SearchIndexedObjectsRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailParams** | Pointer to [**SearchEmailRequestParams**](SearchEmailRequestParams.md) |  | [optional] 
**FileParams** | Pointer to [**SearchFileRequestParams**](SearchFileRequestParams.md) |  | [optional] 
**CassandraParams** | Pointer to [**CassandraOnPremSearchParams**](CassandraOnPremSearchParams.md) |  | [optional] 
**CouchbaseParams** | Pointer to [**CouchBaseOnPremSearchParams**](CouchBaseOnPremSearchParams.md) |  | [optional] 
**HbaseParams** | Pointer to [**HbaseOnPremSearchParams**](HbaseOnPremSearchParams.md) |  | [optional] 
**HiveParams** | Pointer to [**HiveOnPremSearchParams**](HiveOnPremSearchParams.md) |  | [optional] 
**MongodbParams** | Pointer to [**MongoDbOnPremSearchParams**](MongoDbOnPremSearchParams.md) |  | [optional] 
**HdfsParams** | Pointer to [**HDFSOnPremSearchParams**](HDFSOnPremSearchParams.md) |  | [optional] 
**ExchangeParams** | Pointer to [**SearchExchangeObjectsRequestParams**](SearchExchangeObjectsRequestParams.md) |  | [optional] 
**PublicFolderParams** | Pointer to [**SearchPublicFolderRequestParams**](SearchPublicFolderRequestParams.md) |  | [optional] 
**MsTeamsParams** | Pointer to [**SearchMsTeamsRequestParams**](SearchMsTeamsRequestParams.md) |  | [optional] 
**SharepointParams** | Pointer to [**SearchDocumentLibraryRequestParams**](SearchDocumentLibraryRequestParams.md) |  | [optional] 

## Methods

### NewSearchIndexedObjectsRequestAllOf

`func NewSearchIndexedObjectsRequestAllOf() *SearchIndexedObjectsRequestAllOf`

NewSearchIndexedObjectsRequestAllOf instantiates a new SearchIndexedObjectsRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchIndexedObjectsRequestAllOfWithDefaults

`func NewSearchIndexedObjectsRequestAllOfWithDefaults() *SearchIndexedObjectsRequestAllOf`

NewSearchIndexedObjectsRequestAllOfWithDefaults instantiates a new SearchIndexedObjectsRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailParams

`func (o *SearchIndexedObjectsRequestAllOf) GetEmailParams() SearchEmailRequestParams`

GetEmailParams returns the EmailParams field if non-nil, zero value otherwise.

### GetEmailParamsOk

`func (o *SearchIndexedObjectsRequestAllOf) GetEmailParamsOk() (*SearchEmailRequestParams, bool)`

GetEmailParamsOk returns a tuple with the EmailParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailParams

`func (o *SearchIndexedObjectsRequestAllOf) SetEmailParams(v SearchEmailRequestParams)`

SetEmailParams sets EmailParams field to given value.

### HasEmailParams

`func (o *SearchIndexedObjectsRequestAllOf) HasEmailParams() bool`

HasEmailParams returns a boolean if a field has been set.

### GetFileParams

`func (o *SearchIndexedObjectsRequestAllOf) GetFileParams() SearchFileRequestParams`

GetFileParams returns the FileParams field if non-nil, zero value otherwise.

### GetFileParamsOk

`func (o *SearchIndexedObjectsRequestAllOf) GetFileParamsOk() (*SearchFileRequestParams, bool)`

GetFileParamsOk returns a tuple with the FileParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileParams

`func (o *SearchIndexedObjectsRequestAllOf) SetFileParams(v SearchFileRequestParams)`

SetFileParams sets FileParams field to given value.

### HasFileParams

`func (o *SearchIndexedObjectsRequestAllOf) HasFileParams() bool`

HasFileParams returns a boolean if a field has been set.

### GetCassandraParams

`func (o *SearchIndexedObjectsRequestAllOf) GetCassandraParams() CassandraOnPremSearchParams`

GetCassandraParams returns the CassandraParams field if non-nil, zero value otherwise.

### GetCassandraParamsOk

`func (o *SearchIndexedObjectsRequestAllOf) GetCassandraParamsOk() (*CassandraOnPremSearchParams, bool)`

GetCassandraParamsOk returns a tuple with the CassandraParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraParams

`func (o *SearchIndexedObjectsRequestAllOf) SetCassandraParams(v CassandraOnPremSearchParams)`

SetCassandraParams sets CassandraParams field to given value.

### HasCassandraParams

`func (o *SearchIndexedObjectsRequestAllOf) HasCassandraParams() bool`

HasCassandraParams returns a boolean if a field has been set.

### GetCouchbaseParams

`func (o *SearchIndexedObjectsRequestAllOf) GetCouchbaseParams() CouchBaseOnPremSearchParams`

GetCouchbaseParams returns the CouchbaseParams field if non-nil, zero value otherwise.

### GetCouchbaseParamsOk

`func (o *SearchIndexedObjectsRequestAllOf) GetCouchbaseParamsOk() (*CouchBaseOnPremSearchParams, bool)`

GetCouchbaseParamsOk returns a tuple with the CouchbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouchbaseParams

`func (o *SearchIndexedObjectsRequestAllOf) SetCouchbaseParams(v CouchBaseOnPremSearchParams)`

SetCouchbaseParams sets CouchbaseParams field to given value.

### HasCouchbaseParams

`func (o *SearchIndexedObjectsRequestAllOf) HasCouchbaseParams() bool`

HasCouchbaseParams returns a boolean if a field has been set.

### GetHbaseParams

`func (o *SearchIndexedObjectsRequestAllOf) GetHbaseParams() HbaseOnPremSearchParams`

GetHbaseParams returns the HbaseParams field if non-nil, zero value otherwise.

### GetHbaseParamsOk

`func (o *SearchIndexedObjectsRequestAllOf) GetHbaseParamsOk() (*HbaseOnPremSearchParams, bool)`

GetHbaseParamsOk returns a tuple with the HbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseParams

`func (o *SearchIndexedObjectsRequestAllOf) SetHbaseParams(v HbaseOnPremSearchParams)`

SetHbaseParams sets HbaseParams field to given value.

### HasHbaseParams

`func (o *SearchIndexedObjectsRequestAllOf) HasHbaseParams() bool`

HasHbaseParams returns a boolean if a field has been set.

### GetHiveParams

`func (o *SearchIndexedObjectsRequestAllOf) GetHiveParams() HiveOnPremSearchParams`

GetHiveParams returns the HiveParams field if non-nil, zero value otherwise.

### GetHiveParamsOk

`func (o *SearchIndexedObjectsRequestAllOf) GetHiveParamsOk() (*HiveOnPremSearchParams, bool)`

GetHiveParamsOk returns a tuple with the HiveParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveParams

`func (o *SearchIndexedObjectsRequestAllOf) SetHiveParams(v HiveOnPremSearchParams)`

SetHiveParams sets HiveParams field to given value.

### HasHiveParams

`func (o *SearchIndexedObjectsRequestAllOf) HasHiveParams() bool`

HasHiveParams returns a boolean if a field has been set.

### GetMongodbParams

`func (o *SearchIndexedObjectsRequestAllOf) GetMongodbParams() MongoDbOnPremSearchParams`

GetMongodbParams returns the MongodbParams field if non-nil, zero value otherwise.

### GetMongodbParamsOk

`func (o *SearchIndexedObjectsRequestAllOf) GetMongodbParamsOk() (*MongoDbOnPremSearchParams, bool)`

GetMongodbParamsOk returns a tuple with the MongodbParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbParams

`func (o *SearchIndexedObjectsRequestAllOf) SetMongodbParams(v MongoDbOnPremSearchParams)`

SetMongodbParams sets MongodbParams field to given value.

### HasMongodbParams

`func (o *SearchIndexedObjectsRequestAllOf) HasMongodbParams() bool`

HasMongodbParams returns a boolean if a field has been set.

### GetHdfsParams

`func (o *SearchIndexedObjectsRequestAllOf) GetHdfsParams() HDFSOnPremSearchParams`

GetHdfsParams returns the HdfsParams field if non-nil, zero value otherwise.

### GetHdfsParamsOk

`func (o *SearchIndexedObjectsRequestAllOf) GetHdfsParamsOk() (*HDFSOnPremSearchParams, bool)`

GetHdfsParamsOk returns a tuple with the HdfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsParams

`func (o *SearchIndexedObjectsRequestAllOf) SetHdfsParams(v HDFSOnPremSearchParams)`

SetHdfsParams sets HdfsParams field to given value.

### HasHdfsParams

`func (o *SearchIndexedObjectsRequestAllOf) HasHdfsParams() bool`

HasHdfsParams returns a boolean if a field has been set.

### GetExchangeParams

`func (o *SearchIndexedObjectsRequestAllOf) GetExchangeParams() SearchExchangeObjectsRequestParams`

GetExchangeParams returns the ExchangeParams field if non-nil, zero value otherwise.

### GetExchangeParamsOk

`func (o *SearchIndexedObjectsRequestAllOf) GetExchangeParamsOk() (*SearchExchangeObjectsRequestParams, bool)`

GetExchangeParamsOk returns a tuple with the ExchangeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeParams

`func (o *SearchIndexedObjectsRequestAllOf) SetExchangeParams(v SearchExchangeObjectsRequestParams)`

SetExchangeParams sets ExchangeParams field to given value.

### HasExchangeParams

`func (o *SearchIndexedObjectsRequestAllOf) HasExchangeParams() bool`

HasExchangeParams returns a boolean if a field has been set.

### GetPublicFolderParams

`func (o *SearchIndexedObjectsRequestAllOf) GetPublicFolderParams() SearchPublicFolderRequestParams`

GetPublicFolderParams returns the PublicFolderParams field if non-nil, zero value otherwise.

### GetPublicFolderParamsOk

`func (o *SearchIndexedObjectsRequestAllOf) GetPublicFolderParamsOk() (*SearchPublicFolderRequestParams, bool)`

GetPublicFolderParamsOk returns a tuple with the PublicFolderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicFolderParams

`func (o *SearchIndexedObjectsRequestAllOf) SetPublicFolderParams(v SearchPublicFolderRequestParams)`

SetPublicFolderParams sets PublicFolderParams field to given value.

### HasPublicFolderParams

`func (o *SearchIndexedObjectsRequestAllOf) HasPublicFolderParams() bool`

HasPublicFolderParams returns a boolean if a field has been set.

### GetMsTeamsParams

`func (o *SearchIndexedObjectsRequestAllOf) GetMsTeamsParams() SearchMsTeamsRequestParams`

GetMsTeamsParams returns the MsTeamsParams field if non-nil, zero value otherwise.

### GetMsTeamsParamsOk

`func (o *SearchIndexedObjectsRequestAllOf) GetMsTeamsParamsOk() (*SearchMsTeamsRequestParams, bool)`

GetMsTeamsParamsOk returns a tuple with the MsTeamsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsTeamsParams

`func (o *SearchIndexedObjectsRequestAllOf) SetMsTeamsParams(v SearchMsTeamsRequestParams)`

SetMsTeamsParams sets MsTeamsParams field to given value.

### HasMsTeamsParams

`func (o *SearchIndexedObjectsRequestAllOf) HasMsTeamsParams() bool`

HasMsTeamsParams returns a boolean if a field has been set.

### GetSharepointParams

`func (o *SearchIndexedObjectsRequestAllOf) GetSharepointParams() SearchDocumentLibraryRequestParams`

GetSharepointParams returns the SharepointParams field if non-nil, zero value otherwise.

### GetSharepointParamsOk

`func (o *SearchIndexedObjectsRequestAllOf) GetSharepointParamsOk() (*SearchDocumentLibraryRequestParams, bool)`

GetSharepointParamsOk returns a tuple with the SharepointParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharepointParams

`func (o *SearchIndexedObjectsRequestAllOf) SetSharepointParams(v SearchDocumentLibraryRequestParams)`

SetSharepointParams sets SharepointParams field to given value.

### HasSharepointParams

`func (o *SearchIndexedObjectsRequestAllOf) HasSharepointParams() bool`

HasSharepointParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


