# FallbackUserIdMappingParamsFixedTypeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gid** | **NullableInt64** | Specifies the fixed Unix GID, when mapping type is set to kFixed. | 
**Uid** | **NullableInt64** | Specifies the fixed Unix UID, when mapping type is set to kFixed. | 

## Methods

### NewFallbackUserIdMappingParamsFixedTypeParams

`func NewFallbackUserIdMappingParamsFixedTypeParams(gid NullableInt64, uid NullableInt64, ) *FallbackUserIdMappingParamsFixedTypeParams`

NewFallbackUserIdMappingParamsFixedTypeParams instantiates a new FallbackUserIdMappingParamsFixedTypeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFallbackUserIdMappingParamsFixedTypeParamsWithDefaults

`func NewFallbackUserIdMappingParamsFixedTypeParamsWithDefaults() *FallbackUserIdMappingParamsFixedTypeParams`

NewFallbackUserIdMappingParamsFixedTypeParamsWithDefaults instantiates a new FallbackUserIdMappingParamsFixedTypeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGid

`func (o *FallbackUserIdMappingParamsFixedTypeParams) GetGid() int64`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *FallbackUserIdMappingParamsFixedTypeParams) GetGidOk() (*int64, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *FallbackUserIdMappingParamsFixedTypeParams) SetGid(v int64)`

SetGid sets Gid field to given value.


### SetGidNil

`func (o *FallbackUserIdMappingParamsFixedTypeParams) SetGidNil(b bool)`

 SetGidNil sets the value for Gid to be an explicit nil

### UnsetGid
`func (o *FallbackUserIdMappingParamsFixedTypeParams) UnsetGid()`

UnsetGid ensures that no value is present for Gid, not even an explicit nil
### GetUid

`func (o *FallbackUserIdMappingParamsFixedTypeParams) GetUid() int64`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *FallbackUserIdMappingParamsFixedTypeParams) GetUidOk() (*int64, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *FallbackUserIdMappingParamsFixedTypeParams) SetUid(v int64)`

SetUid sets Uid field to given value.


### SetUidNil

`func (o *FallbackUserIdMappingParamsFixedTypeParams) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *FallbackUserIdMappingParamsFixedTypeParams) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


