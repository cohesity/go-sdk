# ResetS3SecretKeyParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **NullableString** | Specifies the fully qualified domain name (FQDN) of an Active Directory or LOCAL for the default LOCAL domain on the Cohesity Cluster. If not specified, it is assumed to be LOCAL. | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the tenant for which the users are to be deleted. | [optional] 
**Username** | Pointer to **NullableString** | Specifies the Cohesity user. | [optional] 

## Methods

### NewResetS3SecretKeyParameters

`func NewResetS3SecretKeyParameters() *ResetS3SecretKeyParameters`

NewResetS3SecretKeyParameters instantiates a new ResetS3SecretKeyParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResetS3SecretKeyParametersWithDefaults

`func NewResetS3SecretKeyParametersWithDefaults() *ResetS3SecretKeyParameters`

NewResetS3SecretKeyParametersWithDefaults instantiates a new ResetS3SecretKeyParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *ResetS3SecretKeyParameters) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ResetS3SecretKeyParameters) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ResetS3SecretKeyParameters) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ResetS3SecretKeyParameters) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *ResetS3SecretKeyParameters) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *ResetS3SecretKeyParameters) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetTenantId

`func (o *ResetS3SecretKeyParameters) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ResetS3SecretKeyParameters) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ResetS3SecretKeyParameters) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ResetS3SecretKeyParameters) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ResetS3SecretKeyParameters) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ResetS3SecretKeyParameters) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetUsername

`func (o *ResetS3SecretKeyParameters) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ResetS3SecretKeyParameters) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ResetS3SecretKeyParameters) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ResetS3SecretKeyParameters) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *ResetS3SecretKeyParameters) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ResetS3SecretKeyParameters) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


