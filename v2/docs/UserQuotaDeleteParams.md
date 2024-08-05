# UserQuotaDeleteParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserIds** | Pointer to [**[]UserId**](UserId.md) | Array of userIds. Specifies the list of user Ids to delete logical user quota override. | [optional] 

## Methods

### NewUserQuotaDeleteParams

`func NewUserQuotaDeleteParams() *UserQuotaDeleteParams`

NewUserQuotaDeleteParams instantiates a new UserQuotaDeleteParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserQuotaDeleteParamsWithDefaults

`func NewUserQuotaDeleteParamsWithDefaults() *UserQuotaDeleteParams`

NewUserQuotaDeleteParamsWithDefaults instantiates a new UserQuotaDeleteParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserIds

`func (o *UserQuotaDeleteParams) GetUserIds() []UserId`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *UserQuotaDeleteParams) GetUserIdsOk() (*[]UserId, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *UserQuotaDeleteParams) SetUserIds(v []UserId)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *UserQuotaDeleteParams) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### SetUserIdsNil

`func (o *UserQuotaDeleteParams) SetUserIdsNil(b bool)`

 SetUserIdsNil sets the value for UserIds to be an explicit nil

### UnsetUserIds
`func (o *UserQuotaDeleteParams) UnsetUserIds()`

UnsetUserIds ensures that no value is present for UserIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


