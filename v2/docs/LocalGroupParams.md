# LocalGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserSids** | Pointer to **[]string** | Specifies the LOCAL users which are part of this group. | [optional] 
**Usernames** | Pointer to **[]string** | Specifies the usernames of the LOCAL users which are part of this group. | [optional] [readonly] 

## Methods

### NewLocalGroupParams

`func NewLocalGroupParams() *LocalGroupParams`

NewLocalGroupParams instantiates a new LocalGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalGroupParamsWithDefaults

`func NewLocalGroupParamsWithDefaults() *LocalGroupParams`

NewLocalGroupParamsWithDefaults instantiates a new LocalGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserSids

`func (o *LocalGroupParams) GetUserSids() []string`

GetUserSids returns the UserSids field if non-nil, zero value otherwise.

### GetUserSidsOk

`func (o *LocalGroupParams) GetUserSidsOk() (*[]string, bool)`

GetUserSidsOk returns a tuple with the UserSids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSids

`func (o *LocalGroupParams) SetUserSids(v []string)`

SetUserSids sets UserSids field to given value.

### HasUserSids

`func (o *LocalGroupParams) HasUserSids() bool`

HasUserSids returns a boolean if a field has been set.

### GetUsernames

`func (o *LocalGroupParams) GetUsernames() []string`

GetUsernames returns the Usernames field if non-nil, zero value otherwise.

### GetUsernamesOk

`func (o *LocalGroupParams) GetUsernamesOk() (*[]string, bool)`

GetUsernamesOk returns a tuple with the Usernames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernames

`func (o *LocalGroupParams) SetUsernames(v []string)`

SetUsernames sets Usernames field to given value.

### HasUsernames

`func (o *LocalGroupParams) HasUsernames() bool`

HasUsernames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


