# PhysicalVolumeProtectionGroupObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableSystemBackup** | Pointer to **NullableBool** | Specifies whether or not to take a system backup. Applicable only for windows sources. | [optional] 
**ExcludedVssWriters** | Pointer to **[]string** | Specifies writer names which should be excluded from physical volume based backups for a given source. | [optional] 
**Id** | **int64** | Specifies the ID of the object protected. | 
**Name** | Pointer to **NullableString** | Specifies the name of the object protected. | [optional] [readonly] 
**VolumeGuids** | Pointer to **[]string** | Specifies the list of GUIDs of volumes protected. If empty, then all volumes will be protected by default. | [optional] 

## Methods

### NewPhysicalVolumeProtectionGroupObjectParams

`func NewPhysicalVolumeProtectionGroupObjectParams(id int64, ) *PhysicalVolumeProtectionGroupObjectParams`

NewPhysicalVolumeProtectionGroupObjectParams instantiates a new PhysicalVolumeProtectionGroupObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalVolumeProtectionGroupObjectParamsWithDefaults

`func NewPhysicalVolumeProtectionGroupObjectParamsWithDefaults() *PhysicalVolumeProtectionGroupObjectParams`

NewPhysicalVolumeProtectionGroupObjectParamsWithDefaults instantiates a new PhysicalVolumeProtectionGroupObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableSystemBackup

`func (o *PhysicalVolumeProtectionGroupObjectParams) GetEnableSystemBackup() bool`

GetEnableSystemBackup returns the EnableSystemBackup field if non-nil, zero value otherwise.

### GetEnableSystemBackupOk

`func (o *PhysicalVolumeProtectionGroupObjectParams) GetEnableSystemBackupOk() (*bool, bool)`

GetEnableSystemBackupOk returns a tuple with the EnableSystemBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSystemBackup

`func (o *PhysicalVolumeProtectionGroupObjectParams) SetEnableSystemBackup(v bool)`

SetEnableSystemBackup sets EnableSystemBackup field to given value.

### HasEnableSystemBackup

`func (o *PhysicalVolumeProtectionGroupObjectParams) HasEnableSystemBackup() bool`

HasEnableSystemBackup returns a boolean if a field has been set.

### SetEnableSystemBackupNil

`func (o *PhysicalVolumeProtectionGroupObjectParams) SetEnableSystemBackupNil(b bool)`

 SetEnableSystemBackupNil sets the value for EnableSystemBackup to be an explicit nil

### UnsetEnableSystemBackup
`func (o *PhysicalVolumeProtectionGroupObjectParams) UnsetEnableSystemBackup()`

UnsetEnableSystemBackup ensures that no value is present for EnableSystemBackup, not even an explicit nil
### GetExcludedVssWriters

`func (o *PhysicalVolumeProtectionGroupObjectParams) GetExcludedVssWriters() []string`

GetExcludedVssWriters returns the ExcludedVssWriters field if non-nil, zero value otherwise.

### GetExcludedVssWritersOk

`func (o *PhysicalVolumeProtectionGroupObjectParams) GetExcludedVssWritersOk() (*[]string, bool)`

GetExcludedVssWritersOk returns a tuple with the ExcludedVssWriters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedVssWriters

`func (o *PhysicalVolumeProtectionGroupObjectParams) SetExcludedVssWriters(v []string)`

SetExcludedVssWriters sets ExcludedVssWriters field to given value.

### HasExcludedVssWriters

`func (o *PhysicalVolumeProtectionGroupObjectParams) HasExcludedVssWriters() bool`

HasExcludedVssWriters returns a boolean if a field has been set.

### SetExcludedVssWritersNil

`func (o *PhysicalVolumeProtectionGroupObjectParams) SetExcludedVssWritersNil(b bool)`

 SetExcludedVssWritersNil sets the value for ExcludedVssWriters to be an explicit nil

### UnsetExcludedVssWriters
`func (o *PhysicalVolumeProtectionGroupObjectParams) UnsetExcludedVssWriters()`

UnsetExcludedVssWriters ensures that no value is present for ExcludedVssWriters, not even an explicit nil
### GetId

`func (o *PhysicalVolumeProtectionGroupObjectParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PhysicalVolumeProtectionGroupObjectParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PhysicalVolumeProtectionGroupObjectParams) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *PhysicalVolumeProtectionGroupObjectParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PhysicalVolumeProtectionGroupObjectParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PhysicalVolumeProtectionGroupObjectParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PhysicalVolumeProtectionGroupObjectParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PhysicalVolumeProtectionGroupObjectParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PhysicalVolumeProtectionGroupObjectParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVolumeGuids

`func (o *PhysicalVolumeProtectionGroupObjectParams) GetVolumeGuids() []string`

GetVolumeGuids returns the VolumeGuids field if non-nil, zero value otherwise.

### GetVolumeGuidsOk

`func (o *PhysicalVolumeProtectionGroupObjectParams) GetVolumeGuidsOk() (*[]string, bool)`

GetVolumeGuidsOk returns a tuple with the VolumeGuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGuids

`func (o *PhysicalVolumeProtectionGroupObjectParams) SetVolumeGuids(v []string)`

SetVolumeGuids sets VolumeGuids field to given value.

### HasVolumeGuids

`func (o *PhysicalVolumeProtectionGroupObjectParams) HasVolumeGuids() bool`

HasVolumeGuids returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


