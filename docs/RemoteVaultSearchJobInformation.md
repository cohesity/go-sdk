# RemoteVaultSearchJobInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterCount** | Pointer to **NullableInt32** | Specifies number of Clusters that have archived to the remote Vault and match the search criteria for this job, up to this point in the search. If the search is complete, the total number of Clusters that have archived to the remote Vault and that match the search criteria for this search Job, are reported. If the search is not complete, a partial number is reported. | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of the search as a Unix epoch Timestamp (in microseconds) if the search Job has completed. | [optional] 
**Error** | Pointer to **NullableString** | Specifies the error message reported when a search fails. | [optional] 
**JobCount** | Pointer to **NullableInt32** | Specifies number of Protection Jobs that have archived to the remote Vault and match the search criteria for this search Job, up to this point in the search. If the search is complete, the total number of Protection Jobs that have archived to the remote Vault and that match the search criteria for this search Job, are reported. If the search is not complete, a partial number is reported. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the search Job. | [optional] 
**SearchJobStatus** | Pointer to **NullableString** | Specifies the status of the search. &#39;kJobRunning&#39; indicates that the Job/task is currently running. &#39;kJobFinished&#39; indicates that the Job/task completed and finished. &#39;kJobFailed&#39; indicates that the Job/task failed and did not complete. &#39;kJobCanceled&#39; indicates that the Job/task was canceled. &#39;kJobPaused&#39; indicates the Job/task is paused. | [optional] 
**SearchJobUid** | Pointer to [**NullableUniversalId**](UniversalId.md) | Specifies the unique id assigned to the search Job by the Cluster. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of the search as a Unix epoch Timestamp (in microseconds). | [optional] 
**VaultId** | Pointer to **NullableInt64** | Specifies the id of the remote Vault (External Target) that was searched. | [optional] 
**VaultName** | Pointer to **NullableString** | Specifies the name of the remote Vault (External Target) that was searched. | [optional] 

## Methods

### NewRemoteVaultSearchJobInformation

`func NewRemoteVaultSearchJobInformation() *RemoteVaultSearchJobInformation`

NewRemoteVaultSearchJobInformation instantiates a new RemoteVaultSearchJobInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteVaultSearchJobInformationWithDefaults

`func NewRemoteVaultSearchJobInformationWithDefaults() *RemoteVaultSearchJobInformation`

NewRemoteVaultSearchJobInformationWithDefaults instantiates a new RemoteVaultSearchJobInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterCount

`func (o *RemoteVaultSearchJobInformation) GetClusterCount() int32`

GetClusterCount returns the ClusterCount field if non-nil, zero value otherwise.

### GetClusterCountOk

`func (o *RemoteVaultSearchJobInformation) GetClusterCountOk() (*int32, bool)`

GetClusterCountOk returns a tuple with the ClusterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCount

`func (o *RemoteVaultSearchJobInformation) SetClusterCount(v int32)`

SetClusterCount sets ClusterCount field to given value.

### HasClusterCount

`func (o *RemoteVaultSearchJobInformation) HasClusterCount() bool`

HasClusterCount returns a boolean if a field has been set.

### SetClusterCountNil

`func (o *RemoteVaultSearchJobInformation) SetClusterCountNil(b bool)`

 SetClusterCountNil sets the value for ClusterCount to be an explicit nil

### UnsetClusterCount
`func (o *RemoteVaultSearchJobInformation) UnsetClusterCount()`

UnsetClusterCount ensures that no value is present for ClusterCount, not even an explicit nil
### GetEndTimeUsecs

`func (o *RemoteVaultSearchJobInformation) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *RemoteVaultSearchJobInformation) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *RemoteVaultSearchJobInformation) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *RemoteVaultSearchJobInformation) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *RemoteVaultSearchJobInformation) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *RemoteVaultSearchJobInformation) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetError

`func (o *RemoteVaultSearchJobInformation) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RemoteVaultSearchJobInformation) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RemoteVaultSearchJobInformation) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *RemoteVaultSearchJobInformation) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *RemoteVaultSearchJobInformation) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *RemoteVaultSearchJobInformation) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetJobCount

`func (o *RemoteVaultSearchJobInformation) GetJobCount() int32`

GetJobCount returns the JobCount field if non-nil, zero value otherwise.

### GetJobCountOk

`func (o *RemoteVaultSearchJobInformation) GetJobCountOk() (*int32, bool)`

GetJobCountOk returns a tuple with the JobCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobCount

`func (o *RemoteVaultSearchJobInformation) SetJobCount(v int32)`

SetJobCount sets JobCount field to given value.

### HasJobCount

`func (o *RemoteVaultSearchJobInformation) HasJobCount() bool`

HasJobCount returns a boolean if a field has been set.

### SetJobCountNil

`func (o *RemoteVaultSearchJobInformation) SetJobCountNil(b bool)`

 SetJobCountNil sets the value for JobCount to be an explicit nil

### UnsetJobCount
`func (o *RemoteVaultSearchJobInformation) UnsetJobCount()`

UnsetJobCount ensures that no value is present for JobCount, not even an explicit nil
### GetName

`func (o *RemoteVaultSearchJobInformation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RemoteVaultSearchJobInformation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RemoteVaultSearchJobInformation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RemoteVaultSearchJobInformation) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RemoteVaultSearchJobInformation) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RemoteVaultSearchJobInformation) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSearchJobStatus

`func (o *RemoteVaultSearchJobInformation) GetSearchJobStatus() string`

GetSearchJobStatus returns the SearchJobStatus field if non-nil, zero value otherwise.

### GetSearchJobStatusOk

`func (o *RemoteVaultSearchJobInformation) GetSearchJobStatusOk() (*string, bool)`

GetSearchJobStatusOk returns a tuple with the SearchJobStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchJobStatus

`func (o *RemoteVaultSearchJobInformation) SetSearchJobStatus(v string)`

SetSearchJobStatus sets SearchJobStatus field to given value.

### HasSearchJobStatus

`func (o *RemoteVaultSearchJobInformation) HasSearchJobStatus() bool`

HasSearchJobStatus returns a boolean if a field has been set.

### SetSearchJobStatusNil

`func (o *RemoteVaultSearchJobInformation) SetSearchJobStatusNil(b bool)`

 SetSearchJobStatusNil sets the value for SearchJobStatus to be an explicit nil

### UnsetSearchJobStatus
`func (o *RemoteVaultSearchJobInformation) UnsetSearchJobStatus()`

UnsetSearchJobStatus ensures that no value is present for SearchJobStatus, not even an explicit nil
### GetSearchJobUid

`func (o *RemoteVaultSearchJobInformation) GetSearchJobUid() UniversalId`

GetSearchJobUid returns the SearchJobUid field if non-nil, zero value otherwise.

### GetSearchJobUidOk

`func (o *RemoteVaultSearchJobInformation) GetSearchJobUidOk() (*UniversalId, bool)`

GetSearchJobUidOk returns a tuple with the SearchJobUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchJobUid

`func (o *RemoteVaultSearchJobInformation) SetSearchJobUid(v UniversalId)`

SetSearchJobUid sets SearchJobUid field to given value.

### HasSearchJobUid

`func (o *RemoteVaultSearchJobInformation) HasSearchJobUid() bool`

HasSearchJobUid returns a boolean if a field has been set.

### SetSearchJobUidNil

`func (o *RemoteVaultSearchJobInformation) SetSearchJobUidNil(b bool)`

 SetSearchJobUidNil sets the value for SearchJobUid to be an explicit nil

### UnsetSearchJobUid
`func (o *RemoteVaultSearchJobInformation) UnsetSearchJobUid()`

UnsetSearchJobUid ensures that no value is present for SearchJobUid, not even an explicit nil
### GetStartTimeUsecs

`func (o *RemoteVaultSearchJobInformation) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *RemoteVaultSearchJobInformation) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *RemoteVaultSearchJobInformation) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *RemoteVaultSearchJobInformation) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *RemoteVaultSearchJobInformation) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *RemoteVaultSearchJobInformation) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetVaultId

`func (o *RemoteVaultSearchJobInformation) GetVaultId() int64`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *RemoteVaultSearchJobInformation) GetVaultIdOk() (*int64, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *RemoteVaultSearchJobInformation) SetVaultId(v int64)`

SetVaultId sets VaultId field to given value.

### HasVaultId

`func (o *RemoteVaultSearchJobInformation) HasVaultId() bool`

HasVaultId returns a boolean if a field has been set.

### SetVaultIdNil

`func (o *RemoteVaultSearchJobInformation) SetVaultIdNil(b bool)`

 SetVaultIdNil sets the value for VaultId to be an explicit nil

### UnsetVaultId
`func (o *RemoteVaultSearchJobInformation) UnsetVaultId()`

UnsetVaultId ensures that no value is present for VaultId, not even an explicit nil
### GetVaultName

`func (o *RemoteVaultSearchJobInformation) GetVaultName() string`

GetVaultName returns the VaultName field if non-nil, zero value otherwise.

### GetVaultNameOk

`func (o *RemoteVaultSearchJobInformation) GetVaultNameOk() (*string, bool)`

GetVaultNameOk returns a tuple with the VaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultName

`func (o *RemoteVaultSearchJobInformation) SetVaultName(v string)`

SetVaultName sets VaultName field to given value.

### HasVaultName

`func (o *RemoteVaultSearchJobInformation) HasVaultName() bool`

HasVaultName returns a boolean if a field has been set.

### SetVaultNameNil

`func (o *RemoteVaultSearchJobInformation) SetVaultNameNil(b bool)`

 SetVaultNameNil sets the value for VaultName to be an explicit nil

### UnsetVaultName
`func (o *RemoteVaultSearchJobInformation) UnsetVaultName()`

UnsetVaultName ensures that no value is present for VaultName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


