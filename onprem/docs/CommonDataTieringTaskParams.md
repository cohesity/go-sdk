# CommonDataTieringTaskParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Specifies the name of the data tiering task. | 
**Description** | Pointer to **NullableString** | Specifies a description of the data tiering task. | [optional] 
**AlertPolicy** | Pointer to [**ProtectionGroupAlertingPolicy**](ProtectionGroupAlertingPolicy.md) |  | [optional] 
**Source** | Pointer to [**DataTieringSource**](DataTieringSource.md) |  | [optional] 
**Target** | Pointer to [**NullableDataTieringTarget**](DataTieringTarget.md) |  | [optional] 
**Schedule** | Pointer to [**DataTieringSchedule**](DataTieringSchedule.md) |  | [optional] 
**Type** | **NullableString** | Type of data tiering task. &#39;Downtier&#39; indicates downtiering task. &#39;Uptier&#39; indicates uptiering task. | 

## Methods

### NewCommonDataTieringTaskParams

`func NewCommonDataTieringTaskParams(name NullableString, type_ NullableString, ) *CommonDataTieringTaskParams`

NewCommonDataTieringTaskParams instantiates a new CommonDataTieringTaskParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonDataTieringTaskParamsWithDefaults

`func NewCommonDataTieringTaskParamsWithDefaults() *CommonDataTieringTaskParams`

NewCommonDataTieringTaskParamsWithDefaults instantiates a new CommonDataTieringTaskParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CommonDataTieringTaskParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommonDataTieringTaskParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommonDataTieringTaskParams) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CommonDataTieringTaskParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CommonDataTieringTaskParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *CommonDataTieringTaskParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CommonDataTieringTaskParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CommonDataTieringTaskParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CommonDataTieringTaskParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CommonDataTieringTaskParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CommonDataTieringTaskParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAlertPolicy

`func (o *CommonDataTieringTaskParams) GetAlertPolicy() ProtectionGroupAlertingPolicy`

GetAlertPolicy returns the AlertPolicy field if non-nil, zero value otherwise.

### GetAlertPolicyOk

`func (o *CommonDataTieringTaskParams) GetAlertPolicyOk() (*ProtectionGroupAlertingPolicy, bool)`

GetAlertPolicyOk returns a tuple with the AlertPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertPolicy

`func (o *CommonDataTieringTaskParams) SetAlertPolicy(v ProtectionGroupAlertingPolicy)`

SetAlertPolicy sets AlertPolicy field to given value.

### HasAlertPolicy

`func (o *CommonDataTieringTaskParams) HasAlertPolicy() bool`

HasAlertPolicy returns a boolean if a field has been set.

### GetSource

`func (o *CommonDataTieringTaskParams) GetSource() DataTieringSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CommonDataTieringTaskParams) GetSourceOk() (*DataTieringSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CommonDataTieringTaskParams) SetSource(v DataTieringSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *CommonDataTieringTaskParams) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTarget

`func (o *CommonDataTieringTaskParams) GetTarget() DataTieringTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CommonDataTieringTaskParams) GetTargetOk() (*DataTieringTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CommonDataTieringTaskParams) SetTarget(v DataTieringTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CommonDataTieringTaskParams) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *CommonDataTieringTaskParams) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *CommonDataTieringTaskParams) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil
### GetSchedule

`func (o *CommonDataTieringTaskParams) GetSchedule() DataTieringSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CommonDataTieringTaskParams) GetScheduleOk() (*DataTieringSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CommonDataTieringTaskParams) SetSchedule(v DataTieringSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CommonDataTieringTaskParams) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetType

`func (o *CommonDataTieringTaskParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CommonDataTieringTaskParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CommonDataTieringTaskParams) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CommonDataTieringTaskParams) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CommonDataTieringTaskParams) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


