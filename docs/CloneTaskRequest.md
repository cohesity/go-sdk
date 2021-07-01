# CloneTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloneViewParameters** | Pointer to [**NullableCloneViewRequest**](CloneViewRequest.md) | Specifies settings for cloning an existing View. This field is required for a &#39;kCloneView&#39; Restore Task. | [optional] 
**ContinueOnError** | Pointer to **NullableBool** | Specifies if the Restore Task should continue when some operations on some objects fail. If true, the Cohesity Cluster ignores intermittent errors and restores as many objects as possible. | [optional] 
**GlacierRetrievalType** | Pointer to **NullableString** | Specifies the way data needs to be retrieved from the external target. This information will be filled in by Iris and Magneto will pass it along to the Icebox as it is to support bulk retrieval from Glacier. Specifies the type of Restore Task.  &#39;kStandard&#39; specifies retrievals that allow to access any of your archives within several hours. Standard retrievals typically complete within 3–5 hours. This is the default option for retrieval requests that do not specify the retrieval option. &#39;kBulk&#39; specifies retrievals that are Glacier’s lowest-cost retrieval option, which can be used to retrieve large amounts, even petabytes, of data inexpensively in a day. Bulk retrieval typically complete within 5–12 hours. &#39;kExpedited&#39; specifies retrievals that allows to quickly access your data when occasional urgent requests for a subset of archives are required. For all but the largest archives (250 MB+), data accessed using Expedited retrievals are typically made available within 1–5 minutes. | [optional] 
**HypervParameters** | Pointer to [**HypervCloneParameters**](HypervCloneParameters.md) |  | [optional] 
**Name** | **NullableString** | Specifies the name of the Restore Task. This field must be set and must be a unique name. | 
**NewParentId** | Pointer to **NullableInt64** | Specify a new registered parent Protection Source. If specified the selected objects are cloned or recovered to this new Protection Source. If not specified, objects are cloned or recovered to the original Protection Source that was managing them. | [optional] 
**Objects** | Pointer to [**[]RestoreObjectDetails**](RestoreObjectDetails.md) | Array of Objects.  Specifies a list of Protection Source objects or Protection Job objects (with specified Protection Source objects). | [optional] 
**TargetViewName** | Pointer to **NullableString** | Specifies the name of the View where the cloned VMs are stored. This field is required for a &#39;kCloneVMs&#39; Restore Task. | [optional] 
**Type** | **NullableString** | Specifies the type of Restore Task such as &#39;kCloneVMs&#39; or &#39;kCloneView&#39;. &#39;kCloneVMs&#39; specifies a Restore Task that clones VMs. &#39;kCloneView&#39; specifies a Restore Task that clones a View. | 
**VlanParameters** | Pointer to [**VlanParameters**](VlanParameters.md) |  | [optional] 
**VmwareParameters** | Pointer to [**VmwareCloneParameters**](VmwareCloneParameters.md) |  | [optional] 

## Methods

### NewCloneTaskRequest

`func NewCloneTaskRequest(name NullableString, type_ NullableString, ) *CloneTaskRequest`

NewCloneTaskRequest instantiates a new CloneTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneTaskRequestWithDefaults

`func NewCloneTaskRequestWithDefaults() *CloneTaskRequest`

NewCloneTaskRequestWithDefaults instantiates a new CloneTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloneViewParameters

`func (o *CloneTaskRequest) GetCloneViewParameters() CloneViewRequest`

GetCloneViewParameters returns the CloneViewParameters field if non-nil, zero value otherwise.

### GetCloneViewParametersOk

`func (o *CloneTaskRequest) GetCloneViewParametersOk() (*CloneViewRequest, bool)`

GetCloneViewParametersOk returns a tuple with the CloneViewParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneViewParameters

`func (o *CloneTaskRequest) SetCloneViewParameters(v CloneViewRequest)`

SetCloneViewParameters sets CloneViewParameters field to given value.

### HasCloneViewParameters

`func (o *CloneTaskRequest) HasCloneViewParameters() bool`

HasCloneViewParameters returns a boolean if a field has been set.

### SetCloneViewParametersNil

`func (o *CloneTaskRequest) SetCloneViewParametersNil(b bool)`

 SetCloneViewParametersNil sets the value for CloneViewParameters to be an explicit nil

### UnsetCloneViewParameters
`func (o *CloneTaskRequest) UnsetCloneViewParameters()`

UnsetCloneViewParameters ensures that no value is present for CloneViewParameters, not even an explicit nil
### GetContinueOnError

`func (o *CloneTaskRequest) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *CloneTaskRequest) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *CloneTaskRequest) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *CloneTaskRequest) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *CloneTaskRequest) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *CloneTaskRequest) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetGlacierRetrievalType

`func (o *CloneTaskRequest) GetGlacierRetrievalType() string`

GetGlacierRetrievalType returns the GlacierRetrievalType field if non-nil, zero value otherwise.

### GetGlacierRetrievalTypeOk

`func (o *CloneTaskRequest) GetGlacierRetrievalTypeOk() (*string, bool)`

GetGlacierRetrievalTypeOk returns a tuple with the GlacierRetrievalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlacierRetrievalType

`func (o *CloneTaskRequest) SetGlacierRetrievalType(v string)`

SetGlacierRetrievalType sets GlacierRetrievalType field to given value.

### HasGlacierRetrievalType

`func (o *CloneTaskRequest) HasGlacierRetrievalType() bool`

HasGlacierRetrievalType returns a boolean if a field has been set.

### SetGlacierRetrievalTypeNil

`func (o *CloneTaskRequest) SetGlacierRetrievalTypeNil(b bool)`

 SetGlacierRetrievalTypeNil sets the value for GlacierRetrievalType to be an explicit nil

### UnsetGlacierRetrievalType
`func (o *CloneTaskRequest) UnsetGlacierRetrievalType()`

UnsetGlacierRetrievalType ensures that no value is present for GlacierRetrievalType, not even an explicit nil
### GetHypervParameters

`func (o *CloneTaskRequest) GetHypervParameters() HypervCloneParameters`

GetHypervParameters returns the HypervParameters field if non-nil, zero value otherwise.

### GetHypervParametersOk

`func (o *CloneTaskRequest) GetHypervParametersOk() (*HypervCloneParameters, bool)`

GetHypervParametersOk returns a tuple with the HypervParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervParameters

`func (o *CloneTaskRequest) SetHypervParameters(v HypervCloneParameters)`

SetHypervParameters sets HypervParameters field to given value.

### HasHypervParameters

`func (o *CloneTaskRequest) HasHypervParameters() bool`

HasHypervParameters returns a boolean if a field has been set.

### GetName

`func (o *CloneTaskRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloneTaskRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloneTaskRequest) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CloneTaskRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CloneTaskRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNewParentId

`func (o *CloneTaskRequest) GetNewParentId() int64`

GetNewParentId returns the NewParentId field if non-nil, zero value otherwise.

### GetNewParentIdOk

`func (o *CloneTaskRequest) GetNewParentIdOk() (*int64, bool)`

GetNewParentIdOk returns a tuple with the NewParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewParentId

`func (o *CloneTaskRequest) SetNewParentId(v int64)`

SetNewParentId sets NewParentId field to given value.

### HasNewParentId

`func (o *CloneTaskRequest) HasNewParentId() bool`

HasNewParentId returns a boolean if a field has been set.

### SetNewParentIdNil

`func (o *CloneTaskRequest) SetNewParentIdNil(b bool)`

 SetNewParentIdNil sets the value for NewParentId to be an explicit nil

### UnsetNewParentId
`func (o *CloneTaskRequest) UnsetNewParentId()`

UnsetNewParentId ensures that no value is present for NewParentId, not even an explicit nil
### GetObjects

`func (o *CloneTaskRequest) GetObjects() []RestoreObjectDetails`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *CloneTaskRequest) GetObjectsOk() (*[]RestoreObjectDetails, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *CloneTaskRequest) SetObjects(v []RestoreObjectDetails)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *CloneTaskRequest) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *CloneTaskRequest) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *CloneTaskRequest) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetTargetViewName

`func (o *CloneTaskRequest) GetTargetViewName() string`

GetTargetViewName returns the TargetViewName field if non-nil, zero value otherwise.

### GetTargetViewNameOk

`func (o *CloneTaskRequest) GetTargetViewNameOk() (*string, bool)`

GetTargetViewNameOk returns a tuple with the TargetViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetViewName

`func (o *CloneTaskRequest) SetTargetViewName(v string)`

SetTargetViewName sets TargetViewName field to given value.

### HasTargetViewName

`func (o *CloneTaskRequest) HasTargetViewName() bool`

HasTargetViewName returns a boolean if a field has been set.

### SetTargetViewNameNil

`func (o *CloneTaskRequest) SetTargetViewNameNil(b bool)`

 SetTargetViewNameNil sets the value for TargetViewName to be an explicit nil

### UnsetTargetViewName
`func (o *CloneTaskRequest) UnsetTargetViewName()`

UnsetTargetViewName ensures that no value is present for TargetViewName, not even an explicit nil
### GetType

`func (o *CloneTaskRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloneTaskRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloneTaskRequest) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CloneTaskRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CloneTaskRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetVlanParameters

`func (o *CloneTaskRequest) GetVlanParameters() VlanParameters`

GetVlanParameters returns the VlanParameters field if non-nil, zero value otherwise.

### GetVlanParametersOk

`func (o *CloneTaskRequest) GetVlanParametersOk() (*VlanParameters, bool)`

GetVlanParametersOk returns a tuple with the VlanParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanParameters

`func (o *CloneTaskRequest) SetVlanParameters(v VlanParameters)`

SetVlanParameters sets VlanParameters field to given value.

### HasVlanParameters

`func (o *CloneTaskRequest) HasVlanParameters() bool`

HasVlanParameters returns a boolean if a field has been set.

### GetVmwareParameters

`func (o *CloneTaskRequest) GetVmwareParameters() VmwareCloneParameters`

GetVmwareParameters returns the VmwareParameters field if non-nil, zero value otherwise.

### GetVmwareParametersOk

`func (o *CloneTaskRequest) GetVmwareParametersOk() (*VmwareCloneParameters, bool)`

GetVmwareParametersOk returns a tuple with the VmwareParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParameters

`func (o *CloneTaskRequest) SetVmwareParameters(v VmwareCloneParameters)`

SetVmwareParameters sets VmwareParameters field to given value.

### HasVmwareParameters

`func (o *CloneTaskRequest) HasVmwareParameters() bool`

HasVmwareParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


