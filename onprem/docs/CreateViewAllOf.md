# CreateViewAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageDomainId** | Pointer to **NullableInt64** | Specifies the id of the Storage Domain (View Box) where the View will be created. | [optional] 
**CaseInsensitiveNamesEnabled** | Pointer to **NullableBool** | Specifies whether to support case insensitive file/folder names. This parameter can only be set during create and cannot be changed. | [optional] 
**ObjectServicesMappingConfig** | Pointer to **NullableString** | Specifies the Object Services key mapping config of the view. This parameter can only be set during create and cannot be changed. Configuration of Object Services key mapping. Specifies the type of Object Services key mapping config. | [optional] 
**ViewProtectionConfig** | Pointer to [**ViewProtectionConfig**](ViewProtectionConfig.md) | Specifies the protection config of the View. | [optional] 

## Methods

### NewCreateViewAllOf

`func NewCreateViewAllOf() *CreateViewAllOf`

NewCreateViewAllOf instantiates a new CreateViewAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateViewAllOfWithDefaults

`func NewCreateViewAllOfWithDefaults() *CreateViewAllOf`

NewCreateViewAllOfWithDefaults instantiates a new CreateViewAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageDomainId

`func (o *CreateViewAllOf) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *CreateViewAllOf) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *CreateViewAllOf) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *CreateViewAllOf) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *CreateViewAllOf) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *CreateViewAllOf) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
### GetCaseInsensitiveNamesEnabled

`func (o *CreateViewAllOf) GetCaseInsensitiveNamesEnabled() bool`

GetCaseInsensitiveNamesEnabled returns the CaseInsensitiveNamesEnabled field if non-nil, zero value otherwise.

### GetCaseInsensitiveNamesEnabledOk

`func (o *CreateViewAllOf) GetCaseInsensitiveNamesEnabledOk() (*bool, bool)`

GetCaseInsensitiveNamesEnabledOk returns a tuple with the CaseInsensitiveNamesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInsensitiveNamesEnabled

`func (o *CreateViewAllOf) SetCaseInsensitiveNamesEnabled(v bool)`

SetCaseInsensitiveNamesEnabled sets CaseInsensitiveNamesEnabled field to given value.

### HasCaseInsensitiveNamesEnabled

`func (o *CreateViewAllOf) HasCaseInsensitiveNamesEnabled() bool`

HasCaseInsensitiveNamesEnabled returns a boolean if a field has been set.

### SetCaseInsensitiveNamesEnabledNil

`func (o *CreateViewAllOf) SetCaseInsensitiveNamesEnabledNil(b bool)`

 SetCaseInsensitiveNamesEnabledNil sets the value for CaseInsensitiveNamesEnabled to be an explicit nil

### UnsetCaseInsensitiveNamesEnabled
`func (o *CreateViewAllOf) UnsetCaseInsensitiveNamesEnabled()`

UnsetCaseInsensitiveNamesEnabled ensures that no value is present for CaseInsensitiveNamesEnabled, not even an explicit nil
### GetObjectServicesMappingConfig

`func (o *CreateViewAllOf) GetObjectServicesMappingConfig() string`

GetObjectServicesMappingConfig returns the ObjectServicesMappingConfig field if non-nil, zero value otherwise.

### GetObjectServicesMappingConfigOk

`func (o *CreateViewAllOf) GetObjectServicesMappingConfigOk() (*string, bool)`

GetObjectServicesMappingConfigOk returns a tuple with the ObjectServicesMappingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectServicesMappingConfig

`func (o *CreateViewAllOf) SetObjectServicesMappingConfig(v string)`

SetObjectServicesMappingConfig sets ObjectServicesMappingConfig field to given value.

### HasObjectServicesMappingConfig

`func (o *CreateViewAllOf) HasObjectServicesMappingConfig() bool`

HasObjectServicesMappingConfig returns a boolean if a field has been set.

### SetObjectServicesMappingConfigNil

`func (o *CreateViewAllOf) SetObjectServicesMappingConfigNil(b bool)`

 SetObjectServicesMappingConfigNil sets the value for ObjectServicesMappingConfig to be an explicit nil

### UnsetObjectServicesMappingConfig
`func (o *CreateViewAllOf) UnsetObjectServicesMappingConfig()`

UnsetObjectServicesMappingConfig ensures that no value is present for ObjectServicesMappingConfig, not even an explicit nil
### GetViewProtectionConfig

`func (o *CreateViewAllOf) GetViewProtectionConfig() ViewProtectionConfig`

GetViewProtectionConfig returns the ViewProtectionConfig field if non-nil, zero value otherwise.

### GetViewProtectionConfigOk

`func (o *CreateViewAllOf) GetViewProtectionConfigOk() (*ViewProtectionConfig, bool)`

GetViewProtectionConfigOk returns a tuple with the ViewProtectionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewProtectionConfig

`func (o *CreateViewAllOf) SetViewProtectionConfig(v ViewProtectionConfig)`

SetViewProtectionConfig sets ViewProtectionConfig field to given value.

### HasViewProtectionConfig

`func (o *CreateViewAllOf) HasViewProtectionConfig() bool`

HasViewProtectionConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


