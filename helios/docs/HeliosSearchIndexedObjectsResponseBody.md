# HeliosSearchIndexedObjectsResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | Pointer to **NullableString** | Specifies the object type. | [optional] 
**Count** | Pointer to **NullableInt32** | Specifies the total number of indexed objects that match the filter and search criteria. Use this value to determine how many additional requests are required to get the full result. | [optional] 
**ClusterErrors** | Pointer to [**[]HeliosSearchIndexedObjectsClusterError**](HeliosSearchIndexedObjectsClusterError.md) | A List of errors that occured on a subset of clusters. | [optional] 
**Emails** | Pointer to [**[]GlobalClusterIdentifier**](GlobalClusterIdentifier.md) | Specifies the indexed emails and email folders. | [optional] 
**Files** | Pointer to [**[]GlobalClusterIdentifier**](GlobalClusterIdentifier.md) | Specifies the indexed files and file folders. | [optional] 
**CassandraObjects** | Pointer to [**[]GlobalClusterIdentifier**](GlobalClusterIdentifier.md) | Specifies the indexed Cassandra objects. | [optional] 
**CouchbaseObjects** | Pointer to [**[]GlobalClusterIdentifier**](GlobalClusterIdentifier.md) | Specifies the indexed Couchbase objects. | [optional] 
**HbaseObjects** | Pointer to [**[]GlobalClusterIdentifier**](GlobalClusterIdentifier.md) | Specifies the indexed Hbase objects. | [optional] 
**HiveObjects** | Pointer to [**[]GlobalClusterIdentifier**](GlobalClusterIdentifier.md) | Specifies the indexed Hive objects. | [optional] 
**MongoObjects** | Pointer to [**[]GlobalClusterIdentifier**](GlobalClusterIdentifier.md) | Specifies the indexed Mongo objects. | [optional] 
**HdfsObjects** | Pointer to [**[]GlobalClusterIdentifier**](GlobalClusterIdentifier.md) | Specifies the indexed HDFS objects. | [optional] 
**ExchangeObjects** | Pointer to [**[]GlobalClusterIdentifier**](GlobalClusterIdentifier.md) | Specifies the indexed HDFS objects. | [optional] 
**PublicFolderItems** | Pointer to [**[]GlobalClusterIdentifier**](GlobalClusterIdentifier.md) | Specifies the indexed Public folder items. | [optional] 

## Methods

### NewHeliosSearchIndexedObjectsResponseBody

`func NewHeliosSearchIndexedObjectsResponseBody() *HeliosSearchIndexedObjectsResponseBody`

NewHeliosSearchIndexedObjectsResponseBody instantiates a new HeliosSearchIndexedObjectsResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosSearchIndexedObjectsResponseBodyWithDefaults

`func NewHeliosSearchIndexedObjectsResponseBodyWithDefaults() *HeliosSearchIndexedObjectsResponseBody`

NewHeliosSearchIndexedObjectsResponseBodyWithDefaults instantiates a new HeliosSearchIndexedObjectsResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *HeliosSearchIndexedObjectsResponseBody) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HeliosSearchIndexedObjectsResponseBody) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HeliosSearchIndexedObjectsResponseBody) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *HeliosSearchIndexedObjectsResponseBody) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### SetObjectTypeNil

`func (o *HeliosSearchIndexedObjectsResponseBody) SetObjectTypeNil(b bool)`

 SetObjectTypeNil sets the value for ObjectType to be an explicit nil

### UnsetObjectType
`func (o *HeliosSearchIndexedObjectsResponseBody) UnsetObjectType()`

UnsetObjectType ensures that no value is present for ObjectType, not even an explicit nil
### GetCount

`func (o *HeliosSearchIndexedObjectsResponseBody) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *HeliosSearchIndexedObjectsResponseBody) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *HeliosSearchIndexedObjectsResponseBody) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *HeliosSearchIndexedObjectsResponseBody) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *HeliosSearchIndexedObjectsResponseBody) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *HeliosSearchIndexedObjectsResponseBody) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetClusterErrors

`func (o *HeliosSearchIndexedObjectsResponseBody) GetClusterErrors() []HeliosSearchIndexedObjectsClusterError`

GetClusterErrors returns the ClusterErrors field if non-nil, zero value otherwise.

### GetClusterErrorsOk

`func (o *HeliosSearchIndexedObjectsResponseBody) GetClusterErrorsOk() (*[]HeliosSearchIndexedObjectsClusterError, bool)`

GetClusterErrorsOk returns a tuple with the ClusterErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterErrors

`func (o *HeliosSearchIndexedObjectsResponseBody) SetClusterErrors(v []HeliosSearchIndexedObjectsClusterError)`

SetClusterErrors sets ClusterErrors field to given value.

### HasClusterErrors

`func (o *HeliosSearchIndexedObjectsResponseBody) HasClusterErrors() bool`

HasClusterErrors returns a boolean if a field has been set.

### GetEmails

`func (o *HeliosSearchIndexedObjectsResponseBody) GetEmails() []GlobalClusterIdentifier`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *HeliosSearchIndexedObjectsResponseBody) GetEmailsOk() (*[]GlobalClusterIdentifier, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *HeliosSearchIndexedObjectsResponseBody) SetEmails(v []GlobalClusterIdentifier)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *HeliosSearchIndexedObjectsResponseBody) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetFiles

`func (o *HeliosSearchIndexedObjectsResponseBody) GetFiles() []GlobalClusterIdentifier`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *HeliosSearchIndexedObjectsResponseBody) GetFilesOk() (*[]GlobalClusterIdentifier, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *HeliosSearchIndexedObjectsResponseBody) SetFiles(v []GlobalClusterIdentifier)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *HeliosSearchIndexedObjectsResponseBody) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetCassandraObjects

`func (o *HeliosSearchIndexedObjectsResponseBody) GetCassandraObjects() []GlobalClusterIdentifier`

GetCassandraObjects returns the CassandraObjects field if non-nil, zero value otherwise.

### GetCassandraObjectsOk

`func (o *HeliosSearchIndexedObjectsResponseBody) GetCassandraObjectsOk() (*[]GlobalClusterIdentifier, bool)`

GetCassandraObjectsOk returns a tuple with the CassandraObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraObjects

`func (o *HeliosSearchIndexedObjectsResponseBody) SetCassandraObjects(v []GlobalClusterIdentifier)`

SetCassandraObjects sets CassandraObjects field to given value.

### HasCassandraObjects

`func (o *HeliosSearchIndexedObjectsResponseBody) HasCassandraObjects() bool`

HasCassandraObjects returns a boolean if a field has been set.

### GetCouchbaseObjects

`func (o *HeliosSearchIndexedObjectsResponseBody) GetCouchbaseObjects() []GlobalClusterIdentifier`

GetCouchbaseObjects returns the CouchbaseObjects field if non-nil, zero value otherwise.

### GetCouchbaseObjectsOk

`func (o *HeliosSearchIndexedObjectsResponseBody) GetCouchbaseObjectsOk() (*[]GlobalClusterIdentifier, bool)`

GetCouchbaseObjectsOk returns a tuple with the CouchbaseObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouchbaseObjects

`func (o *HeliosSearchIndexedObjectsResponseBody) SetCouchbaseObjects(v []GlobalClusterIdentifier)`

SetCouchbaseObjects sets CouchbaseObjects field to given value.

### HasCouchbaseObjects

`func (o *HeliosSearchIndexedObjectsResponseBody) HasCouchbaseObjects() bool`

HasCouchbaseObjects returns a boolean if a field has been set.

### GetHbaseObjects

`func (o *HeliosSearchIndexedObjectsResponseBody) GetHbaseObjects() []GlobalClusterIdentifier`

GetHbaseObjects returns the HbaseObjects field if non-nil, zero value otherwise.

### GetHbaseObjectsOk

`func (o *HeliosSearchIndexedObjectsResponseBody) GetHbaseObjectsOk() (*[]GlobalClusterIdentifier, bool)`

GetHbaseObjectsOk returns a tuple with the HbaseObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseObjects

`func (o *HeliosSearchIndexedObjectsResponseBody) SetHbaseObjects(v []GlobalClusterIdentifier)`

SetHbaseObjects sets HbaseObjects field to given value.

### HasHbaseObjects

`func (o *HeliosSearchIndexedObjectsResponseBody) HasHbaseObjects() bool`

HasHbaseObjects returns a boolean if a field has been set.

### GetHiveObjects

`func (o *HeliosSearchIndexedObjectsResponseBody) GetHiveObjects() []GlobalClusterIdentifier`

GetHiveObjects returns the HiveObjects field if non-nil, zero value otherwise.

### GetHiveObjectsOk

`func (o *HeliosSearchIndexedObjectsResponseBody) GetHiveObjectsOk() (*[]GlobalClusterIdentifier, bool)`

GetHiveObjectsOk returns a tuple with the HiveObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveObjects

`func (o *HeliosSearchIndexedObjectsResponseBody) SetHiveObjects(v []GlobalClusterIdentifier)`

SetHiveObjects sets HiveObjects field to given value.

### HasHiveObjects

`func (o *HeliosSearchIndexedObjectsResponseBody) HasHiveObjects() bool`

HasHiveObjects returns a boolean if a field has been set.

### GetMongoObjects

`func (o *HeliosSearchIndexedObjectsResponseBody) GetMongoObjects() []GlobalClusterIdentifier`

GetMongoObjects returns the MongoObjects field if non-nil, zero value otherwise.

### GetMongoObjectsOk

`func (o *HeliosSearchIndexedObjectsResponseBody) GetMongoObjectsOk() (*[]GlobalClusterIdentifier, bool)`

GetMongoObjectsOk returns a tuple with the MongoObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoObjects

`func (o *HeliosSearchIndexedObjectsResponseBody) SetMongoObjects(v []GlobalClusterIdentifier)`

SetMongoObjects sets MongoObjects field to given value.

### HasMongoObjects

`func (o *HeliosSearchIndexedObjectsResponseBody) HasMongoObjects() bool`

HasMongoObjects returns a boolean if a field has been set.

### GetHdfsObjects

`func (o *HeliosSearchIndexedObjectsResponseBody) GetHdfsObjects() []GlobalClusterIdentifier`

GetHdfsObjects returns the HdfsObjects field if non-nil, zero value otherwise.

### GetHdfsObjectsOk

`func (o *HeliosSearchIndexedObjectsResponseBody) GetHdfsObjectsOk() (*[]GlobalClusterIdentifier, bool)`

GetHdfsObjectsOk returns a tuple with the HdfsObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsObjects

`func (o *HeliosSearchIndexedObjectsResponseBody) SetHdfsObjects(v []GlobalClusterIdentifier)`

SetHdfsObjects sets HdfsObjects field to given value.

### HasHdfsObjects

`func (o *HeliosSearchIndexedObjectsResponseBody) HasHdfsObjects() bool`

HasHdfsObjects returns a boolean if a field has been set.

### GetExchangeObjects

`func (o *HeliosSearchIndexedObjectsResponseBody) GetExchangeObjects() []GlobalClusterIdentifier`

GetExchangeObjects returns the ExchangeObjects field if non-nil, zero value otherwise.

### GetExchangeObjectsOk

`func (o *HeliosSearchIndexedObjectsResponseBody) GetExchangeObjectsOk() (*[]GlobalClusterIdentifier, bool)`

GetExchangeObjectsOk returns a tuple with the ExchangeObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeObjects

`func (o *HeliosSearchIndexedObjectsResponseBody) SetExchangeObjects(v []GlobalClusterIdentifier)`

SetExchangeObjects sets ExchangeObjects field to given value.

### HasExchangeObjects

`func (o *HeliosSearchIndexedObjectsResponseBody) HasExchangeObjects() bool`

HasExchangeObjects returns a boolean if a field has been set.

### GetPublicFolderItems

`func (o *HeliosSearchIndexedObjectsResponseBody) GetPublicFolderItems() []GlobalClusterIdentifier`

GetPublicFolderItems returns the PublicFolderItems field if non-nil, zero value otherwise.

### GetPublicFolderItemsOk

`func (o *HeliosSearchIndexedObjectsResponseBody) GetPublicFolderItemsOk() (*[]GlobalClusterIdentifier, bool)`

GetPublicFolderItemsOk returns a tuple with the PublicFolderItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicFolderItems

`func (o *HeliosSearchIndexedObjectsResponseBody) SetPublicFolderItems(v []GlobalClusterIdentifier)`

SetPublicFolderItems sets PublicFolderItems field to given value.

### HasPublicFolderItems

`func (o *HeliosSearchIndexedObjectsResponseBody) HasPublicFolderItems() bool`

HasPublicFolderItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


