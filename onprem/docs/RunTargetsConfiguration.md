# RunTargetsConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsePolicyDefaults** | Pointer to **NullableBool** | Specifies whether to use default policy settings or not. If specified as true then &#39;replications&#39; and &#39;arcihvals&#39; should not be specified. In case of true value, replicatioan targets congfigured in the policy will be added internally. | [optional] [default to false]
**Replications** | Pointer to [**[]RunReplicationConfig**](RunReplicationConfig.md) | Specifies a list of replication targets configurations. | [optional] 
**Archivals** | Pointer to [**[]RunArchivalConfig**](RunArchivalConfig.md) | Specifies a list of archival targets configurations. | [optional] 

## Methods

### NewRunTargetsConfiguration

`func NewRunTargetsConfiguration() *RunTargetsConfiguration`

NewRunTargetsConfiguration instantiates a new RunTargetsConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunTargetsConfigurationWithDefaults

`func NewRunTargetsConfigurationWithDefaults() *RunTargetsConfiguration`

NewRunTargetsConfigurationWithDefaults instantiates a new RunTargetsConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsePolicyDefaults

`func (o *RunTargetsConfiguration) GetUsePolicyDefaults() bool`

GetUsePolicyDefaults returns the UsePolicyDefaults field if non-nil, zero value otherwise.

### GetUsePolicyDefaultsOk

`func (o *RunTargetsConfiguration) GetUsePolicyDefaultsOk() (*bool, bool)`

GetUsePolicyDefaultsOk returns a tuple with the UsePolicyDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePolicyDefaults

`func (o *RunTargetsConfiguration) SetUsePolicyDefaults(v bool)`

SetUsePolicyDefaults sets UsePolicyDefaults field to given value.

### HasUsePolicyDefaults

`func (o *RunTargetsConfiguration) HasUsePolicyDefaults() bool`

HasUsePolicyDefaults returns a boolean if a field has been set.

### SetUsePolicyDefaultsNil

`func (o *RunTargetsConfiguration) SetUsePolicyDefaultsNil(b bool)`

 SetUsePolicyDefaultsNil sets the value for UsePolicyDefaults to be an explicit nil

### UnsetUsePolicyDefaults
`func (o *RunTargetsConfiguration) UnsetUsePolicyDefaults()`

UnsetUsePolicyDefaults ensures that no value is present for UsePolicyDefaults, not even an explicit nil
### GetReplications

`func (o *RunTargetsConfiguration) GetReplications() []RunReplicationConfig`

GetReplications returns the Replications field if non-nil, zero value otherwise.

### GetReplicationsOk

`func (o *RunTargetsConfiguration) GetReplicationsOk() (*[]RunReplicationConfig, bool)`

GetReplicationsOk returns a tuple with the Replications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplications

`func (o *RunTargetsConfiguration) SetReplications(v []RunReplicationConfig)`

SetReplications sets Replications field to given value.

### HasReplications

`func (o *RunTargetsConfiguration) HasReplications() bool`

HasReplications returns a boolean if a field has been set.

### SetReplicationsNil

`func (o *RunTargetsConfiguration) SetReplicationsNil(b bool)`

 SetReplicationsNil sets the value for Replications to be an explicit nil

### UnsetReplications
`func (o *RunTargetsConfiguration) UnsetReplications()`

UnsetReplications ensures that no value is present for Replications, not even an explicit nil
### GetArchivals

`func (o *RunTargetsConfiguration) GetArchivals() []RunArchivalConfig`

GetArchivals returns the Archivals field if non-nil, zero value otherwise.

### GetArchivalsOk

`func (o *RunTargetsConfiguration) GetArchivalsOk() (*[]RunArchivalConfig, bool)`

GetArchivalsOk returns a tuple with the Archivals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivals

`func (o *RunTargetsConfiguration) SetArchivals(v []RunArchivalConfig)`

SetArchivals sets Archivals field to given value.

### HasArchivals

`func (o *RunTargetsConfiguration) HasArchivals() bool`

HasArchivals returns a boolean if a field has been set.

### SetArchivalsNil

`func (o *RunTargetsConfiguration) SetArchivalsNil(b bool)`

 SetArchivalsNil sets the value for Archivals to be an explicit nil

### UnsetArchivals
`func (o *RunTargetsConfiguration) UnsetArchivals()`

UnsetArchivals ensures that no value is present for Archivals, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


