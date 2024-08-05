# AwsNativeObjectProtectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | Pointer to [**[]AwsObjectLevelParams**](AwsObjectLevelParams.md) | Specifies the objects to be protected. | [optional] 
**VolumeExclusionParams** | Pointer to [**NullableEbsVolumeExclusionParams**](EbsVolumeExclusionParams.md) |  | [optional] 

## Methods

### NewAwsNativeObjectProtectionParams

`func NewAwsNativeObjectProtectionParams() *AwsNativeObjectProtectionParams`

NewAwsNativeObjectProtectionParams instantiates a new AwsNativeObjectProtectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsNativeObjectProtectionParamsWithDefaults

`func NewAwsNativeObjectProtectionParamsWithDefaults() *AwsNativeObjectProtectionParams`

NewAwsNativeObjectProtectionParamsWithDefaults instantiates a new AwsNativeObjectProtectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *AwsNativeObjectProtectionParams) GetObjects() []AwsObjectLevelParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *AwsNativeObjectProtectionParams) GetObjectsOk() (*[]AwsObjectLevelParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *AwsNativeObjectProtectionParams) SetObjects(v []AwsObjectLevelParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *AwsNativeObjectProtectionParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetVolumeExclusionParams

`func (o *AwsNativeObjectProtectionParams) GetVolumeExclusionParams() EbsVolumeExclusionParams`

GetVolumeExclusionParams returns the VolumeExclusionParams field if non-nil, zero value otherwise.

### GetVolumeExclusionParamsOk

`func (o *AwsNativeObjectProtectionParams) GetVolumeExclusionParamsOk() (*EbsVolumeExclusionParams, bool)`

GetVolumeExclusionParamsOk returns a tuple with the VolumeExclusionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeExclusionParams

`func (o *AwsNativeObjectProtectionParams) SetVolumeExclusionParams(v EbsVolumeExclusionParams)`

SetVolumeExclusionParams sets VolumeExclusionParams field to given value.

### HasVolumeExclusionParams

`func (o *AwsNativeObjectProtectionParams) HasVolumeExclusionParams() bool`

HasVolumeExclusionParams returns a boolean if a field has been set.

### SetVolumeExclusionParamsNil

`func (o *AwsNativeObjectProtectionParams) SetVolumeExclusionParamsNil(b bool)`

 SetVolumeExclusionParamsNil sets the value for VolumeExclusionParams to be an explicit nil

### UnsetVolumeExclusionParams
`func (o *AwsNativeObjectProtectionParams) UnsetVolumeExclusionParams()`

UnsetVolumeExclusionParams ensures that no value is present for VolumeExclusionParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


