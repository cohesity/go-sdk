# KmipKmsConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalServerAddress** | Pointer to **[]string** | Additional KMS server IP address or FQDNs for fail over. | [optional] 
**CaCertificate** | **string** | CA certificate. | 
**CertificateExpiryDate** | Pointer to **NullableInt64** | Specifies expiry date of client certificate in msecs. | [optional] [readonly] 
**ClientCertificate** | **string** | Client certificate. | 
**ClientKey** | **string** | Client key. | 
**Port** | Pointer to **NullableInt32** | Port on which the KMS server is listening. | [optional] [default to 5696]
**ProtocolVersion** | **string** | KMIP protocol version used to communicate with the KMS. | 
**Server** | **string** | KMS server IP address or FQDN. | 

## Methods

### NewKmipKmsConfiguration

`func NewKmipKmsConfiguration(caCertificate string, clientCertificate string, clientKey string, protocolVersion string, server string, ) *KmipKmsConfiguration`

NewKmipKmsConfiguration instantiates a new KmipKmsConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmipKmsConfigurationWithDefaults

`func NewKmipKmsConfigurationWithDefaults() *KmipKmsConfiguration`

NewKmipKmsConfigurationWithDefaults instantiates a new KmipKmsConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalServerAddress

`func (o *KmipKmsConfiguration) GetAdditionalServerAddress() []string`

GetAdditionalServerAddress returns the AdditionalServerAddress field if non-nil, zero value otherwise.

### GetAdditionalServerAddressOk

`func (o *KmipKmsConfiguration) GetAdditionalServerAddressOk() (*[]string, bool)`

GetAdditionalServerAddressOk returns a tuple with the AdditionalServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalServerAddress

`func (o *KmipKmsConfiguration) SetAdditionalServerAddress(v []string)`

SetAdditionalServerAddress sets AdditionalServerAddress field to given value.

### HasAdditionalServerAddress

`func (o *KmipKmsConfiguration) HasAdditionalServerAddress() bool`

HasAdditionalServerAddress returns a boolean if a field has been set.

### GetCaCertificate

`func (o *KmipKmsConfiguration) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *KmipKmsConfiguration) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *KmipKmsConfiguration) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.


### GetCertificateExpiryDate

`func (o *KmipKmsConfiguration) GetCertificateExpiryDate() int64`

GetCertificateExpiryDate returns the CertificateExpiryDate field if non-nil, zero value otherwise.

### GetCertificateExpiryDateOk

`func (o *KmipKmsConfiguration) GetCertificateExpiryDateOk() (*int64, bool)`

GetCertificateExpiryDateOk returns a tuple with the CertificateExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateExpiryDate

`func (o *KmipKmsConfiguration) SetCertificateExpiryDate(v int64)`

SetCertificateExpiryDate sets CertificateExpiryDate field to given value.

### HasCertificateExpiryDate

`func (o *KmipKmsConfiguration) HasCertificateExpiryDate() bool`

HasCertificateExpiryDate returns a boolean if a field has been set.

### SetCertificateExpiryDateNil

`func (o *KmipKmsConfiguration) SetCertificateExpiryDateNil(b bool)`

 SetCertificateExpiryDateNil sets the value for CertificateExpiryDate to be an explicit nil

### UnsetCertificateExpiryDate
`func (o *KmipKmsConfiguration) UnsetCertificateExpiryDate()`

UnsetCertificateExpiryDate ensures that no value is present for CertificateExpiryDate, not even an explicit nil
### GetClientCertificate

`func (o *KmipKmsConfiguration) GetClientCertificate() string`

GetClientCertificate returns the ClientCertificate field if non-nil, zero value otherwise.

### GetClientCertificateOk

`func (o *KmipKmsConfiguration) GetClientCertificateOk() (*string, bool)`

GetClientCertificateOk returns a tuple with the ClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificate

`func (o *KmipKmsConfiguration) SetClientCertificate(v string)`

SetClientCertificate sets ClientCertificate field to given value.


### GetClientKey

`func (o *KmipKmsConfiguration) GetClientKey() string`

GetClientKey returns the ClientKey field if non-nil, zero value otherwise.

### GetClientKeyOk

`func (o *KmipKmsConfiguration) GetClientKeyOk() (*string, bool)`

GetClientKeyOk returns a tuple with the ClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKey

`func (o *KmipKmsConfiguration) SetClientKey(v string)`

SetClientKey sets ClientKey field to given value.


### GetPort

`func (o *KmipKmsConfiguration) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *KmipKmsConfiguration) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *KmipKmsConfiguration) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *KmipKmsConfiguration) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *KmipKmsConfiguration) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *KmipKmsConfiguration) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetProtocolVersion

`func (o *KmipKmsConfiguration) GetProtocolVersion() string`

GetProtocolVersion returns the ProtocolVersion field if non-nil, zero value otherwise.

### GetProtocolVersionOk

`func (o *KmipKmsConfiguration) GetProtocolVersionOk() (*string, bool)`

GetProtocolVersionOk returns a tuple with the ProtocolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolVersion

`func (o *KmipKmsConfiguration) SetProtocolVersion(v string)`

SetProtocolVersion sets ProtocolVersion field to given value.


### GetServer

`func (o *KmipKmsConfiguration) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *KmipKmsConfiguration) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *KmipKmsConfiguration) SetServer(v string)`

SetServer sets Server field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


