# GenerateNewCertificateRequest

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

### NewGenerateNewCertificateRequest

`func NewGenerateNewCertificateRequest(city NullableString, commonName NullableString, countryCode NullableString, organization NullableString, organizationUnit NullableString, state NullableString, ) *GenerateNewCertificateRequest`

NewGenerateNewCertificateRequest instantiates a new GenerateNewCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateNewCertificateRequestWithDefaults

`func NewGenerateNewCertificateRequestWithDefaults() *GenerateNewCertificateRequest`

NewGenerateNewCertificateRequestWithDefaults instantiates a new GenerateNewCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *GenerateNewCertificateRequest) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *GenerateNewCertificateRequest) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *GenerateNewCertificateRequest) SetCity(v string)`

SetCity sets City field to given value.


### SetCityNil

`func (o *GenerateNewCertificateRequest) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *GenerateNewCertificateRequest) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCommonName

`func (o *GenerateNewCertificateRequest) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *GenerateNewCertificateRequest) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *GenerateNewCertificateRequest) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.


### SetCommonNameNil

`func (o *GenerateNewCertificateRequest) SetCommonNameNil(b bool)`

 SetCommonNameNil sets the value for CommonName to be an explicit nil

### UnsetCommonName
`func (o *GenerateNewCertificateRequest) UnsetCommonName()`

UnsetCommonName ensures that no value is present for CommonName, not even an explicit nil
### GetCountryCode

`func (o *GenerateNewCertificateRequest) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *GenerateNewCertificateRequest) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *GenerateNewCertificateRequest) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### SetCountryCodeNil

`func (o *GenerateNewCertificateRequest) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *GenerateNewCertificateRequest) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetDuration

`func (o *GenerateNewCertificateRequest) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GenerateNewCertificateRequest) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GenerateNewCertificateRequest) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *GenerateNewCertificateRequest) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *GenerateNewCertificateRequest) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *GenerateNewCertificateRequest) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetEmailAddress

`func (o *GenerateNewCertificateRequest) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *GenerateNewCertificateRequest) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *GenerateNewCertificateRequest) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *GenerateNewCertificateRequest) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *GenerateNewCertificateRequest) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *GenerateNewCertificateRequest) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetKeyType

`func (o *GenerateNewCertificateRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *GenerateNewCertificateRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *GenerateNewCertificateRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *GenerateNewCertificateRequest) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### SetKeyTypeNil

`func (o *GenerateNewCertificateRequest) SetKeyTypeNil(b bool)`

 SetKeyTypeNil sets the value for KeyType to be an explicit nil

### UnsetKeyType
`func (o *GenerateNewCertificateRequest) UnsetKeyType()`

UnsetKeyType ensures that no value is present for KeyType, not even an explicit nil
### GetOrganization

`func (o *GenerateNewCertificateRequest) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *GenerateNewCertificateRequest) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *GenerateNewCertificateRequest) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### SetOrganizationNil

`func (o *GenerateNewCertificateRequest) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *GenerateNewCertificateRequest) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetOrganizationUnit

`func (o *GenerateNewCertificateRequest) GetOrganizationUnit() string`

GetOrganizationUnit returns the OrganizationUnit field if non-nil, zero value otherwise.

### GetOrganizationUnitOk

`func (o *GenerateNewCertificateRequest) GetOrganizationUnitOk() (*string, bool)`

GetOrganizationUnitOk returns a tuple with the OrganizationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationUnit

`func (o *GenerateNewCertificateRequest) SetOrganizationUnit(v string)`

SetOrganizationUnit sets OrganizationUnit field to given value.


### SetOrganizationUnitNil

`func (o *GenerateNewCertificateRequest) SetOrganizationUnitNil(b bool)`

 SetOrganizationUnitNil sets the value for OrganizationUnit to be an explicit nil

### UnsetOrganizationUnit
`func (o *GenerateNewCertificateRequest) UnsetOrganizationUnit()`

UnsetOrganizationUnit ensures that no value is present for OrganizationUnit, not even an explicit nil
### GetSanList

`func (o *GenerateNewCertificateRequest) GetSanList() []string`

GetSanList returns the SanList field if non-nil, zero value otherwise.

### GetSanListOk

`func (o *GenerateNewCertificateRequest) GetSanListOk() (*[]string, bool)`

GetSanListOk returns a tuple with the SanList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanList

`func (o *GenerateNewCertificateRequest) SetSanList(v []string)`

SetSanList sets SanList field to given value.

### HasSanList

`func (o *GenerateNewCertificateRequest) HasSanList() bool`

HasSanList returns a boolean if a field has been set.

### SetSanListNil

`func (o *GenerateNewCertificateRequest) SetSanListNil(b bool)`

 SetSanListNil sets the value for SanList to be an explicit nil

### UnsetSanList
`func (o *GenerateNewCertificateRequest) UnsetSanList()`

UnsetSanList ensures that no value is present for SanList, not even an explicit nil
### GetState

`func (o *GenerateNewCertificateRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GenerateNewCertificateRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GenerateNewCertificateRequest) SetState(v string)`

SetState sets State field to given value.


### SetStateNil

`func (o *GenerateNewCertificateRequest) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *GenerateNewCertificateRequest) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetTenantId

`func (o *GenerateNewCertificateRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *GenerateNewCertificateRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *GenerateNewCertificateRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *GenerateNewCertificateRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *GenerateNewCertificateRequest) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *GenerateNewCertificateRequest) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


