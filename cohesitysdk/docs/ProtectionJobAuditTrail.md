# ProtectionJobAuditTrail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**After** | Pointer to [**ProtectionJob**](ProtectionJob.md) |  | [optional] 
**Before** | Pointer to [**ProtectionJob**](ProtectionJob.md) |  | [optional] 
**Changes** | Pointer to **[]string** | Specifies the list of changed values in a Protection Job. kProtectionJobName implies that protection job has change in the name field kProtectionJobDescription implies that protection job has change in the description field. kProtectionJobSources implies that protection job has change in the source field. kProtectionJobSchedule implies that protection job has change in the schedule field. kProtectionJobFullSchedule implies that protection job has change in the full schedule field. kProtectionJobRetrySettings implies that protection job has change in the retry settings. kProtectionJobRetentionPolicy implies that protection job has change in the retention policy. kProtectionJobIndexingPolicy implies that protection job has change in the indexing policy. kProtectionJobAlertingPolicy implies that protection job has change in the alerting policy. kProtectionJobPriority implies that protection job has change in the alerting policy. kProtectionJobQuiesce implies that protection job has change in the Quiesce. kProtectionJobSla implies that protection job has change in the SLA settings. kProtectionJobPolicyId implies that protection job has change in the poilcy Id settings. kProtectionJobTimezone implies that protection job has change in the timezone settings. kProtectionJobFutureRunsPaused implies that protection job has change in the future run settings. kProtectionJobFutureRunsResumed implies that protection job has change in the future run resume settings. kSnapshotTargetPolicy implies that protection job has change in the snapshot target policy settings. kProtectionJobQOS implies that protection job has change in QOS settings. kProtectionJobInvalidField implies that the changed field is invalid. | [optional] 

## Methods

### NewProtectionJobAuditTrail

`func NewProtectionJobAuditTrail() *ProtectionJobAuditTrail`

NewProtectionJobAuditTrail instantiates a new ProtectionJobAuditTrail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionJobAuditTrailWithDefaults

`func NewProtectionJobAuditTrailWithDefaults() *ProtectionJobAuditTrail`

NewProtectionJobAuditTrailWithDefaults instantiates a new ProtectionJobAuditTrail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfter

`func (o *ProtectionJobAuditTrail) GetAfter() ProtectionJob`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *ProtectionJobAuditTrail) GetAfterOk() (*ProtectionJob, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *ProtectionJobAuditTrail) SetAfter(v ProtectionJob)`

SetAfter sets After field to given value.

### HasAfter

`func (o *ProtectionJobAuditTrail) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetBefore

`func (o *ProtectionJobAuditTrail) GetBefore() ProtectionJob`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *ProtectionJobAuditTrail) GetBeforeOk() (*ProtectionJob, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *ProtectionJobAuditTrail) SetBefore(v ProtectionJob)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *ProtectionJobAuditTrail) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### GetChanges

`func (o *ProtectionJobAuditTrail) GetChanges() []string`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *ProtectionJobAuditTrail) GetChangesOk() (*[]string, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *ProtectionJobAuditTrail) SetChanges(v []string)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *ProtectionJobAuditTrail) HasChanges() bool`

HasChanges returns a boolean if a field has been set.

### SetChangesNil

`func (o *ProtectionJobAuditTrail) SetChangesNil(b bool)`

 SetChangesNil sets the value for Changes to be an explicit nil

### UnsetChanges
`func (o *ProtectionJobAuditTrail) UnsetChanges()`

UnsetChanges ensures that no value is present for Changes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


