# RecoverCouchbaseNoSqlObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | **NullableString** | Specifies the fully qualified name of the object to be restored. | 
**RenameTo** | Pointer to **NullableString** | Specifies the new name to which the object should be renamed, at the time of recovery. | [optional] 

## Methods

### NewRecoverCouchbaseNoSqlObjectParams

`func NewRecoverCouchbaseNoSqlObjectParams(objectName NullableString, ) *RecoverCouchbaseNoSqlObjectParams`

NewRecoverCouchbaseNoSqlObjectParams instantiates a new RecoverCouchbaseNoSqlObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverCouchbaseNoSqlObjectParamsWithDefaults

`func NewRecoverCouchbaseNoSqlObjectParamsWithDefaults() *RecoverCouchbaseNoSqlObjectParams`

NewRecoverCouchbaseNoSqlObjectParamsWithDefaults instantiates a new RecoverCouchbaseNoSqlObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *RecoverCouchbaseNoSqlObjectParams) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *RecoverCouchbaseNoSqlObjectParams) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *RecoverCouchbaseNoSqlObjectParams) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.


### SetObjectNameNil

`func (o *RecoverCouchbaseNoSqlObjectParams) SetObjectNameNil(b bool)`

 SetObjectNameNil sets the value for ObjectName to be an explicit nil

### UnsetObjectName
`func (o *RecoverCouchbaseNoSqlObjectParams) UnsetObjectName()`

UnsetObjectName ensures that no value is present for ObjectName, not even an explicit nil
### GetRenameTo

`func (o *RecoverCouchbaseNoSqlObjectParams) GetRenameTo() string`

GetRenameTo returns the RenameTo field if non-nil, zero value otherwise.

### GetRenameToOk

`func (o *RecoverCouchbaseNoSqlObjectParams) GetRenameToOk() (*string, bool)`

GetRenameToOk returns a tuple with the RenameTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameTo

`func (o *RecoverCouchbaseNoSqlObjectParams) SetRenameTo(v string)`

SetRenameTo sets RenameTo field to given value.

### HasRenameTo

`func (o *RecoverCouchbaseNoSqlObjectParams) HasRenameTo() bool`

HasRenameTo returns a boolean if a field has been set.

### SetRenameToNil

`func (o *RecoverCouchbaseNoSqlObjectParams) SetRenameToNil(b bool)`

 SetRenameToNil sets the value for RenameTo to be an explicit nil

### UnsetRenameTo
`func (o *RecoverCouchbaseNoSqlObjectParams) UnsetRenameTo()`

UnsetRenameTo ensures that no value is present for RenameTo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


