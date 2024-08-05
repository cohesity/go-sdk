# ExperimentalAdapterSourceRegistrationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hosts** | Pointer to **[]string** | Specifies the IPs/hostnames for the nodes forming the Experimental Adapter source cluster. | [optional] 
**WorkflowParams** | Pointer to **NullableString** | Specifies the discover source workflow parameters. This is a stringified JSON representation of the parameters. | [optional] 

## Methods

### NewExperimentalAdapterSourceRegistrationParams

`func NewExperimentalAdapterSourceRegistrationParams() *ExperimentalAdapterSourceRegistrationParams`

NewExperimentalAdapterSourceRegistrationParams instantiates a new ExperimentalAdapterSourceRegistrationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentalAdapterSourceRegistrationParamsWithDefaults

`func NewExperimentalAdapterSourceRegistrationParamsWithDefaults() *ExperimentalAdapterSourceRegistrationParams`

NewExperimentalAdapterSourceRegistrationParamsWithDefaults instantiates a new ExperimentalAdapterSourceRegistrationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHosts

`func (o *ExperimentalAdapterSourceRegistrationParams) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *ExperimentalAdapterSourceRegistrationParams) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *ExperimentalAdapterSourceRegistrationParams) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *ExperimentalAdapterSourceRegistrationParams) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetWorkflowParams

`func (o *ExperimentalAdapterSourceRegistrationParams) GetWorkflowParams() string`

GetWorkflowParams returns the WorkflowParams field if non-nil, zero value otherwise.

### GetWorkflowParamsOk

`func (o *ExperimentalAdapterSourceRegistrationParams) GetWorkflowParamsOk() (*string, bool)`

GetWorkflowParamsOk returns a tuple with the WorkflowParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowParams

`func (o *ExperimentalAdapterSourceRegistrationParams) SetWorkflowParams(v string)`

SetWorkflowParams sets WorkflowParams field to given value.

### HasWorkflowParams

`func (o *ExperimentalAdapterSourceRegistrationParams) HasWorkflowParams() bool`

HasWorkflowParams returns a boolean if a field has been set.

### SetWorkflowParamsNil

`func (o *ExperimentalAdapterSourceRegistrationParams) SetWorkflowParamsNil(b bool)`

 SetWorkflowParamsNil sets the value for WorkflowParams to be an explicit nil

### UnsetWorkflowParams
`func (o *ExperimentalAdapterSourceRegistrationParams) UnsetWorkflowParams()`

UnsetWorkflowParams ensures that no value is present for WorkflowParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


