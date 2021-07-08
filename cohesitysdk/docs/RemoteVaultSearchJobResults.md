# RemoteVaultSearchJobResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterCount** | Pointer to **NullableInt32** | Specifies number of Clusters that have archived to the remote Vault that match the criteria specified in the search Job, up to this point in the search. If the search is complete, the total number of Clusters that have archived to the remote Vault and that match the search criteria for the search Job, are reported. If the search was not complete, a partial number is reported. | [optional] 
**ClusterMatchString** | Pointer to **NullableString** | Specifies the value of the clusterMatchSting if it was set in the original search Job. | [optional] 
**Cookie** | Pointer to **NullableString** | Specifies an opaque string to pass to the next request to get the next set of search results. This is provided to support pagination. If null, this is the last set of search results. | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the value of endTimeUsecs if it was set in the original search Job. End time is recorded as a Unix epoch Timestamp (in microseconds). | [optional] 
**Error** | Pointer to **NullableString** | Specifies the error message if the search fails. | [optional] 
**JobCount** | Pointer to **NullableInt32** | Specifies number of Protection Jobs that have archived to the remote Vault that match the criteria specified in the search Job. If the search is complete, the total number of Protection Jobs that have archived to the remote Vault and match the search criteria for the search Job, are reported. If the search is not complete, a partial number is reported. | [optional] 
**JobMatchString** | Pointer to **NullableString** | Specifies the value of the jobMatchSting if it was set in the original search Job. | [optional] 
**ProtectionJobs** | Pointer to [**[]RemoteProtectionJobRunInformation**](RemoteProtectionJobRunInformation.md) | Array of Protection Jobs.  Specifies a list of Protection Jobs that have archived data to a remote Vault and that also match the filter criteria. | [optional] 
**SearchJobStatus** | Pointer to **NullableString** | Specifies the status of the search Job. &#39;kJobRunning&#39; indicates that the Job/task is currently running. &#39;kJobFinished&#39; indicates that the Job/task completed and finished. &#39;kJobFailed&#39; indicates that the Job/task failed and did not complete. &#39;kJobCanceled&#39; indicates that the Job/task was canceled. &#39;kJobPaused&#39; indicates the Job/task is paused. | [optional] 
**SearchJobUid** | Pointer to [**NullableUniversalId**](UniversalId.md) | Specifies the unique id of the search Job assigned by the Cluster. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the value of startTimeUsecs if it was set in the original search Job. Start time is recorded as a Unix epoch Timestamp (in microseconds). | [optional] 
**VaultId** | Pointer to **NullableInt64** | Specifies the id of the remote Vault that was searched. | [optional] 
**VaultName** | Pointer to **NullableString** | Specifies the name of the remote Vault that was searched. | [optional] 

## Methods

### NewRemoteVaultSearchJobResults

`func NewRemoteVaultSearchJobResults() *RemoteVaultSearchJobResults`

NewRemoteVaultSearchJobResults instantiates a new RemoteVaultSearchJobResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteVaultSearchJobResultsWithDefaults

`func NewRemoteVaultSearchJobResultsWithDefaults() *RemoteVaultSearchJobResults`

NewRemoteVaultSearchJobResultsWithDefaults instantiates a new RemoteVaultSearchJobResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterCount

`func (o *RemoteVaultSearchJobResults) GetClusterCount() int32`

GetClusterCount returns the ClusterCount field if non-nil, zero value otherwise.

### GetClusterCountOk

`func (o *RemoteVaultSearchJobResults) GetClusterCountOk() (*int32, bool)`

GetClusterCountOk returns a tuple with the ClusterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCount

`func (o *RemoteVaultSearchJobResults) SetClusterCount(v int32)`

SetClusterCount sets ClusterCount field to given value.

### HasClusterCount

`func (o *RemoteVaultSearchJobResults) HasClusterCount() bool`

HasClusterCount returns a boolean if a field has been set.

### SetClusterCountNil

`func (o *RemoteVaultSearchJobResults) SetClusterCountNil(b bool)`

 SetClusterCountNil sets the value for ClusterCount to be an explicit nil

### UnsetClusterCount
`func (o *RemoteVaultSearchJobResults) UnsetClusterCount()`

UnsetClusterCount ensures that no value is present for ClusterCount, not even an explicit nil
### GetClusterMatchString

`func (o *RemoteVaultSearchJobResults) GetClusterMatchString() string`

GetClusterMatchString returns the ClusterMatchString field if non-nil, zero value otherwise.

### GetClusterMatchStringOk

`func (o *RemoteVaultSearchJobResults) GetClusterMatchStringOk() (*string, bool)`

GetClusterMatchStringOk returns a tuple with the ClusterMatchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMatchString

`func (o *RemoteVaultSearchJobResults) SetClusterMatchString(v string)`

SetClusterMatchString sets ClusterMatchString field to given value.

### HasClusterMatchString

`func (o *RemoteVaultSearchJobResults) HasClusterMatchString() bool`

HasClusterMatchString returns a boolean if a field has been set.

### SetClusterMatchStringNil

`func (o *RemoteVaultSearchJobResults) SetClusterMatchStringNil(b bool)`

 SetClusterMatchStringNil sets the value for ClusterMatchString to be an explicit nil

### UnsetClusterMatchString
`func (o *RemoteVaultSearchJobResults) UnsetClusterMatchString()`

UnsetClusterMatchString ensures that no value is present for ClusterMatchString, not even an explicit nil
### GetCookie

`func (o *RemoteVaultSearchJobResults) GetCookie() string`

GetCookie returns the Cookie field if non-nil, zero value otherwise.

### GetCookieOk

`func (o *RemoteVaultSearchJobResults) GetCookieOk() (*string, bool)`

GetCookieOk returns a tuple with the Cookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookie

`func (o *RemoteVaultSearchJobResults) SetCookie(v string)`

SetCookie sets Cookie field to given value.

### HasCookie

`func (o *RemoteVaultSearchJobResults) HasCookie() bool`

HasCookie returns a boolean if a field has been set.

### SetCookieNil

`func (o *RemoteVaultSearchJobResults) SetCookieNil(b bool)`

 SetCookieNil sets the value for Cookie to be an explicit nil

### UnsetCookie
`func (o *RemoteVaultSearchJobResults) UnsetCookie()`

UnsetCookie ensures that no value is present for Cookie, not even an explicit nil
### GetEndTimeUsecs

`func (o *RemoteVaultSearchJobResults) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *RemoteVaultSearchJobResults) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *RemoteVaultSearchJobResults) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *RemoteVaultSearchJobResults) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *RemoteVaultSearchJobResults) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *RemoteVaultSearchJobResults) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetError

`func (o *RemoteVaultSearchJobResults) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RemoteVaultSearchJobResults) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RemoteVaultSearchJobResults) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *RemoteVaultSearchJobResults) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *RemoteVaultSearchJobResults) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *RemoteVaultSearchJobResults) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetJobCount

`func (o *RemoteVaultSearchJobResults) GetJobCount() int32`

GetJobCount returns the JobCount field if non-nil, zero value otherwise.

### GetJobCountOk

`func (o *RemoteVaultSearchJobResults) GetJobCountOk() (*int32, bool)`

GetJobCountOk returns a tuple with the JobCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobCount

`func (o *RemoteVaultSearchJobResults) SetJobCount(v int32)`

SetJobCount sets JobCount field to given value.

### HasJobCount

`func (o *RemoteVaultSearchJobResults) HasJobCount() bool`

HasJobCount returns a boolean if a field has been set.

### SetJobCountNil

`func (o *RemoteVaultSearchJobResults) SetJobCountNil(b bool)`

 SetJobCountNil sets the value for JobCount to be an explicit nil

### UnsetJobCount
`func (o *RemoteVaultSearchJobResults) UnsetJobCount()`

UnsetJobCount ensures that no value is present for JobCount, not even an explicit nil
### GetJobMatchString

`func (o *RemoteVaultSearchJobResults) GetJobMatchString() string`

GetJobMatchString returns the JobMatchString field if non-nil, zero value otherwise.

### GetJobMatchStringOk

`func (o *RemoteVaultSearchJobResults) GetJobMatchStringOk() (*string, bool)`

GetJobMatchStringOk returns a tuple with the JobMatchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobMatchString

`func (o *RemoteVaultSearchJobResults) SetJobMatchString(v string)`

SetJobMatchString sets JobMatchString field to given value.

### HasJobMatchString

`func (o *RemoteVaultSearchJobResults) HasJobMatchString() bool`

HasJobMatchString returns a boolean if a field has been set.

### SetJobMatchStringNil

`func (o *RemoteVaultSearchJobResults) SetJobMatchStringNil(b bool)`

 SetJobMatchStringNil sets the value for JobMatchString to be an explicit nil

### UnsetJobMatchString
`func (o *RemoteVaultSearchJobResults) UnsetJobMatchString()`

UnsetJobMatchString ensures that no value is present for JobMatchString, not even an explicit nil
### GetProtectionJobs

`func (o *RemoteVaultSearchJobResults) GetProtectionJobs() []RemoteProtectionJobRunInformation`

GetProtectionJobs returns the ProtectionJobs field if non-nil, zero value otherwise.

### GetProtectionJobsOk

`func (o *RemoteVaultSearchJobResults) GetProtectionJobsOk() (*[]RemoteProtectionJobRunInformation, bool)`

GetProtectionJobsOk returns a tuple with the ProtectionJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionJobs

`func (o *RemoteVaultSearchJobResults) SetProtectionJobs(v []RemoteProtectionJobRunInformation)`

SetProtectionJobs sets ProtectionJobs field to given value.

### HasProtectionJobs

`func (o *RemoteVaultSearchJobResults) HasProtectionJobs() bool`

HasProtectionJobs returns a boolean if a field has been set.

### SetProtectionJobsNil

`func (o *RemoteVaultSearchJobResults) SetProtectionJobsNil(b bool)`

 SetProtectionJobsNil sets the value for ProtectionJobs to be an explicit nil

### UnsetProtectionJobs
`func (o *RemoteVaultSearchJobResults) UnsetProtectionJobs()`

UnsetProtectionJobs ensures that no value is present for ProtectionJobs, not even an explicit nil
### GetSearchJobStatus

`func (o *RemoteVaultSearchJobResults) GetSearchJobStatus() string`

GetSearchJobStatus returns the SearchJobStatus field if non-nil, zero value otherwise.

### GetSearchJobStatusOk

`func (o *RemoteVaultSearchJobResults) GetSearchJobStatusOk() (*string, bool)`

GetSearchJobStatusOk returns a tuple with the SearchJobStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchJobStatus

`func (o *RemoteVaultSearchJobResults) SetSearchJobStatus(v string)`

SetSearchJobStatus sets SearchJobStatus field to given value.

### HasSearchJobStatus

`func (o *RemoteVaultSearchJobResults) HasSearchJobStatus() bool`

HasSearchJobStatus returns a boolean if a field has been set.

### SetSearchJobStatusNil

`func (o *RemoteVaultSearchJobResults) SetSearchJobStatusNil(b bool)`

 SetSearchJobStatusNil sets the value for SearchJobStatus to be an explicit nil

### UnsetSearchJobStatus
`func (o *RemoteVaultSearchJobResults) UnsetSearchJobStatus()`

UnsetSearchJobStatus ensures that no value is present for SearchJobStatus, not even an explicit nil
### GetSearchJobUid

`func (o *RemoteVaultSearchJobResults) GetSearchJobUid() UniversalId`

GetSearchJobUid returns the SearchJobUid field if non-nil, zero value otherwise.

### GetSearchJobUidOk

`func (o *RemoteVaultSearchJobResults) GetSearchJobUidOk() (*UniversalId, bool)`

GetSearchJobUidOk returns a tuple with the SearchJobUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchJobUid

`func (o *RemoteVaultSearchJobResults) SetSearchJobUid(v UniversalId)`

SetSearchJobUid sets SearchJobUid field to given value.

### HasSearchJobUid

`func (o *RemoteVaultSearchJobResults) HasSearchJobUid() bool`

HasSearchJobUid returns a boolean if a field has been set.

### SetSearchJobUidNil

`func (o *RemoteVaultSearchJobResults) SetSearchJobUidNil(b bool)`

 SetSearchJobUidNil sets the value for SearchJobUid to be an explicit nil

### UnsetSearchJobUid
`func (o *RemoteVaultSearchJobResults) UnsetSearchJobUid()`

UnsetSearchJobUid ensures that no value is present for SearchJobUid, not even an explicit nil
### GetStartTimeUsecs

`func (o *RemoteVaultSearchJobResults) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *RemoteVaultSearchJobResults) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *RemoteVaultSearchJobResults) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *RemoteVaultSearchJobResults) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *RemoteVaultSearchJobResults) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *RemoteVaultSearchJobResults) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetVaultId

`func (o *RemoteVaultSearchJobResults) GetVaultId() int64`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *RemoteVaultSearchJobResults) GetVaultIdOk() (*int64, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *RemoteVaultSearchJobResults) SetVaultId(v int64)`

SetVaultId sets VaultId field to given value.

### HasVaultId

`func (o *RemoteVaultSearchJobResults) HasVaultId() bool`

HasVaultId returns a boolean if a field has been set.

### SetVaultIdNil

`func (o *RemoteVaultSearchJobResults) SetVaultIdNil(b bool)`

 SetVaultIdNil sets the value for VaultId to be an explicit nil

### UnsetVaultId
`func (o *RemoteVaultSearchJobResults) UnsetVaultId()`

UnsetVaultId ensures that no value is present for VaultId, not even an explicit nil
### GetVaultName

`func (o *RemoteVaultSearchJobResults) GetVaultName() string`

GetVaultName returns the VaultName field if non-nil, zero value otherwise.

### GetVaultNameOk

`func (o *RemoteVaultSearchJobResults) GetVaultNameOk() (*string, bool)`

GetVaultNameOk returns a tuple with the VaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultName

`func (o *RemoteVaultSearchJobResults) SetVaultName(v string)`

SetVaultName sets VaultName field to given value.

### HasVaultName

`func (o *RemoteVaultSearchJobResults) HasVaultName() bool`

HasVaultName returns a boolean if a field has been set.

### SetVaultNameNil

`func (o *RemoteVaultSearchJobResults) SetVaultNameNil(b bool)`

 SetVaultNameNil sets the value for VaultName to be an explicit nil

### UnsetVaultName
`func (o *RemoteVaultSearchJobResults) UnsetVaultName()`

UnsetVaultName ensures that no value is present for VaultName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


