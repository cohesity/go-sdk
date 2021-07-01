# UserPreferencesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Preferences** | Pointer to **map[string]string** | Preferences is a key-value map of preferences. | [optional] 

## Methods

### NewUserPreferencesResult

`func NewUserPreferencesResult() *UserPreferencesResult`

NewUserPreferencesResult instantiates a new UserPreferencesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPreferencesResultWithDefaults

`func NewUserPreferencesResultWithDefaults() *UserPreferencesResult`

NewUserPreferencesResultWithDefaults instantiates a new UserPreferencesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreferences

`func (o *UserPreferencesResult) GetPreferences() map[string]string`

GetPreferences returns the Preferences field if non-nil, zero value otherwise.

### GetPreferencesOk

`func (o *UserPreferencesResult) GetPreferencesOk() (*map[string]string, bool)`

GetPreferencesOk returns a tuple with the Preferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferences

`func (o *UserPreferencesResult) SetPreferences(v map[string]string)`

SetPreferences sets Preferences field to given value.

### HasPreferences

`func (o *UserPreferencesResult) HasPreferences() bool`

HasPreferences returns a boolean if a field has been set.

### SetPreferencesNil

`func (o *UserPreferencesResult) SetPreferencesNil(b bool)`

 SetPreferencesNil sets the value for Preferences to be an explicit nil

### UnsetPreferences
`func (o *UserPreferencesResult) UnsetPreferences()`

UnsetPreferences ensures that no value is present for Preferences, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


