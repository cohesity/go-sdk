# KmipKmsConfigurationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalServerAddress** | Pointer to **[]string** | Additional KMS server IP address or FQDNs for fail over. | [optional] 
**CertificateExpiryDate** | Pointer to **NullableInt64** | Specifies expiry date of client certificate in msecs. | [optional] [readonly] 
**Port** | Pointer to **NullableInt32** | Port on which the KMS server is listening. | [optional] [default to 5696]
**ProtocolVersion** | Pointer to **NullableString** | KMIP protocol version used to communicate with the KMS. | [optional] 
**Server** | Pointer to **NullableString** | KMS server IP address or FQDN. | [optional] 

## Methods

### NewKmipKmsConfigurationResponse

`func NewKmipKmsConfigurationResponse() *KmipKmsConfigurationResponse`

NewKmipKmsConfigurationResponse instantiates a new KmipKmsConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmipKmsConfigurationResponseWithDefaults

`func NewKmipKmsConfigurationResponseWithDefaults() *KmipKmsConfigurationResponse`

NewKmipKmsConfigurationResponseWithDefaults instantiates a new KmipKmsConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalServerAddress

`func (o *KmipKmsConfigurationResponse) GetAdditionalServerAddress() []string`

GetAdditionalServerAddress returns the AdditionalServerAddress field if non-nil, zero value otherwise.

### GetAdditionalServerAddressOk

`func (o *KmipKmsConfigurationResponse) GetAdditionalServerAddressOk() (*[]string, bool)`

GetAdditionalServerAddressOk returns a tuple with the AdditionalServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalServerAddress

`func (o *KmipKmsConfigurationResponse) SetAdditionalServerAddress(v []string)`

SetAdditionalServerAddress sets AdditionalServerAddress field to given value.

### HasAdditionalServerAddress

`func (o *KmipKmsConfigurationResponse) HasAdditionalServerAddress() bool`

HasAdditionalServerAddress returns a boolean if a field has been set.

### GetCertificateExpiryDate

`func (o *KmipKmsConfigurationResponse) GetCertificateExpiryDate() int64`

GetCertificateExpiryDate returns the CertificateExpiryDate field if non-nil, zero value otherwise.

### GetCertificateExpiryDateOk

`func (o *KmipKmsConfigurationResponse) GetCertificateExpiryDateOk() (*int64, bool)`

GetCertificateExpiryDateOk returns a tuple with the CertificateExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateExpiryDate

`func (o *KmipKmsConfigurationResponse) SetCertificateExpiryDate(v int64)`

SetCertificateExpiryDate sets CertificateExpiryDate field to given value.

### HasCertificateExpiryDate

`func (o *KmipKmsConfigurationResponse) HasCertificateExpiryDate() bool`

HasCertificateExpiryDate returns a boolean if a field has been set.

### SetCertificateExpiryDateNil

`func (o *KmipKmsConfigurationResponse) SetCertificateExpiryDateNil(b bool)`

 SetCertificateExpiryDateNil sets the value for CertificateExpiryDate to be an explicit nil

### UnsetCertificateExpiryDate
`func (o *KmipKmsConfigurationResponse) UnsetCertificateExpiryDate()`

UnsetCertificateExpiryDate ensures that no value is present for CertificateExpiryDate, not even an explicit nil
### GetPort

`func (o *KmipKmsConfigurationResponse) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *KmipKmsConfigurationResponse) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *KmipKmsConfigurationResponse) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *KmipKmsConfigurationResponse) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *KmipKmsConfigurationResponse) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *KmipKmsConfigurationResponse) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetProtocolVersion

`func (o *KmipKmsConfigurationResponse) GetProtocolVersion() string`

GetProtocolVersion returns the ProtocolVersion field if non-nil, zero value otherwise.

### GetProtocolVersionOk

`func (o *KmipKmsConfigurationResponse) GetProtocolVersionOk() (*string, bool)`

GetProtocolVersionOk returns a tuple with the ProtocolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolVersion

`func (o *KmipKmsConfigurationResponse) SetProtocolVersion(v string)`

SetProtocolVersion sets ProtocolVersion field to given value.

### HasProtocolVersion

`func (o *KmipKmsConfigurationResponse) HasProtocolVersion() bool`

HasProtocolVersion returns a boolean if a field has been set.

### SetProtocolVersionNil

`func (o *KmipKmsConfigurationResponse) SetProtocolVersionNil(b bool)`

 SetProtocolVersionNil sets the value for ProtocolVersion to be an explicit nil

### UnsetProtocolVersion
`func (o *KmipKmsConfigurationResponse) UnsetProtocolVersion()`

UnsetProtocolVersion ensures that no value is present for ProtocolVersion, not even an explicit nil
### GetServer

`func (o *KmipKmsConfigurationResponse) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *KmipKmsConfigurationResponse) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *KmipKmsConfigurationResponse) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *KmipKmsConfigurationResponse) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *KmipKmsConfigurationResponse) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *KmipKmsConfigurationResponse) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


