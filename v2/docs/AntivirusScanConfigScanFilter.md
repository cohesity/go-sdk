# AntivirusScanConfigScanFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileExtensionsList** | Pointer to **[]string** | The list of file extensions to apply | [optional] 
**IsEnabled** | Pointer to **NullableBool** | If set, it enables the file extension filter | [optional] 
**Mode** | Pointer to **NullableString** | The mode applied to the list of file extensions &#39;Whitelist&#39; indicates a whitelist extension filter. &#39;Blacklist&#39; indicates a blacklist extension filter. | [optional] 

## Methods

### NewAntivirusScanConfigScanFilter

`func NewAntivirusScanConfigScanFilter() *AntivirusScanConfigScanFilter`

NewAntivirusScanConfigScanFilter instantiates a new AntivirusScanConfigScanFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAntivirusScanConfigScanFilterWithDefaults

`func NewAntivirusScanConfigScanFilterWithDefaults() *AntivirusScanConfigScanFilter`

NewAntivirusScanConfigScanFilterWithDefaults instantiates a new AntivirusScanConfigScanFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileExtensionsList

`func (o *AntivirusScanConfigScanFilter) GetFileExtensionsList() []string`

GetFileExtensionsList returns the FileExtensionsList field if non-nil, zero value otherwise.

### GetFileExtensionsListOk

`func (o *AntivirusScanConfigScanFilter) GetFileExtensionsListOk() (*[]string, bool)`

GetFileExtensionsListOk returns a tuple with the FileExtensionsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExtensionsList

`func (o *AntivirusScanConfigScanFilter) SetFileExtensionsList(v []string)`

SetFileExtensionsList sets FileExtensionsList field to given value.

### HasFileExtensionsList

`func (o *AntivirusScanConfigScanFilter) HasFileExtensionsList() bool`

HasFileExtensionsList returns a boolean if a field has been set.

### SetFileExtensionsListNil

`func (o *AntivirusScanConfigScanFilter) SetFileExtensionsListNil(b bool)`

 SetFileExtensionsListNil sets the value for FileExtensionsList to be an explicit nil

### UnsetFileExtensionsList
`func (o *AntivirusScanConfigScanFilter) UnsetFileExtensionsList()`

UnsetFileExtensionsList ensures that no value is present for FileExtensionsList, not even an explicit nil
### GetIsEnabled

`func (o *AntivirusScanConfigScanFilter) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *AntivirusScanConfigScanFilter) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *AntivirusScanConfigScanFilter) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *AntivirusScanConfigScanFilter) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *AntivirusScanConfigScanFilter) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *AntivirusScanConfigScanFilter) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetMode

`func (o *AntivirusScanConfigScanFilter) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *AntivirusScanConfigScanFilter) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *AntivirusScanConfigScanFilter) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *AntivirusScanConfigScanFilter) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *AntivirusScanConfigScanFilter) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *AntivirusScanConfigScanFilter) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


