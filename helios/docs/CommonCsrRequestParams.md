# CommonCsrRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organization** | **NullableString** | Specifies the organization attribute, which is part of the distinguished name definition. It is used to specify the name of the company. | 
**OrganizationUnit** | **NullableString** | Specifies the organization unit attribute, which is part of the distinguished name definition. It is used to identify the specific department or business unit in the company that is owning the Cluster. | 
**CountryCode** | **NullableString** | Specifies the country attribute, which is part of the distinguished name definition. It is used to identify the country where the state is located. It is specified as two letter code defined by the ISO standard. | 
**State** | **NullableString** | Specifies the state attribute, which is part of the distinguished name definition. It is used to identify the state where the city is located. | 
**City** | **NullableString** | Specifies the locality attribute, which is part of the distinguished name definition. It is used to identify the city where the company is located or the Cluster is installed. | 
**KeyType** | Pointer to **NullableString** | Specifies the algorithm to be used to generate the key pair. RSA is the default value. | [optional] [default to "rsa"]
**KeySizeBits** | Pointer to **NullableInt64** | Specifies the size of the keys in bits. The default is 2048 bits for the RSA keys and 256 bits for ECDSA. | [optional] 
**CommonName** | Pointer to **NullableString** | Specifies the common name attribute, which is part of the distinguished name definition. Common name is used to specify a context for the certificate, for example, the name of the Cluster to which the certificate is to be assigned. Default value is the name of the Cluster. | [optional] 
**DnsNames** | Pointer to **[]string** | Specifies an alternative subject name component to be included in the certificate. It is used to identify the ways the Cluster will be accessed. It is given as a comma separated list of FQDNs. The default value is the Cluster&#39;s VIP hostname. | [optional] 
**HostIps** | Pointer to **[]string** | Specifies an alternative subject name component to be included in the certificate. It is used to identify the ways the Cluster will be accessed. It is given as a comma separated list of IP addresses. The default value is the Cluster&#39;s VIP addresses. | [optional] 
**EmailAddress** | Pointer to **NullableString** | Specifies an alternative subject name component to be included in the certificate. Format is a standard e-mail address, for example joe@company.com. | [optional] 
**ServiceName** | Pointer to **NullableString** | Specifies the Cohesity service name for which the CSR is generated. Default service name is iris. | [optional] [default to "iris"]

## Methods

### NewCommonCsrRequestParams

`func NewCommonCsrRequestParams(organization NullableString, organizationUnit NullableString, countryCode NullableString, state NullableString, city NullableString, ) *CommonCsrRequestParams`

NewCommonCsrRequestParams instantiates a new CommonCsrRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonCsrRequestParamsWithDefaults

`func NewCommonCsrRequestParamsWithDefaults() *CommonCsrRequestParams`

NewCommonCsrRequestParamsWithDefaults instantiates a new CommonCsrRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganization

`func (o *CommonCsrRequestParams) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CommonCsrRequestParams) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CommonCsrRequestParams) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### SetOrganizationNil

`func (o *CommonCsrRequestParams) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *CommonCsrRequestParams) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetOrganizationUnit

`func (o *CommonCsrRequestParams) GetOrganizationUnit() string`

GetOrganizationUnit returns the OrganizationUnit field if non-nil, zero value otherwise.

### GetOrganizationUnitOk

`func (o *CommonCsrRequestParams) GetOrganizationUnitOk() (*string, bool)`

GetOrganizationUnitOk returns a tuple with the OrganizationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationUnit

`func (o *CommonCsrRequestParams) SetOrganizationUnit(v string)`

SetOrganizationUnit sets OrganizationUnit field to given value.


### SetOrganizationUnitNil

`func (o *CommonCsrRequestParams) SetOrganizationUnitNil(b bool)`

 SetOrganizationUnitNil sets the value for OrganizationUnit to be an explicit nil

### UnsetOrganizationUnit
`func (o *CommonCsrRequestParams) UnsetOrganizationUnit()`

UnsetOrganizationUnit ensures that no value is present for OrganizationUnit, not even an explicit nil
### GetCountryCode

`func (o *CommonCsrRequestParams) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *CommonCsrRequestParams) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *CommonCsrRequestParams) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### SetCountryCodeNil

`func (o *CommonCsrRequestParams) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *CommonCsrRequestParams) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetState

`func (o *CommonCsrRequestParams) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CommonCsrRequestParams) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CommonCsrRequestParams) SetState(v string)`

SetState sets State field to given value.


### SetStateNil

`func (o *CommonCsrRequestParams) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *CommonCsrRequestParams) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetCity

`func (o *CommonCsrRequestParams) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CommonCsrRequestParams) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CommonCsrRequestParams) SetCity(v string)`

SetCity sets City field to given value.


### SetCityNil

`func (o *CommonCsrRequestParams) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *CommonCsrRequestParams) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetKeyType

`func (o *CommonCsrRequestParams) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *CommonCsrRequestParams) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *CommonCsrRequestParams) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *CommonCsrRequestParams) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### SetKeyTypeNil

`func (o *CommonCsrRequestParams) SetKeyTypeNil(b bool)`

 SetKeyTypeNil sets the value for KeyType to be an explicit nil

### UnsetKeyType
`func (o *CommonCsrRequestParams) UnsetKeyType()`

UnsetKeyType ensures that no value is present for KeyType, not even an explicit nil
### GetKeySizeBits

`func (o *CommonCsrRequestParams) GetKeySizeBits() int64`

GetKeySizeBits returns the KeySizeBits field if non-nil, zero value otherwise.

### GetKeySizeBitsOk

`func (o *CommonCsrRequestParams) GetKeySizeBitsOk() (*int64, bool)`

GetKeySizeBitsOk returns a tuple with the KeySizeBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySizeBits

`func (o *CommonCsrRequestParams) SetKeySizeBits(v int64)`

SetKeySizeBits sets KeySizeBits field to given value.

### HasKeySizeBits

`func (o *CommonCsrRequestParams) HasKeySizeBits() bool`

HasKeySizeBits returns a boolean if a field has been set.

### SetKeySizeBitsNil

`func (o *CommonCsrRequestParams) SetKeySizeBitsNil(b bool)`

 SetKeySizeBitsNil sets the value for KeySizeBits to be an explicit nil

### UnsetKeySizeBits
`func (o *CommonCsrRequestParams) UnsetKeySizeBits()`

UnsetKeySizeBits ensures that no value is present for KeySizeBits, not even an explicit nil
### GetCommonName

`func (o *CommonCsrRequestParams) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *CommonCsrRequestParams) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *CommonCsrRequestParams) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *CommonCsrRequestParams) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### SetCommonNameNil

`func (o *CommonCsrRequestParams) SetCommonNameNil(b bool)`

 SetCommonNameNil sets the value for CommonName to be an explicit nil

### UnsetCommonName
`func (o *CommonCsrRequestParams) UnsetCommonName()`

UnsetCommonName ensures that no value is present for CommonName, not even an explicit nil
### GetDnsNames

`func (o *CommonCsrRequestParams) GetDnsNames() []string`

GetDnsNames returns the DnsNames field if non-nil, zero value otherwise.

### GetDnsNamesOk

`func (o *CommonCsrRequestParams) GetDnsNamesOk() (*[]string, bool)`

GetDnsNamesOk returns a tuple with the DnsNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNames

`func (o *CommonCsrRequestParams) SetDnsNames(v []string)`

SetDnsNames sets DnsNames field to given value.

### HasDnsNames

`func (o *CommonCsrRequestParams) HasDnsNames() bool`

HasDnsNames returns a boolean if a field has been set.

### SetDnsNamesNil

`func (o *CommonCsrRequestParams) SetDnsNamesNil(b bool)`

 SetDnsNamesNil sets the value for DnsNames to be an explicit nil

### UnsetDnsNames
`func (o *CommonCsrRequestParams) UnsetDnsNames()`

UnsetDnsNames ensures that no value is present for DnsNames, not even an explicit nil
### GetHostIps

`func (o *CommonCsrRequestParams) GetHostIps() []string`

GetHostIps returns the HostIps field if non-nil, zero value otherwise.

### GetHostIpsOk

`func (o *CommonCsrRequestParams) GetHostIpsOk() (*[]string, bool)`

GetHostIpsOk returns a tuple with the HostIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostIps

`func (o *CommonCsrRequestParams) SetHostIps(v []string)`

SetHostIps sets HostIps field to given value.

### HasHostIps

`func (o *CommonCsrRequestParams) HasHostIps() bool`

HasHostIps returns a boolean if a field has been set.

### SetHostIpsNil

`func (o *CommonCsrRequestParams) SetHostIpsNil(b bool)`

 SetHostIpsNil sets the value for HostIps to be an explicit nil

### UnsetHostIps
`func (o *CommonCsrRequestParams) UnsetHostIps()`

UnsetHostIps ensures that no value is present for HostIps, not even an explicit nil
### GetEmailAddress

`func (o *CommonCsrRequestParams) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *CommonCsrRequestParams) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *CommonCsrRequestParams) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *CommonCsrRequestParams) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *CommonCsrRequestParams) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *CommonCsrRequestParams) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetServiceName

`func (o *CommonCsrRequestParams) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *CommonCsrRequestParams) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *CommonCsrRequestParams) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *CommonCsrRequestParams) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### SetServiceNameNil

`func (o *CommonCsrRequestParams) SetServiceNameNil(b bool)`

 SetServiceNameNil sets the value for ServiceName to be an explicit nil

### UnsetServiceName
`func (o *CommonCsrRequestParams) UnsetServiceName()`

UnsetServiceName ensures that no value is present for ServiceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


