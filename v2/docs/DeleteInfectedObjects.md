# DeleteInfectedObjects

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteFailedInfectedObjects** | Pointer to [**[]InfectedObject**](InfectedObject.md) | Specifies the list of infected objects that failed deletion. | [optional] 
**DeleteSucceededInfectedObjects** | Pointer to [**[]InfectedObject**](InfectedObject.md) | Specifies the list of infected objects that are successfully deleted. | [optional] 

## Methods

### NewDeleteInfectedObjects

`func NewDeleteInfectedObjects() *DeleteInfectedObjects`

NewDeleteInfectedObjects instantiates a new DeleteInfectedObjects object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteInfectedObjectsWithDefaults

`func NewDeleteInfectedObjectsWithDefaults() *DeleteInfectedObjects`

NewDeleteInfectedObjectsWithDefaults instantiates a new DeleteInfectedObjects object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteFailedInfectedObjects

`func (o *DeleteInfectedObjects) GetDeleteFailedInfectedObjects() []InfectedObject`

GetDeleteFailedInfectedObjects returns the DeleteFailedInfectedObjects field if non-nil, zero value otherwise.

### GetDeleteFailedInfectedObjectsOk

`func (o *DeleteInfectedObjects) GetDeleteFailedInfectedObjectsOk() (*[]InfectedObject, bool)`

GetDeleteFailedInfectedObjectsOk returns a tuple with the DeleteFailedInfectedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteFailedInfectedObjects

`func (o *DeleteInfectedObjects) SetDeleteFailedInfectedObjects(v []InfectedObject)`

SetDeleteFailedInfectedObjects sets DeleteFailedInfectedObjects field to given value.

### HasDeleteFailedInfectedObjects

`func (o *DeleteInfectedObjects) HasDeleteFailedInfectedObjects() bool`

HasDeleteFailedInfectedObjects returns a boolean if a field has been set.

### SetDeleteFailedInfectedObjectsNil

`func (o *DeleteInfectedObjects) SetDeleteFailedInfectedObjectsNil(b bool)`

 SetDeleteFailedInfectedObjectsNil sets the value for DeleteFailedInfectedObjects to be an explicit nil

### UnsetDeleteFailedInfectedObjects
`func (o *DeleteInfectedObjects) UnsetDeleteFailedInfectedObjects()`

UnsetDeleteFailedInfectedObjects ensures that no value is present for DeleteFailedInfectedObjects, not even an explicit nil
### GetDeleteSucceededInfectedObjects

`func (o *DeleteInfectedObjects) GetDeleteSucceededInfectedObjects() []InfectedObject`

GetDeleteSucceededInfectedObjects returns the DeleteSucceededInfectedObjects field if non-nil, zero value otherwise.

### GetDeleteSucceededInfectedObjectsOk

`func (o *DeleteInfectedObjects) GetDeleteSucceededInfectedObjectsOk() (*[]InfectedObject, bool)`

GetDeleteSucceededInfectedObjectsOk returns a tuple with the DeleteSucceededInfectedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteSucceededInfectedObjects

`func (o *DeleteInfectedObjects) SetDeleteSucceededInfectedObjects(v []InfectedObject)`

SetDeleteSucceededInfectedObjects sets DeleteSucceededInfectedObjects field to given value.

### HasDeleteSucceededInfectedObjects

`func (o *DeleteInfectedObjects) HasDeleteSucceededInfectedObjects() bool`

HasDeleteSucceededInfectedObjects returns a boolean if a field has been set.

### SetDeleteSucceededInfectedObjectsNil

`func (o *DeleteInfectedObjects) SetDeleteSucceededInfectedObjectsNil(b bool)`

 SetDeleteSucceededInfectedObjectsNil sets the value for DeleteSucceededInfectedObjects to be an explicit nil

### UnsetDeleteSucceededInfectedObjects
`func (o *DeleteInfectedObjects) UnsetDeleteSucceededInfectedObjects()`

UnsetDeleteSucceededInfectedObjects ensures that no value is present for DeleteSucceededInfectedObjects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


