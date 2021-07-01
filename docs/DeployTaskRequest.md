# DeployTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Specifies the name of the Deploy Task. This field must be set and must be a unique name. | 
**NewParentId** | Pointer to **NullableInt64** | Specifies a new registered parent Protection Source. If specified the selected objects are cloned or recovered to this new Protection Source. If not specified, objects are cloned or recovered to the original Protection Source that was managing them. | [optional] 
**Objects** | Pointer to [**[]RestoreObjectDetails**](RestoreObjectDetails.md) | Array of Objects.  Specifies a list of Protection Source objects or Protection Job objects (with specified Protection Source objects). | [optional] 
**Target** | Pointer to [**CloudDeployTargetDetails**](CloudDeployTargetDetails.md) |  | [optional] 

## Methods

### NewDeployTaskRequest

`func NewDeployTaskRequest(name NullableString, ) *DeployTaskRequest`

NewDeployTaskRequest instantiates a new DeployTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployTaskRequestWithDefaults

`func NewDeployTaskRequestWithDefaults() *DeployTaskRequest`

NewDeployTaskRequestWithDefaults instantiates a new DeployTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeployTaskRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeployTaskRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeployTaskRequest) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *DeployTaskRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DeployTaskRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNewParentId

`func (o *DeployTaskRequest) GetNewParentId() int64`

GetNewParentId returns the NewParentId field if non-nil, zero value otherwise.

### GetNewParentIdOk

`func (o *DeployTaskRequest) GetNewParentIdOk() (*int64, bool)`

GetNewParentIdOk returns a tuple with the NewParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewParentId

`func (o *DeployTaskRequest) SetNewParentId(v int64)`

SetNewParentId sets NewParentId field to given value.

### HasNewParentId

`func (o *DeployTaskRequest) HasNewParentId() bool`

HasNewParentId returns a boolean if a field has been set.

### SetNewParentIdNil

`func (o *DeployTaskRequest) SetNewParentIdNil(b bool)`

 SetNewParentIdNil sets the value for NewParentId to be an explicit nil

### UnsetNewParentId
`func (o *DeployTaskRequest) UnsetNewParentId()`

UnsetNewParentId ensures that no value is present for NewParentId, not even an explicit nil
### GetObjects

`func (o *DeployTaskRequest) GetObjects() []RestoreObjectDetails`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *DeployTaskRequest) GetObjectsOk() (*[]RestoreObjectDetails, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *DeployTaskRequest) SetObjects(v []RestoreObjectDetails)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *DeployTaskRequest) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *DeployTaskRequest) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *DeployTaskRequest) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetTarget

`func (o *DeployTaskRequest) GetTarget() CloudDeployTargetDetails`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *DeployTaskRequest) GetTargetOk() (*CloudDeployTargetDetails, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *DeployTaskRequest) SetTarget(v CloudDeployTargetDetails)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *DeployTaskRequest) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


