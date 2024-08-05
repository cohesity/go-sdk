# TenantMigrationServiceAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **NullableString** | Specifies the action which will be performed on the tenant. | [optional] 
**ActionIncarnationId** | Pointer to **NullableInt32** | Retry count for the action. If an action needs to be retried, then clients will increment action_incarnation_id and can send the same request again | [optional] [default to 1]
**ErrorMessage** | Pointer to **NullableString** | Specifies the error message if an error occurred. | [optional] 
**ErrorType** | Pointer to **NullableString** | Specifies the type of error which occurred. | [optional] 
**Status** | Pointer to **NullableString** | Action status | [optional] 

## Methods

### NewTenantMigrationServiceAction

`func NewTenantMigrationServiceAction() *TenantMigrationServiceAction`

NewTenantMigrationServiceAction instantiates a new TenantMigrationServiceAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantMigrationServiceActionWithDefaults

`func NewTenantMigrationServiceActionWithDefaults() *TenantMigrationServiceAction`

NewTenantMigrationServiceActionWithDefaults instantiates a new TenantMigrationServiceAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *TenantMigrationServiceAction) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TenantMigrationServiceAction) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TenantMigrationServiceAction) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *TenantMigrationServiceAction) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *TenantMigrationServiceAction) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *TenantMigrationServiceAction) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetActionIncarnationId

`func (o *TenantMigrationServiceAction) GetActionIncarnationId() int32`

GetActionIncarnationId returns the ActionIncarnationId field if non-nil, zero value otherwise.

### GetActionIncarnationIdOk

`func (o *TenantMigrationServiceAction) GetActionIncarnationIdOk() (*int32, bool)`

GetActionIncarnationIdOk returns a tuple with the ActionIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionIncarnationId

`func (o *TenantMigrationServiceAction) SetActionIncarnationId(v int32)`

SetActionIncarnationId sets ActionIncarnationId field to given value.

### HasActionIncarnationId

`func (o *TenantMigrationServiceAction) HasActionIncarnationId() bool`

HasActionIncarnationId returns a boolean if a field has been set.

### SetActionIncarnationIdNil

`func (o *TenantMigrationServiceAction) SetActionIncarnationIdNil(b bool)`

 SetActionIncarnationIdNil sets the value for ActionIncarnationId to be an explicit nil

### UnsetActionIncarnationId
`func (o *TenantMigrationServiceAction) UnsetActionIncarnationId()`

UnsetActionIncarnationId ensures that no value is present for ActionIncarnationId, not even an explicit nil
### GetErrorMessage

`func (o *TenantMigrationServiceAction) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TenantMigrationServiceAction) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TenantMigrationServiceAction) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *TenantMigrationServiceAction) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *TenantMigrationServiceAction) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *TenantMigrationServiceAction) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetErrorType

`func (o *TenantMigrationServiceAction) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *TenantMigrationServiceAction) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *TenantMigrationServiceAction) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *TenantMigrationServiceAction) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.

### SetErrorTypeNil

`func (o *TenantMigrationServiceAction) SetErrorTypeNil(b bool)`

 SetErrorTypeNil sets the value for ErrorType to be an explicit nil

### UnsetErrorType
`func (o *TenantMigrationServiceAction) UnsetErrorType()`

UnsetErrorType ensures that no value is present for ErrorType, not even an explicit nil
### GetStatus

`func (o *TenantMigrationServiceAction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TenantMigrationServiceAction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TenantMigrationServiceAction) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TenantMigrationServiceAction) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TenantMigrationServiceAction) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TenantMigrationServiceAction) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


