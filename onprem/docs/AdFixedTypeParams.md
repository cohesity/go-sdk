# AdFixedTypeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **NullableInt64** | Specifies the fixed Unix UID, when mapping type is set to kFixed. | 
**Gid** | **NullableInt64** | Specifies the fixed Unix GID, when mapping type is set to kFixed. | 

## Methods

### NewAdFixedTypeParams

`func NewAdFixedTypeParams(uid NullableInt64, gid NullableInt64, ) *AdFixedTypeParams`

NewAdFixedTypeParams instantiates a new AdFixedTypeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdFixedTypeParamsWithDefaults

`func NewAdFixedTypeParamsWithDefaults() *AdFixedTypeParams`

NewAdFixedTypeParamsWithDefaults instantiates a new AdFixedTypeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *AdFixedTypeParams) GetUid() int64`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *AdFixedTypeParams) GetUidOk() (*int64, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *AdFixedTypeParams) SetUid(v int64)`

SetUid sets Uid field to given value.


### SetUidNil

`func (o *AdFixedTypeParams) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *AdFixedTypeParams) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil
### GetGid

`func (o *AdFixedTypeParams) GetGid() int64`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *AdFixedTypeParams) GetGidOk() (*int64, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *AdFixedTypeParams) SetGid(v int64)`

SetGid sets Gid field to given value.


### SetGidNil

`func (o *AdFixedTypeParams) SetGidNil(b bool)`

 SetGidNil sets the value for Gid to be an explicit nil

### UnsetGid
`func (o *AdFixedTypeParams) UnsetGid()`

UnsetGid ensures that no value is present for Gid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


