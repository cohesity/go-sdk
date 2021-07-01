# AdDomainController

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupSupported** | Pointer to **NullableBool** | Specifies whether backup of this domain controller is supported. | [optional] 
**BackupUnsupportedReasons** | Pointer to **[]string** | Specifies any reason(s) for domain controller backup not supported. | [optional] 
**Domain** | Pointer to [**AdDomain**](AdDomain.md) |  | [optional] 
**HostName** | Pointer to **NullableString** | Specifies FQDN host name of the domain controller. | [optional] 
**IsGlobalCatalog** | Pointer to **NullableBool** | Specifies whether this domain controller is a global catalog server. | [optional] 
**IsReadOnly** | Pointer to **NullableBool** | Specifies whether this domain controller is read only. | [optional] 
**UtcOffsetMin** | Pointer to **NullableInt32** | Specifies UTC time offset of this domain controller in minutes. | [optional] 

## Methods

### NewAdDomainController

`func NewAdDomainController() *AdDomainController`

NewAdDomainController instantiates a new AdDomainController object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdDomainControllerWithDefaults

`func NewAdDomainControllerWithDefaults() *AdDomainController`

NewAdDomainControllerWithDefaults instantiates a new AdDomainController object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupSupported

`func (o *AdDomainController) GetBackupSupported() bool`

GetBackupSupported returns the BackupSupported field if non-nil, zero value otherwise.

### GetBackupSupportedOk

`func (o *AdDomainController) GetBackupSupportedOk() (*bool, bool)`

GetBackupSupportedOk returns a tuple with the BackupSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSupported

`func (o *AdDomainController) SetBackupSupported(v bool)`

SetBackupSupported sets BackupSupported field to given value.

### HasBackupSupported

`func (o *AdDomainController) HasBackupSupported() bool`

HasBackupSupported returns a boolean if a field has been set.

### SetBackupSupportedNil

`func (o *AdDomainController) SetBackupSupportedNil(b bool)`

 SetBackupSupportedNil sets the value for BackupSupported to be an explicit nil

### UnsetBackupSupported
`func (o *AdDomainController) UnsetBackupSupported()`

UnsetBackupSupported ensures that no value is present for BackupSupported, not even an explicit nil
### GetBackupUnsupportedReasons

`func (o *AdDomainController) GetBackupUnsupportedReasons() []string`

GetBackupUnsupportedReasons returns the BackupUnsupportedReasons field if non-nil, zero value otherwise.

### GetBackupUnsupportedReasonsOk

`func (o *AdDomainController) GetBackupUnsupportedReasonsOk() (*[]string, bool)`

GetBackupUnsupportedReasonsOk returns a tuple with the BackupUnsupportedReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupUnsupportedReasons

`func (o *AdDomainController) SetBackupUnsupportedReasons(v []string)`

SetBackupUnsupportedReasons sets BackupUnsupportedReasons field to given value.

### HasBackupUnsupportedReasons

`func (o *AdDomainController) HasBackupUnsupportedReasons() bool`

HasBackupUnsupportedReasons returns a boolean if a field has been set.

### SetBackupUnsupportedReasonsNil

`func (o *AdDomainController) SetBackupUnsupportedReasonsNil(b bool)`

 SetBackupUnsupportedReasonsNil sets the value for BackupUnsupportedReasons to be an explicit nil

### UnsetBackupUnsupportedReasons
`func (o *AdDomainController) UnsetBackupUnsupportedReasons()`

UnsetBackupUnsupportedReasons ensures that no value is present for BackupUnsupportedReasons, not even an explicit nil
### GetDomain

`func (o *AdDomainController) GetDomain() AdDomain`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *AdDomainController) GetDomainOk() (*AdDomain, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *AdDomainController) SetDomain(v AdDomain)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *AdDomainController) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetHostName

`func (o *AdDomainController) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *AdDomainController) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *AdDomainController) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *AdDomainController) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### SetHostNameNil

`func (o *AdDomainController) SetHostNameNil(b bool)`

 SetHostNameNil sets the value for HostName to be an explicit nil

### UnsetHostName
`func (o *AdDomainController) UnsetHostName()`

UnsetHostName ensures that no value is present for HostName, not even an explicit nil
### GetIsGlobalCatalog

`func (o *AdDomainController) GetIsGlobalCatalog() bool`

GetIsGlobalCatalog returns the IsGlobalCatalog field if non-nil, zero value otherwise.

### GetIsGlobalCatalogOk

`func (o *AdDomainController) GetIsGlobalCatalogOk() (*bool, bool)`

GetIsGlobalCatalogOk returns a tuple with the IsGlobalCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobalCatalog

`func (o *AdDomainController) SetIsGlobalCatalog(v bool)`

SetIsGlobalCatalog sets IsGlobalCatalog field to given value.

### HasIsGlobalCatalog

`func (o *AdDomainController) HasIsGlobalCatalog() bool`

HasIsGlobalCatalog returns a boolean if a field has been set.

### SetIsGlobalCatalogNil

`func (o *AdDomainController) SetIsGlobalCatalogNil(b bool)`

 SetIsGlobalCatalogNil sets the value for IsGlobalCatalog to be an explicit nil

### UnsetIsGlobalCatalog
`func (o *AdDomainController) UnsetIsGlobalCatalog()`

UnsetIsGlobalCatalog ensures that no value is present for IsGlobalCatalog, not even an explicit nil
### GetIsReadOnly

`func (o *AdDomainController) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *AdDomainController) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *AdDomainController) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *AdDomainController) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### SetIsReadOnlyNil

`func (o *AdDomainController) SetIsReadOnlyNil(b bool)`

 SetIsReadOnlyNil sets the value for IsReadOnly to be an explicit nil

### UnsetIsReadOnly
`func (o *AdDomainController) UnsetIsReadOnly()`

UnsetIsReadOnly ensures that no value is present for IsReadOnly, not even an explicit nil
### GetUtcOffsetMin

`func (o *AdDomainController) GetUtcOffsetMin() int32`

GetUtcOffsetMin returns the UtcOffsetMin field if non-nil, zero value otherwise.

### GetUtcOffsetMinOk

`func (o *AdDomainController) GetUtcOffsetMinOk() (*int32, bool)`

GetUtcOffsetMinOk returns a tuple with the UtcOffsetMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtcOffsetMin

`func (o *AdDomainController) SetUtcOffsetMin(v int32)`

SetUtcOffsetMin sets UtcOffsetMin field to given value.

### HasUtcOffsetMin

`func (o *AdDomainController) HasUtcOffsetMin() bool`

HasUtcOffsetMin returns a boolean if a field has been set.

### SetUtcOffsetMinNil

`func (o *AdDomainController) SetUtcOffsetMinNil(b bool)`

 SetUtcOffsetMinNil sets the value for UtcOffsetMin to be an explicit nil

### UnsetUtcOffsetMin
`func (o *AdDomainController) UnsetUtcOffsetMin()`

UnsetUtcOffsetMin ensures that no value is present for UtcOffsetMin, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


