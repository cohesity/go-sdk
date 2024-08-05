# ViewUserQuotaSettingsDefaultQuotaPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertLimitBytes** | Pointer to **NullableInt64** | Specifies if an alert should be triggered when the usage of this resource exceeds this quota limit. This limit is optional and is specified in bytes. If no value is specified, there is no limit. | [optional] 
**AlertThresholdPercentage** | Pointer to **NullableInt64** | Supported only for user quota policy. Specifies when the usage goes above an alert threshold percentage which is: HardLimitBytes * AlertThresholdPercentage, eg: 80% of HardLimitBytes Can only be set if HardLimitBytes is set. Cannot be set if AlertLimitBytes is already set. | [optional] 
**HardLimitBytes** | Pointer to **NullableInt64** | Specifies an optional quota limit on the usage allowed for this resource. This limit is specified in bytes. If no value is specified, there is no limit. | [optional] 

## Methods

### NewViewUserQuotaSettingsDefaultQuotaPolicy

`func NewViewUserQuotaSettingsDefaultQuotaPolicy() *ViewUserQuotaSettingsDefaultQuotaPolicy`

NewViewUserQuotaSettingsDefaultQuotaPolicy instantiates a new ViewUserQuotaSettingsDefaultQuotaPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewUserQuotaSettingsDefaultQuotaPolicyWithDefaults

`func NewViewUserQuotaSettingsDefaultQuotaPolicyWithDefaults() *ViewUserQuotaSettingsDefaultQuotaPolicy`

NewViewUserQuotaSettingsDefaultQuotaPolicyWithDefaults instantiates a new ViewUserQuotaSettingsDefaultQuotaPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertLimitBytes

`func (o *ViewUserQuotaSettingsDefaultQuotaPolicy) GetAlertLimitBytes() int64`

GetAlertLimitBytes returns the AlertLimitBytes field if non-nil, zero value otherwise.

### GetAlertLimitBytesOk

`func (o *ViewUserQuotaSettingsDefaultQuotaPolicy) GetAlertLimitBytesOk() (*int64, bool)`

GetAlertLimitBytesOk returns a tuple with the AlertLimitBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertLimitBytes

`func (o *ViewUserQuotaSettingsDefaultQuotaPolicy) SetAlertLimitBytes(v int64)`

SetAlertLimitBytes sets AlertLimitBytes field to given value.

### HasAlertLimitBytes

`func (o *ViewUserQuotaSettingsDefaultQuotaPolicy) HasAlertLimitBytes() bool`

HasAlertLimitBytes returns a boolean if a field has been set.

### SetAlertLimitBytesNil

`func (o *ViewUserQuotaSettingsDefaultQuotaPolicy) SetAlertLimitBytesNil(b bool)`

 SetAlertLimitBytesNil sets the value for AlertLimitBytes to be an explicit nil

### UnsetAlertLimitBytes
`func (o *ViewUserQuotaSettingsDefaultQuotaPolicy) UnsetAlertLimitBytes()`

UnsetAlertLimitBytes ensures that no value is present for AlertLimitBytes, not even an explicit nil
### GetAlertThresholdPercentage

`func (o *ViewUserQuotaSettingsDefaultQuotaPolicy) GetAlertThresholdPercentage() int64`

GetAlertThresholdPercentage returns the AlertThresholdPercentage field if non-nil, zero value otherwise.

### GetAlertThresholdPercentageOk

`func (o *ViewUserQuotaSettingsDefaultQuotaPolicy) GetAlertThresholdPercentageOk() (*int64, bool)`

GetAlertThresholdPercentageOk returns a tuple with the AlertThresholdPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertThresholdPercentage

`func (o *ViewUserQuotaSettingsDefaultQuotaPolicy) SetAlertThresholdPercentage(v int64)`

SetAlertThresholdPercentage sets AlertThresholdPercentage field to given value.

### HasAlertThresholdPercentage

`func (o *ViewUserQuotaSettingsDefaultQuotaPolicy) HasAlertThresholdPercentage() bool`

HasAlertThresholdPercentage returns a boolean if a field has been set.

### SetAlertThresholdPercentageNil

`func (o *ViewUserQuotaSettingsDefaultQuotaPolicy) SetAlertThresholdPercentageNil(b bool)`

 SetAlertThresholdPercentageNil sets the value for AlertThresholdPercentage to be an explicit nil

### UnsetAlertThresholdPercentage
`func (o *ViewUserQuotaSettingsDefaultQuotaPolicy) UnsetAlertThresholdPercentage()`

UnsetAlertThresholdPercentage ensures that no value is present for AlertThresholdPercentage, not even an explicit nil
### GetHardLimitBytes

`func (o *ViewUserQuotaSettingsDefaultQuotaPolicy) GetHardLimitBytes() int64`

GetHardLimitBytes returns the HardLimitBytes field if non-nil, zero value otherwise.

### GetHardLimitBytesOk

`func (o *ViewUserQuotaSettingsDefaultQuotaPolicy) GetHardLimitBytesOk() (*int64, bool)`

GetHardLimitBytesOk returns a tuple with the HardLimitBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardLimitBytes

`func (o *ViewUserQuotaSettingsDefaultQuotaPolicy) SetHardLimitBytes(v int64)`

SetHardLimitBytes sets HardLimitBytes field to given value.

### HasHardLimitBytes

`func (o *ViewUserQuotaSettingsDefaultQuotaPolicy) HasHardLimitBytes() bool`

HasHardLimitBytes returns a boolean if a field has been set.

### SetHardLimitBytesNil

`func (o *ViewUserQuotaSettingsDefaultQuotaPolicy) SetHardLimitBytesNil(b bool)`

 SetHardLimitBytesNil sets the value for HardLimitBytes to be an explicit nil

### UnsetHardLimitBytes
`func (o *ViewUserQuotaSettingsDefaultQuotaPolicy) UnsetHardLimitBytes()`

UnsetHardLimitBytes ensures that no value is present for HardLimitBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


