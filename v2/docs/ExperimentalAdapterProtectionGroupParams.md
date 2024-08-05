# ExperimentalAdapterProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Concurrency** | Pointer to **NullableInt32** | Specifies the maximum number of concurrent IO streams that will be created to exchange data with the cluster. If not specified, the default value is taken as 1. | [optional] [default to 1]
**ExcludedObjectIds** | Pointer to **[]int32** | Specifies the objects to be excluded in the Protection Group. | [optional] 
**Objects** | Pointer to [**[]ExperimentalAdapaterProtectionGroupObjectParams**](ExperimentalAdapaterProtectionGroupObjectParams.md) | Specifies a list of fully qualified names of the objects to be protected. | [optional] 
**WorkflowParams** | Pointer to **NullableString** | Specifies the discover source workflow parameters. This is a stringified JSON representation of the parameters. | [optional] 

## Methods

### NewExperimentalAdapterProtectionGroupParams

`func NewExperimentalAdapterProtectionGroupParams() *ExperimentalAdapterProtectionGroupParams`

NewExperimentalAdapterProtectionGroupParams instantiates a new ExperimentalAdapterProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentalAdapterProtectionGroupParamsWithDefaults

`func NewExperimentalAdapterProtectionGroupParamsWithDefaults() *ExperimentalAdapterProtectionGroupParams`

NewExperimentalAdapterProtectionGroupParamsWithDefaults instantiates a new ExperimentalAdapterProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConcurrency

`func (o *ExperimentalAdapterProtectionGroupParams) GetConcurrency() int32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *ExperimentalAdapterProtectionGroupParams) GetConcurrencyOk() (*int32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *ExperimentalAdapterProtectionGroupParams) SetConcurrency(v int32)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *ExperimentalAdapterProtectionGroupParams) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### SetConcurrencyNil

`func (o *ExperimentalAdapterProtectionGroupParams) SetConcurrencyNil(b bool)`

 SetConcurrencyNil sets the value for Concurrency to be an explicit nil

### UnsetConcurrency
`func (o *ExperimentalAdapterProtectionGroupParams) UnsetConcurrency()`

UnsetConcurrency ensures that no value is present for Concurrency, not even an explicit nil
### GetExcludedObjectIds

`func (o *ExperimentalAdapterProtectionGroupParams) GetExcludedObjectIds() []int32`

GetExcludedObjectIds returns the ExcludedObjectIds field if non-nil, zero value otherwise.

### GetExcludedObjectIdsOk

`func (o *ExperimentalAdapterProtectionGroupParams) GetExcludedObjectIdsOk() (*[]int32, bool)`

GetExcludedObjectIdsOk returns a tuple with the ExcludedObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedObjectIds

`func (o *ExperimentalAdapterProtectionGroupParams) SetExcludedObjectIds(v []int32)`

SetExcludedObjectIds sets ExcludedObjectIds field to given value.

### HasExcludedObjectIds

`func (o *ExperimentalAdapterProtectionGroupParams) HasExcludedObjectIds() bool`

HasExcludedObjectIds returns a boolean if a field has been set.

### SetExcludedObjectIdsNil

`func (o *ExperimentalAdapterProtectionGroupParams) SetExcludedObjectIdsNil(b bool)`

 SetExcludedObjectIdsNil sets the value for ExcludedObjectIds to be an explicit nil

### UnsetExcludedObjectIds
`func (o *ExperimentalAdapterProtectionGroupParams) UnsetExcludedObjectIds()`

UnsetExcludedObjectIds ensures that no value is present for ExcludedObjectIds, not even an explicit nil
### GetObjects

`func (o *ExperimentalAdapterProtectionGroupParams) GetObjects() []ExperimentalAdapaterProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *ExperimentalAdapterProtectionGroupParams) GetObjectsOk() (*[]ExperimentalAdapaterProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *ExperimentalAdapterProtectionGroupParams) SetObjects(v []ExperimentalAdapaterProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *ExperimentalAdapterProtectionGroupParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetWorkflowParams

`func (o *ExperimentalAdapterProtectionGroupParams) GetWorkflowParams() string`

GetWorkflowParams returns the WorkflowParams field if non-nil, zero value otherwise.

### GetWorkflowParamsOk

`func (o *ExperimentalAdapterProtectionGroupParams) GetWorkflowParamsOk() (*string, bool)`

GetWorkflowParamsOk returns a tuple with the WorkflowParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowParams

`func (o *ExperimentalAdapterProtectionGroupParams) SetWorkflowParams(v string)`

SetWorkflowParams sets WorkflowParams field to given value.

### HasWorkflowParams

`func (o *ExperimentalAdapterProtectionGroupParams) HasWorkflowParams() bool`

HasWorkflowParams returns a boolean if a field has been set.

### SetWorkflowParamsNil

`func (o *ExperimentalAdapterProtectionGroupParams) SetWorkflowParamsNil(b bool)`

 SetWorkflowParamsNil sets the value for WorkflowParams to be an explicit nil

### UnsetWorkflowParams
`func (o *ExperimentalAdapterProtectionGroupParams) UnsetWorkflowParams()`

UnsetWorkflowParams ensures that no value is present for WorkflowParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


