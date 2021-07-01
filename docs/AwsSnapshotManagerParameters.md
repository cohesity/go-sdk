# AwsSnapshotManagerParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmiCreationFrequency** | Pointer to **NullableInt32** | Specifies the frequency of AMI creation. This should be set if the option to create AMI is set. A value of n creates an AMI from the snapshots after every n runs. eg. n &#x3D; 2 implies every alternate backup run starting from the first will create an AMI. | [optional] 
**CreateAmi** | Pointer to **NullableBool** | If true, creates an AMI after taking snapshots of the instance. It should be set only while backing up EC2 instances. CreateAmi creates AMI for the protection job. | [optional] 

## Methods

### NewAwsSnapshotManagerParameters

`func NewAwsSnapshotManagerParameters() *AwsSnapshotManagerParameters`

NewAwsSnapshotManagerParameters instantiates a new AwsSnapshotManagerParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsSnapshotManagerParametersWithDefaults

`func NewAwsSnapshotManagerParametersWithDefaults() *AwsSnapshotManagerParameters`

NewAwsSnapshotManagerParametersWithDefaults instantiates a new AwsSnapshotManagerParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmiCreationFrequency

`func (o *AwsSnapshotManagerParameters) GetAmiCreationFrequency() int32`

GetAmiCreationFrequency returns the AmiCreationFrequency field if non-nil, zero value otherwise.

### GetAmiCreationFrequencyOk

`func (o *AwsSnapshotManagerParameters) GetAmiCreationFrequencyOk() (*int32, bool)`

GetAmiCreationFrequencyOk returns a tuple with the AmiCreationFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmiCreationFrequency

`func (o *AwsSnapshotManagerParameters) SetAmiCreationFrequency(v int32)`

SetAmiCreationFrequency sets AmiCreationFrequency field to given value.

### HasAmiCreationFrequency

`func (o *AwsSnapshotManagerParameters) HasAmiCreationFrequency() bool`

HasAmiCreationFrequency returns a boolean if a field has been set.

### SetAmiCreationFrequencyNil

`func (o *AwsSnapshotManagerParameters) SetAmiCreationFrequencyNil(b bool)`

 SetAmiCreationFrequencyNil sets the value for AmiCreationFrequency to be an explicit nil

### UnsetAmiCreationFrequency
`func (o *AwsSnapshotManagerParameters) UnsetAmiCreationFrequency()`

UnsetAmiCreationFrequency ensures that no value is present for AmiCreationFrequency, not even an explicit nil
### GetCreateAmi

`func (o *AwsSnapshotManagerParameters) GetCreateAmi() bool`

GetCreateAmi returns the CreateAmi field if non-nil, zero value otherwise.

### GetCreateAmiOk

`func (o *AwsSnapshotManagerParameters) GetCreateAmiOk() (*bool, bool)`

GetCreateAmiOk returns a tuple with the CreateAmi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAmi

`func (o *AwsSnapshotManagerParameters) SetCreateAmi(v bool)`

SetCreateAmi sets CreateAmi field to given value.

### HasCreateAmi

`func (o *AwsSnapshotManagerParameters) HasCreateAmi() bool`

HasCreateAmi returns a boolean if a field has been set.

### SetCreateAmiNil

`func (o *AwsSnapshotManagerParameters) SetCreateAmiNil(b bool)`

 SetCreateAmiNil sets the value for CreateAmi to be an explicit nil

### UnsetCreateAmi
`func (o *AwsSnapshotManagerParameters) UnsetCreateAmi()`

UnsetCreateAmi ensures that no value is present for CreateAmi, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


