# ReplicationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedule** | [**TargetSchedule**](TargetSchedule.md) |  | 
**Retention** | [**Retention**](Retention.md) |  | 
**CopyOnRunSuccess** | Pointer to **NullableBool** | Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. &lt;br&gt; If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. &lt;br&gt; If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group. | [optional] 
**ConfigId** | Pointer to **NullableString** | Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated. | [optional] 
**TargetType** | **string** | Specifies the type of target to which replication need to be performed. | 
**RemoteTargetConfig** | Pointer to [**RemoteTargetConfig**](RemoteTargetConfig.md) |  | [optional] 
**AwsTargetConfig** | Pointer to [**AWSTargetConfig**](AWSTargetConfig.md) |  | [optional] 
**AzureTargetConfig** | Pointer to [**AzureTargetConfig**](AzureTargetConfig.md) |  | [optional] 

## Methods

### NewReplicationConfig

`func NewReplicationConfig(schedule TargetSchedule, retention Retention, targetType string, ) *ReplicationConfig`

NewReplicationConfig instantiates a new ReplicationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationConfigWithDefaults

`func NewReplicationConfigWithDefaults() *ReplicationConfig`

NewReplicationConfigWithDefaults instantiates a new ReplicationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedule

`func (o *ReplicationConfig) GetSchedule() TargetSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ReplicationConfig) GetScheduleOk() (*TargetSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ReplicationConfig) SetSchedule(v TargetSchedule)`

SetSchedule sets Schedule field to given value.


### GetRetention

`func (o *ReplicationConfig) GetRetention() Retention`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *ReplicationConfig) GetRetentionOk() (*Retention, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *ReplicationConfig) SetRetention(v Retention)`

SetRetention sets Retention field to given value.


### GetCopyOnRunSuccess

`func (o *ReplicationConfig) GetCopyOnRunSuccess() bool`

GetCopyOnRunSuccess returns the CopyOnRunSuccess field if non-nil, zero value otherwise.

### GetCopyOnRunSuccessOk

`func (o *ReplicationConfig) GetCopyOnRunSuccessOk() (*bool, bool)`

GetCopyOnRunSuccessOk returns a tuple with the CopyOnRunSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyOnRunSuccess

`func (o *ReplicationConfig) SetCopyOnRunSuccess(v bool)`

SetCopyOnRunSuccess sets CopyOnRunSuccess field to given value.

### HasCopyOnRunSuccess

`func (o *ReplicationConfig) HasCopyOnRunSuccess() bool`

HasCopyOnRunSuccess returns a boolean if a field has been set.

### SetCopyOnRunSuccessNil

`func (o *ReplicationConfig) SetCopyOnRunSuccessNil(b bool)`

 SetCopyOnRunSuccessNil sets the value for CopyOnRunSuccess to be an explicit nil

### UnsetCopyOnRunSuccess
`func (o *ReplicationConfig) UnsetCopyOnRunSuccess()`

UnsetCopyOnRunSuccess ensures that no value is present for CopyOnRunSuccess, not even an explicit nil
### GetConfigId

`func (o *ReplicationConfig) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *ReplicationConfig) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *ReplicationConfig) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *ReplicationConfig) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### SetConfigIdNil

`func (o *ReplicationConfig) SetConfigIdNil(b bool)`

 SetConfigIdNil sets the value for ConfigId to be an explicit nil

### UnsetConfigId
`func (o *ReplicationConfig) UnsetConfigId()`

UnsetConfigId ensures that no value is present for ConfigId, not even an explicit nil
### GetTargetType

`func (o *ReplicationConfig) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *ReplicationConfig) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *ReplicationConfig) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.


### GetRemoteTargetConfig

`func (o *ReplicationConfig) GetRemoteTargetConfig() RemoteTargetConfig`

GetRemoteTargetConfig returns the RemoteTargetConfig field if non-nil, zero value otherwise.

### GetRemoteTargetConfigOk

`func (o *ReplicationConfig) GetRemoteTargetConfigOk() (*RemoteTargetConfig, bool)`

GetRemoteTargetConfigOk returns a tuple with the RemoteTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteTargetConfig

`func (o *ReplicationConfig) SetRemoteTargetConfig(v RemoteTargetConfig)`

SetRemoteTargetConfig sets RemoteTargetConfig field to given value.

### HasRemoteTargetConfig

`func (o *ReplicationConfig) HasRemoteTargetConfig() bool`

HasRemoteTargetConfig returns a boolean if a field has been set.

### GetAwsTargetConfig

`func (o *ReplicationConfig) GetAwsTargetConfig() AWSTargetConfig`

GetAwsTargetConfig returns the AwsTargetConfig field if non-nil, zero value otherwise.

### GetAwsTargetConfigOk

`func (o *ReplicationConfig) GetAwsTargetConfigOk() (*AWSTargetConfig, bool)`

GetAwsTargetConfigOk returns a tuple with the AwsTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsTargetConfig

`func (o *ReplicationConfig) SetAwsTargetConfig(v AWSTargetConfig)`

SetAwsTargetConfig sets AwsTargetConfig field to given value.

### HasAwsTargetConfig

`func (o *ReplicationConfig) HasAwsTargetConfig() bool`

HasAwsTargetConfig returns a boolean if a field has been set.

### GetAzureTargetConfig

`func (o *ReplicationConfig) GetAzureTargetConfig() AzureTargetConfig`

GetAzureTargetConfig returns the AzureTargetConfig field if non-nil, zero value otherwise.

### GetAzureTargetConfigOk

`func (o *ReplicationConfig) GetAzureTargetConfigOk() (*AzureTargetConfig, bool)`

GetAzureTargetConfigOk returns a tuple with the AzureTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTargetConfig

`func (o *ReplicationConfig) SetAzureTargetConfig(v AzureTargetConfig)`

SetAzureTargetConfig sets AzureTargetConfig field to given value.

### HasAzureTargetConfig

`func (o *ReplicationConfig) HasAzureTargetConfig() bool`

HasAzureTargetConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


