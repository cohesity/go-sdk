# RecoverExperimentalAdapterObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectId** | Pointer to **NullableInt64** | Specifies the ID of the object. | [optional] 
**Overwrite** | Pointer to **NullableBool** | Set to true to overwrite an existing object at the destination. If set to false, and the same object exists at the destination, then recovery will fail for that object. | [optional] 
**RenameTo** | Pointer to **NullableString** | Specifies the new name to which the object should be renamed to after the recovery. | [optional] 

## Methods

### NewRecoverExperimentalAdapterObjectParams

`func NewRecoverExperimentalAdapterObjectParams() *RecoverExperimentalAdapterObjectParams`

NewRecoverExperimentalAdapterObjectParams instantiates a new RecoverExperimentalAdapterObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverExperimentalAdapterObjectParamsWithDefaults

`func NewRecoverExperimentalAdapterObjectParamsWithDefaults() *RecoverExperimentalAdapterObjectParams`

NewRecoverExperimentalAdapterObjectParamsWithDefaults instantiates a new RecoverExperimentalAdapterObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectId

`func (o *RecoverExperimentalAdapterObjectParams) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *RecoverExperimentalAdapterObjectParams) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *RecoverExperimentalAdapterObjectParams) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *RecoverExperimentalAdapterObjectParams) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### SetObjectIdNil

`func (o *RecoverExperimentalAdapterObjectParams) SetObjectIdNil(b bool)`

 SetObjectIdNil sets the value for ObjectId to be an explicit nil

### UnsetObjectId
`func (o *RecoverExperimentalAdapterObjectParams) UnsetObjectId()`

UnsetObjectId ensures that no value is present for ObjectId, not even an explicit nil
### GetOverwrite

`func (o *RecoverExperimentalAdapterObjectParams) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *RecoverExperimentalAdapterObjectParams) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *RecoverExperimentalAdapterObjectParams) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *RecoverExperimentalAdapterObjectParams) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### SetOverwriteNil

`func (o *RecoverExperimentalAdapterObjectParams) SetOverwriteNil(b bool)`

 SetOverwriteNil sets the value for Overwrite to be an explicit nil

### UnsetOverwrite
`func (o *RecoverExperimentalAdapterObjectParams) UnsetOverwrite()`

UnsetOverwrite ensures that no value is present for Overwrite, not even an explicit nil
### GetRenameTo

`func (o *RecoverExperimentalAdapterObjectParams) GetRenameTo() string`

GetRenameTo returns the RenameTo field if non-nil, zero value otherwise.

### GetRenameToOk

`func (o *RecoverExperimentalAdapterObjectParams) GetRenameToOk() (*string, bool)`

GetRenameToOk returns a tuple with the RenameTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameTo

`func (o *RecoverExperimentalAdapterObjectParams) SetRenameTo(v string)`

SetRenameTo sets RenameTo field to given value.

### HasRenameTo

`func (o *RecoverExperimentalAdapterObjectParams) HasRenameTo() bool`

HasRenameTo returns a boolean if a field has been set.

### SetRenameToNil

`func (o *RecoverExperimentalAdapterObjectParams) SetRenameToNil(b bool)`

 SetRenameToNil sets the value for RenameTo to be an explicit nil

### UnsetRenameTo
`func (o *RecoverExperimentalAdapterObjectParams) UnsetRenameTo()`

UnsetRenameTo ensures that no value is present for RenameTo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


