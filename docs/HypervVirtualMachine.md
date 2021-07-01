# HypervVirtualMachine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsHighlyAvailable** | Pointer to **NullableBool** | Specifies whether the VM is Highly Available or not. | [optional] 
**Version** | Pointer to **NullableString** | Specifies the version of the VM. For example, 8.0, 5.0 etc. | [optional] 
**VmBackupStatus** | Pointer to **NullableString** | Specifies the status of the VM for backup purpose. overrideDescription: true Specifies the backup status of a HyperV Virtual Machine object. &#39;kSupported&#39; indicates the agent on the VM can do backup. &#39;kUnsupportedConfig&#39; indicates the agent on the VM cannot do backup. &#39;kMissing&#39; indicates the VM is not found in SCVMM. | [optional] 
**VmBackupType** | Pointer to **NullableString** | Specifies the type of backup supported by the VM. overrideDescription: true Specifies the type of an HyperV datastore object. &#39;kRctBackup&#39; indicates backup is done using RCT/checkpoints. &#39;kVssBackup&#39; indicates backup is done using VSS. | [optional] 

## Methods

### NewHypervVirtualMachine

`func NewHypervVirtualMachine() *HypervVirtualMachine`

NewHypervVirtualMachine instantiates a new HypervVirtualMachine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervVirtualMachineWithDefaults

`func NewHypervVirtualMachineWithDefaults() *HypervVirtualMachine`

NewHypervVirtualMachineWithDefaults instantiates a new HypervVirtualMachine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsHighlyAvailable

`func (o *HypervVirtualMachine) GetIsHighlyAvailable() bool`

GetIsHighlyAvailable returns the IsHighlyAvailable field if non-nil, zero value otherwise.

### GetIsHighlyAvailableOk

`func (o *HypervVirtualMachine) GetIsHighlyAvailableOk() (*bool, bool)`

GetIsHighlyAvailableOk returns a tuple with the IsHighlyAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHighlyAvailable

`func (o *HypervVirtualMachine) SetIsHighlyAvailable(v bool)`

SetIsHighlyAvailable sets IsHighlyAvailable field to given value.

### HasIsHighlyAvailable

`func (o *HypervVirtualMachine) HasIsHighlyAvailable() bool`

HasIsHighlyAvailable returns a boolean if a field has been set.

### SetIsHighlyAvailableNil

`func (o *HypervVirtualMachine) SetIsHighlyAvailableNil(b bool)`

 SetIsHighlyAvailableNil sets the value for IsHighlyAvailable to be an explicit nil

### UnsetIsHighlyAvailable
`func (o *HypervVirtualMachine) UnsetIsHighlyAvailable()`

UnsetIsHighlyAvailable ensures that no value is present for IsHighlyAvailable, not even an explicit nil
### GetVersion

`func (o *HypervVirtualMachine) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HypervVirtualMachine) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HypervVirtualMachine) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HypervVirtualMachine) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *HypervVirtualMachine) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *HypervVirtualMachine) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetVmBackupStatus

`func (o *HypervVirtualMachine) GetVmBackupStatus() string`

GetVmBackupStatus returns the VmBackupStatus field if non-nil, zero value otherwise.

### GetVmBackupStatusOk

`func (o *HypervVirtualMachine) GetVmBackupStatusOk() (*string, bool)`

GetVmBackupStatusOk returns a tuple with the VmBackupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmBackupStatus

`func (o *HypervVirtualMachine) SetVmBackupStatus(v string)`

SetVmBackupStatus sets VmBackupStatus field to given value.

### HasVmBackupStatus

`func (o *HypervVirtualMachine) HasVmBackupStatus() bool`

HasVmBackupStatus returns a boolean if a field has been set.

### SetVmBackupStatusNil

`func (o *HypervVirtualMachine) SetVmBackupStatusNil(b bool)`

 SetVmBackupStatusNil sets the value for VmBackupStatus to be an explicit nil

### UnsetVmBackupStatus
`func (o *HypervVirtualMachine) UnsetVmBackupStatus()`

UnsetVmBackupStatus ensures that no value is present for VmBackupStatus, not even an explicit nil
### GetVmBackupType

`func (o *HypervVirtualMachine) GetVmBackupType() string`

GetVmBackupType returns the VmBackupType field if non-nil, zero value otherwise.

### GetVmBackupTypeOk

`func (o *HypervVirtualMachine) GetVmBackupTypeOk() (*string, bool)`

GetVmBackupTypeOk returns a tuple with the VmBackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmBackupType

`func (o *HypervVirtualMachine) SetVmBackupType(v string)`

SetVmBackupType sets VmBackupType field to given value.

### HasVmBackupType

`func (o *HypervVirtualMachine) HasVmBackupType() bool`

HasVmBackupType returns a boolean if a field has been set.

### SetVmBackupTypeNil

`func (o *HypervVirtualMachine) SetVmBackupTypeNil(b bool)`

 SetVmBackupTypeNil sets the value for VmBackupType to be an explicit nil

### UnsetVmBackupType
`func (o *HypervVirtualMachine) UnsetVmBackupType()`

UnsetVmBackupType ensures that no value is present for VmBackupType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


