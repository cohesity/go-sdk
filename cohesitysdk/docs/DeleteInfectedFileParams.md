# DeleteInfectedFileParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InfectedFileIds** | Pointer to [**[]InfectedFileParam**](InfectedFileParam.md) | Specifies the list of infected file path. | [optional] 

## Methods

### NewDeleteInfectedFileParams

`func NewDeleteInfectedFileParams() *DeleteInfectedFileParams`

NewDeleteInfectedFileParams instantiates a new DeleteInfectedFileParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteInfectedFileParamsWithDefaults

`func NewDeleteInfectedFileParamsWithDefaults() *DeleteInfectedFileParams`

NewDeleteInfectedFileParamsWithDefaults instantiates a new DeleteInfectedFileParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfectedFileIds

`func (o *DeleteInfectedFileParams) GetInfectedFileIds() []InfectedFileParam`

GetInfectedFileIds returns the InfectedFileIds field if non-nil, zero value otherwise.

### GetInfectedFileIdsOk

`func (o *DeleteInfectedFileParams) GetInfectedFileIdsOk() (*[]InfectedFileParam, bool)`

GetInfectedFileIdsOk returns a tuple with the InfectedFileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfectedFileIds

`func (o *DeleteInfectedFileParams) SetInfectedFileIds(v []InfectedFileParam)`

SetInfectedFileIds sets InfectedFileIds field to given value.

### HasInfectedFileIds

`func (o *DeleteInfectedFileParams) HasInfectedFileIds() bool`

HasInfectedFileIds returns a boolean if a field has been set.

### SetInfectedFileIdsNil

`func (o *DeleteInfectedFileParams) SetInfectedFileIdsNil(b bool)`

 SetInfectedFileIdsNil sets the value for InfectedFileIds to be an explicit nil

### UnsetInfectedFileIds
`func (o *DeleteInfectedFileParams) UnsetInfectedFileIds()`

UnsetInfectedFileIds ensures that no value is present for InfectedFileIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


