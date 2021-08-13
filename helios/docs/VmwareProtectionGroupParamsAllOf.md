# VmwareProtectionGroupParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | Pointer to [**[]VmwareProtectionGroupObjectParams**](VmwareProtectionGroupObjectParams.md) | Specifies the objects to include in the backup. | [optional] 
**GlobalExcludeDisks** | Pointer to [**[]DiskInfo**](DiskInfo.md) | Specifies a list of disks to exclude from the backup. | [optional] 
**StandbyResourceObjects** | Pointer to [**[]VmwareProtectionGroupStandbyResourceParams**](VmwareProtectionGroupStandbyResourceParams.md) | Specifies the standby resource objects for this backup. | [optional] 

## Methods

### NewVmwareProtectionGroupParamsAllOf

`func NewVmwareProtectionGroupParamsAllOf() *VmwareProtectionGroupParamsAllOf`

NewVmwareProtectionGroupParamsAllOf instantiates a new VmwareProtectionGroupParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareProtectionGroupParamsAllOfWithDefaults

`func NewVmwareProtectionGroupParamsAllOfWithDefaults() *VmwareProtectionGroupParamsAllOf`

NewVmwareProtectionGroupParamsAllOfWithDefaults instantiates a new VmwareProtectionGroupParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *VmwareProtectionGroupParamsAllOf) GetObjects() []VmwareProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *VmwareProtectionGroupParamsAllOf) GetObjectsOk() (*[]VmwareProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *VmwareProtectionGroupParamsAllOf) SetObjects(v []VmwareProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *VmwareProtectionGroupParamsAllOf) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetGlobalExcludeDisks

`func (o *VmwareProtectionGroupParamsAllOf) GetGlobalExcludeDisks() []DiskInfo`

GetGlobalExcludeDisks returns the GlobalExcludeDisks field if non-nil, zero value otherwise.

### GetGlobalExcludeDisksOk

`func (o *VmwareProtectionGroupParamsAllOf) GetGlobalExcludeDisksOk() (*[]DiskInfo, bool)`

GetGlobalExcludeDisksOk returns a tuple with the GlobalExcludeDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalExcludeDisks

`func (o *VmwareProtectionGroupParamsAllOf) SetGlobalExcludeDisks(v []DiskInfo)`

SetGlobalExcludeDisks sets GlobalExcludeDisks field to given value.

### HasGlobalExcludeDisks

`func (o *VmwareProtectionGroupParamsAllOf) HasGlobalExcludeDisks() bool`

HasGlobalExcludeDisks returns a boolean if a field has been set.

### SetGlobalExcludeDisksNil

`func (o *VmwareProtectionGroupParamsAllOf) SetGlobalExcludeDisksNil(b bool)`

 SetGlobalExcludeDisksNil sets the value for GlobalExcludeDisks to be an explicit nil

### UnsetGlobalExcludeDisks
`func (o *VmwareProtectionGroupParamsAllOf) UnsetGlobalExcludeDisks()`

UnsetGlobalExcludeDisks ensures that no value is present for GlobalExcludeDisks, not even an explicit nil
### GetStandbyResourceObjects

`func (o *VmwareProtectionGroupParamsAllOf) GetStandbyResourceObjects() []VmwareProtectionGroupStandbyResourceParams`

GetStandbyResourceObjects returns the StandbyResourceObjects field if non-nil, zero value otherwise.

### GetStandbyResourceObjectsOk

`func (o *VmwareProtectionGroupParamsAllOf) GetStandbyResourceObjectsOk() (*[]VmwareProtectionGroupStandbyResourceParams, bool)`

GetStandbyResourceObjectsOk returns a tuple with the StandbyResourceObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyResourceObjects

`func (o *VmwareProtectionGroupParamsAllOf) SetStandbyResourceObjects(v []VmwareProtectionGroupStandbyResourceParams)`

SetStandbyResourceObjects sets StandbyResourceObjects field to given value.

### HasStandbyResourceObjects

`func (o *VmwareProtectionGroupParamsAllOf) HasStandbyResourceObjects() bool`

HasStandbyResourceObjects returns a boolean if a field has been set.

### SetStandbyResourceObjectsNil

`func (o *VmwareProtectionGroupParamsAllOf) SetStandbyResourceObjectsNil(b bool)`

 SetStandbyResourceObjectsNil sets the value for StandbyResourceObjects to be an explicit nil

### UnsetStandbyResourceObjects
`func (o *VmwareProtectionGroupParamsAllOf) UnsetStandbyResourceObjects()`

UnsetStandbyResourceObjects ensures that no value is present for StandbyResourceObjects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


