# IbmIdProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntrospectCacheTimeMins** | Pointer to **NullableInt64** | Specifies the number of minutes the cluster will wait before reperforming token introspection on a token it has already seen. Default value is 15 minutes. Maximum is 24 hours. | [optional] [default to 15]
**OperatorTokenCacheTimeMins** | Pointer to **NullableInt64** | Specifies the number of minutes the cluster will wait before polling for a new operator token. Default value is 15 minutes. Maximum is 24 hours. | [optional] [default to 15]
**RoleCacheTimeMins** | Pointer to **NullableInt64** | Specifies the number of minutes the cluster will wait before fetching roles again for a user it has already seen. Default value is 15 minutes. Maximum is 24 hours. | [optional] [default to 15]

## Methods

### NewIbmIdProvider

`func NewIbmIdProvider() *IbmIdProvider`

NewIbmIdProvider instantiates a new IbmIdProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIbmIdProviderWithDefaults

`func NewIbmIdProviderWithDefaults() *IbmIdProvider`

NewIbmIdProviderWithDefaults instantiates a new IbmIdProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntrospectCacheTimeMins

`func (o *IbmIdProvider) GetIntrospectCacheTimeMins() int64`

GetIntrospectCacheTimeMins returns the IntrospectCacheTimeMins field if non-nil, zero value otherwise.

### GetIntrospectCacheTimeMinsOk

`func (o *IbmIdProvider) GetIntrospectCacheTimeMinsOk() (*int64, bool)`

GetIntrospectCacheTimeMinsOk returns a tuple with the IntrospectCacheTimeMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntrospectCacheTimeMins

`func (o *IbmIdProvider) SetIntrospectCacheTimeMins(v int64)`

SetIntrospectCacheTimeMins sets IntrospectCacheTimeMins field to given value.

### HasIntrospectCacheTimeMins

`func (o *IbmIdProvider) HasIntrospectCacheTimeMins() bool`

HasIntrospectCacheTimeMins returns a boolean if a field has been set.

### SetIntrospectCacheTimeMinsNil

`func (o *IbmIdProvider) SetIntrospectCacheTimeMinsNil(b bool)`

 SetIntrospectCacheTimeMinsNil sets the value for IntrospectCacheTimeMins to be an explicit nil

### UnsetIntrospectCacheTimeMins
`func (o *IbmIdProvider) UnsetIntrospectCacheTimeMins()`

UnsetIntrospectCacheTimeMins ensures that no value is present for IntrospectCacheTimeMins, not even an explicit nil
### GetOperatorTokenCacheTimeMins

`func (o *IbmIdProvider) GetOperatorTokenCacheTimeMins() int64`

GetOperatorTokenCacheTimeMins returns the OperatorTokenCacheTimeMins field if non-nil, zero value otherwise.

### GetOperatorTokenCacheTimeMinsOk

`func (o *IbmIdProvider) GetOperatorTokenCacheTimeMinsOk() (*int64, bool)`

GetOperatorTokenCacheTimeMinsOk returns a tuple with the OperatorTokenCacheTimeMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorTokenCacheTimeMins

`func (o *IbmIdProvider) SetOperatorTokenCacheTimeMins(v int64)`

SetOperatorTokenCacheTimeMins sets OperatorTokenCacheTimeMins field to given value.

### HasOperatorTokenCacheTimeMins

`func (o *IbmIdProvider) HasOperatorTokenCacheTimeMins() bool`

HasOperatorTokenCacheTimeMins returns a boolean if a field has been set.

### SetOperatorTokenCacheTimeMinsNil

`func (o *IbmIdProvider) SetOperatorTokenCacheTimeMinsNil(b bool)`

 SetOperatorTokenCacheTimeMinsNil sets the value for OperatorTokenCacheTimeMins to be an explicit nil

### UnsetOperatorTokenCacheTimeMins
`func (o *IbmIdProvider) UnsetOperatorTokenCacheTimeMins()`

UnsetOperatorTokenCacheTimeMins ensures that no value is present for OperatorTokenCacheTimeMins, not even an explicit nil
### GetRoleCacheTimeMins

`func (o *IbmIdProvider) GetRoleCacheTimeMins() int64`

GetRoleCacheTimeMins returns the RoleCacheTimeMins field if non-nil, zero value otherwise.

### GetRoleCacheTimeMinsOk

`func (o *IbmIdProvider) GetRoleCacheTimeMinsOk() (*int64, bool)`

GetRoleCacheTimeMinsOk returns a tuple with the RoleCacheTimeMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleCacheTimeMins

`func (o *IbmIdProvider) SetRoleCacheTimeMins(v int64)`

SetRoleCacheTimeMins sets RoleCacheTimeMins field to given value.

### HasRoleCacheTimeMins

`func (o *IbmIdProvider) HasRoleCacheTimeMins() bool`

HasRoleCacheTimeMins returns a boolean if a field has been set.

### SetRoleCacheTimeMinsNil

`func (o *IbmIdProvider) SetRoleCacheTimeMinsNil(b bool)`

 SetRoleCacheTimeMinsNil sets the value for RoleCacheTimeMins to be an explicit nil

### UnsetRoleCacheTimeMins
`func (o *IbmIdProvider) UnsetRoleCacheTimeMins()`

UnsetRoleCacheTimeMins ensures that no value is present for RoleCacheTimeMins, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


