# ProtectionSourceSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsProtectionGroupPaused** | Pointer to **NullableBool** | Specifies the status of the protection group | [optional] 
**LastProtectionRun** | Pointer to [**ProtectionGroupRun**](ProtectionGroupRun.md) |  | [optional] 
**NextProtectionRunTimeUsecs** | Pointer to **NullableInt64** | Specifies the time at which the next Protection Run is scheduled for the given Protection Source in Unix epoch Time | [optional] 
**Object** | Pointer to [**Object**](Object.md) |  | [optional] 

## Methods

### NewProtectionSourceSummary

`func NewProtectionSourceSummary() *ProtectionSourceSummary`

NewProtectionSourceSummary instantiates a new ProtectionSourceSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionSourceSummaryWithDefaults

`func NewProtectionSourceSummaryWithDefaults() *ProtectionSourceSummary`

NewProtectionSourceSummaryWithDefaults instantiates a new ProtectionSourceSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsProtectionGroupPaused

`func (o *ProtectionSourceSummary) GetIsProtectionGroupPaused() bool`

GetIsProtectionGroupPaused returns the IsProtectionGroupPaused field if non-nil, zero value otherwise.

### GetIsProtectionGroupPausedOk

`func (o *ProtectionSourceSummary) GetIsProtectionGroupPausedOk() (*bool, bool)`

GetIsProtectionGroupPausedOk returns a tuple with the IsProtectionGroupPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProtectionGroupPaused

`func (o *ProtectionSourceSummary) SetIsProtectionGroupPaused(v bool)`

SetIsProtectionGroupPaused sets IsProtectionGroupPaused field to given value.

### HasIsProtectionGroupPaused

`func (o *ProtectionSourceSummary) HasIsProtectionGroupPaused() bool`

HasIsProtectionGroupPaused returns a boolean if a field has been set.

### SetIsProtectionGroupPausedNil

`func (o *ProtectionSourceSummary) SetIsProtectionGroupPausedNil(b bool)`

 SetIsProtectionGroupPausedNil sets the value for IsProtectionGroupPaused to be an explicit nil

### UnsetIsProtectionGroupPaused
`func (o *ProtectionSourceSummary) UnsetIsProtectionGroupPaused()`

UnsetIsProtectionGroupPaused ensures that no value is present for IsProtectionGroupPaused, not even an explicit nil
### GetLastProtectionRun

`func (o *ProtectionSourceSummary) GetLastProtectionRun() ProtectionGroupRun`

GetLastProtectionRun returns the LastProtectionRun field if non-nil, zero value otherwise.

### GetLastProtectionRunOk

`func (o *ProtectionSourceSummary) GetLastProtectionRunOk() (*ProtectionGroupRun, bool)`

GetLastProtectionRunOk returns a tuple with the LastProtectionRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProtectionRun

`func (o *ProtectionSourceSummary) SetLastProtectionRun(v ProtectionGroupRun)`

SetLastProtectionRun sets LastProtectionRun field to given value.

### HasLastProtectionRun

`func (o *ProtectionSourceSummary) HasLastProtectionRun() bool`

HasLastProtectionRun returns a boolean if a field has been set.

### GetNextProtectionRunTimeUsecs

`func (o *ProtectionSourceSummary) GetNextProtectionRunTimeUsecs() int64`

GetNextProtectionRunTimeUsecs returns the NextProtectionRunTimeUsecs field if non-nil, zero value otherwise.

### GetNextProtectionRunTimeUsecsOk

`func (o *ProtectionSourceSummary) GetNextProtectionRunTimeUsecsOk() (*int64, bool)`

GetNextProtectionRunTimeUsecsOk returns a tuple with the NextProtectionRunTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextProtectionRunTimeUsecs

`func (o *ProtectionSourceSummary) SetNextProtectionRunTimeUsecs(v int64)`

SetNextProtectionRunTimeUsecs sets NextProtectionRunTimeUsecs field to given value.

### HasNextProtectionRunTimeUsecs

`func (o *ProtectionSourceSummary) HasNextProtectionRunTimeUsecs() bool`

HasNextProtectionRunTimeUsecs returns a boolean if a field has been set.

### SetNextProtectionRunTimeUsecsNil

`func (o *ProtectionSourceSummary) SetNextProtectionRunTimeUsecsNil(b bool)`

 SetNextProtectionRunTimeUsecsNil sets the value for NextProtectionRunTimeUsecs to be an explicit nil

### UnsetNextProtectionRunTimeUsecs
`func (o *ProtectionSourceSummary) UnsetNextProtectionRunTimeUsecs()`

UnsetNextProtectionRunTimeUsecs ensures that no value is present for NextProtectionRunTimeUsecs, not even an explicit nil
### GetObject

`func (o *ProtectionSourceSummary) GetObject() Object`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ProtectionSourceSummary) GetObjectOk() (*Object, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ProtectionSourceSummary) SetObject(v Object)`

SetObject sets Object field to given value.

### HasObject

`func (o *ProtectionSourceSummary) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


