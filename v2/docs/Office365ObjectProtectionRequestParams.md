# Office365ObjectProtectionRequestParams

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

### NewOffice365ObjectProtectionRequestParams

`func NewOffice365ObjectProtectionRequestParams(objectProtectionType string, ) *Office365ObjectProtectionRequestParams`

NewOffice365ObjectProtectionRequestParams instantiates a new Office365ObjectProtectionRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffice365ObjectProtectionRequestParamsWithDefaults

`func NewOffice365ObjectProtectionRequestParamsWithDefaults() *Office365ObjectProtectionRequestParams`

NewOffice365ObjectProtectionRequestParamsWithDefaults instantiates a new Office365ObjectProtectionRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupsObjectProtectionParams

`func (o *Office365ObjectProtectionRequestParams) GetGroupsObjectProtectionParams() Office365GroupsObjectProtectionParams`

GetGroupsObjectProtectionParams returns the GroupsObjectProtectionParams field if non-nil, zero value otherwise.

### GetGroupsObjectProtectionParamsOk

`func (o *Office365ObjectProtectionRequestParams) GetGroupsObjectProtectionParamsOk() (*Office365GroupsObjectProtectionParams, bool)`

GetGroupsObjectProtectionParamsOk returns a tuple with the GroupsObjectProtectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsObjectProtectionParams

`func (o *Office365ObjectProtectionRequestParams) SetGroupsObjectProtectionParams(v Office365GroupsObjectProtectionParams)`

SetGroupsObjectProtectionParams sets GroupsObjectProtectionParams field to given value.

### HasGroupsObjectProtectionParams

`func (o *Office365ObjectProtectionRequestParams) HasGroupsObjectProtectionParams() bool`

HasGroupsObjectProtectionParams returns a boolean if a field has been set.

### GetObjectProtectionType

`func (o *Office365ObjectProtectionRequestParams) GetObjectProtectionType() string`

GetObjectProtectionType returns the ObjectProtectionType field if non-nil, zero value otherwise.

### GetObjectProtectionTypeOk

`func (o *Office365ObjectProtectionRequestParams) GetObjectProtectionTypeOk() (*string, bool)`

GetObjectProtectionTypeOk returns a tuple with the ObjectProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectProtectionType

`func (o *Office365ObjectProtectionRequestParams) SetObjectProtectionType(v string)`

SetObjectProtectionType sets ObjectProtectionType field to given value.


### GetSharepointSiteObjectProtectionParams

`func (o *Office365ObjectProtectionRequestParams) GetSharepointSiteObjectProtectionParams() Office365SharepointSiteObjectProtectionParams`

GetSharepointSiteObjectProtectionParams returns the SharepointSiteObjectProtectionParams field if non-nil, zero value otherwise.

### GetSharepointSiteObjectProtectionParamsOk

`func (o *Office365ObjectProtectionRequestParams) GetSharepointSiteObjectProtectionParamsOk() (*Office365SharepointSiteObjectProtectionParams, bool)`

GetSharepointSiteObjectProtectionParamsOk returns a tuple with the SharepointSiteObjectProtectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharepointSiteObjectProtectionParams

`func (o *Office365ObjectProtectionRequestParams) SetSharepointSiteObjectProtectionParams(v Office365SharepointSiteObjectProtectionParams)`

SetSharepointSiteObjectProtectionParams sets SharepointSiteObjectProtectionParams field to given value.

### HasSharepointSiteObjectProtectionParams

`func (o *Office365ObjectProtectionRequestParams) HasSharepointSiteObjectProtectionParams() bool`

HasSharepointSiteObjectProtectionParams returns a boolean if a field has been set.

### GetTeamsObjectProtectionParams

`func (o *Office365ObjectProtectionRequestParams) GetTeamsObjectProtectionParams() Office365TeamsObjectProtectionParams`

GetTeamsObjectProtectionParams returns the TeamsObjectProtectionParams field if non-nil, zero value otherwise.

### GetTeamsObjectProtectionParamsOk

`func (o *Office365ObjectProtectionRequestParams) GetTeamsObjectProtectionParamsOk() (*Office365TeamsObjectProtectionParams, bool)`

GetTeamsObjectProtectionParamsOk returns a tuple with the TeamsObjectProtectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsObjectProtectionParams

`func (o *Office365ObjectProtectionRequestParams) SetTeamsObjectProtectionParams(v Office365TeamsObjectProtectionParams)`

SetTeamsObjectProtectionParams sets TeamsObjectProtectionParams field to given value.

### HasTeamsObjectProtectionParams

`func (o *Office365ObjectProtectionRequestParams) HasTeamsObjectProtectionParams() bool`

HasTeamsObjectProtectionParams returns a boolean if a field has been set.

### GetUserMailboxObjectProtectionParams

`func (o *Office365ObjectProtectionRequestParams) GetUserMailboxObjectProtectionParams() Office365UserMailboxObjectProtectionParams`

GetUserMailboxObjectProtectionParams returns the UserMailboxObjectProtectionParams field if non-nil, zero value otherwise.

### GetUserMailboxObjectProtectionParamsOk

`func (o *Office365ObjectProtectionRequestParams) GetUserMailboxObjectProtectionParamsOk() (*Office365UserMailboxObjectProtectionParams, bool)`

GetUserMailboxObjectProtectionParamsOk returns a tuple with the UserMailboxObjectProtectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMailboxObjectProtectionParams

`func (o *Office365ObjectProtectionRequestParams) SetUserMailboxObjectProtectionParams(v Office365UserMailboxObjectProtectionParams)`

SetUserMailboxObjectProtectionParams sets UserMailboxObjectProtectionParams field to given value.

### HasUserMailboxObjectProtectionParams

`func (o *Office365ObjectProtectionRequestParams) HasUserMailboxObjectProtectionParams() bool`

HasUserMailboxObjectProtectionParams returns a boolean if a field has been set.

### GetUserOneDriveObjectProtectionParams

`func (o *Office365ObjectProtectionRequestParams) GetUserOneDriveObjectProtectionParams() Office365UserOneDriveObjectProtectionParams`

GetUserOneDriveObjectProtectionParams returns the UserOneDriveObjectProtectionParams field if non-nil, zero value otherwise.

### GetUserOneDriveObjectProtectionParamsOk

`func (o *Office365ObjectProtectionRequestParams) GetUserOneDriveObjectProtectionParamsOk() (*Office365UserOneDriveObjectProtectionParams, bool)`

GetUserOneDriveObjectProtectionParamsOk returns a tuple with the UserOneDriveObjectProtectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserOneDriveObjectProtectionParams

`func (o *Office365ObjectProtectionRequestParams) SetUserOneDriveObjectProtectionParams(v Office365UserOneDriveObjectProtectionParams)`

SetUserOneDriveObjectProtectionParams sets UserOneDriveObjectProtectionParams field to given value.

### HasUserOneDriveObjectProtectionParams

`func (o *Office365ObjectProtectionRequestParams) HasUserOneDriveObjectProtectionParams() bool`

HasUserOneDriveObjectProtectionParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


