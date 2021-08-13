# ViewUserQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **NullableString** | Specifies the user sid. | [optional] 
**UnixUid** | Pointer to **NullableInt32** | Specifies the unix UID. | [optional] 
**QuotaPolicy** | Pointer to [**QuotaPolicy**](QuotaPolicy.md) | Specifies the user quota policy. | [optional] 
**UsageBytes** | Pointer to **NullableInt64** | Specifies the user usage in bytes. | [optional] [readonly] 

## Methods

### NewViewUserQuota

`func NewViewUserQuota() *ViewUserQuota`

NewViewUserQuota instantiates a new ViewUserQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewUserQuotaWithDefaults

`func NewViewUserQuotaWithDefaults() *ViewUserQuota`

NewViewUserQuotaWithDefaults instantiates a new ViewUserQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSid

`func (o *ViewUserQuota) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ViewUserQuota) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ViewUserQuota) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ViewUserQuota) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ViewUserQuota) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ViewUserQuota) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUnixUid

`func (o *ViewUserQuota) GetUnixUid() int32`

GetUnixUid returns the UnixUid field if non-nil, zero value otherwise.

### GetUnixUidOk

`func (o *ViewUserQuota) GetUnixUidOk() (*int32, bool)`

GetUnixUidOk returns a tuple with the UnixUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixUid

`func (o *ViewUserQuota) SetUnixUid(v int32)`

SetUnixUid sets UnixUid field to given value.

### HasUnixUid

`func (o *ViewUserQuota) HasUnixUid() bool`

HasUnixUid returns a boolean if a field has been set.

### SetUnixUidNil

`func (o *ViewUserQuota) SetUnixUidNil(b bool)`

 SetUnixUidNil sets the value for UnixUid to be an explicit nil

### UnsetUnixUid
`func (o *ViewUserQuota) UnsetUnixUid()`

UnsetUnixUid ensures that no value is present for UnixUid, not even an explicit nil
### GetQuotaPolicy

`func (o *ViewUserQuota) GetQuotaPolicy() QuotaPolicy`

GetQuotaPolicy returns the QuotaPolicy field if non-nil, zero value otherwise.

### GetQuotaPolicyOk

`func (o *ViewUserQuota) GetQuotaPolicyOk() (*QuotaPolicy, bool)`

GetQuotaPolicyOk returns a tuple with the QuotaPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaPolicy

`func (o *ViewUserQuota) SetQuotaPolicy(v QuotaPolicy)`

SetQuotaPolicy sets QuotaPolicy field to given value.

### HasQuotaPolicy

`func (o *ViewUserQuota) HasQuotaPolicy() bool`

HasQuotaPolicy returns a boolean if a field has been set.

### GetUsageBytes

`func (o *ViewUserQuota) GetUsageBytes() int64`

GetUsageBytes returns the UsageBytes field if non-nil, zero value otherwise.

### GetUsageBytesOk

`func (o *ViewUserQuota) GetUsageBytesOk() (*int64, bool)`

GetUsageBytesOk returns a tuple with the UsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageBytes

`func (o *ViewUserQuota) SetUsageBytes(v int64)`

SetUsageBytes sets UsageBytes field to given value.

### HasUsageBytes

`func (o *ViewUserQuota) HasUsageBytes() bool`

HasUsageBytes returns a boolean if a field has been set.

### SetUsageBytesNil

`func (o *ViewUserQuota) SetUsageBytesNil(b bool)`

 SetUsageBytesNil sets the value for UsageBytes to be an explicit nil

### UnsetUsageBytes
`func (o *ViewUserQuota) UnsetUsageBytes()`

UnsetUsageBytes ensures that no value is present for UsageBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


