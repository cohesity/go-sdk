# VmwareEnvJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludedDisks** | Pointer to [**[]DiskInfo**](DiskInfo.md) | Specifies the list of Disks to be excluded from backing up. These disks are excluded from all Protection Sources in the Protection Job. | [optional] 
**FallbackToCrashConsistent** | Pointer to **NullableBool** | If true, takes a crash-consistent snapshot when app-consistent snapshot fails. Otherwise, the snapshot attempt is marked failed. | [optional] 
**SkipPhysicalRdmDisks** | Pointer to **NullableBool** | If true, skip physical RDM disks when backing up VMs. Otherwise, backup of VMs having physical RDM will fail. | [optional] 

## Methods

### NewVmwareEnvJobParams

`func NewVmwareEnvJobParams() *VmwareEnvJobParams`

NewVmwareEnvJobParams instantiates a new VmwareEnvJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareEnvJobParamsWithDefaults

`func NewVmwareEnvJobParamsWithDefaults() *VmwareEnvJobParams`

NewVmwareEnvJobParamsWithDefaults instantiates a new VmwareEnvJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludedDisks

`func (o *VmwareEnvJobParams) GetExcludedDisks() []DiskInfo`

GetExcludedDisks returns the ExcludedDisks field if non-nil, zero value otherwise.

### GetExcludedDisksOk

`func (o *VmwareEnvJobParams) GetExcludedDisksOk() (*[]DiskInfo, bool)`

GetExcludedDisksOk returns a tuple with the ExcludedDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedDisks

`func (o *VmwareEnvJobParams) SetExcludedDisks(v []DiskInfo)`

SetExcludedDisks sets ExcludedDisks field to given value.

### HasExcludedDisks

`func (o *VmwareEnvJobParams) HasExcludedDisks() bool`

HasExcludedDisks returns a boolean if a field has been set.

### SetExcludedDisksNil

`func (o *VmwareEnvJobParams) SetExcludedDisksNil(b bool)`

 SetExcludedDisksNil sets the value for ExcludedDisks to be an explicit nil

### UnsetExcludedDisks
`func (o *VmwareEnvJobParams) UnsetExcludedDisks()`

UnsetExcludedDisks ensures that no value is present for ExcludedDisks, not even an explicit nil
### GetFallbackToCrashConsistent

`func (o *VmwareEnvJobParams) GetFallbackToCrashConsistent() bool`

GetFallbackToCrashConsistent returns the FallbackToCrashConsistent field if non-nil, zero value otherwise.

### GetFallbackToCrashConsistentOk

`func (o *VmwareEnvJobParams) GetFallbackToCrashConsistentOk() (*bool, bool)`

GetFallbackToCrashConsistentOk returns a tuple with the FallbackToCrashConsistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackToCrashConsistent

`func (o *VmwareEnvJobParams) SetFallbackToCrashConsistent(v bool)`

SetFallbackToCrashConsistent sets FallbackToCrashConsistent field to given value.

### HasFallbackToCrashConsistent

`func (o *VmwareEnvJobParams) HasFallbackToCrashConsistent() bool`

HasFallbackToCrashConsistent returns a boolean if a field has been set.

### SetFallbackToCrashConsistentNil

`func (o *VmwareEnvJobParams) SetFallbackToCrashConsistentNil(b bool)`

 SetFallbackToCrashConsistentNil sets the value for FallbackToCrashConsistent to be an explicit nil

### UnsetFallbackToCrashConsistent
`func (o *VmwareEnvJobParams) UnsetFallbackToCrashConsistent()`

UnsetFallbackToCrashConsistent ensures that no value is present for FallbackToCrashConsistent, not even an explicit nil
### GetSkipPhysicalRdmDisks

`func (o *VmwareEnvJobParams) GetSkipPhysicalRdmDisks() bool`

GetSkipPhysicalRdmDisks returns the SkipPhysicalRdmDisks field if non-nil, zero value otherwise.

### GetSkipPhysicalRdmDisksOk

`func (o *VmwareEnvJobParams) GetSkipPhysicalRdmDisksOk() (*bool, bool)`

GetSkipPhysicalRdmDisksOk returns a tuple with the SkipPhysicalRdmDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipPhysicalRdmDisks

`func (o *VmwareEnvJobParams) SetSkipPhysicalRdmDisks(v bool)`

SetSkipPhysicalRdmDisks sets SkipPhysicalRdmDisks field to given value.

### HasSkipPhysicalRdmDisks

`func (o *VmwareEnvJobParams) HasSkipPhysicalRdmDisks() bool`

HasSkipPhysicalRdmDisks returns a boolean if a field has been set.

### SetSkipPhysicalRdmDisksNil

`func (o *VmwareEnvJobParams) SetSkipPhysicalRdmDisksNil(b bool)`

 SetSkipPhysicalRdmDisksNil sets the value for SkipPhysicalRdmDisks to be an explicit nil

### UnsetSkipPhysicalRdmDisks
`func (o *VmwareEnvJobParams) UnsetSkipPhysicalRdmDisks()`

UnsetSkipPhysicalRdmDisks ensures that no value is present for SkipPhysicalRdmDisks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


