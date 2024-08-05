# UpdateInfectedFilesParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InfectedFiles** | [**[]InfectedFile**](InfectedFile.md) | Specifies a list of infected entities to be updated. | 
**State** | Pointer to **NullableString** | Specifies the state[Quarantined, Unquarantined] of the infected entity. | [optional] 

## Methods

### NewUpdateInfectedFilesParameters

`func NewUpdateInfectedFilesParameters(infectedFiles []InfectedFile, ) *UpdateInfectedFilesParameters`

NewUpdateInfectedFilesParameters instantiates a new UpdateInfectedFilesParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInfectedFilesParametersWithDefaults

`func NewUpdateInfectedFilesParametersWithDefaults() *UpdateInfectedFilesParameters`

NewUpdateInfectedFilesParametersWithDefaults instantiates a new UpdateInfectedFilesParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfectedFiles

`func (o *UpdateInfectedFilesParameters) GetInfectedFiles() []InfectedFile`

GetInfectedFiles returns the InfectedFiles field if non-nil, zero value otherwise.

### GetInfectedFilesOk

`func (o *UpdateInfectedFilesParameters) GetInfectedFilesOk() (*[]InfectedFile, bool)`

GetInfectedFilesOk returns a tuple with the InfectedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfectedFiles

`func (o *UpdateInfectedFilesParameters) SetInfectedFiles(v []InfectedFile)`

SetInfectedFiles sets InfectedFiles field to given value.


### SetInfectedFilesNil

`func (o *UpdateInfectedFilesParameters) SetInfectedFilesNil(b bool)`

 SetInfectedFilesNil sets the value for InfectedFiles to be an explicit nil

### UnsetInfectedFiles
`func (o *UpdateInfectedFilesParameters) UnsetInfectedFiles()`

UnsetInfectedFiles ensures that no value is present for InfectedFiles, not even an explicit nil
### GetState

`func (o *UpdateInfectedFilesParameters) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateInfectedFilesParameters) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateInfectedFilesParameters) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UpdateInfectedFilesParameters) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *UpdateInfectedFilesParameters) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *UpdateInfectedFilesParameters) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


