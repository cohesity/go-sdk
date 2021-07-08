# RecoverVirtualDiskParamsVirtualDiskMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskToOverwrite** | Pointer to [**VirtualDiskId**](VirtualDiskId.md) |  | [optional] 
**SrcDisk** | Pointer to [**VirtualDiskId**](VirtualDiskId.md) |  | [optional] 
**TargetLocation** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 

## Methods

### NewRecoverVirtualDiskParamsVirtualDiskMapping

`func NewRecoverVirtualDiskParamsVirtualDiskMapping() *RecoverVirtualDiskParamsVirtualDiskMapping`

NewRecoverVirtualDiskParamsVirtualDiskMapping instantiates a new RecoverVirtualDiskParamsVirtualDiskMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverVirtualDiskParamsVirtualDiskMappingWithDefaults

`func NewRecoverVirtualDiskParamsVirtualDiskMappingWithDefaults() *RecoverVirtualDiskParamsVirtualDiskMapping`

NewRecoverVirtualDiskParamsVirtualDiskMappingWithDefaults instantiates a new RecoverVirtualDiskParamsVirtualDiskMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskToOverwrite

`func (o *RecoverVirtualDiskParamsVirtualDiskMapping) GetDiskToOverwrite() VirtualDiskId`

GetDiskToOverwrite returns the DiskToOverwrite field if non-nil, zero value otherwise.

### GetDiskToOverwriteOk

`func (o *RecoverVirtualDiskParamsVirtualDiskMapping) GetDiskToOverwriteOk() (*VirtualDiskId, bool)`

GetDiskToOverwriteOk returns a tuple with the DiskToOverwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskToOverwrite

`func (o *RecoverVirtualDiskParamsVirtualDiskMapping) SetDiskToOverwrite(v VirtualDiskId)`

SetDiskToOverwrite sets DiskToOverwrite field to given value.

### HasDiskToOverwrite

`func (o *RecoverVirtualDiskParamsVirtualDiskMapping) HasDiskToOverwrite() bool`

HasDiskToOverwrite returns a boolean if a field has been set.

### GetSrcDisk

`func (o *RecoverVirtualDiskParamsVirtualDiskMapping) GetSrcDisk() VirtualDiskId`

GetSrcDisk returns the SrcDisk field if non-nil, zero value otherwise.

### GetSrcDiskOk

`func (o *RecoverVirtualDiskParamsVirtualDiskMapping) GetSrcDiskOk() (*VirtualDiskId, bool)`

GetSrcDiskOk returns a tuple with the SrcDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcDisk

`func (o *RecoverVirtualDiskParamsVirtualDiskMapping) SetSrcDisk(v VirtualDiskId)`

SetSrcDisk sets SrcDisk field to given value.

### HasSrcDisk

`func (o *RecoverVirtualDiskParamsVirtualDiskMapping) HasSrcDisk() bool`

HasSrcDisk returns a boolean if a field has been set.

### GetTargetLocation

`func (o *RecoverVirtualDiskParamsVirtualDiskMapping) GetTargetLocation() EntityProto`

GetTargetLocation returns the TargetLocation field if non-nil, zero value otherwise.

### GetTargetLocationOk

`func (o *RecoverVirtualDiskParamsVirtualDiskMapping) GetTargetLocationOk() (*EntityProto, bool)`

GetTargetLocationOk returns a tuple with the TargetLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLocation

`func (o *RecoverVirtualDiskParamsVirtualDiskMapping) SetTargetLocation(v EntityProto)`

SetTargetLocation sets TargetLocation field to given value.

### HasTargetLocation

`func (o *RecoverVirtualDiskParamsVirtualDiskMapping) HasTargetLocation() bool`

HasTargetLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


