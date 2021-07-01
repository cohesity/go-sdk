# AppEntityBackupStatusInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **NullableInt64** | Specifies the Id of the App entity. This is typically a database entity in case of SQL, Oracle jobs etc. | [optional] 
**Error** | Pointer to **NullableString** | Specifies if an error occurred (if any) while running this task. This field is populated when the status is equal to &#39;kFailure&#39;. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the app entity. | [optional] 
**OwnerId** | Pointer to **NullableInt64** | Specifies the owner id of the the app. Owner is the host under which the app is located. Example: SQL DB entities can be hosted by Physical host or virtual machine. | [optional] 
**Status** | Pointer to **NullableString** | Specifies the backup status for this app entity. &#39;kAccepted&#39; indicates the task is queued to run but not yet running. &#39;kRunning&#39; indicates the task is running. &#39;kCanceling&#39; indicates a request to cancel the task has occurred but the task is not yet canceled. &#39;kCanceled&#39; indicates the task has been canceled. &#39;kSuccess&#39; indicates the task was successful. &#39;kFailure&#39; indicates the task failed. &#39;kWarning&#39; indicates the task has finished with warning. &#39;kOnHold&#39; indicates the task is kept onHold. &#39;kMissed&#39; indicates the task is missed. | [optional] 
**Warnings** | Pointer to **[]string** | Specifies the warnings that occurred (if any) while running this task. | [optional] 

## Methods

### NewAppEntityBackupStatusInfo

`func NewAppEntityBackupStatusInfo() *AppEntityBackupStatusInfo`

NewAppEntityBackupStatusInfo instantiates a new AppEntityBackupStatusInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppEntityBackupStatusInfoWithDefaults

`func NewAppEntityBackupStatusInfoWithDefaults() *AppEntityBackupStatusInfo`

NewAppEntityBackupStatusInfoWithDefaults instantiates a new AppEntityBackupStatusInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *AppEntityBackupStatusInfo) GetAppId() int64`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AppEntityBackupStatusInfo) GetAppIdOk() (*int64, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AppEntityBackupStatusInfo) SetAppId(v int64)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *AppEntityBackupStatusInfo) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### SetAppIdNil

`func (o *AppEntityBackupStatusInfo) SetAppIdNil(b bool)`

 SetAppIdNil sets the value for AppId to be an explicit nil

### UnsetAppId
`func (o *AppEntityBackupStatusInfo) UnsetAppId()`

UnsetAppId ensures that no value is present for AppId, not even an explicit nil
### GetError

`func (o *AppEntityBackupStatusInfo) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AppEntityBackupStatusInfo) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AppEntityBackupStatusInfo) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AppEntityBackupStatusInfo) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *AppEntityBackupStatusInfo) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *AppEntityBackupStatusInfo) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetName

`func (o *AppEntityBackupStatusInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppEntityBackupStatusInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppEntityBackupStatusInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppEntityBackupStatusInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AppEntityBackupStatusInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AppEntityBackupStatusInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOwnerId

`func (o *AppEntityBackupStatusInfo) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *AppEntityBackupStatusInfo) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *AppEntityBackupStatusInfo) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *AppEntityBackupStatusInfo) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *AppEntityBackupStatusInfo) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *AppEntityBackupStatusInfo) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetStatus

`func (o *AppEntityBackupStatusInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppEntityBackupStatusInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppEntityBackupStatusInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AppEntityBackupStatusInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AppEntityBackupStatusInfo) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AppEntityBackupStatusInfo) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetWarnings

`func (o *AppEntityBackupStatusInfo) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *AppEntityBackupStatusInfo) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *AppEntityBackupStatusInfo) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *AppEntityBackupStatusInfo) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *AppEntityBackupStatusInfo) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *AppEntityBackupStatusInfo) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


