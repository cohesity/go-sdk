# DeleteInfectedFiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteFailedInfectedFiles** | Pointer to [**[]InfectedFile**](InfectedFile.md) | Specifies the list of infected files that failed deletion. | [optional] 
**DeleteSucceededInfectedFiles** | Pointer to [**[]InfectedFile**](InfectedFile.md) | Specifies the list of infected files that are successfully deleted. | [optional] 

## Methods

### NewDeleteInfectedFiles

`func NewDeleteInfectedFiles() *DeleteInfectedFiles`

NewDeleteInfectedFiles instantiates a new DeleteInfectedFiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteInfectedFilesWithDefaults

`func NewDeleteInfectedFilesWithDefaults() *DeleteInfectedFiles`

NewDeleteInfectedFilesWithDefaults instantiates a new DeleteInfectedFiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteFailedInfectedFiles

`func (o *DeleteInfectedFiles) GetDeleteFailedInfectedFiles() []InfectedFile`

GetDeleteFailedInfectedFiles returns the DeleteFailedInfectedFiles field if non-nil, zero value otherwise.

### GetDeleteFailedInfectedFilesOk

`func (o *DeleteInfectedFiles) GetDeleteFailedInfectedFilesOk() (*[]InfectedFile, bool)`

GetDeleteFailedInfectedFilesOk returns a tuple with the DeleteFailedInfectedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteFailedInfectedFiles

`func (o *DeleteInfectedFiles) SetDeleteFailedInfectedFiles(v []InfectedFile)`

SetDeleteFailedInfectedFiles sets DeleteFailedInfectedFiles field to given value.

### HasDeleteFailedInfectedFiles

`func (o *DeleteInfectedFiles) HasDeleteFailedInfectedFiles() bool`

HasDeleteFailedInfectedFiles returns a boolean if a field has been set.

### SetDeleteFailedInfectedFilesNil

`func (o *DeleteInfectedFiles) SetDeleteFailedInfectedFilesNil(b bool)`

 SetDeleteFailedInfectedFilesNil sets the value for DeleteFailedInfectedFiles to be an explicit nil

### UnsetDeleteFailedInfectedFiles
`func (o *DeleteInfectedFiles) UnsetDeleteFailedInfectedFiles()`

UnsetDeleteFailedInfectedFiles ensures that no value is present for DeleteFailedInfectedFiles, not even an explicit nil
### GetDeleteSucceededInfectedFiles

`func (o *DeleteInfectedFiles) GetDeleteSucceededInfectedFiles() []InfectedFile`

GetDeleteSucceededInfectedFiles returns the DeleteSucceededInfectedFiles field if non-nil, zero value otherwise.

### GetDeleteSucceededInfectedFilesOk

`func (o *DeleteInfectedFiles) GetDeleteSucceededInfectedFilesOk() (*[]InfectedFile, bool)`

GetDeleteSucceededInfectedFilesOk returns a tuple with the DeleteSucceededInfectedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteSucceededInfectedFiles

`func (o *DeleteInfectedFiles) SetDeleteSucceededInfectedFiles(v []InfectedFile)`

SetDeleteSucceededInfectedFiles sets DeleteSucceededInfectedFiles field to given value.

### HasDeleteSucceededInfectedFiles

`func (o *DeleteInfectedFiles) HasDeleteSucceededInfectedFiles() bool`

HasDeleteSucceededInfectedFiles returns a boolean if a field has been set.

### SetDeleteSucceededInfectedFilesNil

`func (o *DeleteInfectedFiles) SetDeleteSucceededInfectedFilesNil(b bool)`

 SetDeleteSucceededInfectedFilesNil sets the value for DeleteSucceededInfectedFiles to be an explicit nil

### UnsetDeleteSucceededInfectedFiles
`func (o *DeleteInfectedFiles) UnsetDeleteSucceededInfectedFiles()`

UnsetDeleteSucceededInfectedFiles ensures that no value is present for DeleteSucceededInfectedFiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


