# RegistrationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationStatus** | Pointer to **NullableString** | Specifies the status of the authentication during the registration of a Protection Source. &#39;Pending&#39; indicates the authentication is in progress. &#39;Scheduled&#39; indicates the authentication is scheduled. &#39;Finished&#39; indicates the authentication is completed. &#39;RefreshInProgress&#39; indicates the refresh is in progress. | [optional] [readonly] 
**LastRefreshedTimeMsecs** | Pointer to **NullableInt64** | Specifies the time when the source was last refreshed in milliseconds. | [optional] [readonly] 
**RegistrationTimeMsecs** | Pointer to **NullableInt64** | Specifies the time when the source was registered in milliseconds | [optional] [readonly] 

## Methods

### NewRegistrationInfo

`func NewRegistrationInfo() *RegistrationInfo`

NewRegistrationInfo instantiates a new RegistrationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationInfoWithDefaults

`func NewRegistrationInfoWithDefaults() *RegistrationInfo`

NewRegistrationInfoWithDefaults instantiates a new RegistrationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationStatus

`func (o *RegistrationInfo) GetAuthenticationStatus() string`

GetAuthenticationStatus returns the AuthenticationStatus field if non-nil, zero value otherwise.

### GetAuthenticationStatusOk

`func (o *RegistrationInfo) GetAuthenticationStatusOk() (*string, bool)`

GetAuthenticationStatusOk returns a tuple with the AuthenticationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationStatus

`func (o *RegistrationInfo) SetAuthenticationStatus(v string)`

SetAuthenticationStatus sets AuthenticationStatus field to given value.

### HasAuthenticationStatus

`func (o *RegistrationInfo) HasAuthenticationStatus() bool`

HasAuthenticationStatus returns a boolean if a field has been set.

### SetAuthenticationStatusNil

`func (o *RegistrationInfo) SetAuthenticationStatusNil(b bool)`

 SetAuthenticationStatusNil sets the value for AuthenticationStatus to be an explicit nil

### UnsetAuthenticationStatus
`func (o *RegistrationInfo) UnsetAuthenticationStatus()`

UnsetAuthenticationStatus ensures that no value is present for AuthenticationStatus, not even an explicit nil
### GetLastRefreshedTimeMsecs

`func (o *RegistrationInfo) GetLastRefreshedTimeMsecs() int64`

GetLastRefreshedTimeMsecs returns the LastRefreshedTimeMsecs field if non-nil, zero value otherwise.

### GetLastRefreshedTimeMsecsOk

`func (o *RegistrationInfo) GetLastRefreshedTimeMsecsOk() (*int64, bool)`

GetLastRefreshedTimeMsecsOk returns a tuple with the LastRefreshedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRefreshedTimeMsecs

`func (o *RegistrationInfo) SetLastRefreshedTimeMsecs(v int64)`

SetLastRefreshedTimeMsecs sets LastRefreshedTimeMsecs field to given value.

### HasLastRefreshedTimeMsecs

`func (o *RegistrationInfo) HasLastRefreshedTimeMsecs() bool`

HasLastRefreshedTimeMsecs returns a boolean if a field has been set.

### SetLastRefreshedTimeMsecsNil

`func (o *RegistrationInfo) SetLastRefreshedTimeMsecsNil(b bool)`

 SetLastRefreshedTimeMsecsNil sets the value for LastRefreshedTimeMsecs to be an explicit nil

### UnsetLastRefreshedTimeMsecs
`func (o *RegistrationInfo) UnsetLastRefreshedTimeMsecs()`

UnsetLastRefreshedTimeMsecs ensures that no value is present for LastRefreshedTimeMsecs, not even an explicit nil
### GetRegistrationTimeMsecs

`func (o *RegistrationInfo) GetRegistrationTimeMsecs() int64`

GetRegistrationTimeMsecs returns the RegistrationTimeMsecs field if non-nil, zero value otherwise.

### GetRegistrationTimeMsecsOk

`func (o *RegistrationInfo) GetRegistrationTimeMsecsOk() (*int64, bool)`

GetRegistrationTimeMsecsOk returns a tuple with the RegistrationTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationTimeMsecs

`func (o *RegistrationInfo) SetRegistrationTimeMsecs(v int64)`

SetRegistrationTimeMsecs sets RegistrationTimeMsecs field to given value.

### HasRegistrationTimeMsecs

`func (o *RegistrationInfo) HasRegistrationTimeMsecs() bool`

HasRegistrationTimeMsecs returns a boolean if a field has been set.

### SetRegistrationTimeMsecsNil

`func (o *RegistrationInfo) SetRegistrationTimeMsecsNil(b bool)`

 SetRegistrationTimeMsecsNil sets the value for RegistrationTimeMsecs to be an explicit nil

### UnsetRegistrationTimeMsecs
`func (o *RegistrationInfo) UnsetRegistrationTimeMsecs()`

UnsetRegistrationTimeMsecs ensures that no value is present for RegistrationTimeMsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


