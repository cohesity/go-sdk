# ThrottlingPolicyParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnforceMaxStreams** | Pointer to **NullableBool** | Specifies whether datastore streams are configured for all datastores that are part of the registered entity. If set to true, number of streams from Cohesity cluster to the registered entity will be limited to the value set for maxConcurrentStreams. If not set or set to false, there is no max limit for the number of concurrent streams. | [optional] 
**EnforceRegisteredSourceMaxBackups** | Pointer to **NullableBool** | Specifies whether no. of backups are configured for the registered entity. If set to true, number of backups made by Cohesity cluster in the registered entity will be limited to the value set for RegisteredSourceMaxConcurrentBackups. If not set or set to false, there is no max limit for the number of concurrent backups. | [optional] 
**IsEnabled** | Pointer to **NullableBool** | Indicates whether read operations to the datastores, which are part of the registered Protection Source, are throttled. | [optional] 
**LatencyThresholds** | Pointer to [**LatencyThresholds**](LatencyThresholds.md) |  | [optional] 
**MaxConcurrentStreams** | Pointer to **NullableInt32** | Specifies the limit on the number of streams Cohesity cluster will make concurrently to the datastores of the registered entity. This limit is enforced only when the flag enforceMaxStreams is set to true. | [optional] 
**RegisteredSourceMaxConcurrentBackups** | Pointer to **NullableInt32** | Specifies the limit on the number of backups Cohesity cluster will make concurrently to the registered entity. This limit is enforced only when the flag enforceRegisteredSourceMaxBackups is set to true. | [optional] 

## Methods

### NewThrottlingPolicyParameters

`func NewThrottlingPolicyParameters() *ThrottlingPolicyParameters`

NewThrottlingPolicyParameters instantiates a new ThrottlingPolicyParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThrottlingPolicyParametersWithDefaults

`func NewThrottlingPolicyParametersWithDefaults() *ThrottlingPolicyParameters`

NewThrottlingPolicyParametersWithDefaults instantiates a new ThrottlingPolicyParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnforceMaxStreams

`func (o *ThrottlingPolicyParameters) GetEnforceMaxStreams() bool`

GetEnforceMaxStreams returns the EnforceMaxStreams field if non-nil, zero value otherwise.

### GetEnforceMaxStreamsOk

`func (o *ThrottlingPolicyParameters) GetEnforceMaxStreamsOk() (*bool, bool)`

GetEnforceMaxStreamsOk returns a tuple with the EnforceMaxStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceMaxStreams

`func (o *ThrottlingPolicyParameters) SetEnforceMaxStreams(v bool)`

SetEnforceMaxStreams sets EnforceMaxStreams field to given value.

### HasEnforceMaxStreams

`func (o *ThrottlingPolicyParameters) HasEnforceMaxStreams() bool`

HasEnforceMaxStreams returns a boolean if a field has been set.

### SetEnforceMaxStreamsNil

`func (o *ThrottlingPolicyParameters) SetEnforceMaxStreamsNil(b bool)`

 SetEnforceMaxStreamsNil sets the value for EnforceMaxStreams to be an explicit nil

### UnsetEnforceMaxStreams
`func (o *ThrottlingPolicyParameters) UnsetEnforceMaxStreams()`

UnsetEnforceMaxStreams ensures that no value is present for EnforceMaxStreams, not even an explicit nil
### GetEnforceRegisteredSourceMaxBackups

`func (o *ThrottlingPolicyParameters) GetEnforceRegisteredSourceMaxBackups() bool`

GetEnforceRegisteredSourceMaxBackups returns the EnforceRegisteredSourceMaxBackups field if non-nil, zero value otherwise.

### GetEnforceRegisteredSourceMaxBackupsOk

`func (o *ThrottlingPolicyParameters) GetEnforceRegisteredSourceMaxBackupsOk() (*bool, bool)`

GetEnforceRegisteredSourceMaxBackupsOk returns a tuple with the EnforceRegisteredSourceMaxBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceRegisteredSourceMaxBackups

`func (o *ThrottlingPolicyParameters) SetEnforceRegisteredSourceMaxBackups(v bool)`

SetEnforceRegisteredSourceMaxBackups sets EnforceRegisteredSourceMaxBackups field to given value.

### HasEnforceRegisteredSourceMaxBackups

`func (o *ThrottlingPolicyParameters) HasEnforceRegisteredSourceMaxBackups() bool`

HasEnforceRegisteredSourceMaxBackups returns a boolean if a field has been set.

### SetEnforceRegisteredSourceMaxBackupsNil

`func (o *ThrottlingPolicyParameters) SetEnforceRegisteredSourceMaxBackupsNil(b bool)`

 SetEnforceRegisteredSourceMaxBackupsNil sets the value for EnforceRegisteredSourceMaxBackups to be an explicit nil

### UnsetEnforceRegisteredSourceMaxBackups
`func (o *ThrottlingPolicyParameters) UnsetEnforceRegisteredSourceMaxBackups()`

UnsetEnforceRegisteredSourceMaxBackups ensures that no value is present for EnforceRegisteredSourceMaxBackups, not even an explicit nil
### GetIsEnabled

`func (o *ThrottlingPolicyParameters) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ThrottlingPolicyParameters) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ThrottlingPolicyParameters) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *ThrottlingPolicyParameters) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *ThrottlingPolicyParameters) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *ThrottlingPolicyParameters) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetLatencyThresholds

`func (o *ThrottlingPolicyParameters) GetLatencyThresholds() LatencyThresholds`

GetLatencyThresholds returns the LatencyThresholds field if non-nil, zero value otherwise.

### GetLatencyThresholdsOk

`func (o *ThrottlingPolicyParameters) GetLatencyThresholdsOk() (*LatencyThresholds, bool)`

GetLatencyThresholdsOk returns a tuple with the LatencyThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyThresholds

`func (o *ThrottlingPolicyParameters) SetLatencyThresholds(v LatencyThresholds)`

SetLatencyThresholds sets LatencyThresholds field to given value.

### HasLatencyThresholds

`func (o *ThrottlingPolicyParameters) HasLatencyThresholds() bool`

HasLatencyThresholds returns a boolean if a field has been set.

### GetMaxConcurrentStreams

`func (o *ThrottlingPolicyParameters) GetMaxConcurrentStreams() int32`

GetMaxConcurrentStreams returns the MaxConcurrentStreams field if non-nil, zero value otherwise.

### GetMaxConcurrentStreamsOk

`func (o *ThrottlingPolicyParameters) GetMaxConcurrentStreamsOk() (*int32, bool)`

GetMaxConcurrentStreamsOk returns a tuple with the MaxConcurrentStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentStreams

`func (o *ThrottlingPolicyParameters) SetMaxConcurrentStreams(v int32)`

SetMaxConcurrentStreams sets MaxConcurrentStreams field to given value.

### HasMaxConcurrentStreams

`func (o *ThrottlingPolicyParameters) HasMaxConcurrentStreams() bool`

HasMaxConcurrentStreams returns a boolean if a field has been set.

### SetMaxConcurrentStreamsNil

`func (o *ThrottlingPolicyParameters) SetMaxConcurrentStreamsNil(b bool)`

 SetMaxConcurrentStreamsNil sets the value for MaxConcurrentStreams to be an explicit nil

### UnsetMaxConcurrentStreams
`func (o *ThrottlingPolicyParameters) UnsetMaxConcurrentStreams()`

UnsetMaxConcurrentStreams ensures that no value is present for MaxConcurrentStreams, not even an explicit nil
### GetRegisteredSourceMaxConcurrentBackups

`func (o *ThrottlingPolicyParameters) GetRegisteredSourceMaxConcurrentBackups() int32`

GetRegisteredSourceMaxConcurrentBackups returns the RegisteredSourceMaxConcurrentBackups field if non-nil, zero value otherwise.

### GetRegisteredSourceMaxConcurrentBackupsOk

`func (o *ThrottlingPolicyParameters) GetRegisteredSourceMaxConcurrentBackupsOk() (*int32, bool)`

GetRegisteredSourceMaxConcurrentBackupsOk returns a tuple with the RegisteredSourceMaxConcurrentBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredSourceMaxConcurrentBackups

`func (o *ThrottlingPolicyParameters) SetRegisteredSourceMaxConcurrentBackups(v int32)`

SetRegisteredSourceMaxConcurrentBackups sets RegisteredSourceMaxConcurrentBackups field to given value.

### HasRegisteredSourceMaxConcurrentBackups

`func (o *ThrottlingPolicyParameters) HasRegisteredSourceMaxConcurrentBackups() bool`

HasRegisteredSourceMaxConcurrentBackups returns a boolean if a field has been set.

### SetRegisteredSourceMaxConcurrentBackupsNil

`func (o *ThrottlingPolicyParameters) SetRegisteredSourceMaxConcurrentBackupsNil(b bool)`

 SetRegisteredSourceMaxConcurrentBackupsNil sets the value for RegisteredSourceMaxConcurrentBackups to be an explicit nil

### UnsetRegisteredSourceMaxConcurrentBackups
`func (o *ThrottlingPolicyParameters) UnsetRegisteredSourceMaxConcurrentBackups()`

UnsetRegisteredSourceMaxConcurrentBackups ensures that no value is present for RegisteredSourceMaxConcurrentBackups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


