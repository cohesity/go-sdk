# AWSSnapshotManagerParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmiCreationFrequency** | Pointer to **NullableInt32** | The frequency of AMI creation. This should be set if the option to create AMI is set. A value of n creates an AMI from the snapshots after every n runs. eg. n &#x3D; 2 implies every alternate backup run starting from the first will create an AMI. | [optional] 
**CreateAmiForRun** | Pointer to **NullableBool** | Whether we need to create an AMI for this run. | [optional] 
**ShouldCreateAmi** | Pointer to **NullableBool** | Whether we need to create an AMI after taking snapshots of the instance. | [optional] 

## Methods

### NewAWSSnapshotManagerParams

`func NewAWSSnapshotManagerParams() *AWSSnapshotManagerParams`

NewAWSSnapshotManagerParams instantiates a new AWSSnapshotManagerParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSSnapshotManagerParamsWithDefaults

`func NewAWSSnapshotManagerParamsWithDefaults() *AWSSnapshotManagerParams`

NewAWSSnapshotManagerParamsWithDefaults instantiates a new AWSSnapshotManagerParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmiCreationFrequency

`func (o *AWSSnapshotManagerParams) GetAmiCreationFrequency() int32`

GetAmiCreationFrequency returns the AmiCreationFrequency field if non-nil, zero value otherwise.

### GetAmiCreationFrequencyOk

`func (o *AWSSnapshotManagerParams) GetAmiCreationFrequencyOk() (*int32, bool)`

GetAmiCreationFrequencyOk returns a tuple with the AmiCreationFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmiCreationFrequency

`func (o *AWSSnapshotManagerParams) SetAmiCreationFrequency(v int32)`

SetAmiCreationFrequency sets AmiCreationFrequency field to given value.

### HasAmiCreationFrequency

`func (o *AWSSnapshotManagerParams) HasAmiCreationFrequency() bool`

HasAmiCreationFrequency returns a boolean if a field has been set.

### SetAmiCreationFrequencyNil

`func (o *AWSSnapshotManagerParams) SetAmiCreationFrequencyNil(b bool)`

 SetAmiCreationFrequencyNil sets the value for AmiCreationFrequency to be an explicit nil

### UnsetAmiCreationFrequency
`func (o *AWSSnapshotManagerParams) UnsetAmiCreationFrequency()`

UnsetAmiCreationFrequency ensures that no value is present for AmiCreationFrequency, not even an explicit nil
### GetCreateAmiForRun

`func (o *AWSSnapshotManagerParams) GetCreateAmiForRun() bool`

GetCreateAmiForRun returns the CreateAmiForRun field if non-nil, zero value otherwise.

### GetCreateAmiForRunOk

`func (o *AWSSnapshotManagerParams) GetCreateAmiForRunOk() (*bool, bool)`

GetCreateAmiForRunOk returns a tuple with the CreateAmiForRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAmiForRun

`func (o *AWSSnapshotManagerParams) SetCreateAmiForRun(v bool)`

SetCreateAmiForRun sets CreateAmiForRun field to given value.

### HasCreateAmiForRun

`func (o *AWSSnapshotManagerParams) HasCreateAmiForRun() bool`

HasCreateAmiForRun returns a boolean if a field has been set.

### SetCreateAmiForRunNil

`func (o *AWSSnapshotManagerParams) SetCreateAmiForRunNil(b bool)`

 SetCreateAmiForRunNil sets the value for CreateAmiForRun to be an explicit nil

### UnsetCreateAmiForRun
`func (o *AWSSnapshotManagerParams) UnsetCreateAmiForRun()`

UnsetCreateAmiForRun ensures that no value is present for CreateAmiForRun, not even an explicit nil
### GetShouldCreateAmi

`func (o *AWSSnapshotManagerParams) GetShouldCreateAmi() bool`

GetShouldCreateAmi returns the ShouldCreateAmi field if non-nil, zero value otherwise.

### GetShouldCreateAmiOk

`func (o *AWSSnapshotManagerParams) GetShouldCreateAmiOk() (*bool, bool)`

GetShouldCreateAmiOk returns a tuple with the ShouldCreateAmi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCreateAmi

`func (o *AWSSnapshotManagerParams) SetShouldCreateAmi(v bool)`

SetShouldCreateAmi sets ShouldCreateAmi field to given value.

### HasShouldCreateAmi

`func (o *AWSSnapshotManagerParams) HasShouldCreateAmi() bool`

HasShouldCreateAmi returns a boolean if a field has been set.

### SetShouldCreateAmiNil

`func (o *AWSSnapshotManagerParams) SetShouldCreateAmiNil(b bool)`

 SetShouldCreateAmiNil sets the value for ShouldCreateAmi to be an explicit nil

### UnsetShouldCreateAmi
`func (o *AWSSnapshotManagerParams) UnsetShouldCreateAmi()`

UnsetShouldCreateAmi ensures that no value is present for ShouldCreateAmi, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


