# ViewUserQuotaSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultQuotaPolicy** | Pointer to [**ViewUserQuotaSettingsDefaultQuotaPolicy**](ViewUserQuotaSettingsDefaultQuotaPolicy.md) |  | [optional] 
**Enabled** | **bool** | Specifies whether user quota is enabled for the View. | 

## Methods

### NewViewUserQuotaSettings

`func NewViewUserQuotaSettings(enabled bool, ) *ViewUserQuotaSettings`

NewViewUserQuotaSettings instantiates a new ViewUserQuotaSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewUserQuotaSettingsWithDefaults

`func NewViewUserQuotaSettingsWithDefaults() *ViewUserQuotaSettings`

NewViewUserQuotaSettingsWithDefaults instantiates a new ViewUserQuotaSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultQuotaPolicy

`func (o *ViewUserQuotaSettings) GetDefaultQuotaPolicy() ViewUserQuotaSettingsDefaultQuotaPolicy`

GetDefaultQuotaPolicy returns the DefaultQuotaPolicy field if non-nil, zero value otherwise.

### GetDefaultQuotaPolicyOk

`func (o *ViewUserQuotaSettings) GetDefaultQuotaPolicyOk() (*ViewUserQuotaSettingsDefaultQuotaPolicy, bool)`

GetDefaultQuotaPolicyOk returns a tuple with the DefaultQuotaPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultQuotaPolicy

`func (o *ViewUserQuotaSettings) SetDefaultQuotaPolicy(v ViewUserQuotaSettingsDefaultQuotaPolicy)`

SetDefaultQuotaPolicy sets DefaultQuotaPolicy field to given value.

### HasDefaultQuotaPolicy

`func (o *ViewUserQuotaSettings) HasDefaultQuotaPolicy() bool`

HasDefaultQuotaPolicy returns a boolean if a field has been set.

### GetEnabled

`func (o *ViewUserQuotaSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ViewUserQuotaSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ViewUserQuotaSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


