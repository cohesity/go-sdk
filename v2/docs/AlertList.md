# AlertList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerts** | [**[]AlertInfo**](AlertInfo.md) | Specifies the list of alerts. | 

## Methods

### NewAlertList

`func NewAlertList(alerts []AlertInfo, ) *AlertList`

NewAlertList instantiates a new AlertList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertListWithDefaults

`func NewAlertListWithDefaults() *AlertList`

NewAlertListWithDefaults instantiates a new AlertList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlerts

`func (o *AlertList) GetAlerts() []AlertInfo`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *AlertList) GetAlertsOk() (*[]AlertInfo, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *AlertList) SetAlerts(v []AlertInfo)`

SetAlerts sets Alerts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


