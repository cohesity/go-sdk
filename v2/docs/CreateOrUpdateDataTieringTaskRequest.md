# CreateOrUpdateDataTieringTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertPolicy** | Pointer to [**ProtectionGroupAlertingPolicy**](ProtectionGroupAlertingPolicy.md) |  | [optional] 
**Description** | Pointer to **NullableString** | Specifies a description of the data tiering task. | [optional] 
**Name** | **NullableString** | Specifies the name of the data tiering task. | 
**Schedule** | Pointer to [**DataTieringSchedule**](DataTieringSchedule.md) |  | [optional] 
**Source** | Pointer to [**DataTieringSource**](DataTieringSource.md) |  | [optional] 
**Target** | Pointer to [**NullableDataTieringTarget**](DataTieringTarget.md) |  | [optional] 
**Type** | **NullableString** | Type of data tiering task. &#39;Downtier&#39; indicates downtiering task. &#39;Uptier&#39; indicates uptiering task. | 
**DowntieringPolicy** | Pointer to [**DowntieringPolicy**](DowntieringPolicy.md) |  | [optional] 
**UptieringPolicy** | Pointer to [**UptieringPolicy**](UptieringPolicy.md) |  | [optional] 

## Methods

### NewCreateOrUpdateDataTieringTaskRequest

`func NewCreateOrUpdateDataTieringTaskRequest(name NullableString, type_ NullableString, ) *CreateOrUpdateDataTieringTaskRequest`

NewCreateOrUpdateDataTieringTaskRequest instantiates a new CreateOrUpdateDataTieringTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateDataTieringTaskRequestWithDefaults

`func NewCreateOrUpdateDataTieringTaskRequestWithDefaults() *CreateOrUpdateDataTieringTaskRequest`

NewCreateOrUpdateDataTieringTaskRequestWithDefaults instantiates a new CreateOrUpdateDataTieringTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertPolicy

`func (o *CreateOrUpdateDataTieringTaskRequest) GetAlertPolicy() ProtectionGroupAlertingPolicy`

GetAlertPolicy returns the AlertPolicy field if non-nil, zero value otherwise.

### GetAlertPolicyOk

`func (o *CreateOrUpdateDataTieringTaskRequest) GetAlertPolicyOk() (*ProtectionGroupAlertingPolicy, bool)`

GetAlertPolicyOk returns a tuple with the AlertPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertPolicy

`func (o *CreateOrUpdateDataTieringTaskRequest) SetAlertPolicy(v ProtectionGroupAlertingPolicy)`

SetAlertPolicy sets AlertPolicy field to given value.

### HasAlertPolicy

`func (o *CreateOrUpdateDataTieringTaskRequest) HasAlertPolicy() bool`

HasAlertPolicy returns a boolean if a field has been set.

### GetDescription

`func (o *CreateOrUpdateDataTieringTaskRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOrUpdateDataTieringTaskRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOrUpdateDataTieringTaskRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateOrUpdateDataTieringTaskRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateOrUpdateDataTieringTaskRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateOrUpdateDataTieringTaskRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *CreateOrUpdateDataTieringTaskRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrUpdateDataTieringTaskRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrUpdateDataTieringTaskRequest) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateOrUpdateDataTieringTaskRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateOrUpdateDataTieringTaskRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSchedule

`func (o *CreateOrUpdateDataTieringTaskRequest) GetSchedule() DataTieringSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CreateOrUpdateDataTieringTaskRequest) GetScheduleOk() (*DataTieringSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CreateOrUpdateDataTieringTaskRequest) SetSchedule(v DataTieringSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CreateOrUpdateDataTieringTaskRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetSource

`func (o *CreateOrUpdateDataTieringTaskRequest) GetSource() DataTieringSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CreateOrUpdateDataTieringTaskRequest) GetSourceOk() (*DataTieringSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CreateOrUpdateDataTieringTaskRequest) SetSource(v DataTieringSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *CreateOrUpdateDataTieringTaskRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTarget

`func (o *CreateOrUpdateDataTieringTaskRequest) GetTarget() DataTieringTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CreateOrUpdateDataTieringTaskRequest) GetTargetOk() (*DataTieringTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CreateOrUpdateDataTieringTaskRequest) SetTarget(v DataTieringTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CreateOrUpdateDataTieringTaskRequest) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *CreateOrUpdateDataTieringTaskRequest) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *CreateOrUpdateDataTieringTaskRequest) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil
### GetType

`func (o *CreateOrUpdateDataTieringTaskRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateOrUpdateDataTieringTaskRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateOrUpdateDataTieringTaskRequest) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CreateOrUpdateDataTieringTaskRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CreateOrUpdateDataTieringTaskRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetDowntieringPolicy

`func (o *CreateOrUpdateDataTieringTaskRequest) GetDowntieringPolicy() DowntieringPolicy`

GetDowntieringPolicy returns the DowntieringPolicy field if non-nil, zero value otherwise.

### GetDowntieringPolicyOk

`func (o *CreateOrUpdateDataTieringTaskRequest) GetDowntieringPolicyOk() (*DowntieringPolicy, bool)`

GetDowntieringPolicyOk returns a tuple with the DowntieringPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowntieringPolicy

`func (o *CreateOrUpdateDataTieringTaskRequest) SetDowntieringPolicy(v DowntieringPolicy)`

SetDowntieringPolicy sets DowntieringPolicy field to given value.

### HasDowntieringPolicy

`func (o *CreateOrUpdateDataTieringTaskRequest) HasDowntieringPolicy() bool`

HasDowntieringPolicy returns a boolean if a field has been set.

### GetUptieringPolicy

`func (o *CreateOrUpdateDataTieringTaskRequest) GetUptieringPolicy() UptieringPolicy`

GetUptieringPolicy returns the UptieringPolicy field if non-nil, zero value otherwise.

### GetUptieringPolicyOk

`func (o *CreateOrUpdateDataTieringTaskRequest) GetUptieringPolicyOk() (*UptieringPolicy, bool)`

GetUptieringPolicyOk returns a tuple with the UptieringPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptieringPolicy

`func (o *CreateOrUpdateDataTieringTaskRequest) SetUptieringPolicy(v UptieringPolicy)`

SetUptieringPolicy sets UptieringPolicy field to given value.

### HasUptieringPolicy

`func (o *CreateOrUpdateDataTieringTaskRequest) HasUptieringPolicy() bool`

HasUptieringPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


