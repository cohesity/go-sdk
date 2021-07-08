# SmbPermissionsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerSid** | Pointer to **NullableString** | Specifies the security identifier (SID) of the owner of the SMB share. | [optional] 
**Permissions** | Pointer to [**[]SmbPermission**](SmbPermission.md) | Array of SMB Permissions.  Specifies a list of SMB permissions. | [optional] 

## Methods

### NewSmbPermissionsInfo

`func NewSmbPermissionsInfo() *SmbPermissionsInfo`

NewSmbPermissionsInfo instantiates a new SmbPermissionsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmbPermissionsInfoWithDefaults

`func NewSmbPermissionsInfoWithDefaults() *SmbPermissionsInfo`

NewSmbPermissionsInfoWithDefaults instantiates a new SmbPermissionsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerSid

`func (o *SmbPermissionsInfo) GetOwnerSid() string`

GetOwnerSid returns the OwnerSid field if non-nil, zero value otherwise.

### GetOwnerSidOk

`func (o *SmbPermissionsInfo) GetOwnerSidOk() (*string, bool)`

GetOwnerSidOk returns a tuple with the OwnerSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerSid

`func (o *SmbPermissionsInfo) SetOwnerSid(v string)`

SetOwnerSid sets OwnerSid field to given value.

### HasOwnerSid

`func (o *SmbPermissionsInfo) HasOwnerSid() bool`

HasOwnerSid returns a boolean if a field has been set.

### SetOwnerSidNil

`func (o *SmbPermissionsInfo) SetOwnerSidNil(b bool)`

 SetOwnerSidNil sets the value for OwnerSid to be an explicit nil

### UnsetOwnerSid
`func (o *SmbPermissionsInfo) UnsetOwnerSid()`

UnsetOwnerSid ensures that no value is present for OwnerSid, not even an explicit nil
### GetPermissions

`func (o *SmbPermissionsInfo) GetPermissions() []SmbPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *SmbPermissionsInfo) GetPermissionsOk() (*[]SmbPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *SmbPermissionsInfo) SetPermissions(v []SmbPermission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *SmbPermissionsInfo) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *SmbPermissionsInfo) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *SmbPermissionsInfo) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


