# ActiveDirectoryPrincipal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **NullableString** | Specifies the domain name of the where the principal&#39; account is maintained. | [optional] 
**FullName** | Pointer to **NullableString** | Specifies the full name (first and last names) of the principal. | [optional] 
**ObjectClass** | Pointer to **NullableString** | Specifies the object class of the principal (either &#39;kGroup&#39; or &#39;kUser&#39;). &#39;kUser&#39; specifies a user object class. &#39;kGroup&#39; specifies a group object class. &#39;kComputer&#39; specifies a computer object class. &#39;kWellKnownPrincipal&#39; specifies a well known principal. | [optional] 
**PrincipalName** | Pointer to **NullableString** | Specifies the name of the principal. | [optional] 
**Sid** | Pointer to **NullableString** | Specifies the unique Security id (SID) of the principal. | [optional] 

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

### GetDomain

`func (o *ActiveDirectoryPrincipal) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ActiveDirectoryPrincipal) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ActiveDirectoryPrincipal) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ActiveDirectoryPrincipal) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *ActiveDirectoryPrincipal) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *ActiveDirectoryPrincipal) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
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

### SetObjectClassNil

`func (o *ActiveDirectoryPrincipal) SetObjectClassNil(b bool)`

 SetObjectClassNil sets the value for ObjectClass to be an explicit nil

### UnsetObjectClass
`func (o *ActiveDirectoryPrincipal) UnsetObjectClass()`

UnsetObjectClass ensures that no value is present for ObjectClass, not even an explicit nil
### GetPrincipalName

`func (o *ActiveDirectoryPrincipal) GetPrincipalName() string`

GetPrincipalName returns the PrincipalName field if non-nil, zero value otherwise.

### GetPrincipalNameOk

`func (o *ActiveDirectoryPrincipal) GetPrincipalNameOk() (*string, bool)`

GetPrincipalNameOk returns a tuple with the PrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalName

`func (o *ActiveDirectoryPrincipal) SetPrincipalName(v string)`

SetPrincipalName sets PrincipalName field to given value.

### HasPrincipalName

`func (o *ActiveDirectoryPrincipal) HasPrincipalName() bool`

HasPrincipalName returns a boolean if a field has been set.

### SetPrincipalNameNil

`func (o *ActiveDirectoryPrincipal) SetPrincipalNameNil(b bool)`

 SetPrincipalNameNil sets the value for PrincipalName to be an explicit nil

### UnsetPrincipalName
`func (o *ActiveDirectoryPrincipal) UnsetPrincipalName()`

UnsetPrincipalName ensures that no value is present for PrincipalName, not even an explicit nil
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


