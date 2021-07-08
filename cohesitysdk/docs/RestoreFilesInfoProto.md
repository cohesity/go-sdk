# RestoreFilesInfoProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadFilesPath** | Pointer to **NullableString** | The path that the user should use to download files through the UI. If only a single file needs to be downloaded, this will be the path to that file. If a directory or multiple files &amp; directories need to be downloaded, this will be the path to the .zip file to download. Only applicable to a download files task. | [optional] 
**Error** | Pointer to [**ErrorProto**](ErrorProto.md) |  | [optional] 
**ProxyEntityConnectorParams** | Pointer to [**ConnectorParams**](ConnectorParams.md) |  | [optional] 
**RestoreFilesResultVec** | Pointer to [**[]RestoreFileResultInfo**](RestoreFileResultInfo.md) | Contains the result information of the restored files. | [optional] 
**SlaveTaskStartTimeUsecs** | Pointer to **NullableInt64** | This is the timestamp at which the slave task started. | [optional] 
**TargetType** | Pointer to **NullableInt32** | Specifies the target type for the task. The field is only valid if the task has got a permit. | [optional] 
**TeardownError** | Pointer to [**ErrorProto**](ErrorProto.md) |  | [optional] 
**Type** | Pointer to **NullableInt32** | The type of environment this restore files info pertains to. | [optional] 

## Methods

### NewRestoreFilesInfoProto

`func NewRestoreFilesInfoProto() *RestoreFilesInfoProto`

NewRestoreFilesInfoProto instantiates a new RestoreFilesInfoProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreFilesInfoProtoWithDefaults

`func NewRestoreFilesInfoProtoWithDefaults() *RestoreFilesInfoProto`

NewRestoreFilesInfoProtoWithDefaults instantiates a new RestoreFilesInfoProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadFilesPath

`func (o *RestoreFilesInfoProto) GetDownloadFilesPath() string`

GetDownloadFilesPath returns the DownloadFilesPath field if non-nil, zero value otherwise.

### GetDownloadFilesPathOk

`func (o *RestoreFilesInfoProto) GetDownloadFilesPathOk() (*string, bool)`

GetDownloadFilesPathOk returns a tuple with the DownloadFilesPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadFilesPath

`func (o *RestoreFilesInfoProto) SetDownloadFilesPath(v string)`

SetDownloadFilesPath sets DownloadFilesPath field to given value.

### HasDownloadFilesPath

`func (o *RestoreFilesInfoProto) HasDownloadFilesPath() bool`

HasDownloadFilesPath returns a boolean if a field has been set.

### SetDownloadFilesPathNil

`func (o *RestoreFilesInfoProto) SetDownloadFilesPathNil(b bool)`

 SetDownloadFilesPathNil sets the value for DownloadFilesPath to be an explicit nil

### UnsetDownloadFilesPath
`func (o *RestoreFilesInfoProto) UnsetDownloadFilesPath()`

UnsetDownloadFilesPath ensures that no value is present for DownloadFilesPath, not even an explicit nil
### GetError

`func (o *RestoreFilesInfoProto) GetError() ErrorProto`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RestoreFilesInfoProto) GetErrorOk() (*ErrorProto, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RestoreFilesInfoProto) SetError(v ErrorProto)`

SetError sets Error field to given value.

### HasError

`func (o *RestoreFilesInfoProto) HasError() bool`

HasError returns a boolean if a field has been set.

### GetProxyEntityConnectorParams

`func (o *RestoreFilesInfoProto) GetProxyEntityConnectorParams() ConnectorParams`

GetProxyEntityConnectorParams returns the ProxyEntityConnectorParams field if non-nil, zero value otherwise.

### GetProxyEntityConnectorParamsOk

`func (o *RestoreFilesInfoProto) GetProxyEntityConnectorParamsOk() (*ConnectorParams, bool)`

GetProxyEntityConnectorParamsOk returns a tuple with the ProxyEntityConnectorParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyEntityConnectorParams

`func (o *RestoreFilesInfoProto) SetProxyEntityConnectorParams(v ConnectorParams)`

SetProxyEntityConnectorParams sets ProxyEntityConnectorParams field to given value.

### HasProxyEntityConnectorParams

`func (o *RestoreFilesInfoProto) HasProxyEntityConnectorParams() bool`

HasProxyEntityConnectorParams returns a boolean if a field has been set.

### GetRestoreFilesResultVec

`func (o *RestoreFilesInfoProto) GetRestoreFilesResultVec() []RestoreFileResultInfo`

GetRestoreFilesResultVec returns the RestoreFilesResultVec field if non-nil, zero value otherwise.

### GetRestoreFilesResultVecOk

`func (o *RestoreFilesInfoProto) GetRestoreFilesResultVecOk() (*[]RestoreFileResultInfo, bool)`

GetRestoreFilesResultVecOk returns a tuple with the RestoreFilesResultVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreFilesResultVec

`func (o *RestoreFilesInfoProto) SetRestoreFilesResultVec(v []RestoreFileResultInfo)`

SetRestoreFilesResultVec sets RestoreFilesResultVec field to given value.

### HasRestoreFilesResultVec

`func (o *RestoreFilesInfoProto) HasRestoreFilesResultVec() bool`

HasRestoreFilesResultVec returns a boolean if a field has been set.

### SetRestoreFilesResultVecNil

`func (o *RestoreFilesInfoProto) SetRestoreFilesResultVecNil(b bool)`

 SetRestoreFilesResultVecNil sets the value for RestoreFilesResultVec to be an explicit nil

### UnsetRestoreFilesResultVec
`func (o *RestoreFilesInfoProto) UnsetRestoreFilesResultVec()`

UnsetRestoreFilesResultVec ensures that no value is present for RestoreFilesResultVec, not even an explicit nil
### GetSlaveTaskStartTimeUsecs

`func (o *RestoreFilesInfoProto) GetSlaveTaskStartTimeUsecs() int64`

GetSlaveTaskStartTimeUsecs returns the SlaveTaskStartTimeUsecs field if non-nil, zero value otherwise.

### GetSlaveTaskStartTimeUsecsOk

`func (o *RestoreFilesInfoProto) GetSlaveTaskStartTimeUsecsOk() (*int64, bool)`

GetSlaveTaskStartTimeUsecsOk returns a tuple with the SlaveTaskStartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaveTaskStartTimeUsecs

`func (o *RestoreFilesInfoProto) SetSlaveTaskStartTimeUsecs(v int64)`

SetSlaveTaskStartTimeUsecs sets SlaveTaskStartTimeUsecs field to given value.

### HasSlaveTaskStartTimeUsecs

`func (o *RestoreFilesInfoProto) HasSlaveTaskStartTimeUsecs() bool`

HasSlaveTaskStartTimeUsecs returns a boolean if a field has been set.

### SetSlaveTaskStartTimeUsecsNil

`func (o *RestoreFilesInfoProto) SetSlaveTaskStartTimeUsecsNil(b bool)`

 SetSlaveTaskStartTimeUsecsNil sets the value for SlaveTaskStartTimeUsecs to be an explicit nil

### UnsetSlaveTaskStartTimeUsecs
`func (o *RestoreFilesInfoProto) UnsetSlaveTaskStartTimeUsecs()`

UnsetSlaveTaskStartTimeUsecs ensures that no value is present for SlaveTaskStartTimeUsecs, not even an explicit nil
### GetTargetType

`func (o *RestoreFilesInfoProto) GetTargetType() int32`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *RestoreFilesInfoProto) GetTargetTypeOk() (*int32, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *RestoreFilesInfoProto) SetTargetType(v int32)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *RestoreFilesInfoProto) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *RestoreFilesInfoProto) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *RestoreFilesInfoProto) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
### GetTeardownError

`func (o *RestoreFilesInfoProto) GetTeardownError() ErrorProto`

GetTeardownError returns the TeardownError field if non-nil, zero value otherwise.

### GetTeardownErrorOk

`func (o *RestoreFilesInfoProto) GetTeardownErrorOk() (*ErrorProto, bool)`

GetTeardownErrorOk returns a tuple with the TeardownError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeardownError

`func (o *RestoreFilesInfoProto) SetTeardownError(v ErrorProto)`

SetTeardownError sets TeardownError field to given value.

### HasTeardownError

`func (o *RestoreFilesInfoProto) HasTeardownError() bool`

HasTeardownError returns a boolean if a field has been set.

### GetType

`func (o *RestoreFilesInfoProto) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestoreFilesInfoProto) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestoreFilesInfoProto) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *RestoreFilesInfoProto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *RestoreFilesInfoProto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *RestoreFilesInfoProto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


