# AlertGroupSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Category of alerts by which summary is grouped. | [optional] 
**Type** | Pointer to **string** | Type/bucket that this alert category belongs to. | [optional] 
**WarningCount** | Pointer to **NullableInt64** | Specifies count of warning alerts. | [optional] 
**CriticalCount** | Pointer to **NullableInt64** | Specifies count of critical alerts. | [optional] 
**InfoCount** | Pointer to **NullableInt64** | Specifies count of info alerts. | [optional] 
**TotalCount** | Pointer to **NullableInt64** | Specifies count of total alerts. | [optional] 

## Methods

### NewAlertGroupSummary

`func NewAlertGroupSummary() *AlertGroupSummary`

NewAlertGroupSummary instantiates a new AlertGroupSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertGroupSummaryWithDefaults

`func NewAlertGroupSummaryWithDefaults() *AlertGroupSummary`

NewAlertGroupSummaryWithDefaults instantiates a new AlertGroupSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *AlertGroupSummary) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AlertGroupSummary) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AlertGroupSummary) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AlertGroupSummary) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetType

`func (o *AlertGroupSummary) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlertGroupSummary) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlertGroupSummary) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AlertGroupSummary) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWarningCount

`func (o *AlertGroupSummary) GetWarningCount() int64`

GetWarningCount returns the WarningCount field if non-nil, zero value otherwise.

### GetWarningCountOk

`func (o *AlertGroupSummary) GetWarningCountOk() (*int64, bool)`

GetWarningCountOk returns a tuple with the WarningCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningCount

`func (o *AlertGroupSummary) SetWarningCount(v int64)`

SetWarningCount sets WarningCount field to given value.

### HasWarningCount

`func (o *AlertGroupSummary) HasWarningCount() bool`

HasWarningCount returns a boolean if a field has been set.

### SetWarningCountNil

`func (o *AlertGroupSummary) SetWarningCountNil(b bool)`

 SetWarningCountNil sets the value for WarningCount to be an explicit nil

### UnsetWarningCount
`func (o *AlertGroupSummary) UnsetWarningCount()`

UnsetWarningCount ensures that no value is present for WarningCount, not even an explicit nil
### GetCriticalCount

`func (o *AlertGroupSummary) GetCriticalCount() int64`

GetCriticalCount returns the CriticalCount field if non-nil, zero value otherwise.

### GetCriticalCountOk

`func (o *AlertGroupSummary) GetCriticalCountOk() (*int64, bool)`

GetCriticalCountOk returns a tuple with the CriticalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalCount

`func (o *AlertGroupSummary) SetCriticalCount(v int64)`

SetCriticalCount sets CriticalCount field to given value.

### HasCriticalCount

`func (o *AlertGroupSummary) HasCriticalCount() bool`

HasCriticalCount returns a boolean if a field has been set.

### SetCriticalCountNil

`func (o *AlertGroupSummary) SetCriticalCountNil(b bool)`

 SetCriticalCountNil sets the value for CriticalCount to be an explicit nil

### UnsetCriticalCount
`func (o *AlertGroupSummary) UnsetCriticalCount()`

UnsetCriticalCount ensures that no value is present for CriticalCount, not even an explicit nil
### GetInfoCount

`func (o *AlertGroupSummary) GetInfoCount() int64`

GetInfoCount returns the InfoCount field if non-nil, zero value otherwise.

### GetInfoCountOk

`func (o *AlertGroupSummary) GetInfoCountOk() (*int64, bool)`

GetInfoCountOk returns a tuple with the InfoCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoCount

`func (o *AlertGroupSummary) SetInfoCount(v int64)`

SetInfoCount sets InfoCount field to given value.

### HasInfoCount

`func (o *AlertGroupSummary) HasInfoCount() bool`

HasInfoCount returns a boolean if a field has been set.

### SetInfoCountNil

`func (o *AlertGroupSummary) SetInfoCountNil(b bool)`

 SetInfoCountNil sets the value for InfoCount to be an explicit nil

### UnsetInfoCount
`func (o *AlertGroupSummary) UnsetInfoCount()`

UnsetInfoCount ensures that no value is present for InfoCount, not even an explicit nil
### GetTotalCount

`func (o *AlertGroupSummary) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AlertGroupSummary) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AlertGroupSummary) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *AlertGroupSummary) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### SetTotalCountNil

`func (o *AlertGroupSummary) SetTotalCountNil(b bool)`

 SetTotalCountNil sets the value for TotalCount to be an explicit nil

### UnsetTotalCount
`func (o *AlertGroupSummary) UnsetTotalCount()`

UnsetTotalCount ensures that no value is present for TotalCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


