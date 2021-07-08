# RpoPolicySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertingConfig** | Pointer to [**AlertingConfig**](AlertingConfig.md) |  | [optional] 
**AlertingPolicy** | Pointer to **[]string** | Array of Job Events.  During Job Runs, the following Job Events are generated: 1) Job succeeds 2) Job fails 3) Job violates the SLA These Job Events can cause Alerts to be generated. &#39;kSuccess&#39; means the Protection Job succeeded. &#39;kFailure&#39; means the Protection Job failed. &#39;kSlaViolation&#39; means the Protection Job took longer than the time period specified in the SLA. | [optional] 
**EnvironmentTypeJobParams** | Pointer to [**EnvironmentTypeJobParameters**](EnvironmentTypeJobParameters.md) |  | [optional] 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**QosType** | Pointer to **NullableString** | Specifies the QoS policy type to use. &#39;kBackupHDD&#39; indicates the Cohesity Cluster writes data directly to the HDD tier for this Protection Job. This is the recommended setting. &#39;kBackupSSD&#39; indicates the Cohesity Cluster writes data directly to the SSD tier for this Protection Job. Only specify this policy if you need fast ingest speed for a small number of Protection Jobs. | [optional] 
**StorageDomainId** | Pointer to **NullableInt64** | Specifies the Storage Domain to which data will be written. | [optional] 

## Methods

### NewRpoPolicySettings

`func NewRpoPolicySettings() *RpoPolicySettings`

NewRpoPolicySettings instantiates a new RpoPolicySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpoPolicySettingsWithDefaults

`func NewRpoPolicySettingsWithDefaults() *RpoPolicySettings`

NewRpoPolicySettingsWithDefaults instantiates a new RpoPolicySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertingConfig

`func (o *RpoPolicySettings) GetAlertingConfig() AlertingConfig`

GetAlertingConfig returns the AlertingConfig field if non-nil, zero value otherwise.

### GetAlertingConfigOk

`func (o *RpoPolicySettings) GetAlertingConfigOk() (*AlertingConfig, bool)`

GetAlertingConfigOk returns a tuple with the AlertingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingConfig

`func (o *RpoPolicySettings) SetAlertingConfig(v AlertingConfig)`

SetAlertingConfig sets AlertingConfig field to given value.

### HasAlertingConfig

`func (o *RpoPolicySettings) HasAlertingConfig() bool`

HasAlertingConfig returns a boolean if a field has been set.

### GetAlertingPolicy

`func (o *RpoPolicySettings) GetAlertingPolicy() []string`

GetAlertingPolicy returns the AlertingPolicy field if non-nil, zero value otherwise.

### GetAlertingPolicyOk

`func (o *RpoPolicySettings) GetAlertingPolicyOk() (*[]string, bool)`

GetAlertingPolicyOk returns a tuple with the AlertingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingPolicy

`func (o *RpoPolicySettings) SetAlertingPolicy(v []string)`

SetAlertingPolicy sets AlertingPolicy field to given value.

### HasAlertingPolicy

`func (o *RpoPolicySettings) HasAlertingPolicy() bool`

HasAlertingPolicy returns a boolean if a field has been set.

### SetAlertingPolicyNil

`func (o *RpoPolicySettings) SetAlertingPolicyNil(b bool)`

 SetAlertingPolicyNil sets the value for AlertingPolicy to be an explicit nil

### UnsetAlertingPolicy
`func (o *RpoPolicySettings) UnsetAlertingPolicy()`

UnsetAlertingPolicy ensures that no value is present for AlertingPolicy, not even an explicit nil
### GetEnvironmentTypeJobParams

`func (o *RpoPolicySettings) GetEnvironmentTypeJobParams() EnvironmentTypeJobParameters`

GetEnvironmentTypeJobParams returns the EnvironmentTypeJobParams field if non-nil, zero value otherwise.

### GetEnvironmentTypeJobParamsOk

`func (o *RpoPolicySettings) GetEnvironmentTypeJobParamsOk() (*EnvironmentTypeJobParameters, bool)`

GetEnvironmentTypeJobParamsOk returns a tuple with the EnvironmentTypeJobParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentTypeJobParams

`func (o *RpoPolicySettings) SetEnvironmentTypeJobParams(v EnvironmentTypeJobParameters)`

SetEnvironmentTypeJobParams sets EnvironmentTypeJobParams field to given value.

### HasEnvironmentTypeJobParams

`func (o *RpoPolicySettings) HasEnvironmentTypeJobParams() bool`

HasEnvironmentTypeJobParams returns a boolean if a field has been set.

### GetIndexingPolicy

`func (o *RpoPolicySettings) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *RpoPolicySettings) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *RpoPolicySettings) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *RpoPolicySettings) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetQosType

`func (o *RpoPolicySettings) GetQosType() string`

GetQosType returns the QosType field if non-nil, zero value otherwise.

### GetQosTypeOk

`func (o *RpoPolicySettings) GetQosTypeOk() (*string, bool)`

GetQosTypeOk returns a tuple with the QosType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosType

`func (o *RpoPolicySettings) SetQosType(v string)`

SetQosType sets QosType field to given value.

### HasQosType

`func (o *RpoPolicySettings) HasQosType() bool`

HasQosType returns a boolean if a field has been set.

### SetQosTypeNil

`func (o *RpoPolicySettings) SetQosTypeNil(b bool)`

 SetQosTypeNil sets the value for QosType to be an explicit nil

### UnsetQosType
`func (o *RpoPolicySettings) UnsetQosType()`

UnsetQosType ensures that no value is present for QosType, not even an explicit nil
### GetStorageDomainId

`func (o *RpoPolicySettings) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *RpoPolicySettings) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *RpoPolicySettings) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *RpoPolicySettings) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *RpoPolicySettings) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *RpoPolicySettings) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


