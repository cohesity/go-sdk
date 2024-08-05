# ActiveSessionsCountParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionsPerUser** | Pointer to [**[]UserSessionsCount**](UserSessionsCount.md) | Specifies the sessions count per user. | [optional] 
**TotalSessionsCount** | Pointer to **int64** | Specifies the aggregated sessions count for the user sessions returned. If sids are not given this returns the total system wide sessions count and if the sids are given, this returns the total sessions count for the given sids. | [optional] 

## Methods

### NewActiveSessionsCountParams

`func NewActiveSessionsCountParams() *ActiveSessionsCountParams`

NewActiveSessionsCountParams instantiates a new ActiveSessionsCountParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveSessionsCountParamsWithDefaults

`func NewActiveSessionsCountParamsWithDefaults() *ActiveSessionsCountParams`

NewActiveSessionsCountParamsWithDefaults instantiates a new ActiveSessionsCountParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionsPerUser

`func (o *ActiveSessionsCountParams) GetSessionsPerUser() []UserSessionsCount`

GetSessionsPerUser returns the SessionsPerUser field if non-nil, zero value otherwise.

### GetSessionsPerUserOk

`func (o *ActiveSessionsCountParams) GetSessionsPerUserOk() (*[]UserSessionsCount, bool)`

GetSessionsPerUserOk returns a tuple with the SessionsPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionsPerUser

`func (o *ActiveSessionsCountParams) SetSessionsPerUser(v []UserSessionsCount)`

SetSessionsPerUser sets SessionsPerUser field to given value.

### HasSessionsPerUser

`func (o *ActiveSessionsCountParams) HasSessionsPerUser() bool`

HasSessionsPerUser returns a boolean if a field has been set.

### GetTotalSessionsCount

`func (o *ActiveSessionsCountParams) GetTotalSessionsCount() int64`

GetTotalSessionsCount returns the TotalSessionsCount field if non-nil, zero value otherwise.

### GetTotalSessionsCountOk

`func (o *ActiveSessionsCountParams) GetTotalSessionsCountOk() (*int64, bool)`

GetTotalSessionsCountOk returns a tuple with the TotalSessionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSessionsCount

`func (o *ActiveSessionsCountParams) SetTotalSessionsCount(v int64)`

SetTotalSessionsCount sets TotalSessionsCount field to given value.

### HasTotalSessionsCount

`func (o *ActiveSessionsCountParams) HasTotalSessionsCount() bool`

HasTotalSessionsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


