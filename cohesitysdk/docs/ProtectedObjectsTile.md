# ProtectedObjectsTile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectsProtected** | Pointer to [**[]ProtectedObjectsByEnv**](ProtectedObjectsByEnv.md) | Protected Objects breakdown by object type. | [optional] 
**ProtectedCount** | Pointer to **NullableInt32** | Number of Protected Objects. | [optional] 
**ProtectedSizeBytes** | Pointer to **NullableInt64** | Size of Protected Objects. | [optional] 
**UnprotectedCount** | Pointer to **NullableInt32** | Number of Unprotected Objects. | [optional] 
**UnprotectedSizeBytes** | Pointer to **NullableInt64** | Size of Unprotected Objects. | [optional] 

## Methods

### NewProtectedObjectsTile

`func NewProtectedObjectsTile() *ProtectedObjectsTile`

NewProtectedObjectsTile instantiates a new ProtectedObjectsTile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedObjectsTileWithDefaults

`func NewProtectedObjectsTileWithDefaults() *ProtectedObjectsTile`

NewProtectedObjectsTileWithDefaults instantiates a new ProtectedObjectsTile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectsProtected

`func (o *ProtectedObjectsTile) GetObjectsProtected() []ProtectedObjectsByEnv`

GetObjectsProtected returns the ObjectsProtected field if non-nil, zero value otherwise.

### GetObjectsProtectedOk

`func (o *ProtectedObjectsTile) GetObjectsProtectedOk() (*[]ProtectedObjectsByEnv, bool)`

GetObjectsProtectedOk returns a tuple with the ObjectsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectsProtected

`func (o *ProtectedObjectsTile) SetObjectsProtected(v []ProtectedObjectsByEnv)`

SetObjectsProtected sets ObjectsProtected field to given value.

### HasObjectsProtected

`func (o *ProtectedObjectsTile) HasObjectsProtected() bool`

HasObjectsProtected returns a boolean if a field has been set.

### SetObjectsProtectedNil

`func (o *ProtectedObjectsTile) SetObjectsProtectedNil(b bool)`

 SetObjectsProtectedNil sets the value for ObjectsProtected to be an explicit nil

### UnsetObjectsProtected
`func (o *ProtectedObjectsTile) UnsetObjectsProtected()`

UnsetObjectsProtected ensures that no value is present for ObjectsProtected, not even an explicit nil
### GetProtectedCount

`func (o *ProtectedObjectsTile) GetProtectedCount() int32`

GetProtectedCount returns the ProtectedCount field if non-nil, zero value otherwise.

### GetProtectedCountOk

`func (o *ProtectedObjectsTile) GetProtectedCountOk() (*int32, bool)`

GetProtectedCountOk returns a tuple with the ProtectedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedCount

`func (o *ProtectedObjectsTile) SetProtectedCount(v int32)`

SetProtectedCount sets ProtectedCount field to given value.

### HasProtectedCount

`func (o *ProtectedObjectsTile) HasProtectedCount() bool`

HasProtectedCount returns a boolean if a field has been set.

### SetProtectedCountNil

`func (o *ProtectedObjectsTile) SetProtectedCountNil(b bool)`

 SetProtectedCountNil sets the value for ProtectedCount to be an explicit nil

### UnsetProtectedCount
`func (o *ProtectedObjectsTile) UnsetProtectedCount()`

UnsetProtectedCount ensures that no value is present for ProtectedCount, not even an explicit nil
### GetProtectedSizeBytes

`func (o *ProtectedObjectsTile) GetProtectedSizeBytes() int64`

GetProtectedSizeBytes returns the ProtectedSizeBytes field if non-nil, zero value otherwise.

### GetProtectedSizeBytesOk

`func (o *ProtectedObjectsTile) GetProtectedSizeBytesOk() (*int64, bool)`

GetProtectedSizeBytesOk returns a tuple with the ProtectedSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedSizeBytes

`func (o *ProtectedObjectsTile) SetProtectedSizeBytes(v int64)`

SetProtectedSizeBytes sets ProtectedSizeBytes field to given value.

### HasProtectedSizeBytes

`func (o *ProtectedObjectsTile) HasProtectedSizeBytes() bool`

HasProtectedSizeBytes returns a boolean if a field has been set.

### SetProtectedSizeBytesNil

`func (o *ProtectedObjectsTile) SetProtectedSizeBytesNil(b bool)`

 SetProtectedSizeBytesNil sets the value for ProtectedSizeBytes to be an explicit nil

### UnsetProtectedSizeBytes
`func (o *ProtectedObjectsTile) UnsetProtectedSizeBytes()`

UnsetProtectedSizeBytes ensures that no value is present for ProtectedSizeBytes, not even an explicit nil
### GetUnprotectedCount

`func (o *ProtectedObjectsTile) GetUnprotectedCount() int32`

GetUnprotectedCount returns the UnprotectedCount field if non-nil, zero value otherwise.

### GetUnprotectedCountOk

`func (o *ProtectedObjectsTile) GetUnprotectedCountOk() (*int32, bool)`

GetUnprotectedCountOk returns a tuple with the UnprotectedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnprotectedCount

`func (o *ProtectedObjectsTile) SetUnprotectedCount(v int32)`

SetUnprotectedCount sets UnprotectedCount field to given value.

### HasUnprotectedCount

`func (o *ProtectedObjectsTile) HasUnprotectedCount() bool`

HasUnprotectedCount returns a boolean if a field has been set.

### SetUnprotectedCountNil

`func (o *ProtectedObjectsTile) SetUnprotectedCountNil(b bool)`

 SetUnprotectedCountNil sets the value for UnprotectedCount to be an explicit nil

### UnsetUnprotectedCount
`func (o *ProtectedObjectsTile) UnsetUnprotectedCount()`

UnsetUnprotectedCount ensures that no value is present for UnprotectedCount, not even an explicit nil
### GetUnprotectedSizeBytes

`func (o *ProtectedObjectsTile) GetUnprotectedSizeBytes() int64`

GetUnprotectedSizeBytes returns the UnprotectedSizeBytes field if non-nil, zero value otherwise.

### GetUnprotectedSizeBytesOk

`func (o *ProtectedObjectsTile) GetUnprotectedSizeBytesOk() (*int64, bool)`

GetUnprotectedSizeBytesOk returns a tuple with the UnprotectedSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnprotectedSizeBytes

`func (o *ProtectedObjectsTile) SetUnprotectedSizeBytes(v int64)`

SetUnprotectedSizeBytes sets UnprotectedSizeBytes field to given value.

### HasUnprotectedSizeBytes

`func (o *ProtectedObjectsTile) HasUnprotectedSizeBytes() bool`

HasUnprotectedSizeBytes returns a boolean if a field has been set.

### SetUnprotectedSizeBytesNil

`func (o *ProtectedObjectsTile) SetUnprotectedSizeBytesNil(b bool)`

 SetUnprotectedSizeBytesNil sets the value for UnprotectedSizeBytes to be an explicit nil

### UnsetUnprotectedSizeBytes
`func (o *ProtectedObjectsTile) UnsetUnprotectedSizeBytes()`

UnsetUnprotectedSizeBytes ensures that no value is present for UnprotectedSizeBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


