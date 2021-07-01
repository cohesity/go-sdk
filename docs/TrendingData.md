# TrendingData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cancelled** | Pointer to **NullableInt64** | Specifies number of cancelled runs. | [optional] 
**Failed** | Pointer to **NullableInt64** | Specifies number of failed runs. | [optional] 
**Running** | Pointer to **NullableInt64** | Specifies number of in-progress runs. | [optional] 
**Successful** | Pointer to **NullableInt64** | Specifies number of successful runs. | [optional] 
**Total** | Pointer to **NullableInt64** | Specifies total number of runs. | [optional] 
**TrendName** | Pointer to **NullableString** | Specifies trend name. This is start of the day/week/month. | [optional] 
**TrendStartTimeUsecs** | Pointer to **NullableInt64** | Specifies start of the day/week/month in micro seconds | [optional] 

## Methods

### NewTrendingData

`func NewTrendingData() *TrendingData`

NewTrendingData instantiates a new TrendingData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrendingDataWithDefaults

`func NewTrendingDataWithDefaults() *TrendingData`

NewTrendingDataWithDefaults instantiates a new TrendingData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelled

`func (o *TrendingData) GetCancelled() int64`

GetCancelled returns the Cancelled field if non-nil, zero value otherwise.

### GetCancelledOk

`func (o *TrendingData) GetCancelledOk() (*int64, bool)`

GetCancelledOk returns a tuple with the Cancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelled

`func (o *TrendingData) SetCancelled(v int64)`

SetCancelled sets Cancelled field to given value.

### HasCancelled

`func (o *TrendingData) HasCancelled() bool`

HasCancelled returns a boolean if a field has been set.

### SetCancelledNil

`func (o *TrendingData) SetCancelledNil(b bool)`

 SetCancelledNil sets the value for Cancelled to be an explicit nil

### UnsetCancelled
`func (o *TrendingData) UnsetCancelled()`

UnsetCancelled ensures that no value is present for Cancelled, not even an explicit nil
### GetFailed

`func (o *TrendingData) GetFailed() int64`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *TrendingData) GetFailedOk() (*int64, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *TrendingData) SetFailed(v int64)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *TrendingData) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### SetFailedNil

`func (o *TrendingData) SetFailedNil(b bool)`

 SetFailedNil sets the value for Failed to be an explicit nil

### UnsetFailed
`func (o *TrendingData) UnsetFailed()`

UnsetFailed ensures that no value is present for Failed, not even an explicit nil
### GetRunning

`func (o *TrendingData) GetRunning() int64`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *TrendingData) GetRunningOk() (*int64, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *TrendingData) SetRunning(v int64)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *TrendingData) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### SetRunningNil

`func (o *TrendingData) SetRunningNil(b bool)`

 SetRunningNil sets the value for Running to be an explicit nil

### UnsetRunning
`func (o *TrendingData) UnsetRunning()`

UnsetRunning ensures that no value is present for Running, not even an explicit nil
### GetSuccessful

`func (o *TrendingData) GetSuccessful() int64`

GetSuccessful returns the Successful field if non-nil, zero value otherwise.

### GetSuccessfulOk

`func (o *TrendingData) GetSuccessfulOk() (*int64, bool)`

GetSuccessfulOk returns a tuple with the Successful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessful

`func (o *TrendingData) SetSuccessful(v int64)`

SetSuccessful sets Successful field to given value.

### HasSuccessful

`func (o *TrendingData) HasSuccessful() bool`

HasSuccessful returns a boolean if a field has been set.

### SetSuccessfulNil

`func (o *TrendingData) SetSuccessfulNil(b bool)`

 SetSuccessfulNil sets the value for Successful to be an explicit nil

### UnsetSuccessful
`func (o *TrendingData) UnsetSuccessful()`

UnsetSuccessful ensures that no value is present for Successful, not even an explicit nil
### GetTotal

`func (o *TrendingData) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TrendingData) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TrendingData) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *TrendingData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *TrendingData) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *TrendingData) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetTrendName

`func (o *TrendingData) GetTrendName() string`

GetTrendName returns the TrendName field if non-nil, zero value otherwise.

### GetTrendNameOk

`func (o *TrendingData) GetTrendNameOk() (*string, bool)`

GetTrendNameOk returns a tuple with the TrendName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrendName

`func (o *TrendingData) SetTrendName(v string)`

SetTrendName sets TrendName field to given value.

### HasTrendName

`func (o *TrendingData) HasTrendName() bool`

HasTrendName returns a boolean if a field has been set.

### SetTrendNameNil

`func (o *TrendingData) SetTrendNameNil(b bool)`

 SetTrendNameNil sets the value for TrendName to be an explicit nil

### UnsetTrendName
`func (o *TrendingData) UnsetTrendName()`

UnsetTrendName ensures that no value is present for TrendName, not even an explicit nil
### GetTrendStartTimeUsecs

`func (o *TrendingData) GetTrendStartTimeUsecs() int64`

GetTrendStartTimeUsecs returns the TrendStartTimeUsecs field if non-nil, zero value otherwise.

### GetTrendStartTimeUsecsOk

`func (o *TrendingData) GetTrendStartTimeUsecsOk() (*int64, bool)`

GetTrendStartTimeUsecsOk returns a tuple with the TrendStartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrendStartTimeUsecs

`func (o *TrendingData) SetTrendStartTimeUsecs(v int64)`

SetTrendStartTimeUsecs sets TrendStartTimeUsecs field to given value.

### HasTrendStartTimeUsecs

`func (o *TrendingData) HasTrendStartTimeUsecs() bool`

HasTrendStartTimeUsecs returns a boolean if a field has been set.

### SetTrendStartTimeUsecsNil

`func (o *TrendingData) SetTrendStartTimeUsecsNil(b bool)`

 SetTrendStartTimeUsecsNil sets the value for TrendStartTimeUsecs to be an explicit nil

### UnsetTrendStartTimeUsecs
`func (o *TrendingData) UnsetTrendStartTimeUsecs()`

UnsetTrendStartTimeUsecs ensures that no value is present for TrendStartTimeUsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


