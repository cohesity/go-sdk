# RecoverCassandraNoSqlObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | **NullableString** | Specifies the fully qualified name of the object to be restored. | 
**RenameTo** | Pointer to **NullableString** | Specifies the new name to which the object should be renamed to after the recovery. | [optional] 
**ObjectProperties** | Pointer to [**[]NoSqlObjectProperty**](NoSqlObjectProperty.md) | Specifies the properties to be applied to the object at the time of recovery. | [optional] 

## Methods

### NewRecoverCassandraNoSqlObjectParams

`func NewRecoverCassandraNoSqlObjectParams(objectName NullableString, ) *RecoverCassandraNoSqlObjectParams`

NewRecoverCassandraNoSqlObjectParams instantiates a new RecoverCassandraNoSqlObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverCassandraNoSqlObjectParamsWithDefaults

`func NewRecoverCassandraNoSqlObjectParamsWithDefaults() *RecoverCassandraNoSqlObjectParams`

NewRecoverCassandraNoSqlObjectParamsWithDefaults instantiates a new RecoverCassandraNoSqlObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *RecoverCassandraNoSqlObjectParams) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *RecoverCassandraNoSqlObjectParams) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *RecoverCassandraNoSqlObjectParams) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.


### SetObjectNameNil

`func (o *RecoverCassandraNoSqlObjectParams) SetObjectNameNil(b bool)`

 SetObjectNameNil sets the value for ObjectName to be an explicit nil

### UnsetObjectName
`func (o *RecoverCassandraNoSqlObjectParams) UnsetObjectName()`

UnsetObjectName ensures that no value is present for ObjectName, not even an explicit nil
### GetRenameTo

`func (o *RecoverCassandraNoSqlObjectParams) GetRenameTo() string`

GetRenameTo returns the RenameTo field if non-nil, zero value otherwise.

### GetRenameToOk

`func (o *RecoverCassandraNoSqlObjectParams) GetRenameToOk() (*string, bool)`

GetRenameToOk returns a tuple with the RenameTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameTo

`func (o *RecoverCassandraNoSqlObjectParams) SetRenameTo(v string)`

SetRenameTo sets RenameTo field to given value.

### HasRenameTo

`func (o *RecoverCassandraNoSqlObjectParams) HasRenameTo() bool`

HasRenameTo returns a boolean if a field has been set.

### SetRenameToNil

`func (o *RecoverCassandraNoSqlObjectParams) SetRenameToNil(b bool)`

 SetRenameToNil sets the value for RenameTo to be an explicit nil

### UnsetRenameTo
`func (o *RecoverCassandraNoSqlObjectParams) UnsetRenameTo()`

UnsetRenameTo ensures that no value is present for RenameTo, not even an explicit nil
### GetObjectProperties

`func (o *RecoverCassandraNoSqlObjectParams) GetObjectProperties() []NoSqlObjectProperty`

GetObjectProperties returns the ObjectProperties field if non-nil, zero value otherwise.

### GetObjectPropertiesOk

`func (o *RecoverCassandraNoSqlObjectParams) GetObjectPropertiesOk() (*[]NoSqlObjectProperty, bool)`

GetObjectPropertiesOk returns a tuple with the ObjectProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectProperties

`func (o *RecoverCassandraNoSqlObjectParams) SetObjectProperties(v []NoSqlObjectProperty)`

SetObjectProperties sets ObjectProperties field to given value.

### HasObjectProperties

`func (o *RecoverCassandraNoSqlObjectParams) HasObjectProperties() bool`

HasObjectProperties returns a boolean if a field has been set.

### SetObjectPropertiesNil

`func (o *RecoverCassandraNoSqlObjectParams) SetObjectPropertiesNil(b bool)`

 SetObjectPropertiesNil sets the value for ObjectProperties to be an explicit nil

### UnsetObjectProperties
`func (o *RecoverCassandraNoSqlObjectParams) UnsetObjectProperties()`

UnsetObjectProperties ensures that no value is present for ObjectProperties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


