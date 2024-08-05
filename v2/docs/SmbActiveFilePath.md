# SmbActiveFilePath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveSessions** | Pointer to [**[]SmbActiveSession**](SmbActiveSession.md) | Specifies an active session where the file is open. | [optional] 
**FilePath** | Pointer to **NullableString** | Specifies the filepath in the view. | [optional] 
**ViewId** | Pointer to **NullableInt64** | Specifies the id of the View. | [optional] 
**ViewName** | Pointer to **NullableString** | Specifies the name of the View. | [optional] 

## Methods

### NewSmbActiveFilePath

`func NewSmbActiveFilePath() *SmbActiveFilePath`

NewSmbActiveFilePath instantiates a new SmbActiveFilePath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmbActiveFilePathWithDefaults

`func NewSmbActiveFilePathWithDefaults() *SmbActiveFilePath`

NewSmbActiveFilePathWithDefaults instantiates a new SmbActiveFilePath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveSessions

`func (o *SmbActiveFilePath) GetActiveSessions() []SmbActiveSession`

GetActiveSessions returns the ActiveSessions field if non-nil, zero value otherwise.

### GetActiveSessionsOk

`func (o *SmbActiveFilePath) GetActiveSessionsOk() (*[]SmbActiveSession, bool)`

GetActiveSessionsOk returns a tuple with the ActiveSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSessions

`func (o *SmbActiveFilePath) SetActiveSessions(v []SmbActiveSession)`

SetActiveSessions sets ActiveSessions field to given value.

### HasActiveSessions

`func (o *SmbActiveFilePath) HasActiveSessions() bool`

HasActiveSessions returns a boolean if a field has been set.

### SetActiveSessionsNil

`func (o *SmbActiveFilePath) SetActiveSessionsNil(b bool)`

 SetActiveSessionsNil sets the value for ActiveSessions to be an explicit nil

### UnsetActiveSessions
`func (o *SmbActiveFilePath) UnsetActiveSessions()`

UnsetActiveSessions ensures that no value is present for ActiveSessions, not even an explicit nil
### GetFilePath

`func (o *SmbActiveFilePath) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *SmbActiveFilePath) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *SmbActiveFilePath) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *SmbActiveFilePath) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### SetFilePathNil

`func (o *SmbActiveFilePath) SetFilePathNil(b bool)`

 SetFilePathNil sets the value for FilePath to be an explicit nil

### UnsetFilePath
`func (o *SmbActiveFilePath) UnsetFilePath()`

UnsetFilePath ensures that no value is present for FilePath, not even an explicit nil
### GetViewId

`func (o *SmbActiveFilePath) GetViewId() int64`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *SmbActiveFilePath) GetViewIdOk() (*int64, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *SmbActiveFilePath) SetViewId(v int64)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *SmbActiveFilePath) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### SetViewIdNil

`func (o *SmbActiveFilePath) SetViewIdNil(b bool)`

 SetViewIdNil sets the value for ViewId to be an explicit nil

### UnsetViewId
`func (o *SmbActiveFilePath) UnsetViewId()`

UnsetViewId ensures that no value is present for ViewId, not even an explicit nil
### GetViewName

`func (o *SmbActiveFilePath) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *SmbActiveFilePath) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *SmbActiveFilePath) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *SmbActiveFilePath) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *SmbActiveFilePath) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *SmbActiveFilePath) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


