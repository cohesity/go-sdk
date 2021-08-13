# VmwareObjectProtectionRequestParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]VmwareObjectProtectionRequest**](VmwareObjectProtectionRequest.md) | Specifies the objects to include in the backup. | 
**GlobalExcludeDisks** | Pointer to [**[]DiskInfo**](DiskInfo.md) | Specifies a list of disks to exclude from the backup. | [optional] 

## Methods

### NewVmwareObjectProtectionRequestParamsAllOf

`func NewVmwareObjectProtectionRequestParamsAllOf(objects []VmwareObjectProtectionRequest, ) *VmwareObjectProtectionRequestParamsAllOf`

NewVmwareObjectProtectionRequestParamsAllOf instantiates a new VmwareObjectProtectionRequestParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareObjectProtectionRequestParamsAllOfWithDefaults

`func NewVmwareObjectProtectionRequestParamsAllOfWithDefaults() *VmwareObjectProtectionRequestParamsAllOf`

NewVmwareObjectProtectionRequestParamsAllOfWithDefaults instantiates a new VmwareObjectProtectionRequestParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *VmwareObjectProtectionRequestParamsAllOf) GetObjects() []VmwareObjectProtectionRequest`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *VmwareObjectProtectionRequestParamsAllOf) GetObjectsOk() (*[]VmwareObjectProtectionRequest, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *VmwareObjectProtectionRequestParamsAllOf) SetObjects(v []VmwareObjectProtectionRequest)`

SetObjects sets Objects field to given value.


### GetGlobalExcludeDisks

`func (o *VmwareObjectProtectionRequestParamsAllOf) GetGlobalExcludeDisks() []DiskInfo`

GetGlobalExcludeDisks returns the GlobalExcludeDisks field if non-nil, zero value otherwise.

### GetGlobalExcludeDisksOk

`func (o *VmwareObjectProtectionRequestParamsAllOf) GetGlobalExcludeDisksOk() (*[]DiskInfo, bool)`

GetGlobalExcludeDisksOk returns a tuple with the GlobalExcludeDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalExcludeDisks

`func (o *VmwareObjectProtectionRequestParamsAllOf) SetGlobalExcludeDisks(v []DiskInfo)`

SetGlobalExcludeDisks sets GlobalExcludeDisks field to given value.

### HasGlobalExcludeDisks

`func (o *VmwareObjectProtectionRequestParamsAllOf) HasGlobalExcludeDisks() bool`

HasGlobalExcludeDisks returns a boolean if a field has been set.

### SetGlobalExcludeDisksNil

`func (o *VmwareObjectProtectionRequestParamsAllOf) SetGlobalExcludeDisksNil(b bool)`

 SetGlobalExcludeDisksNil sets the value for GlobalExcludeDisks to be an explicit nil

### UnsetGlobalExcludeDisks
`func (o *VmwareObjectProtectionRequestParamsAllOf) UnsetGlobalExcludeDisks()`

UnsetGlobalExcludeDisks ensures that no value is present for GlobalExcludeDisks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


