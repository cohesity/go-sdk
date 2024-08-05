# DataTieringTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertPolicy** | Pointer to [**ProtectionGroupAlertingPolicy**](ProtectionGroupAlertingPolicy.md) |  | [optional] 
**Description** | Pointer to **NullableString** | Specifies a description of the data tiering task. | [optional] 
**Id** | Pointer to **NullableString** | Specifies the id of the data tiering task. | [optional] 
**IsActive** | Pointer to **NullableBool** | Whether the data tiering task is active or not. | [optional] [default to true]
**IsDeleted** | Pointer to **NullableBool** | Tracks whether the backup job has actually been deleted. | [optional] [default to true]
**IsPaused** | Pointer to **NullableBool** | Whether the data tiering task is paused. New runs are not scheduled for the paused tasks. Active run of the task (if any) is not impacted. | [optional] [default to true]
**LastRun** | Pointer to [**DataTieringTaskRun**](DataTieringTaskRun.md) |  | [optional] 
**Name** | **NullableString** | Specifies the name of the data tiering task. | 
**Schedule** | Pointer to [**DataTieringSchedule**](DataTieringSchedule.md) |  | [optional] 
**Source** | Pointer to [**DataTieringSource**](DataTieringSource.md) |  | [optional] 
**Target** | Pointer to [**NullableDataTieringTarget**](DataTieringTarget.md) |  | [optional] 
**Type** | **NullableString** | Type of data tiering task. &#39;Downtier&#39; indicates downtiering task. &#39;Uptier&#39; indicates uptiering task. | 
**DowntieringPolicy** | Pointer to [**DowntieringPolicy**](DowntieringPolicy.md) |  | [optional] 
**UptieringPolicy** | Pointer to [**UptieringPolicy**](UptieringPolicy.md) |  | [optional] 

## Methods

### NewDataTieringTask

`func NewDataTieringTask(name NullableString, type_ NullableString, ) *DataTieringTask`

NewDataTieringTask instantiates a new DataTieringTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTieringTaskWithDefaults

`func NewDataTieringTaskWithDefaults() *DataTieringTask`

NewDataTieringTaskWithDefaults instantiates a new DataTieringTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertPolicy

`func (o *DataTieringTask) GetAlertPolicy() ProtectionGroupAlertingPolicy`

GetAlertPolicy returns the AlertPolicy field if non-nil, zero value otherwise.

### GetAlertPolicyOk

`func (o *DataTieringTask) GetAlertPolicyOk() (*ProtectionGroupAlertingPolicy, bool)`

GetAlertPolicyOk returns a tuple with the AlertPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertPolicy

`func (o *DataTieringTask) SetAlertPolicy(v ProtectionGroupAlertingPolicy)`

SetAlertPolicy sets AlertPolicy field to given value.

### HasAlertPolicy

`func (o *DataTieringTask) HasAlertPolicy() bool`

HasAlertPolicy returns a boolean if a field has been set.

### GetDescription

`func (o *DataTieringTask) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DataTieringTask) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DataTieringTask) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DataTieringTask) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DataTieringTask) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DataTieringTask) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *DataTieringTask) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataTieringTask) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataTieringTask) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataTieringTask) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DataTieringTask) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DataTieringTask) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsActive

`func (o *DataTieringTask) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *DataTieringTask) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *DataTieringTask) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *DataTieringTask) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *DataTieringTask) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *DataTieringTask) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetIsDeleted

`func (o *DataTieringTask) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *DataTieringTask) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *DataTieringTask) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *DataTieringTask) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *DataTieringTask) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *DataTieringTask) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
### GetIsPaused

`func (o *DataTieringTask) GetIsPaused() bool`

GetIsPaused returns the IsPaused field if non-nil, zero value otherwise.

### GetIsPausedOk

`func (o *DataTieringTask) GetIsPausedOk() (*bool, bool)`

GetIsPausedOk returns a tuple with the IsPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaused

`func (o *DataTieringTask) SetIsPaused(v bool)`

SetIsPaused sets IsPaused field to given value.

### HasIsPaused

`func (o *DataTieringTask) HasIsPaused() bool`

HasIsPaused returns a boolean if a field has been set.

### SetIsPausedNil

`func (o *DataTieringTask) SetIsPausedNil(b bool)`

 SetIsPausedNil sets the value for IsPaused to be an explicit nil

### UnsetIsPaused
`func (o *DataTieringTask) UnsetIsPaused()`

UnsetIsPaused ensures that no value is present for IsPaused, not even an explicit nil
### GetLastRun

`func (o *DataTieringTask) GetLastRun() DataTieringTaskRun`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *DataTieringTask) GetLastRunOk() (*DataTieringTaskRun, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *DataTieringTask) SetLastRun(v DataTieringTaskRun)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *DataTieringTask) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetName

`func (o *DataTieringTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataTieringTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataTieringTask) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *DataTieringTask) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DataTieringTask) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSchedule

`func (o *DataTieringTask) GetSchedule() DataTieringSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *DataTieringTask) GetScheduleOk() (*DataTieringSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *DataTieringTask) SetSchedule(v DataTieringSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *DataTieringTask) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetSource

`func (o *DataTieringTask) GetSource() DataTieringSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DataTieringTask) GetSourceOk() (*DataTieringSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DataTieringTask) SetSource(v DataTieringSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *DataTieringTask) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTarget

`func (o *DataTieringTask) GetTarget() DataTieringTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *DataTieringTask) GetTargetOk() (*DataTieringTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *DataTieringTask) SetTarget(v DataTieringTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *DataTieringTask) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *DataTieringTask) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *DataTieringTask) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil
### GetType

`func (o *DataTieringTask) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DataTieringTask) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DataTieringTask) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *DataTieringTask) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *DataTieringTask) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetDowntieringPolicy

`func (o *DataTieringTask) GetDowntieringPolicy() DowntieringPolicy`

GetDowntieringPolicy returns the DowntieringPolicy field if non-nil, zero value otherwise.

### GetDowntieringPolicyOk

`func (o *DataTieringTask) GetDowntieringPolicyOk() (*DowntieringPolicy, bool)`

GetDowntieringPolicyOk returns a tuple with the DowntieringPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowntieringPolicy

`func (o *DataTieringTask) SetDowntieringPolicy(v DowntieringPolicy)`

SetDowntieringPolicy sets DowntieringPolicy field to given value.

### HasDowntieringPolicy

`func (o *DataTieringTask) HasDowntieringPolicy() bool`

HasDowntieringPolicy returns a boolean if a field has been set.

### GetUptieringPolicy

`func (o *DataTieringTask) GetUptieringPolicy() UptieringPolicy`

GetUptieringPolicy returns the UptieringPolicy field if non-nil, zero value otherwise.

### GetUptieringPolicyOk

`func (o *DataTieringTask) GetUptieringPolicyOk() (*UptieringPolicy, bool)`

GetUptieringPolicyOk returns a tuple with the UptieringPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptieringPolicy

`func (o *DataTieringTask) SetUptieringPolicy(v UptieringPolicy)`

SetUptieringPolicy sets UptieringPolicy field to given value.

### HasUptieringPolicy

`func (o *DataTieringTask) HasUptieringPolicy() bool`

HasUptieringPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


