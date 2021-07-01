# VaultEncryptionKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | Pointer to **NullableString** | Specifies the name of the source Cohesity Cluster that archived the data on the Vault. | [optional] 
**EncryptionKeyData** | Pointer to **NullableString** | Specifies the encryption key data corresponding to the specified keyUid. It contains a Key Encryption Key (KEK) or a Encrypted Data Encryption Key (eDEK). | [optional] 
**KeyUid** | Pointer to [**NullableUniversalId**](UniversalId.md) | Specifies the universal id of the Data Encryption Key. | [optional] 
**VaultId** | Pointer to **NullableInt64** | Specifies the id of the Vault whose data is encrypted by this key. | [optional] 
**VaultName** | Pointer to **NullableString** | Specifies the name of the Vault whose data is encrypted by this key. | [optional] 

## Methods

### NewVaultEncryptionKey

`func NewVaultEncryptionKey() *VaultEncryptionKey`

NewVaultEncryptionKey instantiates a new VaultEncryptionKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultEncryptionKeyWithDefaults

`func NewVaultEncryptionKeyWithDefaults() *VaultEncryptionKey`

NewVaultEncryptionKeyWithDefaults instantiates a new VaultEncryptionKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *VaultEncryptionKey) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *VaultEncryptionKey) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *VaultEncryptionKey) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *VaultEncryptionKey) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### SetClusterNameNil

`func (o *VaultEncryptionKey) SetClusterNameNil(b bool)`

 SetClusterNameNil sets the value for ClusterName to be an explicit nil

### UnsetClusterName
`func (o *VaultEncryptionKey) UnsetClusterName()`

UnsetClusterName ensures that no value is present for ClusterName, not even an explicit nil
### GetEncryptionKeyData

`func (o *VaultEncryptionKey) GetEncryptionKeyData() string`

GetEncryptionKeyData returns the EncryptionKeyData field if non-nil, zero value otherwise.

### GetEncryptionKeyDataOk

`func (o *VaultEncryptionKey) GetEncryptionKeyDataOk() (*string, bool)`

GetEncryptionKeyDataOk returns a tuple with the EncryptionKeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeyData

`func (o *VaultEncryptionKey) SetEncryptionKeyData(v string)`

SetEncryptionKeyData sets EncryptionKeyData field to given value.

### HasEncryptionKeyData

`func (o *VaultEncryptionKey) HasEncryptionKeyData() bool`

HasEncryptionKeyData returns a boolean if a field has been set.

### SetEncryptionKeyDataNil

`func (o *VaultEncryptionKey) SetEncryptionKeyDataNil(b bool)`

 SetEncryptionKeyDataNil sets the value for EncryptionKeyData to be an explicit nil

### UnsetEncryptionKeyData
`func (o *VaultEncryptionKey) UnsetEncryptionKeyData()`

UnsetEncryptionKeyData ensures that no value is present for EncryptionKeyData, not even an explicit nil
### GetKeyUid

`func (o *VaultEncryptionKey) GetKeyUid() UniversalId`

GetKeyUid returns the KeyUid field if non-nil, zero value otherwise.

### GetKeyUidOk

`func (o *VaultEncryptionKey) GetKeyUidOk() (*UniversalId, bool)`

GetKeyUidOk returns a tuple with the KeyUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUid

`func (o *VaultEncryptionKey) SetKeyUid(v UniversalId)`

SetKeyUid sets KeyUid field to given value.

### HasKeyUid

`func (o *VaultEncryptionKey) HasKeyUid() bool`

HasKeyUid returns a boolean if a field has been set.

### SetKeyUidNil

`func (o *VaultEncryptionKey) SetKeyUidNil(b bool)`

 SetKeyUidNil sets the value for KeyUid to be an explicit nil

### UnsetKeyUid
`func (o *VaultEncryptionKey) UnsetKeyUid()`

UnsetKeyUid ensures that no value is present for KeyUid, not even an explicit nil
### GetVaultId

`func (o *VaultEncryptionKey) GetVaultId() int64`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *VaultEncryptionKey) GetVaultIdOk() (*int64, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *VaultEncryptionKey) SetVaultId(v int64)`

SetVaultId sets VaultId field to given value.

### HasVaultId

`func (o *VaultEncryptionKey) HasVaultId() bool`

HasVaultId returns a boolean if a field has been set.

### SetVaultIdNil

`func (o *VaultEncryptionKey) SetVaultIdNil(b bool)`

 SetVaultIdNil sets the value for VaultId to be an explicit nil

### UnsetVaultId
`func (o *VaultEncryptionKey) UnsetVaultId()`

UnsetVaultId ensures that no value is present for VaultId, not even an explicit nil
### GetVaultName

`func (o *VaultEncryptionKey) GetVaultName() string`

GetVaultName returns the VaultName field if non-nil, zero value otherwise.

### GetVaultNameOk

`func (o *VaultEncryptionKey) GetVaultNameOk() (*string, bool)`

GetVaultNameOk returns a tuple with the VaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultName

`func (o *VaultEncryptionKey) SetVaultName(v string)`

SetVaultName sets VaultName field to given value.

### HasVaultName

`func (o *VaultEncryptionKey) HasVaultName() bool`

HasVaultName returns a boolean if a field has been set.

### SetVaultNameNil

`func (o *VaultEncryptionKey) SetVaultNameNil(b bool)`

 SetVaultNameNil sets the value for VaultName to be an explicit nil

### UnsetVaultName
`func (o *VaultEncryptionKey) UnsetVaultName()`

UnsetVaultName ensures that no value is present for VaultName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


