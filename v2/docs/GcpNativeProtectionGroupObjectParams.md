# GcpNativeProtectionGroupObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskExclusionNameParams** | Pointer to **[]string** | Specifies the paramaters to exclude disks attached to GCP VM instances. Here only the name of the disks are taken for exclusion. | [optional] 
**Id** | **NullableInt64** | Specifies the id of the object. | 
**Name** | Pointer to **NullableString** | Specifies the name of the virtual machine. | [optional] [readonly] 

## Methods

### NewGcpNativeProtectionGroupObjectParams

`func NewGcpNativeProtectionGroupObjectParams(id NullableInt64, ) *GcpNativeProtectionGroupObjectParams`

NewGcpNativeProtectionGroupObjectParams instantiates a new GcpNativeProtectionGroupObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpNativeProtectionGroupObjectParamsWithDefaults

`func NewGcpNativeProtectionGroupObjectParamsWithDefaults() *GcpNativeProtectionGroupObjectParams`

NewGcpNativeProtectionGroupObjectParamsWithDefaults instantiates a new GcpNativeProtectionGroupObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskExclusionNameParams

`func (o *GcpNativeProtectionGroupObjectParams) GetDiskExclusionNameParams() []string`

GetDiskExclusionNameParams returns the DiskExclusionNameParams field if non-nil, zero value otherwise.

### GetDiskExclusionNameParamsOk

`func (o *GcpNativeProtectionGroupObjectParams) GetDiskExclusionNameParamsOk() (*[]string, bool)`

GetDiskExclusionNameParamsOk returns a tuple with the DiskExclusionNameParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskExclusionNameParams

`func (o *GcpNativeProtectionGroupObjectParams) SetDiskExclusionNameParams(v []string)`

SetDiskExclusionNameParams sets DiskExclusionNameParams field to given value.

### HasDiskExclusionNameParams

`func (o *GcpNativeProtectionGroupObjectParams) HasDiskExclusionNameParams() bool`

HasDiskExclusionNameParams returns a boolean if a field has been set.

### SetDiskExclusionNameParamsNil

`func (o *GcpNativeProtectionGroupObjectParams) SetDiskExclusionNameParamsNil(b bool)`

 SetDiskExclusionNameParamsNil sets the value for DiskExclusionNameParams to be an explicit nil

### UnsetDiskExclusionNameParams
`func (o *GcpNativeProtectionGroupObjectParams) UnsetDiskExclusionNameParams()`

UnsetDiskExclusionNameParams ensures that no value is present for DiskExclusionNameParams, not even an explicit nil
### GetId

`func (o *GcpNativeProtectionGroupObjectParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GcpNativeProtectionGroupObjectParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GcpNativeProtectionGroupObjectParams) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *GcpNativeProtectionGroupObjectParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GcpNativeProtectionGroupObjectParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *GcpNativeProtectionGroupObjectParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GcpNativeProtectionGroupObjectParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GcpNativeProtectionGroupObjectParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GcpNativeProtectionGroupObjectParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GcpNativeProtectionGroupObjectParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GcpNativeProtectionGroupObjectParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


