# AwsSnapshotManagerParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmiCreationFrequency** | Pointer to **NullableInt32** | Specifies the frequency of AMI creation. This should be set if the option to create AMI is set. A value of n creates an AMI from the snapshots after every n runs. eg. n &#x3D; 2 implies every alternate backup run starting from the first will create an AMI. | [optional] 
**CreateAmi** | Pointer to **NullableBool** | If true, creates an AMI after taking snapshots of the instance. It should be set only while backing up EC2 instances. CreateAmi creates AMI for the protection job. | [optional] 
**VolumeExclusionParams** | Pointer to [**NullableEbsVolumeExclusionParams**](EbsVolumeExclusionParams.md) |  | [optional] 

## Methods

### NewAwsSnapshotManagerParams

`func NewAwsSnapshotManagerParams() *AwsSnapshotManagerParams`

NewAwsSnapshotManagerParams instantiates a new AwsSnapshotManagerParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsSnapshotManagerParamsWithDefaults

`func NewAwsSnapshotManagerParamsWithDefaults() *AwsSnapshotManagerParams`

NewAwsSnapshotManagerParamsWithDefaults instantiates a new AwsSnapshotManagerParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmiCreationFrequency

`func (o *AwsSnapshotManagerParams) GetAmiCreationFrequency() int32`

GetAmiCreationFrequency returns the AmiCreationFrequency field if non-nil, zero value otherwise.

### GetAmiCreationFrequencyOk

`func (o *AwsSnapshotManagerParams) GetAmiCreationFrequencyOk() (*int32, bool)`

GetAmiCreationFrequencyOk returns a tuple with the AmiCreationFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmiCreationFrequency

`func (o *AwsSnapshotManagerParams) SetAmiCreationFrequency(v int32)`

SetAmiCreationFrequency sets AmiCreationFrequency field to given value.

### HasAmiCreationFrequency

`func (o *AwsSnapshotManagerParams) HasAmiCreationFrequency() bool`

HasAmiCreationFrequency returns a boolean if a field has been set.

### SetAmiCreationFrequencyNil

`func (o *AwsSnapshotManagerParams) SetAmiCreationFrequencyNil(b bool)`

 SetAmiCreationFrequencyNil sets the value for AmiCreationFrequency to be an explicit nil

### UnsetAmiCreationFrequency
`func (o *AwsSnapshotManagerParams) UnsetAmiCreationFrequency()`

UnsetAmiCreationFrequency ensures that no value is present for AmiCreationFrequency, not even an explicit nil
### GetCreateAmi

`func (o *AwsSnapshotManagerParams) GetCreateAmi() bool`

GetCreateAmi returns the CreateAmi field if non-nil, zero value otherwise.

### GetCreateAmiOk

`func (o *AwsSnapshotManagerParams) GetCreateAmiOk() (*bool, bool)`

GetCreateAmiOk returns a tuple with the CreateAmi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAmi

`func (o *AwsSnapshotManagerParams) SetCreateAmi(v bool)`

SetCreateAmi sets CreateAmi field to given value.

### HasCreateAmi

`func (o *AwsSnapshotManagerParams) HasCreateAmi() bool`

HasCreateAmi returns a boolean if a field has been set.

### SetCreateAmiNil

`func (o *AwsSnapshotManagerParams) SetCreateAmiNil(b bool)`

 SetCreateAmiNil sets the value for CreateAmi to be an explicit nil

### UnsetCreateAmi
`func (o *AwsSnapshotManagerParams) UnsetCreateAmi()`

UnsetCreateAmi ensures that no value is present for CreateAmi, not even an explicit nil
### GetVolumeExclusionParams

`func (o *AwsSnapshotManagerParams) GetVolumeExclusionParams() EbsVolumeExclusionParams`

GetVolumeExclusionParams returns the VolumeExclusionParams field if non-nil, zero value otherwise.

### GetVolumeExclusionParamsOk

`func (o *AwsSnapshotManagerParams) GetVolumeExclusionParamsOk() (*EbsVolumeExclusionParams, bool)`

GetVolumeExclusionParamsOk returns a tuple with the VolumeExclusionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeExclusionParams

`func (o *AwsSnapshotManagerParams) SetVolumeExclusionParams(v EbsVolumeExclusionParams)`

SetVolumeExclusionParams sets VolumeExclusionParams field to given value.

### HasVolumeExclusionParams

`func (o *AwsSnapshotManagerParams) HasVolumeExclusionParams() bool`

HasVolumeExclusionParams returns a boolean if a field has been set.

### SetVolumeExclusionParamsNil

`func (o *AwsSnapshotManagerParams) SetVolumeExclusionParamsNil(b bool)`

 SetVolumeExclusionParamsNil sets the value for VolumeExclusionParams to be an explicit nil

### UnsetVolumeExclusionParams
`func (o *AwsSnapshotManagerParams) UnsetVolumeExclusionParams()`

UnsetVolumeExclusionParams ensures that no value is present for VolumeExclusionParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


