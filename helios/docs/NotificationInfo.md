# NotificationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnomalyStrengthThreshold** | Pointer to **NullableInt64** | Anomaly strength level beyond which notification has to be sent. | [optional] 

## Methods

### NewNotificationInfo

`func NewNotificationInfo() *NotificationInfo`

NewNotificationInfo instantiates a new NotificationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationInfoWithDefaults

`func NewNotificationInfoWithDefaults() *NotificationInfo`

NewNotificationInfoWithDefaults instantiates a new NotificationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnomalyStrengthThreshold

`func (o *NotificationInfo) GetAnomalyStrengthThreshold() int64`

GetAnomalyStrengthThreshold returns the AnomalyStrengthThreshold field if non-nil, zero value otherwise.

### GetAnomalyStrengthThresholdOk

`func (o *NotificationInfo) GetAnomalyStrengthThresholdOk() (*int64, bool)`

GetAnomalyStrengthThresholdOk returns a tuple with the AnomalyStrengthThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomalyStrengthThreshold

`func (o *NotificationInfo) SetAnomalyStrengthThreshold(v int64)`

SetAnomalyStrengthThreshold sets AnomalyStrengthThreshold field to given value.

### HasAnomalyStrengthThreshold

`func (o *NotificationInfo) HasAnomalyStrengthThreshold() bool`

HasAnomalyStrengthThreshold returns a boolean if a field has been set.

### SetAnomalyStrengthThresholdNil

`func (o *NotificationInfo) SetAnomalyStrengthThresholdNil(b bool)`

 SetAnomalyStrengthThresholdNil sets the value for AnomalyStrengthThreshold to be an explicit nil

### UnsetAnomalyStrengthThreshold
`func (o *NotificationInfo) UnsetAnomalyStrengthThreshold()`

UnsetAnomalyStrengthThreshold ensures that no value is present for AnomalyStrengthThreshold, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


