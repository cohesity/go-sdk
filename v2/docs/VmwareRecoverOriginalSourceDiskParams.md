# VmwareRecoverOriginalSourceDiskParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatastoreId** | Pointer to **NullableInt64** | Specifies the ID of the datastore on which the specified disk will be spun up. | [optional] 
**DiskUuid** | **NullableString** | Specifies the UUID of the source disk being recovered. | 
**OverwriteOriginalDisk** | Pointer to **NullableBool** | Specifies whether or not to overwrite the original disk. If this is set to true, then datastoreId should not be specified. Otherwise, datastoreId must be specified. | [optional] 

## Methods

### NewVmwareRecoverOriginalSourceDiskParams

`func NewVmwareRecoverOriginalSourceDiskParams(diskUuid NullableString, ) *VmwareRecoverOriginalSourceDiskParams`

NewVmwareRecoverOriginalSourceDiskParams instantiates a new VmwareRecoverOriginalSourceDiskParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareRecoverOriginalSourceDiskParamsWithDefaults

`func NewVmwareRecoverOriginalSourceDiskParamsWithDefaults() *VmwareRecoverOriginalSourceDiskParams`

NewVmwareRecoverOriginalSourceDiskParamsWithDefaults instantiates a new VmwareRecoverOriginalSourceDiskParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastoreId

`func (o *VmwareRecoverOriginalSourceDiskParams) GetDatastoreId() int64`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *VmwareRecoverOriginalSourceDiskParams) GetDatastoreIdOk() (*int64, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *VmwareRecoverOriginalSourceDiskParams) SetDatastoreId(v int64)`

SetDatastoreId sets DatastoreId field to given value.

### HasDatastoreId

`func (o *VmwareRecoverOriginalSourceDiskParams) HasDatastoreId() bool`

HasDatastoreId returns a boolean if a field has been set.

### SetDatastoreIdNil

`func (o *VmwareRecoverOriginalSourceDiskParams) SetDatastoreIdNil(b bool)`

 SetDatastoreIdNil sets the value for DatastoreId to be an explicit nil

### UnsetDatastoreId
`func (o *VmwareRecoverOriginalSourceDiskParams) UnsetDatastoreId()`

UnsetDatastoreId ensures that no value is present for DatastoreId, not even an explicit nil
### GetDiskUuid

`func (o *VmwareRecoverOriginalSourceDiskParams) GetDiskUuid() string`

GetDiskUuid returns the DiskUuid field if non-nil, zero value otherwise.

### GetDiskUuidOk

`func (o *VmwareRecoverOriginalSourceDiskParams) GetDiskUuidOk() (*string, bool)`

GetDiskUuidOk returns a tuple with the DiskUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUuid

`func (o *VmwareRecoverOriginalSourceDiskParams) SetDiskUuid(v string)`

SetDiskUuid sets DiskUuid field to given value.


### SetDiskUuidNil

`func (o *VmwareRecoverOriginalSourceDiskParams) SetDiskUuidNil(b bool)`

 SetDiskUuidNil sets the value for DiskUuid to be an explicit nil

### UnsetDiskUuid
`func (o *VmwareRecoverOriginalSourceDiskParams) UnsetDiskUuid()`

UnsetDiskUuid ensures that no value is present for DiskUuid, not even an explicit nil
### GetOverwriteOriginalDisk

`func (o *VmwareRecoverOriginalSourceDiskParams) GetOverwriteOriginalDisk() bool`

GetOverwriteOriginalDisk returns the OverwriteOriginalDisk field if non-nil, zero value otherwise.

### GetOverwriteOriginalDiskOk

`func (o *VmwareRecoverOriginalSourceDiskParams) GetOverwriteOriginalDiskOk() (*bool, bool)`

GetOverwriteOriginalDiskOk returns a tuple with the OverwriteOriginalDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteOriginalDisk

`func (o *VmwareRecoverOriginalSourceDiskParams) SetOverwriteOriginalDisk(v bool)`

SetOverwriteOriginalDisk sets OverwriteOriginalDisk field to given value.

### HasOverwriteOriginalDisk

`func (o *VmwareRecoverOriginalSourceDiskParams) HasOverwriteOriginalDisk() bool`

HasOverwriteOriginalDisk returns a boolean if a field has been set.

### SetOverwriteOriginalDiskNil

`func (o *VmwareRecoverOriginalSourceDiskParams) SetOverwriteOriginalDiskNil(b bool)`

 SetOverwriteOriginalDiskNil sets the value for OverwriteOriginalDisk to be an explicit nil

### UnsetOverwriteOriginalDisk
`func (o *VmwareRecoverOriginalSourceDiskParams) UnsetOverwriteOriginalDisk()`

UnsetOverwriteOriginalDisk ensures that no value is present for OverwriteOriginalDisk, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


