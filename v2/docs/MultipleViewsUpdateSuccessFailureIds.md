# MultipleViewsUpdateSuccessFailureIds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedViewIds** | Pointer to **[]int32** | List of View Ids that has resulted in a failed update. | [optional] 
**SucceededViewIds** | Pointer to **[]int32** | List of View Ids that has resulted in a successful update. | [optional] 

## Methods

### NewMultipleViewsUpdateSuccessFailureIds

`func NewMultipleViewsUpdateSuccessFailureIds() *MultipleViewsUpdateSuccessFailureIds`

NewMultipleViewsUpdateSuccessFailureIds instantiates a new MultipleViewsUpdateSuccessFailureIds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleViewsUpdateSuccessFailureIdsWithDefaults

`func NewMultipleViewsUpdateSuccessFailureIdsWithDefaults() *MultipleViewsUpdateSuccessFailureIds`

NewMultipleViewsUpdateSuccessFailureIdsWithDefaults instantiates a new MultipleViewsUpdateSuccessFailureIds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedViewIds

`func (o *MultipleViewsUpdateSuccessFailureIds) GetFailedViewIds() []int32`

GetFailedViewIds returns the FailedViewIds field if non-nil, zero value otherwise.

### GetFailedViewIdsOk

`func (o *MultipleViewsUpdateSuccessFailureIds) GetFailedViewIdsOk() (*[]int32, bool)`

GetFailedViewIdsOk returns a tuple with the FailedViewIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedViewIds

`func (o *MultipleViewsUpdateSuccessFailureIds) SetFailedViewIds(v []int32)`

SetFailedViewIds sets FailedViewIds field to given value.

### HasFailedViewIds

`func (o *MultipleViewsUpdateSuccessFailureIds) HasFailedViewIds() bool`

HasFailedViewIds returns a boolean if a field has been set.

### SetFailedViewIdsNil

`func (o *MultipleViewsUpdateSuccessFailureIds) SetFailedViewIdsNil(b bool)`

 SetFailedViewIdsNil sets the value for FailedViewIds to be an explicit nil

### UnsetFailedViewIds
`func (o *MultipleViewsUpdateSuccessFailureIds) UnsetFailedViewIds()`

UnsetFailedViewIds ensures that no value is present for FailedViewIds, not even an explicit nil
### GetSucceededViewIds

`func (o *MultipleViewsUpdateSuccessFailureIds) GetSucceededViewIds() []int32`

GetSucceededViewIds returns the SucceededViewIds field if non-nil, zero value otherwise.

### GetSucceededViewIdsOk

`func (o *MultipleViewsUpdateSuccessFailureIds) GetSucceededViewIdsOk() (*[]int32, bool)`

GetSucceededViewIdsOk returns a tuple with the SucceededViewIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceededViewIds

`func (o *MultipleViewsUpdateSuccessFailureIds) SetSucceededViewIds(v []int32)`

SetSucceededViewIds sets SucceededViewIds field to given value.

### HasSucceededViewIds

`func (o *MultipleViewsUpdateSuccessFailureIds) HasSucceededViewIds() bool`

HasSucceededViewIds returns a boolean if a field has been set.

### SetSucceededViewIdsNil

`func (o *MultipleViewsUpdateSuccessFailureIds) SetSucceededViewIdsNil(b bool)`

 SetSucceededViewIdsNil sets the value for SucceededViewIds to be an explicit nil

### UnsetSucceededViewIds
`func (o *MultipleViewsUpdateSuccessFailureIds) UnsetSucceededViewIds()`

UnsetSucceededViewIds ensures that no value is present for SucceededViewIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


