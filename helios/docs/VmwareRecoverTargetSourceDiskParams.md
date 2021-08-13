# VmwareRecoverTargetSourceDiskParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskUuid** | **NullableString** | Specifies the UUID of the source disk being recovered. | 
**DatastoreId** | **NullableInt64** | Specifies the ID of the datastore on which the specified disk will be spun up. | 

## Methods

### NewVmwareRecoverTargetSourceDiskParams

`func NewVmwareRecoverTargetSourceDiskParams(diskUuid NullableString, datastoreId NullableInt64, ) *VmwareRecoverTargetSourceDiskParams`

NewVmwareRecoverTargetSourceDiskParams instantiates a new VmwareRecoverTargetSourceDiskParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareRecoverTargetSourceDiskParamsWithDefaults

`func NewVmwareRecoverTargetSourceDiskParamsWithDefaults() *VmwareRecoverTargetSourceDiskParams`

NewVmwareRecoverTargetSourceDiskParamsWithDefaults instantiates a new VmwareRecoverTargetSourceDiskParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskUuid

`func (o *VmwareRecoverTargetSourceDiskParams) GetDiskUuid() string`

GetDiskUuid returns the DiskUuid field if non-nil, zero value otherwise.

### GetDiskUuidOk

`func (o *VmwareRecoverTargetSourceDiskParams) GetDiskUuidOk() (*string, bool)`

GetDiskUuidOk returns a tuple with the DiskUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUuid

`func (o *VmwareRecoverTargetSourceDiskParams) SetDiskUuid(v string)`

SetDiskUuid sets DiskUuid field to given value.


### SetDiskUuidNil

`func (o *VmwareRecoverTargetSourceDiskParams) SetDiskUuidNil(b bool)`

 SetDiskUuidNil sets the value for DiskUuid to be an explicit nil

### UnsetDiskUuid
`func (o *VmwareRecoverTargetSourceDiskParams) UnsetDiskUuid()`

UnsetDiskUuid ensures that no value is present for DiskUuid, not even an explicit nil
### GetDatastoreId

`func (o *VmwareRecoverTargetSourceDiskParams) GetDatastoreId() int64`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *VmwareRecoverTargetSourceDiskParams) GetDatastoreIdOk() (*int64, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *VmwareRecoverTargetSourceDiskParams) SetDatastoreId(v int64)`

SetDatastoreId sets DatastoreId field to given value.


### SetDatastoreIdNil

`func (o *VmwareRecoverTargetSourceDiskParams) SetDatastoreIdNil(b bool)`

 SetDatastoreIdNil sets the value for DatastoreId to be an explicit nil

### UnsetDatastoreId
`func (o *VmwareRecoverTargetSourceDiskParams) UnsetDatastoreId()`

UnsetDatastoreId ensures that no value is present for DatastoreId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


