# UserQuotaSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultUserQuotaPolicy** | Pointer to [**QuotaPolicy**](QuotaPolicy.md) |  | [optional] 
**EnableUserQuota** | Pointer to **NullableBool** | If set, it enables/disables the user quota overrides for a view. Otherwise, it leaves it at it&#39;s previous state. | [optional] 

## Methods

### NewUserQuotaSettings

`func NewUserQuotaSettings() *UserQuotaSettings`

NewUserQuotaSettings instantiates a new UserQuotaSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserQuotaSettingsWithDefaults

`func NewUserQuotaSettingsWithDefaults() *UserQuotaSettings`

NewUserQuotaSettingsWithDefaults instantiates a new UserQuotaSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultUserQuotaPolicy

`func (o *UserQuotaSettings) GetDefaultUserQuotaPolicy() QuotaPolicy`

GetDefaultUserQuotaPolicy returns the DefaultUserQuotaPolicy field if non-nil, zero value otherwise.

### GetDefaultUserQuotaPolicyOk

`func (o *UserQuotaSettings) GetDefaultUserQuotaPolicyOk() (*QuotaPolicy, bool)`

GetDefaultUserQuotaPolicyOk returns a tuple with the DefaultUserQuotaPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserQuotaPolicy

`func (o *UserQuotaSettings) SetDefaultUserQuotaPolicy(v QuotaPolicy)`

SetDefaultUserQuotaPolicy sets DefaultUserQuotaPolicy field to given value.

### HasDefaultUserQuotaPolicy

`func (o *UserQuotaSettings) HasDefaultUserQuotaPolicy() bool`

HasDefaultUserQuotaPolicy returns a boolean if a field has been set.

### GetEnableUserQuota

`func (o *UserQuotaSettings) GetEnableUserQuota() bool`

GetEnableUserQuota returns the EnableUserQuota field if non-nil, zero value otherwise.

### GetEnableUserQuotaOk

`func (o *UserQuotaSettings) GetEnableUserQuotaOk() (*bool, bool)`

GetEnableUserQuotaOk returns a tuple with the EnableUserQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUserQuota

`func (o *UserQuotaSettings) SetEnableUserQuota(v bool)`

SetEnableUserQuota sets EnableUserQuota field to given value.

### HasEnableUserQuota

`func (o *UserQuotaSettings) HasEnableUserQuota() bool`

HasEnableUserQuota returns a boolean if a field has been set.

### SetEnableUserQuotaNil

`func (o *UserQuotaSettings) SetEnableUserQuotaNil(b bool)`

 SetEnableUserQuotaNil sets the value for EnableUserQuota to be an explicit nil

### UnsetEnableUserQuota
`func (o *UserQuotaSettings) UnsetEnableUserQuota()`

UnsetEnableUserQuota ensures that no value is present for EnableUserQuota, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


