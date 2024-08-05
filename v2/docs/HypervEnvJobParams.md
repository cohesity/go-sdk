# HypervEnvJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeDisks** | Pointer to [**[]HyperVDiskInfo**](HyperVDiskInfo.md) | Specifies a list of disks to exclude from being protected for the object/vm. | [optional] 
**FallbackToCrashConsistent** | Pointer to **NullableBool** | If true, takes a crash-consistent snapshot when app-consistent snapshot fails. Otherwise, the snapshot attempt is marked failed. | [optional] 
**IncludeDisks** | Pointer to [**[]HyperVDiskInfo**](HyperVDiskInfo.md) | Specifies a list of disks to included in the protection for the object/vm. | [optional] 
**ProtectionType** | Pointer to **NullableString** | Specifies the Protection Group type. If not specified, then backup method is auto determined. Specifying RCT will forcibly use RCT backup for all VMs in this Protection Group. Available only for VMs with hardware version 8.0 and above, but is more efficient. Specifying VSS will forcibly use VSS backup for all VMs in this Protection Group. Available for VMs with hardware version 5.0 and above, but is slower than RCT backup. | [optional] 

## Methods

### NewHypervEnvJobParams

`func NewHypervEnvJobParams() *HypervEnvJobParams`

NewHypervEnvJobParams instantiates a new HypervEnvJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervEnvJobParamsWithDefaults

`func NewHypervEnvJobParamsWithDefaults() *HypervEnvJobParams`

NewHypervEnvJobParamsWithDefaults instantiates a new HypervEnvJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeDisks

`func (o *HypervEnvJobParams) GetExcludeDisks() []HyperVDiskInfo`

GetExcludeDisks returns the ExcludeDisks field if non-nil, zero value otherwise.

### GetExcludeDisksOk

`func (o *HypervEnvJobParams) GetExcludeDisksOk() (*[]HyperVDiskInfo, bool)`

GetExcludeDisksOk returns a tuple with the ExcludeDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeDisks

`func (o *HypervEnvJobParams) SetExcludeDisks(v []HyperVDiskInfo)`

SetExcludeDisks sets ExcludeDisks field to given value.

### HasExcludeDisks

`func (o *HypervEnvJobParams) HasExcludeDisks() bool`

HasExcludeDisks returns a boolean if a field has been set.

### SetExcludeDisksNil

`func (o *HypervEnvJobParams) SetExcludeDisksNil(b bool)`

 SetExcludeDisksNil sets the value for ExcludeDisks to be an explicit nil

### UnsetExcludeDisks
`func (o *HypervEnvJobParams) UnsetExcludeDisks()`

UnsetExcludeDisks ensures that no value is present for ExcludeDisks, not even an explicit nil
### GetFallbackToCrashConsistent

`func (o *HypervEnvJobParams) GetFallbackToCrashConsistent() bool`

GetFallbackToCrashConsistent returns the FallbackToCrashConsistent field if non-nil, zero value otherwise.

### GetFallbackToCrashConsistentOk

`func (o *HypervEnvJobParams) GetFallbackToCrashConsistentOk() (*bool, bool)`

GetFallbackToCrashConsistentOk returns a tuple with the FallbackToCrashConsistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackToCrashConsistent

`func (o *HypervEnvJobParams) SetFallbackToCrashConsistent(v bool)`

SetFallbackToCrashConsistent sets FallbackToCrashConsistent field to given value.

### HasFallbackToCrashConsistent

`func (o *HypervEnvJobParams) HasFallbackToCrashConsistent() bool`

HasFallbackToCrashConsistent returns a boolean if a field has been set.

### SetFallbackToCrashConsistentNil

`func (o *HypervEnvJobParams) SetFallbackToCrashConsistentNil(b bool)`

 SetFallbackToCrashConsistentNil sets the value for FallbackToCrashConsistent to be an explicit nil

### UnsetFallbackToCrashConsistent
`func (o *HypervEnvJobParams) UnsetFallbackToCrashConsistent()`

UnsetFallbackToCrashConsistent ensures that no value is present for FallbackToCrashConsistent, not even an explicit nil
### GetIncludeDisks

`func (o *HypervEnvJobParams) GetIncludeDisks() []HyperVDiskInfo`

GetIncludeDisks returns the IncludeDisks field if non-nil, zero value otherwise.

### GetIncludeDisksOk

`func (o *HypervEnvJobParams) GetIncludeDisksOk() (*[]HyperVDiskInfo, bool)`

GetIncludeDisksOk returns a tuple with the IncludeDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDisks

`func (o *HypervEnvJobParams) SetIncludeDisks(v []HyperVDiskInfo)`

SetIncludeDisks sets IncludeDisks field to given value.

### HasIncludeDisks

`func (o *HypervEnvJobParams) HasIncludeDisks() bool`

HasIncludeDisks returns a boolean if a field has been set.

### SetIncludeDisksNil

`func (o *HypervEnvJobParams) SetIncludeDisksNil(b bool)`

 SetIncludeDisksNil sets the value for IncludeDisks to be an explicit nil

### UnsetIncludeDisks
`func (o *HypervEnvJobParams) UnsetIncludeDisks()`

UnsetIncludeDisks ensures that no value is present for IncludeDisks, not even an explicit nil
### GetProtectionType

`func (o *HypervEnvJobParams) GetProtectionType() string`

GetProtectionType returns the ProtectionType field if non-nil, zero value otherwise.

### GetProtectionTypeOk

`func (o *HypervEnvJobParams) GetProtectionTypeOk() (*string, bool)`

GetProtectionTypeOk returns a tuple with the ProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionType

`func (o *HypervEnvJobParams) SetProtectionType(v string)`

SetProtectionType sets ProtectionType field to given value.

### HasProtectionType

`func (o *HypervEnvJobParams) HasProtectionType() bool`

HasProtectionType returns a boolean if a field has been set.

### SetProtectionTypeNil

`func (o *HypervEnvJobParams) SetProtectionTypeNil(b bool)`

 SetProtectionTypeNil sets the value for ProtectionType to be an explicit nil

### UnsetProtectionType
`func (o *HypervEnvJobParams) UnsetProtectionType()`

UnsetProtectionType ensures that no value is present for ProtectionType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


