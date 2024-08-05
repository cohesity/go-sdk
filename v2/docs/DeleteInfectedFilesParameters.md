# DeleteInfectedFilesParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InfectedFiles** | [**[]InfectedFile**](InfectedFile.md) | Specifies a list of infected files to be deleted. | 

## Methods

### NewDeleteInfectedFilesParameters

`func NewDeleteInfectedFilesParameters(infectedFiles []InfectedFile, ) *DeleteInfectedFilesParameters`

NewDeleteInfectedFilesParameters instantiates a new DeleteInfectedFilesParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteInfectedFilesParametersWithDefaults

`func NewDeleteInfectedFilesParametersWithDefaults() *DeleteInfectedFilesParameters`

NewDeleteInfectedFilesParametersWithDefaults instantiates a new DeleteInfectedFilesParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfectedFiles

`func (o *DeleteInfectedFilesParameters) GetInfectedFiles() []InfectedFile`

GetInfectedFiles returns the InfectedFiles field if non-nil, zero value otherwise.

### GetInfectedFilesOk

`func (o *DeleteInfectedFilesParameters) GetInfectedFilesOk() (*[]InfectedFile, bool)`

GetInfectedFilesOk returns a tuple with the InfectedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfectedFiles

`func (o *DeleteInfectedFilesParameters) SetInfectedFiles(v []InfectedFile)`

SetInfectedFiles sets InfectedFiles field to given value.


### SetInfectedFilesNil

`func (o *DeleteInfectedFilesParameters) SetInfectedFilesNil(b bool)`

 SetInfectedFilesNil sets the value for InfectedFiles to be an explicit nil

### UnsetInfectedFiles
`func (o *DeleteInfectedFilesParameters) UnsetInfectedFiles()`

UnsetInfectedFiles ensures that no value is present for InfectedFiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


