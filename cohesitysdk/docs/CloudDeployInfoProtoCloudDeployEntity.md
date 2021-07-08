# CloudDeployInfoProtoCloudDeployEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeployedVmName** | Pointer to **NullableString** | Optional name that should be used for deployed VM. | [optional] 
**Entity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**Error** | Pointer to [**ErrorProto**](ErrorProto.md) |  | [optional] 
**PreviousRelativeCloneDirPath** | Pointer to **NullableString** | Directory where files of the entity&#39;s previous snapshot were cloned to. Path is relative to the destination view. | [optional] 
**PreviousRelativeClonePaths** | Pointer to **[]string** | All the paths that the entity&#39;s previous snapshot files were cloned to. Each path is relative to the destination view. | [optional] 
**ProgressMonitorTaskPath** | Pointer to **NullableString** | Progress monitor task path for this entity which is relative to the root path of the cloud deploy task progress monitor. | [optional] 
**PublicStatus** | Pointer to **NullableInt32** | Iris-facing task state. This field is stamped during the export. | [optional] 
**RelativeClonePaths** | Pointer to **[]string** | All the paths that the entity&#39;s files were cloned to. Each path is relative to the destination view. | [optional] 
**Status** | Pointer to **NullableInt32** | The status of the entity. | [optional] 

## Methods

### NewCloudDeployInfoProtoCloudDeployEntity

`func NewCloudDeployInfoProtoCloudDeployEntity() *CloudDeployInfoProtoCloudDeployEntity`

NewCloudDeployInfoProtoCloudDeployEntity instantiates a new CloudDeployInfoProtoCloudDeployEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudDeployInfoProtoCloudDeployEntityWithDefaults

`func NewCloudDeployInfoProtoCloudDeployEntityWithDefaults() *CloudDeployInfoProtoCloudDeployEntity`

NewCloudDeployInfoProtoCloudDeployEntityWithDefaults instantiates a new CloudDeployInfoProtoCloudDeployEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployedVmName

`func (o *CloudDeployInfoProtoCloudDeployEntity) GetDeployedVmName() string`

GetDeployedVmName returns the DeployedVmName field if non-nil, zero value otherwise.

### GetDeployedVmNameOk

`func (o *CloudDeployInfoProtoCloudDeployEntity) GetDeployedVmNameOk() (*string, bool)`

GetDeployedVmNameOk returns a tuple with the DeployedVmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedVmName

`func (o *CloudDeployInfoProtoCloudDeployEntity) SetDeployedVmName(v string)`

SetDeployedVmName sets DeployedVmName field to given value.

### HasDeployedVmName

`func (o *CloudDeployInfoProtoCloudDeployEntity) HasDeployedVmName() bool`

HasDeployedVmName returns a boolean if a field has been set.

### SetDeployedVmNameNil

`func (o *CloudDeployInfoProtoCloudDeployEntity) SetDeployedVmNameNil(b bool)`

 SetDeployedVmNameNil sets the value for DeployedVmName to be an explicit nil

### UnsetDeployedVmName
`func (o *CloudDeployInfoProtoCloudDeployEntity) UnsetDeployedVmName()`

UnsetDeployedVmName ensures that no value is present for DeployedVmName, not even an explicit nil
### GetEntity

`func (o *CloudDeployInfoProtoCloudDeployEntity) GetEntity() EntityProto`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *CloudDeployInfoProtoCloudDeployEntity) GetEntityOk() (*EntityProto, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *CloudDeployInfoProtoCloudDeployEntity) SetEntity(v EntityProto)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *CloudDeployInfoProtoCloudDeployEntity) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetError

`func (o *CloudDeployInfoProtoCloudDeployEntity) GetError() ErrorProto`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CloudDeployInfoProtoCloudDeployEntity) GetErrorOk() (*ErrorProto, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CloudDeployInfoProtoCloudDeployEntity) SetError(v ErrorProto)`

SetError sets Error field to given value.

### HasError

`func (o *CloudDeployInfoProtoCloudDeployEntity) HasError() bool`

HasError returns a boolean if a field has been set.

### GetPreviousRelativeCloneDirPath

`func (o *CloudDeployInfoProtoCloudDeployEntity) GetPreviousRelativeCloneDirPath() string`

GetPreviousRelativeCloneDirPath returns the PreviousRelativeCloneDirPath field if non-nil, zero value otherwise.

### GetPreviousRelativeCloneDirPathOk

`func (o *CloudDeployInfoProtoCloudDeployEntity) GetPreviousRelativeCloneDirPathOk() (*string, bool)`

GetPreviousRelativeCloneDirPathOk returns a tuple with the PreviousRelativeCloneDirPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousRelativeCloneDirPath

`func (o *CloudDeployInfoProtoCloudDeployEntity) SetPreviousRelativeCloneDirPath(v string)`

SetPreviousRelativeCloneDirPath sets PreviousRelativeCloneDirPath field to given value.

### HasPreviousRelativeCloneDirPath

`func (o *CloudDeployInfoProtoCloudDeployEntity) HasPreviousRelativeCloneDirPath() bool`

HasPreviousRelativeCloneDirPath returns a boolean if a field has been set.

### SetPreviousRelativeCloneDirPathNil

`func (o *CloudDeployInfoProtoCloudDeployEntity) SetPreviousRelativeCloneDirPathNil(b bool)`

 SetPreviousRelativeCloneDirPathNil sets the value for PreviousRelativeCloneDirPath to be an explicit nil

### UnsetPreviousRelativeCloneDirPath
`func (o *CloudDeployInfoProtoCloudDeployEntity) UnsetPreviousRelativeCloneDirPath()`

UnsetPreviousRelativeCloneDirPath ensures that no value is present for PreviousRelativeCloneDirPath, not even an explicit nil
### GetPreviousRelativeClonePaths

`func (o *CloudDeployInfoProtoCloudDeployEntity) GetPreviousRelativeClonePaths() []string`

GetPreviousRelativeClonePaths returns the PreviousRelativeClonePaths field if non-nil, zero value otherwise.

### GetPreviousRelativeClonePathsOk

`func (o *CloudDeployInfoProtoCloudDeployEntity) GetPreviousRelativeClonePathsOk() (*[]string, bool)`

GetPreviousRelativeClonePathsOk returns a tuple with the PreviousRelativeClonePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousRelativeClonePaths

`func (o *CloudDeployInfoProtoCloudDeployEntity) SetPreviousRelativeClonePaths(v []string)`

SetPreviousRelativeClonePaths sets PreviousRelativeClonePaths field to given value.

### HasPreviousRelativeClonePaths

`func (o *CloudDeployInfoProtoCloudDeployEntity) HasPreviousRelativeClonePaths() bool`

HasPreviousRelativeClonePaths returns a boolean if a field has been set.

### SetPreviousRelativeClonePathsNil

`func (o *CloudDeployInfoProtoCloudDeployEntity) SetPreviousRelativeClonePathsNil(b bool)`

 SetPreviousRelativeClonePathsNil sets the value for PreviousRelativeClonePaths to be an explicit nil

### UnsetPreviousRelativeClonePaths
`func (o *CloudDeployInfoProtoCloudDeployEntity) UnsetPreviousRelativeClonePaths()`

UnsetPreviousRelativeClonePaths ensures that no value is present for PreviousRelativeClonePaths, not even an explicit nil
### GetProgressMonitorTaskPath

`func (o *CloudDeployInfoProtoCloudDeployEntity) GetProgressMonitorTaskPath() string`

GetProgressMonitorTaskPath returns the ProgressMonitorTaskPath field if non-nil, zero value otherwise.

### GetProgressMonitorTaskPathOk

`func (o *CloudDeployInfoProtoCloudDeployEntity) GetProgressMonitorTaskPathOk() (*string, bool)`

GetProgressMonitorTaskPathOk returns a tuple with the ProgressMonitorTaskPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressMonitorTaskPath

`func (o *CloudDeployInfoProtoCloudDeployEntity) SetProgressMonitorTaskPath(v string)`

SetProgressMonitorTaskPath sets ProgressMonitorTaskPath field to given value.

### HasProgressMonitorTaskPath

`func (o *CloudDeployInfoProtoCloudDeployEntity) HasProgressMonitorTaskPath() bool`

HasProgressMonitorTaskPath returns a boolean if a field has been set.

### SetProgressMonitorTaskPathNil

`func (o *CloudDeployInfoProtoCloudDeployEntity) SetProgressMonitorTaskPathNil(b bool)`

 SetProgressMonitorTaskPathNil sets the value for ProgressMonitorTaskPath to be an explicit nil

### UnsetProgressMonitorTaskPath
`func (o *CloudDeployInfoProtoCloudDeployEntity) UnsetProgressMonitorTaskPath()`

UnsetProgressMonitorTaskPath ensures that no value is present for ProgressMonitorTaskPath, not even an explicit nil
### GetPublicStatus

`func (o *CloudDeployInfoProtoCloudDeployEntity) GetPublicStatus() int32`

GetPublicStatus returns the PublicStatus field if non-nil, zero value otherwise.

### GetPublicStatusOk

`func (o *CloudDeployInfoProtoCloudDeployEntity) GetPublicStatusOk() (*int32, bool)`

GetPublicStatusOk returns a tuple with the PublicStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicStatus

`func (o *CloudDeployInfoProtoCloudDeployEntity) SetPublicStatus(v int32)`

SetPublicStatus sets PublicStatus field to given value.

### HasPublicStatus

`func (o *CloudDeployInfoProtoCloudDeployEntity) HasPublicStatus() bool`

HasPublicStatus returns a boolean if a field has been set.

### SetPublicStatusNil

`func (o *CloudDeployInfoProtoCloudDeployEntity) SetPublicStatusNil(b bool)`

 SetPublicStatusNil sets the value for PublicStatus to be an explicit nil

### UnsetPublicStatus
`func (o *CloudDeployInfoProtoCloudDeployEntity) UnsetPublicStatus()`

UnsetPublicStatus ensures that no value is present for PublicStatus, not even an explicit nil
### GetRelativeClonePaths

`func (o *CloudDeployInfoProtoCloudDeployEntity) GetRelativeClonePaths() []string`

GetRelativeClonePaths returns the RelativeClonePaths field if non-nil, zero value otherwise.

### GetRelativeClonePathsOk

`func (o *CloudDeployInfoProtoCloudDeployEntity) GetRelativeClonePathsOk() (*[]string, bool)`

GetRelativeClonePathsOk returns a tuple with the RelativeClonePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeClonePaths

`func (o *CloudDeployInfoProtoCloudDeployEntity) SetRelativeClonePaths(v []string)`

SetRelativeClonePaths sets RelativeClonePaths field to given value.

### HasRelativeClonePaths

`func (o *CloudDeployInfoProtoCloudDeployEntity) HasRelativeClonePaths() bool`

HasRelativeClonePaths returns a boolean if a field has been set.

### SetRelativeClonePathsNil

`func (o *CloudDeployInfoProtoCloudDeployEntity) SetRelativeClonePathsNil(b bool)`

 SetRelativeClonePathsNil sets the value for RelativeClonePaths to be an explicit nil

### UnsetRelativeClonePaths
`func (o *CloudDeployInfoProtoCloudDeployEntity) UnsetRelativeClonePaths()`

UnsetRelativeClonePaths ensures that no value is present for RelativeClonePaths, not even an explicit nil
### GetStatus

`func (o *CloudDeployInfoProtoCloudDeployEntity) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudDeployInfoProtoCloudDeployEntity) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudDeployInfoProtoCloudDeployEntity) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudDeployInfoProtoCloudDeployEntity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CloudDeployInfoProtoCloudDeployEntity) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CloudDeployInfoProtoCloudDeployEntity) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


