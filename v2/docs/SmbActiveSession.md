# SmbActiveSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveOpens** | Pointer to [**[]SmbActiveOpen**](SmbActiveOpen.md) | Specifies an active open of an SMB file, its access and sharing information. | [optional] 
**ClientIp** | Pointer to **NullableString** | Specifies the IP address from which the file is open. | [optional] 
**Domain** | Pointer to **NullableString** | Specifies the domain of the user. | [optional] 
**SessionId** | Pointer to **NullableInt64** | Specifies the id of the session. | [optional] 
**UserName** | Pointer to **NullableString** | Specifies the username who keeps the file open. | [optional] 

## Methods

### NewSmbActiveSession

`func NewSmbActiveSession() *SmbActiveSession`

NewSmbActiveSession instantiates a new SmbActiveSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmbActiveSessionWithDefaults

`func NewSmbActiveSessionWithDefaults() *SmbActiveSession`

NewSmbActiveSessionWithDefaults instantiates a new SmbActiveSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveOpens

`func (o *SmbActiveSession) GetActiveOpens() []SmbActiveOpen`

GetActiveOpens returns the ActiveOpens field if non-nil, zero value otherwise.

### GetActiveOpensOk

`func (o *SmbActiveSession) GetActiveOpensOk() (*[]SmbActiveOpen, bool)`

GetActiveOpensOk returns a tuple with the ActiveOpens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveOpens

`func (o *SmbActiveSession) SetActiveOpens(v []SmbActiveOpen)`

SetActiveOpens sets ActiveOpens field to given value.

### HasActiveOpens

`func (o *SmbActiveSession) HasActiveOpens() bool`

HasActiveOpens returns a boolean if a field has been set.

### SetActiveOpensNil

`func (o *SmbActiveSession) SetActiveOpensNil(b bool)`

 SetActiveOpensNil sets the value for ActiveOpens to be an explicit nil

### UnsetActiveOpens
`func (o *SmbActiveSession) UnsetActiveOpens()`

UnsetActiveOpens ensures that no value is present for ActiveOpens, not even an explicit nil
### GetClientIp

`func (o *SmbActiveSession) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *SmbActiveSession) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *SmbActiveSession) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *SmbActiveSession) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### SetClientIpNil

`func (o *SmbActiveSession) SetClientIpNil(b bool)`

 SetClientIpNil sets the value for ClientIp to be an explicit nil

### UnsetClientIp
`func (o *SmbActiveSession) UnsetClientIp()`

UnsetClientIp ensures that no value is present for ClientIp, not even an explicit nil
### GetDomain

`func (o *SmbActiveSession) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *SmbActiveSession) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *SmbActiveSession) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *SmbActiveSession) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *SmbActiveSession) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *SmbActiveSession) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetSessionId

`func (o *SmbActiveSession) GetSessionId() int64`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *SmbActiveSession) GetSessionIdOk() (*int64, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *SmbActiveSession) SetSessionId(v int64)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *SmbActiveSession) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionIdNil

`func (o *SmbActiveSession) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *SmbActiveSession) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
### GetUserName

`func (o *SmbActiveSession) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SmbActiveSession) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SmbActiveSession) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *SmbActiveSession) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *SmbActiveSession) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *SmbActiveSession) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


