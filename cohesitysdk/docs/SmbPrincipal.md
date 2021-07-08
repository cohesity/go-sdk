# SmbPrincipal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **NullableString** | Specifies domain name of the principal. | [optional] 
**Name** | Pointer to **NullableString** | Specifies name of the SMB principal which may be a group or user. | [optional] 
**Sid** | Pointer to **NullableString** | Specifies unique Security ID (SID) of the principal that look similar to windows domain SID. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type. This can be a user or a group. | [optional] 

## Methods

### NewSmbPrincipal

`func NewSmbPrincipal() *SmbPrincipal`

NewSmbPrincipal instantiates a new SmbPrincipal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmbPrincipalWithDefaults

`func NewSmbPrincipalWithDefaults() *SmbPrincipal`

NewSmbPrincipalWithDefaults instantiates a new SmbPrincipal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *SmbPrincipal) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *SmbPrincipal) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *SmbPrincipal) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *SmbPrincipal) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *SmbPrincipal) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *SmbPrincipal) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetName

`func (o *SmbPrincipal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SmbPrincipal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SmbPrincipal) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SmbPrincipal) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SmbPrincipal) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SmbPrincipal) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSid

`func (o *SmbPrincipal) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *SmbPrincipal) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *SmbPrincipal) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *SmbPrincipal) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *SmbPrincipal) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *SmbPrincipal) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetType

`func (o *SmbPrincipal) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SmbPrincipal) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SmbPrincipal) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SmbPrincipal) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SmbPrincipal) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SmbPrincipal) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


