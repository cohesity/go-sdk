# ActiveDirectoryPrincipal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | Pointer to **NullableString** | Specifies the domain name to which the principal belongs to | [optional] 
**FullName** | Pointer to **NullableString** | Specifies the full name of the principal. | [optional] 
**Name** | Pointer to **string** | Specifies the name of the principal which is being added. | [optional] 
**ObjectClass** | Pointer to **string** | Specifies the type of principal, a user or a group | [optional] 
**Sid** | Pointer to **NullableString** | Specifies the unique SID of the principal. | [optional] 

## Methods

### NewActiveDirectoryPrincipal

`func NewActiveDirectoryPrincipal() *ActiveDirectoryPrincipal`

NewActiveDirectoryPrincipal instantiates a new ActiveDirectoryPrincipal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveDirectoryPrincipalWithDefaults

`func NewActiveDirectoryPrincipalWithDefaults() *ActiveDirectoryPrincipal`

NewActiveDirectoryPrincipalWithDefaults instantiates a new ActiveDirectoryPrincipal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainName

`func (o *ActiveDirectoryPrincipal) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *ActiveDirectoryPrincipal) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *ActiveDirectoryPrincipal) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *ActiveDirectoryPrincipal) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *ActiveDirectoryPrincipal) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *ActiveDirectoryPrincipal) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetFullName

`func (o *ActiveDirectoryPrincipal) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ActiveDirectoryPrincipal) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ActiveDirectoryPrincipal) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *ActiveDirectoryPrincipal) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *ActiveDirectoryPrincipal) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *ActiveDirectoryPrincipal) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetName

`func (o *ActiveDirectoryPrincipal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActiveDirectoryPrincipal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActiveDirectoryPrincipal) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ActiveDirectoryPrincipal) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectClass

`func (o *ActiveDirectoryPrincipal) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *ActiveDirectoryPrincipal) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *ActiveDirectoryPrincipal) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *ActiveDirectoryPrincipal) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetSid

`func (o *ActiveDirectoryPrincipal) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ActiveDirectoryPrincipal) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ActiveDirectoryPrincipal) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ActiveDirectoryPrincipal) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ActiveDirectoryPrincipal) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ActiveDirectoryPrincipal) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


