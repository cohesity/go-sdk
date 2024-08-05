# HyperVProtectionGroupObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeDisks** | Pointer to [**[]HyperVDiskInfo**](HyperVDiskInfo.md) | Specifies a list of disks to exclude from being protected for the object/vm. | [optional] 
**Id** | **NullableInt64** | Specifies the id of the object. | 
**IncludeDisks** | Pointer to [**[]HyperVDiskInfo**](HyperVDiskInfo.md) | Specifies a list of disks to included in the protection for the object/vm. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the virtual machine. | [optional] [readonly] 

## Methods

### NewHyperVProtectionGroupObjectParams

`func NewHyperVProtectionGroupObjectParams(id NullableInt64, ) *HyperVProtectionGroupObjectParams`

NewHyperVProtectionGroupObjectParams instantiates a new HyperVProtectionGroupObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperVProtectionGroupObjectParamsWithDefaults

`func NewHyperVProtectionGroupObjectParamsWithDefaults() *HyperVProtectionGroupObjectParams`

NewHyperVProtectionGroupObjectParamsWithDefaults instantiates a new HyperVProtectionGroupObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeDisks

`func (o *HyperVProtectionGroupObjectParams) GetExcludeDisks() []HyperVDiskInfo`

GetExcludeDisks returns the ExcludeDisks field if non-nil, zero value otherwise.

### GetExcludeDisksOk

`func (o *HyperVProtectionGroupObjectParams) GetExcludeDisksOk() (*[]HyperVDiskInfo, bool)`

GetExcludeDisksOk returns a tuple with the ExcludeDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeDisks

`func (o *HyperVProtectionGroupObjectParams) SetExcludeDisks(v []HyperVDiskInfo)`

SetExcludeDisks sets ExcludeDisks field to given value.

### HasExcludeDisks

`func (o *HyperVProtectionGroupObjectParams) HasExcludeDisks() bool`

HasExcludeDisks returns a boolean if a field has been set.

### SetExcludeDisksNil

`func (o *HyperVProtectionGroupObjectParams) SetExcludeDisksNil(b bool)`

 SetExcludeDisksNil sets the value for ExcludeDisks to be an explicit nil

### UnsetExcludeDisks
`func (o *HyperVProtectionGroupObjectParams) UnsetExcludeDisks()`

UnsetExcludeDisks ensures that no value is present for ExcludeDisks, not even an explicit nil
### GetId

`func (o *HyperVProtectionGroupObjectParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HyperVProtectionGroupObjectParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HyperVProtectionGroupObjectParams) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *HyperVProtectionGroupObjectParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HyperVProtectionGroupObjectParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIncludeDisks

`func (o *HyperVProtectionGroupObjectParams) GetIncludeDisks() []HyperVDiskInfo`

GetIncludeDisks returns the IncludeDisks field if non-nil, zero value otherwise.

### GetIncludeDisksOk

`func (o *HyperVProtectionGroupObjectParams) GetIncludeDisksOk() (*[]HyperVDiskInfo, bool)`

GetIncludeDisksOk returns a tuple with the IncludeDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDisks

`func (o *HyperVProtectionGroupObjectParams) SetIncludeDisks(v []HyperVDiskInfo)`

SetIncludeDisks sets IncludeDisks field to given value.

### HasIncludeDisks

`func (o *HyperVProtectionGroupObjectParams) HasIncludeDisks() bool`

HasIncludeDisks returns a boolean if a field has been set.

### SetIncludeDisksNil

`func (o *HyperVProtectionGroupObjectParams) SetIncludeDisksNil(b bool)`

 SetIncludeDisksNil sets the value for IncludeDisks to be an explicit nil

### UnsetIncludeDisks
`func (o *HyperVProtectionGroupObjectParams) UnsetIncludeDisks()`

UnsetIncludeDisks ensures that no value is present for IncludeDisks, not even an explicit nil
### GetName

`func (o *HyperVProtectionGroupObjectParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperVProtectionGroupObjectParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperVProtectionGroupObjectParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperVProtectionGroupObjectParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HyperVProtectionGroupObjectParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HyperVProtectionGroupObjectParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


