# ProtectionSourcesSummaryStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstFailedRunTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of the first failed Job Run protecting this Protection Source Object. The time is specified as a Unix epoch Timestamp (in microseconds). | [optional] 
**FirstSuccessfulRunTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of the first successful Job Run protecting this Protection Source Object. The time is specified as a Unix epoch Timestamp (in microseconds). | [optional] 
**LastFailedRunTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of the last failed Job Run protecting this Protection Source Object. The time is specified as a Unix epoch Timestamp (in microseconds). | [optional] 
**LastRunEndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of the last Job Run protecting this Protection Source Object. The time is specified as a Unix epoch Timestamp (in microseconds). | [optional] 
**LastRunErrorMsg** | Pointer to **NullableString** | Specifies the error message associated with last run, if the last run has failed. | [optional] 
**LastRunStartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of the last Job Run protecting this Protection Source Object. The time is specified as a Unix epoch Timestamp (in microseconds). | [optional] 
**LastRunStatus** | Pointer to **NullableString** | Specifies the Job Run status of the last Job Run protecting this Protection Source Object. &#39;kSuccess&#39; indicates that the Job Run was successful. &#39;kRunning&#39; indicates that the Job Run is currently running. &#39;kWarning&#39; indicates that the Job Run was successful but warnings were issued. &#39;kCancelled&#39; indicates that the Job Run was canceled. &#39;kError&#39; indicates the Job Run encountered an error and did not run to completion. | [optional] 
**LastRunType** | Pointer to **NullableString** | Specifies the Job Run type of the last Job Run protecting this Protection Source Object. &#39;kRegular&#39; indicates an incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a kRegular schedule captures all the blocks. &#39;kFull&#39; indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized. &#39;kLog&#39; indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time. &#39;kSystem&#39; indicates system volume backup. It produces an image for bare metal recovery. | [optional] 
**LastSuccessfulRunTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of the last successful Job Run protecting this Protection Source Object. The time is specified as a Unix epoch Timestamp (in microseconds). | [optional] 
**NumDataReadBytes** | Pointer to **NullableInt64** | Specifies the total number of bytes read while protecting this Protection Source Object. | [optional] 
**NumErrors** | Pointer to **NullableInt32** | Specifies the total number of errors reported during Job Runs of this Protection Source Object. | [optional] 
**NumLogicalBytesProtected** | Pointer to **NullableInt64** | Specifies the total logical bytes protected for this Protection Source Object. The logical size is when the data is fully hydrated or expanded. | [optional] 
**NumSnapshots** | Pointer to **NullableInt32** | Specifies the total number of Snapshots that are backing up this Protection Source Object. | [optional] 
**NumSuccessRuns** | Pointer to **NullableInt32** | Specifies the total number of successful Job Runs protecting this Protection Source Object. | [optional] 
**NumWarnings** | Pointer to **NullableInt32** | Specifies the total number of warnings reported during Job Runs of this Protection Source Object. | [optional] 
**ProtectionSource** | Pointer to [**NullableProtectionSource**](ProtectionSource.md) | Specifies the leaf Protection Source Object (such as VM). Snapshot summary statistics are reported for this Protection Source Object. | [optional] 
**RegisteredSource** | Pointer to **NullableString** | Specifies the name of the Registered Source that is the top level parent of the specified Protection Source Object. | [optional] 
**Tenants** | Pointer to [**[]TenantInfo**](TenantInfo.md) | Specifies basic information about tenants having access to the protection job. | [optional] 

## Methods

### NewProtectionSourcesSummaryStats

`func NewProtectionSourcesSummaryStats() *ProtectionSourcesSummaryStats`

NewProtectionSourcesSummaryStats instantiates a new ProtectionSourcesSummaryStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionSourcesSummaryStatsWithDefaults

`func NewProtectionSourcesSummaryStatsWithDefaults() *ProtectionSourcesSummaryStats`

NewProtectionSourcesSummaryStatsWithDefaults instantiates a new ProtectionSourcesSummaryStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstFailedRunTimeUsecs

`func (o *ProtectionSourcesSummaryStats) GetFirstFailedRunTimeUsecs() int64`

GetFirstFailedRunTimeUsecs returns the FirstFailedRunTimeUsecs field if non-nil, zero value otherwise.

### GetFirstFailedRunTimeUsecsOk

`func (o *ProtectionSourcesSummaryStats) GetFirstFailedRunTimeUsecsOk() (*int64, bool)`

GetFirstFailedRunTimeUsecsOk returns a tuple with the FirstFailedRunTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstFailedRunTimeUsecs

`func (o *ProtectionSourcesSummaryStats) SetFirstFailedRunTimeUsecs(v int64)`

SetFirstFailedRunTimeUsecs sets FirstFailedRunTimeUsecs field to given value.

### HasFirstFailedRunTimeUsecs

`func (o *ProtectionSourcesSummaryStats) HasFirstFailedRunTimeUsecs() bool`

HasFirstFailedRunTimeUsecs returns a boolean if a field has been set.

### SetFirstFailedRunTimeUsecsNil

`func (o *ProtectionSourcesSummaryStats) SetFirstFailedRunTimeUsecsNil(b bool)`

 SetFirstFailedRunTimeUsecsNil sets the value for FirstFailedRunTimeUsecs to be an explicit nil

### UnsetFirstFailedRunTimeUsecs
`func (o *ProtectionSourcesSummaryStats) UnsetFirstFailedRunTimeUsecs()`

UnsetFirstFailedRunTimeUsecs ensures that no value is present for FirstFailedRunTimeUsecs, not even an explicit nil
### GetFirstSuccessfulRunTimeUsecs

`func (o *ProtectionSourcesSummaryStats) GetFirstSuccessfulRunTimeUsecs() int64`

GetFirstSuccessfulRunTimeUsecs returns the FirstSuccessfulRunTimeUsecs field if non-nil, zero value otherwise.

### GetFirstSuccessfulRunTimeUsecsOk

`func (o *ProtectionSourcesSummaryStats) GetFirstSuccessfulRunTimeUsecsOk() (*int64, bool)`

GetFirstSuccessfulRunTimeUsecsOk returns a tuple with the FirstSuccessfulRunTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSuccessfulRunTimeUsecs

`func (o *ProtectionSourcesSummaryStats) SetFirstSuccessfulRunTimeUsecs(v int64)`

SetFirstSuccessfulRunTimeUsecs sets FirstSuccessfulRunTimeUsecs field to given value.

### HasFirstSuccessfulRunTimeUsecs

`func (o *ProtectionSourcesSummaryStats) HasFirstSuccessfulRunTimeUsecs() bool`

HasFirstSuccessfulRunTimeUsecs returns a boolean if a field has been set.

### SetFirstSuccessfulRunTimeUsecsNil

`func (o *ProtectionSourcesSummaryStats) SetFirstSuccessfulRunTimeUsecsNil(b bool)`

 SetFirstSuccessfulRunTimeUsecsNil sets the value for FirstSuccessfulRunTimeUsecs to be an explicit nil

### UnsetFirstSuccessfulRunTimeUsecs
`func (o *ProtectionSourcesSummaryStats) UnsetFirstSuccessfulRunTimeUsecs()`

UnsetFirstSuccessfulRunTimeUsecs ensures that no value is present for FirstSuccessfulRunTimeUsecs, not even an explicit nil
### GetLastFailedRunTimeUsecs

`func (o *ProtectionSourcesSummaryStats) GetLastFailedRunTimeUsecs() int64`

GetLastFailedRunTimeUsecs returns the LastFailedRunTimeUsecs field if non-nil, zero value otherwise.

### GetLastFailedRunTimeUsecsOk

`func (o *ProtectionSourcesSummaryStats) GetLastFailedRunTimeUsecsOk() (*int64, bool)`

GetLastFailedRunTimeUsecsOk returns a tuple with the LastFailedRunTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailedRunTimeUsecs

`func (o *ProtectionSourcesSummaryStats) SetLastFailedRunTimeUsecs(v int64)`

SetLastFailedRunTimeUsecs sets LastFailedRunTimeUsecs field to given value.

### HasLastFailedRunTimeUsecs

`func (o *ProtectionSourcesSummaryStats) HasLastFailedRunTimeUsecs() bool`

HasLastFailedRunTimeUsecs returns a boolean if a field has been set.

### SetLastFailedRunTimeUsecsNil

`func (o *ProtectionSourcesSummaryStats) SetLastFailedRunTimeUsecsNil(b bool)`

 SetLastFailedRunTimeUsecsNil sets the value for LastFailedRunTimeUsecs to be an explicit nil

### UnsetLastFailedRunTimeUsecs
`func (o *ProtectionSourcesSummaryStats) UnsetLastFailedRunTimeUsecs()`

UnsetLastFailedRunTimeUsecs ensures that no value is present for LastFailedRunTimeUsecs, not even an explicit nil
### GetLastRunEndTimeUsecs

`func (o *ProtectionSourcesSummaryStats) GetLastRunEndTimeUsecs() int64`

GetLastRunEndTimeUsecs returns the LastRunEndTimeUsecs field if non-nil, zero value otherwise.

### GetLastRunEndTimeUsecsOk

`func (o *ProtectionSourcesSummaryStats) GetLastRunEndTimeUsecsOk() (*int64, bool)`

GetLastRunEndTimeUsecsOk returns a tuple with the LastRunEndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunEndTimeUsecs

`func (o *ProtectionSourcesSummaryStats) SetLastRunEndTimeUsecs(v int64)`

SetLastRunEndTimeUsecs sets LastRunEndTimeUsecs field to given value.

### HasLastRunEndTimeUsecs

`func (o *ProtectionSourcesSummaryStats) HasLastRunEndTimeUsecs() bool`

HasLastRunEndTimeUsecs returns a boolean if a field has been set.

### SetLastRunEndTimeUsecsNil

`func (o *ProtectionSourcesSummaryStats) SetLastRunEndTimeUsecsNil(b bool)`

 SetLastRunEndTimeUsecsNil sets the value for LastRunEndTimeUsecs to be an explicit nil

### UnsetLastRunEndTimeUsecs
`func (o *ProtectionSourcesSummaryStats) UnsetLastRunEndTimeUsecs()`

UnsetLastRunEndTimeUsecs ensures that no value is present for LastRunEndTimeUsecs, not even an explicit nil
### GetLastRunErrorMsg

`func (o *ProtectionSourcesSummaryStats) GetLastRunErrorMsg() string`

GetLastRunErrorMsg returns the LastRunErrorMsg field if non-nil, zero value otherwise.

### GetLastRunErrorMsgOk

`func (o *ProtectionSourcesSummaryStats) GetLastRunErrorMsgOk() (*string, bool)`

GetLastRunErrorMsgOk returns a tuple with the LastRunErrorMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunErrorMsg

`func (o *ProtectionSourcesSummaryStats) SetLastRunErrorMsg(v string)`

SetLastRunErrorMsg sets LastRunErrorMsg field to given value.

### HasLastRunErrorMsg

`func (o *ProtectionSourcesSummaryStats) HasLastRunErrorMsg() bool`

HasLastRunErrorMsg returns a boolean if a field has been set.

### SetLastRunErrorMsgNil

`func (o *ProtectionSourcesSummaryStats) SetLastRunErrorMsgNil(b bool)`

 SetLastRunErrorMsgNil sets the value for LastRunErrorMsg to be an explicit nil

### UnsetLastRunErrorMsg
`func (o *ProtectionSourcesSummaryStats) UnsetLastRunErrorMsg()`

UnsetLastRunErrorMsg ensures that no value is present for LastRunErrorMsg, not even an explicit nil
### GetLastRunStartTimeUsecs

`func (o *ProtectionSourcesSummaryStats) GetLastRunStartTimeUsecs() int64`

GetLastRunStartTimeUsecs returns the LastRunStartTimeUsecs field if non-nil, zero value otherwise.

### GetLastRunStartTimeUsecsOk

`func (o *ProtectionSourcesSummaryStats) GetLastRunStartTimeUsecsOk() (*int64, bool)`

GetLastRunStartTimeUsecsOk returns a tuple with the LastRunStartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunStartTimeUsecs

`func (o *ProtectionSourcesSummaryStats) SetLastRunStartTimeUsecs(v int64)`

SetLastRunStartTimeUsecs sets LastRunStartTimeUsecs field to given value.

### HasLastRunStartTimeUsecs

`func (o *ProtectionSourcesSummaryStats) HasLastRunStartTimeUsecs() bool`

HasLastRunStartTimeUsecs returns a boolean if a field has been set.

### SetLastRunStartTimeUsecsNil

`func (o *ProtectionSourcesSummaryStats) SetLastRunStartTimeUsecsNil(b bool)`

 SetLastRunStartTimeUsecsNil sets the value for LastRunStartTimeUsecs to be an explicit nil

### UnsetLastRunStartTimeUsecs
`func (o *ProtectionSourcesSummaryStats) UnsetLastRunStartTimeUsecs()`

UnsetLastRunStartTimeUsecs ensures that no value is present for LastRunStartTimeUsecs, not even an explicit nil
### GetLastRunStatus

`func (o *ProtectionSourcesSummaryStats) GetLastRunStatus() string`

GetLastRunStatus returns the LastRunStatus field if non-nil, zero value otherwise.

### GetLastRunStatusOk

`func (o *ProtectionSourcesSummaryStats) GetLastRunStatusOk() (*string, bool)`

GetLastRunStatusOk returns a tuple with the LastRunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunStatus

`func (o *ProtectionSourcesSummaryStats) SetLastRunStatus(v string)`

SetLastRunStatus sets LastRunStatus field to given value.

### HasLastRunStatus

`func (o *ProtectionSourcesSummaryStats) HasLastRunStatus() bool`

HasLastRunStatus returns a boolean if a field has been set.

### SetLastRunStatusNil

`func (o *ProtectionSourcesSummaryStats) SetLastRunStatusNil(b bool)`

 SetLastRunStatusNil sets the value for LastRunStatus to be an explicit nil

### UnsetLastRunStatus
`func (o *ProtectionSourcesSummaryStats) UnsetLastRunStatus()`

UnsetLastRunStatus ensures that no value is present for LastRunStatus, not even an explicit nil
### GetLastRunType

`func (o *ProtectionSourcesSummaryStats) GetLastRunType() string`

GetLastRunType returns the LastRunType field if non-nil, zero value otherwise.

### GetLastRunTypeOk

`func (o *ProtectionSourcesSummaryStats) GetLastRunTypeOk() (*string, bool)`

GetLastRunTypeOk returns a tuple with the LastRunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunType

`func (o *ProtectionSourcesSummaryStats) SetLastRunType(v string)`

SetLastRunType sets LastRunType field to given value.

### HasLastRunType

`func (o *ProtectionSourcesSummaryStats) HasLastRunType() bool`

HasLastRunType returns a boolean if a field has been set.

### SetLastRunTypeNil

`func (o *ProtectionSourcesSummaryStats) SetLastRunTypeNil(b bool)`

 SetLastRunTypeNil sets the value for LastRunType to be an explicit nil

### UnsetLastRunType
`func (o *ProtectionSourcesSummaryStats) UnsetLastRunType()`

UnsetLastRunType ensures that no value is present for LastRunType, not even an explicit nil
### GetLastSuccessfulRunTimeUsecs

`func (o *ProtectionSourcesSummaryStats) GetLastSuccessfulRunTimeUsecs() int64`

GetLastSuccessfulRunTimeUsecs returns the LastSuccessfulRunTimeUsecs field if non-nil, zero value otherwise.

### GetLastSuccessfulRunTimeUsecsOk

`func (o *ProtectionSourcesSummaryStats) GetLastSuccessfulRunTimeUsecsOk() (*int64, bool)`

GetLastSuccessfulRunTimeUsecsOk returns a tuple with the LastSuccessfulRunTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulRunTimeUsecs

`func (o *ProtectionSourcesSummaryStats) SetLastSuccessfulRunTimeUsecs(v int64)`

SetLastSuccessfulRunTimeUsecs sets LastSuccessfulRunTimeUsecs field to given value.

### HasLastSuccessfulRunTimeUsecs

`func (o *ProtectionSourcesSummaryStats) HasLastSuccessfulRunTimeUsecs() bool`

HasLastSuccessfulRunTimeUsecs returns a boolean if a field has been set.

### SetLastSuccessfulRunTimeUsecsNil

`func (o *ProtectionSourcesSummaryStats) SetLastSuccessfulRunTimeUsecsNil(b bool)`

 SetLastSuccessfulRunTimeUsecsNil sets the value for LastSuccessfulRunTimeUsecs to be an explicit nil

### UnsetLastSuccessfulRunTimeUsecs
`func (o *ProtectionSourcesSummaryStats) UnsetLastSuccessfulRunTimeUsecs()`

UnsetLastSuccessfulRunTimeUsecs ensures that no value is present for LastSuccessfulRunTimeUsecs, not even an explicit nil
### GetNumDataReadBytes

`func (o *ProtectionSourcesSummaryStats) GetNumDataReadBytes() int64`

GetNumDataReadBytes returns the NumDataReadBytes field if non-nil, zero value otherwise.

### GetNumDataReadBytesOk

`func (o *ProtectionSourcesSummaryStats) GetNumDataReadBytesOk() (*int64, bool)`

GetNumDataReadBytesOk returns a tuple with the NumDataReadBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDataReadBytes

`func (o *ProtectionSourcesSummaryStats) SetNumDataReadBytes(v int64)`

SetNumDataReadBytes sets NumDataReadBytes field to given value.

### HasNumDataReadBytes

`func (o *ProtectionSourcesSummaryStats) HasNumDataReadBytes() bool`

HasNumDataReadBytes returns a boolean if a field has been set.

### SetNumDataReadBytesNil

`func (o *ProtectionSourcesSummaryStats) SetNumDataReadBytesNil(b bool)`

 SetNumDataReadBytesNil sets the value for NumDataReadBytes to be an explicit nil

### UnsetNumDataReadBytes
`func (o *ProtectionSourcesSummaryStats) UnsetNumDataReadBytes()`

UnsetNumDataReadBytes ensures that no value is present for NumDataReadBytes, not even an explicit nil
### GetNumErrors

`func (o *ProtectionSourcesSummaryStats) GetNumErrors() int32`

GetNumErrors returns the NumErrors field if non-nil, zero value otherwise.

### GetNumErrorsOk

`func (o *ProtectionSourcesSummaryStats) GetNumErrorsOk() (*int32, bool)`

GetNumErrorsOk returns a tuple with the NumErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumErrors

`func (o *ProtectionSourcesSummaryStats) SetNumErrors(v int32)`

SetNumErrors sets NumErrors field to given value.

### HasNumErrors

`func (o *ProtectionSourcesSummaryStats) HasNumErrors() bool`

HasNumErrors returns a boolean if a field has been set.

### SetNumErrorsNil

`func (o *ProtectionSourcesSummaryStats) SetNumErrorsNil(b bool)`

 SetNumErrorsNil sets the value for NumErrors to be an explicit nil

### UnsetNumErrors
`func (o *ProtectionSourcesSummaryStats) UnsetNumErrors()`

UnsetNumErrors ensures that no value is present for NumErrors, not even an explicit nil
### GetNumLogicalBytesProtected

`func (o *ProtectionSourcesSummaryStats) GetNumLogicalBytesProtected() int64`

GetNumLogicalBytesProtected returns the NumLogicalBytesProtected field if non-nil, zero value otherwise.

### GetNumLogicalBytesProtectedOk

`func (o *ProtectionSourcesSummaryStats) GetNumLogicalBytesProtectedOk() (*int64, bool)`

GetNumLogicalBytesProtectedOk returns a tuple with the NumLogicalBytesProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLogicalBytesProtected

`func (o *ProtectionSourcesSummaryStats) SetNumLogicalBytesProtected(v int64)`

SetNumLogicalBytesProtected sets NumLogicalBytesProtected field to given value.

### HasNumLogicalBytesProtected

`func (o *ProtectionSourcesSummaryStats) HasNumLogicalBytesProtected() bool`

HasNumLogicalBytesProtected returns a boolean if a field has been set.

### SetNumLogicalBytesProtectedNil

`func (o *ProtectionSourcesSummaryStats) SetNumLogicalBytesProtectedNil(b bool)`

 SetNumLogicalBytesProtectedNil sets the value for NumLogicalBytesProtected to be an explicit nil

### UnsetNumLogicalBytesProtected
`func (o *ProtectionSourcesSummaryStats) UnsetNumLogicalBytesProtected()`

UnsetNumLogicalBytesProtected ensures that no value is present for NumLogicalBytesProtected, not even an explicit nil
### GetNumSnapshots

`func (o *ProtectionSourcesSummaryStats) GetNumSnapshots() int32`

GetNumSnapshots returns the NumSnapshots field if non-nil, zero value otherwise.

### GetNumSnapshotsOk

`func (o *ProtectionSourcesSummaryStats) GetNumSnapshotsOk() (*int32, bool)`

GetNumSnapshotsOk returns a tuple with the NumSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSnapshots

`func (o *ProtectionSourcesSummaryStats) SetNumSnapshots(v int32)`

SetNumSnapshots sets NumSnapshots field to given value.

### HasNumSnapshots

`func (o *ProtectionSourcesSummaryStats) HasNumSnapshots() bool`

HasNumSnapshots returns a boolean if a field has been set.

### SetNumSnapshotsNil

`func (o *ProtectionSourcesSummaryStats) SetNumSnapshotsNil(b bool)`

 SetNumSnapshotsNil sets the value for NumSnapshots to be an explicit nil

### UnsetNumSnapshots
`func (o *ProtectionSourcesSummaryStats) UnsetNumSnapshots()`

UnsetNumSnapshots ensures that no value is present for NumSnapshots, not even an explicit nil
### GetNumSuccessRuns

`func (o *ProtectionSourcesSummaryStats) GetNumSuccessRuns() int32`

GetNumSuccessRuns returns the NumSuccessRuns field if non-nil, zero value otherwise.

### GetNumSuccessRunsOk

`func (o *ProtectionSourcesSummaryStats) GetNumSuccessRunsOk() (*int32, bool)`

GetNumSuccessRunsOk returns a tuple with the NumSuccessRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSuccessRuns

`func (o *ProtectionSourcesSummaryStats) SetNumSuccessRuns(v int32)`

SetNumSuccessRuns sets NumSuccessRuns field to given value.

### HasNumSuccessRuns

`func (o *ProtectionSourcesSummaryStats) HasNumSuccessRuns() bool`

HasNumSuccessRuns returns a boolean if a field has been set.

### SetNumSuccessRunsNil

`func (o *ProtectionSourcesSummaryStats) SetNumSuccessRunsNil(b bool)`

 SetNumSuccessRunsNil sets the value for NumSuccessRuns to be an explicit nil

### UnsetNumSuccessRuns
`func (o *ProtectionSourcesSummaryStats) UnsetNumSuccessRuns()`

UnsetNumSuccessRuns ensures that no value is present for NumSuccessRuns, not even an explicit nil
### GetNumWarnings

`func (o *ProtectionSourcesSummaryStats) GetNumWarnings() int32`

GetNumWarnings returns the NumWarnings field if non-nil, zero value otherwise.

### GetNumWarningsOk

`func (o *ProtectionSourcesSummaryStats) GetNumWarningsOk() (*int32, bool)`

GetNumWarningsOk returns a tuple with the NumWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumWarnings

`func (o *ProtectionSourcesSummaryStats) SetNumWarnings(v int32)`

SetNumWarnings sets NumWarnings field to given value.

### HasNumWarnings

`func (o *ProtectionSourcesSummaryStats) HasNumWarnings() bool`

HasNumWarnings returns a boolean if a field has been set.

### SetNumWarningsNil

`func (o *ProtectionSourcesSummaryStats) SetNumWarningsNil(b bool)`

 SetNumWarningsNil sets the value for NumWarnings to be an explicit nil

### UnsetNumWarnings
`func (o *ProtectionSourcesSummaryStats) UnsetNumWarnings()`

UnsetNumWarnings ensures that no value is present for NumWarnings, not even an explicit nil
### GetProtectionSource

`func (o *ProtectionSourcesSummaryStats) GetProtectionSource() ProtectionSource`

GetProtectionSource returns the ProtectionSource field if non-nil, zero value otherwise.

### GetProtectionSourceOk

`func (o *ProtectionSourcesSummaryStats) GetProtectionSourceOk() (*ProtectionSource, bool)`

GetProtectionSourceOk returns a tuple with the ProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionSource

`func (o *ProtectionSourcesSummaryStats) SetProtectionSource(v ProtectionSource)`

SetProtectionSource sets ProtectionSource field to given value.

### HasProtectionSource

`func (o *ProtectionSourcesSummaryStats) HasProtectionSource() bool`

HasProtectionSource returns a boolean if a field has been set.

### SetProtectionSourceNil

`func (o *ProtectionSourcesSummaryStats) SetProtectionSourceNil(b bool)`

 SetProtectionSourceNil sets the value for ProtectionSource to be an explicit nil

### UnsetProtectionSource
`func (o *ProtectionSourcesSummaryStats) UnsetProtectionSource()`

UnsetProtectionSource ensures that no value is present for ProtectionSource, not even an explicit nil
### GetRegisteredSource

`func (o *ProtectionSourcesSummaryStats) GetRegisteredSource() string`

GetRegisteredSource returns the RegisteredSource field if non-nil, zero value otherwise.

### GetRegisteredSourceOk

`func (o *ProtectionSourcesSummaryStats) GetRegisteredSourceOk() (*string, bool)`

GetRegisteredSourceOk returns a tuple with the RegisteredSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredSource

`func (o *ProtectionSourcesSummaryStats) SetRegisteredSource(v string)`

SetRegisteredSource sets RegisteredSource field to given value.

### HasRegisteredSource

`func (o *ProtectionSourcesSummaryStats) HasRegisteredSource() bool`

HasRegisteredSource returns a boolean if a field has been set.

### SetRegisteredSourceNil

`func (o *ProtectionSourcesSummaryStats) SetRegisteredSourceNil(b bool)`

 SetRegisteredSourceNil sets the value for RegisteredSource to be an explicit nil

### UnsetRegisteredSource
`func (o *ProtectionSourcesSummaryStats) UnsetRegisteredSource()`

UnsetRegisteredSource ensures that no value is present for RegisteredSource, not even an explicit nil
### GetTenants

`func (o *ProtectionSourcesSummaryStats) GetTenants() []TenantInfo`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ProtectionSourcesSummaryStats) GetTenantsOk() (*[]TenantInfo, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ProtectionSourcesSummaryStats) SetTenants(v []TenantInfo)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ProtectionSourcesSummaryStats) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *ProtectionSourcesSummaryStats) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *ProtectionSourcesSummaryStats) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


