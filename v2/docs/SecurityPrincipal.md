# SecurityPrincipal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | Pointer to **NullableString** | Specifies the domain name where the security principal account is maintained. | [optional] 
**FullName** | Pointer to **NullableString** | Specifies the full name (first and last name) of the security principal. | [optional] 
**ObjectClass** | Pointer to **NullableString** | Specifies the object class of the security principal. | [optional] 
**PrincipalName** | Pointer to **NullableString** | Specifies the name of the security principal. | [optional] 
**Sid** | Pointer to **NullableString** | Specifies the SID of the security principal. | [optional] 

## Methods

### NewSecurityPrincipal

`func NewSecurityPrincipal() *SecurityPrincipal`

NewSecurityPrincipal instantiates a new SecurityPrincipal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityPrincipalWithDefaults

`func NewSecurityPrincipalWithDefaults() *SecurityPrincipal`

NewSecurityPrincipalWithDefaults instantiates a new SecurityPrincipal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainName

`func (o *SecurityPrincipal) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *SecurityPrincipal) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *SecurityPrincipal) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *SecurityPrincipal) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *SecurityPrincipal) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *SecurityPrincipal) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetFullName

`func (o *SecurityPrincipal) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *SecurityPrincipal) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *SecurityPrincipal) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *SecurityPrincipal) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *SecurityPrincipal) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *SecurityPrincipal) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetObjectClass

`func (o *SecurityPrincipal) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *SecurityPrincipal) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *SecurityPrincipal) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *SecurityPrincipal) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### SetObjectClassNil

`func (o *SecurityPrincipal) SetObjectClassNil(b bool)`

 SetObjectClassNil sets the value for ObjectClass to be an explicit nil

### UnsetObjectClass
`func (o *SecurityPrincipal) UnsetObjectClass()`

UnsetObjectClass ensures that no value is present for ObjectClass, not even an explicit nil
### GetPrincipalName

`func (o *SecurityPrincipal) GetPrincipalName() string`

GetPrincipalName returns the PrincipalName field if non-nil, zero value otherwise.

### GetPrincipalNameOk

`func (o *SecurityPrincipal) GetPrincipalNameOk() (*string, bool)`

GetPrincipalNameOk returns a tuple with the PrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalName

`func (o *SecurityPrincipal) SetPrincipalName(v string)`

SetPrincipalName sets PrincipalName field to given value.

### HasPrincipalName

`func (o *SecurityPrincipal) HasPrincipalName() bool`

HasPrincipalName returns a boolean if a field has been set.

### SetPrincipalNameNil

`func (o *SecurityPrincipal) SetPrincipalNameNil(b bool)`

 SetPrincipalNameNil sets the value for PrincipalName to be an explicit nil

### UnsetPrincipalName
`func (o *SecurityPrincipal) UnsetPrincipalName()`

UnsetPrincipalName ensures that no value is present for PrincipalName, not even an explicit nil
### GetSid

`func (o *SecurityPrincipal) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *SecurityPrincipal) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *SecurityPrincipal) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *SecurityPrincipal) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *SecurityPrincipal) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *SecurityPrincipal) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


