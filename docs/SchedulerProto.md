# SchedulerProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchedulerJobs** | Pointer to [**[]SchedulerProtoSchedulerJob**](SchedulerProtoSchedulerJob.md) | The array of the various scheduler jobs. | [optional] 

## Methods

### NewSchedulerProto

`func NewSchedulerProto() *SchedulerProto`

NewSchedulerProto instantiates a new SchedulerProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerProtoWithDefaults

`func NewSchedulerProtoWithDefaults() *SchedulerProto`

NewSchedulerProtoWithDefaults instantiates a new SchedulerProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedulerJobs

`func (o *SchedulerProto) GetSchedulerJobs() []SchedulerProtoSchedulerJob`

GetSchedulerJobs returns the SchedulerJobs field if non-nil, zero value otherwise.

### GetSchedulerJobsOk

`func (o *SchedulerProto) GetSchedulerJobsOk() (*[]SchedulerProtoSchedulerJob, bool)`

GetSchedulerJobsOk returns a tuple with the SchedulerJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulerJobs

`func (o *SchedulerProto) SetSchedulerJobs(v []SchedulerProtoSchedulerJob)`

SetSchedulerJobs sets SchedulerJobs field to given value.

### HasSchedulerJobs

`func (o *SchedulerProto) HasSchedulerJobs() bool`

HasSchedulerJobs returns a boolean if a field has been set.

### SetSchedulerJobsNil

`func (o *SchedulerProto) SetSchedulerJobsNil(b bool)`

 SetSchedulerJobsNil sets the value for SchedulerJobs to be an explicit nil

### UnsetSchedulerJobs
`func (o *SchedulerProto) UnsetSchedulerJobs()`

UnsetSchedulerJobs ensures that no value is present for SchedulerJobs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


