# UpdateUserQuotaSettingsForView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultUserQuotaPolicy** | Pointer to [**QuotaPolicy**](QuotaPolicy.md) |  | [optional] 
**EnableUserQuota** | Pointer to **NullableBool** | If set, it enables/disables the user quota overrides for a view. Otherwise, it leaves it at it&#39;s previous state. | [optional] 
**InheritDefaultPolicyFromViewbox** | Pointer to **NullableBool** | If set to true, the default_policy in view metadata will be cleared and the default policy from viewbox will take effect for all users in the view. | [optional] 
**ViewName** | Pointer to **NullableString** | View name of input view. | [optional] 

## Methods

### NewUpdateUserQuotaSettingsForView

`func NewUpdateUserQuotaSettingsForView() *UpdateUserQuotaSettingsForView`

NewUpdateUserQuotaSettingsForView instantiates a new UpdateUserQuotaSettingsForView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserQuotaSettingsForViewWithDefaults

`func NewUpdateUserQuotaSettingsForViewWithDefaults() *UpdateUserQuotaSettingsForView`

NewUpdateUserQuotaSettingsForViewWithDefaults instantiates a new UpdateUserQuotaSettingsForView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultUserQuotaPolicy

`func (o *UpdateUserQuotaSettingsForView) GetDefaultUserQuotaPolicy() QuotaPolicy`

GetDefaultUserQuotaPolicy returns the DefaultUserQuotaPolicy field if non-nil, zero value otherwise.

### GetDefaultUserQuotaPolicyOk

`func (o *UpdateUserQuotaSettingsForView) GetDefaultUserQuotaPolicyOk() (*QuotaPolicy, bool)`

GetDefaultUserQuotaPolicyOk returns a tuple with the DefaultUserQuotaPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserQuotaPolicy

`func (o *UpdateUserQuotaSettingsForView) SetDefaultUserQuotaPolicy(v QuotaPolicy)`

SetDefaultUserQuotaPolicy sets DefaultUserQuotaPolicy field to given value.

### HasDefaultUserQuotaPolicy

`func (o *UpdateUserQuotaSettingsForView) HasDefaultUserQuotaPolicy() bool`

HasDefaultUserQuotaPolicy returns a boolean if a field has been set.

### GetEnableUserQuota

`func (o *UpdateUserQuotaSettingsForView) GetEnableUserQuota() bool`

GetEnableUserQuota returns the EnableUserQuota field if non-nil, zero value otherwise.

### GetEnableUserQuotaOk

`func (o *UpdateUserQuotaSettingsForView) GetEnableUserQuotaOk() (*bool, bool)`

GetEnableUserQuotaOk returns a tuple with the EnableUserQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUserQuota

`func (o *UpdateUserQuotaSettingsForView) SetEnableUserQuota(v bool)`

SetEnableUserQuota sets EnableUserQuota field to given value.

### HasEnableUserQuota

`func (o *UpdateUserQuotaSettingsForView) HasEnableUserQuota() bool`

HasEnableUserQuota returns a boolean if a field has been set.

### SetEnableUserQuotaNil

`func (o *UpdateUserQuotaSettingsForView) SetEnableUserQuotaNil(b bool)`

 SetEnableUserQuotaNil sets the value for EnableUserQuota to be an explicit nil

### UnsetEnableUserQuota
`func (o *UpdateUserQuotaSettingsForView) UnsetEnableUserQuota()`

UnsetEnableUserQuota ensures that no value is present for EnableUserQuota, not even an explicit nil
### GetInheritDefaultPolicyFromViewbox

`func (o *UpdateUserQuotaSettingsForView) GetInheritDefaultPolicyFromViewbox() bool`

GetInheritDefaultPolicyFromViewbox returns the InheritDefaultPolicyFromViewbox field if non-nil, zero value otherwise.

### GetInheritDefaultPolicyFromViewboxOk

`func (o *UpdateUserQuotaSettingsForView) GetInheritDefaultPolicyFromViewboxOk() (*bool, bool)`

GetInheritDefaultPolicyFromViewboxOk returns a tuple with the InheritDefaultPolicyFromViewbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritDefaultPolicyFromViewbox

`func (o *UpdateUserQuotaSettingsForView) SetInheritDefaultPolicyFromViewbox(v bool)`

SetInheritDefaultPolicyFromViewbox sets InheritDefaultPolicyFromViewbox field to given value.

### HasInheritDefaultPolicyFromViewbox

`func (o *UpdateUserQuotaSettingsForView) HasInheritDefaultPolicyFromViewbox() bool`

HasInheritDefaultPolicyFromViewbox returns a boolean if a field has been set.

### SetInheritDefaultPolicyFromViewboxNil

`func (o *UpdateUserQuotaSettingsForView) SetInheritDefaultPolicyFromViewboxNil(b bool)`

 SetInheritDefaultPolicyFromViewboxNil sets the value for InheritDefaultPolicyFromViewbox to be an explicit nil

### UnsetInheritDefaultPolicyFromViewbox
`func (o *UpdateUserQuotaSettingsForView) UnsetInheritDefaultPolicyFromViewbox()`

UnsetInheritDefaultPolicyFromViewbox ensures that no value is present for InheritDefaultPolicyFromViewbox, not even an explicit nil
### GetViewName

`func (o *UpdateUserQuotaSettingsForView) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *UpdateUserQuotaSettingsForView) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *UpdateUserQuotaSettingsForView) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *UpdateUserQuotaSettingsForView) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *UpdateUserQuotaSettingsForView) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *UpdateUserQuotaSettingsForView) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


