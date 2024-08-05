# AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | Pointer to [**NullableRecoverRDSPostgresToKnownSourceConfigInstance**](RecoverRDSPostgresToKnownSourceConfigInstance.md) |  | [optional] 
**RecoverToNewSource** | Pointer to **NullableBool** | Specifies the parameter whether the recovery should be performed to a new target. | [optional] 
**Region** | Pointer to [**NullableRecoverRDSPostgresToKnownSourceConfigRegion**](RecoverRDSPostgresToKnownSourceConfigRegion.md) |  | [optional] 
**Source** | Pointer to [**NullableRecoverRDSPostgresToKnownSourceConfigSource**](RecoverRDSPostgresToKnownSourceConfigSource.md) |  | [optional] 

## Methods

### NewAwsTargetParamsForRecoverRDSPostgresKnownSourceConfig

`func NewAwsTargetParamsForRecoverRDSPostgresKnownSourceConfig() *AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig`

NewAwsTargetParamsForRecoverRDSPostgresKnownSourceConfig instantiates a new AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsTargetParamsForRecoverRDSPostgresKnownSourceConfigWithDefaults

`func NewAwsTargetParamsForRecoverRDSPostgresKnownSourceConfigWithDefaults() *AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig`

NewAwsTargetParamsForRecoverRDSPostgresKnownSourceConfigWithDefaults instantiates a new AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig) GetInstance() RecoverRDSPostgresToKnownSourceConfigInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig) GetInstanceOk() (*RecoverRDSPostgresToKnownSourceConfigInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig) SetInstance(v RecoverRDSPostgresToKnownSourceConfigInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### SetInstanceNil

`func (o *AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig) SetInstanceNil(b bool)`

 SetInstanceNil sets the value for Instance to be an explicit nil

### UnsetInstance
`func (o *AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig) UnsetInstance()`

UnsetInstance ensures that no value is present for Instance, not even an explicit nil
### GetRecoverToNewSource

`func (o *AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.

### HasRecoverToNewSource

`func (o *AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig) HasRecoverToNewSource() bool`

HasRecoverToNewSource returns a boolean if a field has been set.

### SetRecoverToNewSourceNil

`func (o *AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig) SetRecoverToNewSourceNil(b bool)`

 SetRecoverToNewSourceNil sets the value for RecoverToNewSource to be an explicit nil

### UnsetRecoverToNewSource
`func (o *AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig) UnsetRecoverToNewSource()`

UnsetRecoverToNewSource ensures that no value is present for RecoverToNewSource, not even an explicit nil
### GetRegion

`func (o *AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig) GetRegion() RecoverRDSPostgresToKnownSourceConfigRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig) GetRegionOk() (*RecoverRDSPostgresToKnownSourceConfigRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig) SetRegion(v RecoverRDSPostgresToKnownSourceConfigRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetSource

`func (o *AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig) GetSource() RecoverRDSPostgresToKnownSourceConfigSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig) GetSourceOk() (*RecoverRDSPostgresToKnownSourceConfigSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig) SetSource(v RecoverRDSPostgresToKnownSourceConfigSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


