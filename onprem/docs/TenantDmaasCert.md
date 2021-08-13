# TenantDmaasCert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **NullableString** | The id of the tenant. | [optional] 
**Certificate** | Pointer to **NullableString** | Specifies the tenant certificate. | [optional] 
**PrivateKey** | Pointer to **NullableString** | Specifies the tenant private key. | [optional] 
**ConnectorCaChain** | Pointer to **NullableString** | Specifies the CA chain that is used to sign the connector certificate. | [optional] 
**Passphrase** | Pointer to **NullableString** | Specifies the passphrase used to encode the private key. | [optional] 

## Methods

### NewTenantDmaasCert

`func NewTenantDmaasCert() *TenantDmaasCert`

NewTenantDmaasCert instantiates a new TenantDmaasCert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantDmaasCertWithDefaults

`func NewTenantDmaasCertWithDefaults() *TenantDmaasCert`

NewTenantDmaasCertWithDefaults instantiates a new TenantDmaasCert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *TenantDmaasCert) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantDmaasCert) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantDmaasCert) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantDmaasCert) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TenantDmaasCert) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TenantDmaasCert) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetCertificate

`func (o *TenantDmaasCert) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *TenantDmaasCert) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *TenantDmaasCert) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *TenantDmaasCert) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *TenantDmaasCert) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *TenantDmaasCert) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetPrivateKey

`func (o *TenantDmaasCert) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *TenantDmaasCert) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *TenantDmaasCert) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *TenantDmaasCert) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *TenantDmaasCert) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *TenantDmaasCert) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
### GetConnectorCaChain

`func (o *TenantDmaasCert) GetConnectorCaChain() string`

GetConnectorCaChain returns the ConnectorCaChain field if non-nil, zero value otherwise.

### GetConnectorCaChainOk

`func (o *TenantDmaasCert) GetConnectorCaChainOk() (*string, bool)`

GetConnectorCaChainOk returns a tuple with the ConnectorCaChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorCaChain

`func (o *TenantDmaasCert) SetConnectorCaChain(v string)`

SetConnectorCaChain sets ConnectorCaChain field to given value.

### HasConnectorCaChain

`func (o *TenantDmaasCert) HasConnectorCaChain() bool`

HasConnectorCaChain returns a boolean if a field has been set.

### SetConnectorCaChainNil

`func (o *TenantDmaasCert) SetConnectorCaChainNil(b bool)`

 SetConnectorCaChainNil sets the value for ConnectorCaChain to be an explicit nil

### UnsetConnectorCaChain
`func (o *TenantDmaasCert) UnsetConnectorCaChain()`

UnsetConnectorCaChain ensures that no value is present for ConnectorCaChain, not even an explicit nil
### GetPassphrase

`func (o *TenantDmaasCert) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *TenantDmaasCert) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *TenantDmaasCert) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *TenantDmaasCert) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### SetPassphraseNil

`func (o *TenantDmaasCert) SetPassphraseNil(b bool)`

 SetPassphraseNil sets the value for Passphrase to be an explicit nil

### UnsetPassphrase
`func (o *TenantDmaasCert) UnsetPassphrase()`

UnsetPassphrase ensures that no value is present for Passphrase, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


