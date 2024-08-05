# UpdateProtectionGroupRunResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedRuns** | Pointer to [**[]FailedRunDetails**](FailedRunDetails.md) | Specifies a list of ids of Protection Group Runs that failed to update along with error details. | [optional] 
**SuccessfulRunIds** | Pointer to **[]string** | Specifies a list of ids of Protection Group Runs that are successfully updated. | [optional] 

## Methods

### NewUpdateProtectionGroupRunResponseBody

`func NewUpdateProtectionGroupRunResponseBody() *UpdateProtectionGroupRunResponseBody`

NewUpdateProtectionGroupRunResponseBody instantiates a new UpdateProtectionGroupRunResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProtectionGroupRunResponseBodyWithDefaults

`func NewUpdateProtectionGroupRunResponseBodyWithDefaults() *UpdateProtectionGroupRunResponseBody`

NewUpdateProtectionGroupRunResponseBodyWithDefaults instantiates a new UpdateProtectionGroupRunResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedRuns

`func (o *UpdateProtectionGroupRunResponseBody) GetFailedRuns() []FailedRunDetails`

GetFailedRuns returns the FailedRuns field if non-nil, zero value otherwise.

### GetFailedRunsOk

`func (o *UpdateProtectionGroupRunResponseBody) GetFailedRunsOk() (*[]FailedRunDetails, bool)`

GetFailedRunsOk returns a tuple with the FailedRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRuns

`func (o *UpdateProtectionGroupRunResponseBody) SetFailedRuns(v []FailedRunDetails)`

SetFailedRuns sets FailedRuns field to given value.

### HasFailedRuns

`func (o *UpdateProtectionGroupRunResponseBody) HasFailedRuns() bool`

HasFailedRuns returns a boolean if a field has been set.

### SetFailedRunsNil

`func (o *UpdateProtectionGroupRunResponseBody) SetFailedRunsNil(b bool)`

 SetFailedRunsNil sets the value for FailedRuns to be an explicit nil

### UnsetFailedRuns
`func (o *UpdateProtectionGroupRunResponseBody) UnsetFailedRuns()`

UnsetFailedRuns ensures that no value is present for FailedRuns, not even an explicit nil
### GetSuccessfulRunIds

`func (o *UpdateProtectionGroupRunResponseBody) GetSuccessfulRunIds() []string`

GetSuccessfulRunIds returns the SuccessfulRunIds field if non-nil, zero value otherwise.

### GetSuccessfulRunIdsOk

`func (o *UpdateProtectionGroupRunResponseBody) GetSuccessfulRunIdsOk() (*[]string, bool)`

GetSuccessfulRunIdsOk returns a tuple with the SuccessfulRunIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulRunIds

`func (o *UpdateProtectionGroupRunResponseBody) SetSuccessfulRunIds(v []string)`

SetSuccessfulRunIds sets SuccessfulRunIds field to given value.

### HasSuccessfulRunIds

`func (o *UpdateProtectionGroupRunResponseBody) HasSuccessfulRunIds() bool`

HasSuccessfulRunIds returns a boolean if a field has been set.

### SetSuccessfulRunIdsNil

`func (o *UpdateProtectionGroupRunResponseBody) SetSuccessfulRunIdsNil(b bool)`

 SetSuccessfulRunIdsNil sets the value for SuccessfulRunIds to be an explicit nil

### UnsetSuccessfulRunIds
`func (o *UpdateProtectionGroupRunResponseBody) UnsetSuccessfulRunIds()`

UnsetSuccessfulRunIds ensures that no value is present for SuccessfulRunIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


