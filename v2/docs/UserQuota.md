# UserQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **NullableString** | Specifies the domain name of the user, where the principal&#39; account is maintained. | [optional] 
**Sid** | Pointer to **NullableString** | Specifies the user sid. | [optional] 
**UnixUid** | Pointer to **NullableInt32** | Specifies the unix Uid. | [optional] 
**UserName** | Pointer to **NullableString** | Specifies the full name of the user | [optional] 
**QuotaPolicy** | Pointer to **map[string]interface{}** | Specifies the quota policy for the given user. | [optional] 
**UsageBytes** | Pointer to **NullableInt64** | Specifies the user usage in bytes. | [optional] [readonly] 

## Methods

### NewUserQuota

`func NewUserQuota() *UserQuota`

NewUserQuota instantiates a new UserQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserQuotaWithDefaults

`func NewUserQuotaWithDefaults() *UserQuota`

NewUserQuotaWithDefaults instantiates a new UserQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *UserQuota) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *UserQuota) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *UserQuota) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *UserQuota) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *UserQuota) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *UserQuota) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetSid

`func (o *UserQuota) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *UserQuota) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *UserQuota) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *UserQuota) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *UserQuota) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *UserQuota) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUnixUid

`func (o *UserQuota) GetUnixUid() int32`

GetUnixUid returns the UnixUid field if non-nil, zero value otherwise.

### GetUnixUidOk

`func (o *UserQuota) GetUnixUidOk() (*int32, bool)`

GetUnixUidOk returns a tuple with the UnixUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixUid

`func (o *UserQuota) SetUnixUid(v int32)`

SetUnixUid sets UnixUid field to given value.

### HasUnixUid

`func (o *UserQuota) HasUnixUid() bool`

HasUnixUid returns a boolean if a field has been set.

### SetUnixUidNil

`func (o *UserQuota) SetUnixUidNil(b bool)`

 SetUnixUidNil sets the value for UnixUid to be an explicit nil

### UnsetUnixUid
`func (o *UserQuota) UnsetUnixUid()`

UnsetUnixUid ensures that no value is present for UnixUid, not even an explicit nil
### GetUserName

`func (o *UserQuota) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UserQuota) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UserQuota) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *UserQuota) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *UserQuota) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *UserQuota) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetQuotaPolicy

`func (o *UserQuota) GetQuotaPolicy() map[string]interface{}`

GetQuotaPolicy returns the QuotaPolicy field if non-nil, zero value otherwise.

### GetQuotaPolicyOk

`func (o *UserQuota) GetQuotaPolicyOk() (*map[string]interface{}, bool)`

GetQuotaPolicyOk returns a tuple with the QuotaPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaPolicy

`func (o *UserQuota) SetQuotaPolicy(v map[string]interface{})`

SetQuotaPolicy sets QuotaPolicy field to given value.

### HasQuotaPolicy

`func (o *UserQuota) HasQuotaPolicy() bool`

HasQuotaPolicy returns a boolean if a field has been set.

### GetUsageBytes

`func (o *UserQuota) GetUsageBytes() int64`

GetUsageBytes returns the UsageBytes field if non-nil, zero value otherwise.

### GetUsageBytesOk

`func (o *UserQuota) GetUsageBytesOk() (*int64, bool)`

GetUsageBytesOk returns a tuple with the UsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageBytes

`func (o *UserQuota) SetUsageBytes(v int64)`

SetUsageBytes sets UsageBytes field to given value.

### HasUsageBytes

`func (o *UserQuota) HasUsageBytes() bool`

HasUsageBytes returns a boolean if a field has been set.

### SetUsageBytesNil

`func (o *UserQuota) SetUsageBytesNil(b bool)`

 SetUsageBytesNil sets the value for UsageBytes to be an explicit nil

### UnsetUsageBytes
`func (o *UserQuota) UnsetUsageBytes()`

UnsetUsageBytes ensures that no value is present for UsageBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


