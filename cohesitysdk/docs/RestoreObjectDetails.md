# RestoreObjectDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchivalTarget** | Pointer to [**NullableArchivalExternalTarget**](ArchivalExternalTarget.md) | Specifies settings about the Archival Target (such as Tape or AWS). This field must be set if the object is being recovered or cloned from an archive or if files or folders are being restored from an archive. | [optional] 
**CloudDeployTarget** | Pointer to [**NullableCloudDeployTargetDetails**](CloudDeployTargetDetails.md) | Specifies settings about the Cloud Deploy target. This field must be set if the restore type is kDeployVMs and the object is to be deployed to cloud using a previously converted image. | [optional] 
**Environment** | Pointer to **NullableString** | Specifies the type of the Protection Source. Supported environment types such as &#39;kView&#39;, &#39;kSQL&#39;, &#39;kVMware&#39;, etc. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. &#39;kVMware&#39; indicates the VMware Protection Source environment. &#39;kHyperV&#39; indicates the HyperV Protection Source environment. &#39;kSQL&#39; indicates the SQL Protection Source environment. &#39;kView&#39; indicates the View Protection Source environment. &#39;kPuppeteer&#39; indicates the Cohesity&#39;s Remote Adapter. &#39;kPhysical&#39; indicates the physical Protection Source environment. &#39;kPure&#39; indicates the Pure Storage Protection Source environment. &#39;Nimble&#39; indicates the Nimble Storage Protection Source environment. &#39;kAzure&#39; indicates the Microsoft&#39;s Azure Protection Source environment. &#39;kNetapp&#39; indicates the Netapp Protection Source environment. &#39;kAgent&#39; indicates the Agent Protection Source environment. &#39;kGenericNas&#39; indicates the Generic Network Attached Storage Protection Source environment. &#39;kAcropolis&#39; indicates the Acropolis Protection Source environment. &#39;kPhsicalFiles&#39; indicates the Physical Files Protection Source environment. &#39;kIsilon&#39; indicates the Dell EMC&#39;s Isilon Protection Source environment. &#39;kGPFS&#39; indicates IBM&#39;s GPFS Protection Source environment. &#39;kKVM&#39; indicates the KVM Protection Source environment. &#39;kAWS&#39; indicates the AWS Protection Source environment. &#39;kExchange&#39; indicates the Exchange Protection Source environment. &#39;kHyperVVSS&#39; indicates the HyperV VSS Protection Source environment. &#39;kOracle&#39; indicates the Oracle Protection Source environment. &#39;kGCP&#39; indicates the Google Cloud Platform Protection Source environment. &#39;kFlashBlade&#39; indicates the Flash Blade Protection Source environment. &#39;kAWSNative&#39; indicates the AWS Native Protection Source environment. &#39;kO365&#39; indicates the Office 365 Protection Source environment. &#39;kO365Outlook&#39; indicates Office 365 outlook Protection Source environment. &#39;kHyperFlex&#39; indicates the Hyper Flex Protection Source environment. &#39;kGCPNative&#39; indicates the GCP Native Protection Source environment. &#39;kAzureNative&#39; indicates the Azure Native Protection Source environment. &#39;kKubernetes&#39; indicates a Kubernetes Protection Source environment. &#39;kElastifile&#39; indicates Elastifile Protection Source environment. &#39;kAD&#39; indicates Active Directory Protection Source environment. &#39;kRDSSnapshotManager&#39; indicates AWS RDS Protection Source environment. &#39;kCassandra&#39; indicates Cassandra Protection Source environment. &#39;kMongoDB&#39; indicates MongoDB Protection Source environment. &#39;kCouchbase&#39; indicates Couchbase Protection Source environment. &#39;kHdfs&#39; indicates Hdfs Protection Source environment. &#39;kHive&#39; indicates Hive Protection Source environment. &#39;kHBase&#39; indicates HBase Protection Source environment. | [optional] 
**JobId** | Pointer to **NullableInt64** | Protection Job Id.  Specifies id of the Protection Job that backed up the objects to be restored. | [optional] 
**JobRunId** | Pointer to **NullableInt64** | Specifies the id of the Job Run that captured the snapshot. | [optional] 
**JobUid** | Pointer to [**NullableUniversalId**](UniversalId.md) | Specifies the universal id of the Protection Job that backed up the objects to recover or clone or the objects that contain the files or folders to recover. | [optional] 
**ProtectionSourceId** | Pointer to **NullableInt64** | Specifies the id of the leaf object to recover, clone or recover files/folders from. | [optional] 
**SourceName** | Pointer to **NullableString** | Specifies the name of the Protection Source. | [optional] 
**StartedTimeUsecs** | Pointer to **NullableInt64** | Specifies the time when the Job Run starts capturing a snapshot. Specified as a Unix epoch Timestamp (in microseconds). | [optional] 

## Methods

### NewRestoreObjectDetails

`func NewRestoreObjectDetails() *RestoreObjectDetails`

NewRestoreObjectDetails instantiates a new RestoreObjectDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreObjectDetailsWithDefaults

`func NewRestoreObjectDetailsWithDefaults() *RestoreObjectDetails`

NewRestoreObjectDetailsWithDefaults instantiates a new RestoreObjectDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivalTarget

`func (o *RestoreObjectDetails) GetArchivalTarget() ArchivalExternalTarget`

GetArchivalTarget returns the ArchivalTarget field if non-nil, zero value otherwise.

### GetArchivalTargetOk

`func (o *RestoreObjectDetails) GetArchivalTargetOk() (*ArchivalExternalTarget, bool)`

GetArchivalTargetOk returns a tuple with the ArchivalTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalTarget

`func (o *RestoreObjectDetails) SetArchivalTarget(v ArchivalExternalTarget)`

SetArchivalTarget sets ArchivalTarget field to given value.

### HasArchivalTarget

`func (o *RestoreObjectDetails) HasArchivalTarget() bool`

HasArchivalTarget returns a boolean if a field has been set.

### SetArchivalTargetNil

`func (o *RestoreObjectDetails) SetArchivalTargetNil(b bool)`

 SetArchivalTargetNil sets the value for ArchivalTarget to be an explicit nil

### UnsetArchivalTarget
`func (o *RestoreObjectDetails) UnsetArchivalTarget()`

UnsetArchivalTarget ensures that no value is present for ArchivalTarget, not even an explicit nil
### GetCloudDeployTarget

`func (o *RestoreObjectDetails) GetCloudDeployTarget() CloudDeployTargetDetails`

GetCloudDeployTarget returns the CloudDeployTarget field if non-nil, zero value otherwise.

### GetCloudDeployTargetOk

`func (o *RestoreObjectDetails) GetCloudDeployTargetOk() (*CloudDeployTargetDetails, bool)`

GetCloudDeployTargetOk returns a tuple with the CloudDeployTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudDeployTarget

`func (o *RestoreObjectDetails) SetCloudDeployTarget(v CloudDeployTargetDetails)`

SetCloudDeployTarget sets CloudDeployTarget field to given value.

### HasCloudDeployTarget

`func (o *RestoreObjectDetails) HasCloudDeployTarget() bool`

HasCloudDeployTarget returns a boolean if a field has been set.

### SetCloudDeployTargetNil

`func (o *RestoreObjectDetails) SetCloudDeployTargetNil(b bool)`

 SetCloudDeployTargetNil sets the value for CloudDeployTarget to be an explicit nil

### UnsetCloudDeployTarget
`func (o *RestoreObjectDetails) UnsetCloudDeployTarget()`

UnsetCloudDeployTarget ensures that no value is present for CloudDeployTarget, not even an explicit nil
### GetEnvironment

`func (o *RestoreObjectDetails) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *RestoreObjectDetails) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *RestoreObjectDetails) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *RestoreObjectDetails) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *RestoreObjectDetails) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *RestoreObjectDetails) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetJobId

`func (o *RestoreObjectDetails) GetJobId() int64`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *RestoreObjectDetails) GetJobIdOk() (*int64, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *RestoreObjectDetails) SetJobId(v int64)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *RestoreObjectDetails) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### SetJobIdNil

`func (o *RestoreObjectDetails) SetJobIdNil(b bool)`

 SetJobIdNil sets the value for JobId to be an explicit nil

### UnsetJobId
`func (o *RestoreObjectDetails) UnsetJobId()`

UnsetJobId ensures that no value is present for JobId, not even an explicit nil
### GetJobRunId

`func (o *RestoreObjectDetails) GetJobRunId() int64`

GetJobRunId returns the JobRunId field if non-nil, zero value otherwise.

### GetJobRunIdOk

`func (o *RestoreObjectDetails) GetJobRunIdOk() (*int64, bool)`

GetJobRunIdOk returns a tuple with the JobRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRunId

`func (o *RestoreObjectDetails) SetJobRunId(v int64)`

SetJobRunId sets JobRunId field to given value.

### HasJobRunId

`func (o *RestoreObjectDetails) HasJobRunId() bool`

HasJobRunId returns a boolean if a field has been set.

### SetJobRunIdNil

`func (o *RestoreObjectDetails) SetJobRunIdNil(b bool)`

 SetJobRunIdNil sets the value for JobRunId to be an explicit nil

### UnsetJobRunId
`func (o *RestoreObjectDetails) UnsetJobRunId()`

UnsetJobRunId ensures that no value is present for JobRunId, not even an explicit nil
### GetJobUid

`func (o *RestoreObjectDetails) GetJobUid() UniversalId`

GetJobUid returns the JobUid field if non-nil, zero value otherwise.

### GetJobUidOk

`func (o *RestoreObjectDetails) GetJobUidOk() (*UniversalId, bool)`

GetJobUidOk returns a tuple with the JobUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobUid

`func (o *RestoreObjectDetails) SetJobUid(v UniversalId)`

SetJobUid sets JobUid field to given value.

### HasJobUid

`func (o *RestoreObjectDetails) HasJobUid() bool`

HasJobUid returns a boolean if a field has been set.

### SetJobUidNil

`func (o *RestoreObjectDetails) SetJobUidNil(b bool)`

 SetJobUidNil sets the value for JobUid to be an explicit nil

### UnsetJobUid
`func (o *RestoreObjectDetails) UnsetJobUid()`

UnsetJobUid ensures that no value is present for JobUid, not even an explicit nil
### GetProtectionSourceId

`func (o *RestoreObjectDetails) GetProtectionSourceId() int64`

GetProtectionSourceId returns the ProtectionSourceId field if non-nil, zero value otherwise.

### GetProtectionSourceIdOk

`func (o *RestoreObjectDetails) GetProtectionSourceIdOk() (*int64, bool)`

GetProtectionSourceIdOk returns a tuple with the ProtectionSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionSourceId

`func (o *RestoreObjectDetails) SetProtectionSourceId(v int64)`

SetProtectionSourceId sets ProtectionSourceId field to given value.

### HasProtectionSourceId

`func (o *RestoreObjectDetails) HasProtectionSourceId() bool`

HasProtectionSourceId returns a boolean if a field has been set.

### SetProtectionSourceIdNil

`func (o *RestoreObjectDetails) SetProtectionSourceIdNil(b bool)`

 SetProtectionSourceIdNil sets the value for ProtectionSourceId to be an explicit nil

### UnsetProtectionSourceId
`func (o *RestoreObjectDetails) UnsetProtectionSourceId()`

UnsetProtectionSourceId ensures that no value is present for ProtectionSourceId, not even an explicit nil
### GetSourceName

`func (o *RestoreObjectDetails) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *RestoreObjectDetails) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *RestoreObjectDetails) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *RestoreObjectDetails) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *RestoreObjectDetails) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *RestoreObjectDetails) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetStartedTimeUsecs

`func (o *RestoreObjectDetails) GetStartedTimeUsecs() int64`

GetStartedTimeUsecs returns the StartedTimeUsecs field if non-nil, zero value otherwise.

### GetStartedTimeUsecsOk

`func (o *RestoreObjectDetails) GetStartedTimeUsecsOk() (*int64, bool)`

GetStartedTimeUsecsOk returns a tuple with the StartedTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedTimeUsecs

`func (o *RestoreObjectDetails) SetStartedTimeUsecs(v int64)`

SetStartedTimeUsecs sets StartedTimeUsecs field to given value.

### HasStartedTimeUsecs

`func (o *RestoreObjectDetails) HasStartedTimeUsecs() bool`

HasStartedTimeUsecs returns a boolean if a field has been set.

### SetStartedTimeUsecsNil

`func (o *RestoreObjectDetails) SetStartedTimeUsecsNil(b bool)`

 SetStartedTimeUsecsNil sets the value for StartedTimeUsecs to be an explicit nil

### UnsetStartedTimeUsecs
`func (o *RestoreObjectDetails) UnsetStartedTimeUsecs()`

UnsetStartedTimeUsecs ensures that no value is present for StartedTimeUsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


