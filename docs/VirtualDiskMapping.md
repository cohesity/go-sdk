# VirtualDiskMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskToOverwrite** | Pointer to [**VirtualDiskIdInformation**](VirtualDiskIdInformation.md) |  | [optional] 
**SourceDisk** | Pointer to [**VirtualDiskIdInformation**](VirtualDiskIdInformation.md) |  | [optional] 
**TargetLocationId** | Pointer to **NullableInt64** | Specifies the target location information, for e.g. a datastore in VMware environment. If diskToOverwrite is specified, then the target location is automatically deduced. | [optional] 

## Methods

### NewVirtualDiskMapping

`func NewVirtualDiskMapping() *VirtualDiskMapping`

NewVirtualDiskMapping instantiates a new VirtualDiskMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualDiskMappingWithDefaults

`func NewVirtualDiskMappingWithDefaults() *VirtualDiskMapping`

NewVirtualDiskMappingWithDefaults instantiates a new VirtualDiskMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskToOverwrite

`func (o *VirtualDiskMapping) GetDiskToOverwrite() VirtualDiskIdInformation`

GetDiskToOverwrite returns the DiskToOverwrite field if non-nil, zero value otherwise.

### GetDiskToOverwriteOk

`func (o *VirtualDiskMapping) GetDiskToOverwriteOk() (*VirtualDiskIdInformation, bool)`

GetDiskToOverwriteOk returns a tuple with the DiskToOverwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskToOverwrite

`func (o *VirtualDiskMapping) SetDiskToOverwrite(v VirtualDiskIdInformation)`

SetDiskToOverwrite sets DiskToOverwrite field to given value.

### HasDiskToOverwrite

`func (o *VirtualDiskMapping) HasDiskToOverwrite() bool`

HasDiskToOverwrite returns a boolean if a field has been set.

### GetSourceDisk

`func (o *VirtualDiskMapping) GetSourceDisk() VirtualDiskIdInformation`

GetSourceDisk returns the SourceDisk field if non-nil, zero value otherwise.

### GetSourceDiskOk

`func (o *VirtualDiskMapping) GetSourceDiskOk() (*VirtualDiskIdInformation, bool)`

GetSourceDiskOk returns a tuple with the SourceDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDisk

`func (o *VirtualDiskMapping) SetSourceDisk(v VirtualDiskIdInformation)`

SetSourceDisk sets SourceDisk field to given value.

### HasSourceDisk

`func (o *VirtualDiskMapping) HasSourceDisk() bool`

HasSourceDisk returns a boolean if a field has been set.

### GetTargetLocationId

`func (o *VirtualDiskMapping) GetTargetLocationId() int64`

GetTargetLocationId returns the TargetLocationId field if non-nil, zero value otherwise.

### GetTargetLocationIdOk

`func (o *VirtualDiskMapping) GetTargetLocationIdOk() (*int64, bool)`

GetTargetLocationIdOk returns a tuple with the TargetLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLocationId

`func (o *VirtualDiskMapping) SetTargetLocationId(v int64)`

SetTargetLocationId sets TargetLocationId field to given value.

### HasTargetLocationId

`func (o *VirtualDiskMapping) HasTargetLocationId() bool`

HasTargetLocationId returns a boolean if a field has been set.

### SetTargetLocationIdNil

`func (o *VirtualDiskMapping) SetTargetLocationIdNil(b bool)`

 SetTargetLocationIdNil sets the value for TargetLocationId to be an explicit nil

### UnsetTargetLocationId
`func (o *VirtualDiskMapping) UnsetTargetLocationId()`

UnsetTargetLocationId ensures that no value is present for TargetLocationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


