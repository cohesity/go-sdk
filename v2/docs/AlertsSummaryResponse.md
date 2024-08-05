# AlertsSummaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertsSummary** | Pointer to [**[]AlertGroupSummary**](AlertGroupSummary.md) | Specifies a list of alerts summary grouped by category. | [optional] 

## Methods

### NewAlertsSummaryResponse

`func NewAlertsSummaryResponse() *AlertsSummaryResponse`

NewAlertsSummaryResponse instantiates a new AlertsSummaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertsSummaryResponseWithDefaults

`func NewAlertsSummaryResponseWithDefaults() *AlertsSummaryResponse`

NewAlertsSummaryResponseWithDefaults instantiates a new AlertsSummaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertsSummary

`func (o *AlertsSummaryResponse) GetAlertsSummary() []AlertGroupSummary`

GetAlertsSummary returns the AlertsSummary field if non-nil, zero value otherwise.

### GetAlertsSummaryOk

`func (o *AlertsSummaryResponse) GetAlertsSummaryOk() (*[]AlertGroupSummary, bool)`

GetAlertsSummaryOk returns a tuple with the AlertsSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertsSummary

`func (o *AlertsSummaryResponse) SetAlertsSummary(v []AlertGroupSummary)`

SetAlertsSummary sets AlertsSummary field to given value.

### HasAlertsSummary

`func (o *AlertsSummaryResponse) HasAlertsSummary() bool`

HasAlertsSummary returns a boolean if a field has been set.

### SetAlertsSummaryNil

`func (o *AlertsSummaryResponse) SetAlertsSummaryNil(b bool)`

 SetAlertsSummaryNil sets the value for AlertsSummary to be an explicit nil

### UnsetAlertsSummary
`func (o *AlertsSummaryResponse) UnsetAlertsSummary()`

UnsetAlertsSummary ensures that no value is present for AlertsSummary, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


