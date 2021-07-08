# VmwareEnvJobParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludedDisks** | Pointer to [**[]DiskUnit**](DiskUnit.md) | Specifies the list of Disks to be excluded from backing up. These disks are excluded from all Protection Sources in the Protection Job. | [optional] 
**FallbackToCrashConsistent** | Pointer to **NullableBool** | If true, takes a crash-consistent snapshot when app-consistent snapshot fails. Otherwise, the snapshot attempt is marked failed. | [optional] 
**SkipPhysicalRdmDisks** | Pointer to **NullableBool** | If true, skip physical RDM disks when backing up VMs. Otherwise, backup of VMs having physical RDM will fail. | [optional] 

## Methods

### NewVmwareEnvJobParameters

`func NewVmwareEnvJobParameters() *VmwareEnvJobParameters`

NewVmwareEnvJobParameters instantiates a new VmwareEnvJobParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareEnvJobParametersWithDefaults

`func NewVmwareEnvJobParametersWithDefaults() *VmwareEnvJobParameters`

NewVmwareEnvJobParametersWithDefaults instantiates a new VmwareEnvJobParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludedDisks

`func (o *VmwareEnvJobParameters) GetExcludedDisks() []DiskUnit`

GetExcludedDisks returns the ExcludedDisks field if non-nil, zero value otherwise.

### GetExcludedDisksOk

`func (o *VmwareEnvJobParameters) GetExcludedDisksOk() (*[]DiskUnit, bool)`

GetExcludedDisksOk returns a tuple with the ExcludedDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedDisks

`func (o *VmwareEnvJobParameters) SetExcludedDisks(v []DiskUnit)`

SetExcludedDisks sets ExcludedDisks field to given value.

### HasExcludedDisks

`func (o *VmwareEnvJobParameters) HasExcludedDisks() bool`

HasExcludedDisks returns a boolean if a field has been set.

### SetExcludedDisksNil

`func (o *VmwareEnvJobParameters) SetExcludedDisksNil(b bool)`

 SetExcludedDisksNil sets the value for ExcludedDisks to be an explicit nil

### UnsetExcludedDisks
`func (o *VmwareEnvJobParameters) UnsetExcludedDisks()`

UnsetExcludedDisks ensures that no value is present for ExcludedDisks, not even an explicit nil
### GetFallbackToCrashConsistent

`func (o *VmwareEnvJobParameters) GetFallbackToCrashConsistent() bool`

GetFallbackToCrashConsistent returns the FallbackToCrashConsistent field if non-nil, zero value otherwise.

### GetFallbackToCrashConsistentOk

`func (o *VmwareEnvJobParameters) GetFallbackToCrashConsistentOk() (*bool, bool)`

GetFallbackToCrashConsistentOk returns a tuple with the FallbackToCrashConsistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackToCrashConsistent

`func (o *VmwareEnvJobParameters) SetFallbackToCrashConsistent(v bool)`

SetFallbackToCrashConsistent sets FallbackToCrashConsistent field to given value.

### HasFallbackToCrashConsistent

`func (o *VmwareEnvJobParameters) HasFallbackToCrashConsistent() bool`

HasFallbackToCrashConsistent returns a boolean if a field has been set.

### SetFallbackToCrashConsistentNil

`func (o *VmwareEnvJobParameters) SetFallbackToCrashConsistentNil(b bool)`

 SetFallbackToCrashConsistentNil sets the value for FallbackToCrashConsistent to be an explicit nil

### UnsetFallbackToCrashConsistent
`func (o *VmwareEnvJobParameters) UnsetFallbackToCrashConsistent()`

UnsetFallbackToCrashConsistent ensures that no value is present for FallbackToCrashConsistent, not even an explicit nil
### GetSkipPhysicalRdmDisks

`func (o *VmwareEnvJobParameters) GetSkipPhysicalRdmDisks() bool`

GetSkipPhysicalRdmDisks returns the SkipPhysicalRdmDisks field if non-nil, zero value otherwise.

### GetSkipPhysicalRdmDisksOk

`func (o *VmwareEnvJobParameters) GetSkipPhysicalRdmDisksOk() (*bool, bool)`

GetSkipPhysicalRdmDisksOk returns a tuple with the SkipPhysicalRdmDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipPhysicalRdmDisks

`func (o *VmwareEnvJobParameters) SetSkipPhysicalRdmDisks(v bool)`

SetSkipPhysicalRdmDisks sets SkipPhysicalRdmDisks field to given value.

### HasSkipPhysicalRdmDisks

`func (o *VmwareEnvJobParameters) HasSkipPhysicalRdmDisks() bool`

HasSkipPhysicalRdmDisks returns a boolean if a field has been set.

### SetSkipPhysicalRdmDisksNil

`func (o *VmwareEnvJobParameters) SetSkipPhysicalRdmDisksNil(b bool)`

 SetSkipPhysicalRdmDisksNil sets the value for SkipPhysicalRdmDisks to be an explicit nil

### UnsetSkipPhysicalRdmDisks
`func (o *VmwareEnvJobParameters) UnsetSkipPhysicalRdmDisks()`

UnsetSkipPhysicalRdmDisks ensures that no value is present for SkipPhysicalRdmDisks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


