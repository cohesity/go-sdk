# DataTransferToVaultSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataTransferPerProtectionJob** | Pointer to [**[]DataTransferToVaultPerProtectionJob**](DataTransferToVaultPerProtectionJob.md) | Array of Data Transfer Statistics Per Protection Jobs.  Specifies the data transfer summary statistics for each Protection Job that is transferring data from this Cohesity Cluster to this Vault (External Target). | [optional] 
**LogicalDataTransferredBytesDuringTimeRange** | Pointer to **[]int64** | Array of Logical Data Transferred Per Day.  Specifies the logical data transferred from this Cohesity Cluster to this Vault during the time period specified using the startTimeMsecs and endTimeMsecs parameters. For each day in the time period, an array element is returned, for example if 7 days are specified, 7 array elements are returned. The logical size is when the data is fully hydrated or expanded. | [optional] 
**NumLogicalBytesTransferred** | Pointer to **NullableInt64** | Specifies the total number of logical bytes that are transferred from this Cohesity Cluster to this Vault. The logical size is when the data is fully hydrated or expanded. | [optional] 
**NumPhysicalBytesTransferred** | Pointer to **NullableInt64** | Specifies the total number of physical bytes that are transferred from this Cohesity Cluster to this Vault. | [optional] 
**NumProtectionJobs** | Pointer to **NullableInt64** | Specifies the number of Protection Jobs that transfer data to this Vault. | [optional] 
**PhysicalDataTransferredBytesDuringTimeRange** | Pointer to **[]int64** | Array of Physical Data Transferred Per Day.  Specifies the physical data transferred from this Cohesity Cluster to this Vault during the time period specified using the startTimeMsecs and endTimeMsecs parameters. For each day in the time period, an array element is returned, for example if 7 days are specified, 7 array elements are returned. | [optional] 
**StorageConsumedBytes** | Pointer to **NullableInt64** | Specifies the storage consumed on the Vault as of last day in the specified time range. | [optional] 
**VaultId** | Pointer to **NullableInt64** | The vault Id associated with the vault. | [optional] 
**VaultName** | Pointer to **NullableString** | Specifies the name of the Vault (External Target). | [optional] 
**VaultType** | Pointer to **NullableString** | Specifies the type of Vault. &#39;kNearline&#39; indicates a Google Nearline Vault. &#39;kGlacier&#39; indicates an AWS Glacier Vault. &#39;kS3&#39; indicates an AWS S3 Vault. &#39;kAzureStandard&#39; indicates a Microsoft Azure Standard Vault. &#39;kS3Compatible&#39; indicates an S3 Compatible Vault. (See the online help for supported types.) &#39;kQStarTape&#39; indicates a QStar Tape Vault. &#39;kGoogleStandard&#39; indicates a Google Standard Vault. &#39;kGoogleDRA&#39; indicates a Google DRA Vault. &#39;kAmazonS3StandardIA&#39; indicates an Amazon S3 Standard-IA Vault. &#39;kAWSGovCloud&#39; indicates an AWS Gov Cloud Vault. &#39;kNAS&#39; indicates a NAS Vault. &#39;kColdline&#39; indicates a Google Coldline Vault. &#39;kAzureGovCloud&#39; indicates a Microsoft Azure Gov Cloud Vault. &#39;kAzureArchive&#39; indicates an Azure Archive Vault. &#39;kAzure&#39; indicates an Azure Vault. &#39;kGoogle&#39; indicates a Google Vault. &#39;kAmazon&#39; indicates an Amazon Vault. &#39;kOracle&#39; indicates an Oracle Vault. &#39;kOracleTierStandard&#39; indicates an Oracle Tier Standard Vault. &#39;kOracleTierArchive&#39; indicates an Oracle Tier Archive Vault. &#39;kAmazonC2S&#39; indicates an Amazon Commercial Cloud Services Vault. | [optional] 

## Methods

### NewDataTransferToVaultSummary

`func NewDataTransferToVaultSummary() *DataTransferToVaultSummary`

NewDataTransferToVaultSummary instantiates a new DataTransferToVaultSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTransferToVaultSummaryWithDefaults

`func NewDataTransferToVaultSummaryWithDefaults() *DataTransferToVaultSummary`

NewDataTransferToVaultSummaryWithDefaults instantiates a new DataTransferToVaultSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataTransferPerProtectionJob

`func (o *DataTransferToVaultSummary) GetDataTransferPerProtectionJob() []DataTransferToVaultPerProtectionJob`

GetDataTransferPerProtectionJob returns the DataTransferPerProtectionJob field if non-nil, zero value otherwise.

### GetDataTransferPerProtectionJobOk

`func (o *DataTransferToVaultSummary) GetDataTransferPerProtectionJobOk() (*[]DataTransferToVaultPerProtectionJob, bool)`

GetDataTransferPerProtectionJobOk returns a tuple with the DataTransferPerProtectionJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTransferPerProtectionJob

`func (o *DataTransferToVaultSummary) SetDataTransferPerProtectionJob(v []DataTransferToVaultPerProtectionJob)`

SetDataTransferPerProtectionJob sets DataTransferPerProtectionJob field to given value.

### HasDataTransferPerProtectionJob

`func (o *DataTransferToVaultSummary) HasDataTransferPerProtectionJob() bool`

HasDataTransferPerProtectionJob returns a boolean if a field has been set.

### SetDataTransferPerProtectionJobNil

`func (o *DataTransferToVaultSummary) SetDataTransferPerProtectionJobNil(b bool)`

 SetDataTransferPerProtectionJobNil sets the value for DataTransferPerProtectionJob to be an explicit nil

### UnsetDataTransferPerProtectionJob
`func (o *DataTransferToVaultSummary) UnsetDataTransferPerProtectionJob()`

UnsetDataTransferPerProtectionJob ensures that no value is present for DataTransferPerProtectionJob, not even an explicit nil
### GetLogicalDataTransferredBytesDuringTimeRange

`func (o *DataTransferToVaultSummary) GetLogicalDataTransferredBytesDuringTimeRange() []int64`

GetLogicalDataTransferredBytesDuringTimeRange returns the LogicalDataTransferredBytesDuringTimeRange field if non-nil, zero value otherwise.

### GetLogicalDataTransferredBytesDuringTimeRangeOk

`func (o *DataTransferToVaultSummary) GetLogicalDataTransferredBytesDuringTimeRangeOk() (*[]int64, bool)`

GetLogicalDataTransferredBytesDuringTimeRangeOk returns a tuple with the LogicalDataTransferredBytesDuringTimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalDataTransferredBytesDuringTimeRange

`func (o *DataTransferToVaultSummary) SetLogicalDataTransferredBytesDuringTimeRange(v []int64)`

SetLogicalDataTransferredBytesDuringTimeRange sets LogicalDataTransferredBytesDuringTimeRange field to given value.

### HasLogicalDataTransferredBytesDuringTimeRange

`func (o *DataTransferToVaultSummary) HasLogicalDataTransferredBytesDuringTimeRange() bool`

HasLogicalDataTransferredBytesDuringTimeRange returns a boolean if a field has been set.

### SetLogicalDataTransferredBytesDuringTimeRangeNil

`func (o *DataTransferToVaultSummary) SetLogicalDataTransferredBytesDuringTimeRangeNil(b bool)`

 SetLogicalDataTransferredBytesDuringTimeRangeNil sets the value for LogicalDataTransferredBytesDuringTimeRange to be an explicit nil

### UnsetLogicalDataTransferredBytesDuringTimeRange
`func (o *DataTransferToVaultSummary) UnsetLogicalDataTransferredBytesDuringTimeRange()`

UnsetLogicalDataTransferredBytesDuringTimeRange ensures that no value is present for LogicalDataTransferredBytesDuringTimeRange, not even an explicit nil
### GetNumLogicalBytesTransferred

`func (o *DataTransferToVaultSummary) GetNumLogicalBytesTransferred() int64`

GetNumLogicalBytesTransferred returns the NumLogicalBytesTransferred field if non-nil, zero value otherwise.

### GetNumLogicalBytesTransferredOk

`func (o *DataTransferToVaultSummary) GetNumLogicalBytesTransferredOk() (*int64, bool)`

GetNumLogicalBytesTransferredOk returns a tuple with the NumLogicalBytesTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLogicalBytesTransferred

`func (o *DataTransferToVaultSummary) SetNumLogicalBytesTransferred(v int64)`

SetNumLogicalBytesTransferred sets NumLogicalBytesTransferred field to given value.

### HasNumLogicalBytesTransferred

`func (o *DataTransferToVaultSummary) HasNumLogicalBytesTransferred() bool`

HasNumLogicalBytesTransferred returns a boolean if a field has been set.

### SetNumLogicalBytesTransferredNil

`func (o *DataTransferToVaultSummary) SetNumLogicalBytesTransferredNil(b bool)`

 SetNumLogicalBytesTransferredNil sets the value for NumLogicalBytesTransferred to be an explicit nil

### UnsetNumLogicalBytesTransferred
`func (o *DataTransferToVaultSummary) UnsetNumLogicalBytesTransferred()`

UnsetNumLogicalBytesTransferred ensures that no value is present for NumLogicalBytesTransferred, not even an explicit nil
### GetNumPhysicalBytesTransferred

`func (o *DataTransferToVaultSummary) GetNumPhysicalBytesTransferred() int64`

GetNumPhysicalBytesTransferred returns the NumPhysicalBytesTransferred field if non-nil, zero value otherwise.

### GetNumPhysicalBytesTransferredOk

`func (o *DataTransferToVaultSummary) GetNumPhysicalBytesTransferredOk() (*int64, bool)`

GetNumPhysicalBytesTransferredOk returns a tuple with the NumPhysicalBytesTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPhysicalBytesTransferred

`func (o *DataTransferToVaultSummary) SetNumPhysicalBytesTransferred(v int64)`

SetNumPhysicalBytesTransferred sets NumPhysicalBytesTransferred field to given value.

### HasNumPhysicalBytesTransferred

`func (o *DataTransferToVaultSummary) HasNumPhysicalBytesTransferred() bool`

HasNumPhysicalBytesTransferred returns a boolean if a field has been set.

### SetNumPhysicalBytesTransferredNil

`func (o *DataTransferToVaultSummary) SetNumPhysicalBytesTransferredNil(b bool)`

 SetNumPhysicalBytesTransferredNil sets the value for NumPhysicalBytesTransferred to be an explicit nil

### UnsetNumPhysicalBytesTransferred
`func (o *DataTransferToVaultSummary) UnsetNumPhysicalBytesTransferred()`

UnsetNumPhysicalBytesTransferred ensures that no value is present for NumPhysicalBytesTransferred, not even an explicit nil
### GetNumProtectionJobs

`func (o *DataTransferToVaultSummary) GetNumProtectionJobs() int64`

GetNumProtectionJobs returns the NumProtectionJobs field if non-nil, zero value otherwise.

### GetNumProtectionJobsOk

`func (o *DataTransferToVaultSummary) GetNumProtectionJobsOk() (*int64, bool)`

GetNumProtectionJobsOk returns a tuple with the NumProtectionJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumProtectionJobs

`func (o *DataTransferToVaultSummary) SetNumProtectionJobs(v int64)`

SetNumProtectionJobs sets NumProtectionJobs field to given value.

### HasNumProtectionJobs

`func (o *DataTransferToVaultSummary) HasNumProtectionJobs() bool`

HasNumProtectionJobs returns a boolean if a field has been set.

### SetNumProtectionJobsNil

`func (o *DataTransferToVaultSummary) SetNumProtectionJobsNil(b bool)`

 SetNumProtectionJobsNil sets the value for NumProtectionJobs to be an explicit nil

### UnsetNumProtectionJobs
`func (o *DataTransferToVaultSummary) UnsetNumProtectionJobs()`

UnsetNumProtectionJobs ensures that no value is present for NumProtectionJobs, not even an explicit nil
### GetPhysicalDataTransferredBytesDuringTimeRange

`func (o *DataTransferToVaultSummary) GetPhysicalDataTransferredBytesDuringTimeRange() []int64`

GetPhysicalDataTransferredBytesDuringTimeRange returns the PhysicalDataTransferredBytesDuringTimeRange field if non-nil, zero value otherwise.

### GetPhysicalDataTransferredBytesDuringTimeRangeOk

`func (o *DataTransferToVaultSummary) GetPhysicalDataTransferredBytesDuringTimeRangeOk() (*[]int64, bool)`

GetPhysicalDataTransferredBytesDuringTimeRangeOk returns a tuple with the PhysicalDataTransferredBytesDuringTimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDataTransferredBytesDuringTimeRange

`func (o *DataTransferToVaultSummary) SetPhysicalDataTransferredBytesDuringTimeRange(v []int64)`

SetPhysicalDataTransferredBytesDuringTimeRange sets PhysicalDataTransferredBytesDuringTimeRange field to given value.

### HasPhysicalDataTransferredBytesDuringTimeRange

`func (o *DataTransferToVaultSummary) HasPhysicalDataTransferredBytesDuringTimeRange() bool`

HasPhysicalDataTransferredBytesDuringTimeRange returns a boolean if a field has been set.

### SetPhysicalDataTransferredBytesDuringTimeRangeNil

`func (o *DataTransferToVaultSummary) SetPhysicalDataTransferredBytesDuringTimeRangeNil(b bool)`

 SetPhysicalDataTransferredBytesDuringTimeRangeNil sets the value for PhysicalDataTransferredBytesDuringTimeRange to be an explicit nil

### UnsetPhysicalDataTransferredBytesDuringTimeRange
`func (o *DataTransferToVaultSummary) UnsetPhysicalDataTransferredBytesDuringTimeRange()`

UnsetPhysicalDataTransferredBytesDuringTimeRange ensures that no value is present for PhysicalDataTransferredBytesDuringTimeRange, not even an explicit nil
### GetStorageConsumedBytes

`func (o *DataTransferToVaultSummary) GetStorageConsumedBytes() int64`

GetStorageConsumedBytes returns the StorageConsumedBytes field if non-nil, zero value otherwise.

### GetStorageConsumedBytesOk

`func (o *DataTransferToVaultSummary) GetStorageConsumedBytesOk() (*int64, bool)`

GetStorageConsumedBytesOk returns a tuple with the StorageConsumedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConsumedBytes

`func (o *DataTransferToVaultSummary) SetStorageConsumedBytes(v int64)`

SetStorageConsumedBytes sets StorageConsumedBytes field to given value.

### HasStorageConsumedBytes

`func (o *DataTransferToVaultSummary) HasStorageConsumedBytes() bool`

HasStorageConsumedBytes returns a boolean if a field has been set.

### SetStorageConsumedBytesNil

`func (o *DataTransferToVaultSummary) SetStorageConsumedBytesNil(b bool)`

 SetStorageConsumedBytesNil sets the value for StorageConsumedBytes to be an explicit nil

### UnsetStorageConsumedBytes
`func (o *DataTransferToVaultSummary) UnsetStorageConsumedBytes()`

UnsetStorageConsumedBytes ensures that no value is present for StorageConsumedBytes, not even an explicit nil
### GetVaultId

`func (o *DataTransferToVaultSummary) GetVaultId() int64`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *DataTransferToVaultSummary) GetVaultIdOk() (*int64, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *DataTransferToVaultSummary) SetVaultId(v int64)`

SetVaultId sets VaultId field to given value.

### HasVaultId

`func (o *DataTransferToVaultSummary) HasVaultId() bool`

HasVaultId returns a boolean if a field has been set.

### SetVaultIdNil

`func (o *DataTransferToVaultSummary) SetVaultIdNil(b bool)`

 SetVaultIdNil sets the value for VaultId to be an explicit nil

### UnsetVaultId
`func (o *DataTransferToVaultSummary) UnsetVaultId()`

UnsetVaultId ensures that no value is present for VaultId, not even an explicit nil
### GetVaultName

`func (o *DataTransferToVaultSummary) GetVaultName() string`

GetVaultName returns the VaultName field if non-nil, zero value otherwise.

### GetVaultNameOk

`func (o *DataTransferToVaultSummary) GetVaultNameOk() (*string, bool)`

GetVaultNameOk returns a tuple with the VaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultName

`func (o *DataTransferToVaultSummary) SetVaultName(v string)`

SetVaultName sets VaultName field to given value.

### HasVaultName

`func (o *DataTransferToVaultSummary) HasVaultName() bool`

HasVaultName returns a boolean if a field has been set.

### SetVaultNameNil

`func (o *DataTransferToVaultSummary) SetVaultNameNil(b bool)`

 SetVaultNameNil sets the value for VaultName to be an explicit nil

### UnsetVaultName
`func (o *DataTransferToVaultSummary) UnsetVaultName()`

UnsetVaultName ensures that no value is present for VaultName, not even an explicit nil
### GetVaultType

`func (o *DataTransferToVaultSummary) GetVaultType() string`

GetVaultType returns the VaultType field if non-nil, zero value otherwise.

### GetVaultTypeOk

`func (o *DataTransferToVaultSummary) GetVaultTypeOk() (*string, bool)`

GetVaultTypeOk returns a tuple with the VaultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultType

`func (o *DataTransferToVaultSummary) SetVaultType(v string)`

SetVaultType sets VaultType field to given value.

### HasVaultType

`func (o *DataTransferToVaultSummary) HasVaultType() bool`

HasVaultType returns a boolean if a field has been set.

### SetVaultTypeNil

`func (o *DataTransferToVaultSummary) SetVaultTypeNil(b bool)`

 SetVaultTypeNil sets the value for VaultType to be an explicit nil

### UnsetVaultType
`func (o *DataTransferToVaultSummary) UnsetVaultType()`

UnsetVaultType ensures that no value is present for VaultType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


