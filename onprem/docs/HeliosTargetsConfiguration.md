# HeliosTargetsConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReplicationTargets** | Pointer to [**[]HeliosReplicationConfig**](HeliosReplicationConfig.md) |  | [optional] 
**ArchivalTargets** | Pointer to [**[]HeliosArchivalConfig**](HeliosArchivalConfig.md) |  | [optional] 
**CloudSpinTargets** | Pointer to [**[]HeliosCloudSpinConfig**](HeliosCloudSpinConfig.md) |  | [optional] 

## Methods

### NewHeliosTargetsConfiguration

`func NewHeliosTargetsConfiguration() *HeliosTargetsConfiguration`

NewHeliosTargetsConfiguration instantiates a new HeliosTargetsConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosTargetsConfigurationWithDefaults

`func NewHeliosTargetsConfigurationWithDefaults() *HeliosTargetsConfiguration`

NewHeliosTargetsConfigurationWithDefaults instantiates a new HeliosTargetsConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplicationTargets

`func (o *HeliosTargetsConfiguration) GetReplicationTargets() []HeliosReplicationConfig`

GetReplicationTargets returns the ReplicationTargets field if non-nil, zero value otherwise.

### GetReplicationTargetsOk

`func (o *HeliosTargetsConfiguration) GetReplicationTargetsOk() (*[]HeliosReplicationConfig, bool)`

GetReplicationTargetsOk returns a tuple with the ReplicationTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationTargets

`func (o *HeliosTargetsConfiguration) SetReplicationTargets(v []HeliosReplicationConfig)`

SetReplicationTargets sets ReplicationTargets field to given value.

### HasReplicationTargets

`func (o *HeliosTargetsConfiguration) HasReplicationTargets() bool`

HasReplicationTargets returns a boolean if a field has been set.

### GetArchivalTargets

`func (o *HeliosTargetsConfiguration) GetArchivalTargets() []HeliosArchivalConfig`

GetArchivalTargets returns the ArchivalTargets field if non-nil, zero value otherwise.

### GetArchivalTargetsOk

`func (o *HeliosTargetsConfiguration) GetArchivalTargetsOk() (*[]HeliosArchivalConfig, bool)`

GetArchivalTargetsOk returns a tuple with the ArchivalTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalTargets

`func (o *HeliosTargetsConfiguration) SetArchivalTargets(v []HeliosArchivalConfig)`

SetArchivalTargets sets ArchivalTargets field to given value.

### HasArchivalTargets

`func (o *HeliosTargetsConfiguration) HasArchivalTargets() bool`

HasArchivalTargets returns a boolean if a field has been set.

### GetCloudSpinTargets

`func (o *HeliosTargetsConfiguration) GetCloudSpinTargets() []HeliosCloudSpinConfig`

GetCloudSpinTargets returns the CloudSpinTargets field if non-nil, zero value otherwise.

### GetCloudSpinTargetsOk

`func (o *HeliosTargetsConfiguration) GetCloudSpinTargetsOk() (*[]HeliosCloudSpinConfig, bool)`

GetCloudSpinTargetsOk returns a tuple with the CloudSpinTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudSpinTargets

`func (o *HeliosTargetsConfiguration) SetCloudSpinTargets(v []HeliosCloudSpinConfig)`

SetCloudSpinTargets sets CloudSpinTargets field to given value.

### HasCloudSpinTargets

`func (o *HeliosTargetsConfiguration) HasCloudSpinTargets() bool`

HasCloudSpinTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


