# SearchIndexedObjectsResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt32** | Specifies the total number of indexed objects that match the filter and search criteria. Use this value to determine how many additional requests are required to get the full result. | [optional] 
**ObjectType** | Pointer to **string** | Specifies the object type. | [optional] 
**PaginationCookie** | Pointer to **NullableString** | Specifies cookie for resuming search if pagination is being used. | [optional] 
**CassandraObjects** | Pointer to [**[]CassandraIndexedObject**](CassandraIndexedObject.md) | Specifies the indexed Cassandra objects. | [optional] 
**CouchbaseObjects** | Pointer to [**[]CouchbaseIndexedObject**](CouchbaseIndexedObject.md) | Specifies the indexed Couchbase objects. | [optional] 
**Emails** | Pointer to [**[]Email**](Email.md) | Specifies the indexed emails and email folders. | [optional] 
**ExchangeObjects** | Pointer to [**[]ExchangeIndexedObject**](ExchangeIndexedObject.md) | Specifies the indexed HDFS objects. | [optional] 
**Files** | Pointer to [**[]File**](File.md) | Specifies the indexed files. | [optional] 
**HbaseObjects** | Pointer to [**[]HbaseIndexedObject**](HbaseIndexedObject.md) | Specifies the indexed Hbase objects. | [optional] 
**HdfsObjects** | Pointer to [**[]HDFSIndexedObject**](HDFSIndexedObject.md) | Specifies the indexed HDFS objects. | [optional] 
**HiveObjects** | Pointer to [**[]HiveIndexedObject**](HiveIndexedObject.md) | Specifies the indexed Hive objects. | [optional] 
**MongoObjects** | Pointer to [**[]MongoIndexedObject**](MongoIndexedObject.md) | Specifies the indexed Mongo objects. | [optional] 
**MsGroupItems** | Pointer to [**[]MsGroupItem**](MsGroupItem.md) | Specifies the indexed M365 Groups items like group mail items, files etc. | [optional] 
**OneDriveItems** | Pointer to [**[]DocumentLibraryItem**](DocumentLibraryItem.md) | Specifies the indexed one drive items. | [optional] 
**PublicFolderItems** | Pointer to [**[]PublicFolderItem**](PublicFolderItem.md) | Specifies the indexed Public folder items. | [optional] 
**SfdcRecords** | Pointer to [**SfdcRecords**](SfdcRecords.md) |  | [optional] 
**SharepointItems** | Pointer to [**[]SharepointItem**](SharepointItem.md) | Specifies the indexed Sharepoint items. | [optional] 
**TeamsItems** | Pointer to [**[]TeamsItem**](TeamsItem.md) | Specifies the indexed M365 Teams items like channels, files etc. | [optional] 
**UdaObjects** | Pointer to [**[]UdaIndexedObject**](UdaIndexedObject.md) | Specifies the indexed Universal Data Adapter objects. | [optional] 

## Methods

### NewSearchIndexedObjectsResponseBody

`func NewSearchIndexedObjectsResponseBody() *SearchIndexedObjectsResponseBody`

NewSearchIndexedObjectsResponseBody instantiates a new SearchIndexedObjectsResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchIndexedObjectsResponseBodyWithDefaults

`func NewSearchIndexedObjectsResponseBodyWithDefaults() *SearchIndexedObjectsResponseBody`

NewSearchIndexedObjectsResponseBodyWithDefaults instantiates a new SearchIndexedObjectsResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *SearchIndexedObjectsResponseBody) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SearchIndexedObjectsResponseBody) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SearchIndexedObjectsResponseBody) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *SearchIndexedObjectsResponseBody) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *SearchIndexedObjectsResponseBody) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *SearchIndexedObjectsResponseBody) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetObjectType

`func (o *SearchIndexedObjectsResponseBody) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SearchIndexedObjectsResponseBody) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SearchIndexedObjectsResponseBody) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *SearchIndexedObjectsResponseBody) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetPaginationCookie

`func (o *SearchIndexedObjectsResponseBody) GetPaginationCookie() string`

GetPaginationCookie returns the PaginationCookie field if non-nil, zero value otherwise.

### GetPaginationCookieOk

`func (o *SearchIndexedObjectsResponseBody) GetPaginationCookieOk() (*string, bool)`

GetPaginationCookieOk returns a tuple with the PaginationCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationCookie

`func (o *SearchIndexedObjectsResponseBody) SetPaginationCookie(v string)`

SetPaginationCookie sets PaginationCookie field to given value.

### HasPaginationCookie

`func (o *SearchIndexedObjectsResponseBody) HasPaginationCookie() bool`

HasPaginationCookie returns a boolean if a field has been set.

### SetPaginationCookieNil

`func (o *SearchIndexedObjectsResponseBody) SetPaginationCookieNil(b bool)`

 SetPaginationCookieNil sets the value for PaginationCookie to be an explicit nil

### UnsetPaginationCookie
`func (o *SearchIndexedObjectsResponseBody) UnsetPaginationCookie()`

UnsetPaginationCookie ensures that no value is present for PaginationCookie, not even an explicit nil
### GetCassandraObjects

`func (o *SearchIndexedObjectsResponseBody) GetCassandraObjects() []CassandraIndexedObject`

GetCassandraObjects returns the CassandraObjects field if non-nil, zero value otherwise.

### GetCassandraObjectsOk

`func (o *SearchIndexedObjectsResponseBody) GetCassandraObjectsOk() (*[]CassandraIndexedObject, bool)`

GetCassandraObjectsOk returns a tuple with the CassandraObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraObjects

`func (o *SearchIndexedObjectsResponseBody) SetCassandraObjects(v []CassandraIndexedObject)`

SetCassandraObjects sets CassandraObjects field to given value.

### HasCassandraObjects

`func (o *SearchIndexedObjectsResponseBody) HasCassandraObjects() bool`

HasCassandraObjects returns a boolean if a field has been set.

### GetCouchbaseObjects

`func (o *SearchIndexedObjectsResponseBody) GetCouchbaseObjects() []CouchbaseIndexedObject`

GetCouchbaseObjects returns the CouchbaseObjects field if non-nil, zero value otherwise.

### GetCouchbaseObjectsOk

`func (o *SearchIndexedObjectsResponseBody) GetCouchbaseObjectsOk() (*[]CouchbaseIndexedObject, bool)`

GetCouchbaseObjectsOk returns a tuple with the CouchbaseObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouchbaseObjects

`func (o *SearchIndexedObjectsResponseBody) SetCouchbaseObjects(v []CouchbaseIndexedObject)`

SetCouchbaseObjects sets CouchbaseObjects field to given value.

### HasCouchbaseObjects

`func (o *SearchIndexedObjectsResponseBody) HasCouchbaseObjects() bool`

HasCouchbaseObjects returns a boolean if a field has been set.

### GetEmails

`func (o *SearchIndexedObjectsResponseBody) GetEmails() []Email`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *SearchIndexedObjectsResponseBody) GetEmailsOk() (*[]Email, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *SearchIndexedObjectsResponseBody) SetEmails(v []Email)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *SearchIndexedObjectsResponseBody) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetExchangeObjects

`func (o *SearchIndexedObjectsResponseBody) GetExchangeObjects() []ExchangeIndexedObject`

GetExchangeObjects returns the ExchangeObjects field if non-nil, zero value otherwise.

### GetExchangeObjectsOk

`func (o *SearchIndexedObjectsResponseBody) GetExchangeObjectsOk() (*[]ExchangeIndexedObject, bool)`

GetExchangeObjectsOk returns a tuple with the ExchangeObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeObjects

`func (o *SearchIndexedObjectsResponseBody) SetExchangeObjects(v []ExchangeIndexedObject)`

SetExchangeObjects sets ExchangeObjects field to given value.

### HasExchangeObjects

`func (o *SearchIndexedObjectsResponseBody) HasExchangeObjects() bool`

HasExchangeObjects returns a boolean if a field has been set.

### GetFiles

`func (o *SearchIndexedObjectsResponseBody) GetFiles() []File`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *SearchIndexedObjectsResponseBody) GetFilesOk() (*[]File, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *SearchIndexedObjectsResponseBody) SetFiles(v []File)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *SearchIndexedObjectsResponseBody) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetHbaseObjects

`func (o *SearchIndexedObjectsResponseBody) GetHbaseObjects() []HbaseIndexedObject`

GetHbaseObjects returns the HbaseObjects field if non-nil, zero value otherwise.

### GetHbaseObjectsOk

`func (o *SearchIndexedObjectsResponseBody) GetHbaseObjectsOk() (*[]HbaseIndexedObject, bool)`

GetHbaseObjectsOk returns a tuple with the HbaseObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseObjects

`func (o *SearchIndexedObjectsResponseBody) SetHbaseObjects(v []HbaseIndexedObject)`

SetHbaseObjects sets HbaseObjects field to given value.

### HasHbaseObjects

`func (o *SearchIndexedObjectsResponseBody) HasHbaseObjects() bool`

HasHbaseObjects returns a boolean if a field has been set.

### GetHdfsObjects

`func (o *SearchIndexedObjectsResponseBody) GetHdfsObjects() []HDFSIndexedObject`

GetHdfsObjects returns the HdfsObjects field if non-nil, zero value otherwise.

### GetHdfsObjectsOk

`func (o *SearchIndexedObjectsResponseBody) GetHdfsObjectsOk() (*[]HDFSIndexedObject, bool)`

GetHdfsObjectsOk returns a tuple with the HdfsObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsObjects

`func (o *SearchIndexedObjectsResponseBody) SetHdfsObjects(v []HDFSIndexedObject)`

SetHdfsObjects sets HdfsObjects field to given value.

### HasHdfsObjects

`func (o *SearchIndexedObjectsResponseBody) HasHdfsObjects() bool`

HasHdfsObjects returns a boolean if a field has been set.

### GetHiveObjects

`func (o *SearchIndexedObjectsResponseBody) GetHiveObjects() []HiveIndexedObject`

GetHiveObjects returns the HiveObjects field if non-nil, zero value otherwise.

### GetHiveObjectsOk

`func (o *SearchIndexedObjectsResponseBody) GetHiveObjectsOk() (*[]HiveIndexedObject, bool)`

GetHiveObjectsOk returns a tuple with the HiveObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveObjects

`func (o *SearchIndexedObjectsResponseBody) SetHiveObjects(v []HiveIndexedObject)`

SetHiveObjects sets HiveObjects field to given value.

### HasHiveObjects

`func (o *SearchIndexedObjectsResponseBody) HasHiveObjects() bool`

HasHiveObjects returns a boolean if a field has been set.

### GetMongoObjects

`func (o *SearchIndexedObjectsResponseBody) GetMongoObjects() []MongoIndexedObject`

GetMongoObjects returns the MongoObjects field if non-nil, zero value otherwise.

### GetMongoObjectsOk

`func (o *SearchIndexedObjectsResponseBody) GetMongoObjectsOk() (*[]MongoIndexedObject, bool)`

GetMongoObjectsOk returns a tuple with the MongoObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoObjects

`func (o *SearchIndexedObjectsResponseBody) SetMongoObjects(v []MongoIndexedObject)`

SetMongoObjects sets MongoObjects field to given value.

### HasMongoObjects

`func (o *SearchIndexedObjectsResponseBody) HasMongoObjects() bool`

HasMongoObjects returns a boolean if a field has been set.

### GetMsGroupItems

`func (o *SearchIndexedObjectsResponseBody) GetMsGroupItems() []MsGroupItem`

GetMsGroupItems returns the MsGroupItems field if non-nil, zero value otherwise.

### GetMsGroupItemsOk

`func (o *SearchIndexedObjectsResponseBody) GetMsGroupItemsOk() (*[]MsGroupItem, bool)`

GetMsGroupItemsOk returns a tuple with the MsGroupItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsGroupItems

`func (o *SearchIndexedObjectsResponseBody) SetMsGroupItems(v []MsGroupItem)`

SetMsGroupItems sets MsGroupItems field to given value.

### HasMsGroupItems

`func (o *SearchIndexedObjectsResponseBody) HasMsGroupItems() bool`

HasMsGroupItems returns a boolean if a field has been set.

### GetOneDriveItems

`func (o *SearchIndexedObjectsResponseBody) GetOneDriveItems() []DocumentLibraryItem`

GetOneDriveItems returns the OneDriveItems field if non-nil, zero value otherwise.

### GetOneDriveItemsOk

`func (o *SearchIndexedObjectsResponseBody) GetOneDriveItemsOk() (*[]DocumentLibraryItem, bool)`

GetOneDriveItemsOk returns a tuple with the OneDriveItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneDriveItems

`func (o *SearchIndexedObjectsResponseBody) SetOneDriveItems(v []DocumentLibraryItem)`

SetOneDriveItems sets OneDriveItems field to given value.

### HasOneDriveItems

`func (o *SearchIndexedObjectsResponseBody) HasOneDriveItems() bool`

HasOneDriveItems returns a boolean if a field has been set.

### GetPublicFolderItems

`func (o *SearchIndexedObjectsResponseBody) GetPublicFolderItems() []PublicFolderItem`

GetPublicFolderItems returns the PublicFolderItems field if non-nil, zero value otherwise.

### GetPublicFolderItemsOk

`func (o *SearchIndexedObjectsResponseBody) GetPublicFolderItemsOk() (*[]PublicFolderItem, bool)`

GetPublicFolderItemsOk returns a tuple with the PublicFolderItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicFolderItems

`func (o *SearchIndexedObjectsResponseBody) SetPublicFolderItems(v []PublicFolderItem)`

SetPublicFolderItems sets PublicFolderItems field to given value.

### HasPublicFolderItems

`func (o *SearchIndexedObjectsResponseBody) HasPublicFolderItems() bool`

HasPublicFolderItems returns a boolean if a field has been set.

### GetSfdcRecords

`func (o *SearchIndexedObjectsResponseBody) GetSfdcRecords() SfdcRecords`

GetSfdcRecords returns the SfdcRecords field if non-nil, zero value otherwise.

### GetSfdcRecordsOk

`func (o *SearchIndexedObjectsResponseBody) GetSfdcRecordsOk() (*SfdcRecords, bool)`

GetSfdcRecordsOk returns a tuple with the SfdcRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfdcRecords

`func (o *SearchIndexedObjectsResponseBody) SetSfdcRecords(v SfdcRecords)`

SetSfdcRecords sets SfdcRecords field to given value.

### HasSfdcRecords

`func (o *SearchIndexedObjectsResponseBody) HasSfdcRecords() bool`

HasSfdcRecords returns a boolean if a field has been set.

### GetSharepointItems

`func (o *SearchIndexedObjectsResponseBody) GetSharepointItems() []SharepointItem`

GetSharepointItems returns the SharepointItems field if non-nil, zero value otherwise.

### GetSharepointItemsOk

`func (o *SearchIndexedObjectsResponseBody) GetSharepointItemsOk() (*[]SharepointItem, bool)`

GetSharepointItemsOk returns a tuple with the SharepointItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharepointItems

`func (o *SearchIndexedObjectsResponseBody) SetSharepointItems(v []SharepointItem)`

SetSharepointItems sets SharepointItems field to given value.

### HasSharepointItems

`func (o *SearchIndexedObjectsResponseBody) HasSharepointItems() bool`

HasSharepointItems returns a boolean if a field has been set.

### GetTeamsItems

`func (o *SearchIndexedObjectsResponseBody) GetTeamsItems() []TeamsItem`

GetTeamsItems returns the TeamsItems field if non-nil, zero value otherwise.

### GetTeamsItemsOk

`func (o *SearchIndexedObjectsResponseBody) GetTeamsItemsOk() (*[]TeamsItem, bool)`

GetTeamsItemsOk returns a tuple with the TeamsItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsItems

`func (o *SearchIndexedObjectsResponseBody) SetTeamsItems(v []TeamsItem)`

SetTeamsItems sets TeamsItems field to given value.

### HasTeamsItems

`func (o *SearchIndexedObjectsResponseBody) HasTeamsItems() bool`

HasTeamsItems returns a boolean if a field has been set.

### GetUdaObjects

`func (o *SearchIndexedObjectsResponseBody) GetUdaObjects() []UdaIndexedObject`

GetUdaObjects returns the UdaObjects field if non-nil, zero value otherwise.

### GetUdaObjectsOk

`func (o *SearchIndexedObjectsResponseBody) GetUdaObjectsOk() (*[]UdaIndexedObject, bool)`

GetUdaObjectsOk returns a tuple with the UdaObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdaObjects

`func (o *SearchIndexedObjectsResponseBody) SetUdaObjects(v []UdaIndexedObject)`

SetUdaObjects sets UdaObjects field to given value.

### HasUdaObjects

`func (o *SearchIndexedObjectsResponseBody) HasUdaObjects() bool`

HasUdaObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


