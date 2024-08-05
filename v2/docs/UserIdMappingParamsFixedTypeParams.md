# UserIdMappingParamsFixedTypeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gid** | **NullableInt64** | Specifies the fixed Unix GID, when mapping type is set to kFixed. | 
**Uid** | **NullableInt64** | Specifies the fixed Unix UID, when mapping type is set to kFixed. | 

## Methods

### NewUserIdMappingParamsFixedTypeParams

`func NewUserIdMappingParamsFixedTypeParams(gid NullableInt64, uid NullableInt64, ) *UserIdMappingParamsFixedTypeParams`

NewUserIdMappingParamsFixedTypeParams instantiates a new UserIdMappingParamsFixedTypeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserIdMappingParamsFixedTypeParamsWithDefaults

`func NewUserIdMappingParamsFixedTypeParamsWithDefaults() *UserIdMappingParamsFixedTypeParams`

NewUserIdMappingParamsFixedTypeParamsWithDefaults instantiates a new UserIdMappingParamsFixedTypeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGid

`func (o *UserIdMappingParamsFixedTypeParams) GetGid() int64`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *UserIdMappingParamsFixedTypeParams) GetGidOk() (*int64, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *UserIdMappingParamsFixedTypeParams) SetGid(v int64)`

SetGid sets Gid field to given value.


### SetGidNil

`func (o *UserIdMappingParamsFixedTypeParams) SetGidNil(b bool)`

 SetGidNil sets the value for Gid to be an explicit nil

### UnsetGid
`func (o *UserIdMappingParamsFixedTypeParams) UnsetGid()`

UnsetGid ensures that no value is present for Gid, not even an explicit nil
### GetUid

`func (o *UserIdMappingParamsFixedTypeParams) GetUid() int64`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *UserIdMappingParamsFixedTypeParams) GetUidOk() (*int64, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *UserIdMappingParamsFixedTypeParams) SetUid(v int64)`

SetUid sets Uid field to given value.


### SetUidNil

`func (o *UserIdMappingParamsFixedTypeParams) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *UserIdMappingParamsFixedTypeParams) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


