# TargetsConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchivalTargets** | Pointer to [**[]ArchivalConfig**](ArchivalConfig.md) |  | [optional] 
**CloudSpinTargets** | Pointer to [**[]CloudSpinConfig**](CloudSpinConfig.md) |  | [optional] 
**OnpremDeployTargets** | Pointer to [**[]OnpremDeployConfig**](OnpremDeployConfig.md) |  | [optional] 
**ReplicationTargets** | Pointer to [**[]ReplicationConfig**](ReplicationConfig.md) |  | [optional] 
**RpaasTargets** | Pointer to [**[]RpaasConfig**](RpaasConfig.md) |  | [optional] 

## Methods

### NewTargetsConfiguration

`func NewTargetsConfiguration() *TargetsConfiguration`

NewTargetsConfiguration instantiates a new TargetsConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetsConfigurationWithDefaults

`func NewTargetsConfigurationWithDefaults() *TargetsConfiguration`

NewTargetsConfigurationWithDefaults instantiates a new TargetsConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivalTargets

`func (o *TargetsConfiguration) GetArchivalTargets() []ArchivalConfig`

GetArchivalTargets returns the ArchivalTargets field if non-nil, zero value otherwise.

### GetArchivalTargetsOk

`func (o *TargetsConfiguration) GetArchivalTargetsOk() (*[]ArchivalConfig, bool)`

GetArchivalTargetsOk returns a tuple with the ArchivalTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalTargets

`func (o *TargetsConfiguration) SetArchivalTargets(v []ArchivalConfig)`

SetArchivalTargets sets ArchivalTargets field to given value.

### HasArchivalTargets

`func (o *TargetsConfiguration) HasArchivalTargets() bool`

HasArchivalTargets returns a boolean if a field has been set.

### GetCloudSpinTargets

`func (o *TargetsConfiguration) GetCloudSpinTargets() []CloudSpinConfig`

GetCloudSpinTargets returns the CloudSpinTargets field if non-nil, zero value otherwise.

### GetCloudSpinTargetsOk

`func (o *TargetsConfiguration) GetCloudSpinTargetsOk() (*[]CloudSpinConfig, bool)`

GetCloudSpinTargetsOk returns a tuple with the CloudSpinTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudSpinTargets

`func (o *TargetsConfiguration) SetCloudSpinTargets(v []CloudSpinConfig)`

SetCloudSpinTargets sets CloudSpinTargets field to given value.

### HasCloudSpinTargets

`func (o *TargetsConfiguration) HasCloudSpinTargets() bool`

HasCloudSpinTargets returns a boolean if a field has been set.

### GetOnpremDeployTargets

`func (o *TargetsConfiguration) GetOnpremDeployTargets() []OnpremDeployConfig`

GetOnpremDeployTargets returns the OnpremDeployTargets field if non-nil, zero value otherwise.

### GetOnpremDeployTargetsOk

`func (o *TargetsConfiguration) GetOnpremDeployTargetsOk() (*[]OnpremDeployConfig, bool)`

GetOnpremDeployTargetsOk returns a tuple with the OnpremDeployTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnpremDeployTargets

`func (o *TargetsConfiguration) SetOnpremDeployTargets(v []OnpremDeployConfig)`

SetOnpremDeployTargets sets OnpremDeployTargets field to given value.

### HasOnpremDeployTargets

`func (o *TargetsConfiguration) HasOnpremDeployTargets() bool`

HasOnpremDeployTargets returns a boolean if a field has been set.

### GetReplicationTargets

`func (o *TargetsConfiguration) GetReplicationTargets() []ReplicationConfig`

GetReplicationTargets returns the ReplicationTargets field if non-nil, zero value otherwise.

### GetReplicationTargetsOk

`func (o *TargetsConfiguration) GetReplicationTargetsOk() (*[]ReplicationConfig, bool)`

GetReplicationTargetsOk returns a tuple with the ReplicationTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationTargets

`func (o *TargetsConfiguration) SetReplicationTargets(v []ReplicationConfig)`

SetReplicationTargets sets ReplicationTargets field to given value.

### HasReplicationTargets

`func (o *TargetsConfiguration) HasReplicationTargets() bool`

HasReplicationTargets returns a boolean if a field has been set.

### GetRpaasTargets

`func (o *TargetsConfiguration) GetRpaasTargets() []RpaasConfig`

GetRpaasTargets returns the RpaasTargets field if non-nil, zero value otherwise.

### GetRpaasTargetsOk

`func (o *TargetsConfiguration) GetRpaasTargetsOk() (*[]RpaasConfig, bool)`

GetRpaasTargetsOk returns a tuple with the RpaasTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpaasTargets

`func (o *TargetsConfiguration) SetRpaasTargets(v []RpaasConfig)`

SetRpaasTargets sets RpaasTargets field to given value.

### HasRpaasTargets

`func (o *TargetsConfiguration) HasRpaasTargets() bool`

HasRpaasTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


