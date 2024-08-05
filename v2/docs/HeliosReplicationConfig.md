# HeliosReplicationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigId** | Pointer to **NullableString** | Specifies the unique identifier for the target getting added. This field need to be passed only when helios policies are updated. | [optional] 
**CopyOnRunSuccess** | Pointer to **NullableBool** | Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. &lt;br&gt; If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. &lt;br&gt; If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group. | [optional] 
**Retention** | Pointer to [**HeliosRetention**](HeliosRetention.md) |  | [optional] 
**Schedule** | Pointer to [**HeliosTargetSchedule**](HeliosTargetSchedule.md) |  | [optional] 
**AwsTargetConfig** | Pointer to [**HeliosAWSTargetConfig**](HeliosAWSTargetConfig.md) |  | [optional] 
**AzureTargetConfig** | Pointer to [**HeliosAzureTargetConfig**](HeliosAzureTargetConfig.md) |  | [optional] 
**RemoteTargetConfig** | Pointer to [**HeliosRemoteTargetConfig**](HeliosRemoteTargetConfig.md) |  | [optional] 
**TargetType** | Pointer to **NullableString** | Specifies the type of target to which replication need to be performed. | [optional] 

## Methods

### NewHeliosReplicationConfig

`func NewHeliosReplicationConfig() *HeliosReplicationConfig`

NewHeliosReplicationConfig instantiates a new HeliosReplicationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosReplicationConfigWithDefaults

`func NewHeliosReplicationConfigWithDefaults() *HeliosReplicationConfig`

NewHeliosReplicationConfigWithDefaults instantiates a new HeliosReplicationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigId

`func (o *HeliosReplicationConfig) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *HeliosReplicationConfig) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *HeliosReplicationConfig) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *HeliosReplicationConfig) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### SetConfigIdNil

`func (o *HeliosReplicationConfig) SetConfigIdNil(b bool)`

 SetConfigIdNil sets the value for ConfigId to be an explicit nil

### UnsetConfigId
`func (o *HeliosReplicationConfig) UnsetConfigId()`

UnsetConfigId ensures that no value is present for ConfigId, not even an explicit nil
### GetCopyOnRunSuccess

`func (o *HeliosReplicationConfig) GetCopyOnRunSuccess() bool`

GetCopyOnRunSuccess returns the CopyOnRunSuccess field if non-nil, zero value otherwise.

### GetCopyOnRunSuccessOk

`func (o *HeliosReplicationConfig) GetCopyOnRunSuccessOk() (*bool, bool)`

GetCopyOnRunSuccessOk returns a tuple with the CopyOnRunSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyOnRunSuccess

`func (o *HeliosReplicationConfig) SetCopyOnRunSuccess(v bool)`

SetCopyOnRunSuccess sets CopyOnRunSuccess field to given value.

### HasCopyOnRunSuccess

`func (o *HeliosReplicationConfig) HasCopyOnRunSuccess() bool`

HasCopyOnRunSuccess returns a boolean if a field has been set.

### SetCopyOnRunSuccessNil

`func (o *HeliosReplicationConfig) SetCopyOnRunSuccessNil(b bool)`

 SetCopyOnRunSuccessNil sets the value for CopyOnRunSuccess to be an explicit nil

### UnsetCopyOnRunSuccess
`func (o *HeliosReplicationConfig) UnsetCopyOnRunSuccess()`

UnsetCopyOnRunSuccess ensures that no value is present for CopyOnRunSuccess, not even an explicit nil
### GetRetention

`func (o *HeliosReplicationConfig) GetRetention() HeliosRetention`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *HeliosReplicationConfig) GetRetentionOk() (*HeliosRetention, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *HeliosReplicationConfig) SetRetention(v HeliosRetention)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *HeliosReplicationConfig) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetSchedule

`func (o *HeliosReplicationConfig) GetSchedule() HeliosTargetSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *HeliosReplicationConfig) GetScheduleOk() (*HeliosTargetSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *HeliosReplicationConfig) SetSchedule(v HeliosTargetSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *HeliosReplicationConfig) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetAwsTargetConfig

`func (o *HeliosReplicationConfig) GetAwsTargetConfig() HeliosAWSTargetConfig`

GetAwsTargetConfig returns the AwsTargetConfig field if non-nil, zero value otherwise.

### GetAwsTargetConfigOk

`func (o *HeliosReplicationConfig) GetAwsTargetConfigOk() (*HeliosAWSTargetConfig, bool)`

GetAwsTargetConfigOk returns a tuple with the AwsTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsTargetConfig

`func (o *HeliosReplicationConfig) SetAwsTargetConfig(v HeliosAWSTargetConfig)`

SetAwsTargetConfig sets AwsTargetConfig field to given value.

### HasAwsTargetConfig

`func (o *HeliosReplicationConfig) HasAwsTargetConfig() bool`

HasAwsTargetConfig returns a boolean if a field has been set.

### GetAzureTargetConfig

`func (o *HeliosReplicationConfig) GetAzureTargetConfig() HeliosAzureTargetConfig`

GetAzureTargetConfig returns the AzureTargetConfig field if non-nil, zero value otherwise.

### GetAzureTargetConfigOk

`func (o *HeliosReplicationConfig) GetAzureTargetConfigOk() (*HeliosAzureTargetConfig, bool)`

GetAzureTargetConfigOk returns a tuple with the AzureTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTargetConfig

`func (o *HeliosReplicationConfig) SetAzureTargetConfig(v HeliosAzureTargetConfig)`

SetAzureTargetConfig sets AzureTargetConfig field to given value.

### HasAzureTargetConfig

`func (o *HeliosReplicationConfig) HasAzureTargetConfig() bool`

HasAzureTargetConfig returns a boolean if a field has been set.

### GetRemoteTargetConfig

`func (o *HeliosReplicationConfig) GetRemoteTargetConfig() HeliosRemoteTargetConfig`

GetRemoteTargetConfig returns the RemoteTargetConfig field if non-nil, zero value otherwise.

### GetRemoteTargetConfigOk

`func (o *HeliosReplicationConfig) GetRemoteTargetConfigOk() (*HeliosRemoteTargetConfig, bool)`

GetRemoteTargetConfigOk returns a tuple with the RemoteTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteTargetConfig

`func (o *HeliosReplicationConfig) SetRemoteTargetConfig(v HeliosRemoteTargetConfig)`

SetRemoteTargetConfig sets RemoteTargetConfig field to given value.

### HasRemoteTargetConfig

`func (o *HeliosReplicationConfig) HasRemoteTargetConfig() bool`

HasRemoteTargetConfig returns a boolean if a field has been set.

### GetTargetType

`func (o *HeliosReplicationConfig) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *HeliosReplicationConfig) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *HeliosReplicationConfig) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *HeliosReplicationConfig) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *HeliosReplicationConfig) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *HeliosReplicationConfig) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


