# CertificateRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | **NullableString** | Specifies the locality attribute, which is part of the distinguished name definition. It is used to identify the city where the company is located or the Cluster is installed. | 
**CommonName** | **NullableString** | Specifies the common name attribute, which is part of the distinguished name definition. Common name is used to specify a context for the certificate, for example, the name of the Cluster to which the certificate is to be assigned. Default value is the name of the Cluster. | 
**CountryCode** | **NullableString** | Specifies the country attribute, which is part of the distinguished name definition. It is used to identify the country where the state is located. It is specified as two letter code defined by the ISO standard. | 
**Duration** | Pointer to **NullableString** | Specifies duration of the certificate expiry in(hours). | [optional] 
**EmailAddress** | Pointer to **NullableString** | Specifies an alternative subject name component to be included in the certificate. Format is a standard e-mail address, for example joe@company.com. | [optional] 
**KeyType** | Pointer to **NullableString** | Specifies the algorithm to be used to generate the key pair. RSA is the default value. | [optional] [default to "RSA_4096"]
**Organization** | **NullableString** | Specifies the organization attribute, which is part of the distinguished name definition. It is used to specify the name of the company. | 
**OrganizationUnit** | **NullableString** | Specifies the organization unit attribute, which is part of the distinguished name definition. It is used to identify the specific department or business unit in the company that is owning the Cluster. | 
**SanList** | Pointer to **[]string** | Specifies an alternative subject name component to be included in the certificate. It is used to identify the ways the Cluster will be accessed. It is given as a comma separated list of FQDNs. The default value is the Cluster&#39;s VIP hostname. | [optional] 
**State** | **NullableString** | Specifies the state attribute, which is part of the distinguished name definition. It is used to identify the state where the city is located. | 
**TenantId** | Pointer to **NullableString** | Specifies the tenant id | [optional] 

## Methods

### NewCertificateRequestParams

`func NewCertificateRequestParams(city NullableString, commonName NullableString, countryCode NullableString, organization NullableString, organizationUnit NullableString, state NullableString, ) *CertificateRequestParams`

NewCertificateRequestParams instantiates a new CertificateRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateRequestParamsWithDefaults

`func NewCertificateRequestParamsWithDefaults() *CertificateRequestParams`

NewCertificateRequestParamsWithDefaults instantiates a new CertificateRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *CertificateRequestParams) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CertificateRequestParams) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CertificateRequestParams) SetCity(v string)`

SetCity sets City field to given value.


### SetCityNil

`func (o *CertificateRequestParams) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *CertificateRequestParams) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCommonName

`func (o *CertificateRequestParams) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *CertificateRequestParams) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *CertificateRequestParams) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.


### SetCommonNameNil

`func (o *CertificateRequestParams) SetCommonNameNil(b bool)`

 SetCommonNameNil sets the value for CommonName to be an explicit nil

### UnsetCommonName
`func (o *CertificateRequestParams) UnsetCommonName()`

UnsetCommonName ensures that no value is present for CommonName, not even an explicit nil
### GetCountryCode

`func (o *CertificateRequestParams) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *CertificateRequestParams) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *CertificateRequestParams) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### SetCountryCodeNil

`func (o *CertificateRequestParams) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *CertificateRequestParams) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetDuration

`func (o *CertificateRequestParams) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CertificateRequestParams) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CertificateRequestParams) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *CertificateRequestParams) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *CertificateRequestParams) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *CertificateRequestParams) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetEmailAddress

`func (o *CertificateRequestParams) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *CertificateRequestParams) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *CertificateRequestParams) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *CertificateRequestParams) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *CertificateRequestParams) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *CertificateRequestParams) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetKeyType

`func (o *CertificateRequestParams) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *CertificateRequestParams) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *CertificateRequestParams) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *CertificateRequestParams) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### SetKeyTypeNil

`func (o *CertificateRequestParams) SetKeyTypeNil(b bool)`

 SetKeyTypeNil sets the value for KeyType to be an explicit nil

### UnsetKeyType
`func (o *CertificateRequestParams) UnsetKeyType()`

UnsetKeyType ensures that no value is present for KeyType, not even an explicit nil
### GetOrganization

`func (o *CertificateRequestParams) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CertificateRequestParams) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CertificateRequestParams) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### SetOrganizationNil

`func (o *CertificateRequestParams) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *CertificateRequestParams) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetOrganizationUnit

`func (o *CertificateRequestParams) GetOrganizationUnit() string`

GetOrganizationUnit returns the OrganizationUnit field if non-nil, zero value otherwise.

### GetOrganizationUnitOk

`func (o *CertificateRequestParams) GetOrganizationUnitOk() (*string, bool)`

GetOrganizationUnitOk returns a tuple with the OrganizationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationUnit

`func (o *CertificateRequestParams) SetOrganizationUnit(v string)`

SetOrganizationUnit sets OrganizationUnit field to given value.


### SetOrganizationUnitNil

`func (o *CertificateRequestParams) SetOrganizationUnitNil(b bool)`

 SetOrganizationUnitNil sets the value for OrganizationUnit to be an explicit nil

### UnsetOrganizationUnit
`func (o *CertificateRequestParams) UnsetOrganizationUnit()`

UnsetOrganizationUnit ensures that no value is present for OrganizationUnit, not even an explicit nil
### GetSanList

`func (o *CertificateRequestParams) GetSanList() []string`

GetSanList returns the SanList field if non-nil, zero value otherwise.

### GetSanListOk

`func (o *CertificateRequestParams) GetSanListOk() (*[]string, bool)`

GetSanListOk returns a tuple with the SanList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanList

`func (o *CertificateRequestParams) SetSanList(v []string)`

SetSanList sets SanList field to given value.

### HasSanList

`func (o *CertificateRequestParams) HasSanList() bool`

HasSanList returns a boolean if a field has been set.

### SetSanListNil

`func (o *CertificateRequestParams) SetSanListNil(b bool)`

 SetSanListNil sets the value for SanList to be an explicit nil

### UnsetSanList
`func (o *CertificateRequestParams) UnsetSanList()`

UnsetSanList ensures that no value is present for SanList, not even an explicit nil
### GetState

`func (o *CertificateRequestParams) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CertificateRequestParams) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CertificateRequestParams) SetState(v string)`

SetState sets State field to given value.


### SetStateNil

`func (o *CertificateRequestParams) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *CertificateRequestParams) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetTenantId

`func (o *CertificateRequestParams) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CertificateRequestParams) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CertificateRequestParams) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CertificateRequestParams) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CertificateRequestParams) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CertificateRequestParams) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


