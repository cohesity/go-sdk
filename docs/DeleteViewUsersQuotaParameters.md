# DeleteViewUsersQuotaParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteAll** | Pointer to **NullableBool** | Delete all existing user quota override policies. | [optional] 
**UserIds** | Pointer to [**[]UserId**](UserId.md) | The user ids whose policy needs to be deleted. | [optional] 
**ViewName** | Pointer to **NullableString** | View name of input view. | [optional] 

## Methods

### NewDeleteViewUsersQuotaParameters

`func NewDeleteViewUsersQuotaParameters() *DeleteViewUsersQuotaParameters`

NewDeleteViewUsersQuotaParameters instantiates a new DeleteViewUsersQuotaParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteViewUsersQuotaParametersWithDefaults

`func NewDeleteViewUsersQuotaParametersWithDefaults() *DeleteViewUsersQuotaParameters`

NewDeleteViewUsersQuotaParametersWithDefaults instantiates a new DeleteViewUsersQuotaParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteAll

`func (o *DeleteViewUsersQuotaParameters) GetDeleteAll() bool`

GetDeleteAll returns the DeleteAll field if non-nil, zero value otherwise.

### GetDeleteAllOk

`func (o *DeleteViewUsersQuotaParameters) GetDeleteAllOk() (*bool, bool)`

GetDeleteAllOk returns a tuple with the DeleteAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAll

`func (o *DeleteViewUsersQuotaParameters) SetDeleteAll(v bool)`

SetDeleteAll sets DeleteAll field to given value.

### HasDeleteAll

`func (o *DeleteViewUsersQuotaParameters) HasDeleteAll() bool`

HasDeleteAll returns a boolean if a field has been set.

### SetDeleteAllNil

`func (o *DeleteViewUsersQuotaParameters) SetDeleteAllNil(b bool)`

 SetDeleteAllNil sets the value for DeleteAll to be an explicit nil

### UnsetDeleteAll
`func (o *DeleteViewUsersQuotaParameters) UnsetDeleteAll()`

UnsetDeleteAll ensures that no value is present for DeleteAll, not even an explicit nil
### GetUserIds

`func (o *DeleteViewUsersQuotaParameters) GetUserIds() []UserId`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *DeleteViewUsersQuotaParameters) GetUserIdsOk() (*[]UserId, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *DeleteViewUsersQuotaParameters) SetUserIds(v []UserId)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *DeleteViewUsersQuotaParameters) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### SetUserIdsNil

`func (o *DeleteViewUsersQuotaParameters) SetUserIdsNil(b bool)`

 SetUserIdsNil sets the value for UserIds to be an explicit nil

### UnsetUserIds
`func (o *DeleteViewUsersQuotaParameters) UnsetUserIds()`

UnsetUserIds ensures that no value is present for UserIds, not even an explicit nil
### GetViewName

`func (o *DeleteViewUsersQuotaParameters) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *DeleteViewUsersQuotaParameters) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *DeleteViewUsersQuotaParameters) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *DeleteViewUsersQuotaParameters) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *DeleteViewUsersQuotaParameters) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *DeleteViewUsersQuotaParameters) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


