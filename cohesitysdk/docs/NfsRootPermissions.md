# NfsRootPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gid** | Pointer to **NullableInt32** | Unix GID for the root of the file system. | [optional] 
**Mode** | Pointer to **NullableInt32** | Unix mode bits for the root of the file system. | [optional] 
**Uid** | Pointer to **NullableInt32** | Unix UID for the root of the file system. | [optional] 

## Methods

### NewNfsRootPermissions

`func NewNfsRootPermissions() *NfsRootPermissions`

NewNfsRootPermissions instantiates a new NfsRootPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfsRootPermissionsWithDefaults

`func NewNfsRootPermissionsWithDefaults() *NfsRootPermissions`

NewNfsRootPermissionsWithDefaults instantiates a new NfsRootPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGid

`func (o *NfsRootPermissions) GetGid() int32`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *NfsRootPermissions) GetGidOk() (*int32, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *NfsRootPermissions) SetGid(v int32)`

SetGid sets Gid field to given value.

### HasGid

`func (o *NfsRootPermissions) HasGid() bool`

HasGid returns a boolean if a field has been set.

### SetGidNil

`func (o *NfsRootPermissions) SetGidNil(b bool)`

 SetGidNil sets the value for Gid to be an explicit nil

### UnsetGid
`func (o *NfsRootPermissions) UnsetGid()`

UnsetGid ensures that no value is present for Gid, not even an explicit nil
### GetMode

`func (o *NfsRootPermissions) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *NfsRootPermissions) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *NfsRootPermissions) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *NfsRootPermissions) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *NfsRootPermissions) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *NfsRootPermissions) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetUid

`func (o *NfsRootPermissions) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *NfsRootPermissions) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *NfsRootPermissions) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *NfsRootPermissions) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *NfsRootPermissions) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *NfsRootPermissions) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


