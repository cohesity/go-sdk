# VolumeSecurityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **NullableInt32** | Specifies the Unix group ID for this volume. 0 indicates the root id. | [optional] 
**Permissions** | Pointer to **NullableString** | Specifies the Unix permission bits in octal string format. | [optional] 
**Style** | Pointer to **NullableString** | Specifies the security style associated with this volume. Specifies the type of a NetApp Volume. &#39;kUnix&#39; indicates Unix-style security. &#39;kNtfs&#39; indicates Windows NTFS-style security. &#39;kMixed&#39; indicates mixed-style security. &#39;kUnified&#39; indicates Unified-style security. | [optional] 
**UserId** | Pointer to **NullableInt32** | Specifies the Unix user id for this volume. 0 indicates the root id. | [optional] 

## Methods

### NewVolumeSecurityInfo

`func NewVolumeSecurityInfo() *VolumeSecurityInfo`

NewVolumeSecurityInfo instantiates a new VolumeSecurityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeSecurityInfoWithDefaults

`func NewVolumeSecurityInfoWithDefaults() *VolumeSecurityInfo`

NewVolumeSecurityInfoWithDefaults instantiates a new VolumeSecurityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *VolumeSecurityInfo) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *VolumeSecurityInfo) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *VolumeSecurityInfo) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *VolumeSecurityInfo) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *VolumeSecurityInfo) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *VolumeSecurityInfo) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetPermissions

`func (o *VolumeSecurityInfo) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *VolumeSecurityInfo) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *VolumeSecurityInfo) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *VolumeSecurityInfo) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *VolumeSecurityInfo) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *VolumeSecurityInfo) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetStyle

`func (o *VolumeSecurityInfo) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *VolumeSecurityInfo) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *VolumeSecurityInfo) SetStyle(v string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *VolumeSecurityInfo) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### SetStyleNil

`func (o *VolumeSecurityInfo) SetStyleNil(b bool)`

 SetStyleNil sets the value for Style to be an explicit nil

### UnsetStyle
`func (o *VolumeSecurityInfo) UnsetStyle()`

UnsetStyle ensures that no value is present for Style, not even an explicit nil
### GetUserId

`func (o *VolumeSecurityInfo) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *VolumeSecurityInfo) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *VolumeSecurityInfo) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *VolumeSecurityInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *VolumeSecurityInfo) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *VolumeSecurityInfo) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


