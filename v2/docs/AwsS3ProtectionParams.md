# AwsS3ProtectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupObjectLevelACLs** | Pointer to **NullableBool** | Specifies whether to backup object level acls. Default value is false. | [optional] 
**BaselineIncrementalFrequency** | Pointer to **NullableString** | Specifies the baseline incremental frequency. | [optional] 
**InventoryReportDestination** | Pointer to **NullableString** | ARN of the inventory report destination bucket for S3 backups. | [optional] 
**InventoryReportDestinationPrefix** | Pointer to **NullableString** | The prefix in the S3 destination bucket where inventory reports will be stored. | [optional] 
**InventoryReportFrequency** | Pointer to **NullableString** | Specifies the frequency to generate inventory reports. | [optional] 
**Objects** | Pointer to [**[]AwsS3ObjectLevelParams**](AwsS3ObjectLevelParams.md) | Specifies the objects to be protected. | [optional] 
**SkipOnError** | Pointer to **NullableBool** | Specifies whether to skip files on error or not. Default value is false. | [optional] 
**StorageClass** | Pointer to **[]string** | Specifies the AWS S3 Storage classes to backup. | [optional] 

## Methods

### NewAwsS3ProtectionParams

`func NewAwsS3ProtectionParams() *AwsS3ProtectionParams`

NewAwsS3ProtectionParams instantiates a new AwsS3ProtectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsS3ProtectionParamsWithDefaults

`func NewAwsS3ProtectionParamsWithDefaults() *AwsS3ProtectionParams`

NewAwsS3ProtectionParamsWithDefaults instantiates a new AwsS3ProtectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupObjectLevelACLs

`func (o *AwsS3ProtectionParams) GetBackupObjectLevelACLs() bool`

GetBackupObjectLevelACLs returns the BackupObjectLevelACLs field if non-nil, zero value otherwise.

### GetBackupObjectLevelACLsOk

`func (o *AwsS3ProtectionParams) GetBackupObjectLevelACLsOk() (*bool, bool)`

GetBackupObjectLevelACLsOk returns a tuple with the BackupObjectLevelACLs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupObjectLevelACLs

`func (o *AwsS3ProtectionParams) SetBackupObjectLevelACLs(v bool)`

SetBackupObjectLevelACLs sets BackupObjectLevelACLs field to given value.

### HasBackupObjectLevelACLs

`func (o *AwsS3ProtectionParams) HasBackupObjectLevelACLs() bool`

HasBackupObjectLevelACLs returns a boolean if a field has been set.

### SetBackupObjectLevelACLsNil

`func (o *AwsS3ProtectionParams) SetBackupObjectLevelACLsNil(b bool)`

 SetBackupObjectLevelACLsNil sets the value for BackupObjectLevelACLs to be an explicit nil

### UnsetBackupObjectLevelACLs
`func (o *AwsS3ProtectionParams) UnsetBackupObjectLevelACLs()`

UnsetBackupObjectLevelACLs ensures that no value is present for BackupObjectLevelACLs, not even an explicit nil
### GetBaselineIncrementalFrequency

`func (o *AwsS3ProtectionParams) GetBaselineIncrementalFrequency() string`

GetBaselineIncrementalFrequency returns the BaselineIncrementalFrequency field if non-nil, zero value otherwise.

### GetBaselineIncrementalFrequencyOk

`func (o *AwsS3ProtectionParams) GetBaselineIncrementalFrequencyOk() (*string, bool)`

GetBaselineIncrementalFrequencyOk returns a tuple with the BaselineIncrementalFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineIncrementalFrequency

`func (o *AwsS3ProtectionParams) SetBaselineIncrementalFrequency(v string)`

SetBaselineIncrementalFrequency sets BaselineIncrementalFrequency field to given value.

### HasBaselineIncrementalFrequency

`func (o *AwsS3ProtectionParams) HasBaselineIncrementalFrequency() bool`

HasBaselineIncrementalFrequency returns a boolean if a field has been set.

### SetBaselineIncrementalFrequencyNil

`func (o *AwsS3ProtectionParams) SetBaselineIncrementalFrequencyNil(b bool)`

 SetBaselineIncrementalFrequencyNil sets the value for BaselineIncrementalFrequency to be an explicit nil

### UnsetBaselineIncrementalFrequency
`func (o *AwsS3ProtectionParams) UnsetBaselineIncrementalFrequency()`

UnsetBaselineIncrementalFrequency ensures that no value is present for BaselineIncrementalFrequency, not even an explicit nil
### GetInventoryReportDestination

`func (o *AwsS3ProtectionParams) GetInventoryReportDestination() string`

GetInventoryReportDestination returns the InventoryReportDestination field if non-nil, zero value otherwise.

### GetInventoryReportDestinationOk

`func (o *AwsS3ProtectionParams) GetInventoryReportDestinationOk() (*string, bool)`

GetInventoryReportDestinationOk returns a tuple with the InventoryReportDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryReportDestination

`func (o *AwsS3ProtectionParams) SetInventoryReportDestination(v string)`

SetInventoryReportDestination sets InventoryReportDestination field to given value.

### HasInventoryReportDestination

`func (o *AwsS3ProtectionParams) HasInventoryReportDestination() bool`

HasInventoryReportDestination returns a boolean if a field has been set.

### SetInventoryReportDestinationNil

`func (o *AwsS3ProtectionParams) SetInventoryReportDestinationNil(b bool)`

 SetInventoryReportDestinationNil sets the value for InventoryReportDestination to be an explicit nil

### UnsetInventoryReportDestination
`func (o *AwsS3ProtectionParams) UnsetInventoryReportDestination()`

UnsetInventoryReportDestination ensures that no value is present for InventoryReportDestination, not even an explicit nil
### GetInventoryReportDestinationPrefix

`func (o *AwsS3ProtectionParams) GetInventoryReportDestinationPrefix() string`

GetInventoryReportDestinationPrefix returns the InventoryReportDestinationPrefix field if non-nil, zero value otherwise.

### GetInventoryReportDestinationPrefixOk

`func (o *AwsS3ProtectionParams) GetInventoryReportDestinationPrefixOk() (*string, bool)`

GetInventoryReportDestinationPrefixOk returns a tuple with the InventoryReportDestinationPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryReportDestinationPrefix

`func (o *AwsS3ProtectionParams) SetInventoryReportDestinationPrefix(v string)`

SetInventoryReportDestinationPrefix sets InventoryReportDestinationPrefix field to given value.

### HasInventoryReportDestinationPrefix

`func (o *AwsS3ProtectionParams) HasInventoryReportDestinationPrefix() bool`

HasInventoryReportDestinationPrefix returns a boolean if a field has been set.

### SetInventoryReportDestinationPrefixNil

`func (o *AwsS3ProtectionParams) SetInventoryReportDestinationPrefixNil(b bool)`

 SetInventoryReportDestinationPrefixNil sets the value for InventoryReportDestinationPrefix to be an explicit nil

### UnsetInventoryReportDestinationPrefix
`func (o *AwsS3ProtectionParams) UnsetInventoryReportDestinationPrefix()`

UnsetInventoryReportDestinationPrefix ensures that no value is present for InventoryReportDestinationPrefix, not even an explicit nil
### GetInventoryReportFrequency

`func (o *AwsS3ProtectionParams) GetInventoryReportFrequency() string`

GetInventoryReportFrequency returns the InventoryReportFrequency field if non-nil, zero value otherwise.

### GetInventoryReportFrequencyOk

`func (o *AwsS3ProtectionParams) GetInventoryReportFrequencyOk() (*string, bool)`

GetInventoryReportFrequencyOk returns a tuple with the InventoryReportFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryReportFrequency

`func (o *AwsS3ProtectionParams) SetInventoryReportFrequency(v string)`

SetInventoryReportFrequency sets InventoryReportFrequency field to given value.

### HasInventoryReportFrequency

`func (o *AwsS3ProtectionParams) HasInventoryReportFrequency() bool`

HasInventoryReportFrequency returns a boolean if a field has been set.

### SetInventoryReportFrequencyNil

`func (o *AwsS3ProtectionParams) SetInventoryReportFrequencyNil(b bool)`

 SetInventoryReportFrequencyNil sets the value for InventoryReportFrequency to be an explicit nil

### UnsetInventoryReportFrequency
`func (o *AwsS3ProtectionParams) UnsetInventoryReportFrequency()`

UnsetInventoryReportFrequency ensures that no value is present for InventoryReportFrequency, not even an explicit nil
### GetObjects

`func (o *AwsS3ProtectionParams) GetObjects() []AwsS3ObjectLevelParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *AwsS3ProtectionParams) GetObjectsOk() (*[]AwsS3ObjectLevelParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *AwsS3ProtectionParams) SetObjects(v []AwsS3ObjectLevelParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *AwsS3ProtectionParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetSkipOnError

`func (o *AwsS3ProtectionParams) GetSkipOnError() bool`

GetSkipOnError returns the SkipOnError field if non-nil, zero value otherwise.

### GetSkipOnErrorOk

`func (o *AwsS3ProtectionParams) GetSkipOnErrorOk() (*bool, bool)`

GetSkipOnErrorOk returns a tuple with the SkipOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipOnError

`func (o *AwsS3ProtectionParams) SetSkipOnError(v bool)`

SetSkipOnError sets SkipOnError field to given value.

### HasSkipOnError

`func (o *AwsS3ProtectionParams) HasSkipOnError() bool`

HasSkipOnError returns a boolean if a field has been set.

### SetSkipOnErrorNil

`func (o *AwsS3ProtectionParams) SetSkipOnErrorNil(b bool)`

 SetSkipOnErrorNil sets the value for SkipOnError to be an explicit nil

### UnsetSkipOnError
`func (o *AwsS3ProtectionParams) UnsetSkipOnError()`

UnsetSkipOnError ensures that no value is present for SkipOnError, not even an explicit nil
### GetStorageClass

`func (o *AwsS3ProtectionParams) GetStorageClass() []string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *AwsS3ProtectionParams) GetStorageClassOk() (*[]string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *AwsS3ProtectionParams) SetStorageClass(v []string)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *AwsS3ProtectionParams) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


