# SetupRestoreDiskTaskInfoProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**ProgressMonitorRootTaskPath** | Pointer to **NullableString** | The path to the progress monitor root task if any. | [optional] 
**RootEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**SourceViewName** | Pointer to **NullableString** | The source view which contains the backups for the &#39;entity&#39;. | [optional] 
**TaskId** | Pointer to **NullableInt64** | The id of the associated task. | [optional] 
**ViewBoxId** | Pointer to **NullableInt64** | The view box id containing the backups for the &#39;entity&#39;. | [optional] 
**ViewName** | Pointer to **NullableString** | Destination view into which the files will be cloned. | [optional] 

## Methods

### NewSetupRestoreDiskTaskInfoProto

`func NewSetupRestoreDiskTaskInfoProto() *SetupRestoreDiskTaskInfoProto`

NewSetupRestoreDiskTaskInfoProto instantiates a new SetupRestoreDiskTaskInfoProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetupRestoreDiskTaskInfoProtoWithDefaults

`func NewSetupRestoreDiskTaskInfoProtoWithDefaults() *SetupRestoreDiskTaskInfoProto`

NewSetupRestoreDiskTaskInfoProtoWithDefaults instantiates a new SetupRestoreDiskTaskInfoProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntity

`func (o *SetupRestoreDiskTaskInfoProto) GetEntity() EntityProto`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *SetupRestoreDiskTaskInfoProto) GetEntityOk() (*EntityProto, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *SetupRestoreDiskTaskInfoProto) SetEntity(v EntityProto)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *SetupRestoreDiskTaskInfoProto) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetProgressMonitorRootTaskPath

`func (o *SetupRestoreDiskTaskInfoProto) GetProgressMonitorRootTaskPath() string`

GetProgressMonitorRootTaskPath returns the ProgressMonitorRootTaskPath field if non-nil, zero value otherwise.

### GetProgressMonitorRootTaskPathOk

`func (o *SetupRestoreDiskTaskInfoProto) GetProgressMonitorRootTaskPathOk() (*string, bool)`

GetProgressMonitorRootTaskPathOk returns a tuple with the ProgressMonitorRootTaskPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressMonitorRootTaskPath

`func (o *SetupRestoreDiskTaskInfoProto) SetProgressMonitorRootTaskPath(v string)`

SetProgressMonitorRootTaskPath sets ProgressMonitorRootTaskPath field to given value.

### HasProgressMonitorRootTaskPath

`func (o *SetupRestoreDiskTaskInfoProto) HasProgressMonitorRootTaskPath() bool`

HasProgressMonitorRootTaskPath returns a boolean if a field has been set.

### SetProgressMonitorRootTaskPathNil

`func (o *SetupRestoreDiskTaskInfoProto) SetProgressMonitorRootTaskPathNil(b bool)`

 SetProgressMonitorRootTaskPathNil sets the value for ProgressMonitorRootTaskPath to be an explicit nil

### UnsetProgressMonitorRootTaskPath
`func (o *SetupRestoreDiskTaskInfoProto) UnsetProgressMonitorRootTaskPath()`

UnsetProgressMonitorRootTaskPath ensures that no value is present for ProgressMonitorRootTaskPath, not even an explicit nil
### GetRootEntity

`func (o *SetupRestoreDiskTaskInfoProto) GetRootEntity() EntityProto`

GetRootEntity returns the RootEntity field if non-nil, zero value otherwise.

### GetRootEntityOk

`func (o *SetupRestoreDiskTaskInfoProto) GetRootEntityOk() (*EntityProto, bool)`

GetRootEntityOk returns a tuple with the RootEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootEntity

`func (o *SetupRestoreDiskTaskInfoProto) SetRootEntity(v EntityProto)`

SetRootEntity sets RootEntity field to given value.

### HasRootEntity

`func (o *SetupRestoreDiskTaskInfoProto) HasRootEntity() bool`

HasRootEntity returns a boolean if a field has been set.

### GetSourceViewName

`func (o *SetupRestoreDiskTaskInfoProto) GetSourceViewName() string`

GetSourceViewName returns the SourceViewName field if non-nil, zero value otherwise.

### GetSourceViewNameOk

`func (o *SetupRestoreDiskTaskInfoProto) GetSourceViewNameOk() (*string, bool)`

GetSourceViewNameOk returns a tuple with the SourceViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceViewName

`func (o *SetupRestoreDiskTaskInfoProto) SetSourceViewName(v string)`

SetSourceViewName sets SourceViewName field to given value.

### HasSourceViewName

`func (o *SetupRestoreDiskTaskInfoProto) HasSourceViewName() bool`

HasSourceViewName returns a boolean if a field has been set.

### SetSourceViewNameNil

`func (o *SetupRestoreDiskTaskInfoProto) SetSourceViewNameNil(b bool)`

 SetSourceViewNameNil sets the value for SourceViewName to be an explicit nil

### UnsetSourceViewName
`func (o *SetupRestoreDiskTaskInfoProto) UnsetSourceViewName()`

UnsetSourceViewName ensures that no value is present for SourceViewName, not even an explicit nil
### GetTaskId

`func (o *SetupRestoreDiskTaskInfoProto) GetTaskId() int64`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *SetupRestoreDiskTaskInfoProto) GetTaskIdOk() (*int64, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *SetupRestoreDiskTaskInfoProto) SetTaskId(v int64)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *SetupRestoreDiskTaskInfoProto) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *SetupRestoreDiskTaskInfoProto) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *SetupRestoreDiskTaskInfoProto) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
### GetViewBoxId

`func (o *SetupRestoreDiskTaskInfoProto) GetViewBoxId() int64`

GetViewBoxId returns the ViewBoxId field if non-nil, zero value otherwise.

### GetViewBoxIdOk

`func (o *SetupRestoreDiskTaskInfoProto) GetViewBoxIdOk() (*int64, bool)`

GetViewBoxIdOk returns a tuple with the ViewBoxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewBoxId

`func (o *SetupRestoreDiskTaskInfoProto) SetViewBoxId(v int64)`

SetViewBoxId sets ViewBoxId field to given value.

### HasViewBoxId

`func (o *SetupRestoreDiskTaskInfoProto) HasViewBoxId() bool`

HasViewBoxId returns a boolean if a field has been set.

### SetViewBoxIdNil

`func (o *SetupRestoreDiskTaskInfoProto) SetViewBoxIdNil(b bool)`

 SetViewBoxIdNil sets the value for ViewBoxId to be an explicit nil

### UnsetViewBoxId
`func (o *SetupRestoreDiskTaskInfoProto) UnsetViewBoxId()`

UnsetViewBoxId ensures that no value is present for ViewBoxId, not even an explicit nil
### GetViewName

`func (o *SetupRestoreDiskTaskInfoProto) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *SetupRestoreDiskTaskInfoProto) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *SetupRestoreDiskTaskInfoProto) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *SetupRestoreDiskTaskInfoProto) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *SetupRestoreDiskTaskInfoProto) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *SetupRestoreDiskTaskInfoProto) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


