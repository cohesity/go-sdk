# RestoreCountByObjectType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectCount** | Pointer to **NullableInt32** | Specifies the number of restores of the object type. | [optional] 
**ObjectType** | Pointer to **NullableString** | Specifies the type of the restored object. | [optional] 

## Methods

### NewRestoreCountByObjectType

`func NewRestoreCountByObjectType() *RestoreCountByObjectType`

NewRestoreCountByObjectType instantiates a new RestoreCountByObjectType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreCountByObjectTypeWithDefaults

`func NewRestoreCountByObjectTypeWithDefaults() *RestoreCountByObjectType`

NewRestoreCountByObjectTypeWithDefaults instantiates a new RestoreCountByObjectType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectCount

`func (o *RestoreCountByObjectType) GetObjectCount() int32`

GetObjectCount returns the ObjectCount field if non-nil, zero value otherwise.

### GetObjectCountOk

`func (o *RestoreCountByObjectType) GetObjectCountOk() (*int32, bool)`

GetObjectCountOk returns a tuple with the ObjectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectCount

`func (o *RestoreCountByObjectType) SetObjectCount(v int32)`

SetObjectCount sets ObjectCount field to given value.

### HasObjectCount

`func (o *RestoreCountByObjectType) HasObjectCount() bool`

HasObjectCount returns a boolean if a field has been set.

### SetObjectCountNil

`func (o *RestoreCountByObjectType) SetObjectCountNil(b bool)`

 SetObjectCountNil sets the value for ObjectCount to be an explicit nil

### UnsetObjectCount
`func (o *RestoreCountByObjectType) UnsetObjectCount()`

UnsetObjectCount ensures that no value is present for ObjectCount, not even an explicit nil
### GetObjectType

`func (o *RestoreCountByObjectType) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RestoreCountByObjectType) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RestoreCountByObjectType) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *RestoreCountByObjectType) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### SetObjectTypeNil

`func (o *RestoreCountByObjectType) SetObjectTypeNil(b bool)`

 SetObjectTypeNil sets the value for ObjectType to be an explicit nil

### UnsetObjectType
`func (o *RestoreCountByObjectType) UnsetObjectType()`

UnsetObjectType ensures that no value is present for ObjectType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


