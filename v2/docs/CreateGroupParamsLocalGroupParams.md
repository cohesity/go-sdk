# CreateGroupParamsLocalGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserSids** | Pointer to **[]string** | Specifies the LOCAL users which are part of this group. | [optional] 
**Usernames** | Pointer to **[]string** | Specifies the usernames of the LOCAL users which are part of this group. | [optional] [readonly] 

## Methods

### NewCreateGroupParamsLocalGroupParams

`func NewCreateGroupParamsLocalGroupParams() *CreateGroupParamsLocalGroupParams`

NewCreateGroupParamsLocalGroupParams instantiates a new CreateGroupParamsLocalGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupParamsLocalGroupParamsWithDefaults

`func NewCreateGroupParamsLocalGroupParamsWithDefaults() *CreateGroupParamsLocalGroupParams`

NewCreateGroupParamsLocalGroupParamsWithDefaults instantiates a new CreateGroupParamsLocalGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserSids

`func (o *CreateGroupParamsLocalGroupParams) GetUserSids() []string`

GetUserSids returns the UserSids field if non-nil, zero value otherwise.

### GetUserSidsOk

`func (o *CreateGroupParamsLocalGroupParams) GetUserSidsOk() (*[]string, bool)`

GetUserSidsOk returns a tuple with the UserSids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSids

`func (o *CreateGroupParamsLocalGroupParams) SetUserSids(v []string)`

SetUserSids sets UserSids field to given value.

### HasUserSids

`func (o *CreateGroupParamsLocalGroupParams) HasUserSids() bool`

HasUserSids returns a boolean if a field has been set.

### GetUsernames

`func (o *CreateGroupParamsLocalGroupParams) GetUsernames() []string`

GetUsernames returns the Usernames field if non-nil, zero value otherwise.

### GetUsernamesOk

`func (o *CreateGroupParamsLocalGroupParams) GetUsernamesOk() (*[]string, bool)`

GetUsernamesOk returns a tuple with the Usernames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernames

`func (o *CreateGroupParamsLocalGroupParams) SetUsernames(v []string)`

SetUsernames sets Usernames field to given value.

### HasUsernames

`func (o *CreateGroupParamsLocalGroupParams) HasUsernames() bool`

HasUsernames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


