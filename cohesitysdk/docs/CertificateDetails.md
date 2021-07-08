# CertificateDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertFileName** | Pointer to **NullableString** | Specifies the filename of the certificate. This is unique to each certificate generated. | [optional] 
**ExpiryDate** | Pointer to **NullableString** | Specifies the date in epoch till when the certificate is valid. | [optional] 
**HostIps** | Pointer to **[]string** | Each certificate can be deployed to multiple hosts. List of all hosts is returned after deployment. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of the host such as &#39;kSapHana&#39;, &#39;kSapOracle&#39;, etc. Specifies the host type of host for generating and deploying a Certificate. &#39;kOther&#39; indicates it is any of the other hosts. &#39;kSapOracle&#39; indicates it is a SAP Oracle host. &#39;kSapHana&#39; indicates it is a SAP HANA host. | [optional] 

## Methods

### NewCertificateDetails

`func NewCertificateDetails() *CertificateDetails`

NewCertificateDetails instantiates a new CertificateDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateDetailsWithDefaults

`func NewCertificateDetailsWithDefaults() *CertificateDetails`

NewCertificateDetailsWithDefaults instantiates a new CertificateDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertFileName

`func (o *CertificateDetails) GetCertFileName() string`

GetCertFileName returns the CertFileName field if non-nil, zero value otherwise.

### GetCertFileNameOk

`func (o *CertificateDetails) GetCertFileNameOk() (*string, bool)`

GetCertFileNameOk returns a tuple with the CertFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertFileName

`func (o *CertificateDetails) SetCertFileName(v string)`

SetCertFileName sets CertFileName field to given value.

### HasCertFileName

`func (o *CertificateDetails) HasCertFileName() bool`

HasCertFileName returns a boolean if a field has been set.

### SetCertFileNameNil

`func (o *CertificateDetails) SetCertFileNameNil(b bool)`

 SetCertFileNameNil sets the value for CertFileName to be an explicit nil

### UnsetCertFileName
`func (o *CertificateDetails) UnsetCertFileName()`

UnsetCertFileName ensures that no value is present for CertFileName, not even an explicit nil
### GetExpiryDate

`func (o *CertificateDetails) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *CertificateDetails) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *CertificateDetails) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *CertificateDetails) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDateNil

`func (o *CertificateDetails) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *CertificateDetails) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
### GetHostIps

`func (o *CertificateDetails) GetHostIps() []string`

GetHostIps returns the HostIps field if non-nil, zero value otherwise.

### GetHostIpsOk

`func (o *CertificateDetails) GetHostIpsOk() (*[]string, bool)`

GetHostIpsOk returns a tuple with the HostIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostIps

`func (o *CertificateDetails) SetHostIps(v []string)`

SetHostIps sets HostIps field to given value.

### HasHostIps

`func (o *CertificateDetails) HasHostIps() bool`

HasHostIps returns a boolean if a field has been set.

### SetHostIpsNil

`func (o *CertificateDetails) SetHostIpsNil(b bool)`

 SetHostIpsNil sets the value for HostIps to be an explicit nil

### UnsetHostIps
`func (o *CertificateDetails) UnsetHostIps()`

UnsetHostIps ensures that no value is present for HostIps, not even an explicit nil
### GetType

`func (o *CertificateDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CertificateDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CertificateDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CertificateDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CertificateDetails) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CertificateDetails) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


