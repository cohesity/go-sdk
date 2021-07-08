# VirtualDiskMappingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskToOverwrite** | Pointer to [**VirtualDiskIdInformation**](VirtualDiskIdInformation.md) |  | [optional] 
**SourceDisk** | Pointer to [**VirtualDiskIdInformation**](VirtualDiskIdInformation.md) |  | [optional] 
**TargetLocation** | Pointer to [**ProtectionSource**](ProtectionSource.md) |  | [optional] 

## Methods

### NewVirtualDiskMappingResponse

`func NewVirtualDiskMappingResponse() *VirtualDiskMappingResponse`

NewVirtualDiskMappingResponse instantiates a new VirtualDiskMappingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualDiskMappingResponseWithDefaults

`func NewVirtualDiskMappingResponseWithDefaults() *VirtualDiskMappingResponse`

NewVirtualDiskMappingResponseWithDefaults instantiates a new VirtualDiskMappingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskToOverwrite

`func (o *VirtualDiskMappingResponse) GetDiskToOverwrite() VirtualDiskIdInformation`

GetDiskToOverwrite returns the DiskToOverwrite field if non-nil, zero value otherwise.

### GetDiskToOverwriteOk

`func (o *VirtualDiskMappingResponse) GetDiskToOverwriteOk() (*VirtualDiskIdInformation, bool)`

GetDiskToOverwriteOk returns a tuple with the DiskToOverwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskToOverwrite

`func (o *VirtualDiskMappingResponse) SetDiskToOverwrite(v VirtualDiskIdInformation)`

SetDiskToOverwrite sets DiskToOverwrite field to given value.

### HasDiskToOverwrite

`func (o *VirtualDiskMappingResponse) HasDiskToOverwrite() bool`

HasDiskToOverwrite returns a boolean if a field has been set.

### GetSourceDisk

`func (o *VirtualDiskMappingResponse) GetSourceDisk() VirtualDiskIdInformation`

GetSourceDisk returns the SourceDisk field if non-nil, zero value otherwise.

### GetSourceDiskOk

`func (o *VirtualDiskMappingResponse) GetSourceDiskOk() (*VirtualDiskIdInformation, bool)`

GetSourceDiskOk returns a tuple with the SourceDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDisk

`func (o *VirtualDiskMappingResponse) SetSourceDisk(v VirtualDiskIdInformation)`

SetSourceDisk sets SourceDisk field to given value.

### HasSourceDisk

`func (o *VirtualDiskMappingResponse) HasSourceDisk() bool`

HasSourceDisk returns a boolean if a field has been set.

### GetTargetLocation

`func (o *VirtualDiskMappingResponse) GetTargetLocation() ProtectionSource`

GetTargetLocation returns the TargetLocation field if non-nil, zero value otherwise.

### GetTargetLocationOk

`func (o *VirtualDiskMappingResponse) GetTargetLocationOk() (*ProtectionSource, bool)`

GetTargetLocationOk returns a tuple with the TargetLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLocation

`func (o *VirtualDiskMappingResponse) SetTargetLocation(v ProtectionSource)`

SetTargetLocation sets TargetLocation field to given value.

### HasTargetLocation

`func (o *VirtualDiskMappingResponse) HasTargetLocation() bool`

HasTargetLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


