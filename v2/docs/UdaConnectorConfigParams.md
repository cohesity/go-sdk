# UdaConnectorConfigParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndexConfig** | Pointer to **NullableString** | Specifies the index config as json string. | [optional] 
**OsSpecificConfig** | Pointer to [**[]UdaOSSpecificConfigParams**](UdaOSSpecificConfigParams.md) | Specifies operating system specific configuration. | [optional] 
**Replace** | Pointer to **NullableBool** | If true, any existing connector config with the the same ID is replaced if it exists. | [optional] 
**UiTranslationConfig** | Pointer to [**[]UdaLocaleSpecificTranslations**](UdaLocaleSpecificTranslations.md) | Specifies the translation messages for various locales. It maps locale name to their respective translation json strings. | [optional] 

## Methods

### NewUdaConnectorConfigParams

`func NewUdaConnectorConfigParams() *UdaConnectorConfigParams`

NewUdaConnectorConfigParams instantiates a new UdaConnectorConfigParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUdaConnectorConfigParamsWithDefaults

`func NewUdaConnectorConfigParamsWithDefaults() *UdaConnectorConfigParams`

NewUdaConnectorConfigParamsWithDefaults instantiates a new UdaConnectorConfigParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexConfig

`func (o *UdaConnectorConfigParams) GetIndexConfig() string`

GetIndexConfig returns the IndexConfig field if non-nil, zero value otherwise.

### GetIndexConfigOk

`func (o *UdaConnectorConfigParams) GetIndexConfigOk() (*string, bool)`

GetIndexConfigOk returns a tuple with the IndexConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexConfig

`func (o *UdaConnectorConfigParams) SetIndexConfig(v string)`

SetIndexConfig sets IndexConfig field to given value.

### HasIndexConfig

`func (o *UdaConnectorConfigParams) HasIndexConfig() bool`

HasIndexConfig returns a boolean if a field has been set.

### SetIndexConfigNil

`func (o *UdaConnectorConfigParams) SetIndexConfigNil(b bool)`

 SetIndexConfigNil sets the value for IndexConfig to be an explicit nil

### UnsetIndexConfig
`func (o *UdaConnectorConfigParams) UnsetIndexConfig()`

UnsetIndexConfig ensures that no value is present for IndexConfig, not even an explicit nil
### GetOsSpecificConfig

`func (o *UdaConnectorConfigParams) GetOsSpecificConfig() []UdaOSSpecificConfigParams`

GetOsSpecificConfig returns the OsSpecificConfig field if non-nil, zero value otherwise.

### GetOsSpecificConfigOk

`func (o *UdaConnectorConfigParams) GetOsSpecificConfigOk() (*[]UdaOSSpecificConfigParams, bool)`

GetOsSpecificConfigOk returns a tuple with the OsSpecificConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsSpecificConfig

`func (o *UdaConnectorConfigParams) SetOsSpecificConfig(v []UdaOSSpecificConfigParams)`

SetOsSpecificConfig sets OsSpecificConfig field to given value.

### HasOsSpecificConfig

`func (o *UdaConnectorConfigParams) HasOsSpecificConfig() bool`

HasOsSpecificConfig returns a boolean if a field has been set.

### GetReplace

`func (o *UdaConnectorConfigParams) GetReplace() bool`

GetReplace returns the Replace field if non-nil, zero value otherwise.

### GetReplaceOk

`func (o *UdaConnectorConfigParams) GetReplaceOk() (*bool, bool)`

GetReplaceOk returns a tuple with the Replace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplace

`func (o *UdaConnectorConfigParams) SetReplace(v bool)`

SetReplace sets Replace field to given value.

### HasReplace

`func (o *UdaConnectorConfigParams) HasReplace() bool`

HasReplace returns a boolean if a field has been set.

### SetReplaceNil

`func (o *UdaConnectorConfigParams) SetReplaceNil(b bool)`

 SetReplaceNil sets the value for Replace to be an explicit nil

### UnsetReplace
`func (o *UdaConnectorConfigParams) UnsetReplace()`

UnsetReplace ensures that no value is present for Replace, not even an explicit nil
### GetUiTranslationConfig

`func (o *UdaConnectorConfigParams) GetUiTranslationConfig() []UdaLocaleSpecificTranslations`

GetUiTranslationConfig returns the UiTranslationConfig field if non-nil, zero value otherwise.

### GetUiTranslationConfigOk

`func (o *UdaConnectorConfigParams) GetUiTranslationConfigOk() (*[]UdaLocaleSpecificTranslations, bool)`

GetUiTranslationConfigOk returns a tuple with the UiTranslationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiTranslationConfig

`func (o *UdaConnectorConfigParams) SetUiTranslationConfig(v []UdaLocaleSpecificTranslations)`

SetUiTranslationConfig sets UiTranslationConfig field to given value.

### HasUiTranslationConfig

`func (o *UdaConnectorConfigParams) HasUiTranslationConfig() bool`

HasUiTranslationConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


