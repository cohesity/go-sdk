# UpdateInfectedFilesList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdateFailedInfectedFiles** | Pointer to [**[]InfectedFile**](InfectedFile.md) | Specifies the list of infected entities that failed update. | [optional] 
**UpdateSucceededInfectedFiles** | Pointer to [**[]InfectedFile**](InfectedFile.md) | Specifies the list of infected entities that are successfully updated. | [optional] 

## Methods

### NewUpdateInfectedFilesList

`func NewUpdateInfectedFilesList() *UpdateInfectedFilesList`

NewUpdateInfectedFilesList instantiates a new UpdateInfectedFilesList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInfectedFilesListWithDefaults

`func NewUpdateInfectedFilesListWithDefaults() *UpdateInfectedFilesList`

NewUpdateInfectedFilesListWithDefaults instantiates a new UpdateInfectedFilesList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdateFailedInfectedFiles

`func (o *UpdateInfectedFilesList) GetUpdateFailedInfectedFiles() []InfectedFile`

GetUpdateFailedInfectedFiles returns the UpdateFailedInfectedFiles field if non-nil, zero value otherwise.

### GetUpdateFailedInfectedFilesOk

`func (o *UpdateInfectedFilesList) GetUpdateFailedInfectedFilesOk() (*[]InfectedFile, bool)`

GetUpdateFailedInfectedFilesOk returns a tuple with the UpdateFailedInfectedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateFailedInfectedFiles

`func (o *UpdateInfectedFilesList) SetUpdateFailedInfectedFiles(v []InfectedFile)`

SetUpdateFailedInfectedFiles sets UpdateFailedInfectedFiles field to given value.

### HasUpdateFailedInfectedFiles

`func (o *UpdateInfectedFilesList) HasUpdateFailedInfectedFiles() bool`

HasUpdateFailedInfectedFiles returns a boolean if a field has been set.

### SetUpdateFailedInfectedFilesNil

`func (o *UpdateInfectedFilesList) SetUpdateFailedInfectedFilesNil(b bool)`

 SetUpdateFailedInfectedFilesNil sets the value for UpdateFailedInfectedFiles to be an explicit nil

### UnsetUpdateFailedInfectedFiles
`func (o *UpdateInfectedFilesList) UnsetUpdateFailedInfectedFiles()`

UnsetUpdateFailedInfectedFiles ensures that no value is present for UpdateFailedInfectedFiles, not even an explicit nil
### GetUpdateSucceededInfectedFiles

`func (o *UpdateInfectedFilesList) GetUpdateSucceededInfectedFiles() []InfectedFile`

GetUpdateSucceededInfectedFiles returns the UpdateSucceededInfectedFiles field if non-nil, zero value otherwise.

### GetUpdateSucceededInfectedFilesOk

`func (o *UpdateInfectedFilesList) GetUpdateSucceededInfectedFilesOk() (*[]InfectedFile, bool)`

GetUpdateSucceededInfectedFilesOk returns a tuple with the UpdateSucceededInfectedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSucceededInfectedFiles

`func (o *UpdateInfectedFilesList) SetUpdateSucceededInfectedFiles(v []InfectedFile)`

SetUpdateSucceededInfectedFiles sets UpdateSucceededInfectedFiles field to given value.

### HasUpdateSucceededInfectedFiles

`func (o *UpdateInfectedFilesList) HasUpdateSucceededInfectedFiles() bool`

HasUpdateSucceededInfectedFiles returns a boolean if a field has been set.

### SetUpdateSucceededInfectedFilesNil

`func (o *UpdateInfectedFilesList) SetUpdateSucceededInfectedFilesNil(b bool)`

 SetUpdateSucceededInfectedFilesNil sets the value for UpdateSucceededInfectedFiles to be an explicit nil

### UnsetUpdateSucceededInfectedFiles
`func (o *UpdateInfectedFilesList) UnsetUpdateSucceededInfectedFiles()`

UnsetUpdateSucceededInfectedFiles ensures that no value is present for UpdateSucceededInfectedFiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


