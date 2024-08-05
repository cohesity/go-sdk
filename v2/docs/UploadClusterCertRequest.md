# UploadClusterCertRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificates** | Pointer to [**[]Certificate**](Certificate.md) | Array of certificates | [optional] 

## Methods

### NewUploadClusterCertRequest

`func NewUploadClusterCertRequest() *UploadClusterCertRequest`

NewUploadClusterCertRequest instantiates a new UploadClusterCertRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadClusterCertRequestWithDefaults

`func NewUploadClusterCertRequestWithDefaults() *UploadClusterCertRequest`

NewUploadClusterCertRequestWithDefaults instantiates a new UploadClusterCertRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificates

`func (o *UploadClusterCertRequest) GetCertificates() []Certificate`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *UploadClusterCertRequest) GetCertificatesOk() (*[]Certificate, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *UploadClusterCertRequest) SetCertificates(v []Certificate)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *UploadClusterCertRequest) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


