# SlaRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupRunType** | Pointer to **NullableString** | Specifies the type of run this rule should apply to. | [optional] 
**SlaMinutes** | Pointer to **NullableInt64** | Specifies the number of minutes allotted to a run of the specified type before SLA is considered violated. | [optional] 

## Methods

### NewSlaRule

`func NewSlaRule() *SlaRule`

NewSlaRule instantiates a new SlaRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlaRuleWithDefaults

`func NewSlaRuleWithDefaults() *SlaRule`

NewSlaRuleWithDefaults instantiates a new SlaRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupRunType

`func (o *SlaRule) GetBackupRunType() string`

GetBackupRunType returns the BackupRunType field if non-nil, zero value otherwise.

### GetBackupRunTypeOk

`func (o *SlaRule) GetBackupRunTypeOk() (*string, bool)`

GetBackupRunTypeOk returns a tuple with the BackupRunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRunType

`func (o *SlaRule) SetBackupRunType(v string)`

SetBackupRunType sets BackupRunType field to given value.

### HasBackupRunType

`func (o *SlaRule) HasBackupRunType() bool`

HasBackupRunType returns a boolean if a field has been set.

### SetBackupRunTypeNil

`func (o *SlaRule) SetBackupRunTypeNil(b bool)`

 SetBackupRunTypeNil sets the value for BackupRunType to be an explicit nil

### UnsetBackupRunType
`func (o *SlaRule) UnsetBackupRunType()`

UnsetBackupRunType ensures that no value is present for BackupRunType, not even an explicit nil
### GetSlaMinutes

`func (o *SlaRule) GetSlaMinutes() int64`

GetSlaMinutes returns the SlaMinutes field if non-nil, zero value otherwise.

### GetSlaMinutesOk

`func (o *SlaRule) GetSlaMinutesOk() (*int64, bool)`

GetSlaMinutesOk returns a tuple with the SlaMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaMinutes

`func (o *SlaRule) SetSlaMinutes(v int64)`

SetSlaMinutes sets SlaMinutes field to given value.

### HasSlaMinutes

`func (o *SlaRule) HasSlaMinutes() bool`

HasSlaMinutes returns a boolean if a field has been set.

### SetSlaMinutesNil

`func (o *SlaRule) SetSlaMinutesNil(b bool)`

 SetSlaMinutesNil sets the value for SlaMinutes to be an explicit nil

### UnsetSlaMinutes
`func (o *SlaRule) UnsetSlaMinutes()`

UnsetSlaMinutes ensures that no value is present for SlaMinutes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


