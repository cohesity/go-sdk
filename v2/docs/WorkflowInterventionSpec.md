# WorkflowInterventionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Intervention** | **NullableString** | Specifies the intervention type for ongoing tasks. | 
**WorkflowType** | **NullableString** | Specifies the workflow type for which an intervention would be needed when maintenance mode begins | 

## Methods

### NewWorkflowInterventionSpec

`func NewWorkflowInterventionSpec(intervention NullableString, workflowType NullableString, ) *WorkflowInterventionSpec`

NewWorkflowInterventionSpec instantiates a new WorkflowInterventionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowInterventionSpecWithDefaults

`func NewWorkflowInterventionSpecWithDefaults() *WorkflowInterventionSpec`

NewWorkflowInterventionSpecWithDefaults instantiates a new WorkflowInterventionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntervention

`func (o *WorkflowInterventionSpec) GetIntervention() string`

GetIntervention returns the Intervention field if non-nil, zero value otherwise.

### GetInterventionOk

`func (o *WorkflowInterventionSpec) GetInterventionOk() (*string, bool)`

GetInterventionOk returns a tuple with the Intervention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervention

`func (o *WorkflowInterventionSpec) SetIntervention(v string)`

SetIntervention sets Intervention field to given value.


### SetInterventionNil

`func (o *WorkflowInterventionSpec) SetInterventionNil(b bool)`

 SetInterventionNil sets the value for Intervention to be an explicit nil

### UnsetIntervention
`func (o *WorkflowInterventionSpec) UnsetIntervention()`

UnsetIntervention ensures that no value is present for Intervention, not even an explicit nil
### GetWorkflowType

`func (o *WorkflowInterventionSpec) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *WorkflowInterventionSpec) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *WorkflowInterventionSpec) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.


### SetWorkflowTypeNil

`func (o *WorkflowInterventionSpec) SetWorkflowTypeNil(b bool)`

 SetWorkflowTypeNil sets the value for WorkflowType to be an explicit nil

### UnsetWorkflowType
`func (o *WorkflowInterventionSpec) UnsetWorkflowType()`

UnsetWorkflowType ensures that no value is present for WorkflowType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


