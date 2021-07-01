# GoogleAccountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **NullableString** | Specifies the Account Id assigned by Google. | [optional] 
**UserId** | Pointer to **NullableString** | Specifies the User Id assigned by Google. | [optional] 

## Methods

### NewGoogleAccountInfo

`func NewGoogleAccountInfo() *GoogleAccountInfo`

NewGoogleAccountInfo instantiates a new GoogleAccountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleAccountInfoWithDefaults

`func NewGoogleAccountInfoWithDefaults() *GoogleAccountInfo`

NewGoogleAccountInfoWithDefaults instantiates a new GoogleAccountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *GoogleAccountInfo) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GoogleAccountInfo) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GoogleAccountInfo) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GoogleAccountInfo) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *GoogleAccountInfo) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *GoogleAccountInfo) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetUserId

`func (o *GoogleAccountInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GoogleAccountInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GoogleAccountInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GoogleAccountInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *GoogleAccountInfo) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *GoogleAccountInfo) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


