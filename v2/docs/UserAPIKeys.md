# UserAPIKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeys** | Pointer to [**[]UserAPIKey**](UserAPIKey.md) | List of user owned API Keys. | [optional] 

## Methods

### NewUserAPIKeys

`func NewUserAPIKeys() *UserAPIKeys`

NewUserAPIKeys instantiates a new UserAPIKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAPIKeysWithDefaults

`func NewUserAPIKeysWithDefaults() *UserAPIKeys`

NewUserAPIKeysWithDefaults instantiates a new UserAPIKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeys

`func (o *UserAPIKeys) GetApiKeys() []UserAPIKey`

GetApiKeys returns the ApiKeys field if non-nil, zero value otherwise.

### GetApiKeysOk

`func (o *UserAPIKeys) GetApiKeysOk() (*[]UserAPIKey, bool)`

GetApiKeysOk returns a tuple with the ApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeys

`func (o *UserAPIKeys) SetApiKeys(v []UserAPIKey)`

SetApiKeys sets ApiKeys field to given value.

### HasApiKeys

`func (o *UserAPIKeys) HasApiKeys() bool`

HasApiKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


