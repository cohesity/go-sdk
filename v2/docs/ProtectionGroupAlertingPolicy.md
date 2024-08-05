# ProtectionGroupAlertingPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertTargets** | Pointer to [**[]AlertTarget**](AlertTarget.md) | Specifies a list of targets to receive the alerts. | [optional] 
**BackupRunStatus** | **[]string** | Specifies when to send out alerts. The possible values are kSuccess , kFailure, kSlaViolation and kWarning | 
**RaiseObjectLevelFailureAlert** | Pointer to **bool** | Specifies whether object level alerts are raised for backup failures after the backup run. | [optional] 
**RaiseObjectLevelFailureAlertAfterEachAttempt** | Pointer to **bool** | Specifies whether object level alerts are raised for backup failures after each backup attempt. | [optional] 
**RaiseObjectLevelFailureAlertAfterLastAttempt** | Pointer to **bool** | Specifies whether object level alerts are raised for backup failures after last backup attempt. | [optional] 

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


### GetRaiseObjectLevelFailureAlert

`func (o *ProtectionGroupAlertingPolicy) GetRaiseObjectLevelFailureAlert() bool`

GetRaiseObjectLevelFailureAlert returns the RaiseObjectLevelFailureAlert field if non-nil, zero value otherwise.

### GetRaiseObjectLevelFailureAlertOk

`func (o *ProtectionGroupAlertingPolicy) GetRaiseObjectLevelFailureAlertOk() (*bool, bool)`

GetRaiseObjectLevelFailureAlertOk returns a tuple with the RaiseObjectLevelFailureAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaiseObjectLevelFailureAlert

`func (o *ProtectionGroupAlertingPolicy) SetRaiseObjectLevelFailureAlert(v bool)`

SetRaiseObjectLevelFailureAlert sets RaiseObjectLevelFailureAlert field to given value.

### HasRaiseObjectLevelFailureAlert

`func (o *ProtectionGroupAlertingPolicy) HasRaiseObjectLevelFailureAlert() bool`

HasRaiseObjectLevelFailureAlert returns a boolean if a field has been set.

### GetRaiseObjectLevelFailureAlertAfterEachAttempt

`func (o *ProtectionGroupAlertingPolicy) GetRaiseObjectLevelFailureAlertAfterEachAttempt() bool`

GetRaiseObjectLevelFailureAlertAfterEachAttempt returns the RaiseObjectLevelFailureAlertAfterEachAttempt field if non-nil, zero value otherwise.

### GetRaiseObjectLevelFailureAlertAfterEachAttemptOk

`func (o *ProtectionGroupAlertingPolicy) GetRaiseObjectLevelFailureAlertAfterEachAttemptOk() (*bool, bool)`

GetRaiseObjectLevelFailureAlertAfterEachAttemptOk returns a tuple with the RaiseObjectLevelFailureAlertAfterEachAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaiseObjectLevelFailureAlertAfterEachAttempt

`func (o *ProtectionGroupAlertingPolicy) SetRaiseObjectLevelFailureAlertAfterEachAttempt(v bool)`

SetRaiseObjectLevelFailureAlertAfterEachAttempt sets RaiseObjectLevelFailureAlertAfterEachAttempt field to given value.

### HasRaiseObjectLevelFailureAlertAfterEachAttempt

`func (o *ProtectionGroupAlertingPolicy) HasRaiseObjectLevelFailureAlertAfterEachAttempt() bool`

HasRaiseObjectLevelFailureAlertAfterEachAttempt returns a boolean if a field has been set.

### GetRaiseObjectLevelFailureAlertAfterLastAttempt

`func (o *ProtectionGroupAlertingPolicy) GetRaiseObjectLevelFailureAlertAfterLastAttempt() bool`

GetRaiseObjectLevelFailureAlertAfterLastAttempt returns the RaiseObjectLevelFailureAlertAfterLastAttempt field if non-nil, zero value otherwise.

### GetRaiseObjectLevelFailureAlertAfterLastAttemptOk

`func (o *ProtectionGroupAlertingPolicy) GetRaiseObjectLevelFailureAlertAfterLastAttemptOk() (*bool, bool)`

GetRaiseObjectLevelFailureAlertAfterLastAttemptOk returns a tuple with the RaiseObjectLevelFailureAlertAfterLastAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaiseObjectLevelFailureAlertAfterLastAttempt

`func (o *ProtectionGroupAlertingPolicy) SetRaiseObjectLevelFailureAlertAfterLastAttempt(v bool)`

SetRaiseObjectLevelFailureAlertAfterLastAttempt sets RaiseObjectLevelFailureAlertAfterLastAttempt field to given value.

### HasRaiseObjectLevelFailureAlertAfterLastAttempt

`func (o *ProtectionGroupAlertingPolicy) HasRaiseObjectLevelFailureAlertAfterLastAttempt() bool`

HasRaiseObjectLevelFailureAlertAfterLastAttempt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


