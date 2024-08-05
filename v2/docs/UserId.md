# UserId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **NullableString** | Specifies the domain name of the user, where the principal&#39; account is maintained. | [optional] 
**Sid** | Pointer to **NullableString** | Specifies the user sid. | [optional] 
**UnixUid** | Pointer to **NullableInt32** | Specifies the unix Uid. | [optional] 
**UserName** | Pointer to **NullableString** | Specifies the full name of the user | [optional] 

## Methods

### NewUserId

`func NewUserId() *UserId`

NewUserId instantiates a new UserId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserIdWithDefaults

`func NewUserIdWithDefaults() *UserId`

NewUserIdWithDefaults instantiates a new UserId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *UserId) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *UserId) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *UserId) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *UserId) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *UserId) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *UserId) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetSid

`func (o *UserId) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *UserId) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *UserId) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *UserId) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *UserId) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *UserId) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUnixUid

`func (o *UserId) GetUnixUid() int32`

GetUnixUid returns the UnixUid field if non-nil, zero value otherwise.

### GetUnixUidOk

`func (o *UserId) GetUnixUidOk() (*int32, bool)`

GetUnixUidOk returns a tuple with the UnixUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixUid

`func (o *UserId) SetUnixUid(v int32)`

SetUnixUid sets UnixUid field to given value.

### HasUnixUid

`func (o *UserId) HasUnixUid() bool`

HasUnixUid returns a boolean if a field has been set.

### SetUnixUidNil

`func (o *UserId) SetUnixUidNil(b bool)`

 SetUnixUidNil sets the value for UnixUid to be an explicit nil

### UnsetUnixUid
`func (o *UserId) UnsetUnixUid()`

UnsetUnixUid ensures that no value is present for UnixUid, not even an explicit nil
### GetUserName

`func (o *UserId) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UserId) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UserId) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *UserId) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *UserId) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *UserId) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


