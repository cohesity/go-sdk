# HeliosSearchIndexedObjectsResponseBodyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewHeliosSearchIndexedObjectsResponseBodyAllOf

`func NewHeliosSearchIndexedObjectsResponseBodyAllOf() *HeliosSearchIndexedObjectsResponseBodyAllOf`

NewHeliosSearchIndexedObjectsResponseBodyAllOf instantiates a new HeliosSearchIndexedObjectsResponseBodyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosSearchIndexedObjectsResponseBodyAllOfWithDefaults

`func NewHeliosSearchIndexedObjectsResponseBodyAllOfWithDefaults() *HeliosSearchIndexedObjectsResponseBodyAllOf`

NewHeliosSearchIndexedObjectsResponseBodyAllOfWithDefaults instantiates a new HeliosSearchIndexedObjectsResponseBodyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) GetEmails() []GlobalClusterIdentifier`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) GetEmailsOk() (*[]GlobalClusterIdentifier, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) SetEmails(v []GlobalClusterIdentifier)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetFiles

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) GetFiles() []GlobalClusterIdentifier`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) GetFilesOk() (*[]GlobalClusterIdentifier, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) SetFiles(v []GlobalClusterIdentifier)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetCassandraObjects

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) GetCassandraObjects() []GlobalClusterIdentifier`

GetCassandraObjects returns the CassandraObjects field if non-nil, zero value otherwise.

### GetCassandraObjectsOk

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) GetCassandraObjectsOk() (*[]GlobalClusterIdentifier, bool)`

GetCassandraObjectsOk returns a tuple with the CassandraObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraObjects

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) SetCassandraObjects(v []GlobalClusterIdentifier)`

SetCassandraObjects sets CassandraObjects field to given value.

### HasCassandraObjects

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) HasCassandraObjects() bool`

HasCassandraObjects returns a boolean if a field has been set.

### GetCouchbaseObjects

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) GetCouchbaseObjects() []GlobalClusterIdentifier`

GetCouchbaseObjects returns the CouchbaseObjects field if non-nil, zero value otherwise.

### GetCouchbaseObjectsOk

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) GetCouchbaseObjectsOk() (*[]GlobalClusterIdentifier, bool)`

GetCouchbaseObjectsOk returns a tuple with the CouchbaseObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouchbaseObjects

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) SetCouchbaseObjects(v []GlobalClusterIdentifier)`

SetCouchbaseObjects sets CouchbaseObjects field to given value.

### HasCouchbaseObjects

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) HasCouchbaseObjects() bool`

HasCouchbaseObjects returns a boolean if a field has been set.

### GetHbaseObjects

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) GetHbaseObjects() []GlobalClusterIdentifier`

GetHbaseObjects returns the HbaseObjects field if non-nil, zero value otherwise.

### GetHbaseObjectsOk

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) GetHbaseObjectsOk() (*[]GlobalClusterIdentifier, bool)`

GetHbaseObjectsOk returns a tuple with the HbaseObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseObjects

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) SetHbaseObjects(v []GlobalClusterIdentifier)`

SetHbaseObjects sets HbaseObjects field to given value.

### HasHbaseObjects

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) HasHbaseObjects() bool`

HasHbaseObjects returns a boolean if a field has been set.

### GetHiveObjects

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) GetHiveObjects() []GlobalClusterIdentifier`

GetHiveObjects returns the HiveObjects field if non-nil, zero value otherwise.

### GetHiveObjectsOk

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) GetHiveObjectsOk() (*[]GlobalClusterIdentifier, bool)`

GetHiveObjectsOk returns a tuple with the HiveObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveObjects

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) SetHiveObjects(v []GlobalClusterIdentifier)`

SetHiveObjects sets HiveObjects field to given value.

### HasHiveObjects

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) HasHiveObjects() bool`

HasHiveObjects returns a boolean if a field has been set.

### GetMongoObjects

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) GetMongoObjects() []GlobalClusterIdentifier`

GetMongoObjects returns the MongoObjects field if non-nil, zero value otherwise.

### GetMongoObjectsOk

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) GetMongoObjectsOk() (*[]GlobalClusterIdentifier, bool)`

GetMongoObjectsOk returns a tuple with the MongoObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoObjects

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) SetMongoObjects(v []GlobalClusterIdentifier)`

SetMongoObjects sets MongoObjects field to given value.

### HasMongoObjects

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) HasMongoObjects() bool`

HasMongoObjects returns a boolean if a field has been set.

### GetHdfsObjects

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) GetHdfsObjects() []GlobalClusterIdentifier`

GetHdfsObjects returns the HdfsObjects field if non-nil, zero value otherwise.

### GetHdfsObjectsOk

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) GetHdfsObjectsOk() (*[]GlobalClusterIdentifier, bool)`

GetHdfsObjectsOk returns a tuple with the HdfsObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsObjects

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) SetHdfsObjects(v []GlobalClusterIdentifier)`

SetHdfsObjects sets HdfsObjects field to given value.

### HasHdfsObjects

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) HasHdfsObjects() bool`

HasHdfsObjects returns a boolean if a field has been set.

### GetExchangeObjects

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) GetExchangeObjects() []GlobalClusterIdentifier`

GetExchangeObjects returns the ExchangeObjects field if non-nil, zero value otherwise.

### GetExchangeObjectsOk

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) GetExchangeObjectsOk() (*[]GlobalClusterIdentifier, bool)`

GetExchangeObjectsOk returns a tuple with the ExchangeObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeObjects

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) SetExchangeObjects(v []GlobalClusterIdentifier)`

SetExchangeObjects sets ExchangeObjects field to given value.

### HasExchangeObjects

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) HasExchangeObjects() bool`

HasExchangeObjects returns a boolean if a field has been set.

### GetPublicFolderItems

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) GetPublicFolderItems() []GlobalClusterIdentifier`

GetPublicFolderItems returns the PublicFolderItems field if non-nil, zero value otherwise.

### GetPublicFolderItemsOk

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) GetPublicFolderItemsOk() (*[]GlobalClusterIdentifier, bool)`

GetPublicFolderItemsOk returns a tuple with the PublicFolderItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicFolderItems

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) SetPublicFolderItems(v []GlobalClusterIdentifier)`

SetPublicFolderItems sets PublicFolderItems field to given value.

### HasPublicFolderItems

`func (o *HeliosSearchIndexedObjectsResponseBodyAllOf) HasPublicFolderItems() bool`

HasPublicFolderItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


