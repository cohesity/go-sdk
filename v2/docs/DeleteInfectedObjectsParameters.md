# DeleteInfectedObjectsParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InfectedObjects** | [**[]InfectedObject**](InfectedObject.md) | Specifies a list of infected objects to be deleted. | 

## Methods

### NewDeleteInfectedObjectsParameters

`func NewDeleteInfectedObjectsParameters(infectedObjects []InfectedObject, ) *DeleteInfectedObjectsParameters`

NewDeleteInfectedObjectsParameters instantiates a new DeleteInfectedObjectsParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteInfectedObjectsParametersWithDefaults

`func NewDeleteInfectedObjectsParametersWithDefaults() *DeleteInfectedObjectsParameters`

NewDeleteInfectedObjectsParametersWithDefaults instantiates a new DeleteInfectedObjectsParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfectedObjects

`func (o *DeleteInfectedObjectsParameters) GetInfectedObjects() []InfectedObject`

GetInfectedObjects returns the InfectedObjects field if non-nil, zero value otherwise.

### GetInfectedObjectsOk

`func (o *DeleteInfectedObjectsParameters) GetInfectedObjectsOk() (*[]InfectedObject, bool)`

GetInfectedObjectsOk returns a tuple with the InfectedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfectedObjects

`func (o *DeleteInfectedObjectsParameters) SetInfectedObjects(v []InfectedObject)`

SetInfectedObjects sets InfectedObjects field to given value.


### SetInfectedObjectsNil

`func (o *DeleteInfectedObjectsParameters) SetInfectedObjectsNil(b bool)`

 SetInfectedObjectsNil sets the value for InfectedObjects to be an explicit nil

### UnsetInfectedObjects
`func (o *DeleteInfectedObjectsParameters) UnsetInfectedObjects()`

UnsetInfectedObjects ensures that no value is present for InfectedObjects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


