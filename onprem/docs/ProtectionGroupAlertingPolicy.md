# ProtectionGroupAlertingPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupRunStatus** | **[]string** | Specifies the run status for which the user would like to receive alerts. | 
**AlertTargets** | Pointer to [**[]AlertTarget**](AlertTarget.md) | Specifies a list of targets to receive the alerts. | [optional] 

## Methods

### NewProtectionGroupAlertingPolicy

`func NewProtectionGroupAlertingPolicy(backupRunStatus []string, ) *ProtectionGroupAlertingPolicy`

NewProtectionGroupAlertingPolicy instantiates a new ProtectionGroupAlertingPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionGroupAlertingPolicyWithDefaults

`func NewProtectionGroupAlertingPolicyWithDefaults() *ProtectionGroupAlertingPolicy`

NewProtectionGroupAlertingPolicyWithDefaults instantiates a new ProtectionGroupAlertingPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupRunStatus

`func (o *ProtectionGroupAlertingPolicy) GetBackupRunStatus() []string`

GetBackupRunStatus returns the BackupRunStatus field if non-nil, zero value otherwise.

### GetBackupRunStatusOk

`func (o *ProtectionGroupAlertingPolicy) GetBackupRunStatusOk() (*[]string, bool)`

GetBackupRunStatusOk returns a tuple with the BackupRunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRunStatus

`func (o *ProtectionGroupAlertingPolicy) SetBackupRunStatus(v []string)`

SetBackupRunStatus sets BackupRunStatus field to given value.


### GetAlertTargets

`func (o *ProtectionGroupAlertingPolicy) GetAlertTargets() []AlertTarget`

GetAlertTargets returns the AlertTargets field if non-nil, zero value otherwise.

### GetAlertTargetsOk

`func (o *ProtectionGroupAlertingPolicy) GetAlertTargetsOk() (*[]AlertTarget, bool)`

GetAlertTargetsOk returns a tuple with the AlertTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertTargets

`func (o *ProtectionGroupAlertingPolicy) SetAlertTargets(v []AlertTarget)`

SetAlertTargets sets AlertTargets field to given value.

### HasAlertTargets

`func (o *ProtectionGroupAlertingPolicy) HasAlertTargets() bool`

HasAlertTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


