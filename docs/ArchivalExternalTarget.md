# ArchivalExternalTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VaultId** | Pointer to **NullableInt64** | Specifies the id of Archival Vault assigned by the Cohesity Cluster. | [optional] 
**VaultName** | Pointer to **NullableString** | Name of the Archival Vault. | [optional] 
**VaultType** | Pointer to **NullableString** | Specifies the type of the Archival External Target such as &#39;kCloud&#39;, &#39;kTape&#39; or &#39;kNas&#39;. &#39;kCloud&#39; indicates the archival location as Cloud. &#39;kTape&#39; indicates the archival location as Tape. &#39;kNas&#39; indicates the archival location as Network Attached Storage (Nas). | [optional] 

## Methods

### NewArchivalExternalTarget

`func NewArchivalExternalTarget() *ArchivalExternalTarget`

NewArchivalExternalTarget instantiates a new ArchivalExternalTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchivalExternalTargetWithDefaults

`func NewArchivalExternalTargetWithDefaults() *ArchivalExternalTarget`

NewArchivalExternalTargetWithDefaults instantiates a new ArchivalExternalTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVaultId

`func (o *ArchivalExternalTarget) GetVaultId() int64`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *ArchivalExternalTarget) GetVaultIdOk() (*int64, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *ArchivalExternalTarget) SetVaultId(v int64)`

SetVaultId sets VaultId field to given value.

### HasVaultId

`func (o *ArchivalExternalTarget) HasVaultId() bool`

HasVaultId returns a boolean if a field has been set.

### SetVaultIdNil

`func (o *ArchivalExternalTarget) SetVaultIdNil(b bool)`

 SetVaultIdNil sets the value for VaultId to be an explicit nil

### UnsetVaultId
`func (o *ArchivalExternalTarget) UnsetVaultId()`

UnsetVaultId ensures that no value is present for VaultId, not even an explicit nil
### GetVaultName

`func (o *ArchivalExternalTarget) GetVaultName() string`

GetVaultName returns the VaultName field if non-nil, zero value otherwise.

### GetVaultNameOk

`func (o *ArchivalExternalTarget) GetVaultNameOk() (*string, bool)`

GetVaultNameOk returns a tuple with the VaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultName

`func (o *ArchivalExternalTarget) SetVaultName(v string)`

SetVaultName sets VaultName field to given value.

### HasVaultName

`func (o *ArchivalExternalTarget) HasVaultName() bool`

HasVaultName returns a boolean if a field has been set.

### SetVaultNameNil

`func (o *ArchivalExternalTarget) SetVaultNameNil(b bool)`

 SetVaultNameNil sets the value for VaultName to be an explicit nil

### UnsetVaultName
`func (o *ArchivalExternalTarget) UnsetVaultName()`

UnsetVaultName ensures that no value is present for VaultName, not even an explicit nil
### GetVaultType

`func (o *ArchivalExternalTarget) GetVaultType() string`

GetVaultType returns the VaultType field if non-nil, zero value otherwise.

### GetVaultTypeOk

`func (o *ArchivalExternalTarget) GetVaultTypeOk() (*string, bool)`

GetVaultTypeOk returns a tuple with the VaultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultType

`func (o *ArchivalExternalTarget) SetVaultType(v string)`

SetVaultType sets VaultType field to given value.

### HasVaultType

`func (o *ArchivalExternalTarget) HasVaultType() bool`

HasVaultType returns a boolean if a field has been set.

### SetVaultTypeNil

`func (o *ArchivalExternalTarget) SetVaultTypeNil(b bool)`

 SetVaultTypeNil sets the value for VaultType to be an explicit nil

### UnsetVaultType
`func (o *ArchivalExternalTarget) UnsetVaultType()`

UnsetVaultType ensures that no value is present for VaultType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


