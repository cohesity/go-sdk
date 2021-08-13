# Office365ObjectProtectionUpdateRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectProtectionType** | **string** | Specifies the Microsoft 365 Object Protection type. | 
**UserMailboxObjectProtectionParams** | Pointer to [**Office365UserMailboxObjectProtectionParams**](Office365UserMailboxObjectProtectionParams.md) |  | [optional] 
**UserOneDriveObjectProtectionParams** | Pointer to [**Office365UserOneDriveObjectProtectionParams**](Office365UserOneDriveObjectProtectionParams.md) |  | [optional] 
**SharepointSiteObjectProtectionParams** | Pointer to [**Office365SharepointSiteObjectProtectionParams**](Office365SharepointSiteObjectProtectionParams.md) |  | [optional] 
**TeamsObjectProtectionParams** | Pointer to [**Office365TeamsObjectProtectionParams**](Office365TeamsObjectProtectionParams.md) |  | [optional] 

## Methods

### NewOffice365ObjectProtectionUpdateRequestParams

`func NewOffice365ObjectProtectionUpdateRequestParams(objectProtectionType string, ) *Office365ObjectProtectionUpdateRequestParams`

NewOffice365ObjectProtectionUpdateRequestParams instantiates a new Office365ObjectProtectionUpdateRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffice365ObjectProtectionUpdateRequestParamsWithDefaults

`func NewOffice365ObjectProtectionUpdateRequestParamsWithDefaults() *Office365ObjectProtectionUpdateRequestParams`

NewOffice365ObjectProtectionUpdateRequestParamsWithDefaults instantiates a new Office365ObjectProtectionUpdateRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectProtectionType

`func (o *Office365ObjectProtectionUpdateRequestParams) GetObjectProtectionType() string`

GetObjectProtectionType returns the ObjectProtectionType field if non-nil, zero value otherwise.

### GetObjectProtectionTypeOk

`func (o *Office365ObjectProtectionUpdateRequestParams) GetObjectProtectionTypeOk() (*string, bool)`

GetObjectProtectionTypeOk returns a tuple with the ObjectProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectProtectionType

`func (o *Office365ObjectProtectionUpdateRequestParams) SetObjectProtectionType(v string)`

SetObjectProtectionType sets ObjectProtectionType field to given value.


### GetUserMailboxObjectProtectionParams

`func (o *Office365ObjectProtectionUpdateRequestParams) GetUserMailboxObjectProtectionParams() Office365UserMailboxObjectProtectionParams`

GetUserMailboxObjectProtectionParams returns the UserMailboxObjectProtectionParams field if non-nil, zero value otherwise.

### GetUserMailboxObjectProtectionParamsOk

`func (o *Office365ObjectProtectionUpdateRequestParams) GetUserMailboxObjectProtectionParamsOk() (*Office365UserMailboxObjectProtectionParams, bool)`

GetUserMailboxObjectProtectionParamsOk returns a tuple with the UserMailboxObjectProtectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMailboxObjectProtectionParams

`func (o *Office365ObjectProtectionUpdateRequestParams) SetUserMailboxObjectProtectionParams(v Office365UserMailboxObjectProtectionParams)`

SetUserMailboxObjectProtectionParams sets UserMailboxObjectProtectionParams field to given value.

### HasUserMailboxObjectProtectionParams

`func (o *Office365ObjectProtectionUpdateRequestParams) HasUserMailboxObjectProtectionParams() bool`

HasUserMailboxObjectProtectionParams returns a boolean if a field has been set.

### GetUserOneDriveObjectProtectionParams

`func (o *Office365ObjectProtectionUpdateRequestParams) GetUserOneDriveObjectProtectionParams() Office365UserOneDriveObjectProtectionParams`

GetUserOneDriveObjectProtectionParams returns the UserOneDriveObjectProtectionParams field if non-nil, zero value otherwise.

### GetUserOneDriveObjectProtectionParamsOk

`func (o *Office365ObjectProtectionUpdateRequestParams) GetUserOneDriveObjectProtectionParamsOk() (*Office365UserOneDriveObjectProtectionParams, bool)`

GetUserOneDriveObjectProtectionParamsOk returns a tuple with the UserOneDriveObjectProtectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserOneDriveObjectProtectionParams

`func (o *Office365ObjectProtectionUpdateRequestParams) SetUserOneDriveObjectProtectionParams(v Office365UserOneDriveObjectProtectionParams)`

SetUserOneDriveObjectProtectionParams sets UserOneDriveObjectProtectionParams field to given value.

### HasUserOneDriveObjectProtectionParams

`func (o *Office365ObjectProtectionUpdateRequestParams) HasUserOneDriveObjectProtectionParams() bool`

HasUserOneDriveObjectProtectionParams returns a boolean if a field has been set.

### GetSharepointSiteObjectProtectionParams

`func (o *Office365ObjectProtectionUpdateRequestParams) GetSharepointSiteObjectProtectionParams() Office365SharepointSiteObjectProtectionParams`

GetSharepointSiteObjectProtectionParams returns the SharepointSiteObjectProtectionParams field if non-nil, zero value otherwise.

### GetSharepointSiteObjectProtectionParamsOk

`func (o *Office365ObjectProtectionUpdateRequestParams) GetSharepointSiteObjectProtectionParamsOk() (*Office365SharepointSiteObjectProtectionParams, bool)`

GetSharepointSiteObjectProtectionParamsOk returns a tuple with the SharepointSiteObjectProtectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharepointSiteObjectProtectionParams

`func (o *Office365ObjectProtectionUpdateRequestParams) SetSharepointSiteObjectProtectionParams(v Office365SharepointSiteObjectProtectionParams)`

SetSharepointSiteObjectProtectionParams sets SharepointSiteObjectProtectionParams field to given value.

### HasSharepointSiteObjectProtectionParams

`func (o *Office365ObjectProtectionUpdateRequestParams) HasSharepointSiteObjectProtectionParams() bool`

HasSharepointSiteObjectProtectionParams returns a boolean if a field has been set.

### GetTeamsObjectProtectionParams

`func (o *Office365ObjectProtectionUpdateRequestParams) GetTeamsObjectProtectionParams() Office365TeamsObjectProtectionParams`

GetTeamsObjectProtectionParams returns the TeamsObjectProtectionParams field if non-nil, zero value otherwise.

### GetTeamsObjectProtectionParamsOk

`func (o *Office365ObjectProtectionUpdateRequestParams) GetTeamsObjectProtectionParamsOk() (*Office365TeamsObjectProtectionParams, bool)`

GetTeamsObjectProtectionParamsOk returns a tuple with the TeamsObjectProtectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsObjectProtectionParams

`func (o *Office365ObjectProtectionUpdateRequestParams) SetTeamsObjectProtectionParams(v Office365TeamsObjectProtectionParams)`

SetTeamsObjectProtectionParams sets TeamsObjectProtectionParams field to given value.

### HasTeamsObjectProtectionParams

`func (o *Office365ObjectProtectionUpdateRequestParams) HasTeamsObjectProtectionParams() bool`

HasTeamsObjectProtectionParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


