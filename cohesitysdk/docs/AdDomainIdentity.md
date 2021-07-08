# AdDomainIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dn** | Pointer to **NullableString** | Specifies distinguished name of the domain. | [optional] 
**Guid** | Pointer to **NullableString** | Specifies Unique objectGUID for an AD domain. | [optional] 
**Name** | Pointer to **NullableString** | Specifies display name of the domain. | [optional] 
**Sid** | Pointer to **NullableString** | Specifies domain SID. | [optional] 

## Methods

### NewAdDomainIdentity

`func NewAdDomainIdentity() *AdDomainIdentity`

NewAdDomainIdentity instantiates a new AdDomainIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdDomainIdentityWithDefaults

`func NewAdDomainIdentityWithDefaults() *AdDomainIdentity`

NewAdDomainIdentityWithDefaults instantiates a new AdDomainIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDn

`func (o *AdDomainIdentity) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *AdDomainIdentity) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *AdDomainIdentity) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *AdDomainIdentity) HasDn() bool`

HasDn returns a boolean if a field has been set.

### SetDnNil

`func (o *AdDomainIdentity) SetDnNil(b bool)`

 SetDnNil sets the value for Dn to be an explicit nil

### UnsetDn
`func (o *AdDomainIdentity) UnsetDn()`

UnsetDn ensures that no value is present for Dn, not even an explicit nil
### GetGuid

`func (o *AdDomainIdentity) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *AdDomainIdentity) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *AdDomainIdentity) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *AdDomainIdentity) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *AdDomainIdentity) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *AdDomainIdentity) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetName

`func (o *AdDomainIdentity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdDomainIdentity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdDomainIdentity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdDomainIdentity) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AdDomainIdentity) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AdDomainIdentity) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSid

`func (o *AdDomainIdentity) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *AdDomainIdentity) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *AdDomainIdentity) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *AdDomainIdentity) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *AdDomainIdentity) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *AdDomainIdentity) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


