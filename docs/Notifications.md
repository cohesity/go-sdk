# Notifications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt64** | Notification Count. | [optional] 
**NotificationList** | Pointer to [**[]TaskNotification**](TaskNotification.md) | Notification list. | [optional] 
**UnreadCount** | Pointer to **NullableInt64** | Unread Notification Count. | [optional] 

## Methods

### NewNotifications

`func NewNotifications() *Notifications`

NewNotifications instantiates a new Notifications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsWithDefaults

`func NewNotificationsWithDefaults() *Notifications`

NewNotificationsWithDefaults instantiates a new Notifications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *Notifications) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Notifications) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Notifications) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *Notifications) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *Notifications) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *Notifications) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetNotificationList

`func (o *Notifications) GetNotificationList() []TaskNotification`

GetNotificationList returns the NotificationList field if non-nil, zero value otherwise.

### GetNotificationListOk

`func (o *Notifications) GetNotificationListOk() (*[]TaskNotification, bool)`

GetNotificationListOk returns a tuple with the NotificationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationList

`func (o *Notifications) SetNotificationList(v []TaskNotification)`

SetNotificationList sets NotificationList field to given value.

### HasNotificationList

`func (o *Notifications) HasNotificationList() bool`

HasNotificationList returns a boolean if a field has been set.

### SetNotificationListNil

`func (o *Notifications) SetNotificationListNil(b bool)`

 SetNotificationListNil sets the value for NotificationList to be an explicit nil

### UnsetNotificationList
`func (o *Notifications) UnsetNotificationList()`

UnsetNotificationList ensures that no value is present for NotificationList, not even an explicit nil
### GetUnreadCount

`func (o *Notifications) GetUnreadCount() int64`

GetUnreadCount returns the UnreadCount field if non-nil, zero value otherwise.

### GetUnreadCountOk

`func (o *Notifications) GetUnreadCountOk() (*int64, bool)`

GetUnreadCountOk returns a tuple with the UnreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreadCount

`func (o *Notifications) SetUnreadCount(v int64)`

SetUnreadCount sets UnreadCount field to given value.

### HasUnreadCount

`func (o *Notifications) HasUnreadCount() bool`

HasUnreadCount returns a boolean if a field has been set.

### SetUnreadCountNil

`func (o *Notifications) SetUnreadCountNil(b bool)`

 SetUnreadCountNil sets the value for UnreadCount to be an explicit nil

### UnsetUnreadCount
`func (o *Notifications) UnsetUnreadCount()`

UnsetUnreadCount ensures that no value is present for UnreadCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


