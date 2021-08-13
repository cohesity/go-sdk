# ProgressSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **NullableInt64** | Specifies the successful count. | [optional] 
**Failed** | Pointer to **NullableInt64** | Specifies the failed count. | [optional] 
**Total** | Pointer to **NullableInt64** | Specifies the total count. | [optional] 

## Methods

### NewProgressSummary

`func NewProgressSummary() *ProgressSummary`

NewProgressSummary instantiates a new ProgressSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgressSummaryWithDefaults

`func NewProgressSummaryWithDefaults() *ProgressSummary`

NewProgressSummaryWithDefaults instantiates a new ProgressSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ProgressSummary) GetSuccess() int64`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ProgressSummary) GetSuccessOk() (*int64, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ProgressSummary) SetSuccess(v int64)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ProgressSummary) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### SetSuccessNil

`func (o *ProgressSummary) SetSuccessNil(b bool)`

 SetSuccessNil sets the value for Success to be an explicit nil

### UnsetSuccess
`func (o *ProgressSummary) UnsetSuccess()`

UnsetSuccess ensures that no value is present for Success, not even an explicit nil
### GetFailed

`func (o *ProgressSummary) GetFailed() int64`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ProgressSummary) GetFailedOk() (*int64, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ProgressSummary) SetFailed(v int64)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *ProgressSummary) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### SetFailedNil

`func (o *ProgressSummary) SetFailedNil(b bool)`

 SetFailedNil sets the value for Failed to be an explicit nil

### UnsetFailed
`func (o *ProgressSummary) UnsetFailed()`

UnsetFailed ensures that no value is present for Failed, not even an explicit nil
### GetTotal

`func (o *ProgressSummary) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ProgressSummary) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ProgressSummary) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ProgressSummary) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *ProgressSummary) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *ProgressSummary) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


