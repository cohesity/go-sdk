# ProtectionGroupRuns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsResponseTruncated** | Pointer to **NullableBool** | Indicates whether the result is truncated due to hitting maximum size limit governed by magneto_http_rpc_response_size_limit_bytes | [optional] 
**Runs** | Pointer to [**[]ProtectionGroupRun**](ProtectionGroupRun.md) | Specifies the list of Protection Group runs. | [optional] 
**TotalRuns** | Pointer to **NullableInt32** | Specifies the count of total runs exist for the given set of filters. The number of runs in single API call are limited and this count can be used to estimate query filter values to get next set of remaining runs. Please note that this field will only be populated if startTimeUsecs or endTimeUsecs or both are specified in query parameters. | [optional] 

## Methods

### NewProtectionGroupRuns

`func NewProtectionGroupRuns() *ProtectionGroupRuns`

NewProtectionGroupRuns instantiates a new ProtectionGroupRuns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionGroupRunsWithDefaults

`func NewProtectionGroupRunsWithDefaults() *ProtectionGroupRuns`

NewProtectionGroupRunsWithDefaults instantiates a new ProtectionGroupRuns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsResponseTruncated

`func (o *ProtectionGroupRuns) GetIsResponseTruncated() bool`

GetIsResponseTruncated returns the IsResponseTruncated field if non-nil, zero value otherwise.

### GetIsResponseTruncatedOk

`func (o *ProtectionGroupRuns) GetIsResponseTruncatedOk() (*bool, bool)`

GetIsResponseTruncatedOk returns a tuple with the IsResponseTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsResponseTruncated

`func (o *ProtectionGroupRuns) SetIsResponseTruncated(v bool)`

SetIsResponseTruncated sets IsResponseTruncated field to given value.

### HasIsResponseTruncated

`func (o *ProtectionGroupRuns) HasIsResponseTruncated() bool`

HasIsResponseTruncated returns a boolean if a field has been set.

### SetIsResponseTruncatedNil

`func (o *ProtectionGroupRuns) SetIsResponseTruncatedNil(b bool)`

 SetIsResponseTruncatedNil sets the value for IsResponseTruncated to be an explicit nil

### UnsetIsResponseTruncated
`func (o *ProtectionGroupRuns) UnsetIsResponseTruncated()`

UnsetIsResponseTruncated ensures that no value is present for IsResponseTruncated, not even an explicit nil
### GetRuns

`func (o *ProtectionGroupRuns) GetRuns() []ProtectionGroupRun`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *ProtectionGroupRuns) GetRunsOk() (*[]ProtectionGroupRun, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *ProtectionGroupRuns) SetRuns(v []ProtectionGroupRun)`

SetRuns sets Runs field to given value.

### HasRuns

`func (o *ProtectionGroupRuns) HasRuns() bool`

HasRuns returns a boolean if a field has been set.

### SetRunsNil

`func (o *ProtectionGroupRuns) SetRunsNil(b bool)`

 SetRunsNil sets the value for Runs to be an explicit nil

### UnsetRuns
`func (o *ProtectionGroupRuns) UnsetRuns()`

UnsetRuns ensures that no value is present for Runs, not even an explicit nil
### GetTotalRuns

`func (o *ProtectionGroupRuns) GetTotalRuns() int32`

GetTotalRuns returns the TotalRuns field if non-nil, zero value otherwise.

### GetTotalRunsOk

`func (o *ProtectionGroupRuns) GetTotalRunsOk() (*int32, bool)`

GetTotalRunsOk returns a tuple with the TotalRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRuns

`func (o *ProtectionGroupRuns) SetTotalRuns(v int32)`

SetTotalRuns sets TotalRuns field to given value.

### HasTotalRuns

`func (o *ProtectionGroupRuns) HasTotalRuns() bool`

HasTotalRuns returns a boolean if a field has been set.

### SetTotalRunsNil

`func (o *ProtectionGroupRuns) SetTotalRunsNil(b bool)`

 SetTotalRunsNil sets the value for TotalRuns to be an explicit nil

### UnsetTotalRuns
`func (o *ProtectionGroupRuns) UnsetTotalRuns()`

UnsetTotalRuns ensures that no value is present for TotalRuns, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


