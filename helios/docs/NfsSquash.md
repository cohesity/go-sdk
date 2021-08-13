# NfsSquash

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gid** | Pointer to **NullableInt32** | GID mapped for all clients. | [optional] 
**Uid** | Pointer to **NullableInt32** | UID mapped for all clients. | [optional] 

## Methods

### NewNfsSquash

`func NewNfsSquash() *NfsSquash`

NewNfsSquash instantiates a new NfsSquash object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfsSquashWithDefaults

`func NewNfsSquashWithDefaults() *NfsSquash`

NewNfsSquashWithDefaults instantiates a new NfsSquash object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGid

`func (o *NfsSquash) GetGid() int32`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *NfsSquash) GetGidOk() (*int32, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *NfsSquash) SetGid(v int32)`

SetGid sets Gid field to given value.

### HasGid

`func (o *NfsSquash) HasGid() bool`

HasGid returns a boolean if a field has been set.

### SetGidNil

`func (o *NfsSquash) SetGidNil(b bool)`

 SetGidNil sets the value for Gid to be an explicit nil

### UnsetGid
`func (o *NfsSquash) UnsetGid()`

UnsetGid ensures that no value is present for Gid, not even an explicit nil
### GetUid

`func (o *NfsSquash) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *NfsSquash) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *NfsSquash) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *NfsSquash) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *NfsSquash) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *NfsSquash) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


