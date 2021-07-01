# Principal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **NullableString** | Specifies the domain name of the where the principal&#39; account is maintained. | [optional] 
**FullName** | Pointer to **NullableString** | Specifies the full name (first and last names) of the principal. | [optional] 
**ObjectClass** | Pointer to **NullableString** | Specifies the object class of the principal (either &#39;kGroup&#39; or &#39;kUser&#39;). &#39;kUser&#39; specifies a user object class. &#39;kGroup&#39; specifies a group object class. &#39;kComputer&#39; specifies a computer object class. &#39;kWellKnownPrincipal&#39; specifies a well known principal. | [optional] 
**PrincipalName** | Pointer to **NullableString** | Specifies the name of the principal. | [optional] 
**Sid** | Pointer to **NullableString** | Specifies the unique Security id (SID) of the principal. | [optional] 

## Methods

### NewPrincipal

`func NewPrincipal() *Principal`

NewPrincipal instantiates a new Principal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalWithDefaults

`func NewPrincipalWithDefaults() *Principal`

NewPrincipalWithDefaults instantiates a new Principal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *Principal) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Principal) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Principal) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Principal) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *Principal) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *Principal) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetFullName

`func (o *Principal) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *Principal) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *Principal) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *Principal) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *Principal) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *Principal) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetObjectClass

`func (o *Principal) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *Principal) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *Principal) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *Principal) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### SetObjectClassNil

`func (o *Principal) SetObjectClassNil(b bool)`

 SetObjectClassNil sets the value for ObjectClass to be an explicit nil

### UnsetObjectClass
`func (o *Principal) UnsetObjectClass()`

UnsetObjectClass ensures that no value is present for ObjectClass, not even an explicit nil
### GetPrincipalName

`func (o *Principal) GetPrincipalName() string`

GetPrincipalName returns the PrincipalName field if non-nil, zero value otherwise.

### GetPrincipalNameOk

`func (o *Principal) GetPrincipalNameOk() (*string, bool)`

GetPrincipalNameOk returns a tuple with the PrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalName

`func (o *Principal) SetPrincipalName(v string)`

SetPrincipalName sets PrincipalName field to given value.

### HasPrincipalName

`func (o *Principal) HasPrincipalName() bool`

HasPrincipalName returns a boolean if a field has been set.

### SetPrincipalNameNil

`func (o *Principal) SetPrincipalNameNil(b bool)`

 SetPrincipalNameNil sets the value for PrincipalName to be an explicit nil

### UnsetPrincipalName
`func (o *Principal) UnsetPrincipalName()`

UnsetPrincipalName ensures that no value is present for PrincipalName, not even an explicit nil
### GetSid

`func (o *Principal) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *Principal) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *Principal) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *Principal) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *Principal) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *Principal) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


