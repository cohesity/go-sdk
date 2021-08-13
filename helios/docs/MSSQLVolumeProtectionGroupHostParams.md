# MSSQLVolumeProtectionGroupHostParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | **NullableInt64** | Specifies the id of the host container on which databases are hosted. | 
**HostName** | Pointer to **NullableString** | Specifies the name of the host container on which databases are hosted. | [optional] [readonly] 
**VolumeGuids** | Pointer to **[]string** | Specifies the list of volume GUIDs to be protected. If not specified, all the volumes of the host will be protected. Note that volumes of host on which databases are hosted are protected even if its not mentioned in this list. | [optional] 
**EnableSystemBackup** | Pointer to **NullableBool** | Specifies whether to enable system/bmr backup using 3rd party tools installed on agent host. | [optional] 

## Methods

### NewMSSQLVolumeProtectionGroupHostParams

`func NewMSSQLVolumeProtectionGroupHostParams(hostId NullableInt64, ) *MSSQLVolumeProtectionGroupHostParams`

NewMSSQLVolumeProtectionGroupHostParams instantiates a new MSSQLVolumeProtectionGroupHostParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMSSQLVolumeProtectionGroupHostParamsWithDefaults

`func NewMSSQLVolumeProtectionGroupHostParamsWithDefaults() *MSSQLVolumeProtectionGroupHostParams`

NewMSSQLVolumeProtectionGroupHostParamsWithDefaults instantiates a new MSSQLVolumeProtectionGroupHostParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *MSSQLVolumeProtectionGroupHostParams) GetHostId() int64`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *MSSQLVolumeProtectionGroupHostParams) GetHostIdOk() (*int64, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *MSSQLVolumeProtectionGroupHostParams) SetHostId(v int64)`

SetHostId sets HostId field to given value.


### SetHostIdNil

`func (o *MSSQLVolumeProtectionGroupHostParams) SetHostIdNil(b bool)`

 SetHostIdNil sets the value for HostId to be an explicit nil

### UnsetHostId
`func (o *MSSQLVolumeProtectionGroupHostParams) UnsetHostId()`

UnsetHostId ensures that no value is present for HostId, not even an explicit nil
### GetHostName

`func (o *MSSQLVolumeProtectionGroupHostParams) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *MSSQLVolumeProtectionGroupHostParams) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *MSSQLVolumeProtectionGroupHostParams) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *MSSQLVolumeProtectionGroupHostParams) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### SetHostNameNil

`func (o *MSSQLVolumeProtectionGroupHostParams) SetHostNameNil(b bool)`

 SetHostNameNil sets the value for HostName to be an explicit nil

### UnsetHostName
`func (o *MSSQLVolumeProtectionGroupHostParams) UnsetHostName()`

UnsetHostName ensures that no value is present for HostName, not even an explicit nil
### GetVolumeGuids

`func (o *MSSQLVolumeProtectionGroupHostParams) GetVolumeGuids() []string`

GetVolumeGuids returns the VolumeGuids field if non-nil, zero value otherwise.

### GetVolumeGuidsOk

`func (o *MSSQLVolumeProtectionGroupHostParams) GetVolumeGuidsOk() (*[]string, bool)`

GetVolumeGuidsOk returns a tuple with the VolumeGuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGuids

`func (o *MSSQLVolumeProtectionGroupHostParams) SetVolumeGuids(v []string)`

SetVolumeGuids sets VolumeGuids field to given value.

### HasVolumeGuids

`func (o *MSSQLVolumeProtectionGroupHostParams) HasVolumeGuids() bool`

HasVolumeGuids returns a boolean if a field has been set.

### SetVolumeGuidsNil

`func (o *MSSQLVolumeProtectionGroupHostParams) SetVolumeGuidsNil(b bool)`

 SetVolumeGuidsNil sets the value for VolumeGuids to be an explicit nil

### UnsetVolumeGuids
`func (o *MSSQLVolumeProtectionGroupHostParams) UnsetVolumeGuids()`

UnsetVolumeGuids ensures that no value is present for VolumeGuids, not even an explicit nil
### GetEnableSystemBackup

`func (o *MSSQLVolumeProtectionGroupHostParams) GetEnableSystemBackup() bool`

GetEnableSystemBackup returns the EnableSystemBackup field if non-nil, zero value otherwise.

### GetEnableSystemBackupOk

`func (o *MSSQLVolumeProtectionGroupHostParams) GetEnableSystemBackupOk() (*bool, bool)`

GetEnableSystemBackupOk returns a tuple with the EnableSystemBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSystemBackup

`func (o *MSSQLVolumeProtectionGroupHostParams) SetEnableSystemBackup(v bool)`

SetEnableSystemBackup sets EnableSystemBackup field to given value.

### HasEnableSystemBackup

`func (o *MSSQLVolumeProtectionGroupHostParams) HasEnableSystemBackup() bool`

HasEnableSystemBackup returns a boolean if a field has been set.

### SetEnableSystemBackupNil

`func (o *MSSQLVolumeProtectionGroupHostParams) SetEnableSystemBackupNil(b bool)`

 SetEnableSystemBackupNil sets the value for EnableSystemBackup to be an explicit nil

### UnsetEnableSystemBackup
`func (o *MSSQLVolumeProtectionGroupHostParams) UnsetEnableSystemBackup()`

UnsetEnableSystemBackup ensures that no value is present for EnableSystemBackup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


