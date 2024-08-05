# CertificateBundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaChain** | Pointer to **[]string** | List of certs, starting at the root cert, in the chain of trust corresponding to the &#39;publicCert&#39;. | [optional] 
**PrivateKey** | **string** | Private Key corresponding to the &#39;publicCert&#39;. | 
**PublicCert** | **string** | The leaf certificate. | 

## Methods

### NewCertificateBundle

`func NewCertificateBundle(privateKey string, publicCert string, ) *CertificateBundle`

NewCertificateBundle instantiates a new CertificateBundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateBundleWithDefaults

`func NewCertificateBundleWithDefaults() *CertificateBundle`

NewCertificateBundleWithDefaults instantiates a new CertificateBundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaChain

`func (o *CertificateBundle) GetCaChain() []string`

GetCaChain returns the CaChain field if non-nil, zero value otherwise.

### GetCaChainOk

`func (o *CertificateBundle) GetCaChainOk() (*[]string, bool)`

GetCaChainOk returns a tuple with the CaChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaChain

`func (o *CertificateBundle) SetCaChain(v []string)`

SetCaChain sets CaChain field to given value.

### HasCaChain

`func (o *CertificateBundle) HasCaChain() bool`

HasCaChain returns a boolean if a field has been set.

### GetPrivateKey

`func (o *CertificateBundle) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CertificateBundle) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CertificateBundle) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetPublicCert

`func (o *CertificateBundle) GetPublicCert() string`

GetPublicCert returns the PublicCert field if non-nil, zero value otherwise.

### GetPublicCertOk

`func (o *CertificateBundle) GetPublicCertOk() (*string, bool)`

GetPublicCertOk returns a tuple with the PublicCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicCert

`func (o *CertificateBundle) SetPublicCert(v string)`

SetPublicCert sets PublicCert field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


