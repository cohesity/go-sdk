# UserId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **NullableString** | If interested in a user via smb_client, include SID. Otherwise, If a valid unix-id to SID mappings are available (i.e., when mixed mode is enabled) the server will perform the necessary id mapping and return the correct usage irrespective of whether the unix id / SID is provided. The string is of following format - S-1-IdentifierAuthority-SubAuthority1-SubAuthority2-...-SubAuthorityn. | [optional] 
**UnixUid** | Pointer to **NullableInt32** | If interested in a user via unix-identifier, include UnixUid. Otherwise, If a valid unix-id to SID mappings are available (i.e., when mixed mode is enabled) the server will perform the necessary id mapping and return the correct usage irrespective of whether the unix id / SID is provided. | [optional] 

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


