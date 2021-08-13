# SearchIndexedObjectsResponseBodyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | Pointer to [**[]Email**](Email.md) | Specifies the indexed emails and email folders. | [optional] 
**Files** | Pointer to [**[]File**](File.md) | Specifies the indexed files. | [optional] 
**CassandraObjects** | Pointer to [**[]CassandraIndexedObject**](CassandraIndexedObject.md) | Specifies the indexed Cassandra objects. | [optional] 
**CouchbaseObjects** | Pointer to [**[]CouchbaseIndexedObject**](CouchbaseIndexedObject.md) | Specifies the indexed Couchbase objects. | [optional] 
**HbaseObjects** | Pointer to [**[]HbaseIndexedObject**](HbaseIndexedObject.md) | Specifies the indexed Hbase objects. | [optional] 
**HiveObjects** | Pointer to [**[]HiveIndexedObject**](HiveIndexedObject.md) | Specifies the indexed Hive objects. | [optional] 
**MongoObjects** | Pointer to [**[]MongoIndexedObject**](MongoIndexedObject.md) | Specifies the indexed Mongo objects. | [optional] 
**HdfsObjects** | Pointer to [**[]HDFSIndexedObject**](HDFSIndexedObject.md) | Specifies the indexed HDFS objects. | [optional] 
**ExchangeObjects** | Pointer to [**[]ExchangeIndexedObject**](ExchangeIndexedObject.md) | Specifies the indexed HDFS objects. | [optional] 
**PublicFolderItems** | Pointer to [**[]PublicFolderItem**](PublicFolderItem.md) | Specifies the indexed Public folder items. | [optional] 
**SharepointItems** | Pointer to [**[]SharepointItem**](SharepointItem.md) | Specifies the indexed Sharepoint items. | [optional] 

## Methods

### NewSearchIndexedObjectsResponseBodyAllOf

`func NewSearchIndexedObjectsResponseBodyAllOf() *SearchIndexedObjectsResponseBodyAllOf`

NewSearchIndexedObjectsResponseBodyAllOf instantiates a new SearchIndexedObjectsResponseBodyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchIndexedObjectsResponseBodyAllOfWithDefaults

`func NewSearchIndexedObjectsResponseBodyAllOfWithDefaults() *SearchIndexedObjectsResponseBodyAllOf`

NewSearchIndexedObjectsResponseBodyAllOfWithDefaults instantiates a new SearchIndexedObjectsResponseBodyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *SearchIndexedObjectsResponseBodyAllOf) GetEmails() []Email`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *SearchIndexedObjectsResponseBodyAllOf) GetEmailsOk() (*[]Email, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *SearchIndexedObjectsResponseBodyAllOf) SetEmails(v []Email)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *SearchIndexedObjectsResponseBodyAllOf) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetFiles

`func (o *SearchIndexedObjectsResponseBodyAllOf) GetFiles() []File`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *SearchIndexedObjectsResponseBodyAllOf) GetFilesOk() (*[]File, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *SearchIndexedObjectsResponseBodyAllOf) SetFiles(v []File)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *SearchIndexedObjectsResponseBodyAllOf) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetCassandraObjects

`func (o *SearchIndexedObjectsResponseBodyAllOf) GetCassandraObjects() []CassandraIndexedObject`

GetCassandraObjects returns the CassandraObjects field if non-nil, zero value otherwise.

### GetCassandraObjectsOk

`func (o *SearchIndexedObjectsResponseBodyAllOf) GetCassandraObjectsOk() (*[]CassandraIndexedObject, bool)`

GetCassandraObjectsOk returns a tuple with the CassandraObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraObjects

`func (o *SearchIndexedObjectsResponseBodyAllOf) SetCassandraObjects(v []CassandraIndexedObject)`

SetCassandraObjects sets CassandraObjects field to given value.

### HasCassandraObjects

`func (o *SearchIndexedObjectsResponseBodyAllOf) HasCassandraObjects() bool`

HasCassandraObjects returns a boolean if a field has been set.

### GetCouchbaseObjects

`func (o *SearchIndexedObjectsResponseBodyAllOf) GetCouchbaseObjects() []CouchbaseIndexedObject`

GetCouchbaseObjects returns the CouchbaseObjects field if non-nil, zero value otherwise.

### GetCouchbaseObjectsOk

`func (o *SearchIndexedObjectsResponseBodyAllOf) GetCouchbaseObjectsOk() (*[]CouchbaseIndexedObject, bool)`

GetCouchbaseObjectsOk returns a tuple with the CouchbaseObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouchbaseObjects

`func (o *SearchIndexedObjectsResponseBodyAllOf) SetCouchbaseObjects(v []CouchbaseIndexedObject)`

SetCouchbaseObjects sets CouchbaseObjects field to given value.

### HasCouchbaseObjects

`func (o *SearchIndexedObjectsResponseBodyAllOf) HasCouchbaseObjects() bool`

HasCouchbaseObjects returns a boolean if a field has been set.

### GetHbaseObjects

`func (o *SearchIndexedObjectsResponseBodyAllOf) GetHbaseObjects() []HbaseIndexedObject`

GetHbaseObjects returns the HbaseObjects field if non-nil, zero value otherwise.

### GetHbaseObjectsOk

`func (o *SearchIndexedObjectsResponseBodyAllOf) GetHbaseObjectsOk() (*[]HbaseIndexedObject, bool)`

GetHbaseObjectsOk returns a tuple with the HbaseObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseObjects

`func (o *SearchIndexedObjectsResponseBodyAllOf) SetHbaseObjects(v []HbaseIndexedObject)`

SetHbaseObjects sets HbaseObjects field to given value.

### HasHbaseObjects

`func (o *SearchIndexedObjectsResponseBodyAllOf) HasHbaseObjects() bool`

HasHbaseObjects returns a boolean if a field has been set.

### GetHiveObjects

`func (o *SearchIndexedObjectsResponseBodyAllOf) GetHiveObjects() []HiveIndexedObject`

GetHiveObjects returns the HiveObjects field if non-nil, zero value otherwise.

### GetHiveObjectsOk

`func (o *SearchIndexedObjectsResponseBodyAllOf) GetHiveObjectsOk() (*[]HiveIndexedObject, bool)`

GetHiveObjectsOk returns a tuple with the HiveObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveObjects

`func (o *SearchIndexedObjectsResponseBodyAllOf) SetHiveObjects(v []HiveIndexedObject)`

SetHiveObjects sets HiveObjects field to given value.

### HasHiveObjects

`func (o *SearchIndexedObjectsResponseBodyAllOf) HasHiveObjects() bool`

HasHiveObjects returns a boolean if a field has been set.

### GetMongoObjects

`func (o *SearchIndexedObjectsResponseBodyAllOf) GetMongoObjects() []MongoIndexedObject`

GetMongoObjects returns the MongoObjects field if non-nil, zero value otherwise.

### GetMongoObjectsOk

`func (o *SearchIndexedObjectsResponseBodyAllOf) GetMongoObjectsOk() (*[]MongoIndexedObject, bool)`

GetMongoObjectsOk returns a tuple with the MongoObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoObjects

`func (o *SearchIndexedObjectsResponseBodyAllOf) SetMongoObjects(v []MongoIndexedObject)`

SetMongoObjects sets MongoObjects field to given value.

### HasMongoObjects

`func (o *SearchIndexedObjectsResponseBodyAllOf) HasMongoObjects() bool`

HasMongoObjects returns a boolean if a field has been set.

### GetHdfsObjects

`func (o *SearchIndexedObjectsResponseBodyAllOf) GetHdfsObjects() []HDFSIndexedObject`

GetHdfsObjects returns the HdfsObjects field if non-nil, zero value otherwise.

### GetHdfsObjectsOk

`func (o *SearchIndexedObjectsResponseBodyAllOf) GetHdfsObjectsOk() (*[]HDFSIndexedObject, bool)`

GetHdfsObjectsOk returns a tuple with the HdfsObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsObjects

`func (o *SearchIndexedObjectsResponseBodyAllOf) SetHdfsObjects(v []HDFSIndexedObject)`

SetHdfsObjects sets HdfsObjects field to given value.

### HasHdfsObjects

`func (o *SearchIndexedObjectsResponseBodyAllOf) HasHdfsObjects() bool`

HasHdfsObjects returns a boolean if a field has been set.

### GetExchangeObjects

`func (o *SearchIndexedObjectsResponseBodyAllOf) GetExchangeObjects() []ExchangeIndexedObject`

GetExchangeObjects returns the ExchangeObjects field if non-nil, zero value otherwise.

### GetExchangeObjectsOk

`func (o *SearchIndexedObjectsResponseBodyAllOf) GetExchangeObjectsOk() (*[]ExchangeIndexedObject, bool)`

GetExchangeObjectsOk returns a tuple with the ExchangeObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeObjects

`func (o *SearchIndexedObjectsResponseBodyAllOf) SetExchangeObjects(v []ExchangeIndexedObject)`

SetExchangeObjects sets ExchangeObjects field to given value.

### HasExchangeObjects

`func (o *SearchIndexedObjectsResponseBodyAllOf) HasExchangeObjects() bool`

HasExchangeObjects returns a boolean if a field has been set.

### GetPublicFolderItems

`func (o *SearchIndexedObjectsResponseBodyAllOf) GetPublicFolderItems() []PublicFolderItem`

GetPublicFolderItems returns the PublicFolderItems field if non-nil, zero value otherwise.

### GetPublicFolderItemsOk

`func (o *SearchIndexedObjectsResponseBodyAllOf) GetPublicFolderItemsOk() (*[]PublicFolderItem, bool)`

GetPublicFolderItemsOk returns a tuple with the PublicFolderItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicFolderItems

`func (o *SearchIndexedObjectsResponseBodyAllOf) SetPublicFolderItems(v []PublicFolderItem)`

SetPublicFolderItems sets PublicFolderItems field to given value.

### HasPublicFolderItems

`func (o *SearchIndexedObjectsResponseBodyAllOf) HasPublicFolderItems() bool`

HasPublicFolderItems returns a boolean if a field has been set.

### GetSharepointItems

`func (o *SearchIndexedObjectsResponseBodyAllOf) GetSharepointItems() []SharepointItem`

GetSharepointItems returns the SharepointItems field if non-nil, zero value otherwise.

### GetSharepointItemsOk

`func (o *SearchIndexedObjectsResponseBodyAllOf) GetSharepointItemsOk() (*[]SharepointItem, bool)`

GetSharepointItemsOk returns a tuple with the SharepointItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharepointItems

`func (o *SearchIndexedObjectsResponseBodyAllOf) SetSharepointItems(v []SharepointItem)`

SetSharepointItems sets SharepointItems field to given value.

### HasSharepointItems

`func (o *SearchIndexedObjectsResponseBodyAllOf) HasSharepointItems() bool`

HasSharepointItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


