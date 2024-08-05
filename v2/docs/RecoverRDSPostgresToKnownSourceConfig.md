# RecoverRDSPostgresToKnownSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | Pointer to [**NullableRecoverRDSPostgresToKnownSourceConfigInstance**](RecoverRDSPostgresToKnownSourceConfigInstance.md) |  | [optional] 
**RecoverToNewSource** | Pointer to **NullableBool** | Specifies the parameter whether the recovery should be performed to a new target. | [optional] 
**Region** | Pointer to [**NullableRecoverRDSPostgresToKnownSourceConfigRegion**](RecoverRDSPostgresToKnownSourceConfigRegion.md) |  | [optional] 
**Source** | Pointer to [**NullableRecoverRDSPostgresToKnownSourceConfigSource**](RecoverRDSPostgresToKnownSourceConfigSource.md) |  | [optional] 

## Methods

### NewRecoverRDSPostgresToKnownSourceConfig

`func NewRecoverRDSPostgresToKnownSourceConfig() *RecoverRDSPostgresToKnownSourceConfig`

NewRecoverRDSPostgresToKnownSourceConfig instantiates a new RecoverRDSPostgresToKnownSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverRDSPostgresToKnownSourceConfigWithDefaults

`func NewRecoverRDSPostgresToKnownSourceConfigWithDefaults() *RecoverRDSPostgresToKnownSourceConfig`

NewRecoverRDSPostgresToKnownSourceConfigWithDefaults instantiates a new RecoverRDSPostgresToKnownSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *RecoverRDSPostgresToKnownSourceConfig) GetInstance() RecoverRDSPostgresToKnownSourceConfigInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *RecoverRDSPostgresToKnownSourceConfig) GetInstanceOk() (*RecoverRDSPostgresToKnownSourceConfigInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *RecoverRDSPostgresToKnownSourceConfig) SetInstance(v RecoverRDSPostgresToKnownSourceConfigInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *RecoverRDSPostgresToKnownSourceConfig) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### SetInstanceNil

`func (o *RecoverRDSPostgresToKnownSourceConfig) SetInstanceNil(b bool)`

 SetInstanceNil sets the value for Instance to be an explicit nil

### UnsetInstance
`func (o *RecoverRDSPostgresToKnownSourceConfig) UnsetInstance()`

UnsetInstance ensures that no value is present for Instance, not even an explicit nil
### GetRecoverToNewSource

`func (o *RecoverRDSPostgresToKnownSourceConfig) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *RecoverRDSPostgresToKnownSourceConfig) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *RecoverRDSPostgresToKnownSourceConfig) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.

### HasRecoverToNewSource

`func (o *RecoverRDSPostgresToKnownSourceConfig) HasRecoverToNewSource() bool`

HasRecoverToNewSource returns a boolean if a field has been set.

### SetRecoverToNewSourceNil

`func (o *RecoverRDSPostgresToKnownSourceConfig) SetRecoverToNewSourceNil(b bool)`

 SetRecoverToNewSourceNil sets the value for RecoverToNewSource to be an explicit nil

### UnsetRecoverToNewSource
`func (o *RecoverRDSPostgresToKnownSourceConfig) UnsetRecoverToNewSource()`

UnsetRecoverToNewSource ensures that no value is present for RecoverToNewSource, not even an explicit nil
### GetRegion

`func (o *RecoverRDSPostgresToKnownSourceConfig) GetRegion() RecoverRDSPostgresToKnownSourceConfigRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *RecoverRDSPostgresToKnownSourceConfig) GetRegionOk() (*RecoverRDSPostgresToKnownSourceConfigRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *RecoverRDSPostgresToKnownSourceConfig) SetRegion(v RecoverRDSPostgresToKnownSourceConfigRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *RecoverRDSPostgresToKnownSourceConfig) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *RecoverRDSPostgresToKnownSourceConfig) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *RecoverRDSPostgresToKnownSourceConfig) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetSource

`func (o *RecoverRDSPostgresToKnownSourceConfig) GetSource() RecoverRDSPostgresToKnownSourceConfigSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RecoverRDSPostgresToKnownSourceConfig) GetSourceOk() (*RecoverRDSPostgresToKnownSourceConfigSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RecoverRDSPostgresToKnownSourceConfig) SetSource(v RecoverRDSPostgresToKnownSourceConfigSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *RecoverRDSPostgresToKnownSourceConfig) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *RecoverRDSPostgresToKnownSourceConfig) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *RecoverRDSPostgresToKnownSourceConfig) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


