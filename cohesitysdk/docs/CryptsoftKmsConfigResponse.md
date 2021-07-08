# CryptsoftKmsConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientCertificateExpiryDate** | Pointer to **NullableInt64** | Specifies expiry date of client certificate. | [optional] 
**KmipProtocolVersion** | Pointer to **NullableString** | Specifies protocol version used to communicate with the KMS. | [optional] 
**ServerIp** | Pointer to **NullableString** | Specifies the KMS IP address. | [optional] 
**ServerPort** | Pointer to **NullableInt32** | Specifies port on which the server is listening. Default port is 5696. | [optional] 

## Methods

### NewCryptsoftKmsConfigResponse

`func NewCryptsoftKmsConfigResponse() *CryptsoftKmsConfigResponse`

NewCryptsoftKmsConfigResponse instantiates a new CryptsoftKmsConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptsoftKmsConfigResponseWithDefaults

`func NewCryptsoftKmsConfigResponseWithDefaults() *CryptsoftKmsConfigResponse`

NewCryptsoftKmsConfigResponseWithDefaults instantiates a new CryptsoftKmsConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientCertificateExpiryDate

`func (o *CryptsoftKmsConfigResponse) GetClientCertificateExpiryDate() int64`

GetClientCertificateExpiryDate returns the ClientCertificateExpiryDate field if non-nil, zero value otherwise.

### GetClientCertificateExpiryDateOk

`func (o *CryptsoftKmsConfigResponse) GetClientCertificateExpiryDateOk() (*int64, bool)`

GetClientCertificateExpiryDateOk returns a tuple with the ClientCertificateExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateExpiryDate

`func (o *CryptsoftKmsConfigResponse) SetClientCertificateExpiryDate(v int64)`

SetClientCertificateExpiryDate sets ClientCertificateExpiryDate field to given value.

### HasClientCertificateExpiryDate

`func (o *CryptsoftKmsConfigResponse) HasClientCertificateExpiryDate() bool`

HasClientCertificateExpiryDate returns a boolean if a field has been set.

### SetClientCertificateExpiryDateNil

`func (o *CryptsoftKmsConfigResponse) SetClientCertificateExpiryDateNil(b bool)`

 SetClientCertificateExpiryDateNil sets the value for ClientCertificateExpiryDate to be an explicit nil

### UnsetClientCertificateExpiryDate
`func (o *CryptsoftKmsConfigResponse) UnsetClientCertificateExpiryDate()`

UnsetClientCertificateExpiryDate ensures that no value is present for ClientCertificateExpiryDate, not even an explicit nil
### GetKmipProtocolVersion

`func (o *CryptsoftKmsConfigResponse) GetKmipProtocolVersion() string`

GetKmipProtocolVersion returns the KmipProtocolVersion field if non-nil, zero value otherwise.

### GetKmipProtocolVersionOk

`func (o *CryptsoftKmsConfigResponse) GetKmipProtocolVersionOk() (*string, bool)`

GetKmipProtocolVersionOk returns a tuple with the KmipProtocolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmipProtocolVersion

`func (o *CryptsoftKmsConfigResponse) SetKmipProtocolVersion(v string)`

SetKmipProtocolVersion sets KmipProtocolVersion field to given value.

### HasKmipProtocolVersion

`func (o *CryptsoftKmsConfigResponse) HasKmipProtocolVersion() bool`

HasKmipProtocolVersion returns a boolean if a field has been set.

### SetKmipProtocolVersionNil

`func (o *CryptsoftKmsConfigResponse) SetKmipProtocolVersionNil(b bool)`

 SetKmipProtocolVersionNil sets the value for KmipProtocolVersion to be an explicit nil

### UnsetKmipProtocolVersion
`func (o *CryptsoftKmsConfigResponse) UnsetKmipProtocolVersion()`

UnsetKmipProtocolVersion ensures that no value is present for KmipProtocolVersion, not even an explicit nil
### GetServerIp

`func (o *CryptsoftKmsConfigResponse) GetServerIp() string`

GetServerIp returns the ServerIp field if non-nil, zero value otherwise.

### GetServerIpOk

`func (o *CryptsoftKmsConfigResponse) GetServerIpOk() (*string, bool)`

GetServerIpOk returns a tuple with the ServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIp

`func (o *CryptsoftKmsConfigResponse) SetServerIp(v string)`

SetServerIp sets ServerIp field to given value.

### HasServerIp

`func (o *CryptsoftKmsConfigResponse) HasServerIp() bool`

HasServerIp returns a boolean if a field has been set.

### SetServerIpNil

`func (o *CryptsoftKmsConfigResponse) SetServerIpNil(b bool)`

 SetServerIpNil sets the value for ServerIp to be an explicit nil

### UnsetServerIp
`func (o *CryptsoftKmsConfigResponse) UnsetServerIp()`

UnsetServerIp ensures that no value is present for ServerIp, not even an explicit nil
### GetServerPort

`func (o *CryptsoftKmsConfigResponse) GetServerPort() int32`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *CryptsoftKmsConfigResponse) GetServerPortOk() (*int32, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *CryptsoftKmsConfigResponse) SetServerPort(v int32)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *CryptsoftKmsConfigResponse) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### SetServerPortNil

`func (o *CryptsoftKmsConfigResponse) SetServerPortNil(b bool)`

 SetServerPortNil sets the value for ServerPort to be an explicit nil

### UnsetServerPort
`func (o *CryptsoftKmsConfigResponse) UnsetServerPort()`

UnsetServerPort ensures that no value is present for ServerPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


