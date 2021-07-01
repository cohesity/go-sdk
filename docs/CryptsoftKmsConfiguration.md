# CryptsoftKmsConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaCertificate** | Pointer to **NullableString** | Specifies the CA certificate in PEM format. | [optional] 
**ClientCertificate** | Pointer to **NullableString** | Specifies the client certificate. It is in PEM format. | [optional] 
**ClientKey** | Pointer to **NullableString** | Specifies the client private key. | [optional] 
**KmipProtocolVersion** | Pointer to **NullableString** | Specifies protocol version used to communicate with the KMS. | [optional] 
**ServerIp** | Pointer to **NullableString** | Specifies the KMS IP address. | [optional] 
**ServerPort** | Pointer to **NullableInt32** | Specifies port on which the server is listening. Default port is 5696. | [optional] 

## Methods

### NewCryptsoftKmsConfiguration

`func NewCryptsoftKmsConfiguration() *CryptsoftKmsConfiguration`

NewCryptsoftKmsConfiguration instantiates a new CryptsoftKmsConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptsoftKmsConfigurationWithDefaults

`func NewCryptsoftKmsConfigurationWithDefaults() *CryptsoftKmsConfiguration`

NewCryptsoftKmsConfigurationWithDefaults instantiates a new CryptsoftKmsConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaCertificate

`func (o *CryptsoftKmsConfiguration) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *CryptsoftKmsConfiguration) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *CryptsoftKmsConfiguration) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *CryptsoftKmsConfiguration) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.

### SetCaCertificateNil

`func (o *CryptsoftKmsConfiguration) SetCaCertificateNil(b bool)`

 SetCaCertificateNil sets the value for CaCertificate to be an explicit nil

### UnsetCaCertificate
`func (o *CryptsoftKmsConfiguration) UnsetCaCertificate()`

UnsetCaCertificate ensures that no value is present for CaCertificate, not even an explicit nil
### GetClientCertificate

`func (o *CryptsoftKmsConfiguration) GetClientCertificate() string`

GetClientCertificate returns the ClientCertificate field if non-nil, zero value otherwise.

### GetClientCertificateOk

`func (o *CryptsoftKmsConfiguration) GetClientCertificateOk() (*string, bool)`

GetClientCertificateOk returns a tuple with the ClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificate

`func (o *CryptsoftKmsConfiguration) SetClientCertificate(v string)`

SetClientCertificate sets ClientCertificate field to given value.

### HasClientCertificate

`func (o *CryptsoftKmsConfiguration) HasClientCertificate() bool`

HasClientCertificate returns a boolean if a field has been set.

### SetClientCertificateNil

`func (o *CryptsoftKmsConfiguration) SetClientCertificateNil(b bool)`

 SetClientCertificateNil sets the value for ClientCertificate to be an explicit nil

### UnsetClientCertificate
`func (o *CryptsoftKmsConfiguration) UnsetClientCertificate()`

UnsetClientCertificate ensures that no value is present for ClientCertificate, not even an explicit nil
### GetClientKey

`func (o *CryptsoftKmsConfiguration) GetClientKey() string`

GetClientKey returns the ClientKey field if non-nil, zero value otherwise.

### GetClientKeyOk

`func (o *CryptsoftKmsConfiguration) GetClientKeyOk() (*string, bool)`

GetClientKeyOk returns a tuple with the ClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKey

`func (o *CryptsoftKmsConfiguration) SetClientKey(v string)`

SetClientKey sets ClientKey field to given value.

### HasClientKey

`func (o *CryptsoftKmsConfiguration) HasClientKey() bool`

HasClientKey returns a boolean if a field has been set.

### SetClientKeyNil

`func (o *CryptsoftKmsConfiguration) SetClientKeyNil(b bool)`

 SetClientKeyNil sets the value for ClientKey to be an explicit nil

### UnsetClientKey
`func (o *CryptsoftKmsConfiguration) UnsetClientKey()`

UnsetClientKey ensures that no value is present for ClientKey, not even an explicit nil
### GetKmipProtocolVersion

`func (o *CryptsoftKmsConfiguration) GetKmipProtocolVersion() string`

GetKmipProtocolVersion returns the KmipProtocolVersion field if non-nil, zero value otherwise.

### GetKmipProtocolVersionOk

`func (o *CryptsoftKmsConfiguration) GetKmipProtocolVersionOk() (*string, bool)`

GetKmipProtocolVersionOk returns a tuple with the KmipProtocolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmipProtocolVersion

`func (o *CryptsoftKmsConfiguration) SetKmipProtocolVersion(v string)`

SetKmipProtocolVersion sets KmipProtocolVersion field to given value.

### HasKmipProtocolVersion

`func (o *CryptsoftKmsConfiguration) HasKmipProtocolVersion() bool`

HasKmipProtocolVersion returns a boolean if a field has been set.

### SetKmipProtocolVersionNil

`func (o *CryptsoftKmsConfiguration) SetKmipProtocolVersionNil(b bool)`

 SetKmipProtocolVersionNil sets the value for KmipProtocolVersion to be an explicit nil

### UnsetKmipProtocolVersion
`func (o *CryptsoftKmsConfiguration) UnsetKmipProtocolVersion()`

UnsetKmipProtocolVersion ensures that no value is present for KmipProtocolVersion, not even an explicit nil
### GetServerIp

`func (o *CryptsoftKmsConfiguration) GetServerIp() string`

GetServerIp returns the ServerIp field if non-nil, zero value otherwise.

### GetServerIpOk

`func (o *CryptsoftKmsConfiguration) GetServerIpOk() (*string, bool)`

GetServerIpOk returns a tuple with the ServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIp

`func (o *CryptsoftKmsConfiguration) SetServerIp(v string)`

SetServerIp sets ServerIp field to given value.

### HasServerIp

`func (o *CryptsoftKmsConfiguration) HasServerIp() bool`

HasServerIp returns a boolean if a field has been set.

### SetServerIpNil

`func (o *CryptsoftKmsConfiguration) SetServerIpNil(b bool)`

 SetServerIpNil sets the value for ServerIp to be an explicit nil

### UnsetServerIp
`func (o *CryptsoftKmsConfiguration) UnsetServerIp()`

UnsetServerIp ensures that no value is present for ServerIp, not even an explicit nil
### GetServerPort

`func (o *CryptsoftKmsConfiguration) GetServerPort() int32`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *CryptsoftKmsConfiguration) GetServerPortOk() (*int32, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *CryptsoftKmsConfiguration) SetServerPort(v int32)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *CryptsoftKmsConfiguration) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### SetServerPortNil

`func (o *CryptsoftKmsConfiguration) SetServerPortNil(b bool)`

 SetServerPortNil sets the value for ServerPort to be an explicit nil

### UnsetServerPort
`func (o *CryptsoftKmsConfiguration) UnsetServerPort()`

UnsetServerPort ensures that no value is present for ServerPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


