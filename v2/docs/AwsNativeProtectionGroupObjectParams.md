# AwsNativeProtectionGroupObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableInt64** | Specifies the id of the object. | 
**Name** | Pointer to **NullableString** | Specifies the name of the virtual machine. | [optional] [readonly] 
**VolumeExclusionParams** | Pointer to [**NullableEbsVolumeExclusionParams**](EbsVolumeExclusionParams.md) |  | [optional] 

## Methods

### NewAwsNativeProtectionGroupObjectParams

`func NewAwsNativeProtectionGroupObjectParams(id NullableInt64, ) *AwsNativeProtectionGroupObjectParams`

NewAwsNativeProtectionGroupObjectParams instantiates a new AwsNativeProtectionGroupObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsNativeProtectionGroupObjectParamsWithDefaults

`func NewAwsNativeProtectionGroupObjectParamsWithDefaults() *AwsNativeProtectionGroupObjectParams`

NewAwsNativeProtectionGroupObjectParamsWithDefaults instantiates a new AwsNativeProtectionGroupObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AwsNativeProtectionGroupObjectParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AwsNativeProtectionGroupObjectParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AwsNativeProtectionGroupObjectParams) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *AwsNativeProtectionGroupObjectParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AwsNativeProtectionGroupObjectParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *AwsNativeProtectionGroupObjectParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsNativeProtectionGroupObjectParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsNativeProtectionGroupObjectParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AwsNativeProtectionGroupObjectParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AwsNativeProtectionGroupObjectParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AwsNativeProtectionGroupObjectParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVolumeExclusionParams

`func (o *AwsNativeProtectionGroupObjectParams) GetVolumeExclusionParams() EbsVolumeExclusionParams`

GetVolumeExclusionParams returns the VolumeExclusionParams field if non-nil, zero value otherwise.

### GetVolumeExclusionParamsOk

`func (o *AwsNativeProtectionGroupObjectParams) GetVolumeExclusionParamsOk() (*EbsVolumeExclusionParams, bool)`

GetVolumeExclusionParamsOk returns a tuple with the VolumeExclusionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeExclusionParams

`func (o *AwsNativeProtectionGroupObjectParams) SetVolumeExclusionParams(v EbsVolumeExclusionParams)`

SetVolumeExclusionParams sets VolumeExclusionParams field to given value.

### HasVolumeExclusionParams

`func (o *AwsNativeProtectionGroupObjectParams) HasVolumeExclusionParams() bool`

HasVolumeExclusionParams returns a boolean if a field has been set.

### SetVolumeExclusionParamsNil

`func (o *AwsNativeProtectionGroupObjectParams) SetVolumeExclusionParamsNil(b bool)`

 SetVolumeExclusionParamsNil sets the value for VolumeExclusionParams to be an explicit nil

### UnsetVolumeExclusionParams
`func (o *AwsNativeProtectionGroupObjectParams) UnsetVolumeExclusionParams()`

UnsetVolumeExclusionParams ensures that no value is present for VolumeExclusionParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


