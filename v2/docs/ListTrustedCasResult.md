# ListTrustedCasResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificates** | Pointer to [**[]TrustedCa**](TrustedCa.md) | Array of trusted certificates. Specifies the list of certificates returned in this response. List is not sorted. | [optional] 

## Methods

### NewListTrustedCasResult

`func NewListTrustedCasResult() *ListTrustedCasResult`

NewListTrustedCasResult instantiates a new ListTrustedCasResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTrustedCasResultWithDefaults

`func NewListTrustedCasResultWithDefaults() *ListTrustedCasResult`

NewListTrustedCasResultWithDefaults instantiates a new ListTrustedCasResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificates

`func (o *ListTrustedCasResult) GetCertificates() []TrustedCa`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *ListTrustedCasResult) GetCertificatesOk() (*[]TrustedCa, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *ListTrustedCasResult) SetCertificates(v []TrustedCa)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *ListTrustedCasResult) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### SetCertificatesNil

`func (o *ListTrustedCasResult) SetCertificatesNil(b bool)`

 SetCertificatesNil sets the value for Certificates to be an explicit nil

### UnsetCertificates
`func (o *ListTrustedCasResult) UnsetCertificates()`

UnsetCertificates ensures that no value is present for Certificates, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


