# UpdateDataTieringState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedIds** | Pointer to **[]string** | Specifies a list of data tiering groups ids for which updation of state failed. | [optional] 
**SuccessfulIds** | Pointer to **[]string** | Specifies a list of data tiering groups ids for which updation of state was successful. | [optional] 

## Methods

### NewUpdateDataTieringState

`func NewUpdateDataTieringState() *UpdateDataTieringState`

NewUpdateDataTieringState instantiates a new UpdateDataTieringState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDataTieringStateWithDefaults

`func NewUpdateDataTieringStateWithDefaults() *UpdateDataTieringState`

NewUpdateDataTieringStateWithDefaults instantiates a new UpdateDataTieringState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedIds

`func (o *UpdateDataTieringState) GetFailedIds() []string`

GetFailedIds returns the FailedIds field if non-nil, zero value otherwise.

### GetFailedIdsOk

`func (o *UpdateDataTieringState) GetFailedIdsOk() (*[]string, bool)`

GetFailedIdsOk returns a tuple with the FailedIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedIds

`func (o *UpdateDataTieringState) SetFailedIds(v []string)`

SetFailedIds sets FailedIds field to given value.

### HasFailedIds

`func (o *UpdateDataTieringState) HasFailedIds() bool`

HasFailedIds returns a boolean if a field has been set.

### SetFailedIdsNil

`func (o *UpdateDataTieringState) SetFailedIdsNil(b bool)`

 SetFailedIdsNil sets the value for FailedIds to be an explicit nil

### UnsetFailedIds
`func (o *UpdateDataTieringState) UnsetFailedIds()`

UnsetFailedIds ensures that no value is present for FailedIds, not even an explicit nil
### GetSuccessfulIds

`func (o *UpdateDataTieringState) GetSuccessfulIds() []string`

GetSuccessfulIds returns the SuccessfulIds field if non-nil, zero value otherwise.

### GetSuccessfulIdsOk

`func (o *UpdateDataTieringState) GetSuccessfulIdsOk() (*[]string, bool)`

GetSuccessfulIdsOk returns a tuple with the SuccessfulIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulIds

`func (o *UpdateDataTieringState) SetSuccessfulIds(v []string)`

SetSuccessfulIds sets SuccessfulIds field to given value.

### HasSuccessfulIds

`func (o *UpdateDataTieringState) HasSuccessfulIds() bool`

HasSuccessfulIds returns a boolean if a field has been set.

### SetSuccessfulIdsNil

`func (o *UpdateDataTieringState) SetSuccessfulIdsNil(b bool)`

 SetSuccessfulIdsNil sets the value for SuccessfulIds to be an explicit nil

### UnsetSuccessfulIds
`func (o *UpdateDataTieringState) UnsetSuccessfulIds()`

UnsetSuccessfulIds ensures that no value is present for SuccessfulIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


