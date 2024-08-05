# RecoverUdaObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectId** | Pointer to **NullableInt64** | Specifies the ID of the object. | [optional] 
**ObjectName** | Pointer to **NullableString** | Specifies the fully qualified name of the object to be restored. | [optional] 
**Overwrite** | Pointer to **NullableBool** | Set to true to overwrite an existing object at the destination. If set to false, and the same object exists at the destination, then recovery will fail for that object. | [optional] 
**RenameTo** | Pointer to **NullableString** | Specifies the new name to which the object should be renamed to after the recovery. | [optional] 

## Methods

### NewRecoverUdaObjectParams

`func NewRecoverUdaObjectParams() *RecoverUdaObjectParams`

NewRecoverUdaObjectParams instantiates a new RecoverUdaObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverUdaObjectParamsWithDefaults

`func NewRecoverUdaObjectParamsWithDefaults() *RecoverUdaObjectParams`

NewRecoverUdaObjectParamsWithDefaults instantiates a new RecoverUdaObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectId

`func (o *RecoverUdaObjectParams) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *RecoverUdaObjectParams) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *RecoverUdaObjectParams) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *RecoverUdaObjectParams) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### SetObjectIdNil

`func (o *RecoverUdaObjectParams) SetObjectIdNil(b bool)`

 SetObjectIdNil sets the value for ObjectId to be an explicit nil

### UnsetObjectId
`func (o *RecoverUdaObjectParams) UnsetObjectId()`

UnsetObjectId ensures that no value is present for ObjectId, not even an explicit nil
### GetObjectName

`func (o *RecoverUdaObjectParams) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *RecoverUdaObjectParams) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *RecoverUdaObjectParams) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *RecoverUdaObjectParams) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### SetObjectNameNil

`func (o *RecoverUdaObjectParams) SetObjectNameNil(b bool)`

 SetObjectNameNil sets the value for ObjectName to be an explicit nil

### UnsetObjectName
`func (o *RecoverUdaObjectParams) UnsetObjectName()`

UnsetObjectName ensures that no value is present for ObjectName, not even an explicit nil
### GetOverwrite

`func (o *RecoverUdaObjectParams) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *RecoverUdaObjectParams) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *RecoverUdaObjectParams) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *RecoverUdaObjectParams) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### SetOverwriteNil

`func (o *RecoverUdaObjectParams) SetOverwriteNil(b bool)`

 SetOverwriteNil sets the value for Overwrite to be an explicit nil

### UnsetOverwrite
`func (o *RecoverUdaObjectParams) UnsetOverwrite()`

UnsetOverwrite ensures that no value is present for Overwrite, not even an explicit nil
### GetRenameTo

`func (o *RecoverUdaObjectParams) GetRenameTo() string`

GetRenameTo returns the RenameTo field if non-nil, zero value otherwise.

### GetRenameToOk

`func (o *RecoverUdaObjectParams) GetRenameToOk() (*string, bool)`

GetRenameToOk returns a tuple with the RenameTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameTo

`func (o *RecoverUdaObjectParams) SetRenameTo(v string)`

SetRenameTo sets RenameTo field to given value.

### HasRenameTo

`func (o *RecoverUdaObjectParams) HasRenameTo() bool`

HasRenameTo returns a boolean if a field has been set.

### SetRenameToNil

`func (o *RecoverUdaObjectParams) SetRenameToNil(b bool)`

 SetRenameToNil sets the value for RenameTo to be an explicit nil

### UnsetRenameTo
`func (o *RecoverUdaObjectParams) UnsetRenameTo()`

UnsetRenameTo ensures that no value is present for RenameTo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


