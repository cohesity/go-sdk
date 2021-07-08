# DeleteInfectedFileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteFailedInfectedFiles** | Pointer to [**[]InfectedFileId**](InfectedFileId.md) | Specifies the failed delete infected files. | [optional] 
**DeleteSucceededInfectedFiles** | Pointer to [**[]InfectedFileId**](InfectedFileId.md) | Specifies the successfully deleted infected files. | [optional] 

## Methods

### NewDeleteInfectedFileResponse

`func NewDeleteInfectedFileResponse() *DeleteInfectedFileResponse`

NewDeleteInfectedFileResponse instantiates a new DeleteInfectedFileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteInfectedFileResponseWithDefaults

`func NewDeleteInfectedFileResponseWithDefaults() *DeleteInfectedFileResponse`

NewDeleteInfectedFileResponseWithDefaults instantiates a new DeleteInfectedFileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteFailedInfectedFiles

`func (o *DeleteInfectedFileResponse) GetDeleteFailedInfectedFiles() []InfectedFileId`

GetDeleteFailedInfectedFiles returns the DeleteFailedInfectedFiles field if non-nil, zero value otherwise.

### GetDeleteFailedInfectedFilesOk

`func (o *DeleteInfectedFileResponse) GetDeleteFailedInfectedFilesOk() (*[]InfectedFileId, bool)`

GetDeleteFailedInfectedFilesOk returns a tuple with the DeleteFailedInfectedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteFailedInfectedFiles

`func (o *DeleteInfectedFileResponse) SetDeleteFailedInfectedFiles(v []InfectedFileId)`

SetDeleteFailedInfectedFiles sets DeleteFailedInfectedFiles field to given value.

### HasDeleteFailedInfectedFiles

`func (o *DeleteInfectedFileResponse) HasDeleteFailedInfectedFiles() bool`

HasDeleteFailedInfectedFiles returns a boolean if a field has been set.

### SetDeleteFailedInfectedFilesNil

`func (o *DeleteInfectedFileResponse) SetDeleteFailedInfectedFilesNil(b bool)`

 SetDeleteFailedInfectedFilesNil sets the value for DeleteFailedInfectedFiles to be an explicit nil

### UnsetDeleteFailedInfectedFiles
`func (o *DeleteInfectedFileResponse) UnsetDeleteFailedInfectedFiles()`

UnsetDeleteFailedInfectedFiles ensures that no value is present for DeleteFailedInfectedFiles, not even an explicit nil
### GetDeleteSucceededInfectedFiles

`func (o *DeleteInfectedFileResponse) GetDeleteSucceededInfectedFiles() []InfectedFileId`

GetDeleteSucceededInfectedFiles returns the DeleteSucceededInfectedFiles field if non-nil, zero value otherwise.

### GetDeleteSucceededInfectedFilesOk

`func (o *DeleteInfectedFileResponse) GetDeleteSucceededInfectedFilesOk() (*[]InfectedFileId, bool)`

GetDeleteSucceededInfectedFilesOk returns a tuple with the DeleteSucceededInfectedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteSucceededInfectedFiles

`func (o *DeleteInfectedFileResponse) SetDeleteSucceededInfectedFiles(v []InfectedFileId)`

SetDeleteSucceededInfectedFiles sets DeleteSucceededInfectedFiles field to given value.

### HasDeleteSucceededInfectedFiles

`func (o *DeleteInfectedFileResponse) HasDeleteSucceededInfectedFiles() bool`

HasDeleteSucceededInfectedFiles returns a boolean if a field has been set.

### SetDeleteSucceededInfectedFilesNil

`func (o *DeleteInfectedFileResponse) SetDeleteSucceededInfectedFilesNil(b bool)`

 SetDeleteSucceededInfectedFilesNil sets the value for DeleteSucceededInfectedFiles to be an explicit nil

### UnsetDeleteSucceededInfectedFiles
`func (o *DeleteInfectedFileResponse) UnsetDeleteSucceededInfectedFiles()`

UnsetDeleteSucceededInfectedFiles ensures that no value is present for DeleteSucceededInfectedFiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


