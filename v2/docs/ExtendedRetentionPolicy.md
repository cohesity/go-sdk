# ExtendedRetentionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigId** | Pointer to **NullableString** | Specifies the unique identifier for the target getting added. This field need to be passed olny when policies are updated. | [optional] 
**Retention** | [**Retention**](Retention.md) |  | 
**RunType** | Pointer to **NullableString** | The backup run type to which this extended retention applies to. If this is not set, the extended retention will be applicable to all non-log backup types. Currently, the only value that can be set here is Full. &#39;Regular&#39; indicates a incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a Regular schedule captures all the blocks. &#39;Full&#39; indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized. &#39;Log&#39; indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time. &#39;System&#39; indicates a system backup. System backups are used to do bare metal recovery of the system to a specific point in time. | [optional] 
**Schedule** | [**ExtendedRetentionSchedule**](ExtendedRetentionSchedule.md) |  | 

## Methods

### NewExtendedRetentionPolicy

`func NewExtendedRetentionPolicy(retention Retention, schedule ExtendedRetentionSchedule, ) *ExtendedRetentionPolicy`

NewExtendedRetentionPolicy instantiates a new ExtendedRetentionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedRetentionPolicyWithDefaults

`func NewExtendedRetentionPolicyWithDefaults() *ExtendedRetentionPolicy`

NewExtendedRetentionPolicyWithDefaults instantiates a new ExtendedRetentionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigId

`func (o *ExtendedRetentionPolicy) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *ExtendedRetentionPolicy) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *ExtendedRetentionPolicy) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *ExtendedRetentionPolicy) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### SetConfigIdNil

`func (o *ExtendedRetentionPolicy) SetConfigIdNil(b bool)`

 SetConfigIdNil sets the value for ConfigId to be an explicit nil

### UnsetConfigId
`func (o *ExtendedRetentionPolicy) UnsetConfigId()`

UnsetConfigId ensures that no value is present for ConfigId, not even an explicit nil
### GetRetention

`func (o *ExtendedRetentionPolicy) GetRetention() Retention`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *ExtendedRetentionPolicy) GetRetentionOk() (*Retention, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *ExtendedRetentionPolicy) SetRetention(v Retention)`

SetRetention sets Retention field to given value.


### GetRunType

`func (o *ExtendedRetentionPolicy) GetRunType() string`

GetRunType returns the RunType field if non-nil, zero value otherwise.

### GetRunTypeOk

`func (o *ExtendedRetentionPolicy) GetRunTypeOk() (*string, bool)`

GetRunTypeOk returns a tuple with the RunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunType

`func (o *ExtendedRetentionPolicy) SetRunType(v string)`

SetRunType sets RunType field to given value.

### HasRunType

`func (o *ExtendedRetentionPolicy) HasRunType() bool`

HasRunType returns a boolean if a field has been set.

### SetRunTypeNil

`func (o *ExtendedRetentionPolicy) SetRunTypeNil(b bool)`

 SetRunTypeNil sets the value for RunType to be an explicit nil

### UnsetRunType
`func (o *ExtendedRetentionPolicy) UnsetRunType()`

UnsetRunType ensures that no value is present for RunType, not even an explicit nil
### GetSchedule

`func (o *ExtendedRetentionPolicy) GetSchedule() ExtendedRetentionSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ExtendedRetentionPolicy) GetScheduleOk() (*ExtendedRetentionSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ExtendedRetentionPolicy) SetSchedule(v ExtendedRetentionSchedule)`

SetSchedule sets Schedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


