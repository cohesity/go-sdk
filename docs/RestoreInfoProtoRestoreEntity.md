# RestoreInfoProtoRestoreEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**Error** | Pointer to [**ErrorProto**](ErrorProto.md) |  | [optional] 
**ProgressMonitorTaskPath** | Pointer to **NullableString** | The path relative to the root path of the restore task progress monitor of the progress monitor for this entity. | [optional] 
**PublicStatus** | Pointer to **NullableInt32** | Iris-facing task state. This field is stamped during the export. | [optional] 
**RelativeRestorePaths** | Pointer to **[]string** | All the paths that the entity&#39;s files were restored to. Each path is relative to the destination view. | [optional] 
**ResourcePoolEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**RestoredEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**Status** | Pointer to **NullableInt32** | The restore status of the entity. | [optional] 
**Warnings** | Pointer to [**[]ErrorProto**](ErrorProto.md) | Optional warnings if any. | [optional] 

## Methods

### NewRestoreInfoProtoRestoreEntity

`func NewRestoreInfoProtoRestoreEntity() *RestoreInfoProtoRestoreEntity`

NewRestoreInfoProtoRestoreEntity instantiates a new RestoreInfoProtoRestoreEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreInfoProtoRestoreEntityWithDefaults

`func NewRestoreInfoProtoRestoreEntityWithDefaults() *RestoreInfoProtoRestoreEntity`

NewRestoreInfoProtoRestoreEntityWithDefaults instantiates a new RestoreInfoProtoRestoreEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntity

`func (o *RestoreInfoProtoRestoreEntity) GetEntity() EntityProto`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *RestoreInfoProtoRestoreEntity) GetEntityOk() (*EntityProto, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *RestoreInfoProtoRestoreEntity) SetEntity(v EntityProto)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *RestoreInfoProtoRestoreEntity) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetError

`func (o *RestoreInfoProtoRestoreEntity) GetError() ErrorProto`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RestoreInfoProtoRestoreEntity) GetErrorOk() (*ErrorProto, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RestoreInfoProtoRestoreEntity) SetError(v ErrorProto)`

SetError sets Error field to given value.

### HasError

`func (o *RestoreInfoProtoRestoreEntity) HasError() bool`

HasError returns a boolean if a field has been set.

### GetProgressMonitorTaskPath

`func (o *RestoreInfoProtoRestoreEntity) GetProgressMonitorTaskPath() string`

GetProgressMonitorTaskPath returns the ProgressMonitorTaskPath field if non-nil, zero value otherwise.

### GetProgressMonitorTaskPathOk

`func (o *RestoreInfoProtoRestoreEntity) GetProgressMonitorTaskPathOk() (*string, bool)`

GetProgressMonitorTaskPathOk returns a tuple with the ProgressMonitorTaskPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressMonitorTaskPath

`func (o *RestoreInfoProtoRestoreEntity) SetProgressMonitorTaskPath(v string)`

SetProgressMonitorTaskPath sets ProgressMonitorTaskPath field to given value.

### HasProgressMonitorTaskPath

`func (o *RestoreInfoProtoRestoreEntity) HasProgressMonitorTaskPath() bool`

HasProgressMonitorTaskPath returns a boolean if a field has been set.

### SetProgressMonitorTaskPathNil

`func (o *RestoreInfoProtoRestoreEntity) SetProgressMonitorTaskPathNil(b bool)`

 SetProgressMonitorTaskPathNil sets the value for ProgressMonitorTaskPath to be an explicit nil

### UnsetProgressMonitorTaskPath
`func (o *RestoreInfoProtoRestoreEntity) UnsetProgressMonitorTaskPath()`

UnsetProgressMonitorTaskPath ensures that no value is present for ProgressMonitorTaskPath, not even an explicit nil
### GetPublicStatus

`func (o *RestoreInfoProtoRestoreEntity) GetPublicStatus() int32`

GetPublicStatus returns the PublicStatus field if non-nil, zero value otherwise.

### GetPublicStatusOk

`func (o *RestoreInfoProtoRestoreEntity) GetPublicStatusOk() (*int32, bool)`

GetPublicStatusOk returns a tuple with the PublicStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicStatus

`func (o *RestoreInfoProtoRestoreEntity) SetPublicStatus(v int32)`

SetPublicStatus sets PublicStatus field to given value.

### HasPublicStatus

`func (o *RestoreInfoProtoRestoreEntity) HasPublicStatus() bool`

HasPublicStatus returns a boolean if a field has been set.

### SetPublicStatusNil

`func (o *RestoreInfoProtoRestoreEntity) SetPublicStatusNil(b bool)`

 SetPublicStatusNil sets the value for PublicStatus to be an explicit nil

### UnsetPublicStatus
`func (o *RestoreInfoProtoRestoreEntity) UnsetPublicStatus()`

UnsetPublicStatus ensures that no value is present for PublicStatus, not even an explicit nil
### GetRelativeRestorePaths

`func (o *RestoreInfoProtoRestoreEntity) GetRelativeRestorePaths() []string`

GetRelativeRestorePaths returns the RelativeRestorePaths field if non-nil, zero value otherwise.

### GetRelativeRestorePathsOk

`func (o *RestoreInfoProtoRestoreEntity) GetRelativeRestorePathsOk() (*[]string, bool)`

GetRelativeRestorePathsOk returns a tuple with the RelativeRestorePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeRestorePaths

`func (o *RestoreInfoProtoRestoreEntity) SetRelativeRestorePaths(v []string)`

SetRelativeRestorePaths sets RelativeRestorePaths field to given value.

### HasRelativeRestorePaths

`func (o *RestoreInfoProtoRestoreEntity) HasRelativeRestorePaths() bool`

HasRelativeRestorePaths returns a boolean if a field has been set.

### SetRelativeRestorePathsNil

`func (o *RestoreInfoProtoRestoreEntity) SetRelativeRestorePathsNil(b bool)`

 SetRelativeRestorePathsNil sets the value for RelativeRestorePaths to be an explicit nil

### UnsetRelativeRestorePaths
`func (o *RestoreInfoProtoRestoreEntity) UnsetRelativeRestorePaths()`

UnsetRelativeRestorePaths ensures that no value is present for RelativeRestorePaths, not even an explicit nil
### GetResourcePoolEntity

`func (o *RestoreInfoProtoRestoreEntity) GetResourcePoolEntity() EntityProto`

GetResourcePoolEntity returns the ResourcePoolEntity field if non-nil, zero value otherwise.

### GetResourcePoolEntityOk

`func (o *RestoreInfoProtoRestoreEntity) GetResourcePoolEntityOk() (*EntityProto, bool)`

GetResourcePoolEntityOk returns a tuple with the ResourcePoolEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolEntity

`func (o *RestoreInfoProtoRestoreEntity) SetResourcePoolEntity(v EntityProto)`

SetResourcePoolEntity sets ResourcePoolEntity field to given value.

### HasResourcePoolEntity

`func (o *RestoreInfoProtoRestoreEntity) HasResourcePoolEntity() bool`

HasResourcePoolEntity returns a boolean if a field has been set.

### GetRestoredEntity

`func (o *RestoreInfoProtoRestoreEntity) GetRestoredEntity() EntityProto`

GetRestoredEntity returns the RestoredEntity field if non-nil, zero value otherwise.

### GetRestoredEntityOk

`func (o *RestoreInfoProtoRestoreEntity) GetRestoredEntityOk() (*EntityProto, bool)`

GetRestoredEntityOk returns a tuple with the RestoredEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoredEntity

`func (o *RestoreInfoProtoRestoreEntity) SetRestoredEntity(v EntityProto)`

SetRestoredEntity sets RestoredEntity field to given value.

### HasRestoredEntity

`func (o *RestoreInfoProtoRestoreEntity) HasRestoredEntity() bool`

HasRestoredEntity returns a boolean if a field has been set.

### GetStatus

`func (o *RestoreInfoProtoRestoreEntity) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RestoreInfoProtoRestoreEntity) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RestoreInfoProtoRestoreEntity) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RestoreInfoProtoRestoreEntity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *RestoreInfoProtoRestoreEntity) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *RestoreInfoProtoRestoreEntity) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetWarnings

`func (o *RestoreInfoProtoRestoreEntity) GetWarnings() []ErrorProto`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *RestoreInfoProtoRestoreEntity) GetWarningsOk() (*[]ErrorProto, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *RestoreInfoProtoRestoreEntity) SetWarnings(v []ErrorProto)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *RestoreInfoProtoRestoreEntity) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *RestoreInfoProtoRestoreEntity) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *RestoreInfoProtoRestoreEntity) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


