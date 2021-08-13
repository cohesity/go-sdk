# ViewUserQuotasConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **NullableBool** | Specifies whether user quota is enabled for the View. | 
**DefaultQuotaPolicy** | Pointer to [**QuotaPolicy**](QuotaPolicy.md) | Specifies the default user quota policy. | [optional] 

## Methods

### NewViewUserQuotasConfig

`func NewViewUserQuotasConfig(enabled NullableBool, ) *ViewUserQuotasConfig`

NewViewUserQuotasConfig instantiates a new ViewUserQuotasConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewUserQuotasConfigWithDefaults

`func NewViewUserQuotasConfigWithDefaults() *ViewUserQuotasConfig`

NewViewUserQuotasConfigWithDefaults instantiates a new ViewUserQuotasConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ViewUserQuotasConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ViewUserQuotasConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ViewUserQuotasConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### SetEnabledNil

`func (o *ViewUserQuotasConfig) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *ViewUserQuotasConfig) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetDefaultQuotaPolicy

`func (o *ViewUserQuotasConfig) GetDefaultQuotaPolicy() QuotaPolicy`

GetDefaultQuotaPolicy returns the DefaultQuotaPolicy field if non-nil, zero value otherwise.

### GetDefaultQuotaPolicyOk

`func (o *ViewUserQuotasConfig) GetDefaultQuotaPolicyOk() (*QuotaPolicy, bool)`

GetDefaultQuotaPolicyOk returns a tuple with the DefaultQuotaPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultQuotaPolicy

`func (o *ViewUserQuotasConfig) SetDefaultQuotaPolicy(v QuotaPolicy)`

SetDefaultQuotaPolicy sets DefaultQuotaPolicy field to given value.

### HasDefaultQuotaPolicy

`func (o *ViewUserQuotasConfig) HasDefaultQuotaPolicy() bool`

HasDefaultQuotaPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


