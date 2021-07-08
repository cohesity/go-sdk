# SslCertificateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **NullableString** | Certificate is a SSL certificate used by Iris HTTPS webserver. | [optional] 
**LastUpdateTimeMsecs** | Pointer to **NullableInt64** | LastUpdateTimeMsecs is a time in milliseconds at which certificate was last updated. | [optional] 
**PrivateKey** | Pointer to **NullableString** | PrivateKey is a matching private key of the above certificate. | [optional] 

## Methods

### NewSslCertificateConfig

`func NewSslCertificateConfig() *SslCertificateConfig`

NewSslCertificateConfig instantiates a new SslCertificateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslCertificateConfigWithDefaults

`func NewSslCertificateConfigWithDefaults() *SslCertificateConfig`

NewSslCertificateConfigWithDefaults instantiates a new SslCertificateConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *SslCertificateConfig) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SslCertificateConfig) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SslCertificateConfig) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *SslCertificateConfig) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *SslCertificateConfig) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *SslCertificateConfig) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetLastUpdateTimeMsecs

`func (o *SslCertificateConfig) GetLastUpdateTimeMsecs() int64`

GetLastUpdateTimeMsecs returns the LastUpdateTimeMsecs field if non-nil, zero value otherwise.

### GetLastUpdateTimeMsecsOk

`func (o *SslCertificateConfig) GetLastUpdateTimeMsecsOk() (*int64, bool)`

GetLastUpdateTimeMsecsOk returns a tuple with the LastUpdateTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimeMsecs

`func (o *SslCertificateConfig) SetLastUpdateTimeMsecs(v int64)`

SetLastUpdateTimeMsecs sets LastUpdateTimeMsecs field to given value.

### HasLastUpdateTimeMsecs

`func (o *SslCertificateConfig) HasLastUpdateTimeMsecs() bool`

HasLastUpdateTimeMsecs returns a boolean if a field has been set.

### SetLastUpdateTimeMsecsNil

`func (o *SslCertificateConfig) SetLastUpdateTimeMsecsNil(b bool)`

 SetLastUpdateTimeMsecsNil sets the value for LastUpdateTimeMsecs to be an explicit nil

### UnsetLastUpdateTimeMsecs
`func (o *SslCertificateConfig) UnsetLastUpdateTimeMsecs()`

UnsetLastUpdateTimeMsecs ensures that no value is present for LastUpdateTimeMsecs, not even an explicit nil
### GetPrivateKey

`func (o *SslCertificateConfig) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *SslCertificateConfig) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *SslCertificateConfig) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *SslCertificateConfig) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *SslCertificateConfig) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *SslCertificateConfig) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


