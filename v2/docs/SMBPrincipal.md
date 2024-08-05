# SMBPrincipal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **NullableString** | Specifies the domain of the principal. For active directories, this is the fully qualified domain name (FQDN). | [optional] 
**Name** | Pointer to **NullableString** | Specifies the principal name. | [optional] 
**ObjectClass** | Pointer to **NullableString** | Specifies the principal class. | [optional] 
**Sid** | Pointer to **NullableString** | Specifies the SID of the principal. | [optional] 

## Methods

### NewSMBPrincipal

`func NewSMBPrincipal() *SMBPrincipal`

NewSMBPrincipal instantiates a new SMBPrincipal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMBPrincipalWithDefaults

`func NewSMBPrincipalWithDefaults() *SMBPrincipal`

NewSMBPrincipalWithDefaults instantiates a new SMBPrincipal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *SMBPrincipal) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *SMBPrincipal) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *SMBPrincipal) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *SMBPrincipal) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *SMBPrincipal) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *SMBPrincipal) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetName

`func (o *SMBPrincipal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SMBPrincipal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SMBPrincipal) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SMBPrincipal) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SMBPrincipal) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SMBPrincipal) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetObjectClass

`func (o *SMBPrincipal) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *SMBPrincipal) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *SMBPrincipal) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *SMBPrincipal) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### SetObjectClassNil

`func (o *SMBPrincipal) SetObjectClassNil(b bool)`

 SetObjectClassNil sets the value for ObjectClass to be an explicit nil

### UnsetObjectClass
`func (o *SMBPrincipal) UnsetObjectClass()`

UnsetObjectClass ensures that no value is present for ObjectClass, not even an explicit nil
### GetSid

`func (o *SMBPrincipal) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *SMBPrincipal) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *SMBPrincipal) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *SMBPrincipal) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *SMBPrincipal) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *SMBPrincipal) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


