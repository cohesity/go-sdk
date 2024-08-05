# AwsS3ProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupObjectLevelACLs** | Pointer to **NullableBool** | Specifies whether to backup object level acls. Default value is false. | [optional] 
**BaselineIncrementalFrequency** | Pointer to **NullableString** | Specifies the baseline incremental frequency. | [optional] 
**InventoryReportDestination** | Pointer to **NullableString** | ARN of the inventory report destination bucket for S3 backups. | [optional] 
**InventoryReportDestinationPrefix** | Pointer to **NullableString** | The prefix in the S3 destination bucket where inventory reports will be stored. | [optional] 
**InventoryReportFrequency** | Pointer to **NullableString** | Specifies the frequency to generate inventory reports. | [optional] 
**Objects** | Pointer to [**[]AwsS3ProtectionGroupObjectParams**](AwsS3ProtectionGroupObjectParams.md) | Specifies the objects to be protected. | [optional] 
**SkipOnError** | Pointer to **NullableBool** | Specifies whether to skip files on error or not. Default value is false. | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies the id of the parent of the objects. | [optional] [readonly] 
**SourceName** | Pointer to **NullableString** | Specifies the name of the parent of the objects. | [optional] [readonly] 
**StorageClass** | Pointer to **[]string** | Specifies the AWS S3 Storage classes to backup. | [optional] 

## Methods

### NewAwsS3ProtectionGroupParams

`func NewAwsS3ProtectionGroupParams() *AwsS3ProtectionGroupParams`

NewAwsS3ProtectionGroupParams instantiates a new AwsS3ProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsS3ProtectionGroupParamsWithDefaults

`func NewAwsS3ProtectionGroupParamsWithDefaults() *AwsS3ProtectionGroupParams`

NewAwsS3ProtectionGroupParamsWithDefaults instantiates a new AwsS3ProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupObjectLevelACLs

`func (o *AwsS3ProtectionGroupParams) GetBackupObjectLevelACLs() bool`

GetBackupObjectLevelACLs returns the BackupObjectLevelACLs field if non-nil, zero value otherwise.

### GetBackupObjectLevelACLsOk

`func (o *AwsS3ProtectionGroupParams) GetBackupObjectLevelACLsOk() (*bool, bool)`

GetBackupObjectLevelACLsOk returns a tuple with the BackupObjectLevelACLs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupObjectLevelACLs

`func (o *AwsS3ProtectionGroupParams) SetBackupObjectLevelACLs(v bool)`

SetBackupObjectLevelACLs sets BackupObjectLevelACLs field to given value.

### HasBackupObjectLevelACLs

`func (o *AwsS3ProtectionGroupParams) HasBackupObjectLevelACLs() bool`

HasBackupObjectLevelACLs returns a boolean if a field has been set.

### SetBackupObjectLevelACLsNil

`func (o *AwsS3ProtectionGroupParams) SetBackupObjectLevelACLsNil(b bool)`

 SetBackupObjectLevelACLsNil sets the value for BackupObjectLevelACLs to be an explicit nil

### UnsetBackupObjectLevelACLs
`func (o *AwsS3ProtectionGroupParams) UnsetBackupObjectLevelACLs()`

UnsetBackupObjectLevelACLs ensures that no value is present for BackupObjectLevelACLs, not even an explicit nil
### GetBaselineIncrementalFrequency

`func (o *AwsS3ProtectionGroupParams) GetBaselineIncrementalFrequency() string`

GetBaselineIncrementalFrequency returns the BaselineIncrementalFrequency field if non-nil, zero value otherwise.

### GetBaselineIncrementalFrequencyOk

`func (o *AwsS3ProtectionGroupParams) GetBaselineIncrementalFrequencyOk() (*string, bool)`

GetBaselineIncrementalFrequencyOk returns a tuple with the BaselineIncrementalFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineIncrementalFrequency

`func (o *AwsS3ProtectionGroupParams) SetBaselineIncrementalFrequency(v string)`

SetBaselineIncrementalFrequency sets BaselineIncrementalFrequency field to given value.

### HasBaselineIncrementalFrequency

`func (o *AwsS3ProtectionGroupParams) HasBaselineIncrementalFrequency() bool`

HasBaselineIncrementalFrequency returns a boolean if a field has been set.

### SetBaselineIncrementalFrequencyNil

`func (o *AwsS3ProtectionGroupParams) SetBaselineIncrementalFrequencyNil(b bool)`

 SetBaselineIncrementalFrequencyNil sets the value for BaselineIncrementalFrequency to be an explicit nil

### UnsetBaselineIncrementalFrequency
`func (o *AwsS3ProtectionGroupParams) UnsetBaselineIncrementalFrequency()`

UnsetBaselineIncrementalFrequency ensures that no value is present for BaselineIncrementalFrequency, not even an explicit nil
### GetInventoryReportDestination

`func (o *AwsS3ProtectionGroupParams) GetInventoryReportDestination() string`

GetInventoryReportDestination returns the InventoryReportDestination field if non-nil, zero value otherwise.

### GetInventoryReportDestinationOk

`func (o *AwsS3ProtectionGroupParams) GetInventoryReportDestinationOk() (*string, bool)`

GetInventoryReportDestinationOk returns a tuple with the InventoryReportDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryReportDestination

`func (o *AwsS3ProtectionGroupParams) SetInventoryReportDestination(v string)`

SetInventoryReportDestination sets InventoryReportDestination field to given value.

### HasInventoryReportDestination

`func (o *AwsS3ProtectionGroupParams) HasInventoryReportDestination() bool`

HasInventoryReportDestination returns a boolean if a field has been set.

### SetInventoryReportDestinationNil

`func (o *AwsS3ProtectionGroupParams) SetInventoryReportDestinationNil(b bool)`

 SetInventoryReportDestinationNil sets the value for InventoryReportDestination to be an explicit nil

### UnsetInventoryReportDestination
`func (o *AwsS3ProtectionGroupParams) UnsetInventoryReportDestination()`

UnsetInventoryReportDestination ensures that no value is present for InventoryReportDestination, not even an explicit nil
### GetInventoryReportDestinationPrefix

`func (o *AwsS3ProtectionGroupParams) GetInventoryReportDestinationPrefix() string`

GetInventoryReportDestinationPrefix returns the InventoryReportDestinationPrefix field if non-nil, zero value otherwise.

### GetInventoryReportDestinationPrefixOk

`func (o *AwsS3ProtectionGroupParams) GetInventoryReportDestinationPrefixOk() (*string, bool)`

GetInventoryReportDestinationPrefixOk returns a tuple with the InventoryReportDestinationPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryReportDestinationPrefix

`func (o *AwsS3ProtectionGroupParams) SetInventoryReportDestinationPrefix(v string)`

SetInventoryReportDestinationPrefix sets InventoryReportDestinationPrefix field to given value.

### HasInventoryReportDestinationPrefix

`func (o *AwsS3ProtectionGroupParams) HasInventoryReportDestinationPrefix() bool`

HasInventoryReportDestinationPrefix returns a boolean if a field has been set.

### SetInventoryReportDestinationPrefixNil

`func (o *AwsS3ProtectionGroupParams) SetInventoryReportDestinationPrefixNil(b bool)`

 SetInventoryReportDestinationPrefixNil sets the value for InventoryReportDestinationPrefix to be an explicit nil

### UnsetInventoryReportDestinationPrefix
`func (o *AwsS3ProtectionGroupParams) UnsetInventoryReportDestinationPrefix()`

UnsetInventoryReportDestinationPrefix ensures that no value is present for InventoryReportDestinationPrefix, not even an explicit nil
### GetInventoryReportFrequency

`func (o *AwsS3ProtectionGroupParams) GetInventoryReportFrequency() string`

GetInventoryReportFrequency returns the InventoryReportFrequency field if non-nil, zero value otherwise.

### GetInventoryReportFrequencyOk

`func (o *AwsS3ProtectionGroupParams) GetInventoryReportFrequencyOk() (*string, bool)`

GetInventoryReportFrequencyOk returns a tuple with the InventoryReportFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryReportFrequency

`func (o *AwsS3ProtectionGroupParams) SetInventoryReportFrequency(v string)`

SetInventoryReportFrequency sets InventoryReportFrequency field to given value.

### HasInventoryReportFrequency

`func (o *AwsS3ProtectionGroupParams) HasInventoryReportFrequency() bool`

HasInventoryReportFrequency returns a boolean if a field has been set.

### SetInventoryReportFrequencyNil

`func (o *AwsS3ProtectionGroupParams) SetInventoryReportFrequencyNil(b bool)`

 SetInventoryReportFrequencyNil sets the value for InventoryReportFrequency to be an explicit nil

### UnsetInventoryReportFrequency
`func (o *AwsS3ProtectionGroupParams) UnsetInventoryReportFrequency()`

UnsetInventoryReportFrequency ensures that no value is present for InventoryReportFrequency, not even an explicit nil
### GetObjects

`func (o *AwsS3ProtectionGroupParams) GetObjects() []AwsS3ProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *AwsS3ProtectionGroupParams) GetObjectsOk() (*[]AwsS3ProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *AwsS3ProtectionGroupParams) SetObjects(v []AwsS3ProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *AwsS3ProtectionGroupParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetSkipOnError

`func (o *AwsS3ProtectionGroupParams) GetSkipOnError() bool`

GetSkipOnError returns the SkipOnError field if non-nil, zero value otherwise.

### GetSkipOnErrorOk

`func (o *AwsS3ProtectionGroupParams) GetSkipOnErrorOk() (*bool, bool)`

GetSkipOnErrorOk returns a tuple with the SkipOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipOnError

`func (o *AwsS3ProtectionGroupParams) SetSkipOnError(v bool)`

SetSkipOnError sets SkipOnError field to given value.

### HasSkipOnError

`func (o *AwsS3ProtectionGroupParams) HasSkipOnError() bool`

HasSkipOnError returns a boolean if a field has been set.

### SetSkipOnErrorNil

`func (o *AwsS3ProtectionGroupParams) SetSkipOnErrorNil(b bool)`

 SetSkipOnErrorNil sets the value for SkipOnError to be an explicit nil

### UnsetSkipOnError
`func (o *AwsS3ProtectionGroupParams) UnsetSkipOnError()`

UnsetSkipOnError ensures that no value is present for SkipOnError, not even an explicit nil
### GetSourceId

`func (o *AwsS3ProtectionGroupParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *AwsS3ProtectionGroupParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *AwsS3ProtectionGroupParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *AwsS3ProtectionGroupParams) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *AwsS3ProtectionGroupParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *AwsS3ProtectionGroupParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *AwsS3ProtectionGroupParams) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *AwsS3ProtectionGroupParams) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *AwsS3ProtectionGroupParams) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *AwsS3ProtectionGroupParams) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *AwsS3ProtectionGroupParams) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *AwsS3ProtectionGroupParams) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetStorageClass

`func (o *AwsS3ProtectionGroupParams) GetStorageClass() []string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *AwsS3ProtectionGroupParams) GetStorageClassOk() (*[]string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *AwsS3ProtectionGroupParams) SetStorageClass(v []string)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *AwsS3ProtectionGroupParams) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


