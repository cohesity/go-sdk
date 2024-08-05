# CommonSqlAppSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeepCdc** | Pointer to **NullableBool** | Specifies whether to keep CDC (Change Data Capture) on recovered databases or not. If not passed, this is assumed to be true. If withNoRecovery is passed as true, then this field must not be set to true. Passing this field as true in this scenario will be a invalid request. | [optional] 
**MultiStageRestoreOptions** | Pointer to [**MultiStageRestoreOptions**](MultiStageRestoreOptions.md) |  | [optional] 
**NativeLogRecoveryWithClause** | Pointer to **NullableString** | Specifies the WITH clause to be used in native sql log restore command. This is only applicable for native log restore. | [optional] 
**NativeRecoveryWithClause** | Pointer to **NullableString** | &#39;with_clause&#39; contains &#39;with clause&#39; to be used in native sql restore command. This is only applicable for database restore of native sql backup. Here user can specify multiple restore options. Example: &#39;WITH BUFFERCOUNT &#x3D; 575, MAXTRANSFERSIZE &#x3D; 2097152&#39;. | [optional] 
**OverwritingPolicy** | Pointer to **NullableString** | Specifies a policy to be used while recovering existing databases. | [optional] 
**ReplayEntireLastLog** | Pointer to **NullableBool** | Specifies the option to set replay last log bit while creating the sql restore task and doing restore to latest point-in-time. If this is set to true, we will replay the entire last log without STOPAT. | [optional] 
**RestoreTimeUsecs** | Pointer to **NullableInt64** | Specifies the time in the past to which the Sql database needs to be restored. This allows for granular recovery of Sql databases. If this is not set, the Sql database will be restored from the full/incremental snapshot. | [optional] 
**SecondaryDataFilesDirList** | Pointer to [**[]FilenamePatternToDirectory**](FilenamePatternToDirectory.md) | Specifies the secondary data filename pattern and corresponding direcories of the DB. Secondary data files are optional and are user defined. The recommended file extention for secondary files is \&quot;.ndf\&quot;. If this option is specified and the destination folders do not exist they will be automatically created. | [optional] 
**WithNoRecovery** | Pointer to **NullableBool** | Specifies the flag to bring DBs online or not after successful recovery. If this is passed as true, then it means DBs won&#39;t be brought online. | [optional] 

## Methods

### NewCommonSqlAppSourceConfig

`func NewCommonSqlAppSourceConfig() *CommonSqlAppSourceConfig`

NewCommonSqlAppSourceConfig instantiates a new CommonSqlAppSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonSqlAppSourceConfigWithDefaults

`func NewCommonSqlAppSourceConfigWithDefaults() *CommonSqlAppSourceConfig`

NewCommonSqlAppSourceConfigWithDefaults instantiates a new CommonSqlAppSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeepCdc

`func (o *CommonSqlAppSourceConfig) GetKeepCdc() bool`

GetKeepCdc returns the KeepCdc field if non-nil, zero value otherwise.

### GetKeepCdcOk

`func (o *CommonSqlAppSourceConfig) GetKeepCdcOk() (*bool, bool)`

GetKeepCdcOk returns a tuple with the KeepCdc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepCdc

`func (o *CommonSqlAppSourceConfig) SetKeepCdc(v bool)`

SetKeepCdc sets KeepCdc field to given value.

### HasKeepCdc

`func (o *CommonSqlAppSourceConfig) HasKeepCdc() bool`

HasKeepCdc returns a boolean if a field has been set.

### SetKeepCdcNil

`func (o *CommonSqlAppSourceConfig) SetKeepCdcNil(b bool)`

 SetKeepCdcNil sets the value for KeepCdc to be an explicit nil

### UnsetKeepCdc
`func (o *CommonSqlAppSourceConfig) UnsetKeepCdc()`

UnsetKeepCdc ensures that no value is present for KeepCdc, not even an explicit nil
### GetMultiStageRestoreOptions

`func (o *CommonSqlAppSourceConfig) GetMultiStageRestoreOptions() MultiStageRestoreOptions`

GetMultiStageRestoreOptions returns the MultiStageRestoreOptions field if non-nil, zero value otherwise.

### GetMultiStageRestoreOptionsOk

`func (o *CommonSqlAppSourceConfig) GetMultiStageRestoreOptionsOk() (*MultiStageRestoreOptions, bool)`

GetMultiStageRestoreOptionsOk returns a tuple with the MultiStageRestoreOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiStageRestoreOptions

`func (o *CommonSqlAppSourceConfig) SetMultiStageRestoreOptions(v MultiStageRestoreOptions)`

SetMultiStageRestoreOptions sets MultiStageRestoreOptions field to given value.

### HasMultiStageRestoreOptions

`func (o *CommonSqlAppSourceConfig) HasMultiStageRestoreOptions() bool`

HasMultiStageRestoreOptions returns a boolean if a field has been set.

### GetNativeLogRecoveryWithClause

`func (o *CommonSqlAppSourceConfig) GetNativeLogRecoveryWithClause() string`

GetNativeLogRecoveryWithClause returns the NativeLogRecoveryWithClause field if non-nil, zero value otherwise.

### GetNativeLogRecoveryWithClauseOk

`func (o *CommonSqlAppSourceConfig) GetNativeLogRecoveryWithClauseOk() (*string, bool)`

GetNativeLogRecoveryWithClauseOk returns a tuple with the NativeLogRecoveryWithClause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeLogRecoveryWithClause

`func (o *CommonSqlAppSourceConfig) SetNativeLogRecoveryWithClause(v string)`

SetNativeLogRecoveryWithClause sets NativeLogRecoveryWithClause field to given value.

### HasNativeLogRecoveryWithClause

`func (o *CommonSqlAppSourceConfig) HasNativeLogRecoveryWithClause() bool`

HasNativeLogRecoveryWithClause returns a boolean if a field has been set.

### SetNativeLogRecoveryWithClauseNil

`func (o *CommonSqlAppSourceConfig) SetNativeLogRecoveryWithClauseNil(b bool)`

 SetNativeLogRecoveryWithClauseNil sets the value for NativeLogRecoveryWithClause to be an explicit nil

### UnsetNativeLogRecoveryWithClause
`func (o *CommonSqlAppSourceConfig) UnsetNativeLogRecoveryWithClause()`

UnsetNativeLogRecoveryWithClause ensures that no value is present for NativeLogRecoveryWithClause, not even an explicit nil
### GetNativeRecoveryWithClause

`func (o *CommonSqlAppSourceConfig) GetNativeRecoveryWithClause() string`

GetNativeRecoveryWithClause returns the NativeRecoveryWithClause field if non-nil, zero value otherwise.

### GetNativeRecoveryWithClauseOk

`func (o *CommonSqlAppSourceConfig) GetNativeRecoveryWithClauseOk() (*string, bool)`

GetNativeRecoveryWithClauseOk returns a tuple with the NativeRecoveryWithClause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeRecoveryWithClause

`func (o *CommonSqlAppSourceConfig) SetNativeRecoveryWithClause(v string)`

SetNativeRecoveryWithClause sets NativeRecoveryWithClause field to given value.

### HasNativeRecoveryWithClause

`func (o *CommonSqlAppSourceConfig) HasNativeRecoveryWithClause() bool`

HasNativeRecoveryWithClause returns a boolean if a field has been set.

### SetNativeRecoveryWithClauseNil

`func (o *CommonSqlAppSourceConfig) SetNativeRecoveryWithClauseNil(b bool)`

 SetNativeRecoveryWithClauseNil sets the value for NativeRecoveryWithClause to be an explicit nil

### UnsetNativeRecoveryWithClause
`func (o *CommonSqlAppSourceConfig) UnsetNativeRecoveryWithClause()`

UnsetNativeRecoveryWithClause ensures that no value is present for NativeRecoveryWithClause, not even an explicit nil
### GetOverwritingPolicy

`func (o *CommonSqlAppSourceConfig) GetOverwritingPolicy() string`

GetOverwritingPolicy returns the OverwritingPolicy field if non-nil, zero value otherwise.

### GetOverwritingPolicyOk

`func (o *CommonSqlAppSourceConfig) GetOverwritingPolicyOk() (*string, bool)`

GetOverwritingPolicyOk returns a tuple with the OverwritingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwritingPolicy

`func (o *CommonSqlAppSourceConfig) SetOverwritingPolicy(v string)`

SetOverwritingPolicy sets OverwritingPolicy field to given value.

### HasOverwritingPolicy

`func (o *CommonSqlAppSourceConfig) HasOverwritingPolicy() bool`

HasOverwritingPolicy returns a boolean if a field has been set.

### SetOverwritingPolicyNil

`func (o *CommonSqlAppSourceConfig) SetOverwritingPolicyNil(b bool)`

 SetOverwritingPolicyNil sets the value for OverwritingPolicy to be an explicit nil

### UnsetOverwritingPolicy
`func (o *CommonSqlAppSourceConfig) UnsetOverwritingPolicy()`

UnsetOverwritingPolicy ensures that no value is present for OverwritingPolicy, not even an explicit nil
### GetReplayEntireLastLog

`func (o *CommonSqlAppSourceConfig) GetReplayEntireLastLog() bool`

GetReplayEntireLastLog returns the ReplayEntireLastLog field if non-nil, zero value otherwise.

### GetReplayEntireLastLogOk

`func (o *CommonSqlAppSourceConfig) GetReplayEntireLastLogOk() (*bool, bool)`

GetReplayEntireLastLogOk returns a tuple with the ReplayEntireLastLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayEntireLastLog

`func (o *CommonSqlAppSourceConfig) SetReplayEntireLastLog(v bool)`

SetReplayEntireLastLog sets ReplayEntireLastLog field to given value.

### HasReplayEntireLastLog

`func (o *CommonSqlAppSourceConfig) HasReplayEntireLastLog() bool`

HasReplayEntireLastLog returns a boolean if a field has been set.

### SetReplayEntireLastLogNil

`func (o *CommonSqlAppSourceConfig) SetReplayEntireLastLogNil(b bool)`

 SetReplayEntireLastLogNil sets the value for ReplayEntireLastLog to be an explicit nil

### UnsetReplayEntireLastLog
`func (o *CommonSqlAppSourceConfig) UnsetReplayEntireLastLog()`

UnsetReplayEntireLastLog ensures that no value is present for ReplayEntireLastLog, not even an explicit nil
### GetRestoreTimeUsecs

`func (o *CommonSqlAppSourceConfig) GetRestoreTimeUsecs() int64`

GetRestoreTimeUsecs returns the RestoreTimeUsecs field if non-nil, zero value otherwise.

### GetRestoreTimeUsecsOk

`func (o *CommonSqlAppSourceConfig) GetRestoreTimeUsecsOk() (*int64, bool)`

GetRestoreTimeUsecsOk returns a tuple with the RestoreTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTimeUsecs

`func (o *CommonSqlAppSourceConfig) SetRestoreTimeUsecs(v int64)`

SetRestoreTimeUsecs sets RestoreTimeUsecs field to given value.

### HasRestoreTimeUsecs

`func (o *CommonSqlAppSourceConfig) HasRestoreTimeUsecs() bool`

HasRestoreTimeUsecs returns a boolean if a field has been set.

### SetRestoreTimeUsecsNil

`func (o *CommonSqlAppSourceConfig) SetRestoreTimeUsecsNil(b bool)`

 SetRestoreTimeUsecsNil sets the value for RestoreTimeUsecs to be an explicit nil

### UnsetRestoreTimeUsecs
`func (o *CommonSqlAppSourceConfig) UnsetRestoreTimeUsecs()`

UnsetRestoreTimeUsecs ensures that no value is present for RestoreTimeUsecs, not even an explicit nil
### GetSecondaryDataFilesDirList

`func (o *CommonSqlAppSourceConfig) GetSecondaryDataFilesDirList() []FilenamePatternToDirectory`

GetSecondaryDataFilesDirList returns the SecondaryDataFilesDirList field if non-nil, zero value otherwise.

### GetSecondaryDataFilesDirListOk

`func (o *CommonSqlAppSourceConfig) GetSecondaryDataFilesDirListOk() (*[]FilenamePatternToDirectory, bool)`

GetSecondaryDataFilesDirListOk returns a tuple with the SecondaryDataFilesDirList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDataFilesDirList

`func (o *CommonSqlAppSourceConfig) SetSecondaryDataFilesDirList(v []FilenamePatternToDirectory)`

SetSecondaryDataFilesDirList sets SecondaryDataFilesDirList field to given value.

### HasSecondaryDataFilesDirList

`func (o *CommonSqlAppSourceConfig) HasSecondaryDataFilesDirList() bool`

HasSecondaryDataFilesDirList returns a boolean if a field has been set.

### SetSecondaryDataFilesDirListNil

`func (o *CommonSqlAppSourceConfig) SetSecondaryDataFilesDirListNil(b bool)`

 SetSecondaryDataFilesDirListNil sets the value for SecondaryDataFilesDirList to be an explicit nil

### UnsetSecondaryDataFilesDirList
`func (o *CommonSqlAppSourceConfig) UnsetSecondaryDataFilesDirList()`

UnsetSecondaryDataFilesDirList ensures that no value is present for SecondaryDataFilesDirList, not even an explicit nil
### GetWithNoRecovery

`func (o *CommonSqlAppSourceConfig) GetWithNoRecovery() bool`

GetWithNoRecovery returns the WithNoRecovery field if non-nil, zero value otherwise.

### GetWithNoRecoveryOk

`func (o *CommonSqlAppSourceConfig) GetWithNoRecoveryOk() (*bool, bool)`

GetWithNoRecoveryOk returns a tuple with the WithNoRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithNoRecovery

`func (o *CommonSqlAppSourceConfig) SetWithNoRecovery(v bool)`

SetWithNoRecovery sets WithNoRecovery field to given value.

### HasWithNoRecovery

`func (o *CommonSqlAppSourceConfig) HasWithNoRecovery() bool`

HasWithNoRecovery returns a boolean if a field has been set.

### SetWithNoRecoveryNil

`func (o *CommonSqlAppSourceConfig) SetWithNoRecoveryNil(b bool)`

 SetWithNoRecoveryNil sets the value for WithNoRecovery to be an explicit nil

### UnsetWithNoRecovery
`func (o *CommonSqlAppSourceConfig) UnsetWithNoRecovery()`

UnsetWithNoRecovery ensures that no value is present for WithNoRecovery, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


