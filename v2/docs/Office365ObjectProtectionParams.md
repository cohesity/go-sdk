# Office365ObjectProtectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupsObjectProtectionParams** | Pointer to [**Office365GroupsObjectProtectionParams**](Office365GroupsObjectProtectionParams.md) |  | [optional] 
**ObjectProtectionType** | **string** | Specifies the Microsoft 365 Object Protection type. | 
**SharepointSiteObjectProtectionParams** | Pointer to [**Office365SharepointSiteObjectProtectionParams**](Office365SharepointSiteObjectProtectionParams.md) |  | [optional] 
**TeamsObjectProtectionParams** | Pointer to [**Office365TeamsObjectProtectionParams**](Office365TeamsObjectProtectionParams.md) |  | [optional] 
**UserMailboxObjectProtectionParams** | Pointer to [**Office365UserMailboxObjectProtectionParams**](Office365UserMailboxObjectProtectionParams.md) |  | [optional] 
**UserOneDriveObjectProtectionParams** | Pointer to [**Office365UserOneDriveObjectProtectionParams**](Office365UserOneDriveObjectProtectionParams.md) |  | [optional] 

## Methods

### NewOffice365ObjectProtectionParams

`func NewOffice365ObjectProtectionParams(objectProtectionType string, ) *Office365ObjectProtectionParams`

NewOffice365ObjectProtectionParams instantiates a new Office365ObjectProtectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffice365ObjectProtectionParamsWithDefaults

`func NewOffice365ObjectProtectionParamsWithDefaults() *Office365ObjectProtectionParams`

NewOffice365ObjectProtectionParamsWithDefaults instantiates a new Office365ObjectProtectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupsObjectProtectionParams

`func (o *Office365ObjectProtectionParams) GetGroupsObjectProtectionParams() Office365GroupsObjectProtectionParams`

GetGroupsObjectProtectionParams returns the GroupsObjectProtectionParams field if non-nil, zero value otherwise.

### GetGroupsObjectProtectionParamsOk

`func (o *Office365ObjectProtectionParams) GetGroupsObjectProtectionParamsOk() (*Office365GroupsObjectProtectionParams, bool)`

GetGroupsObjectProtectionParamsOk returns a tuple with the GroupsObjectProtectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsObjectProtectionParams

`func (o *Office365ObjectProtectionParams) SetGroupsObjectProtectionParams(v Office365GroupsObjectProtectionParams)`

SetGroupsObjectProtectionParams sets GroupsObjectProtectionParams field to given value.

### HasGroupsObjectProtectionParams

`func (o *Office365ObjectProtectionParams) HasGroupsObjectProtectionParams() bool`

HasGroupsObjectProtectionParams returns a boolean if a field has been set.

### GetObjectProtectionType

`func (o *Office365ObjectProtectionParams) GetObjectProtectionType() string`

GetObjectProtectionType returns the ObjectProtectionType field if non-nil, zero value otherwise.

### GetObjectProtectionTypeOk

`func (o *Office365ObjectProtectionParams) GetObjectProtectionTypeOk() (*string, bool)`

GetObjectProtectionTypeOk returns a tuple with the ObjectProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectProtectionType

`func (o *Office365ObjectProtectionParams) SetObjectProtectionType(v string)`

SetObjectProtectionType sets ObjectProtectionType field to given value.


### GetSharepointSiteObjectProtectionParams

`func (o *Office365ObjectProtectionParams) GetSharepointSiteObjectProtectionParams() Office365SharepointSiteObjectProtectionParams`

GetSharepointSiteObjectProtectionParams returns the SharepointSiteObjectProtectionParams field if non-nil, zero value otherwise.

### GetSharepointSiteObjectProtectionParamsOk

`func (o *Office365ObjectProtectionParams) GetSharepointSiteObjectProtectionParamsOk() (*Office365SharepointSiteObjectProtectionParams, bool)`

GetSharepointSiteObjectProtectionParamsOk returns a tuple with the SharepointSiteObjectProtectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharepointSiteObjectProtectionParams

`func (o *Office365ObjectProtectionParams) SetSharepointSiteObjectProtectionParams(v Office365SharepointSiteObjectProtectionParams)`

SetSharepointSiteObjectProtectionParams sets SharepointSiteObjectProtectionParams field to given value.

### HasSharepointSiteObjectProtectionParams

`func (o *Office365ObjectProtectionParams) HasSharepointSiteObjectProtectionParams() bool`

HasSharepointSiteObjectProtectionParams returns a boolean if a field has been set.

### GetTeamsObjectProtectionParams

`func (o *Office365ObjectProtectionParams) GetTeamsObjectProtectionParams() Office365TeamsObjectProtectionParams`

GetTeamsObjectProtectionParams returns the TeamsObjectProtectionParams field if non-nil, zero value otherwise.

### GetTeamsObjectProtectionParamsOk

`func (o *Office365ObjectProtectionParams) GetTeamsObjectProtectionParamsOk() (*Office365TeamsObjectProtectionParams, bool)`

GetTeamsObjectProtectionParamsOk returns a tuple with the TeamsObjectProtectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsObjectProtectionParams

`func (o *Office365ObjectProtectionParams) SetTeamsObjectProtectionParams(v Office365TeamsObjectProtectionParams)`

SetTeamsObjectProtectionParams sets TeamsObjectProtectionParams field to given value.

### HasTeamsObjectProtectionParams

`func (o *Office365ObjectProtectionParams) HasTeamsObjectProtectionParams() bool`

HasTeamsObjectProtectionParams returns a boolean if a field has been set.

### GetUserMailboxObjectProtectionParams

`func (o *Office365ObjectProtectionParams) GetUserMailboxObjectProtectionParams() Office365UserMailboxObjectProtectionParams`

GetUserMailboxObjectProtectionParams returns the UserMailboxObjectProtectionParams field if non-nil, zero value otherwise.

### GetUserMailboxObjectProtectionParamsOk

`func (o *Office365ObjectProtectionParams) GetUserMailboxObjectProtectionParamsOk() (*Office365UserMailboxObjectProtectionParams, bool)`

GetUserMailboxObjectProtectionParamsOk returns a tuple with the UserMailboxObjectProtectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMailboxObjectProtectionParams

`func (o *Office365ObjectProtectionParams) SetUserMailboxObjectProtectionParams(v Office365UserMailboxObjectProtectionParams)`

SetUserMailboxObjectProtectionParams sets UserMailboxObjectProtectionParams field to given value.

### HasUserMailboxObjectProtectionParams

`func (o *Office365ObjectProtectionParams) HasUserMailboxObjectProtectionParams() bool`

HasUserMailboxObjectProtectionParams returns a boolean if a field has been set.

### GetUserOneDriveObjectProtectionParams

`func (o *Office365ObjectProtectionParams) GetUserOneDriveObjectProtectionParams() Office365UserOneDriveObjectProtectionParams`

GetUserOneDriveObjectProtectionParams returns the UserOneDriveObjectProtectionParams field if non-nil, zero value otherwise.

### GetUserOneDriveObjectProtectionParamsOk

`func (o *Office365ObjectProtectionParams) GetUserOneDriveObjectProtectionParamsOk() (*Office365UserOneDriveObjectProtectionParams, bool)`

GetUserOneDriveObjectProtectionParamsOk returns a tuple with the UserOneDriveObjectProtectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserOneDriveObjectProtectionParams

`func (o *Office365ObjectProtectionParams) SetUserOneDriveObjectProtectionParams(v Office365UserOneDriveObjectProtectionParams)`

SetUserOneDriveObjectProtectionParams sets UserOneDriveObjectProtectionParams field to given value.

### HasUserOneDriveObjectProtectionParams

`func (o *Office365ObjectProtectionParams) HasUserOneDriveObjectProtectionParams() bool`

HasUserOneDriveObjectProtectionParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


