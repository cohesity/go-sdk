# UpdateInfectedFileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdateFailedInfectedFiles** | Pointer to [**[]InfectedFileId**](InfectedFileId.md) | Specifies the failed update infected files. | [optional] 
**UpdateSucceededInfectedFiles** | Pointer to [**[]InfectedFileId**](InfectedFileId.md) | Specifies the successfully updated infected files. | [optional] 

## Methods

### NewUpdateInfectedFileResponse

`func NewUpdateInfectedFileResponse() *UpdateInfectedFileResponse`

NewUpdateInfectedFileResponse instantiates a new UpdateInfectedFileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInfectedFileResponseWithDefaults

`func NewUpdateInfectedFileResponseWithDefaults() *UpdateInfectedFileResponse`

NewUpdateInfectedFileResponseWithDefaults instantiates a new UpdateInfectedFileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdateFailedInfectedFiles

`func (o *UpdateInfectedFileResponse) GetUpdateFailedInfectedFiles() []InfectedFileId`

GetUpdateFailedInfectedFiles returns the UpdateFailedInfectedFiles field if non-nil, zero value otherwise.

### GetUpdateFailedInfectedFilesOk

`func (o *UpdateInfectedFileResponse) GetUpdateFailedInfectedFilesOk() (*[]InfectedFileId, bool)`

GetUpdateFailedInfectedFilesOk returns a tuple with the UpdateFailedInfectedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateFailedInfectedFiles

`func (o *UpdateInfectedFileResponse) SetUpdateFailedInfectedFiles(v []InfectedFileId)`

SetUpdateFailedInfectedFiles sets UpdateFailedInfectedFiles field to given value.

### HasUpdateFailedInfectedFiles

`func (o *UpdateInfectedFileResponse) HasUpdateFailedInfectedFiles() bool`

HasUpdateFailedInfectedFiles returns a boolean if a field has been set.

### SetUpdateFailedInfectedFilesNil

`func (o *UpdateInfectedFileResponse) SetUpdateFailedInfectedFilesNil(b bool)`

 SetUpdateFailedInfectedFilesNil sets the value for UpdateFailedInfectedFiles to be an explicit nil

### UnsetUpdateFailedInfectedFiles
`func (o *UpdateInfectedFileResponse) UnsetUpdateFailedInfectedFiles()`

UnsetUpdateFailedInfectedFiles ensures that no value is present for UpdateFailedInfectedFiles, not even an explicit nil
### GetUpdateSucceededInfectedFiles

`func (o *UpdateInfectedFileResponse) GetUpdateSucceededInfectedFiles() []InfectedFileId`

GetUpdateSucceededInfectedFiles returns the UpdateSucceededInfectedFiles field if non-nil, zero value otherwise.

### GetUpdateSucceededInfectedFilesOk

`func (o *UpdateInfectedFileResponse) GetUpdateSucceededInfectedFilesOk() (*[]InfectedFileId, bool)`

GetUpdateSucceededInfectedFilesOk returns a tuple with the UpdateSucceededInfectedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSucceededInfectedFiles

`func (o *UpdateInfectedFileResponse) SetUpdateSucceededInfectedFiles(v []InfectedFileId)`

SetUpdateSucceededInfectedFiles sets UpdateSucceededInfectedFiles field to given value.

### HasUpdateSucceededInfectedFiles

`func (o *UpdateInfectedFileResponse) HasUpdateSucceededInfectedFiles() bool`

HasUpdateSucceededInfectedFiles returns a boolean if a field has been set.

### SetUpdateSucceededInfectedFilesNil

`func (o *UpdateInfectedFileResponse) SetUpdateSucceededInfectedFilesNil(b bool)`

 SetUpdateSucceededInfectedFilesNil sets the value for UpdateSucceededInfectedFiles to be an explicit nil

### UnsetUpdateSucceededInfectedFiles
`func (o *UpdateInfectedFileResponse) UnsetUpdateSucceededInfectedFiles()`

UnsetUpdateSucceededInfectedFiles ensures that no value is present for UpdateSucceededInfectedFiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


