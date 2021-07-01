# FixedUnixIdMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gid** | Pointer to **NullableInt64** | Specifies the fixed Unix GID, when mapping type is set to kFixed. | [optional] 
**Uid** | Pointer to **NullableInt64** | Specifies the fixed Unix UID, when mapping type is set to kFixed. | [optional] 

## Methods

### NewFixedUnixIdMapping

`func NewFixedUnixIdMapping() *FixedUnixIdMapping`

NewFixedUnixIdMapping instantiates a new FixedUnixIdMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedUnixIdMappingWithDefaults

`func NewFixedUnixIdMappingWithDefaults() *FixedUnixIdMapping`

NewFixedUnixIdMappingWithDefaults instantiates a new FixedUnixIdMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGid

`func (o *FixedUnixIdMapping) GetGid() int64`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *FixedUnixIdMapping) GetGidOk() (*int64, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *FixedUnixIdMapping) SetGid(v int64)`

SetGid sets Gid field to given value.

### HasGid

`func (o *FixedUnixIdMapping) HasGid() bool`

HasGid returns a boolean if a field has been set.

### SetGidNil

`func (o *FixedUnixIdMapping) SetGidNil(b bool)`

 SetGidNil sets the value for Gid to be an explicit nil

### UnsetGid
`func (o *FixedUnixIdMapping) UnsetGid()`

UnsetGid ensures that no value is present for Gid, not even an explicit nil
### GetUid

`func (o *FixedUnixIdMapping) GetUid() int64`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *FixedUnixIdMapping) GetUidOk() (*int64, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *FixedUnixIdMapping) SetUid(v int64)`

SetUid sets Uid field to given value.

### HasUid

`func (o *FixedUnixIdMapping) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *FixedUnixIdMapping) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *FixedUnixIdMapping) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


