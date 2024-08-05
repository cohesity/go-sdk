# UserSessionsCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionsCount** | Pointer to **int64** | Specifies the number of sessions for the user. | [optional] 
**Sid** | Pointer to **string** | Specifies the user sid. | [optional] 

## Methods

### NewUserSessionsCount

`func NewUserSessionsCount() *UserSessionsCount`

NewUserSessionsCount instantiates a new UserSessionsCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSessionsCountWithDefaults

`func NewUserSessionsCountWithDefaults() *UserSessionsCount`

NewUserSessionsCountWithDefaults instantiates a new UserSessionsCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionsCount

`func (o *UserSessionsCount) GetSessionsCount() int64`

GetSessionsCount returns the SessionsCount field if non-nil, zero value otherwise.

### GetSessionsCountOk

`func (o *UserSessionsCount) GetSessionsCountOk() (*int64, bool)`

GetSessionsCountOk returns a tuple with the SessionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionsCount

`func (o *UserSessionsCount) SetSessionsCount(v int64)`

SetSessionsCount sets SessionsCount field to given value.

### HasSessionsCount

`func (o *UserSessionsCount) HasSessionsCount() bool`

HasSessionsCount returns a boolean if a field has been set.

### GetSid

`func (o *UserSessionsCount) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *UserSessionsCount) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *UserSessionsCount) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *UserSessionsCount) HasSid() bool`

HasSid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


