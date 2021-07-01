# ProtectionJobSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **NullableInt64** | Specifies the id of the cluster on which object is protected. | [optional] 
**ClusterIncarnationId** | Pointer to **NullableInt64** | Specifies the incarnation id of the cluster on which object is protected. | [optional] 
**IsRpoJob** | Pointer to **NullableBool** | Specifies if the Protection Job is created by an RPO policy. | [optional] 
**JobId** | Pointer to **NullableInt64** | Specifies the id of the Protection Job. | [optional] 
**JobName** | Pointer to **NullableString** | Specifies the name of the Protection Job. | [optional] 
**LastProtectionJobRunStatus** | Pointer to **NullableInt32** | Specifies the last job run status. | [optional] 
**PolicyId** | Pointer to **NullableString** | Specifies the id of the policy that is used by a Protection Job. Format of policy id will be like following: clusterid:clusterincarnationid:policyid. | [optional] 
**PolicyName** | Pointer to **NullableString** | Specifies the name of the policy that is used by a Protection Job. | [optional] 

## Methods

### NewProtectionJobSummary

`func NewProtectionJobSummary() *ProtectionJobSummary`

NewProtectionJobSummary instantiates a new ProtectionJobSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionJobSummaryWithDefaults

`func NewProtectionJobSummaryWithDefaults() *ProtectionJobSummary`

NewProtectionJobSummaryWithDefaults instantiates a new ProtectionJobSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *ProtectionJobSummary) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ProtectionJobSummary) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ProtectionJobSummary) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ProtectionJobSummary) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *ProtectionJobSummary) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *ProtectionJobSummary) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *ProtectionJobSummary) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *ProtectionJobSummary) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *ProtectionJobSummary) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *ProtectionJobSummary) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *ProtectionJobSummary) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *ProtectionJobSummary) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetIsRpoJob

`func (o *ProtectionJobSummary) GetIsRpoJob() bool`

GetIsRpoJob returns the IsRpoJob field if non-nil, zero value otherwise.

### GetIsRpoJobOk

`func (o *ProtectionJobSummary) GetIsRpoJobOk() (*bool, bool)`

GetIsRpoJobOk returns a tuple with the IsRpoJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRpoJob

`func (o *ProtectionJobSummary) SetIsRpoJob(v bool)`

SetIsRpoJob sets IsRpoJob field to given value.

### HasIsRpoJob

`func (o *ProtectionJobSummary) HasIsRpoJob() bool`

HasIsRpoJob returns a boolean if a field has been set.

### SetIsRpoJobNil

`func (o *ProtectionJobSummary) SetIsRpoJobNil(b bool)`

 SetIsRpoJobNil sets the value for IsRpoJob to be an explicit nil

### UnsetIsRpoJob
`func (o *ProtectionJobSummary) UnsetIsRpoJob()`

UnsetIsRpoJob ensures that no value is present for IsRpoJob, not even an explicit nil
### GetJobId

`func (o *ProtectionJobSummary) GetJobId() int64`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *ProtectionJobSummary) GetJobIdOk() (*int64, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *ProtectionJobSummary) SetJobId(v int64)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *ProtectionJobSummary) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### SetJobIdNil

`func (o *ProtectionJobSummary) SetJobIdNil(b bool)`

 SetJobIdNil sets the value for JobId to be an explicit nil

### UnsetJobId
`func (o *ProtectionJobSummary) UnsetJobId()`

UnsetJobId ensures that no value is present for JobId, not even an explicit nil
### GetJobName

`func (o *ProtectionJobSummary) GetJobName() string`

GetJobName returns the JobName field if non-nil, zero value otherwise.

### GetJobNameOk

`func (o *ProtectionJobSummary) GetJobNameOk() (*string, bool)`

GetJobNameOk returns a tuple with the JobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobName

`func (o *ProtectionJobSummary) SetJobName(v string)`

SetJobName sets JobName field to given value.

### HasJobName

`func (o *ProtectionJobSummary) HasJobName() bool`

HasJobName returns a boolean if a field has been set.

### SetJobNameNil

`func (o *ProtectionJobSummary) SetJobNameNil(b bool)`

 SetJobNameNil sets the value for JobName to be an explicit nil

### UnsetJobName
`func (o *ProtectionJobSummary) UnsetJobName()`

UnsetJobName ensures that no value is present for JobName, not even an explicit nil
### GetLastProtectionJobRunStatus

`func (o *ProtectionJobSummary) GetLastProtectionJobRunStatus() int32`

GetLastProtectionJobRunStatus returns the LastProtectionJobRunStatus field if non-nil, zero value otherwise.

### GetLastProtectionJobRunStatusOk

`func (o *ProtectionJobSummary) GetLastProtectionJobRunStatusOk() (*int32, bool)`

GetLastProtectionJobRunStatusOk returns a tuple with the LastProtectionJobRunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProtectionJobRunStatus

`func (o *ProtectionJobSummary) SetLastProtectionJobRunStatus(v int32)`

SetLastProtectionJobRunStatus sets LastProtectionJobRunStatus field to given value.

### HasLastProtectionJobRunStatus

`func (o *ProtectionJobSummary) HasLastProtectionJobRunStatus() bool`

HasLastProtectionJobRunStatus returns a boolean if a field has been set.

### SetLastProtectionJobRunStatusNil

`func (o *ProtectionJobSummary) SetLastProtectionJobRunStatusNil(b bool)`

 SetLastProtectionJobRunStatusNil sets the value for LastProtectionJobRunStatus to be an explicit nil

### UnsetLastProtectionJobRunStatus
`func (o *ProtectionJobSummary) UnsetLastProtectionJobRunStatus()`

UnsetLastProtectionJobRunStatus ensures that no value is present for LastProtectionJobRunStatus, not even an explicit nil
### GetPolicyId

`func (o *ProtectionJobSummary) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ProtectionJobSummary) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ProtectionJobSummary) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *ProtectionJobSummary) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *ProtectionJobSummary) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *ProtectionJobSummary) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetPolicyName

`func (o *ProtectionJobSummary) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *ProtectionJobSummary) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *ProtectionJobSummary) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *ProtectionJobSummary) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *ProtectionJobSummary) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *ProtectionJobSummary) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


