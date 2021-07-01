# ListCertResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateList** | Pointer to [**[]CertificateDetails**](CertificateDetails.md) |  | [optional] 

## Methods

### NewListCertResponse

`func NewListCertResponse() *ListCertResponse`

NewListCertResponse instantiates a new ListCertResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCertResponseWithDefaults

`func NewListCertResponseWithDefaults() *ListCertResponse`

NewListCertResponseWithDefaults instantiates a new ListCertResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateList

`func (o *ListCertResponse) GetCertificateList() []CertificateDetails`

GetCertificateList returns the CertificateList field if non-nil, zero value otherwise.

### GetCertificateListOk

`func (o *ListCertResponse) GetCertificateListOk() (*[]CertificateDetails, bool)`

GetCertificateListOk returns a tuple with the CertificateList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateList

`func (o *ListCertResponse) SetCertificateList(v []CertificateDetails)`

SetCertificateList sets CertificateList field to given value.

### HasCertificateList

`func (o *ListCertResponse) HasCertificateList() bool`

HasCertificateList returns a boolean if a field has been set.

### SetCertificateListNil

`func (o *ListCertResponse) SetCertificateListNil(b bool)`

 SetCertificateListNil sets the value for CertificateList to be an explicit nil

### UnsetCertificateList
`func (o *ListCertResponse) UnsetCertificateList()`

UnsetCertificateList ensures that no value is present for CertificateList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


