# UnregisterKerberosRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminPassword** | **NullableString** | Specifies the password of Kerberos admin principal. | 
**AdminPrincipal** | Pointer to **NullableString** | Specifies the name of the Kerberos admin principal. | [optional] 

## Methods

### NewUnregisterKerberosRequest

`func NewUnregisterKerberosRequest(adminPassword NullableString, ) *UnregisterKerberosRequest`

NewUnregisterKerberosRequest instantiates a new UnregisterKerberosRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnregisterKerberosRequestWithDefaults

`func NewUnregisterKerberosRequestWithDefaults() *UnregisterKerberosRequest`

NewUnregisterKerberosRequestWithDefaults instantiates a new UnregisterKerberosRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminPassword

`func (o *UnregisterKerberosRequest) GetAdminPassword() string`

GetAdminPassword returns the AdminPassword field if non-nil, zero value otherwise.

### GetAdminPasswordOk

`func (o *UnregisterKerberosRequest) GetAdminPasswordOk() (*string, bool)`

GetAdminPasswordOk returns a tuple with the AdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPassword

`func (o *UnregisterKerberosRequest) SetAdminPassword(v string)`

SetAdminPassword sets AdminPassword field to given value.


### SetAdminPasswordNil

`func (o *UnregisterKerberosRequest) SetAdminPasswordNil(b bool)`

 SetAdminPasswordNil sets the value for AdminPassword to be an explicit nil

### UnsetAdminPassword
`func (o *UnregisterKerberosRequest) UnsetAdminPassword()`

UnsetAdminPassword ensures that no value is present for AdminPassword, not even an explicit nil
### GetAdminPrincipal

`func (o *UnregisterKerberosRequest) GetAdminPrincipal() string`

GetAdminPrincipal returns the AdminPrincipal field if non-nil, zero value otherwise.

### GetAdminPrincipalOk

`func (o *UnregisterKerberosRequest) GetAdminPrincipalOk() (*string, bool)`

GetAdminPrincipalOk returns a tuple with the AdminPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPrincipal

`func (o *UnregisterKerberosRequest) SetAdminPrincipal(v string)`

SetAdminPrincipal sets AdminPrincipal field to given value.

### HasAdminPrincipal

`func (o *UnregisterKerberosRequest) HasAdminPrincipal() bool`

HasAdminPrincipal returns a boolean if a field has been set.

### SetAdminPrincipalNil

`func (o *UnregisterKerberosRequest) SetAdminPrincipalNil(b bool)`

 SetAdminPrincipalNil sets the value for AdminPrincipal to be an explicit nil

### UnsetAdminPrincipal
`func (o *UnregisterKerberosRequest) UnsetAdminPrincipal()`

UnsetAdminPrincipal ensures that no value is present for AdminPrincipal, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


