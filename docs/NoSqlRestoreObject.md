# NoSqlRestoreObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectRestorePropertiesMap** | Pointer to [**[]NoSqlRestoreObjectObjectRestorePropertiesMapEntry**](NoSqlRestoreObjectObjectRestorePropertiesMapEntry.md) | Key-Value pair for properties to apply on restore object. Ex. Compaction for cassandra or ShardKeyJson for Mongo. | [optional] 
**ObjectUuid** | Pointer to **NullableString** | Uuid of the object to be restored. | [optional] 
**Rename** | Pointer to **NullableString** | The new name of the object after restore. | [optional] 

## Methods

### NewNoSqlRestoreObject

`func NewNoSqlRestoreObject() *NoSqlRestoreObject`

NewNoSqlRestoreObject instantiates a new NoSqlRestoreObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoSqlRestoreObjectWithDefaults

`func NewNoSqlRestoreObjectWithDefaults() *NoSqlRestoreObject`

NewNoSqlRestoreObjectWithDefaults instantiates a new NoSqlRestoreObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectRestorePropertiesMap

`func (o *NoSqlRestoreObject) GetObjectRestorePropertiesMap() []NoSqlRestoreObjectObjectRestorePropertiesMapEntry`

GetObjectRestorePropertiesMap returns the ObjectRestorePropertiesMap field if non-nil, zero value otherwise.

### GetObjectRestorePropertiesMapOk

`func (o *NoSqlRestoreObject) GetObjectRestorePropertiesMapOk() (*[]NoSqlRestoreObjectObjectRestorePropertiesMapEntry, bool)`

GetObjectRestorePropertiesMapOk returns a tuple with the ObjectRestorePropertiesMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectRestorePropertiesMap

`func (o *NoSqlRestoreObject) SetObjectRestorePropertiesMap(v []NoSqlRestoreObjectObjectRestorePropertiesMapEntry)`

SetObjectRestorePropertiesMap sets ObjectRestorePropertiesMap field to given value.

### HasObjectRestorePropertiesMap

`func (o *NoSqlRestoreObject) HasObjectRestorePropertiesMap() bool`

HasObjectRestorePropertiesMap returns a boolean if a field has been set.

### SetObjectRestorePropertiesMapNil

`func (o *NoSqlRestoreObject) SetObjectRestorePropertiesMapNil(b bool)`

 SetObjectRestorePropertiesMapNil sets the value for ObjectRestorePropertiesMap to be an explicit nil

### UnsetObjectRestorePropertiesMap
`func (o *NoSqlRestoreObject) UnsetObjectRestorePropertiesMap()`

UnsetObjectRestorePropertiesMap ensures that no value is present for ObjectRestorePropertiesMap, not even an explicit nil
### GetObjectUuid

`func (o *NoSqlRestoreObject) GetObjectUuid() string`

GetObjectUuid returns the ObjectUuid field if non-nil, zero value otherwise.

### GetObjectUuidOk

`func (o *NoSqlRestoreObject) GetObjectUuidOk() (*string, bool)`

GetObjectUuidOk returns a tuple with the ObjectUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectUuid

`func (o *NoSqlRestoreObject) SetObjectUuid(v string)`

SetObjectUuid sets ObjectUuid field to given value.

### HasObjectUuid

`func (o *NoSqlRestoreObject) HasObjectUuid() bool`

HasObjectUuid returns a boolean if a field has been set.

### SetObjectUuidNil

`func (o *NoSqlRestoreObject) SetObjectUuidNil(b bool)`

 SetObjectUuidNil sets the value for ObjectUuid to be an explicit nil

### UnsetObjectUuid
`func (o *NoSqlRestoreObject) UnsetObjectUuid()`

UnsetObjectUuid ensures that no value is present for ObjectUuid, not even an explicit nil
### GetRename

`func (o *NoSqlRestoreObject) GetRename() string`

GetRename returns the Rename field if non-nil, zero value otherwise.

### GetRenameOk

`func (o *NoSqlRestoreObject) GetRenameOk() (*string, bool)`

GetRenameOk returns a tuple with the Rename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRename

`func (o *NoSqlRestoreObject) SetRename(v string)`

SetRename sets Rename field to given value.

### HasRename

`func (o *NoSqlRestoreObject) HasRename() bool`

HasRename returns a boolean if a field has been set.

### SetRenameNil

`func (o *NoSqlRestoreObject) SetRenameNil(b bool)`

 SetRenameNil sets the value for Rename to be an explicit nil

### UnsetRename
`func (o *NoSqlRestoreObject) UnsetRename()`

UnsetRename ensures that no value is present for Rename, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


