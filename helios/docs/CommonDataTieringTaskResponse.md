# CommonDataTieringTaskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Specifies the id of the data tiering task. | [optional] 
**Name** | **NullableString** | Specifies the name of the data tiering task. | 
**Description** | Pointer to **NullableString** | Specifies a description of the data tiering task. | [optional] 
**AlertPolicy** | Pointer to [**ProtectionGroupAlertingPolicy**](ProtectionGroupAlertingPolicy.md) |  | [optional] 
**Source** | Pointer to [**DataTieringSource**](DataTieringSource.md) |  | [optional] 
**Target** | Pointer to [**NullableDataTieringTarget**](DataTieringTarget.md) |  | [optional] 
**Schedule** | Pointer to [**DataTieringSchedule**](DataTieringSchedule.md) |  | [optional] 
**Type** | **NullableString** | Type of data tiering task. &#39;Downtier&#39; indicates downtiering task. &#39;Uptier&#39; indicates uptiering task. | 
**LastRun** | Pointer to [**DataTieringTaskRun**](DataTieringTaskRun.md) |  | [optional] 
**IsActive** | Pointer to **NullableBool** | Whether the data tiering task is active or not. | [optional] [default to true]
**IsPaused** | Pointer to **NullableBool** | Whether the data tiering task is paused. New runs are not scheduled for the paused tasks. Active run of the task (if any) is not impacted. | [optional] [default to true]
**IsDeleted** | Pointer to **NullableBool** | Tracks whether the backup job has actually been deleted. | [optional] [default to true]

## Methods

### NewCommonDataTieringTaskResponse

`func NewCommonDataTieringTaskResponse(name NullableString, type_ NullableString, ) *CommonDataTieringTaskResponse`

NewCommonDataTieringTaskResponse instantiates a new CommonDataTieringTaskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonDataTieringTaskResponseWithDefaults

`func NewCommonDataTieringTaskResponseWithDefaults() *CommonDataTieringTaskResponse`

NewCommonDataTieringTaskResponseWithDefaults instantiates a new CommonDataTieringTaskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommonDataTieringTaskResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommonDataTieringTaskResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommonDataTieringTaskResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CommonDataTieringTaskResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CommonDataTieringTaskResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CommonDataTieringTaskResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *CommonDataTieringTaskResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommonDataTieringTaskResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommonDataTieringTaskResponse) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CommonDataTieringTaskResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CommonDataTieringTaskResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *CommonDataTieringTaskResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CommonDataTieringTaskResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CommonDataTieringTaskResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CommonDataTieringTaskResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CommonDataTieringTaskResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CommonDataTieringTaskResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAlertPolicy

`func (o *CommonDataTieringTaskResponse) GetAlertPolicy() ProtectionGroupAlertingPolicy`

GetAlertPolicy returns the AlertPolicy field if non-nil, zero value otherwise.

### GetAlertPolicyOk

`func (o *CommonDataTieringTaskResponse) GetAlertPolicyOk() (*ProtectionGroupAlertingPolicy, bool)`

GetAlertPolicyOk returns a tuple with the AlertPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertPolicy

`func (o *CommonDataTieringTaskResponse) SetAlertPolicy(v ProtectionGroupAlertingPolicy)`

SetAlertPolicy sets AlertPolicy field to given value.

### HasAlertPolicy

`func (o *CommonDataTieringTaskResponse) HasAlertPolicy() bool`

HasAlertPolicy returns a boolean if a field has been set.

### GetSource

`func (o *CommonDataTieringTaskResponse) GetSource() DataTieringSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CommonDataTieringTaskResponse) GetSourceOk() (*DataTieringSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CommonDataTieringTaskResponse) SetSource(v DataTieringSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *CommonDataTieringTaskResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTarget

`func (o *CommonDataTieringTaskResponse) GetTarget() DataTieringTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CommonDataTieringTaskResponse) GetTargetOk() (*DataTieringTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CommonDataTieringTaskResponse) SetTarget(v DataTieringTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CommonDataTieringTaskResponse) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *CommonDataTieringTaskResponse) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *CommonDataTieringTaskResponse) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil
### GetSchedule

`func (o *CommonDataTieringTaskResponse) GetSchedule() DataTieringSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CommonDataTieringTaskResponse) GetScheduleOk() (*DataTieringSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CommonDataTieringTaskResponse) SetSchedule(v DataTieringSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CommonDataTieringTaskResponse) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetType

`func (o *CommonDataTieringTaskResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CommonDataTieringTaskResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CommonDataTieringTaskResponse) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CommonDataTieringTaskResponse) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CommonDataTieringTaskResponse) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetLastRun

`func (o *CommonDataTieringTaskResponse) GetLastRun() DataTieringTaskRun`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *CommonDataTieringTaskResponse) GetLastRunOk() (*DataTieringTaskRun, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *CommonDataTieringTaskResponse) SetLastRun(v DataTieringTaskRun)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *CommonDataTieringTaskResponse) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetIsActive

`func (o *CommonDataTieringTaskResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CommonDataTieringTaskResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CommonDataTieringTaskResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *CommonDataTieringTaskResponse) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *CommonDataTieringTaskResponse) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *CommonDataTieringTaskResponse) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetIsPaused

`func (o *CommonDataTieringTaskResponse) GetIsPaused() bool`

GetIsPaused returns the IsPaused field if non-nil, zero value otherwise.

### GetIsPausedOk

`func (o *CommonDataTieringTaskResponse) GetIsPausedOk() (*bool, bool)`

GetIsPausedOk returns a tuple with the IsPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaused

`func (o *CommonDataTieringTaskResponse) SetIsPaused(v bool)`

SetIsPaused sets IsPaused field to given value.

### HasIsPaused

`func (o *CommonDataTieringTaskResponse) HasIsPaused() bool`

HasIsPaused returns a boolean if a field has been set.

### SetIsPausedNil

`func (o *CommonDataTieringTaskResponse) SetIsPausedNil(b bool)`

 SetIsPausedNil sets the value for IsPaused to be an explicit nil

### UnsetIsPaused
`func (o *CommonDataTieringTaskResponse) UnsetIsPaused()`

UnsetIsPaused ensures that no value is present for IsPaused, not even an explicit nil
### GetIsDeleted

`func (o *CommonDataTieringTaskResponse) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *CommonDataTieringTaskResponse) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *CommonDataTieringTaskResponse) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *CommonDataTieringTaskResponse) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *CommonDataTieringTaskResponse) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *CommonDataTieringTaskResponse) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


