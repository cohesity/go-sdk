# CreateCsrRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | **NullableString** | Specifies the locality attribute, which is part of the distinguished name definition. It is used to identify the city where the company is located or the Cluster is installed. | 
**CommonName** | Pointer to **NullableString** | Specifies the common name attribute, which is part of the distinguished name definition. Common name is used to specify a context for the certificate, for example, the name of the Cluster to which the certificate is to be assigned. Default value is the name of the Cluster. | [optional] 
**CountryCode** | **NullableString** | Specifies the country attribute, which is part of the distinguished name definition. It is used to identify the country where the state is located. It is specified as two letter code defined by the ISO standard. | 
**DnsNames** | Pointer to **[]string** | Specifies an alternative subject name component to be included in the certificate. It is used to identify the ways the Cluster will be accessed. It is given as a comma separated list of FQDNs. The default value is the Cluster&#39;s VIP hostname. | [optional] 
**EmailAddress** | Pointer to **NullableString** | Specifies an alternative subject name component to be included in the certificate. Format is a standard e-mail address, for example joe@company.com. | [optional] 
**HostIps** | Pointer to **[]string** | Specifies an alternative subject name component to be included in the certificate. It is used to identify the ways the Cluster will be accessed. It is given as a comma separated list of IP addresses. The default value is the Cluster&#39;s VIP addresses. | [optional] 
**KeySizeBits** | Pointer to **NullableInt64** | Specifies the size of the keys in bits. The default is 2048 bits for the RSA keys and 256 bits for ECDSA. | [optional] 
**KeyType** | Pointer to **NullableString** | Specifies the algorithm to be used to generate the key pair. RSA is the default value. | [optional] [default to "rsa"]
**Organization** | **NullableString** | Specifies the organization attribute, which is part of the distinguished name definition. It is used to specify the name of the company. | 
**OrganizationUnit** | **NullableString** | Specifies the organization unit attribute, which is part of the distinguished name definition. It is used to identify the specific department or business unit in the company that is owning the Cluster. | 
**ServiceName** | Pointer to **NullableString** | Specifies the Cohesity service name for which the CSR is generated. Default service name is iris. | [optional] [default to "iris"]
**State** | **NullableString** | Specifies the state attribute, which is part of the distinguished name definition. It is used to identify the state where the city is located. | 

## Methods

### NewCreateCsrRequest

`func NewCreateCsrRequest(city NullableString, countryCode NullableString, organization NullableString, organizationUnit NullableString, state NullableString, ) *CreateCsrRequest`

NewCreateCsrRequest instantiates a new CreateCsrRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCsrRequestWithDefaults

`func NewCreateCsrRequestWithDefaults() *CreateCsrRequest`

NewCreateCsrRequestWithDefaults instantiates a new CreateCsrRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *CreateCsrRequest) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CreateCsrRequest) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CreateCsrRequest) SetCity(v string)`

SetCity sets City field to given value.


### SetCityNil

`func (o *CreateCsrRequest) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *CreateCsrRequest) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCommonName

`func (o *CreateCsrRequest) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *CreateCsrRequest) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *CreateCsrRequest) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *CreateCsrRequest) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### SetCommonNameNil

`func (o *CreateCsrRequest) SetCommonNameNil(b bool)`

 SetCommonNameNil sets the value for CommonName to be an explicit nil

### UnsetCommonName
`func (o *CreateCsrRequest) UnsetCommonName()`

UnsetCommonName ensures that no value is present for CommonName, not even an explicit nil
### GetCountryCode

`func (o *CreateCsrRequest) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *CreateCsrRequest) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *CreateCsrRequest) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### SetCountryCodeNil

`func (o *CreateCsrRequest) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *CreateCsrRequest) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetDnsNames

`func (o *CreateCsrRequest) GetDnsNames() []string`

GetDnsNames returns the DnsNames field if non-nil, zero value otherwise.

### GetDnsNamesOk

`func (o *CreateCsrRequest) GetDnsNamesOk() (*[]string, bool)`

GetDnsNamesOk returns a tuple with the DnsNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNames

`func (o *CreateCsrRequest) SetDnsNames(v []string)`

SetDnsNames sets DnsNames field to given value.

### HasDnsNames

`func (o *CreateCsrRequest) HasDnsNames() bool`

HasDnsNames returns a boolean if a field has been set.

### SetDnsNamesNil

`func (o *CreateCsrRequest) SetDnsNamesNil(b bool)`

 SetDnsNamesNil sets the value for DnsNames to be an explicit nil

### UnsetDnsNames
`func (o *CreateCsrRequest) UnsetDnsNames()`

UnsetDnsNames ensures that no value is present for DnsNames, not even an explicit nil
### GetEmailAddress

`func (o *CreateCsrRequest) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *CreateCsrRequest) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *CreateCsrRequest) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *CreateCsrRequest) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *CreateCsrRequest) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *CreateCsrRequest) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetHostIps

`func (o *CreateCsrRequest) GetHostIps() []string`

GetHostIps returns the HostIps field if non-nil, zero value otherwise.

### GetHostIpsOk

`func (o *CreateCsrRequest) GetHostIpsOk() (*[]string, bool)`

GetHostIpsOk returns a tuple with the HostIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostIps

`func (o *CreateCsrRequest) SetHostIps(v []string)`

SetHostIps sets HostIps field to given value.

### HasHostIps

`func (o *CreateCsrRequest) HasHostIps() bool`

HasHostIps returns a boolean if a field has been set.

### SetHostIpsNil

`func (o *CreateCsrRequest) SetHostIpsNil(b bool)`

 SetHostIpsNil sets the value for HostIps to be an explicit nil

### UnsetHostIps
`func (o *CreateCsrRequest) UnsetHostIps()`

UnsetHostIps ensures that no value is present for HostIps, not even an explicit nil
### GetKeySizeBits

`func (o *CreateCsrRequest) GetKeySizeBits() int64`

GetKeySizeBits returns the KeySizeBits field if non-nil, zero value otherwise.

### GetKeySizeBitsOk

`func (o *CreateCsrRequest) GetKeySizeBitsOk() (*int64, bool)`

GetKeySizeBitsOk returns a tuple with the KeySizeBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySizeBits

`func (o *CreateCsrRequest) SetKeySizeBits(v int64)`

SetKeySizeBits sets KeySizeBits field to given value.

### HasKeySizeBits

`func (o *CreateCsrRequest) HasKeySizeBits() bool`

HasKeySizeBits returns a boolean if a field has been set.

### SetKeySizeBitsNil

`func (o *CreateCsrRequest) SetKeySizeBitsNil(b bool)`

 SetKeySizeBitsNil sets the value for KeySizeBits to be an explicit nil

### UnsetKeySizeBits
`func (o *CreateCsrRequest) UnsetKeySizeBits()`

UnsetKeySizeBits ensures that no value is present for KeySizeBits, not even an explicit nil
### GetKeyType

`func (o *CreateCsrRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *CreateCsrRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *CreateCsrRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *CreateCsrRequest) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### SetKeyTypeNil

`func (o *CreateCsrRequest) SetKeyTypeNil(b bool)`

 SetKeyTypeNil sets the value for KeyType to be an explicit nil

### UnsetKeyType
`func (o *CreateCsrRequest) UnsetKeyType()`

UnsetKeyType ensures that no value is present for KeyType, not even an explicit nil
### GetOrganization

`func (o *CreateCsrRequest) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CreateCsrRequest) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CreateCsrRequest) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### SetOrganizationNil

`func (o *CreateCsrRequest) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *CreateCsrRequest) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetOrganizationUnit

`func (o *CreateCsrRequest) GetOrganizationUnit() string`

GetOrganizationUnit returns the OrganizationUnit field if non-nil, zero value otherwise.

### GetOrganizationUnitOk

`func (o *CreateCsrRequest) GetOrganizationUnitOk() (*string, bool)`

GetOrganizationUnitOk returns a tuple with the OrganizationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationUnit

`func (o *CreateCsrRequest) SetOrganizationUnit(v string)`

SetOrganizationUnit sets OrganizationUnit field to given value.


### SetOrganizationUnitNil

`func (o *CreateCsrRequest) SetOrganizationUnitNil(b bool)`

 SetOrganizationUnitNil sets the value for OrganizationUnit to be an explicit nil

### UnsetOrganizationUnit
`func (o *CreateCsrRequest) UnsetOrganizationUnit()`

UnsetOrganizationUnit ensures that no value is present for OrganizationUnit, not even an explicit nil
### GetServiceName

`func (o *CreateCsrRequest) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *CreateCsrRequest) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *CreateCsrRequest) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *CreateCsrRequest) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### SetServiceNameNil

`func (o *CreateCsrRequest) SetServiceNameNil(b bool)`

 SetServiceNameNil sets the value for ServiceName to be an explicit nil

### UnsetServiceName
`func (o *CreateCsrRequest) UnsetServiceName()`

UnsetServiceName ensures that no value is present for ServiceName, not even an explicit nil
### GetState

`func (o *CreateCsrRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreateCsrRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreateCsrRequest) SetState(v string)`

SetState sets State field to given value.


### SetStateNil

`func (o *CreateCsrRequest) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *CreateCsrRequest) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


