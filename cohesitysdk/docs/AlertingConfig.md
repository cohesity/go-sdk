# AlertingConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddresses** | Pointer to **[]string** | Exists to maintain backwards compatibility with versions before eff8198. | [optional] 
**EmailDeliveryTargets** | Pointer to [**[]EmailDeliveryTarget**](EmailDeliveryTarget.md) | Specifies additional email addresses where alert notifications (configured in the AlertingPolicy) must be sent. | [optional] 
**RaiseObjectLevelFailureAlert** | Pointer to **NullableBool** | Specifies the boolean to raise per object alert for failures. | [optional] 

## Methods

### NewAlertingConfig

`func NewAlertingConfig() *AlertingConfig`

NewAlertingConfig instantiates a new AlertingConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertingConfigWithDefaults

`func NewAlertingConfigWithDefaults() *AlertingConfig`

NewAlertingConfigWithDefaults instantiates a new AlertingConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddresses

`func (o *AlertingConfig) GetEmailAddresses() []string`

GetEmailAddresses returns the EmailAddresses field if non-nil, zero value otherwise.

### GetEmailAddressesOk

`func (o *AlertingConfig) GetEmailAddressesOk() (*[]string, bool)`

GetEmailAddressesOk returns a tuple with the EmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddresses

`func (o *AlertingConfig) SetEmailAddresses(v []string)`

SetEmailAddresses sets EmailAddresses field to given value.

### HasEmailAddresses

`func (o *AlertingConfig) HasEmailAddresses() bool`

HasEmailAddresses returns a boolean if a field has been set.

### SetEmailAddressesNil

`func (o *AlertingConfig) SetEmailAddressesNil(b bool)`

 SetEmailAddressesNil sets the value for EmailAddresses to be an explicit nil

### UnsetEmailAddresses
`func (o *AlertingConfig) UnsetEmailAddresses()`

UnsetEmailAddresses ensures that no value is present for EmailAddresses, not even an explicit nil
### GetEmailDeliveryTargets

`func (o *AlertingConfig) GetEmailDeliveryTargets() []EmailDeliveryTarget`

GetEmailDeliveryTargets returns the EmailDeliveryTargets field if non-nil, zero value otherwise.

### GetEmailDeliveryTargetsOk

`func (o *AlertingConfig) GetEmailDeliveryTargetsOk() (*[]EmailDeliveryTarget, bool)`

GetEmailDeliveryTargetsOk returns a tuple with the EmailDeliveryTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailDeliveryTargets

`func (o *AlertingConfig) SetEmailDeliveryTargets(v []EmailDeliveryTarget)`

SetEmailDeliveryTargets sets EmailDeliveryTargets field to given value.

### HasEmailDeliveryTargets

`func (o *AlertingConfig) HasEmailDeliveryTargets() bool`

HasEmailDeliveryTargets returns a boolean if a field has been set.

### SetEmailDeliveryTargetsNil

`func (o *AlertingConfig) SetEmailDeliveryTargetsNil(b bool)`

 SetEmailDeliveryTargetsNil sets the value for EmailDeliveryTargets to be an explicit nil

### UnsetEmailDeliveryTargets
`func (o *AlertingConfig) UnsetEmailDeliveryTargets()`

UnsetEmailDeliveryTargets ensures that no value is present for EmailDeliveryTargets, not even an explicit nil
### GetRaiseObjectLevelFailureAlert

`func (o *AlertingConfig) GetRaiseObjectLevelFailureAlert() bool`

GetRaiseObjectLevelFailureAlert returns the RaiseObjectLevelFailureAlert field if non-nil, zero value otherwise.

### GetRaiseObjectLevelFailureAlertOk

`func (o *AlertingConfig) GetRaiseObjectLevelFailureAlertOk() (*bool, bool)`

GetRaiseObjectLevelFailureAlertOk returns a tuple with the RaiseObjectLevelFailureAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaiseObjectLevelFailureAlert

`func (o *AlertingConfig) SetRaiseObjectLevelFailureAlert(v bool)`

SetRaiseObjectLevelFailureAlert sets RaiseObjectLevelFailureAlert field to given value.

### HasRaiseObjectLevelFailureAlert

`func (o *AlertingConfig) HasRaiseObjectLevelFailureAlert() bool`

HasRaiseObjectLevelFailureAlert returns a boolean if a field has been set.

### SetRaiseObjectLevelFailureAlertNil

`func (o *AlertingConfig) SetRaiseObjectLevelFailureAlertNil(b bool)`

 SetRaiseObjectLevelFailureAlertNil sets the value for RaiseObjectLevelFailureAlert to be an explicit nil

### UnsetRaiseObjectLevelFailureAlert
`func (o *AlertingConfig) UnsetRaiseObjectLevelFailureAlert()`

UnsetRaiseObjectLevelFailureAlert ensures that no value is present for RaiseObjectLevelFailureAlert, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


