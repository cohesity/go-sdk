# AwsSnapshotManagerObjectProtectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmiCreationFrequency** | Pointer to **NullableInt32** | The frequency of AMI creation. This should be set if the option to create AMI is set. A value of n creates an AMI from the snapshots after every n runs. eg. n &#x3D; 2 implies every alternate backup run starting from the first will create an AMI. | [optional] 
**CreateAmi** | Pointer to **NullableBool** | Specifies whether AMI should be created after taking snapshots of the instance. | [optional] 
**Objects** | Pointer to [**[]AwsObjectLevelParams**](AwsObjectLevelParams.md) | Specifies the objects to be protected. | [optional] 
**VolumeExclusionParams** | Pointer to [**NullableEbsVolumeExclusionParams**](EbsVolumeExclusionParams.md) |  | [optional] 

## Methods

### NewAwsSnapshotManagerObjectProtectionParams

`func NewAwsSnapshotManagerObjectProtectionParams() *AwsSnapshotManagerObjectProtectionParams`

NewAwsSnapshotManagerObjectProtectionParams instantiates a new AwsSnapshotManagerObjectProtectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsSnapshotManagerObjectProtectionParamsWithDefaults

`func NewAwsSnapshotManagerObjectProtectionParamsWithDefaults() *AwsSnapshotManagerObjectProtectionParams`

NewAwsSnapshotManagerObjectProtectionParamsWithDefaults instantiates a new AwsSnapshotManagerObjectProtectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmiCreationFrequency

`func (o *AwsSnapshotManagerObjectProtectionParams) GetAmiCreationFrequency() int32`

GetAmiCreationFrequency returns the AmiCreationFrequency field if non-nil, zero value otherwise.

### GetAmiCreationFrequencyOk

`func (o *AwsSnapshotManagerObjectProtectionParams) GetAmiCreationFrequencyOk() (*int32, bool)`

GetAmiCreationFrequencyOk returns a tuple with the AmiCreationFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmiCreationFrequency

`func (o *AwsSnapshotManagerObjectProtectionParams) SetAmiCreationFrequency(v int32)`

SetAmiCreationFrequency sets AmiCreationFrequency field to given value.

### HasAmiCreationFrequency

`func (o *AwsSnapshotManagerObjectProtectionParams) HasAmiCreationFrequency() bool`

HasAmiCreationFrequency returns a boolean if a field has been set.

### SetAmiCreationFrequencyNil

`func (o *AwsSnapshotManagerObjectProtectionParams) SetAmiCreationFrequencyNil(b bool)`

 SetAmiCreationFrequencyNil sets the value for AmiCreationFrequency to be an explicit nil

### UnsetAmiCreationFrequency
`func (o *AwsSnapshotManagerObjectProtectionParams) UnsetAmiCreationFrequency()`

UnsetAmiCreationFrequency ensures that no value is present for AmiCreationFrequency, not even an explicit nil
### GetCreateAmi

`func (o *AwsSnapshotManagerObjectProtectionParams) GetCreateAmi() bool`

GetCreateAmi returns the CreateAmi field if non-nil, zero value otherwise.

### GetCreateAmiOk

`func (o *AwsSnapshotManagerObjectProtectionParams) GetCreateAmiOk() (*bool, bool)`

GetCreateAmiOk returns a tuple with the CreateAmi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAmi

`func (o *AwsSnapshotManagerObjectProtectionParams) SetCreateAmi(v bool)`

SetCreateAmi sets CreateAmi field to given value.

### HasCreateAmi

`func (o *AwsSnapshotManagerObjectProtectionParams) HasCreateAmi() bool`

HasCreateAmi returns a boolean if a field has been set.

### SetCreateAmiNil

`func (o *AwsSnapshotManagerObjectProtectionParams) SetCreateAmiNil(b bool)`

 SetCreateAmiNil sets the value for CreateAmi to be an explicit nil

### UnsetCreateAmi
`func (o *AwsSnapshotManagerObjectProtectionParams) UnsetCreateAmi()`

UnsetCreateAmi ensures that no value is present for CreateAmi, not even an explicit nil
### GetObjects

`func (o *AwsSnapshotManagerObjectProtectionParams) GetObjects() []AwsObjectLevelParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *AwsSnapshotManagerObjectProtectionParams) GetObjectsOk() (*[]AwsObjectLevelParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *AwsSnapshotManagerObjectProtectionParams) SetObjects(v []AwsObjectLevelParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *AwsSnapshotManagerObjectProtectionParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetVolumeExclusionParams

`func (o *AwsSnapshotManagerObjectProtectionParams) GetVolumeExclusionParams() EbsVolumeExclusionParams`

GetVolumeExclusionParams returns the VolumeExclusionParams field if non-nil, zero value otherwise.

### GetVolumeExclusionParamsOk

`func (o *AwsSnapshotManagerObjectProtectionParams) GetVolumeExclusionParamsOk() (*EbsVolumeExclusionParams, bool)`

GetVolumeExclusionParamsOk returns a tuple with the VolumeExclusionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeExclusionParams

`func (o *AwsSnapshotManagerObjectProtectionParams) SetVolumeExclusionParams(v EbsVolumeExclusionParams)`

SetVolumeExclusionParams sets VolumeExclusionParams field to given value.

### HasVolumeExclusionParams

`func (o *AwsSnapshotManagerObjectProtectionParams) HasVolumeExclusionParams() bool`

HasVolumeExclusionParams returns a boolean if a field has been set.

### SetVolumeExclusionParamsNil

`func (o *AwsSnapshotManagerObjectProtectionParams) SetVolumeExclusionParamsNil(b bool)`

 SetVolumeExclusionParamsNil sets the value for VolumeExclusionParams to be an explicit nil

### UnsetVolumeExclusionParams
`func (o *AwsSnapshotManagerObjectProtectionParams) UnsetVolumeExclusionParams()`

UnsetVolumeExclusionParams ensures that no value is present for VolumeExclusionParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


