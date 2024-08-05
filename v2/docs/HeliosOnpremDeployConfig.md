# HeliosOnpremDeployConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigId** | Pointer to **NullableString** | Specifies the unique identifier for the target getting added. This field need to be passed only when helios policies are updated. | [optional] 
**CopyOnRunSuccess** | Pointer to **NullableBool** | Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. &lt;br&gt; If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. &lt;br&gt; If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group. | [optional] 
**Retention** | Pointer to [**HeliosRetention**](HeliosRetention.md) |  | [optional] 
**Schedule** | Pointer to [**HeliosTargetSchedule**](HeliosTargetSchedule.md) |  | [optional] 

## Methods

### NewHeliosOnpremDeployConfig

`func NewHeliosOnpremDeployConfig() *HeliosOnpremDeployConfig`

NewHeliosOnpremDeployConfig instantiates a new HeliosOnpremDeployConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosOnpremDeployConfigWithDefaults

`func NewHeliosOnpremDeployConfigWithDefaults() *HeliosOnpremDeployConfig`

NewHeliosOnpremDeployConfigWithDefaults instantiates a new HeliosOnpremDeployConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigId

`func (o *HeliosOnpremDeployConfig) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *HeliosOnpremDeployConfig) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *HeliosOnpremDeployConfig) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *HeliosOnpremDeployConfig) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### SetConfigIdNil

`func (o *HeliosOnpremDeployConfig) SetConfigIdNil(b bool)`

 SetConfigIdNil sets the value for ConfigId to be an explicit nil

### UnsetConfigId
`func (o *HeliosOnpremDeployConfig) UnsetConfigId()`

UnsetConfigId ensures that no value is present for ConfigId, not even an explicit nil
### GetCopyOnRunSuccess

`func (o *HeliosOnpremDeployConfig) GetCopyOnRunSuccess() bool`

GetCopyOnRunSuccess returns the CopyOnRunSuccess field if non-nil, zero value otherwise.

### GetCopyOnRunSuccessOk

`func (o *HeliosOnpremDeployConfig) GetCopyOnRunSuccessOk() (*bool, bool)`

GetCopyOnRunSuccessOk returns a tuple with the CopyOnRunSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyOnRunSuccess

`func (o *HeliosOnpremDeployConfig) SetCopyOnRunSuccess(v bool)`

SetCopyOnRunSuccess sets CopyOnRunSuccess field to given value.

### HasCopyOnRunSuccess

`func (o *HeliosOnpremDeployConfig) HasCopyOnRunSuccess() bool`

HasCopyOnRunSuccess returns a boolean if a field has been set.

### SetCopyOnRunSuccessNil

`func (o *HeliosOnpremDeployConfig) SetCopyOnRunSuccessNil(b bool)`

 SetCopyOnRunSuccessNil sets the value for CopyOnRunSuccess to be an explicit nil

### UnsetCopyOnRunSuccess
`func (o *HeliosOnpremDeployConfig) UnsetCopyOnRunSuccess()`

UnsetCopyOnRunSuccess ensures that no value is present for CopyOnRunSuccess, not even an explicit nil
### GetRetention

`func (o *HeliosOnpremDeployConfig) GetRetention() HeliosRetention`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *HeliosOnpremDeployConfig) GetRetentionOk() (*HeliosRetention, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *HeliosOnpremDeployConfig) SetRetention(v HeliosRetention)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *HeliosOnpremDeployConfig) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetSchedule

`func (o *HeliosOnpremDeployConfig) GetSchedule() HeliosTargetSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *HeliosOnpremDeployConfig) GetScheduleOk() (*HeliosTargetSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *HeliosOnpremDeployConfig) SetSchedule(v HeliosTargetSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *HeliosOnpremDeployConfig) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


