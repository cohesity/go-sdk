# SchedulerProtoSchedulerJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableRecurringEmail** | Pointer to **NullableBool** | The boolean which specifies if this job is to be scheduled or not. | [optional] 
**Id** | Pointer to **NullableInt64** | The unique id for the scheduled job assigned by the cluster. | [optional] 
**Name** | Pointer to **NullableString** | The name of the scheduled job given by the user. | [optional] 
**ScheduleJobParameters** | Pointer to [**SchedulerProtoSchedulerJobScheduleJobParameters**](SchedulerProtoSchedulerJobScheduleJobParameters.md) |  | [optional] 
**Schedules** | Pointer to [**[]SchedulerProtoSchedulerJobSchedule**](SchedulerProtoSchedulerJobSchedule.md) | The frequency of schedule execution. | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies id of tenant who created the scheduler job. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of the job. The enum which defines the Job type of the job. | [optional] 

## Methods

### NewSchedulerProtoSchedulerJob

`func NewSchedulerProtoSchedulerJob() *SchedulerProtoSchedulerJob`

NewSchedulerProtoSchedulerJob instantiates a new SchedulerProtoSchedulerJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerProtoSchedulerJobWithDefaults

`func NewSchedulerProtoSchedulerJobWithDefaults() *SchedulerProtoSchedulerJob`

NewSchedulerProtoSchedulerJobWithDefaults instantiates a new SchedulerProtoSchedulerJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableRecurringEmail

`func (o *SchedulerProtoSchedulerJob) GetEnableRecurringEmail() bool`

GetEnableRecurringEmail returns the EnableRecurringEmail field if non-nil, zero value otherwise.

### GetEnableRecurringEmailOk

`func (o *SchedulerProtoSchedulerJob) GetEnableRecurringEmailOk() (*bool, bool)`

GetEnableRecurringEmailOk returns a tuple with the EnableRecurringEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRecurringEmail

`func (o *SchedulerProtoSchedulerJob) SetEnableRecurringEmail(v bool)`

SetEnableRecurringEmail sets EnableRecurringEmail field to given value.

### HasEnableRecurringEmail

`func (o *SchedulerProtoSchedulerJob) HasEnableRecurringEmail() bool`

HasEnableRecurringEmail returns a boolean if a field has been set.

### SetEnableRecurringEmailNil

`func (o *SchedulerProtoSchedulerJob) SetEnableRecurringEmailNil(b bool)`

 SetEnableRecurringEmailNil sets the value for EnableRecurringEmail to be an explicit nil

### UnsetEnableRecurringEmail
`func (o *SchedulerProtoSchedulerJob) UnsetEnableRecurringEmail()`

UnsetEnableRecurringEmail ensures that no value is present for EnableRecurringEmail, not even an explicit nil
### GetId

`func (o *SchedulerProtoSchedulerJob) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SchedulerProtoSchedulerJob) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SchedulerProtoSchedulerJob) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SchedulerProtoSchedulerJob) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SchedulerProtoSchedulerJob) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SchedulerProtoSchedulerJob) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *SchedulerProtoSchedulerJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SchedulerProtoSchedulerJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SchedulerProtoSchedulerJob) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SchedulerProtoSchedulerJob) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SchedulerProtoSchedulerJob) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SchedulerProtoSchedulerJob) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetScheduleJobParameters

`func (o *SchedulerProtoSchedulerJob) GetScheduleJobParameters() SchedulerProtoSchedulerJobScheduleJobParameters`

GetScheduleJobParameters returns the ScheduleJobParameters field if non-nil, zero value otherwise.

### GetScheduleJobParametersOk

`func (o *SchedulerProtoSchedulerJob) GetScheduleJobParametersOk() (*SchedulerProtoSchedulerJobScheduleJobParameters, bool)`

GetScheduleJobParametersOk returns a tuple with the ScheduleJobParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleJobParameters

`func (o *SchedulerProtoSchedulerJob) SetScheduleJobParameters(v SchedulerProtoSchedulerJobScheduleJobParameters)`

SetScheduleJobParameters sets ScheduleJobParameters field to given value.

### HasScheduleJobParameters

`func (o *SchedulerProtoSchedulerJob) HasScheduleJobParameters() bool`

HasScheduleJobParameters returns a boolean if a field has been set.

### GetSchedules

`func (o *SchedulerProtoSchedulerJob) GetSchedules() []SchedulerProtoSchedulerJobSchedule`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *SchedulerProtoSchedulerJob) GetSchedulesOk() (*[]SchedulerProtoSchedulerJobSchedule, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *SchedulerProtoSchedulerJob) SetSchedules(v []SchedulerProtoSchedulerJobSchedule)`

SetSchedules sets Schedules field to given value.

### HasSchedules

`func (o *SchedulerProtoSchedulerJob) HasSchedules() bool`

HasSchedules returns a boolean if a field has been set.

### SetSchedulesNil

`func (o *SchedulerProtoSchedulerJob) SetSchedulesNil(b bool)`

 SetSchedulesNil sets the value for Schedules to be an explicit nil

### UnsetSchedules
`func (o *SchedulerProtoSchedulerJob) UnsetSchedules()`

UnsetSchedules ensures that no value is present for Schedules, not even an explicit nil
### GetTenantId

`func (o *SchedulerProtoSchedulerJob) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SchedulerProtoSchedulerJob) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SchedulerProtoSchedulerJob) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SchedulerProtoSchedulerJob) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *SchedulerProtoSchedulerJob) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *SchedulerProtoSchedulerJob) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetType

`func (o *SchedulerProtoSchedulerJob) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SchedulerProtoSchedulerJob) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SchedulerProtoSchedulerJob) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SchedulerProtoSchedulerJob) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SchedulerProtoSchedulerJob) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SchedulerProtoSchedulerJob) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


