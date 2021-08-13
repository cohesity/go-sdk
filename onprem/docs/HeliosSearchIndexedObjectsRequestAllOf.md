# HeliosSearchIndexedObjectsRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewHeliosSearchIndexedObjectsRequestAllOf

`func NewHeliosSearchIndexedObjectsRequestAllOf() *HeliosSearchIndexedObjectsRequestAllOf`

NewHeliosSearchIndexedObjectsRequestAllOf instantiates a new HeliosSearchIndexedObjectsRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosSearchIndexedObjectsRequestAllOfWithDefaults

`func NewHeliosSearchIndexedObjectsRequestAllOfWithDefaults() *HeliosSearchIndexedObjectsRequestAllOf`

NewHeliosSearchIndexedObjectsRequestAllOfWithDefaults instantiates a new HeliosSearchIndexedObjectsRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) GetEmailParams() EmailHeliosSearchParams`

GetEmailParams returns the EmailParams field if non-nil, zero value otherwise.

### GetEmailParamsOk

`func (o *HeliosSearchIndexedObjectsRequestAllOf) GetEmailParamsOk() (*EmailHeliosSearchParams, bool)`

GetEmailParamsOk returns a tuple with the EmailParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) SetEmailParams(v EmailHeliosSearchParams)`

SetEmailParams sets EmailParams field to given value.

### HasEmailParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) HasEmailParams() bool`

HasEmailParams returns a boolean if a field has been set.

### GetFileParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) GetFileParams() SearchFileRequestParamsBase`

GetFileParams returns the FileParams field if non-nil, zero value otherwise.

### GetFileParamsOk

`func (o *HeliosSearchIndexedObjectsRequestAllOf) GetFileParamsOk() (*SearchFileRequestParamsBase, bool)`

GetFileParamsOk returns a tuple with the FileParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) SetFileParams(v SearchFileRequestParamsBase)`

SetFileParams sets FileParams field to given value.

### HasFileParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) HasFileParams() bool`

HasFileParams returns a boolean if a field has been set.

### GetCassandraParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) GetCassandraParams() CassandraSearchParams`

GetCassandraParams returns the CassandraParams field if non-nil, zero value otherwise.

### GetCassandraParamsOk

`func (o *HeliosSearchIndexedObjectsRequestAllOf) GetCassandraParamsOk() (*CassandraSearchParams, bool)`

GetCassandraParamsOk returns a tuple with the CassandraParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) SetCassandraParams(v CassandraSearchParams)`

SetCassandraParams sets CassandraParams field to given value.

### HasCassandraParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) HasCassandraParams() bool`

HasCassandraParams returns a boolean if a field has been set.

### GetCouchbaseParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) GetCouchbaseParams() CouchbaseSearchParams`

GetCouchbaseParams returns the CouchbaseParams field if non-nil, zero value otherwise.

### GetCouchbaseParamsOk

`func (o *HeliosSearchIndexedObjectsRequestAllOf) GetCouchbaseParamsOk() (*CouchbaseSearchParams, bool)`

GetCouchbaseParamsOk returns a tuple with the CouchbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouchbaseParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) SetCouchbaseParams(v CouchbaseSearchParams)`

SetCouchbaseParams sets CouchbaseParams field to given value.

### HasCouchbaseParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) HasCouchbaseParams() bool`

HasCouchbaseParams returns a boolean if a field has been set.

### GetHbaseParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) GetHbaseParams() HbaseSearchParams`

GetHbaseParams returns the HbaseParams field if non-nil, zero value otherwise.

### GetHbaseParamsOk

`func (o *HeliosSearchIndexedObjectsRequestAllOf) GetHbaseParamsOk() (*HbaseSearchParams, bool)`

GetHbaseParamsOk returns a tuple with the HbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) SetHbaseParams(v HbaseSearchParams)`

SetHbaseParams sets HbaseParams field to given value.

### HasHbaseParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) HasHbaseParams() bool`

HasHbaseParams returns a boolean if a field has been set.

### GetHiveParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) GetHiveParams() HiveSearchParams`

GetHiveParams returns the HiveParams field if non-nil, zero value otherwise.

### GetHiveParamsOk

`func (o *HeliosSearchIndexedObjectsRequestAllOf) GetHiveParamsOk() (*HiveSearchParams, bool)`

GetHiveParamsOk returns a tuple with the HiveParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) SetHiveParams(v HiveSearchParams)`

SetHiveParams sets HiveParams field to given value.

### HasHiveParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) HasHiveParams() bool`

HasHiveParams returns a boolean if a field has been set.

### GetMongodbParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) GetMongodbParams() MongodbSearchParams`

GetMongodbParams returns the MongodbParams field if non-nil, zero value otherwise.

### GetMongodbParamsOk

`func (o *HeliosSearchIndexedObjectsRequestAllOf) GetMongodbParamsOk() (*MongodbSearchParams, bool)`

GetMongodbParamsOk returns a tuple with the MongodbParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) SetMongodbParams(v MongodbSearchParams)`

SetMongodbParams sets MongodbParams field to given value.

### HasMongodbParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) HasMongodbParams() bool`

HasMongodbParams returns a boolean if a field has been set.

### GetHdfsParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) GetHdfsParams() HdfsSearchParams`

GetHdfsParams returns the HdfsParams field if non-nil, zero value otherwise.

### GetHdfsParamsOk

`func (o *HeliosSearchIndexedObjectsRequestAllOf) GetHdfsParamsOk() (*HdfsSearchParams, bool)`

GetHdfsParamsOk returns a tuple with the HdfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) SetHdfsParams(v HdfsSearchParams)`

SetHdfsParams sets HdfsParams field to given value.

### HasHdfsParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) HasHdfsParams() bool`

HasHdfsParams returns a boolean if a field has been set.

### GetExchangeParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) GetExchangeParams() SearchExchangeObjectsRequestParams`

GetExchangeParams returns the ExchangeParams field if non-nil, zero value otherwise.

### GetExchangeParamsOk

`func (o *HeliosSearchIndexedObjectsRequestAllOf) GetExchangeParamsOk() (*SearchExchangeObjectsRequestParams, bool)`

GetExchangeParamsOk returns a tuple with the ExchangeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) SetExchangeParams(v SearchExchangeObjectsRequestParams)`

SetExchangeParams sets ExchangeParams field to given value.

### HasExchangeParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) HasExchangeParams() bool`

HasExchangeParams returns a boolean if a field has been set.

### GetPublicFolderParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) GetPublicFolderParams() SearchPublicFolderRequestParams`

GetPublicFolderParams returns the PublicFolderParams field if non-nil, zero value otherwise.

### GetPublicFolderParamsOk

`func (o *HeliosSearchIndexedObjectsRequestAllOf) GetPublicFolderParamsOk() (*SearchPublicFolderRequestParams, bool)`

GetPublicFolderParamsOk returns a tuple with the PublicFolderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicFolderParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) SetPublicFolderParams(v SearchPublicFolderRequestParams)`

SetPublicFolderParams sets PublicFolderParams field to given value.

### HasPublicFolderParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) HasPublicFolderParams() bool`

HasPublicFolderParams returns a boolean if a field has been set.

### GetMsTeamsParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) GetMsTeamsParams() SearchMsTeamsRequestParams`

GetMsTeamsParams returns the MsTeamsParams field if non-nil, zero value otherwise.

### GetMsTeamsParamsOk

`func (o *HeliosSearchIndexedObjectsRequestAllOf) GetMsTeamsParamsOk() (*SearchMsTeamsRequestParams, bool)`

GetMsTeamsParamsOk returns a tuple with the MsTeamsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsTeamsParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) SetMsTeamsParams(v SearchMsTeamsRequestParams)`

SetMsTeamsParams sets MsTeamsParams field to given value.

### HasMsTeamsParams

`func (o *HeliosSearchIndexedObjectsRequestAllOf) HasMsTeamsParams() bool`

HasMsTeamsParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


