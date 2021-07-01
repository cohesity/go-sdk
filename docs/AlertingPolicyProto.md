# AlertingPolicyProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryTargetVec** | Pointer to [**[]DeliveryRuleProtoDeliveryTarget**](DeliveryRuleProtoDeliveryTarget.md) | The delivery targets to be alerted. | [optional] 
**Emails** | Pointer to **[]string** | The email addresses to send alerts to. This field has been deprecated in favor of the field delivery_target_vec. The clients should take care to ensure that the emails stored in here are migrated to that field, or else utilise both the fields when trying to obtain the complete list of delivery targets. | [optional] 
**Policy** | Pointer to **NullableInt32** | &#39;policy&#39; is declared as int32 because ORing the enums will generate values which are invalid as enums. Protobuf doesn&#39;t allow those invalid enums to be set. | [optional] 
**RaiseObjectLevelFailureAlert** | Pointer to **NullableBool** | Raise per object alert for failures. | [optional] 

## Methods

### NewAlertingPolicyProto

`func NewAlertingPolicyProto() *AlertingPolicyProto`

NewAlertingPolicyProto instantiates a new AlertingPolicyProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertingPolicyProtoWithDefaults

`func NewAlertingPolicyProtoWithDefaults() *AlertingPolicyProto`

NewAlertingPolicyProtoWithDefaults instantiates a new AlertingPolicyProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveryTargetVec

`func (o *AlertingPolicyProto) GetDeliveryTargetVec() []DeliveryRuleProtoDeliveryTarget`

GetDeliveryTargetVec returns the DeliveryTargetVec field if non-nil, zero value otherwise.

### GetDeliveryTargetVecOk

`func (o *AlertingPolicyProto) GetDeliveryTargetVecOk() (*[]DeliveryRuleProtoDeliveryTarget, bool)`

GetDeliveryTargetVecOk returns a tuple with the DeliveryTargetVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTargetVec

`func (o *AlertingPolicyProto) SetDeliveryTargetVec(v []DeliveryRuleProtoDeliveryTarget)`

SetDeliveryTargetVec sets DeliveryTargetVec field to given value.

### HasDeliveryTargetVec

`func (o *AlertingPolicyProto) HasDeliveryTargetVec() bool`

HasDeliveryTargetVec returns a boolean if a field has been set.

### SetDeliveryTargetVecNil

`func (o *AlertingPolicyProto) SetDeliveryTargetVecNil(b bool)`

 SetDeliveryTargetVecNil sets the value for DeliveryTargetVec to be an explicit nil

### UnsetDeliveryTargetVec
`func (o *AlertingPolicyProto) UnsetDeliveryTargetVec()`

UnsetDeliveryTargetVec ensures that no value is present for DeliveryTargetVec, not even an explicit nil
### GetEmails

`func (o *AlertingPolicyProto) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *AlertingPolicyProto) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *AlertingPolicyProto) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *AlertingPolicyProto) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### SetEmailsNil

`func (o *AlertingPolicyProto) SetEmailsNil(b bool)`

 SetEmailsNil sets the value for Emails to be an explicit nil

### UnsetEmails
`func (o *AlertingPolicyProto) UnsetEmails()`

UnsetEmails ensures that no value is present for Emails, not even an explicit nil
### GetPolicy

`func (o *AlertingPolicyProto) GetPolicy() int32`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *AlertingPolicyProto) GetPolicyOk() (*int32, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *AlertingPolicyProto) SetPolicy(v int32)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *AlertingPolicyProto) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicyNil

`func (o *AlertingPolicyProto) SetPolicyNil(b bool)`

 SetPolicyNil sets the value for Policy to be an explicit nil

### UnsetPolicy
`func (o *AlertingPolicyProto) UnsetPolicy()`

UnsetPolicy ensures that no value is present for Policy, not even an explicit nil
### GetRaiseObjectLevelFailureAlert

`func (o *AlertingPolicyProto) GetRaiseObjectLevelFailureAlert() bool`

GetRaiseObjectLevelFailureAlert returns the RaiseObjectLevelFailureAlert field if non-nil, zero value otherwise.

### GetRaiseObjectLevelFailureAlertOk

`func (o *AlertingPolicyProto) GetRaiseObjectLevelFailureAlertOk() (*bool, bool)`

GetRaiseObjectLevelFailureAlertOk returns a tuple with the RaiseObjectLevelFailureAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaiseObjectLevelFailureAlert

`func (o *AlertingPolicyProto) SetRaiseObjectLevelFailureAlert(v bool)`

SetRaiseObjectLevelFailureAlert sets RaiseObjectLevelFailureAlert field to given value.

### HasRaiseObjectLevelFailureAlert

`func (o *AlertingPolicyProto) HasRaiseObjectLevelFailureAlert() bool`

HasRaiseObjectLevelFailureAlert returns a boolean if a field has been set.

### SetRaiseObjectLevelFailureAlertNil

`func (o *AlertingPolicyProto) SetRaiseObjectLevelFailureAlertNil(b bool)`

 SetRaiseObjectLevelFailureAlertNil sets the value for RaiseObjectLevelFailureAlert to be an explicit nil

### UnsetRaiseObjectLevelFailureAlert
`func (o *AlertingPolicyProto) UnsetRaiseObjectLevelFailureAlert()`

UnsetRaiseObjectLevelFailureAlert ensures that no value is present for RaiseObjectLevelFailureAlert, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


