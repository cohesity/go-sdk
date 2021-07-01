# UpdateNotifications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **NullableString** | Specifies the operation to be performed on the resource. Eg. \&quot;action\&quot;:\&quot;dismiss\&quot; | [optional] 
**NotificationIds** | Pointer to **[]string** | Specifies the list of NotificationIds to be operated upon. | [optional] 

## Methods

### NewUpdateNotifications

`func NewUpdateNotifications() *UpdateNotifications`

NewUpdateNotifications instantiates a new UpdateNotifications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNotificationsWithDefaults

`func NewUpdateNotificationsWithDefaults() *UpdateNotifications`

NewUpdateNotificationsWithDefaults instantiates a new UpdateNotifications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *UpdateNotifications) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UpdateNotifications) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UpdateNotifications) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *UpdateNotifications) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *UpdateNotifications) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *UpdateNotifications) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetNotificationIds

`func (o *UpdateNotifications) GetNotificationIds() []string`

GetNotificationIds returns the NotificationIds field if non-nil, zero value otherwise.

### GetNotificationIdsOk

`func (o *UpdateNotifications) GetNotificationIdsOk() (*[]string, bool)`

GetNotificationIdsOk returns a tuple with the NotificationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationIds

`func (o *UpdateNotifications) SetNotificationIds(v []string)`

SetNotificationIds sets NotificationIds field to given value.

### HasNotificationIds

`func (o *UpdateNotifications) HasNotificationIds() bool`

HasNotificationIds returns a boolean if a field has been set.

### SetNotificationIdsNil

`func (o *UpdateNotifications) SetNotificationIdsNil(b bool)`

 SetNotificationIdsNil sets the value for NotificationIds to be an explicit nil

### UnsetNotificationIds
`func (o *UpdateNotifications) UnsetNotificationIds()`

UnsetNotificationIds ensures that no value is present for NotificationIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


