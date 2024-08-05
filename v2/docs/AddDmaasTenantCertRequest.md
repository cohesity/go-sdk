# AddDmaasTenantCertRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | **NullableString** | Specifies the tenant certificate. | 
**ConnectorCaChain** | **NullableString** | Specifies the CA chain that is used to sign the connector certificate. | 
**Passphrase** | Pointer to **NullableString** | Specifies the passphrase used to encode the private key. | [optional] 
**PrivateKey** | **NullableString** | Specifies the tenant private key. | 
**TenantId** | **NullableString** | The id of the tenant. | 

## Methods

### NewAddDmaasTenantCertRequest

`func NewAddDmaasTenantCertRequest(certificate NullableString, connectorCaChain NullableString, privateKey NullableString, tenantId NullableString, ) *AddDmaasTenantCertRequest`

NewAddDmaasTenantCertRequest instantiates a new AddDmaasTenantCertRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDmaasTenantCertRequestWithDefaults

`func NewAddDmaasTenantCertRequestWithDefaults() *AddDmaasTenantCertRequest`

NewAddDmaasTenantCertRequestWithDefaults instantiates a new AddDmaasTenantCertRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *AddDmaasTenantCertRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *AddDmaasTenantCertRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *AddDmaasTenantCertRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### SetCertificateNil

`func (o *AddDmaasTenantCertRequest) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *AddDmaasTenantCertRequest) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetConnectorCaChain

`func (o *AddDmaasTenantCertRequest) GetConnectorCaChain() string`

GetConnectorCaChain returns the ConnectorCaChain field if non-nil, zero value otherwise.

### GetConnectorCaChainOk

`func (o *AddDmaasTenantCertRequest) GetConnectorCaChainOk() (*string, bool)`

GetConnectorCaChainOk returns a tuple with the ConnectorCaChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorCaChain

`func (o *AddDmaasTenantCertRequest) SetConnectorCaChain(v string)`

SetConnectorCaChain sets ConnectorCaChain field to given value.


### SetConnectorCaChainNil

`func (o *AddDmaasTenantCertRequest) SetConnectorCaChainNil(b bool)`

 SetConnectorCaChainNil sets the value for ConnectorCaChain to be an explicit nil

### UnsetConnectorCaChain
`func (o *AddDmaasTenantCertRequest) UnsetConnectorCaChain()`

UnsetConnectorCaChain ensures that no value is present for ConnectorCaChain, not even an explicit nil
### GetPassphrase

`func (o *AddDmaasTenantCertRequest) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *AddDmaasTenantCertRequest) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *AddDmaasTenantCertRequest) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *AddDmaasTenantCertRequest) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### SetPassphraseNil

`func (o *AddDmaasTenantCertRequest) SetPassphraseNil(b bool)`

 SetPassphraseNil sets the value for Passphrase to be an explicit nil

### UnsetPassphrase
`func (o *AddDmaasTenantCertRequest) UnsetPassphrase()`

UnsetPassphrase ensures that no value is present for Passphrase, not even an explicit nil
### GetPrivateKey

`func (o *AddDmaasTenantCertRequest) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *AddDmaasTenantCertRequest) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *AddDmaasTenantCertRequest) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### SetPrivateKeyNil

`func (o *AddDmaasTenantCertRequest) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *AddDmaasTenantCertRequest) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
### GetTenantId

`func (o *AddDmaasTenantCertRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AddDmaasTenantCertRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AddDmaasTenantCertRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### SetTenantIdNil

`func (o *AddDmaasTenantCertRequest) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *AddDmaasTenantCertRequest) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


