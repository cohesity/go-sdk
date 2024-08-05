# GenerateCaCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommonName** | **string** | Common Name | 
**Country** | **string** | Country | 
**Expiry** | Pointer to **string** | Duration (in hours) of the CA certificate. | [optional] 
**Locality** | **string** | Locality | 
**Organization** | **string** | Organization | 
**OrganizationalUnit** | Pointer to **string** | Organizational Unit | [optional] 

## Methods

### NewGenerateCaCertificateRequest

`func NewGenerateCaCertificateRequest(commonName string, country string, locality string, organization string, ) *GenerateCaCertificateRequest`

NewGenerateCaCertificateRequest instantiates a new GenerateCaCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateCaCertificateRequestWithDefaults

`func NewGenerateCaCertificateRequestWithDefaults() *GenerateCaCertificateRequest`

NewGenerateCaCertificateRequestWithDefaults instantiates a new GenerateCaCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommonName

`func (o *GenerateCaCertificateRequest) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *GenerateCaCertificateRequest) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *GenerateCaCertificateRequest) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.


### GetCountry

`func (o *GenerateCaCertificateRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GenerateCaCertificateRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GenerateCaCertificateRequest) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetExpiry

`func (o *GenerateCaCertificateRequest) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *GenerateCaCertificateRequest) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *GenerateCaCertificateRequest) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *GenerateCaCertificateRequest) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetLocality

`func (o *GenerateCaCertificateRequest) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *GenerateCaCertificateRequest) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *GenerateCaCertificateRequest) SetLocality(v string)`

SetLocality sets Locality field to given value.


### GetOrganization

`func (o *GenerateCaCertificateRequest) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *GenerateCaCertificateRequest) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *GenerateCaCertificateRequest) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetOrganizationalUnit

`func (o *GenerateCaCertificateRequest) GetOrganizationalUnit() string`

GetOrganizationalUnit returns the OrganizationalUnit field if non-nil, zero value otherwise.

### GetOrganizationalUnitOk

`func (o *GenerateCaCertificateRequest) GetOrganizationalUnitOk() (*string, bool)`

GetOrganizationalUnitOk returns a tuple with the OrganizationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnit

`func (o *GenerateCaCertificateRequest) SetOrganizationalUnit(v string)`

SetOrganizationalUnit sets OrganizationalUnit field to given value.

### HasOrganizationalUnit

`func (o *GenerateCaCertificateRequest) HasOrganizationalUnit() bool`

HasOrganizationalUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


