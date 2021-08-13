# PhysicalSnapshotParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectionType** | Pointer to **NullableString** | Specifies the protection type of Physical snapshots. | [optional] 
**EnableSystemBackup** | Pointer to **NullableBool** | Specifies if system backup was enabled for the source in that particular run. | [optional] 

## Methods

### NewPhysicalSnapshotParams

`func NewPhysicalSnapshotParams() *PhysicalSnapshotParams`

NewPhysicalSnapshotParams instantiates a new PhysicalSnapshotParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalSnapshotParamsWithDefaults

`func NewPhysicalSnapshotParamsWithDefaults() *PhysicalSnapshotParams`

NewPhysicalSnapshotParamsWithDefaults instantiates a new PhysicalSnapshotParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtectionType

`func (o *PhysicalSnapshotParams) GetProtectionType() string`

GetProtectionType returns the ProtectionType field if non-nil, zero value otherwise.

### GetProtectionTypeOk

`func (o *PhysicalSnapshotParams) GetProtectionTypeOk() (*string, bool)`

GetProtectionTypeOk returns a tuple with the ProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionType

`func (o *PhysicalSnapshotParams) SetProtectionType(v string)`

SetProtectionType sets ProtectionType field to given value.

### HasProtectionType

`func (o *PhysicalSnapshotParams) HasProtectionType() bool`

HasProtectionType returns a boolean if a field has been set.

### SetProtectionTypeNil

`func (o *PhysicalSnapshotParams) SetProtectionTypeNil(b bool)`

 SetProtectionTypeNil sets the value for ProtectionType to be an explicit nil

### UnsetProtectionType
`func (o *PhysicalSnapshotParams) UnsetProtectionType()`

UnsetProtectionType ensures that no value is present for ProtectionType, not even an explicit nil
### GetEnableSystemBackup

`func (o *PhysicalSnapshotParams) GetEnableSystemBackup() bool`

GetEnableSystemBackup returns the EnableSystemBackup field if non-nil, zero value otherwise.

### GetEnableSystemBackupOk

`func (o *PhysicalSnapshotParams) GetEnableSystemBackupOk() (*bool, bool)`

GetEnableSystemBackupOk returns a tuple with the EnableSystemBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSystemBackup

`func (o *PhysicalSnapshotParams) SetEnableSystemBackup(v bool)`

SetEnableSystemBackup sets EnableSystemBackup field to given value.

### HasEnableSystemBackup

`func (o *PhysicalSnapshotParams) HasEnableSystemBackup() bool`

HasEnableSystemBackup returns a boolean if a field has been set.

### SetEnableSystemBackupNil

`func (o *PhysicalSnapshotParams) SetEnableSystemBackupNil(b bool)`

 SetEnableSystemBackupNil sets the value for EnableSystemBackup to be an explicit nil

### UnsetEnableSystemBackup
`func (o *PhysicalSnapshotParams) UnsetEnableSystemBackup()`

UnsetEnableSystemBackup ensures that no value is present for EnableSystemBackup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


