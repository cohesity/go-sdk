# SanEnvJobParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxSnapshotsOnPrimary** | Pointer to **NullableInt64** | Specifies how many recent snapshots of each backed up entity to retain on the primary environment. If not specified, then snapshots will not be be deleted from the primary environment. | [optional] 

## Methods

### NewSanEnvJobParameters

`func NewSanEnvJobParameters() *SanEnvJobParameters`

NewSanEnvJobParameters instantiates a new SanEnvJobParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSanEnvJobParametersWithDefaults

`func NewSanEnvJobParametersWithDefaults() *SanEnvJobParameters`

NewSanEnvJobParametersWithDefaults instantiates a new SanEnvJobParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxSnapshotsOnPrimary

`func (o *SanEnvJobParameters) GetMaxSnapshotsOnPrimary() int64`

GetMaxSnapshotsOnPrimary returns the MaxSnapshotsOnPrimary field if non-nil, zero value otherwise.

### GetMaxSnapshotsOnPrimaryOk

`func (o *SanEnvJobParameters) GetMaxSnapshotsOnPrimaryOk() (*int64, bool)`

GetMaxSnapshotsOnPrimaryOk returns a tuple with the MaxSnapshotsOnPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSnapshotsOnPrimary

`func (o *SanEnvJobParameters) SetMaxSnapshotsOnPrimary(v int64)`

SetMaxSnapshotsOnPrimary sets MaxSnapshotsOnPrimary field to given value.

### HasMaxSnapshotsOnPrimary

`func (o *SanEnvJobParameters) HasMaxSnapshotsOnPrimary() bool`

HasMaxSnapshotsOnPrimary returns a boolean if a field has been set.

### SetMaxSnapshotsOnPrimaryNil

`func (o *SanEnvJobParameters) SetMaxSnapshotsOnPrimaryNil(b bool)`

 SetMaxSnapshotsOnPrimaryNil sets the value for MaxSnapshotsOnPrimary to be an explicit nil

### UnsetMaxSnapshotsOnPrimary
`func (o *SanEnvJobParameters) UnsetMaxSnapshotsOnPrimary()`

UnsetMaxSnapshotsOnPrimary ensures that no value is present for MaxSnapshotsOnPrimary, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


