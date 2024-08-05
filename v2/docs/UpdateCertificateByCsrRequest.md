# UpdateCertificateByCsrRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | **NullableString** | Specifies the certificate to be imported. | 
**CsrId** | **NullableString** | Specifies the id of the csr corresponding to the certificate. | 

## Methods

### NewUpdateCertificateByCsrRequest

`func NewUpdateCertificateByCsrRequest(certificate NullableString, csrId NullableString, ) *UpdateCertificateByCsrRequest`

NewUpdateCertificateByCsrRequest instantiates a new UpdateCertificateByCsrRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCertificateByCsrRequestWithDefaults

`func NewUpdateCertificateByCsrRequestWithDefaults() *UpdateCertificateByCsrRequest`

NewUpdateCertificateByCsrRequestWithDefaults instantiates a new UpdateCertificateByCsrRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *UpdateCertificateByCsrRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *UpdateCertificateByCsrRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *UpdateCertificateByCsrRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### SetCertificateNil

`func (o *UpdateCertificateByCsrRequest) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *UpdateCertificateByCsrRequest) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetCsrId

`func (o *UpdateCertificateByCsrRequest) GetCsrId() string`

GetCsrId returns the CsrId field if non-nil, zero value otherwise.

### GetCsrIdOk

`func (o *UpdateCertificateByCsrRequest) GetCsrIdOk() (*string, bool)`

GetCsrIdOk returns a tuple with the CsrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrId

`func (o *UpdateCertificateByCsrRequest) SetCsrId(v string)`

SetCsrId sets CsrId field to given value.


### SetCsrIdNil

`func (o *UpdateCertificateByCsrRequest) SetCsrIdNil(b bool)`

 SetCsrIdNil sets the value for CsrId to be an explicit nil

### UnsetCsrId
`func (o *UpdateCertificateByCsrRequest) UnsetCsrId()`

UnsetCsrId ensures that no value is present for CsrId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


