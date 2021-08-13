# UserUiConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Preferences** | Pointer to **NullableString** | Specifies the user&#39;s preferences for UI customization. | [optional] 
**Locale** | Pointer to **NullableString** | Specifies the locale. | [optional] 

## Methods

### NewUserUiConfig

`func NewUserUiConfig() *UserUiConfig`

NewUserUiConfig instantiates a new UserUiConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUiConfigWithDefaults

`func NewUserUiConfigWithDefaults() *UserUiConfig`

NewUserUiConfigWithDefaults instantiates a new UserUiConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreferences

`func (o *UserUiConfig) GetPreferences() string`

GetPreferences returns the Preferences field if non-nil, zero value otherwise.

### GetPreferencesOk

`func (o *UserUiConfig) GetPreferencesOk() (*string, bool)`

GetPreferencesOk returns a tuple with the Preferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferences

`func (o *UserUiConfig) SetPreferences(v string)`

SetPreferences sets Preferences field to given value.

### HasPreferences

`func (o *UserUiConfig) HasPreferences() bool`

HasPreferences returns a boolean if a field has been set.

### SetPreferencesNil

`func (o *UserUiConfig) SetPreferencesNil(b bool)`

 SetPreferencesNil sets the value for Preferences to be an explicit nil

### UnsetPreferences
`func (o *UserUiConfig) UnsetPreferences()`

UnsetPreferences ensures that no value is present for Preferences, not even an explicit nil
### GetLocale

`func (o *UserUiConfig) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *UserUiConfig) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *UserUiConfig) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *UserUiConfig) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *UserUiConfig) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *UserUiConfig) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


