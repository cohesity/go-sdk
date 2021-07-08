# SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllUnderHierarchy** | Pointer to **NullableBool** | Specifies if subtenants of the given tenants should be considered for report generation. | [optional] 
**CompactVersion** | Pointer to **NullableString** | Specifies the Cohesity Agent software version. | [optional] 
**ConsecutiveFailures** | Pointer to **NullableInt32** | Specifies the number of consecutive failures. | [optional] 
**Environment** | Pointer to **NullableString** | Specifies the Environment for the entity being protected. | [optional] 
**GroupBy** | Pointer to **NullableInt32** | Specifies if the report should be grouped by any field. | [optional] 
**HealthStatus** | Pointer to **[]string** | Specifies the Cohesity Agent health status. | [optional] 
**HostOsType** | Pointer to **[]string** | Specifies the OS type on which Cohesity Agent is installed. | [optional] 
**JobId** | Pointer to **NullableInt64** | Specifies the id of the job for which to get the report data. | [optional] 
**JobName** | Pointer to **NullableString** | Specifies the name of the job for which to get the report data. | [optional] 
**LastNDays** | Pointer to **NullableInt32** | Specifies the number of days from current date for which the report data is to be fetched. | [optional] 
**ObjectIds** | Pointer to **[]int64** | Specifies the object ids for which to get the report data. | [optional] 
**ObjectType** | Pointer to **NullableString** | Specifies the object type for which to get the report data. | [optional] 
**RegisteredSourceId** | Pointer to **NullableInt64** | Specifies the registered source for which to get the report data. | [optional] 
**RegisteredSourceIds** | Pointer to **[]int64** | Specifies the registered sources for which to get the report data. | [optional] 
**Rollup** | Pointer to **NullableString** | Specifies the rollup(day/week/month) for protected object trends report. | [optional] 
**RunStatus** | Pointer to **[]string** | Specifies the run status for which to get the report data. | [optional] 
**Sid** | Pointer to **NullableString** | Specifies the sid of the user. | [optional] 
**TenantIdVec** | Pointer to **[]string** | Specifies tenant ids for which report needs to be generated. | [optional] 
**Timezone** | Pointer to **NullableString** | Specifies the timezone. | [optional] 
**UnixUid** | Pointer to **NullableInt32** | Specifies the unix uid of the user. | [optional] 
**VaultIds** | Pointer to **[]int64** | Specifies the vault ids for which to get the report data. | [optional] 
**ViewBoxId** | Pointer to **NullableInt64** | Specifies the view box for which to get the report data. | [optional] 
**ViewName** | Pointer to **NullableString** | Specifies the view name for which the report is required. | [optional] 
**ViewboxIds** | Pointer to **[]int64** | Specifies the viewbox ids to filter by. | [optional] 
**VmName** | Pointer to **NullableString** | Specifies the VM name for which to get the report data. | [optional] 

## Methods

### NewSchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters

`func NewSchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters() *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters`

NewSchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters instantiates a new SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParametersWithDefaults

`func NewSchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParametersWithDefaults() *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters`

NewSchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParametersWithDefaults instantiates a new SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllUnderHierarchy

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetAllUnderHierarchy() bool`

GetAllUnderHierarchy returns the AllUnderHierarchy field if non-nil, zero value otherwise.

### GetAllUnderHierarchyOk

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetAllUnderHierarchyOk() (*bool, bool)`

GetAllUnderHierarchyOk returns a tuple with the AllUnderHierarchy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllUnderHierarchy

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetAllUnderHierarchy(v bool)`

SetAllUnderHierarchy sets AllUnderHierarchy field to given value.

### HasAllUnderHierarchy

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) HasAllUnderHierarchy() bool`

HasAllUnderHierarchy returns a boolean if a field has been set.

### SetAllUnderHierarchyNil

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetAllUnderHierarchyNil(b bool)`

 SetAllUnderHierarchyNil sets the value for AllUnderHierarchy to be an explicit nil

### UnsetAllUnderHierarchy
`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) UnsetAllUnderHierarchy()`

UnsetAllUnderHierarchy ensures that no value is present for AllUnderHierarchy, not even an explicit nil
### GetCompactVersion

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetCompactVersion() string`

GetCompactVersion returns the CompactVersion field if non-nil, zero value otherwise.

### GetCompactVersionOk

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetCompactVersionOk() (*string, bool)`

GetCompactVersionOk returns a tuple with the CompactVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompactVersion

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetCompactVersion(v string)`

SetCompactVersion sets CompactVersion field to given value.

### HasCompactVersion

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) HasCompactVersion() bool`

HasCompactVersion returns a boolean if a field has been set.

### SetCompactVersionNil

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetCompactVersionNil(b bool)`

 SetCompactVersionNil sets the value for CompactVersion to be an explicit nil

### UnsetCompactVersion
`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) UnsetCompactVersion()`

UnsetCompactVersion ensures that no value is present for CompactVersion, not even an explicit nil
### GetConsecutiveFailures

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetConsecutiveFailures() int32`

GetConsecutiveFailures returns the ConsecutiveFailures field if non-nil, zero value otherwise.

### GetConsecutiveFailuresOk

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetConsecutiveFailuresOk() (*int32, bool)`

GetConsecutiveFailuresOk returns a tuple with the ConsecutiveFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutiveFailures

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetConsecutiveFailures(v int32)`

SetConsecutiveFailures sets ConsecutiveFailures field to given value.

### HasConsecutiveFailures

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) HasConsecutiveFailures() bool`

HasConsecutiveFailures returns a boolean if a field has been set.

### SetConsecutiveFailuresNil

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetConsecutiveFailuresNil(b bool)`

 SetConsecutiveFailuresNil sets the value for ConsecutiveFailures to be an explicit nil

### UnsetConsecutiveFailures
`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) UnsetConsecutiveFailures()`

UnsetConsecutiveFailures ensures that no value is present for ConsecutiveFailures, not even an explicit nil
### GetEnvironment

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetGroupBy

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetGroupBy() int32`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetGroupByOk() (*int32, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetGroupBy(v int32)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### SetGroupByNil

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetGroupByNil(b bool)`

 SetGroupByNil sets the value for GroupBy to be an explicit nil

### UnsetGroupBy
`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) UnsetGroupBy()`

UnsetGroupBy ensures that no value is present for GroupBy, not even an explicit nil
### GetHealthStatus

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetHealthStatus() []string`

GetHealthStatus returns the HealthStatus field if non-nil, zero value otherwise.

### GetHealthStatusOk

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetHealthStatusOk() (*[]string, bool)`

GetHealthStatusOk returns a tuple with the HealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatus

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetHealthStatus(v []string)`

SetHealthStatus sets HealthStatus field to given value.

### HasHealthStatus

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) HasHealthStatus() bool`

HasHealthStatus returns a boolean if a field has been set.

### SetHealthStatusNil

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetHealthStatusNil(b bool)`

 SetHealthStatusNil sets the value for HealthStatus to be an explicit nil

### UnsetHealthStatus
`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) UnsetHealthStatus()`

UnsetHealthStatus ensures that no value is present for HealthStatus, not even an explicit nil
### GetHostOsType

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetHostOsType() []string`

GetHostOsType returns the HostOsType field if non-nil, zero value otherwise.

### GetHostOsTypeOk

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetHostOsTypeOk() (*[]string, bool)`

GetHostOsTypeOk returns a tuple with the HostOsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostOsType

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetHostOsType(v []string)`

SetHostOsType sets HostOsType field to given value.

### HasHostOsType

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) HasHostOsType() bool`

HasHostOsType returns a boolean if a field has been set.

### SetHostOsTypeNil

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetHostOsTypeNil(b bool)`

 SetHostOsTypeNil sets the value for HostOsType to be an explicit nil

### UnsetHostOsType
`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) UnsetHostOsType()`

UnsetHostOsType ensures that no value is present for HostOsType, not even an explicit nil
### GetJobId

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetJobId() int64`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetJobIdOk() (*int64, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetJobId(v int64)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### SetJobIdNil

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetJobIdNil(b bool)`

 SetJobIdNil sets the value for JobId to be an explicit nil

### UnsetJobId
`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) UnsetJobId()`

UnsetJobId ensures that no value is present for JobId, not even an explicit nil
### GetJobName

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetJobName() string`

GetJobName returns the JobName field if non-nil, zero value otherwise.

### GetJobNameOk

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetJobNameOk() (*string, bool)`

GetJobNameOk returns a tuple with the JobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobName

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetJobName(v string)`

SetJobName sets JobName field to given value.

### HasJobName

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) HasJobName() bool`

HasJobName returns a boolean if a field has been set.

### SetJobNameNil

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetJobNameNil(b bool)`

 SetJobNameNil sets the value for JobName to be an explicit nil

### UnsetJobName
`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) UnsetJobName()`

UnsetJobName ensures that no value is present for JobName, not even an explicit nil
### GetLastNDays

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetLastNDays() int32`

GetLastNDays returns the LastNDays field if non-nil, zero value otherwise.

### GetLastNDaysOk

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetLastNDaysOk() (*int32, bool)`

GetLastNDaysOk returns a tuple with the LastNDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNDays

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetLastNDays(v int32)`

SetLastNDays sets LastNDays field to given value.

### HasLastNDays

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) HasLastNDays() bool`

HasLastNDays returns a boolean if a field has been set.

### SetLastNDaysNil

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetLastNDaysNil(b bool)`

 SetLastNDaysNil sets the value for LastNDays to be an explicit nil

### UnsetLastNDays
`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) UnsetLastNDays()`

UnsetLastNDays ensures that no value is present for LastNDays, not even an explicit nil
### GetObjectIds

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetObjectIds() []int64`

GetObjectIds returns the ObjectIds field if non-nil, zero value otherwise.

### GetObjectIdsOk

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetObjectIdsOk() (*[]int64, bool)`

GetObjectIdsOk returns a tuple with the ObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectIds

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetObjectIds(v []int64)`

SetObjectIds sets ObjectIds field to given value.

### HasObjectIds

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) HasObjectIds() bool`

HasObjectIds returns a boolean if a field has been set.

### SetObjectIdsNil

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetObjectIdsNil(b bool)`

 SetObjectIdsNil sets the value for ObjectIds to be an explicit nil

### UnsetObjectIds
`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) UnsetObjectIds()`

UnsetObjectIds ensures that no value is present for ObjectIds, not even an explicit nil
### GetObjectType

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### SetObjectTypeNil

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetObjectTypeNil(b bool)`

 SetObjectTypeNil sets the value for ObjectType to be an explicit nil

### UnsetObjectType
`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) UnsetObjectType()`

UnsetObjectType ensures that no value is present for ObjectType, not even an explicit nil
### GetRegisteredSourceId

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetRegisteredSourceId() int64`

GetRegisteredSourceId returns the RegisteredSourceId field if non-nil, zero value otherwise.

### GetRegisteredSourceIdOk

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetRegisteredSourceIdOk() (*int64, bool)`

GetRegisteredSourceIdOk returns a tuple with the RegisteredSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredSourceId

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetRegisteredSourceId(v int64)`

SetRegisteredSourceId sets RegisteredSourceId field to given value.

### HasRegisteredSourceId

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) HasRegisteredSourceId() bool`

HasRegisteredSourceId returns a boolean if a field has been set.

### SetRegisteredSourceIdNil

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetRegisteredSourceIdNil(b bool)`

 SetRegisteredSourceIdNil sets the value for RegisteredSourceId to be an explicit nil

### UnsetRegisteredSourceId
`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) UnsetRegisteredSourceId()`

UnsetRegisteredSourceId ensures that no value is present for RegisteredSourceId, not even an explicit nil
### GetRegisteredSourceIds

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetRegisteredSourceIds() []int64`

GetRegisteredSourceIds returns the RegisteredSourceIds field if non-nil, zero value otherwise.

### GetRegisteredSourceIdsOk

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetRegisteredSourceIdsOk() (*[]int64, bool)`

GetRegisteredSourceIdsOk returns a tuple with the RegisteredSourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredSourceIds

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetRegisteredSourceIds(v []int64)`

SetRegisteredSourceIds sets RegisteredSourceIds field to given value.

### HasRegisteredSourceIds

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) HasRegisteredSourceIds() bool`

HasRegisteredSourceIds returns a boolean if a field has been set.

### SetRegisteredSourceIdsNil

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetRegisteredSourceIdsNil(b bool)`

 SetRegisteredSourceIdsNil sets the value for RegisteredSourceIds to be an explicit nil

### UnsetRegisteredSourceIds
`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) UnsetRegisteredSourceIds()`

UnsetRegisteredSourceIds ensures that no value is present for RegisteredSourceIds, not even an explicit nil
### GetRollup

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetRollup() string`

GetRollup returns the Rollup field if non-nil, zero value otherwise.

### GetRollupOk

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetRollupOk() (*string, bool)`

GetRollupOk returns a tuple with the Rollup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollup

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetRollup(v string)`

SetRollup sets Rollup field to given value.

### HasRollup

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) HasRollup() bool`

HasRollup returns a boolean if a field has been set.

### SetRollupNil

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetRollupNil(b bool)`

 SetRollupNil sets the value for Rollup to be an explicit nil

### UnsetRollup
`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) UnsetRollup()`

UnsetRollup ensures that no value is present for Rollup, not even an explicit nil
### GetRunStatus

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetRunStatus() []string`

GetRunStatus returns the RunStatus field if non-nil, zero value otherwise.

### GetRunStatusOk

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetRunStatusOk() (*[]string, bool)`

GetRunStatusOk returns a tuple with the RunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunStatus

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetRunStatus(v []string)`

SetRunStatus sets RunStatus field to given value.

### HasRunStatus

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) HasRunStatus() bool`

HasRunStatus returns a boolean if a field has been set.

### SetRunStatusNil

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetRunStatusNil(b bool)`

 SetRunStatusNil sets the value for RunStatus to be an explicit nil

### UnsetRunStatus
`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) UnsetRunStatus()`

UnsetRunStatus ensures that no value is present for RunStatus, not even an explicit nil
### GetSid

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetTenantIdVec

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetTenantIdVec() []string`

GetTenantIdVec returns the TenantIdVec field if non-nil, zero value otherwise.

### GetTenantIdVecOk

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetTenantIdVecOk() (*[]string, bool)`

GetTenantIdVecOk returns a tuple with the TenantIdVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdVec

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetTenantIdVec(v []string)`

SetTenantIdVec sets TenantIdVec field to given value.

### HasTenantIdVec

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) HasTenantIdVec() bool`

HasTenantIdVec returns a boolean if a field has been set.

### SetTenantIdVecNil

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetTenantIdVecNil(b bool)`

 SetTenantIdVecNil sets the value for TenantIdVec to be an explicit nil

### UnsetTenantIdVec
`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) UnsetTenantIdVec()`

UnsetTenantIdVec ensures that no value is present for TenantIdVec, not even an explicit nil
### GetTimezone

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetUnixUid

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetUnixUid() int32`

GetUnixUid returns the UnixUid field if non-nil, zero value otherwise.

### GetUnixUidOk

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetUnixUidOk() (*int32, bool)`

GetUnixUidOk returns a tuple with the UnixUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixUid

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetUnixUid(v int32)`

SetUnixUid sets UnixUid field to given value.

### HasUnixUid

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) HasUnixUid() bool`

HasUnixUid returns a boolean if a field has been set.

### SetUnixUidNil

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetUnixUidNil(b bool)`

 SetUnixUidNil sets the value for UnixUid to be an explicit nil

### UnsetUnixUid
`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) UnsetUnixUid()`

UnsetUnixUid ensures that no value is present for UnixUid, not even an explicit nil
### GetVaultIds

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetVaultIds() []int64`

GetVaultIds returns the VaultIds field if non-nil, zero value otherwise.

### GetVaultIdsOk

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetVaultIdsOk() (*[]int64, bool)`

GetVaultIdsOk returns a tuple with the VaultIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultIds

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetVaultIds(v []int64)`

SetVaultIds sets VaultIds field to given value.

### HasVaultIds

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) HasVaultIds() bool`

HasVaultIds returns a boolean if a field has been set.

### SetVaultIdsNil

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetVaultIdsNil(b bool)`

 SetVaultIdsNil sets the value for VaultIds to be an explicit nil

### UnsetVaultIds
`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) UnsetVaultIds()`

UnsetVaultIds ensures that no value is present for VaultIds, not even an explicit nil
### GetViewBoxId

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetViewBoxId() int64`

GetViewBoxId returns the ViewBoxId field if non-nil, zero value otherwise.

### GetViewBoxIdOk

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetViewBoxIdOk() (*int64, bool)`

GetViewBoxIdOk returns a tuple with the ViewBoxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewBoxId

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetViewBoxId(v int64)`

SetViewBoxId sets ViewBoxId field to given value.

### HasViewBoxId

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) HasViewBoxId() bool`

HasViewBoxId returns a boolean if a field has been set.

### SetViewBoxIdNil

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetViewBoxIdNil(b bool)`

 SetViewBoxIdNil sets the value for ViewBoxId to be an explicit nil

### UnsetViewBoxId
`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) UnsetViewBoxId()`

UnsetViewBoxId ensures that no value is present for ViewBoxId, not even an explicit nil
### GetViewName

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil
### GetViewboxIds

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetViewboxIds() []int64`

GetViewboxIds returns the ViewboxIds field if non-nil, zero value otherwise.

### GetViewboxIdsOk

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetViewboxIdsOk() (*[]int64, bool)`

GetViewboxIdsOk returns a tuple with the ViewboxIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewboxIds

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetViewboxIds(v []int64)`

SetViewboxIds sets ViewboxIds field to given value.

### HasViewboxIds

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) HasViewboxIds() bool`

HasViewboxIds returns a boolean if a field has been set.

### SetViewboxIdsNil

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetViewboxIdsNil(b bool)`

 SetViewboxIdsNil sets the value for ViewboxIds to be an explicit nil

### UnsetViewboxIds
`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) UnsetViewboxIds()`

UnsetViewboxIds ensures that no value is present for ViewboxIds, not even an explicit nil
### GetVmName

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetVmName() string`

GetVmName returns the VmName field if non-nil, zero value otherwise.

### GetVmNameOk

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) GetVmNameOk() (*string, bool)`

GetVmNameOk returns a tuple with the VmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmName

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetVmName(v string)`

SetVmName sets VmName field to given value.

### HasVmName

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) HasVmName() bool`

HasVmName returns a boolean if a field has been set.

### SetVmNameNil

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) SetVmNameNil(b bool)`

 SetVmNameNil sets the value for VmName to be an explicit nil

### UnsetVmName
`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters) UnsetVmName()`

UnsetVmName ensures that no value is present for VmName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


