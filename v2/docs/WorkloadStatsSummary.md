# WorkloadStatsSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workloads** | Pointer to [**[]WorkloadStatsSchema**](WorkloadStatsSchema.md) | Specifies the Workload types. | [optional] 

## Methods

### NewWorkloadStatsSummary

`func NewWorkloadStatsSummary() *WorkloadStatsSummary`

NewWorkloadStatsSummary instantiates a new WorkloadStatsSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadStatsSummaryWithDefaults

`func NewWorkloadStatsSummaryWithDefaults() *WorkloadStatsSummary`

NewWorkloadStatsSummaryWithDefaults instantiates a new WorkloadStatsSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkloads

`func (o *WorkloadStatsSummary) GetWorkloads() []WorkloadStatsSchema`

GetWorkloads returns the Workloads field if non-nil, zero value otherwise.

### GetWorkloadsOk

`func (o *WorkloadStatsSummary) GetWorkloadsOk() (*[]WorkloadStatsSchema, bool)`

GetWorkloadsOk returns a tuple with the Workloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloads

`func (o *WorkloadStatsSummary) SetWorkloads(v []WorkloadStatsSchema)`

SetWorkloads sets Workloads field to given value.

### HasWorkloads

`func (o *WorkloadStatsSummary) HasWorkloads() bool`

HasWorkloads returns a boolean if a field has been set.

### SetWorkloadsNil

`func (o *WorkloadStatsSummary) SetWorkloadsNil(b bool)`

 SetWorkloadsNil sets the value for Workloads to be an explicit nil

### UnsetWorkloads
`func (o *WorkloadStatsSummary) UnsetWorkloads()`

UnsetWorkloads ensures that no value is present for Workloads, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


