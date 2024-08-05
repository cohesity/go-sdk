# LockRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsExclusive** | Pointer to **NullableBool** | Specifies if entity lock is exclusive. | [optional] 
**Length** | Pointer to **NullableInt32** | Specifies the length of an entity lock. | [optional] 
**Offset** | Pointer to **NullableInt32** | Specifies the offset of an entity lock. | [optional] 

## Methods

### NewLockRange

`func NewLockRange() *LockRange`

NewLockRange instantiates a new LockRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLockRangeWithDefaults

`func NewLockRangeWithDefaults() *LockRange`

NewLockRangeWithDefaults instantiates a new LockRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsExclusive

`func (o *LockRange) GetIsExclusive() bool`

GetIsExclusive returns the IsExclusive field if non-nil, zero value otherwise.

### GetIsExclusiveOk

`func (o *LockRange) GetIsExclusiveOk() (*bool, bool)`

GetIsExclusiveOk returns a tuple with the IsExclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExclusive

`func (o *LockRange) SetIsExclusive(v bool)`

SetIsExclusive sets IsExclusive field to given value.

### HasIsExclusive

`func (o *LockRange) HasIsExclusive() bool`

HasIsExclusive returns a boolean if a field has been set.

### SetIsExclusiveNil

`func (o *LockRange) SetIsExclusiveNil(b bool)`

 SetIsExclusiveNil sets the value for IsExclusive to be an explicit nil

### UnsetIsExclusive
`func (o *LockRange) UnsetIsExclusive()`

UnsetIsExclusive ensures that no value is present for IsExclusive, not even an explicit nil
### GetLength

`func (o *LockRange) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *LockRange) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *LockRange) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *LockRange) HasLength() bool`

HasLength returns a boolean if a field has been set.

### SetLengthNil

`func (o *LockRange) SetLengthNil(b bool)`

 SetLengthNil sets the value for Length to be an explicit nil

### UnsetLength
`func (o *LockRange) UnsetLength()`

UnsetLength ensures that no value is present for Length, not even an explicit nil
### GetOffset

`func (o *LockRange) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *LockRange) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *LockRange) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *LockRange) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *LockRange) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *LockRange) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


