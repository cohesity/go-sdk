# UserDeleteParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **NullableString** | Specifies the domain associated with the users to delete. Only users associated with the same domain can be deleted by a single request. If no domain is specified, the specified users are deleted from the LOCAL domain on the Cohesity Cluster. If a non-LOCAL domain is specified, the specified users are deleted on the Cohesity Cluster. However, the referenced user principals on the Active Directory are not deleted. | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the tenant for which the users are to be deleted. | [optional] 
**Users** | Pointer to **[]string** | Array of Users.  Specifies the list of users to delete on Cohesity Cluster. Only users from the same domain can be deleted by a single request. | [optional] 

## Methods

### NewUserDeleteParameters

`func NewUserDeleteParameters() *UserDeleteParameters`

NewUserDeleteParameters instantiates a new UserDeleteParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDeleteParametersWithDefaults

`func NewUserDeleteParametersWithDefaults() *UserDeleteParameters`

NewUserDeleteParametersWithDefaults instantiates a new UserDeleteParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *UserDeleteParameters) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *UserDeleteParameters) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *UserDeleteParameters) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *UserDeleteParameters) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *UserDeleteParameters) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *UserDeleteParameters) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetTenantId

`func (o *UserDeleteParameters) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UserDeleteParameters) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UserDeleteParameters) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *UserDeleteParameters) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *UserDeleteParameters) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *UserDeleteParameters) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetUsers

`func (o *UserDeleteParameters) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UserDeleteParameters) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UserDeleteParameters) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *UserDeleteParameters) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *UserDeleteParameters) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *UserDeleteParameters) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


